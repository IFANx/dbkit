package internal

type DBKitConfig struct {
	Oracle     string
	Target     string
	DataSource DataSourceConfig
	MySQL      MySQLConfig
	TiDB       TiDBConfig
	Cockroach  CockroachConfig
	SQLite     SQLiteConfig
}

type DataSourceConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

type MySQLConfig struct {
	DBName   string
	Port     int
	Host     string
	Username string
	Password string
}

type TiDBConfig struct {
	DBName   string
	Port     int
	Host     string
	Username string
	Password string
}

type CockroachConfig struct {
	DBName   string
	Port     int
	Host     string
	Username string
	Password string
}

type SQLiteConfig struct {
	DBName   string
	Port     int
	Host     string
	Username string
	Password string
}
