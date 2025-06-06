package api

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/app"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/kafka"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc"
	appSrv "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http"
)

type App struct {
	Log        ports.Log
	Pool       *postgres.Pool
	HTTPServer *appSrv.Server
	GRPCServer *grpc.Server
	Shutdowner ports.Shutdowner
}

func NewApp(ctx context.Context) *App {
	cfg := config.NewConfig()

	log := slog.Default()

	pool := postgres.NewPool(ctx, &cfg.Storage)

	producer := kafka.NewProducer(&cfg.MessageBroker)

	HTTPServer := setupHTTP(
		&cfg.HTTPServer,
		log,
		pool,
		producer,
	)

	GRPCServer := setupGRPC(
		log,
		&cfg.GRPCServer,
		pool,
	)

	shutdowner := app.NewShutdowner(
		HTTPServer,
		GRPCServer,
		producer,
		pool,
	)

	return &App{
		Log:        log,
		Pool:       pool,
		HTTPServer: HTTPServer,
		GRPCServer: GRPCServer,
		Shutdowner: shutdowner,
	}
}

func (a *App) Run(ctx context.Context) {
	serversWg := &sync.WaitGroup{}
	errChan := make(chan error, 1)

	serversWg.Add(1)
	go func() {
		defer serversWg.Done()
		if err := a.GRPCServer.StartServer(); err != nil {
			errChan <- fmt.Errorf("grpc server error: %w", err)
		}
	}()

	serversWg.Add(1)
	go func() {
		defer serversWg.Done()
		if err := a.HTTPServer.StartServer(); err != nil {
			errChan <- fmt.Errorf("http server error: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		a.Log.Info("Received shutdown signal, shutting down...")
	case err := <-errChan:
		a.Log.Error("App error: " + err.Error())
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Shutdowner.ShutdownComponents(shutdownCtx); err != nil {
		a.Log.Error("Failed to shutdown components: " + err.Error())
	} else {
		a.Log.Info("All components shutdown successfully")
	}

	serversWg.Wait()
	a.Log.Info("App stopped gracefully")
}
