package d2remoteclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/google/uuid"
)

// RemoteClientConnection is the implementation of ClientConnection
// for a remote client.
type RemoteClientConnection struct {
	asset          *d2asset.AssetManager
	heroState      *d2hero.HeroStateFactory
	clientListener d2networking.ClientListener // The GameClient
	uniqueID       string                      // Unique ID generated on construction
	tcpConnection  *net.TCPConn                // UDP connection to the server
	active         bool                        // The connection is currently open
}

// Create constructs a new RemoteClientConnection
// and returns a pointer to it.
func Create(asset *d2asset.AssetManager) (*RemoteClientConnection, error) {
	heroStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	result := &RemoteClientConnection{
		heroState: heroStateFactory,
		uniqueID:  uuid.New().String(),
	}

	return result, nil
}

// Open runs serverListener() in a goroutine to continuously read UDP packets.
// It also sends a PlayerConnectionRequestPacket packet to the server (see d2netpacket).
func (r *RemoteClientConnection) Open(connectionString, saveFilePath string) error {
	if !strings.Contains(connectionString, ":") {
		connectionString += ":6669"
	}

	// TODO: Connect to the server
	tcpAddress, err := net.ResolveTCPAddr("tcp", connectionString)

	// TODO: Show connection error screen if connection fails
	if err != nil {
		return err
	}

	r.tcpConnection, err = net.DialTCP("tcp", nil, tcpAddress)
	// TODO: Show connection error screen if connection fails
	if err != nil {
		return err
	}

	r.active = true
	go r.serverListener()

	log.Printf("Connected to server at %s", r.tcpConnection.RemoteAddr().String())

	gameState := r.heroState.LoadHeroState(saveFilePath)
	packet := d2netpacket.CreatePlayerConnectionRequestPacket(r.GetUniqueID(), gameState)
	err = r.SendPacketToServer(packet)

	if err != nil {
		log.Print("RemoteClientConnection: error sending PlayerConnectionRequestPacket to server.")
		return err
	}

	return nil
}

// Close informs the server that this client has disconnected and sets
// RemoteClientConnection.active to false.
func (r *RemoteClientConnection) Close() error {
	r.active = false
	err := r.SendPacketToServer(d2netpacket.CreatePlayerDisconnectRequestPacket(r.GetUniqueID()))

	if err != nil {
		return err
	}

	return nil
}

// GetUniqueID returns RemoteClientConnection.uniqueID.
func (r RemoteClientConnection) GetUniqueID() string {
	return r.uniqueID
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype
func (r RemoteClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.LANClient
}

// SetClientListener sets RemoteClientConnection.clientListener to the given value.
func (r *RemoteClientConnection) SetClientListener(listener d2networking.ClientListener) {
	r.clientListener = listener
}

// SendPacketToServer compresses the JSON encoding of a NetPacket and
// sends it to the server.
func (r *RemoteClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	data, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	if _, err = r.tcpConnection.Write(data); err != nil {
		return err
	}

	return nil
}

// serverListener runs a while loop, reading from the GameServer's TCP
// connection.
func (r *RemoteClientConnection) serverListener() {
	var packet d2netpacket.NetPacket

	decoder := json.NewDecoder(r.tcpConnection)

	for {
		err := decoder.Decode(&packet)
		if err != nil {
			log.Printf("failed to decode the packet, err: %v\n", err)
			return
		}

		p, err := r.decodeToPacket(packet.PacketType, string(packet.PacketData))
		if err != nil {
			log.Println(packet.PacketType, err)
		}

		err = r.clientListener.OnPacketReceived(p)
		if err != nil {
			log.Println(packet.PacketType, err)
		}
	}
}

// bytesToJSON reads the packet type, decompresses the packet and returns a JSON string.
func (r *RemoteClientConnection) bytesToJSON(buffer []byte) (string, d2netpackettype.NetPacketType, error) {
	packet, err := d2netpacket.UnmarshalNetPacket(buffer)
	if err != nil {
		return "", 0, err
	}

	return string(packet.PacketData), packet.PacketType, nil
}

// decodeToPacket unmarshals the JSON string into the correct struct
// and returns a NetPacket declaring that struct.
func (r *RemoteClientConnection) decodeToPacket(t d2netpackettype.NetPacketType, data string) (d2netpacket.NetPacket, error) {
	var np = d2netpacket.NetPacket{}

	var err error

	switch t {
	case d2netpackettype.GenerateMap:
		var p d2netpacket.GenerateMapPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: d2netpacket.MarshalPacket(p)}

	case d2netpackettype.MovePlayer:
		var p d2netpacket.MovePlayerPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: d2netpacket.MarshalPacket(p)}

	case d2netpackettype.UpdateServerInfo:
		var p d2netpacket.UpdateServerInfoPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: d2netpacket.MarshalPacket(p)}

	case d2netpackettype.AddPlayer:
		var p d2netpacket.AddPlayerPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: d2netpacket.MarshalPacket(p)}

	case d2netpackettype.Ping:
		var p d2netpacket.PingPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: d2netpacket.MarshalPacket(p)}

	case d2netpackettype.PlayerDisconnectionNotification:
		var p d2netpacket.PlayerDisconnectRequestPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: d2netpacket.MarshalPacket(p)}

	default:
		err = fmt.Errorf("RemoteClientConnection: unrecognized packet type: %v", t)
	}

	if err != nil {
		return np, err
	}

	return np, nil
}
