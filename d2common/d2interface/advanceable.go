package d2interface

type Advanceable interface {
	Advance(elapsedTime, currentTime float64) error
}
