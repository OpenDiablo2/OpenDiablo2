project(OpenDiablo2.SDL2 VERSION 0.1 LANGUAGES CXX)

add_library(OpenDiablo2.SDL2 STATIC
    src/D2Graphics.cpp
    src/D2Input.cpp
)

target_include_directories(OpenDiablo2.SDL2
    PUBLIC
        include
    PRIVATE
        include
)

target_link_libraries(OpenDiablo2.SDL2
        spdlog::spdlog
        SDL2
)
