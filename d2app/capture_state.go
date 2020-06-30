package d2app

type captureState int

const (
	captureStateNone captureState = iota
	captureStateFrame
	captureStateGif
)
