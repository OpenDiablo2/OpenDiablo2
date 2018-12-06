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
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IMPQProvider mpqProvider;

        public List<LevelPreset> LevelPresets { get; internal set; }
        public List<LevelType> LevelTypes { get; internal set; }
        public List<LevelDetail> LevelDetails { get; internal set; }

        public EngineDataManager(IMPQProvider mpqProvider)
        {
            this.mpqProvider = mpqProvider;

            LoadLevelPresets();
            LoadLevelTypes();
            LoadLevelDetails();
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
                .ToArray()
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
                .ToArray()
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
                .ToArray()
                .Select(x => x.ToLevelDetail());

            LevelDetails = new List<LevelDetail>(data);
        }
    }
}
