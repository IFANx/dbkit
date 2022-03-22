package main

import (
	"dbkit/internal"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer logFile.Close()

	state := internal.GetState()

	dataSource := state.Config.DataSource
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/dbkit",
		dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Port)
	conn, _ := sqlx.Connect("mysql", connStr)
	rows, _ := conn.Query("desc t")
	types, _ := rows.Columns()
	for _, c := range types {
		fmt.Printf("%v\n", c)
	}

}
