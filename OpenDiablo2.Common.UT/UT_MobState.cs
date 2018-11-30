using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Models.Mobs;

namespace OpenDiablo2.Common.UT
{
    [TestClass]
    public class UT_MobState
    {
        [TestMethod]
        public void MobDeathTest()
        {
            MobState mob = new MobState("test1", 1, 1, 100,  0, 0);
            mob.TakeDamage(1000, Enums.Mobs.eDamageTypes.NONE);
            Assert.AreEqual(false, mob.Alive);
        }

        [TestMethod]
        public void MobFlagsTest()
        {
            MobState mob = new MobState("test1", 1, 1, 100,  0, 0);
            mob.AddFlag(Enums.Mobs.eMobFlags.ENEMY);
            Assert.AreEqual(true, mob.HasFlag(Enums.Mobs.eMobFlags.ENEMY));
            Assert.AreEqual(false, mob.HasFlag(Enums.Mobs.eMobFlags.INVULNERABLE));
            mob.RemoveFlag(Enums.Mobs.eMobFlags.ENEMY);
            Assert.AreEqual(false, mob.HasFlag(Enums.Mobs.eMobFlags.ENEMY));
        }

        [TestMethod]
        public void MobResistancesTest()
        {
            MobState mob = new MobState("test1", 1, 1, 100, 0, 0);
            mob.SetResistance(eDamageTypes.COLD, 0.5);
            int dam = mob.TakeDamage(100, eDamageTypes.COLD); // now try 100 cold damage
            Assert.AreEqual(50, dam); // b/c mob has 50% resistance, should only take 50 damage 
            Assert.IsTrue(mob.Alive); // mob should not have taken enough damage to die
            int dam2 = mob.TakeDamage(100, eDamageTypes.FIRE); // now try 100 fire damage
            Assert.AreEqual(100, dam2); // b/c mob has no fire resistance, should take full damage 
            Assert.IsFalse(mob.Alive); // mob should be dead
        }

        [TestMethod]
        public void MobImmunityTest()
        {
            MobState mob = new MobState("test1", 1, 1, 100, 0, 0);
            mob.AddImmunitiy(eDamageTypes.COLD);
            int dam = mob.TakeDamage(100, eDamageTypes.COLD); // now try 100 cold damage
            Assert.AreEqual(0, dam); // b/c mob has immunity, should not take damage 
            Assert.IsTrue(mob.Alive); // mob should not have taken enough damage to die
            int dam2 = mob.TakeDamage(100, eDamageTypes.FIRE); // now try 100 fire damage
            Assert.AreEqual(100, dam2); // b/c mob has no fire immunity, should take full damage 
            Assert.IsFalse(mob.Alive); // mob should be dead
        }
    }
}
