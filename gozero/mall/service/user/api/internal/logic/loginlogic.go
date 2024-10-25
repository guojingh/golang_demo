package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"mall/service/user/model"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func passwordMd5(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	h.Write(secret)
	return hex.EncodeToString(h.Sum(nil))
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 实现登陆功能
	// 1.处理用户发过来的请求，拿到用户名和密码
	username := req.Username
	password := req.Password

	// 2.判断用户名和密码跟数据库中是不是一致
	// 两种方式：
	// 1.用 用户输入的用户名和密码去查数据库
	// 2.用用户名查到结果，再判断密码
	user, err := l.svcCtx.UsermModel.FindOneByUsername(l.ctx, username)

	if err != nil && !errors.Is(err, model.ErrNotFound) {
		logx.Errorw(
			"user_Login_UsermModel_FindOneByUsername",
			logx.Field("err", err),
		)
		return &types.LoginResponse{Message: "系统错误，请稍后重试"}, nil
	}

	if user == nil {
		return &types.LoginResponse{Message: "用户还未注册，请去注册"}, nil
	}

	// 3.如果一致登陆成功，否则登陆失败
	if passwordMd5(password) != user.Password {
		logx.Errorf("user_Login_UsermModel_Login:%s密码输入错误", user.Username)
		return &types.LoginResponse{Message: "用户名或密码错误"}, nil
	}

	return &types.LoginResponse{Message: "Success"}, nil
}
