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
    public class UT_PlayerState
    {
        private PlayerState MakePlayer()
        {
            HeroTypeConfig herotypeconfig = new HeroTypeConfig(vitality: 20, strength: 20, dexterity: 20,
                energy: 20, health: 50, mana: 40, stamina: 30, manaRegen: 20,
                perlevelhealth: 3, perlevelmana: 1.5, perlevelstamina: 1,
                pervitalityhealth: 2, pervitalitystamina: 1, perenergymana: 1.5,
                perLevelStatPoints: 5,
                baseatkrating: -30, basedefrating: -30, perdexterityatkrating: 5, perdexteritydefrating: 4,
                walkVelocity: 6, runVelocity: 9, runDrain: 20, walkFrames: 8, runFrames: 8, swingFrames: 8,
                spellFrames: 8, getHitFrames: 8, bowFrames: 8, startingSkill: "",
                startingSkills: new string[] { },
                allSkillsBonusString: "", firstTabBonusString: "", secondTabBonusString: "",
                thirdTabBonusString: "", classOnlyBonusString: "", baseWeaponClass: "hth", 
                initialEquipment: new List<InitialEquipment>());
            LevelExperienceConfig expconfig = new LevelExperienceConfig(new List<long>()
            {
                0, // level 0
                0, // level 1
                500, // level 2
                1500, // level 3
                2250,
                4125,
            });
            PlayerState ps = new PlayerState(0, "player1", id: 1, level: 1, x: 0, y: 0,
                vitality: herotypeconfig.StartingVitality,
                strength: herotypeconfig.StartingStrength,
                energy: herotypeconfig.StartingEnergy,
                dexterity: herotypeconfig.StartingDexterity,
                experience: 0,
                herotype: Enums.eHero.Amazon,
                heroconfig: herotypeconfig,
                expconfig: expconfig);
            return ps;
        }

        [TestMethod]
        public void PlayerLevelUpTest()
        {
            PlayerState ps = MakePlayer();
            
            int level1hp = ps.GetHealthMax();
            int level1mana = ps.GetManaMax();
            int level1stamina = ps.GetStaminaMax();
            Assert.IsFalse(ps.AddExperience(499)); // not quite enough
            Assert.IsTrue(ps.AddExperience(1)); // there we go
            int level2hp = ps.GetHealthMax();
            int level2mana = ps.GetManaMax();
            int level2stamina = ps.GetStaminaMax();
            Assert.IsFalse(ps.AddExperience(999)); // not quite enough
            Assert.IsTrue(ps.AddExperience(1)); // there we go
            int level3hp = ps.GetHealthMax();
            int level3mana = ps.GetManaMax();
            int level3stamina = ps.GetStaminaMax();

            // this should update the player's health, mana, and stamina
            Assert.AreEqual(level1hp + 3, level2hp);
            Assert.AreEqual(level1mana + 1, level2mana);
            Assert.AreEqual(level1stamina + 1, level2stamina);

            Assert.AreEqual(level2hp + 3, level3hp);
            Assert.AreEqual(level2mana + 2, level3mana); // because 2->3 is an odd levelup, should get 2
            // more mana on this level (1.5 = +1 on even levels, +2 on odd)
            Assert.AreEqual(level2stamina + 1, level3stamina);
        }
    }
}
