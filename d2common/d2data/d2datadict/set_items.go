package d2datadict

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"log"
)

const (
	numPropertiesOnSetItem = 9
	numBonusPropertiesOnSetItem = 5
	bonusToken1 = "a"
	bonusToken2 = "b"
	propCodeFmt = "prop%d"
	propParamFmt = "par%d"
	propMinFmt = "min%d"
	propMaxFmt = "max%d"
	bonusCodeFmt = "aprop%d%s"
	bonusParamFmt = "apar%d%s"
	bonusMinFmt = "amin%d%s"
	bonusMaxFmt = "amax%d%s"
)

// SetItemRecord represents a set item
type SetItemRecord struct {
	// SetItemKey (index)
	// string key to item's name in a .tbl file
	SetItemKey string

	// SetKey (set)
	// string key to the index field in Sets.txt - the set the item is a part of.
	SetKey string

	// ItemCode (item)
	// base item code of this set item (matches code field in Weapons.txt, Armor.txt or Misc.txt files).
	ItemCode string

	// Rarity
	// Chance to pick this set item if more then one set item of the same base item exist,
	// this uses the common rarity/total_rarity formula, so if you have two set rings,
	// one with a rarity of 100 the other with a rarity of 1,
	// then the first will drop 100/101 percent of the time (
	// 99%) and the other will drop 1/101 percent of the time (1%),
	// rarity can be anything between 0 and 255.
	Rarity int

	// QualityLevel (lvl)
	// The quality level of this set item, monsters, cube recipes, vendors,
	// objects and the like most be at least this level or higher to be able to drop this item,
	// otherwise they would drop a magical item with twice normal durability.
	QualityLevel int

	// RequiredLevel ("lvl req")
	// The character level required to use this set item.
	RequiredLevel int

	// CharacterPaletteTransform (chrtransform)
	// Palette shift to apply to the the DCC component-file and the DC6 flippy-file (
	// whenever or not the color shift will apply is determined by Weapons.txt,
	// Armor.txt or Misc.txt). This is an ID pointer from Colors.txt.
	CharacterPaletteTransform int

	// InventoryPaletteTransform (invtransform)
	// Palette shift to apply to the the DC6 inventory-file (
	// whenever or not the color shift will apply is determined by Weapons.txt,
	// Armor.txt or Misc.txt). This is an ID pointer from Colors.txt.
	InventoryPaletteTransform int

	// InvFile
	// Overrides the invfile and setinvfile specified in Weapons.txt,
	// Armor.txt or Misc.txt for the base item.
	// This field contains the file name of the DC6 inventory graphic (without the .dc6 extension).
	InvFile string

	// FlippyFile
	// Overrides the flippyfile specified in Weapons.txt, Armor.txt or Misc.txt for the base item.
	// This field contains the file name of the DC6 flippy animation (without the .dc6 extension).
	FlippyFile string

	// DropSound
	// Overrides the dropsound (the sound played when the item hits the ground) specified in Weapons.
	// txt, Armor.txt or Misc.txt for the base item. This field contains an ID pointer from Sounds.txt.
	DropSound string

	// DropSfxFrame
	// How many frames after the flippy animation starts playing will the associated drop sound start
	// to play. This overrides the values in Weapons.txt, Armor.txt or Misc.txt.
	DropSfxFrame int

	// UseSound
	// Overrides the usesound (the sound played when the item is consumed by the player) specified in
	// Weapons.txt, Armor.txt or Misc.txt for the base item.
	// This field contains an ID pointer from Sounds.txt.
	UseSound string

	// CostMult ("cost mult")
	// The base item's price is multiplied by this value when sold, repaired or bought from a vendor.
	CostMult int

	// CostAdd ("cost add")
	// After the price has been multiplied, this amount of gold is added to the price on top.
	CostAdd int

	// AddFn ("add func")
	// a property mode field that controls how the variable attributes will appear and be functional
	// on a set item. See the appendix for further details about this field's effects.
	AddFn int

	// Properties are a propert code, parameter, min, max for generating an item propert
	Properties [numPropertiesOnSetItem]*SetItemProperty

	// SetPropertiesLevel1 is the first version of bonus properties for the set
	SetPropertiesLevel1 [numBonusPropertiesOnSetItem]*SetItemProperty

	// SetPropertiesLevel2 is the second version of bonus properties for the set
	SetPropertiesLevel2 [numBonusPropertiesOnSetItem]*SetItemProperty
}

// SetItemProperty is describes a property of a set item
type SetItemProperty struct {
	Code      string
	Parameter string // depending on the property, this may be an int (usually), or a string
	Min       int
	Max       int
}

// SetItems holds all of the SetItemRecords
var SetItems map[string]*SetItemRecord //nolint:gochecknoglobals // Currently global by design,
// only written once

// LoadSetItems loads all of the SetItemRecords from SetItems.txt
func LoadSetItems(file []byte) {
	SetItems = make(map[string]*SetItemRecord)

	d := d2common.LoadDataDictionary(file)

	for d.Next() {
		record := &SetItemRecord{
			SetItemKey:                d.String("index"),
			SetKey:                    d.String("set"),
			ItemCode:                  d.String("item"),
			Rarity:                    d.Number("rarity"),
			QualityLevel:              d.Number("lvl"),
			RequiredLevel:             d.Number("lvl req"),
			CharacterPaletteTransform: d.Number("chrtransform"),
			InventoryPaletteTransform: d.Number("invtransform"),
			InvFile:                   d.String("invfile"),
			FlippyFile:                d.String("flippyfile"),
			DropSound:                 d.String("dropsound"),
			DropSfxFrame:              d.Number("dropsfxframe"),
			UseSound:                  d.String("usesound"),
			CostMult:                  d.Number("cost mult"),
			CostAdd:                   d.Number("cost add"),
			AddFn:                     d.Number("add func"),

		}

		// normal properties
		props := [numPropertiesOnSetItem]*SetItemProperty{}

		for idx := 0; idx < numPropertiesOnSetItem; idx++ {
			num := idx + 1
			props[idx] = &SetItemProperty{
				d.String(fmt.Sprintf(propCodeFmt, num)),
				d.String(fmt.Sprintf(propParamFmt, num)),
				d.Number(fmt.Sprintf(propMinFmt, num)),
				d.Number(fmt.Sprintf(propMaxFmt, num)),
			}
		}

		// set bonus properties
		bonus1 := [numBonusPropertiesOnSetItem]*SetItemProperty{}
		bonus2 := [numBonusPropertiesOnSetItem]*SetItemProperty{}

		for idx := 0; idx < numBonusPropertiesOnSetItem; idx++ {
			num := idx + 1

			bonus1[idx] = &SetItemProperty{
				d.String(fmt.Sprintf(bonusCodeFmt, num, bonusToken1)),
				d.String(fmt.Sprintf(bonusParamFmt, num, bonusToken1)),
				d.Number(fmt.Sprintf(bonusMinFmt, num,bonusToken1)),
				d.Number(fmt.Sprintf(bonusMaxFmt, num, bonusToken1)),
			}

			bonus2[idx] = &SetItemProperty{
				d.String(fmt.Sprintf(bonusCodeFmt, num, bonusToken2)),
				d.String(fmt.Sprintf(bonusParamFmt, num, bonusToken2)),
				d.Number(fmt.Sprintf(bonusMinFmt, num,bonusToken2)),
				d.Number(fmt.Sprintf(bonusMaxFmt, num, bonusToken2)),
			}
		}

		record.Properties = props

		SetItems[record.SetItemKey] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d SetItem records", len(SetItems))
}
