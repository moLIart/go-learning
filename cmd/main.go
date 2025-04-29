package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moLIart/go-course/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		cancel()
	}()

	service.StartProcessing(ctx, 19*time.Millisecond)
}
