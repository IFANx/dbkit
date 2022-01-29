package common


type Index struct {
	Name			string
	IndexedCols		[]string
	IsPrimary		bool
	IsUnique		bool
}
