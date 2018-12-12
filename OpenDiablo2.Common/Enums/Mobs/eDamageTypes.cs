using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums.Mobs
{
    public enum eDamageTypes
    {
        NONE = -1, // no resistances apply
        PHYSICAL = 0,
        MAGIC = 3, //1=fire, 2=lightning, 4=cold, 5=poison
        FIRE = 1,
        COLD = 4,
        LIGHTNING = 2,
        POISON = 5,
        LIFE_STEAL = 6,
        MANA_STEAL = 7,
        STAMINA_STEAL = 8,
        STUN = 9,
        RANDOM = 10, // random between fire/cold/lightning/poison
        BURN = 11,
        FREEZE = 12,
    }
}
