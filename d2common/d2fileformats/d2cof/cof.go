package d2cof

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	unknownByteCount = 21
	numHeaderBytes   = 4 + unknownByteCount
	numLayerBytes    = 9
)

const (
	headerNumLayers = iota
	headerFramesPerDir
	headerNumDirs
	headerSpeed = numHeaderBytes - 1
)

const (
	layerType = iota
	layerShadow
	layerSelectable
	layerTransparent
	layerDrawEffect
	layerWeaponClass
)

const (
	badCharacter = string(byte(0))
)

// COF is a structure that represents a COF file.
type COF struct {
	// unknown bytes for header
	unknownHeaderBytes []byte
	// unknown bytes (first "body's" bytes)
	unknown1           []byte
	NumberOfDirections int
	FramesPerDirection int
	NumberOfLayers     int
	Speed              int
	CofLayers          []CofLayer
	CompositeLayers    map[d2enum.CompositeType]int
	AnimationFrames    []d2enum.AnimationFrame
	Priority           [][][]d2enum.CompositeType
}

// Load loads a COF file.
func Load(fileData []byte) (*COF, error) {
	result := &COF{}
	streamReader := d2datautils.CreateStreamReader(fileData)

	var b []byte

	var err error

	b, err = streamReader.ReadBytes(numHeaderBytes)
	if err != nil {
		return nil, err
	}

	result.NumberOfLayers = int(b[headerNumLayers])
	result.FramesPerDirection = int(b[headerFramesPerDir])
	result.NumberOfDirections = int(b[headerNumDirs])
	result.unknownHeaderBytes = b[headerNumDirs+1 : headerSpeed]
	result.Speed = int(b[headerSpeed])

	// read unknown bytes
	// previous streamReader.SkipBytes(3)
	for i := 0; i < 3; i++ {
		b, errSR := streamReader.ReadByte()
		if errSR != nil {
			return nil, errSR
		}

		result.unknown1 = append(result.unknown1, b)
	}

	result.CofLayers = make([]CofLayer, result.NumberOfLayers)
	result.CompositeLayers = make(map[d2enum.CompositeType]int)

	for i := 0; i < result.NumberOfLayers; i++ {
		layer := CofLayer{}

		b, err = streamReader.ReadBytes(numLayerBytes)
		if err != nil {
			return nil, err
		}

		layer.Type = d2enum.CompositeType(b[layerType])
		layer.Shadow = b[layerShadow]
		layer.Selectable = b[layerSelectable] > 0
		layer.Transparent = b[layerTransparent] > 0
		layer.DrawEffect = d2enum.DrawEffect(b[layerDrawEffect])

		layer.weaponClassByte = b[layerWeaponClass:]
		layer.WeaponClass = d2enum.WeaponClassFromString(strings.TrimSpace(strings.ReplaceAll(
			string(b[layerWeaponClass:]), badCharacter, "")))

		result.CofLayers[i] = layer
		result.CompositeLayers[layer.Type] = i
	}

	b, err = streamReader.ReadBytes(result.FramesPerDirection)
	if err != nil {
		return nil, err
	}

	result.AnimationFrames = make([]d2enum.AnimationFrame, result.FramesPerDirection)

	for i := range b {
		result.AnimationFrames[i] = d2enum.AnimationFrame(b[i])
	}

	priorityLen := result.FramesPerDirection * result.NumberOfDirections * result.NumberOfLayers
	result.Priority = make([][][]d2enum.CompositeType, result.NumberOfDirections)

	priorityBytes, err := streamReader.ReadBytes(priorityLen)
	if err != nil {
		return nil, err
	}

	priorityIndex := 0

	for direction := 0; direction < result.NumberOfDirections; direction++ {
		result.Priority[direction] = make([][]d2enum.CompositeType, result.FramesPerDirection)
		for frame := 0; frame < result.FramesPerDirection; frame++ {
			result.Priority[direction][frame] = make([]d2enum.CompositeType, result.NumberOfLayers)
			for i := 0; i < result.NumberOfLayers; i++ {
				result.Priority[direction][frame][i] = d2enum.CompositeType(priorityBytes[priorityIndex])
				priorityIndex++
			}
		}
	}

	return result, nil
}

// Marshals encodes COF back into byte slince
func (c *COF) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	sw.PushByte(byte(c.NumberOfLayers))
	sw.PushByte(byte(c.FramesPerDirection))
	sw.PushByte(byte(c.NumberOfDirections))
	sw.PushBytes(c.unknownHeaderBytes)
	sw.PushByte(byte(c.Speed))
	sw.PushBytes(c.unknown1)

	for i := range c.CofLayers {
		sw.PushByte(byte(c.CofLayers[i].Type.Int()))
		sw.PushByte(c.CofLayers[i].Shadow)

		if c.CofLayers[i].Selectable {
			sw.PushByte(byte(1))
		} else {
			sw.PushByte(byte(0))
		}

		if c.CofLayers[i].Transparent {
			sw.PushByte(byte(1))
		} else {
			sw.PushByte(byte(0))
		}

		sw.PushByte(byte(c.CofLayers[i].DrawEffect))

		sw.PushBytes(c.CofLayers[i].weaponClassByte)
	}

	for _, i := range c.AnimationFrames {
		sw.PushByte(byte(i))
	}

	for direction := 0; direction < c.NumberOfDirections; direction++ {
		for frame := 0; frame < c.FramesPerDirection; frame++ {
			for i := 0; i < c.NumberOfLayers; i++ {
				sw.PushByte(byte(c.Priority[direction][frame][i]))
			}
		}
	}

	return sw.GetBytes()
}
