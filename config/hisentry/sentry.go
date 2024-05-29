package hisentry

import (
	"log"

	env "api.kalbe.crm/config/viper"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

type hiSentry struct {
	env viper.Viper
}

func NewHisentry() *hiSentry {
	return &hiSentry{
		env: *env.NewViper(),
	}
}

func (s *hiSentry) Init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:        s.env.GetString("SENTRY_DSN"),
		ServerName: "histore-pub",
		// BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
		// 	return &sentry.Event{
		// 		Level: sentry.Level(logger.Level.String()),
		// 		Extra: map[string]interface{}{
		// 			"log": logger.Data,
		// 		},
		// 		Timestamp: logger.Time,
		// 	}
		// },
		Debug: false,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

}
