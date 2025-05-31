package worker

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/shared/interfaces"
	"log/slog"
)

type App struct {
	log interfaces.Log
}

func NewApp() *App {
	log := slog.Default()

	return &App{
		log: log,
	}
}

func (a *App) Run() {}
