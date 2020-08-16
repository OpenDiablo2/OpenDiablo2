package d2parser

import (
	"log"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation/d2lexer"
)

// Parser is a parser for calculations used for skill and missiles.
type Parser struct {
	lex *d2lexer.Lexer

	binaryOperations  map[string]binaryOperation
	unaryOperations   map[string]unaryOperation
	ternaryOperations map[string]ternaryOperation
	fixedFunctions    map[string]func(v1, v2 int) int

	currentType string
	currentName string
}

// New creates a new parser.
func New() *Parser {
	return &Parser{
		binaryOperations:  getBinaryOperations(),
		unaryOperations:   getUnaryOperations(),
		ternaryOperations: getTernaryOperations(),
		fixedFunctions:    getFunctions(),
	}
}

// SetCurrentReference sets the current reference type and name, such as "skill" and skill name.
func (parser *Parser) SetCurrentReference(propType, propName string) {
	parser.currentType = propType
	parser.currentName = propName
}

// Parse parses the calculation string and creates a Calculation tree.
func (parser *Parser) Parse(calc string) d2calculation.Calculation {
	calc = strings.TrimSpace(calc)
	if calc == "" {
		return &d2calculation.ConstantCalculation{Value: 0}
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error parsing calculation: %v", calc)
		}
	}()

	parser.lex = d2lexer.New([]byte(calc))

	return parser.parseLevel(0)
}

func (parser *Parser) peek() d2lexer.Token {
	return parser.lex.Peek()
}

func (parser *Parser) consume() d2lexer.Token {
	return parser.lex.NextToken()
}

func (parser *Parser) parseLevel(level int) d2calculation.Calculation {
	node := parser.parseProduction()

	t := parser.peek()
	if t.Type == d2lexer.EOF {
		return node
	}

	for {
		if t.Type != d2lexer.Symbol {
			break
		}

		op, ok := parser.binaryOperations[t.Value]
		if !ok || op.Precedence < level {
			break
		}

		parser.consume()

		var nextLevel int

		if op.IsRightAssociated {
			nextLevel = op.Precedence
		} else {
			nextLevel = op.Precedence + 1
		}

		otherCalculation := parser.parseLevel(nextLevel)
		node = &d2calculation.BinaryCalculation{
			Left:  node,
			Right: otherCalculation,
			Op:    op.Function,
		}
		t = parser.peek()
	}

	for {
		if t.Type != d2lexer.Symbol {
			break
		}

		op, ok := parser.ternaryOperations[t.Value]
		if !ok || op.Precedence < level {
			break
		}

		parser.consume()

		var nextLevel int

		if op.IsRightAssociated {
			nextLevel = op.Precedence
		} else {
			nextLevel = op.Precedence + 1
		}

		middleCalculation := parser.parseLevel(nextLevel)

		t = parser.peek()
		if t.Type != d2lexer.Symbol || t.Value != op.Marker {
			panic("Invalid ternary! " + t.Value + ", expected: " + op.Marker)
		}

		parser.consume()
		rightCalculation := parser.parseLevel(nextLevel)

		node = &d2calculation.TernaryCalculation{
			Left:   node,
			Middle: middleCalculation,
			Right:  rightCalculation,
			Op:     op.Function,
		}
		t = parser.peek()
	}

	return node
}

func (parser *Parser) parseProduction() d2calculation.Calculation {
	t := parser.peek()

	switch {
	case t.Type == d2lexer.Symbol:
		if t.Value == "(" {
			parser.consume()
			node := parser.parseLevel(0)

			t = parser.peek()
			if t.Type != d2lexer.Symbol ||
				t.Value != ")" {
				if t.Type == d2lexer.EOF { // Ignore unclosed final parenthesis due to syntax error in original Fire Wall calculation.
					return node
				}

				panic("Parenthesis not closed!")
			}

			parser.consume()

			return node
		}

		op, ok := parser.unaryOperations[t.Value]
		if !ok {
			panic("Invalid unary symbol: " + t.Value)
		}

		parser.consume()
		node := parser.parseLevel(op.Precedence)

		return &d2calculation.UnaryCalculation{
			Child: node,
			Op:    op.Function,
		}

	case t.Type == d2lexer.Name || t.Type == d2lexer.Number:
		return parser.parseLeafCalculation()
	default:
		panic("Expected parenthesis, unary operator, function or value!")
	}
}

func (parser *Parser) parseLeafCalculation() d2calculation.Calculation {
	t := parser.peek()

	if t.Type == d2lexer.Number {
		val, err := strconv.Atoi(t.Value)
		if err != nil {
			panic("Invalid number: " + t.Value)
		}

		parser.consume()

		return &d2calculation.ConstantCalculation{Value: val}
	}

	if t.Value == "skill" ||
		t.Value == "miss" ||
		t.Value == "stat" {
		return parser.parseProperty()
	}

	if parser.fixedFunctions[t.Value] != nil {
		return parser.parseFunction(t.Value)
	}

	if t.Type == d2lexer.Name {
		parser.consume()

		return &d2calculation.PropertyReferenceCalculation{
			Type:      parser.currentType,
			Name:      parser.currentName,
			Qualifier: t.Value,
		}
	}

	panic(t.Value + " is not a function, property, or number!")
}

func (parser *Parser) parseFunction(name string) d2calculation.Calculation {
	function := parser.fixedFunctions[name]
	parser.consume()

	t := parser.peek()
	if t.Value != "(" {
		panic("Invalid function!")
	}

	parser.consume()

	firstParam := parser.parseLevel(0)

	t = parser.peek()
	if t.Type != d2lexer.Symbol || t.Value != "," {
		panic("Invalid function!")
	}

	parser.consume()

	secondParam := parser.parseLevel(0)

	t = parser.peek()
	if t.Value != ")" {
		panic("Invalid function!")
	}

	parser.consume()

	return &d2calculation.BinaryCalculation{
		Left:  firstParam,
		Right: secondParam,
		Op:    function,
	}
}

func (parser *Parser) parseProperty() d2calculation.Calculation {
	t := parser.peek()
	propType := t.Value
	t = parser.consume()

	t = parser.peek()
	if t.Value != "(" {
		panic("Invalid property: " + propType + ", open parenthesis missing.")
	}

	parser.consume()

	t = parser.peek()
	if t.Type != d2lexer.String {
		panic("Property name must be in quotes: " + propType)
	}

	propName := t.Value

	parser.consume()

	t = parser.peek()
	if t.Type != d2lexer.Symbol || t.Value != "." {
		panic("Property name must be followed by dot: " + propType)
	}

	parser.consume()

	t = parser.peek()
	if t.Type != d2lexer.Name {
		panic("Invalid propery qualifier: " + propType)
	}

	propQual := t.Value

	parser.consume()

	t = parser.peek()
	if t.Value != ")" {
		panic("Invalid property: " + propType + ", closed parenthesis missing.")
	}

	parser.consume()

	return &d2calculation.PropertyReferenceCalculation{
		Type:      propType,
		Name:      propName,
		Qualifier: propQual,
	}
}
