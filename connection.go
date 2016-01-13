package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type connection struct {
	// Buffered channel of outbound messages.
	send chan []byte
	// The hub.
	h *hub
}

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fn, _ := strconv.ParseUint(string(message[:]), 10, 32) //Getting FieldValue From Player Action
		fmt.Println("Reader: ", fn)
		gs.Fields[fn].Set = true
		gs.Fields[fn].Symbol = "X"
		c.h.broadcast <- message
	}
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for message := range c.send {
		err := wsConn.WriteMessage(websocket.TextMessage, stateToJSON())
		fmt.Println("Writer: ", string(message[:]))
		if err != nil {
			break
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, //TODO: Remove in production. Needed for gin proxy
}

type wsHandler struct {
	hub *hub
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}
	gs.NumberOfPlayers++
	if gs.NumberOfPlayers == 2 {
		gs.StatusMessage = "starting game"
	}
	err = wsConn.WriteMessage(websocket.TextMessage, stateToJSON())
	if err != nil {
		fmt.Println(err)
	}
	c := &connection{send: make(chan []byte, 256), h: wsh.hub}
	c.h.addConnection(c)
	defer c.h.removeConnection(c)
	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, wsConn)
	go c.reader(&wg, wsConn)
	wg.Wait()
	wsConn.Close()
}

/*
func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	state.AddPlayer()

	json, err := json.Marshal(state)
	if err != nil {
		log.Fatal("Error in marshalling json:", err)
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}

	fmt.Println("Number of Players:", state.NumberOfPlayers)
	err = wsConn.WriteMessage(websocket.TextMessage, json)

}
*/
