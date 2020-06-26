package d2server

import (
	"encoding/json"
	"log"
	"net"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	packet "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	packettype "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	d2udp "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server/d2udpclientconnection"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

type GameServer struct {
	sync.RWMutex
	version           string
	clientConnections map[string]ClientConnection
	manager           *ConnectionManager
	realm             *d2mapengine.MapRealm
	scriptEngine      *d2script.ScriptEngine
	udpConnection     *net.UDPConn
	seed              int64
	running           bool
	maxClients        int
	lastAdvance       float64
}

func (srv *GameServer) handlePlayerConnRequest(addr *net.UDPAddr, data string) {
	packetData := &packet.PlayerConnectionRequestPacket{}
	json.Unmarshal([]byte(data), packetData)

	srvCon := srv.udpConnection
	packetId := packetData.Id
	clientCon := d2udp.CreateUDPClientConnection(srvCon, packetId, addr)

	state := packetData.PlayerState
	clientCon.SetPlayerState(state)
	OnClientConnected(clientCon)
}

func (srv *GameServer) handleMovePlayer(addr *net.UDPAddr, data string) {
	packetData := &packet.MovePlayerPacket{}
	json.Unmarshal([]byte(data), packetData)

	netPacket := packet.NetPacket{
		PacketType: packettype.MovePlayer,
		PacketData: packetData,
	}

	for _, player := range srv.clientConnections {
		player.SendPacketToClient(netPacket)
	}
}

func (srv *GameServer) handlePong(addr *net.UDPAddr, data string) {
	packetData := packet.PlayerConnectionRequestPacket{}
	json.Unmarshal([]byte(data), &packetData)
	srv.manager.Recv(packetData.Id)
}

func (srv *GameServer) handlePlayerDisconnectNotification(data string) {
	var packet packet.PlayerDisconnectRequestPacket
	json.Unmarshal([]byte(data), &packet)
	log.Printf("Received disconnect: %s", packet.Id)
}
