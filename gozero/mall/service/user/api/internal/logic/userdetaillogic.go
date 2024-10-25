package logic

import (
	"context"
	"errors"
	"mall/service/user/model"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {

	user, err := l.svcCtx.UsermModel.FindOneByUserId(l.ctx, int64(req.UserID))

	if err != nil && errors.Is(err, model.ErrNotFound) {
		logx.Errorw(
			"user_detail_UserDetail_FindOneByUserId",
			logx.Field("err", err),
		)
		return nil, nil
	}

	return &types.UserDetailResponse{
		Username: user.Username,
		Email:    user.Email.String,
		Gender:   int(user.Gender),
	}, nil
}
