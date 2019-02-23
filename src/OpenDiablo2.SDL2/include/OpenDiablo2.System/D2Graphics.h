#ifndef OPENDIABLO2_D2GRAPHICS_H
#define OPENDIABLO2_D2GRAPHICS_H

#include <memory>
#include <SDL2/SDL.h>
#include <spdlog/spdlog.h>

namespace OpenDiablo2 {
	namespace System {

		struct SDLWindowDestroyer
		{
			void operator()(SDL_Window* w) const
			{
				spdlog::debug("Destroying SDL window");
				if (w) SDL_DestroyWindow(w);
			}
		};

		struct SDLRendererDestroyer
		{
			void operator()(SDL_Renderer* r) const
			{
				spdlog::debug("Destroying SDL renderer");
				if (r) SDL_DestroyRenderer(r);
			}
		};

		class D2Graphics {
		public:
			typedef std::unique_ptr<D2Graphics> Ptr;
			D2Graphics();
			void InitializeWindow();
			void Clear();
			void Present();
		private:
			std::unique_ptr<SDL_Window, SDLWindowDestroyer> window;
			std::unique_ptr<SDL_Renderer, SDLRendererDestroyer> renderer;
		};

	}
}

#endif //OPENDIABLO2_D2GRAPHICS_H
