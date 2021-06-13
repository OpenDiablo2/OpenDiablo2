package d2cof

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	numUnknownHeaderBytes = 21
	numUnknownBodyBytes   = 3
	numHeaderBytes        = 4 + numUnknownHeaderBytes
	numLayerBytes         = 9
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

// New creates a new COF
func New() *COF {
	return &COF{
		unknownHeaderBytes: make([]byte, numUnknownHeaderBytes),
		unknownBodyBytes:   make([]byte, numUnknownBodyBytes),
		NumberOfDirections: 0,
		FramesPerDirection: 0,
		NumberOfLayers:     0,
		Speed:              0,
		CofLayers:          make([]CofLayer, 0),
		CompositeLayers:    make(map[d2enum.CompositeType]int),
		AnimationFrames:    make([]d2enum.AnimationFrame, 0),
		Priority:           make([][][]d2enum.CompositeType, 0),
	}
}

// Marshal a COF to a new byte slice
func Marshal(c *COF) []byte {
	return c.Marshal()
}

// Unmarshal a byte slice to a new COF
func Unmarshal(data []byte) (*COF, error) {
	c := New()
	err := c.Unmarshal(data)

	return c, err
}

// COF is a structure that represents a COF file.
type COF struct {
	CompositeLayers    map[d2enum.CompositeType]int
	unknownHeaderBytes []byte
	AnimationFrames    []d2enum.AnimationFrame
	Priority           [][][]d2enum.CompositeType
	CofLayers          []CofLayer
	unknownBodyBytes   []byte
	NumberOfLayers     int
	Speed              int
	NumberOfDirections int
	FramesPerDirection int
}

// Unmarshal a byte slice to this COF
func (c *COF) Unmarshal(fileData []byte) error {
	var err error

	streamReader := d2datautils.CreateStreamReader(fileData)

	headerBytes, err := streamReader.ReadBytes(numHeaderBytes)
	if err != nil {
		return err
	}

	c.loadHeader(headerBytes)

	c.unknownBodyBytes, err = streamReader.ReadBytes(numUnknownBodyBytes)
	if err != nil {
		return err
	}

	c.CofLayers = make([]CofLayer, c.NumberOfLayers)
	c.CompositeLayers = make(map[d2enum.CompositeType]int)

	err = c.loadCOFLayers(streamReader)
	if err != nil {
		return err
	}

	animationFramesData, err := streamReader.ReadBytes(c.FramesPerDirection)
	if err != nil {
		return err
	}

	c.loadAnimationFrames(animationFramesData)

	priorityLen := c.FramesPerDirection * c.NumberOfDirections * c.NumberOfLayers
	c.Priority = make([][][]d2enum.CompositeType, c.NumberOfDirections)

	priorityBytes, err := streamReader.ReadBytes(priorityLen)
	if err != nil {
		return err
	}

	c.loadPriority(priorityBytes)

	return nil
}

func (c *COF) loadHeader(b []byte) {
	c.NumberOfLayers = int(b[headerNumLayers])
	c.FramesPerDirection = int(b[headerFramesPerDir])
	c.NumberOfDirections = int(b[headerNumDirs])
	c.unknownHeaderBytes = b[headerNumDirs+1 : headerSpeed]
	c.Speed = int(b[headerSpeed])
}

func (c *COF) loadCOFLayers(streamReader *d2datautils.StreamReader) error {
	for i := 0; i < c.NumberOfLayers; i++ {
		layer := CofLayer{}

		b, err := streamReader.ReadBytes(numLayerBytes)
		if err != nil {
			return err
		}

		layer.Type = d2enum.CompositeType(b[layerType])
		layer.Shadow = b[layerShadow]
		layer.Selectable = b[layerSelectable] > 0
		layer.Transparent = b[layerTransparent] > 0
		layer.DrawEffect = d2enum.DrawEffect(b[layerDrawEffect])

		layer.WeaponClass = d2enum.WeaponClassFromString(strings.TrimSpace(strings.ReplaceAll(
			string(b[layerWeaponClass:]), badCharacter, "")))

		c.CofLayers[i] = layer
		c.CompositeLayers[layer.Type] = i
	}

	return nil
}

func (c *COF) loadAnimationFrames(b []byte) {
	c.AnimationFrames = make([]d2enum.AnimationFrame, c.FramesPerDirection)

	for i := range b {
		c.AnimationFrames[i] = d2enum.AnimationFrame(b[i])
	}
}

func (c *COF) loadPriority(priorityBytes []byte) {
	priorityIndex := 0

	for direction := 0; direction < c.NumberOfDirections; direction++ {
		c.Priority[direction] = make([][]d2enum.CompositeType, c.FramesPerDirection)
		for frame := 0; frame < c.FramesPerDirection; frame++ {
			c.Priority[direction][frame] = make([]d2enum.CompositeType, c.NumberOfLayers)
			for i := 0; i < c.NumberOfLayers; i++ {
				c.Priority[direction][frame][i] = d2enum.CompositeType(priorityBytes[priorityIndex])
				priorityIndex++
			}
		}
	}
}

// Marshal this COF to a byte slice
func (c *COF) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes(byte(c.NumberOfLayers))
	sw.PushBytes(byte(c.FramesPerDirection))
	sw.PushBytes(byte(c.NumberOfDirections))
	sw.PushBytes(c.unknownHeaderBytes...)
	sw.PushBytes(byte(c.Speed))
	sw.PushBytes(c.unknownBodyBytes...)

	for i := range c.CofLayers {
		sw.PushBytes(byte(c.CofLayers[i].Type))
		sw.PushBytes(c.CofLayers[i].Shadow)

		if c.CofLayers[i].Selectable {
			sw.PushBytes(byte(1))
		} else {
			sw.PushBytes(byte(0))
		}

		if c.CofLayers[i].Transparent {
			sw.PushBytes(byte(1))
		} else {
			sw.PushBytes(byte(0))
		}

		sw.PushBytes(byte(c.CofLayers[i].DrawEffect))

		const (
			maxCodeLength = 3 // we assume item codes to look like 'hax' or 'kit'
			terminator    = 0
		)

		weaponCode := c.CofLayers[i].WeaponClass.String()

		for idx, letter := range weaponCode {
			if idx > maxCodeLength {
				break
			}

			sw.PushBytes(byte(letter))
		}

		sw.PushBytes(terminator)
	}

	for _, i := range c.AnimationFrames {
		sw.PushBytes(byte(i))
	}

	for direction := 0; direction < c.NumberOfDirections; direction++ {
		for frame := 0; frame < c.FramesPerDirection; frame++ {
			for i := 0; i < c.NumberOfLayers; i++ {
				sw.PushBytes(byte(c.Priority[direction][frame][i]))
			}
		}
	}

	return sw.GetBytes()
}
