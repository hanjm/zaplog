## A wrapped zap log tool

zaplog是包装了[zap](https://github.com/uber-go/zap)、用于生产环境的日志输出工具, 它通过设置log.SetOutput()使所有使用标准库log的包的日志输出标准化为行json格式.
NewLogger()返回一个会记录日志发生行的文件名的logger, 格式为 文件名:行号:函数名, 如"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"
NewNoCallerLogger()返回一个不记录日志发生行的logger, 用于http api server的access log.
## Document 

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
