using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class EnemyTypeAppearanceConfig : IEnemyTypeAppearanceConfig
    {
        public bool HasDeathAnimation { get; private set; }
        public bool HasNeutralAnimation { get; private set; }
        public bool HasWalkAnimation { get; private set; }
        public bool HasGetHitAnimation { get; private set; }
        public bool[] HasAttackAnimation { get; private set; } // 1-2
        public bool HasBlockAnimation { get; private set; }
        public bool HasCastAnimation { get; private set; }
        public bool[] HasSkillAnimation { get; private set; } // 1-4
        public bool HasCorpseAnimation { get; private set; }
        public bool HasKnockbackAnimation { get; private set; }
        // HasSkillSequenceAnimation skipped due to being unused
        public bool HasRunAnimation { get; private set; }

        public bool HasLifeBar { get; private set; } // does this monster show a lifebar when you scroll over it
        public bool HasNameBar { get; private set; } // does this monster show a name when you scroll over it
        public bool CannotBeSelected { get; private set; } // if true, monster can never be highlighted
        public bool CanCorpseBeSelected { get; private set; } // if true, the corpse can be highlighted

        public int BleedType { get; private set; } // how does this monster bleed when hit?
        // 0 = it doesn't, 1 = small blood, 2 = large blood, 3+ = random missiles from missiles.txt, the larger
        // the number, the more missiles (??? this needs to be tested...)
        public bool HasShadow { get; private set; } // does this have a shadow?
        public int LightRadius { get; private set; } // how large of a light does this emit?
        public bool HasUniqueBossColors { get; private set; } // if true, has unique colors when spawned 
        // as a boss.
        public bool CompositeDeath { get; private set; } // if true, death animation uses multiple components
        // (not clear if this is just for reference or if it actually has an effect...)

        public byte[] LightRGB { get; private set; } // 0: R, 1: G, 2: B

        public EnemyTypeAppearanceConfig(bool HasDeathAnimation, bool HasNeutralAnimation,
            bool HasWalkAnimation, bool HasGetHitAnimation, bool HasAttack1Animation,
            bool HasAttack2Animation, bool HasBlockAnimation, bool HasCastAnimation,
            bool[] HasSkillAnimation, bool HasCorpseAnimation, bool HasKnockbackAnimation,
            bool HasRunAnimation,
            bool HasLifeBar, bool HasNameBar, bool CannotBeSelected, bool CanCorpseBeSelected,
            int BleedType, bool HasShadow, int LightRadius, bool HasUniqueBossColors, bool CompositeDeath,
            byte[] LightRGB)
        {
            this.HasDeathAnimation = HasDeathAnimation;
            this.HasNeutralAnimation = HasNeutralAnimation;
            this.HasWalkAnimation = HasWalkAnimation;
            this.HasGetHitAnimation = HasGetHitAnimation;
            this.HasAttackAnimation = new bool[] { HasAttack1Animation, HasAttack2Animation };
            this.HasBlockAnimation = HasBlockAnimation;
            this.HasCastAnimation = HasCastAnimation;
            this.HasSkillAnimation = HasSkillAnimation;
            this.HasCorpseAnimation = HasCorpseAnimation;
            this.HasKnockbackAnimation = HasKnockbackAnimation;
            this.HasRunAnimation = HasRunAnimation;

            this.HasLifeBar = HasLifeBar;
            this.HasNameBar = HasNameBar;
            this.CannotBeSelected = CannotBeSelected;
            this.CanCorpseBeSelected = CanCorpseBeSelected;

            this.BleedType = BleedType;
            this.HasShadow = HasShadow;
            this.LightRadius = LightRadius;
            this.HasUniqueBossColors = HasUniqueBossColors;
            this.CompositeDeath = CompositeDeath;

            this.LightRGB = LightRGB;
        }
    }
}
