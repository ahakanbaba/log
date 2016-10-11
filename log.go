// log package is a simple lightweight wrapper adding severity specification
// around the log package from the standard library. This package enables,
// users to make calls like log.Debugf(...) log.Warningf(...), etc. It uses the
// log package from the standard library as the underlying logger
// implementation. It can be passed a user specified log.Logging object or one
// can use a DefaultLog that has reasonable defaults. See the Examples for more
// information.
package log

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Severity type represent different severrity levels to log.
type Severity uint8

// The different severities are defined in their order of severeness.
const (
	Debug Severity = iota
	Info
	Warning
	Error
	Fatal
)

// Log type enables severity specification and filtering and wraps the log.Logger
// type from the standard library.
type Log struct {
	// enabledSeverity specifies the lowest enabled severity.
	// Any severity greater or equal to enabledSeverity are enabled.
	enabledSeverity Severity

	// Impl is the log object from standard library that is used for actual log
	// generation.
	Impl *log.Logger
}

// DefaultLog creates a log.Logger object and initializes it with reasonable
// default parameters The logfile is created in the run directory.  File name
// looks like this: fileNamePrefix_timestamp_pid.log If the file cannot be
// opened DefaultLog panics The log lines also contain timestamps
func DefaultLog(fileNamePrefix string) *log.Logger {

	// An interesting way of specifying the time. There is a well known date and
	// the way you spedicy that date will be followed by the Format function.
	now := time.Now().Format("2006-01-02--15-04-05.99999999")
	pid := strconv.Itoa(os.Getpid())
	fileName := fileNamePrefix + "_" + now + "_" + pid + ".log"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicln("Failed to open log file", fileName, ":", err)
	}
	l := log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return l
}

// New constructs a new Log object and returns a pointer to it.
// The passed sev is the lowest severity that is enabled. For example
// passing Info to the New function enables all Debug, Info, Warning and Fatal
// levels and disables the Warning level
// The user may pass her own log.Logger type if they have one set up already.
// If the user is happy with the default settings she can use
// myLog := New(desiredSeverity,log.DefaultLog())
func New(sev Severity, l *log.Logger) *Log {
	return &Log{sev, l}
}

// IsDebugEnabled returns true if Debug level logs are emitted, returns false
// if they are supressed.
func (l *Log) IsDebugEnabled() bool {
	return l != nil && l.enabledSeverity <= Debug
}

// IsInfoEnabled returns true if Info level logs are emitted, returns false if
// they are supressed.
func (l *Log) IsInfoEnabled() bool {
	return l != nil && l.enabledSeverity <= Info
}

// IsWarningEnabled returns true if Warning level logs are emitted, returns
// false if they are supressed.
func (l *Log) IsWarningEnabled() bool {
	return l != nil && l.enabledSeverity <= Warning
}

// IsErrorEnabled returns true if Error level logs are emitted, returns
// false if they are supressed.
func (l *Log) IsErrorEnabled() bool {
	return l != nil && l.enabledSeverity <= Error
}

// IsFatalEnabled returns true if Fatal level logs are emitted, returns
// false if they are supressed.
func (l *Log) IsFatalEnabled() bool {
	return l != nil && l.enabledSeverity <= Fatal
}

// Debugf generates a Debug level log with a "DEBUG" annotation. Arguments are
// handled in the manner of fmt.Printf.
func (l *Log) Debugf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.IsDebugEnabled() {
		a = append([]interface{}{"DEBUG"}, a...)
		l.Impl.Printf("%s "+f, a...)
	}
}

// Infof generates an Info level log with an "INFO" annotation. Arguments are
// handled in the manner of fmt.Printf.
func (l *Log) Infof(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.IsInfoEnabled() {
		a = append([]interface{}{"INFO"}, a...)
		l.Impl.Printf("%s "+f, a...)
	}
}

// Warningf generates a Warning level log with a "WARNING" annotation. Arguments are
// handled in the manner of fmt.Printf.
func (l *Log) Warningf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.IsWarningEnabled() {
		a = append([]interface{}{"WARNING"}, a...)
		l.Impl.Printf("%s "+f, a...)
	}
}

// ErrorF generates an Error level log with an "ERROR" annotation. Arguments are
// handled in the manner of fmt.Printf.
func (l *Log) Errorf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.IsErrorEnabled() {
		a = append([]interface{}{"ERROR"}, a...)
		l.Impl.Printf("%s "+f, a...)
	}
}

// Fatalf generates a Fatal  level log with a "FATAL" annotation. Arguments are
// handled in the manner of fmt.Printf.
func (l *Log) Fatalf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.IsFatalEnabled() {
		a = append([]interface{}{"FATAL"}, a...)
		l.Impl.Printf("%s "+f, a...)
	}
}

// FixMe: Add some examples.
