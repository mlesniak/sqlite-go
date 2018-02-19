package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	// This is the usual way to include an SQL driver in golang. Actually we are not using
	// any imports from the package explictly.
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./demo.db")

	initialize(database)
	fill(database)
	query(database)
}

func initialize(database *sql.DB) {
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS storage (id INTEGER PRIMARY KEY, key TEXT, value TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Exec()
}

func fill(database *sql.DB) {
	statement, _ := database.Prepare("INSERT INTO storage (key, value) VALUES (?, ?)")
	statement.Exec("timestamp", time.Now())
}

func query(database *sql.DB) {
	rows, _ := database.Query("SELECT id, key, value FROM storage")
	var id int
	var key string
	var value string
	for rows.Next() {
		rows.Scan(&id, &key, &value)
		fmt.Println(strconv.Itoa(id) + ":" + key + ":" + value)
	}
}
