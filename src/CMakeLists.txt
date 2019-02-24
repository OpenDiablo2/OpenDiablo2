# Initialize the CMake version stuff
cmake_minimum_required(VERSION 3.1...3.13)

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

# Pull down submodules -------------------------------------------------------------------------------------------------
find_package(Git QUIET)
if(GIT_FOUND AND EXISTS "${PROJECT_SOURCE_DIR}/.git")
    # Update submodules as needed
    option(GIT_SUBMODULE "Check submodules during build" ON)
    if(GIT_SUBMODULE)
        message(STATUS "Submodule update")
        execute_process(COMMAND ${GIT_EXECUTABLE} submodule update --init --recursive
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                RESULT_VARIABLE GIT_SUBMOD_RESULT)
        if(NOT GIT_SUBMOD_RESULT EQUAL "0")
            message(FATAL_ERROR "git submodule update --init failed with ${GIT_SUBMOD_RESULT}, please checkout submodules")
        endif()
    endif()
endif()

# Freaky relocatable exe stuff -----------------------------------------------------------------------------------------
set(CMAKE_POSITION_INDEPENDENT_CODE ON CACHE BOOL "")

# Import the libraries -------------------------------------------------------------------------------------------------
set( BGFX_BUILD_TOOLS      OFF )
set( BGFX_BUILD_EXAMPLES   OFF )
set( BGFX_INSTALL          OFF )
set( BGFX_INSTALL_EXAMPLES OFF )
set( BGFX_CUSTOM_TARGETS   OFF )
set( BGFX_USE_OVR          OFF )
set( BGFX_AMALGAMATED      ON  )
set( BX_AMALGAMATED        ON  )
set( BGFX_CONFIG_DEBUG     OFF )
set( SDL_STATIC_PIC        ON  )
set( SDL_SHARED            OFF )
set( BUILD_SHARED_LIBS     OFF )

add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../3rdparty/spdlog   spdlog)
add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../3rdparty/cli11    cli11)
add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../3rdparty/sdl2     sdl2)
add_subdirectory(${CMAKE_CURRENT_SOURCE_DIR}/../3rdparty/stormlib stormlib)

# Import the actual projects -------------------------------------------------------------------------------------------
add_subdirectory(OpenDiablo2.SDL2)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/src/OpenDiablo2.SDL2/include)

add_subdirectory(OpenDiablo2.Game)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/src/OpenDiablo2.Game/include)


