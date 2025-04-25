package stdloger

import "fmt"

type StdLogger struct {
}

func New() *StdLogger {
	return &StdLogger{}
}

func (l *StdLogger) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", p)
	return len(p), nil
}
