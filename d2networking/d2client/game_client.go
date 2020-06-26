package d2client

import (
	"log"
	"os"

	// "github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2maprenderer"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	d2cct "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type GameClient struct {
	clientConnection ClientConnection
	connectionType   d2cct.ClientConnectionType
	GameState        *d2player.PlayerState
	MapEngine        *d2mapengine.MapEngine
	MapRenderer      *d2maprenderer.MapRenderer
	realm            *d2mapengine.MapRealm
	PlayerId         string
	Players          map[string]*d2mapentity.Player
	Seed             int64
	RegenMap         bool
}

// Using the `clientConnection`, opens a connection and passes the savefile path
func (g *GameClient) Open(connectionString string, saveFilePath string) error {
	return g.clientConnection.Open(connectionString, saveFilePath)
}

// Closes the `clientConnection`
func (g *GameClient) Close() error {
	return g.clientConnection.Close()
}

// Closes the `clientConnection`
func (g *GameClient) Destroy() error {
	return g.clientConnection.Close()
}

// Routes the incoming packets to the packet handlers
func (g *GameClient) OnPacketReceived(packet d2netpacket.NetPacket) error {
	switch packet.PacketType {

	// UNSURE: should we be bubbling up errors from these handler calls?
	case d2netpackettype.UpdateServerInfo:
		g.handleUpdateServerInfo(packet)

	case d2netpackettype.AddPlayer:
		g.handleAddPlayer(packet)

	case d2netpackettype.GenerateMap:
		g.handleGenerateMap(packet)

	case d2netpackettype.MovePlayer:
		g.handleMovePlayer(packet)

	case d2netpackettype.Ping:
		g.handlePong(packet)

	case d2netpackettype.ServerClosed:
		g.handleServerClosed(packet)

	default:
		log.Fatalf("Invalid packet type: %d", packet.PacketType)
	}
	return nil
}

// Using the `clientConnection`, sends a packet to the server
func (g *GameClient) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return g.clientConnection.SendPacketToServer(packet)
}

func (g *GameClient) handleUpdateServerInfo(p d2netpacket.NetPacket) {
	serverInfo := p.PacketData.(d2netpacket.UpdateServerInfoPacket)
	seed := serverInfo.Seed
	playerId := serverInfo.PlayerId

	g.Seed = seed
	g.realm.Init(seed)
	g.PlayerId = playerId

	log.Printf("Player id set to %s", playerId)
}

func (g *GameClient) handleAddPlayer(p d2netpacket.NetPacket) {
	player := p.PacketData.(d2netpacket.AddPlayerPacket)
	levelId := g.realm.GetFirstActLevelId(player.Act)
	g.MapEngine = g.realm.GetMapEngine(player.Act, levelId)

	pId := player.Id
	pName := player.Name
	pAct := player.Act
	pLvlId := levelId
	pX := player.X
	pY := player.Y
	pDir := 0
	pHero := player.HeroType
	pStat := player.Stats
	pEquip := player.Equipment

	// UNSURE: maybe we should be passing a struct instead of all the vars?
	newPlayer := d2mapentity.CreatePlayer(
		pId, pName, pAct, pLvlId, pX, pY, pDir, pHero, pStat, pEquip,
	)

	g.Players[newPlayer.Id] = newPlayer
	g.realm.AddPlayer(pId, pAct)
	g.MapEngine.AddEntity(newPlayer)
}

func (g *GameClient) handleGenerateMap(p d2netpacket.NetPacket) {
	mapData := p.PacketData.(d2netpacket.GenerateMapPacket)
	g.realm.GenerateMap(mapData.ActId, mapData.LevelId)
	engine := g.realm.GetMapEngine(mapData.ActId, mapData.LevelId)
	g.MapRenderer = d2maprenderer.CreateMapRenderer(engine)
	g.RegenMap = true
}

func (g *GameClient) handleMovePlayer(p d2netpacket.NetPacket) {
	movePlayer := p.PacketData.(d2netpacket.MovePlayerPacket)

	player := g.Players[movePlayer.PlayerId]
	x1, y1 := movePlayer.StartX, movePlayer.StartY
	x2, y2 := movePlayer.DestX, movePlayer.DestY

	path, _, _ := g.MapEngine.PathFind(x1, y1, x2, y2)

	if len(path) > 0 {
		player.SetPath(path, func() {
			tile := g.MapEngine.TileAt(player.TileX, player.TileY)
			if tile == nil {
				return
			}

			regionType := tile.RegionType
			if regionType == d2enum.RegionAct1Town {
				player.SetIsInTown(true)
			} else {
				player.SetIsInTown(false)
			}
			player.SetAnimationMode(player.GetAnimationMode().String())
		})
	}
}

func (g *GameClient) handlePong(p d2netpacket.NetPacket) {
	pong := d2netpacket.CreatePongPacket(g.PlayerId)
	g.clientConnection.SendPacketToServer(pong)
}

func (g *GameClient) handleServerClosed(p d2netpacket.NetPacket) {
	// TODO: Need to be tied into a character save and exit
	log.Print("Server has been closed")
	os.Exit(0)
}
