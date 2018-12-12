﻿using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyState : MobState
    {
        public IEnemyTypeConfig EnemyConfig { get; protected set; }

        public int ExperienceGiven { get; protected set; }
        
        public Stat[] AttackRating { get; protected set; }
        public Stat DefenseRating { get; protected set; }
        public Stat HealthRegen { get; protected set; }

        public EnemyState() : base() { }

        public EnemyState(string name, int id, float x, float y,
            IEnemyTypeConfig EnemyConfig, eDifficulty Difficulty, Random Rand)
            : base(name, id, 0, 0, x, y)
        {
            this.EnemyConfig = EnemyConfig;
            IEnemyTypeDifficultyConfig difficulty = EnemyConfig.GetDifficultyConfig(Difficulty);

            int maxhp = Rand.Next(difficulty.MinHP, difficulty.MaxHP + 1);
            Health = new Stat(0, maxhp, maxhp, true);
            ExperienceGiven = difficulty.Exp;
            Level = difficulty.Level;

            // stats
            AttackRating = new Stat[2];
            AttackRating[0] = new Stat(0, difficulty.AttackChanceToHit[0], difficulty.AttackChanceToHit[0], false);
            AttackRating[1] = new Stat(0, difficulty.AttackChanceToHit[1], difficulty.AttackChanceToHit[1], false);

            DefenseRating = new Stat(0, EnemyConfig.CombatConfig.ChanceToBlock, EnemyConfig.CombatConfig.ChanceToBlock, false);

            HealthRegen = new Stat(0, EnemyConfig.HealthRegen, EnemyConfig.HealthRegen, true);

            // handle immunities / resistances
            SetResistance(eDamageTypes.COLD, difficulty.ColdResist);
            SetResistance(eDamageTypes.FIRE, difficulty.FireResist);
            SetResistance(eDamageTypes.MAGIC, difficulty.MagicResist);
            SetResistance(eDamageTypes.PHYSICAL, difficulty.DamageResist);
            SetResistance(eDamageTypes.POISON, difficulty.PoisonResist);
            SetResistance(eDamageTypes.LIGHTNING, difficulty.LightningResist);
            // interestingly, monsters don't actually have immunity flags
            // TODO: should we treat a resistance of 1 differently?
            // should a resistance of 1 be an immunity (e.g. the main difference is that
            // effects / debuffs couldn't reduce the resistance below 100%)

            // handle flags
            if (EnemyConfig.IsBoss)
            {
                AddFlag(eMobFlags.BOSS);
            }
            if (EnemyConfig.IsCritter)
            {
                AddFlag(eMobFlags.CRITTER);
            }
            if (EnemyConfig.IsNPC)
            {
                AddFlag(eMobFlags.NPC);
            }
            if (EnemyConfig.IsLarge)
            {
                AddFlag(eMobFlags.LARGE);
            }
            if (EnemyConfig.IsSmall)
            {
                AddFlag(eMobFlags.SMALL);
            }
            if (EnemyConfig.IsInteractable)
            {
                AddFlag(eMobFlags.INTERACTABLE);
            }
            if (!EnemyConfig.IsKillable)
            {
                AddFlag(eMobFlags.INVULNERABLE);
            }
            if (EnemyConfig.IsDemon)
            {
                AddFlag(eMobFlags.DEMON);
            }
            if (EnemyConfig.CombatConfig.IgnoredBySummons)
            {
                AddFlag(eMobFlags.IGNORED_BY_SUMMONS);
            }
            AddFlag(eMobFlags.ENEMY);
        }
    }
}
