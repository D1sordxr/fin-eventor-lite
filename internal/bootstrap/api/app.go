package api

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc"
	appSrv "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http"
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	log        ports.Log
	HttpServer *appSrv.Server
	GrpcServer *grpc.Server
}

func NewApp() *App {
	log := slog.Default()

	// TODO: ReadConfig()

	server := setupHTTP(
		log,
		"9090", // TODO: Read from config
	)

	return &App{
		log:        log,
		HttpServer: server,
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
		err := a.HttpServer.StartServer()
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		a.log.Info("Received shutdown signal, shutting down...")
		if err := a.HttpServer.Shutdown(ctx); err != nil {
			a.log.Error("Failed to shutdown HTTP server: " + err.Error())
		}
	case err := <-errChan:
		a.log.Error("App error: " + err.Error())
	}

	appsWg.Wait()
	a.log.Info("App stopped gracefully")
}
