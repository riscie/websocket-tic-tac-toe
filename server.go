package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//ip := "192.168.0.20"
	host := "192.168.0.20"
	port := "8080"
	conn := fmt.Sprint(host, ":", port)
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./webroot")))
	router.Handle("/ws", wsHandler{})

	log.Printf("serving tic-tac-toe on %v", conn)

	log.Fatal(http.ListenAndServe(conn, router))
}
