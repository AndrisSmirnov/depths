package main

import (
	"context"
	"depths/voc"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"depths/app"
	"depths/pkg/log"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	configLogger()

	log.Info("Depths microservice launch...")

	application, err := app.CreateApplication(ctx, cancelFunc)
	if err != nil {
		log.Panic(err)
	}

	if err := application.Launch(); err != nil {
		log.Panic(err)
	}

	log.Info("Depths microservice launched (/◔ ◡ ◔)/")

	go syscallWait(cancelFunc)
	<-ctx.Done()

	log.Info("Depths microservice stopping...")

	if err := application.Stop(); err != nil {
		log.Warning(
			fmt.Sprintf("%s:%v",
				voc.ErrorApplicationStop,
				err,
			))
	}

	log.Info("Depths microservice stopped.")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}

func configLogger() {
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = int(log.DEBUG)
	}

	log.SetLevel(log.LogLevel(logLevel))
}
