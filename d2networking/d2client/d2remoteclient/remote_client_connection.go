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

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	uuid "github.com/satori/go.uuid"
)

type RemoteClientConnection struct {
	clientListener d2networking.ClientListener
	uniqueId       string
	udpConnection  *net.UDPConn
	active         bool
}

func (l RemoteClientConnection) GetUniqueId() string {
	return l.uniqueId
}

func (l RemoteClientConnection) GetConnectionType() string {
	return "Remote Client"
}

func (l *RemoteClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error { // WHAT IS THIS
	return l.clientListener.OnPacketReceived(packet)
}

func Create() *RemoteClientConnection {
	result := &RemoteClientConnection{
		uniqueId: uuid.NewV4().String(),
	}

	return result
}

func (l *RemoteClientConnection) Open(connectionString string, saveFilePath string) error {
	if strings.Index(connectionString, ":") == -1 {
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
	gameState := d2gamestate.LoadGameState(saveFilePath)
	l.SendPacketToServer(d2netpacket.CreatePlayerConnectionRequestPacket(l.GetUniqueId(), gameState))

	return nil
}

func (l *RemoteClientConnection) Close() error {
	l.active = false
	// TODO: Disconnect from the server - send a disconnect packet
	return nil
}

func (l *RemoteClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	data, err := json.Marshal(packet.PacketData)
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	buff.WriteByte(byte(packet.PacketType))
	writer, _ := gzip.NewWriterLevel(&buff, gzip.BestCompression)
	writer.Write(data)
	writer.Close()
	if _, err = l.udpConnection.Write(buff.Bytes()); err != nil {
		return err
	}
	return nil
}

func (l *RemoteClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}

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
		buff := bytes.NewBuffer(buffer)
		packetTypeId, err := buff.ReadByte()
		packetType := d2netpackettype.NetPacketType(packetTypeId)
		reader, err := gzip.NewReader(buff)
		sb := new(strings.Builder)
		io.Copy(sb, reader)
		stringData := sb.String()
		switch packetType {
		case d2netpackettype.GenerateMap:
			var packet d2netpacket.GenerateMapPacket
			json.Unmarshal([]byte(stringData), &packet)
			l.SendPacketToClient(d2netpacket.NetPacket{
				PacketType: packetType,
				PacketData: packet,
			})
		case d2netpackettype.MovePlayer:
			var packet d2netpacket.MovePlayerPacket
			json.Unmarshal([]byte(stringData), &packet)
			l.SendPacketToClient(d2netpacket.NetPacket{
				PacketType: packetType,
				PacketData: packet,
			})
		case d2netpackettype.UpdateServerInfo:
			var packet d2netpacket.UpdateServerInfoPacket
			json.Unmarshal([]byte(stringData), &packet)
			l.SendPacketToClient(d2netpacket.NetPacket{
				PacketType: packetType,
				PacketData: packet,
			})
		case d2netpackettype.AddPlayer:
			var packet d2netpacket.AddPlayerPacket
			json.Unmarshal([]byte(stringData), &packet)
			l.SendPacketToClient(d2netpacket.NetPacket{
				PacketType: packetType,
				PacketData: packet,
			})
		default:
			fmt.Printf("Unknown packet type %d\n", packetType)
		}

	}
}
