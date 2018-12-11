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
        public string Name { get; internal set; }
        public int Def { get; internal set; }
        public int LevelId { get; internal set; }
        public bool Populate { get; internal set; }
        public int Logicals { get; internal set; }
        public int Outdoors { get; internal set; }
        public int Animate { get; internal set; }
        public int KillEdge { get; internal set; }
        public int FillBlanks { get; internal set; }
        public int SizeX { get; internal set; }
        public int SizeY { get; internal set; }
        public int AutoMap { get; internal set; }
        public int Scan { get; internal set; }
        public int Pops { get; internal set; }
        public int PopPad { get; internal set; }
        public int Files { get; internal set; }
        public string File1 { get; internal set; }
        public string File2 { get; internal set; }
        public string File3 { get; internal set; }
        public string File4 { get; internal set; }
        public string File5 { get; internal set; }
        public string File6 { get; internal set; }
        public UInt32 Dt1Mask { get; internal set; }
        public bool Beta { get; internal set; }
    }

    public static class LevelPresetHelper
    {
        public static LevelPreset ToLevelPreset(this string[] row)
            => new LevelPreset
            {
                Name = row[0],
                Def = Convert.ToInt32(row[1]),
                LevelId = Convert.ToInt32(row[2]),
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
                Beta = Convert.ToInt32(row[23]) == 1
            };
    }
}
