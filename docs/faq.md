# FAQ

> Q: Should I have experience in game engine programming to be able to contribute?

Absolutely not!
Most of our contributors are _not_ game developers.
Actually, for lots of us, this is both our first _Go_ project, and our first _game_ project.

> Q: I am a developer, but I don't have experience writing Go. How much of a problem could this be?

For people who are absolutely new to Golang, we usually recommend going through the [Tour of Go] tutorial.
Even with half completion of this course, you should feel confident enough to submit your first PRs.

> Q: Is it difficult to understand how the code works?

Currently, the barrier to entry can be quite high due to coupling among various parts of the codebase.
That said, to understand how things are working behind the scenes, you have to go through a lot of lines of code.
Getting a grip on the current state of the project requires a lot of reading and debugging.

> Q: How can I become familiar with the codebase?

You can practice building small apps in [Ebiten] (current engine) and [SDL2] (future engine).
Also feel free to read about the [Entity Component System] model (or read [Akara] code).
Until we further decouple the larger systems in the codebase, we're going to have this issue where you need to be familiar with a huge chunk of the codebase.
Until we rectify this issue, your best bet is to just ask a lot of questions, and those that know will be happy to chime in and try to explain.

> What do I need to know about Diablo 2?

First off, you should be familiar with the gameplay.
Apart from that, you only need to know some basic information about Diablo 2 mod-making.
For example, you have to know the basic game's file formats, such as `MPQ`, `DC6`, `DS1`, `DT1`, etc.
You can find a lot of good info in `#file-formats` channel on our Discord server and on [PhrozenKeep] (the d2 modding community forum).
There are [plenty useful information for the Diablo II file formats](https://d2mods.info/forum/viewtopic.php?f=7&t=724) and other info such as the [Levels.txt](https://d2mods.info/forum/viewtopic.php?t=6754) file.

> Q: What's [Akara]?

Writing tests (in isolation) for something in the codebase is difficult because of the high level of coupling.
The solution was found at [ECS], which is fantastic at de-coupling game code.
[Akara] is the [ECS] implementation for Go we have selected.
For more information visit the [Akara] repository.
Using [Akara] we are able to test individual components in isolation without problems.

> Q: What's [AbyssEngine]?

Currently, everything in OpenDiablo2 is hard coded.
We are moving the pieces, so we can make everything into a custom `MPQ` (a 'shim' `MPQ`) -- what will be `OpenDiablo2.mpq`.
The engine itself will be stripped of hard coded logic and moved to [AbyssEngine].
That way there's a clean separation between the game, and the engine code.

> Q: Are there any other project related information?

Project Boards:

- [OpenDiablo2 Project](https://github.com/OpenDiablo2/OpenDiablo2/projects/4)
- [HellSpawner Project](https://github.com/orgs/OpenDiablo2/projects/7)

Epics:

- [OpenDiablo2 Milestones](https://github.com/OpenDiablo2/OpenDiablo2/milestones)

Other docs:

- [Diablo 2 Data Tables](https://docs.google.com/spreadsheets/d/13Wo58CNxDQlQiZm066dAWVVU4kmgKayn0zdyvmU18AM/edit#gid=330752700)
- [Official Documentation PDF](https://github.com/OpenDiablo2/SystemRequirementsSpecs/releases)

[HellSpawner]: https://github.com/OpenDiablo2/HellSpawner
[Entity Component System]: https://en.wikipedia.org/wiki/Entity_component_system
[ECS]: https://en.wikipedia.org/wiki/Entity_component_system
[Akara]: https://github.com/gravestench/akara
[Ebiten]: https://ebiten.org/
[SDL2]: https://github.com/veandco/go-sdl2
[Tour of Go]: https://tour.golang.org/welcome/1
[PhrozenKeep]: https://d2mods.info/home.php
[AbyssEngine]: https://github.com/OpenDiablo2/AbyssEngine
