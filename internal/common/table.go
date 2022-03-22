package common

import "dbkit/internal"

type Table struct {
	DBMS        internal.DBMS
	Name        string
	ColumnNames []string
	Columns     map[string]*Column
	Indexes     map[string]*Index

	indexCount    int
	hasPrimaryKey bool
}

func (table *Table) UpdateSchema() {
	
}
