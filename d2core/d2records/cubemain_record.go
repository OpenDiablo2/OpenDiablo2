package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// CubeRecipes contains all rows in CubeMain.txt.
type CubeRecipes []*CubeRecipeRecord

// CubeRecipeRecord represents one row from CubeMain.txt.
// It is one possible recipe for the Horadric Cube, with
// requirements and output items.
// See: https://d2mods.info/forum/kb/viewarticle?a=284
type CubeRecipeRecord struct {
	Description  string
	Class        []d2enum.Hero
	Inputs       []CubeRecipeItem
	Outputs      []CubeRecipeResult
	ReqValue     int
	MinDiff      int
	ReqOperation int
	NumInputs    int
	Version      int
	ReqStatID    int
	Enabled      bool
	Ladder       bool
}

// CubeRecipeResult is an item generated on use of a
// cube recipe.
type CubeRecipeResult struct {
	Properties []CubeRecipeItemProperty
	Item       CubeRecipeItem
	Level      int
	PLevel     int
	ILevel     int
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
