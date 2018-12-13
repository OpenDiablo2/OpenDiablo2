using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Models.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.UT
{
    [TestClass]
    public class UT_EnemyState
    {
        public EnemyState MakeEnemyState(int id, eDifficulty difficulty, Random rand)
        {
            // does not correspond to any particular enemy in the actual data, just for testing
            EnemyTypeConfig config = new EnemyTypeConfig(
                InternalName: "TestEnemy",
                Name: "TestEnemy1",
                Type: "Skeleton", 
                Descriptor: "",
                BaseId: 1,
                PopulateId: 1,
                Spawned: true,
                Beta: false,
                Code: "SK",
                ClientOnly: false,
                NoMap: false,
                SizeX: 2,
                SizeY: 2,
                Height: 3,
                NoOverlays: false,
                OverlayHeight: 2,
                Velocity: 3,
                RunVelocity: 6,
                CanStealFrom: false,
                ColdEffect: 50,
                Rarity: true,
                MinimumGrouping: 2,
                MaximumGrouping: 5,
                BaseWeapon: "1hs",
                AIParams: new int[] {75, 85, 9, 50, 0 },
                Allied: false,
                IsNPC: false,
                IsCritter: false,
                CanEnterTown: false,
                HealthRegen: 80,
                IsDemon: false,
                IsLarge: false,
                IsSmall: false,
                IsFlying: false,
                CanOpenDoors: false,
                SpawningColumn: 0,
                IsBoss: false,
                IsInteractable: false,
                IsKillable: true,
                CanBeConverted: true,
                HitClass: 3,
                HasSpecialEndDeath: false,
                DeadCollision: false,
                CanBeRevivedByOtherMonsters: true,
                AppearanceConfig: new EnemyTypeAppearanceConfig(
                    HasDeathAnimation: true,
                    HasNeutralAnimation: true,
                    HasWalkAnimation: true,
                    HasGetHitAnimation: true,
                    HasAttack1Animation: true,
                    HasAttack2Animation: false,
                    HasBlockAnimation: true,
                    HasCastAnimation: false,
                    HasSkillAnimation: new bool[] { false, false, false, false },
                    HasCorpseAnimation: false, 
                    HasKnockbackAnimation: true,
                    HasRunAnimation: true,
                    HasLifeBar: true,
                    HasNameBar: false,
                    CannotBeSelected: false,
                    CanCorpseBeSelected: false,
                    BleedType: 0,
                    HasShadow: true,
                    LightRadius: 0,
                    HasUniqueBossColors: false,
                    CompositeDeath: true,
                    LightRGB: new byte[] { 255, 255, 255 }
                    ),
                CombatConfig: new EnemyTypeCombatConfig(
                    ElementalAttackMode: 4,
                    ElementalAttackType: eDamageTypes.PHYSICAL,
                    ElementalOverlayId: 0,
                    ElementalChance: 0,
                    ElementalMinDamage: 0,
                    ElementalMaxDamage: 0,
                    ElementalDuration: 0,
                    MissileForAttack: new int[] { 0, 0 },
                    MissileForSkill: new int[] { 0, 0, 0, 0, },
                    MissileForCase: 0,
                    MissileForSequence: 0,
                    CanMoveAttack: new bool[] { false, false },
                    CanMoveSkill: new bool[] { false, false, false, false },
                    CanMoveCast: false,
                    IsMelee: true,
                    IsAttackable: true,
                    MeleeRange: 0,
                    SkillType: new int[] { 0, 0, 0, 0 },
                    SkillSequence: new int[] { 0, 0, 0, 0 },
                    SkillLevel: new int[] { 0, 0, 0, 0 },
                    IsUndeadWithPhysicalAttacks: true,
                    IsUndeadWithMagicAttacks: false,
                    UsesMagicAttacks: false,
                    ChanceToBlock: 20,
                    DoesDeathDamage: false,
                    IgnoredBySummons: false
                    ),
                NormalDifficultyConfig: new EnemyTypeDifficultyConfig(
                    Level: 5,
                    DamageResist: 0,
                    MagicResist: 0,
                    FireResist: 0.5,
                    LightResist: 0,
                    ColdResist: 0,
                    PoisonResist: 0,
                    MinHP: 10,
                    MaxHP: 15,
                    AC: 3,
                    Exp: 100,
                    AttackMinDamage: new int[] { 5, 0 },
                    AttackMaxDamage: new int[] { 10, 0 },
                    AttackChanceToHit: new int[] { 35, 0 },
                    Skill1MinDamage: 0,
                    Skill1MaxDamage: 0,
                    Skill1ChanceToHit: 0,
                    TreasureClass: new string[] { "Swarm 1", "Act 2 Champ A", "Act 2 Unique A", "" }
                    ),
                NightmareDifficultyConfig: new EnemyTypeDifficultyConfig(
                    Level: 10,
                    DamageResist: 0,
                    MagicResist: 0,
                    FireResist: 0.5,
                    LightResist: 0.5,
                    ColdResist: 0,
                    PoisonResist: 0,
                    MinHP: 20,
                    MaxHP: 25,
                    AC: 3,
                    Exp: 1000,
                    AttackMinDamage: new int[] { 5, 0 },
                    AttackMaxDamage: new int[] { 10, 0 },
                    AttackChanceToHit: new int[] { 35, 0 },
                    Skill1MinDamage: 0,
                    Skill1MaxDamage: 0,
                    Skill1ChanceToHit: 0,
                    TreasureClass: new string[] { "Swarm 1", "Act 3 Champ A", "Act 3 Unique A", "" }
                    ),
                HellDifficultyConfig: new EnemyTypeDifficultyConfig(
                    Level: 15,
                    DamageResist: 0,
                    MagicResist: 0,
                    FireResist: 0.5,
                    LightResist: 0.5,
                    ColdResist: 0.5,
                    PoisonResist: 0,
                    MinHP: 30,
                    MaxHP: 45,
                    AC: 3,
                    Exp: 10000,
                    AttackMinDamage: new int[] { 5, 0 },
                    AttackMaxDamage: new int[] { 10, 0 },
                    AttackChanceToHit: new int[] { 35, 0 },
                    Skill1MinDamage: 0,
                    Skill1MaxDamage: 0,
                    Skill1ChanceToHit: 0,
                    TreasureClass: new string[] { "Swarm 2", "Act 4 Champ A", "Act 4 Unique A", "" }
                    )
                );

            EnemyState en = new EnemyState("Fallen", id, 0, 0, config, difficulty, rand);

            return en;
        }

        [TestMethod]
        public void EnemyDifficultyTest()
        {
            Random rand = new Random();
            EnemyState en = MakeEnemyState(1, eDifficulty.NORMAL, rand);
            Assert.AreEqual(5, en.Level);
            Assert.AreEqual(100, en.ExperienceGiven);

            EnemyState en2 = MakeEnemyState(2, eDifficulty.NIGHTMARE, rand);
            Assert.AreEqual(10, en2.Level);
            Assert.AreEqual(1000, en2.ExperienceGiven);

            EnemyState en3 = MakeEnemyState(3, eDifficulty.HELL, rand);
            Assert.AreEqual(15, en3.Level);
            Assert.AreEqual(10000, en3.ExperienceGiven);
        }
    }
}
