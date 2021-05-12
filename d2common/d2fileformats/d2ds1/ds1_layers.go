package d2ds1

const (
	maxWallLayers         = 4
	maxFloorLayers        = 2
	maxShadowLayers       = 1
	maxSubstitutionLayers = 1
)

// LayerGroupType represents a type of layer (floor, wall, shadow, etc)
type LayerGroupType int

// Layer group types
const (
	FloorLayerGroup LayerGroupType = iota
	WallLayerGroup
	ShadowLayerGroup
	SubstitutionLayerGroup
)

func (l LayerGroupType) String() string {
	switch l {
	case FloorLayerGroup:
		return "floor"
	case WallLayerGroup:
		return "wall"
	case ShadowLayerGroup:
		return "shadow"
	case SubstitutionLayerGroup:
		return "substitution"
	}

	// should not be reached
	return "unknown"
}

type layerGroup []*Layer

type ds1Layers struct {
	width, height int
	Floors        layerGroup
	Walls         layerGroup
	Shadows       layerGroup
	Substitutions layerGroup
}

func (l *ds1Layers) ensureInit() {
	if l.Floors == nil {
		l.Floors = make(layerGroup, 0)
	}

	if l.Walls == nil {
		l.Walls = make(layerGroup, 0)
	}

	if l.Shadows == nil {
		l.Shadows = make(layerGroup, 0)
	}

	if l.Substitutions == nil {
		l.Substitutions = make(layerGroup, 0)
	}
}

// removes nil layers from all layer groups
func (l *ds1Layers) cull() {
	l.cullNilLayers(FloorLayerGroup)
	l.cullNilLayers(WallLayerGroup)
	l.cullNilLayers(ShadowLayerGroup)
	l.cullNilLayers(SubstitutionLayerGroup)
}

// removes nil layers of given layer group type
func (l *ds1Layers) cullNilLayers(t LayerGroupType) {
	group := l.GetLayersGroup(t)
	if group == nil {
		return
	}

	// from last to first layer, remove first encountered nil layer and restart the culling procedure.
	// exit culling procedure when no nil layers are found in entire group.
culling:
	for {
		for idx := len(*group) - 1; idx >= 0; idx-- {
			if (*group)[idx] == nil {
				*group = append((*group)[:idx], (*group)[idx+1:]...)
				continue culling
			}
		}

		break culling // encountered no new nil layers
	}
}

func (l *ds1Layers) Size() (w, h int) {
	l.ensureInit()
	l.cull()

	return l.width, l.height
}

func (l *ds1Layers) SetSize(w, h int) {
	l.width, l.height = w, h

	l.enforceSize(FloorLayerGroup)
	l.enforceSize(WallLayerGroup)
	l.enforceSize(ShadowLayerGroup)
	l.enforceSize(SubstitutionLayerGroup)
}

func (l *ds1Layers) enforceSize(t LayerGroupType) {
	l.ensureInit()
	l.cull()

	group := l.GetLayersGroup(t)
	if group == nil {
		return
	}

	for idx := range *group {
		(*group)[idx].SetSize(l.width, l.height)
	}
}

func (l *ds1Layers) Width() int {
	w, _ := l.Size()
	return w
}

func (l *ds1Layers) SetWidth(w int) {
	l.SetSize(w, l.height)
}

func (l *ds1Layers) Height() int {
	_, h := l.Size()
	return h
}

func (l *ds1Layers) SetHeight(h int) {
	l.SetSize(l.width, h)
}

// generic push func for all layer types
func (l *ds1Layers) push(t LayerGroupType, layer *Layer) {
	l.ensureInit()
	l.cull()
	layer.SetSize(l.Size())

	group := l.GetLayersGroup(t)

	max := GetMaxGroupLen(t)

	if len(*group) < max {
		*group = append(*group, layer)
	}
}

// generic pop func for all layer types
func (l *ds1Layers) pop(t LayerGroupType) *Layer {
	l.ensureInit()
	l.cull()

	group := l.GetLayersGroup(t)
	if group == nil {
		return nil
	}

	var theLayer *Layer

	// remove last layer of slice and return it
	if len(*group) > 0 {
		lastIdx := len(*group) - 1
		theLayer = (*group)[lastIdx]
		*group = (*group)[:lastIdx]

		return theLayer
	}

	return nil
}

func (l *ds1Layers) get(t LayerGroupType, idx int) *Layer {
	l.ensureInit()
	l.cull()

	group := l.GetLayersGroup(t)
	if group == nil {
		return nil
	}

	if idx >= len(*group) || idx < 0 {
		return nil
	}

	return (*group)[idx]
}

