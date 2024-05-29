package logger

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type sentryHook struct {
}

func NewSentryHook() *sentryHook {
	return &sentryHook{}
}

func (s *sentryHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *sentryHook) Fire(logger *logrus.Entry) error {

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetContext("log_data", logger.Data)
		scope.SetLevel(sentry.Level(logger.Level.String()))
	})

	sentry.CaptureMessage(logger.Message)
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	return nil
}
