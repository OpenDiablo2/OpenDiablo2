
add_executable(OpenDiablo2
    src/main.cpp
    src/D2Engine.cpp
)

include_directories(OpenDiablo2
    include
    ../OpenDiablo2.OS/include
    ../OpenDiablo2.Graphics/include
    ../3rdparty/spdlog/include
)
target_link_libraries(OpenDiablo2
    PUBLIC
        OpenDiablo2.OS
        OpenDiablo2.Graphics
)

target_compile_features(OpenDiablo2 PRIVATE cxx_std_17)

if(MSVC)
    target_compile_options(OpenDiablo2 PRIVATE /W4 /WX)
else(MSVC)
    target_compile_options(OpenDiablo2 PRIVATE -Wall -Wextra -pedantic -Werror)
endif(MSVC)
