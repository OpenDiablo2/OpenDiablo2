using OpenDiablo2.Common.Enums;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IEnemyTypeConfig
    {
        string InternalName { get; }
        int InternalId { get; }
        string DisplayName { get; }
        string Type { get; }
        string Descriptor { get; }

        int BaseId { get; }
        int PopulateId { get; }
        bool Spawned { get; }
        bool Beta { get; }
        bool ClientOnly { get; }
        bool NoMap { get; }

        int SizeX { get; }
        int SizeY { get; }
        int Height { get; }
        bool NoOverlays { get; }
        int OverlayHeight { get; }

        int Velocity { get; }
        int RunVelocity { get; }

        bool CanStealFrom { get; }
        int ColdEffect { get; }

        bool Rarity { get; }

        int MinimumGrouping { get; }
        int MaximumGrouping { get; }

        string BaseWeapon { get; }

        int[] AIParams { get; } // up to 5

        bool Allied { get; } // Is this enemy on the player's side?
        bool IsNPC { get; } // is this an NPC?
        bool IsCritter { get; } // is this a critter? (e.g. the chickens)
        bool CanEnterTown { get; } // can this enter town or not?

        int HealthRegen { get; } // hp regen per minute

        bool IsDemon { get; } // demon?
        bool IsLarge { get; } // size large (e.g. bosses)
        bool IsSmall { get; } // size small (e.g. fallen)
        bool IsFlying { get; } // can move over water for instance
        bool CanOpenDoors { get; }
        int SpawningColumn { get; } // is the monster area restricted
        // 0 = no, spawns through levels.txt, 1-3 unknown? 
        // TODO: understand spawningcolumn
        bool IsBoss { get; }
        bool IsInteractable { get; } // can this be interacted with? like an NPC

        bool IsKillable { get; }
        bool CanBeConverted { get; } // if true, can be switched to Allied by spells like Conversion

        int HitClass { get; } // TODO: find out what this is

        bool HasSpecialEndDeath { get; } // marks if a monster dies a special death
        bool DeadCollision { get; } // if true, corpse has collision   

        bool CanBeRevivedByOtherMonsters { get; }

        // appearance config
        IEnemyTypeAppearanceConfig AppearanceConfig { get; }

        // combat config
        IEnemyTypeCombatConfig CombatConfig { get; }

        // difficulty configs
        IEnemyTypeDifficultyConfig NormalDifficultyConfig { get; }
        IEnemyTypeDifficultyConfig NightmareDifficultyConfig { get; }
        IEnemyTypeDifficultyConfig HellDifficultyConfig { get; }

        IEnemyTypeDifficultyConfig GetDifficultyConfig(eDifficulty Difficulty);
    }
}
