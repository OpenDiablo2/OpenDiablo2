#include <spdlog/spdlog.h>
#include <OpenDiablo2.System/D2Graphics.h>
#include <string>
#include <SDL2/SDL.h>


#include "OpenDiablo2.System/D2Graphics.h"

OpenDiablo2::System::D2Graphics::D2Graphics()
{
	atexit(SDL_Quit);

	if (SDL_Init(SDL_INIT_VIDEO) < 0) {
		spdlog::error("Could not initialize sdl2: " + std::string(SDL_GetError()));
		exit(1);
	}
}

void
OpenDiablo2::System::D2Graphics::Clear()
{
	SDL_RenderClear(renderer.get());
}

void
OpenDiablo2::System::D2Graphics::InitializeWindow()
{
	spdlog::debug("Initializing SDL window");
	window = std::unique_ptr<SDL_Window, SDLWindowDestroyer>(SDL_CreateWindow("OpenDiablo 2", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED,
		800, 600, SDL_WINDOW_SHOWN));
	if (window == nullptr) {
		spdlog::error("Could not create sdl2 window: " + std::string(SDL_GetError()));
		SDL_Quit();
		exit(1);
	}

	spdlog::debug("Initializing SDL renderer");
	renderer = std::unique_ptr<SDL_Renderer, SDLRendererDestroyer>(SDL_CreateRenderer(window.get(), -1, SDL_RENDERER_ACCELERATED | SDL_RENDERER_PRESENTVSYNC));
	if (renderer == nullptr){
		spdlog::error("Could not create sdl2 window: " + std::string(SDL_GetError()));
		SDL_Quit();
		exit(1);
	}
}

void
OpenDiablo2::System::D2Graphics::Present()
{
	SDL_RenderPresent(renderer.get());
}

