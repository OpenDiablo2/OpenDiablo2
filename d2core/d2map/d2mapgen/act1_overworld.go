package d2mapgen

// magic number suppression has been added because most of the work done here
// is experiemental, and mapgen will likely change dramatically in the future.

import (
	"math/rand"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen/d2wilderness"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

const (
	presetA = iota
	presetB
	presetC
	presetD
	presetE
	presetF
)

const (
	wildernessDetailsRecordID = 2
)

const (
	mapWidth      = 150
	mapHeight     = mapWidth
	mapMargin     = 9
	autoFileIndex = -1
)

// GenerateAct1Overworld generates the map and entities for the first town and surrounding area.
func (g *MapGenerator) GenerateAct1Overworld() {
	rand.Seed(g.engine.Seed())

	wilderness1Details := g.asset.Records.GetLevelDetails(wildernessDetailsRecordID)

	g.engine.ResetMap(d2enum.RegionAct1Town, mapWidth, mapHeight)
	mapWidth := g.engine.Size().Width
	mapHeight := g.engine.Size().Height

	townStamp := g.engine.LoadStamp(d2enum.RegionAct1Town, presetB, autoFileIndex)
	townStamp.RegionPath()
	townSize := townStamp.Size()

	g.Infof("Region Path: %s", townStamp.RegionPath())

	switch {
	case strings.Contains(townStamp.RegionPath(), "E1"):
		g.engine.PlaceStamp(townStamp, 0, 0)
		g.generateWilderness1TownEast(townSize.Width, 0)
	case strings.Contains(townStamp.RegionPath(), "S1"):
		g.engine.PlaceStamp(townStamp, mapWidth-townSize.Width, 0)
		rightWaterBorderStamp := g.loadPreset(d2wilderness.WaterBorderEast, 0)
		rightWaterBorderStamp2 := g.loadPreset(d2wilderness.WaterBorderWest, 0)

		mapInnerHeight := mapHeight - mapMargin
		mapInnerWidth := mapWidth - mapMargin

		for y := townSize.Height; y < mapInnerHeight; y += mapMargin {
			g.engine.PlaceStamp(rightWaterBorderStamp, mapInnerWidth-mapMargin, y)
			g.engine.PlaceStamp(rightWaterBorderStamp2, mapInnerWidth, y)
		}

		edgeOffset := 14
		startX := mapWidth - wilderness1Details.SizeXNormal - edgeOffset
		startY := townSize.Height

		g.generateWilderness1TownSouth(startX, startY)
	case strings.Contains(townStamp.RegionPath(), "W1"):
		g.engine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)
		startX := mapWidth - townSize.Width - wilderness1Details.SizeXNormal
		startY := mapHeight - wilderness1Details.SizeYNormal
		g.generateWilderness1TownWest(startX, startY)
	default:
		g.engine.PlaceStamp(townStamp, mapWidth-townSize.Width, mapHeight-townSize.Height)
	}
}

// nolint:gosec,gomnd // we dont need crypto-strong randomness, mapgen will get a refactor soon
func (g *MapGenerator) generateWilderness1TownEast(startX, startY int) {
	levelDetails := g.asset.Records.GetLevelDetails(wildernessDetailsRecordID)

	fenceNorthStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderNorth, presetA),
		g.loadPreset(d2wilderness.TreeBorderNorth, presetB),
		g.loadPreset(d2wilderness.TreeBorderNorth, presetC),
	}

	fenceSouthStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderSouth, presetA),
		g.loadPreset(d2wilderness.TreeBorderSouth, presetB),
		g.loadPreset(d2wilderness.TreeBorderSouth, presetC),
	}

	fenceWestStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderWest, presetA),
		g.loadPreset(d2wilderness.TreeBorderWest, presetB),
		g.loadPreset(d2wilderness.TreeBorderWest, presetC),
	}

	fenceEastStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderEast, presetA),
		g.loadPreset(d2wilderness.TreeBorderEast, presetB),
		g.loadPreset(d2wilderness.TreeBorderEast, presetC),
	}

	fenceSouthWestStamp := g.loadPreset(d2wilderness.TreeBorderSouthWest, presetA)
	fenceNorthEastStamp := g.loadPreset(d2wilderness.TreeBorderNorthEast, presetA)
	fenceSouthEastStamp := g.loadPreset(d2wilderness.TreeBorderSouthEast, presetA)
	fenceWestEdge := g.loadPreset(d2wilderness.TreeBoxNorthEast, presetA)

	areaRect := d2geom.Rectangle{
		Left:   startX,
		Top:    startY + 9,
		Width:  levelDetails.SizeXNormal,
		Height: levelDetails.SizeYNormal - 3,
	}

	g.generateWilderness1Contents(areaRect)

	// Draw the north and south fence
	for i := 0; i < 9; i++ {
		g.engine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9), startY)
		g.engine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9),
			startY+(levelDetails.SizeYNormal+6))
	}

	// West fence
	for i := 1; i < 6; i++ {
		g.engine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX,
			startY+(levelDetails.SizeYNormal+6)-(i*9))
	}

	// East Fence
	for i := 1; i < 10; i++ {
		g.engine.PlaceStamp(fenceEastStamp[rand.Intn(3)], startX+levelDetails.SizeXNormal, startY+(i*9))
	}

	g.engine.PlaceStamp(fenceSouthWestStamp, startX, startY+levelDetails.SizeYNormal+6)
	g.engine.PlaceStamp(fenceWestEdge, startX, startY+(levelDetails.SizeYNormal-3)-45)
	g.engine.PlaceStamp(fenceNorthEastStamp, startX+levelDetails.SizeXNormal, startY)
	g.engine.PlaceStamp(fenceSouthEastStamp, startX+levelDetails.SizeXNormal, startY+levelDetails.SizeYNormal+6)
}

