package d2interface

type Serializer interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
