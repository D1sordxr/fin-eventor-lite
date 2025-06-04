package main

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/bootstrap/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := api.NewApp(ctx)
	app.Run(ctx)
}
