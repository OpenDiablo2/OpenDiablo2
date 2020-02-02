package d2asset

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Composite struct {
	object      *d2datadict.ObjectLookupRecord
	palettePath string
	mode        *compositeMode
}

func CreateComposite(object *d2datadict.ObjectLookupRecord, palettePath string) *Composite {
	return &Composite{object: object, palettePath: palettePath}
}

func (c *Composite) Advance(elapsed float64) error {
	if c.mode == nil {
		return nil
	}

	c.mode.lastFrameTime += elapsed
	framesToAdd := int(c.mode.lastFrameTime / c.mode.animationSpeed)
	c.mode.lastFrameTime -= float64(framesToAdd) * c.mode.animationSpeed
	c.mode.frameIndex += framesToAdd
	c.mode.playedCount += c.mode.frameIndex / c.mode.frameCount
	c.mode.frameIndex %= c.mode.frameCount

	for _, layer := range c.mode.layers {
		if layer != nil {
			if err := layer.Advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Composite) Render(target d2render.Surface) error {
	if c.mode == nil {
		return nil
	}

	for _, layerIndex := range c.mode.drawOrder[c.mode.frameIndex] {
		layer := c.mode.layers[layerIndex]
		if layer != nil {
			if err := layer.Render(target); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Composite) SetMode(animationMode, weaponClass string, direction int) error {
	if c.mode != nil && c.mode.animationMode == animationMode && c.mode.weaponClass == weaponClass && c.mode.direction == direction {
		return nil
	}

	mode, err := c.createMode(animationMode, weaponClass, direction)
	if err != nil {
		return err
	}

	c.mode = mode
	return nil
}

func (c *Composite) GetDirectionCount() int {
	if c.mode == nil {
		return 0
	}

	return c.mode.directionCount
}

func (c *Composite) GetPlayedCount() int {
	if c.mode == nil {
		return 0
	}

	return c.mode.playedCount
}

func (c *Composite) ResetPlayedCount() {
	if c.mode != nil {
		c.mode.playedCount = 0
	}
}

type compositeMode struct {
	animationMode  string
	weaponClass    string
	direction      int
	directionCount int
	playedCount    int

	layers    []*Animation
	drawOrder [][]d2enum.CompositeType

	frameCount     int
	frameIndex     int
	animationSpeed float64
	lastFrameTime  float64
}

func (c *Composite) createMode(animationMode, weaponClass string, direction int) (*compositeMode, error) {
	cofPath := fmt.Sprintf("%s/%s/COF/%s%s%s.COF", c.object.Base, c.object.Token, c.object.Token, animationMode, weaponClass)
	if exists, _ := FileExists(cofPath); !exists {
		return nil, errors.New("composite not found")
	}

	cof, err := loadCOF(cofPath)
	if err != nil {
		return nil, err
	}

	if direction >= cof.NumberOfDirections {
		return nil, errors.New("invalid direction")
	}

	animationKey := strings.ToLower(c.object.Token + animationMode + weaponClass)
	animationData := d2data.AnimationData[animationKey]
	if len(animationData) == 0 {
		return nil, errors.New("could not find animation data")
	}

	mode := &compositeMode{
		animationMode:  animationMode,
		weaponClass:    weaponClass,
		direction:      direction,
		directionCount: cof.NumberOfDirections,
		layers:         make([]*Animation, d2enum.CompositeTypeMax),
		frameCount:     animationData[0].FramesPerDirection,
		animationSpeed: 1.0 / ((float64(animationData[0].AnimationSpeed) * 25.0) / 256.0),
	}

	mode.drawOrder = make([][]d2enum.CompositeType, mode.frameCount)
	for frame := 0; frame < mode.frameCount; frame++ {
		mode.drawOrder[frame] = cof.Priority[direction][frame]
	}

	var layerDirection int
	switch cof.NumberOfDirections {
	case 4:
		layerDirection = d2dcc.CofToDir4[mode.direction]
	case 8:
		layerDirection = d2dcc.CofToDir8[mode.direction]
	case 16:
		layerDirection = d2dcc.CofToDir16[mode.direction]
	case 32:
		layerDirection = d2dcc.CofToDir32[mode.direction]
	}

	for _, cofLayer := range cof.CofLayers {
		var layerKey, layerValue string
		switch cofLayer.Type {
		case d2enum.CompositeTypeHead:
			layerKey = "HD"
			layerValue = c.object.HD
		case d2enum.CompositeTypeTorso:
			layerKey = "TR"
			layerValue = c.object.TR
		case d2enum.CompositeTypeLegs:
			layerKey = "LG"
			layerValue = c.object.LG
		case d2enum.CompositeTypeRightArm:
			layerKey = "RA"
			layerValue = c.object.RA
		case d2enum.CompositeTypeLeftArm:
			layerKey = "LA"
			layerValue = c.object.LA
		case d2enum.CompositeTypeRightHand:
			layerKey = "RH"
			layerValue = c.object.RH
		case d2enum.CompositeTypeLeftHand:
			layerKey = "LH"
			layerValue = c.object.LH
		case d2enum.CompositeTypeShield:
			layerKey = "SH"
			layerValue = c.object.SH
		case d2enum.CompositeTypeSpecial1:
			layerKey = "S1"
			layerValue = c.object.S1
		case d2enum.CompositeTypeSpecial2:
			layerKey = "S2"
			layerValue = c.object.S2
		case d2enum.CompositeTypeSpecial3:
			layerKey = "S3"
			layerValue = c.object.S3
		case d2enum.CompositeTypeSpecial4:
			layerKey = "S4"
			layerValue = c.object.S4
		case d2enum.CompositeTypeSpecial5:
			layerKey = "S5"
			layerValue = c.object.S5
		case d2enum.CompositeTypeSpecial6:
			layerKey = "S6"
			layerValue = c.object.S6
		case d2enum.CompositeTypeSpecial7:
			layerKey = "S7"
			layerValue = c.object.S7
		case d2enum.CompositeTypeSpecial8:
			layerKey = "S8"
			layerValue = c.object.S8
		default:
			return nil, errors.New("unknown layer type")
		}

		blend := false
		transparency := 255
		if cofLayer.Transparent {
			switch cofLayer.DrawEffect {
			case d2enum.DrawEffectPctTransparency25:
				transparency = 64
			case d2enum.DrawEffectPctTransparency50:
				transparency = 128
			case d2enum.DrawEffectPctTransparency75:
				transparency = 192
			case d2enum.DrawEffectModulate:
				blend = true
			}
		}

		layer, err := loadCompositeLayer(c.object, layerKey, layerValue, animationMode, weaponClass, c.palettePath, transparency)
		if err == nil {
			layer.SetPlaySpeed(mode.animationSpeed)
			layer.PlayForward()
			layer.SetBlend(blend)
			layer.SetDirection(layerDirection)
			mode.layers[cofLayer.Type] = layer
		}
	}

	return mode, nil
}

func loadCompositeLayer(object *d2datadict.ObjectLookupRecord, layerKey, layerValue, animationMode, weaponClass, palettePath string, transparency int) (*Animation, error) {
	animationPaths := []string{
		fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dcc", object.Base, object.Token, layerKey, object.Token, layerKey, layerValue, animationMode, weaponClass),
		fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dcc", object.Base, object.Token, layerKey, object.Token, layerKey, layerValue, animationMode, "HTH"),
		fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dc6", object.Base, object.Token, layerKey, object.Token, layerKey, layerValue, animationMode, weaponClass),
		fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dc6", object.Base, object.Token, layerKey, object.Token, layerKey, layerValue, animationMode, "HTH"),
	}

	for _, animationPath := range animationPaths {
		if exists, _ := FileExists(animationPath); exists {
			animation, err := LoadAnimationWithTransparency(animationPath, palettePath, transparency)
			if err == nil {
				return animation, nil
			}
		}
	}

	return nil, errors.New("animation not found")
}
