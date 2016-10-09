package log

// logger is a simple wrapper around the log package in the standard library.
// It enables severities and very basic severity filtering.
// One can also have an easy way to enable logging some reasonable default
// settings.

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Severity uint8

const (
	Debug Severity = iota
	Info
	Warning
	Error
	Fatal
)

type Log struct {
	enabledSeverity Severity
	// enabledSeverity specifies the lowest enabled severity.
	// Any severity greater or equal to enabledSeverity are enabled.

	log *log.Logger
	// The log object from standard library that is used for actual log generation.
}

func DefaultLog(fileNamePrefix string) *log.Logger {
	// DefaultLog creates a logger object and initializes with reasonable default
	// parameters
	// The logfile is created  in the run directory.
	// File name looks like this: fileNamePrefix_timestamp_pid.log
	// If the file cannot be opened DefaultLog panics
	// The log lines also contain timestamps

	now := time.Now().Format("2006-01-02--15-04-05.99999999")
	// An interesting way of specifying the time. There is a well known date and
	// the way you spedicy that date will be followed by the Format function.
	pid := strconv.Itoa(os.Getpid())
	fileName := fileNamePrefix + "_" + now + "_" + pid + ".log"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicln("Failed to open log file", fileName, ":", err)
	}
	l := log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return l
}

func NewLogger(sev Severity, l *log.Logger) *Log {
	return &Log{sev, l}
}

func (l *Log) isDebugEnabled() bool {
	return l != nil && l.enabledSeverity <= Debug
}
func (l *Log) isInfoEnabled() bool {
	return l != nil && l.enabledSeverity <= Info
}
func (l *Log) isWarningEnabled() bool {
	return l != nil && l.enabledSeverity <= Warning
}
func (l *Log) isErrorEnabled() bool {
	return l != nil && l.enabledSeverity <= Error
}
func (l *Log) isFatalEnabled() bool {
	return l != nil && l.enabledSeverity <= Fatal
}

func (l *Log) Debugf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.isDebugEnabled() {
		a = append([]interface{}{"DEBUG"}, a...)
		l.log.Printf("%s "+f, a...)
	}
}
func (l *Log) Infof(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.isInfoEnabled() {
		a = append([]interface{}{"INFO"}, a...)
		l.log.Printf("%s "+f, a...)
	}
}
func (l *Log) Warningf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.isWarningEnabled() {
		a = append([]interface{}{"WARNING"}, a...)
		l.log.Printf("%s "+f, a...)
	}
}
func (l *Log) Errorf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.isErrorEnabled() {
		a = append([]interface{}{"ERROR"}, a...)
		l.log.Printf("%s "+f, a...)
	}
}
func (l *Log) Fatalf(f string, a ...interface{}) {
	if l == nil {
		return
	}
	if l.isFatalEnabled() {
		a = append([]interface{}{"FATAL"}, a...)
		l.log.Printf("%s "+f, a...)
	}
}

// FixMe: We should make the comments go doc friendly.
// FixMe: Add some examples.
// FixMe: Add some readme and documentation.
