using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class LevelExperienceConfig : ILevelExperienceConfig
    {
        private List<int> ExperiencePerLevel = new List<int>();

        public LevelExperienceConfig(List<int> expperlevel)
        {
            ExperiencePerLevel = expperlevel;
        }

        public int GetTotalExperienceForLevel(int level)
        {
            if(ExperiencePerLevel.Count <= level)
            {
                return -1; // note: a value of -1 means this level is unattainable!
            }
            return ExperiencePerLevel[level];
        }

        public int GetMaxLevel()
        {
            return ExperiencePerLevel.Count - 1;
        }
    }
}
