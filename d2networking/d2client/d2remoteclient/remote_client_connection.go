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

	unmarshalErr := "RemoteClientConnection: error unmarshalling %T: %s"
	sendToClientErr := "RemoteClientConnection: error processing packet %v: %s"

	for l.active {
		n, _, err := l.udpConnection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Socket error: %s\n", err)
			continue
		}

		if n <= 0 {
			continue
		}

		buff := bytes.NewBuffer(buffer)

		packetTypeID, err := buff.ReadByte()
		if err != nil {
			log.Printf("RemoteClientConnection: error reading packet type: %s", err)
		}

		packetType := d2netpackettype.NetPacketType(packetTypeID)
		reader, err := gzip.NewReader(buff)
		if err != nil {
			log.Printf("RemoteClientConnection: error reading packet type: %s", err)
		}

		sb := new(strings.Builder)

		// This will throw errors where packets are not compressed. This doesn't
		// break anything, so the gzip.ErrHeader error, is currently ignored to
		// avoid noisy logging.
		written, _ := io.Copy(sb, reader)

		if err != nil && err != gzip.ErrHeader {
			log.Printf("RemoteClientConnection: error copying bytes from %v packet: %s", packetType, err)
		}

		if written == 0 {
			log.Printf("RemoteClientConnection: empty packet %v packet received", packetType)
			continue
		}

		data := sb.String()

		switch packetType {
		case d2netpackettype.GenerateMap:
			var packet d2netpacket.GenerateMapPacket
			if err := json.Unmarshal([]byte(data), &packet); err != nil {
				log.Printf(unmarshalErr, packetType, err)
				continue
			}

			p := d2netpacket.NetPacket{PacketType: packetType, PacketData: packet}
			if err := l.SendPacketToClient(p); err != nil {
				log.Printf(sendToClientErr, packetType, err)
			}

		case d2netpackettype.MovePlayer:
			var packet d2netpacket.MovePlayerPacket
			if err := json.Unmarshal([]byte(data), &packet); err != nil {
				log.Printf(unmarshalErr, packetType, err)
				continue
			}

			p := d2netpacket.NetPacket{PacketType: packetType, PacketData: packet}
			if err := l.SendPacketToClient(p); err != nil {
				log.Printf(sendToClientErr, packetType, err)
			}

		case d2netpackettype.UpdateServerInfo:
			var packet d2netpacket.UpdateServerInfoPacket
			if err := json.Unmarshal([]byte(data), &packet); err != nil {
				log.Printf(unmarshalErr, packetType, err)
				continue
			}

			p := d2netpacket.NetPacket{PacketType: packetType, PacketData: packet}
			if err := l.SendPacketToClient(p); err != nil {
				log.Printf(sendToClientErr, packetType, err)
			}

		case d2netpackettype.AddPlayer:
			var packet d2netpacket.AddPlayerPacket
			if err := json.Unmarshal([]byte(data), &packet); err != nil {
				log.Printf(unmarshalErr, packetType, err)
				continue
			}

			p := d2netpacket.NetPacket{PacketType: packetType, PacketData: packet}
			if err := l.SendPacketToClient(p); err != nil {
				log.Printf(sendToClientErr, packetType, err)
			}

		case d2netpackettype.Ping:
			err := l.SendPacketToServer(d2netpacket.CreatePongPacket(l.uniqueID))
			if err != nil {
				log.Printf(sendToClientErr, packetType, err)
			}

		case d2netpackettype.PlayerDisconnectionNotification:
			var packet d2netpacket.PlayerDisconnectRequestPacket
			if err := json.Unmarshal([]byte(data), &packet); err != nil {
				log.Printf(unmarshalErr, packetType, err)
				continue
			}

			log.Printf("RemoteClientConnection: received disconnect: %s", packet.Id)

		default:
			fmt.Printf("RemoteClientConnection: unknown packet type %v", packetType)
		}
	}
}
