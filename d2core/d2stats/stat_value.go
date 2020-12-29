package d2stats

// StatNumberType is a value type for a stat value
type StatNumberType int

//  Stat value types
const (
	StatValueInt StatNumberType = iota
	StatValueFloat
)

// ValueCombineType is a rule for combining stat values
type ValueCombineType int

const (
	// StatValueCombineSum means that the values are simply summed
	StatValueCombineSum ValueCombineType = iota

	// StatValueCombineStatic means that values can be combined only if they
	// have the same number value, and that the combination does not alter
	// the number value. This is typically for things like static skill level
	// monster/skill index for on proc stats where it doesnt make sense to sum
	// the values
	// example 1:
	//	if
	//		Stat_A := `25% chance to cast level 2 Frozen Orb on attack`
	//		Stat_B := `25% chance to cast level 3 Frozen Orb on attack`
	// then
	// 		Stat_A can NOT be combined with Stat_B
	//		even though the skills are the same, the levels are different
	//
	// example 2:
	//	if
	//		Stat_A := `25% chance to cast level 20 Frost Nova on attack`
	//		Stat_B := `25% chance to cast level 20 Frost Nova on attack`
	// then
	// 		the skills and skill levels are the same, so it can be combined
	//		(Stat_A + Stat_B) == `50% chance to cast level 20 Frost Nova on attack`
	StatValueCombineStatic
)

// StatValue is something that can have both integer and float
// number components, as well as a means of retrieving a string for
// its values.
type StatValue interface {
	NumberType() StatNumberType
	CombineType() ValueCombineType

	Clone() StatValue

	SetInt(int) StatValue
	SetFloat(float64) StatValue
	SetStringer(func(StatValue) string) StatValue

	Int() int
	Float() float64
	String() string
	Stringer() func(StatValue) string
}
