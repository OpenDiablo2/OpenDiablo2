using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class PlayerState : MobState
    {
        public Guid Id { get; protected set; }
        public eHero HeroType { get; protected set; }
        private IHeroTypeConfig HeroTypeConfig;
        private ILevelExperienceConfig ExperienceConfig;

        // Player character stats
        protected Stat Vitality;
        protected Stat Strength;
        protected Stat Energy;
        protected Stat Dexterity;
        
        protected Stat DefenseRating;
        protected Stat AttackRating;

        protected Stat Stamina;
        protected Stat Mana;

        public int Experience { get; protected set; }

        public PlayerState(string name, int id, int level, float x, float y,
            int vitality, int strength, int energy, int dexterity, int experience, eHero herotype, 
            IHeroTypeConfig heroconfig, ILevelExperienceConfig expconfig)
            : base(name, id, level, 0, x, y)
        {
            Stamina = new Stat(0, 0, 0, true);
            Mana = new Stat(0, 0, 0, true);

            Vitality = new Stat(0, vitality, vitality, true);
            Strength = new Stat(0, strength, strength, true);
            Energy = new Stat(0, energy, energy, true);
            Dexterity = new Stat(0, dexterity, dexterity, true);

            AttackRating = new Stat(0, 0, 0, false); 
            DefenseRating = new Stat(0, 0, 0, false);

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

        #region Level and Experience
        public int GetExperienceToLevel()
        {
            return GetExperienceTotalToLevel() - Experience;
        }
        public int GetExperienceTotalToLevel()
        {
            return ExperienceConfig.GetTotalExperienceForLevel(Level + 1);
        }
        public int GetMaxLevel()
        {
            return ExperienceConfig.GetMaxLevel();
        }
        public bool AddExperience(int amount)
        {
            // returns true if you level up from this
            Experience += amount;
            return TryLevelUp();
        }
        public bool TryLevelUp()
        {
            if(Level >= GetMaxLevel())
            {
                return false; // can't level up anymore
            }
            if(GetExperienceToLevel() > 0)
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
            return (int)(((Level % 2) * Math.Ceiling(increase)) + ((1 - (Level % 2)) * Math.Floor(increase)));
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

        // TODO: when a player equips an item, apply the relevant modifiers to their stats
        // TODO: when a player unequips an item, remove the relevant modifiers from their stats
    }
}
