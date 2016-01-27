package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//reading environment specific settings .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	//preparing mux and server
	conn := fmt.Sprint(host, ":", port)
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./webroot")))
	router.Handle("/ws", wsHandler{})

	//serving
	log.Printf("serving tic-tac-toe on %v", conn)
	log.Fatal(http.ListenAndServe(conn, router))
}
