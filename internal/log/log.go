package log

import (
	"github.com/khulnasoft/binpack/internal/redact"
	"github.com/khulnasoft-lab/go-logger"
	"github.com/khulnasoft-lab/go-logger/adapter/discard"
	redactLogger "github.com/khulnasoft-lab/go-logger/adapter/redact"
)

// log is the singleton used to facilitate logging internally within
var log = discard.New()

func Set(l logger.Logger) {
	// though the application will automatically have a redaction logger, library consumers may not be doing this.
	// for this reason we additionally ensure there is a redaction logger configured for any logger passed. The
	// source of truth for redaction values is still in the internal redact package. If the passed logger is already
	// redacted, then this is a no-op.
	store := redact.Get()
	if store != nil {
		l = redactLogger.New(l, store)
	}
	log = l
}

func Get() logger.Logger {
	return log
}

// Errorf takes a formatted template string and template arguments for the error logging level.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Error logs the given arguments at the error logging level.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Warnf takes a formatted template string and template arguments for the warning logging level.
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Warn logs the given arguments at the warning logging level.
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Infof takes a formatted template string and template arguments for the info logging level.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Info logs the given arguments at the info logging level.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Debugf takes a formatted template string and template arguments for the debug logging level.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Debug logs the given arguments at the debug logging level.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Tracef takes a formatted template string and template arguments for the trace logging level.
func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

// Trace logs the given arguments at the trace logging level.
func Trace(args ...interface{}) {
	log.Trace(args...)
}

// WithFields returns a message logger with multiple key-value fields.
func WithFields(fields ...interface{}) logger.MessageLogger {
	return log.WithFields(fields...)
}

// Nested returns a new logger with hard coded key-value pairs
func Nested(fields ...interface{}) logger.Logger {
	return log.Nested(fields...)
}
