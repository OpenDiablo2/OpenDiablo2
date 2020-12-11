package d2gamescreen

import (
	"errors"
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

const (
	black50alpha = 0x0000007f // rgba
)

// CreateGame creates the Gameplay screen and returns a pointer to it
func CreateGame(
	navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	ui *d2ui.UIManager,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	gameClient *d2client.GameClient,
	term d2interface.Terminal,
	l d2util.LogLevel,
	guiManager *d2gui.GuiManager,
) (*Game, error) {
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

	keyMap := d2player.GetDefaultKeyMap(asset)

	game := &Game{
		asset:                asset,
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastRegionType:       d2enum.RegionNone,
		ticksSinceLevelCheck: 0,
		mapRenderer: d2maprenderer.CreateMapRenderer(asset, renderer,
			gameClient.MapEngine, term, l, startX, startY),
		escapeMenu:    d2player.NewEscapeMenu(navigator, renderer, audioProvider, ui, guiManager, asset, l, keyMap),
		inputManager:  inputManager,
		audioProvider: audioProvider,
		renderer:      renderer,
		terminal:      term,
		soundEngine:   d2audio.NewSoundEngine(audioProvider, asset, l, term),
		uiManager:     ui,
		guiManager:    guiManager,
		keyMap:        keyMap,
		logLevel:      l,
	}
	game.Logger = d2util.NewLogger()
	game.Logger.SetLevel(l)
	game.Logger.SetPrefix(logPrefix)

	game.soundEnv = d2audio.NewSoundEnvironment(game.soundEngine)

	game.escapeMenu.OnLoad()

	if err := inputManager.BindHandler(game.escapeMenu); err != nil {
		return nil, errors.New("failed to add gameplay screen as event handler")
	}

	return game, nil
}

// Game represents the Gameplay screen
type Game struct {
	*d2mapentity.MapEntityFactory
	asset                *d2asset.AssetManager
	gameClient           *d2client.GameClient
	mapRenderer          *d2maprenderer.MapRenderer
	uiManager            *d2ui.UIManager
	gameControls         *d2player.GameControls
	localPlayer          *d2mapentity.Player
	lastRegionType       d2enum.RegionIdType
	ticksSinceLevelCheck float64
	escapeMenu           *d2player.EscapeMenu
	soundEngine          *d2audio.SoundEngine
	soundEnv             d2audio.SoundEnvironment
	guiManager           *d2gui.GuiManager
	keyMap               *d2player.KeyMap

	renderer      d2interface.Renderer
	inputManager  d2interface.InputManager
	audioProvider d2interface.AudioProvider
	terminal      d2interface.Terminal

	*d2util.Logger
	logLevel d2util.LogLevel
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
		v.Errorf("failed to bind the '%s' action, err: %v\n", "spawnitem", err)
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
		v.Errorf("failed to bind the '%s' action, err: %v\n", "spawnitemat", err)
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

			monster, npcErr := v.gameClient.MapEngine.NewNPC(x, y, monstat, 0)
			if npcErr != nil {
				v.terminal.OutputErrorf("error generating monster \"%s\": %v", name, npcErr)
				return
			}

			v.gameClient.MapEngine.AddEntity(monster)
		},
	)
	if err != nil {
		v.Errorf("failed to bind the '%s' action, err: %v\n", "spawnmon", err)
	}
}

// OnUnload releases the resources of Gameplay screen
func (v *Game) OnUnload() error {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/792
	if err := v.inputManager.UnbindHandler(v.gameControls); err != nil {
		return err
	}

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/792
	if err := v.inputManager.UnbindHandler(v.escapeMenu); err != nil {
		return err
	}

	if err := v.terminal.UnbindAction("spawnItemAt"); err != nil {
		return err
	}

	if err := v.terminal.UnbindAction("spawnItem"); err != nil {
		return err
	}

	if err := v.OnPlayerSave(); err != nil {
		return err
	}

	if err := v.gameClient.Close(); err != nil {
		return err
	}

	v.soundEngine.Reset()

	return nil
}

