using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Models.Mobs;
using OpenDiablo2.Core.GameState_;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

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
            MobState mob2 = new MobState("test2", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            MobState mob3 = new MobState("test3", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            mob1.AddFlag(eMobFlags.ENEMY);
            mob2.AddFlag(eMobFlags.ENEMY);
            mob2.AddFlag(eMobFlags.INVULNERABLE);
            mobman.AddMob(mob1);
            mobman.AddMob(mob2);
            mobman.AddMob(mob3);

            Assert.IsTrue(mobman.FindMobs(new MobConditionFlags(
                new Dictionary<eMobFlags, bool> {
                    { eMobFlags.ENEMY, true }
                })).Count == 2);

            Assert.IsTrue(mobman.FindMobs(new MobConditionFlags(
                new Dictionary<eMobFlags, bool> {
                    { eMobFlags.INVULNERABLE, true }
                })).Count == 1);

            Assert.IsTrue(mobman.FindMobs(new MobConditionFlags(
                new Dictionary<eMobFlags, bool> {
                    { eMobFlags.PLAYER, true }
                })).Count == 0);

            Assert.IsTrue(mobman.FindMobs(new MobConditionFlags(
                new Dictionary<eMobFlags, bool> {
                    { eMobFlags.PLAYER, false }
                })).Count == 3);
        }

        [TestMethod]
        public void FindMobsInRadiusTest()
        {
            MobManager mobman = new MobManager();
            MobState mob1 = new MobState("test1", mobman.GetNextAvailableMobId(), 1, 100, 0, 0);
            MobState mob2 = new MobState("test2", mobman.GetNextAvailableMobId(), 1, 100, 10, 10);
            MobState mob3 = new MobState("test3", mobman.GetNextAvailableMobId(), 1, 100, 3, 1);
            mob1.AddFlag(eMobFlags.ENEMY);
            mob2.AddFlag(eMobFlags.ENEMY);
            mob2.AddFlag(eMobFlags.INVULNERABLE);
            mobman.AddMob(mob1);
            mobman.AddMob(mob2);
            mobman.AddMob(mob3);

            Assert.IsTrue(mobman.FindMobsInRadius(0, 0, 1, null).Count == 1);
            Assert.IsTrue(mobman.FindMobsInRadius(0, 0, 7, null).Count == 2);
            Assert.IsTrue(mobman.FindMobsInRadius(0, 0, 20, null).Count == 3);
            Assert.IsTrue(mobman.FindMobsInRadius(10, 10, 1, null).Count == 1);
        }
    }
}
