package d2render

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2data"

	"github.com/OpenDiablo2/D2Shared/d2data/d2cof"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dcc"

	"github.com/OpenDiablo2/D2Shared/d2helper"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"

	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"

	"github.com/hajimehoshi/ebiten"
)

var DccLayerNames = []string{"HD", "TR", "LG", "RA", "LA", "RH", "LH", "SH", "S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8"}

// DirectionLookup is used to decode the direction offset indexes
var DirectionLookup = []int{9, 15, 5, 6, 4, 12, 10, 2, 8, 13, 1, 7, 0, 14, 11, 3}

type LayerCacheEntry struct {
	frames           []*ebiten.Image
	compositeMode    ebiten.CompositeMode
	offsetX, offsetY int32
}

// AnimatedEntity represents an entity on the map that can be animated
type AnimatedEntity struct {
	fileProvider d2interface.FileProvider
	// LocationX represents the tile X position of the entity
	LocationX float64
	// LocationY represents the tile Y position of the entity
	subcellX, subcellY float64 // Subcell coordinates within the current tile
	LocationY          float64
	dccLayers          map[string]d2dcc.DCC
	Cof                *d2cof.COF
	palette            d2enum.PaletteType
	base               string
	token              string
	animationMode      string
	weaponClass        string
	lastFrameTime      float64
	framesToAnimate    int
	animationSpeed     float64
	direction          int
	currentFrame       int
	offsetX, offsetY   int32
	//frameLocations     []d2common.Rectangle
	object     *d2datadict.ObjectLookupRecord
	layerCache []LayerCacheEntry
	drawOrder  [][]d2enum.CompositeType
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int32, object *d2datadict.ObjectLookupRecord, fileProvider d2interface.FileProvider, palette d2enum.PaletteType) AnimatedEntity {
	result := AnimatedEntity{
		fileProvider: fileProvider,
		base:         object.Base,
		token:        object.Token,
		object:       object,
		palette:      palette,
		layerCache:   make([]LayerCacheEntry, d2enum.CompositeTypeMax),
		//frameLocations: []d2common.Rectangle{},
	}
	result.dccLayers = make(map[string]d2dcc.DCC)
	result.LocationX = float64(x) / 5
	result.LocationY = float64(y) / 5

	result.subcellX = 1 + math.Mod(float64(x), 5)
	result.subcellY = 1 + math.Mod(float64(y), 5)

	return result
}

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedEntity) SetMode(animationMode, weaponClass string, direction int) {
	cofPath := fmt.Sprintf("%s/%s/COF/%s%s%s.COF", v.base, v.token, v.token, animationMode, weaponClass)
	v.Cof = d2cof.LoadCOF(cofPath, v.fileProvider)
	if v.Cof.NumberOfDirections == 0 || v.Cof.NumberOfLayers == 0 || v.Cof.FramesPerDirection == 0 {
		return
	}
	v.animationMode = animationMode
	v.weaponClass = weaponClass
	v.direction = direction
	if v.direction >= v.Cof.NumberOfDirections {
		v.direction = v.Cof.NumberOfDirections - 1
	}
	v.dccLayers = make(map[string]d2dcc.DCC)
	for _, cofLayer := range v.Cof.CofLayers {
		layerName := DccLayerNames[cofLayer.Type]
		v.dccLayers[layerName] = v.LoadLayer(layerName, v.fileProvider)
		if !v.dccLayers[layerName].IsValid() {
			continue
		}
	}

	v.updateFrameCache()
}

