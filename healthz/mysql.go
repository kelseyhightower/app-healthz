package healthz

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLChecker struct {
	db *sql.DB
}

func NewMySQLChecker(driverName, dataSourceName string) (*MySQLChecker, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &MySQLChecker{db}, nil
}

func (mc *MySQLChecker) Ping() error {
	err := mc.db.Ping()
	if err != nil {
		return err
	}
	return nil
}
