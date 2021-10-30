package d2parser

import (
	"math"
	"math/rand"
)

type binaryOperation struct {
	Function          func(v1, v2 int) int
	Operator          string
	Precedence        int
	IsRightAssociated bool
}

type unaryOperation struct {
	Function   func(v int) int
	Operator   string
	Precedence int
}

type ternaryOperation struct {
	Function          func(v1, v2, v3 int) int
	Operator          string
	Marker            string
	Precedence        int
	IsRightAssociated bool
}

func getUnaryOperations() map[string]unaryOperation {
	return map[string]unaryOperation{
		"+": {
			func(v int) int {
				return v
			},
			"+",
			4,
		},
		"-": {
			func(v int) int {
				return -v
			},
			"-",
			4,
		},
	}
}

func getTernaryOperations() map[string]ternaryOperation {
	return map[string]ternaryOperation{
		"?": {
			func(v1, v2, v3 int) int {
				if v1 != 0 {
					return v2
				}
				return v3
			},
			"?",
			":",
			0,
			true,
		},
	}
}

func getBinaryOperations() map[string]binaryOperation { //nolint:funlen // No reason to split function, just creates the operations.
	return map[string]binaryOperation{
		"==": {
			func(v1, v2 int) int {
				if v1 == v2 {
					return 1
				}
				return 0
			},
			"==",
			1,
			false,
		},
		"!=": {
			func(v1, v2 int) int {
				if v1 != v2 {
					return 1
				}
				return 0
			},
			"!=",
			1,
			false,
		},
		"<": {
			func(v1, v2 int) int {
				if v1 < v2 {
					return 1
				}
				return 0
			},
			"<",
			2,
			false,
		},
		">": {
			func(v1, v2 int) int {
				if v1 > v2 {
					return 1
				}
				return 0
			},
			">",
			2,
			false,
		},
		"<=": {
			func(v1, v2 int) int {
				if v1 <= v2 {
					return 1
				}
				return 0
			},
			"<=",
			2,
			false,
		},
		">=": {
			func(v1, v2 int) int {
				if v1 >= v2 {
					return 1
				}
				return 0
			},
			">=",
			2,
			false,
		},
		"+": {
			func(v1, v2 int) int {
				return v1 + v2
			},
			"+",
			3,
			false,
		},
		"-": {
			func(v1, v2 int) int {
				return v1 - v2
			},
			"-",
			3,
			false,
		},
		"*": {
			func(v1, v2 int) int {
				return v1 * v2
			},
			"*",
			5,
			false,
		},
		"/": {
			func(v1, v2 int) int {
				return v1 / v2
			},
			"/",
			5,
			false,
		},
		"^": {
			func(v1, v2 int) int {
				return int(math.Pow(float64(v1), float64(v2)))
			},
			"^",
			6,
			true,
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
