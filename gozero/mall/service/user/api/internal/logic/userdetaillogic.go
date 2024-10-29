package logic

import (
	"context"
	"errors"
	"fmt"
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
	// jwt 鉴权后，解析出来的数据
	fmt.Printf("JWT userId:%v", l.ctx.Value("userId"))
	// 遇事不决先写注释
	// 1.拿到请求参数
	// 2.根据用户id查询数据库
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, int64(req.UserID))

	if err != nil {
		logx.Errorw(
			"user_detail_UserDetail_FindOneByUserId",
			logx.Field("err", err),
		)
		if !errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("系统错误请稍后重试")
		}

		return nil, errors.New("用户不存在")
	}

	// 3.格式化数据（数据库里的字段可能和前端要求的字段格式不一致）
	// 4.返回响应
	return &types.UserDetailResponse{
		Username: user.Username,
		Email:    user.Email.String,
		Gender:   int(user.Gender),
	}, nil
}
