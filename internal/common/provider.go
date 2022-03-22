package common

import "github.com/jmoiron/sqlx"

type Provider interface {
	GenerateConn() *sqlx.Conn
	GenerateTable() *Table
	UpdateSchema()
}
