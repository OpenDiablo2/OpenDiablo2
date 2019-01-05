/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.Linq;

namespace OpenDiablo2.Common.Enums
{
    public enum eHero
    {
        None,
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
        public readonly static eHero[] all = {
            eHero.Barbarian,
            eHero.Necromancer,
            eHero.Paladin,
            eHero.Assassin,
            eHero.Sorceress,
            eHero.Amazon,
            eHero.Druid,
        };

        public readonly static Dictionary<eHero, string> tokens = new Dictionary<eHero, string>
        {
            { eHero.Barbarian   , "BA" },
            { eHero.Necromancer , "NE" },
            { eHero.Paladin     , "PA" },
            { eHero.Assassin    , "AI" },
            { eHero.Sorceress   , "SO" },
            { eHero.Amazon      , "AM" },
            { eHero.Druid       , "DZ" },
        };

        public static string ToToken(this eHero source) => tokens[source];
        public static eHero ToHero(this string source) => tokens.First(x => x.Value.ToUpper() == source.ToUpper()).Key;
    }
}
