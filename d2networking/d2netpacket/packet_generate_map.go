package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// GenerateMapPacket contains an enumerable representing a region. It
// is sent by the server to generate the map for the given region on
// a client.
type GenerateMapPacket struct {
	RegionType d2enum.RegionIdType `json:"regionType"`
}

// CreateGenerateMapPacket returns a NetPacket which declares a
// GenerateMapPacket with the given regionType.
func CreateGenerateMapPacket(regionType d2enum.RegionIdType) NetPacket {
	generateMapPacket := GenerateMapPacket{
		RegionType: regionType,
	}
	b, _ := json.Marshal(generateMapPacket)

	return NetPacket{
		PacketType: d2netpackettype.GenerateMap,
		PacketData: b,
	}
}

func UnmarshalGenerateMap(packet []byte) (GenerateMapPacket, error) {
	var p GenerateMapPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
