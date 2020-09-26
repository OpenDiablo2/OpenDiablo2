package d2gamescreen

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

const hideZoneTextAfterSeconds = 2.0

const (
	moveErrStr         = "failed to send MovePlayer packet to the server, playerId: %s, x: %g, x: %g\n"
	bindControlsErrStr = "failed to add gameControls as input handler for player: %s\n"
	castErrStr         = "failed to send CastSkill packet to the server, playerId: %s, skillId: %d, x: %g, x: %g\n"
	spawnItemErrStr    = "failed to send SpawnItem packet to the server: (%d, %d) %+v"
)

// Game represents the Gameplay screen
type Game struct {
	*d2mapentity.MapEntityFactory
	asset                *d2asset.AssetManager
	gameClient           *d2client.GameClient
	mapRenderer          *d2maprenderer.MapRenderer
	uiManager            *d2ui.UIManager
	gameControls         *d2player.GameControls // TODO: Hack
	localPlayer          *d2mapentity.Player
	lastRegionType       d2enum.RegionIdType
	ticksSinceLevelCheck float64
	escapeMenu           *EscapeMenu
	soundEngine          *d2audio.SoundEngine
	soundEnv             d2audio.SoundEnvironment
	guiManager           *d2gui.GuiManager

	renderer      d2interface.Renderer
	inputManager  d2interface.InputManager
	audioProvider d2interface.AudioProvider
	terminal      d2interface.Terminal
}

// CreateGame creates the Gameplay screen and returns a pointer to it
func CreateGame(
	navigator Navigator,
	asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	gameClient *d2client.GameClient,
	term d2interface.Terminal,
	guiManager *d2gui.GuiManager,
) *Game {
	// find the local player and its initial location
	var startX, startY float64

	for _, player := range gameClient.Players {
		if player.ID() != gameClient.PlayerID {
			continue
		}

		worldPosition := player.Position.World()
		startX, startY = worldPosition.X(), worldPosition.Y()

		break
	}

	result := &Game{
		asset:                asset,
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastRegionType:       d2enum.RegionNone,
		ticksSinceLevelCheck: 0,
		mapRenderer: d2maprenderer.CreateMapRenderer(asset, renderer,
			gameClient.MapEngine, term, startX, startY),
		escapeMenu:    NewEscapeMenu(navigator, renderer, audioProvider, guiManager, asset),
		inputManager:  inputManager,
		audioProvider: audioProvider,
		renderer:      renderer,
		terminal:      term,
		soundEngine:   d2audio.NewSoundEngine(audioProvider, asset, term),
		uiManager:     ui,
		guiManager:    guiManager,
	}
	result.soundEnv = d2audio.NewSoundEnvironment(result.soundEngine)

	result.escapeMenu.onLoad()

	if err := inputManager.BindHandler(result.escapeMenu); err != nil {
		fmt.Println("failed to add gameplay screen as event handler")
	}

	return result
}

// OnLoad loads the resources for the Gameplay screen
func (v *Game) OnLoad(_ d2screen.LoadingState) {
	v.audioProvider.PlayBGM("")

	err := v.terminal.BindAction(
		"spawnitem",
		"spawns an item at the local player position",
		func(code1, code2, code3, code4, code5 string) {
			codes := []string{code1, code2, code3, code4, code5}
			v.debugSpawnItemAtPlayer(codes...)
		},
	)
	if err != nil {
		fmt.Printf("failed to bind the '%s' action, err: %v\n", "spawnitem", err)
	}

	err = v.terminal.BindAction(
		"spawnitemat",
		"spawns an item at the x,y coordinates",
		func(x, y int, code1, code2, code3, code4, code5 string) {
			codes := []string{code1, code2, code3, code4, code5}
			v.debugSpawnItemAtLocation(x, y, codes...)
		},
	)
	if err != nil {
		fmt.Printf("failed to bind the '%s' action, err: %v\n", "spawnitemat", err)
	}

	err = v.terminal.BindAction(
		"spawnmon",
		"spawn monster at the local player position",
		func(name string) {
			x := int(v.localPlayer.Position.X())
			y := int(v.localPlayer.Position.Y())
			monstat := v.asset.Records.Monster.Stats[name]
			if monstat == nil {
				v.terminal.OutputErrorf("no monstat entry for \"%s\"", name)
				return
			}

			monster, err := v.gameClient.MapEngine.NewNPC(x, y, monstat, 0)
			if err != nil {
				v.terminal.OutputErrorf("error generating monster \"%s\": %v", name, err)
				return
			}

			v.gameClient.MapEngine.AddEntity(monster)
		},
	)
	if err != nil {
		fmt.Printf("failed to bind the '%s' action, err: %v\n", "spawnmon", err)
	}
}

// OnUnload releases the resources of Gameplay screen
func (v *Game) OnUnload() error {
	if err := v.inputManager.UnbindHandler(v.gameControls); err != nil { // TODO: hack
		return err
	}

	if err := v.inputManager.UnbindHandler(v.escapeMenu); err != nil { // TODO: hack
		return err
	}

	if err := v.terminal.UnbindAction("spawnItemAt"); err != nil {
		return err
	}

	if err := v.terminal.UnbindAction("spawnItem"); err != nil {
		return err
	}

	if err := v.gameClient.Close(); err != nil {
		return err
	}

	v.soundEngine.Reset()

	return nil
}

