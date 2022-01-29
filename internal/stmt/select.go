package stmt

import (
	"fmt"
	"strings"
)

// SelectStmt e.g. SELECT c1+6,c2 FROM t WHERE c3 > 16 FOR UPDATE
type SelectStmt struct {
	Table     string
	Columns   []string
	Predicate string
	ForShare  bool
	ForUpdate bool
}

func (stmt *SelectStmt) String() string {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s",
		strings.Join(stmt.Columns, ","), stmt.Table, stmt.Predicate)
	if stmt.ForShare {
		sql += " FOR SHARE"
	}
	if stmt.ForUpdate {
		sql += " FOR UPDATE"
	}
	return sql
}

func (stmt *SelectStmt) StringInMode() string {
	return strings.Replace(stmt.String(), "FOR SHARE", "IN SHARE MODE", -1)
}
