#include <OpenDiablo2.System/D2Sprite.h>
#include <OpenDiablo2.Game/Common/D2Palette.h>

OpenDiablo2::System::D2Sprite::D2Sprite()
: Frame(-1)
, TotalFrames(-1)
, Blend(false)
, Darken(false)
, Location(OpenDiablo2::Game::Common::D2Point {0, 0})
, FrameSize(OpenDiablo2::Game::Common::D2Size {0, 0})
, LocalFrameSize(OpenDiablo2::Game::Common::D2Size {0, 0})
, CurrentPalette(OpenDiablo2::Game::Common::D2Palette::Static)
{ }

OpenDiablo2::System::D2Sprite *
OpenDiablo2::System::D2Sprite::Load(
	std::string           resourcePath,
	std::string           palette,
	bool                  cacheFrames)
{
	return Load(resourcePath, palette, OpenDiablo2::Game::Common::D2Point {0, 0}, cacheFrames);
}

OpenDiablo2::System::D2Sprite *
OpenDiablo2::System::D2Sprite::Load(
	std::string           resourcePath,
	std::string           palette,
	Game::Common::D2Point location,
	bool                  cacheFrames)
{
	auto result = new D2Sprite();

	return result;
}
