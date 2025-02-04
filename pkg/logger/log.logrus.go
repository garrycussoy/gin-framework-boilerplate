package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	// Create a new logger instance
	log = logrus.New()

	// Set output to stdout
	log.SetOutput(os.Stdout)

	// Set log formatter
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// Define some log functions based on level
func Debug(message string, fields logrus.Fields) {
	log.WithFields(fields).Debug(message)
}

func DebugF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Debugf(format, args...)
}

func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

func InfoF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Infof(format, args...)
}

func Error(message string, fields logrus.Fields) {
	log.WithFields(fields).Error(message)
}

func ErrorF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Errorf(format, args...)
}

func Fatal(message string, fields logrus.Fields) {
	log.WithFields(fields).Fatal(message)
}

func FatalF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Fatalf(format, args...)
}

func Panic(message string, fields logrus.Fields) {
	log.WithFields(fields).Panic(message)
}

func PanicF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Panicf(format, args...)
}
