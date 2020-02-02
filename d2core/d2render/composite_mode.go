package d2render

type CompositeMode int

const (
	// Regular alpha blending
	// c_out = c_src + c_dst × (1 - α_src)
	CompositeModeSourceOver = CompositeMode(1)

	// c_out = 0
	CompositeModeClear = CompositeMode(2)

	// c_out = c_src
	CompositeModeCopy = CompositeMode(3)

	// c_out = c_dst
	CompositeModeDestination = CompositeMode(4)

	// c_out = c_src × (1 - α_dst) + c_dst
	CompositeModeDestinationOver = CompositeMode(5)

	// c_out = c_src × α_dst
	CompositeModeSourceIn = CompositeMode(6)

	// c_out = c_dst × α_src
	CompositeModeDestinationIn = CompositeMode(7)

	// c_out = c_src × (1 - α_dst)
	CompositeModeSourceOut = CompositeMode(8)

	// c_out = c_dst × (1 - α_src)
	CompositeModeDestinationOut = CompositeMode(9)

	// c_out = c_src × α_dst + c_dst × (1 - α_src)
	CompositeModeSourceAtop = CompositeMode(10)

	// c_out = c_src × (1 - α_dst) + c_dst × α_src
	CompositeModeDestinationAtop = CompositeMode(11)

	// c_out = c_src × (1 - α_dst) + c_dst × (1 - α_src)
	CompositeModeXor = CompositeMode(12)

	// Sum of source and destination (a.k.a. 'plus' or 'additive')
	// c_out = c_src + c_dst
	CompositeModeLighter = CompositeMode(13)
)