func (l *ds1Layers) insert(t LayerGroupType, idx int, newLayer *Layer) {
	l.ensureInit()
	l.cull()

	if newLayer == nil {
		return
	}

	newLayer.SetSize(l.Size())

	group := l.GetLayersGroup(t)
	if group == nil {
		return
	}

	if len(*group)+1 > GetMaxGroupLen(t) {
		return
	}

	if len(*group) == 0 {
		*group = append(*group, newLayer) // nolint:staticcheck // we possibly use group later
		return
	}

	if l := len(*group) - 1; idx > l {
		idx = l
	}

	// example:
	// suppose
	//		idx=1
	//		newLayer=c
	// 		existing layerGroup is [a, b]
	*group = append((*group)[:idx], append([]*Layer{newLayer}, (*group)[idx:]...)...)
}

func (l *ds1Layers) delete(t LayerGroupType, idx int) {
	l.ensureInit()
	l.cull()

	group := l.GetLayersGroup(t)
	if group == nil {
		return
	}

	if idx >= len(*group) || idx < 0 {
		return
	}

	(*group)[idx] = nil

	l.cull()
}

func (l *ds1Layers) GetFloor(idx int) *Layer {
	return l.get(FloorLayerGroup, idx)
}

func (l *ds1Layers) PushFloor(floor *Layer) *ds1Layers {
	l.push(FloorLayerGroup, floor)
	return l
}

func (l *ds1Layers) PopFloor() *Layer {
	return l.pop(FloorLayerGroup)
}

func (l *ds1Layers) InsertFloor(idx int, newFloor *Layer) {
	l.insert(FloorLayerGroup, idx, newFloor)
}

func (l *ds1Layers) DeleteFloor(idx int) {
	l.delete(FloorLayerGroup, idx)
}

func (l *ds1Layers) GetWall(idx int) *Layer {
	return l.get(WallLayerGroup, idx)
}

func (l *ds1Layers) PushWall(wall *Layer) *ds1Layers {
	l.push(WallLayerGroup, wall)
	return l
}

func (l *ds1Layers) PopWall() *Layer {
	return l.pop(WallLayerGroup)
}

func (l *ds1Layers) InsertWall(idx int, newWall *Layer) {
	l.insert(WallLayerGroup, idx, newWall)
}

func (l *ds1Layers) DeleteWall(idx int) {
	l.delete(WallLayerGroup, idx)
}

func (l *ds1Layers) GetShadow(idx int) *Layer {
	return l.get(ShadowLayerGroup, idx)
}

func (l *ds1Layers) PushShadow(shadow *Layer) *ds1Layers {
	l.push(ShadowLayerGroup, shadow)
	return l
}

func (l *ds1Layers) PopShadow() *Layer {
	return l.pop(ShadowLayerGroup)
}

func (l *ds1Layers) InsertShadow(idx int, newShadow *Layer) {
	l.insert(ShadowLayerGroup, idx, newShadow)
}

func (l *ds1Layers) DeleteShadow(idx int) {
	l.delete(ShadowLayerGroup, idx)
}

func (l *ds1Layers) GetSubstitution(idx int) *Layer {
	return l.get(SubstitutionLayerGroup, idx)
}

func (l *ds1Layers) PushSubstitution(sub *Layer) *ds1Layers {
	l.push(SubstitutionLayerGroup, sub)
	return l
}

func (l *ds1Layers) PopSubstitution() *Layer {
	return l.pop(SubstitutionLayerGroup)
}

func (l *ds1Layers) InsertSubstitution(idx int, newSubstitution *Layer) {
	l.insert(SubstitutionLayerGroup, idx, newSubstitution)
}

func (l *ds1Layers) DeleteSubstitution(idx int) {
	l.delete(ShadowLayerGroup, idx)
}

// GetLayersGroup returns layer group depending on type given
func (l *ds1Layers) GetLayersGroup(t LayerGroupType) (group *layerGroup) {
	switch t {
	case FloorLayerGroup:
		group = &l.Floors
	case WallLayerGroup:
		group = &l.Walls
	case ShadowLayerGroup:
		group = &l.Shadows
	case SubstitutionLayerGroup:
		group = &l.Substitutions
	default:
		return nil
	}

	return group
}

// GetMaxGroupLen returns maximum length of layer group of type given
func GetMaxGroupLen(t LayerGroupType) (max int) {
	switch t {
	case FloorLayerGroup:
		max = maxFloorLayers
	case WallLayerGroup:
		max = maxWallLayers
	case ShadowLayerGroup:
		max = maxShadowLayers
	case SubstitutionLayerGroup:
		max = maxSubstitutionLayers
	default:
		return 0
	}

	return max
}
