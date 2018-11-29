using Microsoft.VisualStudio.TestTools.UnitTesting;
using OpenDiablo2.Common.Models.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.UT
{
    [TestClass]
    public class UT_Stat
    {
        [TestMethod]
        public void BasicStatTest()
        {
            Stat hp = new Stat(0, 100, 100, false);
            Assert.AreEqual(100, hp.GetCurrent());
            Assert.AreEqual(100, hp.GetMax());
            Assert.AreEqual(0, hp.GetMin());

            hp.AddCurrent(10);
            Assert.AreEqual(100, hp.GetCurrent()); // not allowed to go over max

            hp.AddCurrent(-30);
            Assert.AreEqual(70, hp.GetCurrent());

            hp.SetCurrent(10);
            Assert.AreEqual(10, hp.GetCurrent());

            hp.AddCurrent(-30);
            Assert.AreEqual(0, hp.GetCurrent()); // not allowed to go under min
        }

        [TestMethod]
        public void ModifierStatTest()
        {
            Stat hp = new Stat(0, 100, 100, false);

            StatModifierAddition mod1 = new StatModifierAddition("mod1", 10);
            hp.AddModifier(mod1);

            Assert.AreEqual(100, hp.GetCurrent()); // not allowed to overflow from mods
            hp.AllowedToOverflowFromModifiers = true;
            Assert.AreEqual(110, hp.GetCurrent()); // if we flip overflow to true, should overflow from mods

            StatModifierAddition mod2 = new StatModifierAddition("mod2", -20);
            hp.AddModifier(mod2);
            Assert.AreEqual(90, hp.GetCurrent());

            hp.RemoveModifier("mod1");
            Assert.AreEqual(80, hp.GetCurrent()); // if we remove mod1, we should see its effect go away

            hp.RemoveModifier("mod2");
            Assert.AreEqual(100, hp.GetCurrent()); // similarly with mod2
        }

        [TestMethod]
        public void ModifierPriorityTest()
        {
            Stat hp = new Stat(0, 100, 50, true);

            StatModifierMultiplication mod1 = new StatModifierMultiplication("mod1", 0.5, priority: 1);
            hp.AddModifier(mod1);
            Assert.AreEqual(75, hp.GetCurrent());

            StatModifierMultiplication mod2 = new StatModifierMultiplication("mod2", 0.5, priority: 1);
            hp.AddModifier(mod2);
            Assert.AreEqual(100, hp.GetCurrent()); // because these have the same priority,
            // they don't see eachother's effects (e.g. we get 50 + 50*0.5 + 50*0.5, NOT 50 + 50*0.5 + 75*0.5

            hp.RemoveModifier("mod2");
            Assert.AreEqual(75, hp.GetCurrent());
            StatModifierMultiplication mod3 = new StatModifierMultiplication("mod3", -0.2, priority: 0);
            hp.AddModifier(mod3);
            Assert.AreEqual(60, hp.GetCurrent()); // because this one has a lower priority, it happens AFTER mod1
            // so we get 50 + 50*0.5 = 75, then 75 + 75*-0.2 = 60
        }
    }
}
