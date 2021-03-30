## Status

We are currently working on features necessary to play Diablo 2 in its entirety.
After this is completed, we will work on expanding the project to include tools and plugin support for modding, as well as writing completely new games with the engine.

At the moment (March 2021) the game starts, you can select any character and run around Act1 town.
You can also open any of the game's panels.

Much work has been made in the background, but a lot of work still has to be done for the game to be playable.

> Q: Where help is currently most needed?

The best way to help right now is to focus all of your energy on [HellSpawner].
This is a GUI application for viewing, editing, and creating Diablo II files.
This will help us work on `OpenDiablo2.mpq` using [Hellspawner].
In other words, the [OpenDiablo2] codebase will be soon obsoleted.
We're currently moving the core features into the [AbyssEngine] (game engine), and the implementation specific stuff will go into a _shim_ `MPQ` which later will become the OpenDiablo 2 project.
The [OpenDiablo2] repository will become the new place for the `OpenDiablo2.mpq` file, which [AbyssEngine] is meant to use it.
So it's not that the current logic is going to be removed/wasted, but we will just need to translated its appropriate parts:

* In-game logic and interactivity will be translated to Javascript
* while things like entity composition and map rendering will implemented by the game engine.

Feel free to contribute!

[AbyssEngine]: https://github.com/OpenDiablo2/AbyssEngine
[OpenDiablo2]: https://github.com/OpenDiablo2/OpenDiablo2
[HellSpawner]: https://github.com/OpenDiablo2/HellSpawner
