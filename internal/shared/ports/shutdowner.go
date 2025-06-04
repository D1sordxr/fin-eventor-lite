package ports

import "context"

type AppComponent interface {
	Shutdown(ctx context.Context) error
}
