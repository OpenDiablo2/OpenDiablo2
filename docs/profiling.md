## Profiling

There are many profiler options to debug performance issues.
These can be enabled by supplying the following command-line option and are saved in the `pprof` directory:

`go run . --profile=cpu`

Available profilers:\
`cpu` `mem` `block` `goroutine` `trace` `thread` `mutex`

You can export the profiler output with the following command:\
`go tool pprof --pdf ./OpenDiablo2 pprof/profiler.pprof > file.pdf`

In game you can create a heap dump by pressing `~` and typing `dumpheap`. A heap.pprof is written to the `pprof` directory.

You may need to install [Graphviz](http://www.graphviz.org/download/) in order to convert the profiler output.
