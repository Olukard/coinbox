package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DBfile = "users.db"

const DBinitCommand = `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(256) NOT NULL DEFAULT "", 
	lastName VARCHAR(256) NOT NULL DEFAULT "",
	balance INTEGER NOT NULL DEFAULT 0,
	role VARCHAR(128) NOT NULL DEFAULT "",
	status VARCHAR(128) NOT NULL DEFAULT ""
	);`

const DBindexCommand = `
	CREATE INDEX id_indx ON users (date)
	`

//checkDBexists проверяет существование файла базы данных в директории проекта

func CheckDBexists() bool {

	_, err := os.Stat(DBfile)

	return err == nil
}

//CreateDB создает файл базы данных с индексакцией в соотвествии с заданными константами DBinitCommand и DBindexCommand

func CreateDB() {

	db, err := sql.Open("sqlite3", "./"+DBfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(DBinitCommand)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("База данных создана, проводим индексацию...")

	_, err = db.Exec(DBindexCommand)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Индексация завершена.")

}
