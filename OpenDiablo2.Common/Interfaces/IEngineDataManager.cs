using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models;
using System.Collections.Generic;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IEngineDataManager
    {
        List<LevelDetail> Levels { get; }
        List<LevelPreset> LevelPresets { get; }
        List<LevelType> LevelTypes { get; }
        List<Item> Items { get; }
        Dictionary<eHero, ILevelExperienceConfig> ExperienceConfigs { get; }
        Dictionary<eHero, IHeroTypeConfig> HeroTypeConfigs { get; }
        List<IEnemyTypeConfig> EnemyTypeConfigs { get; }
        List<ObjectInfo> Objects { get; }
        List<ObjectTypeInfo> ObjectTypes { get; }
        Dictionary<int, IMissileTypeConfig> MissileTypeConfigs { get; }
        Dictionary<string, int> MissileTypeConfigsLookup { get; }
    }
}
