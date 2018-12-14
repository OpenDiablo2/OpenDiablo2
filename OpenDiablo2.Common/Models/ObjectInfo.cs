using System;

namespace OpenDiablo2.Common.Models
{
    public sealed class ObjectInfo
    {
        public string Name { get; internal set; }
        public int Id { get; internal set; }
        public string Token { get; internal set; }
        public int SpawnMax { get; internal set; }
        public bool[] Selectable0_7 { get; internal set; } = new bool[8];
        public int TrapProb { get; internal set; }
        public int SizeX { get; internal set; }
        public int SizeY { get; internal set; }
        public int nTgtFX { get; internal set; }
        public int nTgtFY { get; internal set; }
        public int nTgtBX { get; internal set; }
        public int nTgtBY { get; internal set; }
        public int[] FrameCnt0_7 { get; internal set; } = new int[8];
        public int[] FrameDelta0_7 { get; internal set; } = new int[8];
        public bool[] CycleAnim0_7 { get; internal set; } = new bool[8];
        public int[] Lit0_7 { get; internal set; } = new int[8];
        public bool[] BlocksLight0_7 { get; internal set; } = new bool[8];
        public bool[] HasCollision0_7 { get; internal set; } = new bool[8];
        public bool IsAttackable0 { get; internal set; }
        public int[] Start0_7 { get; internal set; } = new int[8];
        public bool EnvEffect { get; internal set; }
        public bool IsDoor { get; internal set; }
        public bool BlocksVis { get; internal set; }
        public int Orientation { get; internal set; }
        public int Trans { get; internal set; }
        public int[] OrderFlag0_7 { get; internal set; } = new int[8];
        public bool PreOperate { get; internal set; }
        public bool[] Mode0_7 { get; internal set; } = new bool[8];
        public int Yoffset { get; internal set; }
        public int Xoffset { get; internal set; }
        public int Draw { get; internal set; }
        public byte Red { get; internal set; }
        public byte Green { get; internal set; }
        public byte Blue { get; internal set; }
        public bool HD { get; internal set; }
        public bool TR { get; internal set; }
        public bool LG { get; internal set; }
        public bool RA { get; internal set; }
        public bool LA { get; internal set; }
        public bool RH { get; internal set; }
        public bool LH { get; internal set; }
        public bool SH { get; internal set; }
        public bool[] S1_8 { get; internal set; } = new bool[8]; // (S# equipment slot)
        public int TotalPieces { get; internal set; }
        public int SubClass { get; internal set; }
        public int Xspace { get; internal set; }
        public int Yspace { get; internal set; }
        public int NameOffset { get; internal set; }
        public bool MonsterOK { get; internal set; }
        public int OperateRange { get; internal set; }
        public bool ShrineFunction { get; internal set; }
        public bool Restore { get; internal set; }
        public int[] Parm0_7 { get; internal set; } = new int[8];
        public int Act { get; internal set; }
        public bool Lockable { get; internal set; }
        public bool Gore { get; internal set; }
        public bool Sync { get; internal set; }
        public bool Flicker { get; internal set; }
        public int Damage { get; internal set; }
        public bool Beta { get; internal set; }
        public bool Overlay { get; internal set; }
        public bool CollisionSubst { get; internal set; }
        public int Left { get; internal set; }
        public int Top { get; internal set; }
        public int Width { get; internal set; }
        public int Height { get; internal set; }
        public int OperateFn { get; internal set; }
        public int PopulateFn { get; internal set; }
        public int InitFn { get; internal set; }
        public int ClientFn { get; internal set; }
        public bool RestoreVirgins { get; internal set; }
        public bool BlockMissile { get; internal set; }
        public int DrawUnder { get; internal set; }
        public bool OpenWarp { get; internal set; }
        public int AutoMap { get; internal set; }
    }

