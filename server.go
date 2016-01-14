package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		h1 := newHub() //Creating the first hub here which holds the first websocket connection pair
		h2 := newHub() //Creating the first hub here which holds the first websocket connection pair
		hubList := []*hub{}
		hubList = append(hubList, h1,h2)
		fmt.Println("hub", hubList)
	*/
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./webroot")))
	router.Handle("/ws", wsHandler{})

	log.Printf("serving tic-tac-toe on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
