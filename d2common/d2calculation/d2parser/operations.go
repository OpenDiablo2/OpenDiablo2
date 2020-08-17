package d2parser

import (
	"math"
	"math/rand"
)

type binaryOperation struct {
	Operator          string
	Precedence        int
	IsRightAssociated bool
	Function          func(v1, v2 int) int
}

type unaryOperation struct {
	Operator   string
	Precedence int
	Function   func(v int) int
}

type ternaryOperation struct {
	Operator          string
	Marker            string
	Precedence        int
	IsRightAssociated bool
	Function          func(v1, v2, v3 int) int
}

func getUnaryOperations() map[string]unaryOperation {
	return map[string]unaryOperation{
		"+": {
			"+",
			4,
			func(v int) int {
				return v
			},
		},
		"-": {
			"-",
			4,
			func(v int) int {
				return -v
			},
		},
	}
}

func getTernaryOperations() map[string]ternaryOperation {
	return map[string]ternaryOperation{
		"?": {
			"?",
			":",
			0,
			true,
			func(v1, v2, v3 int) int {
				if v1 != 0 {
					return v2
				}
				return v3
			},
		},
	}
}

func getBinaryOperations() map[string]binaryOperation { //nolint:funlen // No reason to split function, just creates the operations.
	return map[string]binaryOperation{
		"==": {
			"==",
			1,
			false,
			func(v1, v2 int) int {
				if v1 == v2 {
					return 1
				}
				return 0
			},
		},
		"!=": {
			"!=",
			1,
			false,
			func(v1, v2 int) int {
				if v1 != v2 {
					return 1
				}
				return 0
			},
		},
		"<": {
			"<",
			2,
			false,
			func(v1, v2 int) int {
				if v1 < v2 {
					return 1
				}
				return 0
			},
		},
		">": {
			">",
			2,
			false,
			func(v1, v2 int) int {
				if v1 > v2 {
					return 1
				}
				return 0
			},
		},
		"<=": {
			"<=",
			2,
			false,
			func(v1, v2 int) int {
				if v1 <= v2 {
					return 1
				}
				return 0
			},
		},
		">=": {
			">=",
			2,
			false,
			func(v1, v2 int) int {
				if v1 >= v2 {
					return 1
				}
				return 0
			},
		},
		"+": {
			"+",
			3,
			false,
			func(v1, v2 int) int {
				return v1 + v2
			},
		},
		"-": {
			"-",
			3,
			false,
			func(v1, v2 int) int {
				return v1 - v2
			},
		},
		"*": {
			"*",
			5,
			false,
			func(v1, v2 int) int {
				return v1 * v2
			},
		},
		"/": {
			"/",
			5,
			false,
			func(v1, v2 int) int {
				return v1 / v2
			},
		},
		"^": {
			"^",
			6,
			true,
			func(v1, v2 int) int {
				return int(math.Pow(float64(v1), float64(v2)))
			},
		},
	}
}

func getFunctions() map[string]func(v1, v2 int) int {
	return map[string]func(v1, v2 int) int{
		"min": func(v1, v2 int) int {
			if v1 < v2 {
				return v1
			}
			return v2
		},
		"max": func(v1, v2 int) int {
			if v1 > v2 {
				return v1
			}
			return v2
		},
		"rand": func(v1, v2 int) int {
			if rand.Int()%2 == 0 { //nolint:gosec // Secure random not necessary.
				return v1
			}
			return v2
		},
	}
}
