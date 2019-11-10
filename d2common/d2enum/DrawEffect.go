package d2enum

type DrawEffect int

const (
	DrawEffectPctTransparency75  = 0 //75 % transparency (colormaps 561-816 in a .pl2)
	DrawEffectPctTransparency50  = 1 //50 % transparency (colormaps 305-560 in a .pl2)
	DrawEffectPctTransparency25  = 2 //25 % transparency (colormaps 49-304 in a .pl2)
	DrawEffectScreen             = 3 //Screen (colormaps 817-1072 in a .pl2)
	DrawEffectLuminance          = 4 //luminance (colormaps 1073-1328 in a .pl2)
	DrawEffectBringAlphaBlending = 5 //bright alpha blending (colormaps 1457-1712 in a .pl2)
)
