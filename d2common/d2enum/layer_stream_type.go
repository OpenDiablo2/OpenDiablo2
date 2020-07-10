package d2enum

// LayerStreamType represents a layer stream type
type LayerStreamType int

// Layer stream types
const (
	LayerStreamWall1 LayerStreamType = iota
	LayerStreamWall2
	LayerStreamWall3
	LayerStreamWall4
	LayerStreamOrientation1
	LayerStreamOrientation2
	LayerStreamOrientation3
	LayerStreamOrientation4
	LayerStreamFloor1
	LayerStreamFloor2
	LayerStreamShadow
	LayerStreamSubstitute
)
