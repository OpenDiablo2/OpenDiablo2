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


        public int InternalId { get; private set; }
        public string BaseName { get; private set; } // name of the monster it is a subtype of
        public string NextInClass { get; private set; } // name of the next monster in the subtype chain
        public bool Enabled { get; private set; } // if false, monster cannot be used in any way

        public bool IsRanged { get; private set; }

        public bool IsSpawner { get; private set; } // tells the game this is a nest / spawner that makes other monsters
        public string SpawnName { get; private set; } // name of what this nest spawns
        public int SpawnOffsetX { get; private set; } // offset of spawned monsters
        public int SpawnOffsetY { get; private set; } // offset of spawned monsters
        public string SpawnAnimationMode { get; private set; } // what animation mode the monsters spawn in

        public string[] MinionNames { get; private set; } // 1-2; when a monster is spawned, it will
        // create minions with these names around it
        public bool SpawnsBoss { get; private set; } // is the unit it spawns a boss?
        public bool BossTransfers { get; private set; } // if true, boss status is passed to one of this
        // monster's minions when it is killed
        public int MinionMin { get; private set; } // how many minions to spawn
        public int MinionMax { get; private set; }

        public int MinimumGrouping { get; private set; } // how many of this monster to spawn
        public int MaximumGrouping { get; private set; }
        public int SpawnChance { get; private set; } // when game chooses to spawn this monster, it will
        // have this chance of spawning something else instead 
        // NOTE SPECIAL CASE: if spawnchance is 0, it should be considered as 100%

        public int Velocity { get; private set; }
        public int RunVelocity { get; private set; }

        public int Rarity { get; private set; }

        public string Sound { get; private set; } // name of this monster's sound bank
        public string UniqueSound { get; private set; } // name of the sound bank if unique

        public bool Allied { get; private set; } // Is this enemy on the player's side?
        public bool Neutral { get; private set; } // If this is set, not allied nor a monster
        public bool Spawned { get; private set; }
        public bool IsNPC { get; private set; } // is this an NPC?
        public bool IsInteractable { get; private set; } // can this be interacted with? like an NPC

        public bool HasInventory { get; private set; } // does this have an inventory?
        // needed if monster has equipment data in monequip.txt or if this is an npc whose inventory
        // you can access
        public bool CanEnterTown { get; private set; } // can this enter town or not?
        public bool LowUndead { get; private set; } // low undead can be rez'd by high undead
        public bool HighUndead { get; private set; } // cannot be rez'd
        public bool IsDemon { get; private set; } // demon?
        public bool IsFlying { get; private set; } // can move over water for instance
        public bool CanOpenDoors { get; private set; }
        public bool IsBoss { get; private set; } // Bosses ignore the level settings and always spawn with
        // their monstats level. Bosses are also immune to being stunned. (TODO: research full extent of boss powers)

        public bool IsPrimeEvil { get; private set; } // Primevils have a 300% bonus vs hirelings and summons
        public bool IsKillable { get; private set; }
        public bool CanBeConverted { get; private set; } // if true, can be switched to Allied by spells like Conversion


        // OLD 

        public int PopulateId { get; private set; }
        
        public bool Beta { get; private set; }
        public bool ClientOnly { get; private set; }
        public bool NoMap { get; private set; }

        public int SizeX { get; private set; }
        public int SizeY { get; private set; }
        public int Height { get; private set; }
        public bool NoOverlays { get; private set; }
        public int OverlayHeight { get; private set; }
        
        public bool CanStealFrom { get; private set; }
        public int ColdEffect { get; private set; }
        
        public string BaseWeapon { get; private set; }

        
        
        public bool IsCritter { get; private set; } // is this a critter? (e.g. the chickens)
        

        public int HealthRegen { get; private set; } // hp regen per minute

        
        public bool IsLarge { get; private set; } // size large (e.g. bosses)
        public bool IsSmall { get; private set; } // size small (e.g. fallen)
        
        
        public int SpawningColumn { get; private set; } // is the monster area restricted
        // 0 = no, spawns through levels.txt, 1-3 unknown? 
        // TODO: understand spawningcolumn
        
        
        
        
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

        public EnemyTypeConfig(string InternalName, int InternalId, string BaseId, string NextInClass,
            string Name, string Type, string Descriptor,
            bool IsRanged, bool IsSpawner, string SpawnName, int SpawnOffsetX, int SpawnOffsetY,
            string SpawnAnimationMode, string[] MinionNames, bool SpawnsBoss, bool BossTransfers,
            int MinionMin, int MinionMax, int MinimumGrouping, int MaximumGrouping,
            int SpawnChance, int Velocity, int RunVelocity, int Rarity,
            string Sound, string UniqueSound,
            bool Allied, bool Neutral, bool Spawned, bool IsNPC, bool IsInteractable,
            bool HasInventory, bool CanEnterTown, bool LowUndead, bool HighUndead,
            bool IsDemon, bool IsFlying, bool CanOpenDoors, bool IsBoss,
            bool IsPrimeEvil, bool IsKillable, bool CanBeConverted,


            int PopulateId, bool Beta, bool ClientOnly, bool NoMap,
            int SizeX, int SizeY, int Height, bool NoOverlays, int OverlayHeight,
            bool CanStealFrom, int ColdEffect,
            string BaseWeapon,
            bool IsCritter, 
            int HealthRegen,
            bool IsLarge, bool IsSmall,  int SpawningColumn,
            
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
            this.InternalId = InternalId;
            this.DisplayName = Name;
            this.Type = Type;
            this.Descriptor = Descriptor;

            this.BaseName = BaseId;
            this.NextInClass = NextInClass;

            this.IsRanged = IsRanged;
            this.IsSpawner = IsSpawner;
            this.SpawnName = SpawnName;
            this.SpawnOffsetX = SpawnOffsetX;
            this.SpawnOffsetY = SpawnOffsetY;
            this.SpawnAnimationMode = SpawnAnimationMode;
            this.MinionNames = MinionNames;
            this.SpawnsBoss = SpawnsBoss;
            this.BossTransfers = BossTransfers;
            this.MinionMin = MinionMin;
            this.MinionMax = MinionMax;
            this.MinimumGrouping = MinimumGrouping;
            this.MaximumGrouping = MaximumGrouping;
            this.SpawnChance = SpawnChance;
            if(SpawnChance == 0)
            {
                this.SpawnChance = 100; // override behavior
            }
            this.Velocity = Velocity;
            this.RunVelocity = RunVelocity;
            this.Rarity = Rarity;

            this.Sound = Sound;
            this.UniqueSound = UniqueSound;

            this.Allied = Allied;
            this.Neutral = Neutral;
            this.Spawned = Spawned;
            this.IsNPC = IsNPC;
            this.IsInteractable = IsInteractable;

            this.HasInventory = HasInventory;
            this.CanEnterTown = CanEnterTown;
            this.LowUndead = LowUndead;
            this.HighUndead = HighUndead;
            this.IsDemon = IsDemon;
            this.IsFlying = IsFlying;
            this.CanOpenDoors = CanOpenDoors;
            this.IsBoss = IsBoss;
            this.IsPrimeEvil = IsPrimeEvil;
            this.IsKillable = IsKillable;
            this.CanBeConverted = CanBeConverted;

            // vvv
            this.PopulateId = PopulateId;
            this.Beta = Beta;
            this.ClientOnly = ClientOnly;
            this.NoMap = NoMap;

            this.SizeX = SizeX;
            this.SizeY = SizeY;
            this.Height = Height;
            this.NoOverlays = NoOverlays;
            this.OverlayHeight = OverlayHeight;

            this.CanStealFrom = CanStealFrom;
            this.ColdEffect = ColdEffect;

            this.BaseWeapon = BaseWeapon;
            
            
            this.IsCritter = IsCritter;
            this.CanEnterTown = CanEnterTown;

            this.HealthRegen = HealthRegen;
            
            this.IsLarge = IsLarge;
            this.IsSmall = IsSmall;
            this.SpawningColumn = SpawningColumn;
            

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
        
        private static int IntConvert(string s)
        {
            // this is a convenience because sometimes they forget to put a 0 in the D2 monstats.txt
            if (string.IsNullOrEmpty(s))
            {
                return 0;
            }
            return Convert.ToInt32(s);
        }

        public static IEnemyTypeConfig MakeEnemyTypeConfig(string[] row, string[] row2, string[] rowsprop)
        {
            // important to note: rows2 and rowsprop may be empty
            return new EnemyTypeConfig(
                InternalName: row[0],
                InternalId: IntConvert(row[1]),
                BaseId: row[2],
                NextInClass: row[3],
                Name: row[5],
                Type: row[9],
                Descriptor: row[10],

                IsRanged: (row[13] == "1"),
                IsSpawner: (row[14] == "1"),
                SpawnName: row[15],
                SpawnOffsetX: IntConvert(row[16]),
                SpawnOffsetY: IntConvert(row[17]),
                SpawnAnimationMode: row[18],
                MinionNames: new string[] { row[19], row[20] },
                SpawnsBoss: (row[21] == "1"),
                BossTransfers: (row[22] == "1"),
                MinionMin: IntConvert(row[23]),
                MinionMax: IntConvert(row[24]),
                MinimumGrouping: IntConvert(row[25]),
                MaximumGrouping: IntConvert(row[26]),
                SpawnChance: IntConvert(row[27]),
                Velocity: IntConvert(row[28]),
                RunVelocity: IntConvert(row[29]),
                Rarity: IntConvert(row[30]),

                Sound: row[34],
                UniqueSound: row[35],

                Allied: (row[75] == "1"),
                Neutral: (row[75] == "2"),
                IsNPC: (row[78] == "1"),
                IsInteractable: (row[79] == "1"),

                HasInventory: (row[80] == "1"),
                CanEnterTown: (row[81] == "1"),
                LowUndead: (row[82] == "1"),
                HighUndead: (row[83] == "1"),
                IsDemon: (row[84] == "1"),
                IsFlying: (row[85] == "1"),
                CanOpenDoors: (row[86] == "1"),
                IsBoss: (row[87] == "1"),
                IsPrimeEvil: (row[88] == "1"),
                IsKillable: (row[89] == "1"),
                CanBeConverted: (row[90] == "1"),

                CombatConfig: new EnemyTypeCombatConfig(
                    Threat: IntConvert(row[36]),
                    MissileForAttack: new string[]
                    {
                        row[67],
                        row[68]
                    },
                    MissileForSkill: new string[]
                    {
                        row[69],
                        row[70],
                        row[71],
                        row[72]
                    },
                    MissileForCast: row[73],
                    MissileForSequence: row[74],
                    IsMelee: (row[77] == "1"),
                    ),

                AppearanceConfig: new EnemyTypeAppearanceConfig(
                    PalleteNo: IntConvert(row[4]),
                    Code: row[11],
                    ),
                
                NormalDifficultyConfig: new EnemyTypeDifficultyConfig(
                    Level: IntConvert(row[31]),
                    AiDelay: IntConvert(row[37]),
                    AiActivationDistance: IntConvert(row[40]),
                    AiParams: new int[]
                    {
                        IntConvert(row[43]),
                        IntConvert(row[46]),
                        IntConvert(row[49]),
                        IntConvert(row[52]),
                        IntConvert(row[55]),
                        IntConvert(row[58]),
                        IntConvert(row[61]),
                        IntConvert(row[64])
                    },
                    ),
                NightmareDifficultyConfig: new EnemyTypeDifficultyConfig(
                    Level: IntConvert(row[32]),
                    AiDelay: IntConvert(row[38]),
                    AiActivationDistance: IntConvert(row[41]),
                    AiParams: new int[]
                    {
                        IntConvert(row[44]),
                        IntConvert(row[47]),
                        IntConvert(row[50]),
                        IntConvert(row[53]),
                        IntConvert(row[56]),
                        IntConvert(row[59]),
                        IntConvert(row[62]),
                        IntConvert(row[65])
                    },
                    ),
                HellDifficultyConfig: new EnemyTypeDifficultyConfig(
                    Level: IntConvert(row[33]),
                    AiDelay: IntConvert(row[39]),
                    AiActivationDistance: IntConvert(row[42]),
                    AiParams: new int[]
                    {
                        IntConvert(row[45]),
                        IntConvert(row[48]),
                        IntConvert(row[51]),
                        IntConvert(row[54]),
                        IntConvert(row[57]),
                        IntConvert(row[60]),
                        IntConvert(row[63]),
                        IntConvert(row[66])
                    },
                    )
                );
        }


        public static IEnemyTypeConfig ToEnemyTypeConfig(this string[] row)
        {
            return new EnemyTypeConfig(
                InternalName: row[0],
                Name: row[1],
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
