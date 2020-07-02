package d2interface

import "math/big"

type Vector2 interface {
	X() *big.Float
	Y() *big.Float
	Marshal() ([]byte, error)
	Unmarshal(buf []byte) error
	Clone() Vector2
	Copy(src Vector2) Vector2
	SetFromEntity(entity WorldEntity) Vector2
	Set(x, y *big.Float) Vector2
	SetToPolar(azimuth, radius *big.Float) Vector2
	Equals(src Vector2) bool
	FuzzyEquals(src Vector2) bool
	Abs() Vector2
	Angle() *big.Float
	SetAngle(angle *big.Float) Vector2
	Add(src Vector2) Vector2
	Subtract(src Vector2) Vector2
	Multiply(src Vector2) Vector2
	Scale(value *big.Float) Vector2
	Divide(src Vector2) Vector2
	Negate() Vector2
	Distance(src Vector2) *big.Float
	DistanceSq(src Vector2) *big.Float
	Length() *big.Float
	SetLength(length *big.Float) Vector2
	LengthSq() (*big.Float, *big.Float)
	Normalize() Vector2
	NormalizeRightHand() Vector2
	NormalizeLeftHand() Vector2
	Dot(src Vector2) *big.Float
	Cross(src Vector2) *big.Float
	Lerp(src Vector2, t *big.Float) Vector2
	Reset() Vector2
	Limit(max *big.Float) Vector2
	Reflect(normal Vector2) Vector2
	Mirror(axis Vector2) Vector2
	Rotate(delta *big.Float) Vector2
}
