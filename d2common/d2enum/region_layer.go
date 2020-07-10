package d2enum

// RegionLayerType represents a region layer
type RegionLayerType int

// Region layer types
const (
	RegionLayerTypeFloors RegionLayerType = iota
	RegionLayerTypeWalls
	RegionLayerTypeShadows
)
