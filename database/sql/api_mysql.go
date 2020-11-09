package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	db = database
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer db.Close()

	queryForList()
	queryForObject()
	insert()
	transaction()

}

func queryForList() {
	rows, err := db.Query("select id, name from tb_user where id = ?", 1)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("无数据")
		}
		log.Fatal(err)
	}

	defer rows.Close()

	var result = make(map[string]interface{})

	for rows.Next() {
		var (
			id   int
			name string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		result[name] = id
	}

	if err := rows.Err(); err != nil {
		log.Println("rows err: ", err)
	}

	log.Println("queryForList#", result)
}

func queryForObject() {
	var (
		id   int
		name string
	)
	row := db.QueryRow("select id, name from tb_user where id = ? limit 1", 1)
	err := row.Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("queryForObject#", id, name)
}

func insert() {
	stmt, err := db.Prepare("insert into tb_user(id,name) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Exec(1, "test_name")
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := result.RowsAffected()

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

}

func transaction() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("insert into tb_user(id,name) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(1, "test_name2")

	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
