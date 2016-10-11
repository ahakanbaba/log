package log_test

import (
	"github.com/ahakanbaba/log"
	stdlog "log"
	"os"
)

// An long and expensive function that should not be called during live production
func expensiveFunctionGeneratesDebugDump() string {
	return "Hard to generate string"
}

func Example_expensiveLogParameter() {
	// The user wants simple logging to stdout. Does not care about timestamps
	// in the log messages
	logImpl := stdlog.New(os.Stdout, "", 0)
	// Pass the user specific logImpl to the New function to use its settings
	l := log.New(log.Debug, logImpl)

	// Sometimes the log function parameters may be expensive to calculate.
	// In order to avoid evaluation of the paramters for supressed messages, one
	// can check beforehand if the log message will be emitted or not.
	if l.IsDebugEnabled() {
		l.Debugf("The state of the process is: %s",
			expensiveFunctionGeneratesDebugDump())
	}

	// Output:
	// DEBUG The state of the process is: Hard to generate string
}
