package d2maprenderer

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var imageCacheRecords map[uint32]d2interface.Surface

// Invalidates the global region image cache. Call this when you are changing regions
func InvalidateImageCache() {
	imageCacheRecords = nil
}

func (mr *MapRenderer) getImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte) d2interface.Surface {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	return imageCacheRecords[lookupIndex]
}

func (mr *MapRenderer) setImageCacheRecord(style, sequence byte, tileType d2enum.TileType, randomIndex byte, image d2interface.Surface) {
	lookupIndex := uint32(style)<<24 | uint32(sequence)<<16 | uint32(tileType)<<8 | uint32(randomIndex)
	if imageCacheRecords == nil {
		imageCacheRecords = make(map[uint32]d2interface.Surface)
	}
	imageCacheRecords[lookupIndex] = image
}
