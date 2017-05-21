## zaplog是包装了[zap](https://github.com/uber-go/zap)、用于生产环境的日志输出工具
[![GoDoc](https://godoc.org/github.com/hanjm/zaplog?status.svg)](https://godoc.org/github.com/hanjm/zaplog)
[![Go Report Card](https://goreportcard.com/badge/github.com/hanjm/zaplog)](https://goreportcard.com/report/github.com/hanjm/zaplog)
[![code-coverage](http://gocover.io/_badge/github.com/hanjm/zaplog)](http://gocover.io/github.com/hanjm/zaplog)

### Document

```go
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder)
    CallerEncoder will add caller to log. format is
    "filename:lineNum:funcName",
    e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"

func FormatStdLog()
    FormatStdLog set the output of stand package log to zaplog

func NewCustomLoggers(debugLevel bool) (logger, noCallerLogger *zap.Logger)
    NewCustomLoggers is a shortcut to get normal logger, noCallerLogger.

func NewLogger(debugLevel bool) (logger *zap.Logger)
    NewLogger return a normal logger

func NewNoCallerLogger(debugLevel bool) (noCallerLogger *zap.Logger)
    NewNoCallerLogger return a no caller key value, will be faster

TYPES

type CompatibleLogger struct {
    // contains filtered or unexported fields
}
    CompatibleLogger is a logger which compatible to logrus/std
    log/prometheus. it implements Print() Println() Printf() Dbug()
    Debugln() Debugf() Info() Infoln() Infof() Warn() Warnln() Warnf()
    Error() Errorln() Errorf() Fatal() Fataln() Fatalf() Panic() Panicln()
    Panicf() With()

func NewCompatibleLogger(debugLevel bool) *CompatibleLogger
    NewCompatibleLogger return CompatibleLogger with caller field

func (l CompatibleLogger) Debug(args ...interface{})
    Debug logs a message at level Debug on the compatibleLogger.

func (l CompatibleLogger) Debugf(format string, args ...interface{})
    Debugf logs a message at level Debug on the compatibleLogger.

func (l CompatibleLogger) Debugln(args ...interface{})
    Debugln logs a message at level Debug on the compatibleLogger.

func (l CompatibleLogger) Error(args ...interface{})
    Error logs a message at level Error on the compatibleLogger.

func (l CompatibleLogger) Errorf(format string, args ...interface{})
    Errorf logs a message at level Error on the compatibleLogger.

func (l CompatibleLogger) Errorln(args ...interface{})
    Errorln logs a message at level Error on the compatibleLogger.

func (l CompatibleLogger) Fatal(args ...interface{})
    Fatal logs a message at level Fatal on the compatibleLogger.

func (l CompatibleLogger) Fatalf(format string, args ...interface{})
    Fatalf logs a message at level Fatal on the compatibleLogger.

func (l CompatibleLogger) Fatalln(args ...interface{})
    Fatalln logs a message at level Fatal on the compatibleLogger.

func (l CompatibleLogger) Info(args ...interface{})
    Info logs a message at level Info on the compatibleLogger.

func (l CompatibleLogger) Infof(format string, args ...interface{})
    Infof logs a message at level Info on the compatibleLogger.

func (l CompatibleLogger) Infoln(args ...interface{})
    Infoln logs a message at level Info on the compatibleLogger.

func (l CompatibleLogger) Panic(args ...interface{})
    Panic logs a message at level Painc on the compatibleLogger.

func (l CompatibleLogger) Panicf(format string, args ...interface{})
    Panicf logs a message at level Painc on the compatibleLogger.

func (l CompatibleLogger) Panicln(args ...interface{})
    Panicln logs a message at level Painc on the compatibleLogger.

func (l CompatibleLogger) Print(args ...interface{})
    Print logs a message at level Info on the compatibleLogger.

func (l CompatibleLogger) Printf(format string, args ...interface{})
    Printf logs a message at level Info on the compatibleLogger.

func (l CompatibleLogger) Println(args ...interface{})
    Println logs a message at level Info on the compatibleLogger.

func (l CompatibleLogger) Warn(args ...interface{})
    Warn logs a message at level Warn on the compatibleLogger.

func (l CompatibleLogger) Warnf(format string, args ...interface{})
    Warnf logs a message at level Warn on the compatibleLogger.

func (l CompatibleLogger) Warnln(args ...interface{})
    Warnln logs a message at level Warn on the compatibleLogger.

func (l *CompatibleLogger) With(key string, value interface{}) *CompatibleLogger
    With return a logger with a extra field.

func (l *CompatibleLogger) WithField(key string, value interface{}) *CompatibleLogger
    WithField return a logger with a extra field.

func (l *CompatibleLogger) WithFields(fields map[string]interface{}) *CompatibleLogger
    WithField return a logger with a extra field.


```

