using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eCompositType
    {
        Head,
        Torso,
        Legs,
        RightArm,
        LeftArm,
        RightHand,
        LeftHand,
        Shield,
        Special1,
        Special2,
        Special3,
        Special4,
        Special5,
        Special6,
        Special7,
        Special8
    }

    public static class eCompositeTypeHelper
    {
        private readonly static Dictionary<eCompositType, string> tokens = new Dictionary<eCompositType, string>
        {
            { eCompositType.Head        , "HD" },
            { eCompositType.Torso       , "TR" },
            { eCompositType.Legs        , "LG" },
            { eCompositType.RightArm    , "RA" },
            { eCompositType.LeftArm     , "LA" },
            { eCompositType.RightHand   , "RH" },
            { eCompositType.LeftHand    , "LH" },
            { eCompositType.Shield      , "SH" },
            { eCompositType.Special1    , "S1" },
            { eCompositType.Special2    , "S2" },
            { eCompositType.Special3    , "S3" },
            { eCompositType.Special4    , "S4" },
            { eCompositType.Special5    , "S5" },
            { eCompositType.Special6    , "S6" },
            { eCompositType.Special7    , "S7" },
            { eCompositType.Special8    , "S8" }
        };

        public static string ToToken(this eCompositType source) => tokens[source];
        public static eCompositType ToCompositeType(this string source) => tokens.First(x => x.Value.ToUpper() == source.ToUpper()).Key;

    }
}
