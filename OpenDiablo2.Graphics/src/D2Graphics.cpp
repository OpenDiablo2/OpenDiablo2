#include <OpenDiablo2.Graphics/D2Graphics.h>
#include <GL/gl.h>

#include "OpenDiablo2.Graphics/D2Graphics.h"

namespace OpenDiablo2 {
	namespace Graphics {

		D2Graphics::D2Graphics() = default;

		void
		D2Graphics::Clear() {
			glClear(GL_COLOR_BUFFER_BIT);
		}

	}
}

