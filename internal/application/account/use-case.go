package account

import (
	"context"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/dto"
	"github.com/D1sordxr/fin-eventor-lite/internal/application/account/ports"
)

type UseCase struct {
	svc      ports.Svc
	repo     ports.Repository
	es       ports.EventStore
	producer ports.MsgProducer
}

func NewUseCase(
	svc ports.Svc,
	repo ports.Repository,
	es ports.EventStore,
	producer ports.MsgProducer,
) *UseCase {
	return &UseCase{
		svc:      svc,
		repo:     repo,
		es:       es,
		producer: producer,
	}
}

func (uc *UseCase) Create(ctx context.Context, dto dto.DTO) (string, error) {
	entity, err := uc.svc.CreateEntity(dto.UserID)
	if err != nil {
		return "", err
	}

	if err = uc.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}

func (uc *UseCase) Deposit(ctx context.Context, dto dto.DTO) error {
	event, err := uc.svc.CreateDepositEvent(dto.ID, dto.Balance)
	if err != nil {
		return err
	}

	if err = uc.es.Save(ctx, event); err != nil {
		return err
	}

	payload, err := uc.svc.PayloadEvent(event)
	if err != nil {
		return err
	}

	if err = uc.producer.Publish(ctx, payload); err != nil {
		return err
	}

	return nil
}
