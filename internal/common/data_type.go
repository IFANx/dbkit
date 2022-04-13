package common

type DataType interface {
	DBMS() DBMS
	Name() string
	HasSize() bool
	IsString() bool
	IsNumeric() bool
	CanBePrimary() bool
	GenRandomVal() string
}
