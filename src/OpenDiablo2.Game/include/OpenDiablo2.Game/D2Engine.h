#ifndef OPENDIABLO2_D2ENGINE_H
#define OPENDIABLO2_D2ENGINE_H

#include <OpenDiablo2.System/D2Graphics.h>
#include <OpenDiablo2.System/D2Input.h>
#include "D2EngineConfig.h"

namespace OpenDiablo2 {
	namespace Game {
		class D2Engine {
		public:
			D2Engine(const D2EngineConfig& config);
			void Run();
		private:
			const D2EngineConfig config;
			OpenDiablo2::System::D2Graphics::Ptr gfx;
			OpenDiablo2::System::D2Input::Ptr input;
			bool isRunning = true;

		};
	}
}


#endif //OPENDIABLO2_D2ENGINE_H

