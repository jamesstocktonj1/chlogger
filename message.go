package chlogger

import "fmt"

type Level string

const (
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	Fatal Level = "FATAL"
)

type Message struct {
	Level   Level
	Message string
	Args    []interface{}
}

func (m *Message) String() string {
	msg := fmt.Sprintf(m.Message, m.Args...)
	return fmt.Sprintf("[%s]: %s", m.Level, msg)
}
