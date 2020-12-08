package d2systems

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2sprite"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	fmtCreateSpriteErr = "could not create sprite from image `%s` and palette `%s`"
)

const (
	spriteCacheBudget = 1024
)

// NewSpriteFactory creates a new sprite factory which is intended
// to be embedded in the game object factory system.
func NewSpriteFactory(b akara.BaseSystem, l *d2util.Logger) *SpriteFactory {
	sys := &SpriteFactory{
		Logger: l,
		cache: d2cache.CreateCache(spriteCacheBudget),
	}

	sys.BaseSystem = b

	sys.World.AddSystem(sys)

	return sys
}

type spriteLoadQueueEntry struct {
	spriteImage, spritePalette akara.EID
}

type spriteLoadQueue = map[akara.EID]spriteLoadQueueEntry

// SpriteFactory is responsible for queueing sprites to be loaded (as sprites),
// as well as binding the spriteation to a renderer if one is present (which generates the sprite surfaces).
type SpriteFactory struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	RenderSystem *RenderSystem
	d2components.FileFactory
	d2components.TransformFactory
	d2components.Dc6Factory
	d2components.DccFactory
	d2components.PaletteFactory
	d2components.SpriteFactory
	d2components.TextureFactory
	d2components.OriginFactory
	d2components.SegmentedSpriteFactory
	loadQueue       spriteLoadQueue
	spritesToRender *akara.Subscription
	spritesToUpdate *akara.Subscription
	cache d2interface.Cache
}

// Init the sprite factory, injecting the necessary components
func (t *SpriteFactory) Init(world *akara.World) {
	t.World = world

	t.Debug("initializing sprite factory ...")

	t.setupFactories()
	t.setupSubscriptions()

	t.loadQueue = make(spriteLoadQueue)
}

func (t *SpriteFactory) setupFactories() {
	t.InjectComponent(&d2components.File{}, &t.File)
	t.InjectComponent(&d2components.Transform{}, &t.Transform)
	t.InjectComponent(&d2components.Dc6{}, &t.Dc6)
	t.InjectComponent(&d2components.Dcc{}, &t.Dcc)
	t.InjectComponent(&d2components.Palette{}, &t.Palette)
	t.InjectComponent(&d2components.Texture{}, &t.Texture)
	t.InjectComponent(&d2components.Origin{}, &t.Origin)
	t.InjectComponent(&d2components.Sprite{}, &t.SpriteFactory.Sprite)
	t.InjectComponent(&d2components.SegmentedSprite{}, &t.SegmentedSpriteFactory.SegmentedSprite)
}

func (t *SpriteFactory) setupSubscriptions() {
	spritesToRender := t.NewComponentFilter().
		Require(&d2components.Sprite{}). // we want to process entities that have an sprite ...
		Forbid(&d2components.Texture{}). // ... but are missing a surface
		Build()

	spritesToUpdate := t.NewComponentFilter().
		Require(&d2components.Sprite{}).  // we want to process entities that have an sprite ...
		Require(&d2components.Texture{}). // ... but are missing a surface
		Build()

	t.spritesToRender = t.AddSubscription(spritesToRender)
	t.spritesToUpdate = t.AddSubscription(spritesToUpdate)
}

// Update processes the load queue which attempting to create sprites, as well as
// binding existing sprites to a renderer if one is present.
func (t *SpriteFactory) Update() {
	for spriteID := range t.loadQueue {
		t.tryCreatingSprite(spriteID)
	}

	for _, eid := range t.spritesToUpdate.GetEntities() {
		t.updateSprite(eid)
	}

	for _, eid := range t.spritesToRender.GetEntities() {
		t.tryRenderingSprite(eid)
	}
}

// Sprite queues a sprite spriteation to be loaded
func (t *SpriteFactory) Sprite(x, y float64, imgPath, palPath string) akara.EID {
	spriteID := t.NewEntity()

	transform := t.AddTransform(spriteID)
	transform.Translation.X, transform.Translation.Y = x, y

	imgID, palID := t.NewEntity(), t.NewEntity()
	t.AddFile(imgID).Path = imgPath
	t.AddFile(palID).Path = palPath

	t.loadQueue[spriteID] = spriteLoadQueueEntry{
		spriteImage:   imgID,
		spritePalette: palID,
	}

	return spriteID
}

