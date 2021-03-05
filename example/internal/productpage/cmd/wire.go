// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/loki-zhou/arthas/app"
	"github.com/loki-zhou/arthas/config"
	"github.com/loki-zhou/arthas/discovery/consul"
	"github.com/loki-zhou/arthas/http"
	"github.com/loki-zhou/arthas/grpc"
	"github.com/loki-zhou/arthas/log"
	"github.com/loki-zhou/arthas/trace/jaeger"

	"github.com/loki-zhou/arthas/example/internal/productpage/internal"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	grpc.ProviderSet,
	http.ProviderSet,
	jaeger.ProviderSet,
	consul.ProviderSet,
	internal.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}


