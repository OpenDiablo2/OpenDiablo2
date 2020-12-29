package d2cof

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// COF is a structure that represents a COF file.
type COF struct {
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
	result.NumberOfLayers = int(streamReader.GetByte())
	result.FramesPerDirection = int(streamReader.GetByte())
	result.NumberOfDirections = int(streamReader.GetByte())

	streamReader.SkipBytes(21) //nolint:gomnd // Unknown data

	result.Speed = int(streamReader.GetByte())

	streamReader.SkipBytes(3) //nolint:gomnd // Unknown data

	result.CofLayers = make([]CofLayer, result.NumberOfLayers)
	result.CompositeLayers = make(map[d2enum.CompositeType]int)

	for i := 0; i < result.NumberOfLayers; i++ {
		layer := CofLayer{}
		layer.Type = d2enum.CompositeType(streamReader.GetByte())
		layer.Shadow = streamReader.GetByte()
		layer.Selectable = streamReader.GetByte() != 0
		layer.Transparent = streamReader.GetByte() != 0
		layer.DrawEffect = d2enum.DrawEffect(streamReader.GetByte())
		weaponClassStr := streamReader.ReadBytes(4) //nolint:gomnd // Binary data
		layer.WeaponClass = d2enum.WeaponClassFromString(strings.TrimSpace(strings.ReplaceAll(string(weaponClassStr), string(byte(0)), "")))
		result.CofLayers[i] = layer
		result.CompositeLayers[layer.Type] = i
	}

	animationFrameBytes := streamReader.ReadBytes(result.FramesPerDirection)
	result.AnimationFrames = make([]d2enum.AnimationFrame, result.FramesPerDirection)

	for i := range animationFrameBytes {
		result.AnimationFrames[i] = d2enum.AnimationFrame(animationFrameBytes[i])
	}

	priorityLen := result.FramesPerDirection * result.NumberOfDirections * result.NumberOfLayers
	result.Priority = make([][][]d2enum.CompositeType, result.NumberOfDirections)
	priorityBytes := streamReader.ReadBytes(priorityLen)
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
