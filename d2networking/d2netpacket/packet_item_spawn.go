package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// SpawnItemPacket contains the data required to create a Item entity
type SpawnItemPacket struct {
	X     int      `json:"x"`
	Y     int      `json:"y"`
	Codes []string `json:"codes"`
}

// CreateSpawnItemPacket returns a NetPacket which declares a
// SpawnItemPacket with the data in given parameters.
func CreateSpawnItemPacket(x, y int, codes ...string) NetPacket {
	spawnItemPacket := SpawnItemPacket{
		X:     x,
		Y:     y,
		Codes: codes,
	}
	b, _ := json.Marshal(spawnItemPacket)

	return NetPacket{
		PacketType: d2netpackettype.SpawnItem,
		PacketData: b,
	}
}

func UnmarshalSpawnItem(packet []byte) (SpawnItemPacket, error) {
	var p SpawnItemPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
