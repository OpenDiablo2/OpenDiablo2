package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2animation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	fmtCreateSpriteErr = "could not create sprite from image `%s` and palette `%s`"
)

// NewSpriteFactorySubsystem creates a new sprite factory which is intended
// to be embedded in the game object factory system.
func NewSpriteFactorySubsystem(b *akara.BaseSystem, l *d2util.Logger) *SpriteFactory {
	spritesToRender := akara.NewFilter().
		Require(d2components.Animation). // we want to process entities that have an animation ...
		Forbid(d2components.Renderable). // ... but are missing a surface
		Build()

	spritesToUpdate := akara.NewFilter().
		Require(d2components.Animation).  // we want to process entities that have an animation ...
		Require(d2components.Renderable). // ... but are missing a surface
		Build()

	sys := &SpriteFactory{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(spritesToRender, spritesToUpdate),
		Logger:               l,
	}

	sys.BaseSystem = b

	sys.World.AddSystem(sys)

	return sys
}

type spriteLoadQueueEntry struct {
	spriteImage, spritePalette akara.EID
}

type spriteLoadQueue = map[akara.EID]spriteLoadQueueEntry

// SpriteFactory is responsible for queueing sprites to be loaded (as animations),
// as well as binding the animation to a renderer if one is present (which generates the sprite surfaces).
type SpriteFactory struct {
	*akara.BaseSubscriberSystem
	*d2util.Logger
	*RenderSystem
	*d2components.FilePathMap
	*d2components.PositionMap
	*d2components.Dc6Map
	*d2components.DccMap
	*d2components.PaletteMap
	*d2components.AnimationMap
	*d2components.RenderableMap
	*d2components.SegmentedSpriteMap
	loadQueue       spriteLoadQueue
	spritesToRender *akara.Subscription
	spritesToUpdate *akara.Subscription
}

// Init the sprite factory, injecting the necessary components
func (t *SpriteFactory) Init(_ *akara.World) {
	t.Info("initializing sprite factory ...")

	t.loadQueue = make(spriteLoadQueue)

	t.spritesToRender = t.Subscriptions[0]
	t.spritesToUpdate = t.Subscriptions[1]

	t.FilePathMap = t.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	t.PositionMap = t.InjectMap(d2components.Position).(*d2components.PositionMap)
	t.Dc6Map = t.InjectMap(d2components.Dc6).(*d2components.Dc6Map)
	t.DccMap = t.InjectMap(d2components.Dcc).(*d2components.DccMap)
	t.PaletteMap = t.InjectMap(d2components.Palette).(*d2components.PaletteMap)
	t.AnimationMap = t.InjectMap(d2components.Animation).(*d2components.AnimationMap)
	t.RenderableMap = t.InjectMap(d2components.Renderable).(*d2components.RenderableMap)
	t.SegmentedSpriteMap = t.InjectMap(d2components.SegmentedSprite).(*d2components.SegmentedSpriteMap)
}

// Update processes the load queue which attempting to create animations, as well as
// binding existing animations to a renderer if one is present.
func (t *SpriteFactory) Update() {
	for _, eid := range t.spritesToUpdate.GetEntities() {
		t.updateSprite(eid)
	}

	for _, eid := range t.spritesToRender.GetEntities() {
		t.tryRenderingSprite(eid)
	}

	for spriteID := range t.loadQueue {
		t.tryCreatingAnimation(spriteID)
	}
}

// Sprite queues a sprite animation to be loaded
func (t *SpriteFactory) Sprite(x, y float64, imgPath, palPath string) akara.EID {
	spriteID := t.NewEntity()

	t.AddPosition(spriteID).Set(x, y)

	imgID, palID := t.NewEntity(), t.NewEntity()
	t.AddFilePath(imgID).Path = imgPath
	t.AddFilePath(palID).Path = palPath

	t.loadQueue[spriteID] = spriteLoadQueueEntry{
		spriteImage:   imgID,
		spritePalette: palID,
	}

	return spriteID
}

// SegmentedSprite queues a segmented sprite animation to be loaded.
// A segmented sprite is a sprite that has many frames that form the entire sprite.
func (t *SpriteFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	spriteID := t.Sprite(x, y, imgPath, palPath)

	s := t.AddSegmentedSprite(spriteID)
	s.Xsegments = xseg
	s.Ysegments = yseg
	s.FrameOffset = frame

	return spriteID
}

func (t *SpriteFactory) tryCreatingAnimation(id akara.EID) {
	entry := t.loadQueue[id]
	imageID, paletteID := entry.spriteImage, entry.spritePalette

	imagePath, found := t.GetFilePath(imageID)
	if !found {
		return
	}

	palettePath, found := t.GetFilePath(paletteID)
	if !found {
		return
	}

	palette, found := t.GetPalette(paletteID)
	if !found {
		return
	}

	var anim d2interface.Animation

	var err error

	if dc6, found := t.GetDc6(imageID); found {
		anim, err = t.createDc6Animation(dc6, palette)
	}

	if dcc, found := t.GetDcc(imageID); found {
		anim, err = t.createDccAnimation(dcc, palette)
	}

	if err != nil {
		t.Errorf(fmtCreateSpriteErr, imagePath.Path, palettePath.Path)

		t.RemoveEntity(id)
		t.RemoveEntity(imageID)
		t.RemoveEntity(paletteID)
	}

	t.AddAnimation(id).Animation = anim

	delete(t.loadQueue, id)
}

func (t *SpriteFactory) tryRenderingSprite(eid akara.EID) {
	if t.RenderSystem == nil {
		return
	}

	if t.RenderSystem.renderer == nil {
		return
	}

	anim, found := t.GetAnimation(eid)
	if !found {
		return
	}

	if anim.Animation == nil {
		return
	}

	anim.BindRenderer(t.renderer)

	sfc := anim.GetCurrentFrameSurface()

	t.AddRenderable(eid).Surface = sfc
}

func (t *SpriteFactory) updateSprite(eid akara.EID) {
	if t.RenderSystem == nil {
		return
	}

	if t.RenderSystem.renderer == nil {
		return
	}

	anim, found := t.GetAnimation(eid)
	if !found {
		return
	}

	if anim.Animation == nil {
		return
	}

	renderable, found := t.GetRenderable(eid)
	if !found {
		return
	}

	renderable.Surface = anim.GetCurrentFrameSurface()
}

func (t *SpriteFactory) createDc6Animation(
	dc6 *d2components.Dc6Component,
	pal *d2components.PaletteComponent,
) (d2interface.Animation, error) {
	return d2animation.NewDC6Animation(dc6.DC6, pal.Palette, 0)
}

func (t *SpriteFactory) createDccAnimation(
	dcc *d2components.DccComponent,
	pal *d2components.PaletteComponent,
) (d2interface.Animation, error) {
	return d2animation.NewDCCAnimation(dcc.DCC, pal.Palette, 0)
}
