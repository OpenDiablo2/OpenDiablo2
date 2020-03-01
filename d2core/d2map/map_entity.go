package d2map

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/beefsack/go-astar"
)

// mapEntity represents an entity on the map that can be animated
type mapEntity struct {
	LocationX          float64
	LocationY          float64
	TileX, TileY       int     // Coordinates of the tile the unit is within
	subcellX, subcellY float64 // Subcell coordinates within the current tile
	weaponClass        string
	offsetX, offsetY   int
	TargetX            float64
	TargetY            float64
	Speed              float64
	path               []astar.Pather
	done               func()
	directioner        func(angle float64)
}

// createMapEntity creates an instance of mapEntity
func createMapEntity(x, y int) mapEntity {
	locX, locY := float64(x), float64(y)
	return mapEntity{
		LocationX: locX,
		LocationY: locY,
		TargetX:   locX,
		TargetY:   locY,
		TileX:     x / 5,
		TileY:     y / 5,
		subcellX:  1 + math.Mod(locX, 5),
		subcellY:  1 + math.Mod(locY, 5),
		Speed:     6,
		path:      []astar.Pather{},
	}
}

func (m *mapEntity) SetPath(path []astar.Pather, done func()) {
	m.path = path
	m.done = done
}

func (m *mapEntity) getStepLength(tickTime float64) (float64, float64) {
	length := tickTime * m.Speed

	angle := 359 - d2common.GetAngleBetween(
		m.LocationX,
		m.LocationY,
		m.TargetX,
		m.TargetY,
	)
	radians := (math.Pi / 180.0) * float64(angle)
	oneStepX := length * math.Cos(radians)
	oneStepY := length * math.Sin(radians)
	return oneStepX, oneStepY
}

func (m *mapEntity) IsAtTarget() bool {
	return m.LocationX == m.TargetX && m.LocationY == m.TargetY && !m.HasPathFinding()
}

func (m *mapEntity) Step(tickTime float64) {
	if m.IsAtTarget() {
		if m.done != nil {
			m.done()
			m.done = nil
		}
		return
	}

	stepX, stepY := m.getStepLength(tickTime)

	if d2common.AlmostEqual(m.LocationX, m.TargetX, stepX) {
		m.LocationX = m.TargetX
	}
	if d2common.AlmostEqual(m.LocationY, m.TargetY, stepY) {
		m.LocationY = m.TargetY
	}
	if m.LocationX != m.TargetX {
		m.LocationX += stepX
	}
	if m.LocationY != m.TargetY {
		m.LocationY += stepY
	}

	m.subcellX = 1 + math.Mod(m.LocationX, 5)
	m.subcellY = 1 + math.Mod(m.LocationY, 5)
	m.TileX = int(m.LocationX / 5)
	m.TileY = int(m.LocationY / 5)

	if (m.LocationX != m.TargetX) || (m.LocationY != m.TargetY) {
		return
	}

	if len(m.path) > 0 {
		m.SetTarget(m.path[0].(*PathTile).X*5, m.path[0].(*PathTile).Y*5, m.done)

		if len(m.path) > 1 {
			m.path = m.path[1:]
		} else {
			m.path = []astar.Pather{}
		}
		return
	}

}

func (m *mapEntity) HasPathFinding() bool {
	return len(m.path) > 0
}

// SetTarget sets target coordinates and changes animation based on proximity and direction
func (m *mapEntity) SetTarget(tx, ty float64, done func()) {
	m.TargetX, m.TargetY = tx, ty
	m.done = done

	if m.directioner != nil {
		angle := 359 - d2common.GetAngleBetween(
			m.LocationX,
			m.LocationY,
			tx,
			ty,
		)
		m.directioner(float64(angle))
	}
}

func angleToDirection(angle float64) int {
	degreesPerDirection := 360.0 / 64.0
	offset := 45.0 - (degreesPerDirection / 2)

	newDirection := int((angle - offset) / degreesPerDirection)

	if newDirection >= 64 {
		newDirection = newDirection - 64
	} else if newDirection < 0 {
		newDirection = 64 + newDirection
	}

	return newDirection
}

func (m *mapEntity) GetPosition() (float64, float64) {
	return float64(m.TileX), float64(m.TileY)
}
