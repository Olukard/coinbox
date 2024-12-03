package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"coinbox/db"
	"coinbox/handlers"
)

func main() {

	fmt.Println("Проверяем наличие базы данных...")
	if !db.CheckDBexists() {
		fmt.Println("База данных не найдена, создаем...")
		db.CreateDB()
	} else {
		fmt.Println("База данных найдена.")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.AddUserHandler).Methods("GET")

	// Graceful Shutdown
	server := &http.Server{Addr: ":8080", Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
	// Close database connection here if needed (db.Close())
}
