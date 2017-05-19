package zaplog

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
	"log"
	"runtime"
	"strings"
)

func init() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(&zapWriter{NewNoCallerLogger(false)})
}

type zapWriter struct {
	logger *zap.Logger
}

// Write implement io.Writer, as std log's output
func (w zapWriter) Write(p []byte) (n int, err error) {
	i := bytes.Index(p, []byte(":")) + 1
	j := bytes.Index(p[i:], []byte(":")) + 1 + i
	caller := bytes.TrimRight(p[:j], ":")
	// last index of /
	i = bytes.LastIndex(caller, []byte("/"))
	// penultimate index of /
	i = bytes.LastIndex(caller[:i], []byte("/"))
	w.logger.Info("stdLog", zap.ByteString("caller", caller[i+1:]), zap.ByteString("log", bytes.TrimSpace(p[j:])))
	return len(p), nil
}

// CallerEncoder will add caller to log. format is "filename:lineNum:funcName", e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

func newLoggerConfig(debugLevel bool) (loggerConfig zap.Config) {
	loggerConfig = zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.EncodeCaller = CallerEncoder
	if debugLevel {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return
}

// NewCustomLoggers is a shortcut to get normal logger, noCallerLogger and compatibleLogger
func NewCustomLoggers(debugLevel bool) (logger, noCallerLogger *zap.Logger, compatibleLogger *zapgrpc.Logger) {
	loggerConfig := newLoggerConfig(debugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	loggerConfig.DisableCaller = true
	noCallerLogger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	compatibleLogger = zapgrpc.NewLogger(logger.WithOptions(zap.AddCallerSkip(1)))
	return
}

// NewLogger return a normal logger
func NewLogger(debugLevel bool) (logger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewNoCallerLogger return a no caller key value, will be faster
func NewNoCallerLogger(debugLevel bool) (noCallerLogger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel)
	loggerConfig.DisableCaller = true
	noCallerLogger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewCompatibleLogger return a logger which compatible to std log package function. log level is info
// log.Print(v ...interface) log.Printf(format string, v ...interface) log.Println(v ...interface{})
// log.Panic() log.Panicf()
// log.Fatal() log.Fatalf()
func NewCompatibleLogger(debugLevel bool) (compatibleLogger *zapgrpc.Logger) {
	compatibleLogger = zapgrpc.NewLogger(NewNoCallerLogger(debugLevel).WithOptions(zap.AddCallerSkip(1)))
	return
}
