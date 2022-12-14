package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Conectar() (*sql.DB, error) {
	db, err := sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
