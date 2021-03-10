package grpc

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	pb "github.com/loki-zhou/arthas/example/internal/details/gen"
	"github.com/loki-zhou/arthas/example/internal/details/internal/domain"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	stdgrpc "google.golang.org/grpc"
)

type DetailHandler struct {
	DetailUsecase domain.DetailUsecase
	logger     *zap.Logger
	pb.UnimplementedDetailServiceServer
}

func NewDetailHandler(logger *zap.Logger,detailUsecase domain.DetailUsecase) *DetailHandler{
	handler := &DetailHandler{
		DetailUsecase: detailUsecase,
		logger: logger,
	}
	return handler
}

type DetailGRPCDeliveryFn func(s *stdgrpc.Server)

func NewDetailGRPCDeliveryFn(handler *DetailHandler) DetailGRPCDeliveryFn {
	return func(s *stdgrpc.Server) {
		pb.RegisterDetailServiceServer(s,handler)
	}
}

var ProviderSet = wire.NewSet(NewDetailHandler, NewDetailGRPCDeliveryFn)


func (d *DetailHandler) GetDetail(ctx context.Context, req *pb.GetDetailRequest) (*pb.GetDetailResponse, error) {
	detail, err := d.DetailUsecase.GetByID(ctx, req.GetId())
	if err != nil {
		d.logger.Error("get detail id error", zap.Error(err))
		return nil, nil
	}
	resp := pb.GetDetailResponse{}
	if err := copier.Copy(&resp, &detail); err != nil {
		return nil, errors.Wrap(err, "detail grpc copy error")
	}
	return &resp, nil
}