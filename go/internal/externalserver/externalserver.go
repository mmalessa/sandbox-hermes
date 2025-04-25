package externalserver

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

type ExternalServer struct {
	execCmd *exec.Cmd
	pid     int
	process *process.Process
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
	e.execCmd.Stdout = stdLogger
}

func (e *ExternalServer) Stop() error {
	if e.execCmd.ProcessState != nil && e.execCmd.ProcessState.Exited() {
		e.pid = 0
		e.process = nil

		return nil
	}

	time.Sleep(1 * time.Second) // ??
	if err := e.execCmd.Process.Signal(syscall.SIGINT); err != nil {
		return err
	}

	time.Sleep(2 * time.Second) // waiting more

	if e.execCmd.Process != nil {
		if err := e.execCmd.Process.Kill(); err != nil {
			return err
		}
	}

	e.pid = 0
	e.process = nil
	return nil
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
	e.Stop()
	return nil
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
