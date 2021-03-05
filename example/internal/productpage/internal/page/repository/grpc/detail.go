package grpc

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/loki-zhou/arthas/grpc"
	pb "github.com/loki-zhou/arthas/example/internal/productpage/gen"
	"github.com/loki-zhou/arthas/example/internal/productpage/internal/domain"
)

type grpcDetailRepository struct {
	detailGRPC pb.DetailServiceClient
	logger     *zap.Logger
}

func (g *grpcDetailRepository)GetByID(ctx context.Context, id int32) (*domain.Detail, error) {
	response, err := g.detailGRPC.GetDetail(ctx, &pb.GetDetailRequest{Id:id})
	if err != nil {
		return nil, errors.Wrap(err, "detail grpc service error")
	}
	detail := domain.Detail{}
	if err := copier.Copy(&detail, response); err != nil {
		return nil, errors.Wrap(err, "detail grpc copy error")
	}
	return &detail, nil
}


func NewgrpcDetailRepository(client *grpc.Client, logger *zap.Logger) (domain.DetailRepository, error) {
	conn, err := client.Dial("Details")
	if err != nil {
		return nil,errors.Wrap(err,"detail client dial error")
	}
	return &grpcDetailRepository{
		detailGRPC: pb.NewDetailServiceClient(conn),
		logger: logger,
	}, nil
}



