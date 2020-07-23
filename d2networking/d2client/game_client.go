package d2client

import (
	"fmt"
	"log"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2localclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2remoteclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

const (
	numSubtilesPerTile = 5
)

// GameClient manages a connection to d2server.GameServer
// and keeps a synchronized copy of the map and entities.
type GameClient struct {
	clientConnection ServerConnection                            // Abstract local/remote connection
	connectionType   d2clientconnectiontype.ClientConnectionType // Type of connection (local or remote)
	scriptEngine     *d2script.ScriptEngine
	GameState        *d2player.PlayerState          // local player state
	MapEngine        *d2mapengine.MapEngine         // Map and entities
	PlayerID         string                         // ID of the local player
	Players          map[string]*d2mapentity.Player // IDs of the other players
	Seed             int64                          // Map seed
	RegenMap         bool                           // Regenerate tile cache on render (map has changed)
}

// Create constructs a new GameClient and returns a pointer to it.
func Create(connectionType d2clientconnectiontype.ClientConnectionType, scriptEngine *d2script.ScriptEngine) (*GameClient, error) {
	result := &GameClient{
		MapEngine:      d2mapengine.CreateMapEngine(), // TODO: Mapgen - Needs levels.txt stuff
		Players:        make(map[string]*d2mapentity.Player),
		connectionType: connectionType,
		scriptEngine:   scriptEngine,
	}

	switch connectionType {
	case d2clientconnectiontype.LANClient:
		result.clientConnection = d2remoteclient.Create()
	case d2clientconnectiontype.LANServer:
		result.clientConnection = d2localclient.Create(true)
	case d2clientconnectiontype.Local:
		result.clientConnection = d2localclient.Create(false)
	default:
		return nil, fmt.Errorf("unknown client connection type specified: %d", connectionType)
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
	case d2netpackettype.Ping:
		if err := g.handlePingPacket(); err != nil {
			log.Printf("GameClient: error responding to server ping: %s", err)
		}
	case d2netpackettype.PlayerDisconnectionNotification:
		// Not implemented
		log.Printf("RemoteClientConnection: received disconnect: %s", packet.PacketData)
	case d2netpackettype.ServerClosed:
		// TODO: Need to be tied into a character save and exit
		log.Print("Server has been closed")
		os.Exit(0)
	default:
		log.Fatalf("Invalid packet type: %d", packet.PacketType)
	}

	return nil
}

// SendPacketToServer calls server.OnPacketReceived if the client is local.
// If it is remote the NetPacket sent over a UDP connection to the server.
func (g *GameClient) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return g.clientConnection.SendPacketToServer(packet)
}

func (g *GameClient) handleGenerateMapPacket(packet d2netpacket.NetPacket) error {
	mapData := packet.PacketData.(d2netpacket.GenerateMapPacket)

	if mapData.RegionType == d2enum.RegionAct1Town {
		d2mapgen.GenerateAct1Overworld(g.MapEngine)
	}

	g.RegenMap = true

	return nil
}

func (g *GameClient) handleUpdateServerInfoPacket(packet d2netpacket.NetPacket) error {
	serverInfo := packet.PacketData.(d2netpacket.UpdateServerInfoPacket)
	g.MapEngine.SetSeed(serverInfo.Seed)
	g.PlayerID = serverInfo.PlayerID
	g.Seed = serverInfo.Seed
	log.Printf("Player id set to %s", serverInfo.PlayerID)

	return nil
}

func (g *GameClient) handleAddPlayerPacket(packet d2netpacket.NetPacket) error {
	player := packet.PacketData.(d2netpacket.AddPlayerPacket)
	newPlayer := d2mapentity.CreatePlayer(player.ID, player.Name, player.X, player.Y, 0,
		player.HeroType, player.Stats, &player.Equipment)

	g.Players[newPlayer.ID] = newPlayer
	g.MapEngine.AddEntity(newPlayer)

	return nil
}

func (g *GameClient) handleMovePlayerPacket(packet d2netpacket.NetPacket) error {
	movePlayer := packet.PacketData.(d2netpacket.MovePlayerPacket)
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

			regionType := tile.RegionType
			if regionType == d2enum.RegionAct1Town {
				player.SetIsInTown(true)
			} else {
				player.SetIsInTown(false)
			}

			err := player.SetAnimationMode(player.GetAnimationMode())

			if err != nil {
				log.Printf("GameClient: error setting animation mode for player %s: %s", player.ID, err)
			}
		})
	}

	return nil
}

func (g *GameClient) handleCastSkillPacket(packet d2netpacket.NetPacket) error {
	playerCast := packet.PacketData.(d2netpacket.CastPacket)
	player := g.Players[playerCast.SourceEntityID]

	player.SetCasting()
	player.ClearPath()

	// currently hardcoded to missile skill
	missile, err := d2mapentity.CreateMissile(
		int(player.Position.X()),
		int(player.Position.Y()),
		d2datadict.Missiles[playerCast.SkillID],
	)

	if err != nil {
		return err
	}

	rads := d2common.GetRadiansBetween(
		player.Position.X(),
		player.Position.Y(),
		playerCast.TargetX*numSubtilesPerTile,
		playerCast.TargetY*numSubtilesPerTile,
	)

	missile.SetRadians(rads, func() {
		g.MapEngine.RemoveEntity(missile)
	})

	g.MapEngine.AddEntity(missile)

	return nil
}

func (g *GameClient) handlePingPacket() error {
	pongPacket := d2netpacket.CreatePongPacket(g.PlayerID)
	err := g.clientConnection.SendPacketToServer(pongPacket)

	if err != nil {
		return err
	}

	return nil
}
