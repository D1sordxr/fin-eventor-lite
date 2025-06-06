package api

import (
	"net/http"

	accountUC "github.com/D1sordxr/fin-eventor-lite/internal/application/account"
	userUC "github.com/D1sordxr/fin-eventor-lite/internal/application/user"
	accountDomainSvc "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/account/services"
	userDomain "github.com/D1sordxr/fin-eventor-lite/internal/domain/core/user"
	"github.com/D1sordxr/fin-eventor-lite/internal/domain/ports"
	cfg "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/http"
	midSetup "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/http/middleware"
	accountStore "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres/account"
	userStore "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres/user"
	httpSrv "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/delivery/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/delivery/user"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/middleware"
)

func setupHTTP(
	config *cfg.Config,
	log ports.Log,
	storage ports.Storage,
	msgProducer ports.Producer,
) *httpSrv.Server {

	accountRepo := accountStore.NewRepository(storage, new(accountStore.Converter))
	userRepo := userStore.NewRepository(storage, new(userStore.Converter))

	userUseCase := userUC.NewUseCase(
		new(userDomain.Svc),
		userRepo,
	)
	accountUseCase := accountUC.NewUseCase(
		new(accountDomainSvc.Svc),
		accountRepo,
		msgProducer,
	)

	chainer := new(midSetup.ChainerImpl)

	traceMid := new(middleware.TracingMid)
	logMid := middleware.NewLogMid(log)
	methodPostMid := middleware.NewMethodMid(http.MethodPost)
	semaphoreMid := middleware.NewSemaphoreMid()
	retryMid := new(middleware.RetryMid)

	userHandler := user.NewHandler(
		userUseCase,
		chainer,
		traceMid.Trace,
		logMid.Log,
		methodPostMid.OnlyPost,
		semaphoreMid.Limit,
		retryMid.RetryWithBackoff,
	)
	accountHandler := account.NewHandler(
		accountUseCase,
		chainer,
		traceMid.Trace,
		logMid.Log,
		methodPostMid.OnlyPost,
		semaphoreMid.Limit,
		retryMid.RetryWithBackoff,
	)

	return httpSrv.NewServer(
		log,
		config,
		userHandler,
		accountHandler,
	)
}
