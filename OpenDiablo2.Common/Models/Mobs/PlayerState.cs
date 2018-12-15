/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Mobs;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class PlayerState : MobState, IMobLocation
    {
        private readonly IHeroTypeConfig HeroTypeConfig;
        private readonly ILevelExperienceConfig ExperienceConfig;

        public Guid UID { get; protected set; } = Guid.NewGuid();
        public eHero HeroType { get; protected set; }
        public int ClientHash { get; protected set; }
        public List<PointF> Waypoints { get; set; } = new List<PointF>();
        public eMovementType MovementType { get; set; } = eMovementType.Stopped;
        public eMobMode MobMode { get; set; } = eMobMode.PlayerTownWalk;
        public PlayerEquipment Equipment { get; set; } = new PlayerEquipment();


        // Player character stats
        protected Stat Vitality;
        protected Stat Strength;
        protected Stat Energy;
        protected Stat Dexterity;

        protected Stat DefenseRating;
        protected Stat AttackRating;

        protected Stat Stamina;
        protected Stat Mana;
        protected Stat ManaRegen;

        protected Stat WalkVelocity;
        protected Stat RunVelocity;
        protected Stat RunDrain;
        
        public long Experience { get; protected set; }

        public PlayerState() : base() { }

        public PlayerState(int clientHash, string name, int id, int level, float x, float y,
            int vitality, int strength, int energy, int dexterity, long experience, eHero herotype,
            IHeroTypeConfig heroconfig, ILevelExperienceConfig expconfig)
            : base(name, id, level, 0, x, y)
        {
            this.ClientHash = clientHash;
            Stamina = new Stat(0, 0, 0, true);
            Mana = new Stat(0, 0, 0, true);
            ManaRegen = new Stat(0, heroconfig.StartingManaRegen, heroconfig.StartingManaRegen, true);

            Vitality = new Stat(0, vitality, vitality, true);
            Strength = new Stat(0, strength, strength, true);
            Energy = new Stat(0, energy, energy, true);
            Dexterity = new Stat(0, dexterity, dexterity, true);

            AttackRating = new Stat(0, 0, 0, false);
            DefenseRating = new Stat(0, 0, 0, false);

            WalkVelocity = new Stat(0, heroconfig.WalkVelocity, 100, false); // TODO: what should max velocity be? (replace the 100)
            RunVelocity = new Stat(0, heroconfig.RunVelocity, 100, false); // TODO: what should max velocity be?
            RunDrain = new Stat(0, heroconfig.RunDrain, heroconfig.RunDrain, true);

            Experience = experience; // how much total exp do they have


            HeroType = herotype;
            HeroTypeConfig = heroconfig;
            ExperienceConfig = expconfig;

            AddFlag(eMobFlags.PLAYER);

            RefreshMaxes(); // initialize the max health / mana / energy
            Health.SetCurrent(Health.GetMax());
            Mana.SetCurrent(Mana.GetMax());
            Energy.SetCurrent(Energy.GetMax());
            RefreshDerived();
        }

        public void UpdateEquipment(string slot, ItemInstance item) => Equipment.EquipItem(slot, item);
        
        #region Level and Experience
        public long GetExperienceToLevel()
        {
            return GetExperienceTotalToLevel() - Experience;
        }
        public long GetExperienceTotalToLevel()
        {
            return ExperienceConfig.GetTotalExperienceForLevel(Level + 1);
        }
        public int GetMaxLevel()
        {
            return ExperienceConfig.GetMaxLevel();
        }
        public bool AddExperience(long amount)
        {
            // returns true if you level up from this
            Experience += amount;
            return TryLevelUp();
        }
        public bool TryLevelUp()
        {
            if (Level >= GetMaxLevel())
            {
                return false; // can't level up anymore
            }
            if (GetExperienceToLevel() > 0)
            {
                return false; // not enough exp
            }
            LevelUp();
            return true;
        }
        public void LevelUp()
        {
            Level += 1;

            // instead of adding to max here, let's just refresh all the maxes upon level up
            //Health.AddMax(CalculatePerLevelStatIncrease(HeroTypeConfig.PerLevelHealth));
            //Mana.AddMax(CalculatePerLevelStatIncrease(HeroTypeConfig.PerLevelMana));
            //Stamina.AddMax(CalculatePerLevelStatIncrease(HeroTypeConfig.PerLevelStamina));
            RefreshMaxes();
        }
        private int CalculatePerLevelStatIncrease(double increase)
        {
            // on even levels, e.g. 2, 4, 6, you gain 1 from an increase of 1.5
            // on odd levels, you gain 2 from an increase of 1.5
            return (int)((Level % 2 * Math.Ceiling(increase)) + ((1 - (Level % 2)) * Math.Floor(increase)));
        }
        #endregion Level and Experience

        #region Stat Calculation
        public void RefreshDerived()
        {
            RefreshDexterityDerived();
        }
        // determine attack rating / defense rating
        public void RefreshDexterityDerived()
        {
            // this needs to be called whenever dexterity changes
            RefreshAttackRating();
            RefreshDefenseRating();
        }
        public void RefreshAttackRating()
        {
            int atk = HeroTypeConfig.BaseAttackRating + Dexterity.GetCurrent() * HeroTypeConfig.PerDexterityAttackRating;
            AttackRating.SetMax(atk);
            AttackRating.SetCurrent(atk);
        }
        public void RefreshDefenseRating()
        {
            int def = HeroTypeConfig.BaseDefenseRating + Dexterity.GetCurrent() * HeroTypeConfig.PerDexterityDefenseRating;
            DefenseRating.SetMax(def);
            DefenseRating.SetCurrent(def);
        }
        // determine the max health/ mana/ stamina
        public void RefreshMaxes()
        {
            RefreshMaxHealth();
            RefreshMaxStamina();
            RefreshMaxMana();
        }
        public void RefreshMaxHealth()
        {
            int health = HeroTypeConfig.StartingHealth;
            health += (int)Math.Floor((Level - 1) * HeroTypeConfig.PerLevelHealth);
            health += (int)Math.Floor((Vitality.GetCurrent() - HeroTypeConfig.StartingVitality) * HeroTypeConfig.PerVitalityHealth);
            Health.SetMax(health);
        }
        public void RefreshMaxStamina()
        {
            int stamina = HeroTypeConfig.StartingStamina;
            stamina += (int)Math.Floor((Level - 1) * HeroTypeConfig.PerLevelStamina);
            stamina += (int)Math.Floor((Vitality.GetCurrent() - HeroTypeConfig.StartingVitality) * HeroTypeConfig.PerVitalityStamina);
            Stamina.SetMax(stamina);
        }
        public void RefreshMaxMana()
        {
            int mana = HeroTypeConfig.StartingMana;
            mana += (int)Math.Floor((Level - 1) * HeroTypeConfig.PerLevelMana);
            mana += (int)Math.Floor((Energy.GetCurrent() - HeroTypeConfig.StartingEnergy) * HeroTypeConfig.PerEnergyMana);
            Mana.SetMax(mana);
        }
        #endregion Stat Calculation

        #region Mana
        public int GetMana()
        {
            return Mana.GetCurrent();
        }
        public int GetManaMax()
        {
            return Mana.GetMax();
        }
        public void RecoverMana(int mana)
        {
            Mana.AddCurrent(mana);
        }
        public void UseMana(int mana)
        {
            Mana.AddCurrent(-mana);
        }
        public int GetManaRegen()
        {
            return ManaRegen.GetCurrent();
        }
        #endregion Mana

        #region Stamina
        public int GetStamina()
        {
            return Stamina.GetCurrent();
        }
        public int GetStaminaMax()
        {
            return Stamina.GetMax();
        }
        public void RecoverStamina(int stamina)
        {
            Stamina.AddCurrent(stamina);
        }
        public void UseStamina(int stamina)
        {
            Stamina.AddCurrent(-stamina);
        }
        #endregion Stamina

        #region Movement
        public int GetRunVelocity()
        {
            return RunVelocity.GetCurrent();
        }
        public int GetWalkVeloicty()
        {
            return WalkVelocity.GetCurrent();
        }
        public int GetRunDrain()
        {
            return RunDrain.GetCurrent();
        }
        #endregion Movement

        // TODO: when a player equips an item, apply the relevant modifiers to their stats
        // TODO: when a player unequips an item, remove the relevant modifiers from their stats
    }
}
