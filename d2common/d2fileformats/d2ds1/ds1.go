package d2ds1

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2path"
)

const (
	subType1 = 1
	subType2 = 2
)

const (
	wallZeroBitmask = 0xFFFFFF00
	wallZeroOffset  = 8
	wallTypeBitmask = 0x000000FF
)

const (
	unknown1BytesCount = 8
)

// DS1 represents the "stamp" data that is used to build up maps.
type DS1 struct {
	*ds1Layers
	Files              []string            // FilePtr table of file string pointers
	Objects            []Object            // Objects
	substitutionGroups []SubstitutionGroup // Substitution groups for the DS1

	version          ds1version
	Act              int32 // Act, from 1 to 5. This tells which Act table to use for the Objects list
	substitutionType int32 // SubstitutionType (layer type): 0 if no layer, else type 1 or type 2
	unknown1         []byte
	unknown2         uint32
}

const (
	defaultNumFloors        = 1
	defaultNumShadows       = maxShadowLayers
	defaultNumSubstitutions = 0
)

// Unmarshal the given bytes to a DS1 struct
func Unmarshal(fileData []byte) (*DS1, error) {
	return (&DS1{}).Unmarshal(fileData)
}

// Unmarshal the given bytes to a DS1 struct
func (ds1 *DS1) Unmarshal(fileData []byte) (*DS1, error) {
	ds1.ds1Layers = &ds1Layers{}

	stream := d2datautils.CreateStreamReader(fileData)

	if err := ds1.loadHeader(stream); err != nil {
		return nil, fmt.Errorf("loading header: %w", err)
	}

	if err := ds1.loadBody(stream); err != nil {
		return nil, fmt.Errorf("loading body: %w", err)
	}

	return ds1, nil
}

func (ds1 *DS1) loadHeader(br *d2datautils.StreamReader) error {
	var err error

	var width, height int32

	v, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading version: %w", err)
	}

	ds1.version = ds1version(v)

	width, err = br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading width: %w", err)
	}

	height, err = br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading height: %w", err)
	}

	width++
	height++

	ds1.SetSize(int(width), int(height))

	if ds1.version.specifiesAct() {
		ds1.Act, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading Act: %w", err)
		}

		ds1.Act = d2math.MinInt32(d2enum.ActsNumber, ds1.Act+1)
	}

	if ds1.version.specifiesSubstitutionType() {
		ds1.substitutionType, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution type: %w", err)
		}

		switch ds1.substitutionType {
		case subType1, subType2:
			ds1.PushSubstitution(&layer{})
		}
	}

	err = ds1.loadFileList(br)
	if err != nil {
		return fmt.Errorf("loading file list: %w", err)
	}

	return nil
}

func (ds1 *DS1) loadBody(stream *d2datautils.StreamReader) error {
	var numWalls, numFloors, numShadows, numSubstitutions int32

	numFloors = defaultNumFloors
	numShadows = defaultNumShadows
	numSubstitutions = defaultNumSubstitutions

	if ds1.version.hasUnknown1Bytes() {
		var err error

		ds1.unknown1, err = stream.ReadBytes(unknown1BytesCount)
		if err != nil {
			return fmt.Errorf("reading unknown1: %w", err)
		}
	}

	if ds1.version.specifiesWalls() {
		var err error

		numWalls, err = stream.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading wall number: %w", err)
		}

		if ds1.version.specifiesFloors() {
			numFloors, err = stream.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading number of Floors: %w", err)
			}
		}
	}

	for ; numWalls > 0; numWalls-- {
		ds1.PushWall(&layer{})
	}

	for ; numShadows > 0; numShadows-- {
		ds1.PushShadow(&layer{})
	}

	for ; numFloors > 0; numFloors-- {
		ds1.PushFloor(&layer{})
	}

	for ; numSubstitutions > 0; numSubstitutions-- {
		ds1.PushSubstitution(&layer{})
	}

	ds1.SetSize(ds1.width, ds1.height)

	if err := ds1.loadLayerStreams(stream); err != nil {
		return fmt.Errorf("loading layer streams: %w", err)
	}

	if err := ds1.loadObjects(stream); err != nil {
		return fmt.Errorf("loading Objects: %w", err)
	}

	if err := ds1.loadSubstitutions(stream); err != nil {
		return fmt.Errorf("loading Substitutions: %w", err)
	}

	if err := ds1.loadNPCs(stream); err != nil {
		return fmt.Errorf("loading npc's: %w", err)
	}

	return nil
}

