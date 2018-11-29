using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.UT
{
    [TestClass]
    public class UT_MobState
    {
        [TestMethod]
        public void MobDeathTest()
        {
            MobState mob = new MobState("test1", 1, 100, 100, 100, 0, 0);
            mob.TakeDamage(1000, Enums.Mobs.eDamageTypes.NONE);
            Assert.AreEqual(false, mob.Alive);
        }

        [TestMethod]
        public void MobFlagsTest()
        {
            MobState mob = new MobState("test1", 1, 100, 100, 100, 0, 0);
            mob.AddFlag(Enums.Mobs.eMobFlags.ENEMY);
            Assert.AreEqual(true, mob.HasFlag(Enums.Mobs.eMobFlags.ENEMY));
            Assert.AreEqual(false, mob.HasFlag(Enums.Mobs.eMobFlags.INVULNERABLE));
            mob.RemoveFlag(Enums.Mobs.eMobFlags.ENEMY);
            Assert.AreEqual(false, mob.HasFlag(Enums.Mobs.eMobFlags.ENEMY));
        }
    }
}
