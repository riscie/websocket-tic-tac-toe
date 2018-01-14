package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	host, port := "", ""
	//reading environment specific settings .env file. If missing; assume dev env
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
		host = "127.0.0.1"
		port = "8080"
	} else {
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")
	}

	//preparing mux and server
	conn := fmt.Sprint(host, ":", port)
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./webroot")))
	router.Handle("/ws", wsHandler{})

	//serving
	log.Printf("serving tic-tac-toe on %v", conn)
	log.Fatal(http.ListenAndServe(conn, router))
}
