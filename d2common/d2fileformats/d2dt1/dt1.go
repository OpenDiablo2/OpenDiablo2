package d2dt1

import (
	"fmt"
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

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

// DT1 represents a DT1 file.
type DT1 struct {
	majorVersion       int32
	minorVersion       int32
	unknownHeaderBytes []byte
	numberOfTiles      int32
	bodyPosition       int32
	Tiles              []Tile
}

// LoadDT1 loads a DT1 record
//nolint:funlen,gocognit,gocyclo // Can't reduce
func LoadDT1(fileData []byte) (*DT1, error) {
	result := &DT1{}
	br := d2datautils.CreateStreamReader(fileData)

	var err error

	result.majorVersion, err = br.ReadInt32()
	if err != nil {
		return nil, err
	}

	result.minorVersion, err = br.ReadInt32()
	if err != nil {
		return nil, err
	}

	if result.majorVersion != knownMajorVersion || result.minorVersion != knownMinorVersion {
		const fmtErr = "expected to have a version of 7.6, but got %d.%d instead"
		return nil, fmt.Errorf(fmtErr, result.majorVersion, result.minorVersion)
	}

	result.unknownHeaderBytes, err = br.ReadBytes(numUnknownHeaderBytes)
	if err != nil {
		return nil, err
	}

	result.numberOfTiles, err = br.ReadInt32()
	if err != nil {
		return nil, err
	}

	result.bodyPosition, err = br.ReadInt32()
	if err != nil {
		return nil, err
	}

	br.SetPosition(uint64(result.bodyPosition))

	result.Tiles = make([]Tile, result.numberOfTiles)

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

		tile.unknown1, err = br.ReadBytes(numUnknownTileBytes1)
		if err != nil {
			return nil, err
		}

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

		tile.unknown2, err = br.ReadBytes(numUnknownTileBytes2)
		if err != nil {
			return nil, err
		}

		for i := range tile.SubTileFlags {
			var subtileFlagBytes byte

			subtileFlagBytes, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			tile.SubTileFlags[i] = NewSubTileFlags(subtileFlagBytes)
		}

		tile.unknown3, err = br.ReadBytes(numUnknownTileBytes3)
		if err != nil {
			return nil, err
		}

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

		tile.unknown4, err = br.ReadBytes(numUnknownTileBytes4)
		if err != nil {
			return nil, err
		}

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

			//nolint:gomnd // Unknown data
			result.Tiles[tileIdx].Blocks[blockIdx].unknown1, err = br.ReadBytes(2)
			if err != nil {
				return nil, err
			}

			result.Tiles[tileIdx].Blocks[blockIdx].GridX, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			result.Tiles[tileIdx].Blocks[blockIdx].GridY, err = br.ReadByte()
			if err != nil {
				return nil, err
			}

			result.Tiles[tileIdx].Blocks[blockIdx].format, err = br.ReadInt16()
			if err != nil {
				return nil, err
			}

			result.Tiles[tileIdx].Blocks[blockIdx].Length, err = br.ReadInt32()
			if err != nil {
				return nil, err
			}

			//nolint:gomnd // Unknown data
			result.Tiles[tileIdx].Blocks[blockIdx].unknown2, err = br.ReadBytes(2)
			if err != nil {
				return nil, err
			}

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

// Marshal encodes dt1 data back to byte slice
func (d *DT1) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	// header
	sw.PushInt32(d.majorVersion)
	sw.PushInt32(d.minorVersion)
	sw.PushBytes(d.unknownHeaderBytes...)
	sw.PushInt32(d.numberOfTiles)
	sw.PushInt32(d.bodyPosition)

	// Step 1 - encoding tiles headers
	for i := 0; i < len(d.Tiles); i++ {
		sw.PushInt32(d.Tiles[i].Direction)
		sw.PushInt16(d.Tiles[i].RoofHeight)
		sw.PushUint16(d.Tiles[i].MaterialFlags.Encode())
		sw.PushInt32(d.Tiles[i].Height)
		sw.PushInt32(d.Tiles[i].Width)
		sw.PushBytes(d.Tiles[i].unknown1...)
		sw.PushInt32(d.Tiles[i].Type)
		sw.PushInt32(d.Tiles[i].Style)
		sw.PushInt32(d.Tiles[i].Sequence)
		sw.PushInt32(d.Tiles[i].RarityFrameIndex)
		sw.PushBytes(d.Tiles[i].unknown2...)

		for _, j := range d.Tiles[i].SubTileFlags {
			sw.PushBytes(j.Encode())
		}

		sw.PushBytes(d.Tiles[i].unknown3...)
		sw.PushInt32(d.Tiles[i].blockHeaderPointer)
		sw.PushInt32(d.Tiles[i].blockHeaderSize)
		sw.PushInt32(int32(len(d.Tiles[i].Blocks)))
		sw.PushBytes(d.Tiles[i].unknown4...)
	}

	// we must sort blocks first
	blocks := make(map[int][]Block)
	for i := range d.Tiles {
		blocks[int(d.Tiles[i].blockHeaderPointer)] = d.Tiles[i].Blocks
	}

	keys := make([]int, 0, len(blocks))
	for i := range blocks {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	// Step 2 - encoding blocks
	for i := 0; i < len(keys); i++ {
		// Step 2.1 - encoding blocks' header
		for j := range blocks[keys[i]] {
			sw.PushInt16(blocks[keys[i]][j].X)
			sw.PushInt16(blocks[keys[i]][j].Y)
			sw.PushBytes(blocks[keys[i]][j].unknown1...)
			sw.PushBytes(blocks[keys[i]][j].GridX)
			sw.PushBytes(blocks[keys[i]][j].GridY)
			sw.PushInt16(blocks[keys[i]][j].format)
			sw.PushInt32(blocks[keys[i]][j].Length)
			sw.PushBytes(blocks[keys[i]][j].unknown2...)
			sw.PushInt32(blocks[keys[i]][j].FileOffset)
		}

		// Step 2.2 - encoding blocks' data
		for j := range blocks[keys[i]] {
			sw.PushBytes(blocks[keys[i]][j].EncodedData...)
		}
	}

	return sw.GetBytes()
}
