package externalserver

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

type ExternalServer struct {
	execCmd *exec.Cmd
	pid     int
	process *process.Process
	stdin   io.WriteCloser
	stdout  io.ReadCloser
	reader  *bufio.Reader
}

func New(env map[string]string, cmd []string, stdLogger io.Writer) *ExternalServer {
	c := &ExternalServer{}
	c.init(env, cmd, stdLogger)

	return c
}

func (e *ExternalServer) init(env map[string]string, cmd []string, stdLogger io.Writer) {
	var cmdArgs []string
	switch len(cmd) {
	case 1:
		cmdArgs = append(cmdArgs, strings.Split(cmd[0], " ")...)
	default:
		cmdArgs = cmd
	}
	if len(cmdArgs) == 1 {
		e.execCmd = exec.Command(cmd[0])
	} else {
		e.execCmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
	}

	if len(env) > 0 {
		for k, v := range env {
			e.execCmd.Env = append(e.execCmd.Env, fmt.Sprintf("%s=%s", strings.ToUpper(k), os.Expand(v, os.Getenv)))
		}
	}
	e.execCmd.Env = append(e.execCmd.Env, os.Environ()...)

	e.execCmd.Stderr = stdLogger
	e.stdin, _ = e.execCmd.StdinPipe()

	// e.execCmd.Stdout = e
	e.stdout, _ = e.execCmd.StdoutPipe()
	e.reader = bufio.NewReader(e.stdout)
}

func (e *ExternalServer) Start() error {
	err := e.execCmd.Start()
	if err != nil {
		return err
	}
	e.pid = e.execCmd.Process.Pid

	e.process, err = process.NewProcess(int32(e.pid))
	if err != nil {
		return err
	}

	return nil
}

func (e *ExternalServer) Wait() error {
	if err := e.execCmd.Wait(); err != nil {
		return err
	}
	logrus.Debug("Extenral serwer finished")
	e.destruct()
	return nil
}

func (e *ExternalServer) Stop() error {
	logrus.Debug("ExternalServer -> STOP")
	defer e.destruct()
	e.stdin.Close()
	e.stdout.Close()

	// First check
	if e.execCmd.ProcessState != nil && e.execCmd.ProcessState.Exited() {
		logrus.Debug("OK, was dead")
		return nil
	}

	// We send SIGINT
	if err := e.execCmd.Process.Signal(syscall.SIGINT); err != nil {
		return err
	}

	// Waiting
	timeout := 10 * time.Second
	interval := 1000 * time.Millisecond
	elapsed := time.Duration(0)

	for elapsed < timeout {
		if e.execCmd.ProcessState != nil && e.execCmd.ProcessState.Exited() {
			logrus.Debug("OK, is dead")
			// Process has exited gracefully
			return nil
		}
		logrus.Debug("Still alive")
		time.Sleep(interval)
		elapsed += interval
	}

	// If it is still alive...
	if e.execCmd.ProcessState == nil || !e.execCmd.ProcessState.Exited() {
		logrus.Debug("Still alive -> KILL")
		if err := e.execCmd.Process.Kill(); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (e *ExternalServer) destruct() {
	e.pid = 0
	e.process = nil
	e.stdin.Close()
	e.stdout.Close()
}

func (e *ExternalServer) Pid() (int, error) {
	if e.pid == 0 {
		return 0, errors.New("external server not started")
	}

	return e.pid, nil
}

func (e *ExternalServer) MemoryUsage() (uint64, error) {

	if e.process == nil || e.pid == 0 {
		return 0, errors.New("external server not started")
	}

	memInfo, err := e.process.MemoryInfo()
	if err != nil {
		return 0, err
	}

	return memInfo.RSS, nil
}

// PHP STDIN
func (e *ExternalServer) Send(req InternalRequest) error {
	if e.pid == 0 {
		return nil
	}

	packed, _ := msgpack.Marshal(req)
	binary.Write(e.stdin, binary.BigEndian, uint32(len(packed)))
	_, err := e.stdin.Write(packed)
	return err
}

// PHP STDOUT
// func (e *ExternalServer) Write(msg []byte) (n int, err error) {
// 	logrus.Infof("From ExternalServer <<<: %s", msg)
// 	return len(msg), nil
// }

// PHP STDOUT
func (e *ExternalServer) Receive() (InternalResponse, error) {
	var resp InternalResponse

	if e.pid == 0 {
		return resp, nil
	}

	// Get length
	lenBuf := make([]byte, 4)
	if _, err := e.stdout.Read(lenBuf); err != nil {
		return resp, err
	}
	length := binary.BigEndian.Uint32(lenBuf)

	// Get response
	respBuffer := make([]byte, length)
	if _, err := e.stdout.Read(respBuffer); err != nil {
		return resp, err
	}

	// var generic map[string]interface{}
	// if err := msgpack.Unmarshal(respBuffer, &generic); err == nil {
	// 	logrus.Debugf("Decoded map: %+v", generic)
	// }

	if err := msgpack.Unmarshal(respBuffer, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
