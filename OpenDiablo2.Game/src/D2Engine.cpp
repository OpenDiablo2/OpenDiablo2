#include <OpenDiablo2.Game/D2Engine.h>


OpenDiablo2::Game::D2Engine::D2Engine() {
	window = std::make_unique<OpenDiablo2::OS::D2Window>();
	gfx = std::make_unique<OpenDiablo2::Graphics::D2Graphics>();
}

void
OpenDiablo2::Game::D2Engine::Run() {
	window->Initialize();

	while (window->WindowStillOpen()) {
		window->PollEvents();
		gfx->Clear();

		window->FlipBuffer();
	}

	window->Finalize();
}
