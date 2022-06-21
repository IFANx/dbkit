package statement

import (
	"dbkit/internal/randomly"
	"strconv"
	"strings"
)

type CreateTableStmt struct {
	TableName       string
	Columns         []Column
	TableOptions    []TableOption
	PartitionOption PartitionOptions
}

func (stmt *CreateTableStmt) String() string {
	sql := "CREATE"
	//TODO support temporary tables in the schema
	sql += " TABLE"
	if randomly.RandBool() {
		sql += " IF NOT EXISTS"
	}
	sql += " " + stmt.TableName
	//TODO support LIKE
	sql += "("
	for idx, col := range stmt.Columns {
		if idx != 0 {
			sql += ", "
		}
		sql += col.String()
	}
	sql += ") "
	for _, tabOpt := range stmt.TableOptions {
		if tabOpt == TabOptAutoIncrement {
			sql += "AUTO_INCREMENT = "
			sql += strconv.Itoa(randomly.RandIntGap(1, 10))
		} else if tabOpt == TabOptAvgRowLength {
			sql += "AVG_ROW_LENGTH = "
			sql += strconv.Itoa(randomly.RandIntGap(1, 10))
		} else if tabOpt == TabOptChecksum {
			sql += "CHECKSUM = 1"
		} else if tabOpt == TabOptCompression {
			sql += "COMPRESSION = '"
			sql += randomly.RandPickOneStr([]string{"ZLIB", "LZ4", "NONE"})
			sql += "'"
		} else if tabOpt == TabOptDelayKeyWrite { //Set this to 1 if you want to delay key updates for the table until the table is closed.
			sql += "DELAY_KEY_WRITE = "
			sql += strconv.Itoa(randomly.RandPickNInt([]int{0, 1}, 1)[0])
		} else if tabOpt == TabOptEngine {
			sql += "ENGINE = "
			sql += randomly.RandPickOneStr([]string{"InnoDB", "MyISAM", "MEMORY", "HEAP", "CSV", "ARCHIVE"})
		} else if tabOpt == TabOptInsertMethod {
			sql += "INSERT_METHOD = "
			sql += randomly.RandPickOneStr([]string{"NO", "FIRST", "LAST"})
		} else if tabOpt == TabOptKeyBlockSize {
			sql += "KEY_BLOCK_SIZE = "
			sql += strconv.Itoa(randomly.RandIntGap(1, 10))
		} else if tabOpt == TabOptMaxRows {
			sql += "MAX_ROWS = "
			sql += strconv.Itoa(randomly.RandIntGap(5, 10))
		} else if tabOpt == TabOptMinRows {
			sql += "MIN_ROWS = "
			sql += strconv.Itoa(randomly.RandIntGap(1, 5))
		} else if tabOpt == TabOptPackKeys { //Set this option to 1 if you want to have smaller indexes.
			sql += "PACK_KEYS = "
			sql += randomly.RandPickOneStr([]string{"1", "0", "DEFAULT"})
		} else if tabOpt == TabOptStatsAutoRecalc {
			sql += "STATS_AUTO_RECALC = "
			sql += randomly.RandPickOneStr([]string{"1", "0", "DEFAULT"})
		} else if tabOpt == TabOptStatsPersistent {
			sql += "STATS_PERSISTENT = "
			sql += randomly.RandPickOneStr([]string{"1", "0", "DEFAULT"})
		} else if tabOpt == TabOptStatsSamplePages {
			sql += "STATS_SAMPLE_PAGES = "
			sql += strconv.Itoa(randomly.RandIntGap(1, 10))
		} else {
			panic("Unsupported table option")
		}
	}

	if randomly.RandBool() {
		return sql
	}

	sql += " PARTITION BY"
	if stmt.PartitionOption == PartOptHASH {
		if randomly.RandBool() {
			sql += " LINEAR"
		}
		sql += " HASH("
		sql += stmt.Columns[randomly.RandIntGap(0, len(stmt.Columns)-1)].Name
		sql += ")"
	} else if stmt.PartitionOption == PartOptKEY {
		if randomly.RandBool() {
			sql += " LINEAR"
		}
		sql += " KEY"
		if randomly.RandBool() {
			sql += " ALGORITHM="
			sql += strconv.Itoa(randomly.RandPickNInt([]int{1, 2}, 1)[0])
		}
		sql += " ("
		colNames := make([]string, 0)
		for _, col := range stmt.Columns {
			if randomly.RandBool() {
				colNames = append(colNames, col.Name)
			}
		}
		sql += strings.Join(colNames, ",")
		sql += ")"
	} else {
		panic("Unsupported partition option")
	}

	return sql
}

type Column struct {
	Name       string
	Type       ColumnType
	Constraint []ColumnOptions
}

func (col *Column) String() string {
	colExpr := col.Name + " "
	for _, colOpt := range col.Constraint {
		if colOpt == ColOptColumnFormat {
			colExpr += "COLUMN_FORMAT "
			colExpr += randomly.RandPickOneStr([]string{"FIXED", "DYNAMIC", "DEFAULT"})
		} else if colOpt == ColOptComment {
			colExpr += "COMMENT " + randomly.RandAlphabetStrLen(4)
		} else if colOpt == ColOptNullOrNotNull {
			if randomly.RandBool() {
				colExpr += "NULL"
			} else {
				colExpr += "NOT NULL"
			}
		} else if colOpt == ColOptPrimaryKey {
			colExpr += "PRIMARY KEY"
		} else if colOpt == ColOptStorage {
			colExpr += "STORAGE "
			colExpr += randomly.RandPickOneStr([]string{"DISK", "MEMORY"})
		} else if colOpt == ColOptUnique {
			colExpr += "UNIQUE"
			if randomly.RandBool() {
				colExpr += " KEY"
			}
		} else {
			panic("Unsupported column constraint")
		}
	}
	return colExpr
}

type ColumnType = int

type ColumnOptions = int

const (
	ColOptColumnFormat = iota
	ColOptComment
	ColOptNullOrNotNull
	ColOptPrimaryKey
	ColOptStorage
	ColOptUnique
)

type TableOption = int

const (
	TabOptAutoIncrement    = iota //The initial AUTO_INCREMENT value for the table.
	TabOptAvgRowLength            //An approximation of the average row length for your table.
	TabOptChecksum                //MyISAM only
	TabOptCompression             //The compression algorithm used for page level compression for InnoDB tables.
	TabOptDelayKeyWrite           //MyISAM only
	TabOptEngine                  //
	TabOptInsertMethod            //INSERT_METHOD is an option useful for MERGE tables only
	TabOptKeyBlockSize            //not support temporary table
	TabOptMaxRows                 //The maximum number of rows you plan to store in the table.
	TabOptMinRows                 //The minimum number of rows you plan to store in the table.
	TabOptPackKeys                //MyISAM only
	TabOptStatsAutoRecalc         //Specifies whether to automatically recalculate persistent statistics for an InnoDB table.
	TabOptStatsPersistent         //Specifies whether to enable persistent statistics for an InnoDB table.
	TabOptStatsSamplePages        //The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by ANALYZE TABLE.
)

type PartitionOptions int

const (
	PartOptHASH = iota
	PartOptKEY
)
