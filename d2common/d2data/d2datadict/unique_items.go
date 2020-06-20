package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// UniqueItemRecord is a representation of a row from uniqueitems.txt
type UniqueItemRecord struct {
	Name    string
	Version int  // 0 = classic pre 1.07, 1 = 1.07-1.11, 100 = expansion
	Enabled bool // if false, this record won't be loaded (should always be true...)
	Ladder  bool // if true, can only be found on ladder and not in single player / tcp/ip
	Rarity  int  // 1-255, higher is more common (ironically...)
	NoLimit bool // if true, can drop more than once per game
	// (if false, can only drop once per game; if it would drop,
	//	 instead a rare item with enhanced durability drops)

	Level int // item's level, can only be dropped by monsters / recipes / vendors / objects of this level or higher
	// otherwise they would drop a rare item with enhanced durability
	RequiredLevel int    // character must have this level to use this item
	Code          string // three letter code, points to a record in Weapons, Armor, or Misc

	TypeDescription string
	UberDescription string
	SingleCopy      bool // if true, player can only hold one of these. can't trade it or pick it up
	CostMultiplier  int  // base price is multiplied by this when sold, repaired or bought
	CostAdd         int  // after multiplied by above, this much is added to the price

	CharacterGfxTransform string // palette shift applied to this items gfx when held and when
	// on the ground (flippy). Points to a record in Colors.txt
	InventoryGfxTransform string // palette shift applied to the inventory gfx
	FlippyFile            string // if non-empty, overrides the base item's dropped gfx
	InventoryFile         string // if non-empty, overrides the base item's inventory gfx

	DropSound    string // if non-empty, overrides the base item's drop sound
	DropSfxFrame int    // if non-empty, overrides the base item's frame at which the drop sound plays
	UseSound     string // if non-empty, overrides the sound played when item is used

	Properties [12]UniqueItemProperty
}

// UniqueItemProperty is describes a property of a unique item
type UniqueItemProperty struct {
	Property  string
	Parameter d2common.CalcString // depending on the property, this may be an int (usually), or a string
	Min       int
	Max       int
}

func createUniqueItemRecord(r []string) UniqueItemRecord {
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := UniqueItemRecord{
		Name:    r[inc()],
		Version: d2common.StringToInt(d2common.EmptyToZero(r[inc()])),
		Enabled: d2common.StringToInt(d2common.EmptyToZero(r[inc()])) == 1,

		Ladder:  d2common.StringToInt(d2common.EmptyToZero(r[inc()])) == 1,
		Rarity:  d2common.StringToInt(d2common.EmptyToZero(r[inc()])),
		NoLimit: d2common.StringToInt(d2common.EmptyToZero(r[inc()])) == 1,

		Level:         d2common.StringToInt(d2common.EmptyToZero(r[inc()])),
		RequiredLevel: d2common.StringToInt(d2common.EmptyToZero(r[inc()])),
		Code:          r[inc()],

		TypeDescription: r[inc()],
		UberDescription: r[inc()],
		SingleCopy:      d2common.StringToInt(d2common.EmptyToZero(r[inc()])) == 1,
		CostMultiplier:  d2common.StringToInt(d2common.EmptyToZero(r[inc()])),
		CostAdd:         d2common.StringToInt(d2common.EmptyToZero(r[inc()])),

		CharacterGfxTransform: r[inc()],
		InventoryGfxTransform: r[inc()],
		FlippyFile:            r[inc()],
		InventoryFile:         r[inc()],

		DropSound:    r[inc()],
		DropSfxFrame: d2common.StringToInt(d2common.EmptyToZero(r[inc()])),
		UseSound:     r[inc()],

		Properties: [12]UniqueItemProperty{
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),

			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),

			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
			createUniqueItemProperty(&r, inc),
		},
	}

	return result
}

func createUniqueItemProperty(r *[]string, inc func() int) UniqueItemProperty {
	result := UniqueItemProperty{
		Property:  (*r)[inc()],
		Parameter: d2common.CalcString((*r)[inc()]),
		Min:       d2common.StringToInt(d2common.EmptyToZero((*r)[inc()])),
		Max:       d2common.StringToInt(d2common.EmptyToZero((*r)[inc()])),
	}

	return result
}

// UniqueItems stores all of the UniqueItemRecords
//nolint:gochecknoglobals // Currently global by design
var UniqueItems map[string]*UniqueItemRecord

// LoadUniqueItems loadsUniqueItemRecords fro uniqueitems.txt
func LoadUniqueItems(file []byte) {
	UniqueItems = make(map[string]*UniqueItemRecord)
	data := strings.Split(string(file), "\r\n")[1:]

	for _, line := range data {
		if line == "" {
			continue
		}

		r := strings.Split(line, "\t")
		// skip rows that are not enabled
		if r[2] != "1" {
			continue
		}

		rec := createUniqueItemRecord(r)
		UniqueItems[rec.Code] = &rec
	}

	log.Printf("Loaded %d unique items", len(UniqueItems))
}
