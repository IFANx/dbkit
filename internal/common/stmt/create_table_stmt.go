package stmt

import (
	"fmt"
	"strconv"
	"strings"
)

type CreateTableStmt struct {
	TableName       string
	Columns         []Column
	TableOptions    map[string]string
	PartitionOption string
}

func (stmt *CreateTableStmt) String() string {
	colDefs := make([]string, len(stmt.Columns))
	for idx, col := range stmt.Columns {
		colDefs[idx] = col.Name + " " + strconv.Itoa(col.Type)
		colConstrDic := []string{" UNIQUE ", " NOT NULL ", " PRIMARY KEY "}
		for _, colConstr := range col.Constraint {
			colDefs[idx] += colConstrDic[colConstr]
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

type Column struct {
	Table      *CreateTableStmt
	Name       string
	Type       ColumnType
	Constraint []ColumnConstraint
	Length     int
	ValueCache []string
}

type ColumnType = int

const (
	INT = iota
	DOUBLE
	CHAR
	VARCHAR
	BLOB
)

type ColumnConstraint = int

const (
	NotNull = iota
	Unique
	Primary
)
