package internal

import (
	"context"
	userSvc "fin-eventor-lite/internal/application/user"
	appSrv "fin-eventor-lite/internal/presentation/http"
	"fin-eventor-lite/internal/presentation/http/delivery/middleware"
	"fin-eventor-lite/internal/presentation/http/delivery/user"
	"fin-eventor-lite/pkg"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	log    pkg.Log
	Server *appSrv.Server
}

func NewApp() *App {
	log := slog.Default()

	chainer := new(middleware.Chainer)

	logMid := middleware.NewLogMid(log)
	methodPostMid := middleware.NewMethodMid(http.MethodPost)
	semaphoreMid := middleware.NewSemaphoreMid()
	retryMid := new(middleware.RetryMid)

	userHandler := user.NewHandler(
		new(userSvc.MockUseCase),
		chainer,
		logMid.Log,
		methodPostMid.OnlyPost,
		semaphoreMid.Limit,
		retryMid.RetryWithBackoff,
	)

	server := appSrv.NewServer(
		"8080",
		userHandler,
	)

	return &App{
		log:    log,
		Server: server,
	}
}

func (a *App) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	appsWg := &sync.WaitGroup{}
	errChan := make(chan error, 1)

	appsWg.Add(1)
	go func() {
		defer appsWg.Done()
		err := a.Server.StartServer()
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		a.log.Info("Received shutdown signal, shutting down...")
	case err := <-errChan:
		_ = err
		a.log.Error("Server error: %v", err)
	}

	appsWg.Wait()
	a.log.Info("Shutting down gracefully...")
}
