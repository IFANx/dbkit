package common

import (
	"dbkit/internal/common/stmt"
	"github.com/jmoiron/sqlx"
)

type SqlExecutor interface {
	Queryx(query string) (*sqlx.Rows, error)
	QuerySQL(query string) ([][]interface{}, error)
	Query(stmt stmt.SelectStmt) ([][]interface{}, error)
	ExecSQLIgnoreError(sql string)
	ExecSQL(sql string) error
	ExecSQLAffectedRow(sql string) (int, error)
	ExecUpdate(stmt stmt.UpdateStmt) (int, error)
	ExecDelete(stmt stmt.DeleteStmt) (int, error)
	ExecInsert(stmt stmt.InsertStmt) error
}