func (v *AnimatedEntity) LoadLayer(layer string, fileProvider d2interface.FileProvider) d2dcc.DCC {
	layerName := "TR"
	switch strings.ToUpper(layer) {
	case "HD": // Head
		layerName = v.object.HD
	case "TR": // Torso
		layerName = v.object.TR
	case "LG": // Legs
		layerName = v.object.LG
	case "RA": // RightArm
		layerName = v.object.RA
	case "LA": // LeftArm
		layerName = v.object.LA
	case "RH": // RightHand
		layerName = v.object.RH
	case "LH": // LeftHand
		layerName = v.object.LH
	case "SH": // Shield
		layerName = v.object.SH
	case "S1": // Special1
		layerName = v.object.S1
	case "S2": // Special2
		layerName = v.object.S2
	case "S3": // Special3
		layerName = v.object.S3
	case "S4": // Special4
		layerName = v.object.S4
	case "S5": // Special5
		layerName = v.object.S5
	case "S6": // Special6
		layerName = v.object.S6
	case "S7": // Special7
		layerName = v.object.S7
	case "S8": // Special8
		layerName = v.object.S8
	}
	if len(layerName) == 0 {
		return d2dcc.DCC{}
	}
	dccPath := fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dcc", v.base, v.token, layer, v.token, layer, layerName, v.animationMode, v.weaponClass)
	result := d2dcc.LoadDCC(dccPath, fileProvider)
	if !result.IsValid() {
		dccPath = fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dcc", v.base, v.token, layer, v.token, layer, layerName, v.animationMode, "HTH")
		result = d2dcc.LoadDCC(dccPath, fileProvider)
	}
	return result
}

// Render draws this animated entity onto the target
func (v *AnimatedEntity) Render(target *ebiten.Image, offsetX, offsetY int) {
	if v.animationSpeed > 0 {
		now := d2helper.Now()
		framesToAdd := math.Floor((now - v.lastFrameTime) / v.animationSpeed)
		if framesToAdd > 0 {
			v.lastFrameTime += v.animationSpeed * framesToAdd
			v.currentFrame += int(math.Floor(framesToAdd))
			for v.currentFrame >= v.framesToAnimate {
				v.currentFrame -= v.framesToAnimate
			}
		}
	}
	localX := (v.subcellX - v.subcellY) * 16
	localY := ((v.subcellX + v.subcellY) * 8) - 5

	if v.drawOrder == nil {
		return
	}
	for _, layerIdx := range v.drawOrder[v.currentFrame] {
		if v.currentFrame < 0 || v.layerCache[layerIdx].frames == nil || v.currentFrame >= len(v.layerCache[layerIdx].frames) || v.layerCache[layerIdx].frames[v.currentFrame] == nil {
			continue
		}
		opts := &ebiten.DrawImageOptions{}
		x := float64(v.offsetX) + float64(offsetX) + localX + float64(v.layerCache[layerIdx].offsetX)
		y := float64(v.offsetY) + float64(offsetY) + localY + float64(v.layerCache[layerIdx].offsetY)
		opts.GeoM.Translate(x, y)
		opts.CompositeMode = v.layerCache[layerIdx].compositeMode
		if err := target.DrawImage(v.layerCache[layerIdx].frames[v.currentFrame], opts); err != nil {
			log.Panic(err.Error())
		}
	}
}

