package database

import "database/sql"

type Database interface {
	Connect(url string) error
	GetDB() *sql.DB
	Close() error
}
