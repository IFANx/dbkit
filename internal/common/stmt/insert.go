package stmt

import (
	"fmt"
	"strings"
)

// InsertStmt e.g. INSERT INTO t(c2, c3) VALUES (16, 'a'), (22, 'a')
type InsertStmt struct {
	TableName string
	Columns   []string
	Values    [][]string
}

func (stmt *InsertStmt) String() string {
	var pairs []string
	for _, cells := range stmt.Values {
		pairs = append(pairs, strings.Join(cells, ","))
	}
	rows := "(" + strings.Join(pairs, "),(") + ")"
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES %s",
		stmt.TableName, strings.Join(stmt.Columns, ","), rows)
	return sql
}
