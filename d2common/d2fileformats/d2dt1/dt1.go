package d2dt1

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

// DT1 represents a DT1 file.
type DT1 struct {
	Tiles []Tile
}

// BlockDataFormat represents the format of the block data
type BlockDataFormat int16

const (
	// BlockFormatRLE specifies the block format is RLE encoded
	BlockFormatRLE BlockDataFormat = 0

	// BlockFormatIsometric specifies the block format isometrically encoded
	BlockFormatIsometric BlockDataFormat = 1
)

// LoadDT1 loads a DT1 record
//nolint:funlen // Can't reduce
func LoadDT1(fileData []byte) (*DT1, error) {
	result := &DT1{}
	br := d2datautils.CreateStreamReader(fileData)
	ver1 := br.GetInt32()
	ver2 := br.GetInt32()

	if ver1 != 7 || ver2 != 6 {
		return nil, fmt.Errorf("expected to have a version of 7.6, but got %d.%d instead", ver1, ver2)
	}

	br.SkipBytes(260) //nolint:gomnd // Unknown data

	numberOfTiles := br.GetInt32()
	br.SetPosition(uint64(br.GetInt32()))

	result.Tiles = make([]Tile, numberOfTiles)

	for tileIdx := range result.Tiles {
		newTile := Tile{}
		newTile.Direction = br.GetInt32()
		newTile.RoofHeight = br.GetInt16()
		newTile.MaterialFlags = NewMaterialFlags(br.GetUInt16())
		newTile.Height = br.GetInt32()
		newTile.Width = br.GetInt32()

		br.SkipBytes(4) //nolint:gomnd // Unknown data

		newTile.Type = br.GetInt32()
		newTile.Style = br.GetInt32()
		newTile.Sequence = br.GetInt32()
		newTile.RarityFrameIndex = br.GetInt32()

		br.SkipBytes(4) //nolint:gomnd // Unknown data

		for i := range newTile.SubTileFlags {
			newTile.SubTileFlags[i] = NewSubTileFlags(br.GetByte())
		}

		br.SkipBytes(7) //nolint:gomnd // Unknown data

		newTile.blockHeaderPointer = br.GetInt32()
		newTile.blockHeaderSize = br.GetInt32()
		newTile.Blocks = make([]Block, br.GetInt32())

		br.SkipBytes(12) //nolint:gomnd // Unknown data

		result.Tiles[tileIdx] = newTile
	}

	for tileIdx := range result.Tiles {
		tile := &result.Tiles[tileIdx]
		br.SetPosition(uint64(tile.blockHeaderPointer))

		for blockIdx := range tile.Blocks {
			result.Tiles[tileIdx].Blocks[blockIdx].X = br.GetInt16()
			result.Tiles[tileIdx].Blocks[blockIdx].Y = br.GetInt16()

			br.SkipBytes(2) //nolint:gomnd // Unknown data

			result.Tiles[tileIdx].Blocks[blockIdx].GridX = br.GetByte()
			result.Tiles[tileIdx].Blocks[blockIdx].GridY = br.GetByte()
			formatValue := br.GetInt16()

			if formatValue == 1 {
				result.Tiles[tileIdx].Blocks[blockIdx].Format = BlockFormatIsometric
			} else {
				result.Tiles[tileIdx].Blocks[blockIdx].Format = BlockFormatRLE
			}

			result.Tiles[tileIdx].Blocks[blockIdx].Length = br.GetInt32()

			br.SkipBytes(2) //nolint:gomnd // Unknown data

			result.Tiles[tileIdx].Blocks[blockIdx].FileOffset = br.GetInt32()
		}

		for blockIndex, block := range tile.Blocks {
			br.SetPosition(uint64(tile.blockHeaderPointer + block.FileOffset))
			encodedData := br.ReadBytes(int(block.Length))
			tile.Blocks[blockIndex].EncodedData = encodedData
		}
	}

	return result, nil
}
