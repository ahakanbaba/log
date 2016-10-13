package log_test

import (
	"github.com/ahakanbaba/log"
)

// The user is happy with the default settings of the log module.
// Sets up log inteface with the DefaultLog and
func ExampleDefaultLog() {

	// The DefaultLog results in a file with name similar to
	// SimpleMath_2016-10-10--18-09-42.44439598_32859.log
	// SimpleMapth_<timestamp>_<pid>.log
	l := log.New(log.Info, log.DefaultLog("SimpleMath"))

	a := 3
	b := 5
	kk := 31
	c := 3 + 5
	// Prints:
	// 2016/10/10 18:09:42.444874 INFO 3 + 5 is equal to 8
	l.Infof("%d + %d is equal to %d", a, b, c)

	// Does not emit any log because the lowest enabled severity passed to New
	// call is Info.  Debug is less severe than Info, there it is not printed
	// out to the log file
	l.Debugf("%s", "Addition is complete")

}

// FixMe: Add using isDebugEnabled. With expensive command line arguments to be
// evaluated to the log function.
// FixMe: Add different module's logs get accumulated in the same log file.
