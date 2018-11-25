using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models
{
    public sealed class LevelDetail
    {
        public string Name { get; set; }
        public int Id { get; set; }
        public int Pal { get; set; }
        public int Act { get; set; }
        public int Layer { get; set; }
        public int SizeX { get; set; }
        public int SizeY { get; set; }
        public int OffsetX { get; set; }
        public int OffsetY { get; set; }
        public int Depend { get; set; }
        public int Rain { get; set; }
        public int Mud { get; set; }
        public int NoPer { get; set; }
        public int LOSDraw { get; set; }
        public int FloorFilter { get; set; }
        public int BlankScreen { get; set; }
        public int DrawEdges { get; set; }
        public int IsInside { get; set; }
        public int DrlgType { get; set; }
        public int LevelType { get; set; }
        public int SubType { get; set; }
        public int SubTheme { get; set; }
        public int SubWaypoint { get; set; }
        public int SubShrine { get; set; }
        public int Vis0 { get; set; }
        public int Vis1 { get; set; }
        public int Vis2 { get; set; }
        public int Vis3 { get; set; }
        public int Vis4 { get; set; }
        public int Vis5 { get; set; }
        public int Vis6 { get; set; }
        public int Vis7 { get; set; }
        public int Warp0 { get; set; }
        public int Warp1 { get; set; }
        public int Warp2 { get; set; }
        public int Warp3 { get; set; }
        public int Warp4 { get; set; }
        public int Warp5 { get; set; }
        public int Warp6 { get; set; }
        public int Warp7 { get; set; }
        public int Intensity { get; set; }
        public int Red { get; set; }
        public int Green { get; set; }
        public int Blue { get; set; }
        public int Portal { get; set; }
        public int Position { get; set; }
        public int SaveMonsters { get; set; }
        public int Quest { get; set; }
        public int WarpDist { get; set; }
        public int MonLvl1 { get; set; }
        public int MonLvl2 { get; set; }
        public int MonLvl3 { get; set; }
        public int MonDen { get; set; }
        public int MonUMin { get; set; }
        public int MonUMax { get; set; }
        public int MonWndr { get; set; }
        public int MonSpcWalk { get; set; }
        public int Mtot { get; set; }
        public int[] M1_25 { get; set; }
        public int[] S1_25 { get; set; }
        public int Utot { get; set; }
        public int[] U1_25 { get; set; }
        public int[] C1_5 { get; set; }
        public int[] CA1_5 { get; set; }
        public int[] CD1_5 { get; set; }
        public int Themes { get; set; }
        public int SoundEnv { get; set; }
        public int Waypoint { get; set; }
        public string LevelName { get; set; }
        public string LevelWarp { get; set; }
        public string EntryFile { get; set; }
        public int[] ObjGrp0_7 { get; set; }
        public int[] ObjPrb0_7 { get; set; }
        public bool Beta { get; set; }
    }

    public static class LevelDetailHelper
    {
        public static LevelDetail ToLevelDetail(this string[] v)
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
            result.LevelType       = Convert.ToInt32(v[i++]);
            result.SubType         = Convert.ToInt32(v[i++]);
            result.SubTheme        = Convert.ToInt32(v[i++]);
            result.SubWaypoint     = Convert.ToInt32(v[i++]);
            result.SubShrine       = Convert.ToInt32(v[i++]);
            result.Vis0            = Convert.ToInt32(v[i++]);
            result.Vis1            = Convert.ToInt32(v[i++]);
            result.Vis2            = Convert.ToInt32(v[i++]);
            result.Vis3            = Convert.ToInt32(v[i++]);
            result.Vis4            = Convert.ToInt32(v[i++]);
            result.Vis5            = Convert.ToInt32(v[i++]);
            result.Vis6            = Convert.ToInt32(v[i++]);
            result.Vis7            = Convert.ToInt32(v[i++]);
            result.Warp0           = Convert.ToInt32(v[i++]);
            result.Warp1           = Convert.ToInt32(v[i++]);
            result.Warp2           = Convert.ToInt32(v[i++]);
            result.Warp3           = Convert.ToInt32(v[i++]);
            result.Warp4           = Convert.ToInt32(v[i++]);
            result.Warp5           = Convert.ToInt32(v[i++]);
            result.Warp6           = Convert.ToInt32(v[i++]);
            result.Warp7           = Convert.ToInt32(v[i++]);
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
            result.Beta = Convert.ToInt32(v[i++]) == 1;


            return result;
        }
    }
}
