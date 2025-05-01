package stdloger

import (
	"github.com/sirupsen/logrus"
)

type StdLogger struct {
}

func New() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) Write(p []byte) (n int, err error) {
	logrus.Debugf("StdLoger: %s", p)
	return len(p), nil
}
