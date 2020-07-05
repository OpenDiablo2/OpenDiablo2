package d2interface

// App is the OpenDiablo2 application. It creates all of the AppComponents
// and creates two-way references for dependency resolution.
type App interface {
	Run()
	BindAppComponent(AppComponent) error

	// The AppComponents
	Input() (InputManager, error)
	Audio() (AudioProvider, error)
	Renderer() (Renderer, error)
	Terminal() (Terminal, error)
	Asset() (AssetManager, error)
}

// AppComponent defines a high-level part of the app.
// AppComponents are bound and initialized by the App.
// This ensures that component dependencies don't make the app explode.
type AppComponent interface {
	BindApp(App) error
	Initialize() error
}
