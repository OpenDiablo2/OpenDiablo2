project(libOpenDiablo2.Graphics VERSION 0.1 LANGUAGES CXX)

find_package(OpenGL)

include_directories(
    include
    ../3rdparty/bgfx/include
    ../3rdparty/spdlog/include
    ${OPENGL_INCLUDE_DIRS}
)

add_library(OpenDiablo2.Graphics SHARED
    src/D2Graphics.cpp
)


target_link_libraries(OpenDiablo2.Graphics ${OPENGL_LIBRARIES})

if(MSVC)
    target_compile_options(OpenDiablo2.Graphics PRIVATE /W4 /WX)
else(MSVC)
    target_compile_options(OpenDiablo2.Graphics PRIVATE -Wall -Wextra -pedantic -Werror)
endif(MSVC)
