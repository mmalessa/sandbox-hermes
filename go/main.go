package main

import (
	"hermes/internal/externalserver"
	"hermes/internal/stdloger"
	"os"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Hermes start")

	cmdLine := []string{"php", "bin/console", "hermes:test-worker"}
	envs := map[string]string{}
	stdl := stdloger.New()

	es := externalserver.New(envs, cmdLine, stdl)

	es.Start()

	var wg sync.WaitGroup
	wg.Add(2)
	defer func() {
		if err := es.Stop(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Info("Hermes ended gracefully")
		wg.Done()
	}()

	go func() {
		if err := es.Wait(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Warning("External server ended gracefully")
		wg.Done()
	}()

	for i := 0; i < 3; i++ {

		message := "Hello from GO! " + strconv.Itoa(i)
		if err := es.Send([]byte(message)); err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("Request to PHP: %s", message)
		}

		if response, err := es.Receive(); err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("Response from PHP: %s", response)
		}
	}
	logrus.Info("END")
}
