package ast

import (
	"dbkit/internal"
	"strconv"
	"strings"
)

type UpdateStmt struct {
	Options    []UpdateOption
	Table      internal.Table
	Partitions []string
	UpdateCol  []*internal.Column
	UpdateExpr []AstNode // 结构待调整
	Where      AstNode
	OrderBy    []*internal.Column
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *UpdateStmt) String() string {
	res := "UPDATE "
	selOptDict := []string{"LOW_PRIORITY", "IGNORE"}
	optionStrList := make([]string, 0)
	for _, opt := range stmt.Options {
		optionStrList = append(optionStrList, selOptDict[opt])
	}
	res += strings.Join(optionStrList, " ")
	res += stmt.Table.Name
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += "PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	res += " SET "
	for i, expr := range stmt.UpdateExpr {
		if i != 0 {
			res += ", "
		}
		res += stmt.UpdateCol[i].Name + " = "
		res += expr.String()
	}
	if stmt.Where != nil {
		res += " WHERE " + stmt.Where.String() + " "
	}
	if stmt.OrderBy != nil {
		orderByList := make([]string, 0)
		for _, col := range stmt.OrderBy {
			optionStrList = append(orderByList, col.Name)
		}
		res += "ORDER BY " + strings.Join(optionStrList, " ")
	}
	if stmt.OrderOpt > -1 {
		orderOptDict := []string{"ASC", "DESC"}
		res += orderOptDict[stmt.OrderOpt] + " "
	}
	if stmt.Limit > -1 {
		res += "LIMIT " + strconv.Itoa(stmt.Limit) + " "
	}
	return res
}

type UpdateOption = int

const (
	UptOptLowPriority = iota
	UptOptIgnore
)
