package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"log"
)

// SetItemRecord represents a set item
type SetItemRecord struct {
	// StringTableKey (index)
	// string key to item's name in a .tbl file
	StringTableKey string

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

	// Prop (prop1 to prop9)
	// An ID pointer of a property from Properties.txt,
	// these columns control each of the nine different fixed (
	// blue) modifiers a set item can grant you at most.
	Prop [9]string

	// Par (par1 to par9)
	// The parameter passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	Par [9]int

	// Min, Max (min1 to min9, max1 to max9)
	// Minimum value to assign to the associated (blue) property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guide for Properties.txt and ItemStatCost.
	// txt for further details.
	Min [9]int
	Max [9]int

	// APropA, APropB (aprop1a,aprop1b to aprop5a,aprop5b)
	// An ID pointer of a property from Properties.txt,
	// these columns control each of the five pairs of different variable (
	// green) modifiers a set item can grant you at most.
	APropA [5]string
	APropB [5]string

	// AParA, AParB (apar1a,apar1b to apar5a,apar5b)
	// The parameter passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	AParA [5]int
	AParB [5]int

	// AMinA, AMinB, AMaxA, AMaxB (amin1a,amin1b to amin5a,amin5b)
	// Minimum value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guide for Properties.txt and ItemStatCost.
	// txt for further details.
	AMinA [5]int
	AMinB [5]int
	AMaxA [5]int
	AMaxB [5]int
}

// SetItems holds all of the SetItemRecords
var SetItems []*SetItemRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadSetItems loads all of the SetItemRecords from SetItems.txt
func LoadSetItems(file []byte) {
	SetItems = make([]*SetItemRecord, 0)

	d := d2common.LoadDataDictionary(file)

	for d.Next() {
		record := &SetItemRecord{
			StringTableKey:            d.String("index"),
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
			Prop: [9]string{
				d.String("prop1"),
				d.String("prop2"),
				d.String("prop3"),
				d.String("prop4"),
				d.String("prop5"),
				d.String("prop6"),
				d.String("prop7"),
				d.String("prop8"),
				d.String("prop9"),
			},
			Par: [9]int{
				d.Number("par1"),
				d.Number("par2"),
				d.Number("par3"),
				d.Number("par4"),
				d.Number("par5"),
				d.Number("par6"),
				d.Number("par7"),
				d.Number("par8"),
				d.Number("par9"),
			},
			Min: [9]int{
				d.Number("min1"),
				d.Number("min2"),
				d.Number("min3"),
				d.Number("min4"),
				d.Number("min5"),
				d.Number("min6"),
				d.Number("min7"),
				d.Number("min8"),
				d.Number("min9"),
			},
			Max: [9]int{
				d.Number("max1"),
				d.Number("max2"),
				d.Number("max3"),
				d.Number("max4"),
				d.Number("max5"),
				d.Number("max6"),
				d.Number("max7"),
				d.Number("max8"),
				d.Number("max9"),
			},
			APropA: [5]string{
				d.String("aprop1a"),
				d.String("aprop2a"),
				d.String("aprop3a"),
				d.String("aprop4a"),
				d.String("aprop5a"),
			},
			APropB: [5]string{
				d.String("aprop1b"),
				d.String("aprop2b"),
				d.String("aprop3b"),
				d.String("aprop4b"),
				d.String("aprop5b"),
			},
			AParA: [5]int{
				d.Number("apar1a"),
				d.Number("apar2a"),
				d.Number("apar3a"),
				d.Number("apar4a"),
				d.Number("apar5a"),
			},
			AParB: [5]int{
				d.Number("apar1b"),
				d.Number("apar2b"),
				d.Number("apar3b"),
				d.Number("apar4b"),
				d.Number("apar5b"),
			},
			AMinA: [5]int{
				d.Number("amin1a"),
				d.Number("amin2a"),
				d.Number("amin3a"),
				d.Number("amin4a"),
				d.Number("amin5a"),
			},
			AMinB: [5]int{
				d.Number("amin1b"),
				d.Number("amin2b"),
				d.Number("amin3b"),
				d.Number("amin4b"),
				d.Number("amin5b"),
			},
			AMaxA: [5]int{
				d.Number("amax1a"),
				d.Number("amax2a"),
				d.Number("amax3a"),
				d.Number("amax4a"),
				d.Number("amax5a"),
			},
			AMaxB: [5]int{
				d.Number("amax1b"),
				d.Number("amax2b"),
				d.Number("amax3b"),
				d.Number("amax4b"),
				d.Number("amax5b"),
			},
		}

		SetItems = append(SetItems, record)
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d SetItem records", len(SetItems))
}
