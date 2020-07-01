// Package d2remoteclient facilitates communication between a remote client and server.
package d2remoteclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	uuid "github.com/satori/go.uuid"
)

// RemoteClientConnection is the implementation of ClientConnection
// for a remote client.
type RemoteClientConnection struct {
	clientListener d2networking.ClientListener // The GameClient
	uniqueID       string                      // Unique ID generated on construction
	udpConnection  *net.UDPConn                // UDP connection to the server
	active         bool                        // The connection is currently open
}

// Create constructs a new RemoteClientConnection
// and returns a pointer to it.
func Create() *RemoteClientConnection {
	result := &RemoteClientConnection{
		uniqueID: uuid.NewV4().String(),
	}

	return result
}

// Open runs serverListener() in a goroutine to continuously read UDP packets.
// It also sends a PlayerConnectionRequestPacket packet to the server (see d2netpacket).
func (l *RemoteClientConnection) Open(connectionString, saveFilePath string) error {
	if !strings.Contains(connectionString, ":") {
		connectionString += ":6669"
	}

	// TODO: Connect to the server
	udpAddress, err := net.ResolveUDPAddr("udp", connectionString)

	// TODO: Show connection error screen if connection fails
	if err != nil {
		return err
	}

	l.udpConnection, err = net.DialUDP("udp", nil, udpAddress)
	// TODO: Show connection error screen if connection fails
	if err != nil {
		return err
	}

	l.active = true
	go l.serverListener()

	log.Printf("Connected to server at %s", l.udpConnection.RemoteAddr().String())

	gameState := d2player.LoadPlayerState(saveFilePath)
	err = l.SendPacketToServer(d2netpacket.CreatePlayerConnectionRequestPacket(l.GetUniqueID(), gameState))

	if err != nil {
		log.Print("RemoteClientConnection: error sending PlayerConnectionRequestPacket to server.")
		return err
	}

	return nil
}

// Close informs the server that this client has disconnected and sets
// RemoteClientConnection.active to false.
func (l *RemoteClientConnection) Close() error {
	l.active = false
	err := l.SendPacketToServer(d2netpacket.CreatePlayerDisconnectRequestPacket(l.GetUniqueID()))

	if err != nil {
		return err
	}

	return nil
}

// GetUniqueID returns RemoteClientConnection.uniqueID.
func (l RemoteClientConnection) GetUniqueID() string {
	return l.uniqueID
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype
func (l RemoteClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.LANClient
}

// SetClientListener sets RemoteClientConnection.clientListener to the given value.
func (l *RemoteClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}

// SendPacketToClient passes a packet to the game client for processing.
func (l *RemoteClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	return l.clientListener.OnPacketReceived(packet)
}

// SendPacketToServer compresses the JSON encoding of a NetPacket and
// sends it to the server.
func (l *RemoteClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	data, err := json.Marshal(packet.PacketData)
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	buff.WriteByte(byte(packet.PacketType))
	writer, _ := gzip.NewWriterLevel(&buff, gzip.BestCompression)
	var written int

	if written, err = writer.Write(data); err != nil {
		return err
	} else if written == 0 {
		return fmt.Errorf("remoteClientConnection: attempted to send empty %v packet body", packet.PacketType)
	}
	if err := writer.Close(); err != nil {
		return err
	}

	if _, err = l.udpConnection.Write(buff.Bytes()); err != nil {
		return err
	}

	return nil
}

// serverListener runs a while loop, reading from the GameServer's UDP
// connection.
func (l *RemoteClientConnection) serverListener() {
	buffer := make([]byte, 4096)

	for l.active {
		n, _, err := l.udpConnection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Socket error: %s\n", err)
			continue
		}

		if n <= 0 {
			continue
		}

		// Bytes to JSON
		data, packetType, err := l.readPacket(buffer)
		if err != nil {
			log.Println("RemoteClientConnection: error", packetType, err)
		}

		// JSON to struct
		packet, err := l.decodeToPacket(packetType, data)
		if err != nil {
			log.Println("RemoteClientConnection: error", packetType, err)
		}

		// Pass struct to GameClient
		if err := l.SendPacketToClient(packet); err != nil {
			log.Println("RemoteClientConnection: error", packetType, err)
		}
	}
}

func (l *RemoteClientConnection) readPacket(buffer []byte) (string, d2netpackettype.NetPacketType, error) {
	buff := bytes.NewBuffer(buffer)

	packetTypeID, err := buff.ReadByte()
	if err != nil {
		// The packet type here will be UpdateServerInfo. That shouldn't matter
		// but perhaps we should have a 'None' packet type anyway.
		return "", d2netpackettype.NetPacketType(0), fmt.Errorf("error reading packet type: %s", err)
	}

	packetType := d2netpackettype.NetPacketType(packetTypeID)
	reader, err := gzip.NewReader(buff)
	if err != nil {
		return "", packetType, fmt.Errorf("error creating reader for %v packet: %s", packetType, err)
	}

	sb := new(strings.Builder)

	// This will throw errors where packets are not compressed. This doesn't
	// break anything, so the gzip.ErrHeader error, is currently ignored to
	// avoid noisy logging.
	written, _ := io.Copy(sb, reader)

	if err != nil && err != gzip.ErrHeader {
		return "", packetType, fmt.Errorf("error copying bytes from %v packet: %s", packetType, err)
	}

	if written == 0 {
		return "", packetType, fmt.Errorf("empty %v packet received", packetType)
	}

	return sb.String(), packetType, nil
}

func (l *RemoteClientConnection) decodeToPacket(t d2netpackettype.NetPacketType, data string) (d2netpacket.NetPacket, error) {
	var np = d2netpacket.NetPacket{}

	var err error

	switch t {
	case d2netpackettype.GenerateMap:
		var p d2netpacket.GenerateMapPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: p}

	case d2netpackettype.MovePlayer:
		var p d2netpacket.MovePlayerPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: p}

	case d2netpackettype.UpdateServerInfo:
		var p d2netpacket.UpdateServerInfoPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: p}

	case d2netpackettype.AddPlayer:
		var p d2netpacket.AddPlayerPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: p}

	case d2netpackettype.Ping:
		var p d2netpacket.PingPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: p}

	case d2netpackettype.PlayerDisconnectionNotification:
		var p d2netpacket.PlayerDisconnectRequestPacket
		if err = json.Unmarshal([]byte(data), &p); err != nil {
			break
		}

		np = d2netpacket.NetPacket{PacketType: t, PacketData: p}

	default:
		err = fmt.Errorf("RemoteClientConnection: unrecognized packet type: %v", t)
	}

	if err != nil {
		return np, err
	}

	return np, nil
}
