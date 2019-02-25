
#include <OpenDiablo2.System/D2Input.h>
#include <SDL2/SDL.h>

OpenDiablo2::System::D2Input::D2Input()
{
	SDL_Init(SDL_INIT_EVENTS);
}

void
OpenDiablo2::System::D2Input::ProcessEvents()
{
	SDL_Event event;
	while (SDL_PollEvent(&event)) {
		if( event.type == SDL_QUIT ) {
			quitIsRequested = true;
		}
	}
}

bool
OpenDiablo2::System::D2Input::QuitIsRequested()
{
	return quitIsRequested;
}
