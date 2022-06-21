package stmt

import (
	"fmt"
	"strings"
)

type CreateTableStmt struct {
	TableName       string
	Columns         []string
	ColumnTypes     map[string]string
	ColumnOptions   map[string]string
	TableOptions    map[string]string
	PartitionOption string
}

func (stmt *CreateTableStmt) String() string {
	colDefs := make([]string, len(stmt.Columns))
	for idx, colName := range stmt.Columns {
		colDefs[idx] = colName + " " + stmt.ColumnTypes[colName]
		if stmt.ColumnOptions[colName] != "" {
			colDefs[idx] += " " + stmt.ColumnOptions[colName]
		}
	}
	tabOps := make([]string, 0)
	for key, val := range stmt.TableOptions {
		tabOps = append(tabOps, key+"="+val)
	}
	sql := fmt.Sprintf("CREATE TABLE %s(%s) %s %s", stmt.TableName,
		strings.Join(colDefs, ", "), strings.Join(tabOps, ", "), stmt.PartitionOption)
	return sql
}
