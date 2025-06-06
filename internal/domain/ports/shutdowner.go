package ports

import "context"

type Shutdowner interface {
	ShutdownComponents(ctx context.Context) error
}

type AppComponent interface {
	Shutdown(ctx context.Context) error
}
