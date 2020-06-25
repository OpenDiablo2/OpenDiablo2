package d2client

import (
	"fmt"
	"log"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2localclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2remoteclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type GameClient struct {
	clientConnection ClientConnection
	connectionType   d2clientconnectiontype.ClientConnectionType
	GameState        *d2player.PlayerState
	MapEngine        *d2mapengine.MapEngine
	PlayerId         string
	Players          map[string]*d2mapentity.Player
	Seed             int64
	RegenMap         bool
}

func Create(connectionType d2clientconnectiontype.ClientConnectionType) (*GameClient, error) {
	result := &GameClient{
		MapEngine:      d2mapengine.CreateMapEngine(), // TODO: Mapgen - Needs levels.txt stuff
		Players:        make(map[string]*d2mapentity.Player),
		connectionType: connectionType,
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

func (g *GameClient) Open(connectionString string, saveFilePath string) error {
	return g.clientConnection.Open(connectionString, saveFilePath)
}

func (g *GameClient) Close() error {
	return g.clientConnection.Close()
}

func (g *GameClient) Destroy() error {
	return g.clientConnection.Close()
}

func (g *GameClient) OnPacketReceived(packet d2netpacket.NetPacket) error {
	switch packet.PacketType {
	case d2netpackettype.GenerateMap:
		mapData := packet.PacketData.(d2netpacket.GenerateMapPacket)
		switch mapData.RegionType {
		case d2enum.RegionAct1Town:
			d2mapgen.GenerateAct1Overworld(g.MapEngine)
		}
		g.RegenMap = true
	case d2netpackettype.UpdateServerInfo:
		serverInfo := packet.PacketData.(d2netpacket.UpdateServerInfoPacket)
		g.MapEngine.SetSeed(serverInfo.Seed)
		g.PlayerId = serverInfo.PlayerId
		g.Seed = serverInfo.Seed
		log.Printf("Player id set to %s", serverInfo.PlayerId)
	case d2netpackettype.AddPlayer:
		player := packet.PacketData.(d2netpacket.AddPlayerPacket)
		newPlayer := d2mapentity.CreatePlayer(player.Id, player.Name, player.X, player.Y, 0, player.HeroType, player.Stats, player.Equipment)
		g.Players[newPlayer.Id] = newPlayer
		g.MapEngine.AddEntity(newPlayer)
	case d2netpackettype.MovePlayer:
		movePlayer := packet.PacketData.(d2netpacket.MovePlayerPacket)
		player := g.Players[movePlayer.PlayerId]
		path, _, _ := g.MapEngine.PathFind(movePlayer.StartX, movePlayer.StartY, movePlayer.DestX, movePlayer.DestY)
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
	case d2netpackettype.Ping:
		g.clientConnection.SendPacketToServer(d2netpacket.CreatePongPacket(g.PlayerId))
	case d2netpackettype.ServerClosed:
		// TODO: Need to be tied into a character save and exit
		log.Print("Server has been closed")
		os.Exit(0)
	default:
		log.Fatalf("Invalid packet type: %d", packet.PacketType)
	}
	return nil
}

func (g *GameClient) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return g.clientConnection.SendPacketToServer(packet)
}
