package d2components

import (
	"os/user"
	"path"
	"runtime"

	"github.com/gravestench/akara"
)

// static check that GameConfigComponent implements Component
var _ akara.Component = &GameConfigComponent{}

// static check that GameConfigMap implements ComponentMap
var _ akara.ComponentMap = &GameConfigMap{}

type GameConfigComponent struct {
	MpqLoadOrder    []string
	Language        string
	MpqPath         string
	TicksPerSecond  int
	FpsCap          int
	SfxVolume       float64
	BgmVolume       float64
	FullScreen      bool
	RunInBackground bool
	VsyncEnabled    bool
	Backend         string
}

// ID returns a unique identifier for the component type
func (*GameConfigComponent) ID() akara.ComponentID {
	return GameConfigCID
}

// NewMap returns a new component map for the component type
func (*GameConfigComponent) NewMap() akara.ComponentMap {
	return NewGameConfigMap()
}

// GameConfig is a convenient reference to be used as a component identifier
var GameConfig = (*GameConfigComponent)(nil) // nolint:gochecknoglobals // global by design

// NewGameConfigMap creates a new map of entity ID's to GameConfigComponent components
func NewGameConfigMap() *GameConfigMap {
	cm := &GameConfigMap{
		components: make(map[akara.EID]*GameConfigComponent),
	}

	return cm
}

// GameConfigMap is a map of entity ID's to GameConfigComponent components
type GameConfigMap struct {
	world      *akara.World
	components map[akara.EID]*GameConfigComponent
}

// Init initializes the component map with the given world
func (cm *GameConfigMap) Init(world *akara.World) {
	cm.world = world
}

// Add a new GameConfigComponent for the given entity id, return that component.
// If the entity already has a component, just return that one.
func (cm *GameConfigMap) Add(id akara.EID) akara.Component {
	if com, has := cm.components[id]; has {
		return com
	}

	cm.components[id] = defaultConfig()

	cm.world.UpdateEntity(id)

	return cm.components[id]
}

// ID returns a unique identifier for the component type
func (*GameConfigMap) ID() akara.ComponentID {
	return GameConfigCID
}

// NewMap returns a new component map for the component type
func (*GameConfigMap) NewMap() akara.ComponentMap {
	return NewGameConfigMap()
}

// AddGameConfig adds a new GameConfigComponent for the given entity id and returns it.
// If the entity already has a GameConfigComponent component, just return that one.
// this is a convenience method for the generic Add method, as it returns a
// *GameConfigComponent instead of an akara.Component
func (cm *GameConfigMap) AddGameConfig(id akara.EID) *GameConfigComponent {
	return cm.Add(id).(*GameConfigComponent)
}

// Get returns the component associated with the given entity id
func (cm *GameConfigMap) Get(id akara.EID) (akara.Component, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// GetGameConfig returns the GameConfigComponent component associated with the given entity id
func (cm *GameConfigMap) GetGameConfig(id akara.EID) (*GameConfigComponent, bool) {
	entry, found := cm.components[id]
	return entry, found
}

// Remove a component for the given entity id, return the component.
func (cm *GameConfigMap) Remove(id akara.EID) {
	delete(cm.components, id)
	cm.world.UpdateEntity(id)
}

func defaultConfig() *GameConfigComponent {
	const (
		defaultSfxVolume = 1.0
		defaultBgmVolume = 0.3
	)

	config := &GameConfigComponent{
		Language:        "ENG",
		FullScreen:      false,
		TicksPerSecond:  -1,
		RunInBackground: true,
		VsyncEnabled:    true,
		SfxVolume:       defaultSfxVolume,
		BgmVolume:       defaultBgmVolume,
		MpqPath:         "C:/Program Files (x86)/Diablo II",
		Backend:         "Ebiten",
		MpqLoadOrder: []string{
			"Patch_D2.mpq",
			"d2exp.mpq",
			"d2xmusic.mpq",
			"d2xtalk.mpq",
			"d2xvideo.mpq",
			"d2data.mpq",
			"d2char.mpq",
			"d2music.mpq",
			"d2sfx.mpq",
			"d2video.mpq",
			"d2speech.mpq",
		},
	}

	switch runtime.GOOS {
	case "windows":
		if runtime.GOARCH == "386" {
			config.MpqPath = "C:/Program Files/Diablo II"
		}
	case "darwin":
		config.MpqPath = "/Applications/Diablo II/"
		config.MpqLoadOrder = []string{
			"Diablo II Patch",
			"Diablo II Expansion Data",
			"Diablo II Expansion Movies",
			"Diablo II Expansion Music",
			"Diablo II Expansion Speech",
			"Diablo II Game Data",
			"Diablo II Graphics",
			"Diablo II Movies",
			"Diablo II Music",
			"Diablo II Sounds",
			"Diablo II Speech",
		}
	case "linux":
		if usr, err := user.Current(); err == nil {
			config.MpqPath = path.Join(usr.HomeDir, ".wine/drive_c/Program Files (x86)/Diablo II")
		}
	}

	return config
}
