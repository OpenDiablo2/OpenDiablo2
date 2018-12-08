using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eMobMode
    {
        PlayerDeath,
        PlayerNeutral,
        PlayerWalk,
        PlayerRun,
        PlayerGetHit,
        PlayerTownNeutral,
        PlayerTownWalk,
        PlayerAttack1,
        PlayerAttack2,
        PlayerBlock,
        PlayerCast,
        PlayerThrow,
        PlayerKick,
        PlayerSkill1,
        PlayerSkill2,
        PlayerSkill3,
        PlayerSkill4,
        PlayerDead,
        PlayerSequence,
        PlayerKnockBack,
        MonsterDeath,
        MonsterNeutral,
        MonsterWalk,
        MonsterGetHit,
        MonsterAttack1,
        MonsterAttack2,
        MonsterBlock,
        MonsterCast,
        MonsterSkill1,
        MonsterSkill2,
        MonsterSkill3,
        MonsterSkill4,
        MonsterDead,
        MonsterKnockback,
        MonsterSequence,
        MonsterRun,
        ObjectNeutral,
        ObjectOperating,
        ObjectOpened,
        ObjectSpecial1,
        ObjectSpecial2,
        ObjectSpecial3,
        ObjectSpecial4,
        ObjectSpecial5


    }

    public static class eMobModeExtensions
    {
        private static readonly Dictionary<eMobMode, string> mobModes = new Dictionary<eMobMode, string>
        {
            { eMobMode.PlayerDeath          ,"DT" },
            { eMobMode.PlayerNeutral        ,"NU" },
            { eMobMode.PlayerWalk           ,"WL" },
            { eMobMode.PlayerRun            ,"RN" },
            { eMobMode.PlayerGetHit         ,"GH" },
            { eMobMode.PlayerTownNeutral    ,"TN" },
            { eMobMode.PlayerTownWalk       ,"TW" },
            { eMobMode.PlayerAttack1        ,"A1" },
            { eMobMode.PlayerAttack2        ,"A2" },
            { eMobMode.PlayerBlock          ,"BL" },
            { eMobMode.PlayerCast           ,"SC" },
            { eMobMode.PlayerThrow          ,"TH" },
            { eMobMode.PlayerKick           ,"KK" },
            { eMobMode.PlayerSkill1         ,"S1" },
            { eMobMode.PlayerSkill2         ,"S2" },
            { eMobMode.PlayerSkill3         ,"S3" },
            { eMobMode.PlayerSkill4         ,"S4" },
            { eMobMode.PlayerDead           ,"DD" },
            { eMobMode.PlayerSequence       ,"GH" },
            { eMobMode.PlayerKnockBack      ,"GH" },
              
            { eMobMode.MonsterDeath         , "DT" },
            { eMobMode.MonsterNeutral       , "NU" },
            { eMobMode.MonsterWalk          , "WL" },
            { eMobMode.MonsterGetHit        , "GH" },
            { eMobMode.MonsterAttack1       , "A1" },
            { eMobMode.MonsterAttack2       , "A2" },
            { eMobMode.MonsterBlock         , "BL" },
            { eMobMode.MonsterCast          , "SC" },
            { eMobMode.MonsterSkill1        , "S1" },
            { eMobMode.MonsterSkill2        , "S2" },
            { eMobMode.MonsterSkill3        , "S3" },
            { eMobMode.MonsterSkill4        , "S4" },
            { eMobMode.MonsterDead          , "DD" },
            { eMobMode.MonsterKnockback     , "GH" },
            { eMobMode.MonsterSequence      , "xx" },
            { eMobMode.MonsterRun           , "RN" },

            { eMobMode.ObjectNeutral        , "NU" },
            { eMobMode.ObjectOperating      , "OP" },
            { eMobMode.ObjectOpened         , "ON" },
            { eMobMode.ObjectSpecial1       , "S1" },
            { eMobMode.ObjectSpecial2       , "S2" },
            { eMobMode.ObjectSpecial3       , "S3" },
            { eMobMode.ObjectSpecial4       , "S4" },
            { eMobMode.ObjectSpecial5       , "S5" }


        };

        public static string ToToken(this eMobMode src) => mobModes[src];
        public static eMobMode FromToken(this string token) => mobModes.First(x => x.Value.ToUpper() == token.ToUpper()).Key;
    }

}
