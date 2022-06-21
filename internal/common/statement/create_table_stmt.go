package statement

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
		colDefs[idx] += colConstrDic[col.Constraint[0]] //TODO
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
	Name       string
	Type       ColumnType
	Constraint []ColumnOptions
}

type ColumnType = int

const (
	ColTypeBlob = iota
	ColTypeChar
	ColTypeDouble
	ColTypeInt
	ColTypeVarchar
)

type ColumnOptions = int

const (
	ColOptColumnFormat = iota
	ColOptComment
	ColOptNullOrNotNull
	ColOptPrimaryKey
	ColOptStorage
	ColOptUnique
)
