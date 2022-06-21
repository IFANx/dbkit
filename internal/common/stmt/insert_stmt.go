package stmt

import (
	"fmt"
	"strings"
)

type InsertStmt struct {
	TableName string
	Ignore    bool
	Columns   []string
	Values    []string
}

func (stmt *InsertStmt) String() string {
	ignore := ""
	if stmt.Ignore {
		ignore = " IGNORE"
	}
	sql := fmt.Sprintf("INSERT%s INTO %s(%s) VALUES (%s)",
		ignore, stmt.TableName, strings.Join(stmt.Columns, ","),
		strings.Join(stmt.Values, ","))
	return sql
}
