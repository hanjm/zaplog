A wrapped zap log tool
```
package zaplog
    import "zaplog"


FUNCTIONS

func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder)
    CallerEncoder will add caller to log. format is
    "filename:lineNum:funcName",
    e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"

func NewCompatibleLogger(debugLevel bool) (compatibleLogger *zapgrpc.Logger)
    NewCompatibleLogger return a logger which compatible to std log package
    function. log level is info log.Print(v ...interface) log.Printf(format
    string, v ...interface) log.Println(v ...interface{}) log.Panic()
    log.Panicf() log.Fatal() log.Fatalf()

func NewCustomLoggers(debugLevel bool) (logger, noCallerLogger *zap.Logger, compatibleLogger *zapgrpc.Logger)
    NewCustomLoggers is a shortcut to get normal logger, noCallerLogger and
    compatibleLogger

func NewLogger(debugLevel bool) (logger *zap.Logger)
    NewLogger return a normal logger

func NewNoCallerLogger(debugLevel bool) (noCallerLogger *zap.Logger)
    NewNoCallerLogger return a no caller key value, will be faster

```
