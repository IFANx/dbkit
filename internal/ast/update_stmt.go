package ast

import (
	"dbkit/internal"
	"strconv"
	"strings"
)

type UpdateStmt struct {
	Options    []UpdateOption
	Table      internal.Table
	UpdateExpr []AstNode
	Where      AstNode
	OrderBy    AstNode
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *UpdateStmt) String() string {
	res := "UPDATE "
	selOptDict := []string{"LOW_PRIORITY", "IGNORE"}
	optionStrList := make([]string, 0)
	for opt := range stmt.Options {
		optionStrList = append(optionStrList, selOptDict[opt])
	}
	res += strings.Join(optionStrList, " ")
	res += stmt.Table.Name
	res += " SET "
	if stmt.Where != nil {
		res += "WHERE " + stmt.Where.String() + " "
	}
	if stmt.OrderBy != nil {
		res += "ORDER BY " + stmt.OrderBy.String() + " "
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
