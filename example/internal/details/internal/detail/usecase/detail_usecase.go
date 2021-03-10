package usecase

import (
	"context"
	"github.com/google/wire"
	"github.com/loki-zhou/arthas/example/internal/details/internal/domain"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type  detailUsecase struct {
	detailRepo domain.DetailRepository
	logger     *zap.Logger
}


func (du *detailUsecase) GetByID(ctx context.Context, id int32) (*domain.Detail, error) {
	detail, err := du.detailRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "detail grpc service error")
	}
	return detail, nil
}

func NewDetailUsecase(detailRepo domain.DetailRepository, logger *zap.Logger) domain.DetailUsecase {
	return &detailUsecase{
		detailRepo: detailRepo,
		logger: logger,
		}
}

var DetailUsecaseProviderSet = wire.NewSet(NewDetailUsecase)