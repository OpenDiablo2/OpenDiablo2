package d2item

// Item describes all types of item that can be placed in the
// player inventory grid (not just things that can be equipped!)
type Item interface {
	Context() StatContext
	SetContext(StatContext)

	Label() string
	Description() string
}
