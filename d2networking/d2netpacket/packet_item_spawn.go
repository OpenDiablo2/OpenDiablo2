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
func CreateSpawnItemPacket(x, y int, codes ...string) (NetPacket, error) {
	spawnItemPacket := SpawnItemPacket{
		X:     x,
		Y:     y,
		Codes: codes,
	}

	b, err := json.Marshal(spawnItemPacket)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.SpawnItem}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.SpawnItem,
		PacketData: b,
	}, nil
}

// UnmarshalSpawnItem unmarshals the given data to a SpawnItemPacket struct
func UnmarshalSpawnItem(packet []byte) (SpawnItemPacket, error) {
	var p SpawnItemPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
