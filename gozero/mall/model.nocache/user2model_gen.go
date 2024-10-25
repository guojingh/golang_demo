// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model_nocache

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	user2FieldNames          = builder.RawFieldNames(&User2{})
	user2Rows                = strings.Join(user2FieldNames, ",")
	user2RowsExpectAutoSet   = strings.Join(stringx.Remove(user2FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	user2RowsWithPlaceHolder = strings.Join(stringx.Remove(user2FieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	user2Model interface {
		Insert(ctx context.Context, data *User2) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User2, error)
		FindOneByUserId(ctx context.Context, userId int64) (*User2, error)
		FindOneByUsername(ctx context.Context, username string) (*User2, error)
		Update(ctx context.Context, data *User2) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUser2Model struct {
		conn  sqlx.SqlConn
		table string
	}

	User2 struct {
		Id         int64          `db:"id"`
		UserId     int64          `db:"user_id"`
		Username   string         `db:"username"`
		Password   string         `db:"password"`
		Email      sql.NullString `db:"email"`
		Gender     int64          `db:"gender"`
		CreateTime time.Time      `db:"create_time"`
		UpdateTime time.Time      `db:"update_time"`
	}
)

func newUser2Model(conn sqlx.SqlConn) *defaultUser2Model {
	return &defaultUser2Model{
		conn:  conn,
		table: "`user2`",
	}
}

func (m *defaultUser2Model) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUser2Model) FindOne(ctx context.Context, id int64) (*User2, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", user2Rows, m.table)
	var resp User2
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUser2Model) FindOneByUserId(ctx context.Context, userId int64) (*User2, error) {
	var resp User2
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", user2Rows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUser2Model) FindOneByUsername(ctx context.Context, username string) (*User2, error) {
	var resp User2
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", user2Rows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUser2Model) Insert(ctx context.Context, data *User2) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, user2RowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Username, data.Password, data.Email, data.Gender)
	return ret, err
}

func (m *defaultUser2Model) Update(ctx context.Context, newData *User2) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, user2RowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.UserId, newData.Username, newData.Password, newData.Email, newData.Gender, newData.Id)
	return err
}

func (m *defaultUser2Model) tableName() string {
	return m.table
}
