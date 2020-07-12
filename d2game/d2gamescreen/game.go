package d2gamescreen

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

const hideZoneTextAfterSeconds = 2.0

// Game represents the Gameplay screen
type Game struct {
	gameClient           *d2client.GameClient
	mapRenderer          *d2maprenderer.MapRenderer
	gameControls         *d2player.GameControls // TODO: Hack
	localPlayer          *d2mapentity.Player
	lastRegionType       d2enum.RegionIdType
	ticksSinceLevelCheck float64
	escapeMenu           *EscapeMenu

	renderer      d2interface.Renderer
	audioProvider d2interface.AudioProvider
	terminal      d2interface.Terminal
}

// CreateGame creates the Gameplay screen and returns a pointer to it
func CreateGame(renderer d2interface.Renderer, audioProvider d2interface.AudioProvider, gameClient *d2client.GameClient,
	term d2interface.Terminal, scriptEngine *d2script.ScriptEngine) *Game {
	result := &Game{
		gameClient:           gameClient,
		gameControls:         nil,
		localPlayer:          nil,
		lastRegionType:       d2enum.RegionNone,
		ticksSinceLevelCheck: 0,
		mapRenderer:          d2maprenderer.CreateMapRenderer(renderer, gameClient.MapEngine, term),
		escapeMenu:           NewEscapeMenu(renderer, audioProvider, term, scriptEngine),
		audioProvider:        audioProvider,
		renderer:             renderer,
		terminal:             term,
	}
	result.escapeMenu.onLoad()

	if err := d2input.BindHandler(result.escapeMenu); err != nil {
		fmt.Println("failed to add gameplay screen as event handler")
	}

	return result
}

// OnLoad loads the resources for the Gameplay screen
func (v *Game) OnLoad(_ d2screen.LoadingState) {
	v.audioProvider.PlayBGM("")
}

// OnUnload releases the resources of Gameplay screen
func (v *Game) OnUnload() error {
	if err := d2input.UnbindHandler(v.gameControls); err != nil { // TODO: hack
		return err
	}

	if err := d2input.UnbindHandler(v.escapeMenu); err != nil { // TODO: hack
		return err
	}

	if err := v.gameClient.Close(); err != nil {
		return err
	}

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
		v.gameControls.Render(screen)
	}

	return nil
}

// Advance runs the update logic on the Gameplay screen
func (v *Game) Advance(tickTime float64) error {
	if (v.escapeMenu != nil && !v.escapeMenu.isOpen) || len(v.gameClient.Players) != 1 {
		v.gameClient.MapEngine.Advance(tickTime) // TODO: Hack
	}

	if v.gameControls != nil {
		if err := v.gameControls.Advance(tickTime); err != nil {
			return err
		}
	}

	v.ticksSinceLevelCheck += tickTime
	if v.ticksSinceLevelCheck > 1.0 {
		v.ticksSinceLevelCheck = 0
		if v.localPlayer != nil {
			tilePosition := v.localPlayer.Position.Tile()
			tile := v.gameClient.MapEngine.TileAt(int(tilePosition.X()), int(tilePosition.Y()))

			if tile != nil {
				musicInfo := d2common.GetMusicDef(tile.RegionType)
				v.audioProvider.PlayBGM(musicInfo.MusicFile)

				// skip showing zone change text the first time we enter the world
				if v.lastRegionType != d2enum.RegionNone && v.lastRegionType != tile.RegionType {
					//TODO: Should not be using RegionType as an index - this will return incorrect LevelDetails record for most of the zones.
					v.gameControls.SetZoneChangeText(fmt.Sprintf("Entering The %s", d2datadict.LevelDetails[int(tile.RegionType)].LevelDisplayName))
					v.gameControls.ShowZoneChangeText()
					v.gameControls.HideZoneChangeTextAfter(hideZoneTextAfterSeconds)
				}

				v.lastRegionType = tile.RegionType
			}
		}
	}

	// Bind the game controls to the player once it exists
	if v.gameControls == nil {
		v.bindGameControls()
	}

	// Update the camera to focus on the player
	if v.localPlayer != nil && !v.gameControls.FreeCam {
		worldPosition := v.localPlayer.Position.World()
		rx, ry := v.mapRenderer.WorldToOrtho(worldPosition.X(), worldPosition.Y())
		v.mapRenderer.MoveCameraTo(rx, ry)
	}

	return nil
}

func (v *Game) bindGameControls() {
	for _, player := range v.gameClient.Players {
		if player.Id != v.gameClient.PlayerId {
			continue
		}

		v.localPlayer = player
		v.gameControls = d2player.NewGameControls(v.renderer, player, v.gameClient.MapEngine, v.mapRenderer, v, v.terminal)
		v.gameControls.Load()

		if err := d2input.BindHandler(v.gameControls); err != nil {
			fmt.Printf("failed to add gameControls as input handler for player: %s\n", player.Id)
		}

		break
	}
}

// OnPlayerMove sends the player move action to the server
func (v *Game) OnPlayerMove(x, y float64) {
	worldPosition := v.localPlayer.Position.World()

	err := v.gameClient.SendPacketToServer(d2netpacket.CreateMovePlayerPacket(v.gameClient.PlayerId, worldPosition.X(), worldPosition.Y(), x, y))
	if err != nil {
		fmt.Printf("failed to send MovePlayer packet to the server, playerId: %s, x: %g, x: %g\n", v.gameClient.PlayerId, x, y)
	}
}

// OnPlayerCast sends the casting skill action to the server
func (v *Game) OnPlayerCast(missileID int, targetX, targetY float64) {
	err := v.gameClient.SendPacketToServer(d2netpacket.CreateCastPacket(v.gameClient.PlayerId, missileID, targetX, targetY))
	if err != nil {
		fmt.Printf(
			"failed to send CastSkill packet to the server, playerId: %s, missileId: %d, x: %g, x: %g\n",
			v.gameClient.PlayerId, missileID, targetX, targetY,
		)
	}
}
