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

func NewSpriteFactorySubsystem(b *akara.BaseSystem, l *d2util.Logger) *SpriteFactory {
	spritesToRender := akara.NewFilter().
		Require(d2components.Animation). // we want to process entities that have an animation ...
		Forbid(d2components.Surface). // ... but are missing a surface
		Build()

	sys := &SpriteFactory{
		BaseSubscriberSystem: akara.NewBaseSubscriberSystem(spritesToRender),
		Logger:           l,
	}

	sys.BaseSystem = b

	sys.World.AddSystem(sys)

	return sys
}

type spriteLoadQueueEntry struct {
	spriteImage, spritePalette akara.EID
}

type spriteLoadQueue = map[akara.EID]spriteLoadQueueEntry

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
	*d2components.SurfaceMap
	loadQueue       spriteLoadQueue
	spritesToRender *akara.Subscription
}

func (t *SpriteFactory) Init(world *akara.World) {
	t.Info("initializing sprite factory ...")

	t.loadQueue = make(spriteLoadQueue, 0)

	t.spritesToRender = t.Subscriptions[0]

	t.FilePathMap = t.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	t.PositionMap = t.InjectMap(d2components.Position).(*d2components.PositionMap)
	t.Dc6Map = t.InjectMap(d2components.Dc6).(*d2components.Dc6Map)
	t.DccMap = t.InjectMap(d2components.Dcc).(*d2components.DccMap)
	t.PaletteMap = t.InjectMap(d2components.Palette).(*d2components.PaletteMap)
	t.AnimationMap = t.InjectMap(d2components.Animation).(*d2components.AnimationMap)
	t.SurfaceMap = t.InjectMap(d2components.Surface).(*d2components.SurfaceMap)
}

func (t *SpriteFactory) Update() {
	for _, eid := range t.spritesToRender.GetEntities() {
		t.tryRenderingSprite(eid)
	}

	for spriteID := range t.loadQueue {
		t.tryCreatingAnimation(spriteID)
	}
}

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

	anim.BindRenderer(t.renderer)

	sfc := anim.GetCurrentFrameSurface()

	t.AddSurface(eid).Surface = sfc
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
