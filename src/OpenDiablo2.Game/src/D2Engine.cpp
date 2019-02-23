#include <OpenDiablo2.Game/D2Engine.h>


OpenDiablo2::Game::D2Engine::D2Engine(const D2EngineConfig &config)
: config(config) {
	gfx = std::make_unique<OpenDiablo2::System::D2Graphics>();
	input = std::make_unique<OpenDiablo2::System::D2Input>();
}

void
OpenDiablo2::Game::D2Engine::Run() {
	gfx->InitializeWindow();

	while (isRunning) {
		input->ProcessEvents();
		if (input->QuitIsRequested()) {
			isRunning = false;
			break;
		}
		gfx->Clear();

		gfx->Present();
	}
}

