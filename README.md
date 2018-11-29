# OpenDiablo2 [![Build status](https://ci.appveyor.com/api/projects/status/jx83as0ku784r5vh?svg=true)](https://ci.appveyor.com/project/essial/opendiablo2)

An open source re-implementation of Diablo 2 in C# 

[Join us on Discord!](https://discord.gg/pRy8tdc)\
[Development Live stream](https://www.twitch.tv/essial/)

<img src="https://raw.githubusercontent.com/essial/OpenDiablo2/master/Screenshot.png" />
<img src="https://raw.githubusercontent.com/essial/OpenDiablo2/master/Screenshot2.png" />

## About this project

This is an attempt to re-create Diablo 2's game engine in C#, and make it cross platform as well. This project does not ship with the assets or content required to work. You must have a legally purchased copy of [Diablo 2](https://us.shop.battle.net/en-us/product/diablo-ii) and its expansion [Lord of Destruction](https://us.shop.battle.net/en-us/product/diablo-ii-lord-of-destruction) installed on your computer in order to run this engine. If you have an original copy of the disks, those files should work fine as well.

Please note that **this game is neither developed by, nor endorsed by Blizzard or its parent company Activision**.

This game is a clean-room implementation based on observations of how the original game works. Aside from the data file formats themselves, we have not and will not reverse engineer the original binaries of the game in an attempt to copy or duplicate intellectual property.

Diablo 2 and its content is ©2000 Blizzard Entertainment, Inc. All rights reserved. Diablo and Blizzard Entertainment are trademarks or registered trademarks of Blizzard Entertainment, Inc. in the U.S. and/or other countries.

ALL OTHER TRADEMARKS ARE THE PROPERTY OF THEIR RESPECTIVE OWNERS.

## Building On Windows
To build this engine, you simply need to have [Microsoft Visual Studio 2017](https://visualstudio.microsoft.com/downloads/) installed with C#/Windows support. Also make sure that only the x64 architecture is selected as we are not shipping 32-bit versions of SDL currently.

## Building On Linux
You need to have MonoDevelop installed, as well as any depenencies for that. You also need LibSDL2 installed (installing via your favorite package manager should be fine).

## Command Line Parameters
| Long Name    | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| --datapath   | Defines the path where the data files can be found           |
| --hwmouse    | Use the hardware mouse instead of software                   |
| --mousescale | When hardware mouse is enabled, this defines the pixel scale of the mouse. No effect for software mode |

By default OpenDiablo2 assumes the media files are in the same folder as itself. At minimum you'll most likely want to specify this option:

```-p "C:\Program Files (x86)\Diablo II"```

Substitute the path with wherever you have installed Diablo 2 and its expansions.

When running via VisualStudio, go to the debug section of the OpenDiablo2 project and specify the command line options there. At minimum you'll want to specify the path to the data files.



## Contributing
If you find something you'd like to fix thats obviously broken, create a branch, commit your code, and submit a pull request. If it's a new or missing feature you'd like to see, add an issue, and be descriptive! 
If you'd like to help out and are not quite sure how, you can look through any open issues and tasks.
