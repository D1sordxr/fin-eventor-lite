package worker

import (
	"context"
	"log/slog"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
)

type App struct {
	log ports.Log
}

func NewApp(ctx context.Context) *App {
	log := slog.Default()

	_ = ctx

	return &App{
		log: log,
	}
}

func (a *App) Run(ctx context.Context) {
	a.log.Info("Worker application is starting...")

	<-ctx.Done()
	a.log.Info("Worker application is shutting down gracefully")
}
