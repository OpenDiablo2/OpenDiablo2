using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class StatModifierAddition : IStatModifier
    {
        public double Value = 0;
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
        public double Value = 0;
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
                    return (current * Value);
                case eStatModifierType.MAX:
                    return (max * Value);
                case eStatModifierType.MIN:
                    return (min * Value);
            }
            return 0; // shouldn't reach this
        }
    }
}
