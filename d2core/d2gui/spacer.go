package d2gui

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

func (s *SpacerStatic) getSize() (int, int) {
	return s.width, s.height
}

type SpacerDynamic struct {
	widgetBase
}

func createSpacerDynamic() *SpacerDynamic {
	spacer := &SpacerDynamic{}
	spacer.SetVisible(true)
	spacer.SetExpanding(true)

	return spacer
}
