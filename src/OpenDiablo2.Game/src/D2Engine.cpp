#include <OpenDiablo2.Game/D2Engine.h>
#include <OpenDiablo2.Game/Scenes/D2MainMenu.h>


OpenDiablo2::Game::D2Engine::D2Engine(const D2EngineConfig &config)
: config(config) {
	gfx = std::make_unique<OpenDiablo2::System::D2Graphics>();
	input = std::make_unique<OpenDiablo2::System::D2Input>();
}

void
OpenDiablo2::Game::D2Engine::Run() {
	gfx->InitializeWindow();

	sceneStack.emplace(std::make_shared<Scenes::MainMenu>(shared_from_this()));

	while (isRunning) {
		input->ProcessEvents();
		sceneStack.top()->Update();
		if (input->QuitIsRequested()) {
			isRunning = false;
			break;
		}
		gfx->Clear();
		sceneStack.top()->Render();
		gfx->Present();
	}
}

