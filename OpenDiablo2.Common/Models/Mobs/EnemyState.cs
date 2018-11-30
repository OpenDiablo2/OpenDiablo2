using OpenDiablo2.Common.Enums.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyState : MobState
    {
        public int ExperienceGiven { get; protected set; }

        public EnemyState(string name, int id, int level, int maxhealth, float x, float y, int experiencegiven)
            : base(name, id, level, maxhealth, x, y)
        {
            ExperienceGiven = experiencegiven;
            AddFlag(eMobFlags.ENEMY);
        }
    }
}
