package svc

import (
	"github.com/guojinghu/shortenter/internal/config"
	"github.com/guojinghu/shortenter/model"
	"github.com/guojinghu/shortenter/sequence"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map

	// Sequence *sequence.MySQL
	Sequence sequence.Sequence
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)

	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
		//Sequence:      sequence.NewMySQL(c.Sequence.DSN), // sequence
		Sequence: sequence.NewRedis(c.Redis),
	}
}
