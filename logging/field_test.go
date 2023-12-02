package logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "bambang").Info("Hello World")

	logger.WithField("username", "bambang").
		WithField("name", "Bambang Cahyo").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "bambang",
		"name":     "Bambang Cahyo",
	}).Info("Hello World")
}
