package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := "localhost"
	port := "8080"
	conn := fmt.Sprint(host, ":", port)
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./webroot")))
	router.Handle("/ws", wsHandler{})

	log.Printf("serving tic-tac-toe on %v", conn)

	log.Fatal(http.ListenAndServe(conn, router))
}
