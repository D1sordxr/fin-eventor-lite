package api

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/balance"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
	cfg "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/grpc"
	balanceStore "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres/account/balance"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/grpc/services/account"
)

func setupGRPC(
	log ports.Log,
	config *cfg.Config,
	storage ports.Storage,
) *grpc.Server {
	dao := balanceStore.NewDAO(storage)

	useCase := balance.NewUseCase(dao)

	accountService := account.NewService(useCase)

	server := grpc.NewServer(
		log,
		config,
		accountService,
	)

	return server
}
