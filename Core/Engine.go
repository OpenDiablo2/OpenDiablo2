package Core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"strings"
	"sync"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/PaletteDefs"

	"github.com/OpenDiablo2/OpenDiablo2/Sound"

	"github.com/OpenDiablo2/OpenDiablo2/Common"
	"github.com/OpenDiablo2/OpenDiablo2/MPQ"
	"github.com/OpenDiablo2/OpenDiablo2/ResourcePaths"
	"github.com/OpenDiablo2/OpenDiablo2/Scenes"
	"github.com/OpenDiablo2/OpenDiablo2/UI"

	"github.com/hajimehoshi/ebiten"
	"github.com/mitchellh/go-homedir"
)

// EngineConfig defines the configuration for the engine, loaded from config.json
type EngineConfig struct {
	Language        string
	FullScreen      bool
	Scale           float64
	RunInBackground bool
	TicksPerSecond  int
	FpsCap          int
	VsyncEnabled    bool
	MpqPath         string
	MpqLoadOrder    []string
	SfxVolume       float64
	BgmVolume       float64
}

// Engine is the core OpenDiablo2 engine
type Engine struct {
	Settings        *EngineConfig                    // Engine configuration settings from json file
	Files           map[string]*Common.MpqFileRecord // Map that defines which files are in which MPQs
	CheckedPatch    map[string]bool					 // First time we check a file, we'll check if it's in the patch. This notes that we've already checked that.
	LoadingSprite   *Common.Sprite                   // The sprite shown when loading stuff
	loadingProgress float64                          // LoadingProcess is a range between 0.0 and 1.0. If set, loading screen displays.
	stepLoadingSize float64                          // The size for each loading step
	CurrentScene    Scenes.Scene                     // The current scene being rendered
	UIManager       *UI.Manager                      // The UI manager
	SoundManager    *Sound.Manager                   // The sound manager
	nextScene       Scenes.Scene                     // The next scene to be loaded at the end of the game loop
	fullscreenKey   bool                             // When true, the fullscreen toggle is still being pressed
}

// CreateEngine creates and instance of the OpenDiablo2 engine
func CreateEngine() *Engine {
	result := &Engine{
		CurrentScene: nil,
		nextScene:    nil,
	}
	result.loadConfigurationFile()
	ResourcePaths.LanguageCode = result.Settings.Language
	result.mapMpqFiles()
	Common.LoadPalettes(result.Files, result)
	Common.LoadTextDictionary(result)
	Common.LoadLevelTypes(result)
	Common.LoadLevelPresets(result)
	Common.LoadLevelWarps(result)
	Common.LoadObjectTypes(result)
	Common.LoadObjects(result)
	Common.LoadMissiles(result)
	Common.LoadSounds(result)
	result.SoundManager = Sound.CreateManager(result)
	result.SoundManager.SetVolumes(result.Settings.BgmVolume, result.Settings.SfxVolume)
	result.UIManager = UI.CreateManager(result, *result.SoundManager)
	result.LoadingSprite = result.LoadSprite(ResourcePaths.LoadingScreen, PaletteDefs.Loading)
	loadingSpriteSizeX, loadingSpriteSizeY := result.LoadingSprite.GetSize()
	result.LoadingSprite.MoveTo(int(400-(loadingSpriteSizeX/2)), int(300+(loadingSpriteSizeY/2)))
	result.SetNextScene(Scenes.CreateMainMenu(result, result, result.UIManager, result.SoundManager))
	//result.SetNextScene(Scenes.CreateBlizzardIntro(result, result))
	return result
}

