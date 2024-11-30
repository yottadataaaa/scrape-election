package loglib

import (
	"context"
	"fmt"
	"log"
)

var LoggerKey = struct{}{}

func UseLogger(ctx context.Context) (*log.Logger, error) {
	logger, ok := ctx.Value(LoggerKey).(*log.Logger)
	if !ok {
		return nil, fmt.Errorf("logger not found by LoggerKey: %v", LoggerKey)
	}

	return logger, nil
}
