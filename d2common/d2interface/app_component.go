package d2interface

// AppComponent is something that has a unique name and binds to the App.
// The App is responsible for calling the AppComponent's Advance and Render
// methods. Some components have functionality that depend on the presence of
// other components, so components will contain a reference to the App that
// governs them.
type AppComponent interface {
	Advanceable
	Renderable
	Initializer
	BindApp(app App) error
	UnbindApp(app App) error
}
