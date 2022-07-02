package troc

import "dbkit/internal/common/stmt"

type StatementType int

const (
	StmtTypeUnknown = iota
	StmtTypeSelect
	StmtTypeSelectShare
	StmtTypeSelectUpdate
	StmtTypeUpdate
	StmtTypeDelete
	StmtTypeInsert
	StmtTypeBegin
	StmtTypeCommit
	StmtTypeRollback
)

type StatementCell struct {
	tx       TransactionWrap
	stmtId   int
	stmt     stmt.Statement
	stmtType StatementType
	view     View
	blocked  bool
	result   []interface{}
	newRowId int
}
