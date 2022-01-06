package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "modernc.org/sqlite"
)

var db *sql.DB
var lock sync.Mutex

func Start() {
	var err error
	fmt.Println("Init Database")
	db, err = sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db.SetMaxOpenConns(1)
	db.Exec(`CREATE TABLE IF NOT EXISTS "users" (
		"id"	TEXT NOT NULL,
		"login"	TEXT NOT NULL,
		"displayName"	TEXT,
		PRIMARY KEY("id")
	);`)
}

func Stop() {
	lock.Lock()
	defer lock.Unlock()
	db.Close()
}
