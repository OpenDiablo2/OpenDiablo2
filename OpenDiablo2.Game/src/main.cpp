#include <memory>
#include <D2Window.h>

int
main() {
	auto window = std::make_unique<OpenDiablo2::OS::D2Window>();
	window->Run();
	return 0;
}


