package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/kafka"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc"
	appSrv "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http"
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"log/slog"
	"sync"
	"time"
)

type App struct {
	Log        ports.Log
	Pool       *postgres.Pool
	HttpServer *appSrv.Server
	GrpcServer *grpc.Server
	components []ports.AppComponent
}

func NewApp(ctx context.Context) *App {
	cfg := config.NewConfig()

	log := slog.Default()

	pool := postgres.NewPool(ctx, &cfg.Storage)

	producer := kafka.NewProducer(&cfg.MessageBroker)

	httpServer := setupHTTP(
		&cfg.HttpServer,
		log,
		pool,
		producer,
	)

	grpcServer := setupGRPC(
		log,
		&cfg.GrpcServer,
		pool,
	)

	components := setupComponents(
		httpServer,
		grpcServer,
		producer,
		pool,
	)

	return &App{
		Log:        log,
		Pool:       pool,
		HttpServer: httpServer,
		GrpcServer: grpcServer,
		components: components,
	}
}

func (a *App) Run(ctx context.Context) {
	serversWg := &sync.WaitGroup{}
	errChan := make(chan error, 1)

	serversWg.Add(1)
	go func() {
		defer serversWg.Done()
		if err := a.GrpcServer.StartServer(); err != nil {
			errChan <- fmt.Errorf("grpc server error: %w", err)
		}
	}()

	serversWg.Add(1)
	go func() {
		defer serversWg.Done()
		if err := a.HttpServer.StartServer(); err != nil {
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

	if err := a.shutdownComponents(shutdownCtx); err != nil {
		a.Log.Error("Failed to shutdown components: " + err.Error())
	} else {
		a.Log.Info("All components shutdown successfully")
	}

	serversWg.Wait()
	a.Log.Info("App stopped gracefully")
}

func (a *App) shutdownComponents(ctx context.Context) error {
	var errs []error
	for _, component := range a.components {
		err := component.Shutdown(ctx)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to shutdown component %T: %w\n", component, err))
		}
	}

	return errors.Join(errs...)
}
