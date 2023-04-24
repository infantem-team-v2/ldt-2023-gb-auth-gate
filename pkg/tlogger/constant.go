package tlogger

const (
	labels string = "source=%s job=error_logging statusCode=%d internalError=%s"
)

type LogLevel uint8

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)
