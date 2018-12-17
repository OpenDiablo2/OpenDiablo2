using OpenDiablo2.Common.Enums.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Interfaces.Mobs
{
    public interface IMissileTypeConfig
    {
        string Name { get; }
        int Id { get; }

        int ClientMovementFunctionId { get; }
        int ClientHitFunctionId { get; }
        int ServerMovementFunctionId { get; }
        int ServerHitFunctionId { get; }
        int ServerDamageFunctionId { get; }
        
        string ServerMovementCalc { get; }
        int[] ServerMovementParams { get; }

        string ClientMovementCalc { get; }
        int[] ClientMovementParams { get; }

        string ServerHitCalc { get; }
        int[] ServerHitParams { get; }

        string ClientHitCalc { get; }
        int[] ClientHitParams { get; }

        string ServerDamageCalc { get; }
        int[] ServerDamageParams { get; }

        int BaseVelocity { get; }
        int MaxVelocity { get; }
        int VelocityPerSkillLevel { get; }

        int Acceleration { get; }
        int Range { get; }
        int ExtraRangePerLevel { get; }

        int LightRadius { get; }
        bool LightFlicker { get; }
        byte[] LightColor { get; }

        int FramesBeforeVisible { get; }
        int FramesBeforeActive { get; }
        bool LoopAnimation { get; }

        string CelFilePath { get; }
        int AnimationRate { get; }
        int AnimationLength { get; }
        int AnimationSpeed { get; }
        int StartingFrame { get; }
        bool AnimationHasSubLoop { get; }
        int AnimationSubLoopStart { get; }
        int AnimationSubLoopEnd { get; }


        int CollisionType { get; }
        bool CollideKill { get; }
        bool CollideFriend { get; }
        bool CollideLast { get; }
        bool CollisionUnknown { get; }
        bool CollisionClient { get; }
        bool CollisionClientSend { get; }
        bool CollisionUseTimer { get; }
        int CollisionTimerLength { get; }
        
        int XOffset { get; }
        int YOffset { get; }
        int ZOffset { get; }
        int Size { get; }

        bool SrcTown { get; }
        bool SrcTownClient { get; }
        bool CanDestroy { get; }

        bool UseAttackRating { get; }

        bool AlwaysExplode { get; }

        bool IsClientExplosion { get; }
        bool AllowedInTown { get; }
        bool NoUniqueMod { get; }
        bool NoMultishotMod { get; }
        int Holy { get; }

        bool CanSlow { get; }
        bool ReturnFire { get; }

        bool KnockIntoGetHit { get; }
        bool SoftHit { get; }
        int KnockbackChance { get; }
        int TransparencyType { get; }
        bool UsesQuantity { get; }
        bool Pierce { get; }
        bool SpecialSetup { get; }

        bool MissileSkill { get; }
        string SkillName { get; }

        int ResultFlags { get; }
        int HitFlags { get; }
        int HitShift { get; }


        bool ApplyMastery { get; }
        int SourceDamage { get; }
        bool HalfWithTwoHandedWeapon { get; }
        int SourceMissileDamage { get; }

        int MinDamage { get; }
        int MaxDamage { get; }
        int[] MinDamagePerLevel { get; }
        int[] MaxDamagePerLevel { get; }
        string PhysicalDamageSynergyCalc { get; }

        eDamageTypes ElementalDamageType { get; }
        int ElementalMinDamage { get; }
        int ElementalMaxDamage { get; }
        int[] ElementalMinDamagePerLevel { get; }
        int[] ElementalMaxDamagePerLevel { get; }
        string ElementalDamageSynergyCalc { get; }
        int ElementalDuration { get; }
        int[] ElementalDurationPerLevel { get; }

        int HitClass { get; }
        bool CanBleed { get; }
        bool AffectedByOpenWounds { get; }
        int DamageRate { get; }

        string SoundTravel { get; }
        string SoundHit { get; }
        string SoundProg { get; }

        string ProgOverlay { get; }
        string ExplosionMissile { get; }
        
        string[] SubMissiles { get; }
        string[] HitSubMissiles { get; }
        string[] ClientSubMissiles { get; }
        string[] ClientHitSubMissiles { get; }
    }
}
