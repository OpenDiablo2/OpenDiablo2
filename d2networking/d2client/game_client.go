package d2client

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2localclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type GameClient struct {
	clientConnection ClientConnection
	GameState        *d2gamestate.GameState
	MapEngine        *d2map.MapEngine
	playerId         string
	players          []*d2map.Player
	GameControls     *d2player.GameControls // TODO: Hack
}

func Create(connectionType d2clientconnectiontype.ClientConnectionType) (*GameClient, error) {
	result := &GameClient{
		MapEngine: d2map.CreateMapEngine(),
		players:   make([]*d2map.Player, 0),
	}

	switch connectionType {
	case d2clientconnectiontype.Local:
		result.clientConnection = d2localclient.Create()
		result.clientConnection.SetClientListener(result)
	default:
		return nil, fmt.Errorf("unknown client connection type specified: %d", connectionType)
	}

	return result, nil
}

func (g *GameClient) Open(connectionString string) error {
	return g.clientConnection.Open(connectionString)
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
		g.MapEngine.GenerateMap(mapData.RegionType, mapData.LevelPreset, mapData.FileIndex)
		break
	case d2netpackettype.UpdateServerInfo:
		serverInfo := packet.PacketData.(d2netpacket.UpdateServerInfoPacket)
		g.MapEngine.SetSeed(serverInfo.Seed)
		g.playerId = serverInfo.PlayerId
		log.Printf("Player id set to %s", serverInfo.PlayerId)
		break
	case d2netpackettype.AddPlayer:
		player := packet.PacketData.(d2netpacket.AddPlayerPacket)
		newPlayer := d2map.CreatePlayer(player.Id, (player.X*5)+3, (player.Y*5)+3, 0, player.HeroType, player.Equipment)
		g.MapEngine.AddEntity(newPlayer)
		if newPlayer.Id == g.playerId {
			g.GameControls = d2player.NewGameControls(newPlayer, g.MapEngine)
			g.GameControls.Load()
			d2input.BindHandler(g.GameControls)

			// TODO: Temporary stuff
			rx, ry := g.MapEngine.WorldToOrtho(newPlayer.AnimatedComposite.LocationX/5, newPlayer.AnimatedComposite.LocationY/5)
			g.MapEngine.MoveCameraTo(rx, ry)
		}
		break
	default:
		log.Fatalf("Invalid packet type: %d", packet.PacketType)
	}
	return nil
}

func (g *GameClient) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return g.clientConnection.SendPacketToServer(packet)
}
