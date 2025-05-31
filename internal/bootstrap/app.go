package bootstrap

import (
	"context"
	userSvc "github.com/D1sordxr/fin-eventor-lite/internal/application/user"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/shared"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/shared/interfaces"
	appSrv "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/delivery/user"
	middleware2 "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	log    interfaces.Log
	Server *appSrv.Server
}

func NewApp() *App {
	log := slog.Default()

	chainer := new(shared.Chainer)

	logMid := middleware2.NewLogMid(log)
	methodPostMid := middleware2.NewMethodMid(http.MethodPost)
	semaphoreMid := middleware2.NewSemaphoreMid()
	retryMid := new(middleware2.RetryMid)

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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

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
		if err := a.Server.Shutdown(ctx); err != nil {
			a.log.Error("Failed to shutdown server: " + err.Error())
		}
	case err := <-errChan:
		a.log.Error("Server error: " + err.Error())
	}

	appsWg.Wait()
	a.log.Info("Server stopped gracefully")
}
