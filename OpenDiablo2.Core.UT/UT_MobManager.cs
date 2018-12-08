using System.Linq;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Extensions;
using OpenDiablo2.Common.Models.Mobs;
using OpenDiablo2.Core.GameState_;

namespace OpenDiablo2.Core.UT
{
    [TestClass]
    public class UT_MobManager
    {
        [TestMethod]
        public void FindMobsTest()
        {
            MobManager mobman = new MobManager();
            MobState mob1 = new MobState("test1", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            mob1.AddFlag(eMobFlags.ENEMY);
            mobman.AddMob(mob1);

            MobState mob2 = new MobState("test2", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            mob2.AddFlag(eMobFlags.ENEMY);
            mob2.AddFlag(eMobFlags.INVULNERABLE);
            mobman.AddMob(mob2);

            MobState mob3 = new MobState("test3", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            mobman.AddMob(mob3);

            Assert.IsTrue(mobman.Mobs.Count(x => x.HasFlag(eMobFlags.ENEMY)) == 2);
            Assert.IsTrue(mobman.Mobs.Count(x => x.HasFlag(eMobFlags.INVULNERABLE)) == 1);
            Assert.IsTrue(!mobman.Mobs.Any(x => x.HasFlag(eMobFlags.PLAYER)));
            Assert.IsTrue(mobman.Mobs.Count(x => !x.HasFlag(eMobFlags.PLAYER)) == 3);
        }

        [TestMethod]
        public void FindMobsInRadiusTest()
        {
            MobManager mobman = new MobManager();
            MobState mob1 = new MobState("test1", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            mob1.AddFlag(eMobFlags.ENEMY);
            mobman.AddMob(mob1);

            MobState mob2 = new MobState("test2", mobman.GetNextAvailableMobId(), 1, 100, 10, 10);
            mob2.AddFlag(eMobFlags.ENEMY);
            mob2.AddFlag(eMobFlags.INVULNERABLE);
            mobman.AddMob(mob2);

            MobState mob3 = new MobState("test3", mobman.GetNextAvailableMobId(), 1, 100, 3, 1);
            mobman.AddMob(mob3);

            Assert.IsTrue(mobman.Mobs.FindInRadius(0, 0, 1).Count() == 1);
            Assert.IsTrue(mobman.Mobs.FindInRadius(0, 0, 7).Count() == 2);
            Assert.IsTrue(mobman.Mobs.FindInRadius(0, 0, 20).Count() == 3);
            Assert.IsTrue(mobman.Mobs.FindInRadius(10, 10, 1).Count() == 1);
        }
    }
}
