/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using OpenDiablo2.Common.Attributes;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;

namespace OpenDiablo2.Common.Enums
{
    public static class SkillsHelper
    {
        public static SkillInfoAttribute GetSkillInfo(this eSkill skill)
        {
            var attr = skill.GetType().GetCustomAttributes(typeof(SkillInfoAttribute), false).FirstOrDefault() as SkillInfoAttribute;
            Debug.Assert(attr != null, $"Skill {skill} does not contain SkillInfo attribute.");
            return attr;
        }

        public static IEnumerable<eSkill> GetHeroSkills(eHero hero)
        {
            foreach(var skill in (eSkill[])Enum.GetValues(typeof(eSkill)))
            {
                if(skill.GetSkillInfo()?.Hero == hero)
                    yield return skill;
            }
        }
    }

    //public int[] eSkillLevelTier
    //{
    //    1,
    //    6,
    //    12,
    //    18,
    //    24,
    //    30
    //};

    public enum eSkill
    {
        /// Generic Skills
        [SkillInfo(eHero.None, 1)]
        Attack = 0,
        [SkillInfo(eHero.None, 3)]
        Throw = 2,
        [SkillInfo(eHero.None, 2)]
        Unsummon = 3,

        /// Amazon
        // Javelin and Spear Skills
        [SkillInfo(eHero.Amazon, 4, 1)]
        Jab = 10,
        [SkillInfo(eHero.Amazon, 8, 6, Jab)]
        PowerStrike = 14,
        [SkillInfo(eHero.Amazon, 9, 6)]
        PoisonJavelin = 15,
        [SkillInfo(eHero.Amazon, 13, 12, Jab)]
        Impale = 19,
        [SkillInfo(eHero.Amazon, 14, 12, PoisonJavelin)]
        LightningBolt = 20,
        [SkillInfo(eHero.Amazon, 18, 18, LightningBolt, PowerStrike)]
        ChargedStrike = 24,
        [SkillInfo(eHero.Amazon, 19, 18, LightningBolt)]
        PlagueJavelin = 25,
        [SkillInfo(eHero.Amazon, 24, 24, Impale)]
        Fend = 30,
        [SkillInfo(eHero.Amazon, 28, 30, ChargedStrike)]
        LightningStrike = 34,
        [SkillInfo(eHero.Amazon, 29, 30, PlagueJavelin)]
        LightningFury = 35,
        // Passive and Magic Skills
        [SkillInfo(eHero.Amazon, 2, 1)]
        InnerSight = 8,
        [SkillInfo(eHero.Amazon, 3, 1)]
        CriticalStrike = 9,
        [SkillInfo(eHero.Amazon, 7, 6)]
        Dodge = 13,
        [SkillInfo(eHero.Amazon, 11, 12, InnerSight)]
        SlowMissles = 17,
        [SkillInfo(eHero.Amazon, 12, 12, Dodge)]
        Avoid = 18,
        [SkillInfo(eHero.Amazon, 17, 18, CriticalStrike)]
        Penetrate = 23,
        [SkillInfo(eHero.Amazon, 22, 24, SlowMissles)]
        Decoy = 28, // called Dopplezon
        [SkillInfo(eHero.Amazon, 23, 24, Avoid)]
        Evade = 29,
        [SkillInfo(eHero.Amazon, 26, 30, Decoy, Evade)]
        Valkyrie = 32,
        [SkillInfo(eHero.Amazon, 27, 30, Penetrate)]
        Pierce = 33,
        // Bow and Crossbow Skills
        [SkillInfo(eHero.Amazon, 0, 1)]
        MagicArrow = 6,
        [SkillInfo(eHero.Amazon, 1, 1)]
        FireArrow = 7,
        [SkillInfo(eHero.Amazon, 5, 6)]
        ColdArrow = 11,
        [SkillInfo(eHero.Amazon, 6, 6, MagicArrow)]
        MultipleShot = 12,
        [SkillInfo(eHero.Amazon, 10, 12, FireArrow)]
        ExplodingArrow = 16,
        [SkillInfo(eHero.Amazon, 15, 18, ColdArrow)]
        IceArrow = 21,
        [SkillInfo(eHero.Amazon, 16, 18, MultipleShot)]
        GuidedArrow = 22,
        [SkillInfo(eHero.Amazon, 20, 24, GuidedArrow)]
        Strafe = 26,
        [SkillInfo(eHero.Amazon, 21, 24, ExplodingArrow)]
        ImmolationArrow = 27,
        [SkillInfo(eHero.Amazon, 25, 30, IceArrow)]
        FreezingArrow = 31,

