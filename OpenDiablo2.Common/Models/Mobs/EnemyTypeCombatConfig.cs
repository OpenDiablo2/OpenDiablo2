using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyTypeCombatConfig : IEnemyTypeCombatConfig
    {
        public int Threat { get; private set; } // the higher this is, the higher priority mercs give to killing this first
        
        public string[] MissileForAttack { get; private set; } // 1-2, which missile is used per attack
        // important note: 65535 = NONE. See missiles.txt, corresponds to a row number -2
        public string[] MissileForSkill { get; private set; } // 1-4, which missile is used for a skill
        // important note: 65535 = NONE. See missiles.txt, corresponds to a row number -2
        public string MissileForCast { get; private set; } // see above, specifies a missile for Case
        public string MissileForSequence { get; private set; } // see above, specifies a missile for Sequence

        public bool IsMelee { get; private set; } // is this a melee attacker?

        // vvv old

        public int ElementalAttackMode { get; private set; } // 4 = on hit, rest are unknown?
        public eDamageTypes ElementalAttackType { get; private set; } // 1=fire, 2=lightning, 4=cold, 5=poison
        public int ElementalOverlayId { get; private set; } // see overlays.txt, corresponds to a row number -2
        public int ElementalChance { get; private set; } // chance of use on a normal attack, 0 = never 100 = always
        public int ElementalMinDamage { get; private set; }
        public int ElementalMaxDamage { get; private set; }
        public int ElementalDuration { get; private set; } // duration of effects like cold and poison

        
        public bool[] CanMoveAttack { get; private set; } // 1-2, can move while using attack 1 / 2
        public bool[] CanMoveSkill { get; private set; } // 1-4, can move while using skill 1/2/3/4
        public bool CanMoveCast { get; private set; } // can move while casting?

        
        public bool IsAttackable { get; private set; } // is this monster attackable?

        public int MeleeRange { get; private set; }

        public int[] SkillType { get; private set; } // 1-5, what type of skill is it? 0 = none
        public int[] SkillSequence { get; private set; } // 1-5 animation sequence number 
        // should not be 0 if skill is not 0; see animation column in skills.txt
        public int[] SkillLevel { get; private set; } // 1-5, level of the skill

        public bool IsUndeadWithPhysicalAttacks { get; private set; } // if true, is an undead attacking with physical attacks
        public bool IsUndeadWithMagicAttacks { get; private set; } // if true '' magic atttacks
        public bool UsesMagicAttacks { get; private set; }

        public int ChanceToBlock { get; private set; }

        public bool DoesDeathDamage { get; private set; } // if true, does damage when it dies
                                                          // like undead flayers (TODO: understand exactly how this works)

        public bool IgnoredBySummons { get; private set; } // if true, is ignored by summons in combat

        // (87) A1Move	A2Move	S1Move	S2Move	S3Move	S4Move	Cmove
        // (109) Skill1	Skill1Seq	Skill1Lvl	Skill2	Skill2Seq	Skill2Lvl	Skill3	Skill3Seq	Skill3Lvl	Skill4	Skill4Seq	Skill4Lvl	Skill5	Skill5Seq	Skill5Lvl
        // (146) eLUndead	eHUndead	eDemon	eMagicUsing	eLarge	eSmall	eFlying	eOpenDoors	eSpawnCol	eBoss


        public EnemyTypeCombatConfig(
            int Threat,
            string[] MissileForAttack, string[] MissileForSkill, string MissileForCast, string MissileForSequence,
            bool IsMelee,


            int ElementalAttackMode, eDamageTypes ElementalAttackType, int ElementalOverlayId,
            int ElementalChance, int ElementalMinDamage, int ElementalMaxDamage, int ElementalDuration,
            
            bool[] CanMoveAttack, bool[] CanMoveSkill, bool CanMoveCast,
             bool IsAttackable,
            int MeleeRange,
            int[] SkillType, int[] SkillSequence, int[] SkillLevel,
            bool IsUndeadWithPhysicalAttacks, bool IsUndeadWithMagicAttacks, bool UsesMagicAttacks,
            int ChanceToBlock,
            bool DoesDeathDamage,
            bool IgnoredBySummons
            )
        {
            this.Threat = Threat;

            this.MissileForAttack = MissileForAttack;
            this.MissileForSkill = MissileForSkill;
            this.MissileForCast = MissileForCast;
            this.MissileForSequence = MissileForSequence;

            this.IsMelee = IsMelee;


            // vvv
            this.ElementalAttackMode = ElementalAttackMode;
            this.ElementalAttackType = ElementalAttackType;
            this.ElementalOverlayId = ElementalOverlayId;
            this.ElementalChance = ElementalChance;
            this.ElementalMinDamage = ElementalMinDamage;
            this.ElementalMaxDamage = ElementalMaxDamage;
            this.ElementalDuration = ElementalDuration;

            

            this.CanMoveAttack = CanMoveAttack;
            this.CanMoveSkill = CanMoveSkill;
            this.CanMoveCast = CanMoveCast;

            
            this.IsAttackable = IsAttackable;

            this.MeleeRange = MeleeRange;

            this.SkillType = SkillType;
            this.SkillSequence = SkillSequence;
            this.SkillLevel = SkillLevel;

            this.IsUndeadWithPhysicalAttacks = IsUndeadWithPhysicalAttacks;
            this.IsUndeadWithMagicAttacks = IsUndeadWithMagicAttacks;
            this.UsesMagicAttacks = UsesMagicAttacks;

            this.ChanceToBlock = ChanceToBlock;

            this.DoesDeathDamage = DoesDeathDamage;

            this.IgnoredBySummons = IgnoredBySummons;
        }
    }
}
