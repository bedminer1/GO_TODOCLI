//go:build !inmemory
// +build !inmemory

package repository

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createTableInterval string = `CREATE TABLE IF NOT EXISTS "interval" (
"id" INTEGER,
"start_time" DATETIME NOT NULL,
"planned_duration" INTEGER DEFAULT 0,
"actual_duration" INTEGER DEFAULT 0,
"category" TEXT NOT NULL,
"state" INTEGER DEFAULT 1,
PRIMARY KEY("id")
);`
)

type dbRepo struct {
	db *sql.DB
	sync.RWMutex
}

func NewSQLite3Repo(dbfile string) (*dbRepo, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(30*time.Minute)
	db.SetMaxOpenConns(1)

	// verify connection established
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createTableInterval); err != nil {
		return nil, err
	}

	return &dbRepo{
		db: db,
	}, nil
}