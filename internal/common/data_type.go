package common

import "dbkit/internal/common/dbms"

type DataType interface {
	DBMS() dbms.DBMS
	Name() string
	HasSize() bool
	IsString() bool
	IsNumeric() bool
	CanBePrimary() bool
	GenRandomVal() string
}
