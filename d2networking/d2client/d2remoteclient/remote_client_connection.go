package d2remoteclient

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/google/uuid"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

const logPrefix = "Remote Client"

// RemoteClientConnection is the implementation of ClientConnection
// for a remote client.
type RemoteClientConnection struct {
	asset          *d2asset.AssetManager
	heroState      *d2hero.HeroStateFactory
	clientListener d2networking.ClientListener // The GameClient
	uniqueID       string                      // Unique ID generated on construction
	tcpConnection  *net.TCPConn                // UDP connection to the server
	active         bool                        // The connection is currently open

	*d2util.Logger
}

// Create constructs a new RemoteClientConnection
// and returns a pointer to it.
func Create(l d2util.LogLevel, asset *d2asset.AssetManager) (*RemoteClientConnection, error) {
	heroStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	result := &RemoteClientConnection{
		asset:     asset,
		heroState: heroStateFactory,
		uniqueID:  uuid.New().String(),
	}

	result.Logger = d2util.NewLogger()
	result.Logger.SetPrefix(logPrefix)
	result.Logger.SetLevel(l)

	return result, nil
}

// Open runs serverListener() in a goroutine to continuously read UDP packets.
// It also sends a PlayerConnectionRequestPacket packet to the server (see d2netpacket).
func (r *RemoteClientConnection) Open(connectionString, saveFilePath string) error {
	if !strings.Contains(connectionString, ":") {
		connectionString += ":6669"
	}

	tcpAddress, err := net.ResolveTCPAddr("tcp", connectionString)

	if err != nil {
		return err
	}

	r.tcpConnection, err = net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		return err
	}

	r.active = true
	go r.serverListener()

	r.Infof("Connected to server at %s", r.tcpConnection.RemoteAddr().String())

	gameState := r.heroState.LoadHeroState(saveFilePath)

	packet, err := d2netpacket.CreatePlayerConnectionRequestPacket(r.GetUniqueID(), gameState)
	if err != nil {
		r.Errorf("PlayerConnectionRequestPacket: %v", err)
	}

	err = r.SendPacketToServer(packet)

	if err != nil {
		r.Errorf("RemoteClientConnection: error sending PlayerConnectionRequestPacket to server.")
		return err
	}

	return nil
}

// Close informs the server that this client has disconnected and sets
// RemoteClientConnection.active to false.
func (r *RemoteClientConnection) Close() error {
	r.active = false

	pd, err := d2netpacket.CreatePlayerDisconnectRequestPacket(r.GetUniqueID())
	if err != nil {
		return fmt.Errorf("PlayerDisconnectRequestPacket: %v", err)
	}

	err = r.SendPacketToServer(pd)

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
			r.Errorf("failed to decode the packet, err: %v\n", err)
			return
		}

		p, err := r.decodeToPacket(packet.PacketType, string(packet.PacketData))
		if err != nil {
			r.Errorf("%v %v", packet.PacketType, err)
		}

		err = r.clientListener.OnPacketReceived(p)
		if err != nil {
			r.Errorf("%v %v", packet.PacketType, err)
		}
	}
}

// bytesToJSON reads the packet type, decompresses the packet and returns a JSON string.
// nolint:unused // WIP
func (r *RemoteClientConnection) bytesToJSON(buffer []byte) (string, d2netpackettype.NetPacketType, error) {
	packet, err := d2netpacket.UnmarshalNetPacket(buffer)
	if err != nil {
		return "", 0, err
	}

	return string(packet.PacketData), packet.PacketType, nil
}

// decodeToPacket unmarshals the JSON string into the correct struct
// and returns a NetPacket declaring that struct.
// nolint:gocyclo,funlen // switch statement on packet type makes sense, no need to change
func (r *RemoteClientConnection) decodeToPacket(
	t d2netpackettype.NetPacketType,
	data string) (d2netpacket.NetPacket, error) {
	var np = d2netpacket.NetPacket{}

	var err error

	switch t {
	case d2netpackettype.GenerateMap:
		var p d2netpacket.GenerateMapPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	case d2netpackettype.MovePlayer:
		var p d2netpacket.MovePlayerPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	case d2netpackettype.UpdateServerInfo:
		var p d2netpacket.UpdateServerInfoPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	case d2netpackettype.AddPlayer:
		var p d2netpacket.AddPlayerPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	case d2netpackettype.CastSkill:
		var p d2netpacket.CastPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	case d2netpackettype.Ping:
		var p d2netpacket.PingPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	case d2netpackettype.PlayerDisconnectionNotification:
		var p d2netpacket.PlayerDisconnectRequestPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		mp, marshalErr := d2netpacket.MarshalPacket(p)
		if marshalErr != nil {
			r.Errorf("MarshalPacket: %v", marshalErr)
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: mp}

	default:
		err = fmt.Errorf("RemoteClientConnection: unrecognized packet type: %v", t)
	}

	if err != nil {
		return np, err
	}

	return np, nil
}
