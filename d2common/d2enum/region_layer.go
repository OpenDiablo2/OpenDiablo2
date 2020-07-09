package d2enum

// RegionLayerType ...
type RegionLayerType int

// Region layer types
const (
	RegionLayerTypeFloors RegionLayerType = iota
	RegionLayerTypeWalls
	RegionLayerTypeShadows
)
