package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	sigctx, cancel1 := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel1()

	toctx, cancel2 := context.WithTimeout(ctx, time.Second*5)
	defer cancel2()

	select {
	case <-sigctx.Done():
		fmt.Println("signal")
	case <-toctx.Done():
		fmt.Println("timeout")
	}
}
