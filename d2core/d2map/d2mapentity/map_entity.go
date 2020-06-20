package d2mapentity

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"
)

type MapEntity interface {
	Render(target d2interface.Surface)
	Advance(tickTime float64)
	GetPosition() (float64, float64)
	GetLayer() int
	GetPositionF() (float64, float64)
	Name() string
	Selectable() bool
	Highlight()
}

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
	path               []d2astar.Pather
	drawLayer          int

	done        func()
	directioner func(direction int)
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
		drawLayer: 0,
		path:      []d2astar.Pather{},
	}
}

func (m *mapEntity) GetLayer() int {
	return m.drawLayer
}

func (m *mapEntity) SetPath(path []d2astar.Pather, done func()) {
	m.path = path
	m.done = done
}

func (m *mapEntity) ClearPath() {
	m.path = nil
}

func (m *mapEntity) SetSpeed(speed float64) {
	m.Speed = speed
}

func (m *mapEntity) GetSpeed() float64 {
	return m.Speed
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
	return math.Abs(m.LocationX-m.TargetX) < 0.0001 && math.Abs(m.LocationY-m.TargetY) < 0.0001 && !m.HasPathFinding()
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
	for {
		if d2common.AlmostEqual(m.LocationX-m.TargetX, 0, 0.0001) {
			stepX = 0
		}
		if d2common.AlmostEqual(m.LocationY-m.TargetY, 0, 0.0001) {
			stepY = 0
		}
		m.LocationX, stepX = d2common.AdjustWithRemainder(m.LocationX, stepX, m.TargetX)
		m.LocationY, stepY = d2common.AdjustWithRemainder(m.LocationY, stepY, m.TargetY)

		m.subcellX = 1 + math.Mod(m.LocationX, 5)
		m.subcellY = 1 + math.Mod(m.LocationY, 5)
		m.TileX = int(m.LocationX / 5)
		m.TileY = int(m.LocationY / 5)

		if d2common.AlmostEqual(m.LocationX, m.TargetX, 0.01) && d2common.AlmostEqual(m.LocationY, m.TargetY, 0.01) {
			if len(m.path) > 0 {
				m.SetTarget(m.path[0].(*d2common.PathTile).X*5, m.path[0].(*d2common.PathTile).Y*5, m.done)

				if len(m.path) > 1 {
					m.path = m.path[1:]
				} else {
					m.path = []d2astar.Pather{}
				}
			} else {
				m.LocationX = m.TargetX
				m.LocationY = m.TargetY
				m.subcellX = 1 + math.Mod(m.LocationX, 5)
				m.subcellY = 1 + math.Mod(m.LocationY, 5)
				m.TileX = int(m.LocationX / 5)
				m.TileY = int(m.LocationY / 5)
			}
		}

		if stepX == 0 && stepY == 0 {
			break
		}

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
		m.directioner(angleToDirection(float64(angle)))
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

func (m *mapEntity) GetPositionF() (float64, float64) {
	return float64(m.TileX) + (float64(m.subcellX) / 5.0), float64(m.TileY) + (float64(m.subcellY) / 5.0)
}

func (m *mapEntity) Name() string {
	return ""
}

func (m *mapEntity) Highlight() {
}

func (m *mapEntity) Selectable() bool {
	return false
}
