package d2render

import (
	"fmt"
	"image"
	"log"
	"math"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2data/d2cof"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dcc"

	"github.com/OpenDiablo2/D2Shared/d2helper"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"

	"github.com/OpenDiablo2/D2Shared/d2common"

	"github.com/OpenDiablo2/D2Shared/d2data"

	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"

	"github.com/hajimehoshi/ebiten"
)

var DccLayerNames = []string{"HD", "TR", "LG", "RA", "LA", "RH", "LH", "SH", "S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8"}

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
	frames             map[string][]*ebiten.Image
	frameLocations     map[string][]d2common.Rectangle
	object             *d2datadict.ObjectLookupRecord
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(x, y int32, object *d2datadict.ObjectLookupRecord, fileProvider d2interface.FileProvider, palette d2enum.PaletteType) AnimatedEntity {
	result := AnimatedEntity{
		fileProvider: fileProvider,
		base:         object.Base,
		token:        object.Token,
		object:       object,
		palette:      palette,
	}
	result.dccLayers = make(map[string]d2dcc.DCC)
	result.LocationX = float64(x) / 5
	result.LocationY = float64(y) / 5

	result.subcellX = 1 + math.Mod(float64(x), 5)
	result.subcellY = 1 + math.Mod(float64(y), 5)

	return result
}

// DirectionLookup is used to decode the direction offset indexes
var DirectionLookup = []int{3, 15, 4, 8, 0, 9, 5, 10, 1, 11, 6, 12, 2, 13, 7, 14}

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedEntity) SetMode(animationMode, weaponClass string, direction int) {
	cofPath := fmt.Sprintf("%s/%s/COF/%s%s%s.COF", v.base, v.token, v.token, animationMode, weaponClass)
	v.Cof = d2cof.LoadCOF(cofPath, v.fileProvider)
	v.animationMode = animationMode
	v.weaponClass = weaponClass
	v.direction = direction
	if v.direction >= v.Cof.NumberOfDirections {
		v.direction = v.Cof.NumberOfDirections - 1
	}
	v.frames = make(map[string][]*ebiten.Image)
	v.frameLocations = make(map[string][]d2common.Rectangle)
	v.dccLayers = make(map[string]d2dcc.DCC)
	for _, cofLayer := range v.Cof.CofLayers {
		layerName := DccLayerNames[cofLayer.Type]
		v.dccLayers[layerName] = v.LoadLayer(layerName, v.fileProvider)
		if !v.dccLayers[layerName].IsValid() {
			continue
		}
		v.cacheFrames(layerName)
	}

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
	for idx := 0; idx < v.Cof.NumberOfLayers; idx++ {
		priority := v.Cof.Priority[v.direction][v.currentFrame][idx]
		if int(priority) >= len(DccLayerNames) {
			continue
		}
		frameName := DccLayerNames[priority]
		if v.frames[frameName] == nil {
			continue
		}

		// Location within the current tile
		localX := (v.subcellX - v.subcellY) * 16
		localY := ((v.subcellX + v.subcellY) * 8) - 5

		// TODO: Transparency op maybe, but it'l murder batch calls
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(v.frameLocations[frameName][v.currentFrame].Left+offsetX)+localX,
			float64(v.frameLocations[frameName][v.currentFrame].Top+offsetY)+localY)
		if err := target.DrawImage(v.frames[frameName][v.currentFrame], opts); err != nil {
			log.Panic(err.Error())
		}
	}
}

func (v *AnimatedEntity) cacheFrames(layerName string) {
	dcc := v.dccLayers[layerName]
	v.currentFrame = 0
	animationData := d2data.AnimationData[strings.ToLower(v.token+v.animationMode+v.weaponClass)][0]
	v.animationSpeed = 1.0 / ((float64(animationData.AnimationSpeed) * 25.0) / 256.0)
	v.framesToAnimate = animationData.FramesPerDirection
	v.lastFrameTime = d2helper.Now()
	minX := int32(10000)
	minY := int32(10000)
	maxX := int32(-10000)
	maxY := int32(-10000)
	for _, layer := range dcc.Directions {
		minX = d2helper.MinInt32(minX, int32(layer.Box.Left))
		minY = d2helper.MinInt32(minY, int32(layer.Box.Top))
		maxX = d2helper.MaxInt32(maxX, int32(layer.Box.Right()))
		maxY = d2helper.MaxInt32(maxY, int32(layer.Box.Bottom()))
	}
	frameW := maxX - minX
	frameH := maxY - minY
	v.frames[layerName] = make([]*ebiten.Image, v.framesToAnimate)
	v.frameLocations[layerName] = make([]d2common.Rectangle, v.framesToAnimate)
	for frameIndex := range v.frames[layerName] {
		v.frames[layerName][frameIndex], _ = ebiten.NewImage(int(frameW), int(frameH), ebiten.FilterNearest)
		for layerIdx := 0; layerIdx < v.Cof.NumberOfLayers; layerIdx++ {
			transparency := byte(255)
			if v.Cof.CofLayers[layerIdx].Transparent {
				transparency = byte(128)
			}

			direction := dcc.Directions[v.direction]
			if frameIndex >= len(direction.Frames) {
				continue
			}
			frame := direction.Frames[frameIndex]
			img := image.NewRGBA(image.Rect(0, 0, int(frameW), int(frameH)))
			for y := 0; y < direction.Box.Height; y++ {
				for x := 0; x < direction.Box.Width; x++ {
					paletteIndex := frame.PixelData[x+(y*direction.Box.Width)]

					if paletteIndex == 0 {
						continue
					}
					color := d2datadict.Palettes[v.palette].Colors[paletteIndex]
					actualX := x + direction.Box.Left - int(minX)
					actualY := y + direction.Box.Top - int(minY)
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)] = color.R
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)+1] = color.G
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)+2] = color.B
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)+3] = transparency
				}
			}
			newImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
			img = nil
			v.frames[layerName][frameIndex] = newImage
			v.frameLocations[layerName][frameIndex] = d2common.Rectangle{
				Left:   int(minX),
				Top:    int(minY),
				Width:  int(frameW),
				Height: int(frameH),
			}
		}
	}
}
