package common

import (
	"dbkit/internal/common/isolation"
	"dbkit/internal/common/stmt"
)

type Transaction struct {
	Isolation isolation.IsolationLevel
	Stmts     []stmt.Statement
}
