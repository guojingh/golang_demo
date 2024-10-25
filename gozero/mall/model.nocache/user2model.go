package model_nocache

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ User2Model = (*customUser2Model)(nil)

type (
	// User2Model is an interface to be customized, add more methods here,
	// and implement the added methods in customUser2Model.
	User2Model interface {
		user2Model
		withSession(session sqlx.Session) User2Model
	}

	customUser2Model struct {
		*defaultUser2Model
	}
)

// NewUser2Model returns a model.nocache for the database table.
func NewUser2Model(conn sqlx.SqlConn) User2Model {
	return &customUser2Model{
		defaultUser2Model: newUser2Model(conn),
	}
}

func (m *customUser2Model) withSession(session sqlx.Session) User2Model {
	return NewUser2Model(sqlx.NewSqlConnFromSession(session))
}
