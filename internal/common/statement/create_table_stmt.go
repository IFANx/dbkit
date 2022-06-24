package statement

import (
	"dbkit/internal/mysql/gen"
	"strconv"
	"strings"
)

type CreateTableStmt struct {
	TableName          string
	Columns            []Column
	TableOptions       []TableOption
	AutoIncrement      int
	AvgRowLength       int
	Compression        Compressions
	DelayKeyWrite      bool
	TableEngine        TableEngines
	InsertMethod       InsertMethods
	KeyBlockSize       int
	MaxRows            int
	MinRows            int
	PackKey            int
	StatsAutoRecalc    int
	StatsPersistent    int
	StatsSamplePages   int
	PartitionOption    PartitionOptions
	PartitionColumn    Column
	PartitionAlgorithm int
	PartitionColumns   []Column
}

func (stmt *CreateTableStmt) String() string {
	sql := "CREATE"
	//TODO support temporary tables in the schema
	sql += " TABLE"
	// if randomly.RandBool() {
	// 	sql += " IF NOT EXISTS"
	// }
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
			sql += strconv.Itoa(stmt.AutoIncrement)
		} else if tabOpt == TabOptAvgRowLength {
			sql += "AVG_ROW_LENGTH = "
			sql += strconv.Itoa(stmt.AvgRowLength)
		} else if tabOpt == TabOptChecksum {
			sql += "CHECKSUM = 1"
		} else if tabOpt == TabOptCompression {
			sql += "COMPRESSION = '"
			compreDic := []string{"ZLIB", "LZ4", "NONE"}
			sql += compreDic[stmt.Compression]
			sql += "'"
		} else if tabOpt == TabOptDelayKeyWrite { //Set this to 1 if you want to delay key updates for the table until the table is closed.
			sql += "DELAY_KEY_WRITE = "
			if stmt.DelayKeyWrite {
				sql += strconv.Itoa(1)
			} else {
				sql += strconv.Itoa(0)
			}
		} else if tabOpt == TabOptEngine {
			sql += "ENGINE = "
			tabEngDic := []string{"ARCHIVE", "CSV", "HEAP", "InnoDB", "MEMORY", "MyISAM"}
			sql += tabEngDic[stmt.TableEngine]
		} else if tabOpt == TabOptInsertMethod {
			sql += "INSERT_METHOD = "
			insertMethods := []string{"NO", "FIRST", "LAST"}
			sql += insertMethods[stmt.InsertMethod]
		} else if tabOpt == TabOptKeyBlockSize {
			sql += "KEY_BLOCK_SIZE = "
			sql += strconv.Itoa(stmt.KeyBlockSize)
		} else if tabOpt == TabOptMaxRows {
			sql += "MAX_ROWS = "
			sql += strconv.Itoa(stmt.MaxRows)
		} else if tabOpt == TabOptMinRows {
			sql += "MIN_ROWS = "
			sql += strconv.Itoa(stmt.MinRows)
		} else if tabOpt == TabOptPackKeys { //Set this option to 1 if you want to have smaller indexes.
			sql += "PACK_KEYS = "
			if stmt.PackKey == -1 {
				sql += "DEFAULT"
			} else {
				sql += strconv.Itoa(stmt.PackKey)
			}
		} else if tabOpt == TabOptStatsAutoRecalc {
			sql += "STATS_AUTO_RECALC = "
			if stmt.StatsAutoRecalc == -1 {
				sql += "DEFAULT"
			} else {
				sql += strconv.Itoa(stmt.PackKey)
			}
		} else if tabOpt == TabOptStatsPersistent {
			sql += "STATS_PERSISTENT = "
			if stmt.StatsPersistent == -1 {
				sql += "DEFAULT"
			} else {
				sql += strconv.Itoa(stmt.PackKey)
			}
		} else if tabOpt == TabOptStatsSamplePages {
			sql += "STATS_SAMPLE_PAGES = "
			sql += strconv.Itoa(stmt.StatsSamplePages)
		} else if tabOpt == -1 {
			//do nothing
		} else {
			panic("Unsupported table option")
		}
	}

	if stmt.PartitionOption == -1 {
		return sql
	}

	sql += " PARTITION BY"
	if stmt.PartitionOption == PartOptHASH {
		// if randomly.RandBool() {
		// 	sql += " LINEAR"
		// }
		sql += " HASH("
		sql += stmt.PartitionColumn.Name
		sql += ")"
	} else if stmt.PartitionOption == PartOptKEY {
		// if randomly.RandBool() {
		// 	sql += " LINEAR"
		// }
		sql += " KEY"
		if stmt.PartitionAlgorithm == -1 {

		} else if stmt.PartitionAlgorithm == 0 {
			sql += " ALGORITHM="
			sql += strconv.Itoa(1)
		} else if stmt.PartitionAlgorithm == 1 {
			sql += " ALGORITHM="
			sql += strconv.Itoa(2)
		} else {
			panic("Unsupported algorithm")
		}
		sql += " ("
		colNames := make([]string, 0)
		for _, col := range stmt.PartitionColumns {
			colNames = append(colNames, col.Name)
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
	Type       gen.MySQLDataType
	Constraint []ColumnOptions
	ColForMat  ColFormats
	Comment    string
	Storage    StorageOption
}

func (col *Column) String() string {
	colExpr := col.Name + " "
	for _, colOpt := range col.Constraint {
		if colOpt == ColOptColumnFormat {
			colExpr += "COLUMN_FORMAT "
			colFormatDic := []string{"FIXED", "DYNAMIC", "DEFAULT"}
			colExpr += colFormatDic[col.ColForMat]
		} else if colOpt == ColOptComment {
			colExpr += "COMMENT " + col.Comment
		} else if colOpt == ColOptNotNull {
			colExpr += "NOT NULL"
		} else if colOpt == ColOptNull {
			colExpr += "NULL"
		} else if colOpt == ColOptPrimaryKey {
			colExpr += "PRIMARY KEY"
		} else if colOpt == ColOptStorage {
			colExpr += "STORAGE "
			storageOptionDic := []string{"DISK", "MEMORY"}
			colExpr += storageOptionDic[col.Storage]
		} else if colOpt == ColOptUnique {
			colExpr += "UNIQUE"
			// if randomly.RandBool() {
			// 	colExpr += " KEY"
			// }
		} else if colOpt == -1 {
			//do nothing
		} else {
			panic("Unsupported column constraint")
		}
	}
	return colExpr
}

type ColumnOptions = int

const (
	ColOptColumnFormat = iota
	ColOptComment
	ColOptNotNull
	ColOptNull
	ColOptPrimaryKey
	ColOptStorage
	ColOptUnique
)

type ColFormats int

const (
	ColForMatDEFAULT = iota
	ColForMatDYNAMIC
	ColForMatFixed
)

type StorageOption int

const (
	StorageDisk = iota
	StorageMemory
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

type Compressions int

const (
	CompreLZ4 = iota
	CompreNONE
	CompreZLIB
)

type TableEngines int

const (
	TabEngArchive = iota
	TabEngCSV
	TabEngHeap
	TabEngInnoDB
	TabEngMemory
	TabEngMyISAM
)

type InsertMethods int

const (
	InsertFirst = iota
	InsertLast
	InsertNo
)

type PartitionOptions int

const (
	PartOptHASH = iota
	PartOptKEY
)
