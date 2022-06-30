package internal

import "dbkit/internal/common"

type TaskRunner interface {
	RunTask(ctx common.OracleRuntime)
}
