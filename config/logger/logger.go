package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"api.kalbe.crm/config/constants"
	"api.kalbe.crm/config/viper"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
	env = viper.NewViper()
)

func LogInit(log *logrus.Logger) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	log.WithTime(now)

	log.AddHook(NewSentryHook())

	if !strings.EqualFold(env.GetString("APP_ENV"), "production") {
		log.SetLevel(logrus.TraceLevel)
	}
	log.SetFormatter(&logrus.JSONFormatter{})
}

func LogError(ctx context.Context, err error, statusCode int) {
	if err == nil {
		return
	}
	_, file, line, _ := runtime.Caller(1)

	go func(file string, line int) {
		log.SetLevel(logrus.ErrorLevel)
		LogInit(log)

		var (
			email      string
			membership string
		)

		emailRes := ctx.Value(constants.Key("email"))
		if emailRes != nil {
			email = emailRes.(string)
		}
		membershipRes := ctx.Value(constants.Key("membership"))
		if membershipRes != nil {
			membership = membershipRes.(string)
		}

		log.WithFields(logrus.Fields{
			"msg":           err.Error(),
			"file":          fmt.Sprintf("%s:%d", file, line),
			"response_code": statusCode,
			"user": map[string]string{
				"email":      email,
				"membership": membership,
			},
		}).Error(err)
	}(file, line)
}

func LogErrorStr(ctx context.Context, err string, statusCode int) {
	_, file, line, _ := runtime.Caller(1)

	go func() {
		log.SetLevel(logrus.ErrorLevel)
		LogInit(log)

		var (
			email      string
			membership string
		)

		emailRes := ctx.Value(constants.Key("email"))
		if emailRes != nil {
			email = emailRes.(string)
		}
		membershipRes := ctx.Value(constants.Key("membership"))
		if membershipRes != nil {
			membership = membershipRes.(string)
		}

		log.WithFields(logrus.Fields{
			"msg":           err,
			"file":          fmt.Sprintf("%s:%d", file, line),
			"response_code": statusCode,
			"user": map[string]string{
				"email":      email,
				"membership": membership,
			},
		}).Error(err)
	}()
}

func LogWarn(ctx context.Context, err error, statusCode int) {
	_, file, line, _ := runtime.Caller(1)

	go func() {
		log.SetLevel(logrus.WarnLevel)
		LogInit(log)

		var (
			email      string
			membership string
		)

		emailRes := ctx.Value(constants.Key("email"))
		if emailRes != nil {
			email = emailRes.(string)
		}
		membershipRes := ctx.Value(constants.Key("membership"))
		if membershipRes != nil {
			membership = membershipRes.(string)
		}

		log.WithFields(logrus.Fields{
			"msg":           err.Error(),
			"file":          fmt.Sprintf("%s:%d", file, line),
			"response_code": statusCode,
			"user": map[string]string{
				"email":      email,
				"membership": membership,
			},
		}).Warn(err)
	}()
}

func LogWarnStr(ctx context.Context, err string, statusCode int) {
	_, file, line, _ := runtime.Caller(1)

	go func() {
		log.SetLevel(logrus.WarnLevel)
		LogInit(log)

		var (
			email      string
			membership string
		)

		emailRes := ctx.Value(constants.Key("email"))
		if emailRes != nil {
			email = emailRes.(string)
		}
		membershipRes := ctx.Value(constants.Key("membership"))
		if membershipRes != nil {
			membership = membershipRes.(string)
		}

		log.WithFields(logrus.Fields{
			"msg":           err,
			"file":          fmt.Sprintf("%s:%d", file, line),
			"response_code": statusCode,
			"user": map[string]string{
				"email":      email,
				"membership": membership,
			},
		}).Warn(err)
	}()
}

func LogTrace(ctx context.Context, err interface{}) {
	_, file, line, _ := runtime.Caller(1)

	go func() {
		LogInit(log)

		var (
			email      string
			membership string
		)

		emailRes := ctx.Value(constants.Key("email"))
		if emailRes != nil {
			email = emailRes.(string)
		}
		membershipRes := ctx.Value(constants.Key("membership"))
		if membershipRes != nil {
			membership = membershipRes.(string)
		}

		log.WithFields(logrus.Fields{
			"msg":  err,
			"file": fmt.Sprintf("%s:%d", file, line),
			"user": map[string]string{
				"email":      email,
				"membership": membership,
			},
		}).Trace(err)
	}()
}

func LogDebug(ctx context.Context, err error, statusCode int) {
	_, file, line, _ := runtime.Caller(1)

	go func() {
		LogInit(log)

		var (
			email      string
			membership string
		)

		emailRes := ctx.Value(constants.Key("email"))
		if emailRes != nil {
			email = emailRes.(string)
		}
		membershipRes := ctx.Value(constants.Key("membership"))
		if membershipRes != nil {
			membership = membershipRes.(string)
		}

		log.WithFields(logrus.Fields{
			"msg":  err,
			"file": fmt.Sprintf("%s:%d", file, line),
			"user": map[string]string{
				"email":      email,
				"membership": membership,
			},
		}).Debug(err)
	}()
}
