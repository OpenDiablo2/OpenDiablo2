package d2math

import "testing"

func TestRangedNumber_Clone(t *testing.T) {
	rn1 := &RangedNumber{1, 1}
	rn2 := rn1.Clone()

	if &rn1 == &rn2 {
		t.Errorf("Cloned ranged number is not unique: *%d == *%d", &rn1, &rn2)
	}
}

func TestRangedNumber_Copy(t *testing.T) {
	rn1 := &RangedNumber{1, 1}
	rn2 := rn1.Clone().Set(-1, -1)
	rn1.Copy(rn2)

	if rn1.Min() != rn2.Min() || rn1.Max() != rn2.Max() {
		t.Errorf("Min/Max values were not copied: %d != %d", rn1, rn2)
	}
}

func TestRangedNumber_Set(t *testing.T) {
	badMin, badMax := 10, -10
	rn := &RangedNumber{badMin, badMax} // should get swapped when used

	if rn.Min() == badMin || rn.Max() == badMax {
		t.Errorf("Min/Max values were not ordered: %d", rn)
	}

	rn.Set(badMin, badMax)

	if rn.Min() == badMin || rn.Max() == badMax {
		t.Errorf("Min/Max values were not ordered: %d", rn)
	}
}

func TestRangedNumber_Equals(t *testing.T) {
	min, max := 1, 2
	rn1 := &RangedNumber{min, max}
	rn2 := rn1.Clone()

	if !rn1.Equals(rn2) {
		t.Errorf("Not equal when they should be: %d == %d", rn1, rn2)
	}

	rn1.Set(3, 4)

	if rn1.Equals(rn2) {
		t.Errorf("equal when they should NOT be: %d != %d", rn1, rn2)
	}
}

func TestRangedNumber_Add(t *testing.T) {
	min, max := 1, 2
	rn1 := &RangedNumber{min, max}
	rn2 := rn1.Clone()

	rn1.Add(rn2)

	if rn1.Min() != 2 || rn1.Max() != 4 {
		t.Errorf("Unexpected value after addition: %d ", rn1)
	}
}

func TestRangedNumber_Sub(t *testing.T) {
	min, max := 1, 2
	rn1 := &RangedNumber{min, max}
	rn2 := rn1.Clone()

	rn1.Sub(rn2)

	if rn1.Min() != 0 || rn1.Max() != 0 {
		t.Errorf("Unexpected value after subtraction: %d ", rn1)
	}
}

func TestRangedNumber_Mul(t *testing.T) {
	min, max := 1, 2
	rn1 := &RangedNumber{min, max}
	rn2 := rn1.Clone()

	rn1.Mul(rn2)

	if rn1.Min() != 1 || rn1.Max() != 4 {
		t.Errorf("Unexpected value after multiplication: %d ", rn1)
	}
}

func TestRangedNumber_Div(t *testing.T) {
	min, max := 1, 2
	rn1 := &RangedNumber{min, max}
	rn2 := rn1.Clone()

	rn1.Div(rn2)

	if rn1.Min() != 1 || rn1.Max() != 1 {
		t.Errorf("Unexpected value after division: %d ", rn1)
	}
}

func TestRangedNumber_String(t *testing.T) {
	rn := &RangedNumber{}

	if rn.String() != "0" {
		t.Errorf("Unexpected string output: %d ", rn)
	}

	if rn.Set(0, -1).String() != "-1 to 0" {
		t.Errorf("Unexpected string output: %d ", rn)
	}

	if rn.Set(-10, 20).String() != "-10 to 20" {
		t.Errorf("Unexpected string output: %d ", rn)
	}
}
