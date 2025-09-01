package database

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	Gorm *gorm.DB
	SQL  *sql.DB
}

func NewPostgresDatabase() *PostgresDB {
	return &PostgresDB{}
}

func (p *PostgresDB) Connect(dsn string) error {
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	p.Gorm = gdb
	p.SQL = sqlDB
	return nil
}

func (p *PostgresDB) GetDB() *gorm.DB {
	return p.Gorm
}

func (p *PostgresDB) Close() error {
	if p.SQL != nil {
		return p.SQL.Close()
	}
	return nil
}
