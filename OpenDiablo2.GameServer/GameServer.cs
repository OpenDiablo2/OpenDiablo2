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
        private readonly IEngineDataManager engineDataManager;

        public int Seed { get; private set; }
        public IEnumerable<PlayerState> Players => mobManager.Players;

        const double Deg2Rad = Math.PI / 180.0;

        public GameServer(IMobManager mobManager, IEngineDataManager engineDataManager)
        {
            this.mobManager = mobManager;
            this.engineDataManager = engineDataManager;
        }

        public void InitializeNewGame()
        {
            log.Info("Initializing a new game");
            Seed = (new Random()).Next();
        }

        public int SpawnNewPlayer(int clientHash, string playerName, eHero heroType)
        {
            ILevelExperienceConfig expConfig = null;
            try
            {
                expConfig = engineDataManager.ExperienceConfigs[heroType];
            }
            catch(Exception e)
            {
                log.Error("Error: Experience Config not loaded for '" + heroType.ToString() + "'.");
                expConfig = new LevelExperienceConfig(new List<long>() { 100 });
                // TODO: should we have a more robust default experience config?
                // or should we just fail in some way here?
            }
            var newPlayer = new PlayerState(clientHash, playerName, mobManager.GetNextAvailableMobId(), 1, 20.0f, 20.0f, 10, 10, 10, 10, 0, heroType, 
                new HeroTypeConfig(10, 10, 10, 50, 50, 50, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1), expConfig);

            mobManager.AddPlayer(newPlayer);
            return newPlayer.Id;
        }

        public void Update(int ms)
        {
            var seconds = (float)ms / 1000f;
            foreach(var player in Players)
            {
                UpdatePlayerMovement(player, seconds);
            }
        }

        private void UpdatePlayerMovement(PlayerState player, float seconds)
        {
            // TODO: We need to do collision detection here...
            if (player.MovementType == eMovementType.Stopped)
                return;

            var rads = (float)player.MovementDirection * 22 * (float)Deg2Rad;

            var moveX = (float)Math.Cos(rads) * seconds * 2f;
            var moveY = (float)Math.Sin(rads) * seconds * 2f;

            player.X += moveX;
            player.Y += moveY;
        }

        public void Dispose()
        {
        }
    }
}
