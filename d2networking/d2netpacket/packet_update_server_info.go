package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

// UpdateServerInfoPacket contains the ID for a player and the map seed.
// It is sent by the server to synchronise these values on the client.
type UpdateServerInfoPacket struct {
	Seed     int64  `json:"seed"`
	PlayerId string `json:"playerId"`
}

// CreateUpdateServerInfoPacket returns a NetPacket which declares an
// UpdateServerInfoPacket with the given player ID and map seed.
func CreateUpdateServerInfoPacket(seed int64, playerId string) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.UpdateServerInfo,
		PacketData: UpdateServerInfoPacket{
			Seed:     seed,
			PlayerId: playerId,
		},
	}
}
