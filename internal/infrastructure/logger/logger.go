package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

type Field struct {
	Key   string
	Value interface{}
}

func NewLogger() (*Logger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{logger: zapLogger}, nil
}

func toZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = zap.Any(field.Key, field.Value)
	}
	return zapFields
}

func (z *Logger) Debug(msg string, fields ...Field) {
	z.logger.Debug(msg, toZapFields(fields)...)
}

func (z *Logger) Info(msg string, fields ...Field) {
	z.logger.Info(msg, toZapFields(fields)...)
}

func (z *Logger) Warn(msg string, fields ...Field) {
	z.logger.Warn(msg, toZapFields(fields)...)
}

func (z *Logger) Error(msg string, fields ...Field) {
	z.logger.Error(msg, toZapFields(fields)...)
}

func (z *Logger) Fatal(msg string, fields ...Field) {
	z.logger.Fatal(msg, toZapFields(fields)...)
}

func (z *Logger) With(key string, value interface{}) *Logger {
	newLogger := z.logger.With(zap.Any(key, value))
	return &Logger{newLogger}
}

func (z *Logger) WithError(err error) *Logger {
	newLogger := z.logger.With(zap.Error(err))
	return &Logger{logger: newLogger}
}

func (z *Logger) Close() {
	_ = z.logger.Sync()
}