func (v *Engine) loadConfigurationFile() {
	log.Println("Loading configuration file")
	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config EngineConfig

	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		log.Fatal(err)
	}
	v.Settings = &config
	// Path fixup for wine-installed diablo 2 in linux
	if v.Settings.MpqPath[0] != '/' {
		if _, err := os.Stat(v.Settings.MpqPath); os.IsNotExist(err) {
			homeDir, _ := homedir.Dir()
			newPath := strings.ReplaceAll(v.Settings.MpqPath, `C:\`, homeDir+"/.wine/drive_c/")
			newPath = strings.ReplaceAll(newPath, "C:/", homeDir+"/.wine/drive_c/")
			newPath = strings.ReplaceAll(newPath, `\`, "/")
			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				log.Printf("Detected linux wine installation, path updated to wine prefix path.")
				v.Settings.MpqPath = newPath
			}
		}
	}
}

func (v *Engine) mapMpqFiles() {
	log.Println("mapping mpq file structure")
	v.Files = make(map[string]*Common.MpqFileRecord)
	v.CheckedPatch = make(map[string]bool)
	for _, mpqFileName := range v.Settings.MpqLoadOrder {
		mpqPath := path.Join(v.Settings.MpqPath, mpqFileName)
		mpq, err := MPQ.Load(mpqPath)

		if err != nil {
			log.Fatal(err)
		}
		fileListText, err := mpq.ReadFile("(listfile)")
		if err != nil || fileListText == nil {
			// Super secret patch file activate!
			continue
		}
		fileList := strings.Split(string(fileListText), "\r\n")

		for _, filePath := range fileList {
			transFilePath := `/`+strings.ReplaceAll(strings.ToLower(filePath), `\`, `/`)
			if _, exists := v.Files[transFilePath]; exists {
				if v.Files[transFilePath].IsPatch {
					v.Files[transFilePath].UnpatchedMpqFile = mpqPath
				}
				continue
			}
			v.Files[transFilePath] = &Common.MpqFileRecord{
				mpqPath, false, ""}
			v.CheckedPatch[transFilePath] = false
		}
	}
}

var mutex sync.Mutex

// LoadFile loads a file from the specified mpq and returns the data as a byte array
func (v *Engine) LoadFile(fileName string) []byte {
	fileName = strings.ReplaceAll(fileName, "{LANG}", ResourcePaths.LanguageCode)
	fileName = strings.ReplaceAll(fileName, `\`, `/`)
	var mpqLookupFileName string
	if strings.HasPrefix(fileName, "/") || strings.HasPrefix(fileName,"\\") {
		mpqLookupFileName = strings.ReplaceAll(fileName, `/`, `\`)[1:]
	} else {
		mpqLookupFileName = strings.ReplaceAll(fileName, `/`, `\`)
	}
	
	mutex.Lock()
	// TODO: May want to cache some things if performance becomes an issue
	mpqFile := v.Files[strings.ToLower(fileName)]
	var mpq MPQ.MPQ
	var err error

	// always try to load from patch first
	checked, checkok := v.CheckedPatch[strings.ToLower(fileName)]
	patchLoaded := false
	if !checked || !checkok {
		patchMpqFilePath := path.Join(v.Settings.MpqPath, v.Settings.MpqLoadOrder[0])
		mpq, err = MPQ.Load(patchMpqFilePath)
		if err == nil {
			// loaded patch mpq. check if this file exists in it
			fileInPatch := mpq.FileExists(mpqLookupFileName)
			if fileInPatch {
				patchLoaded = true
				// set the path to the patch so it will be loaded there in the future
				mpqFile = &Common.MpqFileRecord{patchMpqFilePath, false, ""}
				v.Files[strings.ToLower(fileName)] = mpqFile
			}
		}
		v.CheckedPatch[strings.ToLower(fileName)] = true
	}
	
	if patchLoaded {
		// if we already loaded the correct mpq from the patch check, don't bother reloading it
	} else if mpqFile == nil {
		// Super secret non-listed file?
		found := false
		for _, mpqFile := range v.Settings.MpqLoadOrder {
			mpqFilePath := path.Join(v.Settings.MpqPath, mpqFile)
			mpq, err = MPQ.Load(mpqFilePath)
			if err != nil {
				continue
			}
			if !mpq.FileExists(fileName) {
				continue
			}
			// We found the super-secret file!
			found = true
			v.Files[strings.ToLower(fileName)] = &Common.MpqFileRecord{mpqFilePath, false, ""}
			break
		}
		if !found {
			log.Fatal(fmt.Sprintf("File '%s' not found during preload of listfiles, and could not be located in any MPQ checking manually.", fileName))
		}
	} else if mpqFile.IsPatch {
		log.Fatal("Tried to load a patchfile")
	} else {
		mpq, err = MPQ.Load(mpqFile.MpqFile)
		if err != nil {
			log.Printf("Error loading file '%s'", fileName)
			log.Fatal(err)
		}
	}

	blockTableEntry, err := mpq.GetFileBlockData(mpqLookupFileName)
	if err != nil {
		log.Printf("Error locating block data entry for '%s' in mpq file '%s'", mpqLookupFileName, mpq.FileName)
		log.Fatal(err)
	}
	mpqStream := MPQ.CreateStream(mpq, blockTableEntry, mpqLookupFileName)
	result := make([]byte, blockTableEntry.UncompressedFileSize)
	mpqStream.Read(result, 0, blockTableEntry.UncompressedFileSize)
	mutex.Unlock()
	return result
}

// IsLoading returns true if the engine is currently in a loading state
func (v *Engine) IsLoading() bool {
	return v.loadingProgress < 1.0
}

// LoadSprite loads a sprite from the game's data files
func (v *Engine) LoadSprite(fileName string, palette PaletteDefs.PaletteType) *Common.Sprite {
	data := v.LoadFile(fileName)
	sprite := Common.CreateSprite(data, Common.Palettes[palette])
	return sprite
}

// updateScene handles the scene maintenance for the engine
func (v *Engine) updateScene() {
	if v.nextScene == nil {
		return
	}
	if v.CurrentScene != nil {
		v.CurrentScene.Unload()
	}
	v.CurrentScene = v.nextScene
	v.nextScene = nil
	v.UIManager.Reset()
	thingsToLoad := v.CurrentScene.Load()
	v.SetLoadingStepSize(1.0 / float64(len(thingsToLoad)))
	v.ResetLoading()
	go func() {
		for _, f := range thingsToLoad {
			f()
			v.StepLoading()
		}
		v.FinishLoading()
	}()
}

// Update updates the internal state of the engine
func (v *Engine) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyAlt) && ebiten.IsKeyPressed(ebiten.KeyEnter) {
		if !v.fullscreenKey {
			ebiten.SetFullscreen(!ebiten.IsFullscreen())
		}
		v.fullscreenKey = true
	} else {
		v.fullscreenKey = false
	}

	v.updateScene()
	if v.CurrentScene == nil {
		log.Fatal("no scene loaded")
	}

	if v.IsLoading() {
		return
	}

	v.CurrentScene.Update(float64(1) / ebiten.CurrentTPS())
	v.UIManager.Update()
}

// Draw draws the game
func (v *Engine) Draw(screen *ebiten.Image) {
	if v.loadingProgress < 1.0 {
		v.LoadingSprite.Frame = uint8(Common.Max(0, Common.Min(uint32(len(v.LoadingSprite.Frames)-1), uint32(float64(len(v.LoadingSprite.Frames)-1)*v.loadingProgress))))
		v.LoadingSprite.Draw(screen)
	} else {
		if v.CurrentScene == nil {
			log.Fatal("no scene loaded")
		}
		v.CurrentScene.Render(screen)
		v.UIManager.Draw(screen)
	}
}

// SetNextScene tells the engine what scene to load on the next update cycle
func (v *Engine) SetNextScene(nextScene Scenes.Scene) {
	v.nextScene = nextScene
}

// SetLoadingStepSize sets the size of the loading step
func (v *Engine) SetLoadingStepSize(size float64) {
	v.stepLoadingSize = size
}

// ResetLoading resets the loading progress
func (v *Engine) ResetLoading() {
	v.loadingProgress = 0.0
}

// StepLoading increments the loading progress
func (v *Engine) StepLoading() {
	v.loadingProgress = math.Min(1.0, v.loadingProgress+v.stepLoadingSize)
}

// FinishLoading terminates the loading phase
func (v *Engine) FinishLoading() {
	v.loadingProgress = 1.0
}
