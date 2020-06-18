package d2remoteclient

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"log"
	"net"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	uuid "github.com/satori/go.uuid"
)

type RemoteClientConnection struct {
	clientListener d2networking.ClientListener
	uniqueId       string
	udpConnection  *net.UDPConn
}

func (l RemoteClientConnection) GetUniqueId() string {
	return l.uniqueId
}

func (l RemoteClientConnection) GetConnectionType() string {
	return "Remote Client"
}

func (l *RemoteClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error { // WHAT IS THIS
	//return l.clientListener.OnPacketReceived(packet)
	return nil
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

	log.Printf("Connected to server at %s", l.udpConnection.RemoteAddr().String())
	gameState := d2gamestate.LoadGameState(saveFilePath)
	l.SendPacketToServer(d2netpacket.CreatePlayerConnectionRequestPacket(l.GetUniqueId(), gameState))

	return nil
}

func (l *RemoteClientConnection) Close() error {
	// TODO: Disconnect from the server - send a disconnect packet
	return nil
}

func (l *RemoteClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	data, err := json.Marshal(packet)
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	writer, _ := gzip.NewWriterLevel(&buff, gzip.BestCompression)
	writer.Write(data)
	writer.Flush()
	writer.Close()
	if _, err = l.udpConnection.Write(buff.Bytes()); err != nil {
		return err
	}
	return nil
}

func (l *RemoteClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}
