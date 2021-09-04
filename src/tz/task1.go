package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main()  {
	db, err := sql.Open("sqlite3", "./tz.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, name TEXT, birthday TEXT)")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO people (name, birthday) VALUES (?, ?)")
	var name string
	fmt.Println("Name: ")
	fmt.Scanf("%s", &name)

	var birthday string
	fmt.Println("Birthday: ")
	fmt.Scanf("%s", &birthday)
	fmt.Println("The data about the person was successfully added to the table called \"People\" in the database")

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	stmt.Exec(&name, &birthday)

}

