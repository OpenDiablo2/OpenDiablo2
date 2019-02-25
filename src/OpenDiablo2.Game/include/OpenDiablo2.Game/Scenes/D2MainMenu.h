#ifndef OPENDIALOB2_SCENES_D2MAINMENU_H
#define OPENDIALOB2_SCENES_D2MAINMENU_H

#include <OpenDiablo2.Game/Scenes/D2Scene.h>
#include <OpenDiablo2.Game/D2Engine.h>

namespace OpenDiablo2::Game::Scenes {

class MainMenu : public D2Scene {
public:
	MainMenu(std::shared_ptr<D2Engine> engine);
	void Render();
	void Update();
private:
	std::shared_ptr<D2Engine> engine;
};

}
#endif // OPENDIALOB2_SCENES_D2MAINMENU_H
