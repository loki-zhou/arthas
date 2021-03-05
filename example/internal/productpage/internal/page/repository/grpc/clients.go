package grpc

import (
	"github.com/google/wire"
)


var DetailRpcProviderSet = wire.NewSet(NewgrpcDetailRepository)