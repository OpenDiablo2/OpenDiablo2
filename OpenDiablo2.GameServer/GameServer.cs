/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.Linq;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;
using OpenDiablo2.Common.Services;

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

            var newPlayer = new PlayerState(clientHash, playerName, mobManager.GetNextAvailableMobId(), 1, 20.5f, 20.5f, 10, 10, 10, 10, 0, heroType, 
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

            // TODO: Default torso for testing. Remove when... we're done testing.
            newPlayer.UpdateEquipment("tors", itemManager.getItemInstance("aar"));
        
            mobManager.AddPlayer(newPlayer);
            return newPlayer.Id;
        }

        public PlayerEquipment UpdateEquipment(int clienthash, string slot, ItemInstance itemInstance)
        {
            var player = mobManager.Players.FirstOrDefault(x => x.ClientHash == clienthash);

            player.Equipment.EquipItem(slot, itemInstance);

            return player.Equipment;
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
            if (player.Waypoints.Count == 0)
                return;

            (new MobMovementService(player)).CalculateMovement(seconds);
        }

        public void Dispose()
        {
        }
    }
}
