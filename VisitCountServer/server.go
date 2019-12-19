package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB
func init() {
	db, _ = sql.Open("mysql","root:qwertyui@tcp(172.19.0.2:3306)/")
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS visitors")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("USE visitors")
	if err != nil {
		log.Println(err)
	}
	db.Close()

	db, err = sql.Open("mysql","root:qwertyui@tcp(172.19.0.2:3306)/visitors")
	_,err = db.Exec("CREATE TABLE IF NOT EXISTS visits ( visitors INT )")
	if err != nil {
		log.Println(err)
	}
	_,err = db.Exec("INSERT INTO visits values (0)")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(writer http.ResponseWriter, request *http.Request) {
	rows,err := db.Query("SELECT visitors FROM visits")
	if err != nil {
		log.Println(err)
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Fprintf(writer,"Number of visits are %d", count)
	_,err = db.Exec("UPDATE visits SET visitors = ? ", count+1)
	if err != nil {
		log.Println(err)
	}
}
