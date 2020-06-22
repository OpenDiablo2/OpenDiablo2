package d2server

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"log"
	"sync"
	"time"
)

type ConnectionManager struct {
	sync.RWMutex
	Retries    int
	interval   time.Duration
	gameServer *GameServer
	status     map[string]int
}

func CreateConnectionManager(gameServer *GameServer) *ConnectionManager {
	manager := &ConnectionManager{
		Retries:    3,
		interval:   time.Millisecond * 1000,
		gameServer: gameServer,
		status:     make(map[string]int),
	}

	go manager.Run()

	return manager
}

func (c *ConnectionManager) Run() {
	log.Print("Starting connection manager...")
	for {
		for id, connection := range c.gameServer.clientConnections {
			if connection.GetConnectionType() != d2clientconnectiontype.Local {
				if err := connection.SendPacketToClient(d2netpacket.CreatePingPacket()); err != nil {
					log.Printf("Cannot ping client id: %s", id)
				}
				c.RWMutex.Lock()
				c.status[id] += 1

				if c.status[id] >= c.Retries {
					delete(c.status, id)
					c.Drop(id)
				}

				c.RWMutex.Unlock()
			}
		}

		time.Sleep(c.interval)
	}
}

func (c *ConnectionManager) Recv(id string) {
	c.status[id] = 0
}

func (c *ConnectionManager) Drop(id string) {
	c.gameServer.RWMutex.Lock()
	defer c.gameServer.RWMutex.Unlock()
	delete(c.gameServer.clientConnections, id)
	log.Printf("%s has been disconnected...", id)
}
