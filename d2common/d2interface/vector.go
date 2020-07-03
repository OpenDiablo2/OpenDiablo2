package d2interface

import "math/big"

// Vector is a 2-dimensional vector implementation using big.Float
type Vector interface {
	X() *big.Float
	Y() *big.Float
	Marshal() ([]byte, error)
	Unmarshal(buf []byte) error
	Clone() Vector
	Copy(src Vector) Vector
	// SetFromEntity(entity WorldEntity) Vector
	Set(x, y *big.Float) Vector
	SetToPolar(azimuth, radius *big.Float) Vector
	Equals(src Vector) bool
	FuzzyEquals(src Vector) bool
	Abs() Vector
	Angle() *big.Float
	SetAngle(angle *big.Float) Vector
	Add(src Vector) Vector
	Subtract(src Vector) Vector
	Multiply(src Vector) Vector
	Scale(value *big.Float) Vector
	Divide(src Vector) Vector
	Negate() Vector
	Distance(src Vector) *big.Float
	DistanceSq(src Vector) *big.Float
	Length() *big.Float
	SetLength(length *big.Float) Vector
	LengthSq() (*big.Float, *big.Float)
	Normalize() Vector
	NormalizeRightHand() Vector
	NormalizeLeftHand() Vector
	Dot(src Vector) *big.Float
	Cross(src Vector) *big.Float
	Lerp(src Vector, t *big.Float) Vector
	Reset() Vector
	Limit(max *big.Float) Vector
	Reflect(normal Vector) Vector
	Mirror(axis Vector) Vector
	Rotate(delta *big.Float) Vector
}
