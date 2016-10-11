
log package is a simple lightweight wrapper adding severity specification
around the log package from the standard library. This package enables,
users to make calls like log.Debugf(...) log.Warningf(...), etc. It uses the
log package from the standard library as the underlying logger
implementation. It can be passed a user specified log.Logging object or one
can use a DefaultLog that has reasonable defaults.

For more information see: https://godoc.org/github.com/ahakanbaba/log
