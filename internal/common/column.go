package common

type Column struct {
	Table      Table
	Name       string
	Type       DataType
	NotNull    bool
	Unique     bool
	Length     int
	ValueCache []string
}
