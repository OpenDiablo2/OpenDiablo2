﻿using OpenDiablo2.Common.Enums;
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
        private List<long> ExperiencePerLevel = new List<long>();

        public LevelExperienceConfig(List<long> expperlevel)
        {
            ExperiencePerLevel = expperlevel;
        }

        public long GetTotalExperienceForLevel(int level)
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

    public static class LevelExperienceConfigHelper
    {
        public static Dictionary<eHero, ILevelExperienceConfig> ToLevelExperienceConfigs(this string[][] data)
        {
            Dictionary<eHero, ILevelExperienceConfig> result = new Dictionary<eHero, ILevelExperienceConfig>();
            for (int i = 1; i < data[0].Length; i++)
            {
                // i starts at 1 because we want to skip the first column
                // the first column is just the row titles
                string heroname = data[i][0]; // first row is the hero name
                // LEFT OFF HERE:
                //need to figure out ehero from heroname
                eHero herotype = eHero.Amazon;
                try
                {
                    herotype = (eHero)Enum.Parse(typeof(eHero), heroname);
                }
                catch
                {
                    continue; // skip this hero if we can't parse the name into a valid hero type
                }
                int maxlevel = -1;
                try
                {
                    maxlevel = int.Parse(data[1][i]);
                }
                catch
                {
                    maxlevel = -1; // we don't need to fail in this case since maxlevel 
                    // can be inferred from the number of experience listings
                }
                List<long> expperlevel = new List<long>();
                for (int o = 2; o < data.Length && (o-2 < maxlevel || maxlevel == -1); o++)
                {
                    try
                    {
                        expperlevel.Add(long.Parse(data[o][i]));
                    }
                    catch
                    {
                        throw new Exception("Could not parse experience number '" + data[o][i] + "'.");
                    }
                }
                result.Add(herotype, new LevelExperienceConfig(expperlevel));
            }

            return result;
        }
    }
}