// nolint:gosec,gomnd // we dont need crypto-strong randomness, mapgen will get a refactor soon
func (g *MapGenerator) generateWilderness1TownSouth(startX, startY int) {
	levelDetails := g.asset.Records.GetLevelDetails(wildernessDetailsRecordID)

	fenceNorthStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderNorth, presetA),
		g.loadPreset(d2wilderness.TreeBorderNorth, presetB),
		g.loadPreset(d2wilderness.TreeBorderNorth, presetC),
	}

	fenceWestStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderWest, presetA),
		g.loadPreset(d2wilderness.TreeBorderWest, presetB),
		g.loadPreset(d2wilderness.TreeBorderWest, presetC),
	}

	fenceSouthStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderSouth, presetA),
		g.loadPreset(d2wilderness.TreeBorderSouth, presetB),
		g.loadPreset(d2wilderness.TreeBorderSouth, presetC),
	}

	fenceNorthWestStamp := g.loadPreset(d2wilderness.TreeBorderNorthWest, presetA)
	fenceSouthWestStamp := g.loadPreset(d2wilderness.TreeBorderSouthWest, presetB)
	fenceWaterBorderSouthEast := g.loadPreset(d2wilderness.WaterBorderEast, presetC)

	areaRect := d2geom.Rectangle{
		Left:   startX + 2,
		Top:    startY,
		Width:  levelDetails.SizeXNormal - 2,
		Height: levelDetails.SizeYNormal - 3,
	}
	g.generateWilderness1Contents(areaRect)

	// Draw the north fence
	for i := 0; i < 4; i++ {
		g.engine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9)+5, startY-6)
	}

	// Draw the west fence
	for i := 0; i < 8; i++ {
		g.engine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX, startY+(i*9)+3)
	}

	// Draw the south fence
	for i := 1; i < 9; i++ {
		g.engine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9), startY+(8*9)+3)
	}

	g.engine.PlaceStamp(fenceNorthWestStamp, startX, startY-6)
	g.engine.PlaceStamp(fenceSouthWestStamp, startX, startY+(8*9)+3)
	g.engine.PlaceStamp(fenceWaterBorderSouthEast, startX+(9*9)-4, startY+(8*9)+1)
}

