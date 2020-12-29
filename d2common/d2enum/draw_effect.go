package d2enum

// DrawEffect is a draw effect
type DrawEffect int

// Names courtesy of Necrolis
const (
	// DrawEffectPctTransparency25 is a draw effect that implements the following function:
	// GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA 25 % transparency (colormaps 49-304 in a .pl2)
	DrawEffectPctTransparency25 DrawEffect = iota

	// DrawEffectPctTransparency50 is a draw effect that implements the following function:
	// GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA 50 % transparency (colormaps 305-560 in a .pl2)
	DrawEffectPctTransparency50

	// DrawEffectPctTransparency75 is a draw effect that implements the following function:
	// GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA 75 % transparency (colormaps 561-816 in a .pl2)
	DrawEffectPctTransparency75

	// DrawEffectModulate is a draw effect that implements the following function:
	// GL_MODULATE; GL_SRC_ALPHA, GL_DST_ALPHA (colormaps 817-1072 in a .pl2)
	DrawEffectModulate

	// DrawEffectBurn is a draw effect that implements the following function:
	// GL_MODULATE; GL_DST_COLOR, GL_ONE_MINUS_SRC_ALPHA (colormaps 1073-1328 in a .pl2)
	DrawEffectBurn

	// DrawEffectNormal is a draw effect that implements the following function:
	// GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA (colormaps 1457-1712 in a .pl2)
	DrawEffectNormal

	// DrawEffectMod2XTrans is a draw effect that implements the following function:
	// GL_MODULATE; GL_SRC_COLOR, GL_DST_ALPHA (colormaps 1457-1712 in a .pl2)
	DrawEffectMod2XTrans

	// DrawEffectMod2X is a draw effect that implements the following function:
	// GL_COMBINE_ARB; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA (colormaps 1457-1712 in a .pl2)
	DrawEffectMod2X

	// no effect
	DrawEffectNone
)

// Transparent returns true if there is no effect, false otherwise
func (d DrawEffect) Transparent() bool {
	return d != DrawEffectNone
}
