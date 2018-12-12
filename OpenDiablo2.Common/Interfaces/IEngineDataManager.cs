using System.Collections.Generic;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IEngineDataManager
    {
        List<LevelPreset> LevelPresets { get; }
        List<LevelType> LevelTypes { get; }
        List<LevelDetail> LevelDetails { get; }
        List<Item> Items { get; }
        Dictionary<eHero, ILevelExperienceConfig> ExperienceConfigs { get; }
        Dictionary<eHero, IHeroTypeConfig> HeroTypeConfigs { get; }
        List<IEnemyTypeConfig> EnemyTypeConfigs { get; }
    }
}
