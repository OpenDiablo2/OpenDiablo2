using System.Collections.Generic;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IMobManager
    {
        IEnumerable<MobState> Mobs { get; }
        IEnumerable<PlayerState> Players { get; }
        IEnumerable<EnemyState> Enemies { get; }

        void AddPlayer(PlayerState player);
        void RemovePlayer(PlayerState player);

        void AddMob(MobState mob);
        void RemoveMob(MobState mob);
        int GetNextAvailableMobId();

        void AddEnemy(EnemyState enemy);
        void RemoveEnemy(EnemyState enemy);
    }
}
