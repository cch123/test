package main

import "time"
import "go.uber.org/zap"

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()

}

func logto() {
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "aaa",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "aaa")
}

func main() {
	logto()
}
