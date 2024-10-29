package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"mall/service/user/model"
	"time"

	"api/internal/svc"
	"api/internal/types"

	"github.com/golang-jwt/jwt/v4"
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
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)

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

	if passwordMd5(password) != user.Password {
		logx.Errorf("user_Login_UsermModel_Login:%s密码输入错误", user.Username)
		return &types.LoginResponse{Message: "用户名或密码错误"}, nil
	}

	// 生成 JWT
	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	secretKey := l.svcCtx.Config.Auth.AccessSecret
	token, err := l.getJwtToken(secretKey, now, expire, user.UserId)
	if err != nil {
		logx.Errorw("l.getJwtToken failed", logx.Field("err", err))
		return nil, errors.New("内部错误")
	}

	// 3.如果一致登陆成功，否则登陆失败
	return &types.LoginResponse{
		Message:      "Success",
		AccessToken:  token,
		AccessExpire: int(now + expire),
		RefreshAfter: int(now + expire/2),
	}, nil
}

// 生成 JWT 方法
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
