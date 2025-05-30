package account

import (
	"context"
)

type UseCase struct {
	svc      svc
	repo     repository
	es       eventStore
	producer msgProducer
}

func NewUseCase(
	svc svc,
) *UseCase {
	return &UseCase{
		svc: svc,
	}
}

func (uc *UseCase) Create(ctx context.Context, dto DTO) (string, error) {
	entity, err := uc.svc.CreateEntity(dto.UserID)
	if err != nil {
		return "", err
	}

	if err = uc.repo.Save(ctx, entity); err != nil {
		return "", err
	}

	return entity.ID.String(), nil
}

func (uc *UseCase) Deposit(ctx context.Context, dto DTO) error {
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
