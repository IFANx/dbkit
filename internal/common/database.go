package common

import "dbkit/internal/common/dbms"

type Database struct {
	DBMS   dbms.DBMS
	DBName string
	Tables []*Table
}
