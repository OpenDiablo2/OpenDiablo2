add_executable(OpenDiablo2 src/main.cpp)

include_directories(OpenDiablo2 ../OpenDiablo2.OS/include)
target_link_libraries(OpenDiablo2 OpenDiablo2.OS)
