using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Core.GameState_
{
    public class MobManager : IMobManager
    {
        private List<MobState> Mobs = new List<MobState>(); // all mobs (including players!)
        private List<PlayerState> Players = new List<PlayerState>();
        private List<EnemyState> Enemies = new List<EnemyState>();

        #region Player Controls
        public void AddPlayer(PlayerState player)
        {
            Players.Add(player);
            Mobs.Add(player);
        }
        public void RemovePlayer(PlayerState player)
        {
            Players.Remove(player);
            Mobs.Remove(player);
        }
        #endregion Player Controls

        #region Mob Controls
        public void AddMob(MobState mob)
        {
            Mobs.Add(mob);
        }
        public void RemoveMob(MobState mob)
        {
            Mobs.Remove(mob);
        }
        #endregion Mob Controls

        #region Enemy Controls
        public void AddEnemy(EnemyState enemy)
        {
            Enemies.Add(enemy);
            Mobs.Add(enemy);
        }
        public void RemoveEnemy(EnemyState enemy)
        {
            Enemies.Remove(enemy);
            Mobs.Remove(enemy);
        }
        #endregion Enemy Controls

        #region Searching and Filtering
        public List<MobState> FilterMobs(IEnumerable<MobState> mobs, IMobCondition condition)
        {
            // note: if condition is null, returns full list
            List<MobState> filtered = new List<MobState>();
            foreach(MobState mob in mobs)
            {
                if (condition == null || condition.Evaluate(mob))
                {
                    filtered.Add(mob);
                }
            }
            return filtered;
        }

        public List<MobState> FindMobs(IMobCondition condition)
        {
            return FilterMobs(Mobs, condition);
        }
        public List<MobState> FindEnemies(IMobCondition condition)
        {
            return FilterMobs(Enemies, condition);
        }
        public List<MobState> FindPlayers(IMobCondition condition)
        {
            return FilterMobs(Players, condition);
        }

        public List<MobState> FindInRadius(IEnumerable<MobState> mobs, float centerx, float centery, float radius)
        {
            List<MobState> filtered = new List<MobState>();
            foreach(MobState mob in mobs)
            {
                if(mob.GetDistance(centerx, centery) <= radius)
                {
                    filtered.Add(mob);
                }
            }
            return filtered;
        }

        public List<MobState> FindMobsInRadius(float centerx, float centery, float radius, IMobCondition condition)
        {
            return FilterMobs(FindInRadius(Mobs, centerx, centery, radius), condition);
        }
        public List<MobState> FindEnemiesInRadius(float centerx, float centery, float radius, IMobCondition condition)
        {
            return FilterMobs(FindInRadius(Enemies, centerx, centery, radius), condition);
        }
        public List<MobState> FindPlayersInRadius(float centerx, float centery, float radius, IMobCondition condition)
        {
            return FilterMobs(FindInRadius(Players, centerx, centery, radius), condition);
        }
        #endregion Searching and Filtering
    }
}
