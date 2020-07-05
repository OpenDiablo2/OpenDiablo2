package d2interface

// App is the OpenDiablo2 application
// As the app starts, it creates all of the AppComponents
// and binds to them.
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

// AppComponent is part of the app
// components are first bound, then initialized.
// This ensures that component dependencies don't make the app explode
type AppComponent interface {
	BindApp(App) error
	Initialize() error
}
