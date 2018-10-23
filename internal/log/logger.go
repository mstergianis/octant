package log

import "go.uber.org/zap"

// Logger is an interface for logging
type Logger interface {
	// Debugf uses fmt.Sprintf to log a templated message.
	Debugf(template string, args ...interface{})

	// Infof uses fmt.Sprintf to log a templated message.
	Infof(template string, args ...interface{})

	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(template string, args ...interface{})

	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(template string, args ...interface{})

	With(args ...interface{}) Logger

	Named(name string) Logger
}

// sugaredLogWrapper adapts a zap.SuggaredLogger to the Logger interface
type sugaredLogWrapper struct {
	*zap.SugaredLogger
}

func (s *sugaredLogWrapper) With(args ...interface{}) Logger {
	return &sugaredLogWrapper{s.SugaredLogger.With(args...)}
}

func (s *sugaredLogWrapper) Named(name string) Logger {
	return &sugaredLogWrapper{s.SugaredLogger.Named(name)}
}

var _ Logger = (*sugaredLogWrapper)(nil)

// Wrap zap.SugaredLogger as Logger interface
func Wrap(z *zap.SugaredLogger) Logger {
	return &sugaredLogWrapper{z}
}

// NopLogger constructs a nop logger
func NopLogger() Logger {
	return Wrap(zap.NewNop().Sugar())
}