using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class HeroTypeConfig : IHeroTypeConfig
    {
        public int StartingVitality { get; protected set; }
        public int StartingStrength { get; protected set; }
        public int StartingDexterity { get; protected set; }
        public int StartingEnergy { get; protected set; }
        public int StartingHealth { get; protected set; }
        public int StartingMana { get; protected set; }
        public int StartingStamina { get; protected set; }

        public double PerLevelHealth { get; protected set; }
        public double PerLevelMana { get; protected set; }
        public double PerLevelStamina { get; protected set; }

        public double PerVitalityHealth { get; protected set; }
        public double PerVitalityStamina { get; protected set; }
        public double PerEnergyMana { get; protected set; }

        public int BaseAttackRating { get; protected set; }
        public int PerDexterityAttackRating { get; protected set; }

        public int BaseDefenseRating { get; protected set; }
        public int PerDexterityDefenseRating { get; protected set; }

        public HeroTypeConfig(int vitality, int strength, int dexterity, int energy,
            int health, int mana, int stamina,
            double perlevelhealth, double perlevelmana, double perlevelstamina,
            double pervitalityhealth, double pervitalitystamina, double perenergymana,
            int baseatkrating, int basedefrating, int perdexterityatkrating, int perdexteritydefrating)
        {
            StartingDexterity = dexterity;
            StartingVitality = vitality;
            StartingStrength = strength;
            StartingEnergy = energy;
            StartingMana = mana;
            StartingHealth = health;
            StartingStamina = stamina;

            PerLevelHealth = perlevelhealth;
            PerLevelMana = perlevelmana;
            PerLevelStamina = perlevelstamina;

            PerVitalityHealth = pervitalityhealth;
            PerVitalityStamina = pervitalitystamina;
            PerEnergyMana = perenergymana;

            BaseAttackRating = baseatkrating;
            BaseDefenseRating = basedefrating;
            PerDexterityAttackRating = perdexterityatkrating;
            PerDexterityDefenseRating = perdexteritydefrating;
        }
    }
}
