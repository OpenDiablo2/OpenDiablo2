project(libOpenDiablo2.OS VERSION 0.1 LANGUAGES CXX)

add_library(OpenDiablo2.OS SHARED
    src/D2Window.cpp
)

target_include_directories(OpenDiablo2.OS
    PUBLIC
        include
    PRIVATE
        include
        ../../3rdparty/glfw/include
        ../../3rdparty/spdlog/include
)

target_link_libraries(OpenDiablo2.OS
    PRIVATE
        glfw
)

if(MSVC)
    target_compile_options(OpenDiablo2.OS PRIVATE /W4 /WX)
else(MSVC)
    target_compile_options(OpenDiablo2.OS PRIVATE -Wall -Wextra -pedantic -Werror)
endif(MSVC)
