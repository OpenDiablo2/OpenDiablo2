package d2map

import (
	"math"
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type MapEntitiesSearcher interface {
	// Returns all map entities.
	All() []MapEntity
	// Add adds an entity to the index and re-sorts.
	Add(entities ...MapEntity)
	// Remove finds and removes the entity from the index.
	Remove(entity MapEntity)
	// SearchByRect get entities in a rectangle, results will be sorted top left to bottom right.
	// Elements with equal Y will be sorted by X
	SearchByRect(rect d2common.Rectangle) []MapEntity
	// SearchByRadius get entities in a circle, results will be sorted top left to bottom right.
	// Elements with equal Y will be sorted by X
	SearchByRadius(originX, originY, radius float64) []MapEntity
	// Update re-sorts the index, must be ran after each update.
	Update()
}

// rangeSearcher a basic index of entity locations using a slice ordered by Y then X coordinates.
// Eventually this should be probably replaced with a proper spatial index.
type rangeSearcher struct {
	entities []MapEntity
}

func NewRangeSearcher() MapEntitiesSearcher {
	return &rangeSearcher{
		entities: make([]MapEntity, 0, 64),
	}
}

func (r *rangeSearcher) All() []MapEntity {
	return r.entities
}

func (r *rangeSearcher) Add(entities ...MapEntity) {
	r.entities = append(r.entities, entities...)

	r.Update()
}

func (r *rangeSearcher) Remove(entity MapEntity) {
	if entity == nil {
		return
	}

	// In-place filter to remove the given entity.
	n := 0
	for _, check := range r.entities {
		if check != entity {
			r.entities[n] = check
			n++
		}
	}
	r.entities = r.entities[:n]
}

func (r *rangeSearcher) SearchByRect(rect d2common.Rectangle) []MapEntity {
	left, top, right, bottom := float64(rect.Left), float64(rect.Top), float64(rect.Right()), float64(rect.Bottom())
	topIndex := sort.Search(len(r.entities), func(i int) bool {
		x, y := r.entities[i].GetPosition()
		if y == top {
			return x >= left
		}
		return y >= top
	})

	matches := make([]MapEntity, 0, 16)

	for i := topIndex; i < len(r.entities); i++ {
		x, y := r.entities[i].GetPosition()
		if y > bottom {
			break
		}

		if x >= left && x <= right {
			matches = append(matches, r.entities[i])
		}

	}

	return matches
}

func (r *rangeSearcher) SearchByRadius(originX, originY, radius float64) []MapEntity {
	left, right := originX-radius, originX+radius
	top, bottom := originY-radius, originY+radius
	inRect := r.SearchByRect(d2common.Rectangle{
		Left:   int(left),
		Top:    int(top),
		Width:  int(right - left),
		Height: int(bottom - top),
	})

	// In-place filter to remove entities outside the radius.
	n := 0
	for _, check := range inRect {
		x, y := check.GetPosition()
		if distance(originX, originY, x, y) <= radius {
			inRect[n] = check
			n++
		}
	}
	return inRect[:n]
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Abs(math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)))
}

// Re-sorts the index after entities have moved.
// Uses bubble sort to target O(n) sort time, in most cases no entities will be swapped.
func (r *rangeSearcher) Update() {
	bubbleSort(r.entities, func(i, j int) bool {
		ix, iy := r.entities[i].GetPosition()
		jx, jy := r.entities[j].GetPosition()
		if iy == jy {
			return ix < jx
		}
		return iy < jy
	})
}

func bubbleSort(items []MapEntity, less func(i, j int) bool) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if less(i+1, i) {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
