package d2render

import (
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2data"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"

	"github.com/hajimehoshi/ebiten"
)

var DccLayerNames = []string{"HD", "TR", "LG", "RA", "LA", "RH", "LH", "SH", "S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8"}

// AnimatedEntity represents an entity on the map that can be animated
type AnimatedEntity struct {
	// LocationX represents the tile X position of the entity
	LocationX float64
	// LocationY represents the tile Y position of the entity
	LocationY       float64
	dccLayers       map[string]*d2data.DCC
	Cof             *d2data.Cof
	palette         d2enum.PaletteType
	base            string
	token           string
	animationMode   string
	weaponClass     string
	lastFrameTime   time.Time
	framesToAnimate int
	animationSpeed  int
	direction       int
	currentFrame    int
	frames          map[string][]*ebiten.Image
	frameLocations  map[string][]d2common.Rectangle
	object          d2data.Object
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(object d2data.Object, fileProvider d2interface.FileProvider, palette d2enum.PaletteType) *AnimatedEntity {
	result := &AnimatedEntity{
		base:    object.Lookup.Base,
		token:   object.Lookup.Token,
		object:  object,
		palette: palette,
	}
	result.dccLayers = make(map[string]*d2data.DCC)
	result.LocationX = float64(object.X) / 5
	result.LocationY = float64(object.Y) / 5
	return result
}

// DirectionLookup is used to decode the direction offset indexes
var DirectionLookup = []int{3, 15, 4, 8, 0, 9, 5, 10, 1, 11, 6, 12, 2, 13, 7, 14}

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedEntity) SetMode(animationMode, weaponClass string, direction int, provider d2interface.FileProvider) {
	cofPath := fmt.Sprintf("%s/%s/Cof/%s%s%s.Cof", v.base, v.token, v.token, animationMode, weaponClass)
	v.Cof = d2data.LoadCof(cofPath, provider)
	v.animationMode = animationMode
	v.weaponClass = weaponClass
	v.direction = direction
	if v.direction >= v.Cof.NumberOfDirections {
		v.direction = v.Cof.NumberOfDirections - 1
	}
	v.frames = make(map[string][]*ebiten.Image)
	v.frameLocations = make(map[string][]d2common.Rectangle)
	v.dccLayers = make(map[string]*d2data.DCC)
	for _, cofLayer := range v.Cof.CofLayers {
		layerName := DccLayerNames[cofLayer.Type]
		v.dccLayers[layerName] = v.LoadLayer(layerName, provider)
		if v.dccLayers[layerName] == nil {
			continue
		}
		v.cacheFrames(layerName)
	}

}

func (v *AnimatedEntity) LoadLayer(layer string, fileProvider d2interface.FileProvider) *d2data.DCC {
	layerName := "tr"
	switch strings.ToUpper(layer) {
	case "HD": // Head
		layerName = v.object.Lookup.HD
	case "TR": // Torso
		layerName = v.object.Lookup.TR
	case "LG": // Legs
		layerName = v.object.Lookup.LG
	case "RA": // RightArm
		layerName = v.object.Lookup.RA
	case "LA": // LeftArm
		layerName = v.object.Lookup.LA
	case "RH": // RightHand
		layerName = v.object.Lookup.RH
	case "LH": // LeftHand
		layerName = v.object.Lookup.LH
	case "SH": // Shield
		layerName = v.object.Lookup.SH
	case "S1": // Special1
		layerName = v.object.Lookup.S1
	case "S2": // Special2
		layerName = v.object.Lookup.S2
	case "S3": // Special3
		layerName = v.object.Lookup.S3
	case "S4": // Special4
		layerName = v.object.Lookup.S4
	case "S5": // Special5
		layerName = v.object.Lookup.S5
	case "S6": // Special6
		layerName = v.object.Lookup.S6
	case "S7": // Special7
		layerName = v.object.Lookup.S7
	case "S8": // Special8
		layerName = v.object.Lookup.S8
	}
	if len(layerName) == 0 {
		return nil
	}
	dccPath := fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dcc", v.base, v.token, layer, v.token, layer, layerName, v.animationMode, v.weaponClass)
	return d2data.LoadDCC(dccPath, fileProvider)
}

// Render draws this animated entity onto the target
func (v *AnimatedEntity) Render(target *ebiten.Image, offsetX, offsetY int) {
	for v.lastFrameTime.Add(time.Millisecond * time.Duration(v.animationSpeed)).Before(time.Now()) {
		v.lastFrameTime = v.lastFrameTime.Add(time.Millisecond * time.Duration(v.animationSpeed))
		v.currentFrame++
		if v.currentFrame >= v.framesToAnimate {
			v.currentFrame = 0
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
		// TODO: Transparency op maybe, but it'l murder batch calls
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(v.frameLocations[frameName][v.currentFrame].Left+offsetX),
			float64(v.frameLocations[frameName][v.currentFrame].Top+offsetY+40))
		target.DrawImage(v.frames[frameName][v.currentFrame], opts)
	}
}

func (v *AnimatedEntity) cacheFrames(layerName string) {
	dcc := v.dccLayers[layerName]
	v.currentFrame = 0
	animationData := d2data.AnimationData[strings.ToLower(v.token+v.animationMode+v.weaponClass)][0]
	v.animationSpeed = int(1000.0 / ((float64(animationData.AnimationSpeed) * 25.0) / 256.0))
	v.framesToAnimate = animationData.FramesPerDirection
	v.lastFrameTime = time.Now()
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
