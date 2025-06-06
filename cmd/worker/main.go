package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/D1sordxr/fin-eventor-lite/internal/bootstrap/worker"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := worker.NewApp(ctx)
	app.Run(ctx)
}
