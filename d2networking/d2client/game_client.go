package d2client

import (
	"fmt"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2localclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2remoteclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

const logPrefix = "Game Client"

const (
	numSubtilesPerTile = 5
)

// GameClient manages a connection to d2server.GameServer
// and keeps a synchronized copy of the map and entities.
type GameClient struct {
	clientConnection ServerConnection                            // Abstract local/remote connection
	connectionType   d2clientconnectiontype.ClientConnectionType // Type of connection (local or remote)
	asset            *d2asset.AssetManager
	scriptEngine     *d2script.ScriptEngine
	GameState        *d2hero.HeroState              // local player state
	MapEngine        *d2mapengine.MapEngine         // Map and entities
	mapGen           *d2mapgen.MapGenerator         // map generator
	PlayerID         string                         // ID of the local player
	Players          map[string]*d2mapentity.Player // IDs of the other players
	Seed             int64                          // Map seed
	RegenMap         bool                           // Regenerate tile cache on render (map has changed)

	*d2util.Logger
}

// Create constructs a new GameClient and returns a pointer to it.
func Create(connectionType d2clientconnectiontype.ClientConnectionType,
	asset *d2asset.AssetManager,
	l d2util.LogLevel,
	scriptEngine *d2script.ScriptEngine) (*GameClient, error) {
	result := &GameClient{
		asset:          asset,
		MapEngine:      d2mapengine.CreateMapEngine(l, asset),
		Players:        make(map[string]*d2mapentity.Player),
		connectionType: connectionType,
		scriptEngine:   scriptEngine,
	}

	result.Logger = d2util.NewLogger()
	result.Logger.SetPrefix(logPrefix)
	result.Logger.SetLevel(l)

	// for a remote client connection, set loading to true - wait until we process the GenerateMapPacket
	// before we start updating map entites
	result.MapEngine.IsLoading = connectionType == d2clientconnectiontype.LANClient

	mapGen, err := d2mapgen.NewMapGenerator(asset, l, result.MapEngine)
	if err != nil {
		return nil, err
	}

	result.mapGen = mapGen

	switch connectionType {
	case d2clientconnectiontype.LANClient:
		result.clientConnection, err = d2remoteclient.Create(l, asset)
	case d2clientconnectiontype.LANServer:
		result.clientConnection, err = d2localclient.Create(asset, l, true)
	case d2clientconnectiontype.Local:
		result.clientConnection, err = d2localclient.Create(asset, l, false)
	default:
		err = fmt.Errorf("unknown client connection type specified: %d", connectionType)
	}

	if err != nil {
		return nil, err
	}

	result.clientConnection.SetClientListener(result)

	return result, nil
}

// Open creates the server and connects to it if the client is local.
// If the client is remote it sends a PlayerConnectionRequestPacket to the
// server (see d2netpacket).
func (g *GameClient) Open(connectionString, saveFilePath string) error {
	switch g.connectionType {
	case d2clientconnectiontype.LANServer, d2clientconnectiontype.Local:
		g.scriptEngine.AllowEval()
	}

	return g.clientConnection.Open(connectionString, saveFilePath)
}

// Close destroys the server if the client is local. For remote clients
// it sends a DisconnectRequestPacket (see d2netpacket).
func (g *GameClient) Close() error {
	switch g.connectionType {
	case d2clientconnectiontype.LANServer, d2clientconnectiontype.Local:
		g.scriptEngine.DisallowEval()
	}

	return g.clientConnection.Close()
}

// Destroy does the same thing as Close.
func (g *GameClient) Destroy() error {
	return g.Close()
}

// OnPacketReceived is called by the ClientConection and processes incoming
// packets.
// nolint:gocyclo // switch statement on packet type makes sense, no need to change
func (g *GameClient) OnPacketReceived(packet d2netpacket.NetPacket) error {
	switch packet.PacketType {
	case d2netpackettype.GenerateMap:
		if err := g.handleGenerateMapPacket(packet); err != nil {
			return err
		}
	case d2netpackettype.UpdateServerInfo:
		if err := g.handleUpdateServerInfoPacket(packet); err != nil {
			return err
		}
	case d2netpackettype.AddPlayer:
		if err := g.handleAddPlayerPacket(packet); err != nil {
			return err
		}
	case d2netpackettype.MovePlayer:
		if err := g.handleMovePlayerPacket(packet); err != nil {
			return err
		}
	case d2netpackettype.CastSkill:
		if err := g.handleCastSkillPacket(packet); err != nil {
			return err
		}
	case d2netpackettype.SpawnItem:
		if err := g.handleSpawnItemPacket(packet); err != nil {
			return err
		}
	case d2netpackettype.Ping:
		if err := g.handlePingPacket(); err != nil {
			g.Errorf("GameClient: error responding to server ping: %s", err)
		}
	case d2netpackettype.PlayerDisconnectionNotification:
		// Not implemented
		g.Infof("RemoteClientConnection: received disconnect: %s", packet.PacketData)
	case d2netpackettype.ServerClosed:
		// https://github.com/OpenDiablo2/OpenDiablo2/issues/802
		g.Infof("Server has been closed")
		os.Exit(0)
	case d2netpackettype.ServerFull:
		g.Infof("Server is full") // need to be verified
		os.Exit(0)
	default:
		g.Fatalf("Invalid packet type: %d", packet.PacketType)
	}

	return nil
}

// SendPacketToServer calls server.OnPacketReceived if the client is local.
// If it is remote the NetPacket sent over a UDP connection to the server.
func (g *GameClient) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return g.clientConnection.SendPacketToServer(packet)
}

