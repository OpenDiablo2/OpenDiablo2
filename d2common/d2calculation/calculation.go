package d2calculation

import (
	"fmt"
	"strconv"
)

type Calculation interface {
	fmt.Stringer
	Eval() int
}

type BinaryCalculation struct {
	Left  Calculation
	Right Calculation
	Op    func(v1, v2 int) int
}

func (node *BinaryCalculation) Eval() int {
	return node.Op(node.Left.Eval(), node.Right.Eval())
}

func (node *BinaryCalculation) String() string {
	return "Binary(" + node.Left.String() + "," + node.Right.String() + ")"
}

type UnaryCalculation struct {
	Child Calculation
	Op    func(v int) int
}

func (node *UnaryCalculation) Eval() int {
	return node.Op(node.Child.Eval())
}

func (node *UnaryCalculation) String() string {
	return "Unary(" + node.Child.String() + ")"
}

type TernaryCalculation struct {
	Left   Calculation
	Middle Calculation
	Right  Calculation
	Op     func(v1, v2, v3 int) int
}

func (node *TernaryCalculation) Eval() int {
	return node.Op(node.Left.Eval(), node.Middle.Eval(), node.Right.Eval())
}

func (node *TernaryCalculation) String() string {
	return "Ternary(" + node.Left.String() + "," + node.Middle.String() + "," + node.Right.String() + ")"
}

type PropertyReferenceCalculation struct {
	Type      string
	Name      string
	Qualifier string
}

func (node *PropertyReferenceCalculation) Eval() int {
	return 1
}

func (node *PropertyReferenceCalculation) String() string {
	return "Property(" + node.Type + "," + node.Name + "," + node.Qualifier + ")"
}

type ConstantCalculation struct {
	Value int
}

func (node *ConstantCalculation) Eval() int {
	return node.Value
}

func (node *ConstantCalculation) String() string {
	return strconv.Itoa(node.Value)
}
