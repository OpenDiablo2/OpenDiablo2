package d2interface

type WorldEntity interface {
	Serializer
	Position() Vector
}
