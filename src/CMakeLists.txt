# Initialize the CMake version stuff
cmake_minimum_required(VERSION 3.1...3.13)
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${project_SOURCE_DIR}/cmake")

if(${CMAKE_VERSION} VERSION_LESS 3.12)
    cmake_policy(VERSION ${CMAKE_MAJOR_VERSION}.${CMAKE_MINOR_VERSION})
endif()

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/output)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/output)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/output)

# Add our nifty CMake scripts ------------------------------------------------------------------------------------------
list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}/cmake")

# Prevent users from shooting themselves in the foot -------------------------------------------------------------------
if(CMAKE_SOURCE_DIR STREQUAL CMAKE_BINARY_DIR)
    message(FATAL_ERROR
            "Do not build in-source.
             Please remove CMakeCache.txt and the CMakeFiles/ directory. Then build out-of-source.")
endif()

# ----------------------------------------------------------------------------------------------------------------------
project(OpenDiablo2 VERSION 0.1
        DESCRIPTION "An open source Diablo2 engine.")
# ----------------------------------------------------------------------------------------------------------------------

enable_language(CXX)

# Freaky relocatable exe stuff -----------------------------------------------------------------------------------------
set(CMAKE_POSITION_INDEPENDENT_CODE ON CACHE BOOL "")

set_property(GLOBAL PROPERTY USE_FOLDERS ON)

# Import the actual projects -------------------------------------------------------------------------------------------
add_subdirectory(OpenDiablo2.Common)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/src/OpenDiablo2.Common/include)

add_subdirectory(OpenDiablo2.SDL2)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/src/OpenDiablo2.SDL2/include)

add_subdirectory(OpenDiablo2.Game)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/src/OpenDiablo2.Game/include)

