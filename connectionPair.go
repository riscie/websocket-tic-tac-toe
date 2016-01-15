package main

import (
	"sync"
	"time"
)

// connectionPair handles the update of the gameState between two players
type connectionPair struct {
	// the mutex to protect connections
	connectionsMx sync.RWMutex
	// Registered connections.
	connections map[*connection]struct{}
	// Inbound messages from the connections.
	shouldBroadcast chan bool
	logMx           sync.RWMutex
	log             [][]byte
	gs              gameState
}

// newConnectionPair is the constructor for the connectionPair struct
func newConnectionPair() *connectionPair {
	cp := &connectionPair{
		connectionsMx:   sync.RWMutex{},
		shouldBroadcast: make(chan bool),
		connections:     make(map[*connection]struct{}),
		gs:              newGameState(),
	}

	go func() {
		for {
			//waiting for an update of one of the clients in the connection pair
			<-cp.shouldBroadcast
			cp.connectionsMx.RLock()
			for c := range cp.connections {
				select {
				case c.doBroadcast <- true:
				// stop trying to send to this connection after trying for 1 second.
				// if we have to stop, it means that a reader died so remove the connection also.
				case <-time.After(1 * time.Second):
					cp.removeConnection(c)
				}
			}
			cp.connectionsMx.RUnlock()
		}
	}()
	return cp
}

// addConnection adds a players connection to the connectionPair
func (h *connectionPair) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	h.connections[conn] = struct{}{}
}

// removeConnection removes a players connection from the connectionPair
// TODO: Needs fixing. Connections are note removed atm.
func (h *connectionPair) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.doBroadcast)
	}
}
