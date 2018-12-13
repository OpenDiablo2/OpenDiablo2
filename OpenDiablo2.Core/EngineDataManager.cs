﻿using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Core
{
    public sealed class EngineDataManager : IEngineDataManager
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IMPQProvider mpqProvider;

        public ImmutableList<LevelDetail> Levels { get; internal set; }
        public ImmutableList<LevelPreset> LevelPresets { get; internal set; }
        public ImmutableList<LevelType> LevelTypes { get; internal set; }

        public ImmutableList<Item> Items { get; internal set; }
        public ImmutableDictionary<eHero, ILevelExperienceConfig> ExperienceConfigs { get; internal set; }
        public ImmutableDictionary<eHero, IHeroTypeConfig> HeroTypeConfigs { get; internal set; }
        public ImmutableList<IEnemyTypeConfig> EnemyTypeConfigs { get; internal set; } 

        public EngineDataManager(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;

            LoadLevelDetails();
            LoadCharacterData();
            LoadEnemyData();

            Items = LoadItemData();
        }


        private void LoadLevelDetails()
        {
            log.Info("Loading level types");
            LevelTypes = mpqProvider
                .GetTextFile(ResourcePaths.LevelType)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x.Count() == 36 && x[0] != "Expansion")
                .Select(x => x.ToLevelType())
                .ToImmutableList();

            
            log.Info("Loading level presets");
            LevelPresets = mpqProvider
                .GetTextFile(ResourcePaths.LevelPreset)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x.Count() == 24 && x[0] != "Expansion")
                .Select(x => x.ToLevelPreset())
                .ToImmutableList();
            
            log.Info("Loading level details");
            Levels = mpqProvider
                .GetTextFile(ResourcePaths.LevelDetails)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x.Count() > 80 && x[0] != "Expansion")
                .Select(x => x.ToLevelDetail(LevelPresets, LevelTypes))
                .ToImmutableList();
        }

        private ImmutableList<Item> LoadItemData()
            => new List<Item>()
                .Concat(LoadWeaponData())
                .Concat(LoadArmorData())
                .Concat(LoadMiscData())
                .ToImmutableList();

        private IEnumerable<Weapon> LoadWeaponData()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.Weapons)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                //.Where(x => !String.IsNullOrWhiteSpace(x[27]))
                .Select(x => x.ToWeapon());

            return data;
        }

        private IEnumerable<Armor> LoadArmorData()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.Armor)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                //.Where(x => !String.IsNullOrWhiteSpace(x[27]))
                .Select(x => x.ToArmor());

            return data;
        }

        private IEnumerable<Misc> LoadMiscData()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.Misc)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                //.Where(x => !String.IsNullOrWhiteSpace(x[27]))
                .Select(x => x.ToMisc());

            return data;
        }

        private void LoadCharacterData()
        {
            ExperienceConfigs = LoadExperienceConfig();
            HeroTypeConfigs = LoadHeroTypeConfig();
        }

        private ImmutableDictionary<eHero, ILevelExperienceConfig> LoadExperienceConfig()
            => mpqProvider
                .GetTextFile(ResourcePaths.Experience)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .ToArray()
                .ToLevelExperienceConfigs();

        private ImmutableDictionary<eHero, IHeroTypeConfig> LoadHeroTypeConfig()
            => mpqProvider
                .GetTextFile(ResourcePaths.CharStats)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x[0] != "Expansion")
                .ToArray()
                .ToImmutableDictionary(x => (eHero)Enum.Parse(typeof(eHero), x[0]), x => x.ToHeroTypeConfig());

        private void LoadEnemyData()
        {
            EnemyTypeConfigs = LoadEnemyTypeConfig();
        }

        private ImmutableList<IEnemyTypeConfig> LoadEnemyTypeConfig()
            => mpqProvider
                .GetTextFile(ResourcePaths.MonStats)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x[0] != "Expansion" && x[0] != "unused")
                .ToArray()
                .Select(x => x.ToEnemyTypeConfig())
                .ToImmutableList();
    }
}
