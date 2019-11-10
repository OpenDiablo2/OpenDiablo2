# OpenDiablo2 Enums
Items in this folder are compiled with two programs. You can obtain them
by running the following:
```
go get golang.org/x/tools/cmd/stringer
go get github.com/mewspring/tools/cmd/string2enum
```
Once you have the tools installed, simply run the following command in this
folder to regenerate the support files:
```
generate
```
If you add any new files, make sure to add the following to the end of the
file:
```go
//go:generate stringer -linecomment -type AnimationMode
```