func (ds1 *DS1) loadFileList(br *d2datautils.StreamReader) error {
	if !ds1.version.hasFileList() {
		return nil
	}

	// These Files reference things that don't exist anymore :-?
	numberOfFiles, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading number of Files: %w", err)
	}

	ds1.Files = make([]string, numberOfFiles)

	for i := 0; i < int(numberOfFiles); i++ {
		ds1.Files[i] = ""

		for {
			ch, err := br.ReadByte()
			if err != nil {
				return fmt.Errorf("reading file character: %w", err)
			}

			if ch == 0 {
				break
			}

			ds1.Files[i] += string(ch)
		}
	}

	return nil
}

func (ds1 *DS1) loadObjects(br *d2datautils.StreamReader) error {
	if !ds1.version.hasObjects() {
		ds1.Objects = make([]Object, 0)
		return nil
	}

	numObjects, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading number of Objects: %w", err)
	}

	ds1.Objects = make([]Object, numObjects)

	for objIdx := 0; objIdx < int(numObjects); objIdx++ {
		obj := Object{}

		objType, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading object's %d type: %v", objIdx, err)
		}

		objID, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading object's %d ID: %v", objIdx, err)
		}

		objX, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading object's %d X: %v", objIdx, err)
		}

		objY, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading object's %d Y: %v", objY, err)
		}

		objFlags, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading object's %d flags: %v", objIdx, err)
		}

		obj.Type = int(objType)
		obj.ID = int(objID)
		obj.X = int(objX)
		obj.Y = int(objY)
		obj.Flags = int(objFlags)

		ds1.Objects[objIdx] = obj
	}

	return nil
}

func (ds1 *DS1) loadSubstitutions(br *d2datautils.StreamReader) error {
	var err error

	hasSubstitutions := ds1.version.hasSubstitutions() &&
		(ds1.substitutionType == subType1 || ds1.substitutionType == subType2)

	if !hasSubstitutions {
		ds1.substitutionGroups = make([]SubstitutionGroup, 0)
		return nil
	}

	if ds1.version.hasUnknown2Bytes() {
		ds1.unknown2, err = br.ReadUInt32()
		if err != nil {
			return fmt.Errorf("reading unknown 2: %w", err)
		}
	}

	numberOfSubGroups, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading number of sub groups: %w", err)
	}

	ds1.substitutionGroups = make([]SubstitutionGroup, numberOfSubGroups)

	for subIdx := 0; subIdx < int(numberOfSubGroups); subIdx++ {
		newSub := SubstitutionGroup{}

		newSub.TileX, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d X: %v", subIdx, err)
		}

		newSub.TileY, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d Y: %v", subIdx, err)
		}

		newSub.WidthInTiles, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d W: %v", subIdx, err)
		}

		newSub.HeightInTiles, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d H: %v", subIdx, err)
		}

		newSub.Unknown, err = br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading substitution's %d unknown: %v", subIdx, err)
		}

		ds1.substitutionGroups[subIdx] = newSub
	}

	return err
}

func (ds1 *DS1) getLayerSchema() []layerStreamType {
	var layerStream []layerStreamType

	if ds1.version.hasStandardLayers() {
		layerStream = []layerStreamType{
			layerStreamWall1,
			layerStreamFloor1,
			layerStreamOrientation1,
			layerStreamSubstitute1,
			layerStreamShadow1,
		}

		return layerStream
	}

	numWalls := len(ds1.Walls)
	numOrientations := numWalls
	numFloors := len(ds1.Floors)
	numShadows := len(ds1.Shadows)
	numSubs := len(ds1.Substitutions)
	numLayers := numWalls + numOrientations + numFloors + numShadows + numSubs

	layerStream = make([]layerStreamType, numLayers)

	layerIdx := 0

	for i := 0; i < numWalls; i++ {
		layerStream[layerIdx] = layerStreamType(int(layerStreamWall1) + i)
		layerIdx++

		layerStream[layerIdx] = layerStreamType(int(layerStreamOrientation1) + i)
		layerIdx++
	}

	for i := 0; i < numFloors; i++ {
		layerStream[layerIdx] = layerStreamType(int(layerStreamFloor1) + i)
		layerIdx++
	}

	if numShadows > 0 {
		layerStream[layerIdx] = layerStreamShadow1
		layerIdx++
	}

	if numSubs > 0 {
		layerStream[layerIdx] = layerStreamSubstitute1
	}

	return layerStream
}

