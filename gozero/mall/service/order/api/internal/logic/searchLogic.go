package logic

import (
	"context"
	"errors"
	"strconv"

	"mall/service/order/api/internal/logic/interceptor"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/model"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据订单号查询订单信息+用户信息
func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// 1.根据请求参数中的订单号查询数据库找到的记录
	orderID, err := strconv.ParseUint(req.OrderID, 10, 64)
	if err != nil {
		logx.Errorw("order server Search strconv.ParseUint failed", logx.Field("err", err))
		return nil, errors.New("服务器异常")
	}

	order, err := l.svcCtx.OrderModel.FindOneByOrderId(l.ctx, orderID)
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.New("查找订单不存在")
	}

	if err != nil {
		logx.Errorw("order server FindOneByOrderId failed", logx.Field("err", err))
		return nil, errors.New("服务器异常")
	}

	// 2.根据订单记录中的 user_id 去查询用户数据（通过RPC调用user服务）
	// 假设：user_id = 162731831328769
	// 如何存入 adminID?
	l.ctx = context.WithValue(l.ctx, interceptor.CtxKeyAdminID, "33")
	userResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userclient.GetUserReq{UserID: int64(order.UserId)})
	if err != nil {
		logx.Errorw("UserRPC.GetUser failed", logx.Field("err", err))
		return nil, errors.New("内部错误")
	}

	// 3.拼接返回结果（因为我们这个接口的数据不是由我们一个服务组成）
	return &types.SearchResponse{
		OrderID:  req.OrderID,       // 实际查询的订单记录来赋值
		Status:   int(order.Status), //  实际查询的订单记录来赋值
		Username: userResp.Username, // RPC调用User服务拿到的数据
	}, nil
}
