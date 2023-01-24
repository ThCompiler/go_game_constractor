package context

import (
	"context"

	"github.com/ThCompiler/go_game_constractor/pkg/logger"
)

type LogObject struct {
	log logger.Interface
}

func NewLogObject(log logger.Interface) LogObject {
	return LogObject{log: log}
}

func (l *LogObject) BaseLog() logger.Interface {
	return l.log
}

func (l *LogObject) Log(ctx context.Context) logger.Interface {
	if ctx == nil {
		return l.log.With("type", "base_log")
	}

	ctxLogger := ctx.Value(LoggerField)
	log := l.log

	if ctxLogger != nil {
		if ctxLog, ok := ctxLogger.(logger.Interface); ok {
			log = ctxLog
		}
	}

	return log
}

func GetLogger(ctx context.Context) logger.Interface {
	ctxLogger := ctx.Value(LoggerField)
	if ctxLog, ok := ctxLogger.(logger.Interface); ok {
		return ctxLog
	}

	return nil
}

func SaveLogger(ctx context.Context, log logger.Interface) context.Context {
	return context.WithValue(ctx, LoggerField, log)
}
