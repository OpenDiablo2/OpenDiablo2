package main

import (
	"github.com/pkg/profile"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2app"
)

// GitBranch is set by the CI build process to the name of the branch
//nolint:gochecknoglobals This is filled in by the build system
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
//nolint:gochecknoglobals This is filled in by the build system
var GitCommit string

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")

	profileOption := kingpin.Flag("profile", "Profiles the program, one of (cpu, mem, block, goroutine, trace, thread, mutex)").String()
	kingpin.Parse()

	if len(*profileOption) > 0 {
		profiler := enableProfiler(*profileOption)
		if profiler != nil {
			defer profiler.Stop()
		}
	}

	app := d2app.Create(GitBranch, GitCommit)
	if err := app.Run(); err != nil {
		return
	}
}

func enableProfiler(profileOption string) interface{ Stop() } {
	var options []func(*profile.Profile)

	switch strings.ToLower(strings.Trim(profileOption, " ")) {
	case "cpu":
		log.Printf("CPU profiling is enabled.")

		options = append(options, profile.CPUProfile)
	case "mem":
		log.Printf("Memory profiling is enabled.")

		options = append(options, profile.MemProfile)
	case "block":
		log.Printf("Block profiling is enabled.")

		options = append(options, profile.BlockProfile)
	case "goroutine":
		log.Printf("Goroutine profiling is enabled.")

		options = append(options, profile.GoroutineProfile)
	case "trace":
		log.Printf("Trace profiling is enabled.")

		options = append(options, profile.TraceProfile)
	case "thread":
		log.Printf("Thread creation profiling is enabled.")

		options = append(options, profile.ThreadcreationProfile)
	case "mutex":
		log.Printf("Mutex profiling is enabled.")

		options = append(options, profile.MutexProfile)
	}

	options = append(options, profile.ProfilePath("./pprof/"))

	if len(options) > 1 {
		return profile.Start(options...)
	}

	return nil
}
