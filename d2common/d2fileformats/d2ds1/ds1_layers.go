package d2ds1

const (
	maxWallLayers         = 4
	maxFloorLayers        = 2
	maxShadowLayers       = 1
	maxSubstitutionLayers = 1
)

type layerGroupType int

const (
	floorLayerGroup layerGroupType = iota
	wallLayerGroup
	shadowLayerGroup
	substitutionLayerGroup
)

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
	l.cullNilLayers(floorLayerGroup)
	l.cullNilLayers(wallLayerGroup)
	l.cullNilLayers(shadowLayerGroup)
	l.cullNilLayers(substitutionLayerGroup)
}

// removes nil layers of given layer group type
func (l *ds1Layers) cullNilLayers(t layerGroupType) {
	group := l.getLayersGroup(t)
	if group == nil {
		return
	}

	// from last to first layer, remove first encountered nil layer and restart the culling procedure.
	// exit culling procedure when no nil layers are found in entire group.
culling:
	for {
		for idx := len(*group) - 1; idx > 0; idx-- {
			if (*group)[idx] == nil {
				*group = (*group)[:idx]
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

	l.enforceSize(floorLayerGroup)
	l.enforceSize(wallLayerGroup)
	l.enforceSize(shadowLayerGroup)
	l.enforceSize(substitutionLayerGroup)
}

func (l *ds1Layers) enforceSize(t layerGroupType) {
	l.ensureInit()
	l.cull()

	group := l.getLayersGroup(t)
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
func (l *ds1Layers) push(t layerGroupType, layer *Layer) {
	l.ensureInit()
	l.cull()
	layer.SetSize(l.Size())

	group := l.getLayersGroup(t)

	max := getMaxGroupLen(t)

	if len(*group) < max {
		*group = append(*group, layer)
	}
}

// generic pop func for all layer types
func (l *ds1Layers) pop(t layerGroupType) *Layer {
	l.ensureInit()
	l.cull()

	group := l.getLayersGroup(t)
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

func (l *ds1Layers) get(t layerGroupType, idx int) *Layer {
	l.ensureInit()
	l.cull()

	group := l.getLayersGroup(t)
	if group == nil {
		return nil
	}

	if idx >= len(*group) || idx < 0 {
		return nil
	}

	return (*group)[idx]
}

func (l *ds1Layers) insert(t layerGroupType, idx int, newLayer *Layer) {
	l.ensureInit()
	l.cull()

	if newLayer == nil {
		return
	}

	newLayer.SetSize(l.Size())

	group := l.getLayersGroup(t)
	if group == nil {
		return
	}

	if len(*group)+1 > getMaxGroupLen(t) {
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
	newGroup := append((*group)[:idx], append([]*Layer{newLayer}, (*group)[idx:]...)...)
	*group = newGroup
}

func (l *ds1Layers) delete(t layerGroupType, idx int) {
	l.ensureInit()
	l.cull()

	group := l.getLayersGroup(t)
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
	return l.get(floorLayerGroup, idx)
}

func (l *ds1Layers) PushFloor(floor *Layer) *ds1Layers {
	l.push(floorLayerGroup, floor)
	return l
}

func (l *ds1Layers) PopFloor() *Layer {
	return l.pop(floorLayerGroup)
}

func (l *ds1Layers) InsertFloor(idx int, newFloor *Layer) {
	l.insert(floorLayerGroup, idx, newFloor)
}

func (l *ds1Layers) DeleteFloor(idx int) {
	l.delete(floorLayerGroup, idx)
}

func (l *ds1Layers) GetWall(idx int) *Layer {
	return l.get(wallLayerGroup, idx)
}

func (l *ds1Layers) PushWall(wall *Layer) *ds1Layers {
	l.push(wallLayerGroup, wall)
	return l
}

func (l *ds1Layers) PopWall() *Layer {
	return l.pop(wallLayerGroup)
}

func (l *ds1Layers) InsertWall(idx int, newWall *Layer) {
	l.insert(wallLayerGroup, idx, newWall)
}

func (l *ds1Layers) DeleteWall(idx int) {
	l.delete(wallLayerGroup, idx)
}

func (l *ds1Layers) GetShadow(idx int) *Layer {
	return l.get(shadowLayerGroup, idx)
}

func (l *ds1Layers) PushShadow(shadow *Layer) *ds1Layers {
	l.push(shadowLayerGroup, shadow)
	return l
}

func (l *ds1Layers) PopShadow() *Layer {
	return l.pop(shadowLayerGroup)
}

func (l *ds1Layers) InsertShadow(idx int, newShadow *Layer) {
	l.insert(shadowLayerGroup, idx, newShadow)
}

func (l *ds1Layers) DeleteShadow(idx int) {
	l.delete(shadowLayerGroup, idx)
}

func (l *ds1Layers) GetSubstitution(idx int) *Layer {
	return l.get(substitutionLayerGroup, idx)
}

func (l *ds1Layers) PushSubstitution(sub *Layer) *ds1Layers {
	l.push(substitutionLayerGroup, sub)
	return l
}

func (l *ds1Layers) PopSubstitution() *Layer {
	return l.pop(substitutionLayerGroup)
}

func (l *ds1Layers) InsertSubstitution(idx int, newSubstitution *Layer) {
	l.insert(substitutionLayerGroup, idx, newSubstitution)
}

func (l *ds1Layers) DeleteSubstitution(idx int) {
	l.delete(shadowLayerGroup, idx)
}

func (l *ds1Layers) getLayersGroup(t layerGroupType) (group *layerGroup) {
	switch t {
	case floorLayerGroup:
		group = &l.Floors
	case wallLayerGroup:
		group = &l.Walls
	case shadowLayerGroup:
		group = &l.Shadows
	case substitutionLayerGroup:
		group = &l.Substitutions
	default:
		return nil
	}

	return group
}

func getMaxGroupLen(t layerGroupType) (max int) {
	switch t {
	case floorLayerGroup:
		max = maxFloorLayers
	case wallLayerGroup:
		max = maxWallLayers
	case shadowLayerGroup:
		max = maxShadowLayers
	case substitutionLayerGroup:
		max = maxSubstitutionLayers
	default:
		return 0
	}

	return max
}
