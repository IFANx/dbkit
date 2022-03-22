package stmt

import (
	"fmt"
)

// DeleteStmt e.g. DELETE FROM t WHERE c3 > 16
type DeleteStmt struct {
	TableName string
	Predicate string
}

func (stmt *DeleteStmt) String() string {
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s",
		stmt.TableName, stmt.Predicate)
	return sql
}