    public static class ObjectInfoHelper
    {
        public static ObjectInfo ToObjectInfo(this string[] row)
        {
            var result = new ObjectInfo();
            var idx = 0;
            result.Name= row[idx++];
            idx++; // Description is 'unused'
            result.Id = Convert.ToInt32(row[idx++]);
            result.Token = row[idx++];
            result.SpawnMax = Convert.ToInt32(row[idx++]);
            for (var i = 0; i < 8; i++)
                result.Selectable0_7[i] = Convert.ToInt32(row[idx++]) == 1;
            result.TrapProb = Convert.ToInt32(row[idx++]);
            result.SizeX = Convert.ToInt32(row[idx++]);
            result.SizeY = Convert.ToInt32(row[idx++]);
            result.nTgtFX = Convert.ToInt32(row[idx++]);
            result.nTgtFY = Convert.ToInt32(row[idx++]);
            result.nTgtBX = Convert.ToInt32(row[idx++]);
            result.nTgtBY = Convert.ToInt32(row[idx++]);
            for (var i = 0; i < 8; i++)
                result.FrameCnt0_7[i] = Convert.ToInt32(row[idx++]);
            for (var i = 0; i < 8; i++)
                result.FrameDelta0_7[i] = Convert.ToInt32(row[idx++]);
            for (var i = 0; i < 8; i++)
                result.CycleAnim0_7[i] = Convert.ToInt32(row[idx++]) == 1;
            for (var i = 0; i < 8; i++)
                result.Lit0_7[i] = Convert.ToInt32(row[idx++]);
            for (var i = 0; i < 8; i++)
                result.BlocksLight0_7[i] = Convert.ToInt32(row[idx++]) == 1;
            for (var i = 0; i < 8; i++)
                result.HasCollision0_7[i] = Convert.ToInt32(row[idx++]) == 1;
            result.IsAttackable0 = Convert.ToInt32(row[idx++]) == 1;
            for (var i = 0; i < 8; i++)
                result.Start0_7[i] = Convert.ToInt32(row[idx++]);
            result.EnvEffect = Convert.ToInt32(row[idx++]) == 1;
            result.IsDoor = Convert.ToInt32(row[idx++]) == 1;
            result.BlocksVis = Convert.ToInt32(row[idx++]) == 1;
            result.Orientation = Convert.ToInt32(row[idx++]);
            result.Trans = Convert.ToInt32(row[idx++]);
            for (var i = 0; i < 8; i++)
                result.OrderFlag0_7[i] = Convert.ToInt32(row[idx++]);
            result.PreOperate = Convert.ToInt32(row[idx++]) == 1;
            for (var i = 0; i < 8; i++)
                result.Mode0_7[i] = Convert.ToInt32(row[idx++]) == 1;
            result.Yoffset = Convert.ToInt32(row[idx++]);
            result.Xoffset = Convert.ToInt32(row[idx++]);
            result.Draw = Convert.ToInt32(row[idx++]);
            result.Red = Convert.ToByte(row[idx++]);
            result.Green = Convert.ToByte(row[idx++]);
            result.Blue = Convert.ToByte(row[idx++]);
            result.HD = Convert.ToInt32(row[idx++]) == 1;
            result.TR = Convert.ToInt32(row[idx++]) == 1;
            result.LG = Convert.ToInt32(row[idx++]) == 1;
            result.RA = Convert.ToInt32(row[idx++]) == 1;
            result.LA = Convert.ToInt32(row[idx++]) == 1;
            result.RH = Convert.ToInt32(row[idx++]) == 1;
            result.LH = Convert.ToInt32(row[idx++]) == 1;
            result.SH = Convert.ToInt32(row[idx++]) == 1;
            for (var i = 0; i < 8; i++)
                result.S1_8[i] = Convert.ToInt32(row[idx++]) == 1;
            result.TotalPieces = Convert.ToInt32(row[idx++]);
            result.SubClass = Convert.ToInt32(row[idx++]);
            result.Xspace = Convert.ToInt32(row[idx++]);
            result.Yspace = Convert.ToInt32(row[idx++]);
            result.NameOffset = Convert.ToInt32(row[idx++]);
            result.MonsterOK = Convert.ToInt32(row[idx++]) == 1;
            result.OperateRange = Convert.ToInt32(row[idx++]);
            result.ShrineFunction = Convert.ToInt32(row[idx++]) == 1;
            result.Restore = Convert.ToInt32(row[idx++]) == 1;
            for (var i = 0; i < 8; i++)
                result.Parm0_7[i] = Convert.ToInt32(row[idx++]);
            result.Act = Convert.ToInt32(row[idx++]);
            result.Lockable = Convert.ToInt32(row[idx++]) == 1;
            result.Gore = Convert.ToInt32(row[idx++]) == 1;
            result.Sync = Convert.ToInt32(row[idx++]) == 1;
            result.Flicker = Convert.ToInt32(row[idx++]) == 1;
            result.Damage = Convert.ToInt32(row[idx++]);
            result.Beta = Convert.ToInt32(row[idx++]) == 1;
            result.Overlay = Convert.ToInt32(row[idx++]) == 1;
            result.CollisionSubst = Convert.ToInt32(row[idx++]) == 1;
            result.Left = Convert.ToInt32(row[idx++]);
            result.Top = Convert.ToInt32(row[idx++]);
            result.Width = Convert.ToInt32(row[idx++]);
            result.Height = Convert.ToInt32(row[idx++]);
            result.OperateFn = Convert.ToInt32(row[idx++]);
            result.PopulateFn = Convert.ToInt32(row[idx++]);
            result.InitFn = Convert.ToInt32(row[idx++]);
            result.ClientFn = Convert.ToInt32(row[idx++]);
            result.RestoreVirgins = Convert.ToInt32(row[idx++]) == 1;
            result.BlockMissile = Convert.ToInt32(row[idx++]) == 1;
            result.DrawUnder = Convert.ToInt32(row[idx++]);
            result.OpenWarp = Convert.ToInt32(row[idx++]) == 1;
            result.AutoMap = Convert.ToInt32(row[idx++]);
            return result;
        }
    }
}
