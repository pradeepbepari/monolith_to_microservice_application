package database

import (
	"database/sql"
	"fmt"
)

func ConnectDB(dbUser, dbPass, dbHost, dbPort, dbNaame string) (*sql.DB, error) {
	postgresDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbNaame,
	)
	conn, err := sql.Open("postgres", postgresDsn)
	if err != nil {
		return nil, err
	}
	if err := PingDB(conn); err != nil {
		return nil, err
	}
	return conn, nil
}
func PingDB(db *sql.DB) error {
	return db.Ping()
}
func CloseDB(db *sql.DB) error {
	return db.Close()
}
