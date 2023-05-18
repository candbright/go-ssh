package ssh

import (
	"github.com/candbright/go-log/log"
	"github.com/sirupsen/logrus"
)

func Success(logger *log.Logger, name string, arg []string, output string) {
	if logger != nil {
		logger.WithFields(logrus.Fields{
			"status": "success",
			"cmd":    Command(name, arg...),
		}).Info(output)
	}
}

func Fail(logger *log.Logger, name string, arg []string, output string, err error) {
	if logger != nil {
		logger.WithFields(logrus.Fields{
			"status": "fail",
			"cmd":    Command(name, arg...),
			"error":  err,
		}).Info(output)
	}
}
