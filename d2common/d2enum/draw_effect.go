package d2enum

type DrawEffect int

// Names courtesy of Necrolis
const (
	DrawEffectPctTransparency25 = 0 //GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA 25 % transparency (colormaps 49-304 in a .pl2)
	DrawEffectPctTransparency50 = 1 //GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA 50 % transparency (colormaps 305-560 in a .pl2)
	DrawEffectPctTransparency75 = 2 //GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA 75 % transparency (colormaps 561-816 in a .pl2)
	DrawEffectModulate          = 3 //GL_MODULATE; GL_SRC_ALPHA, GL_DST_ALPHA (colormaps 817-1072 in a .pl2)
	DrawEffectBurn              = 4 //GL_MODULATE; GL_DST_COLOR, GL_ONE_MINUS_SRC_ALPHA (colormaps 1073-1328 in a .pl2)
	DrawEffectNormal            = 5 //GL_MODULATE; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA (colormaps 1457-1712 in a .pl2)
	DrawEffectMod2XTrans        = 6 //GL_MODULATE; GL_SRC_COLOR, GL_DST_ALPHA (colormaps 1457-1712 in a .pl2)
	DrawEffectMod2X             = 7 //GL_COMBINE_ARB; GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA (colormaps 1457-1712 in a .pl2)
)
