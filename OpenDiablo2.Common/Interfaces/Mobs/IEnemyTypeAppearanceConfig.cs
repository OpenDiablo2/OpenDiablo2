using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IEnemyTypeAppearanceConfig
    {
        bool HasDeathAnimation { get; }
        bool HasNeutralAnimation { get; }
        bool HasWalkAnimation { get; }
        bool HasGetHitAnimation { get; }
        bool[] HasAttackAnimation { get; } // 1-2
        bool HasBlockAnimation { get; }
        bool HasCastAnimation { get; }
        bool[] HasSkillAnimation { get; } // 1-4
        bool HasCorpseAnimation { get; }
        bool HasKnockbackAnimation { get; }
        // HasSkillSequenceAnimation skipped due to being unused
        bool HasRunAnimation { get; }

        bool HasLifeBar { get; } // does this monster show a lifebar when you scroll over it
        bool HasNameBar { get; } // does this monster show a name when you scroll over it
        bool CannotBeSelected { get; } // if true, monster can never be highlighted
        bool CanCorpseBeSelected { get; } // if true, the corpse can be highlighted

        int BleedType { get; } // how does this monster bleed when hit?
        // 0 = it doesn't, 1 = small blood, 2 = large blood, 3+ = random missiles from missiles.txt, the larger
        // the number, the more missiles (??? this needs to be tested...)
        bool HasShadow { get; } // does this have a shadow?
        int LightRadius { get; } // how large of a light does this emit?
        bool HasUniqueBossColors { get; } // if true, has unique colors when spawned 
        // as a boss.
        bool CompositeDeath { get; } // if true, death animation uses multiple components
        // (not clear if this is just for reference or if it actually has an effect...)

        byte[] LightRGB { get; } // 0: R, 1: G, 2: B
    }
}
