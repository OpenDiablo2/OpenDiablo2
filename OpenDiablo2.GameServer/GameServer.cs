using System;
using System.Collections.Generic;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.GameServer_
{
    public sealed class GameServer : IGameServer
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IMobManager mobManager;

        public int Seed { get; private set; }
        public IEnumerable<PlayerState> Players => mobManager.Players;

        public GameServer(IMobManager mobManager)
        {
            this.mobManager = mobManager;
        }

        public void InitializeNewGame()
        {
            log.Info("Initializing a new game");
            Seed = (new Random()).Next();
        }

        public int SpawnNewPlayer(int clientHash, string playerName, eHero heroType)
        {
            var newPlayer = new PlayerState(clientHash, playerName, mobManager.GetNextAvailableMobId(), 1, 20.0f, 20.0f, 10, 10, 10, 10, 0, heroType, 
                new HeroTypeConfig(10, 10, 10, 50, 50, 50, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1), new LevelExperienceConfig(new List<int>() { 100 }));

            mobManager.AddPlayer(newPlayer);
            return newPlayer.Id;
        }

        public void Dispose()
        {
        }
    }
}
