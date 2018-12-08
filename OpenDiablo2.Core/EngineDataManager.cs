using System;
using System.Collections.Generic;
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

        public List<LevelPreset> LevelPresets { get; internal set; }
        public List<LevelType> LevelTypes { get; internal set; }
        public List<LevelDetail> LevelDetails { get; internal set; }
        public List<Item> Items { get; internal set; } = new List<Item>();
        public Dictionary<eHero, ILevelExperienceConfig> ExperienceConfigs { get; internal set; } = new Dictionary<eHero, ILevelExperienceConfig>();
        public Dictionary<eHero, IHeroTypeConfig> HeroTypeConfigs { get; internal set; } = new Dictionary<eHero, IHeroTypeConfig>();

        public EngineDataManager(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;

            LoadLevelPresets();
            LoadLevelTypes();
            LoadLevelDetails();

            LoadItemData();

            LoadCharacterData();
        }

        private void LoadLevelTypes()
        {
            log.Info("Loading level types");
            var data = mpqProvider
                .GetTextFile(ResourcePaths.LevelType)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x.Count() == 36 && x[0] != "Expansion")
                .Select(x => x.ToLevelType());

            LevelTypes = new List<LevelType>(data);
        }

        private void LoadLevelPresets()
        {
            log.Info("Loading level presets");
            var data = mpqProvider
                .GetTextFile(ResourcePaths.LevelPreset)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x.Count() == 24 && x[0] != "Expansion")
                .Select(x => x.ToLevelPreset());

            LevelPresets = new List<LevelPreset>(data);
        }

        private void LoadLevelDetails()
        {
            log.Info("Loading level details");
            var data = mpqProvider
                .GetTextFile(ResourcePaths.LevelDetails)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x.Count() > 80 && x[0] != "Expansion")
                .Select(x => x.ToLevelDetail());

            LevelDetails = new List<LevelDetail>(data);
        }

        private void LoadItemData()
        {
            var weaponData = LoadWeaponData();
            var armorData = LoadArmorData();
            var miscData = LoadMiscData();

            Items.AddRange(weaponData);
            Items.AddRange(armorData);
            Items.AddRange(miscData);
        }

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
            LoadExperienceConfig();
            LoadHeroTypeConfig();
        }

        private void LoadExperienceConfig()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.Experience)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .ToArray()
                .ToLevelExperienceConfigs();
            
            ExperienceConfigs = data;
        }

        private void LoadHeroTypeConfig()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.CharStats)
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .Where(x => x[0] != "Expansion")
                .ToDictionary(x => (eHero)Enum.Parse(typeof(eHero),x[0]), x => x.ToHeroTypeConfig());

            HeroTypeConfigs = data;
        }
    }
}
