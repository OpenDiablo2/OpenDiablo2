# Locate SDL_image library
#
# This module defines:
#
# ::
#
#   SDL2_IMAGE_LIBRARIES, the name of the library to link against
#   SDL2_IMAGE_INCLUDE_DIRS, where to find the headers
#   SDL2_IMAGE_FOUND, if false, do not try to link against
#   SDL2_IMAGE_VERSION_STRING - human-readable string containing the version of SDL_image
#
#
#
# For backward compatibility the following variables are also set:
#
# ::
#
#   SDLIMAGE_LIBRARY (same value as SDL2_IMAGE_LIBRARIES)
#   SDLIMAGE_INCLUDE_DIR (same value as SDL2_IMAGE_INCLUDE_DIRS)
#   SDLIMAGE_FOUND (same value as SDL2_IMAGE_FOUND)
#
#
#
# $SDLDIR is an environment variable that would correspond to the
# ./configure --prefix=$SDLDIR used in building SDL.
#
# Created by Eric Wing.  This was influenced by the FindSDL.cmake
# module, but with modifications to recognize OS X frameworks and
# additional Unix paths (FreeBSD, etc).

#=============================================================================
# Copyright 2005-2009 Kitware, Inc.
# Copyright 2012 Benjamin Eikel
#
# Distributed under the OSI-approved BSD License (the "License");
# see accompanying file Copyright.txt for details.
#
# This software is distributed WITHOUT ANY WARRANTY; without even the
# implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
# See the License for more information.
#=============================================================================
# (To distribute this file outside of CMake, substitute the full
#  License text for the above reference.)

find_path(SDL2_IMAGE_INCLUDE_DIR SDL_image.h
        HINTS
        ENV SDL2IMAGEDIR
        ENV SDL2DIR
        PATH_SUFFIXES SDL2
        # path suffixes to search inside ENV{SDLDIR}
        include/SDL2 include
        PATHS ${SDL2_IMAGE_PATH}
        )

if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(VC_LIB_PATH_SUFFIX lib/x64)
else()
    set(VC_LIB_PATH_SUFFIX lib/x86)
endif()

find_library(SDL2_IMAGE_LIBRARY
        NAMES SDL2_image
        HINTS
        ENV SDL2IMAGEDIR
        ENV SDL2DIR
        PATH_SUFFIXES lib ${VC_LIB_PATH_SUFFIX}
        PATHS ${SDL2_IMAGE_PATH}
        )

if(SDL2_IMAGE_INCLUDE_DIR AND EXISTS "${SDL2_IMAGE_INCLUDE_DIR}/SDL_image.h")
    file(STRINGS "${SDL2_IMAGE_INCLUDE_DIR}/SDL_image.h" SDL2_IMAGE_VERSION_MAJOR_LINE REGEX "^#define[ \t]+SDL_IMAGE_MAJOR_VERSION[ \t]+[0-9]+$")
    file(STRINGS "${SDL2_IMAGE_INCLUDE_DIR}/SDL_image.h" SDL2_IMAGE_VERSION_MINOR_LINE REGEX "^#define[ \t]+SDL_IMAGE_MINOR_VERSION[ \t]+[0-9]+$")
    file(STRINGS "${SDL2_IMAGE_INCLUDE_DIR}/SDL_image.h" SDL2_IMAGE_VERSION_PATCH_LINE REGEX "^#define[ \t]+SDL_IMAGE_PATCHLEVEL[ \t]+[0-9]+$")
    string(REGEX REPLACE "^#define[ \t]+SDL_IMAGE_MAJOR_VERSION[ \t]+([0-9]+)$" "\\1" SDL2_IMAGE_VERSION_MAJOR "${SDL2_IMAGE_VERSION_MAJOR_LINE}")
    string(REGEX REPLACE "^#define[ \t]+SDL_IMAGE_MINOR_VERSION[ \t]+([0-9]+)$" "\\1" SDL2_IMAGE_VERSION_MINOR "${SDL2_IMAGE_VERSION_MINOR_LINE}")
    string(REGEX REPLACE "^#define[ \t]+SDL_IMAGE_PATCHLEVEL[ \t]+([0-9]+)$" "\\1" SDL2_IMAGE_VERSION_PATCH "${SDL2_IMAGE_VERSION_PATCH_LINE}")
    set(SDL2_IMAGE_VERSION_STRING ${SDL2_IMAGE_VERSION_MAJOR}.${SDL2_IMAGE_VERSION_MINOR}.${SDL2_IMAGE_VERSION_PATCH})
    unset(SDL2_IMAGE_VERSION_MAJOR_LINE)
    unset(SDL2_IMAGE_VERSION_MINOR_LINE)
    unset(SDL2_IMAGE_VERSION_PATCH_LINE)
    unset(SDL2_IMAGE_VERSION_MAJOR)
    unset(SDL2_IMAGE_VERSION_MINOR)
    unset(SDL2_IMAGE_VERSION_PATCH)
endif()

set(SDL2_IMAGE_LIBRARIES ${SDL2_IMAGE_LIBRARY})
set(SDL2_IMAGE_INCLUDE_DIRS ${SDL2_IMAGE_INCLUDE_DIR})

include(FindPackageHandleStandardArgs)

FIND_PACKAGE_HANDLE_STANDARD_ARGS(SDL2_image
        REQUIRED_VARS SDL2_IMAGE_LIBRARIES SDL2_IMAGE_INCLUDE_DIRS
        VERSION_VAR SDL2_IMAGE_VERSION_STRING)

# for backward compatibility
set(SDLIMAGE_LIBRARY ${SDL2_IMAGE_LIBRARIES})
set(SDLIMAGE_INCLUDE_DIR ${SDL2_IMAGE_INCLUDE_DIRS})
set(SDLIMAGE_FOUND ${SDL2_IMAGE_FOUND})

mark_as_advanced(SDL2_IMAGE_LIBRARY SDL2_IMAGE_INCLUDE_DIR)