// nolint:gosec,gomnd // we dont need crypto-strong randomness, mapgen will get a refactor soon
func (g *MapGenerator) generateWilderness1TownWest(startX, startY int) {
	levelDetails := g.asset.Records.GetLevelDetails(wildernessDetailsRecordID)

	fenceEastEdge := g.loadPreset(d2wilderness.TreeBoxSouthWest, presetA)
	fenceNorthWestStamp := g.loadPreset(d2wilderness.TreeBorderNorthWest, presetA)
	fenceNorthEastStamp := g.loadPreset(d2wilderness.TreeBorderNorthEast, presetA)
	fenceSouthWestStamp := g.loadPreset(d2wilderness.TreeBorderSouthWest, presetA)

	fenceSouthStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderSouth, presetA),
		g.loadPreset(d2wilderness.TreeBorderSouth, presetB),
		g.loadPreset(d2wilderness.TreeBorderSouth, presetC),
	}

	fenceNorthStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderNorth, presetA),
		g.loadPreset(d2wilderness.TreeBorderNorth, presetB),
		g.loadPreset(d2wilderness.TreeBorderNorth, presetC),
	}

	fenceEastStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderEast, presetA),
		g.loadPreset(d2wilderness.TreeBorderEast, presetB),
		g.loadPreset(d2wilderness.TreeBorderEast, presetC),
	}

	fenceWestStamp := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.TreeBorderWest, presetA),
		g.loadPreset(d2wilderness.TreeBorderWest, presetB),
		g.loadPreset(d2wilderness.TreeBorderWest, presetC),
	}

	// Draw the north and south fences
	for i := 0; i < 9; i++ {
		if i > 0 && i < 8 {
			g.engine.PlaceStamp(fenceNorthStamp[rand.Intn(3)], startX+(i*9)-1, startY-15)
		}

		g.engine.PlaceStamp(fenceSouthStamp[rand.Intn(3)], startX+(i*9)-1, startY+levelDetails.SizeYNormal-12)
	}

	// Draw the east fence
	for i := 0; i < 6; i++ {
		g.engine.PlaceStamp(fenceEastStamp[rand.Intn(3)], startX+levelDetails.SizeXNormal-9, startY+(i*9)-6)
	}

	// Draw the west fence
	for i := 0; i < 9; i++ {
		g.engine.PlaceStamp(fenceWestStamp[rand.Intn(3)], startX, startY+(i*9)-6)
	}

	// Draw the west fence
	g.engine.PlaceStamp(fenceEastEdge, startX+levelDetails.SizeXNormal-9, startY+39)
	g.engine.PlaceStamp(fenceNorthWestStamp, startX, startY-15)
	g.engine.PlaceStamp(fenceSouthWestStamp, startX, startY+levelDetails.SizeYNormal-12)
	g.engine.PlaceStamp(fenceNorthEastStamp, startX+levelDetails.SizeXNormal-9, startY-15)

	areaRect := d2geom.Rectangle{
		Left:   startX + 9,
		Top:    startY - 10,
		Width:  levelDetails.SizeXNormal - 9,
		Height: levelDetails.SizeYNormal - 2,
	}
	g.generateWilderness1Contents(areaRect)
}

// nolint:gosec,gomnd // we dont need crypto-strong randomness, mapgen will get a refactor soon
func (g *MapGenerator) generateWilderness1Contents(rect d2geom.Rectangle) {
	levelDetails := g.asset.Records.GetLevelDetails(wildernessDetailsRecordID)

	denOfEvil := g.loadPreset(d2wilderness.DenOfEvilEntrance, 0)
	denOfEvilLoc := d2geom.Point{
		X: rect.Left + (rect.Width / 2) + rand.Intn(10),
		Y: rect.Top + (rect.Height / 2) + rand.Intn(10),
	}

	// Fill in the grass
	for y := 0; y < rect.Height; y++ {
		for x := 0; x < rect.Width; x++ {
			tile := g.engine.Tile(rect.Left+x, rect.Top+y)
			tile.RegionType = d2enum.RegionIdType(levelDetails.LevelType)
			tile.Components.Floors = []d2ds1.FloorShadowRecord{{Prop1: 1, Style: 0, Sequence: 0}} // wildernessGrass
			tile.PrepareTile(x, y, g.engine)
		}
	}

	stuff := []*d2mapstamp.Stamp{
		g.loadPreset(d2wilderness.StoneFill1, presetA),
		g.loadPreset(d2wilderness.StoneFill1, presetB),
		g.loadPreset(d2wilderness.StoneFill1, presetC),
		g.loadPreset(d2wilderness.StoneFill2, presetA),
		g.loadPreset(d2wilderness.StoneFill2, presetB),
		g.loadPreset(d2wilderness.StoneFill2, presetC),
		g.loadPreset(d2wilderness.Cottages1, presetA),
		g.loadPreset(d2wilderness.Cottages1, presetB),
		g.loadPreset(d2wilderness.Cottages1, presetC),
		g.loadPreset(d2wilderness.Cottages1, presetD),
		g.loadPreset(d2wilderness.Cottages1, presetE),
		g.loadPreset(d2wilderness.Cottages1, presetF),
		g.loadPreset(d2wilderness.FallenCamp1, presetA),
		g.loadPreset(d2wilderness.FallenCamp1, presetB),
		g.loadPreset(d2wilderness.FallenCamp1, presetC),
		g.loadPreset(d2wilderness.FallenCamp1, presetD),
		g.loadPreset(d2wilderness.Pond, presetA),
		g.loadPreset(d2wilderness.SwampFill1, presetA),
		g.loadPreset(d2wilderness.SwampFill2, presetA),
	}

	g.engine.PlaceStamp(denOfEvil, denOfEvilLoc.X, denOfEvilLoc.Y)

	numPlaced := 0
	for numPlaced < 25 {
		stamp := stuff[rand.Intn(len(stuff))]

		stampRect := d2geom.Rectangle{
			Left:   rect.Left + rand.Intn(rect.Width) - stamp.Size().Width,
			Top:    rect.Top + rand.Intn(rect.Height) - stamp.Size().Height,
			Width:  stamp.Size().Width,
			Height: stamp.Size().Height,
		}

		if areaEmpty(g.engine, stampRect) {
			g.engine.PlaceStamp(stamp, stampRect.Left, stampRect.Top)
			numPlaced++
		}
	}
}
