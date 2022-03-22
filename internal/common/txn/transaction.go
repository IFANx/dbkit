package txn

import (
	"dbkit/internal/common/stmt"
)

type Transaction struct {
	Isolation IsolationLevel
	Stmts     []stmt.Statement
}
