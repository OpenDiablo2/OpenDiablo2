package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// AddPlayerPacket contains the data required to create a Player entity.
// It is sent by the server to create the entity for a newly connected
// player on a client.
type AddPlayerPacket struct {
	ID        string                         `json:"id"`
	Name      string                         `json:"name"`
	X         int                            `json:"x"`
	Y         int                            `json:"y"`
	HeroType  d2enum.Hero                    `json:"hero"`
	Equipment d2inventory.CharacterEquipment `json:"equipment"`
	Stats     *d2hero.HeroStatsState         `json:"heroStats"`
}

// CreateAddPlayerPacket returns a NetPacket which declares an
// AddPlayerPacket with the data in given parameters.
func CreateAddPlayerPacket(id, name string, x, y int, heroType d2enum.Hero,
	stats *d2hero.HeroStatsState, equipment d2inventory.CharacterEquipment) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.AddPlayer,
		PacketData: AddPlayerPacket{
			ID:        id,
			Name:      name,
			X:         x,
			Y:         y,
			HeroType:  heroType,
			Equipment: equipment,
			Stats:     stats,
		},
	}
}
