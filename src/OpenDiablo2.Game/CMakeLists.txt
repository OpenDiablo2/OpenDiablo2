set (SOURCES
	src/main.cpp
	src/D2DataManager.cpp
	src/D2Engine.cpp
	src/Scenes/D2MainMenu.cpp
)

set (HEADERS
	include/OpenDiablo2.Game/D2DataManager.h
	include/OpenDiablo2.Game/D2Engine.h
	include/OpenDiablo2.Game/D2EngineConfig.h
	include/OpenDiablo2.Game/Scenes/D2Scene.h
	include/OpenDiablo2.Game/Scenes/D2MainMenu.h
)

add_executable(OpenDiablo2.Game ${SOURCES} ${HEADERS})

target_include_directories(OpenDiablo2.SDL2
	PUBLIC
		include
	PRIVATE
		include
)

target_link_libraries(OpenDiablo2.Game
	OpenDiablo2.SDL2
	storm
	CLI11::CLI11
	stdc++fs
	spdlog::spdlog
)

set_target_properties(OpenDiablo2.Game PROPERTIES
	CXX_STANDARD 17
)
