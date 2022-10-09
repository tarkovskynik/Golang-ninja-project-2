package logger

import log "github.com/sirupsen/logrus"

func LogrusLogger(err error, handler, problem string) {
	log.WithFields(log.Fields{
		"handler": handler,
		"problem": problem,
	}).Error(err)
}
