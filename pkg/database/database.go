package database

import "database/sql"

type Database interface {
	Connect(dsn string) error
	GetDB() *sql.DB
	Close() error
}
