using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eHero
    {
        Barbarian,
        Necromancer,
        Paladin,
        Assassin,
        Sorceress,
        Amazon,
        Druid
    }

    public static class eHeroExtensions
    {
        public readonly static Dictionary<eHero, string> tokens = new Dictionary<eHero, string>
        {
            { eHero.Amazon      , "AM" },
            { eHero.Sorceress   , "SO" },
            { eHero.Necromancer , "NE" },
            { eHero.Paladin     , "PA" },
            { eHero.Barbarian   , "BA" },
            { eHero.Druid       , "DZ" },
            { eHero.Assassin    , "AI" }
        };

        public static string ToToken(this eHero source) => tokens[source];
        public static eHero ToHero(this string source) => tokens.First(x => x.Value.ToUpper() == source.ToUpper()).Key;
    }
}
