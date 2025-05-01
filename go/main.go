package main

import (
	"hermes/internal/externalserver"
	"hermes/internal/stdloger"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Hermes start")

	cmdLine := []string{"php", "bin/console", "hermes:test-worker"}
	envs := map[string]string{
		"TEST_ENV": "something from Go",
	}
	stdl := stdloger.New()

	es := externalserver.New(envs, cmdLine, stdl)
	es.Start()

	// Handling Ctrl+C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		logrus.Infof("Signal received: '%s'", sig)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := es.Wait(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Info("External server ended gracefully")
		wg.Done()
	}()

	// example only
	for i := 0; i < 3; i++ {
		text := "Hello from GO! " + strconv.Itoa(i)
		id := uuid.New().String()
		msg := map[string]string{"text": text, "id": id}
		packed, _ := msgpack.Marshal(msg)

		if err := es.Send(packed); err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("Request to PHP: %#v", msg)
		}

		if response, err := es.Receive(); err != nil {
			logrus.Errorf("Response from PHP Error: %s", err)
		} else {
			logrus.Infof("Response from PHP: %s", response)
		}
	}
	es.Stop()
	// end of example

	wg.Wait()
	logrus.Info("END")
}
