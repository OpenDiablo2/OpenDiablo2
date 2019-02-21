project(libOpenDiablo2.OS VERSION 0.1 LANGUAGES CXX)

include_directories(
    include
    ../3rdparty/glfw/include
    ../3rdparty/spdlog/include
)

add_library(OpenDiablo2.OS SHARED
    src/D2Window.cpp
)

target_link_libraries(OpenDiablo2.OS glfw)

if(MSVC)
    target_compile_options(OpenDiablo2.OS PRIVATE /W4 /WX)
else(MSVC)
    target_compile_options(OpenDiablo2.OS PRIVATE -Wall -Wextra -pedantic -Werror)
endif(MSVC)
