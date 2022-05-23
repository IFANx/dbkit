package gen

import (
	"dbkit/internal/common"

	log "github.com/sirupsen/logrus"
)

type CockroachDataType int8

const (
	TypeInt2 = iota
	TypeInt4
	TypeInt8
	TypeSerial2
	TypeSerial4
	TypeSerial8
	TypeBit
	TypeVarBit
	TypeArray
	TypeBool
	TypeString
	TypeFloat
	TypeBytes
	TypeDecimal
	TypeTimestamp
)

func (dataType CockroachDataType) DBMS() common.DBMS {
	return common.COCKROACHDB
}

func (dataType CockroachDataType) Name() string {
	switch dataType {
	case TypeInt2:
		return "int2"
	case TypeInt4:
		return "int4"
	case TypeInt8:
		return "int8"
	case TypeSerial2:
		return "serial2"
	case TypeSerial4:
		return "serial4"
	case TypeSerial8:
		return "serial8"
	case TypeBit:
		return "bit"
	case TypeVarBit:
		return "varbit"
	case TypeArray:
		return "array"
	case TypeBool:
		return "bool"
	case TypeString:
		return "string"
	case TypeFloat:
		return "float"
	case TypeBytes:
		return "bytes"
	case TypeDecimal:
		return "decimal"
	case TypeTimestamp:
		return "timestamp"
	default:
		log.Infof("Unsupported data type: %v", dataType)
		panic("Unreachable")
	}
}

func (dataType CockroachDataType) HasSize() bool {
	switch dataType {
	case TypeVarBit, TypeBytes, TypeArray, TypeString:
		return true
	default:
		return false
	}
}

func (dataType CockroachDataType) IsString() bool {
	switch dataType {
	case TypeString:
		return true
	default:
		return false
	}
}

func (dataType CockroachDataType) IsNumeric() bool {
	switch dataType {
	case TypeInt2, TypeInt4, TypeInt8, TypeSerial2, TypeSerial4, TypeSerial8, TypeFloat, TypeDecimal:
		return true
	default:
		return false
	}
}

func (dataType CockroachDataType) canBePrimaryKey() bool {
	switch dataType {
	case TypeInt2, TypeInt4, TypeInt8, TypeSerial2, TypeSerial4, TypeSerial8:
		return true
	default:
		return false
	}
}

// GenRandomVal TODO
func (dataType CockroachDataType) GenRandomVal() string {
	switch dataType {
	case TypeInt2:
		return "int2"
	case TypeInt4:
		return "int4"
	case TypeInt8:
		return "int8"
	case TypeSerial2:
		return "serial2"
	case TypeSerial4:
		return "serial4"
	case TypeSerial8:
		return "serial8"
	case TypeBit:
		return "bit"
	case TypeVarBit:
		return "varbit"
	case TypeArray:
		return "array"
	case TypeBool:
		return "bool"
	case TypeString:
		return "string"
	case TypeFloat:
		return "float"
	case TypeBytes:
		return "bytes"
	case TypeDecimal:
		return "decimal"
	case TypeTimestamp:
		return "timestamp"
	default:
		log.Infof("Unsupported data type")
		panic("Unreachable")
	}
}
