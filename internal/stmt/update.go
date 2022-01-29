package stmt

import (
	"fmt"
	"strings"
)

// UpdateStmt e.g. UPDATE t SET c1=2, c2=3 WHERE c3 > 16
type UpdateStmt struct {
	Table       string
	ColValPairs map[string]string
	Predicate   string
}

func (stmt *UpdateStmt) String() string {
	var pairs []string
	for _, col := range stmt.ColValPairs {
		pairs = append(pairs, col+"="+stmt.ColValPairs[col])
	}
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		stmt.Table, strings.Join(pairs, ","), stmt.Predicate)
	return sql
}
