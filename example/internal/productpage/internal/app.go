package internal

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/loki-zhou/arthas/app"
	"github.com/loki-zhou/arthas/http"

	_pageHTTPDelivery "github.com/loki-zhou/arthas/example/internal/productpage/internal/page/delivery/http"
	_pageUsecase "github.com/loki-zhou/arthas/example/internal/productpage/internal/page/usecase"
	_pageGRPCRepo "github.com/loki-zhou/arthas/example/internal/productpage/internal/page/repository/grpc"

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

func NewApp(o *Options, logger *zap.Logger, hs *http.Server, detailfn _pageHTTPDelivery.DetailHTTPDeliveryFn) (*app.Application, error) {
	a, err := app.New(o.Name, logger, app.HttpServerOption(hs))
	hs.AddRoute(http.InitControllers(detailfn))
	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions,
	_pageHTTPDelivery.ProviderSet,
	_pageUsecase.DetailUsecaseProviderSet,
	_pageGRPCRepo.DetailRpcProviderSet,
)

