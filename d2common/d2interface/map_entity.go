package d2interface

type MapEntity interface {
	Renderable
	Advanceable
	WorldEntity
	Name() string
	GetLayer() int
	Selectable() bool
	Highlight()
}

type AnimatedEntity interface {
	MapEntity
	Direction()
	SetDirection()
}

type MovingEntity interface {
	AnimatedEntity
	SetVelocity()
}

type Missile interface {
	MovingEntity
}

type ModalAnimatedEntity interface {
	MovingEntity
	SetMode(string)
}

type CompositeModalAnimatedEntity interface {
	ModalAnimatedEntity
}
