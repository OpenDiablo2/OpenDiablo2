package d2systems

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2bitmapfont"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	"github.com/gravestench/akara"
	"image/color"
)

const (
	fontCacheBudget = 64
)

// NewWidgetFactory creates a new ui widget factory which is intended
// to be embedded in the game object factory system.
func NewUIWidgetFactory(
	b akara.BaseSystem,
	l *d2util.Logger,
	spriteFactory *SpriteFactory,
	shapeFactory *ShapeSystem,
) *UIWidgetFactory {
	sys := &UIWidgetFactory{
		Logger:          l,
		SpriteFactory:   spriteFactory,
		ShapeSystem:     shapeFactory,
		bitmapFontCache: d2cache.CreateCache(fontCacheBudget),
		buttonLoadQueue: make(buttonLoadQueue),
		labelLoadQueue:  make(labelLoadQueue),
	}

	sys.BaseSystem = b

	sys.World.AddSystem(sys)

	return sys
}

type buttonLoadQueueEntry struct {
	sprite, palette akara.EID
}

type buttonLoadQueue = map[akara.EID]buttonLoadQueueEntry

type labelLoadQueueEntry struct {
	table, sprite akara.EID
}

type labelLoadQueue = map[akara.EID]labelLoadQueueEntry

// UIWidgetFactory is responsible for creating UI widgets like buttons and tabs
type UIWidgetFactory struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	*RenderSystem
	*SpriteFactory
	*ShapeSystem
	buttonLoadQueue
	labelLoadQueue
	bitmapFontCache d2interface.Cache
	d2components.FileFactory
	d2components.TransformFactory
	d2components.InteractiveFactory
	d2components.FontTableFactory
	d2components.PaletteFactory
	d2components.BitmapFontFactory
	d2components.LabelFactory
	labelsToUpdate *akara.Subscription
	booted bool
}

// Init the ui widget factory, injecting the necessary components
func (t *UIWidgetFactory) Init(world *akara.World) {
	t.World = world

	t.Debug("initializing ui widget factory ...")

	t.setupFactories()
	t.setupSubscriptions()
}

func (t *UIWidgetFactory) setupFactories() {
	t.InjectComponent(&d2components.File{}, &t.File)
	t.InjectComponent(&d2components.Transform{}, &t.Transform)
	t.InjectComponent(&d2components.Interactive{}, &t.Interactive)
	t.InjectComponent(&d2components.FontTable{}, &t.FontTable)
	t.InjectComponent(&d2components.Palette{}, &t.Palette)
	t.InjectComponent(&d2components.BitmapFont{}, &t.BitmapFont)
	t.InjectComponent(&d2components.Label{}, &t.LabelFactory.Label)
}

func (t *UIWidgetFactory) setupSubscriptions() {
	labelsToUpdate := t.NewComponentFilter().
		Require(&d2components.Label{}).
		Build()

	t.labelsToUpdate = t.AddSubscription(labelsToUpdate)
}

func (t *UIWidgetFactory) boot() {
	if t.RenderSystem == nil {
		return
	}

	if t.RenderSystem.renderer == nil {
		return
	}

	t.booted = true
}

// Update processes the load queues and update the widgets. The load queues are necessary because
// UI widgets are composed of a bunch of things, which each need to be loaded by other systems (like the asset loader)
func (t *UIWidgetFactory) Update() {
	if !t.booted {
		t.boot()
		return
	}

	for labelEID := range t.labelLoadQueue {
		t.processLabel(labelEID)
	}

	for _, labelEID := range t.labelsToUpdate.GetEntities() {
		t.renderLabel(labelEID)
	}
}

// Label creates a label widget.
//
// The font is assumed to be a path for two files, omiting the file extension
//
// Basically, diablo2 stored bitmap fonts as two files, a glyph table and sprite.
//
// For example, specifying this font: /data/local/FONT/ENG/fontexocet10
//
// will use these two files:
//
// /data/local/FONT/ENG/fontexocet10.dc6
//
// /data/local/FONT/ENG/fontexocet10.tbl
func (t *UIWidgetFactory) Label(text, font, palettePath string) akara.EID {
	tablePath := font + ".tbl"
	spritePath := font + ".dc6"

	labelEID := t.NewEntity()

	tableEID := t.NewEntity()
	t.AddFile(tableEID).Path = tablePath

	spriteEID := t.SpriteFactory.Sprite(0, 0, spritePath, palettePath)

	t.labelLoadQueue[labelEID] = labelLoadQueueEntry{
		table:   tableEID,
		sprite:  spriteEID,
	}

	label := t.AddLabel(labelEID)
	label.SetText(text)

	return labelEID
}

