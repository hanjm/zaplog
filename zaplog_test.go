package zaplog

import (
	"log"
	"testing"
)

func TestFormatStdLog(t *testing.T) {
	FormatStdLog()
	log.Print("std log")
	log.Print("")
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger(true)
	defer logger.Sync()
	logger.Debug("zap log debug msg")
}

func TestNewCustomLoggers(t *testing.T) {
	logger, noCallerLogger, compatibleLogger := NewCustomLoggers(true)
	logger.Info("normalLogger")
	noCallerLogger.Info("noCallerLogger")
	compatibleLogger.Print("compatibleLogger")
}

func BenchmarkStdLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print("std log printf")
	}
}

func BenchmarkNewLogger(b *testing.B) {
	logger := NewLogger(true)
	defer logger.Sync()
	for i := 0; i < b.N; i++ {
		logger.Debug("zap log debug msg")
	}
}
