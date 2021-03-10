package mysql

import (
	"context"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/loki-zhou/arthas/example/internal/details/internal/domain"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	_detail_info = `select * from detailtest`
)

type mysqlDetailRepositroy struct {
	Conn *sqlx.DB
	logger     *zap.Logger
}

func NewMysqlDetailRepositroy(conn *sqlx.DB, logger *zap.Logger) domain.DetailRepository {
	return &mysqlDetailRepositroy{
		Conn: conn,
		logger: logger,
	}
}

func (m *mysqlDetailRepositroy) GetByID(ctx context.Context, id int32) (*domain.Detail, error) {
	detail := domain.Detail{}
	err := m.Conn.Get(detail, _detail_info)
	if err != nil {
		return nil, errors.Wrap(err, "detail sql get error")
	}
	return &detail, nil
}

var DetailMySQLProviderSet = wire.NewSet(NewMysqlDetailRepositroy)