package d2cof

import "testing"

func TestCOF_New(t *testing.T) {
	c := New()

	if c == nil {
		t.Error("method New created nil instance")
	}
}

func TestCOF_Marshal_Unmarshal(t *testing.T) {
	cof1 := New()
	cof2 := New()

	var err error

	err = cof1.Unmarshal(make([]byte, 1000))
	if err != nil {
		t.Error(err)
	}

	cof1.Speed = 255
	data1 := cof1.Marshal()

	err = cof2.Unmarshal(data1)
	if err != nil {
		t.Error(err)
	}

	if cof2.Speed != cof1.Speed {
		t.Error("marshaled data does not match unmarshaled data")
	}
}
