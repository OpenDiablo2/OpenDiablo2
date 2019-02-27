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
	include/OpenDiablo2.Game/Common/D2Palette.h
	include/OpenDiablo2.Game/Common/D2Point.h
	include/OpenDiablo2.Game/Common/D2ResourcePath.h
	include/OpenDiablo2.Game/Common/D2Size.h
	include/OpenDiablo2.Game/Scenes/D2Scene.h
	include/OpenDiablo2.Game/Scenes/D2MainMenu.h
)

find_package(CLI11 REQUIRED)

add_executable(OpenDiablo2.Game ${SOURCES} ${HEADERS})

include_directories(OpenDiablo2.Game
	PUBLIC
		../Common
		../OpenDiablo2.SDL2/include
		include
)

target_link_libraries(OpenDiablo2.Game
	OpenDiablo2.SDL2
	storm
	CLI11::CLI11
	stdc++fs
)

set_target_properties(OpenDiablo2.Game PROPERTIES
	CXX_STANDARD 17
)
