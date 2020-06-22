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

	// Requirement holds the parameters for a comparison which can cause the recipe to be skipped or fail.
	Requirement CubeRecipeRequirement

	// Class Can be used to make recipes class specific. Example class codes given are:
	// ama bar pal nec sor dru ass
	//
	// Since this field isn't used in the game data, classField has been implemented based on that example. It understands the following syntax, which may be incorrect:
	// "ama,bar,dru"
	Class []d2enum.Hero

	// NumInputs is the total count of input items required, including counts in item stacks.
	NumInputs int

	// Inputs is a list of CubeRecipeInput structs representing the data from the input 1-7 field
	Inputs []CubeRecipeInput

	//output
	//lvl
	//plvl
	//ilvl
	//mod 1
	//mod 1 chance
	//mod 1 param
	//mod 1 min
	//mod 1 max
	//mod 2
	//mod 2 chance
	//mod 2 param
	//mod 2 min
	//mod 2 max
	//mod 3
	//mod 3 chance
	//mod 3 param
	//mod 3 min
	//mod 3 max
	//mod 4
	//mod 4 chance
	//mod 4 param
	//mod 4 min
	//mod 4 max
	//mod 5
	//mod 5 chance
	//mod 5 param
	//mod 5 min
	//mod 5 max
	//output b
	//b lvl
	//b plvl
	//b ilvl
	//b mod 1
	//b mod 1 chance
	//b mod 1 param
	//b mod 1 min
	//b mod 1 max
	//b mod 2
	//b mod 2 chance
	//b mod 2 param
	//b mod 2 min
	//b mod 2 max
	//b mod 3
	//b mod 3 chance
	//b mod 3 param
	//b mod 3 min
	//b mod 3 max
	//b mod 4
	//b mod 4 chance
	//b mod 4 param
	//b mod 4 min
	//b mod 4 max
	//b mod 5
	//b mod 5 chance
	//b mod 5 param
	//b mod 5 min
	//b mod 5 max
	//output c
	//c lvl
	//c plvl
	//c ilvl
	//c mod 1
	//c mod 1 chance
	//c mod 1 param
	//c mod 1 min
	//c mod 1 max
	//c mod 2
	//c mod 2 chance
	//c mod 2 param
	//c mod 2 min
	//c mod 2 max
	//c mod 3
	//c mod 3 chance
	//c mod 3 param
	//c mod 3 min
	//c mod 3 max
	//c mod 4
	//c mod 4 chance
	//c mod 4 param
	//c mod 4 min
	//c mod 4 max
	//c mod 5
	//c mod 5 chance
	//c mod 5 param
	//c mod 5 min
	//c mod 5 max
	//*eol
}

// CubeRecipeRequirement represents the op, param and value fields in cubemain.go. These form a comparison: if <Paremeter> <Operator> <Value> then recipe is allowed.
// See: https://d2mods.info/forum/kb/viewarticle?a=284
type CubeRecipeRequirement struct {
	// Parameter is an ID value from the ItemStatsCost data set specifying the stat to compare. Whether this references a player or item stat depends on the Operator.
	StatID int
	// Operation is a number describing the comparison operator and the action to take if it evaluates to true. See Appendix A in the linked article and note that 1, 2, 27 and 28 are unconventional.
	Operation int // 1 - 28
	// Value is the number the stat is compared against.
	Value int
}

// CubeRecipeInput represents the data from the input 1-7 fields in cubemain.txt. Each cell represents a different item to be placed in the cube, stackable items can also have a required count. Originally a comma-delimited string of parameters.
// See: https://d2mods.info/forum/kb/viewarticle?a=284
type CubeRecipeInput struct {
	Code   string   // item code e.g. 'weap'
	Params []string // list of argument parameters e.g. 'sock'
	Count  int      // required stack count
}

// CubeRecipes contains all rows CubeMain.txt.
var CubeRecipes []*CubeRecipeRecord

