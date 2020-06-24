package d2mapengine

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2astar"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func (m *MapEngine) RegenerateWalkPaths() {
	for subTileY := 0; subTileY < m.size.Height*5; subTileY++ {
		tileY := int(float64(subTileY) / 5.0)
		for subTileX := 0; subTileX < m.size.Width*5; subTileX++ {
			tileX := int(float64(subTileX) / 5.0)
			tile := m.TileAt(tileX, tileY)
			isBlocked := false
			for _, floor := range tile.Floors {
				tileData := m.GetTileData(int32(floor.Style), int32(floor.Sequence), d2enum.Floor)
				if tileData == nil {
					continue
				}
				tileSubAttributes := tileData.GetSubTileFlags(subTileX%5, subTileY%5)
				isBlocked = isBlocked || tileSubAttributes.BlockWalk
				if isBlocked {
					break
				}
			}
			if !isBlocked {
				for _, wall := range tile.Walls {
					tileData := m.GetTileData(int32(wall.Style), int32(wall.Sequence), wall.Type)
					if tileData == nil {
						continue
					}
					tileSubAttributes := tileData.GetSubTileFlags(subTileX%5, subTileY%5)
					isBlocked = isBlocked || tileSubAttributes.BlockWalk
					if isBlocked {
						break
					}
				}
			}

			index := subTileX + (subTileY * m.size.Width * 5)
			m.walkMesh[index] = d2common.PathTile{
				Walkable: !isBlocked,
				X:        float64(subTileX) / 5.0,
				Y:        float64(subTileY) / 5.0,
			}

			ySkew := m.size.Width * 5
			if !isBlocked && subTileY > 0 && m.walkMesh[index-ySkew].Walkable {
				m.walkMesh[index].Up = &m.walkMesh[index-ySkew]
				m.walkMesh[index-ySkew].Down = &m.walkMesh[index]
			}
			if !isBlocked && subTileX > 0 && m.walkMesh[index-1].Walkable {
				m.walkMesh[index].Left = &m.walkMesh[index-1]
				m.walkMesh[index-1].Right = &m.walkMesh[index]
			}
			if !isBlocked && subTileX > 0 && subTileY > 0 && m.walkMesh[(index-ySkew)-1].Walkable {
				m.walkMesh[index].UpLeft = &m.walkMesh[(index-ySkew)-1]
				m.walkMesh[(index-ySkew)-1].DownRight = &m.walkMesh[index]
			}
			if !isBlocked && subTileY > 0 && subTileX < (m.size.Width*5)-1 && m.walkMesh[(index-ySkew)+1].Walkable {
				m.walkMesh[index].UpRight = &m.walkMesh[(index-ySkew)+1]
				m.walkMesh[(index-ySkew)+1].DownLeft = &m.walkMesh[index]
			}
		}
	}
}

// Finds a walkable path between two points
func (m *MapEngine) PathFind(startX, startY, endX, endY float64) (path []d2astar.Pather, distance float64, found bool) {
	startTileX := int(math.Floor(startX))
	startTileY := int(math.Floor(startY))
	if !m.TileExists(startTileX, startTileY) {
		return
	}
	startSubtileX := int((startX - float64(int(startX))) * 5)
	startSubtileY := int((startY - float64(int(startY))) * 5)
	startNodeIndex := ((startSubtileY + (startTileY * 5)) * m.size.Width * 5) + startSubtileX + ((startTileX) * 5)
	if startNodeIndex < 0 || startNodeIndex >= len(m.walkMesh) {
		return
	}
	startNode := &m.walkMesh[startNodeIndex]

	endTileX := int(math.Floor(endX))
	endTileY := int(math.Floor(endY))
	if !m.TileExists(endTileX, endTileY) {
		return
	}
	endSubtileX := int((endX - float64(int(endX))) * 5)
	endSubtileY := int((endY - float64(int(endY))) * 5)

	endNodeIndex := ((endSubtileY + (endTileY * 5)) * m.size.Width * 5) + endSubtileX + ((endTileX) * 5)
	if endNodeIndex < 0 || endNodeIndex >= len(m.walkMesh) {
		return
	}
	endNode := &m.walkMesh[endNodeIndex]

	path, distance, found = d2astar.Path(startNode, endNode, 80)
	if path != nil {
		// Reverse the path to fit what the game expects.
		for i := len(path)/2-1; i >= 0; i-- {
			opp := len(path)-1-i
			path[i], path[opp] = path[opp], path[i]
		}

		path = path[1:]
	}
	return
}
