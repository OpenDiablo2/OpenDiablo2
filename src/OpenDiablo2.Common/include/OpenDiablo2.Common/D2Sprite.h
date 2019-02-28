#ifndef OPENDIABLO2_COMMON_D2SPRITE_H
#define OPENDIABLO2_COMMON_D2SPRITE_H

#include <OpenDiablo2.Common/D2Palette.h>
#include <OpenDiablo2.Common/D2Point.h>
#include <OpenDiablo2.Common/D2Size.h>

namespace OpenDiablo2::Common {

class D2Sprite {
public:
	D2Sprite();
	D2Sprite(std::string resourcePath, std::string palette, D2Point location, bool cacheFrames = false);
	D2Sprite(std::string resourcePath, std::string palette, bool cacheFrames = false);

	int Frame;
	int TotalFrames;
	bool Blend;
	bool Darken;
	D2Point Location;
	D2Size FrameSize;
	D2Size LocalFrameSize;
	D2Palette::Entry CurrentPalette;
};

}

#endif // OPENDIABLO2_COMMON_D2SPRITE_H
