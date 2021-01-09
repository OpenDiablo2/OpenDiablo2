package d2records

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func cubeRecipeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make([]*CubeRecipeRecord, 0)

	// There are repeated fields and sections in this file, some
	// of which have inconsistent naming conventions. These slices
	// are a simple way to handle them.
	var outputFields = []string{"output", "output b", "output c"}

	var outputLabels = []string{"", "b ", "c "}

	var propLabels = []string{"mod 1", "mod 2", "mod 3", "mod 4", "mod 5"}

	var inputFields = []string{"input 1", "input 2", "input 3", "input 4", "input 5", "input 6", "input 7"}

	for d.Next() {
		class, err := classFieldToEnum(d.String("class"))
		if err != nil {
			return err
		}

		record := &CubeRecipeRecord{
			Description: d.String("description"),

			Enabled: d.Bool("enabled"),
			Ladder:  d.Bool("ladder"),

			MinDiff: d.Number("min diff"),
			Version: d.Number("version"),

			ReqStatID:    d.Number("param"),
			ReqOperation: d.Number("op"),
			ReqValue:     d.Number("value"),

			Class: class,

			NumInputs: d.Number("numinputs"),
		}

		// Create inputs - input 1-7
		record.Inputs = make([]CubeRecipeItem, len(inputFields))
		for i := range inputFields {
			record.Inputs[i], err = newCubeRecipeItem(
				d.String(inputFields[i]))
			if err != nil {
				return err
			}
		}

		// Create outputs - output "", b, c
		record.Outputs = make([]CubeRecipeResult, len(outputLabels))

		for o, outLabel := range outputLabels {
			item, err := newCubeRecipeItem(
				d.String(outputFields[o]))
			if err != nil {
				return err
			}

			record.Outputs[o] = CubeRecipeResult{
				Item:   item,
				Level:  d.Number(outLabel + "lvl"),
				ILevel: d.Number(outLabel + "plvl"),
				PLevel: d.Number(outLabel + "ilvl"),
			}

			// Create properties - mod 1-5
			properties := make([]CubeRecipeItemProperty, len(propLabels))
			for p, prop := range propLabels {
				properties[p] = CubeRecipeItemProperty{
					Code:   d.String(outLabel + prop),
					Chance: d.Number(outLabel + prop + " chance"),
					Param:  d.Number(outLabel + prop + " param"),
					Min:    d.Number(outLabel + prop + " min"),
					Max:    d.Number(outLabel + prop + " max"),
				}
			}

			record.Outputs[o].Properties = properties
		}

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d CubeMainRecord records", len(records))

	r.Item.Cube.Recipes = records

	return nil
}

// newCubeRecipeItem constructs a CubeRecipeItem from a string of
// arguments. arguments include at least an item and sometimes
// parameters and/or a count (qty parameter). For example:
// "weap,sock,mag,qty=10"
func newCubeRecipeItem(f string) (CubeRecipeItem, error) {
	args := splitFieldValue(f)

	item := CubeRecipeItem{
		Code:  args[0], // the first argument is always the item count
		Count: 1,       // default to a count of 1 (no qty parameter)
	}

	// Ignore the first argument
	args = args[1:]

	// Find the qty parameter if it was provided,
	// convert to int and assign to item.Count
	for idx, arg := range args {
		if !strings.HasPrefix(arg, "qty") {
			continue
		}

		count, err := strconv.Atoi(strings.Split(arg, "=")[1])

		if err != nil {
			// need to be verified
			return item, fmt.Errorf("error parsing item count %e", err)
		}

		item.Count = count

		// Remove the qty parameter
		if idx != len(args)-1 {
			args[idx] = args[len(args)-1]
		}

		args = args[:len(args)-1]

		break
	}

	// No other arguments were provided
	if len(args) == 0 {
		return item, nil
	}

	// Record the argument strings
	item.Params = make([]string, len(args))
	for idx, arg := range args {
		item.Params[idx] = arg
	}

	return item, nil
}

// classFieldToEnum converts class tokens to s2enum.Hero.
func classFieldToEnum(f string) ([]d2enum.Hero, error) {
	split := splitFieldValue(f)
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
			return nil, fmt.Errorf("unknown hero token: '%s'", class)
		}
	}

	return enums, nil
}

// splitFieldValue splits a string array from the following format:
// "one,two,three"
func splitFieldValue(s string) []string {
	return strings.Split(strings.Trim(s, "\""), ",")
}