// SegmentedSprite queues a segmented sprite spriteation to be loaded.
// A segmented sprite is a sprite that has many frames that form the entire sprite.
func (t *SpriteFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	spriteID := t.Sprite(x, y, imgPath, palPath)

	s := t.AddSegmentedSprite(spriteID)
	s.Xsegments = xseg
	s.Ysegments = yseg
	s.FrameOffset = frame

	return spriteID
}

func (t *SpriteFactory) tryCreatingSprite(id akara.EID) {
	entry := t.loadQueue[id]
	imageID, paletteID := entry.spriteImage, entry.spritePalette

	imageFile, found := t.GetFile(imageID)
	if !found {
		return
	}

	paletteFile, found := t.GetFile(paletteID)
	if !found {
		return
	}

	palette, found := t.GetPalette(paletteID)
	if !found {
		return
	}

	var sprite d2interface.Sprite

	var err error

	cacheKey := spriteCacheKey(imageFile.Path, paletteFile.Path)
	if iface, found := t.cache.Retrieve(cacheKey); found {
		sprite = iface.(d2interface.Sprite)
	}

	if dc6, found := t.GetDc6(imageID); found && sprite == nil {
		sprite, err = t.createDc6Sprite(dc6, palette)
		_ = t.cache.Insert(cacheKey, sprite, 1)
	}

	if dcc, found := t.GetDcc(imageID); found && sprite == nil {
		sprite, err = t.createDccSprite(dcc, palette)
		_ = t.cache.Insert(cacheKey, sprite, 1)
	}

	if err != nil {
		t.Errorf(fmtCreateSpriteErr, imageFile.Path, paletteFile.Path)

		t.RemoveEntity(id)
		t.RemoveEntity(imageID)
		t.RemoveEntity(paletteID)
	}

	spriteComponent := t.AddSprite(id)
	spriteComponent.Sprite = sprite
	spriteComponent.SpritePath = imageFile.Path
	spriteComponent.PalettePath = paletteFile.Path

	delete(t.loadQueue, id)
}

func (t *SpriteFactory) tryRenderingSprite(eid akara.EID) {
	if t.RenderSystem == nil {
		return
	}

	if t.RenderSystem.renderer == nil {
		return
	}

	sprite, found := t.GetSprite(eid)
	if !found {
		return
	}

	if sprite.Sprite == nil {
		return
	}

	sprite.BindRenderer(t.RenderSystem.renderer)

	sfc := sprite.GetCurrentFrameSurface()

	t.AddTexture(eid).Texture = sfc
}

func (t *SpriteFactory) updateSprite(eid akara.EID) {
	if t.RenderSystem == nil {
		return
	}

	if t.RenderSystem.renderer == nil {
		return
	}

	sprite, found := t.GetSprite(eid)
	if !found {
		return
	}

	if sprite.Sprite == nil {
		return
	}

	texture, found := t.GetTexture(eid)
	if !found {
		return
	}

	origin, found := t.GetOrigin(eid)
	if !found {
		origin = t.AddOrigin(eid)
	}

	_ = sprite.Sprite.Advance(t.World.TimeDelta)

	texture.Texture = sprite.GetCurrentFrameSurface()

	ox, oy := sprite.GetCurrentFrameOffset()
	origin.X, origin.Y = float64(ox), float64(oy)

	if _, isSegmented := t.GetSegmentedSprite(eid); !isSegmented {
		_, frameHeight := sprite.GetCurrentFrameSize()
		origin.Y -= float64(frameHeight)
	}
}

func (t *SpriteFactory) createDc6Sprite(
	dc6 *d2components.Dc6,
	pal *d2components.Palette,
) (d2interface.Sprite, error) {
	return d2sprite.NewDC6Sprite(dc6.DC6, pal.Palette, 0)
}

func (t *SpriteFactory) createDccSprite(
	dcc *d2components.Dcc,
	pal *d2components.Palette,
) (d2interface.Sprite, error) {
	return d2sprite.NewDCCSprite(dcc.DCC, pal.Palette, 0)
}

func spriteCacheKey(imgpath, palpath string) string {
	return fmt.Sprintf("%s::%s", imgpath, palpath)
}
