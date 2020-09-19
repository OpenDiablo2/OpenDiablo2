package d2records

// SetItems holds all of the SetItemRecords
type SetItems map[string]*SetItemRecord

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
