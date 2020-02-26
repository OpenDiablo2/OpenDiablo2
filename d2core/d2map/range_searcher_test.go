package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockEntity struct {
	x float64
	y float64
}

func (m *mockEntity) Render(target d2render.Surface) {
	panic("implement me")
}

func (m *mockEntity) Advance(tickTime float64) {
	panic("implement me")
}

func (m *mockEntity) GetPosition() (float64, float64) {
	return m.x, m.y
}

func newMockEntity(x, y float64) MapEntity {
	return &mockEntity{
		x: x,
		y: y,
	}
}

func TestRangeSearcher_Add(t *testing.T) {
	searcher := &rangeSearcher{
		entities: make([]MapEntity, 0, 64),
	}

	searcher.Add(
		newMockEntity(0, 9),
		newMockEntity(8, 1),
		newMockEntity(1, 8),
		newMockEntity(3, 6),
		newMockEntity(5, 4),
		newMockEntity(6, 3),
		newMockEntity(9, 0),
		newMockEntity(4, 5),
		newMockEntity(2, 7),
		newMockEntity(7, 2),
	)

	for i := 0; i <= 9; i++ {
		_, pos := searcher.entities[i].GetPosition()
		assert.Equal(t, float64(i), pos)
	}

}

func TestRangeSearcher_SearchByRect(t *testing.T) {
	searcher := &rangeSearcher{
		entities: make([]MapEntity, 0, 64),
	}

	searcher.Add(
		newMockEntity(0, 9),
		newMockEntity(8, 1),
		newMockEntity(1, 8),
		newMockEntity(3, 6),
		newMockEntity(5, 4),
		newMockEntity(6, 3),
		newMockEntity(9, 0),
		newMockEntity(4, 5),
		newMockEntity(2, 7),
		newMockEntity(7, 2),
	)

	matches := searcher.SearchByRect(d2common.Rectangle{
		Left:   3,
		Top:    0,
		Width:  4,
		Height: 9,
	})

	valsX := make([]float64, 0)
	for _, match := range matches {
		x, _ := match.GetPosition()
		valsX = append(valsX, x)
	}

	assert.ElementsMatch(t, []float64{3, 4, 5, 6, 7}, valsX)

	matches = searcher.SearchByRect(d2common.Rectangle{
		Left:   0,
		Top:    1,
		Width:  9,
		Height: 4,
	})

	valsY := make([]float64, 0)
	for _, match := range matches {
		_, y := match.GetPosition()
		valsY = append(valsY, y)
	}

	assert.ElementsMatch(t, []float64{1, 2, 3, 4, 5}, valsY)

	matches = searcher.SearchByRect(d2common.Rectangle{
		Left:   3,
		Top:    3,
		Width:  2,
		Height: 2,
	})

	valsY = make([]float64, 0)
	valsX = make([]float64, 0)
	for _, match := range matches {
		x, y := match.GetPosition()
		valsX = append(valsX, x)
		valsY = append(valsY, y)
	}

	assert.ElementsMatch(t, []float64{4, 5}, valsY)
	assert.ElementsMatch(t, []float64{4, 5}, valsX)
}
