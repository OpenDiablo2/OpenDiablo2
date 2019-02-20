project(libOpenDiablo2.OS VERSION 0.1 LANGUAGES CXX)

set(GLFW_BUILD_DOCS OFF CACHE BOOL "" FORCE)
set(GLFW_BUILD_TESTS OFF CACHE BOOL "" FORCE)
set(GLFW_BUILD_EXAMPLES OFF CACHE BOOL "" FORCE)
add_subdirectory(../3rdparty/glfw build)

include_directories(
    include
    ../3rdparty/glfw/include
)

add_library(OpenDiablo2.OS
    src/D2Window.cpp
)

target_link_libraries(OpenDiablo2.OS
    glfw
)
