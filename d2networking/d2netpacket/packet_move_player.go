package d2netpacket

import (
	"encoding/json"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// MovePlayerPacket contains a movement command for a specific player entity.
// It is sent by the server to move a player entity on a client.
// TODO: Need to handle being on different maps
type MovePlayerPacket struct {
	PlayerID string  `json:"playerId"`
	StartX   float64 `json:"startX"`
	StartY   float64 `json:"startY"`
	DestX    float64 `json:"destX"`
	DestY    float64 `json:"destY"`
}

// CreateMovePlayerPacket returns a NetPacket which declares a MovePlayerPacket
// with the given ID and movement command.
func CreateMovePlayerPacket(playerID string, startX, startY, destX, destY float64) NetPacket {
	movePlayerPacket := MovePlayerPacket{
		PlayerID: playerID,
		StartX:   startX,
		StartY:   startY,
		DestX:    destX,
		DestY:    destY,
	}
	b, err := json.Marshal(movePlayerPacket)
	if err != nil {
		log.Print(err)
	}

	return NetPacket{
		PacketType: d2netpackettype.MovePlayer,
		PacketData: b,
	}
}

func UnmarshalMovePlayer(packet []byte) (MovePlayerPacket, error) {
	var p MovePlayerPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
