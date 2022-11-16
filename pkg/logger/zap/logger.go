package zap

import (
    "fmt"
    log "github.com/ThCompiler/go_game_constractor/pkg/logger"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "io"
    "os"
    "strings"
)

type Params struct {
    AppName                  string
    LogDir                   string
    Level                    log.LogLevel
    UseStdAndFIle            bool
    AddLowPriorityLevelToCmd bool
}

// Logger -.
type Logger struct {
    logger *zap.SugaredLogger
}

var _ log.Interface = (*Logger)(nil)

// New -.
func New(param Params, out io.Writer) *Logger {
    core := newZapCore(param, out)
    zap.NewProductionConfig()
    logger := zap.New(core)

    sugLogger := logger.Sugar()

    return &Logger{
        logger: sugLogger.With(log.AppName, param.AppName),
    }
}

func toZapLevel(level log.LogLevel) zapcore.Level {
    switch log.LogLevel(strings.ToLower(string(level))) {
    case log.ErrorLevel:
        return zap.ErrorLevel
    case log.WarnLevel:
        return zap.WarnLevel
    case log.InfoLevel:
        return zap.InfoLevel
    case log.DebugLevel:
        return zap.DebugLevel
    case log.PanicLevel:
        return zap.PanicLevel
    case log.FatalLevel:
        return zap.FatalLevel
    default:
        return zap.InfoLevel
    }
}

func newZapCore(param Params, out io.Writer) (core zapcore.Core) {
    // First, define our level-handling logic.
    highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= toZapLevel(param.Level)
    })

    if param.AddLowPriorityLevelToCmd { // separate levels
        lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
            return lvl < toZapLevel(param.Level)
        })

        topicDebugging := zapcore.AddSync(out)
        topicErrors := zapcore.AddSync(out)
        fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

        if param.UseStdAndFIle && param.LogDir != "" {
            consoleDebugging := zapcore.Lock(os.Stdout)
            consoleErrors := zapcore.Lock(os.Stderr)
            consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

            core = zapcore.NewTee(
                zapcore.NewCore(fileEncoder, topicErrors, highPriority),
                zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
                zapcore.NewCore(fileEncoder, topicDebugging, lowPriority),
                zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
            )
        } else {
            core = zapcore.NewTee(
                zapcore.NewCore(fileEncoder, topicErrors, highPriority),
                zapcore.NewCore(fileEncoder, topicDebugging, lowPriority),
            )
        }
    } else { // not separate levels
        topicErrors := zapcore.AddSync(out)
        fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

        if param.UseStdAndFIle && param.LogDir != "" {
            consoleErrors := zapcore.Lock(os.Stderr)
            consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

            core = zapcore.NewTee(
                zapcore.NewCore(fileEncoder, topicErrors, highPriority),
                zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
            )
        } else {
            core = zapcore.NewTee(
                zapcore.NewCore(fileEncoder, topicErrors, highPriority),
            )
        }
    }
    return core
}

func (l *Logger) Sync() error {
    return l.logger.Sync()
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
    l.log(l.logger.Debugf, message, args...)
}

// Info -.
func (l *Logger) Info(message interface{}, args ...interface{}) {
    l.log(l.logger.Infof, message, args...)
}

// Warn -.
func (l *Logger) Warn(message interface{}, args ...interface{}) {
    l.log(l.logger.Warnf, message, args...)
}

// Panic -.
func (l *Logger) Panic(message interface{}, args ...interface{}) {
    l.log(l.logger.Panicf, message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
    l.log(l.logger.Errorf, message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
    l.log(l.logger.Fatalf, message, args...)
}

func (l *Logger) log(lg func(message string, args ...interface{}), message interface{}, args ...interface{}) {
    switch tp := message.(type) {
    case error:
        lg(tp.Error(), args...)
    case string:
        lg(tp, args...)
    default:
        lg(fmt.Sprintf("message %v has unknown type %v", message, tp), args...)
    }
}

func (l *Logger) With(key log.Field, value interface{}) log.Interface {
    return &Logger{l.logger.With(string(key), value)}
}
