package d2server

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"log"
	"sync"
	"time"
)

// ConnectionManager is responsible for cleanup up connections accepted by the game server. As the server communicates over
// UDP and is stateless we need to implement some loose state management via a ping/pong system. ConnectionManager also handles
// communication for graceful shutdowns.
//
// retries: # of attempts before the dropping the client
// interval: How long to wait before each ping/pong test
// gameServer: The *GameServer is argument provided for the connection manager to watch over
// status: map of inflight ping/pong requests
type ConnectionManager struct {
	sync.RWMutex
	retries    int
	interval   time.Duration
	gameServer *GameServer
	status     map[string]int
}

func CreateConnectionManager(gameServer *GameServer) *ConnectionManager {
	manager := &ConnectionManager{
		retries:    3,
		interval:   time.Millisecond * 1000,
		gameServer: gameServer,
		status:     make(map[string]int),
	}

	go manager.Run()

	return manager
}

// Run starts up any watchers for for the connection manager
func (c *ConnectionManager) Run() {
	log.Print("Starting connection manager...")
	for {
		c.checkPeers()
		time.Sleep(c.interval)
	}
}

// checkPeers manages connection validation and cleanup for all peers.
func (c *ConnectionManager) checkPeers() {
	for id, connection := range c.gameServer.clientConnections {
		if connection.GetConnectionType() != d2clientconnectiontype.Local {
			if err := connection.SendPacketToClient(d2netpacket.CreatePingPacket()); err != nil {
				log.Printf("Cannot ping client id: %s", id)
			}
			c.RWMutex.Lock()
			c.status[id] += 1

			if c.status[id] >= c.retries {
				delete(c.status, id)
				c.Drop(id)
			}

			c.RWMutex.Unlock()
		}
	}
}

// Recv simply resets the counter, acknowledging we have received a pong from the client.
func (c *ConnectionManager) Recv(id string) {
	c.status[id] = 0
}

// Drop removes the client id from the connection pool of the game server.
func (c *ConnectionManager) Drop(id string) {
	c.gameServer.RWMutex.Lock()
	defer c.gameServer.RWMutex.Unlock()
	delete(c.gameServer.clientConnections, id)
	log.Printf("%s has been disconnected...", id)
}

// Shutdown will notify all of the clients that the server has been shutdown.
func (c *ConnectionManager) Shutdown() {
	// TODO: Currently this will never actually get called as the go routines are never signaled about the application termination.
	// Things can be done more cleanly once we have graceful exits however we still need to account for other OS Signals
	log.Print("Notifying clients server is shutting down...")
	for _, connection := range c.gameServer.clientConnections {
		connection.SendPacketToClient(d2netpacket.CreateServerClosedPacket())
	}
	Stop()
}
