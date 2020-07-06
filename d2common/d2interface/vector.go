package d2interface

import "math/big"

// Vector is a 2-dimensional vector implementation using big.Float
type Vector interface {

	// Return multiple numeric types

	// XYBigFloat returns the values as big.Float.
	XYBigFloat() (*big.Float, *big.Float)

	// XYFloat64 returns the values as float64.
	XYFloat64() (*float64, *float64)

	Equals(v Vector) bool
	EqualsF(v Vector) bool
	CompareF(Vector) (int, int)

	Set(x, y float64) Vector
	Clone() Vector

	Floor() Vector
	Add(v Vector) Vector
	Subtract(v Vector) Vector
	Multiply(v Vector) Vector
	Scale(s float64) Vector
	Divide(v Vector) Vector
	Abs() Vector
	Negate() Vector

	Distance(v Vector) float64
	Length() float64

	String() string

	/*

		Marshal() ([]byte, error)
		Unmarshal(buf []byte) error
		Copy(src Vector) Vector
		// SetFromEntity(entity WorldEntity) Vector
		SetToPolar(azimuth, radius *big.Float) Vector
		Angle() *big.Float
		SetAngle(angle *big.Float) Vector
		Scale(value *big.Float) Vector
		DistanceSq(src Vector) *big.Float
		LengthSq() (*big.Float, *big.Float)
		Normalize() Vector
		NormalizeRightHand() Vector
		NormalizeLeftHand() Vector
		Cross(src Vector) *big.Float
		Lerp(src Vector, t *big.Float) Vector
		Reset() Vector
		Limit(max *big.Float) Vector
		Reflect(normal Vector) Vector
		Mirror(axis Vector) Vector
		Rotate(delta *big.Float) Vector
	*/
}
