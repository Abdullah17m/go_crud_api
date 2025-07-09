package logger

import (
	log "github.com/sirupsen/logrus"
	sentry "github.com/evalphobia/logrus_sentry"
)

func InitLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	hook, err := sentry.NewSentryHook("https://your_sentry_dsn", []log.Level{log.ErrorLevel})
	if err == nil {
		log.AddHook(hook)
	}
}
