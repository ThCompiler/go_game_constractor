package http

import "github.com/ThCompiler/go_game_constractor/pkg/logger"

const ContextLoggerField = "logger"

const (
    RequestId logger.Field = "request_id"
    Method    logger.Field = "method"
    URL       logger.Field = "url"
)
