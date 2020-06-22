package d2datadict

import (
	"log"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// CubeRecipeRecord represents one row from CubeMain.txt.
type CubeRecipeRecord struct {
	// Description has no function, it just describes the recipe.
	Description string

	// Enabled is true if the recipe is active in game.
	Enabled bool

	// Ladder is true if the recipe is only allowed in ladder on realms. Also works for single player TCP/IP.
	Ladder bool

	// MinDiff sets the minimum difficulty level required to use this recipe.
	MinDiff int // 0, 1, 2 = normal, nightmare, hell

	// Version specifies whether the recipe is old classic, new classic or expansion.
	Version int // 0, 1, 100 = old cl, new cl, expansion

	// The following three values form a comparison: if <StatID> <Operation> <Value> then recipe is allowed.
	//
	// Parameter is an ID value from the ItemStatsCost data set specifying the stat to compare. Whether this references a player or item stat depends on the Operator.
	StatID int
	// Operation is a number describing the comparison operator and the action to take if it evaluates to true. See Appendix A in the linked article and note that 1, 2, 27 and 28 are unconventional.
	Operation int // 1 - 28
	// Value is the number the stat is compared against.
	Value int

	// Class Can be used to make recipes class specific. Example class codes given are:
	// ama bar pal nec sor dru ass
	//
	// Since this field isn't used in the game data, classField has been implemented based on that example. It understands the following syntax, which may be incorrect:
	// "ama,bar,dru"
	Class []d2enum.Hero

	// NumInputs is the total count of input items required, including counts in item stacks.
	NumInputs int

	// Inputs is the actual recipe, a collection of items/stacks with parameters required to obtain the items defined in Outputs.
	Inputs []CubeRecipeItem

	// Outputs are the items created when the recipe is used.
	Outputs []CubeRecipeResult
}

// CubeRecipeResult represents the output, lvl, plvl, ilvl, mod #, mod # chance, mod # param, mod # min, mod # max fields in cubemain.txt. It is an item generated on use of a cube recipe.
type CubeRecipeResult struct {
	// Note: The original data allows for only three output items/stacks.
	// Item the item, with a count and parameters.
	Item CubeRecipeItem

	// Note that this value force spawns the item at this specific level. Its also used in the formula for the next two fields.
	Level int // the item level of Item

	// PLevel uses a portion of the players level for the output level.
	PLevel int

	// ILevel uses a portion of the first input's level for the output level.
	ILevel int

	// Note: The original data allows for only five properties.
	// Properties is a list of properties which may be attached to Item.
	Properties []CubeRecipeItemProperty
}

// CubeRecipeItem represents the data from the input 1-7 fields in cubemain.txt. Each cell represents a different item to be placed in the cube, stackable items can also have a required count. Originally a comma-delimited string of parameters.
// See: https://d2mods.info/forum/kb/viewarticle?a=284
type CubeRecipeItem struct {
	Code   string   // item code e.g. 'weap'
	Params []string // list of argument parameters e.g. 'sock'
	Count  int      // required stack count
}

// CubeRecipeItemProperty represents a property attached to an output item. Properties are defined in properties.txt. See:
// https://d2mods.info/forum/kb/viewarticle?a=345
type CubeRecipeItemProperty struct {
	Code string // the code field from properties.txt

	// Note: I can't find any example value for this so I've made it an int for now
	Chance int // the chance to apply the property

	// Note: The few examples in cubemain.go are integers, however d2datadict.UniqueItemProperty is a similar struct which handles a similar field that may be a string or an integer.
	//"the parameter passed on to the associated property, this is used to pass skill IDs, state IDs, monster IDs, montype IDs and the like on to the properties that require them, these fields support calculations."
	Param int // for properties that use parameters

	Min int // the minimum value of the property stat
	Max int // the maximum value of the property stat
}

// CubeRecipes contains all rows CubeMain.txt.
var CubeRecipes []*CubeRecipeRecord

// The fields and sections in this file have inconsistent naming conventions. These arrays are a simple way to handle them.
var outputFields = []string{"output", "output b", "output c"}
var outputLabels = []string{"", "b ", "c "}
var propLabels = []string{"mod 1", "mod 2", "mod 3", "mod 4", "mod 5"}
var inputFields = []string{"input 1", "input 2", "input 3", "input 4", "input 5", "input 6", "input 7"}

// LoadCubeRecipes populates CubeRecipes with the data from CubeMain.txt.
func LoadCubeRecipes(file []byte) {

	// Load data
	d := d2common.LoadDataDictionary(string(file))

	// Create records
	CubeRecipes = make([]*CubeRecipeRecord, len(d.Data))
	for idx := range d.Data {
		CubeRecipes[idx] = &CubeRecipeRecord{
			Description: d.GetString("description", idx),

			Enabled: d.GetNumber("enabled", idx) == 1,
			Ladder:  d.GetNumber("ladder", idx) == 1,

			MinDiff: d.GetNumber("min diff", idx),
			Version: d.GetNumber("version", idx),

			StatID:    d.GetNumber("param", idx),
			Operation: d.GetNumber("op", idx),
			Value:     d.GetNumber("value", idx),

			Class: classField(d.GetString("class", idx)),

			NumInputs: d.GetNumber("numinputs", idx),
		}

		// Create inputs
		CubeRecipes[idx].Inputs = make([]CubeRecipeItem, 7)
		for i := range inputFields {
			CubeRecipes[idx].Inputs[i] = newCubeRecipeItem(
				d.GetString(inputFields[i], idx))
		}

		// Create outputs
		CubeRecipes[idx].Outputs = make([]CubeRecipeResult, 3)
		for o, outLabel := range outputLabels {

			CubeRecipes[idx].Outputs[o] = CubeRecipeResult{
				Item: newCubeRecipeItem(
					d.GetString(outputFields[o], idx)),

				Level:  d.GetNumber(outLabel+"lvl", idx),
				ILevel: d.GetNumber(outLabel+"plvl", idx),
				PLevel: d.GetNumber(outLabel+"ilvl", idx),
			}

			// Create properties
			properties := make([]CubeRecipeItemProperty, 5)
			for p, prop := range propLabels {

				properties[p] = CubeRecipeItemProperty{
					Code:   d.GetString(outLabel+prop, idx),
					Chance: d.GetNumber(outLabel+prop+" chance", idx),
					Param:  d.GetNumber(outLabel+prop+" param", idx),
					Min:    d.GetNumber(outLabel+prop+" min", idx),
					Max:    d.GetNumber(outLabel+prop+" max", idx),
				}
			}

			CubeRecipes[idx].Outputs[o].Properties = properties
		}

	}

	log.Printf( /*"Loaded %d CubeMainRecord records"*/ "LoadCubeRecipes ran - %d", len(CubeRecipes))
}

// newCubeRecipeItem constructs a CubeRecipeItem from a string of arguments. arguments include at least an item and sometimes parameters and/or a count (qty parameter).
func newCubeRecipeItem(f string) CubeRecipeItem {
	args := stringArray(f)

	item := CubeRecipeItem{
		Code:  args[0], // the first argument is always the item count
		Count: 1,       // default to a count of 1 (no qty parameter)
	}

	// Ignore the first argument
	args = args[1:]

	// Find the qty parameter if it was provided, convert to int and assign to item.Count
	for idx, arg := range args {
		if strings.HasPrefix(arg, "qty") {
			count, err := strconv.Atoi(strings.Split(arg, "=")[1])
			if err != nil {
				log.Fatal("Error parsing item count:", err)
			}
			item.Count = count

			// Remove the qty parameter
			if idx != len(args)-1 {
				args[idx] = args[len(args)-1]
			}
			args = args[:len(args)-1]

			break
		}
	}

	// No other arguments were provided
	if len(args) == 0 {
		return item
	}

	// Record the parameter strings
	item.Params = make([]string, len(args))
	for idx, arg := range args {
		item.Params[idx] = arg
	}

	return item
}

// classField converts class tokens to s2enum.Hero.
func classField(f string) []d2enum.Hero {
	split := stringArray(f)
	enums := make([]d2enum.Hero, len(split))
	for idx, class := range split {
		if class == "" {
			continue
		}
		switch class {
		case "bar":
			enums[idx] = d2enum.HeroBarbarian
		case "nec":
			enums[idx] = d2enum.HeroNecromancer
		case "pal":
			enums[idx] = d2enum.HeroPaladin
		case "ass":
			enums[idx] = d2enum.HeroAssassin
		case "sor":
			enums[idx] = d2enum.HeroSorceress
		case "ama":
			enums[idx] = d2enum.HeroAmazon
		case "dru":
			enums[idx] = d2enum.HeroDruid
		default:
			log.Fatalf("Unknown hero token: '%s'", class)
		}
	}
	return enums
}

// stringArray splits a string array from the following format:
// "one,two,three"
func stringArray(s string) []string {
	return strings.Split(strings.Trim(s, "\""), ",")
}
