using System;
using System.Collections.Generic;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Core.GameState_
{
    public class MobManager : IMobManager
    {
        public HashSet<MobState> Mobs { get; private set; } = new HashSet<MobState>(); // all mobs (including players!)
        public HashSet<PlayerState> Players { get; private set; } = new HashSet<PlayerState>();
        public HashSet<EnemyState> Enemies { get; private set;} = new HashSet<EnemyState>();

        IEnumerable<MobState> IMobManager.Mobs => Mobs;
        IEnumerable<PlayerState> IMobManager.Players => Players;
        IEnumerable<EnemyState> IMobManager.Enemies => Enemies;

        private HashSet<int> IdsUsed = new HashSet<int>();

        #region Player Controls
        public void AddPlayer(PlayerState player)
        {
            Players.Add(player);
            AddMob(player);
        }
        public void RemovePlayer(PlayerState player)
        {
            Players.Remove(player);
            RemoveMob(player);
        }
        #endregion Player Controls

        #region Mob Controls
        public void AddMob(MobState mob)
        {
            Mobs.Add(mob);
            if (IdsUsed.Contains(mob.Id))
                throw new ApplicationException("Tried to insert an existing mob id!");
            IdsUsed.Add(mob.Id);
        }
        public void RemoveMob(MobState mob)
        {
            Mobs.Remove(mob);
            IdsUsed.Remove(mob.Id);
        }
        public int GetNextAvailableMobId()
        {
            for (var i = 1; i < int.MaxValue; i++)
                if (!IdsUsed.Contains(i))
                    return i;

            throw new ApplicationException("Ran out of IDs. How did this even happen?!");
        }
        #endregion Mob Controls

        #region Enemy Controls
        public void AddEnemy(EnemyState enemy)
        {
            Enemies.Add(enemy);
            AddMob(enemy);
        }
        public void RemoveEnemy(EnemyState enemy)
        {
            Enemies.Remove(enemy);
            RemoveMob(enemy);
        }
        #endregion Enemy Controls

    }
}
