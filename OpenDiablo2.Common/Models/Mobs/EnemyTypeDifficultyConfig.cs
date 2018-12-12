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
        // (127) DamageResist	MagicResist	FireResist	LightResist	ColdResist	PoisonResist
        // (133) DamageResist(N)	MagicResist(N)	FireResist(N)	LightResist(N)	ColdResist(N)	PoisonResist(N)
        // (139) DamageResist(H)	MagicResist(H)	FireResist(H)	LightResist(H)	ColdResist(H)	PoisonResist(H)
        // (158) MinHP	MaxHP	AC	Exp	ToBlock
        // (163) A1MinD	A1MaxD	A1ToHit	A2MinD	A2MaxD	A2ToHit	S1MinD	S1MaxD	S1ToHit
        // (175) MinHP(N)	MaxHP(N)	AC(N)	Exp(N)	A1MinD(N)	A1MaxD(N)	A1ToHit(N)	A2MinD(N)	A2MaxD(N)	A2ToHit(N)	S1MinD(N)	S1MaxD(N)	S1ToHit(N)
        // (188) MinHP(H)	MaxHP(H)	AC(H)	Exp(H)	A1MinD(H)	A1MaxD(H)	A1ToHit(H)	A2MinD(H)	A2MaxD(H)	A2ToHit(H)	S1MinD(H)	S1MaxD(H)	S1ToHit(H)

        // (201) TreasureClass1	TreasureClass2	TreasureClass3	TreasureClass4
        // (205) TreasureClass1(N)	TreasureClass2(N)	TreasureClass3(N)	TreasureClass4(N)
        // (209) TreasureClass1(H)	TreasureClass2(H)	TreasureClass3(H)	TreasureClass4(H)

        public int Level { get; private set; }

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
            double DamageResist, double MagicResist, double FireResist, double LightResist,
            double ColdResist, double PoisonResist,
            int MinHP, int MaxHP, int AC, int Exp,
            int[] AttackMinDamage, int[] AttackMaxDamage, int[] AttackChanceToHit,
            int Skill1MinDamage, int Skill1MaxDamage, int Skill1ChanceToHit,
            string[] TreasureClass)
        {
            this.Level = Level;

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
