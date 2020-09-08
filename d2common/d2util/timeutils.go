package d2util

import "time"

const (
	nanoseconds = 1000000000.0
)

// Now returns how many seconds have elapsed since Unix time (January 1, 1970 UTC)
func Now() float64 {
	// Unix time in nanoseconds divided by how many nanoseconds in a second
	return float64(time.Now().UnixNano()) / nanoseconds
}
