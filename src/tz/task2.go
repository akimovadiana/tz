package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func main() {
	db, err := sql.Open("sqlite3", "./tz.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var text string
	fmt.Println("Type the name of the person whose data you want to find in the database: ")
	fmt.Scanf("%s", &text)

	if text != ""{
		rows, err := db.Query("SELECT id, name, birthday FROM people WHERE name LIKE '" + text + "%'")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		printRows(rows)
	} else {
		rows, err := db.Query("SELECT id, name, birthday FROM people")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		printRows(rows)
	}
}

func printRows(rows *sql.Rows) {
	var id int
	var name string
	var birthday string
	for rows.Next() {
		rows.Scan(&id, &name, &birthday)
		fmt.Println(strconv.Itoa(id) + ": " + name + " " + birthday)
	}
}