package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// UpdateServerInfoPacket contains the ID for a player and the map seed.
// It is sent by the server to synchronize these values on the client.
type UpdateServerInfoPacket struct {
	Seed     int64  `json:"seed"`
	PlayerID string `json:"playerId"`
}

// CreateUpdateServerInfoPacket returns a NetPacket which declares an
// UpdateServerInfoPacket with the given player ID and map seed.
func CreateUpdateServerInfoPacket(seed int64, playerID string) (NetPacket, error) {
	updateServerInfo := UpdateServerInfoPacket{
		Seed:     seed,
		PlayerID: playerID,
	}

	b, err := json.Marshal(updateServerInfo)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.UpdateServerInfo}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.UpdateServerInfo,
		PacketData: b,
	}, nil
}

// UnmarshalUpdateServerInfo unmarshals the data to a UpdateServerInfoPacket struct
func UnmarshalUpdateServerInfo(packet []byte) (UpdateServerInfoPacket, error) {
	var resp UpdateServerInfoPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
