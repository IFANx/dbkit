package stmt

import (
	"dbkit/internal/common/isolation"
)

type BeginStmt struct{}

func (stmt *BeginStmt) String() string {
	return "BEGIN"
}

type RollbackStmt struct{}

func (stmt *RollbackStmt) String() string {
	return "ROLLBACK"
}

type SetIsolationStmt struct {
	Isolation isolation.IsolationLevel
}

func (stmt *SetIsolationStmt) String() string {
	return "SET TRANSACTION ISOLATION LEVEL " + stmt.Isolation.Name
}
