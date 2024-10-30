package sequence

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 建立 MySQL 连接执行 REPLACE INTO  语句
// REPLACE INTO sequence (stub) VALUES ('a');
// SELECT LAST_INSERT_ID();

const sqlReplaceStub = `REPLACE INTO sequence (stub) VALUES ('a')`

type MySQL struct {
	conn sqlx.SqlConn
}

func NewMySQL(dns string) Sequence {
	conn := sqlx.NewMysql(dns)
	return &MySQL{
		conn: conn,
	}
}

// Next 取下一个号
func (m *MySQL) Next() (seq uint64, err error) {

	// prepare
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	defer stmt.Close()

	// 执行
	var rest sql.Result
	rest, err = stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec() failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	// 获取获取插入的id
	var lid int64
	lid, err = rest.LastInsertId()
	if err != nil {
		logx.Errorw("rest.LastInsertId() failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	return uint64(lid), nil
}
