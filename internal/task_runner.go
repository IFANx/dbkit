package internal

type TaskRunner interface {
	RunTask(ctx *TaskContext)
}
