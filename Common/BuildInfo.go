package Common

type BuildInfoRecord struct {
	Branch string
	Commit string
}

var BuildInfo BuildInfoRecord

func SetBuildInfo(branch, commit string) {
	BuildInfo = BuildInfoRecord{
		Branch: branch,
		Commit: commit,
	}
}
