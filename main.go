package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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
	bytes, err := ioutil.ReadFile("init.sql")
	if err != nil {
		log.Println("No initialization file found.")
		return
	}

	cmds := strings.Split(string(bytes), ";")
	for _, cmd := range cmds {
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		statement, err := database.Prepare(cmd)
		if err != nil {
			log.Println("Unable to execute statement: ", cmd)
			return
		}
		log.Println(cmd)
		statement.Exec()
	}
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
		log.Println(strconv.Itoa(id) + ":" + key + ":" + value)
	}
}
