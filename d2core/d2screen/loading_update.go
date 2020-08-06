package d2screen

const progressCompleted = 1.0

type loadingUpdate struct {
	progress float64
	err      error
	done     bool
}

// Error provides a way for callers to report an error during loading.
func (l *LoadingState) Error(err error) {
	l.updates <- loadingUpdate{err: err}
}

// Progress provides a way for callers to report the ratio between `0` and `1` of the progress made loading a screen.
func (l *LoadingState) Progress(ratio float64) {
	l.updates <- loadingUpdate{progress: ratio}
}

// Done provides a way for callers to report that screen loading has been completed.
func (l *LoadingState) Done() {
	l.updates <- loadingUpdate{progress: progressCompleted}
	l.updates <- loadingUpdate{done: true}
}
