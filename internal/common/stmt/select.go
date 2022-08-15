package stmt

import (
	"fmt"
	"strings"
)

type SelectStmt struct {
	TableName string
	Targets   []string
	Predicate string
	ForShare  bool
	ForUpdate bool
}

func (stmt *SelectStmt) String() string {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s",
		strings.Join(stmt.Targets, ","), strings.ToUpper(stmt.TableName), stmt.Predicate)
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
