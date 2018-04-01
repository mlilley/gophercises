package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	urlshort "github.com/mlilley/gophercises/ex2/bonus3"
)

type config struct {
	paths map[string]string
}

func initDb() {
	db, err := sql.Open("sqlite3", "./config.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("drop table if exists config; create table config (path text, url text)")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into config (path, url) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("/urlshort-json", "https://github.com/gophercises/urlshort")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("/urlshort-final-json", "https://github.com/gophercises/urlshort")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func loadConfig() (config, error) {
	var c config
	c.paths = make(map[string]string)

	db, err := sql.Open("sqlite3", "./config.db")
	if err != nil {
		return c, err
	}
	defer db.Close()

	rows, err := db.Query("select path, url from config;")
	if err != nil {
		return c, err
	}
	defer rows.Close()

	for rows.Next() {
		var path string
		var url string
		err = rows.Scan(&path, &url)
		if err != nil {
			return c, err
		}
		c.paths[path] = url
	}

	err = rows.Err()
	if err != nil {
		return c, err
	}

	return c, nil
}

func main() {
	mux := defaultMux()

	// initDb()
	c, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mapHandler := urlshort.MapHandler(c.paths, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
