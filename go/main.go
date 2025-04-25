package main

import (
	"hermes/internal/externalserver"
	"hermes/internal/stdloger"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hermes start")

	cmdLine := []string{"php", "bin/console", "hermes:test-worker"}
	envs := map[string]string{}
	stdl := stdloger.New()

	es := externalserver.New(envs, cmdLine, stdl)

	es.Start()
	defer func() {
		if err := es.Stop(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	}()

	go func() {
		if err := es.Wait(); err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Warning("External server end")
	}()

	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(time.Second)
		}
		printMemoryUsage(es)
	}

	logrus.Info("Hermes end")
}

func printMemoryUsage(es *externalserver.ExternalServer) {
	if mu, err := es.MemoryUsage(); err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("Memory usage %d\n", mu)
	}
}
