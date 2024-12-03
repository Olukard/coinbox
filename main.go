package main

import (
	"fmt"
	"net/http"

	"coinbox/db"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Получен запрос")
}

func main() {

	fmt.Println("Проверяем наличие базы данных...")
	if !db.CheckDBexists() {
		fmt.Println("База данных не найдена, создаем...")
		db.CreateDB()
	} else {
		fmt.Println("База данных найдена.")
	}

	fmt.Println("Запускаем сервер")
	http.HandleFunc(`/`, mainHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}
