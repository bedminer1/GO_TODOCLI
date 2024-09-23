//go:build !inmemory
// +build !inmemory

package repository

import (
	"database/sql"
	"sync"
	"time"

	"github.com/bedminer1/pomo/pomodoro"
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

func (r *dbRepo) Create(i pomodoro.Interval) (int64, error) {
	r.Lock()
	defer r.Unlock()

	// prepare INSERT statement
	insStmt, err := r.db.Prepare("INSERT INTO interval VALUES(NULL, ?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer insStmt.Close()

	// Exec INSERT statement
	res, err := insStmt.Exec(i.StartTime, i.PlannedDuration, i.ActualDuration, i.Category, i.State)
	if err != nil {
		return 0, err
	}

	// INSERT res
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *dbRepo) Update(i pomodoro.Interval) error {
	r.Lock()
	defer r.Unlock()

	updStmt, err := r.db.Prepare("UPDATE interval SET start_time=?, actual_duration=?, state=? WHERE id=?")
	if err != nil {
		return err
	}
	defer updStmt.Close()

	res, err := updStmt.Exec(i.StartTime, i.ActualDuration, i.State, i.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	return err
}