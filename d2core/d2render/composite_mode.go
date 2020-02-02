package d2render

type CompositeMode int

const (
	// Regular alpha blending
	// c_out = c_src + c_dst × (1 - α_src)
	CompositeModeSourceOver CompositeMode = CompositeMode(1)

	// c_out = 0
	CompositeModeClear CompositeMode = CompositeMode(2)

	// c_out = c_src
	CompositeModeCopy CompositeMode = CompositeMode(3)

	// c_out = c_dst
	CompositeModeDestination CompositeMode = CompositeMode(4)

	// c_out = c_src × (1 - α_dst) + c_dst
	CompositeModeDestinationOver CompositeMode = CompositeMode(5)

	// c_out = c_src × α_dst
	CompositeModeSourceIn CompositeMode = CompositeMode(6)

	// c_out = c_dst × α_src
	CompositeModeDestinationIn CompositeMode = CompositeMode(7)

	// c_out = c_src × (1 - α_dst)
	CompositeModeSourceOut CompositeMode = CompositeMode(8)

	// c_out = c_dst × (1 - α_src)
	CompositeModeDestinationOut CompositeMode = CompositeMode(9)

	// c_out = c_src × α_dst + c_dst × (1 - α_src)
	CompositeModeSourceAtop CompositeMode = CompositeMode(10)

	// c_out = c_src × (1 - α_dst) + c_dst × α_src
	CompositeModeDestinationAtop CompositeMode = CompositeMode(11)

	// c_out = c_src × (1 - α_dst) + c_dst × (1 - α_src)
	CompositeModeXor CompositeMode = CompositeMode(12)

	// Sum of source and destination (a.k.a. 'plus' or 'additive')
	// c_out = c_src + c_dst
	CompositeModeLighter CompositeMode = CompositeMode(13)
)
