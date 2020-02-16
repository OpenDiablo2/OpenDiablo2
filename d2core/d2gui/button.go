package d2gui

// // Button defines a standard wide UI button
// type Button struct {
// 	enabled               bool
// 	x, y                  int
// 	width, height         int
// 	visible               bool
// 	pressed               bool
// 	toggled               bool
// 	normalSurface         d2render.Surface
// 	pressedSurface        d2render.Surface
// 	toggledSurface        d2render.Surface
// 	pressedToggledSurface d2render.Surface
// 	disabledSurface       d2render.Surface
// 	buttonLayout          ButtonLayout
// 	onClick               func()
// }
//
// // CreateButton creates an instance of Button
// func CreateButton(buttonType ButtonType, text string) Button {
// 	result := Button{
// 		width:   0,
// 		height:  0,
// 		visible: true,
// 		enabled: true,
// 		pressed: false,
// 	}
// 	buttonLayout := ButtonLayouts[buttonType]
// 	result.buttonLayout = buttonLayout
// 	font := GetFont(buttonLayout.FontPath, d2resource.PaletteUnits)
//
// 	animation, _ := d2asset.LoadAnimation(buttonLayout.ResourceName, buttonLayout.PaletteName)
// 	buttonSprite, _ := LoadSprite(animation)
// 	totalButtonTypes := buttonSprite.GetFrameCount() / (buttonLayout.XSegments * buttonLayout.YSegments)
// 	for i := 0; i < buttonLayout.XSegments; i++ {
// 		w, _, _ := buttonSprite.GetFrameSize(i)
// 		result.width += w
// 	}
// 	for i := 0; i < buttonLayout.YSegments; i++ {
// 		_, h, _ := buttonSprite.GetFrameSize(i * buttonLayout.YSegments)
// 		result.height += h
// 	}
//
// 	result.normalSurface, _ = d2render.NewSurface(result.width, result.height, d2render.FilterNearest)
// 	_, fontHeight := font.GetTextMetrics(text)
// 	textY := (result.height / 2) - (fontHeight / 2) + buttonLayout.TextOffset
//
// 	buttonSprite.SetPosition(0, 0)
// 	buttonSprite.SetBlend(true)
// 	buttonSprite.RenderSegmented(result.normalSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame)
// 	font.Render(0, textY, text, color.RGBA{R: 100, G: 100, B: 100, A: 255}, result.normalSurface)
// 	if buttonLayout.AllowFrameChange {
// 		if totalButtonTypes > 1 {
// 			result.pressedSurface, _ = d2render.NewSurface(result.width, result.height, d2render.FilterNearest)
// 			buttonSprite.RenderSegmented(result.pressedSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+1)
// 			font.Render(-2, textY+2, text, color.RGBA{R: 100, G: 100, B: 100, A: 255}, result.pressedSurface)
// 		}
// 		if totalButtonTypes > 2 {
// 			result.toggledSurface, _ = d2render.NewSurface(result.width, result.height, d2render.FilterNearest)
// 			buttonSprite.RenderSegmented(result.toggledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+2)
// 			font.Render(0, textY, text, color.RGBA{R: 100, G: 100, B: 100, A: 255}, result.toggledSurface)
// 		}
// 		if totalButtonTypes > 3 {
// 			result.pressedToggledSurface, _ = d2render.NewSurface(result.width, result.height, d2render.FilterNearest)
// 			buttonSprite.RenderSegmented(result.pressedToggledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.BaseFrame+3)
// 			font.Render(0, textY, text, color.RGBA{R: 100, G: 100, B: 100, A: 255}, result.pressedToggledSurface)
// 		}
// 		if buttonLayout.DisabledFrame != -1 {
// 			result.disabledSurface, _ = d2render.NewSurface(result.width, result.height, d2render.FilterNearest)
// 			buttonSprite.RenderSegmented(result.disabledSurface, buttonLayout.XSegments, buttonLayout.YSegments, buttonLayout.DisabledFrame)
// 			font.Render(0, textY, text, color.RGBA{R: 100, G: 100, B: 100, A: 255}, result.disabledSurface)
// 		}
// 	}
// 	return result
// }
//
// // OnActivated defines the callback handler for the activate event
// func (v *Button) OnActivated(callback func()) {
// 	v.onClick = callback
// }
//
// // Activate calls the on activated callback handler, if any
// func (v *Button) Activate() {
// 	if v.onClick == nil {
// 		return
// 	}
// 	v.onClick()
// }
//
// // Render renders the button
// func (v *Button) Render(target d2render.Surface) {
// 	target.PushCompositeMode(d2render.CompositeModeSourceAtop)
// 	target.PushFilter(d2render.FilterNearest)
// 	target.PushTranslation(v.x, v.y)
// 	defer target.PopN(3)
//
// 	if !v.enabled {
// 		target.PushColor(color.RGBA{R: 128, G: 128, B: 128, A: 195})
// 		defer target.Pop()
// 		target.Render(v.disabledSurface)
// 	} else if v.toggled && v.pressed {
// 		target.Render(v.pressedToggledSurface)
// 	} else if v.pressed {
// 		target.Render(v.pressedSurface)
// 	} else if v.toggled {
// 		target.Render(v.toggledSurface)
// 	} else {
// 		target.Render(v.normalSurface)
// 	}
// }
//
// func (v *Button) Advance(elapsed float64) {
//
// }
//
// // GetEnabled returns the enabled state
// func (v *Button) GetEnabled() bool {
// 	return v.enabled
// }
//
// // SetEnabled sets the enabled state
// func (v *Button) SetEnabled(enabled bool) {
// 	v.enabled = enabled
// }
//
// // GetSize returns the size of the button
// func (v *Button) GetSize() (int, int) {
// 	return v.width, v.height
// }
//
// // SetPosition moves the button
// func (v *Button) SetPosition(x, y int) {
// 	v.x = x
// 	v.y = y
// }
//
// // GetPosition returns the location of the button
// func (v *Button) GetPosition() (x, y int) {
// 	return v.x, v.y
// }
//
// // GetVisible returns the visibility of the button
// func (v *Button) GetVisible() bool {
// 	return v.visible
// }
//
// // SetVisible sets the visibility of the button
// func (v *Button) SetVisible(visible bool) {
// 	v.visible = visible
// }
//
// // GetPressed returns the pressed state of the button
// func (v *Button) GetPressed() bool {
// 	return v.pressed
// }
//
// // SetPressed sets the pressed state of the button
// func (v *Button) SetPressed(pressed bool) {
// 	v.pressed = pressed
// }
