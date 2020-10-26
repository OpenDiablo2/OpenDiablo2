package d2records

// PropertyDescriptor is a generic description of a property, used to create properties.
// Sets, SetItems, UniqueItems will all use this generic form of a property
type PropertyDescriptor struct {
	// Code is an ID pointer of a property from Properties.txt,
	// these columns control each of the eight different full set modifiers a set item can grant you
	// at most.
	Code string

	// Param is the passed on to the associated property, this is used to pass skill IDs, state IDs,
	// monster IDs, montype IDs and the like on to the properties that require them,
	// these fields support calculations.
	Parameter string

	// Min value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	Min int

	// Max value to assign to the associated property.
	// Certain properties have special interpretations based on stat encoding (e.g.
	// chance-to-cast and charged skills). See the File Guides for Properties.txt and ItemStatCost.
	// txt for further details.
	Max int
}
