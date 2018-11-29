using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyState : MobState
    {
        public EnemyState(string name, int id, int maxhealth, int maxmana, int maxstamina, float x, float y)
            : base(name, id, maxhealth, maxmana, maxstamina, x, y)
        {

        }
    }
}
