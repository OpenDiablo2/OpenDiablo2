using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class MobConditionAnd : IMobCondition
    {
        public List<IMobCondition> Conditions = new List<IMobCondition>();

        public MobConditionAnd(List<IMobCondition> conditions)
        {
            Conditions = conditions;
        }

        public bool Evaluate(MobState mob)
        {
            foreach(IMobCondition condition in Conditions)
            {
                if (!condition.Evaluate(mob))
                {
                    return false;
                }
            }
            return true;
        }
    }

    public class MobConditionOr : IMobCondition
    {
        public List<IMobCondition> Conditions = new List<IMobCondition>();

        public MobConditionOr(List<IMobCondition> conditions)
        {
            Conditions = conditions;
        }

        public bool Evaluate(MobState mob)
        {
            foreach (IMobCondition condition in Conditions)
            {
                if (condition.Evaluate(mob))
                {
                    return true;
                }
            }
            return false;
        }
    }

    public class MobConditionNot : IMobCondition
    {
        public IMobCondition Condition = null;

        public MobConditionNot(IMobCondition condition)
        {
            Condition = condition;
        }

        public bool Evaluate(MobState mob)
        {
            if(Condition == null)
            {
                return false;
            }
            return !Condition.Evaluate(mob);
        }
    }

    public class MobConditionFlags : IMobCondition
    {
        public Dictionary<eMobFlags, bool> Flags = new Dictionary<eMobFlags, bool>();

        public MobConditionFlags(Dictionary<eMobFlags, bool> flags)
        {
            Flags = flags;
        }

        public bool Evaluate(MobState mob)
        {
            foreach(eMobFlags flag in Flags.Keys)
            {
                if(Flags[flag] != mob.HasFlag(flag))
                {
                    return false;
                }
            }
            return true;
        }
    }

    // TODO: implement more of these
}
