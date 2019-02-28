project(OpenDiablo2.SDL2 VERSION 0.1 LANGUAGES CXX)

set (SOURCES
	src/D2Graphics.cpp
	src/D2Input.cpp
	src/D2Sprite.cpp
)

set (HEADERS
	include/OpenDiablo2.System/D2Graphics.h
	include/OpenDiablo2.System/D2Input.h
)

find_package(SDL2 REQUIRED)

add_library(OpenDiablo2.SDL2 SHARED ${SOURCES} ${HEADERS})

target_include_directories(OpenDiablo2.SDL2
	PUBLIC
		include
	PRIVATE
		../Common
		../OpenDiablo2.Game/include
		include
		sdl2::sdl2
		SDL2::SDL2main
)

target_link_libraries(OpenDiablo2.SDL2
	${SDL2_LIBRARY}
)

set_target_properties(OpenDiablo2.SDL2 PROPERTIES
	CXX_STANDARD 17
)
