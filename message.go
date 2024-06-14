package chlogger

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
