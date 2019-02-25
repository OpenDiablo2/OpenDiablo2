project(OpenDiablo2.SDL2 VERSION 0.1 LANGUAGES CXX)

set (SOURCES
	src/D2Graphics.cpp
	src/D2Input.cpp
)

set (HEADERS
	include/OpenDiablo2.System/D2Graphics.h
	include/OpenDiablo2.System/D2Input.h
)

add_library(OpenDiablo2.SDL2 SHARED ${SOURCES} ${HEADERS})

target_include_directories(OpenDiablo2.SDL2
	PUBLIC
		include
	PRIVATE
		include
)

target_link_libraries(OpenDiablo2.SDL2
	spdlog::spdlog
	SDL2-static
)

set_target_properties(OpenDiablo2.SDL2 PROPERTIES
	CXX_STANDARD 17
)
