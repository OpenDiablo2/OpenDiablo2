using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces.Mobs;
using OpenDiablo2.Common.Models;
using System.Collections.Immutable;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IEngineDataManager
    {
        ImmutableList<LevelDetail> Levels { get; }
        ImmutableList<LevelPreset> LevelPresets { get; }
        ImmutableList<LevelType> LevelTypes { get; }
        ImmutableList<Item> Items { get; }
        ImmutableDictionary<eHero, ILevelExperienceConfig> ExperienceConfigs { get; }
        ImmutableDictionary<eHero, IHeroTypeConfig> HeroTypeConfigs { get; }
        ImmutableList<IEnemyTypeConfig> EnemyTypeConfigs { get; }
        ImmutableList<ObjectInfo> Objects { get; }
        ImmutableList<ObjectTypeInfo> ObjectTypes { get; }
        ImmutableDictionary<int, IMissileTypeConfig> MissileTypeConfigs { get; }
        ImmutableDictionary<string, int> MissileTypeConfigsLookup { get; }
    }
}
