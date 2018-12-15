using OpenDiablo2.Common.Models.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IHeroTypeConfig
    {
        int StartingVitality { get; }
        int StartingStrength { get; }
        int StartingDexterity { get; }
        int StartingEnergy { get; }
        int StartingHealth { get; }
        int StartingMana { get; }
        int StartingStamina { get; }
        int StartingManaRegen { get; }

        double PerLevelHealth { get; } // NOTE: these are doubles because some classes have
        // e.g. 1.5 mana per level, which means they get 1 mana on even levels (e.g. -> 2)
        // and 2 mana on odd levels (e.g. -> 3)
        double PerLevelMana { get; }
        double PerLevelStamina { get; }

        double PerVitalityHealth { get; }
        double PerVitalityStamina { get; }
        double PerEnergyMana { get; }

        int PerLevelStatPoints { get; }

        int BaseAttackRating { get; }
        int PerDexterityAttackRating { get; }

        int BaseDefenseRating { get; }
        int PerDexterityDefenseRating { get; }

        int WalkVelocity { get; }
        int RunVelocity { get; }
        int RunDrain { get; }

        int WalkFrames { get; }
        int RunFrames { get; }
        int SwingFrames { get; }
        int SpellFrames { get; }
        int GetHitFrames { get; }
        int BowFrames { get; }

        string StartingSkill { get; }
        string[] StartingSkills { get; }

        string AllSkillsBonusString { get; }
        string FirstTabBonusString { get; }
        string SecondTabBonusString { get; }
        string ThirdTabBonusString { get; }
        string ClassOnlyBonusString { get; }
        string BaseWeaponClass { get; }

        List<InitialEquipment> InitialEquipment { get; }
    }
}
