package d2asset

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	d2iface "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type dccAnimationFrame struct {
	width   int
	height  int
	offsetX int
	offsetY int

	image d2iface.Surface
}

type dccDirection struct {
	decoded bool
	frames  []*dccAnimationFrame
}

// DCCAnimation represens an animation decoded from DCC
type DCCAnimation struct {
	dccPath          string
	transparency     int
	palette          d2iface.Palette
	renderer         d2iface.Renderer
	directions       []dccDirection
	colorMod         color.Color
	frameIndex       int
	directionIndex   int
	lastFrameTime    float64
	playedCount      int
	compositeMode    d2enum.CompositeMode
	playMode         playMode
	playLength       float64
	subStartingFrame int
	subEndingFrame   int
	playLoop         bool
	hasSubLoop       bool // runs after first animation ends
}

// CreateAnimationFromDCC creates an animation from d2dcc.DCC and d2dat.DATPalette
func CreateDCCAnimation(renderer d2iface.Renderer, dccPath string, palette d2iface.Palette,
	transparency int) (d2iface.Animation, error) {
	anim := DCCAnimation{
		playLength:   defaultPlayLength,
		dccPath:      dccPath,
		palette:      palette,
		renderer:     renderer,
		transparency: transparency,
		playLoop:     true,
	}

	dcc, err := loadDCC(dccPath)
	if err != nil {
		return nil, err
	}

	anim.directions = make([]dccDirection, dcc.NumberOfDirections)

	err = anim.SetDirection(0)
	if err != nil {
		return nil, err
	}

	return &anim, nil
}

// Clone creates a copy of the animation
func (a *DCCAnimation) Clone() d2iface.Animation {
	animation := *a
	return &animation
}

// SetSubLoop sets a sub loop for the animation
func (a *DCCAnimation) SetSubLoop(startFrame, endFrame int) {
	a.subStartingFrame = startFrame
	a.subEndingFrame = endFrame
	a.hasSubLoop = true
}

// Advance advances the animation state
func (a *DCCAnimation) Advance(elapsed float64) error {
	if a.playMode == playModePause {
		return nil
	}

	frameCount := a.GetFrameCount()
	frameLength := a.playLength / float64(frameCount)
	a.lastFrameTime += elapsed
	framesAdvanced := int(a.lastFrameTime / frameLength)
	a.lastFrameTime -= float64(framesAdvanced) * frameLength

	for i := 0; i < framesAdvanced; i++ {
		startIndex := 0
		endIndex := frameCount

		if a.hasSubLoop && a.playedCount > 0 {
			startIndex = a.subStartingFrame
			endIndex = a.subEndingFrame
		}

		switch a.playMode {
		case playModeForward:
			a.frameIndex++
			if a.frameIndex >= endIndex {
				a.playedCount++
				if a.playLoop {
					a.frameIndex = startIndex
				} else {
					a.frameIndex = endIndex - 1
					break
				}
			}
		case playModeBackward:
			a.frameIndex--
			if a.frameIndex < startIndex {
				a.playedCount++
				if a.playLoop {
					a.frameIndex = endIndex - 1
				} else {
					a.frameIndex = startIndex
					break
				}
			}
		}
	}

	return nil
}

// Render renders the animation to the given surface
func (a *DCCAnimation) Render(target d2iface.Surface) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushCompositeMode(a.compositeMode)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	return target.Render(frame.image)
}

// RenderFromOrigin renders the animation from the animation origin
func (a *DCCAnimation) RenderFromOrigin(target d2iface.Surface) error {
	return a.Render(target)
}

