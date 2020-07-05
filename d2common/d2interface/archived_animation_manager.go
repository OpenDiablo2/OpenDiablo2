package d2interface

type ArchivedAnimationManager interface {
	Cacher
	AssetManagerSubordinate
	LoadAnimation(animationPath, palettePath string, transparency int) (Animation, error)
}
