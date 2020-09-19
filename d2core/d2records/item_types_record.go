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
	// Name (ItemType)
	// A comment field that contains the “internal name” of this iType,
	// you can basically  enter anything you wish here,
	// but since you can add as many comment columns as you wish,
	// there is no reason to use it for another purpose .
	Name string

	// Code
	// The ID pointer of this ItemType, this pointer is used in many txt files  (armor.txt,
	// cubemain.txt, misc.txt, skills.txt, treasureclassex.txt, weapons.txt),
	// never use the same ID pointer twice,
	// the game will only use the first instance and ignore all other occurrences.
	// ID pointers are case sensitive, 3-4 chars long and can contain numbers, letters and symbols.
	Code string

	// Equiv1-2
	// This is used to define the parent iType, note that an iType can have multiple parents (
	// as will be shown in the cladogram – link below),
	// the only thing you must avoid at all cost is creating infinite loops.
	// I haven't ever tested what happens when you create an iType loop,
	// but infinite loops are something you should always avoid.
	Equiv1 string
	Equiv2 string

	// Shoots
	// This column specifies which type of quiver (“ammo”) this iType (
	// in case it is a weapon) requires in order to shoot (
	// you use the ID pointer of the quiver iType here).
	// Caution: The place it checks which missile to pick (either arrow, bolt,
	// explosive arrow or magic arrow) is buried deep within D2Common.dll,
	// the section can be modified, there is an extensive post discussing this in Code Editing.
	// - Thanks go to Kingpin for spotting a silly little mistake in here.
	Shoots string

	// Quiver
	// The equivalent to the previous column,
	// in here you specify which weapon this quiver is linked to. Make sure the two columns match. (
	// this also uses the ID pointer of course).
	Quiver string

	// InvGfx1-6
	// This column contains the file names of the inventory graphics that are randomly picked for
	// this iType, so if you use columns 1-3, you will set VarInvGfx to 3 (duh).
	InvGfx1 string
	InvGfx2 string
	InvGfx3 string
	InvGfx4 string
	InvGfx5 string
	InvGfx6 string

	// StorePage
	// The page code for the page a vendor should place this iType in when sold,
	// if you enable the magic tab in D2Client.dll,
	// you need to use the proper code here to put items in that tab.
	// Right now the ones used are weap = weapons1 and 2, armo = armor and misc = miscellaneous.
	StorePage string

	// BodyLoc1-2
	// If you have set the previous column to 1,
	// you need to specify the inventory slots in which the item has to be equipped. (
	// the codes used by this field are read from BodyLocs.txt)
	BodyLoc1 int
	BodyLoc2 int

	// MaxSock1, MaxSock25, MaxSock40
	// Maximum sockets for iLvl 1-25,
	// 26-40 and 40+. The range is hardcoded but the location is known,
	// so you can alter around the range to your liking. On normal,
	// items dropped from monsters are limited to 3, on nightmare to 4 and on hell to 6 sockets,
	// irregardless of this columns content.
	MaxSock1  int
	MaxSock25 int
	MaxSock40 int

	// TreasureClass
	// Can this iType ID Pointer be used as an auto TC in TreasureClassEx.txt. 1=Yes,
	// 0=No. *Such as armo3-99 and weap3-99 etc.
	TreasureClass int

	// Rarity
	// Dunno what it does, may have to do with the chance that an armor or weapon rack will pick
	// items of this iType. If it works like other rarity fields,
	// the chance is rarity / total_rarity * 100.
	Rarity int

	// StaffMods
	// Contains the class code for the character class that should get +skills from this iType (
	// such as wands that can spawn with +Necromancer skills). Note,
	// this only works if the item is not low quality, set or unique. Note,
	// that this uses the vanilla min/max skill IDs for each class as the range for the skill pool,
	// so if you add new class skills to the end of the file, you should use automagic.txt instead
	StaffMods d2enum.Hero

	// CostFormula
	// Does the game generate the sell/repair/buy prices of this iType based on its modifiers or does
	// it use only the cost specific in the respective item txt files. 2=Organ (
	// probably higher price based on unit that dropped the organ), 1=Yes, 0=No.
	// Note: Only applies to items that are not unique or set, for those the price is solely controlled
	// by the base item file and by the bonus to price given in SetItems and UniqueItems txt files.
	// The exact functionality remains unknown, as for example charms, have this disabled.
	CostFormula int

	// Class
	// Contains the class code for the class that should be able to use this iType (
	// for class specific items).
	Class d2enum.Hero

	// VarInvGfx
	// This column contains the sum of randomly picked inventory graphics this iType can have.
	VarInvGfx int

	// Repair
	// Boolean, 1=Merchants can repair this item type, 0=Merchants cannot repair this iType (note,
	// this also refers to charges being rechargeable).
	Repair bool

	// Body
	// Boolean, 1=The character can wear this iType,
	// 0=This iType can only be carried in the inventory,
	// cube or stash (and belt if it is set as “beltable” in the other item related txt files)
	Body bool

	// Throwable
	// Can this iType be thrown (determines whenever it uses the quantity and throwing damage columns
	// in Weapons.txt for example).
	Throwable bool

	// Reload
	// Can the this item be re-stacked via drag and drop. 1=Yes, 0=No.
	Reload bool

	// ReEquip
	// If the ammo runs out the game will automatically pick the next item of the same iType to
	// be equipped in it's place.
	// 1=Yes, 0=No. (more clearly, when you use up all the arrows in a quiver, the next quiver,
	// if available, will be equipped in its place).
	ReEquip bool

	// AutoStack
	// Are identical stacks automatically combined when you pick the up? 1=Yes, 0=No. (for example,
	// which you pick up throwing potions or normal javelins,
	// they are automatically combined with those you already have)
	AutoStack bool

	// Magic
	// Is this iType always Magic? 1=Yes, 0=No.
	Magic bool

	// Rare
	// Can this iType spawn as a rare item?
	// 1=Yes, 0=No.
	// Note: If you want an item that spawns only as magic or rare,
	// you need to set the previous column to 1 as well.
	Rare bool

	// Normal
	// Is this iType always Normal? 1=Yes, 0=No.
	Normal bool

	// Charm
	// Does this iType function as a charm? 1=Yes, 0=No. Note: This effect is hardcoded,
	// if you need a new charm type, you must use the char iType in one of the equivs.
	Charm bool

	// Gem
	// Can this iType be inserted into sockets? 1=Yes,
	// 0=No (Link your item to the sock iType instead to achieve this).
	Gem bool

	// Beltable
	// Can this iType be placed in your characters belt slots? 1=Yes,
	// 0=No. (This requires further tweaking in other txt files).
	Beltable bool

	EquivalentItems ItemEquivalenceList
}
