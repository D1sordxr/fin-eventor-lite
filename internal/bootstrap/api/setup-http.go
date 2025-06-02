package api

import (
	accountUC "github.com/D1sordxr/fin-eventor-lite/internal/application/account"
	userUC "github.com/D1sordxr/fin-eventor-lite/internal/application/user"
	accountDomainSvc "github.com/D1sordxr/fin-eventor-lite/internal/domain/account/services"
	userDomain "github.com/D1sordxr/fin-eventor-lite/internal/domain/user"
	midSetup "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/http/middleware"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/kafka/mocks"
	accountMocks "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres/account/mocks"
	userStore "github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/postgres/user"
	httpSrv "github.com/D1sordxr/fin-eventor-lite/internal/presentation/http"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/delivery/account"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/delivery/user"
	"github.com/D1sordxr/fin-eventor-lite/internal/presentation/http/middleware"
	"github.com/D1sordxr/fin-eventor-lite/internal/shared/ports"
	"net/http"
)

func setupHTTP(
	log ports.Log,
	port string,
) *httpSrv.Server {

	accountUseCase := accountUC.NewUseCase(
		new(accountDomainSvc.Svc),
		accountMocks.NewMockRepo(), // Mock
		new(mocks.Producer),        // Mock
	)
	userUseCase := userUC.NewUseCase(
		new(userDomain.Svc),
		userStore.NewMockRepo(), // Mock
	)

	chainer := new(midSetup.ChainerImpl)

	logMid := middleware.NewLogMid(log)
	methodPostMid := middleware.NewMethodMid(http.MethodPost)
	semaphoreMid := middleware.NewSemaphoreMid()
	retryMid := new(middleware.RetryMid)

	userHandler := user.NewHandler(
		userUseCase,
		chainer,
		logMid.Log,
		methodPostMid.OnlyPost,
		semaphoreMid.Limit,
		retryMid.RetryWithBackoff,
	)
	accountHandler := account.NewHandler(
		accountUseCase,
		chainer,
		logMid.Log,
		methodPostMid.OnlyPost,
		semaphoreMid.Limit,
		retryMid.RetryWithBackoff,
	)

	return httpSrv.NewServer(
		port,
		userHandler,
		accountHandler,
	)
}
