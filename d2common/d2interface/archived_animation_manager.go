package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

type ArchivedAnimationManager interface {
	Cacher
	LoadAnimation(animationPath, palettePath string, drawEffect d2enum.DrawEffect) (Animation, error)
}
