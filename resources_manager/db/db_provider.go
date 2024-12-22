package db

import (
	"database/sql"
)

type DBProvider struct {
	db *sql.DB
}

func NewDBProvider(db *sql.DB) *DBProvider {
	return &DBProvider{db: db}
}

func (p *DBProvider) GetDB() DB {
	return p.db
}
