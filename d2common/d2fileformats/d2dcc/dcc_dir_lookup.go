package d2dcc

type directionCount int

const (
	four directionCount = 4 << iota
	eight
	sixteen
	thirtyTwo
	sixtyFour
)

// Dir64ToDcc returns the DCC direction based on the actual direction.
// Special thanks for Necrolis for these tables!
func Dir64ToDcc(direction, numDirections int) int {
	var dir4 = [64]int{
		0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
		3, 3, 3, 3, 3, 3, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0}

	var dir8 = [64]int{
		4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5, 5, 5,
		5, 5, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 6, 6, 6, 6,
		6, 6, 6, 6, 2, 2, 2, 2, 2, 2, 2, 2, 7, 7, 7, 7,
		7, 7, 7, 7, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4}

	var dir16 = [64]int{
		4, 4, 8, 8, 8, 8, 0, 0, 0, 0, 9, 9, 9, 9, 5, 5,
		5, 5, 10, 10, 10, 10, 1, 1, 1, 1, 11, 11, 11, 11, 6, 6,
		6, 6, 12, 12, 12, 12, 2, 2, 2, 2, 13, 13, 13, 13, 7, 7,
		7, 7, 14, 14, 14, 14, 3, 3, 3, 3, 15, 15, 15, 15, 4, 4}

	var dir32 = [64]int{
		4, 16, 16, 8, 8, 17, 17, 0, 0, 18, 18, 9, 9, 19, 19, 5,
		5, 20, 20, 10, 10, 21, 21, 1, 1, 22, 22, 11, 11, 23, 23, 6,
		6, 24, 24, 12, 12, 25, 25, 2, 2, 26, 26, 13, 13, 27, 27, 7,
		7, 28, 28, 14, 14, 29, 29, 3, 3, 30, 30, 15, 15, 31, 31, 4}

	var dir64 = [64]int{
		4, 32, 16, 33, 8, 34, 17, 35, 0, 36, 18, 37, 9, 38, 19, 39,
		5, 40, 20, 41, 10, 42, 21, 43, 1, 44, 22, 45, 11, 46, 23, 47,
		6, 48, 24, 49, 12, 50, 25, 51, 2, 52, 26, 53, 13, 54, 27, 55,
		7, 56, 28, 57, 14, 58, 29, 59, 3, 60, 30, 61, 15, 62, 31, 63}

	switch directionCount(numDirections) {
	case four:
		return dir4[direction]
	case eight:
		return dir8[direction]
	case sixteen:
		return dir16[direction]
	case thirtyTwo:
		return dir32[direction]
	case sixtyFour:
		return dir64[direction]
	default:
		return 0
	}
}
