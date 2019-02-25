#ifndef OPENDIABLO2_D2SCENE_H
#define OPENDIABLO2_D2SCENE_H

#include <memory>

namespace OpenDiablo2::Game::Scenes
{

class D2Scene
{
public:
	virtual void Render() = 0;
	virtual void Update() = 0;
};

}

#endif
