#include <OpenDiablo2.Common/D2Sprite.h>
#include <OpenDiablo2.Common/D2Palette.h>

OpenDiablo2::Common::D2Sprite::D2Sprite()
{

}

OpenDiablo2::Common::D2Sprite::D2Sprite(
	std::string resourcePath,
	std::string palette,
	bool        cacheFrames)
: Frame          (0)
, TotalFrames    (0)
, Blend          (false)
, Darken         (false)
, Location       (D2Point {0, 0})
, FrameSize      (D2Size  {0, 0})
, LocalFrameSize (D2Size  {0, 0})
, CurrentPalette (palette)
{

}

OpenDiablo2::Common::D2Sprite::D2Sprite(
	std::string resourcePath,
	std::string palette,
	D2Point     location,
	bool                  cacheFrames)
: Frame          (0)
, TotalFrames    (0)
, Blend          (false)
, Darken         (false)
, Location       (location)
, FrameSize      (D2Size {0, 0})
, LocalFrameSize (D2Size {0, 0})
, CurrentPalette (palette)
{

}
