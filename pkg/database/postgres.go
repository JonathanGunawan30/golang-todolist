package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type postgresDB struct {
	DB *sql.DB
}

func NewPostgresDatabase() Database {
	return &postgresDB{}
}

func (postgres *postgresDB) Connect(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	if err := db.Ping(); err != nil {
		db.Close()
		return err
	}
	postgres.DB = db
	return nil
}
func (postgres *postgresDB) GetDB() *sql.DB {
	return postgres.DB
}

func (postgres *postgresDB) Close() error {
	return postgres.DB.Close()
}
