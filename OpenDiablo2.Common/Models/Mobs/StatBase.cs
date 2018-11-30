using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public abstract class StatBase
    {
        protected List<int> OrderedModifierKeys = new List<int>();
        protected Dictionary<int, List<IStatModifier>> Modifiers = new Dictionary<int, List<IStatModifier>>();

        public void AddModifier(IStatModifier mod)
        {
            int priority = mod.Priority;
            if (!OrderedModifierKeys.Contains(priority))
            {
                // if this priority doesn't exist yet, create it
                // we must add the new key in order of highest to lowest
                int i = 0;
                while (i < OrderedModifierKeys.Count)
                {
                    if (OrderedModifierKeys[i] < priority)
                    {
                        OrderedModifierKeys.Insert(i, priority);
                        break;
                    }
                    i++;
                }
                if (i >= OrderedModifierKeys.Count)
                {
                    OrderedModifierKeys.Add(priority);
                }
                Modifiers.Add(priority, new List<IStatModifier>() { mod });
            }
            else
            {
                // if priority already exists, just add it to the existing list
                Modifiers[priority].Add(mod);
            }
        }

        public void RemoveModifier(string name)
        {
            foreach (int k in Modifiers.Keys)
            {
                int i = 0;
                while (i < Modifiers[k].Count)
                {
                    IStatModifier mod = Modifiers[k][i];
                    if (mod.Name == name)
                    {
                        Modifiers[k].RemoveAt(i);
                        continue; // don't increment i
                    }
                    i++;
                }
            }

            // now do cleanup: if any modifier list is empty, remove it and its priority
            int o = 0;
            while (o < Modifiers.Keys.Count)
            {
                int k = Modifiers.Keys.ElementAt(o);
                if (Modifiers[k].Count <= 0)
                {
                    Modifiers.Remove(k);
                    OrderedModifierKeys.Remove(k);
                    continue; // don't increment o
                }
                o++;
            }
        }
    }
}
