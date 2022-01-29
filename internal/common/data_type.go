package common

import "dbkit/internal"

type DataType interface {
	DBMS() internal.DBMS
	Name() string
	HasSize() bool
	IsString() bool
	IsNumeric() bool
	CanBePrimary() bool
	GenRandomVal() string
}
