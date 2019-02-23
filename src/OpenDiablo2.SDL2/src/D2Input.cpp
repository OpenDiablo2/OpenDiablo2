
#include <OpenDiablo2.System/D2Input.h>
#include <SDL2/SDL.h>

namespace OpenDiablo2 {
	namespace System {

		D2Input::D2Input() {
			SDL_Init(SDL_INIT_EVENTS);
		}

		void
		D2Input::ProcessEvents() {
			SDL_Event event;
			while (SDL_PollEvent(&event)) {
				if( event.type == SDL_QUIT ) {
					quitIsRequested = true;
				}
			}
		}

		bool
		D2Input::QuitIsRequested() {
			return quitIsRequested;
		}
	}
}
