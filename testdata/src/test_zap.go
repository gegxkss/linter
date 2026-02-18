package src

import (
	"testing"

	"go.uber.org/zap/zaptest"
)

func TestZapLogger(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Info("Starting server")
	sugar := logger.Sugar()
	sugar.Infof("Server started on port %d", 8080)
	logger.Error("Failed to connect")
}
