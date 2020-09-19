package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// CubeRecipes contains all rows in CubeMain.txt.
type CubeRecipes []*CubeRecipeRecord

// CubeRecipeRecord represents one row from CubeMain.txt.
// It is one possible recipe for the Horadric Cube, with
// requirements and output items.
// See: https://d2mods.info/forum/kb/viewarticle?a=284
type CubeRecipeRecord struct {
	// Description has no function, it just describes the
	// recipe.
	Description string

	// Enabled is true if the recipe is active in game.
	Enabled bool

	// Ladder is true if the recipe is only allowed in
	// ladder on realms. Also works for single player
	// TCP/IP.
	Ladder bool

	// MinDiff sets the minimum difficulty level required
	// to use this recipe.
	MinDiff int // 0, 1, 2 = normal, nightmare, hell

	// Version specifies whether the recipe is old
	// classic, new classic or expansion.
	Version int // 0, 1, 100 = old cl, new cl, expansion

	// The following three 'Req' values form a comparison:
	// if <ReqStatID> <ReqOperation> <ReqValue> then recipe
	// is allowed.
	//
	// ReqStatID is an ID value from the ItemStatsCost
	// data set specifying the stat to compare. Whether
	// this references a player or item stat depends on
	// the Operator.
	ReqStatID int
	// ReqOperation is a number describing the
	// comparison operator and the action to take if
	// it evaluates to true. See Appendix A in the
	// linked article and note that 1, 2, 27 and 28
	// are unusual.
	ReqOperation int // 1 - 28
	// ReqValue is the number the stat is compared against.
	ReqValue int

	// Class Can be used to make recipes class
	// specific. Example class codes given are:
	// ama bar pal nec sor dru ass
	//
	// Since this field isn't used in the game data,
	// classFieldToEnum has been implemented based on that
	// example. It understands the following syntax,
	// which may be incorrect:
	// "ama,bar,dru"
	Class []d2enum.Hero

	// NumInputs is the total count of input items
	// required, including counts in item stacks.
	NumInputs int

	// Inputs is the actual recipe, a collection of
	// items/stacks with parameters required to
	// obtain the items defined in Outputs.
	Inputs []CubeRecipeItem

	// Outputs are the items created when the recipe
	// is used.
	Outputs []CubeRecipeResult
}

// CubeRecipeResult is an item generated on use of a
// cube recipe.
type CubeRecipeResult struct {
	// Item is the item, with a count and parameters.
	Item CubeRecipeItem

	// Level causes the item to be a specific level.
	//
	// Note that this value force spawns the item at
	// this specific level. Its also used in the
	// formula for the next two fields.
	Level int // the item level of Item

	// PLevel uses a portion of the players level for
	// the output level.
	PLevel int

	// ILevel uses a portion of the first input's
	// level for the output level.
	ILevel int

	// Properties is a list of properties which may
	// be attached to Item.
	Properties []CubeRecipeItemProperty
}

// CubeRecipeItem represents an item, with a stack count
// and parameters. Here it is used to describe the
// required ingredients of the recipe and the output
// result. See:
// https://d2mods.info/forum/kb/viewarticle?a=284
type CubeRecipeItem struct {
	Code   string   // item code e.g. 'weap'
	Params []string // list of parameters e.g. 'sock'
	Count  int      // required stack count
}

// CubeRecipeItemProperty represents the mod #,
// mod # chance, mod # param, mod # min, mod # max
// fields in cubemain.txt
type CubeRecipeItemProperty struct {
	Code string // the code field from properties.txt

	// Note: I can't find any example value for this
	// so I've made it an int for now
	Chance int // the chance to apply the property

	// Note: The few examples in cubemain.txt are integers,
	// however d2records.UniqueItemProperty is a similar
	// struct which handles a similar field that may be a
	// string or an integer.
	//
	// See: https://d2mods.info/forum/kb/viewarticle?a=345
	// "the parameter passed on to the associated property, this is used to pass skill IDs,
	// state IDs, monster IDs, montype IDs and the like on to the properties that require
	// them, these fields support calculations."
	Param int // for properties that use parameters

	Min int // the minimum value of the property stat
	Max int // the maximum value of the property stat
}
