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
        public int StartingManaRegen { get; protected set; }

        public double PerLevelHealth { get; protected set; }
        public double PerLevelMana { get; protected set; }
        public double PerLevelStamina { get; protected set; }

        public double PerVitalityHealth { get; protected set; }
        public double PerVitalityStamina { get; protected set; }
        public double PerEnergyMana { get; protected set; }

        public int PerLevelStatPoints { get; protected set; }

        public int BaseAttackRating { get; protected set; }
        public int PerDexterityAttackRating { get; protected set; }

        public int BaseDefenseRating { get; protected set; }
        public int PerDexterityDefenseRating { get; protected set; }

        public int WalkVelocity { get; protected set; }
        public int RunVelocity { get; protected set; }
        public int RunDrain { get; protected set; }

        public int WalkFrames { get; protected set; }
        public int RunFrames { get; protected set; }
        public int SwingFrames { get; protected set; }
        public int SpellFrames { get; protected set; }
        public int GetHitFrames { get; protected set; }
        public int BowFrames { get; protected set; }

        public string StartingSkill { get; protected set; }
        public string[] StartingSkills { get; protected set; }

        public string AllSkillsBonusString { get; protected set; }
        public string FirstTabBonusString { get; protected set; }
        public string SecondTabBonusString { get; protected set; }
        public string ThirdTabBonusString { get; protected set; }
        public string ClassOnlyBonusString { get; protected set; }

        public string BaseWeaponClass { get; protected set; }

        public List<InitialEquipment> InitialEquipment { get; }

        public HeroTypeConfig(int vitality, int strength, int dexterity, int energy,
            int health, int mana, int stamina, int manaRegen,
            double perlevelhealth, double perlevelmana, double perlevelstamina,
            double pervitalityhealth, double pervitalitystamina, double perenergymana,
            int perLevelStatPoints,
            int baseatkrating, int basedefrating, int perdexterityatkrating, int perdexteritydefrating,
            int walkVelocity, int runVelocity, int runDrain,
            int walkFrames, int runFrames, int swingFrames, int spellFrames, int getHitFrames, int bowFrames,
            string startingSkill, string[] startingSkills,
            string allSkillsBonusString, string firstTabBonusString, string secondTabBonusString,
            string thirdTabBonusString, string classOnlyBonusString,
            string baseWeaponClass,
            List<InitialEquipment> initialEquipment)
        {
            StartingDexterity = dexterity;
            StartingVitality = vitality;
            StartingStrength = strength;
            StartingEnergy = energy;
            StartingMana = mana;
            StartingHealth = health;
            StartingStamina = stamina;
            StartingManaRegen = manaRegen;

            PerLevelHealth = perlevelhealth;
            PerLevelMana = perlevelmana;
            PerLevelStamina = perlevelstamina;

            PerVitalityHealth = pervitalityhealth;
            PerVitalityStamina = pervitalitystamina;
            PerEnergyMana = perenergymana;

            PerLevelStatPoints = perLevelStatPoints;

            BaseAttackRating = baseatkrating;
            BaseDefenseRating = basedefrating;
            PerDexterityAttackRating = perdexterityatkrating;
            PerDexterityDefenseRating = perdexteritydefrating;

            WalkVelocity = walkVelocity;
            RunVelocity = runVelocity;
            RunDrain = runDrain;

            WalkFrames = walkFrames;
            RunFrames = runFrames;
            SwingFrames = swingFrames;
            SpellFrames = spellFrames;
            GetHitFrames = getHitFrames;
            BowFrames = bowFrames;

            StartingSkill = startingSkill;
            StartingSkills = startingSkills;
            AllSkillsBonusString = allSkillsBonusString;
            FirstTabBonusString = firstTabBonusString;
            SecondTabBonusString = secondTabBonusString;
            ThirdTabBonusString = thirdTabBonusString;
            ClassOnlyBonusString = classOnlyBonusString;

            BaseWeaponClass = baseWeaponClass;

            InitialEquipment = initialEquipment;
        }
    }

    public struct InitialEquipment
    {
        public string name;
        public string location;
        public int count;
    };

    public static class HeroTypeConfigHelper
    {
        public static IHeroTypeConfig ToHeroTypeConfig(this string[] row)
        {
            // rows:

            // 0        1   2   3   4   5   6       7
            // class	str	dex	int	vit	tot	stamina	hpadd
            // 8            9           10          11
            // PercentStr	PercentDex	PercentInt	PercentVit
            // 12           13          14              15
            // ManaRegen	ToHitFactor	WalkVelocity	RunVelocity
            // 16       17      18              19              20
            // RunDrain	Comment	LifePerLevel	StaminaPerLevel	ManaPerLevel
            // 21               22                  23
            // LifePerVitality	StaminaPerVitality	ManaPerMagic	
            // 24
            // StatPerLevel
            // 25       26      27      28      29      30
            // #walk	#run	#swing	#spell	#gethit	#bow	
            // 31           32          
            // BlockFactor	StartSkill	
            // 33 ... 42
            // Skill1 ... Skill10
            // 43           44              45              46              47
            // StrAllSkills StrSkillTab1    StrSkillTab2    StrSkillTab3    StrClassOnly
            // 48
            // baseWClass	
            // 49       50          51          49+3    50+3        51+3
            // item1	item1loc	item1count	item2	item2loc	item2count
            // item3	item3loc	item3count	item4	item4loc	item4count
            // item5	item5loc	item5count	item6	item6loc	item6count
            // item7	item7loc	item7count	item8	item8loc	item8count
            // item9	item9loc	item9count	item10	item10loc	item10count
            // itemn = 49 + (n-1)*3, itemnloc = 50 + (n-1)*3, itemncount = 51 + (n-1)*3

            // NEW TODO: StatPerLevel Skill1 ... Skill10
            //  StrAllSkills StrSkillTab1    StrSkillTab2    StrSkillTab3    StrClassOnly

            List<string> itemNames = new List<string>();
            List<string> itemLocs = new List<string>();
            List<int> itemCounts = new List<int>();
            List<InitialEquipment> initialEquipment = new List<InitialEquipment>();

            for(int i = 49; i <= 78; i+=3)
            {
                var item = new InitialEquipment();
                item.name = row[i];
                item.location = row[i + 1];
                item.count = Convert.ToInt32(row[i + 2]);

                initialEquipment.Add(item);
            }

            return new HeroTypeConfig(
                vitality: Convert.ToInt32(row[4]),
                strength: Convert.ToInt32(row[1]),
                dexterity: Convert.ToInt32(row[2]),
                energy: Convert.ToInt32(row[3]),
                health: Convert.ToInt32(row[7]),
                mana: 0,
                stamina: Convert.ToInt32(row[6]),
                manaRegen: Convert.ToInt32(row[12]),
                perlevelhealth: Convert.ToInt32(row[18]) / 4.0,
                perlevelmana: Convert.ToInt32(row[20]) / 4.0,
                perlevelstamina: Convert.ToInt32(row[19]) / 4.0,
                pervitalityhealth: Convert.ToInt32(row[21]) / 4.0,
                pervitalitystamina: Convert.ToInt32(row[22]) / 4.0,
                perenergymana: Convert.ToInt32(row[23]) / 4.0,
                perLevelStatPoints: Convert.ToInt32(row[24]),
                baseatkrating: Convert.ToInt32(row[13]),
                basedefrating: Convert.ToInt32(row[31]),
                perdexterityatkrating: 5,
                perdexteritydefrating: 4,
                walkVelocity: Convert.ToInt32(row[14]),
                runVelocity: Convert.ToInt32(row[15]),
                runDrain: Convert.ToInt32(row[16]),
                walkFrames: Convert.ToInt32(row[25]),
                runFrames: Convert.ToInt32(row[26]),
                swingFrames: Convert.ToInt32(row[27]),
                spellFrames: Convert.ToInt32(row[28]),
                getHitFrames: Convert.ToInt32(row[29]),
                bowFrames: Convert.ToInt32(row[30]),
                startingSkill: row[32],
                startingSkills: new string[]
                {
                    row[33], row[34], row[35], row[36], row[37], row[38], row[39], row[40], row[41], row[42]
                },
                allSkillsBonusString: row[43],
                firstTabBonusString: row[44],
                secondTabBonusString: row[45],
                thirdTabBonusString: row[46],
                classOnlyBonusString: row[47],
                baseWeaponClass: row[48],
                initialEquipment: initialEquipment);
        }
    }
}
