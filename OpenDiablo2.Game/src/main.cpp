#include <memory>
#include <spdlog/spdlog.h>
#include <spdlog/sinks/stdout_color_sinks.h>
#include <OpenDiablo2.Game/D2Engine.h>

int
main() {
	spdlog::set_level(spdlog::level::trace);
	spdlog::set_pattern("[%^%l%$] %v");

	spdlog::info("OpenDiablo 2 has started");

	auto engine = std::make_unique<OpenDiablo2::Game::D2Engine>();
	engine->Run();

	return 0;
}


