package common

import (
	"dbkit/internal/common/stmt"
	"dbkit/internal/common/txn"
)

type Transaction struct {
	Isolation txn.IsolationLevel
	Stmts     []stmt.Statement
}
