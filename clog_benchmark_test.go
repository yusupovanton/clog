package clog_test

import (
	"bytes"
	"context"
	"github.com/yusupovanton/clog"
	"log/slog"
	"testing"
	"time"
)

func BenchmarkCustomLogger(b *testing.B) {
	var buf bytes.Buffer

	logger := clog.NewCustomLogger(&buf, false, slog.LevelDebug)

	ctx := logger.AddKeysValuesToCtx(context.Background(), map[string]interface{}{
		"userID":    12345,
		"userName":  "testuser",
		"timestamp": time.Now(),
		"data":      []int{1, 2, 3},
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.InfoCtx(ctx, "Some test message")
	}
}
