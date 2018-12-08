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

using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class StatModifierAddition : IStatModifier
    {
        public double Value { get; set; } = 0;
        public int Priority { get; private set; }
        public string Name { get; private set; }
        public eStatModifierType ModifierType { get; private set; }
        public StatModifierAddition(string name, double value, int priority = 0, eStatModifierType modifiertype = eStatModifierType.CURRENT)
        {
            Value = value;
            Priority = priority;
            Name = name;
            ModifierType = modifiertype;
        }
        public int GetValue(int min, int max, int current)
        {
            return (int)Value; 
        }
        public double GetValue(double min, double max, double current)
        {
            return Value;
        }
    }

    public class StatModifierMultiplication : IStatModifier
    {
        public double Value { get; set; } = 0;
        public int Priority { get; private set; }
        public string Name { get; private set; }
        public eStatModifierType ModifierType { get; private set; }
        public StatModifierMultiplication(string name, double value, int priority = 0, eStatModifierType modifiertype = eStatModifierType.CURRENT)
        {
            Value = value;
            Priority = priority;
            Name = name;
            ModifierType = modifiertype;
        }
        public int GetValue(int min, int max, int current)
        {
            switch (ModifierType)
            {
                case eStatModifierType.CURRENT:
                    return (int)(current * Value);
                case eStatModifierType.MAX:
                    return (int)(max * Value);
                case eStatModifierType.MIN:
                    return (int)(min * Value);
            }
            return 0; // shouldn't reach this
        }
        public double GetValue(double min, double max, double current)
        {
            switch (ModifierType)
            {
                case eStatModifierType.CURRENT:
                    return current * Value;
                case eStatModifierType.MAX:
                    return max * Value;
                case eStatModifierType.MIN:
                    return min * Value;
            }
            return 0; // shouldn't reach this
        }
    }
}
