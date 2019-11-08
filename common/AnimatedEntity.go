package common

import (
	"fmt"
	"image"
	"math"
	"strings"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/palettedefs"

	"github.com/hajimehoshi/ebiten"
)

// AnimatedEntity represents an entity on the map that can be animated
type AnimatedEntity struct {
	// LocationX represents the tile X position of the entity
	LocationX float64
	// LocationY represents the tile Y position of the entity
	LocationY       float64
	dcc             *DCC
	cof             *Cof
	palette         palettedefs.PaletteType
	base            string
	token           string
	tr              string
	animationMode   string
	weaponClass     string
	lastFrameTime   time.Time
	framesToAnimate int
	animationSpeed  int
	direction       int
	currentFrame    int
	frames          []*ebiten.Image
	frameLocations  []Rectangle
}

// CreateAnimatedEntity creates an instance of AnimatedEntity
func CreateAnimatedEntity(object Object, fileProvider FileProvider, palette palettedefs.PaletteType) *AnimatedEntity {
	result := &AnimatedEntity{
		base:    object.Lookup.Base,
		token:   object.Lookup.Token,
		tr:      object.Lookup.TR,
		palette: palette,
	}
	result.LocationX = math.Floor(float64(object.X) / 5)
	result.LocationY = math.Floor(float64(object.Y) / 5)
	return result
}

// DirectionLookup is used to decode the direction offset indexes
var DirectionLookup = []int{3, 15, 4, 8, 0, 9, 5, 10, 1, 11, 6, 12, 2, 13, 7, 14}

// SetMode changes the graphical mode of this animated entity
func (v *AnimatedEntity) SetMode(animationMode, weaponClass string, direction int, provider FileProvider) {
	dccPath := fmt.Sprintf("%s/%s/tr/%str%s%s%s.dcc", v.base, v.token, v.token, v.tr, animationMode, weaponClass)
	v.dcc = LoadDCC(dccPath, provider)
	cofPath := fmt.Sprintf("%s/%s/cof/%s%s%s.cof", v.base, v.token, v.token, animationMode, weaponClass)
	v.cof = LoadCof(cofPath, provider)
	v.animationMode = animationMode
	v.weaponClass = weaponClass
	v.direction = direction
	v.cacheFrames()
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

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(v.frameLocations[v.currentFrame].Left+offsetX), float64(v.frameLocations[v.currentFrame].Top+offsetY+40))
	target.DrawImage(v.frames[v.currentFrame], opts)
}

func (v *AnimatedEntity) cacheFrames() {
	v.currentFrame = 0
	animationData := AnimationData[strings.ToLower(v.token+v.animationMode+v.weaponClass)][0]
	v.animationSpeed = int(1000.0 / ((float64(animationData.AnimationSpeed) * 25.0) / 256.0))
	v.framesToAnimate = animationData.FramesPerDirection
	v.lastFrameTime = time.Now()
	minX := int32(2147483647)
	minY := int32(2147483647)
	maxX := int32(-2147483648)
	maxY := int32(-2147483648)
	for _, layer := range v.dcc.Directions {
		minX = MinInt32(minX, int32(layer.Box.Left))
		minY = MinInt32(minY, int32(layer.Box.Top))
		maxX = MaxInt32(maxX, int32(layer.Box.Right()))
		maxY = MaxInt32(maxY, int32(layer.Box.Bottom()))
	}
	frameW := maxX - minX
	frameH := maxY - minY
	v.frames = make([]*ebiten.Image, v.framesToAnimate)
	v.frameLocations = make([]Rectangle, v.framesToAnimate)
	for frameIndex := range v.frames {
		v.frames[frameIndex], _ = ebiten.NewImage(int(frameW), int(frameH), ebiten.FilterNearest)
		priorityBase := (v.direction * animationData.FramesPerDirection * v.cof.NumberOfLayers) + (frameIndex * v.cof.NumberOfLayers)
		for layerIdx := 0; layerIdx < v.cof.NumberOfLayers; layerIdx++ {
			comp := v.cof.Priority[priorityBase+layerIdx]
			if _, found := v.cof.CompositeLayers[comp]; !found {
				continue
			}
			direction := v.dcc.Directions[v.direction]
			frame := direction.Frames[frameIndex]
			img := image.NewRGBA(image.Rect(0, 0, int(frameW), int(frameH)))
			for y := 0; y < direction.Box.Height; y++ {
				for x := 0; x < direction.Box.Width; x++ {
					paletteIndex := frame.PixelData[x+(y*direction.Box.Width)]

					if paletteIndex == 0 {
						continue
					}
					color := Palettes[v.palette].Colors[paletteIndex]
					actualX := x + direction.Box.Left - int(minX)
					actualY := y + direction.Box.Top - int(minY)
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)] = color.R
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)+1] = color.G
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)+2] = color.B
					img.Pix[(actualX*4)+(actualY*int(frameW)*4)+3] = 255
				}
			}
			newImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
			img = nil
			v.frames[frameIndex] = newImage
			v.frameLocations[frameIndex] = Rectangle{
				Left:   int(minX),
				Top:    int(minY),
				Width:  int(frameW),
				Height: int(frameH),
			}
		}
	}
}
