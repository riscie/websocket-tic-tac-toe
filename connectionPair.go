package main

import (
	"log"
	"sync"
	"time"
)

type connectionPair struct {
	// the mutex to protect connections
	connectionsMx sync.RWMutex

	// Registered connections.
	connections map[*connection]struct{}

	// Inbound messages from the connections.
	broadcast chan bool

	logMx sync.RWMutex
	log   [][]byte

	gs gameState
}

func newConnectionPair() *connectionPair {
	cp := &connectionPair{
		connectionsMx: sync.RWMutex{},
		broadcast:     make(chan bool),
		connections:   make(map[*connection]struct{}),
		gs:            newGameState(),
	}

	go func() {
		for {
			msg := <-cp.broadcast
			cp.connectionsMx.RLock()
			for c := range cp.connections {
				select {
				case c.send <- msg:
				// stop trying to send to this connection after trying for 1 second.
				// if we have to stop, it means that a reader died so remove the connection also.
				case <-time.After(1 * time.Second):
					log.Printf("shutting down connection %s", c)
					cp.removeConnection(c)
				}
			}
			cp.connectionsMx.RUnlock()
		}
	}()
	return cp
}

func (h *connectionPair) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	h.connections[conn] = struct{}{}
}

func (h *connectionPair) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.send)
	}
}
