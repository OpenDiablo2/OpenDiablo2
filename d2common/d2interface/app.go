package d2interface

// App is the core OpenDiablo2 interface. It is comprised of AppComponents which
// all work together, but the App governs the order in which they operate.
type App interface {
	Advanceable
	Renderable
	BindComponent(component AppComponent, key string) error
	UnbindComponent(string) error
	GetComponent(string) (AppComponent, error)
	Run()
}
