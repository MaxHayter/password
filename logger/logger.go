package logger

import (
	"context"
	"log"
	"os"
)

type ctxlog struct{}

func WithLogger(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, ctxlog{}, logger)
}

var DefaultLogger = log.New(os.Stdout, "TgBot: ", log.LstdFlags)

// GetLogger get logger from context, or DefaultLogger if not exists
func GetLogger(ctx context.Context) *log.Logger {
	return GetLoggerDefault(ctx, DefaultLogger)
}

// GetLoggerDefault get logger from context
// defLogger - logger that is used if logger is not found in context
func GetLoggerDefault(ctx context.Context, defLogger *log.Logger) *log.Logger {
	l, ok := ctx.Value(ctxlog{}).(*log.Logger)
	if !ok {
		l = defLogger
	}
	return l
}
