package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// ItemTypes stores all of the ItemTypeRecords
type ItemTypes map[string]*ItemTypeRecord

// ItemEquivalenceMap describes item equivalencies for ItemTypes
type ItemEquivalenceMap map[string]ItemEquivalenceList

// ItemEquivalenceList is an equivalence map that each ItemTypeRecord will have
type ItemEquivalenceList []*ItemCommonRecord

// ItemEquivalenceByRecord is used for getting equivalent item codes using an ItemCommonRecord
type ItemEquivalenceByRecord map[*ItemCommonRecord][]string

// ItemTypeRecord describes the types for items
type ItemTypeRecord struct {
	Name            string
	Code            string
	Equiv1          string
	Equiv2          string
	Shoots          string
	Quiver          string
	InvGfx1         string
	InvGfx2         string
	InvGfx3         string
	InvGfx4         string
	InvGfx5         string
	InvGfx6         string
	StorePage       string
	EquivalentItems ItemEquivalenceList
	BodyLoc2        int
	MaxSock1        int
	MaxSock25       int
	MaxSock40       int
	BodyLoc1        int
	Rarity          int
	StaffMods       d2enum.Hero
	CostFormula     int
	Class           d2enum.Hero
	VarInvGfx       int
	TreasureClass   int
	Repair          bool
	Throwable       bool
	Reload          bool
	ReEquip         bool
	AutoStack       bool
	Magic           bool
	Rare            bool
	Normal          bool
	Charm           bool
	Gem             bool
	Beltable        bool
	Body            bool
}
