package statement

import (
	"dbkit/internal"
	"dbkit/internal/common/ast"
	"strconv"
	"strings"
)

type UpdateStmt struct {
	Option     UpdateOption
	Table      internal.Table
	Partitions []string
	UpdateCol  []*internal.Column
	UpdateExpr []ast.AstNode // 结构待调整
	Where      ast.AstNode
	OrderBy    []*internal.Column
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *UpdateStmt) String() string {
	res := "UPDATE "
	if stmt.Option == 1 {
		res += "IGNORE "
	}
	res += stmt.Table.Name
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += " PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
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

type UpdateOption = int

const (
	UpdOptIgnore = 1
)
