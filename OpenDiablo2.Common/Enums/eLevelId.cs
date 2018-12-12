using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eLevelId
    {
        None,
        Act1_Town1 = 1,
        Act1_CaveTreasure2 = 13,
        Act1_CaveTreasure3 = 14,
        Act1_CaveTreasure4 = 15,
        Act1_CaveTreasure5 = 16,
        Act1_CryptCountessX = 25,
        Act1_Tower2 = 20,
        Act1_MonFront = 26,
        Act1_Courtyard1 = 27,
        Act1_Courtyard2 = 32,
        Act1_Cathedral = 33,
        Act1_Andariel = 37,
        Act1_Tristram = 38,
        Act2_Town = 40,
        Act2_Harem = 50,
        Act2_DurielsLair = 73,
        Act3_Town = 75,
        Act3_DungeonTreasure1 = 90,
        Act3_DungeonTreasure2 = 91,
        Act3_SewerTreasureX = 93,
        Act3_Temple1 = 94,
        Act3_Temple2 = 95,
        Act3_Temple3 = 96,
        Act3_Temple4 = 97,
        Act3_Temple5 = 98,
        Act3_Temple6 = 99,
        Act3_MephistoComplex = 102,
        Act4_Fortress = 103,
        Act5_Town = 109,
        Act5_TempleFinalRoom = 124,
        Act5_ThroneRoom = 131,
        Act5_WorldStone = 132,
        Act5_TempleEntrance = 121,
        Act5_BaalEntrance = 120,
    }

    public static class ELevelIdHelper
    {
        public static string GenerateEnum(List<LevelPreset> levelPresets)
        {
            var output = new StringBuilder();
            foreach (LevelPreset lp in levelPresets)
            {
                // need to convert e.g. 'Act 1 - Town 1' to 'Act1_Town'
                if (lp.LevelId == 0)
                {
                    continue;
                }
                string name = lp.Name.Replace(" - ", "_").Replace(" ", "").Replace("'", "");
                output.AppendLine($"{name}={lp.LevelId}");
            }
            return output.ToString();
        }
    }
}
