package goq

import "github.com/sirupsen/logrus"

var l *logrus.Logger

func init() {
	l = logrus.New()
}

func SetLogger(newL *logrus.Logger) {
	l = newL
}
