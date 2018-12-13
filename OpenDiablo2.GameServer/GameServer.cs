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
        private readonly IItemManager itemManager;

        public int Seed { get; private set; }
        public IEnumerable<PlayerState> Players => mobManager.Players;

        const double Deg2Rad = Math.PI / 180.0;

        public GameServer(IMobManager mobManager, IEngineDataManager engineDataManager, IItemManager itemManager)
        {
            this.mobManager = mobManager;
            this.engineDataManager = engineDataManager;
            this.itemManager = itemManager;
        }

        public void InitializeNewGame()
        {
            log.Info("Initializing a new game");
            Seed = new Random().Next();
        }

        public int SpawnNewPlayer(int clientHash, string playerName, eHero heroType)
        {
            ILevelExperienceConfig expConfig = null;
            IHeroTypeConfig heroConfig = null;
            if (engineDataManager.ExperienceConfigs.ContainsKey(heroType))
            {
                expConfig = engineDataManager.ExperienceConfigs[heroType];
            }
            else
            {
                log.Error("Error: Experience Config not loaded for '" + heroType.ToString() + "'.");
                expConfig = new LevelExperienceConfig(new List<long>() { 100 });
                // TODO: should we have a more robust default experience config?
                // or should we just fail in some way here?
            }
            if (engineDataManager.HeroTypeConfigs.ContainsKey(heroType))
            {
                heroConfig = engineDataManager.HeroTypeConfigs[heroType];
            }
            else
            {
                log.Error("Error: Hero Config not loaded for '" + heroType.ToString() + "'.");
                // Do we even need a default?
                //heroConfig = new HeroTypeConfig(10, 10, 10, 10, 10, 10, 10, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 9,
                //    1, 10, 10, 10, 10, 10, 10, 0, "hth");
                // TODO: should we have a more robust default hero config?
                // or should we just fail in some way here?
                // ... we should probably just fail here
            }

            var newPlayer = new PlayerState(clientHash, playerName, mobManager.GetNextAvailableMobId(), 1, 20.0f, 20.0f, 10, 10, 10, 10, 0, heroType, 
                heroConfig, expConfig);

            // This is probably not the right place to do this.
            // Only add items with a location set, the other ones go into the inventory - that we do not support yet
            foreach (var item in heroConfig.InitialEquipment)
            {
                if (item.location.Length > 0)
                {
                    newPlayer.UpdateEquipment(item.location, itemManager.getItemInstance(item.name));
                }
            }

        
            mobManager.AddPlayer(newPlayer);
            return newPlayer.Id;
        }

        public void Update(int ms)
        {
            var seconds = ms / 1000f;
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

            var speed = (player.MovementType == eMovementType.Running ? player.GetRunVelocity() : player.GetWalkVeloicty()) / 4f;

            var moveX = (float)Math.Cos(rads) * seconds * speed;
            var moveY = (float)Math.Sin(rads) * seconds * speed;

            player.X += moveX;
            player.Y += moveY;
        }

        public void Dispose()
        {
        }
    }
}