// Render renders the Gameplay screen
func (v *Game) Render(screen d2interface.Surface) {
	if v.gameClient.RegenMap {
		v.gameClient.RegenMap = false
		v.mapRenderer.RegenerateTileCache()
		v.gameClient.MapEngine.IsLoading = false
	}

	screen.Clear(color.Black)
	v.mapRenderer.Render(screen)

	if v.gameControls != nil {
		if v.gameControls.HelpOverlay != nil && v.gameControls.HelpOverlay.IsOpen() {
			screen.DrawRect(screenWidth, screenHeight, d2util.Color(black50alpha))
		}

		if err := v.gameControls.Render(screen); err != nil {
			return
		}
	}
}

// Advance runs the update logic on the Gameplay screen
func (v *Game) Advance(elapsed float64) error {
	v.soundEngine.Advance(elapsed)

	if (v.escapeMenu != nil && !v.escapeMenu.IsOpen()) || len(v.gameClient.Players) != 1 {
		v.gameClient.MapEngine.Advance(elapsed)
	}

	if v.gameControls != nil {
		if err := v.gameControls.Advance(elapsed); err != nil {
			return err
		}
	}

	v.ticksSinceLevelCheck += elapsed
	if v.ticksSinceLevelCheck > 1 {
		v.ticksSinceLevelCheck = 0
		if v.localPlayer != nil {
			tilePosition := v.localPlayer.Position.Tile()
			tile := v.gameClient.MapEngine.TileAt(int(tilePosition.X()), int(tilePosition.Y()))

			if tile != nil {
				levelDetails := v.asset.Records.Level.Details[int(tile.RegionType)]
				v.soundEnv.SetEnv(levelDetails.SoundEnvironmentID)

				// skip showing zone change text the first time we enter the world
				if v.lastRegionType != d2enum.RegionNone && v.lastRegionType != tile.RegionType {
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
			v.escapeMenu, v.mapRenderer, v, v.terminal, v.uiManager, v.keyMap, v.logLevel, v.gameClient.IsSinglePlayer())

		if err != nil {
			return err
		}

		v.gameControls.Load()

		if err := v.inputManager.BindHandler(v.gameControls); err != nil {
			v.Error(bindControlsErrStr + player.ID())
		}

		break
	}

	return nil
}

// OnPlayerMove sends the player move action to the server
func (v *Game) OnPlayerMove(targetX, targetY float64) {
	worldPosition := v.localPlayer.Position.World()

	playerID, worldX, worldY := v.gameClient.PlayerID, worldPosition.X(), worldPosition.Y()

	createMovePlayerPacket, err := d2netpacket.CreateMovePlayerPacket(playerID, worldX, worldY, targetX, targetY)
	if err != nil {
		v.Errorf("MovePlayerPacket: %v", err)
	}

	err = v.gameClient.SendPacketToServer(createMovePlayerPacket)

	if err != nil {
		v.Errorf(moveErrStr, v.gameClient.PlayerID, targetX, targetY)
	}
}

// OnPlayerSave instructs the server to save our player data
func (v *Game) OnPlayerSave() error {
	playerState := v.gameClient.Players[v.gameClient.PlayerID]

	sp, err := d2netpacket.CreateSavePlayerPacket(playerState, d2enum.DifficultyNormal)
	if err != nil {
		return fmt.Errorf("SavePlayerPacket: %v", err)
	}

	err = v.gameClient.SendPacketToServer(sp)

	if err != nil {
		return err
	}

	return nil
}

// OnPlayerCast sends the casting skill action to the server
func (v *Game) OnPlayerCast(skillID int, targetX, targetY float64) {
	cp, err := d2netpacket.CreateCastPacket(v.gameClient.PlayerID, skillID, targetX, targetY)
	if err != nil {
		v.Errorf("CastPacket: %v", err)
	}

	err = v.gameClient.SendPacketToServer(cp)
	if err != nil {
		v.Errorf(castErrStr, v.gameClient.PlayerID, skillID, targetX, targetY)
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
	packet, err := d2netpacket.CreateSpawnItemPacket(x, y, codes...)
	if err != nil {
		v.Errorf("SpawnItemPacket: %v", err)
	}

	err = v.gameClient.SendPacketToServer(packet)
	if err != nil {
		v.Errorf(spawnItemErrStr, x, y, codes)
	}
}
