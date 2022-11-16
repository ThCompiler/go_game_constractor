package http

import (
    "context"
    "github.com/ThCompiler/go_game_constractor/pkg/logger"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "time"
)

func GinRequestLogger(l logger.Interface) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start timer
        start := time.Now()
        path := c.Request.URL.Path
        raw := c.Request.URL.RawQuery

        method := c.Request.Method
        requestId := uuid.New()

        if raw != "" {
            path = path + "?" + raw
        }

        lg := l.With(URL, path).With(RequestId, requestId).With(Method, method)
        c.Set(ContextLoggerField, lg)

        // Process request
        c.Next()

        // Stop timer
        timeStamp := time.Now()
        latency := timeStamp.Sub(start)

        clientIP := c.ClientIP()

        statusCode := c.Writer.Status()
        errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

        if latency > time.Minute {
            latency = latency.Truncate(time.Second)
        }

        l.Info("[GIN] %v | %d | %v | %s | %s  %v | %s |",
            timeStamp.Format("2006/01/02 - 15:04:05"),
            statusCode,
            latency,
            clientIP,
            method,
            path,
            errorMessage,
        )
    }
}

func GetLogger(ctx context.Context) logger.Interface {
    ctxLogger := ctx.Value(ContextLoggerField)
    if ctxLog, ok := ctxLogger.(logger.Interface); ok {
        return ctxLog
    }
    return nil
}
