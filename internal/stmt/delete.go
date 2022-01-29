package stmt

import (
	"fmt"
)

// DeleteStmt e.g. DELETE FROM t WHERE c3 > 16
type DeleteStmt struct {
	Table     string
	Predicate string
}

func (stmt *DeleteStmt) String() string {
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s",
		stmt.Table, stmt.Predicate)
	return sql
}
