package d2common

import "time"

func Now() float64 {
	return float64(time.Now().UnixNano()) / 1000000000.0
}
