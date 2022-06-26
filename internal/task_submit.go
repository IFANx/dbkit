package internal

import (
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/oracle"
	"github.com/jmoiron/sqlx"
	"log"
)

type TaskType int

const (
	TaskTypeTest = iota
	TaskTypeDiff
	TaskTypeVerify
)

type TaskSubmit struct {
	Type        TaskType
	OracleList  []oracle.Oracle
	TargetTypes []dbms.DBMS
	ConnList    []*sqlx.DB
	Limit       float32
	Model       string // TODO: Change to enumerate
	Comments    string
}

func BuildTaskFromSubmit(submit *TaskSubmit) (int, error) {
	log.Printf("%+v\n", submit)
	return 1, nil
}
