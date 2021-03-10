package internal

import (
	"github.com/google/wire"
	_detailGRPCDelivery "github.com/loki-zhou/arthas/example/internal/details/internal/detail/delivery/grpc"
	_detailMySQLRepo "github.com/loki-zhou/arthas/example/internal/details/internal/detail/repository/mysql"
	_detailUsecase "github.com/loki-zhou/arthas/example/internal/details/internal/detail/usecase"

	"github.com/loki-zhou/arthas/grpc"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/loki-zhou/arthas/app"
	"github.com/loki-zhou/arthas/http"
)

type Options struct {
	Name string

}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, hs *http.Server, gs *grpc.Server, detailfn _detailGRPCDelivery.DetailGRPCDeliveryFn) (*app.Application, error) {
	a, err := app.New(o.Name, logger, app.GrpcServerOption(gs), app.HttpServerOption(hs) )
	gs.AddServers(grpc.InitServers(detailfn))
	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions,
	_detailGRPCDelivery.ProviderSet,
	_detailMySQLRepo.DetailMySQLProviderSet,
	_detailUsecase.DetailUsecaseProviderSet,
)



