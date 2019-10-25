using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyTypeDifficultyConfig : IEnemyTypeDifficultyConfig
    {
        public int Level { get; private set; }

        public int AiDelay { get; private set; } // Delay between AI Ticks, lower = they act faster
        public int AiActivationDistance { get; private set; } // distance at which ai becomes active
        // if left blank, should default to 35
        public int[] AiParams { get; private set; } // 1-8, used in AI Functions

        //vvv

        public double DamageResist { get; private set; }
        public double MagicResist { get; private set; }
        public double FireResist { get; private set; }
        public double LightningResist { get; private set; }
        public double ColdResist { get; private set; }
        public double PoisonResist { get; private set; }

        public int MinHP { get; private set; }
        public int MaxHP { get; private set; }
        public int AC { get; private set; } // armor class
        public int Exp { get; private set; }

        public int[] AttackMinDamage { get; private set; } // 1-2, min damage attack can roll
        public int[] AttackMaxDamage { get; private set; } // 1-2 max damage attack can roll
        public int[] AttackChanceToHit { get; private set; } // 1-2 chance attack has to hit (out of 100??)
        public int Skill1MinDamage { get; private set; } // min damage skill 1 can do (why only skill 1?)
        public int Skill1MaxDamage { get; private set; } // max damage for skill 1
        public int Skill1ChanceToHit { get; private set; } // chance skill has to hit

        public string[] TreasureClass { get; private set; } // 1-4
        
        public EnemyTypeDifficultyConfig (int Level,
            int AiDelay, int AiActivationDistance, int[] AiParams,

            double DamageResist, double MagicResist, double FireResist, double LightResist,
            double ColdResist, double PoisonResist,
            int MinHP, int MaxHP, int AC, int Exp,
            int[] AttackMinDamage, int[] AttackMaxDamage, int[] AttackChanceToHit,
            int Skill1MinDamage, int Skill1MaxDamage, int Skill1ChanceToHit,
            string[] TreasureClass)
        {
            this.Level = Level;

            this.AiDelay = AiDelay;
            if(AiActivationDistance == 0)
            {
                AiActivationDistance = 35; // should default to 35 if not specified
            }
            this.AiActivationDistance = AiActivationDistance;
            this.AiParams = AiParams;



            this.DamageResist = DamageResist;
            this.MagicResist = MagicResist;
            this.FireResist = FireResist;
            this.LightningResist = LightResist;
            this.ColdResist = ColdResist;
            this.PoisonResist = PoisonResist;

            this.MinHP = MinHP;
            this.MaxHP = MaxHP;
            this.AC = AC;
            this.Exp = Exp;

            this.AttackMinDamage = AttackMinDamage;
            this.AttackMaxDamage = AttackMaxDamage;
            this.AttackChanceToHit = AttackChanceToHit;
            this.Skill1MinDamage = Skill1MinDamage;
            this.Skill1MaxDamage = Skill1MaxDamage;
            this.Skill1ChanceToHit = Skill1ChanceToHit;

            this.TreasureClass = TreasureClass;
        }
    }
}
