package logger

import "go.uber.org/zap"

type ZapLogger struct {
	l *zap.Logger
}

func NewZapLogger(l *zap.Logger) *ZapLogger {
	return &ZapLogger{
		l: l,
	}
}

func (t *ZapLogger) Debug(msg string, data []LogData) {
	zapData := t.logDataToZapData(data)
	t.l.Debug(msg, zapData...)
}

func (t *ZapLogger) Info(msg string, data []LogData) {
	zapData := t.logDataToZapData(data)
	t.l.Info(msg, zapData...)
}

func (t *ZapLogger) Warn(msg string, data []LogData) {
	zapData := t.logDataToZapData(data)
	t.l.Warn(msg, zapData...)
}

func (t *ZapLogger) Error(msg string, data []LogData) {
	zapData := t.logDataToZapData(data)
	t.l.Error(msg, zapData...)
}

func (t *ZapLogger) saveReqLog(msg string, data []LogData) {
	zapData := t.logDataToZapData(data)
	t.l.Error(msg, zapData...)
}

func (t *ZapLogger) logDataToZapData(data []LogData) []zap.Field {
	zapData := make([]zap.Field, 0, len(data))
	for _, d := range data {
		zapData = append(zapData, zap.Any(d.Key, d.Val))
	}
	return zapData
}
