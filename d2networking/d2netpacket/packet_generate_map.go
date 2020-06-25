package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type GenerateMapPacket struct {
	ActId   int `json:"actId"`
	LevelId int `json:"levelId"`
}

func CreateGenerateMapPacket(actId, levelId int) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.GenerateMap,
		PacketData: GenerateMapPacket{
			ActId:   actId,
			LevelId: levelId,
		},
	}

}
