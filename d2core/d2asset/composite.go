package d2asset

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// Composite is a composite entity animation
type Composite struct {
	baseType    d2enum.ObjectType
	basePath    string
	token       string
	palettePath string
	direction   int
	equipment   [d2enum.CompositeTypeMax]string
	mode        *compositeMode
}

// CreateComposite creates a Composite from a given ObjectLookupRecord and palettePath.
func CreateComposite(baseType d2enum.ObjectType, token, palettePath string) *Composite {
	return &Composite{baseType: baseType, basePath: baseString(baseType),
		token: token, palettePath: palettePath}
}

// Advance moves the composite animation forward for a given elapsed time in nanoseconds.
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

// Render performs drawing of the Composite on the rendered d2interface.Surface.
func (c *Composite) Render(target d2interface.Surface) error {
	if c.mode == nil {
		return nil
	}

	direction := d2cof.Dir64ToCof(c.direction, c.mode.cof.NumberOfDirections)
	for _, layerIndex := range c.mode.cof.Priority[direction][c.mode.frameIndex] {
		layer := c.mode.layers[layerIndex]
		if layer != nil {
			if err := layer.RenderFromOrigin(target); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetAnimationMode returns the animation mode the Composite should render with.
func (c *Composite) GetAnimationMode() string {
	return c.mode.animationMode
}

// GetWeaponClass returns the currently loaded weapon class
func (c *Composite) GetWeaponClass() string {
	return c.mode.weaponClass
}

// SetMode sets the Composite's animation mode weapon class and direction
func (c *Composite) SetMode(animationMode, weaponClass string) error {
	if c.mode != nil && c.mode.animationMode == animationMode && c.mode.weaponClass == weaponClass {
		return nil
	}

	mode, err := c.createMode(animationMode, weaponClass)
	if err != nil {
		return err
	}

	c.resetPlayedCount()
	c.mode = mode

	return nil
}

// Equip changes the current layer configuration
func (c *Composite) Equip(equipment *[d2enum.CompositeTypeMax]string) error {
	c.equipment = *equipment
	if c.mode == nil {
		return nil
	}

	mode, err := c.createMode(c.mode.animationMode, c.mode.weaponClass)

	if err != nil {
		return err
	}

	c.mode = mode

	return nil
}

// SetAnimSpeed sets the speed at which the Composite's animation should advance through its frames
func (c *Composite) SetAnimSpeed(speed int) {
	c.mode.animationSpeed = 1.0 / ((float64(speed) * 25.0) / 256.0)
	for layerIdx := range c.mode.layers {
		layer := c.mode.layers[layerIdx]
		if layer != nil {
			layer.SetPlaySpeed(c.mode.animationSpeed)
		}
	}
}

// SetDirection sets the direction of the composite and its layers
func (c *Composite) SetDirection(direction int) {
	c.direction = direction
	for layerIdx := range c.mode.layers {
		layer := c.mode.layers[layerIdx]
		if layer != nil {
			layer.SetDirection(c.direction)
		}
	}
}

// GetDirection returns the current direction the composite is facing
func (c *Composite) GetDirection() int {
	return c.direction
}

// GetPlayedCount returns the number of times the current animation mode has completed all its distinct frames
func (c *Composite) GetPlayedCount() int {
	if c.mode == nil {
		return 0
	}

	return c.mode.playedCount
}

func (c *Composite) resetPlayedCount() {
	if c.mode != nil {
		c.mode.playedCount = 0
	}
}

type compositeMode struct {
	cof           *d2cof.COF
	animationMode string
	weaponClass   string
	playedCount   int

	layers []*Animation

	frameCount     int
	frameIndex     int
	animationSpeed float64
	lastFrameTime  float64
}

func (c *Composite) createMode(animationMode, weaponClass string) (*compositeMode, error) {
	cofPath := fmt.Sprintf("%s/%s/COF/%s%s%s.COF", c.basePath, c.token, c.token, animationMode, weaponClass)
	if exists, _ := FileExists(cofPath); !exists {
		return nil, errors.New("composite not found")
	}

	cof, err := loadCOF(cofPath)
	if err != nil {
		return nil, err
	}

	animationKey := strings.ToLower(c.token + animationMode + weaponClass)

	animationData := d2data.AnimationData[animationKey]
	if len(animationData) == 0 {
		return nil, errors.New("could not find animation data")
	}

	mode := &compositeMode{
		cof:            cof,
		animationMode:  animationMode,
		weaponClass:    weaponClass,
		layers:         make([]*Animation, d2enum.CompositeTypeMax),
		frameCount:     animationData[0].FramesPerDirection,
		animationSpeed: 1.0 / ((float64(animationData[0].AnimationSpeed) * 25.0) / 256.0),
	}

	for _, cofLayer := range cof.CofLayers {
		layerValue := c.equipment[cofLayer.Type]
		if layerValue == "" {
			layerValue = "lit"
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

		layer, err := c.loadCompositeLayer(cofLayer.Type.String(), layerValue, animationMode,
			cofLayer.WeaponClass.String(), c.palettePath, transparency)
		if err == nil {
			layer.SetPlaySpeed(mode.animationSpeed)
			layer.PlayForward()
			layer.SetBlend(blend)

			if err := layer.SetDirection(c.direction); err != nil {
				return nil, err
			}

			mode.layers[cofLayer.Type] = layer
		}
	}

	return mode, nil
}

func (c *Composite) loadCompositeLayer(layerKey, layerValue, animationMode, weaponClass,
	palettePath string, transparency int) (*Animation, error) {
	animationPaths := []string{
		fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dcc", c.basePath, c.token, layerKey, c.token, layerKey, layerValue, animationMode, weaponClass),
		fmt.Sprintf("%s/%s/%s/%s%s%s%s%s.dc6", c.basePath, c.token, layerKey, c.token, layerKey, layerValue, animationMode, weaponClass),
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

func baseString(baseType d2enum.ObjectType) string {
	switch baseType {
	case d2enum.ObjectTypePlayer:
		return "/data/global/chars"
	case d2enum.ObjectTypeCharacter:
		return "/data/global/monsters"
	case d2enum.ObjectTypeItem:
		return "/data/global/objects"
	default:
		return ""
	}
}
