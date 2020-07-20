package main

import (
	"testing"

	"go.uber.org/zap"
)

func BenchmarkSugarLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer sugar.Sync()
	msg := &Message{"message_text", 1}

	for i := 0; i < b.N; i++ {
		sugarLogger(sugar, msg)
	}
}
func BenchmarkStructuredLogger(b *testing.B) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	msg := &Message{"message_text", 1}

	for i := 0; i < b.N; i++ {
		structuredLogger(logger, msg)
	}
}
