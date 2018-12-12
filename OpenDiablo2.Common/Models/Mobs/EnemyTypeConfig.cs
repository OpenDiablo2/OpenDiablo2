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
    public class EnemyTypeConfig : IEnemyTypeConfig
    {
        public string InternalName { get; private set; }
        public string DisplayName { get; private set; } 
        public string Type { get; private set; }
        public string Descriptor { get; private set; }

        public int BaseId { get; private set; }
        public int PopulateId { get; private set; }
        public bool Spawned { get; private set; }
        public bool Beta { get; private set; }
        public string Code { get; private set; }
        public bool ClientOnly { get; private set; }
        public bool NoMap { get; private set; }

        public int SizeX { get; private set; }
        public int SizeY { get; private set; }
        public int Height { get; private set; }
        public bool NoOverlays { get; private set; }
        public int OverlayHeight { get; private set; }

        public int Velocity { get; private set; }
        public int RunVelocity { get; private set; }

        public bool CanStealFrom { get; private set; }
        public int ColdEffect { get; private set; }

        public bool Rarity { get; private set; }
        
        public int MinimumGrouping { get; private set; }
        public int MaximumGrouping { get; private set; }

        public string BaseWeapon { get; private set; }

        public int[] AIParams { get; private set; } // up to 5

        public bool Allied { get; private set; } // Is this enemy on the player's side?
        public bool IsNPC { get; private set; } // is this an NPC?
        public bool IsCritter { get; private set; } // is this a critter? (e.g. the chickens)
        public bool CanEnterTown { get; private set; } // can this enter town or not?

        public int HealthRegen { get; private set; } // hp regen per minute

        public bool IsDemon { get; private set; } // demon?
        public bool IsLarge { get; private set; } // size large (e.g. bosses)
        public bool IsSmall { get; private set; } // size small (e.g. fallen)
        public bool IsFlying { get; private set; } // can move over water for instance
        public bool CanOpenDoors { get; private set; } 
        public int SpawningColumn { get; private set; } // is the monster area restricted
        // 0 = no, spawns through levels.txt, 1-3 unknown? 
        // TODO: understand spawningcolumn
        public bool IsBoss { get; private set; }
        public bool IsInteractable { get; private set; } // can this be interacted with? like an NPC

        public bool IsKillable { get; private set; }
        public bool CanBeConverted { get; private set; } // if true, can be switched to Allied by spells like Conversion

        public int HitClass { get; private set; } // TODO: find out what this is

        public bool HasSpecialEndDeath { get; private set; } // marks if a monster dies a special death
        public bool DeadCollision { get; private set; } // if true, corpse has collision   
        
        public bool CanBeRevivedByOtherMonsters { get; private set; }

        // appearance config
        public IEnemyTypeAppearanceConfig AppearanceConfig { get; private set; }

        // combat config
        public IEnemyTypeCombatConfig CombatConfig { get; private set; }

        // difficulty configs
        public IEnemyTypeDifficultyConfig NormalDifficultyConfig { get; private set; }
        public IEnemyTypeDifficultyConfig NightmareDifficultyConfig { get; private set; }
        public IEnemyTypeDifficultyConfig HellDifficultyConfig { get; private set; }

        public EnemyTypeConfig(string InternalName, string Name, string Type, string Descriptor,
            int BaseId, int PopulateId, bool Spawned, bool Beta, string Code, bool ClientOnly, bool NoMap,
            int SizeX, int SizeY, int Height, bool NoOverlays, int OverlayHeight,
            int Velocity, int RunVelocity,
            bool CanStealFrom, int ColdEffect,
            bool Rarity,
            int MinimumGrouping, int MaximumGrouping,
            string BaseWeapon,
            int[] AIParams,
            bool Allied, bool IsNPC, bool IsCritter, bool CanEnterTown,
            int HealthRegen,
            bool IsDemon, bool IsLarge, bool IsSmall, bool IsFlying, bool CanOpenDoors, int SpawningColumn,
            bool IsBoss, bool IsInteractable,
            bool IsKillable, bool CanBeConverted,
            int HitClass,
            bool HasSpecialEndDeath, bool DeadCollision,
            bool CanBeRevivedByOtherMonsters,
            IEnemyTypeAppearanceConfig AppearanceConfig,
            IEnemyTypeCombatConfig CombatConfig,
            IEnemyTypeDifficultyConfig NormalDifficultyConfig,
            IEnemyTypeDifficultyConfig NightmareDifficultyConfig,
            IEnemyTypeDifficultyConfig HellDifficultyConfig)
        {
            this.InternalName = InternalName;
            this.DisplayName = Name;
            this.Type = Type;
            this.Descriptor = Descriptor;

            this.BaseId = BaseId;
            this.PopulateId = PopulateId;
            this.Spawned = Spawned;
            this.Beta = Beta;
            this.Code = Code;
            this.ClientOnly = ClientOnly;
            this.NoMap = NoMap;

            this.SizeX = SizeX;
            this.SizeY = SizeY;
            this.Height = Height;
            this.NoOverlays = NoOverlays;
            this.OverlayHeight = OverlayHeight;

            this.Velocity = Velocity;
            this.RunVelocity = RunVelocity;

            this.CanStealFrom = CanStealFrom;
            this.ColdEffect = ColdEffect;

            this.Rarity = Rarity;

            this.MinimumGrouping = MinimumGrouping;
            this.MaximumGrouping = MaximumGrouping;

            this.BaseWeapon = BaseWeapon;

            this.AIParams = AIParams;

            this.Allied = Allied;
            this.IsNPC = IsNPC;
            this.IsCritter = IsCritter;
            this.CanEnterTown = CanEnterTown;

            this.HealthRegen = HealthRegen;

            this.IsDemon = IsDemon;
            this.IsLarge = IsLarge;
            this.IsSmall = IsSmall;
            this.IsFlying = IsFlying;
            this.CanOpenDoors = CanOpenDoors;
            this.SpawningColumn = SpawningColumn;
            this.IsBoss = IsBoss;
            this.IsInteractable = IsInteractable;

            this.IsKillable = IsKillable;
            this.CanBeConverted = CanBeConverted;

            this.HitClass = HitClass;

            this.HasSpecialEndDeath = HasSpecialEndDeath;
            this.DeadCollision = DeadCollision;

            this.CanBeRevivedByOtherMonsters = CanBeRevivedByOtherMonsters;

            this.AppearanceConfig = AppearanceConfig;

            this.CombatConfig = CombatConfig;

            this.NormalDifficultyConfig = NormalDifficultyConfig;
            this.NightmareDifficultyConfig = NightmareDifficultyConfig;
            this.HellDifficultyConfig = HellDifficultyConfig;
        }

        public IEnemyTypeDifficultyConfig GetDifficultyConfig(eDifficulty Difficulty)
        {
            switch (Difficulty)
            {
                case eDifficulty.HELL:
                    return HellDifficultyConfig;
                case eDifficulty.NIGHTMARE:
                    return NightmareDifficultyConfig;
                case eDifficulty.NORMAL:
                    return NormalDifficultyConfig;
                default:
                    return NormalDifficultyConfig;
            }
        }
    }

    public static class EnemyTypeConfigHelper
    {

        // (0) Class	namco	Type	Descriptor
        // (4) BaseId	PopulateId	Spawned	Beta	Code	ClientOnly	NoMap
        // (11) SizeX	SizeY	Height
        // (14) NoOverlays	OverlayHeight
        // (16) Velocity	Run
        // (18) CanStealFrom	ColdEffect
        // (20) Rarity	Level	Level(N)	Level(H)
        // (24) MeleeRng	MinGrp	MaxGrp
        // (27) HD	TR	LG	RA	LA	RH	LH	SH	S1	S2	S3	S4	S5	S6	S7	S8
        // (43) TotalPieces	SpawnComponents	BaseW
        // (46) AIParam1	Comment	AIParam2	Comment	AIParam3	Comment	AIParam4	Comment	AIParam5	Comment
        // (56) ModeDH	ModeN	ModeW	ModeGH	ModeA1	ModeA2	ModeB	ModeC
        // (64) ModeS1	ModeS2	ModeS3	ModeS4	ModeDD	ModeKB	ModeSQ	ModeRN
        // (72) ElMode	ElType	ElOver	ElPct	ElMinD	ElMaxD	ElDur
        // (79) MissA1	MissA2	MissS1	MissS2	MissS3	MissS4	MissC	MissSQ
        // (87) A1Move	A2Move	S1Move	S2Move	S3Move	S4Move	Cmove
        // (94) Align	
        // (95) IsMelee	IsSel	IsSel2	NeverSel	CorpseSel	IsAtt	IsNPC	IsCritter	InTown
        // (104) Bleed	Shadow	Light	NoUniqueShift	CompositeDeath
        // (109) Skill1	Skill1Seq	Skill1Lvl	Skill2	Skill2Seq	Skill2Lvl	Skill3	Skill3Seq	Skill3Lvl	Skill4	Skill4Seq	Skill4Lvl	Skill5	Skill5Seq	Skill5Lvl
        // (124) LightR	LightG	LightB
        // (127) DamageResist	MagicResist	FireResist	LightResist	ColdResist	PoisonResist
        // (133) DamageResist(N)	MagicResist(N)	FireResist(N)	LightResist(N)	ColdResist(N)	PoisonResist(N)
        // (139) DamageResist(H)	MagicResist(H)	FireResist(H)	LightResist(H)	ColdResist(H)	PoisonResist(H)
        // (145) DamageRegen
        // (146) eLUndead	eHUndead	eDemon	eMagicUsing	eLarge	eSmall	eFlying	eOpenDoors	eSpawnCol	eBoss
        // (156) PixHeight	Interact
        // (158) MinHP	MaxHP	AC	Exp	ToBlock
        // (163) A1MinD	A1MaxD	A1ToHit	A2MinD	A2MaxD	A2ToHit	S1MinD	S1MaxD	S1ToHit


        // (172) MinHP(N)	MaxHP(N)	AC(N)	Exp(N)	A1MinD(N)	A1MaxD(N)	A1ToHit(N)	A2MinD(N)	A2MaxD(N)	A2ToHit(N)	S1MinD(N)	S1MaxD(N)	S1ToHit(N)
        // (185) MinHP(H)	MaxHP(H)	AC(H)	Exp(H)	A1MinD(H)	A1MaxD(H)	A1ToHit(H)	A2MinD(H)	A2MaxD(H)	A2ToHit(H)	S1MinD(H)	S1MaxD(H)	S1ToHit(H)
        // (198) TreasureClass1	TreasureClass2	TreasureClass3	TreasureClass4
        // (202) TreasureClass1(N)	TreasureClass2(N)	TreasureClass3(N)	TreasureClass4(N)
        // (206) TreasureClass1(H)	TreasureClass2(H)	TreasureClass3(H)	TreasureClass4(H)
        // (210) SpawnPctBonus	Soft	Heart	BodyPart	Killable	Switch	Restore	NeverCount	HitClass
        // (219) SplEndDeath	SplGetModeChart	SplEndGeneric	SplClientEnd
        // (223) DeadCollision	UnflatDead	BloodLocal	DeathDamage
        // (227) PetIgnore	NoGfxHitTest
        // (229) HitTestTop	HitTestLeft	HitTestWidth	HitTestHeight
        // (233) GenericSpawn	AutomapCel	SparsePopulate	Zoo	ObjectCollision	Inert
        private static int IntConvert(string s)
        {
            // this is a convenience because sometimes they forget to put a 0 in the D2 monstats.txt
            if (string.IsNullOrEmpty(s))
            {
                return 0;
            }
            return Convert.ToInt32(s);
        }

        public static IEnemyTypeConfig ToEnemyTypeConfig(this string[] row)
        {
            return new EnemyTypeConfig(
                Name: row[0],
                UniqueName: row[1],
                Type: row[2],
                Descriptor: row[3],

                BaseId: IntConvert(row[4]),
                PopulateId: IntConvert(row[5]),
                Spawned: (row[6] == "1"),
                Beta: (row[7] == "1"),
                Code: row[8],
                ClientOnly: (row[9] == "1"),
                NoMap: (row[10] == "1"),

                SizeX: IntConvert(row[11]),
                SizeY: IntConvert(row[12]),
                Height: IntConvert(row[13]),
                NoOverlays: (row[14] == "1"),
                OverlayHeight: IntConvert(row[15]),

                Velocity: IntConvert(row[16]),
                RunVelocity: IntConvert(row[17]),

                CanStealFrom: (row[18] == "1"),
                ColdEffect: IntConvert(row[19]),

                Rarity: (row[20] == "1"),

                
                MinimumGrouping: IntConvert(row[25]),
                MaximumGrouping: IntConvert(row[26]),

                BaseWeapon: row[45],

                AIParams: new int[]
                {
                    IntConvert(row[46]),
                    IntConvert(row[48]),
                    IntConvert(row[50]),
                    IntConvert(row[52]),
                    IntConvert(row[54])
                },

                Allied: (row[94] == "1"),
                IsNPC: (row[101] == "1"),
                IsCritter: (row[102] == "1"),
                CanEnterTown: (row[103] == "1"),

                HealthRegen: IntConvert(row[145]),

                // (146) eLUndead	eHUndead	eDemon	eMagicUsing	eLarge	eSmall	eFlying	eOpenDoors	eSpawnCol	eBoss
                IsDemon: (row[148] == "1"),
                IsLarge: (row[150] == "1"),
                IsSmall: (row[151] == "1"),
                IsFlying: (row[152] == "1"),
                CanOpenDoors: (row[153] == "1"),
                SpawningColumn: IntConvert(row[154]),
                IsBoss: (row[155] == "1"),
                IsInteractable: (row[157] == "1"),

                IsKillable: (row[215] == "1"),
                CanBeConverted: (row[216] == "1"),

                HitClass: IntConvert(row[218]),

                HasSpecialEndDeath: (row[219] == "1"),
                DeadCollision: (row[223] == "1"),

                CanBeRevivedByOtherMonsters: (row[236] == "1"),

                // (56) ModeDH	ModeN	ModeW	ModeGH	ModeA1	ModeA2	ModeB	ModeC
                // (64) ModeS1	ModeS2	ModeS3	ModeS4	ModeDD	ModeKB	ModeSQ	ModeRN
                // (95) IsMelee	IsSel	IsSel2	NeverSel	CorpseSel	IsAtt	IsNPC	IsCritter	InTown
                // (104) Bleed	Shadow	Light	NoUniqueShift	CompositeDeath
                // (124) LightR	LightG	LightB
                AppearanceConfig: new EnemyTypeAppearanceConfig
                (
                    HasDeathAnimation: (row[56] == "1"),
                    HasNeutralAnimation: (row[57] == "1"),
                    HasWalkAnimation: (row[58] == "1"),
                    HasGetHitAnimation: (row[59] == "1"),
                    HasAttack1Animation: (row[60] == "1"),
                    HasAttack2Animation: (row[61] == "1"),
                    HasBlockAnimation: (row[62] == "1"),
                    HasCastAnimation: (row[63] == "1"),
                    HasSkillAnimation: new bool[]
                    {
                        (row[64] == "1"),
                        (row[65] == "1"),
                        (row[66] == "1"),
                        (row[67] == "1")
                    },
                    HasCorpseAnimation: (row[68] == "1"),
                    HasKnockbackAnimation: (row[69] == "1"),
                    HasRunAnimation: (row[71] == "1"),

                    HasLifeBar: (row[96] == "1"),
                    HasNameBar: (row[97] == "1"),
                    CannotBeSelected: (row[98] == "1"),
                    CanCorpseBeSelected: (row[99] == "1"),

                    BleedType: IntConvert(row[104]),
                    HasShadow: (row[105] == "1"),
                    LightRadius: IntConvert(row[106]),
                    HasUniqueBossColors: (row[107] == "0"), // note: inverting the bool here
                    // since in the table it is "NoUniqueShift" but I want to store it as "HasUniqueShift"
                    CompositeDeath: (row[108] == "1"),

                    LightRGB: new byte[]
                    {
                        Convert.ToByte(row[124]),
                        Convert.ToByte(row[125]),
                        Convert.ToByte(row[126])
                    }
                ),

                // (72) ElMode	ElType	ElOver	ElPct	ElMinD	ElMaxD	ElDur
                // (79) MissA1	MissA2	MissS1	MissS2	MissS3	MissS4	MissC	MissSQ
                // (87) A1Move	A2Move	S1Move	S2Move	S3Move	S4Move	Cmove
                // (95) IsMelee	IsSel	IsSel2	NeverSel	CorpseSel	IsAtt	IsNPC	IsCritter	InTown
                // (109) Skill1	Skill1Seq	Skill1Lvl	Skill2	Skill2Seq	Skill2Lvl	Skill3	Skill3Seq	Skill3Lvl	Skill4	Skill4Seq	Skill4Lvl	Skill5	Skill5Seq	Skill5Lvl
                CombatConfig: new EnemyTypeCombatConfig
                (
                    ElementalAttackMode: IntConvert(row[72]),
                    ElementalAttackType: (eDamageTypes)IntConvert(row[73]),
                    ElementalOverlayId: IntConvert(row[74]),
                    ElementalChance: IntConvert(row[75]),
                    ElementalMinDamage: IntConvert(row[76]),
                    ElementalMaxDamage: IntConvert(row[77]),
                    ElementalDuration: IntConvert(row[78]),

                    MissileForAttack: new int[]
                    {
                        IntConvert(row[79]),
                        IntConvert(row[80])
                    },
                    MissileForSkill: new int[]
                    {
                        IntConvert(row[81]),
                        IntConvert(row[82]),
                        IntConvert(row[83]),
                        IntConvert(row[84])
                    },
                    MissileForCase: IntConvert(row[85]),
                    MissileForSequence: IntConvert(row[86]),

                    CanMoveAttack: new bool[]
                    {
                        (row[87] == "1"),
                        (row[88] == "1")
                    },
                    CanMoveSkill: new bool[]
                    {
                        (row[89] == "1"),
                        (row[90] == "1"),
                        (row[91] == "1"),
                        (row[92] == "1")
                    },
                    CanMoveCast: (row[93] == "1"),

                    IsMelee: (row[95] == "1"),
                    IsAttackable: (row[100] == "1"),

                    MeleeRange: IntConvert(row[24]),

                    SkillType: new int[]
                    {
                        IntConvert(row[109]),
                        IntConvert(row[112]),
                        IntConvert(row[115]),
                        IntConvert(row[118]),
                        IntConvert(row[121]),
                    },
                    SkillSequence: new int[]
                    {
                        IntConvert(row[110]),
                        IntConvert(row[113]),
                        IntConvert(row[116]),
                        IntConvert(row[119]),
                        IntConvert(row[122]),
                    },
                    SkillLevel: new int[]
                    {
                        IntConvert(row[111]),
                        IntConvert(row[114]),
                        IntConvert(row[117]),
                        IntConvert(row[120]),
                        IntConvert(row[123]),
                    },

                    IsUndeadWithPhysicalAttacks: (row[146] == "1"),
                    IsUndeadWithMagicAttacks: (row[147] == "1"),
                    UsesMagicAttacks: (row[149] == "1"),

                    ChanceToBlock: IntConvert(row[162]),

                    DoesDeathDamage: (row[226] == "1"),

                    IgnoredBySummons: (row[227] == "1")
                ),

                // (127) DamageResist	MagicResist	FireResist	LightResist	ColdResist	PoisonResist
                // (133) DamageResist(N)	MagicResist(N)	FireResist(N)	LightResist(N)	ColdResist(N)	PoisonResist(N)
                // (139) DamageResist(H)	MagicResist(H)	FireResist(H)	LightResist(H)	ColdResist(H)	PoisonResist(H)
                // (158) MinHP	MaxHP	AC	Exp	ToBlock
                // (163) A1MinD	A1MaxD	A1ToHit	A2MinD	A2MaxD	A2ToHit	S1MinD	S1MaxD	S1ToHit
                // (172) MinHP(N)	MaxHP(N)	AC(N)	Exp(N)	A1MinD(N)	A1MaxD(N)	A1ToHit(N)	A2MinD(N)	A2MaxD(N)	A2ToHit(N)	S1MinD(N)	S1MaxD(N)	S1ToHit(N)
                // (185) MinHP(H)	MaxHP(H)	AC(H)	Exp(H)	A1MinD(H)	A1MaxD(H)	A1ToHit(H)	A2MinD(H)	A2MaxD(H)	A2ToHit(H)	S1MinD(H)	S1MaxD(H)	S1ToHit(H)
                // (192) TreasureClass1	TreasureClass2	TreasureClass3	TreasureClass4
                // (202) TreasureClass1(N)	TreasureClass2(N)	TreasureClass3(N)	TreasureClass4(N)
                // (206) TreasureClass1(H)	TreasureClass2(H)	TreasureClass3(H)	TreasureClass4(H)
                NormalDifficultyConfig: new EnemyTypeDifficultyConfig
                (
                    Level: IntConvert(row[21]),

                    DamageResist: (IntConvert(row[127]) / 100.0),
                    MagicResist: (IntConvert(row[128]) / 100.0),
                    FireResist: (IntConvert(row[129]) / 100.0),
                    LightResist: (IntConvert(row[130]) / 100.0),
                    ColdResist: (IntConvert(row[131]) / 100.0),
                    PoisonResist: (IntConvert(row[132]) / 100.0),

                    MinHP: IntConvert(row[158]),
                    MaxHP: IntConvert(row[159]),
                    AC: IntConvert(row[160]),
                    Exp: IntConvert(row[161]),

                    AttackMinDamage: new int[]
                    {
                        IntConvert(row[163]),
                        IntConvert(row[166])
                    },
                    AttackMaxDamage: new int[]
                    {
                        IntConvert(row[164]),
                        IntConvert(row[167])
                    },
                    AttackChanceToHit: new int[]
                    {
                        IntConvert(row[165]),
                        IntConvert(row[168])
                    },
                    Skill1MinDamage: IntConvert(row[169]),
                    Skill1MaxDamage: IntConvert(row[170]),
                    Skill1ChanceToHit: IntConvert(row[171]),

                    TreasureClass: new string[]
                    {
                        row[198], row[199], row[200], row[201]
                    }
                ),
                NightmareDifficultyConfig: new EnemyTypeDifficultyConfig
                (
                    Level: IntConvert(row[22]),

                    DamageResist: (IntConvert(row[133]) / 100.0),
                    MagicResist: (IntConvert(row[134]) / 100.0),
                    FireResist: (IntConvert(row[135]) / 100.0),
                    LightResist: (IntConvert(row[136]) / 100.0),
                    ColdResist: (IntConvert(row[137]) / 100.0),
                    PoisonResist: (IntConvert(row[138]) / 100.0),

                    MinHP: IntConvert(row[172]),
                    MaxHP: IntConvert(row[173]),
                    AC: IntConvert(row[174]),
                    Exp: IntConvert(row[175]),

                    AttackMinDamage: new int[]
                    {
                        IntConvert(row[176]),
                        IntConvert(row[179])
                    },
                    AttackMaxDamage: new int[]
                    {
                        IntConvert(row[177]),
                        IntConvert(row[180])
                    },
                    AttackChanceToHit: new int[]
                    {
                        IntConvert(row[178]),
                        IntConvert(row[181])
                    },
                    Skill1MinDamage: IntConvert(row[182]),
                    Skill1MaxDamage: IntConvert(row[183]),
                    Skill1ChanceToHit: IntConvert(row[184]),

                    TreasureClass: new string[]
                    {
                        row[202], row[203], row[204], row[205]
                    }
                ),
                HellDifficultyConfig: new EnemyTypeDifficultyConfig
                (
                    Level: IntConvert(row[23]),

                    DamageResist: (IntConvert(row[139]) / 100.0),
                    MagicResist: (IntConvert(row[140]) / 100.0),
                    FireResist: (IntConvert(row[141]) / 100.0),
                    LightResist: (IntConvert(row[142]) / 100.0),
                    ColdResist: (IntConvert(row[143]) / 100.0),
                    PoisonResist: (IntConvert(row[144]) / 100.0),

                    MinHP: IntConvert(row[185]),
                    MaxHP: IntConvert(row[186]),
                    AC: IntConvert(row[187]),
                    Exp: IntConvert(row[188]),

                    AttackMinDamage: new int[]
                    {
                        IntConvert(row[189]),
                        IntConvert(row[192])
                    },
                    AttackMaxDamage: new int[]
                    {
                        IntConvert(row[190]),
                        IntConvert(row[193])
                    },
                    AttackChanceToHit: new int[]
                    {
                        IntConvert(row[191]),
                        IntConvert(row[194])
                    },
                    Skill1MinDamage: IntConvert(row[195]),
                    Skill1MaxDamage: IntConvert(row[196]),
                    Skill1ChanceToHit: IntConvert(row[197]),

                    TreasureClass: new string[]
                    {
                        row[206], row[207], row[208], row[209]
                    }
                )
                );
        }
    }
}
