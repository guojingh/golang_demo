package svc

import (
	"mall/service/order/api/internal/config"
	"mall/service/order/api/internal/logic/interceptor"
	"mall/service/order/model"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrdersModel

	UserRPC userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrdersModel(sqlxConn, c.CacheRedis),
		// 初始化user服务的RPC客户端
		UserRPC: userclient.NewUser(
			zrpc.MustNewClient(c.UserRPC,
				zrpc.WithUnaryClientInterceptor(interceptor.UnaryInterceptor),
			),
		),
	}
}
