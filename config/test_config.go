package config

import "github.com/jmoiron/sqlx"

type TestConfig struct {
	Oracle string
	Target string
	Conn   *sqlx.DB
}
