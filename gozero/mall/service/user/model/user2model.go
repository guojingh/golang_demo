package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ User2Model = (*customUser2Model)(nil)

type (
	// User2Model is an interface to be customized, add more methods here,
	// and implement the added methods in customUser2Model.
	User2Model interface {
		user2Model
	}

	customUser2Model struct {
		*defaultUser2Model
	}
)

// NewUser2Model returns a model for the database table.
func NewUser2Model(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) User2Model {
	return &customUser2Model{
		defaultUser2Model: newUser2Model(conn, c, opts...),
	}
}