// RenderSection renders the section of the animation frame enclosed by bounds
func (a *DCCAnimation) RenderSection(sfc d2iface.Surface, bound image.Rectangle) error {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	sfc.PushTranslation(frame.offsetX, frame.offsetY)
	sfc.PushCompositeMode(a.compositeMode)
	sfc.PushColor(a.colorMod)

	defer sfc.PopN(3)

	return sfc.RenderSection(frame.image, bound)
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (a *DCCAnimation) GetFrameSize(frameIndex int) (width, height int, err error) {
	direction := a.directions[a.directionIndex]
	if frameIndex >= len(direction.frames) {
		return 0, 0, errors.New("invalid frame index")
	}

	frame := direction.frames[frameIndex]

	return frame.width, frame.height, nil
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (a *DCCAnimation) GetCurrentFrameSize() (width, height int) {
	width, height, _ = a.GetFrameSize(a.frameIndex)
	return width, height
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (a *DCCAnimation) GetFrameBounds() (maxWidth, maxHeight int) {
	maxWidth, maxHeight = 0, 0

	direction := a.directions[a.directionIndex]
	for _, frame := range direction.frames {
		maxWidth = d2common.MaxInt(maxWidth, frame.width)
		maxHeight = d2common.MaxInt(maxHeight, frame.height)
	}

	return maxWidth, maxHeight
}

// GetCurrentFrame gets index of current frame in animation
func (a *DCCAnimation) GetCurrentFrame() int {
	return a.frameIndex
}

// GetFrameCount gets number of frames in animation
func (a *DCCAnimation) GetFrameCount() int {
	direction := a.directions[a.directionIndex]
	return len(direction.frames)
}

// IsOnFirstFrame gets if the animation on its first frame
func (a *DCCAnimation) IsOnFirstFrame() bool {
	return a.frameIndex == 0
}

// IsOnLastFrame gets if the animation on its last frame
func (a *DCCAnimation) IsOnLastFrame() bool {
	return a.frameIndex == a.GetFrameCount()-1
}

// GetDirectionCount gets the number of animation direction
func (a *DCCAnimation) GetDirectionCount() int {
	return len(a.directions)
}

// SetDirection places the animation in the direction of an animation
func (a *DCCAnimation) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	direction := d2dcc.Dir64ToDcc(directionIndex, len(a.directions))
	if a.directions[direction].decoded == false {
		err := a.decodeDirection(direction)
		if err != nil {
			return err
		}
	}

	a.directionIndex = direction
	a.frameIndex = 0

	return nil
}

// GetDirection get the current animation direction
func (a *DCCAnimation) GetDirection() int {
	return a.directionIndex
}

// SetCurrentFrame sets animation at a specific frame
func (a *DCCAnimation) SetCurrentFrame(frameIndex int) error {
	if frameIndex >= a.GetFrameCount() {
		return errors.New("invalid frame index")
	}

	a.frameIndex = frameIndex
	a.lastFrameTime = 0

	return nil
}

// Rewind animation to beginning
func (a *DCCAnimation) Rewind() {
	_ = a.SetCurrentFrame(0)
}

// PlayForward plays animation forward
func (a *DCCAnimation) PlayForward() {
	a.playMode = playModeForward
	a.lastFrameTime = 0
}

// PlayBackward plays animation backward
func (a *DCCAnimation) PlayBackward() {
	a.playMode = playModeBackward
	a.lastFrameTime = 0
}

// Pause animation
func (a *DCCAnimation) Pause() {
	a.playMode = playModePause
	a.lastFrameTime = 0
}

// SetPlayLoop sets whether to loop the animation
func (a *DCCAnimation) SetPlayLoop(loop bool) {
	a.playLoop = loop
}

// SetPlaySpeed sets play speed of the animation
func (a *DCCAnimation) SetPlaySpeed(playSpeed float64) {
	a.SetPlayLength(playSpeed * float64(a.GetFrameCount()))
}

// SetPlayLength sets the Animation's play length in seconds
func (a *DCCAnimation) SetPlayLength(playLength float64) {
	// TODO refactor to use time.Duration instead of float64
	a.playLength = playLength
	a.lastFrameTime = 0
}

// SetPlayLengthMs sets the Animation's play length in milliseconds
func (a *DCCAnimation) SetPlayLengthMs(playLengthMs int) {
	// TODO remove this method
	const millisecondsPerSecond = 1000.0
	a.SetPlayLength(float64(playLengthMs) / millisecondsPerSecond)
}

// SetColorMod sets the Animation's color mod
func (a *DCCAnimation) SetColorMod(colorMod color.Color) {
	a.colorMod = colorMod
}

// GetPlayedCount gets the number of times the application played
func (a *DCCAnimation) GetPlayedCount() int {
	return a.playedCount
}

// ResetPlayedCount resets the play count
func (a *DCCAnimation) ResetPlayedCount() {
	a.playedCount = 0
}

// SetBlend sets the Animation alpha blending status
func (a *DCCAnimation) SetBlend(blend bool) {
	if blend {
		a.compositeMode = d2enum.CompositeModeLighter
	} else {
		a.compositeMode = d2enum.CompositeModeSourceOver
	}
}

func (a *DCCAnimation) decodeDirection(directionIndex int) error {
	dcc, err := loadDCC(a.dccPath)
	if err != nil {
		return err
	}

	direction := dcc.DecodeDirection(directionIndex)

	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for _, dccFrame := range direction.Frames {
		minX = d2common.MinInt(minX, dccFrame.Box.Left)
		minY = d2common.MinInt(minY, dccFrame.Box.Top)
		maxX = d2common.MaxInt(maxX, dccFrame.Box.Right())
		maxY = d2common.MaxInt(maxY, dccFrame.Box.Bottom())
	}

	for _, dccFrame := range direction.Frames {
		frameWidth := maxX - minX
		frameHeight := maxY - minY

		const bytesPerPixel = 4
		pixels := make([]byte, frameWidth*frameHeight*bytesPerPixel)

		for y := 0; y < frameHeight; y++ {
			for x := 0; x < frameWidth; x++ {
				paletteIndex := dccFrame.PixelData[y*frameWidth+x]

				if paletteIndex == 0 {
					continue
				}

				palColor := a.palette.GetColors()[paletteIndex]
				offset := (x + y*frameWidth) * bytesPerPixel
				pixels[offset] = palColor.R()
				pixels[offset+1] = palColor.G()
				pixels[offset+2] = palColor.B()
				pixels[offset+3] = byte(a.transparency)
			}
		}

		sfc, err := a.renderer.NewSurface(frameWidth, frameHeight, d2iface.FilterNearest)
		if err != nil {
			return err
		}

		if err := sfc.ReplacePixels(pixels); err != nil {
			return err
		}

		a.directions[directionIndex].decoded = true
		a.directions[directionIndex].frames = append(a.directions[directionIndex].frames, &dccAnimationFrame{
			width:   dccFrame.Width,
			height:  dccFrame.Height,
			offsetX: minX,
			offsetY: minY,
			image:   sfc,
		})
	}

	return nil
}
