package main

import (
	"fmt"
	"go.uber.org/zap"
	"loggo/helper"
	"time"
)

var (
	url  = "https://error.com"
	name = "GoLang"
)

func LoggerOne(l *zap.Logger) {
	l.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func LoggerTwo(l *zap.Logger) {
	sugar := l.Sugar()

	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "",
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Infof("Failed to fetch URL: %s", url)
}

func previewHelperFunction() {
	helper.Desc = "Public Variable"
	helperName := helper.GetHelperName()

	fmt.Printf(helperName)
}

func previewStopWatch() {
	clock := new(helper.StopWatch)
	clock.Desc = "My Stop Watch"
	clock.Start()

	time.Sleep(1 * time.Second)

	clock.Stop()
	duration := clock.GetPassingTime()

	fmt.Printf("%s Duration = %f\n", clock.Desc, duration)
}

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	LoggerTwo(logger)
	LoggerTwo(logger)

	previewStopWatch()
	previewHelperFunction()
}
