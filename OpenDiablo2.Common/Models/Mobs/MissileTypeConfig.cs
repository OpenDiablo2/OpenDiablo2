using OpenDiablo2.Common.Enums.Mobs;
using OpenDiablo2.Common.Interfaces.Mobs;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Models.Mobs
{
    public class MissileTypeConfig : IMissileTypeConfig
    {
        public string Name { get; private set; }
        public int Id { get; private set; }

        public int ClientMovementFunctionId { get; private set; } // defines how the missile acts graphically when it moves
        public int ClientHitFunctionId { get; private set; } // defines how the missile acts graphically when it hits
        public int ServerMovementFunctionId { get; private set; } // defines how the missile actually moves
        public int ServerHitFunctionId { get; private set; } // defines how the missile actually hits
        public int ServerDamageFunctionId { get; private set; } // defines if the missile has special damage effect
        // like "mana burn" or "stuns", etc.

        public string ServerMovementCalc { get; private set; } // used by movement function
        // what this does really needs to be researched
        // values for this are like "0", "3", "lvl*2" and the comment describes this as "# subloops"
        public int[] ServerMovementParams { get; private set; } // what these do depend on the movement function id chosen

        public string ClientMovementCalc { get; private set; } // the client equivalent of ServerMovementCalc
        public int[] ClientMovementParams { get; private set; } // affect client movement func

        public string ServerHitCalc { get; private set; } // used by hit function
        public int[] ServerHitParams { get; private set; } // used by hit function

        public string ClientHitCalc { get; private set; } // used by hit function
        public int[] ClientHitParams { get; private set; } // used by hit function

        public string ServerDamageCalc { get; private set; } // used when damage is dealt
        public int[] ServerDamageParams { get; private set; } // used by damage function

        public int BaseVelocity { get; private set; }
        public int MaxVelocity { get; private set; }
        public int VelocityPerSkillLevel { get; private set; } // extra vel per skill level
        // NOTE: may be unused! requires some research
        public int Acceleration { get; private set; }
        public int Range { get; private set; }
        public int ExtraRangePerLevel { get; private set; }

        public int LightRadius { get; private set; }
        public bool LightFlicker { get; private set; }
        public byte[] LightColor { get; private set; } // rgb

        public int FramesBeforeVisible { get; private set; } // how many anim frames before it becomes visible
        public int FramesBeforeActive { get; private set; } // anim frames before it can do anything, e.g. collide
        public bool LoopAnimation { get; private set; } // true: repeat animation until missile hits range or is otherwise destroyed
        // false: repeat animation once, then vanish (WARNING: only becomes invisible, not deleted)
        public string CelFilePath { get; private set; }
        public int AnimationRate { get; private set; } // does not seem to be used
        public int AnimationLength { get; private set; } // length of animation for one direction
        public int AnimationSpeed { get; private set; } // actually used, frames per second
        public int StartingFrame { get; private set; } // 'randstart' is a misnomer, actually just starts at this frame
        public bool AnimationHasSubLoop { get; private set; } // if true, will repeat in a certain range of frames
        //instead of repeating the whole animation again
        public int AnimationSubLoopStart { get; private set; } // what frame to start at if it has a subloop
        public int AnimationSubLoopEnd { get; private set; } // what frame to end at in subloop
        // when it hits this, goes back to subloop start or goes to end if missile is set to die (out of range)

        public int CollisionType { get; private set; }
        //0: no collision
        //1: cols with units
        //3: cols with units and walls
        //6: cols with walls
        //8: cols with units, walls, and floors
        public bool CollideKill { get; private set; } // if true, this is killed on collision
        public bool CollideFriend { get; private set; } // if true, can collide with allies
        public bool CollideLast { get; private set; } // unknown
        public bool CollisionUnknown { get; private set; } // unknown
        public bool CollisionClient { get; private set; } // seems to be unused; unknown
        public bool CollisionClientSend { get; private set; } // unknown
        public bool CollisionUseTimer { get; private set; } // use collision timer after first collision
        public int CollisionTimerLength { get; private set; } // how long the above timer is
        // for instance, if timerlength is 40, a missile will be unable to damage another unit 
        // until 40 frames have passed

        public int XOffset { get; private set; }
        public int YOffset { get; private set; }
        public int ZOffset { get; private set; }
        public int Size { get; private set; } // diameter in subtiles

        public bool SrcTown { get; private set; } // unknown, probably means missile dies if caster enters town
        public bool SrcTownClient { get; private set; } // client version of above
        public bool CanDestroy { get; private set; } // unknown

        public bool UseAttackRating { get; private set; } // if true, uses attack rating to calculate hit
        // if false, has a fixed 95% chance of hitting
        public bool AlwaysExplode { get; private set; } // if true, always explodes on death even if it
        // just runs out of range and doesn't hit anything
        public bool IsClientExplosion { get; private set; } // if true, doesn't exist on server, only for graphical effect
        public bool AllowedInTown { get; private set; } // if true, doesn't vanish in town
        public bool NoUniqueMod { get; private set; } // if true, does not receive bonuses from unique monster modifiers
        public bool NoMultishotMod { get; private set; } // if true, not affected by multishot modifier of bosses
        public int Holy { get; private set; } // Controls what this missile can hit
        // 0 = any unit, 1 = only undead, 2 = demons, 3 = any units (COULD USE EXTRA TESTING)
        public bool CanSlow { get; private set; } // specifically means (if true) that this is effected by skill_handofathena
        public bool ReturnFire { get; private set; } // if true, can trigger collision events on target
        // collision events are things that happen when getting hit
        public bool KnockIntoGetHit { get; private set; } // if true, can knock target into hit recovery mode
        public bool SoftHit { get; private set; } // unknown
        public int KnockbackChance { get; private set; } // how often to trigger hit recovery mode on  target
        public int TransparencyType { get; private set; } // 0= normal, 1 = alpha blending, 2 = special blending (?)
        public bool UsesQuantity { get; private set; } //does it use up quantity in conjunction with its skill
        public bool Pierce { get; private set; } // if true, affected by amazon's Pierce
        public bool SpecialSetup { get; private set; } // unknown. associated with potions

        public bool MissileSkill { get; private set; } // if true, the splash radius of this missile 
        // gains the elemental damage of items and loses all other damage modifiers (???)
        public string SkillName { get; private set; } // if set to a non-empty value, this missile
        // inherits its damage info from the skill with this name
        public int ResultFlags { get; private set; } // unknown, probably just for reference
        public int HitFlags { get; private set; } // unknown, probably just for reference
        public int HitShift { get; private set; } // hitshift is a bitwise shift that is applied
        // to all of the damage values. Life is stored as 256ths. In a typical scenario, a 
        // hitshift of 8 and a damage of 1 would become a real damage of 256
        public bool ApplyMastery { get; private set; } // not 100% clear, but seems to indicate
        // whether or not mastery bonuses are applied for MeteorFire
        public int SourceDamage { get; private set; } // this value is in 128ths for some reason
        // so 128 = 100%. This has to do with whether the damage you would do with a normal attack
        // is added to the damage of this missile. This includes any modifiers on your normal attack,
        // such as lifesteal and other elemental damage, or bonuses from skills on your normal attack.
        public bool HalfWithTwoHandedWeapon { get; private set; } // if true, does half the damage if
        // used with a two-handed weapon
        public int SourceMissileDamage { get; private set; } // if this is created by another missile,
        // tells this one how much damage to carry over (128 = 100%) 
        public int MinDamage { get; private set; } // remember, these are shifted by hitshift 
        public int MaxDamage { get; private set; }
        public int[] MinDamagePerLevel { get; private set; } // damage to add per level in the skill
        public int[] MaxDamagePerLevel { get; private set; } // or if no skill is associated, level
        // of the attacking unit. There are 5 min/max damage per level values, and they correspond
        // to different level ranges: 2-8, 9-16, 17-22, 23-28 and 29+
        public string PhysicalDamageSynergyCalc { get; private set; } // a percentage bonus is applied
        // to physical damage dealt by this missile based on this calculation. Does not include base damage
        // from srcdamage
        public eDamageTypes ElementalDamageType { get; private set; }
        public int ElementalMinDamage { get; private set; } // remember, these are shifted by hitshift 
        public int ElementalMaxDamage { get; private set; }
        public int[] ElementalMinDamagePerLevel { get; private set; } // damage to add per level in the skill
        public int[] ElementalMaxDamagePerLevel { get; private set; } // or if no skill is associated, level
        // of the attacking unit. There are 5 min/max damage per level values, and they correspond
        // to different level ranges: 2-8, 9-16, 17-22, 23-28 and 29+
        public string ElementalDamageSynergyCalc { get; private set; } // a percentage bonus is applied
        // to elemental damage dealt by this missile based on this calculation.
        public int ElementalDuration { get; private set; } // in frames (25 = 1 second), duration used
        // for stun poison and burning damage
        public int[] ElementalDurationPerLevel { get; private set; } // extra duration per level in
        // the skill (or unit's level if none is specified). There are 3 of these values and they correspond
        // to different level ranges, but these level ranges are currently unknown (TODO)

        public int HitClass { get; private set; } // a lot is unknown about how hitclass works,
        // but it appears to mostly influence client-side effects like what sound effect plays on hit
        // NumDirections is just for reference
        public bool CanBleed { get; private set; }
        public bool AffectedByOpenWounds { get; private set; }
        public int DamageRate { get; private set; } // this needs to be better understood, but seems to 
        // indicate how often the magic_damage_reduced stat can be applied to this missile.

        public string SoundTravel { get; private set; }
        public string SoundHit { get; private set; }
        public string SoundProg { get; private set; } // plays at special events determined by client side
        // collision functions

        public string ProgOverlay { get; private set; } // overlay used for special events determined by
        // client side collision functions
        public string ExplosionMissile { get; private set; } // explosion missile created when missile
        // successfully collides with a unit or obstacle (or whenever it dies if AlwaysExplodes is true)

        public string[] SubMissiles { get; private set; } // missile names spawned by server movement functions (1-3)
        public string[] HitSubMissiles { get; private set; } // spawned by server collision functions (1-4)
        public string[] ClientSubMissiles { get; private set; } // client version of submissiles 1-3
        public string[] ClientHitSubMissiles { get; private set; } // client version of hitsubmissiles 1-4
        

        public MissileTypeConfig(string Name, int Id,
            int ClientMovementFunctionId, int ClientHitFunctionId, int ServerMovementFunctionId, 
            int ServerHitFunctionId, int ServerDamageFunctionId, 
            string ServerMovementCalc, int[] ServerMovementParams, 
            string ClientMovementCalc, int[] ClientMovementParams, 
            string ServerHitCalc, int[] ServerHitParams,
            string ClientHitCalc, int[] ClientHitParams,
            string ServerDamageCalc, int[] ServerDamageParams,
            int BaseVelocity, int MaxVelocity, int VelocityPerSkillLevel, int Acceleration, int Range, int ExtraRangePerLevel,
            int LightRadius, bool LightFlicker, byte[] LightColor,
            int FramesBeforeVisible, int FramesBeforeActive, bool LoopAnimation, string CelFilePath,
            int AnimationRate, int AnimationLength, int AnimationSpeed, int StartingFrame, 
            bool AnimationHasSubLoop, int AnimationSubLoopStart, int AnimationSubLoopEnd,
            int CollisionType, bool CollideKill, bool CollideFriend, bool CollideLast,
            bool CollisionUnknown, bool CollisionClient, bool CollisionClientSend,
            bool CollisionUseTimer, int CollisionTimerLength,
            int XOffset, int YOffset, int ZOffset, int Size,
            bool SrcTown, bool SrcTownClient, bool CanDestroy,
            bool UseAttackRating, bool AlwaysExplode, bool IsClientExplosion,
            bool AllowedInTown, bool NoUniqueMod, bool NoMultishotMod, int Holy, 
            bool CanSlow, bool ReturnFire, bool KnockIntoGetHit, bool SoftHit,
            int KnockbackChance, int TransparencyType, bool UsesQuantity, bool Pierce, bool SpecialSetup,
            bool MissileSkill, string SkillName, int ResultFlags, int HitFlags, int HitShift,
            bool ApplyMastery, int SourceDamage, bool HalfWithTwoHandedWeapon, int SourceMissileDamage,
            int MinDamage, int MaxDamage, int[] MinDamagePerLevel, int[] MaxDamagePerLevel,
            string PhysicalDamageSynergyCalc, eDamageTypes ElementalDamageType, int ElementalMinDamage, 
            int ElementalMaxDamage, int[] ElementalMinDamagePerLevel, int[] ElementalMaxDamagePerLevel,
            string ElementalDamageSynergyCalc, int ElementalDuration, int[] ElementalDurationPerLevel,
            int HitClass, bool CanBleed, bool AffectedByOpenWounds, int DamageRate,
            string SoundTravel, string SoundHit, string SoundProg,
            string ProgOverlay, string ExplosionMissile,
            string[] SubMissiles, string[] HitSubMissiles,
            string[] ClientSubMissiles, string[] ClientHitSubMissiles
            )
        {
            this.Name = Name;
            this.Id = Id;

            this.ClientMovementFunctionId = ClientMovementFunctionId;
            this.ClientHitFunctionId = ClientHitFunctionId;
            this.ServerMovementFunctionId = ServerMovementFunctionId;
            this.ServerHitFunctionId = ServerHitFunctionId;
            this.ServerDamageFunctionId = ServerDamageFunctionId;
            this.ServerMovementCalc = ServerMovementCalc;
            this.ServerMovementParams = ServerMovementParams;
            this.ClientMovementCalc = ClientMovementCalc;
            this.ClientMovementParams = ClientMovementParams;
            this.ServerHitCalc = ServerHitCalc;
            this.ServerHitParams = ServerHitParams;
            this.ClientHitCalc = ClientHitCalc;
            this.ClientHitParams = ClientHitParams;
            this.ServerDamageCalc = ServerDamageCalc;
            this.ServerDamageParams = ServerDamageParams;

            this.BaseVelocity = BaseVelocity;
            this.MaxVelocity = MaxVelocity;
            this.VelocityPerSkillLevel = VelocityPerSkillLevel;
            this.Acceleration = Acceleration;
            this.Range = Range;
            this.ExtraRangePerLevel = ExtraRangePerLevel;

            this.LightRadius = LightRadius;
            this.LightFlicker = LightFlicker;
            this.LightColor = LightColor;

            this.FramesBeforeVisible = FramesBeforeVisible;
            this.FramesBeforeActive = FramesBeforeActive;
            this.LoopAnimation = LoopAnimation;
            this.CelFilePath = CelFilePath;
            this.AnimationRate = AnimationRate;
            this.AnimationLength = AnimationLength;
            this.AnimationSpeed = AnimationSpeed;
            this.StartingFrame = StartingFrame;
            this.AnimationHasSubLoop = AnimationHasSubLoop;
            this.AnimationSubLoopStart = AnimationSubLoopStart;
            this.AnimationSubLoopEnd = AnimationSubLoopEnd;

            this.CollisionType = CollisionType;
            this.CollideKill = CollideKill;
            this.CollideFriend = CollideFriend;
            this.CollideLast = CollideLast;
            this.CollisionUnknown = CollisionUnknown;
            this.CollisionClient = CollisionClient;
            this.CollisionClientSend = CollisionClientSend;
            this.CollisionUseTimer = CollisionUseTimer;
            this.CollisionTimerLength = CollisionTimerLength;

            this.XOffset = XOffset;
            this.YOffset = YOffset;
            this.ZOffset = ZOffset;
            this.Size = Size;
            this.SrcTown = SrcTown;
            this.SrcTownClient = SrcTownClient;
            this.CanDestroy = CanDestroy;
            this.UseAttackRating = UseAttackRating;
            this.AlwaysExplode = AlwaysExplode;
            this.IsClientExplosion = IsClientExplosion;
            this.AllowedInTown = AllowedInTown;
            this.NoUniqueMod = NoUniqueMod;
            this.NoMultishotMod = NoMultishotMod;
            this.Holy = Holy;
            this.CanSlow = CanSlow;
            this.ReturnFire = ReturnFire;
            this.KnockIntoGetHit = KnockIntoGetHit;
            this.SoftHit = SoftHit;
            this.KnockbackChance = KnockbackChance;
            this.TransparencyType = TransparencyType;
            this.UsesQuantity = UsesQuantity;
            this.Pierce = Pierce;
            this.SpecialSetup = SpecialSetup;

            this.MissileSkill = MissileSkill;
            this.SkillName = SkillName;
            this.ResultFlags = ResultFlags;
            this.HitFlags = HitFlags;
            this.HitShift = HitShift;
            this.ApplyMastery = ApplyMastery;
            this.SourceDamage = SourceDamage;
            this.HalfWithTwoHandedWeapon = HalfWithTwoHandedWeapon;
            this.SourceMissileDamage = SourceMissileDamage;
            this.MinDamage = MinDamage;
            this.MaxDamage = MaxDamage;
            this.MinDamagePerLevel = MinDamagePerLevel;
            this.MaxDamagePerLevel = MaxDamagePerLevel;
            this.PhysicalDamageSynergyCalc = PhysicalDamageSynergyCalc;
            this.ElementalDamageType = ElementalDamageType;
            this.ElementalMinDamage = ElementalMinDamage;
            this.ElementalMaxDamage = ElementalMaxDamage;
            this.ElementalMinDamagePerLevel = ElementalMinDamagePerLevel;
            this.ElementalMaxDamagePerLevel = ElementalMaxDamagePerLevel;
            this.ElementalDamageSynergyCalc = ElementalDamageSynergyCalc;
            this.ElementalDuration = ElementalDuration;
            this.ElementalDurationPerLevel = ElementalDurationPerLevel;

            this.HitClass = HitClass;
            this.CanBleed = CanBleed;
            this.AffectedByOpenWounds = AffectedByOpenWounds;
            this.DamageRate = DamageRate;
            this.SoundTravel = SoundTravel;
            this.SoundHit = SoundHit;
            this.SoundProg = SoundProg;
            this.ProgOverlay = ProgOverlay;
            this.ExplosionMissile = ExplosionMissile;
            this.SubMissiles = SubMissiles;
            this.HitSubMissiles = HitSubMissiles;
            this.ClientSubMissiles = ClientSubMissiles;
            this.ClientHitSubMissiles = ClientHitSubMissiles;
        }
    }

    public static class MissileTypeConfigHelper
    {
        // Missile | Id | pCltDoFunc | pCltHitFunc | pSrvDoFunc | pSrvHitFunc | pSrvDmgFunc | SrvCalc1 | *srv calc 1 desc | Param1 | 
        // 0         1    2            3             4            5             6             7          8                  9        
        // *param1 desc | Param2 | *param2 desc | Param3 | *param3 desc | Param4 | *param4 desc | Param5 | *param5 desc | CltCalc1 | 
        // 10             11       12             13       14             15       16             17       18             19         
        // *client calc 1 desc | CltParam1 | *client param1 desc | CltParam2 | *client param2 desc | CltParam3 | *client param3 desc | 
        // 20                    21          22                    23          24                    25          26                    
        // CltParam4 | *client param4 desc | CltParam5 | *client param5 desc | SHitCalc1 | *server hit calc 1 desc | sHitPar1 | *server hit param1 desc | 
        // 27          28                    29          30                    31          32                        33         34                        
        // sHitPar2 | *server hit param2 desc | sHitPar3 | *server hit param3 desc | CHitCalc1 | *client hit calc1 desc | cHitPar1 | 
        // 35         36                        37         38                        39          40                       41         
        // *client hit param1 desc | cHitPar2 | *client hit param2 desc | cHitPar3 | *client hit param3 desc | DmgCalc1 | *damage calc 1 | 
        // 42                        43         44                        45         46                        47         48               
        // dParam1 | *damage param1 desc | dParam2 | *damage param2 desc | Vel | MaxVel | VelLev | Accel | Range | LevRange | Light | 
        // 49        50                    51        52                    53    54       55       56      57      58         59      
        // Flicker | Red | Green | Blue | InitSteps | Activate | LoopAnim | CelFile | animrate | AnimLen | AnimSpeed | RandStart | SubLoop | 
        // 60        61    62      63     64          65         66         67        68         69        70          71          72        
        // SubStart | SubStop | CollideType | CollideKill | CollideFriend | LastCollide | Collision | ClientCol | ClientSend | NextHit | 
        // 73         74        75            76            77              78            79          80          81           82        
        // NextDelay | xoffset | yoffset | zoffset | Size | SrcTown | CltSrcTown | CanDestroy | ToHit | AlwaysExplode | Explosion | 
        // 83          84        85        86        87     88        89           90           91      92              93          
        // Town | NoUniqueMod | NoMultiShot | Holy | CanSlow | ReturnFire | GetHit | SoftHit | KnockBack | Trans | Qty | Pierce | SpecialSetup | 
        // 94     95            96            97     98        99           100      101       102         103     104   105      106            
        // MissileSkill | Skill | ResultFlags | HitFlags | HitShift | ApplyMastery | SrcDamage | Half2HSrc | SrcMissDmg | MinDamage | 
        // 107            108     109           110        111        112            113         114         115          116         
        // MinLevDam1 | MinLevDam2 | MinLevDam3 | MinLevDam4 | MinLevDam5 | MaxDamage | MaxLevDam1 | MaxLevDam2 | MaxLevDam3 | MaxLevDam4 | 
        // 117          118          119          120          121          122         123          124          125          126          
        // MaxLevDam5 | DmgSymPerCalc | EType | EMin | MinELev1 | MinELev2 | MinELev3 | MinELev4 | MinELev5 | Emax | MaxELev1 | MaxELev2 | 
        // 127          128             129     130    131        132        133        134        135        136    137        138        
        // MaxELev3 | MaxELev4 | MaxELev5 | EDmgSymPerCalc | ELen | ELevLen1 | ELevLen2 | ELevLen3 | HitClass | NumDirections | LocalBlood | 
        // 139        140        141        142              143    144        145        146        147        148             149          
        // DamageRate | TravelSound | HitSound | ProgSound | ProgOverlay | ExplosionMissile | SubMissile1 | SubMissile2 | SubMissile3 | 
        // 150          151           152        153         154           155                156           157           158           
        // HitSubMissile1 | HitSubMissile2 | HitSubMissile3 | HitSubMissile4 | CltSubMissile1 | CltSubMissile2 | CltSubMissile3 | CltHitSubMissile1 | 
        // 159              160              161              162              163              164              165              166                 

        private static eDamageTypes DamageTypeFromString(string s)
        {
            switch (s)
            {
                case "fire":
                    return eDamageTypes.FIRE;
                case "ltng":
                    return eDamageTypes.LIGHTNING;
                case "mag":
                    return eDamageTypes.MAGIC;
                case "cold":
                    return eDamageTypes.COLD;
                case "pois":
                    return eDamageTypes.POISON;
                case "life":
                    return eDamageTypes.LIFE_STEAL;
                case "mana":
                    return eDamageTypes.MANA_STEAL;
                case "stam":
                    return eDamageTypes.STAMINA_STEAL;
                case "stun":
                    return eDamageTypes.STUN; // note: stun damage is ignored, only uses duration
                case "rand":
                    return eDamageTypes.RANDOM;
                case "burn":
                    return eDamageTypes.BURN;
                case "frze":
                    return eDamageTypes.FREEZE;
                default:
                    return eDamageTypes.NONE;
            }
        }

        private static int IntOrZero(string s)
        {
            if (string.IsNullOrWhiteSpace(s) || s.StartsWith("*")) // comments begin with *
            {
                return 0;
            }

            return Convert.ToInt32(s);
        }

        public static IMissileTypeConfig ToMissileTypeConfig(this string[] row)
        {
            return new MissileTypeConfig(
                Name: row[0],
                Id: Convert.ToInt32(row[1]),

                ClientMovementFunctionId: IntOrZero(row[2]),
                ClientHitFunctionId: IntOrZero(row[3]),
                ServerMovementFunctionId: IntOrZero(row[4]),
                ServerHitFunctionId: IntOrZero(row[5]),
                ServerDamageFunctionId: IntOrZero(row[6]),
                ServerMovementCalc: row[7],
                ServerMovementParams: new int[] {
                    IntOrZero(row[9]),
                    IntOrZero(row[11]),
                    IntOrZero(row[13]),
                    IntOrZero(row[15]),
                    IntOrZero(row[17]),
                },
                ClientMovementCalc: row[19],
                ClientMovementParams: new int[]{
                    IntOrZero(row[21]),
                    IntOrZero(row[23]),
                    IntOrZero(row[25]),
                    IntOrZero(row[27]),
                    IntOrZero(row[29])
                },
                ServerHitCalc: row[31],
                ServerHitParams: new int[]{
                    IntOrZero(row[33]),
                    IntOrZero(row[35]),
                    IntOrZero(row[37])
                },
                ClientHitCalc: row[39],
                ClientHitParams: new int[]{
                    IntOrZero(row[41]),
                    IntOrZero(row[43]),
                    IntOrZero(row[45])
                },
                ServerDamageCalc: row[47],
                ServerDamageParams: new int[]{
                    IntOrZero(row[49]),
                    IntOrZero(row[51])
                },


                BaseVelocity: IntOrZero(row[53]),
                MaxVelocity: IntOrZero(row[54]),
                VelocityPerSkillLevel: IntOrZero(row[55]),
                Acceleration: IntOrZero(row[56]),
                Range: IntOrZero(row[57]),
                ExtraRangePerLevel: IntOrZero(row[58]),

                LightRadius: IntOrZero(row[59]),
                LightFlicker: (row[60] == "1"),
                LightColor: new byte[] {
                    Convert.ToByte(IntOrZero(row[61])),
                    Convert.ToByte(IntOrZero(row[62])),
                    Convert.ToByte(IntOrZero(row[63]))
                },

                FramesBeforeVisible: IntOrZero(row[64]),
                FramesBeforeActive: IntOrZero(row[65]),
                LoopAnimation: (row[66] == "1"),
                CelFilePath: row[67],

                AnimationRate: IntOrZero(row[68]),
                AnimationLength: IntOrZero(row[69]),
                AnimationSpeed: IntOrZero(row[70]),
                StartingFrame: IntOrZero(row[71]),
                AnimationHasSubLoop: (row[72] == "1"),
                AnimationSubLoopStart: IntOrZero(row[73]),
                AnimationSubLoopEnd: IntOrZero(row[74]),

                CollisionType: IntOrZero(row[75]),
                CollideKill: (row[76] == "1"),
                CollideFriend: (row[77] == "1"),
                CollideLast: (row[78] == "1"),
                CollisionUnknown: (row[79] == "1"),
                CollisionClient: (row[80] == "1"),
                CollisionClientSend: (row[81] == "1"),
                CollisionUseTimer: (row[82] == "1"),
                CollisionTimerLength: IntOrZero(row[83]),

                XOffset: IntOrZero(row[84]),
                YOffset: IntOrZero(row[85]),
                ZOffset: IntOrZero(row[86]),
                Size: IntOrZero(row[87]),

                SrcTown: (row[88] == "1"),
                SrcTownClient: (row[89] == "1"),
                CanDestroy: (row[90] == "1"),

                UseAttackRating: (row[91] == "1"),

                AlwaysExplode: (row[92] == "1"),

                IsClientExplosion: (row[93] == "1"),
                AllowedInTown: (row[94] == "1"),
                NoUniqueMod: (row[95] == "1"),
                NoMultishotMod: (row[96] == "1"),
                Holy: IntOrZero(row[97]),

                CanSlow: (row[98] == "1"),
                ReturnFire: (row[99] == "1"),

                KnockIntoGetHit: (row[100] == "1"),
                SoftHit: (row[101] == "1"),
                KnockbackChance: IntOrZero(row[102]),
                TransparencyType: IntOrZero(row[103]),
                UsesQuantity: (row[104] == "1"),
                Pierce: (row[105] == "1"),
                SpecialSetup: (row[106] == "1"),


                MissileSkill: (row[107] == "1"),
                SkillName: row[108],

                ResultFlags: IntOrZero(row[109]),
                HitFlags: IntOrZero(row[110]),
                HitShift: IntOrZero(row[111]),

                ApplyMastery: (row[112] == "1"),
                SourceDamage: IntOrZero(row[113]),
                HalfWithTwoHandedWeapon: (row[114] == "1"),
                SourceMissileDamage: IntOrZero(row[115]),

                MinDamage: IntOrZero(row[116]),
                MaxDamage: IntOrZero(row[122]),
                MinDamagePerLevel: new int[] 
                {
                    IntOrZero(row[117]),
                    IntOrZero(row[118]),
                    IntOrZero(row[119]),
                    IntOrZero(row[120]),
                    IntOrZero(row[121]),
                },
                MaxDamagePerLevel: new int[] {
                    IntOrZero(row[123]),
                    IntOrZero(row[124]),
                    IntOrZero(row[125]),
                    IntOrZero(row[126]),
                    IntOrZero(row[127]),
                },
                PhysicalDamageSynergyCalc: row[128],

                ElementalDamageType: DamageTypeFromString(row[129]),
                ElementalMinDamage: IntOrZero(row[130]),
                ElementalMaxDamage: IntOrZero(row[136]),
                ElementalMinDamagePerLevel: new int[] {
                    IntOrZero(row[131]),
                    IntOrZero(row[132]),
                    IntOrZero(row[133]),
                    IntOrZero(row[134]),
                    IntOrZero(row[135]),
                },
                ElementalMaxDamagePerLevel: new int[] {
                    IntOrZero(row[137]),
                    IntOrZero(row[138]),
                    IntOrZero(row[139]),
                    IntOrZero(row[140]),
                    IntOrZero(row[141]),
                },
                ElementalDamageSynergyCalc: row[142],
                ElementalDuration: IntOrZero(row[143]),
                ElementalDurationPerLevel: new int[] {
                    IntOrZero(row[144]),
                    IntOrZero(row[145]),
                    IntOrZero(row[146]),
                },

                HitClass: IntOrZero(row[147]),
                CanBleed: (row[149] == "1" || row[149] == "2"),
                AffectedByOpenWounds: (row[149] == "2"),
                DamageRate: IntOrZero(row[150]),

                SoundTravel: row[151],
                SoundHit: row[152],
                SoundProg: row[153],

                ProgOverlay: row[154],
                ExplosionMissile: row[155],

                SubMissiles: new string[]
                {
                    row[156],
                    row[157],
                    row[158]
                },
                HitSubMissiles: new string[]
                {
                    row[159],
                    row[160],
                    row[161],
                    row[162]
                },
                ClientSubMissiles: new string[]
                {
                    row[163],
                    row[164],
                    row[165]
                },
                ClientHitSubMissiles: new string[]
                {
                    row[166],
                    row[167],
                    row[168],
                    row[169]
                }
                );
        }
    }
}
