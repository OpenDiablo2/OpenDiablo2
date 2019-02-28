project(OpenDiablo2.Common VERSION 0.1 LANGUAGES CXX)

set (SOURCES
    src/D2Sprite.cpp
    src/D2DataManager.cpp
)

set (HEADERS
    include/OpenDiablo2.Common/D2EngineConfig.h
    include/OpenDiablo2.Common/D2DataManager.h
	include/OpenDiablo2.Common/D2Palette.h
	include/OpenDiablo2.Common/D2Point.h
    include/OpenDiablo2.Common/D2ResourcePath.h
    include/OpenDiablo2.Common/D2Size.h
    include/OpenDiablo2.Common/D2Sprite.h
)

add_library(OpenDiablo2.Common SHARED ${SOURCES} ${HEADERS})

target_include_directories(OpenDiablo2.Common
	PUBLIC
		include
	PRIVATE
        include
        ../ExtraUtils
)


set_target_properties(OpenDiablo2.Common PROPERTIES
	CXX_STANDARD 17
)
