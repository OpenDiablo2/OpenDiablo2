using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core
{
    public sealed class EngineDataManager : IEngineDataManager
    {
        private readonly IMPQProvider mpqProvider;

        public List<LevelPreset> LevelPresets { get; internal set; }
        public List<LevelType> LevelTypes { get; internal set; }

        public EngineDataManager(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;

            LoadLevelPresets();
            LoadLevelTypes();
        }

        private void LoadLevelTypes()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.LevelType)
                .First()
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .ToArray()
                .Select(x => x.ToLevelType());

            LevelTypes = new List<LevelType>(data);
        }

        private void LoadLevelPresets()
        {
            var data = mpqProvider
                .GetTextFile(ResourcePaths.LevelPreset)
                .First()
                .Skip(1)
                .Where(x => !String.IsNullOrWhiteSpace(x))
                .Select(x => x.Split('\t'))
                .ToArray()
                .Select(x => x.ToLevelPreset());

            LevelPresets = new List<LevelPreset>(data);
        }
    }
}
