package d2gui

// SpacerStatic is a spacer with explicit width and height, meaning
// that it wont dynamically expand within a layout
type SpacerStatic struct {
	widgetBase

	width  int
	height int
}

func createSpacerStatic(width, height int) *SpacerStatic {
	spacer := &SpacerStatic{width: width, height: height}
	spacer.SetVisible(true)

	return spacer
}

func (s *SpacerStatic) getSize() (width, height int) {
	return s.width, s.height
}

// SpacerDynamic is a spacer that will expand within a layout,
// depending on the layout position and alignment types
type SpacerDynamic struct {
	widgetBase
}

func createSpacerDynamic() *SpacerDynamic {
	spacer := &SpacerDynamic{}
	spacer.SetVisible(true)
	spacer.SetExpanding(true)

	return spacer
}
