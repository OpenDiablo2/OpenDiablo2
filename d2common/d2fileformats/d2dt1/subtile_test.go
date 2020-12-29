package d2dt1

import (
	"testing"

	testify "github.com/stretchr/testify/assert"
)

func TestNewSubTile(t *testing.T) {
	assert := testify.New(t)
	data := []byte{1, 2, 4, 8, 16, 32, 64, 128}

	for i, b := range data {
		tile := NewSubTileFlags(b)
		assert.Equal(i == 0, tile.BlockWalk)
		assert.Equal(i == 1, tile.BlockLOS)
		assert.Equal(i == 2, tile.BlockJump)
		assert.Equal(i == 3, tile.BlockPlayerWalk)
		assert.Equal(i == 4, tile.Unknown1)
		assert.Equal(i == 5, tile.BlockLight)
		assert.Equal(i == 6, tile.Unknown2)
		assert.Equal(i == 7, tile.Unknown3)
	}
}
