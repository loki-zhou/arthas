package database

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"

)

type Options struct {
	DSN   string
	Debug bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.DSN))

	return o, err
}

func New(o *Options) ( *sqlx.DB, error) {
	db, err := sqlx.Open("mysql", o.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx open database connection error")
	}
	return db, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
