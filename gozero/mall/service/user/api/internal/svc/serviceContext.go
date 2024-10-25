package svc

import (
	"api/internal/config"
	"mall/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UsermModel model.User2Model // ??User???????model.nocache
}

func NewServiceContext(c config.Config) *ServiceContext {
	// UserModel -> ????
	// *defaultUserModel ?????
	// ???????? *model.nocache.defaultUserModel
	// NewUser2Model(conn sqlx.SqlConn)
	// ?? sqlx.SqlConn ??????

	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		UsermModel: model.NewUser2Model(sqlxConn, c.CacheRedis),
	}
}
