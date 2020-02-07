package rds

import (
	"database/sql"

	// Postgres Driver
	_ "github.com/lib/pq"
)

// ConnectPool Connect to a postgres instance
func ConnectPool(connectionString string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
