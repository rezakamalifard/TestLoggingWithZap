package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Message is a general user defined type that we can use.
type Message struct {
	s string
	i int64
}

// Add MarshalLogObject to implement ObjectEncoder interface to Marshal Message type for logging.
func (m *Message) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("s", m.s)
	enc.AddInt64("index", m.i)
	return nil
}

func sugarLogger(sugar *zap.SugaredLogger, msg *Message) {
	sugar.Infow("logger",
		"type", "logger",
		"status", "OK",
		"error", 0,
		"message", msg,
	)
}

func structuredLogger(logger *zap.Logger, msg *Message) {
	logger.Info("logger",
		zap.String("type", "logger"),
		zap.String("status", "OK"),
		zap.Int("error", 0),
		zap.Object("message", msg))
}
