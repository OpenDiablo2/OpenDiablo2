using System.Collections.Generic;
using System.Linq;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.Extensions
{
    public static class MobManagerExtensions
    {
        public static IEnumerable<MobState> FindInRadius(this IEnumerable<MobState> mobs, float centerx, float centery, float radius)
            => mobs.Where(x => x.GetDistance(centerx, centery) <= radius);
    }
}
