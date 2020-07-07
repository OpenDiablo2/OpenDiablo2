package d2interface

type ArchivedAnimationManager interface {
	Cacher
	LoadAnimation(animationPath, palettePath string, transparency int) (Animation, error)
}
