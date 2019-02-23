project(OpenDiablo2.SDL2 VERSION 0.1 LANGUAGES CXX)

add_library(OpenDiablo2.SDL2 SHARED
    src/D2Graphics.cpp
    src/D2Input.cpp
)

find_package(SDL2 REQUIRED)
find_package(SDL2_image REQUIRED)

target_include_directories(OpenDiablo2.SDL2
    PUBLIC
        include
    PRIVATE
        include
        ../../3rdparty/bgfx/include
        ../../3rdparty/spdlog/include
        ${SDL2_INCLUDE_DIR}
        ${SDL2_IMAGE_INCLUDE_DIR}
)

target_link_libraries(OpenDiablo2.SDL2
    PUBLIC
        ${SDL2_LIBRARY}
        ${SDL2_IMAGE_LIBRARIES}
)

if(MSVC)
    target_compile_options(OpenDiablo2.SDL2 PRIVATE /W4 /WX)
else(MSVC)
    target_compile_options(OpenDiablo2.SDL2 PRIVATE -Wall -Wextra -pedantic -Werror)
endif(MSVC)
