package chlogger

// Logger is an interface that defines the methods for logging messages
type Logger interface {
	Init() error

	// Print logs a string
	Print(string)

	// Println logs a string message followed by a new line
	Println(string)

	// Debug logs a debug string
	Debug(string)

	// Info logs an informational string
	Info(string)

	// Warn logs a warning string
	Warn(string)

	// Error logs an error string
	Error(string)

	// Fatal logs a fatal string message
	Fatal(string)

	// Printf logs a formatted string
	Printf(string, ...interface{})

	// Debugf logs a formatted debug string
	Debugf(string, ...interface{})

	// Infof logs a formatted informational string
	Infof(string, ...interface{})

	// Warnf logs a formatted warning string
	Warnf(string, ...interface{})

	// Errorf logs a formatted error string
	Errorf(string, ...interface{})

	// Fatalf logs a formatted fatal string message
	Fatalf(string, ...interface{})
}
