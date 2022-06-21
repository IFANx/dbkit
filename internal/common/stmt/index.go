package stmt

import (
	"fmt"
	"strings"
)

// CreateIndexStmt e.g. CREATE UNIQUE INDEX i1 ON t0(c1, c2)
type CreateIndexStmt struct {
	IndexName string
	TableName string
	Columns   []string
	Unique    bool
}

func (stmt *CreateIndexStmt) String() string {
	uniqueStr := " "
	if stmt.Unique {
		uniqueStr = " UNIQUE "
	}
	sql := fmt.Sprintf("CREATE%sINDEX %s ON %s(%s)", uniqueStr,
		stmt.IndexName, stmt.TableName, strings.Join(stmt.Columns, ","))
	return sql
}
