package mysql

import "dbkit/internal/common"

type MySQLProvider struct {
}

func NewMySQLProvider() *MySQLProvider {
	provider := MySQLProvider{}
	return &provider
}

func (provider MySQLProvider) GenerateTable() common.Table {
	return common.Table{}
}
