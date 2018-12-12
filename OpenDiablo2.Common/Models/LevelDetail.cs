using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class LevelDetail
    {
        /// <summary>Internal level name</summary>
        public string Name { get; internal set; }

        /// <summary>Level ID (Used in columns like VIS0-1)</summary>
        public int Id { get; internal set; }

        /// <summary>Palette</summary>
        public int Pal { get; internal set; }

        /// <summary>Act</summary>
        public int Act { get; internal set; }

        /// <summary>What layer the level is on (surface levels are always 0)</summary>
        public int Layer { get; internal set; }

        /// <summary>Horizontal size of the level</summary>
        public int SizeX { get; internal set; }

        /// <summary>Vertical size of the level</summary>
        public int SizeY { get; internal set; }

        /// <summary>Horizontal Placement Offset</summary>
        public int OffsetX { get; internal set; }

        /// <summary>Vertical placement offset</summary>
        public int OffsetY { get; internal set; }

        /// <summary>Special setting for levels that aren't random or preset (like Outer Cloister and Arcane Sancturary)</summary>
        public int Depend { get; internal set; }

        /// <summary>If true, it rains (or snows)</summary>
        public int Rain { get; internal set; }

        /// <summary>Unused</summary>
        public int Mud { get; internal set; }

        /// <summary>Perspective mode forced to off if set to 1</summary>
        public int NoPer { get; internal set; }

        /// <summary>Level of sight drawing</summary>
        public int LOSDraw { get; internal set; }

        /// <summary>Unknown</summary>
        public int FloorFilter { get; internal set; }

        /// <summary>Unknown</summary>
        public int BlankScreen { get; internal set; }

        /// <summary>For levels bordered with mountains or walls</summary>
        public int DrawEdges { get; internal set; }

        /// <summary>Set to 1 if this is underground or inside</summary>
        public int IsInside { get; internal set; }

        /// <summary> Setting for Level Generation: 1=Random Size, amount of rooms defined by LVLMAZE.TXT, 2=pre set map (example: catacombs lvl4) and 3=Random Area with preset size (wildernesses)</summary>
        public int DrlgType { get; internal set; }

        /// <summary>The level id to reference in lvltypes.txt</summary>
        public int LevelTypeId { get; internal set; }

        /// <summary>Setting Regarding Level Type for lvlsub.txt (6=wilderness, 9=desert etc, -1=no subtype)</summary>
        public int SubType { get; internal set; }

        /// <summary></summary>
        public int SubTheme { get; internal set; }

        /// <summary></summary>
        public int SubWaypoint { get; internal set; }

        /// <summary></summary>
        public int SubShrine { get; internal set; }

        /// <summary>Entry/Exit to level #1-#8</summary>
        public int[] Vis0_7 { get; internal set; }

        /// <summary>ID into lvlwarp.txt</summary>
        public int[] Warp0_7 { get; internal set; }

        /// <summary>Light intensity (0-255)</summary>
        public int Intensity { get; internal set; }

        /// <summary></summary>
        public int Red { get; internal set; }

        /// <summary></summary>
        public int Green { get; internal set; }

        /// <summary></summary>
        public int Blue { get; internal set; }

        /// <summary>Unknown</summary>
        public int Portal { get; internal set; }

        /// <summary>Settings for preset levels</summary>
        public int Position { get; internal set; }

        /// <summary>If true, monster/creatures get saved/loaded instead of despawning.</summary>
        public int SaveMonsters { get; internal set; }

        /// <summary>Quest flags</summary>
        public int Quest { get; internal set; }

        /// <summary>Usually 2025, unknown</summary>
        public int WarpDist { get; internal set; }

        /// <summary>Level on Normal (controls the item level of items that drop from chests etc)</summary>
        public int MonLvl1 { get; internal set; }

        /// <summary>Level on Nightmare (controls the item level of items that drop from chests etc)</summary>
        public int MonLvl2 { get; internal set; }

        /// <summary> Level on Hell (controls the item level of items that drop from chests etc)</summary>
        public int MonLvl3 { get; internal set; }

        /// <summary>The Density of Monsters</summary>
        public int MonDen { get; internal set; }

        /// <summary>Minimum Unique and Champion Monsters Spawned in this Level</summary>
        public int MonUMin { get; internal set; }

        /// <summary>Maximum Unique and Champion Monsters Spawned in this Level</summary>
        public int MonUMax { get; internal set; }

        /// <summary></summary>
        public int MonWndr { get; internal set; }

        /// <summary></summary>
        public int MonSpcWalk { get; internal set; }

        /// <summary>How many different Species of Monsters can occur in this area (example: if you use M1-25 then set Mtot to 25 etc)</summary>
        public int Mtot { get; internal set; }

        /// <summary>Monster Species 1-25 (use ID from MonStats.txt)</summary>
        public int[] M1_25 { get; internal set; }

        /// <summary>Spawned Species 1-25 (use ID from MonStats for Monsters which have eSpawnCol set to 2, related to M1-25, eg: if M1 Spawns S1 will Spawn)</summary>
        public int[] S1_25 { get; internal set; }

        /// <summary>How many different Species of Monsters can spawn as Uniques and Champions in this Area (works like Mtot)</summary>
        public int Utot { get; internal set; }

        /// <summary> Unique Species 1-25 (same as M1-M25 just for Monsters that you want to appear as Unique/Champions)</summary>
        public int[] U1_25 { get; internal set; }

        /// <summary>Critter Species 1-5 (For monsters set to 1 in the IsCritter Column in MonStats.txt)</summary>
        public int[] C1_5 { get; internal set; }

        /// <summary>Related to C1-5, eg: if you spawn a critter thru C1 then set this column to 30 etc. (function unknown)</summary>
        public int[] CA1_5 { get; internal set; }

        /// <summary>Unknown</summary>
        public int[] CD1_5 { get; internal set; }

        /// <summary>unknown</summary>
        public int Themes { get; internal set; }

        /// <summary>Referes to a entry in SoundEnviron.txt (for the Levels Music)</summary>
        public int SoundEnv { get; internal set; }

        /// <summary>255=No way Point, other #'s Waypoint ID</summary>
        public int Waypoint { get; internal set; }

        /// <summary>String Code for the Display name of the Level</summary>
        public string LevelName { get; internal set; }

        /// <summary>String Code for the Display name of a entrance to this Level</summary>
        public string LevelWarp { get; internal set; }

        /// <summary>Which *.DC6 Title Image is loaded when you enter this area</summary>
        public string EntryFile { get; internal set; }

        /// <summary>Use the ID of the ObjectGroup you want to Spawn in this Area (from ObjectGroups.txt)</summary>
        public int[] ObjGrp0_7 { get; internal set; }

        /// <summary>Object Spawn Possibility: the Chance for this object to occur (if you use ObjGrp0 then set ObjPrb0 to a value below 100)</summary>
        public int[] ObjPrb0_7 { get; internal set; }

        /// <summary>Unused</summary>
        public bool Beta { get; internal set; }

        public LevelPreset LevelPreset { get; internal set; }

        public LevelType LevelType { get; internal set; }
    }

    public static class LevelDetailHelper
    {
        public static LevelDetail ToLevelDetail(this string[] v, IEnumerable<LevelPreset> levelPresets, IEnumerable<LevelType> levelTypes)
        {
            var result = new LevelDetail();
            int i = 0;
            result.Name            = v[i++];
            result.Id              = Convert.ToInt32(v[i++]);
            result.Pal             = Convert.ToInt32(v[i++]);
            result.Act             = Convert.ToInt32(v[i++]);
            result.Layer           = Convert.ToInt32(v[i++]);
            result.SizeX           = Convert.ToInt32(v[i++]);
            result.SizeY           = Convert.ToInt32(v[i++]);
            result.OffsetX         = Convert.ToInt32(v[i++]);
            result.OffsetY         = Convert.ToInt32(v[i++]);
            result.Depend          = Convert.ToInt32(v[i++]);
            result.Rain            = Convert.ToInt32(v[i++]);
            result.Mud             = Convert.ToInt32(v[i++]);
            result.NoPer           = Convert.ToInt32(v[i++]);
            result.LOSDraw         = Convert.ToInt32(v[i++]);
            result.FloorFilter     = Convert.ToInt32(v[i++]);
            result.BlankScreen     = Convert.ToInt32(v[i++]);
            result.DrawEdges       = Convert.ToInt32(v[i++]);
            result.IsInside        = Convert.ToInt32(v[i++]);
            result.DrlgType        = Convert.ToInt32(v[i++]);
            result.LevelTypeId     = Convert.ToInt32(v[i++]);
            result.SubType         = Convert.ToInt32(v[i++]);
            result.SubTheme        = Convert.ToInt32(v[i++]);
            result.SubWaypoint     = Convert.ToInt32(v[i++]);
            result.SubShrine       = Convert.ToInt32(v[i++]);
            result.Vis0_7 = new int[8];
            for (int j = 0; j < 8; j++) result.Vis0_7[j] = Convert.ToInt32(v[i++]);
            result.Warp0_7 = new int[8];
            for (int j = 0; j < 8; j++) result.Warp0_7[j] = Convert.ToInt32(v[i++]);
            result.Intensity       = Convert.ToInt32(v[i++]);
            result.Red             = Convert.ToInt32(v[i++]);
            result.Green           = Convert.ToInt32(v[i++]);
            result.Blue            = Convert.ToInt32(v[i++]);
            result.Portal          = Convert.ToInt32(v[i++]);
            result.Position        = Convert.ToInt32(v[i++]);
            result.SaveMonsters    = Convert.ToInt32(v[i++]);
            result.Quest           = Convert.ToInt32(v[i++]);
            result.WarpDist        = Convert.ToInt32(v[i++]);
            result.MonLvl1         = Convert.ToInt32(v[i++]);
            result.MonLvl2         = Convert.ToInt32(v[i++]);
            result.MonLvl3         = Convert.ToInt32(v[i++]);
            result.MonDen          = Convert.ToInt32(v[i++]);
            result.MonUMin         = Convert.ToInt32(v[i++]);
            result.MonUMax         = Convert.ToInt32(v[i++]);
            result.MonWndr         = Convert.ToInt32(v[i++]);
            result.MonSpcWalk      = Convert.ToInt32(v[i++]);
            result.Mtot            = Convert.ToInt32(v[i++]);
            result.M1_25 = new int[25];
            for (int j = 0; j < 25; j++) result.M1_25[j] = Convert.ToInt32(v[i++]);
            result.S1_25 = new int[25];
            for (int j = 0; j < 25; j++) result.S1_25[j] = Convert.ToInt32(v[i++]);
            result.Utot            = Convert.ToInt32(v[i++]);
            result.U1_25 = new int[25];
            for (int j = 0; j < 25; j++) result.U1_25[j] = Convert.ToInt32(v[i++]);
            result.C1_5 = new int[5];
            for (int j = 0; j < 5; j++) result.C1_5[j] = Convert.ToInt32(v[i++]);
            result.CA1_5 = new int[5];
            for (int j = 0; j < 5; j++) result.CA1_5[j] = Convert.ToInt32(v[i++]);
            result.CD1_5 = new int[5];
            for (int j = 0; j < 5; j++) result.CD1_5[j] = Convert.ToInt32(v[i++]);
            result.Themes          = Convert.ToInt32(v[i++]);
            result.SoundEnv        = Convert.ToInt32(v[i++]);
            result.Waypoint        = Convert.ToInt32(v[i++]);
            result.LevelName       = v[i++];
            result.LevelWarp       = v[i++];
            result.EntryFile       = v[i++];
            result.ObjGrp0_7 = new int[8];
            for (int j = 0; j < 5; j++) result.ObjGrp0_7[j] = Convert.ToInt32(v[i++]);
            result.ObjPrb0_7 = new int[8];
            for (int j = 0; j < 5; j++) result.ObjPrb0_7[j] = Convert.ToInt32(v[i++]);
            result.Beta = Convert.ToInt32(v[i]) == 1;
            result.LevelPreset = levelPresets.FirstOrDefault(x => x.LevelId == result.Id);
            result.LevelType = levelTypes.FirstOrDefault(x => x.Id == result.LevelTypeId);
            return result;
        }
    }
}
