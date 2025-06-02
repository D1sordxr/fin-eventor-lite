package worker

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"log/slog"
)

type App struct {
	log ports.Log
}

func NewApp() *App {
	log := slog.Default()

	return &App{
		log: log,
	}
}

func (a *App) Run() {}
