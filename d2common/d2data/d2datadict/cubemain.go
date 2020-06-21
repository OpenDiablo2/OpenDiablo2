package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// CubeRecipeRecord represents one row from CubeMain.txt.
type CubeRecipeRecord struct {
	//description
	//enabled
	//ladder
	//min diff
	//version
	//op
	//param
	//value
	//class
	//numinputs
	//input 1
	//input 2
	//input 3
	//input 4
	//input 5
	//input 6
	//input 7
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
			// Populate struct fields
		}
	}

	log.Printf( /*"Loaded %d CubeMainRecord records"*/ "LoadCubeRecipes ran - %d", len(CubeRecipes))
}
