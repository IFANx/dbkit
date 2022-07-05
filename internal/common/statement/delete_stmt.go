package statement

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"strconv"
	"strings"
)

type DeleteStmt struct {
	Options    []DeleteOption
	Tables     []*common.Table
	Join       ast.AstNode
	JoinOn     ast.AstNode
	Partitions []string
	Where      ast.AstNode
	OrderBy    []*common.Column
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *DeleteStmt) String() string {
	res := "DELETE "
	if stmt.Options != nil && len(stmt.Options) > 0 {
		delOptDict := []string{"IGNORE", "LOW_PRIORITY", "QUICK"}
		optionStrList := make([]string, 0)
		for _, opt := range stmt.Options {
			optionStrList = append(optionStrList, delOptDict[opt])
		}
		res += strings.Join(optionStrList, " ")
		res += " "
	}
	res += "FROM "
	tableNameList := make([]string, 0)
	for _, tab := range stmt.Tables {
		tableNameList = append(tableNameList, tab.Name)
	}
	res += strings.Join(tableNameList, ",")
	res += " "
	// res += stmt.Tables[0].Name
	if stmt.Join != nil {
		res += "USING "
		res += stmt.Join.String()
		res += " ON " + stmt.JoinOn.String()
		res += " "
	}
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += "PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	if stmt.Where != nil {
		res += "WHERE " + stmt.Where.String()
		res += " "
	}
	if stmt.OrderBy != nil && len(stmt.OrderBy) > 0 {
		orderByList := make([]string, 0)
		for _, col := range stmt.OrderBy {
			orderByList = append(orderByList, col.Name)
		}
		res += "ORDER BY " + strings.Join(orderByList, ", ")
		res += " "
	}
	if stmt.OrderOpt > -1 {
		orderOptDict := []string{"ASC", "DESC"}
		res += orderOptDict[stmt.OrderOpt]
		res += " "
	}
	if stmt.Limit > -1 {
		res += "LIMIT " + strconv.Itoa(stmt.Limit) + " "
	}
	return res
}

type DeleteOption = int

const (
	DelOptIgnore      = iota // MySQL ignores ignorable errors during the process of deleting rows.
	DelOptLowPriority        // The storage engine does not merge index leaves during delete for MyISAM tables.
	DelOptQuick              // The server delays DELETE execution until no other clients are reading from the table. Affect using only table-level locking engines(MyISAM, MEMORY, and MERGE)
)
