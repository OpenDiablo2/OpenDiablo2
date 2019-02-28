#ifndef OPENDIABLO2_D2ENGINE_H
#define OPENDIABLO2_D2ENGINE_H

#include <stack>
#include <memory>
#include <OpenDiablo2.Common/D2DataManager.h>
#include <OpenDiablo2.System/D2Graphics.h>
#include <OpenDiablo2.System/D2Input.h>
#include <OpenDiablo2.Game/Scenes/D2Scene.h>
#include "OpenDiablo2.Common/D2EngineConfig.h"

namespace OpenDiablo2::Game
{

// The main OpenDiablo2 engine
class D2Engine : public std::enable_shared_from_this<D2Engine>
{
public:
	D2Engine(const Common::D2EngineConfig &config);

	// Runs the engine
	void Run();

	OpenDiablo2::Common::D2DataManager::Ptr dataManager;
	OpenDiablo2::System::D2Graphics::Ptr gfx;
	OpenDiablo2::System::D2Input::Ptr input;

private:
	// Represents the engine configuration
	const Common::D2EngineConfig config;

	// Indicates the system should keep running (if set to true)
	bool isRunning = true;

	std::stack<std::shared_ptr<OpenDiablo2::Game::Scenes::D2Scene>> sceneStack;
};

}

#endif //OPENDIABLO2_D2ENGINE_H
