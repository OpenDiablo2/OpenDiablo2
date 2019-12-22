package d2asset

import (
	"errors"
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2data"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/D2Shared/d2data/d2dcc"
	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/hajimehoshi/ebiten"
)

type paperdollCacheEntry struct {
	sheetImage    *ebiten.Image
	compositeMode ebiten.CompositeMode
	width         int
	height        int
	offsetX       int
	offsetY       int
}

type Paperdoll struct {
	object  *d2datadict.ObjectLookupRecord
	palette *d2datadict.PaletteRec

	mode *paperdollMode
}

func createPaperdoll(object *d2datadict.ObjectLookupRecord, palette *d2datadict.PaletteRec) *Paperdoll {
	return &Paperdoll{object: object, palette: palette}
}

func (p *Paperdoll) Render(target *ebiten.Image, offsetX, offsetY int) {
	if p.mode == nil {
		return
	}

	if p.mode.animationSpeed > 0 {
		frameTime := d2helper.Now()
		framesToAdd := int(math.Floor((frameTime - p.mode.lastFrameTime) / p.mode.animationSpeed))
		if framesToAdd > 0 {
			p.mode.lastFrameTime += p.mode.animationSpeed * float64(framesToAdd)
			p.mode.currentFrame = (p.mode.currentFrame + framesToAdd) % p.mode.frameCount
		}
	}

	for _, layerIndex := range p.mode.drawOrder[p.mode.currentFrame] {
		cacheEntry := p.mode.layerCache[layerIndex]

		x := float64(offsetX) + float64(p.mode.layerCache[layerIndex].offsetX)
		y := float64(offsetY) + float64(p.mode.layerCache[layerIndex].offsetY)

		sheetOffset := cacheEntry.width * p.mode.currentFrame
		sheetRect := image.Rect(sheetOffset, 0, sheetOffset+cacheEntry.width, cacheEntry.height)

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(x, y)
		opts.CompositeMode = cacheEntry.compositeMode
		target.DrawImage(cacheEntry.sheetImage.SubImage(sheetRect).(*ebiten.Image), opts)
	}
}

func (p *Paperdoll) SetMode(animationMode, weaponClass string, direction int) error {
	mode, err := p.createMode(animationMode, weaponClass, direction)
	if err != nil {
		return err
	}

	p.mode = mode
	return nil
}

type paperdollMode struct {
	animationMode string
	weaponClass   string
	direction     int

	layers     []*d2dcc.DCC
	layerCache []*paperdollCacheEntry
	drawOrder  [][]d2enum.CompositeType

	frameCount     int
	animationSpeed float64
	currentFrame   int
	lastFrameTime  float64
}

