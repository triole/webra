package logging

import (
	"github.com/sirupsen/logrus"
)

// Info logs successful tests
func (lg Logging) Info(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Info(msg)
	default:
		lg.Logrus.Info(msg)
	}
}

// Error logs failed tests
func (lg Logging) Error(msg interface{}, fields interface{}) {
	var msgStr string
	switch val := msg.(type) {
	case error:
		msgStr = val.Error()
	default:
		msgStr = val.(string)
	}
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Error(msgStr)
	default:
		lg.Logrus.Error(msgStr)
	}
}

// LogFatal logs fatal and exits
func (lg Logging) LogFatal(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Fatal(msg)
	default:
		lg.Logrus.Fatal(msg)
	}
}
