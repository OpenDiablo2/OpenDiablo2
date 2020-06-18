package d2udpclientconnection

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type UDPClientConnection struct {
	id            string
	address       *net.UDPAddr
	udpConnection *net.UDPConn
	playerState   *d2player.PlayerState
}

func CreateUDPClientConnection(udpConnection *net.UDPConn, id string, address *net.UDPAddr) *UDPClientConnection {
	result := &UDPClientConnection{
		id:            id,
		address:       address,
		udpConnection: udpConnection,
	}

	return result
}

func (u UDPClientConnection) GetUniqueId() string {
	return u.id
}

func (u UDPClientConnection) GetConnectionType() string {
	return "Remote Client"
}

func (u *UDPClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	data, err := json.Marshal(packet.PacketData)
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	buff.WriteByte(byte(packet.PacketType))
	writer, _ := gzip.NewWriterLevel(&buff, gzip.BestCompression)
	writer.Write(data)
	writer.Close()
	u.udpConnection.WriteToUDP(buff.Bytes(), u.address)

	return nil
}

func (u *UDPClientConnection) SetPlayerState(playerState *d2player.PlayerState) {
	u.playerState = playerState
}

func (u *UDPClientConnection) GetPlayerState() *d2player.PlayerState {
	return u.playerState
}
