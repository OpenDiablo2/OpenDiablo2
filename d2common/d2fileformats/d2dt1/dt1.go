package d2dt1

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// https://d2mods.info/forum/viewtopic.php?t=65163

type DT1 struct {
	Tiles []Tile
}

type BlockDataFormat int16

const (
	BlockFormatRLE       BlockDataFormat = 0 // Not 1
	BlockFormatIsometric BlockDataFormat = 1
)

func LoadDT1(fileData []byte) DT1 {
	result := DT1{}
	br := d2common.CreateStreamReader(fileData)
	ver1 := br.GetInt32()
	ver2 := br.GetInt32()
	if ver1 != 7 || ver2 != 6 {
		log.Panicf("Expected to have a version of 7.6, but got %d.%d instead", ver1, ver2)
	}
	br.SkipBytes(260)
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
		br.SkipBytes(4)
		newTile.Type = br.GetInt32()
		newTile.Style = br.GetInt32()
		newTile.Sequence = br.GetInt32()
		newTile.RarityFrameIndex = br.GetInt32()
		br.SkipBytes(4)
		for i := range newTile.SubTileFlags {
			newTile.SubTileFlags[i] = NewSubTileFlags(br.GetByte())
		}
		br.SkipBytes(7)
		newTile.blockHeaderPointer = br.GetInt32()
		newTile.blockHeaderSize = br.GetInt32()
		newTile.Blocks = make([]Block, br.GetInt32())
		br.SkipBytes(12)
		result.Tiles[tileIdx] = newTile
	}
	for tileIdx, tile := range result.Tiles {
		br.SetPosition(uint64(tile.blockHeaderPointer))
		for blockIdx := range tile.Blocks {
			result.Tiles[tileIdx].Blocks[blockIdx].X = br.GetInt16()
			result.Tiles[tileIdx].Blocks[blockIdx].Y = br.GetInt16()
			br.SkipBytes(2)
			result.Tiles[tileIdx].Blocks[blockIdx].GridX = br.GetByte()
			result.Tiles[tileIdx].Blocks[blockIdx].GridY = br.GetByte()
			formatValue := br.GetInt16()
			if formatValue == 1 {
				result.Tiles[tileIdx].Blocks[blockIdx].Format = BlockFormatIsometric
			} else {
				result.Tiles[tileIdx].Blocks[blockIdx].Format = BlockFormatRLE
			}
			result.Tiles[tileIdx].Blocks[blockIdx].Length = br.GetInt32()
			br.SkipBytes(2)
			result.Tiles[tileIdx].Blocks[blockIdx].FileOffset = br.GetInt32()
		}
		for blockIndex, block := range tile.Blocks {
			br.SetPosition(uint64(tile.blockHeaderPointer + block.FileOffset))
			encodedData, _ := br.ReadBytes(int(block.Length))
			tile.Blocks[blockIndex].EncodedData = encodedData
		}

	}
	return result
}
