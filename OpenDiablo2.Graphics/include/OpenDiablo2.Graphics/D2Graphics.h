#ifndef OPENDIABLO2_D2GRAPHICS_H
#define OPENDIABLO2_D2GRAPHICS_H

#include <memory>

namespace OpenDiablo2 {
	namespace Graphics {

		class D2Graphics {
		public:
			D2Graphics();
			void Clear();
		private:
		};

		typedef std::unique_ptr<D2Graphics> D2GraphicsPtr;

	}
}

#endif //OPENDIABLO2_D2GRAPHICS_H