func (g *GameClient) handleGenerateMapPacket(packet d2netpacket.NetPacket) error {
	mapData, err := d2netpacket.UnmarshalGenerateMap(packet.PacketData)
	if err != nil {
		return err
	}

	if mapData.RegionType == d2enum.RegionAct1Town {
		g.mapGen.GenerateAct1Overworld()
	}

	g.RegenMap = true

	return nil
}

func (g *GameClient) handleUpdateServerInfoPacket(packet d2netpacket.NetPacket) error {
	serverInfo, err := d2netpacket.UnmarshalUpdateServerInfo(packet.PacketData)
	if err != nil {
		return err
	}

	g.MapEngine.SetSeed(serverInfo.Seed)
	g.PlayerID = serverInfo.PlayerID
	g.Seed = serverInfo.Seed
	g.Infof("Player id set to %s", serverInfo.PlayerID)

	return nil
}

func (g *GameClient) handleAddPlayerPacket(packet d2netpacket.NetPacket) error {
	player, err := d2netpacket.UnmarshalAddPlayer(packet.PacketData)
	if err != nil {
		return err
	}

	d2hero.HydrateSkills(player.Skills, g.asset)

	newPlayer := g.MapEngine.NewPlayer(player.ID, player.Name, player.X, player.Y, 0,
		player.HeroType, player.Stats, player.Skills, &player.Equipment, player.LeftSkill, player.RightSkill, player.Gold)

	g.Players[newPlayer.ID()] = newPlayer
	g.MapEngine.AddEntity(newPlayer)

	return nil
}

func (g *GameClient) handleSpawnItemPacket(packet d2netpacket.NetPacket) error {
	item, err := d2netpacket.UnmarshalSpawnItem(packet.PacketData)
	if err != nil {
		return err
	}

	itemEntity, err := g.MapEngine.NewItem(item.X, item.Y, item.Codes...)

	if err == nil {
		g.MapEngine.AddEntity(itemEntity)
	}

	return err
}

func (g *GameClient) handleMovePlayerPacket(packet d2netpacket.NetPacket) error {
	movePlayer, err := d2netpacket.UnmarshalMovePlayer(packet.PacketData)
	if err != nil {
		return err
	}

	player := g.Players[movePlayer.PlayerID]
	start := d2vector.NewPositionTile(movePlayer.StartX, movePlayer.StartY)
	dest := d2vector.NewPositionTile(movePlayer.DestX, movePlayer.DestY)
	path := g.MapEngine.PathFind(start, dest)

	if len(path) > 0 {
		player.SetPath(path, func() {
			tilePosition := player.Position.Tile()
			tile := g.MapEngine.TileAt(int(tilePosition.X()), int(tilePosition.Y()))

			if tile == nil {
				return
			}

			player.SetIsInTown(tile.RegionType == d2enum.RegionAct1Town)

			err := player.SetAnimationMode(player.GetAnimationMode())

			if err != nil {
				fmtStr := "GameClient: error setting animation mode for player %s: %s"
				g.Errorf(fmtStr, player.ID(), err)
			}
		})
	}

	return nil
}