// Render renders the Gameplay screen
func (v *Game) Render(screen d2interface.Surface) error {
	if v.gameClient.RegenMap {
		v.gameClient.RegenMap = false
		v.mapRenderer.RegenerateTileCache()
	}

	if err := screen.Clear(color.Black); err != nil {
		return err
	}

	v.mapRenderer.Render(screen)

	if v.gameControls != nil {
		if v.gameControls.HelpOverlay != nil && v.gameControls.HelpOverlay.IsOpen() {
			// When help overlay is open, put transparent black screen. Magic noumber is hex for RGBA.
			screen.DrawRect(800, 600, d2util.Color(0x0000007f))
		}
		if err := v.gameControls.Render(screen); err != nil {
			return err
		}
	}

	return nil
}

// Advance runs the update logic on the Gameplay screen
func (v *Game) Advance(elapsed float64) error {
	v.soundEngine.Advance(elapsed)

	if (v.escapeMenu != nil && !v.escapeMenu.isOpen) || len(v.gameClient.Players) != 1 {
		v.gameClient.MapEngine.Advance(elapsed) // TODO: Hack
	}

	if v.gameControls != nil {
		if err := v.gameControls.Advance(elapsed); err != nil {
			return err
		}
	}

	v.ticksSinceLevelCheck += elapsed
	if v.ticksSinceLevelCheck > 1.0 {
		v.ticksSinceLevelCheck = 0
		if v.localPlayer != nil {
			tilePosition := v.localPlayer.Position.Tile()
			tile := v.gameClient.MapEngine.TileAt(int(tilePosition.X()), int(tilePosition.Y()))

			if tile != nil {
				levelDetails := v.asset.Records.Level.Details[int(tile.RegionType)]
				v.soundEnv.SetEnv(levelDetails.SoundEnvironmentID)

				// skip showing zone change text the first time we enter the world
				if v.lastRegionType != d2enum.RegionNone && v.lastRegionType != tile.RegionType {
					//TODO: Should not be using RegionType as an index - this will return incorrect LevelDetails record for most of the zones.
					areaName := levelDetails.LevelDisplayName
					areaChgStr := fmt.Sprintf("Entering The %s", areaName)
					v.gameControls.SetZoneChangeText(areaChgStr)
					v.gameControls.ShowZoneChangeText()
					v.gameControls.HideZoneChangeTextAfter(hideZoneTextAfterSeconds)
				}

				v.lastRegionType = tile.RegionType
			}
		}
	}

	// Bind the game controls to the player once it exists
	if v.gameControls == nil {
		if err := v.bindGameControls(); err != nil {
			return err
		}
	}

	// Update the camera to focus on the player
	if v.localPlayer != nil && !v.gameControls.FreeCam {
		worldPosition := v.localPlayer.Position.World()
		rx, ry := v.mapRenderer.WorldToOrtho(worldPosition.X(), worldPosition.Y())
		position := d2vector.NewPosition(rx, ry)
		v.mapRenderer.SetCameraTarget(&position)
	}

	v.soundEnv.Advance(elapsed)

	return nil
}

func (v *Game) bindGameControls() error {
	for _, player := range v.gameClient.Players {
		if player.ID() != v.gameClient.PlayerID {
			continue
		}

		v.localPlayer = player

		var err error
		v.gameControls, err = d2player.NewGameControls(v.asset, v.renderer, player, v.gameClient.MapEngine,
			v.mapRenderer, v, v.terminal, v.uiManager, v.guiManager, v.gameClient.IsSinglePlayer())

		if err != nil {
			return err
		}

		v.gameControls.Load()

		if err := v.inputManager.BindHandler(v.gameControls); err != nil {
			fmt.Printf(bindControlsErrStr, player.ID())
		}

		break
	}

	return nil
}

// OnPlayerMove sends the player move action to the server
func (v *Game) OnPlayerMove(targetX, targetY float64) {
	worldPosition := v.localPlayer.Position.World()

	playerID, worldX, worldY := v.gameClient.PlayerID, worldPosition.X(), worldPosition.Y()
	createPlayerPacket := d2netpacket.CreateMovePlayerPacket(playerID, worldX, worldY, targetX, targetY)
	err := v.gameClient.SendPacketToServer(createPlayerPacket)

	if err != nil {
		fmt.Printf(moveErrStr, v.gameClient.PlayerID, targetX, targetY)
	}
}

// OnPlayerCast sends the casting skill action to the server
func (v *Game) OnPlayerCast(skillID int, targetX, targetY float64) {
	err := v.gameClient.SendPacketToServer(d2netpacket.CreateCastPacket(v.gameClient.PlayerID, skillID, targetX, targetY))
	if err != nil {
		fmt.Printf(castErrStr, v.gameClient.PlayerID, skillID, targetX, targetY)
	}
}

func (v *Game) debugSpawnItemAtPlayer(codes ...string) {
	if v.localPlayer == nil {
		return
	}

	pos := v.localPlayer.GetPosition()
	tile := pos.Tile()
	x, y := int(tile.X()), int(tile.Y())

	v.debugSpawnItemAtLocation(x, y, codes...)
}

func (v *Game) debugSpawnItemAtLocation(x, y int, codes ...string) {
	packet := d2netpacket.CreateSpawnItemPacket(x, y, codes...)

	err := v.gameClient.SendPacketToServer(packet)
	if err != nil {
		fmt.Printf(spawnItemErrStr, x, y, codes)
	}
}