// LoadCubeRecipes populates CubeRecipes with the data from CubeMain.txt.
func LoadCubeRecipes(file []byte) {
	// Load data
	d := d2common.LoadDataDictionary(string(file))

	// Construct records
	CubeRecipes = make([]*CubeRecipeRecord, len(d.Data))
	for idx := range d.Data {
		CubeRecipes[idx] = &CubeRecipeRecord{
			Description: d.GetString("description", idx),

			Enabled: d.GetNumber("enabled", idx) == 1,
			Ladder:  d.GetNumber("ladder", idx) == 1,

			MinDiff: d.GetNumber("min diff", idx),
			Version: d.GetNumber("version", idx),

			Requirement: CubeRecipeRequirement{
				StatID:    d.GetNumber("param", idx),
				Operation: d.GetNumber("op", idx),
				Value:     d.GetNumber("value", idx),
			},

			Class: classField(d.GetString("class", idx)),

			NumInputs: d.GetNumber("numinputs", idx),

			Inputs: inputFieldsArray(
				d.GetString("input 1", idx),
				d.GetString("input 2", idx),
				d.GetString("input 3", idx),
				d.GetString("input 4", idx),
				d.GetString("input 5", idx),
				d.GetString("input 6", idx),
				d.GetString("input 7", idx),
			),

			//output
			//lvl
			//plvl
			//ilvl
			//mod 1
			//mod 1 chance
			//mod 1 param
			//mod 1 min
			//mod 1 max
			//mod 2
			//mod 2 chance
			//mod 2 param
			//mod 2 min
			//mod 2 max
			//mod 3
			//mod 3 chance
			//mod 3 param
			//mod 3 min
			//mod 3 max
			//mod 4
			//mod 4 chance
			//mod 4 param
			//mod 4 min
			//mod 4 max
			//mod 5
			//mod 5 chance
			//mod 5 param
			//mod 5 min
			//mod 5 max
			//output b
			//b lvl
			//b plvl
			//b ilvl
			//b mod 1
			//b mod 1 chance
			//b mod 1 param
			//b mod 1 min
			//b mod 1 max
			//b mod 2
			//b mod 2 chance
			//b mod 2 param
			//b mod 2 min
			//b mod 2 max
			//b mod 3
			//b mod 3 chance
			//b mod 3 param
			//b mod 3 min
			//b mod 3 max
			//b mod 4
			//b mod 4 chance
			//b mod 4 param
			//b mod 4 min
			//b mod 4 max
			//b mod 5
			//b mod 5 chance
			//b mod 5 param
			//b mod 5 min
			//b mod 5 max
			//output c
			//c lvl
			//c plvl
			//c ilvl
			//c mod 1
			//c mod 1 chance
			//c mod 1 param
			//c mod 1 min
			//c mod 1 max
			//c mod 2
			//c mod 2 chance
			//c mod 2 param
			//c mod 2 min
			//c mod 2 max
			//c mod 3
			//c mod 3 chance
			//c mod 3 param
			//c mod 3 min
			//c mod 3 max
			//c mod 4
			//c mod 4 chance
			//c mod 4 param
			//c mod 4 min
			//c mod 4 max
			//c mod 5
			//c mod 5 chance
			//c mod 5 param
			//c mod 5 min
			//c mod 5 max
			//*eol
		}
	}

	log.Printf( /*"Loaded %d CubeMainRecord records"*/ "LoadCubeRecipes ran - %d", len(CubeRecipes))
}

// inputFieldsArray constructs an array of CubeRecipeInput from an array of strings, ignoring empty strings.
func inputFieldsArray(fields ...string) []CubeRecipeInput {
	// Count fields with values
	count := 0
	for _, f := range fields {
		if f != "" {
			count++
		}
	}

	// Construct CubeRecipeInputs
	inputs := make([]CubeRecipeInput, count)
	index := 0
	for _, f := range fields {
		if f != "" {
			inputs[index] = newCubeRecipeInput(f)
			index++
		}
	}
	return inputs
}

// newCubeRecipe constructs a CubeRecipeInput from a string of arguments.
func newCubeRecipeInput(f string) CubeRecipeInput {
	args := stringArray(f)

	cube := CubeRecipeInput{
		Code:  args[0], // the first argument is always the item count
		Count: 1,       // default to a count of 1 (no qty parameter)
	}

	// Ignore the first argument
	args = args[1:]

	// Find the qty parameter if it was provided, convert to int and assign to cube.Count
	for idx, arg := range args {
		if strings.HasPrefix(arg, "qty") {
			count, err := strconv.Atoi(strings.Split(arg, "=")[1])
			if err != nil {
				log.Fatal("Error parsing item count:", err)
			}
			cube.Count = count

			// If qty isn't the last item, move the last item to its place
			if idx != len(args)-1 {
				args[idx] = args[len(args)-1]
			}
			// Ignore the last item
			args = args[:len(args)-1]
		}
	}

	// No other arguments were provided
	if len(args) == 0 {
		return cube
	}

	// Record the parameter strings
	cube.Params = make([]string, len(args))
	for idx, arg := range args {
		cube.Params[idx] = arg
	}

	return cube
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
