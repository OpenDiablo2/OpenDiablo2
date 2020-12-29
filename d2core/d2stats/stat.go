package d2stats

// Stat a generic interface for a stat. It is something which can be
// combined with other stats, holds one or more values, and handles the
// way that it is printed as a string
type Stat interface {
	Name() string
	Clone() Stat
	Copy(Stat) Stat
	Combine(Stat) (combined Stat, err error)
	String() string
	Values() []StatValue
	SetValues(...StatValue)
	Priority() int
}
