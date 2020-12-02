package d2systems

import (
	"image/color"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

// NewShapeSystem creates a new sprite factory which is intended
// to be embedded in the game object factory system.
func NewShapeSystem(b akara.BaseSystem, l *d2util.Logger) *ShapeSystem {
	sys := &ShapeSystem{
		Logger: l,
	}

	sys.BaseSystem = b

	sys.World.AddSystem(sys)

	return sys
}

// ShapeSystem is responsible for queueing sprites to be loaded (as spriteations),
// as well as binding the spriteation to a renderer if one is present (which generates the sprite surfaces).
type ShapeSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	RenderSystem *RenderSystem
	d2components.PositionFactory
	d2components.ColorFactory
	d2components.RectangleFactory
	d2components.TextureFactory
	d2components.SizeFactory
	d2components.OriginFactory
	loadQueue      spriteLoadQueue
	shapesToRender *akara.Subscription
	shapesToUpdate *akara.Subscription
}

// Init the sprite factory, injecting the necessary components
func (t *ShapeSystem) Init(world *akara.World) {
	t.World = world

	t.Info("initializing sprite factory ...")

	t.setupFactories()
	t.setupSubscriptions()

	t.loadQueue = make(spriteLoadQueue)
}

func (t *ShapeSystem) setupFactories() {
	t.InjectComponent(&d2components.Color{}, &t.ColorFactory.Color)
	t.InjectComponent(&d2components.Position{}, &t.PositionFactory.Position)
	t.InjectComponent(&d2components.Texture{}, &t.TextureFactory.Texture)
	t.InjectComponent(&d2components.Origin{}, &t.OriginFactory.Origin)
	t.InjectComponent(&d2components.Size{}, &t.SizeFactory.Size)
	t.InjectComponent(&d2components.Rectangle{}, &t.RectangleFactory.Rectangle)
}

func (t *ShapeSystem) setupSubscriptions() {
	shapesToRender := t.NewComponentFilter().
		RequireOne(&d2components.Rectangle{}).
		Require(&d2components.Texture{}).
		Build()

	shapesToUpdate := t.NewComponentFilter().
		RequireOne(&d2components.Rectangle{}).
		Require(&d2components.Position{}, &d2components.Size{}).
		Build()

	t.shapesToRender = t.AddSubscription(shapesToRender)
	t.shapesToUpdate = t.AddSubscription(shapesToUpdate)
}

// Update processes the load queue which attempting to create spriteations, as well as
// binding existing spriteations to a renderer if one is present.
func (t *ShapeSystem) Update() {
	for _, id := range t.shapesToUpdate.GetEntities() {
		t.updateShape(id)
	}

	for _, id := range t.shapesToRender.GetEntities() {
		t.renderShape(id)
	}
}

// Sprite queues a sprite spriteation to be loaded
func (t *ShapeSystem) Rectangle(x, y, width, height int, color color.Color) akara.EID {
	t.Info("creating rectangle")

	eid := t.NewEntity()
	r := t.AddRectangle(eid)

	r.X, r.Y = float64(x), float64(y)
	r.Width, r.Height = float64(width), float64(height)

	c := t.AddColor(eid)
	c.Color = color

	texture := t.AddTexture(eid)
	texture.Texture = t.RenderSystem.renderer.NewSurface(width, height)
	texture.Texture.Clear(c.Color)

	return eid
}

func (t *ShapeSystem) updateShape(eid akara.EID) {
	position, found := t.GetPosition(eid)
	if !found {
		return
	}

	size, found := t.GetSize(eid)
	if !found {
		return
	}

	texture, found := t.GetTexture(eid)
	if !found || texture.Texture == nil {
		return
	}

	rectangle, rectangleFound := t.GetRectangle(eid)
	if rectangleFound {
		position.X, position.Y = rectangle.X, rectangle.Y
		size.X, size.Y = rectangle.Width, rectangle.Height

		tw, th := texture.Texture.GetSize()
		if tw != int(size.X) || th != int(size.Y) {
			texture.Texture.Renderer().NewSurface(int(size.X), int(size.Y))
		}
	}
}

func (t *ShapeSystem) renderShape(eid akara.EID) {
	texture, found := t.GetTexture(eid)
	if !found || texture.Texture == nil {
		return
	}

	col, found := t.GetColor(eid)
	if !found {
		return
	}

	texture.Texture.Clear(col.Color)
}
