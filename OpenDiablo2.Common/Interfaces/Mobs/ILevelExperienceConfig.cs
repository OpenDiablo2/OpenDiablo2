using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface ILevelExperienceConfig
    {
        long GetTotalExperienceForLevel(int level);
        int GetMaxLevel();
    }
}
