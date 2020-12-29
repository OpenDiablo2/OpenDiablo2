package d2enum

// Filter represents the type of texture filter to be used when an image is magnified or minified.
type Filter int

const (
	// FilterDefault represents the default filter.
	FilterDefault Filter = iota

	// FilterNearest represents nearest (crisp-edged) filter
	FilterNearest

	// FilterLinear represents linear filter
	FilterLinear
)
