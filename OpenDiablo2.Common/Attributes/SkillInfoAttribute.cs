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

using OpenDiablo2.Common.Enums;
using System;
using System.Collections.Generic;

namespace OpenDiablo2.Common.Attributes
{
    [AttributeUsage(AttributeTargets.Field, AllowMultiple = false)]
    public class SkillInfoAttribute : Attribute
    {
        int a = 5;
        public SkillInfoAttribute(eHero hero, int spriteIndex = 0, int level/*levelGroup*/ = 0, params eSkill[] skillsRequired)
        {
            Hero = hero;
            SpriteIndex = spriteIndex;
            Level = level;
            SkillsRequired = skillsRequired ?? Array.Empty<eSkill>();
        }

        public eHero Hero { get; }
        public int SpriteIndex { get; }
        public int Level { get; }
        public IReadOnlyList<eSkill> SkillsRequired { get; }
    }
}
