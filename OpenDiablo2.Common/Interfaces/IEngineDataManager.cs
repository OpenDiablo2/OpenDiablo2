using System.Collections.Generic;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IEngineDataManager
    {
        List<LevelPreset> LevelPresets { get; }
        List<LevelType> LevelTypes { get; }
        List<LevelDetail> LevelDetails { get; }
    }
}
