package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"mall/service/user/model"

	"api/internal/svc"
	"api/internal/types"

	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var secret = []byte("夏天夏天悄悄过去")

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {
	// 参数校验
	if req.Password != req.RePassword {
		return nil, errors.New("两次输入的密码不一致")
	}
	logx.Debugv(req) // json.Marshal(req)
	logx.Debugf("req:%#v", req)

	// 业务逻辑写在这里
	// 把用户的注册信息保存到数据库
	// 0.查询username是否已经被注册
	exist, err := l.svcCtx.UsermModel.FindOneByUsername(l.ctx, req.Username)
	// 0.1 查询数据库失败
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(
			"user_Signup_UsermModel.FindOneByUsername failed failed",
			logx.Field("err", err),
		)
		return nil, err
	}

	// 0.2 查到记录
	if exist != nil {
		return &types.SignupResponse{Message: "用户名已经存在"}, nil
	}

	// 1.生成userId（雪花算法）
	userID, err := GetID()
	if err != nil {
		return &types.SignupResponse{Message: "用户生成UUID失败"}, nil
	}

	// 2.加密密码（加盐 MD5）
	h := md5.New()
	h.Write([]byte(req.Password)) // 密码计算MD5
	h.Write(secret)               // 加盐
	passwordStr := hex.EncodeToString(h.Sum(nil))

	//password, _ := EncryptPassword(req.Password)
	user := &model.User2{
		UserId:   int64(userID),
		Username: req.Username,
		Password: passwordStr, //不能存明文
		Gender:   int64(req.Gender),
	}

	_, err = l.svcCtx.UsermModel.Insert(context.Background(), user)
	if err != nil {
		logx.Errorf("user_Signup_UsermModel.Insert failed, err:%v", err)
		return nil, err
	}

	return &types.SignupResponse{Message: "success"}, nil
}

var (
	sonyFlake     *sonyflake.Sonyflake // 实例
	sonyMachineID uint16               // 机器ID
)

// 生成雪花算法的一个节点
func InitSonyflakeNode(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2024-10-24") // 初始化一个开始的时间
	settings := sonyflake.Settings{                // 生成全局配置
		StartTime: t,
		MachineID: getMachineID, // 指定机器ID
	}
	sonyFlake = sonyflake.NewSonyflake(settings) // 用配置生成sonyflake节点
	return
}

func getMachineID() (uint16, error) { // 返回全局定义的机器ID
	return sonyMachineID, nil
}

// GetID 返回生成的id值
func GetID() (id uint64, err error) { // 拿到sonyFlake节点生成id值
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
