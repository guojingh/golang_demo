package logic

import (
	"context"
	"errors"
	"fmt"

	"mall/service/user/model"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line
	// 更具 userID 查询数据库返回用户信息
	one, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserID)

	// 1.查询数据库失败
	// 2.根据UserID查不到用户
	if errors.Is(err, model.ErrNotFound) {
		fmt.Printf("无效的UserID:%v\n", in.UserID)
		return nil, errors.New("无效的UserID")
	}

	if err != nil {
		logx.Errorw("FindOneByUserId failed", logx.Field("err", err))
		return nil, errors.New("查询失败")
	}

	// 返回响应
	return &user.GetUserResp{
		UserID:   one.UserId,
		Username: one.Username,
		Gender:   one.Gender,
	}, nil
}
