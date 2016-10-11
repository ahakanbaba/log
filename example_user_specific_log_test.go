package log_test

import (
	"github.com/ahakanbaba/log"
	stdlog "log"
	"os"
)

const logPrefix = ""

// If the user does not like the parameters in DefaultLog, she can pass their
// own log object to use.
func Example_userSpecificLog() {
	// The user wants simple logging to stdout. Does not care about timestamps
	// in the log messages
	logImpl := stdlog.New(os.Stdout, logPrefix, 0)
	// Pass the user specific logImpl to the New function to use its settings
	l := log.New(log.Warning, logImpl)

	// Error and warning logs get emitted because the severity passed to New
	// function is warning.
	l.Errorf("%s", "This is an error log")
	l.Warningf("%s", "This is a warning log")

	// Debug log gets supressed because it is less severe than what is passed
	// to the New function.
	l.Debugf("%s", "This is a debug log")

	// Output:
	// ERROR This is an error log
	// WARNING This is a warning log
}
