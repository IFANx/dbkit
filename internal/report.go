package internal

type IssueReport struct {
	ReportID   int64
	TestID     int64
	InputStmt  string
	InputRes   string
	OracleStmt string
	OracleRes  string
	Category   string
	State      string
}
