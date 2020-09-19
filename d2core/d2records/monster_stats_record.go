package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// https://d2mods.info/forum/kb/viewarticle?a=360

// MonStats stores all of the MonStat Records
type MonStats map[string]*MonStatsRecord

type (
	// MonStatsRecord represents a single row from `data/global/excel/monstats.txt` in the MPQ files.
	// These records are used for creating monsters.
	MonStatsRecord struct {

		// Key contains the pointer that will be used in other txt files
		// such as levels.txt and superuniques.txt.
		Key string // called `Id` in monstats.txt

		// Id is the actual internal ID of the unit (this is what the ID pointer
		// actually points at) remember that no two units can have the same ID,
		// this will result in lots of unpredictable behavior and crashes so please
		// don’t do it. This 'HarcCodedInDeX' is used for several things, such as
		// determining whenever the unit uses DCC or DC6 graphics (like mephisto
		// and the death animations of Diablo, the Maggoc Queen etc.), the hcIdx
		// column also links other hardcoded effects to the units, such as the
		// transparency on necro summons and the name-color change on unique boss
		// units (thanks to Kingpin for the info)
		ID int // called `hcIdx` in monstats.txt

		// BaseKey is an ID pointer of the “base” unit for this specific
		// monster type (ex. There are five types of “Fallen”; all of them have
		// fallen1 as their “base” unit).
		BaseKey string // called `BaseId` in monstats.txt

		// NextKey is the ID of the next unit in the chain. (fallen1 has the ID pointer of fallen2 in here).
		// The game uses this for “map generated” monsters such as the fallen in the fallen camps,
		// which get picked based on area level.
		NextKey string // called `NextInClass` in monstats.txt

		// NameStringTableKey the string-key used in the TBL (string.tbl,
		// expansionstring.tbl and patchstring.tbl) files to make this monsters
		// name appear when you highlight it.
		NameString string // called `NameStr` in monstats.txt

		// ExtraDataKey the ID pointer to an entry in MonStats2.txt.
		ExtraDataKey string // called `MonStatsEx` in monstats.txt

		// PropertiesKey contains the ID pointer to an entry in MonProp.txt which
		// controls what special modifiers are appended to the unit
		PropertiesKey string // called `MonProp` in monstats.txt

		// MonsterGroup contains the group ID of the “super group” this monster
		// belongs to, IE all skeletons belong to the "super group" skeleton. The
		MonsterGroup string // called `MonType` in monstats.txt

		// AiKey tells the game which AI to use for this monster. Every AI
		// needs a specific set of animation modes (GH, A1, A2, S1, WL, RN etc).
		AiKey string // called `AI` in monstats.txt

		// DescriptionStringTableKey contains the string-key used in the TBL (string.tbl,
		// expansionstring.tbl and patchstring.tbl) files for the monsters
		// description (leave it blank for no description).
		// NOTE: ever wondered how to make it say something below the monster
		// name (such as “Drains Mana and Stamina etc), well this is how you do it.
		// Just put the string-key of the string you want to display below the
		// monsters name in here.
		DescriptionStringTableKey string // called `DescStr` in monstats.txt

		// AnimationDirectoryToken controls which token (IE name of a folder that
		// contains animations) the game uses for this monster.
		AnimationDirectoryToken string // called `Code` in monstats.txt

		// SpawnKey contains the key of the unit to spawn.
		SpawnKey string // called `spawn` in monstats.txt

		// SpawnAnimationKey
		// which animation mode will the spawned monster be spawned in.
		SpawnAnimationKey string // called `spawnmode` in monstats.txt

		// MinionId1 is an Id of a minion that spawns when this monster is created
		MinionId1 string //nolint:golint,stylecheck // called `minion1` in monstats.txt

		// MinionId2 is an Id of a minion that spawns when this monster is created
		MinionId2 string //nolint:golint,stylecheck // called `minion2` in monstats.txt

		// SoundKeyNormal, SoundKeySpecial
		// specifies the ID pointer to this monsters “Sound Bank” in MonSound.txt
		// when this monster is normal.
		SoundKeyNormal  string // called `MonSound` in monstats.txt
		SoundKeySpecial string // called `UMonSound` in monstats.txt

		// MissileA1 -- MissileSQ
		// these columns control “non-skill-related” missiles used by the monster.
		// For example if you enter a missile ID pointer (from Missiles.txt) in
		// MissA1 then, whenever the monster uses its A1 mode, it will shoot a
		// missile, this however will successfully prevent it from dealing any damage
		// with the swing of A1.
		// NOTE: for the beginners, A1=Attack1, A2=Attack2, S1=Skill1, S2=Skill2,
		// S3=Skill3, S4=Skill4, C=Cast, SQ=Sequence.
		MissileA1 string // called `MissA1` in monstats.txt
		MissileA2 string // called `MissA2` in monstats.txt
		MissileS1 string // called `MissS1` in monstats.txt
		MissileS2 string // called `MissS2` in monstats.txt
		MissileS3 string // called `MissS3` in monstats.txt
		MissileS4 string // called `MissS4` in monstats.txt
		MissileC  string // called `MissC` in monstats.txt
		MissileSQ string // called `MissSQ` in monstats.txt

		// SkillId1 -- SkillId8
		// the ID Pointer to the skill (from Skills.txt) the monster will cast when
		// this specific slot is accessed by the AI. Which slots are used is
		// determined by the units AI.
		SkillId1 string //nolint:golint,stylecheck // called `Skill1` in monstats.txt
		SkillId2 string //nolint:golint,stylecheck // called `Skill2` in monstats.txt
		SkillId3 string //nolint:golint,stylecheck // called `Skill3` in monstats.txt
		SkillId4 string //nolint:golint,stylecheck // called `Skill4` in monstats.txt
		SkillId5 string //nolint:golint,stylecheck // called `Skill5` in monstats.txt
		SkillId6 string //nolint:golint,stylecheck // called `Skill6` in monstats.txt
		SkillId7 string //nolint:golint,stylecheck // called `Skill7` in monstats.txt
		SkillId8 string //nolint:golint,stylecheck // called `Skill8` in monstats.txt

		// SkillAnimation1 -- SkillAnimation8
		// the graphical MODE (or SEQUENCE) this unit uses when it uses this skill.
		SkillAnimation1 string // called `Sk1mode` in monstats.txt
		SkillAnimation2 string // called `Sk2mode` in monstats.txt
		SkillAnimation3 string // called `Sk3mode` in monstats.txt
		SkillAnimation4 string // called `Sk4mode` in monstats.txt
		SkillAnimation5 string // called `Sk5mode` in monstats.txt
		SkillAnimation6 string // called `Sk6mode` in monstats.txt
		SkillAnimation7 string // called `Sk7mode` in monstats.txt
		SkillAnimation8 string // called `Sk8mode` in monstats.txt

		// DamageSkillId
		// ID Pointer to the skill that controls this units damage. This is used for
		// the druids summons. IE their damage is specified solely by Skills.txt and
		// not by MonStats.txt.
		DamageSkillId string //nolint:golint,stylecheck // called `SkillDamage` in monstats.txt

		// ElementAttackMode1 -- ElementAttackMode3
		// the mode to which the elemental damage is appended. The modes to which
		// you would usually attack elemental damage are A1, A2, S1, S2, S3, S4, SQ
		// or C as these are the only ones that naturally contain trigger bytes.
		ElementAttackMode1 string // called `El1Mode` in monstats.txt
		ElementAttackMode2 string // called `El2Mode` in monstats.txt
		ElementAttackMode3 string // called `El3Mode` in monstats.txt

		// ElementType1 -- ElementType3
		// the type of the elemental damage appended to an attack. There are several
		// elements: fire=Fire Damage, ltng=Lightning Damage, cold=Cold Damage
		// (uses duration), pois = Poison Damage (uses duration), mag=Magic Damage,
		// life=Life Drain (the monster heals the specified amount when it hits
		// you), mana=Mana Drain (the monster steals the specified amount of mana
		// when it hits you), stam=Stamina Drain (the monster steals the specified
		// amount of stamina when it hits you), stun=Stun Damage (uses duration,
		// damage is not used, this only effects pets and mercs, players will not
		// get immobilized but they will get thrown into hit recovery whenever they
		// get hit by an attack, no matter what type of attack it is, thanks to
		// Brother Laz clearing this one up), rand=Random Damage (uses duration,
		// either does Poison, Cold, Fire or Lightning damage, randomly picked for
		// every attack), burn=Burning Damage (uses duration, this damage type
		// cannot be resisted or reduced in any way), frze=Freezing Damage (uses
		// duration, this will effect players like normal cold damage but will
		// freeze and shatter pets). If you want to give your monster knockback use
		// MonProp.txt.
		ElementType1 string // called `El1Type` in monstats.txt
		ElementType2 string // called `El2Type` in monstats.txt
		ElementType3 string // called `El3Type` in monstats.txt

		// TreasureClassNormal
		// Treasure class for normal monsters, champions, uniques, and quests
		// on the respective difficulties.
		TreasureClassNormal            string // called `TreasureClass1` in monstats.txt
		TreasureClassNightmare         string // called `TreasureClass1(N)` in monstats.txt
		TreasureClassHell              string // called `TreasureClass1(H)` in monstats.txt
		TreasureClassChampionNormal    string // called `TreasureClass2` in monstats.txt
		TreasureClassChampionNightmare string // called `TreasureClass2(N)` in monstats.txt
		TreasureClassChampionHell      string // called `TreasureClass2(H)` in monstats.txt
		TreasureClass3UniqueNormal     string // called `TreasureClass3` in monstats.txt
		TreasureClass3UniqueNightmare  string // called `TreasureClass3(N)` in monstats.txt
		TreasureClass3UniqueHell       string // called `TreasureClass3(H)` in monstats.txt
		TreasureClassQuestNormal       string // called `TreasureClass4` in monstats.txt
		TreasureClassQuestNightmare    string // called `TreasureClass4(N)` in monstats.txt
		TreasureClassQuestHell         string // called `TreasureClass4(H)` in monstats.txt

		// TreasureClassQuestTriggerId
		// the ID of the Quest that triggers the Quest Treasureclass drop.
		TreasureClassQuestTriggerId string //nolint:golint,stylecheck // called `TCQuestId` in monstats.txt

		// TreasureClassQuestCompleteId
		//  the ID of the Quest State that you need to complete to trigger the Quest
		// Treasureclass trop.
		TreasureClassQuestCompleteId string //nolint:golint,stylecheck // called `TCQuestCP` in monstats.txt

		// PaletteId indicates which palette (color) entry the unit will use, most
		// monsters have a palshift.dat file in their COF folder, this file
		// contains 8 palettes, starting from index 0. These palettes are used by
		// the game to make the various monster sub-types appear with color
		// variations. The game with use the palette from the palettes file
		// corresponding to the value in this column plus 2; eg: translvl = 0 will
		// use the third palette in the file.
		// NOTE: some tokens (token = IE name of a folder that contains animations)
		// such as FC do not accept their palettes.
		// NOTE no 2: some monsters got unused palettes, ZM (zombie) for example
		// will turn light-rotten-green with palette nr 5 and pink-creamy with 6.
		PaletteId int //nolint:golint,stylecheck // called `TransLvl` in monstats.txt

		// SpawnOffsetX, SpawnOffsetY
		// are the x/y offsets at which spawned monsters are placed. IE this prevents
		// the spawned monsters from being created at the same x/y coordinates as
		// the spawner itself.
		SpawnOffsetX int // called `spawnx` in monstats.txt
		SpawnOffsetY int // called `spawny` in monstats.txt

		// MinionPartyMin, MinionPartyMax controls how many minions are spawned together with this unit.
		MinionPartyMin int // called `PartyMin` in monstats.txt
		MinionPartyMax int // called `PartyMax` in monstats.txt

		// MinionGroupMin, MinionGroupMax
		// controls how many units of the base unit to spawn.
		MinionGroupMin int // called `MinGrp` in monstats.txt
		MinionGroupMax int // called `MaxGrp` in monstats.txt

		// PopulationReductionPercent controls the overall chance something will spawn in
		// percentages. Blank entries are the same as 100%.
		PopulationReductionPercent int // called `sparsePopulate` in monstats.txt

		// SpeedBase, SpeedRun
		// controls the walking and running speed of this monster respectively.
		// NOTE: RUN is only used if the monster has a RN mode and its AI uses that
		// mode.
		SpeedBase int // called `Velocity` in monstats.txt
		SpeedRun  int // called `Run` in monstats.txt

		// Rarity controls the overall odds that this monster will be spawned.
		// IE Lets say in Levels.txt you have two monsters set to spawn - Monster A
		// has rarity of 10 whereas Monster B has rarity of 1 and the level in
		// question is limited to 1 monster type. First the game sums up the
		// chances (11) and then calculates the odds of the monster spawning. Which
		// would be 1/11 (9% chance) for Monster B and 10/11 (91% chance) for
		// Monster A, thus Monster A is a lot more common than monster B. If you set
		// this column to 0 then the monster will never be selected by Levels.txt
		// for obvious reasons.
		Rarity int // called `Rarity` in monstats.txt

		// LevelNormal, LevelNightmare, LevelHell
		// controls the monsters level on the specified difficulty. This setting is
		// only used on normal. On nightmare and hell the monsters level is
		// identical with the area level from Levels.txt, unless your monster has
		// BOSS column set to 1, in this case its level will be always taken from
		// these 3 columns.
		LevelNormal    int // called `Level` in monstats.txt
		LevelNightmare int // called `Level(N)` in monstats.txt
		LevelHell      int // called `Level(H)` in monstats.txt

		// used by the game to tell AIs which unit to target first. The higher this
		// is the higher the threat level. Setting this to 25 or so on Maggot Eggs
		// would make your Mercenary NPC try to destroy those first.
		ThreatLevel int // called `threat` in monstats.txt

		// AiDelayNormal, AiDelayNightmare, AiDelayHell
		// this controls delays between AI ticks (on normal, nightmare and hell).
		// The lower the number, the faster the AI's will attack thanks to reduced
		// delay between swings, casting spells, throwing missiles etc. Please
		// remember that some AI's got individual delays between attacks, this will
		// still make them faster and seemingly more deadly though.
		AiDelayNormal    int // called `aidel` in monstats.txt
		AiDelayNightmare int // called `aidel(N)` in monstats.txt
		AiDelayHell      int // called `aidel(H)` in monstats.txt

		// AiDistanceNormal, AiDistanceNightmare, AiDistanceHell
		// the distance in cells from which AI is activated. Most AI"s have base
		// hardcoded activation radius of 35 which stands for a distamnce of about
		// 1 screen, thus leaving these fields blank sets this to 35 automatically.
		AiDistanceNormal    int // called `aidist` in monstats.txt
		AiDistanceNightmare int // called `aidist(N)` in monstats.txt
		AiDistanceHell      int // called `aidist(H)` in monstats.txt

		// AiParameterNormal1, AiParameterNightmare1, AiParameterHell1
		// these cells are very important, they pass on parameters (in percentage)
		// to the AI code. For descriptions about what all these AI's do, check
		// The AI Compendium. https://d2mods.info/forum/viewtopic.php?t=36230
		// Warning: many people have trouble with the AI of the Imps, this AI is
		// special and uses multiple rows.
		AiParameterNormal1    int // called `aip1` in monstats.txt
		AiParameterNormal2    int // called `aip2` in monstats.txt
		AiParameterNormal3    int // called `aip3` in monstats.txt
		AiParameterNormal4    int // called `aip4` in monstats.txt
		AiParameterNormal5    int // called `aip5` in monstats.txt
		AiParameterNormal6    int // called `aip6` in monstats.txt
		AiParameterNormal7    int // called `aip7` in monstats.txt
		AiParameterNormal8    int // called `aip8` in monstats.txt
		AiParameterNightmare1 int // called `aip1(N)` in monstats.txt
		AiParameterNightmare2 int // called `aip2(N)` in monstats.txt
		AiParameterNightmare3 int // called `aip3(N)` in monstats.txt
		AiParameterNightmare4 int // called `aip4(N)` in monstats.txt
		AiParameterNightmare5 int // called `aip5(N)` in monstats.txt
		AiParameterNightmare6 int // called `aip6(N)` in monstats.txt
		AiParameterNightmare7 int // called `aip7(N)` in monstats.txt
		AiParameterNightmare8 int // called `aip8(N)` in monstats.txt
		AiParameterHell1      int // called `aip1(H)` in monstats.txt
		AiParameterHell2      int // called `aip2(H)` in monstats.txt
		AiParameterHell3      int // called `aip3(H)` in monstats.txt
		AiParameterHell4      int // called `aip4(H)` in monstats.txt
		AiParameterHell5      int // called `aip5(H)` in monstats.txt
		AiParameterHell6      int // called `aip6(H)` in monstats.txt
		AiParameterHell7      int // called `aip7(H)` in monstats.txt
		AiParameterHell8      int // called `aip8(H)` in monstats.txt

		// Alignment controls whenever the monster fights on your side or
		// fights against you (or if it just walks around, IE a critter).
		// If you want to turn some obsolete NPCs into enemies, this is
		// one of the settings you will need to modify. Setting it to 2
		// without adjusting other settings (related to AI and also some
		// in MonStats2) it will simply attack everything.
		Alignment d2enum.MonsterAlignmentType // called `Align` in monstats.txt

		// SkillLevel1 -- SkillLevel8
		// the skill level of the skill in question. This gets a bonus on nightmare
		// and hell which you can modify in DifficultyLevels.txt.
		SkillLevel1 int // called `Sk1lvl` in monstats.txt
		SkillLevel2 int // called `Sk2lvl` in monstats.txt
		SkillLevel3 int // called `Sk3lvl` in monstats.txt
		SkillLevel4 int // called `Sk4lvl` in monstats.txt
		SkillLevel5 int // called `Sk5lvl` in monstats.txt
		SkillLevel6 int // called `Sk6lvl` in monstats.txt
		SkillLevel7 int // called `Sk7lvl` in monstats.txt
		SkillLevel8 int // called `Sk8lvl` in monstats.txt

		// LeechSensitivityNormal / Nightmare / Hell
		// controls the effectiveness of Life and Mana steal from equipment on this
		// unit on the respective difficulties. 0=Can’t leech at all. Remember that
		// besides this, Life and Mana Steal is further limited by DifficultyLevels.txt.
		LeechSensitivityNormal    int // called `Drain` in monstats.txt
		LeechSensitivityNightmare int // called `Drain(N)` in monstats.txt
		LeechSensitivityHell      int // called `Drain(H)` in monstats.txt

		// ColdSensitivityNormal / Nightmare / Hell
		// controls the effectiveness of cold effect and its duration and freeze
		// duration on this unit. The lower this value is, the more speed this unit
		// looses when its under the effect of cold, also freezing/cold effect will
		// stay for longer. Positive values will make the unit faster (thanks to
		// Brother Laz for confirming my assumption), and 0 will make it
		// unfreezeable. Besides this, cold length and freeze length settings are
		// also set in DifficultyLevels.txt.
		ColdSensitivityNormal    int // called `coldeffect` in monstats.txt
		ColdSensitivityNightmare int // called `coldeffect(N)` in monstats.txt
		ColdSensitivityHell      int // called `coldeffect(H)` in monstats.txt

		// ResistancePhysicalNormal
		// Damage resistance on the respective difficulties. Negative values mean
		// that the unit takes more damage from this element, values at or above 100
		// will result in immunity.
		ResistancePhysicalNormal     int // called `ResDm` in monstats.txt
		ResistancePhysicalNightmare  int // called `ResDm(N)` in monstats.txt
		ResistancePhysicalHell       int // called `ResDm(H)` in monstats.txt
		ResistanceMagicNormal        int // called `ResMa` in monstats.txt
		ResistanceMagicNightmare     int // called `ResMa(N)` in monstats.txt
		ResistanceMagicHell          int // called `ResMa(H)` in monstats.txt
		ResistanceFireNormal         int // called `ResFi` in monstats.txt
		ResistanceFireNightmare      int // called `ResFi(N)` in monstats.txt
		ResistanceFireHell           int // called `ResFi(H)` in monstats.txt
		ResistanceLightningNormal    int // called `ResLi` in monstats.txt
		ResistanceLightningNightmare int // called `ResLi(N)` in monstats.txt
		ResistanceLightningHell      int // called `ResLi(H)` in monstats.txt
		ResistanceColdNormal         int // called `ResCo` in monstats.txt
		ResistanceColdNightmare      int // called `ResCo(N)` in monstats.txt
		ResistanceColdHell           int // called `ResCo(H)` in monstats.txt
		ResistancePoisonNormal       int // called `ResPo` in monstats.txt
		ResistancePoisonNightmare    int // called `ResPo(N)` in monstats.txt
		ResistancePoisonHell         int // called `ResPo(H)` in monstats.txt

		// HealthRegenPerFrame
		// this controls how much health this unit regenerates per frame. Sometimes
		// this is altered by the units AI. The formula is (REGEN * HP) / 4096. So
		// a monster with 200 hp and a regen rate of 10 would regenerate ~0,5 HP
		// (~12 per second) every frame (1 second = 25 frames).
		HealthRegenPerFrame int // called `DamageRegen` in monstats.txt

		// ChanceToBlockNormal / Nightmare / Hell
		// this units chance to block. See the above column for details when this
		// applies or not. Monsters are capped at 75% block as players are AFAIK.
		ChanceToBlockNormal    int // called `ToBlock` in monstats.txt
		ChanceToBlockNightmare int // called `ToBlock(N)` in monstats.txt
		ChanceToBlockHell      int // called `ToBlock(H)` in monstats.txt

		// ChanceDeadlyStrike
		// this units chance of scoring a critical hit (dealing double the damage).
		ChanceDeadlyStrike int // called `Crit` in monstats.txt

		// MinHPNormal -- MaxHPHell
		// minHp, maxHp, minHp(N), maxHp(N), minHp(H), maxHp(H): this units minimum
		// and maximum HP on the respective difficulties.
		// NOTE: Monster HitPoints are calculated as the following: (minHp * Hp from
		// MonLvl.txt)/100 for minimal hp and (maxHp * Hp from MonLvl.txt)/100 for
		// maximum hp.
		// To make this guide idiot-proof, we will calculate the hit points of a
		// Hungry Dead from vanilla on Normal difficulty and Single Player mode.
		// It has minHp = 101 and maxHp = 186 and level 2. Hp for level 2 in
		// MonLvl.txt = 9
		// It means Hungry Dead has (101*9)/100 ~ 9 of minimum hp and
		// (186*9)/100 ~ 17 maximum hit points. You have to remember monsters on
		// nightmare and hell take their level (unless Boss = 1) from area level of
		// Levels.txt instead of Level column of MonStats.txt. I hope this is clear.
		MinHPNormal    int // called `minHP` in monstats.txt
		MinHPNightmare int // called `MinHP(N)` in monstats.txt
		MinHPHell      int // called `MinHP(H)` in monstats.txt
		MaxHPNormal    int // called `maxHP` in monstats.txt
		MaxHPNightmare int // called `MaxHP(N)` in monstats.txt
		MaxHPHell      int // called `MaxHP(H)` in monstats.txt

		// ArmorClassNormal -- Hell
		// this units Armor Class on the respective difficulties. The calculation is
		// the same (analogical) as for hit points.
		ArmorClassNormal    int // called `AC` in monstats.txt
		ArmorClassNightmare int // called `AC(N)` in monstats.txt
		ArmorClassHell      int // called `AC(H)` in monstats.txt

		// ExperienceNormal -- Hell
		// the experience you get when killing this unit on the respective
		// difficulty. The calculation is the same (analogical) as for hit points.
		ExperienceNormal    int // called `Exp` in monstats.txt
		ExperienceNightmare int // called `Exp(N)` in monstats.txt
		ExperienceHell      int // called `Exp(H)` in monstats.txt

		// DamageMinA1Normal / Nightmare / Hell
		// DamageMaxA1Normal / Nightmare /Hell
		// this units minimum and maximum damage when it uses A1/A2/S1 mode.
		// The calculation is the same (analogical) as for hit points.
		DamageMinA1Normal    int // called `A1MinD` in monstats.txt
		DamageMinA1Nightmare int // called `A1MinD(N)` in monstats.txt
		DamageMinA1Hell      int // called `A1MinD(H)` in monstats.txt
		DamageMaxA1Normal    int // called `A1MaxD` in monstats.txt
		DamageMaxA1Nightmare int // called `A1MaxD(N)` in monstats.txt
		DamageMaxA1Hell      int // called `A1MaxD(H)` in monstats.txt
		DamageMinA2Normal    int // called `A2MinD` in monstats.txt
		DamageMinA2Nightmare int // called `A2MinD(N)` in monstats.txt
		DamageMinA2Hell      int // called `A2MinD(H)` in monstats.txt
		DamageMaxA2Normal    int // called `A2MaxD` in monstats.txt
		DamageMaxA2Nightmare int // called `A2MaxD(N)` in monstats.txt
		DamageMaxA2Hell      int // called `A2MaxD(H)` in monstats.txt
		DamageMinS1Normal    int // called `S1MinD` in monstats.txt
		DamageMinS1Nightmare int // called `S1MinD(N)` in monstats.txt
		DamageMinS1Hell      int // called `S1MinD(H)` in monstats.txt
		DamageMaxS1Normal    int // called `S1MaxD` in monstats.txt
		DamageMaxS1Nightmare int // called `S1MaxD(N)` in monstats.txt
		DamageMaxS1Hell      int // called `S1MaxD(H)` in monstats.txt

		// AttackRatingA1Normal AttackRatingS1Hell
		// this units attack rating for A1/A2/S1 mode on the respective difficulties
		// The calculation is the same (analogical) as for hit points.
		AttackRatingA1Normal    int // called `A1TH` in monstats.txt
		AttackRatingA1Nightmare int // called `A1TH(N)` in monstats.txt
		AttackRatingA1Hell      int // called `A1TH(H)` in monstats.txt
		AttackRatingA2Normal    int // called `A2TH` in monstats.txt
		AttackRatingA2Nightmare int // called `A2TH(N)` in monstats.txt
		AttackRatingA2Hell      int // called `A2TH(H)` in monstats.txt
		AttackRatingS1Normal    int // called `S1TH` in monstats.txt
		AttackRatingS1Nightmare int // called `S1TH(N)` in monstats.txt
		AttackRatingS1Hell      int // called `S1TH(H)` in monstats.txt

		// ElementChance1Normal -- ElementChance3Hell
		// chance to append elemental damage to an attack on the respective
		// difficulties. 0=Never append, 100=Always append.
		ElementChance1Normal    int // called `El1Pct` in monstats.txt
		ElementChance1Nightmare int // called `El1Pct(N)` in monstats.txt
		ElementChance1Hell      int // called `El1Pct(H)` in monstats.txt
		ElementChance2Normal    int // called `El2Pct` in monstats.txt
		ElementChance2Nightmare int // called `El2Pct(N)` in monstats.txt
		ElementChance2Hell      int // called `El2Pct(H)` in monstats.txt
		ElementChance3Normal    int // called `El3Pct` in monstats.txt
		ElementChance3Nightmare int // called `El3Pct(N)` in monstats.txt
		ElementChance3Hell      int // called `El3Pct(H)` in monstats.txt

		// ElementDamageMin1Normal -- ElementDamageMax3Hell
		// minimum and Maximum elemental damage to append to the attack on the
		// respective difficulties. Note that you should only append elemental
		// damage to those missiles that don’t have any set in Missiles.txt. The
		// calculation is the same (analogical) as for hit points.
		ElementDamageMin1Normal    int // called `El1MinD` in monstats.txt
		ElementDamageMin1Nightmare int // called `El1MinD(N)` in monstats.txt
		ElementDamageMin1Hell      int // called `El1MinD(H)` in monstats.txt
		ElementDamageMin2Normal    int // called `El2MinD` in monstats.txt
		ElementDamageMin2Nightmare int // called `El2MinD(N)` in monstats.txt
		ElementDamageMin2Hell      int // called `El2MinD(H)` in monstats.txt
		ElementDamageMin3Normal    int // called `El3MinD` in monstats.txt
		ElementDamageMin3Nightmare int // called `El3MinD(N)` in monstats.txt
		ElementDamageMin3Hell      int // called `El3MinD(H)` in monstats.txt
		ElementDamageMax1Normal    int // called `El1MaxD` in monstats.txt
		ElementDamageMax1Nightmare int // called `El1MaxD(N)` in monstats.txt
		ElementDamageMax1Hell      int // called `El1MaxD(H)` in monstats.txt
		ElementDamageMax2Normal    int // called `El2MaxD` in monstats.txt
		ElementDamageMax2Nightmare int // called `El2MaxD(N)` in monstats.txt
		ElementDamageMax2Hell      int // called `El2MaxD(H)` in monstats.txt
		ElementDamageMax3Normal    int // called `El3MaxD` in monstats.txt
		ElementDamageMax3Nightmare int // called `El3MaxD(N)` in monstats.txt
		ElementDamageMax3Hell      int // called `El3MaxD(H)` in monstats.txt

		// ElementDuration1Normal -- ElementDuration3Hell
		// duration of the elemental effect (for freeze, burn, cold, poison and
		// stun) on the respective difficulties.
		ElementDuration1Normal    int // called `El1Dur` in monstats.txt
		ElementDuration1Nightmare int // called `El1Dur(N)` in monstats.txt
		ElementDuration1Hell      int // called `El1Dur(H)` in monstats.txt
		ElementDuration2Normal    int // called `El2Dur` in monstats.txt
		ElementDuration2Nightmare int // called `El2Dur(N)` in monstats.txt
		ElementDuration2Hell      int // called `El2Dur(H)` in monstats.txt
		ElementDuration3Normal    int // called `El3Dur` in monstats.txt
		ElementDuration3Nightmare int // called `El3Dur(N)` in monstats.txt
		ElementDuration3Hell      int // called `El3Dur(H)` in monstats.txt

		// SpecialEndDeath
		// 0 == no special death
		// 1 == spawn minion1 on death
		// 2 == kill mounted minion on death (ie the guard tower)
		SpecialEndDeath int // called `SplEndDeath` in monstats.txt

		// Enabled controls whenever the unit can be
		// used at all for any purpose whatsoever. This is not the only setting
		// that controls this; there are some other things that can also disable
		// the unit (Rarity and isSpawn columns see those for description).
		Enabled bool // called `enabled` in monstats.txt

		// SpawnsMinions tells the game whenever this
		// unit is a “nest”. IE, monsters that spawn new monsters have this set to
		// 1. Note that you can make any monster spawn new monsters, irregardless of
		// its AI, all you need to do is adjust spawn related columns and make sure
		// one of its skills is either “Nest” or “Minion Spawner”.
		SpawnsMinions bool // called `placespawn` in monstats.txt

		// IsLeader controls if a monster is the leader of minions it spawns
		// a leadercan order "raid on target" it causes group members to use
		// SK1 instead of A1 and A2 modes while raiding.
		IsLeader bool // called `SetBoss` in monstats.txt

		// TransferLeadership is connected with the previous one,
		// when "boss of the group" is killed, the "leadership" is passed to one of
		// his minions.
		TransferLeadership bool // called `BossXfer` in monstats.txt

		// Boolean, 1=spawnable, 0=not spawnable. This controls whenever this unit
		// can be spawned via Levels.txt.
		IsLevelSpawnable bool // called `isSpawn` in monstats.txt

		// IsMelee controls whenever
		// this unit can spawn with boss modifiers such as multiple shot or not.
		IsMelee bool // called `isMelee` in monstats.txt

		// IsNPC controls whenever the unit is a NPC or not.
		IsNpc bool // called `npc` in monstats.txt

		// IsInteractable
		// controls whenever you can interact with this unit. IE this controls
		// whenever it opens a speech-box or menu when you click on the unit. To
		// turn units like Kaeleen or Flavie into enemies you will need to set this
		// to 0 (you will also need to set NPC to 0 for that).
		IsInteractable bool // called `interact` in monstats.txt

		// IsRanged tells the game whenever this is a ranged attacker. It will make it possible for
		// monsters to spawn with multiple shot modifier.
		IsRanged bool // called `rangedtype` in monstats.txt

		// HasInventory Controls whenever this
		// NPC or UNIT can carry items with it. For NPCs this means that you can
		// access their Inventory and buy items (if you disable this and then try to
		// access this feature it will cause a crash so don’t do it unless you know
		// what you’re doing). For Monsters this means that they can access their
		// equipment data in MonEquip.txt.
		HasInventory bool // called `inventory` in monstats.txt

		// CanEnterTown
		// controls whenever enemies can follow you into a town or not. This should be set to
		// 1 for everything that spawns in a town for obvious reasons. According to
		// informations from Ogodei, it also disables/enables collision in
		// singleplayer and allows pets to walk/not walk in city in multiplayer.
		// In multiplayer collision is always set to 0 for pets.
		CanEnterTown bool // called `inTown` in monstats.txt

		// IsUndeadLow, IsUndeadHigh
		// Blizzard used this to differentiate High and Low Undead (IE low
		// undead like Zombies, Skeletons etc are set to 1 here), both this and
		// HUNDEAD will make the unit be considered undead. Low undeads can be
		// resurrected by high undeads. High undeads can't resurrect eachother.
		IsUndeadLow  bool // called `lUndead` in monstats.txt
		IsUndeadHigh bool // called `hUndead` in monstats.txt

		// IsDemon makes the game consider this unit a demon.
		IsDemon bool // called `demon` in monstats.txt

		// IsFlying If you set this to 1 the monster will be able to move fly over
		// obstacles such as puddles and rivers.
		IsFlying bool // called `flying` in monstats.txt

		// CanOpenDoors controls whether monsters can open doors or not
		CanOpenDoors bool // called `opendoors` in monstats.txt

		// IsSpecialBoss controls whenever this unit
		// is a special boss, as mentioned already, monsters set as boss IGNORE the
		// level settings, IE they will always spawn with the levels specified in
		// MonStats.txt. Boss will gain some special resistances, such as immunity
		// to being stunned (!!!), also it will not be affected by things like
		// deadly strike the way normal monsters are.
		IsSpecialBoss bool // called `boss` in monstats.txt

		// IsActBoss
		// Setting this to 1 will give your monsters huge (300% IIRC) damage bonus
		// against hirelings and summons. Ever wondered why Diablo destroys your
		// skeletons with 1 fire nova while barely doing anything to you? Here is
		// your answer.
		IsActBoss bool // called `primeevil` in monstats.txt

		// IsKillable will make the monster absolutely unkillable.
		IsKillable bool // called `killable` in monstats.txt

		// IsAiSwitchable Gives controls if this units mind may
		// be altered by “mind altering skills” like Attract, Conversion, Revive
		IsAiSwitchable bool // called `switchai` in monstats.txt

		// DisableAura Monsters set to 0 here
		// will not be effected by friendly auras
		DisableAura bool // called `noAura` in monstats.txt

		// DisableMultiShot
		// This is another layer of security to prevent this modifier from spawning,
		// besides the ISMELEE layer.
		DisableMultiShot bool // called `nomultishot` in monstats.txt

		// DisableCounting
		// prevents your pets from being counted as population in said area, for
		// example thanks to this you can finish The Den Of Evil quest while having
		// pets summoned.
		DisableCounting bool // called `neverCount` in monstats.txt

		// IgnorePets
		// Summons and hirelings are ignored by this unit, 0=Summons and
		// hirelings are noticed by this unit. If you set this to 1 you will the
		// monsters going directly for the player.
		IgnorePets bool // called `petIgnore` in monstats.txt

		// DealsDamageOnDeath This works similar to corpse explosion (its based on
		// hitpoints) and damages the surrounding players when the unit dies. (Ever
		// wanted to prevent those undead stygian dolls from doing damage when they
		// die, this is all there is to it)
		DealsDamageOnDeath bool // called `deathDmg` in monstats.txt

		// GenericSpawn Has to do
		// something is with minions being transformed into suicide minions, the
		// exact purpose of this is a mystery.
		GenericSpawn bool // called `genericSpawn` in monstats.txt

		// IgnoreMonLevelTxt Does this unit use
		// MonLevel.txt or does it use the stats listed in MonStats.txt as is.
		// Setting this to 1 will result in an array of problems, such as the
		// appended elemental damage being completely ignored, irregardless of the
		// values in it.
		IgnoreMonLevelTxt bool // called `noRatio` in monstats.txt

		// CanBlockWithoutShield in order for a unit to
		// block it needs the BL mode, if this is set to 1 then it will block
		// irregardless of what modes it has.
		CanBlockWithoutShield bool // called `NoShldBlock` in monstats.txt

		// SpecialGetModeChart
		// Unknown but could be telling the game to look at some internal table.
		// This is used for some Act Bosses and monsters like Putrid Defilers.
		SpecialGetModeChart bool // called `SplGetModeChar` in monstats.txt

		// SpecialEndGeneric Works in conjunction with SPLCLIENTEND, this
		// makes the unit untargetable when it is first spawned (used for those monsters that are under water, under ground or fly above you)
		SpecialEndGeneric bool // called `SplEndGeneric` in monstats.txt

		// SpecialClientEnd Works in conjunction with SPLENDGENERIC, this
		// makes the unit invisible when it is first spawned (used for those
		// monsters that are under water, under ground or fly above you), this is
		// also used for units that have other special drawing setups.
		SpecialClientEnd bool // called `SplClientEnd` in monstats.txt

	}
)
