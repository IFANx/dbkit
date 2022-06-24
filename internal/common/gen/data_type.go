package gen

import (
	"dbkit/internal/common"
	"dbkit/internal/common/dbms"
	"dbkit/internal/randomly"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetRandomMySQLDataType() common.DataType {
	return MySQLDataType(randomly.RandIntGap(TypeBigInt, TypeYear))
}

type MySQLDataType int

const (
	TypeBigInt = iota
	TypeBinary
	TypeBit
	TypeBlob
	TypeBool
	TypeChar
	TypeDate
	TypeDateTime
	TypeDecimal
	TypeDouble
	TypeEnum
	TypeFloat
	TypeInt
	TypeJson
	TypeLongBlob
	TypeLongText
	TypeMediumInt
	TypeMediumText
	TypeSet
	TypeSmallInt
	TypeText
	TypeTime
	TypeTimestamp
	TypeTinyBlob
	TypeTinyInt
	TypeTinyText
	TypeVarBinary
	TypeVarchar
	TypeYear
)

func (dataType MySQLDataType) DBMS() dbms.DBMS {
	return dbms.MYSQL
}

func RandMySQLType() MySQLDataType {
	types := []MySQLDataType{TypeInt, TypeVarchar, TypeFloat}
	return types[randomly.RandIntGap(0, len(types))]
}

func (dataType MySQLDataType) Name() string {
	switch dataType {
	case TypeBit:
		return "bit"
	case TypeBool:
		return "boolean"
	case TypeTinyInt:
		return "tinyint"
	case TypeSmallInt:
		return "smallint"
	case TypeMediumInt:
		return "mediumint"
	case TypeInt:
		return "int"
	case TypeBigInt:
		return "bigint"
	case TypeFloat:
		return "float"
	case TypeDouble:
		return "double"
	case TypeDecimal:
		return "decimal"
	case TypeChar:
		return "char"
	case TypeVarchar:
		return "varchar"
	case TypeText:
		return "text"
	case TypeTinyText:
		return "tinytext"
	case TypeMediumText:
		return "mediumtext"
	case TypeLongText:
		return "longtext"
	case TypeBinary:
		return "binary"
	case TypeVarBinary:
		return "varbinary"
	case TypeBlob:
		return "blob"
	case TypeTinyBlob:
		return "tinyblob"
	case TypeLongBlob:
		return "longblob"
	case TypeEnum:
		return "enum"
	case TypeSet:
		return "set"
	case TypeDate:
		return "date"
	case TypeTime:
		return "time"
	case TypeDateTime:
		return "datetime"
	case TypeTimestamp:
		return "timestamp"
	case TypeYear:
		return "year"
	case TypeJson:
		return "json"
	default:
		log.Infof("Unsupported data type: %v", dataType)
		panic("Unreachable")
	}
}

func ParseDataType(name string) MySQLDataType {
	idx := strings.Index(name, "(")
	if idx > 0 {
		name = name[:idx]
	}
	switch name {
	case "bit":
		return TypeBit
	case "bool":
		return TypeBool
	case "tinyint":
		return TypeTinyInt
	case "smallint":
		return TypeSmallInt
	case "mediumint":
		return TypeMediumInt
	case "int":
		return TypeInt
	case "bigint":
		return TypeBigInt
	case "float":
		return TypeFloat
	case "double":
		return TypeDouble
	case "decimal":
		return TypeDecimal
	case "char":
		return TypeChar
	case "varchar":
		return TypeVarchar
	case "text":
		return TypeText
	case "tinytext":
		return TypeTinyText
	case "mediumtext":
		return TypeMediumText
	case "longtext":
		return TypeLongText
	case "binary":
		return TypeBinary
	case "varbinary":
		return TypeVarBinary
	case "blob":
		return TypeBlob
	case "tinyblob":
		return TypeTinyBlob
	case "longblob":
		return TypeLongBlob
	case "enum":
		return TypeEnum
	case "set":
		return TypeSet
	case "date":
		return TypeDate
	case "time":
		return TypeTime
	case "datetime":
		return TypeDateTime
	case "timestamp":
		return TypeTimestamp
	case "year":
		return TypeYear
	case "json":
		return TypeJson
	default:
		log.Infof("Unsupported data type: %v", name)
		panic("Unreachable")
	}
}

func (dataType MySQLDataType) CanBePrimary() bool {
	switch dataType {
	case TypeInt, TypeTinyInt, TypeMediumInt, TypeBigInt, TypeFloat, TypeDouble, TypeDecimal:
		return true
	default:
		return false
	}
}

func (dataType MySQLDataType) HasSize() bool {
	switch dataType {
	case TypeBit, TypeChar, TypeVarchar, TypeVarBinary, TypeText, TypeBlob:
		return true
	default:
		return false
	}
}

func (dataType MySQLDataType) IsString() bool {
	switch dataType {
	case TypeChar, TypeVarchar, TypeText, TypeTinyText, TypeMediumText, TypeLongText:
		return true
	default:
		return false
	}
}

func (dataType MySQLDataType) IsNumeric() bool {
	switch dataType {
	case TypeInt, TypeTinyInt, TypeMediumInt, TypeBigInt, TypeFloat, TypeDouble, TypeDecimal:
		return true
	default:
		return false
	}
}

func (dataType MySQLDataType) canBePrimaryKey() bool {
	switch dataType {
	case TypeInt, TypeTinyInt, TypeMediumInt, TypeBigInt, TypeFloat, TypeDouble, TypeDecimal,
		TypeChar, TypeVarchar, TypeText:
		return true
	default:
		return false
	}
}

func (dataType MySQLDataType) GenRandomVal() string {
	switch dataType {
	case TypeBit:
		if randomly.RandBool() {
			return "1"
		} else {
			return "0"
		}
	case TypeBool:
		if randomly.RandBool() {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case TypeTinyInt:
		return strconv.Itoa(randomly.RandIntGap(-128, 127))
	case TypeSmallInt:
		//return strconv.Itoa(randomly.RandIntGap(-32768, 32767))
		return strconv.Itoa(randomly.RandIntGap(-128, 127))
	case TypeMediumInt:
		//return strconv.Itoa(randomly.RandIntGap(-8388608, 8388607))
		return strconv.Itoa(randomly.RandIntGap(-128, 127))
	case TypeInt, TypeBigInt:
		//return strconv.Itoa(int(rand.Int31()))
		return strconv.Itoa(randomly.RandIntGap(-128, 127))
	case TypeDecimal, TypeFloat:
		return strconv.FormatFloat(float64(randomly.RandFloat()), 'f',
			randomly.RandIntGap(0, 5), 32)
	case TypeDouble:
		return strconv.FormatFloat(randomly.RandDouble(), 'f',
			randomly.RandIntGap(0, 10), 64)
	case TypeChar, TypeVarchar, TypeText, TypeTinyText, TypeMediumText, TypeLongText:
		return "'" + randomly.RandNormStrLen(randomly.RandIntGap(5, 10)) + "'"
	case TypeBinary, TypeVarBinary:
		return "'" + randomly.RandNormStrLen(randomly.RandIntGap(5, 10)) + "'"
	case TypeBlob, TypeTinyBlob, TypeLongBlob:
		return "'" + randomly.RandHexStrLen(randomly.RandIntGap(5, 10)) + "'"
	case TypeEnum:
		return "enum"
	case TypeSet:
		return "set"
	case TypeDate:
		return "'" + randomly.RandDateStr() + "'"
	case TypeTime:
		return "'" + randomly.RandTimeStr() + "'"
	case TypeDateTime:
		return "'" + randomly.RandDateTimeStr() + "'"
	case TypeTimestamp:
		return "'" + randomly.RandDateTimeStr() + "'"
	case TypeYear:
		return strconv.Itoa(randomly.RandIntGap(1901, 2155))
	case TypeJson:
		return "{type: json}"
	default:
		log.Infof("Unsupported data type: %v", dataType)
		panic("Unreachable")
	}
}
