package healthz

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseChecker struct {
	db *sql.DB
}

func NewDatabaseChecker(driverName, dataSourceName string) (*DatabaseChecker, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DatabaseChecker{db}, nil
}

func (dc *DatabaseChecker) Ping() error {
	err := dc.db.Ping()
	if err != nil {
		return err
	}
	return nil
}
