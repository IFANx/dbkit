package config

type DBKitConfig struct {
	DataSource DataSourceConfig
}

type DataSourceConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}
