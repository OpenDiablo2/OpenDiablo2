package d2scene

import (
	"math"
	"runtime"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// Scene defines the function necessary for scene management
type Scene interface {
	Load() []func()
	Unload()
	Render(target d2render.Surface)
	Advance(tickTime float64)
}

var nextScene Scene         // The next scene to be loaded at the end of the game loop
var currentScene Scene      // The current scene being rendered
var loadingIndex int        // Determines which load function is currently being called
var thingsToLoad []func()   // The load functions for the next scene
var loadingProgress float64 // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
var stepLoadingSize float64 // The size for each loading step

// SetNextScene tells the engine what scene to load on the next update cycle
func SetNextScene(scene Scene) {
	nextScene = scene
}

func GetCurrentScene() Scene {
	return currentScene
}

// updateScene handles the scene maintenance for the engine
func UpdateScene() {
	if nextScene == nil {
		if thingsToLoad != nil {
			if loadingIndex < len(thingsToLoad) {
				thingsToLoad[loadingIndex]()
				loadingIndex++
				if loadingIndex < len(thingsToLoad) {
					StepLoading()
				} else {
					FinishLoading()
					thingsToLoad = nil
				}
				return
			}
		}
		return
	}
	if currentScene != nil {
		currentScene.Unload()
		runtime.GC()
	}
	currentScene = nextScene
	nextScene = nil
	d2ui.Reset()
	thingsToLoad = currentScene.Load()
	loadingIndex = 0
	SetLoadingStepSize(1.0 / float64(len(thingsToLoad)))
	ResetLoading()
}

func Advance(time float64) {
	if currentScene == nil {
		return
	}
	currentScene.Advance(time)
}

func Render(surface d2render.Surface) {
	if currentScene == nil {
		return
	}
	currentScene.Render(surface)
}

// SetLoadingStepSize sets the size of the loading step
func SetLoadingStepSize(size float64) {
	stepLoadingSize = size
}

// ResetLoading resets the loading progress
func ResetLoading() {
	loadingProgress = 0.0
}

// StepLoading increments the loading progress
func StepLoading() {
	loadingProgress = math.Min(1.0, loadingProgress+stepLoadingSize)
}

// FinishLoading terminates the loading phase
func FinishLoading() {
	loadingProgress = 1.0
}

// IsLoading returns true if the engine is currently in a loading state
func IsLoading() bool {
	return loadingProgress < 1.0
}

func GetLoadingProgress() float64 {
	return loadingProgress
}