// Label creates a label widget
func (t *UIWidgetFactory) processLabel(labelEID akara.EID) {
	bmfComponent, found := t.GetBitmapFont(labelEID)
	if !found {
		t.addBitmapFontForLabel(labelEID)
		return
	}

	bmfComponent.Sprite.BindRenderer(t.renderer)

	label, found := t.GetLabel(labelEID)
	if !found {
		label = t.AddLabel(labelEID)
	}

	label.Font = bmfComponent.BitmapFont

	t.RemoveEntity(t.labelLoadQueue[labelEID].table)

	delete(t.labelLoadQueue, labelEID)
}

func (t *UIWidgetFactory) renderLabel(labelEID akara.EID) {
	label, found := t.GetLabel(labelEID)
	if !found {
		return
	}

	bmf, found := t.GetBitmapFont(labelEID)
	if !found {
		return
	}

	if label.Font != bmf.BitmapFont {
		label.Font = bmf.BitmapFont
	}

	if !label.IsDirty() {
		return
	}

	texture, found := t.RenderSystem.GetTexture(labelEID)
	if !found {
		texture = t.RenderSystem.AddTexture(labelEID)
	}

	texture.Texture = t.renderer.NewSurface(label.GetSize())

	label.Render(texture.Texture)
}

func (t *UIWidgetFactory) addBitmapFontForLabel(labelEID akara.EID) {
	// get the load queue
	entry, found := t.labelLoadQueue[labelEID]
	if !found {
		return
	}

	// make sure the components have been loaded (by the asset loader)
	_, tableFound := t.GetFontTable(entry.table)
	_, spriteFound := t.GetSprite(entry.sprite)

	if !(tableFound && spriteFound) {
		return
	}

	// now we check the cache, see if we can just pull a pre-rendered bitmap font
	tableFile, found := t.GetFile(entry.table)
	if !found {
		return
	}

	sprite, found := t.GetSprite(entry.sprite)
	if !found {
		return
	}

	cacheKey := fontCacheKey(tableFile.Path, sprite.SpritePath, sprite.PalettePath)

	if iface, found := t.bitmapFontCache.Retrieve(cacheKey); found {
		// we found it, add the bitmap font component and set the embedded struct to what we retrieved
		t.AddBitmapFont(labelEID).BitmapFont = iface.(*d2bitmapfont.BitmapFont)
		delete(t.labelLoadQueue, labelEID)

		return
	}

	bmf := t.createBitmapFont(entry)
	if bmf == nil {
		return
	}

	// we need to create and cache the bitmap font
	if err := t.bitmapFontCache.Insert(cacheKey, bmf, 1); err != nil {
		t.Warning(err.Error())
	}

	t.AddBitmapFont(labelEID).BitmapFont = bmf
}

func (t *UIWidgetFactory) createBitmapFont(entry labelLoadQueueEntry) *d2bitmapfont.BitmapFont {
	// make sure the components have been loaded (by the asset loader)
	table, tableFound := t.GetFontTable(entry.table)
	sprite, spriteFound := t.GetSprite(entry.sprite)

	if !(tableFound && spriteFound) || sprite.Sprite == nil || table.Data == nil {
		return nil
	}

	return d2bitmapfont.New(sprite.Sprite, table.Data, color.White)
}

func fontCacheKey(t, s, p string) string {
	return fmt.Sprintf("%s::%s::%s", t, s, p)
}

// Button creates a button ui widget
func (t *UIWidgetFactory) Button(x, y float64, imgPath, palPath string) akara.EID {
	buttonEID := t.NewEntity()

	//transform := t.AddTransform(buttonEID)
	//transform.Translation.X, transform.Translation.Y = x, y
	//
	//imgID, palID := t.NewEntity(), t.NewEntity()
	//t.AddFile(imgID).Path = imgPath
	//t.AddFile(palID).Path = palPath
	//
	//t.buttonLoadQueue[buttonEID] = buttonLoadQueueEntry{
	//	spriteImage:   imgID,
	//	spritePalette: palID,
	//}

	return buttonEID
}
