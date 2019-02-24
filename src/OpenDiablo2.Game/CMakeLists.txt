add_executable(OpenDiablo2.Game
    src/main.cpp
    src/D2Engine.cpp
)

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

set(CMAKE_CXX_STANDARD 17)
set(CXX_STANDARD 17)

set_property(GLOBAL PROPERTY USE_FOLDERS ON)