func (p *Paperdoll) createMode(animationMode, weaponClass string, direction int) (*paperdollMode, error) {
	mode := &paperdollMode{
		animationMode: animationMode,
		weaponClass:   weaponClass,
		direction:     direction,
	}

	cofPath := fmt.Sprintf(
		"%s/%s/COF/%s%s%s.COF",
		p.object.Base,
		p.object.Token,
		p.object.Token,
		mode.animationMode,
		mode.weaponClass,
	)

	cof, err := LoadCOF(cofPath)
	if err != nil {
		return nil, err
	}

	if mode.direction >= cof.NumberOfDirections {
		return nil, errors.New("invalid direction")
	}

	mode.layers = make([]*d2dcc.DCC, d2enum.CompositeTypeMax)
	for _, cofLayer := range cof.CofLayers {
		var layerKey, layerValue string
		switch cofLayer.Type {
		case d2enum.CompositeTypeHead:
			layerKey = "HD"
			layerValue = p.object.HD
		case d2enum.CompositeTypeTorso:
			layerKey = "TR"
			layerValue = p.object.TR
		case d2enum.CompositeTypeLegs:
			layerKey = "LG"
			layerValue = p.object.LG
		case d2enum.CompositeTypeRightArm:
			layerKey = "RA"
			layerValue = p.object.RA
		case d2enum.CompositeTypeLeftArm:
			layerKey = "LA"
			layerValue = p.object.LA
		case d2enum.CompositeTypeRightHand:
			layerKey = "RH"
			layerValue = p.object.RH
		case d2enum.CompositeTypeLeftHand:
			layerKey = "LH"
			layerValue = p.object.LH
		case d2enum.CompositeTypeShield:
			layerKey = "SH"
			layerValue = p.object.SH
		case d2enum.CompositeTypeSpecial1:
			layerKey = "S1"
			layerValue = p.object.S1
		case d2enum.CompositeTypeSpecial2:
			layerKey = "S2"
			layerValue = p.object.S2
		case d2enum.CompositeTypeSpecial3:
			layerKey = "S3"
			layerValue = p.object.S3
		case d2enum.CompositeTypeSpecial4:
			layerKey = "S4"
			layerValue = p.object.S4
		case d2enum.CompositeTypeSpecial5:
			layerKey = "S5"
			layerValue = p.object.S5
		case d2enum.CompositeTypeSpecial6:
			layerKey = "S6"
			layerValue = p.object.S6
		case d2enum.CompositeTypeSpecial7:
			layerKey = "S7"
			layerValue = p.object.S7
		case d2enum.CompositeTypeSpecial8:
			layerKey = "S8"
			layerValue = p.object.S8
		default:
			return nil, errors.New("unknown layer type")
		}

		layerPath := fmt.Sprintf(
			"%s/%s/%s/%s%s%s%s%s.dcc",
			p.object.Base,
			p.object.Token,
			layerKey,
			p.object.Token,
			layerKey,
			layerValue,
			mode.animationMode,
			mode.weaponClass,
		)

		dcc, err := LoadDCC(layerPath)
		if err != nil {
			return nil, err
		}

		mode.layers[cofLayer.Type] = dcc
	}

	animationKey := strings.ToLower(p.object.Token + mode.animationMode + mode.weaponClass)
	animationData := d2data.AnimationData[animationKey]
	if len(animationData) == 0 {
		return nil, errors.New("could not find animation data")
	}

	mode.animationSpeed = 1.0 / ((float64(animationData[0].AnimationSpeed) * 25.0) / 256.0)
	mode.lastFrameTime = d2helper.Now()
	mode.frameCount = animationData[0].FramesPerDirection

	var dccDirection int
	switch cof.NumberOfDirections {
	case 4:
		dccDirection = d2dcc.CofToDir4[mode.direction]
	case 8:
		dccDirection = d2dcc.CofToDir8[mode.direction]
	case 16:
		dccDirection = d2dcc.CofToDir16[mode.direction]
	case 32:
		dccDirection = d2dcc.CofToDir32[mode.direction]
	}

	mode.drawOrder = make([][]d2enum.CompositeType, mode.frameCount)
	for frame := 0; frame < mode.frameCount; frame++ {
		mode.drawOrder[frame] = cof.Priority[direction][frame]
	}

	mode.layerCache = make([]*paperdollCacheEntry, d2enum.CompositeTypeMax)
	for _, cofLayer := range cof.CofLayers {
		layer := mode.layers[cofLayer.Type]

		minX, minY := math.MaxInt32, math.MaxInt32
		maxX, maxY := math.MinInt32, math.MinInt32
		for _, frame := range layer.Directions[dccDirection].Frames {
			minX = d2helper.MinInt(minX, frame.Box.Left)
			minY = d2helper.MinInt(minY, frame.Box.Top)
			maxX = d2helper.MaxInt(maxX, frame.Box.Right())
			maxY = d2helper.MaxInt(maxY, frame.Box.Bottom())
		}

		cacheEntry := &paperdollCacheEntry{
			offsetX: minX,
			offsetY: minY,
			width:   maxX - minX,
			height:  maxY - minY,
		}

		if cacheEntry.width <= 0 || cacheEntry.height <= 0 {
			return nil, errors.New("invalid animation size")
		}

		var transparency int
		if cofLayer.Transparent {
			switch cofLayer.DrawEffect {
			case d2enum.DrawEffectPctTransparency25:
				transparency = 64
			case d2enum.DrawEffectPctTransparency50:
				transparency = 128
			case d2enum.DrawEffectPctTransparency75:
				transparency = 192
			case d2enum.DrawEffectModulate:
				cacheEntry.compositeMode = ebiten.CompositeModeLighter
			default:
				transparency = 255
			}
		}

		pixels := make([]byte, mode.frameCount*cacheEntry.width*cacheEntry.height*4)

		for i := 0; i < mode.frameCount; i++ {
			direction := layer.Directions[dccDirection]
			if i >= len(direction.Frames) {
				return nil, errors.New("invalid animation index")
			}

			sheetOffset := cacheEntry.width * i
			sheetWidth := cacheEntry.height * mode.frameCount

			frame := direction.Frames[i]
			for y := 0; y < direction.Box.Height; y++ {
				for x := 0; x < direction.Box.Width; x++ {
					if paletteIndex := frame.PixelData[x+(y*direction.Box.Width)]; paletteIndex != 0 {
						color := p.palette.Colors[paletteIndex]
						frameX := (x + direction.Box.Left) - minX
						frameY := (y + direction.Box.Top) - minY
						offset := (sheetOffset + frameX + (frameY * sheetWidth)) * 4
						pixels[offset] = color.R
						pixels[offset+1] = color.G
						pixels[offset+2] = color.B
						pixels[offset+3] = byte(transparency)
					}
				}
			}
		}

		cacheEntry.sheetImage, err = ebiten.NewImage(cacheEntry.width*mode.frameCount, cacheEntry.height, ebiten.FilterNearest)
		if err != nil {
			return nil, err
		}

		if err := cacheEntry.sheetImage.ReplacePixels(pixels); err != nil {
			return nil, err
		}

		mode.layerCache[cofLayer.Type] = cacheEntry
	}

	return mode, nil
}
