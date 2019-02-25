#include <OpenDiablo2.System/D2Sprite.h>
#include <OpenDiablo2.Game/Common/D2Palette.h>

OpenDiablo2::System::D2Sprite::D2Sprite()
{

}

OpenDiablo2::System::D2Sprite::D2Sprite(
	std::string resourcePath,
	std::string palette,
	bool        cacheFrames)
: Frame          (0)
, TotalFrames    (0)
, Blend          (false)
, Darken         (false)
, Location       (OpenDiablo2::Game::Common::D2Point {0, 0})
, FrameSize      (OpenDiablo2::Game::Common::D2Size  {0, 0})
, LocalFrameSize (OpenDiablo2::Game::Common::D2Size  {0, 0})
, CurrentPalette (palette)
{

}

OpenDiablo2::System::D2Sprite::D2Sprite(
	std::string           resourcePath,
	std::string           palette,
	Game::Common::D2Point location,
	bool                  cacheFrames)
: Frame          (0)
, TotalFrames    (0)
, Blend          (false)
, Darken         (false)
, Location       (location)
, FrameSize      (OpenDiablo2::Game::Common::D2Size {0, 0})
, LocalFrameSize (OpenDiablo2::Game::Common::D2Size {0, 0})
, CurrentPalette (palette)
{

}
