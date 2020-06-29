package d2netpacket

import (
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
	return NetPacket{
		PacketType: d2netpackettype.GenerateMap,
		PacketData: GenerateMapPacket{
			RegionType: regionType,
		},
	}

}
