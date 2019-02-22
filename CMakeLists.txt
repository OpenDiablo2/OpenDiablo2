cmake_minimum_required(VERSION 3.0)

if(CMAKE_SOURCE_DIR STREQUAL CMAKE_BINARY_DIR)
    message(FATAL_ERROR "Do not build in-source. Please remove CMakeCache.txt and the CMakeFiles/ directory. Then build out-of-source.")
endif()

project(OpenDiablo2 CXX)

include(CTest)
include(CMakeDependentOption)

set(CMAKE_CXX_STANDARD 17)
set(CXX_STANDARD 17)
set(OpenGL_GL_PREFERENCE "GLVND")

set_property(GLOBAL PROPERTY USE_FOLDERS ON)

add_subdirectory(3rdparty)

add_subdirectory(src/OpenDiablo2.OS)
add_subdirectory(src/OpenDiablo2.Graphics)
add_subdirectory(src/OpenDiablo2.Game)