func (ds1 *DS1) loadNPCs(br *d2datautils.StreamReader) error {
	var err error

	if !ds1.version.specifiesNPCs() {
		return nil
	}

	numberOfNpcs, err := br.ReadInt32()
	if err != nil {
		return fmt.Errorf("reading number of npcs: %w", err)
	}

	for npcIdx := 0; npcIdx < int(numberOfNpcs); npcIdx++ {
		numPaths, err := br.ReadInt32() // nolint:govet // I want to re-use this error variable
		if err != nil {
			return fmt.Errorf("reading number of paths for npc %d: %v", npcIdx, err)
		}

		npcX, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading X pos for NPC %d: %v", npcIdx, err)
		}

		npcY, err := br.ReadInt32()
		if err != nil {
			return fmt.Errorf("reading Y pos for NPC %d: %v", npcIdx, err)
		}

		objIdx := -1

		for idx, ds1Obj := range ds1.Objects {
			if ds1Obj.X == int(npcX) && ds1Obj.Y == int(npcY) {
				objIdx = idx

				break
			}
		}

		if objIdx > -1 {
			err = ds1.loadNpcPaths(br, objIdx, int(numPaths))
			if err != nil {
				return fmt.Errorf("loading paths for NPC %d: %v", npcIdx, err)
			}
		} else {
			if ds1.version.specifiesNPCActions() {
				br.SkipBytes(int(numPaths) * 3) //nolint:gomnd // Unknown data
			} else {
				br.SkipBytes(int(numPaths) * 2) //nolint:gomnd // Unknown data
			}
		}
	}

	return err
}

func (ds1 *DS1) loadNpcPaths(br *d2datautils.StreamReader, objIdx, numPaths int) error {
	if ds1.Objects[objIdx].Paths == nil {
		ds1.Objects[objIdx].Paths = make([]d2path.Path, numPaths)
	}

	for pathIdx := 0; pathIdx < numPaths; pathIdx++ {
		newPath := d2path.Path{}

		px, err := br.ReadInt32() //nolint:govet // i want to re-use the err variable...
		if err != nil {
			return fmt.Errorf("reading X point for path %d: %v", pathIdx, err)
		}

		py, err := br.ReadInt32() //nolint:govet // i want to re-use the err variable...
		if err != nil {
			return fmt.Errorf("reading Y point for path %d: %v", pathIdx, err)
		}

		newPath.Position = d2vector.NewPosition(float64(px), float64(py))

		if ds1.version.specifiesNPCActions() {
			action, err := br.ReadInt32()
			if err != nil {
				return fmt.Errorf("reading action for path %d: %v", pathIdx, err)
			}

			newPath.Action = int(action)
		}

		ds1.Objects[objIdx].Paths[pathIdx] = newPath
	}

	return nil
}

func (ds1 *DS1) loadLayerStreams(br *d2datautils.StreamReader) error {
	dirLookup := []int32{
		0x00, 0x01, 0x02, 0x01, 0x02, 0x03, 0x03, 0x05, 0x05, 0x06,
		0x06, 0x07, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
		0x0F, 0x10, 0x11, 0x12, 0x14,
	}

	layerStreamTypes := ds1.getLayerSchema()

	for _, layerStreamType := range layerStreamTypes {
		for y := 0; y < ds1.height; y++ {
			for x := 0; x < ds1.width; x++ {
				dw, err := br.ReadUInt32()
				if err != nil {
					return fmt.Errorf("reading layer's dword: %w", err)
				}

				switch layerStreamType {
				case layerStreamWall1, layerStreamWall2, layerStreamWall3, layerStreamWall4:
					wallIndex := int(layerStreamType) - int(layerStreamWall1)
					ds1.Walls[wallIndex].Tile(x, y).DecodeWall(dw)
				case layerStreamOrientation1, layerStreamOrientation2,
					layerStreamOrientation3, layerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(layerStreamOrientation1)
					c := int32(dw & wallTypeBitmask)

					if ds1.version < v7 {
						if c < int32(len(dirLookup)) {
							c = dirLookup[c]
						}
					}

					tile := ds1.Walls[wallIndex].Tile(x, y)
					tile.Type = d2enum.TileType(c)
					tile.Zero = byte((dw & wallZeroBitmask) >> wallZeroOffset)
				case layerStreamFloor1, layerStreamFloor2:
					floorIndex := int(layerStreamType) - int(layerStreamFloor1)
					ds1.Floors[floorIndex].Tile(x, y).DecodeFloor(dw)
				case layerStreamShadow1:
					ds1.Shadows[0].Tile(x, y).DecodeShadow(dw)
				case layerStreamSubstitute1:
					ds1.Substitutions[0].Tile(x, y).Substitution = dw
				}
			}
		}
	}

	return nil
}

// SetSize sets the size of all layers in the DS1
func (ds1 *DS1) SetSize(w, h int) {
	ds1.ds1Layers.SetSize(w, h)
}

