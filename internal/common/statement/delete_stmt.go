package statement

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"strconv"
	"strings"
)

type DeleteStmt struct {
	Option     DeleteOption
	Table      common.Table
	Partitions []string
	Where      ast.AstNode
	OrderBy    []*common.Column
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *DeleteStmt) String() string {
	res := "DELETE "
	if stmt.Option == 1 {
		res += "IGNORE "
	}
	res += " FROM "
	res += stmt.Table.Name
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += " PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	if stmt.Where != nil {
		res += "WHERE " + stmt.Where.String() + " "
	}
	if stmt.OrderBy != nil {
		orderByList := make([]string, 0)
		for _, col := range stmt.OrderBy {
			orderByList = append(orderByList, col.Name)
		}
		res += "ORDER BY " + strings.Join(orderByList, " ")
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
	DelOptIgnore = 1
)
