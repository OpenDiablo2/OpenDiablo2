#ifndef OPENDIABLO2_D2ENGINE_H
#define OPENDIABLO2_D2ENGINE_H

#include <OpenDiablo2.OS/D2Window.h>
#include <OpenDiablo2.Graphics/D2Graphics.h>

namespace OpenDiablo2 {
	namespace Game {
		class D2Engine {
		public:
			D2Engine();
			void Run();
		private:
			OpenDiablo2::OS::D2WindowPtr window;
			OpenDiablo2::Graphics::D2GraphicsPtr gfx;
		};
	}
}


#endif //OPENDIABLO2_D2ENGINE_H
