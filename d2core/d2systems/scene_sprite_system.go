package d2systems

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/gravestench/akara"
	"time"

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
		cache:  d2cache.CreateCache(spriteCacheBudget),
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
	RenderSystem    *RenderSystem
	loadQueue       spriteLoadQueue
	spritesToRender *akara.Subscription
	spritesToUpdate *akara.Subscription
	cache           d2interface.Cache
	Components      struct {
		File            d2components.FileFactory
		Transform       d2components.TransformFactory
		Dc6             d2components.Dc6Factory
		Dcc             d2components.DccFactory
		Palette         d2components.PaletteFactory
		Sprite          d2components.SpriteFactory
		Texture         d2components.TextureFactory
		Origin          d2components.OriginFactory
		Ready           d2components.ReadyFactory
		SegmentedSprite d2components.SegmentedSpriteFactory
	}
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
	t.InjectComponent(&d2components.File{}, &t.Components.File.ComponentFactory)
	t.InjectComponent(&d2components.Transform{}, &t.Components.Transform.ComponentFactory)
	t.InjectComponent(&d2components.Dc6{}, &t.Components.Dc6.ComponentFactory)
	t.InjectComponent(&d2components.Dcc{}, &t.Components.Dcc.ComponentFactory)
	t.InjectComponent(&d2components.Palette{}, &t.Components.Palette.ComponentFactory)
	t.InjectComponent(&d2components.Texture{}, &t.Components.Texture.ComponentFactory)
	t.InjectComponent(&d2components.Origin{}, &t.Components.Origin.ComponentFactory)
	t.InjectComponent(&d2components.Sprite{}, &t.Components.Sprite.ComponentFactory)
	t.InjectComponent(&d2components.SegmentedSprite{}, &t.Components.SegmentedSprite.ComponentFactory)
	t.InjectComponent(&d2components.Ready{}, &t.Components.Ready.ComponentFactory)
}

func (t *SpriteFactory) setupSubscriptions() {
	spritesToRender := t.NewComponentFilter().
		Require(&d2components.Sprite{}). // we want to process entities that have an sprite ...
		Forbid(&d2components.Texture{}). // ... but are missing a surface
		Build()

	spritesToUpdate := t.NewComponentFilter().
		Require(&d2components.Sprite{}). // we want to process entities that have an sprite ...
		Require(&d2components.Texture{}). // ... but are missing a surface
		Build()

	t.spritesToRender = t.AddSubscription(spritesToRender)
	t.spritesToUpdate = t.AddSubscription(spritesToUpdate)
}

// Update processes the load queue which attempting to create sprites, as well as
// binding existing sprites to a renderer if one is present.
func (t *SpriteFactory) Update() {
	start := time.Now()

	for spriteID := range t.loadQueue {
		t.tryCreatingSprite(spriteID)
	}

	if time.Since(start) > maxTimePerUpdate {
		return
	}

	for _, eid := range t.spritesToUpdate.GetEntities() {
		t.updateSprite(eid)
	}

	if time.Since(start) > maxTimePerUpdate {
		return
	}

	for _, eid := range t.spritesToRender.GetEntities() {
		t.tryRenderingSprite(eid)
	}

	if time.Since(start) > maxTimePerUpdate {
		return
	}
}

// ComponentFactory queues a sprite spriteation to be loaded
func (t *SpriteFactory) Sprite(x, y float64, imgPath, palPath string) akara.EID {
	spriteID := t.NewEntity()

	transform := t.Components.Transform.Add(spriteID)
	transform.Translation.X, transform.Translation.Y = x, y

	imgID, palID := t.NewEntity(), t.NewEntity()
	t.Components.File.Add(imgID).Path = imgPath
	t.Components.File.Add(palID).Path = palPath

	t.loadQueue[spriteID] = spriteLoadQueueEntry{
		spriteImage:   imgID,
		spritePalette: palID,
	}

	return spriteID
}

// ComponentFactory queues a segmented sprite spriteation to be loaded.
// A segmented sprite is a sprite that has many frames that form the entire sprite.
func (t *SpriteFactory) SegmentedSprite(x, y float64, imgPath, palPath string, xseg, yseg, frame int) akara.EID {
	spriteID := t.Sprite(x, y, imgPath, palPath)

	s := t.Components.SegmentedSprite.Add(spriteID)
	s.Xsegments = xseg
	s.Ysegments = yseg
	s.FrameOffset = frame

	return spriteID
}

