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
	Components struct {
		Transform d2components.TransformFactory
		Color d2components.ColorFactory
		Rectangle d2components.RectangleFactory
		Texture d2components.TextureFactory
		Size d2components.SizeFactory
		Origin d2components.OriginFactory
	}
	loadQueue      spriteLoadQueue
	shapesToRender *akara.Subscription
	shapesToUpdate *akara.Subscription
}

// Init the sprite factory, injecting the necessary components
func (t *ShapeSystem) Init(world *akara.World) {
	t.World = world

	t.Debug("initializing sprite factory ...")

	t.setupFactories()
	t.setupSubscriptions()

	t.loadQueue = make(spriteLoadQueue)
}

func (t *ShapeSystem) setupFactories() {
	t.InjectComponent(&d2components.Color{}, &t.Components.Color.ComponentFactory)
	t.InjectComponent(&d2components.Transform{}, &t.Components.Transform.ComponentFactory)
	t.InjectComponent(&d2components.Texture{}, &t.Components.Texture.ComponentFactory)
	t.InjectComponent(&d2components.Origin{}, &t.Components.Origin.ComponentFactory)
	t.InjectComponent(&d2components.Size{}, &t.Components.Size.ComponentFactory)
	t.InjectComponent(&d2components.Rectangle{}, &t.Components.Rectangle.ComponentFactory)
}

func (t *ShapeSystem) setupSubscriptions() {
	shapesToRender := t.NewComponentFilter().
		RequireOne(&d2components.Rectangle{}).
		Require(&d2components.Texture{}).
		Build()

	shapesToUpdate := t.NewComponentFilter().
		RequireOne(&d2components.Rectangle{}).
		Require(&d2components.Transform{}, &d2components.Size{}).
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

// ComponentFactory queues a sprite spriteation to be loaded
func (t *ShapeSystem) Rectangle(x, y, width, height int, color color.Color) akara.EID {
	t.Debug("creating rectangle")

	eid := t.NewEntity()
	r := t.Components.Rectangle.Add(eid)

	r.X, r.Y = float64(x), float64(y)
	r.Width, r.Height = float64(width), float64(height)

	c := t.Components.Color.Add(eid)
	c.Color = color

	texture := t.Components.Texture.Add(eid)
	texture.Texture = t.RenderSystem.renderer.NewSurface(width, height)
	texture.Texture.Clear(c.Color)

	return eid
}

func (t *ShapeSystem) updateShape(eid akara.EID) {
	transform, found := t.Components.Transform.Get(eid)
	if !found {
		return
	}

	size, found := t.Components.Size.Get(eid)
	if !found {
		return
	}

	texture, found := t.Components.Texture.Get(eid)
	if !found || texture.Texture == nil {
		return
	}

	rectangle, rectangleFound := t.Components.Rectangle.Get(eid)
	if rectangleFound {
		transform.Translation.X, transform.Translation.Y = rectangle.X, rectangle.Y
		size.X, size.Y = rectangle.Width, rectangle.Height

		tw, th := texture.Texture.GetSize()
		if tw != int(size.X) || th != int(size.Y) {
			texture.Texture.Renderer().NewSurface(int(size.X), int(size.Y))
		}
	}
}

func (t *ShapeSystem) renderShape(eid akara.EID) {
	texture, found := t.Components.Texture.Get(eid)
	if !found || texture.Texture == nil {
		return
	}

	col, found := t.Components.Color.Get(eid)
	if !found {
		return
	}

	texture.Texture.Clear(col.Color)
}
