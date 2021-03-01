package d2ds1

const (
	maxWallLayers         = 4
	maxOrientationLayers  = 4
	maxFloorLayers        = 2
	maxShadowLayers       = 1
	maxSubstitutionLayers = 1
)

type layerGroupType int

const (
	floorLayerGroup layerGroupType = iota
	wallLayerGroup
	orientationLayerGroup
	shadowLayerGroup
	substitutionLayerGroup
)

type layerGroup []*layer

type ds1Layers struct {
	width, height int
	Floors        layerGroup
	Walls         layerGroup
	Orientations  layerGroup
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

	if l.Orientations == nil {
		l.Orientations = make(layerGroup, 0)
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
	l.cullNilLayers(orientationLayerGroup)
	l.cullNilLayers(shadowLayerGroup)
	l.cullNilLayers(substitutionLayerGroup)
}

// removes nil layers of given layer group type
func (l *ds1Layers) cullNilLayers(t layerGroupType) {
	var group *layerGroup

	switch t {
	case floorLayerGroup:
		group = &l.Floors
	case wallLayerGroup:
		group = &l.Walls
	case orientationLayerGroup:
		group = &l.Orientations
	case shadowLayerGroup:
		group = &l.Shadows
	case substitutionLayerGroup:
		group = &l.Substitutions
	default:
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
	l.enforceSize(orientationLayerGroup)
	l.enforceSize(shadowLayerGroup)
	l.enforceSize(substitutionLayerGroup)
}

func (l *ds1Layers) enforceSize(t layerGroupType) {
	l.ensureInit()
	l.cull()

	var group *layerGroup

	switch t {
	case floorLayerGroup:
		group = &l.Floors
	case wallLayerGroup:
		group = &l.Walls
	case orientationLayerGroup:
		group = &l.Orientations
	case shadowLayerGroup:
		group = &l.Shadows
	case substitutionLayerGroup:
		group = &l.Substitutions
	default:
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
func (l *ds1Layers) push(t layerGroupType, layer *layer) {
	l.ensureInit()
	l.cull()

	var group *layerGroup

	var max int

	switch t {
	case floorLayerGroup:
		group = &l.Floors
		max = maxFloorLayers
	case wallLayerGroup:
		group = &l.Walls
		max = maxWallLayers
	case orientationLayerGroup:
		group = &l.Orientations
		max = maxOrientationLayers
	case shadowLayerGroup:
		group = &l.Shadows
		max = maxShadowLayers
	case substitutionLayerGroup:
		group = &l.Substitutions
		max = maxSubstitutionLayers
	default:
		return
	}

	if len(*group) < max {
		*group = append(*group, layer)
	}
}

// generic pop func for all layer types
func (l *ds1Layers) pop(t layerGroupType) *layer {
	l.ensureInit()
	l.cull()

	var group *layerGroup

	var theLayer *layer

	switch t {
	case floorLayerGroup:
		group = &l.Floors
	case wallLayerGroup:
		group = &l.Walls
	case orientationLayerGroup:
		group = &l.Orientations
	case shadowLayerGroup:
		group = &l.Shadows
	case substitutionLayerGroup:
		group = &l.Substitutions
	default:
		return nil
	}

	// remove last layer of slice and return it
	if len(*group) > 0 {
		lastIdx := len(*group) - 1
		theLayer = (*group)[lastIdx]
		*group = (*group)[:lastIdx]

		return theLayer
	}

	return nil
}

func (l *ds1Layers) get(t layerGroupType, idx int) *layer {
	l.ensureInit()
	l.cull()

	var group *layerGroup

	switch t {
	case floorLayerGroup:
		group = &l.Floors
	case wallLayerGroup:
		group = &l.Walls
	case orientationLayerGroup:
		group = &l.Orientations
	case shadowLayerGroup:
		group = &l.Shadows
	case substitutionLayerGroup:
		group = &l.Substitutions
	default:
		return nil
	}

	if idx >= len(*group) || idx < 0 {
		return nil
	}

	return (*group)[idx]
}

func (l *ds1Layers) insert(t layerGroupType, idx int, newLayer *layer) {
	l.ensureInit()
	l.cull()

	if newLayer == nil {
		return
	}

	var group layerGroup

	switch t {
	case floorLayerGroup:
		group = l.Floors
	case wallLayerGroup:
		group = l.Walls
	case orientationLayerGroup:
		group = l.Orientations
	case shadowLayerGroup:
		group = l.Shadows
	case substitutionLayerGroup:
		group = l.Substitutions
	default:
		return
	}

	if len(group) == 0 {
		group = append(group, newLayer) // nolint:staticcheck // we possibly use group later
		return
	}

	if idx > len(group)-1 {
		idx = len(group) - 1
	}

	// example:
	// suppose
	//		idx=1
	//		newLayer=c
	// 		existing layerGroup is [a, b]
	group = append(group, group[idx:]...) // [a, b] becomes [a, b, b]
	group[idx] = newLayer                 // [a, b, b] becomes [a, c, b]
}

func (l *ds1Layers) delete(t layerGroupType, idx int) {
	l.ensureInit()
	l.cull()

	var group layerGroup

	switch t {
	case floorLayerGroup:
		group = l.Floors
	case wallLayerGroup:
		group = l.Walls
	case orientationLayerGroup:
		group = l.Orientations
	case shadowLayerGroup:
		group = l.Shadows
	case substitutionLayerGroup:
		group = l.Substitutions
	default:
		return
	}

	if idx >= len(group) || idx < 0 {
		return
	}

	group[idx] = nil

	l.cull()
}

func (l *ds1Layers) GetFloor(idx int) *layer {
	return l.get(floorLayerGroup, idx)
}

func (l *ds1Layers) PushFloor(floor *layer) *ds1Layers {
	l.push(floorLayerGroup, floor)
	return l
}

func (l *ds1Layers) PopFloor() *layer {
	return l.pop(floorLayerGroup)
}

func (l *ds1Layers) InsertFloor(idx int, newFloor *layer) {
	l.insert(floorLayerGroup, idx, newFloor)
}

func (l *ds1Layers) DeleteFloor(idx int) {
	l.delete(floorLayerGroup, idx)
}

func (l *ds1Layers) GetWall(idx int) *layer {
	return l.get(wallLayerGroup, idx)
}

func (l *ds1Layers) PushWall(wall *layer) *ds1Layers {
	l.push(wallLayerGroup, wall)
	return l
}

func (l *ds1Layers) PopWall() *layer {
	return l.pop(wallLayerGroup)
}

func (l *ds1Layers) InsertWall(idx int, newWall *layer) {
	l.insert(wallLayerGroup, idx, newWall)
}

func (l *ds1Layers) DeleteWall(idx int) {
	l.delete(wallLayerGroup, idx)
}

func (l *ds1Layers) GetOrientation(idx int) *layer {
	return l.get(orientationLayerGroup, idx)
}

func (l *ds1Layers) PushOrientation(orientation *layer) *ds1Layers {
	l.push(orientationLayerGroup, orientation)
	return l
}

func (l *ds1Layers) PopOrientation() *layer {
	return l.pop(orientationLayerGroup)
}

func (l *ds1Layers) InsertOrientation(idx int, newOrientation *layer) {
	l.insert(orientationLayerGroup, idx, newOrientation)
}

func (l *ds1Layers) DeleteOrientation(idx int) {
	l.delete(orientationLayerGroup, idx)
}

func (l *ds1Layers) GetShadow(idx int) *layer {
	return l.get(shadowLayerGroup, idx)
}

func (l *ds1Layers) PushShadow(shadow *layer) *ds1Layers {
	l.push(shadowLayerGroup, shadow)
	return l
}

func (l *ds1Layers) PopShadow() *layer {
	return l.pop(shadowLayerGroup)
}

func (l *ds1Layers) InsertShadow(idx int, newShadow *layer) {
	l.insert(shadowLayerGroup, idx, newShadow)
}

func (l *ds1Layers) DeleteShadow(idx int) {
	l.delete(shadowLayerGroup, idx)
}

func (l *ds1Layers) GetSubstitution(idx int) *layer {
	return l.get(substitutionLayerGroup, idx)
}

func (l *ds1Layers) PushSubstitution(sub *layer) *ds1Layers {
	l.push(substitutionLayerGroup, sub)
	return l
}

func (l *ds1Layers) PopSubstitution() *layer {
	return l.pop(substitutionLayerGroup)
}

func (l *ds1Layers) InsertSubstitution(idx int, newSubstitution *layer) {
	l.insert(shadowLayerGroup, idx, newSubstitution)
}

func (l *ds1Layers) DeleteSubstitution(idx int) {
	l.delete(shadowLayerGroup, idx)
}
