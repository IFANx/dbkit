package stmt

import "dbkit/internal/txn"

type BeginStmt struct{}

func (stmt *BeginStmt) String() string {
	return "BEGIN"
}

type RollbackStmt struct{}

func (stmt *RollbackStmt) String() string {
	return "ROLLBACK"
}

type SetIsolationStmt struct {
	Isolation txn.IsolationLevel
}

func (stmt *SetIsolationStmt) String() string {
	return "SET TRANSACTION ISOLATION LEVEL " + stmt.Isolation.Name
}
