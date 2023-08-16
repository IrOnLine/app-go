package utils

import (
  "os"

  "github.com/sirupsen/logrus"
)

// ErrLoggerInit occurs when logger initialization fails
var ErrLoggerInit = errors.New("failed to initialize logger")

// Log is the global logger instance
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

// Error wraps logger.Error() and adds message 
func Error(msg string, err error) {
  if err != nil {
    Log.WithError(err).Error(msg)
  }
} 

// Info wraps logger.Info() and adds message
func Info(msg string) {
  Log.Info(msg)
}

// Debug wraps logger.Debug() and adds message  
func Debug(msg string) {
  Log.Debug(msg) 
}