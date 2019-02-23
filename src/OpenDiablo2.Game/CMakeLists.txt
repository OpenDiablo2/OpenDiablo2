
add_executable(OpenDiablo2
    src/main.cpp
    src/D2Engine.cpp
)

target_link_libraries(OpenDiablo2
    PRIVATE
        OpenDiablo2.SDL2
        storm
        CLI11
        stdc++fs
)

target_include_directories(OpenDiablo2
    PUBLIC
        include
    PRIVATE
        include
        ../OpenDiablo2.OS/include
        ../OpenDiablo2.SDL2/include
        ../../3rdparty/spdlog/include
        ../../3rdparty/stormlib/src
        ../../3rdparty/cli11/include
)

if(MSVC)
    target_compile_options(OpenDiablo2 PRIVATE /W4 /WX)
else(MSVC)
    target_compile_options(OpenDiablo2 PRIVATE -Wall -Wextra -pedantic -Werror)
endif(MSVC)
