package chlogger

// Logger is the interface that all loggers must implement.
var _ Logger = &Noop{}

// Noop is a logger implementation that does nothing.
type Noop struct{}

// Init initializes the Noop logger.
func (n *Noop) Init() error {
	return nil
}

// Close closes the Noop logger.
func (n *Noop) Close() error {
	return nil
}

// Print does nothing.
func (n *Noop) Print(string) {}

// Println does nothing.
func (n *Noop) Println(string) {}

// Debug does nothing.
func (n *Noop) Debug(string) {}

// Info does nothing.
func (n *Noop) Info(string) {}

// Warn does nothing.
func (n *Noop) Warn(string) {}

// Error does nothing.
func (n *Noop) Error(string) {}

// Fatal does nothing.
func (n *Noop) Fatal(string) {}

// Printf does nothing.
func (n *Noop) Printf(string, ...interface{}) {}

// Debugf does nothing.
func (n *Noop) Debugf(string, ...interface{}) {}

// Infof does nothing.
func (n *Noop) Infof(string, ...interface{}) {}

// Warnf does nothing.
func (n *Noop) Warnf(string, ...interface{}) {}

// Errorf does nothing.
func (n *Noop) Errorf(string, ...interface{}) {}

// Fatalf does nothing.
func (n *Noop) Fatalf(string, ...interface{}) {}
