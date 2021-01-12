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

const (
	numUnknownHeaderBytes = 260
	knownMajorVersion     = 7
	knownMinorVersion     = 6
	numUnknownTileBytes1  = 4
	numUnknownTileBytes2  = 4
	numUnknownTileBytes3  = 7
	numUnknownTileBytes4  = 12
)

// LoadDT1 loads a DT1 record
//nolint:funlen,gocognit,gocyclo // Can't reduce
func LoadDT1(fileData []byte) (*DT1, error) {
	result := &DT1{}
	br := d2datautils.CreateStreamReader(fileData)

	var err error

	majorVersion, err := br.ReadInt32()
	if err != nil {
		return nil, err
	}

	minorVersion, err := br.ReadInt32()
	if err != nil {
		return nil, err
	}

	if majorVersion != knownMajorVersion || minorVersion != knownMinorVersion {
		const fmtErr = "expected to have a version of 7.6, but got %d.%d instead"
		return nil, fmt.Errorf(fmtErr, majorVersion, minorVersion)
	}

	br.SkipBytes(numUnknownHeaderBytes)

	numberOfTiles, err := br.ReadInt32()
	if err != nil {
		return nil, err
	}

	position, err := br.ReadInt32()
	if err != nil {
		return nil, err
	}

	br.SetPosition(uint64(position))

	result.Tiles = make([]Tile, numberOfTiles)

	for tileIdx := range result.Tiles {
		tile := Tile{}

		tile.Direction, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.RoofHeight, err = br.ReadInt16()
		if err != nil {
			return nil, err
		}

		var matFlagBytes uint16

		matFlagBytes, err = br.ReadUInt16()
		if err != nil {
			return nil, err
		}

		tile.MaterialFlags = NewMaterialFlags(matFlagBytes)

		tile.Height, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.Width, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		br.SkipBytes(numUnknownTileBytes1)

		tile.Type, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.Style, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.Sequence, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.RarityFrameIndex, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		br.SkipBytes(numUnknownTileBytes2)

		for i := range tile.SubTileFlags {
			var subtileFlagBytes byte

			subtileFlagBytes, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			tile.SubTileFlags[i] = NewSubTileFlags(subtileFlagBytes)
		}

		br.SkipBytes(numUnknownTileBytes3)

		tile.blockHeaderPointer, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.blockHeaderSize, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		var numBlocks int32

		numBlocks, err = br.ReadInt32()
		if err != nil {
			return nil, err
		}

		tile.Blocks = make([]Block, numBlocks)

		br.SkipBytes(numUnknownTileBytes4)

		result.Tiles[tileIdx] = tile
	}

	for tileIdx := range result.Tiles {
		tile := &result.Tiles[tileIdx]
		br.SetPosition(uint64(tile.blockHeaderPointer))

		for blockIdx := range tile.Blocks {
			result.Tiles[tileIdx].Blocks[blockIdx].X, err = br.ReadInt16()
			if err != nil {
				return nil, err
			}

			result.Tiles[tileIdx].Blocks[blockIdx].Y, err = br.ReadInt16()
			if err != nil {
				return nil, err
			}

			br.SkipBytes(2) //nolint:gomnd // Unknown data

			result.Tiles[tileIdx].Blocks[blockIdx].GridX, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			result.Tiles[tileIdx].Blocks[blockIdx].GridY, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			formatValue, err := br.ReadInt16()
			if err != nil {
				return nil, err
			}

			if formatValue == 1 {
				result.Tiles[tileIdx].Blocks[blockIdx].Format = BlockFormatIsometric
			} else {
				result.Tiles[tileIdx].Blocks[blockIdx].Format = BlockFormatRLE
			}

			result.Tiles[tileIdx].Blocks[blockIdx].Length, err = br.ReadInt32()
			if err != nil {
				return nil, err
			}

			br.SkipBytes(2) //nolint:gomnd // Unknown data

			result.Tiles[tileIdx].Blocks[blockIdx].FileOffset, err = br.ReadInt32()
			if err != nil {
				return nil, err
			}
		}

		for blockIndex, block := range tile.Blocks {
			br.SetPosition(uint64(tile.blockHeaderPointer + block.FileOffset))

			encodedData, err := br.ReadBytes(int(block.Length))
			if err != nil {
				return nil, err
			}

			tile.Blocks[blockIndex].EncodedData = encodedData
		}
	}

	return result, nil
}
