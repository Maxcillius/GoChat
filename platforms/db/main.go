package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Maxcillius/GoChat/pkg/logger"
	"github.com/Maxcillius/GoChat/platforms/db/grpc"
	"golang.org/x/sys/unix"
)

func Main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			panic(fmt.Sprintf("failed to write log: `%s` original error is: %s", ferr, err))
		}
		return 1
	}

	clogger := l.WithName("db")

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.RunServer(ctx, 5000, clogger.WithName("grpc"))
	}()

	select {
	case err := <-errCh:
		fmt.Println(err.Error())
		return 1
	case <-ctx.Done():
		fmt.Println("shutting down...")
		return 0
	}
}
