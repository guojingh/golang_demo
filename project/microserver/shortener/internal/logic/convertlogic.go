package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/guojinghu/shortenter/internal/svc"
	"github.com/guojinghu/shortenter/internal/types"
	"github.com/guojinghu/shortenter/model"
	"github.com/guojinghu/shortenter/pkg/base62"
	"github.com/guojinghu/shortenter/pkg/connect"
	"github.com/guojinghu/shortenter/pkg/md5"
	"github.com/guojinghu/shortenter/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 转链：输入一个长链接 --> 转为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1.校验输入的数据
	// 1.1 数据不能为空
	// if len(req.LongUrl) == 0 {}
	// 使用 validator 包用来做参数校验

	// 1.2 输入的长链接能请求通的网址
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效的链接")
	}

	// 1.3 判断之前是否已经转链过（数据库中是否已经存在该长链接）
	// 1.3.1 给长链接生成MD5值
	md5Value := md5.Sum([]byte(req.LongUrl)) // 这里使用的是项目中封装的 pkg/md5 包
	// 1.3.2 拿MD5值取数据库中查是否存在
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if !errors.Is(err, model.ErrNotFound) {
		if err == nil {
			return nil, fmt.Errorf("该链接已被转为%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	// 1.4 避免循环转链（输入的不能是一个短链接）
	// 输入的是一个完整的url q1mi.cn/1d12a?name=q1mi
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("url.Parse failed", logx.LogField{Key: "lurl", Value: req.LongUrl}, logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if !errors.Is(err, model.ErrNotFound) {
		if err != nil {
			return nil, errors.New("该链接已经是断链了")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	var short string
	for {
		// 2.取号
		// 每来一个转链请求，我们就使用 REPLACE INTO语句往 sequence 表插入一条语句，并且取出主键id作为号码
		seq, err := l.svcCtx.Sequence.Next()
		if err != nil {
			logx.Errorw("Sequence.Next() failed", logx.LogField{Key: "err", Value: err.Error()})
			return nil, err
		}

		fmt.Println(seq)

		// 3.号码转短链
		// 3.1 安全性
		short = base62.Int2String(seq)
		// 3.2 短域名黑名单避免某些特殊词如：api health fuck...
		if _, ok := l.svcCtx.ShortUrlBlackList[short]; !ok {
			break // 生成不在黑名单里的短链接就跳出for循环
		}
	}

	fmt.Printf("short:%v\n", short)

	// 4.存储长链接短链接的映射
	if _, err := l.svcCtx.ShortUrlModel.Insert(
		l.ctx,
		&model.ShortUrlMap{
			Lurl: sql.NullString{String: req.LongUrl, Valid: true},
			Md5:  sql.NullString{String: md5Value, Valid: true},
			Surl: sql.NullString{String: short, Valid: true},
		},
	); err != nil {
		logx.Errorw("ShortUrlModel.Insert", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	// 4.2 将生成的短链接加到布隆过滤器中
	if err := l.svcCtx.Filter.Add([]byte(short)); err != nil {
		logx.Errorw("BloomFilter.Add() failed", logx.LogField{Key: "err", Value: err.Error()})
	}

	// 5.返回响应
	// 5.1 返回的是短域名+短链接 qimi.cn/1En
	shortUrl := l.svcCtx.Config.ShortDoamin + "/" + short

	return &types.ConvertResponse{ShortUrl: shortUrl}, nil
}
