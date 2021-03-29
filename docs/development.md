## VS Code Extensions

The following extensions are recommended for working with this project:

*   ms-vscode.go
*   defaltd.go-coverage-viewer

When you open the workspace for the first time, Visual Studio Code will automatically suggest these extensions for installation.

Alternatively you can get to it by going to settings <kbd>Ctrl+,</kbd>, expanding `Extensions` and selecting `Go configuration`, then clicking on `Edit in settings.json`.
Just paste that section where appropriate.

We use `golangci-lint` to catch lint errors, and we require all contributors to install and use it.
As of Oct. 26, 2020 we will no longer be accepting pull requests that introduce lint errors.
Installation instructions can be found [here](https://golangci-lint.run/usage/install/).
