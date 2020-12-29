// Package d2calculation contains code for calculation nodes.
package d2calculation

import (
	"fmt"
	"strconv"
)

// Calculation is the interface of every evaluatable calculation.
type Calculation interface {
	fmt.Stringer
	Eval() int
}

// BinaryCalculation is a calculation with a binary function or operator.
type BinaryCalculation struct {
	// Left is the left operand.
	Left Calculation

	// Right is the right operand.
	Right Calculation

	// Op is the actual operation.
	Op func(v1, v2 int) int
}

// Eval evaluates the calculation.
func (node *BinaryCalculation) Eval() int {
	return node.Op(node.Left.Eval(), node.Right.Eval())
}

func (node *BinaryCalculation) String() string {
	return "Binary(" + node.Left.String() + "," + node.Right.String() + ")"
}

// UnaryCalculation is a calculation with a unary function or operator.
type UnaryCalculation struct {
	// Child is the operand.
	Child Calculation

	// Op is the operation.
	Op func(v int) int
}

// Eval evaluates the calculation.
func (node *UnaryCalculation) Eval() int {
	return node.Op(node.Child.Eval())
}

func (node *UnaryCalculation) String() string {
	return "Unary(" + node.Child.String() + ")"
}

// TernaryCalculation is a calculation with a ternary function or operator.
type TernaryCalculation struct {
	// Left is the left operand.
	Left Calculation

	// Middle is the middle operand.
	Middle Calculation

	// Right is the right operand.
	Right Calculation
	Op    func(v1, v2, v3 int) int
}

// Eval evaluates the calculation.
func (node *TernaryCalculation) Eval() int {
	return node.Op(node.Left.Eval(), node.Middle.Eval(), node.Right.Eval())
}

func (node *TernaryCalculation) String() string {
	return "Ternary(" + node.Left.String() + "," + node.Middle.String() + "," + node.Right.String() + ")"
}

// PropertyReferenceCalculation is the calculation representing a property.
type PropertyReferenceCalculation struct {
	Type      string
	Name      string
	Qualifier string
}

// Eval evaluates the calculation.
func (node *PropertyReferenceCalculation) Eval() int {
	return 1
}

func (node *PropertyReferenceCalculation) String() string {
	return "Property(" + node.Type + "," + node.Name + "," + node.Qualifier + ")"
}

// ConstantCalculation is a constant value.
type ConstantCalculation struct {
	// Value is the constant value.
	Value int
}

// Eval evaluates the calculation.
func (node *ConstantCalculation) Eval() int {
	return node.Value
}

func (node *ConstantCalculation) String() string {
	return strconv.Itoa(node.Value)
}
