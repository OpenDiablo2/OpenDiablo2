package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type GenerateMapPacket struct {
	RegionType d2enum.RegionIdType `json:"regionType"`
}

func CreateGenerateMapPacket(regionType d2enum.RegionIdType) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.GenerateMap,
		PacketData: GenerateMapPacket{
			RegionType: regionType,
		},
	}

}
