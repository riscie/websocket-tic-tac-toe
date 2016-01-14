package main

import (
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
	broadcast chan bool
	// The connectionPair. Holds up to 2 connections.
	cp        *connectionPair
	// playerNum represents the players Slot. Either 0 or 1
	playerNum int
}

//wsHandler implements the Handler Interface
type wsHandler struct{}

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for {
		//Reading next move from connection here
		_, clientMoveMessage, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		field, _ := strconv.ParseInt(string(clientMoveMessage[:]), 10, 32) //Getting FieldValue From Player Action
		c.cp.gs.makeMove(c.playerNum, int(field))
		c.cp.updateFromClient <- true //telling cp to broadcast the gamestate
	}
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for range c.broadcast {
		sendGameStateToConnection(wsConn, c)
	}
}

func getConnectionPairWithEmptySlot() (*connectionPair, int) {
	sizeBefore := len(connections)
	for _, h := range connections {
		if len(h.connections) <= 1 {
			return h, len(h.connections)
		}
	}
	h := newConnectionPair()
	connections = append(connections, h)
	return connections[sizeBefore], 0
}

func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Upgrading HTTP Connection to websocket connection
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading %s", err)
		return
	}

	//Adding Connection to connectionpair
	cp, pn := getConnectionPairWithEmptySlot()
	c := &connection{broadcast: make(chan bool), cp: cp, playerNum: pn}
	c.cp.addConnection(c)
	defer c.cp.removeConnection(c)

	c.cp.gs.NumberOfPlayers++
	if c.cp.gs.NumberOfPlayers == 2 {
		c.cp.gs.StatusMessage = "starting game"
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
	err := wsConn.WriteMessage(websocket.TextMessage, stateToJSON(c.cp.gs))
	//removing connection if updating gamestate fails
	if err != nil {
		c.cp.removeConnection(c)
	}
}