        /// Assassin
        // Martial Arts
        TigerStrike,
        DragonTalon,
        FistsOfFire,
        DragonClaw,
        CobraStrike,
        ClawsOfThunder,
        DragonTail,
        BladesOfIce,
        DragonFlight,
        PhoenixStrike,
        // Shadow Disciplines
        ClawMastery,
        PsychicHammer,
        BurstOfSpeed,
        WeaponBlock,
        CloakOfShadows,
        Fade,
        ShadowWarrior,
        MindBlast,
        Venom,
        ShadowMaster,
        // Traps
        FireBlast,
        ShockWeb,
        BladeSentinel,
        ChargedBoltSentry,
        WakeOfFire,
        BladeFury,
        LightningSentry,
        WakeOfInferno,
        DeathSentry,
        BladeShield,

        /// Necromancer
        // Summoning Spells
        SkeletonMastery,
        RaiseSkeleton,
        ClayGolem,
        GolemMastery,
        RaiseSkeletalMage,
        BloodGolem,
        SummonResist,
        IronGolem,
        FireGolem,
        Revive,
        // Poison and Bone Spells
        Teeth,
        BoneArmor,
        PoisonDagger,
        CorpseExplosion,
        BoneWall,
        PoisonExplosion,
        BoneSpear,
        BonePrision,
        PoisonNova,
        BoneSpirit,
        // Curses
        AmplifyDamage,
        DimVision,
        Weaken,
        IronMaiden,
        Terror,
        Confuse,
        LifeTap,
        Attract,
        Decrepify,
        LowerResist,

        /// Barbarian
        // Warcries
        Howl,
        FindPotion,
        Taunt,
        Shout,
        FindItem,
        BattleCry,
        BattleOrders,
        GrimWard,
        WarCry,
        BattleCommand,
        // Combat Masteries
        SwordMastery,
        AxeMastery,
        MaceMastery,
        PoleArmMastery,
        ThrowingMastery,
        SpearMastery,
        IncreasedStamina,
        IronSkin,
        IncreasedSpeed,
        NaturalResistance,
        // Combat Skills
        Bash,
        Leap,
        DoubleSwing,
        Stun,
        DoubleThrow,
        LeapAttack,
        Concentrate,
        Frenzy,
        Whirlwind,
        Berserk,

        /// Paladin
        // Defensive Auras
        Prayer,
        ResistFire,
        Defiance,
        ResistCold,
        Cleansing,
        ResistLightning,
        Vigor,
        Meditation,
        Redemption,
        Salvation,
        // Offensive Auras
        Might,
        HolyFire,
        Thorns,
        BlessedAim,
        Concentration,
        HolyFreeze,
        HolyShock,
        Sanctuary,
        Fanaticism,
        Conviction,
        // Combat Skills
        Sacrafice,
        Smite,
        HolyBolt,
        Zeal,
        Charge,
        Vengeance,
        BlessedHammer,
        Conversion,
        HolyShield,
        FistOfTheHeavens,

        /// Sorceress
        // Cold Spells
        IceBolt,
        FrozenArmor,
        FrostNova,
        IceBlast,
        ShiverArmor,
        GlacialSpike,
        Blizzard,
        ChillingArmor,
        FrozenOrb,
        ColdMastery,
        // Lightning Spells
        ChargedBolt,
        StaticField,
        Telekinesis,
        Nova,
        Lightning,
        ChainLightning,
        Teleport,
        ThunderStorm,
        EnergyShield,
        LightningMastery,
        // Fire Spells
        FireBolt,
        Warmth,
        Inferno,
        Blaze,
        FireBall,
        FireWall,
        Enchant,
        Meteor,
        FireMastery,
        Hydra,
        
        /// Druid
        // Elemental
        Firestorm,
        MoltenBoulder,
        ArcticBlast,
        Fissure,
        CycloneArmor,
        Twister,
        Volcano,
        Tornado,
        Armageddon,
        Hurricane,
        // Shape Shifting
        Werewolf,
        Lycanthropy,
        Werebear,
        FeralRage,
        Maul,
        Rabies,
        FireClaw,
        Hunger,
        ShockWave,
        Fury,
        // Summoning
        Raven,
        PoisonCreeper,
        OakSage,
        SummonSpiritWolf,
        CarrionVine,
        heartOfWolverine,
        SummonDireWolf,
        SolarCreeper,
        SpiritOfBarbs,
        SummonGrizzly
    }
}
