package tlogger

type LoggerMessenger interface {
	LoggerMessage() string
}

type Logger interface {
	// sendLog Send log to remote logs storage
	sendLog(msg string)

	// Default logger methods
	Infof(msgf string, args ...interface{})
	Debugf(msgf string, args ...interface{})
	Warnf(msgf string, args ...interface{})
	Errorf(msgf string, args ...interface{})
	// ErrorFull Logging error w/ stacktrace
	ErrorFull(err error)
}
