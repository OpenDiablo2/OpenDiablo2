package d2ds1

import "testing"

func Test_layers(t *testing.T) {
	const (
		fmtWidthHeightError = "unexpected wall layer width/height: %dx%d"
	)

	l := &layer{}

	l.SetSize(0, 0)

	if l.Width() != 1 || l.Height() != 1 {
		t.Fatalf(fmtWidthHeightError, l.Width(), l.Height())
	}

	l.SetSize(4, 5)

	if l.Width() != 4 || l.Height() != 5 {
		t.Fatalf(fmtWidthHeightError, l.Width(), l.Height())
	}

	l.SetSize(4, 3)

	if l.Width() != 4 || l.Height() != 3 {
		t.Fatalf(fmtWidthHeightError, l.Width(), l.Height())
	}
}
