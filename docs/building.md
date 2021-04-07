## Building

To pull the project down, run `go get github.com/OpenDiablo2/OpenDiablo2`

On windows this folder will most likely be in `%USERPROFILE%\go\src\github.com\OpenDiablo2\OpenDiablo2`

In the root folder, run `go get -d` to pull down all dependencies.

To run the project, run `go run .` from the root folder.

You can also open the root folder in VSCode. Make sure you have the `ms-vscode.go` plugin installed.

### Linux

There are several dependencies which need to be installed additionally.
To install them you can use `./build.sh` in the project root folder - this script takes care of the installation for you.

### Windows

### MacOS

1. Before start, download and install Go programming language [here](https://golang.org/doc/install) if needed.
2. Launch Terminal (Note: Restart any open Terminal console if needed)
3. Fetch the OpenDiablo2 project with this comand: `go get github.com/OpenDiablo2/OpenDiablo2`
4. The OpenDiablo2 will be at `~/go/bin/OpenDiablo2` but we still need to download the official Diablo 2 and LoD.
5. Purchase and download Diablo II (2000) and Diablo II: Lord of Destruction (2001) from Battle.net. The downloaders will be in form of EXE.
6. Find a Windows PC, or use Virtual Machine (e.g. VirtualBox), or Bootcamp for Mac, or Parallels, to install both games using the downloaders.
7. Copy the installed games from Windows to macOS's `/Applications/Diablo II/` folder. Requires administrator permission.
8. Make sure the `/Applications/Diablo II/patch_d2.mpq` is in place.
9. Run OpenDialo2 by `~/go/bin/OpenDiablo2` in Terminal.
10. Enjoy.