func (t *SpriteFactory) tryCreatingSprite(id akara.EID) {
	entry := t.loadQueue[id]
	imageID, paletteID := entry.spriteImage, entry.spritePalette

	imageFile, found := t.Components.File.Get(imageID)
	if !found {
		return
	}

	paletteFile, found := t.Components.File.Get(paletteID)
	if !found {
		return
	}

	palette, found := t.Components.Palette.Get(paletteID)
	if !found {
		return
	}

	var sprite d2interface.Sprite

	var err error

	cacheKey := spriteCacheKey(imageFile.Path, paletteFile.Path)
	if iface, found := t.cache.Retrieve(cacheKey); found {
		sprite = iface.(d2interface.Sprite)
	}

	if dc6, found := t.Components.Dc6.Get(imageID); found && sprite == nil {
		sprite, err = t.createDc6Sprite(dc6, palette)
		_ = t.cache.Insert(cacheKey, sprite, 1)
	}

	if dcc, found := t.Components.Dcc.Get(imageID); found && sprite == nil {
		sprite, err = t.createDccSprite(dcc, palette)
		_ = t.cache.Insert(cacheKey, sprite, 1)
	}

	if err != nil {
		t.Errorf(fmtCreateSpriteErr, imageFile.Path, paletteFile.Path)

		t.RemoveEntity(id)
		t.RemoveEntity(imageID)
		t.RemoveEntity(paletteID)
	}

	spriteComponent := t.Components.Sprite.Add(id)
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

	sprite, found := t.Components.Sprite.Get(eid)
	if !found {
		return
	}

	if sprite.Sprite == nil {
		return
	}

	sprite.BindRenderer(t.RenderSystem.renderer)

	// if it's a segmented sprite, we need to handle that
	if segmented, isSegmented := t.Components.SegmentedSprite.Get(eid); isSegmented {
		_, isReady := t.Components.Ready.Get(eid)
		if !isReady {
			t.renderSegmentedSprite(eid, segmented)
		}

		return
	}

	// otherwise, it's just a normal sprite
	sfc := sprite.GetCurrentFrameSurface()

	t.Components.Texture.Add(eid).Texture = sfc
	t.Components.Ready.Add(eid)
}

func (t *SpriteFactory) renderSegmentedSprite(id akara.EID, seg *d2components.SegmentedSprite) {
	sprite, found := t.Components.Sprite.Get(id)
	if !found {
		return
	}

	var offsetY int

	segmentsX, segmentsY := seg.Xsegments, seg.Ysegments
	frameOffset := seg.FrameOffset
	numFrames := sprite.GetFrameCount()

	// find max width+height
	fullWidth, fullHeight := 0, 0

	// first, we're going to determine the width and height of the texture we need
	for y := 0; y < segmentsY; y++ {
		fullWidth = 0
		for x := 0; x < segmentsX; x++ {
			idx := x + y*segmentsX + frameOffset
			if idx >= numFrames {
				continue
			}

			if err := sprite.SetCurrentFrame(idx); err != nil {
				fmtErr := "SetCurrentFrame error %s: \n\tsprite: %v\n\tframe count: %v\n\tframe tried: %v\n\t%v"

				t.Errorf(fmtErr, err.Error(), sprite.SpritePath, sprite.GetFrameCount(), idx, seg)
			}

			frameWidth, _ := sprite.GetCurrentFrameSize()
			fullWidth += frameWidth
		}

		_, frameHeight := sprite.GetCurrentFrameSize()
		fullHeight += frameHeight
	}

	target := t.RenderSystem.renderer.NewSurface(fullWidth, fullHeight)

	// now, we'll actually render the sprite to the texture
	for y := 0; y < segmentsY; y++ {
		var offsetX, maxFrameHeight int

		for x := 0; x < segmentsX; x++ {
			idx := x + y*segmentsX + frameOffset
			if idx >= numFrames {
				continue
			}

			if err := sprite.SetCurrentFrame(idx); err != nil {
				fmtErr := "SetCurrentFrame error %s: \n\tsprite: %v\n\tframe count: %v\n\tframe tried: %v\n\t%v"

				t.Errorf(fmtErr, err.Error(), sprite.SpritePath, sprite.GetFrameCount(), idx, seg)
			}

			target.PushTranslation(x+offsetX, y+offsetY)
			target.Render(sprite.GetCurrentFrameSurface())
			target.Pop()

			frameWidth, frameHeight := sprite.GetCurrentFrameSize()
			maxFrameHeight = d2math.MaxInt(maxFrameHeight, frameHeight)
			offsetX += frameWidth - 1
		}

		offsetY += maxFrameHeight - 1
	}

	// now, we're putting the texture into the component
	t.Components.Texture.Add(id).Texture = target
	t.Components.Ready.Add(id)
}

func (t *SpriteFactory) updateSprite(eid akara.EID) {
	if t.RenderSystem == nil {
		return
	}

	if t.RenderSystem.renderer == nil {
		return
	}

	sprite, found := t.Components.Sprite.Get(eid)
	if !found {
		return
	}

	if sprite.Sprite == nil {
		return
	}

	texture, found := t.Components.Texture.Get(eid)
	if !found {
		return
	}

	origin, found := t.Components.Origin.Get(eid)
	if !found {
		origin = t.Components.Origin.Add(eid)
	}

	_ = sprite.Sprite.Advance(t.World.TimeDelta)

	texture.Texture = sprite.GetCurrentFrameSurface()

	ox, oy := sprite.GetCurrentFrameOffset()
	origin.X, origin.Y = float64(ox), float64(oy)

	if _, isSegmented := t.Components.SegmentedSprite.Get(eid); isSegmented {
		return
	}

	// for some shitty reason we need to offset vertically by the sprite height.
	// this comes from the original implementation, we need to fix this shit.
	_, frameHeight := sprite.GetCurrentFrameSize()
	origin.Y -= float64(frameHeight)
}

func (t *SpriteFactory) createDc6Sprite(dc6 *d2components.Dc6, pal *d2components.Palette) (d2interface.Sprite, error) {
	return d2sprite.NewDC6Sprite(dc6.DC6, pal.Palette, 0)
}

func (t *SpriteFactory) createDccSprite(dcc *d2components.Dcc, pal *d2components.Palette) (d2interface.Sprite, error) {
	return d2sprite.NewDCCSprite(dcc.DCC, pal.Palette, 0)
}

func spriteCacheKey(imgpath, palpath string) string {
	return fmt.Sprintf("%s::%s", imgpath, palpath)
}
