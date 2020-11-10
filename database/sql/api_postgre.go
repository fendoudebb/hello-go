package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var pg *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "z-blog", "", "z-blog")
	db, err := sql.Open("postgres", psqlInfo)

	//connStr := "postgres://testuser:testpwd@127.0.0.1/testdb?sslmode=verify-full"
	//db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pg = db

	pg.SetConnMaxLifetime(time.Minute * 3)
	pg.SetMaxOpenConns(10)
	pg.SetMaxIdleConns(10)

	err = pg.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	defer pg.Close()

	var (
		id   int
		title string
	)
	row := pg.QueryRow("select id, title from post where id = $1 limit 1", 1)
	err := row.Scan(&id, &title)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("queryForObject#", id, title)


}