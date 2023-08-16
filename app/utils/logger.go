package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// ErrLoggerInit возникает при сбое инициализации регистратора
var ErrLoggerInit = errors.New("failed to initialize logger")

// Log - экземпляр глобального регистратора
var Log *logrus.Logger

// InitLogger initializes the global logger
func InitLogger() error {

	Log = logrus.New()

	// Set log level
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return ErrLoggerInit
	}
	Log.Level = logLevel

	// Set log format to JSON
	Log.Formatter = &logrus.JSONFormatter{}

	return nil
}

// Error оборачивает logger.Error() и добавляет сообщение
func Error(msg string, err error) {
	if err != nil {
		Log.WithError(err).Error(msg)
	}
}

// Info оборачивает logger.Info() и добавляет сообщение
func Info(msg string) {
	Log.Info(msg)
}

// Debug оборачивает logger.Debug() и добавляет сообщение
func Debug(msg string) {
	Log.Debug(msg)
}
