using OpenDiablo2.Common.Enums.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IStatModifier
    {
        string Name { get; }
        int Priority { get; } // what priority does this modifier happen in? higher = occurs before other modifiers
        // modifiers at the same priority level occur simultaneously
        eStatModifierType ModifierType { get; } // does it affect current, min or max?
        int GetValue(int min, int max, int current); // what does this modifier add to the stat's current value?
        double GetValue(double min, double max, double current);
    }
}
