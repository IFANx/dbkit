package common

import (
	"dbkit/internal/common/dbms"
	"dbkit/internal/common/stmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	DBMS   dbms.DBMS
	DBName string
	Conn   *sqlx.DB
	Tables []*Table
}

func (db *Database) Refresh() error {
	var err error
	err = db.ExecSQL("DROP DATABASE IF EXISTS " + db.DBName)
	if err != nil {
		return err
	}
	err = db.ExecSQL("CREATE DATABASE " + db.DBName)
	if err != nil {
		return err
	}
	err = db.ExecSQL("USE " + db.DBName)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Queryx(query string) (*sqlx.Rows, error) {
	return db.Conn.Queryx(query)
}

func (db *Database) QuerySQL(query string) ([][]interface{}, error) {
	rows, err := db.Conn.Queryx(query)
	if err != nil {
		log.Warnf("Fail to query: %s, cause: %s", query, err)
		return nil, err
	}
	defer rows.Close()
	res := make([][]interface{}, 0)
	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			log.Warnf("Fail to query: %s, cause: %s", query, err)
			return nil, err
		}
		res = append(res, cols)
	}
	return res, nil
}

func (db *Database) Query(stmt stmt.SelectStmt) ([][]interface{}, error) {
	var query string
	if db.DBMS == dbms.TIDB {
		query = stmt.StringInMode()
	} else {
		query = stmt.String()
	}
	return db.QuerySQL(query)
}

func (db *Database) ExecSQLIgnoreError(sql string) {
	_, err := db.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
	}
}

func (db *Database) ExecSQL(sql string) error {
	_, err := db.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
	}
	return err
}

func (db *Database) ExecSQLAffectedRow(sql string) (int, error) {
	res, err := db.Conn.Exec(sql)
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", sql, err)
		return 0, err
	}
	return int(count), err
}

func (db *Database) ExecUpdate(stmt stmt.UpdateStmt) (int, error) {
	return db.ExecSQLAffectedRow(stmt.String())
}

func (db *Database) ExecDelete(stmt stmt.DeleteStmt) (int, error) {
	return db.ExecSQLAffectedRow(stmt.String())
}

func (db *Database) ExecInsert(stmt stmt.InsertStmt) error {
	_, err := db.Conn.Exec(stmt.String())
	if err != nil {
		log.Warnf("Fail to execute: %s, cause: %s", stmt.String(), err)
	}
	return err
}