func (v *AnimatedEntity) updateFrameCache() {
	v.currentFrame = 0
	// TODO: This animation data madness is incorrect, yet tasty
	animDataTemp := d2data.AnimationData[strings.ToLower(v.token+v.animationMode+v.weaponClass)]
	if animDataTemp == nil {
		return
	}
	animationData := animDataTemp[0]
	v.animationSpeed = 1.0 / ((float64(animationData.AnimationSpeed) * 25.0) / 256.0)
	v.framesToAnimate = animationData.FramesPerDirection
	v.lastFrameTime = d2helper.Now()

	v.drawOrder = make([][]d2enum.CompositeType, v.framesToAnimate)
	// TODO/HACK: Get draw order from direction 0 until we get a handle on directions.
	for frame := 0; frame < v.framesToAnimate; frame++ {
		v.drawOrder[frame] = v.Cof.Priority[0][frame]
	}

	for cofLayerIdx := range v.Cof.CofLayers {
		layerType := v.Cof.CofLayers[cofLayerIdx].Type
		layerName := DccLayerNames[layerType]
		dccLayer := v.dccLayers[layerName]
		if !dccLayer.IsValid() {
			continue
		}
		v.layerCache[layerType].frames = make([]*ebiten.Image, v.framesToAnimate)

		minX := int32(10000)
		minY := int32(10000)
		maxX := int32(-10000)
		maxY := int32(-10000)
		for frameIdx := range dccLayer.Directions[v.direction].Frames {
			minX = d2helper.MinInt32(minX, int32(dccLayer.Directions[v.direction].Frames[frameIdx].Box.Left))
			minY = d2helper.MinInt32(minY, int32(dccLayer.Directions[v.direction].Frames[frameIdx].Box.Top))
			maxX = d2helper.MaxInt32(maxX, int32(dccLayer.Directions[v.direction].Frames[frameIdx].Box.Right()))
			maxY = d2helper.MaxInt32(maxY, int32(dccLayer.Directions[v.direction].Frames[frameIdx].Box.Bottom()))
		}

		v.layerCache[layerType].offsetX = minX
		v.layerCache[layerType].offsetY = minY
		actualWidth := maxX - minX
		actualHeight := maxY - minY

		if (actualWidth <= 0) || (actualHeight < 0) {
			log.Printf("Animated entity created with an invalid size of (%d, %d)", actualWidth, actualHeight)
			return
		}

		transparency := byte(255)
		if v.Cof.CofLayers[cofLayerIdx].Transparent {
			switch v.Cof.CofLayers[cofLayerIdx].DrawEffect {
			//Lets pick whatever we have that's closest.
			case d2enum.DrawEffectPctTransparency25:
				transparency = byte(64)
			case d2enum.DrawEffectPctTransparency50:
				transparency = byte(128)
			case d2enum.DrawEffectPctTransparency75:
				transparency = byte(192)
			case d2enum.DrawEffectModulate:
				v.layerCache[layerType].compositeMode = ebiten.CompositeModeLighter
			case d2enum.DrawEffectBurn:
				// Flies in tal rasha's tomb use this
			case d2enum.DrawEffectNormal:
			}
		}

		pixels := make([]byte, actualWidth*actualHeight*4)

		for animationIdx := 0; animationIdx < v.framesToAnimate; animationIdx++ {
			for i := 0; i < int(actualWidth*actualHeight); i++ {
				pixels[(i*4)+3] = 0
			}
			if animationIdx >= len(dccLayer.Directions[v.direction].Frames) {
				log.Printf("Invalid animation index of %d for animated entity", animationIdx)
				continue
			}

			frame := dccLayer.Directions[v.direction].Frames[animationIdx]
			for y := 0; y < dccLayer.Directions[v.direction].Box.Height; y++ {
				for x := 0; x < dccLayer.Directions[v.direction].Box.Width; x++ {
					paletteIndex := frame.PixelData[x+(y*dccLayer.Directions[v.direction].Box.Width)]
					if paletteIndex == 0 {
						continue
					}
					color := d2datadict.Palettes[v.palette].Colors[paletteIndex]
					actualX := (x + dccLayer.Directions[v.direction].Box.Left) - int(minX)
					actualY := (y + dccLayer.Directions[v.direction].Box.Top) - int(minY)
					pixels[(actualX*4)+(actualY*int(actualWidth)*4)] = color.R
					pixels[(actualX*4)+(actualY*int(actualWidth)*4)+1] = color.G
					pixels[(actualX*4)+(actualY*int(actualWidth)*4)+2] = color.B
					pixels[(actualX*4)+(actualY*int(actualWidth)*4)+3] = transparency
				}
			}
			v.layerCache[layerType].frames[animationIdx], _ = ebiten.NewImage(int(actualWidth), int(actualHeight), ebiten.FilterNearest)
			_ = v.layerCache[layerType].frames[animationIdx].ReplacePixels(pixels)
		}
	}
}

func (v AnimatedEntity) GetDirection() int {
	return v.direction
}
