using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class StatDouble : StatBase
    {
        protected double Min = 0;
        protected double Max = 0;
        protected double Current = 0; // the current value BEFORE modifiers

        public bool AllowedToOverflowFromModifiers { get; set; } = false; // if true, can return a value greater than Max 
        // if a modifier is increasing the current value

        public StatDouble(double min, double max, double current, bool allowedToOverflowFromModifiers)
        {
            Min = min;
            Max = max;
            Current = current;
            AllowedToOverflowFromModifiers = allowedToOverflowFromModifiers;
        }

        public void AddCurrent(double value)
        {
            Current += value;
            ClampCurrent();
        }

        public void SetCurrent(double value)
        {
            Current = value;
            ClampCurrent();
        }

        private void ClampCurrent()
        {
            double currentmax = GetMax();
            double currentmin = GetMin();
            if (Current > currentmax)
            {
                Current = currentmax;
            }
            if (Current < currentmin)
            {
                Current = currentmin;
            }
        }

        public void AddMax(double value)
        {
            Max += value;
            ClampCurrent();
        }

        public void SetMax(double value)
        {
            Max = value;
            ClampCurrent();
        }

        public void AddMin(double value)
        {
            Min += value;
            ClampCurrent();
        }

        public void SetMin(double value)
        {
            Min = value;
            ClampCurrent();
        }

        public double GetCurrent()
        {
            double val = Current;
            // take the current value and apply each modifier to it
            // the modifiers are done in priority order
            // OrderedModifierKeys holds the list of priorities in order from highest to lowest
            // highest is operated first
            // Items on the same priority level don't see eachother
            // e.g. if Health is 20 and we have two modifiers with priority 5 that add 50% health,
            //      then we would get 20 + 10 + 10, NOT 20 + 10 + 15
            //      However, if one of the modifiers now has a different priority, say 4, then we would get
            //      20 + 10 + 15
            foreach (int k in OrderedModifierKeys)
            {
                double newval = val;
                foreach (IStatModifier mod in Modifiers[k])
                {
                    if (mod.ModifierType == eStatModifierType.CURRENT)
                    {
                        newval += mod.GetValue(Min, Max, val);
                    }
                }
                val = newval;
            }

            if (!AllowedToOverflowFromModifiers)
            {
                // if we aren't allowed to go over max, even with modifiers, we must double check
                // that we're within our bounds
                double currentmax = GetMax();
                double currentmin = GetMin();
                if (val > currentmax)
                {
                    val = currentmax;
                }
                if (val < currentmin)
                {
                    val = currentmin;
                }
            }

            return val;
        }

        public double GetMin()
        {
            double val = Min;
            foreach (int k in OrderedModifierKeys)
            {
                double newval = val;
                foreach (IStatModifier mod in Modifiers[k])
                {
                    if (mod.ModifierType == eStatModifierType.MIN)
                    {
                        newval += mod.GetValue(val, Max, Current);
                    }
                }
                val = newval;
            }
            return val;
        }

        public double GetMax()
        {
            double val = Max;
            foreach (int k in OrderedModifierKeys)
            {
                double newval = val;
                foreach (IStatModifier mod in Modifiers[k])
                {
                    if (mod.ModifierType == eStatModifierType.MAX)
                    {
                        newval += mod.GetValue(Min, val, Current);
                    }
                }
                val = newval;
            }

            return val;
        }
    }
}
