package tlogger

import (
	"bank_api/config"
	"github.com/carlware/promtail-go"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
)

type TLogger struct {
	Config     *config.Config `di:"config"`
	lokiLogger promtail.HttpClient
	logger     *logrus.Logger
}

func NewLogger() ILogger {

	return &TLogger{}
}

func BuildLogger(ctn di.Container) (interface{}, error) {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	return &TLogger{
		Config: ctn.Get("config").(*config.Config),
		logger: logger,
	}, nil
}

func (T *TLogger) log(level logrus.Level, msg string, args ...interface{}) {
	T.logger.Logf(level, msg, args...)
}

func (T *TLogger) sendLog(msg string) {
	//TODO implement me
	panic("implement me")
}

func (T *TLogger) Infof(msgf string, args ...interface{}) {
	T.log(logrus.InfoLevel, msgf, args...)
}

func (T *TLogger) Debugf(msgf string, args ...interface{}) {
	T.log(logrus.DebugLevel, msgf, args...)
}

func (T *TLogger) Warnf(msgf string, args ...interface{}) {
	T.log(logrus.WarnLevel, msgf, args...)
}

func (T *TLogger) Errorf(msgf string, args ...interface{}) {
	T.log(logrus.ErrorLevel, msgf, args...)
}

func (T *TLogger) ErrorFull(err error) {
	T.Errorf("Error: %s", err.Error())
}

// Write implementing for putting it into fiber logger
func (T *TLogger) Write(p []byte) (n int, err error) {

	return 0, nil
}
