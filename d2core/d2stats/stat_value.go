package d2stats

// StatValueType is a value type for a stat value
type StatValueType int

//  Stat value types
const (
	StatValueInt StatValueType = iota
	StatValueFloat
)

// StatValue is something that can have both integer and float
// number components, as well as a means of retrieving a string for
// its values.
type StatValue interface {
	Type() StatValueType
	Clone() StatValue

	SetInt(int) StatValue
	SetFloat(float64) StatValue
	SetStringer(func(StatValue) string) StatValue

	Int() int
	Float() float64
	String() string
	Stringer() func(StatValue) string
}
