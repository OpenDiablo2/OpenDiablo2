package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// CreateItemPacket contains the data required to create a Item entity
type SpawnItemPacket struct {
	X     int      `json:"x"`
	Y     int      `json:"y"`
	Codes []string `json:"codes"`
}

// CreateSpawnItemPacket returns a NetPacket which declares a
// SpawnItemPacket with the data in given parameters.
func CreateSpawnItemPacket(x, y int, codes ...string) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.SpawnItem,
		PacketData: SpawnItemPacket{
			X:     x,
			Y:     y,
			Codes: codes,
		},
	}
}
