package statement

import (
	"dbkit/internal/common"
	"dbkit/internal/common/ast"
	"strconv"
	"strings"
)

type UpdateStmt struct {
	Options    []UpdateOption
	Tables     []*common.Table
	Join       ast.AstNode
	JoinOn     ast.AstNode
	Partitions []string
	UpdateCol  []ast.ColRefNode
	UpdateExpr []string // 结构待调整
	Where      ast.AstNode
	OrderBy    []ast.ColRefNode
	OrderOpt   OrderOption
	Limit      int
}

func (stmt *UpdateStmt) String() string {
	res := "UPDATE "
	if stmt.Options != nil && len(stmt.Options) > 0 {
		updOptDict := []string{"IGNORE", "LOW_PRIORITY"}
		optionStrList := make([]string, 0)
		for _, opt := range stmt.Options {
			optionStrList = append(optionStrList, updOptDict[opt])
		}
		res += strings.Join(optionStrList, " ")
		res += " "
	}
	if stmt.Join == nil {
		tableNameList := make([]string, 0)
		for _, tab := range stmt.Tables {
			tableNameList = append(tableNameList, tab.Name)
		}
		res += strings.Join(tableNameList, ",")
		res += " "
		// res += stmt.Tables[0].Name
	} else {
		res += stmt.Join.String()
		res += " ON " + stmt.JoinOn.String()
		res += " "
	}
	if stmt.Partitions != nil && len(stmt.Partitions) > 0 {
		res += "PARTITION(" + strings.Join(stmt.Partitions, ",") + ") "
	}
	res += "SET "
	for i, expr := range stmt.UpdateExpr {
		if i != 0 {
			res += ", "
		}
		res += stmt.UpdateCol[i].String() + " = "
		res += expr
	}
	res += " "
	if stmt.Where != nil {
		res += "WHERE " + stmt.Where.String() + " "
	}
	if stmt.OrderBy != nil && len(stmt.OrderBy) > 0 {
		orderByList := make([]string, 0)
		for _, col := range stmt.OrderBy {
			orderByList = append(orderByList, col.String())
		}
		res += "ORDER BY " + strings.Join(orderByList, ", ")
		res += " "
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
	UpdOptIgnore = iota
	UpdOptLowPriority
)
