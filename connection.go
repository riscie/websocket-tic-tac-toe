package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//upgrader is needed to upgrade the HTTP Connection to a websocket Connection
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, //TODO: Remove in production. Needed for gin proxy
}

//connections stores all the hubs
var connections []*connectionPair

type connection struct {
	// Channel which triggers the connection to update the gameState
	send chan bool
	// The hub.
	h *connectionPair
}

//wsHandler implements the Handler Interface
type wsHandler struct{}

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		//Reading next move from connection here
		_, chosenField, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fn, _ := strconv.ParseUint(string(chosenField[:]), 10, 32) //Getting FieldValue From Player Action

		c.h.gs.Fields[fn].Set = true
		c.h.gs.Fields[fn].Symbol = "X"

		c.h.broadcast <- true //telling hub to broadcast the gamestate
	}
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for range c.send {
		sendGameStateToConnection(wsConn, c)
	}
}

func getHubWithEmptySlot() *connectionPair {
	sizeBefore := len(connections)
	for i, h := range connections {
		fmt.Println("Hub", i, "has", len(h.connections), "connections")
		if len(h.connections) <= 1 {
			fmt.Println("returning Hub", i)
			return h
		}
	}
	h := newConnectionPair()
	connections = append(connections, h)
	fmt.Println("Hubs: ", len(connections), "Before: ", sizeBefore)
	return connections[sizeBefore]
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Upgrading HTTP Connection to websocket connection
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}

	//Adding Connection to hub
	c := &connection{send: make(chan bool), h: getHubWithEmptySlot()}
	c.h.addConnection(c)
	defer c.h.removeConnection(c)

	c.h.gs.NumberOfPlayers++
	if c.h.gs.NumberOfPlayers == 2 {
		c.h.gs.StatusMessage = "starting game"
	}

	//Sending initial Gamestate to client
	sendGameStateToConnection(wsConn, c)

	var wg sync.WaitGroup
	wg.Add(2)
	go c.writer(&wg, wsConn)
	go c.reader(&wg, wsConn)
	wg.Wait()
	wsConn.Close()
}

func sendGameStateToConnection(wsConn *websocket.Conn, c *connection) {
	err := wsConn.WriteMessage(websocket.TextMessage, stateToJSON(c.h.gs))
	//removing connection if updating gamestate fails
	if err != nil {
		c.h.removeConnection(c)
	}
}
