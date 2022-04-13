package internal

import "dbkit/internal/common"

type Column struct {
	Table      *Table
	Name       string
	Type       common.DataType
	NotNull    bool
	Unique     bool
	Primary    bool
	Length     int
	ValueCache []string
}
