using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eArmorType
    {
        Lite,
        Medium,
        Heavy
    }

    public static class eArmorTypeHelper
    {
        private static readonly ImmutableDictionary<eArmorType, string> tokens = new Dictionary<eArmorType, string>()
        {
            { eArmorType.Lite   , "lit" },
            { eArmorType.Medium , "med" },
            { eArmorType.Heavy  , "hvy" }
        }.ToImmutableDictionary();
        public static string ToToken(this eArmorType armorType) => tokens[armorType];
        public static eArmorType ToArmorType(this string source) => tokens.First(x => x.Value.ToUpper() == source.ToUpper()).Key;
    }
}
