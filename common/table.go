package common


type Table struct {
	Name			string
	ColumnNames		[]string
	Columns			map[string]Column
	Indexes			map[string]Index

	indexCount		int
	hasPrimaryKey	bool
}