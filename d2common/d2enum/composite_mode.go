package d2enum

// CompositeMode defines the composite mode
type CompositeMode int

const (
	// CompositeModeSourceOver applies a composite based on:
	// c_out = c_src + c_dst × (1 - α_src) (Regular alpha blending)
	CompositeModeSourceOver CompositeMode = iota + 1

	// CompositeModeClear applies a composite based on: c_out = 0
	CompositeModeClear

	// CompositeModeCopy applies a composite based on: c_out = c_src
	CompositeModeCopy

	// CompositeModeDestination applies a composite based on: c_out = c_dst
	CompositeModeDestination

	// CompositeModeDestinationOver applies a composite based on: c_out = c_src × (1 - α_dst) + c_dst
	CompositeModeDestinationOver

	// CompositeModeSourceIn applies a composite based on: c_out = c_src × α_dst
	CompositeModeSourceIn

	// CompositeModeDestinationIn applies a composite based on: c_out = c_dst × α_src
	CompositeModeDestinationIn

	// CompositeModeSourceOut applies a composite based on: c_out = c_src × (1 - α_dst)
	CompositeModeSourceOut

	// CompositeModeDestinationOut applies a composite based on: c_out = c_dst × (1 - α_src)
	CompositeModeDestinationOut

	// CompositeModeSourceAtop applies a composite based on: c_out = c_src × α_dst + c_dst × (1 - α_src)
	CompositeModeSourceAtop

	// CompositeModeDestinationAtop applies a composite based on: c_out = c_src × (1 - α_dst) + c_dst × α_src
	CompositeModeDestinationAtop

	// CompositeModeXor applies a composite based on: c_out = c_src × (1 - α_dst) + c_dst × (1 - α_src)
	CompositeModeXor

	// CompositeModeLighter applies a composite based on:
	// c_out = c_src + c_dst Sum of source and destination (a.k.a. 'plus' or 'additive')
	CompositeModeLighter
)