func (g *GameClient) handleCastSkillPacket(packet d2netpacket.NetPacket) error {
	playerCast, err := d2netpacket.UnmarshalCast(packet.PacketData)
	if err != nil {
		return err
	}

	player := g.Players[playerCast.SourceEntityID]
	player.StopMoving()

	castX := playerCast.TargetX * numSubtilesPerTile
	castY := playerCast.TargetY * numSubtilesPerTile

	direction := player.Position.DirectionTo(*d2vector.NewVector(castX, castY))
	player.SetDirection(direction)

	skillRecord := g.asset.Records.Skill.Details[playerCast.SkillID]

	missileEntities, err := g.createMissileEntities(skillRecord, player, castX, castY)
	if err != nil {
		return err
	}

	var summonedNpcEntity *d2mapentity.NPC
	if skillRecord.Summon != "" {
		summonedNpcEntity, err = g.createSummonedNpcEntity(skillRecord, int(castX), int(castY))

		if err != nil {
			return err
		}
	}

	player.StartCasting(skillRecord.Anim, func() {
		if len(missileEntities) > 0 {
			// shoot the missiles of the skill after the player has finished casting
			for _, missileEntity := range missileEntities {
				g.MapEngine.AddEntity(missileEntity)
			}
		}

		if summonedNpcEntity != nil {
			// summon the referenced NPC after the player has finished casting
			g.MapEngine.AddEntity(summonedNpcEntity)
		}
	})

	overlayRecord := g.asset.Records.Layout.Overlays[skillRecord.Castoverlay]

	return g.playCastOverlay(overlayRecord, int(player.Position.X()), int(player.Position.Y()))
}

func (g *GameClient) createSummonedNpcEntity(skillRecord *d2records.SkillRecord, x, y int) (*d2mapentity.NPC, error) {
	monsterStatsRecord := g.asset.Records.Monster.Stats[skillRecord.Summon]

	if monsterStatsRecord == nil {
		fmtErr := "cannot cast skill - No monstat entry for \"%s\""
		return nil, fmt.Errorf(fmtErr, skillRecord.Summon)
	}

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/803
	summonedNpcEntity, err := g.MapEngine.NewNPC(x, y, monsterStatsRecord, 0)
	if err != nil {
		return nil, err
	}

	return summonedNpcEntity, nil
}

func (g *GameClient) createMissileEntities(
	skillRecord *d2records.SkillRecord,
	player *d2mapentity.Player,
	castX, castY float64,
) ([]*d2mapentity.Missile, error) {
	missileRecords := []*d2records.MissileRecord{
		g.asset.Records.GetMissileByName(skillRecord.Cltmissile),
		g.asset.Records.GetMissileByName(skillRecord.Cltmissilea),
		g.asset.Records.GetMissileByName(skillRecord.Cltmissileb),
		g.asset.Records.GetMissileByName(skillRecord.Cltmissilec),
		g.asset.Records.GetMissileByName(skillRecord.Cltmissiled),
	}

	missileEntities := make([]*d2mapentity.Missile, 0)

	for _, missileRecord := range missileRecords {
		if missileRecord == nil {
			continue
		}

		missileEntity, err := g.createMissileEntity(missileRecord, player, castX, castY)
		if err != nil {
			return nil, err
		}

		missileEntities = append(missileEntities, missileEntity)
	}

	return missileEntities, nil
}

func (g *GameClient) createMissileEntity(
	missileRecord *d2records.MissileRecord,
	player *d2mapentity.Player,
	castX, castY float64,
) (*d2mapentity.Missile, error) {
	if missileRecord == nil {
		return nil, nil
	}

	radians := d2math.GetRadiansBetween(
		player.Position.X(),
		player.Position.Y(),
		castX,
		castY,
	)

	missileEntity, err := g.MapEngine.NewMissile(
		int(player.Position.X()),
		int(player.Position.Y()),
		g.asset.Records.Missiles[missileRecord.Id],
	)

	if err != nil {
		return nil, err
	}

	missileEntity.SetRadians(radians, func() {
		g.MapEngine.RemoveEntity(missileEntity)
	})

	return missileEntity, nil
}

func (g *GameClient) playCastOverlay(overlayRecord *d2records.OverlayRecord, x, y int) error {
	if overlayRecord == nil {
		return nil
	}

	overlayEntity, err := g.MapEngine.NewCastOverlay(
		x,
		y,
		overlayRecord,
	)
	if err != nil {
		return err
	}

	overlayEntity.SetOnDoneFunc(func() {
		g.MapEngine.RemoveEntity(overlayEntity)
	})

	g.MapEngine.AddEntity(overlayEntity)

	return nil
}

func (g *GameClient) handlePingPacket() error {
	pongPacket, err := d2netpacket.CreatePongPacket(g.PlayerID)
	if err != nil {
		return err
	}

	err = g.clientConnection.SendPacketToServer(pongPacket)

	if err != nil {
		return err
	}

	return nil
}

// IsSinglePlayer returns a bool for whether the game is a single-player game
func (g *GameClient) IsSinglePlayer() bool {
	return g.connectionType == d2clientconnectiontype.Local
}
