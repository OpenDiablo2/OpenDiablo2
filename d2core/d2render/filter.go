package d2render

// Filter represents the type of texture filter to be used when an image is maginified or minified.
type Filter int

const (
	// FilterDefault represents the default filter.
	FilterDefault Filter = 0

	// FilterNearest represents nearest (crisp-edged) filter
	FilterNearest = Filter(1)

	// FilterLinear represents linear filter
	FilterLinear = Filter(2)
)