// Marshal encodes ds1 back to byte slice
func (ds1 *DS1) Marshal() []byte {
	// create stream writer
	sw := d2datautils.CreateStreamWriter()

	// Step 1 - encode header
	sw.PushInt32(int32(ds1.version))
	sw.PushInt32(int32(ds1.width - 1))
	sw.PushInt32(int32(ds1.height - 1))

	if ds1.version.specifiesAct() {
		sw.PushInt32(ds1.Act - 1)
	}

	if ds1.version.specifiesSubstitutionType() {
		sw.PushInt32(ds1.substitutionType)
	}

	if ds1.version.hasFileList() {
		sw.PushInt32(int32(len(ds1.Files)))

		for _, i := range ds1.Files {
			sw.PushBytes([]byte(i)...)

			// separator
			sw.PushBytes(0)
		}
	}

	if ds1.version.hasUnknown1Bytes() {
		sw.PushBytes(ds1.unknown1...)
	}

	if ds1.version.specifiesWalls() {
		sw.PushInt32(int32(len(ds1.Walls)))

		if ds1.version.specifiesFloors() {
			sw.PushInt32(int32(len(ds1.Walls)))
		}
	}

	// Step 2 - encode grid
	ds1.encodeLayers(sw)

	// Step 3 - encode Objects
	if ds1.version.hasObjects() {
		sw.PushInt32(int32(len(ds1.Objects)))

		for _, i := range ds1.Objects {
			sw.PushUint32(uint32(i.Type))
			sw.PushUint32(uint32(i.ID))
			sw.PushUint32(uint32(i.X))
			sw.PushUint32(uint32(i.Y))
			sw.PushUint32(uint32(i.Flags))
		}
	}

	// Step 4 - encode Substitutions
	hasSubstitutions := ds1.version.hasSubstitutions() &&
		(ds1.substitutionType == subType1 || ds1.substitutionType == subType2)

	if hasSubstitutions {
		sw.PushUint32(ds1.unknown2)

		sw.PushUint32(uint32(len(ds1.substitutionGroups)))

		for _, i := range ds1.substitutionGroups {
			sw.PushInt32(i.TileX)
			sw.PushInt32(i.TileY)
			sw.PushInt32(i.WidthInTiles)
			sw.PushInt32(i.HeightInTiles)
			sw.PushInt32(i.Unknown)
		}
	}

	// Step 5 - encode NPC's and its paths
	ds1.encodeNPCs(sw)

	return sw.GetBytes()
}

func (ds1 *DS1) encodeLayers(sw *d2datautils.StreamWriter) {
	layerStreamTypes := ds1.getLayerSchema()

	for _, layerStreamType := range layerStreamTypes {
		for y := 0; y < ds1.height; y++ {
			for x := 0; x < ds1.width; x++ {
				dw := uint32(0)

				switch layerStreamType {
				case layerStreamWall1, layerStreamWall2, layerStreamWall3, layerStreamWall4:
					wallIndex := int(layerStreamType) - int(layerStreamWall1)
					ds1.Walls[wallIndex].Tile(x, y).EncodeWall(sw)
				case layerStreamOrientation1, layerStreamOrientation2,
					layerStreamOrientation3, layerStreamOrientation4:
					wallIndex := int(layerStreamType) - int(layerStreamOrientation1)
					dw |= uint32(ds1.Walls[wallIndex].Tile(x, y).Type)
					dw |= (uint32(ds1.Walls[wallIndex].Tile(x, y).Zero) & wallZeroBitmask) << wallZeroOffset

					sw.PushUint32(dw)
				case layerStreamFloor1, layerStreamFloor2:
					floorIndex := int(layerStreamType) - int(layerStreamFloor1)
					ds1.Floors[floorIndex].Tile(x, y).EncodeFloor(sw)
				case layerStreamShadow1:
					ds1.Shadows[0].Tile(x, y).EncodeShadow(sw)
				case layerStreamSubstitute1:
					sw.PushUint32(ds1.Substitutions[0].Tile(x, y).Substitution)
				}
			}
		}
	}
}

func (ds1 *DS1) encodeNPCs(sw *d2datautils.StreamWriter) {
	objectsWithPaths := make([]int, 0)

	for n, obj := range ds1.Objects {
		if len(obj.Paths) != 0 {
			objectsWithPaths = append(objectsWithPaths, n)
		}
	}

	// Step 5.1 - encode npc's
	sw.PushUint32(uint32(len(objectsWithPaths)))

	// Step 5.2 - enoce npcs' paths
	for objectIdx := range objectsWithPaths {
		sw.PushUint32(uint32(len(ds1.Objects[objectIdx].Paths)))
		sw.PushUint32(uint32(ds1.Objects[objectIdx].X))
		sw.PushUint32(uint32(ds1.Objects[objectIdx].Y))

		for _, path := range ds1.Objects[objectIdx].Paths {
			sw.PushUint32(uint32(path.Position.X()))
			sw.PushUint32(uint32(path.Position.Y()))

			if ds1.version >= v15 {
				sw.PushUint32(uint32(path.Action))
			}
		}
	}
}

// Version returns the ds1 version
func (ds1 *DS1) Version() int {
	return int(ds1.version)
}

// SetVersion sets the ds1 version, can not be negative.
func (ds1 *DS1) SetVersion(v int) {
	if v < 0 {
		v = 0
	}

	ds1.version = ds1version(v)
}
