package main

import (
	"log"
	"net/http"
)

func main() {
	h := newHub()

	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./webroot")))
	router.Handle("/ws", wsHandler{hub: h})

	log.Printf("serving tic-tac-toe on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
