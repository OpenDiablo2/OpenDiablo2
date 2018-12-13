using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IEnemyTypeDifficultyConfig
    {
        int Level { get; }

        double DamageResist { get; }
        double MagicResist { get; }
        double FireResist { get; }
        double LightningResist { get; }
        double ColdResist { get; }
        double PoisonResist { get; }

        int MinHP { get; }
        int MaxHP { get; }
        int AC { get; } // armor class
        int Exp { get; }

        int[] AttackMinDamage { get; } // 1-2, min damage attack can roll
        int[] AttackMaxDamage { get; } // 1-2 max damage attack can roll
        int[] AttackChanceToHit { get; } // 1-2 chance attack has to hit (out of 100??)
        int Skill1MinDamage { get; } // min damage skill 1 can do (why only skill 1?)
        int Skill1MaxDamage { get; } // max damage for skill 1
        int Skill1ChanceToHit { get; } // chance skill has to hit

        string[] TreasureClass { get; } // 1-4
    }
}
