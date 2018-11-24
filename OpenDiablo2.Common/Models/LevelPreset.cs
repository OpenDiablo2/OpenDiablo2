using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    //
    public sealed class LevelPreset
    {
        public string Name { get; set; }
        public int Def { get; set; }
        public int LevelId { get; set; }
        public bool Populate { get; set; }
        public int Logicals { get; set; }
        public int Outdoors { get; set; }
        public int Animate { get; set; }
        public int KillEdge { get; set; }
        public int FillBlanks { get; set; }
        public int SizeX { get; set; }
        public int SizeY { get; set; }
        public int AutoMap { get; set; }
        public int Scan { get; set; }
        public int Pops { get; set; }
        public int PopPad { get; set; }
        public int Files { get; set; }
        public string File1 { get; set; }
        public string File2 { get; set; }
        public string File3 { get; set; }
        public string File4 { get; set; }
        public string File5 { get; set; }
        public string File6 { get; set; }
        public UInt32 Dt1Mask { get; set; }
        public bool Beta { get; set; }
    }

    public static class LevelPresetHelper
    {
        public static LevelPreset ToLevelPreset(this string[] row)
            => new LevelPreset
            {
                Name = row[0],
                Def = Convert.ToInt32(row[1]),
                LevelId = row[2] == "#REF!" ? 0 : Convert.ToInt32(row[2]), // TODO: Why is this a thing?
                Populate = Convert.ToInt32(row[3]) == 1,
                Logicals = Convert.ToInt32(row[4]),
                Outdoors = Convert.ToInt32(row[5]),
                Animate = Convert.ToInt32(row[6]),
                KillEdge = Convert.ToInt32(row[7]),
                FillBlanks = Convert.ToInt32(row[8]),
                SizeX = Convert.ToInt32(row[9]),
                SizeY = Convert.ToInt32(row[10]),
                AutoMap = Convert.ToInt32(row[11]),
                Scan = Convert.ToInt32(row[12]),
                Pops = Convert.ToInt32(row[13]),
                PopPad = Convert.ToInt32(row[14]),
                Files = Convert.ToInt32(row[15]),
                File1 = row[16],
                File2 = row[17],
                File3 = row[18],
                File4 = row[19],
                File5 = row[20],
                File6 = row[21],
                Dt1Mask = Convert.ToUInt32(row[22]),
                Beta = Convert.ToInt32(row[23]) == 1,
            };
    }   
}
