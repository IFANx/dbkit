package ast

import (
	"dbkit/internal"
	"strconv"
	"strings"
)

type DeleteStmt struct {
	Options    []DeleteOption
	Table      internal.Table
	Partitions []string
	Where      AstNode
	OrderBy    []*internal.Column
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *DeleteStmt) String() string {
	res := "DELETE "
	delOptDict := []string{"LOW_PRIORITY", "IGNORE"}
	optionStrList := make([]string, 0)
	for _, opt := range stmt.Options {
		optionStrList = append(optionStrList, delOptDict[opt])
	}
	res += strings.Join(optionStrList, " ")
	res += " FROM "
	res += stmt.Table.Name
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += "PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	if stmt.Where != nil {
		res += "WHERE " + stmt.Where.String() + " "
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

type DeleteOption = int

const (
	DelOptLowPriority = iota
	DelOptIgnore
)
