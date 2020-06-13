package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

type UpdateServerInfoPacket struct {
	Seed     int64  `json:"seed"`
	PlayerId string `json:"playerId"`
}

func CreateUpdateServerInfoPacket(seed int64, playerId string) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.UpdateServerInfo,
		PacketData: UpdateServerInfoPacket{
			Seed:     seed,
			PlayerId: playerId,
		},
	}
}
