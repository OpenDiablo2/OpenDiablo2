package d2common

// BuildInfoRecord is the structure used to hold information about the current build
type BuildInfoRecord struct {
	// Branch is the branch this build is based on (or 'Local' if built locally)
	Branch string
	// Commit is the commit hash of the build (or blank if built locally)
	Commit string
}

// BuildInfo contains information about the build currently being ran
var BuildInfo BuildInfoRecord

// SetBuildInfo is called at the start of the application to generate the global BuildInfo value
func SetBuildInfo(branch, commit string) {
	BuildInfo = BuildInfoRecord{
		Branch: branch,
		Commit: commit,
	}
}
