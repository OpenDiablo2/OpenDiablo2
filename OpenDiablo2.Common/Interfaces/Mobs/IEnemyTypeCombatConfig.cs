using OpenDiablo2.Common.Enums.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IEnemyTypeCombatConfig
    {
        int ElementalAttackMode { get;} // 4 = on hit, rest are unknown?
        eDamageTypes ElementalAttackType { get;} // 1=fire, 2=lightning, 4=cold, 5=poison
        int ElementalOverlayId { get;} // see overlays.txt, corresponds to a row number -2
        int ElementalChance { get;} // chance of use on a normal attack, 0 = never 100 = always
        int ElementalMinDamage { get;}
        int ElementalMaxDamage { get;}
        int ElementalDuration { get;} // duration of effects like cold and poison

        // TODO: these could be switched to SHORTS, is there any benefit to that?
        int[] MissileForAttack { get;} // 1-2, which missile is used per attack
        // important note: 65535 = NONE. See missiles.txt, corresponds to a row number -2
        int[] MissileForSkill { get;} // 1-4, which missile is used for a skill
        // important note: 65535 = NONE. See missiles.txt, corresponds to a row number -2
        int MissileForCase { get;} // see above, specifies a missile for Case
        int MissileForSequence { get;} // see above, specifies a missile for Sequence

        bool[] CanMoveAttack { get;} // 1-2, can move while using attack 1 / 2
        bool[] CanMoveSkill { get;} // 1-4, can move while using skill 1/2/3/4
        bool CanMoveCast { get;} // can move while casting?

        bool IsMelee { get;} // is this a melee attacker?
        bool IsAttackable { get;} // is this monster attackable?

        int MeleeRange { get; }

        int[] SkillType { get;} // 1-5, what type of skill is it? 0 = none
        int[] SkillSequence { get;} // 1-5 animation sequence number 
        // should not be 0 if skill is not 0; see animation column in skills.txt
        int[] SkillLevel { get;} // 1-5, level of the skill

        bool IsUndeadWithPhysicalAttacks { get;} // if true, is an undead attacking with physical attacks
        bool IsUndeadWithMagicAttacks { get;} // if true '' magic atttacks
        bool UsesMagicAttacks { get;}

        int ChanceToBlock { get;}

        bool DoesDeathDamage { get;} // if true, does damage when it dies
                                                          // like undead flayers (TODO: understand exactly how this works)

        bool IgnoredBySummons { get;} // if true, is ignored by summons in combat

    }
}
