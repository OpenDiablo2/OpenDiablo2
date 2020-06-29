// d2datadict contains loaders for the txt file data
package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// https://d2mods.info/forum/kb/viewarticle?a=360

// TODO: these types are described in the notes below, but they aren't used yet
type MonsterCombatType int

const (
	// see notes on column `rangedtype`
	MonsterMelee MonsterCombatType = iota
	MonsterRanged
)

type MonsterAlignmentType int

const (
	MonsterEnemy MonsterAlignmentType = iota
	MonsterFriend
	MonsterNeutral
)

// i've tried to choose better field names instead of what's inside of the txt
// file. after the definition, if there is a comment after the type, it will
// be the original column name corresponding to that field.
// i've tried to preserve the order of the columns, as well.
type MonStatsRecord struct {

	// this column contains the pointer that will be used in other txt files
	// such as levels.txt and superuniques.txt.
	Key string // Id <-- this is the original field name

	// this is the actual internal ID of the unit (this is what the ID pointer
	// actually points at) remember that no two units can have the same ID,
	// this will result in lots of unpredictable behavior and crashes so please
	// don’t do it. This 'HarcCodedInDeX' is used for several things, such as
	// determining whenever the unit uses DCC or DC6 graphics (like mephisto
	// and the death animations of Diablo, the Maggoc Queen etc.), the hcIdx
	// column also links other hardcoded effects to the units, such as the
	// transparency on necro summons and the name-color change on unique boss
	// units (thanks to Kingpin for the info)
	Id string // hcIdx

	// this column contains the ID pointer of the “base” unit for this specific
	// monster type (ex. There are five types of “Fallen”; all of them have
	// fallen1 as their “base” unit). The baseID is responsible for some
	// hardcoded behaviors, for example moving thru walls (ghosts), knowing
	// what units to resurrect, create etc (putrid defilers, shamans etc), the
	// explosion appended to suicide minions (either cold, fire or ice). Thanks
	// to Kingpin for additional info on this column.
	BaseKey string // BaseId

	// this column contains the ID of the next unit in the chain. (Continuing
	// on the above example, fallen1 has the ID pointer of fallen2 in here).
	// If you want to make a monster subtype less you should simply leave this
	// blank and make BaseId point at itself (its ID pointer). The game uses
	// this for “map generated” monsters such as the fallen in the fallen camps,
	// which get picked based on area level (the same camp, that in the cold
	// plains contains normal fallen will contain carvers and devilkin
	// elsewhere, read Levels.txt to check how to adjust area level).
	NextKey string // NextInClass

	// this indicates which palette (color) entry the unit will use, most
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
	PaletteId int // TransLvl

	// this column contains the string-key used in the TBL (string.tbl,
	// expansionstring.tbl and patchstring.tbl) files to make this monsters
	// name appear when you highlight it. Without that your monster will be
	// displayed as "not used - tell ken" or "an evil force".
	// Warning: string keys are case sensitive, so if you enter a key like “Foo”
	// in monstats.txt you must enter exactly that in the TBL file! (IE if you
	// enter “foo” or “fOO” you will get "not used - tell ken" or "an evil force").
	NameStringTableKey string // NameStr

	// this column contains the ID pointer to an entry in MonStats2.txt.
	// In 1.10 Blizzard has moved all the graphical aspects (light radius,
	// bleeding etc) to a new file to conserve space (MonStats.txt is one column
	// short of 256, the maximum MS Excel can handle, and that’s what they
	// probably used for their files).
	ExtraDataKey string // MonStatsEx

	// this column contains the ID pointer to an entry in MonProp.txt which
	// controls what special modifiers are appended to the unit, for example
	// you can use it to give your monsters auras, states, random resistances
	// or immunities, give them “get hit skills” and almost anything else.
	PropertiesKey string // MonProp

	// this column contains the group ID of the “super group” this monster
	// belongs to, IE all skeletons belong to the "super group" skeleton. The
	// 1.10 MonType.txt works exactly like ItemTypes.txt, furthermore this file
	// is used for special modifiers such as additional damage vs. monster-class.
	MonsterGroup string // MonType

	// this column tells the game which AI to use for this monster. Every AI
	// needs a specific set of animation modes (GH, A1, A2, S1, WL, RN etc).
	// Most of AI's require a configuration in aip columns (read about them
	// below), without that they (most of them) will do absolutely nothing.
	AiKey string // AI

	// this column contains the string-key used in the TBL (string.tbl,
	// expansionstring.tbl and patchstring.tbl) files for the monsters
	// description (leave it blank for no description).
	// NOTE: ever wondered how to make it say something below the monster
	// name (such as “Drains Mana and Stamina etc), well this is how you do it.
	// Just put the string-key of the string you want to display below the
	// monsters name in here.
	DescriptionStringTableKey string // DescStr

	// this is the only graphical setting (besides TransLvl) left in
	// MonStats.txt, this controls which token (IE name of a folder that
	// contains animations) the game uses for this monster.
	AnimationDirectoryToken string // Code

	// Boolean, 1=enabled, 0=disabled. This controls whenever the unit can be
	// used at all for any purpose whatsoever. This is not the only setting
	// that controls this; there are some other things that can also disable
	// the unit (Rarity and isSpawn columns see those for description).
	Enabled bool // enabled

	// Boolean, 1=ranged attacker, 0=melee attacker. This tells the game
	// whenever this is a ranged attacker. It will make it possible for
	// monsters to spawn with multiple shot modifier. Also, I suspect this has
	// to do with the RANGEDSPAWN column in Levels.txt. (Could it be the game
	// uses this for preference settings when spawning monsters to avoid areas
	// being populated only with melee monsters, IE the game picks a set amount
	// of monsters for every level, randomly, only based on their rarity values,
	// from those specified in Levels.txt, now I assume that it could pick 4
	// melee monsters, however in 1.10 Blizzard added a check to prevent this
	// from happening AFAIK and this could be how they control it.)
	IsRanged int // rangedtype

	// Boolean, 1=spawner, 0=not a spawner. This tells the game whenever this
	// unit is a “nest”. IE, monsters that spawn new monsters have this set to
	// 1. Note that you can make any monster spawn new monsters, irregardless of
	// its AI, all you need to do is adjust spawn related columns and make sure
	// one of its skills is either “Nest” or “Minion Spawner”.
	SpawnsMinions bool // placespawn

	// this column contains the ID pointer of the unit to spawn. (in case it is
	// a spawner that is), so if you want to make a new monster that generates
	// Balrogs this is where you would put the Balrog ID pointer.
	SpawnId string // spawn

	// the x/y offsets at which spawned monsters are placed. IE this prevents
	// the spawned monsters from being created at the same x/y coordinates as
	// the spawner itself, albeit its not needed, Blizzards collision detection
	// system is good enough to prevent them from getting stuck.
	SpawnOffsetX int // spawnx
	SpawnOffsetY int // spawny

	// which animation mode will the spawned monster be spawned in. IE. If you
	// make a golem summoner (yes I know, “very original”) you could put S1 in
	// here to make it look as if the golems are really summoned (otherwise they
	// would just appear), in most cases you will probably want to use NU mode
	// or a sequence though.
	SpawnAnimationKey string // spawnmode

	// these columns contain the ID pointers to the minions that spawn around
	// this unit when it is created.
	// NOTE + MINI TUTORIAL: lets say you want your super-strong boss to spawn
	// with 5 Oblivion Knights. To do this you would simply enter the Oblivion
	// Knights ID pointer in the MINION1 column. And set PARTYMIN and PARTYMAX
	// both to 5. MINION1/2 are used for several other things. If the monster
	// spawns as unique or superunique then it will have the unit from MINION1/2
	// set as its minion instead of monsters of its own type. That’s why
	// Lord De Seis doesn’t spawn with other oblivion knights anymore. To
	// semi-circumvent this I suggest you simply put the monsters ID pointer in
	// the MINION2 column (I.E. if you give the Oblivion Knights their own ID
	// pointer in MINION2, Lord De Seis should spawn with both Doom Knights and
	// Oblivion Knights again). The other use controls what monster is created
	// when this unit dies. For example Flayer Shamans will spawn a regular
	// Flayer when they are killed. To enable this you must set SPLENDDEATH to
	// 1, make sure the unit you spawn this was has a raise or resurrect
	// sequence otherwise it will look weird (but it works).
	// NOTE: THESE ARE SPAWNED ON MONSTER CREATION, NOTHING TO DO WITH ABOVE
	MinionId1 string // minion1
	MinionId2 string // minion2

	// Boolean, 1=set unit as boss, 0=don’t set unit as boss. This is related to
	// hardcoded behavior of some of the AI's. IE Scarabs are spawned a bit
	// different than most of other monsters. There are 1 or 2 with minions of
	// own kind. Thanks to this colums, the "bosses" of a group can have
	// ('can" because you may set the chance in percentages ia aip5 column of
	// Scarab AI) a chance to order "raid on tagret", what does it change is
	// they will always use SK1 instead of A1 and A2 modes while raiding.
	IsLeader bool // SetBoss

	// Boolean, 1=true, 0=false. This field is connected with the previous one,
	// when "boss of the group" is killed, his "leadership" is passed to one of
	// his minions.
	TransferLeadership bool // BossXfer

	// how many minions are spawned together with this unit. As mentioned above
	// in the MINION1/2 columns, this controls the quantity of minions this
	// unit has.
	MinionPartyMin int // PartyMin
	MinionPartyMax int // PartyMax

	// exactly like the previous two columns, just that this controls how many
	// units of the base unit to spawn. In versions 1.00-1.06, setting the
	// minimum to more then 99 would crash the game.
	MinionGroupMin int // MinGrp
	MinionGroupMax int // MaxGrp

	// this column controls the overall chance something will spawn in
	// percentages, leaving it blank is the same as 100%. If you enter "80" in
	// this column then whenever the game chooses to spawn this unit it will
	// first roll out the chances. So in 2 out of 10 cases the monster will not
	// be spawned and other one may take its place. If you use some low number
	// there can be a situation when the game doesn't roll out the monster and
	// if this monster has units specified in minion1/minion2 columns it will
	// only spawn those minions without the "main" unit.
	PopulationReductionPercent int // sparsePopulate

	// controls the walking and running speed of this monster respectively.
	// NOTE: RUN is only used if the monster has a RN mode and its AI uses that
	// mode.
	SpeedBase int // Velocity
	SpeedRun  int // Run

	// this column controls the overall odds that this monster will be spawned.
	// IE Lets say in Levels.txt you have two monsters set to spawn - Monster A
	// has rarity of 10 whereas Monster B has rarity of 1 and the level in
	// question is limited to 1 monster type. First the game sums up the
	// chances (11) and then calculates the odds of the monster spawning. Which
	// would be 1/11 (9% chance) for Monster B and 10/11 (91% chance) for
	// Monster A, thus Monster A is a lot more common than monster B. If you set
	// this column to 0 then the monster will never be selected by Levels.txt
	// for obvious reasons.
	Rarity int // Rarity

	// controls the monsters level on the specified difficulty. This setting is
	// only used on normal. On nightmare and hell the monsters level is
	// identical with the area level from Levels.txt, unless your monster has
	// BOSS column set to 1, in this case its level will be always taken from
	// these 3 columns.
	LevelNormal    int // Level
	LevelNightmare int // Level(N)
	LevelHell      int // Level(H)

	// specifies the ID pointer to this monsters “Sound Bank” in MonSound.txt
	// when this monster is normal.
	SoundKeyNormal  string // MonSound
	SoundKeySpecial string // UMonSound

	// used by the game to tell AIs which unit to target first. The higher this
	// is the higher the threat level. Setting this to 25 or so on Maggot Eggs
	// would make your Merc try to destroy those first.
	ThreatLevel int // threat

	// this controls delays between AI ticks (on normal, nightmare and hell).
	// The lower the number, the faster the AI's will attack thanks to reduced
	// delay between swings, casting spells, throwing missiles etc. Please
	// remember that some AI's got individual delays between attacks, this will
	// still make them faster and seemingly more deadly though.
	AiDelayNormal    int // aidel
	AiDelayNightmare int // aidel(N)
	AiDelayHell      int // aidel(H)

	// the distance in cells from which AI is activated. Most AI"s have base
	// hardcoded activation radius of 35 which stands for a distamnce of about
	// 1 screen, thus leaving these fields blank sets this to 35 automatically.
	AiDistanceNormal    int // aidist
	AiDistanceNightmare int // aidist(N)
	AiDistanceHell      int // aidist(H)

	// these cells are very important, they pass on parameters (in percentage)
	// to the AI code. For descriptions about what all these AI's do, check
	// The AI Compendium. https://d2mods.info/forum/viewtopic.php?t=36230
	// Warning: many people have trouble with the AI of the Imps, this AI is
	// special and uses multiple rows.
	AiParameterNormal1    int // aip1
	AiParameterNormal2    int // aip2
	AiParameterNormal3    int // aip3
	AiParameterNormal4    int // aip4
	AiParameterNormal5    int // aip5
	AiParameterNormal6    int // aip6
	AiParameterNormal7    int // aip7
	AiParameterNormal8    int // aip8
	AiParameterNightmare1 int // aip1
	AiParameterNightmare2 int // aip2
	AiParameterNightmare3 int // aip3
	AiParameterNightmare4 int // aip4
	AiParameterNightmare5 int // aip5
	AiParameterNightmare6 int // aip6
	AiParameterNightmare7 int // aip7
	AiParameterNightmare8 int // aip8
	AiParameterHell1      int // aip1
	AiParameterHell2      int // aip2
	AiParameterHell3      int // aip3
	AiParameterHell4      int // aip4
	AiParameterHell5      int // aip5
	AiParameterHell6      int // aip6
	AiParameterHell7      int // aip7
	AiParameterHell8      int // aip8

	// these columns control “non-skill-related” missiles used by the monster.
	// For example if you enter a missile ID pointer (from Missiles.txt) in
	// MissA1 then, whenever the monster uses its A1 mode, it will shoot a
	// missile, this however will successfully prevent it from dealing any damage
	// with the swing of A1.
	// NOTE: for the beginners, A1=Attack1, A2=Attack2, S1=Skill1, S2=Skill2,
	// S3=Skill3, S4=Skill4, C=Cast, SQ=Sequence.
	MissileA1 string // MissA1
	MissileA2 string // MissA2
	MissileS1 string // MissS1
	MissileS2 string // MissS2
	MissileS3 string // MissS3
	MissileS4 string // MissS4
	MissileC  string // MissC
	MissileSQ string // MissSQ

	// Switch, 0=enemy, 1=aligned, 2=neutral. This setting controls whenever the
	// monster fights on your side or fights against you (or if it just walks
	// around, IE a critter). If you want to turn some obsolete NPCs into
	// enemies, this is one of the settings you will need to modify. Setting it
	// to 2 without adjusting other settings (related to AI and also some in
	// MonStats2) it will simply attack everything.
	Alignment int // Align

	// Boolean, 1=spawnable, 0=not spawnable. This controls whenever this unit
	// can be spawned via Levels.txt.
	IsLevelSpawnable bool // isSpawn

	// Boolean, 1=melee attacker, 0=not a melee attacker. This controls whenever
	// this unit can spawn with boss modifiers such as multiple shot or not.
	// IE melee monsters will never spawn with multiple shot.
	IsMelee bool // isMelee

	// Boolean, 1=I’m a NPC, 0=I’m not. This controls whenever the unit is a NPC
	// or not.
	IsNpc bool // npc

	// Boolean, 1=Special NPC features enabled, 0=No special NPC features. This
	// controls whenever you can interact with this unit. IE this controls
	// whenever it opens a speech-box or menu when you click on the unit. To
	// turn units like Kaeleen or Flavie into enemies you will need to set this
	// to 0 (you will also need to set NPC to 0 for that).
	IsInteractable bool // interact

	// Boolean, 1=Has an inventory, 0=Has no inventory. Controls whenever this
	// NPC or UNIT can carry items with it. For NPCs this means that you can
	// access their Inventory and buy items (if you disable this and then try to
	// access this feature it will cause a crash so don’t do it unless you know
	// what you’re doing). For Monsters this means that they can access their
	// equipment data in MonEquip.txt.
	HasInventory bool // inventory

	// Boolean, 1=I can enter towns, 0=I can’t enter towns. This controls
	// whenever enemies can follow you into a town or not. This should be set to
	// 1 for everything that spawns in a town for obvious reasons. According to
	// informations from Ogodei, it also disables/enables collision in
	// singleplayer and allows pets to walk/not walk in city in multiplayer.
	// In multiplayer collision is always set to 0 for pets.
	CanEnterTown bool // inTown

	// Boolean. Blizzard used this to differentiate High and Low Undead (IE low
	// undead like Zombies, Skeletons etc are set to 1 here), both this and
	// HUNDEAD will make the unit be considered undead. Low undeads can be
	// resurrected by high undeads. High undeads can't resurrect eachother.
	IsUndeadLow  bool // lUndead
	IsUndeadHigh bool // hUndead

	// Boolean, This makes the game consider this unit a demon.
	IsDemon bool // demon

	// Boolean, If you set this to 1 the monster will be able to move fly over
	// obstacles such as puddles and rivers.
	// TIP: Setting this to 1 for Mephisto will prevent the possibility of
	// making him stuck and being slaughtered by blizzard or other ranged spells
	// easily.
	IsFlying bool // flying

	// Boolean, 1=I can open doors, 0=I’m too damn retarded to open doors. Ever
	// wanted to make the game more like D1 (where closing doors could actually
	// protect you), then this column is all you need. By setting this to 0 you
	// will successfully lobotomize the monster, thus he will not be able to open
	// doors any more.
	CanOpenDoors bool // opendoors

	// Boolean, 1=I’m a boss, 0=I’m not a boss. This controls whenever this unit
	// is a special boss, as mentioned already, monsters set as boss IGNORE the
	// level settings, IE they will always spawn with the levels specified in
	// MonStats.txt. Boss will gain some special resistances, such as immunity
	// to being stunned (!!!), also it will not be affected by things like
	// deadly strike the way normal monsters are.
	IsSpecialBoss bool // boss

	// Boolean, 1=I’m a prime evil, 0=I’m not a prime evil. (=Act Boss).
	// Setting this to 1 will give your monsters huge (300% IIRC) damage bonus
	// against hirelings and summons. Ever wondered why Diablo destroys your
	// skeletons with 1 fire nova while barely doing anything to you? Here is
	// your answer.
	IsActBoss bool // primeevil

	// Setting this to 0 will make the monster absolutely unkillable.
	IsKillable bool // killable

	// Boolean, 1=Can change sides, 0=Cannot change sides. Gives your monster
	// something what I call "Strong Mind", it decides if this units mind may
	// be altered by “mind altering skills” like Attract, Conversion, Revive
	// etc or not.
	IsAiSwitchable bool // switchai

	// Boolean, 1=Can’t get an aura, 0=Can get an aura. Monsters set to 0 here
	// will not be effected by friendly auras
	DisableAura bool // noAura

	// Boolean, 1=Can’t get multishot modifier, 0=Can get multishot modifier.
	// This is another layer of security to prevent this modifier from spawning,
	// besides the ISMELEE layer.
	DisableMultiShot bool // nomultishot

	// Boolean, 1=Never accounted for, 0=Accounted for. Setting this to 1
	// prevents your pets from being counted as population in said area, for
	// example thanks to this you can finish The Den Of Evil quest while having
	// pets summoned.
	DisableCounting bool // neverCount

	// Boolean 1=Summons and hirelings are ignored by this unit, 0=Summons and
	// hirelings are noticed by this unit. If you set this to 1 you will the
	// monsters going directly for the player.
	IgnorePets bool // petIgnore

	// Boolean, 1=Damage players colliding with my death animation, 0=Don’t
	// damage anything. This works similar to corpse explosion (its based on
	// hitpoints) and damages the surrounding players when the unit dies. (Ever
	// wanted to prevent those undead stygian dolls from doing damage when they
	// die, this is all there is to it)
	DealsDamageOnDeath bool // deathDmg

	// Boolean, 1=Use generic spawning, 0=Don’t use generic spawning. Has to do
	// something is with minions being transformed into suicide minions, the
	// exact purpose of this is a mystery to me though.
	GenericSpawn bool // genericSpawn

	// Boolean, 1=true, 0=false. Unused.
	// zoo

	// Switch, 1=Unknown, 2=Used for assassin traps, 0=Don’t send skills. This
	// is only used by two of the Assassin traps, however it doesn't serve any
	// purpose anymore.
	// SendSkills

	// the ID Pointer to the skill (from Skills.txt) the monster will cast when
	// this specific slot is accessed by the AI. Which slots are used is
	// determined by the units AI.
	SkillId1 string // Skill1
	SkillId2 string // Skill2
	SkillId3 string // Skill3
	SkillId4 string // Skill4
	SkillId5 string // Skill5
	SkillId6 string // Skill6
	SkillId7 string // Skill7
	SkillId8 string // Skill8

	// the graphical MODE (or SEQUENCE) this unit uses when it uses this skill.
	SkillAnimation1 string // Sk1mode
	SkillAnimation2 string // Sk2mode
	SkillAnimation3 string // Sk3mode
	SkillAnimation4 string // Sk4mode
	SkillAnimation5 string // Sk5mode
	SkillAnimation6 string // Sk6mode
	SkillAnimation7 string // Sk7mode
	SkillAnimation8 string // Sk8mode

	// the skill level of the skill in question. This gets a bonus on nightmare
	// and hell which you can modify in DifficultyLevels.txt.
	SkillLevel1 int // Sk1lvl
	SkillLevel2 int // Sk2lvl
	SkillLevel3 int // Sk3lvl
	SkillLevel4 int // Sk4lvl
	SkillLevel5 int // Sk5lvl
	SkillLevel6 int // Sk6lvl
	SkillLevel7 int // Sk7lvl
	SkillLevel8 int // Sk8lvl

	// controls the effectiveness of Life and Mana steal from equipment on this
	// unit on the respective difficulties. 0=Can’t leech at all. (negative
	// values don't damage you, thanks to Doombreed-x for testing this), setting
	// it to more then 100 would make LL and ML more effective. Remember that
	// besides this, Life and Mana Steal is further limited by
	// DifficultyLevels.txt.
	LeechSensitivityNormal    int // Drain
	LeechSensitivityNightmare int // Drain(N)
	LeechSensitivityHell      int // Drain(H)

	// controls the effectiveness of cold effect and its duration and freeze
	// duration on this unit. The lower this value is, the more speed this unit
	// looses when its under the effect of cold, also freezing/cold effect will
	// stay for longer. Positive values will make the unit faster (thanks to
	// Brother Laz for confirming my assumption), and 0 will make it
	// unfreezeable. Besides this, cold length and freeze length settings are
	// also set in DifficultyLevels.txt.
	ColdSensitivityNormal    int // coldeffect
	ColdSensitivityNightmare int // coldeffect(N)
	ColdSensitivityHell      int // coldeffect(H)

	// damage resistance on the respective difficulties. Negative values mean
	// that the unit takes more damage from this element, values at or above 100
	// will result in immunity.
	// NOTE: even though it may be quite obvious, I already met many people who
	// do not know about it. Each point of resistance means 1% of reduction (or
	// increase if the value is <0) of damage from said source. The same stands
	// for player characters. Same stands for all other resistances. 81% of fire
	// resist means 81% of incoming fire damage reduction.
	// NOTE 2: when resistance is >100 (immunity), you need 5 points of
	// resistance reducing stat applied DIRECTLY to the unit to reduce 1% of
	// immunity. It means if you want to break 105% immunity, you will need at
	// least 30 resist reduction stat (105 - 6 = 99 and 6 * 5 = 30).
	// NOTE 3: yes, yes, indeed, "damage reduced by x%" stat that may spawn on
	// items such as The Crown Of Ages or Ber rune is nothing else than physical
	// resistance.
	ResistancePhysicalNormal    int // ResDm
	ResistancePhysicalNightmare int // ResDm(N)
	ResistancePhysicalHell      int // ResDm(H)

	ResistanceMagicNormal    int // ResMa
	ResistanceMagicNightmare int // ResMa(N)
	ResistanceMagicHell      int // ResMa(H)

	ResistanceFireNormal    int // ResFi
	ResistanceFireNightmare int // ResFi(N)
	ResistanceFireHell      int // ResFi(H)

	ResistanceLightningNormal    int // ResLi
	ResistanceLightningNightmare int // ResLi(N)
	ResistanceLightningHell      int // ResLi(H)

	ResistanceColdNormal    int // ResCo
	ResistanceColdNightmare int // ResCo(N)
	ResistanceColdHell      int // ResCo(H)

	ResistancePoisonNormal    int // ResPo
	ResistancePoisonNightmare int // ResPo(N)
	ResistancePoisonHell      int // ResPo(H)

	// this controls how much health this unit regenerates per frame. Sometimes
	// this is altered by the units AI. The formula is (REGEN * HP) / 4096. So
	// a monster with 200 hp and a regen rate of 10 would regenerate ~0,5 HP
	// (~12 per second) every frame (1 second = 25 frames).
	HealthRegenPerFrame int // DamageRegen

	// ID Pointer to the skill that controls this units damage. This is used for
	// the druids summons. IE their damage is specified solely by Skills.txt and
	// not by MonStats.txt.
	DamageSkillId string // SkillDamage

	// Boolean, 1=Don’t use MonLevel.txt, 0=Use MonLevel.txt. Does this unit use
	// MonLevel.txt or does it use the stats listed in MonStats.txt as is.
	// Setting this to 1 will result in an array of problems, such as the
	// appended elemental damage being completely ignored, irregardless of the
	// values in it.
	IgnoreMonLevelTxt bool // noRatio

	// Boolean, 1=Can block without a blocking animation, 0=Can’t block without
	// a blocking animation. Quite self explanatory, in order for a unit to
	// block it needs the BL mode, if this is set to 1 then it will block
	// irregardless of what modes it has.
	CanBlockWithoutShield bool // NoShldBlock

	// this units chance to block. See the above column for details when this
	// applies or not. Monsters are capped at 75% block as players are AFAIK.
	ChanceToBlockNormal    int // ToBlock
	ChanceToBlockNightmare int // ToBlock(N)
	ChanceToBlockHell      int // ToBlock(H)

	// this units chance of scoring a critical hit (dealing double the damage).
	ChanceDeadlyStrike int // Crit

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
	MinHPNormal    int // minHP
	MinHPNightmare int // MinHP(N)
	MinHPHell      int // MinHP(H)

	MaxHPNormal    int // maxHP
	MaxHPNightmare int // MaxHP(N)
	MaxHPHell      int // MaxHP(H)

	// this units Armor Class on the respective difficulties. The calculation is
	// the same (analogical) as for hit points.
	ArmorClassNormal    int // AC
	ArmorClassNightmare int // AC(N)
	ArmorClassHell      int // AC(H)

	// the experience you get when killing this unit on the respective
	// difficulty. The calculation is the same (analogical) as for hit points.
	ExperienceNormal    int // Exp
	ExperienceNightmare int // Exp(N)
	ExperienceHell      int // Exp(H)

	// this units minimum and maximum damage when it uses A1/A2/S1 mode.
	// The calculation is the same (analogical) as for hit points.
	DamageMinA1Normal    int // A1MinD
	DamageMinA1Nightmare int // A1MinD(N)
	DamageMinA1Hell      int // A1MinD(H)

	DamageMaxA1Normal    int // A1MaxD
	DamageMaxA1Nightmare int // A1MaxD(N)
	DamageMaxA1Hell      int // A1MaxD(H)

	DamageMinA2Normal    int // A2MinD
	DamageMinA2Nightmare int // A2MinD(N)
	DamageMinA2Hell      int // A2MinD(H)

	DamageMaxA2Normal    int // A2MaxD
	DamageMaxA2Nightmare int // A2MaxD(N)
	DamageMaxA2Hell      int // A2MaxD(H)

	DamageMinS1Normal    int // S1MinD
	DamageMinS1Nightmare int // S1MinD(N)
	DamageMinS1Hell      int // S1MinD(H)

	DamageMaxS1Normal    int // S1MaxD
	DamageMaxS1Nightmare int // S1MaxD(N)
	DamageMaxS1Hell      int // S1MaxD(H)

	// this units attack rating for A1/A2/S1 mode on the respective difficulties
	// The calculation is the same (analogical) as for hit points.
	AttackRatingA1Normal    int // A1TH
	AttackRatingA1Nightmare int // A1TH(N)
	AttackRatingA1Hell      int // A1TH(H)

	AttackRatingA2Normal    int // A2TH
	AttackRatingA2Nightmare int // A2TH(N)
	AttackRatingA2Hell      int // A2TH(H)

	AttackRatingS1Normal    int // S1TH
	AttackRatingS1Nightmare int // S1TH(N)
	AttackRatingS1Hell      int // S1TH(H)

	// the mode to which the elemental damage is appended. The modes to which
	// you would usually attack elemental damage are A1, A2, S1, S2, S3, S4, SQ
	// or C as these are the only ones that naturally contain trigger bytes.
	ElementAttackMode1 string // El1Mode
	ElementAttackMode2 string // El2Mode
	ElementAttackMode3 string // El3Mode

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
	ElementType1 string // El1Type
	ElementType2 string // El2Type
	ElementType3 string // El3Type

	// chance to append elemental damage to an attack on the respective
	// difficulties. 0=Never append, 100=Always append.
	ElementChance1Normal    int // El1Pct
	ElementChance1Nightmare int // El1Pct(N)
	ElementChance1Hell      int // El1Pct(H)

	ElementChance2Normal    int // El2Pct
	ElementChance2Nightmare int // El2Pct(N)
	ElementChance2Hell      int // El2Pct(H)

	ElementChance3Normal    int // El3Pct
	ElementChance3Nightmare int // El3Pct(N)
	ElementChance3Hell      int // El3Pct(H)

	// minimum and Maximum elemental damage to append to the attack on the
	// respective difficulties. Note that you should only append elemental
	// damage to those missiles that don’t have any set in Missiles.txt. The
	// calculation is the same (analogical) as for hit points.
	ElementDamageMin1Normal    int // El1MinD
	ElementDamageMin1Nightmare int // El1MinD(N)
	ElementDamageMin1Hell      int // El1MinD(H)

	ElementDamageMin2Normal    int // El2MinD
	ElementDamageMin2Nightmare int // El2MinD(N)
	ElementDamageMin2Hell      int // El2MinD(H)

	ElementDamageMin3Normal    int // El3MinD
	ElementDamageMin3Nightmare int // El3MinD(N)
	ElementDamageMin3Hell      int // El3MinD(H)

	ElementDamageMax1Normal    int // El1MaxD
	ElementDamageMax1Nightmare int // El1MaxD(N)
	ElementDamageMax1Hell      int // El1MaxD(H)

	ElementDamageMax2Normal    int // El2MaxD
	ElementDamageMax2Nightmare int // El2MaxD(N)
	ElementDamageMax2Hell      int // El2MaxD(H)

	ElementDamageMax3Normal    int // El3MaxD
	ElementDamageMax3Nightmare int // El3MaxD(N)
	ElementDamageMax3Hell      int // El3MaxD(H)

	// duration of the elemental effect (for freeze, burn, cold, poison and
	// stun) on the respective difficulties.
	ElementDuration1Normal    int // El1Dur
	ElementDuration1Nightmare int // El1Dur(N)
	ElementDuration1Hell      int // El1Dur(H)

	ElementDuration2Normal    int // El2Dur
	ElementDuration2Nightmare int // El2Dur(N)
	ElementDuration2Hell      int // El2Dur(H)

	ElementDuration3Normal    int // El3Dur
	ElementDuration3Nightmare int // El3Dur(N)
	ElementDuration3Hell      int // El3Dur(H)

	// the TreasureClass used by this unit as a normal monster on the respective
	// difficulties.
	// NOTE: because of the new TreasureClass system introduced in 1.10 and
	// later patches, TC entries are only of minor influence regarding what TC
	// is being selected unless you change the system by editing
	// TreasureClassEX.txt.
	TreasureClassNormal    string // TreasureClass1
	TreasureClassNightmare string // TreasureClass1(N)
	TreasureClassHell      string // TreasureClass1(H)

	TreasureClassChampionNormal    string // TreasureClass2
	TreasureClassChampionNightmare string // TreasureClass2(N)
	TreasureClassChampionHell      string // TreasureClass2(H)

	TreasureClass3UniqueNormal    string // TreasureClass3
	TreasureClass3UniqueNightmare string // TreasureClass3(N)
	TreasureClass3UniqueHell      string // TreasureClass3(H)

	TreasureClassQuestNormal    string // TreasureClass4
	TreasureClassQuestNightmare string // TreasureClass4(N)
	TreasureClassQuestHell      string // TreasureClass4(H)

	// the ID of the Quest that triggers the Quest Treasureclass drop.
	TreasureClassQuestTriggerId string // TCQuestId

	//  the ID of the Quest State that you need to complete to trigger the Quest
	// Treasureclass trop.
	TreasureClassQuestComleteId string // TCQuestCP

	// TODO: fix these last 4 field names, they're fucking ridiculous
	// Switch, 0=no special death, 1=spawn the monster in the MINION1 column
	// when I die, 2=kill whatever monster is mounted to me when I die (used by
	// guard towers that kill the imps that are on top of them when they die).
	SpecialEndDeath int // SplEndDeath

	// Boolean, 1=Get Special Mode Chart, 0=Don’t get special mode chart.
	// Unknown but could be telling the game to look at some internal table.
	// This is used for some Act Bosses and monsters like Putrid Defilers.
	SpecialGetModeChart bool // SplGetModeChart

	// Boolean, 1=true, 0=false. Works in conjunction with SPLCLIENTEND, this
	// makes the unit untargetable when it is first spawned (used for those monsters that are under water, under ground or fly above you)
	SpecialEndGeneric bool // SplEndGeneric

	// Boolean, 1=true, 0=false. Works in conjunction with SPLENDGENERIC, this
	// makes the unit invisible when it is first spawned (used for those
	// monsters that are under water, under ground or fly above you), this is
	// also used for units that have other special drawing setups.
	SpecialClientEnd bool // SplClientEnd
}

var MonStats map[string]*MonStatsRecord

//nolint:funlen // Makes no sense to split
func LoadMonStats(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	MonStats = make(map[string]*MonStatsRecord, numRecords)

	for idx := range dict.Data {
		record := &MonStatsRecord{
			Key:                            dict.GetString("Id", idx),
			Id:                             dict.GetString("hcIdx", idx),
			BaseKey:                        dict.GetString("BaseId", idx),
			NextKey:                        dict.GetString("NextInClass", idx),
			PaletteId:                      dict.GetNumber("TransLvl", idx),
			NameStringTableKey:             dict.GetString("NameStr", idx),
			ExtraDataKey:                   dict.GetString("MonStatsEx", idx),
			PropertiesKey:                  dict.GetString("MonProp", idx),
			MonsterGroup:                   dict.GetString("MonType", idx),
			AiKey:                          dict.GetString("AI", idx),
			DescriptionStringTableKey:      dict.GetString("DescStr", idx),
			AnimationDirectoryToken:        dict.GetString("Code", idx),
			Enabled:                        dict.GetNumber("enabled", idx) > 0,
			IsRanged:                       dict.GetNumber("rangedtype", idx),
			SpawnsMinions:                  dict.GetNumber("placespawn", idx) > 0,
			SpawnId:                        dict.GetString("spawn", idx),
			SpawnOffsetX:                   dict.GetNumber("spawnx", idx),
			SpawnOffsetY:                   dict.GetNumber("spawny", idx),
			SpawnAnimationKey:              dict.GetString("spawnmode", idx),
			MinionId1:                      dict.GetString("minion1", idx),
			MinionId2:                      dict.GetString("minion2", idx),
			IsLeader:                       dict.GetNumber("SetBoss", idx) > 0,
			TransferLeadership:             dict.GetNumber("BossXfer", idx) > 0,
			MinionPartyMin:                 dict.GetNumber("PartyMin", idx),
			MinionPartyMax:                 dict.GetNumber("PartyMax", idx),
			MinionGroupMin:                 dict.GetNumber("MinGrp", idx),
			MinionGroupMax:                 dict.GetNumber("MaxGrp", idx),
			PopulationReductionPercent:     dict.GetNumber("sparsePopulate", idx),
			SpeedBase:                      dict.GetNumber("Velocity", idx),
			SpeedRun:                       dict.GetNumber("Run", idx),
			Rarity:                         dict.GetNumber("Rarity", idx),
			LevelNormal:                    dict.GetNumber("Level", idx),
			LevelNightmare:                 dict.GetNumber("Level(N)", idx),
			LevelHell:                      dict.GetNumber("Level(H)", idx),
			SoundKeyNormal:                 dict.GetString("MonSound", idx),
			SoundKeySpecial:                dict.GetString("UMonSound", idx),
			ThreatLevel:                    dict.GetNumber("threat", idx),
			AiDelayNormal:                  dict.GetNumber("aidel", idx),
			AiDelayNightmare:               dict.GetNumber("aidel(N)", idx),
			AiDelayHell:                    dict.GetNumber("aidel(H)", idx),
			AiDistanceNormal:               dict.GetNumber("aidist", idx),
			AiDistanceNightmare:            dict.GetNumber("aidist(N)", idx),
			AiDistanceHell:                 dict.GetNumber("aidist(H)", idx),
			AiParameterNormal1:             dict.GetNumber("aip1", idx),
			AiParameterNormal2:             dict.GetNumber("aip2", idx),
			AiParameterNormal3:             dict.GetNumber("aip3", idx),
			AiParameterNormal4:             dict.GetNumber("aip4", idx),
			AiParameterNormal5:             dict.GetNumber("aip5", idx),
			AiParameterNormal6:             dict.GetNumber("aip6", idx),
			AiParameterNormal7:             dict.GetNumber("aip7", idx),
			AiParameterNormal8:             dict.GetNumber("aip8", idx),
			AiParameterNightmare1:          dict.GetNumber("aip1", idx),
			AiParameterNightmare2:          dict.GetNumber("aip2", idx),
			AiParameterNightmare3:          dict.GetNumber("aip3", idx),
			AiParameterNightmare4:          dict.GetNumber("aip4", idx),
			AiParameterNightmare5:          dict.GetNumber("aip5", idx),
			AiParameterNightmare6:          dict.GetNumber("aip6", idx),
			AiParameterNightmare7:          dict.GetNumber("aip7", idx),
			AiParameterNightmare8:          dict.GetNumber("aip8", idx),
			AiParameterHell1:               dict.GetNumber("aip1", idx),
			AiParameterHell2:               dict.GetNumber("aip2", idx),
			AiParameterHell3:               dict.GetNumber("aip3", idx),
			AiParameterHell4:               dict.GetNumber("aip4", idx),
			AiParameterHell5:               dict.GetNumber("aip5", idx),
			AiParameterHell6:               dict.GetNumber("aip6", idx),
			AiParameterHell7:               dict.GetNumber("aip7", idx),
			AiParameterHell8:               dict.GetNumber("aip8", idx),
			MissileA1:                      dict.GetString("MissA1", idx),
			MissileA2:                      dict.GetString("MissA2", idx),
			MissileS1:                      dict.GetString("MissS1", idx),
			MissileS2:                      dict.GetString("MissS2", idx),
			MissileS3:                      dict.GetString("MissS3", idx),
			MissileS4:                      dict.GetString("MissS4", idx),
			MissileC:                       dict.GetString("MissC", idx),
			MissileSQ:                      dict.GetString("MissSQ", idx),
			Alignment:                      dict.GetNumber("Align", idx),
			IsLevelSpawnable:               dict.GetNumber("isSpawn", idx) > 0,
			IsMelee:                        dict.GetNumber("isMelee", idx) > 0,
			IsNpc:                          dict.GetNumber("npc", idx) > 0,
			IsInteractable:                 dict.GetNumber("interact", idx) > 0,
			HasInventory:                   dict.GetNumber("inventory", idx) > 0,
			CanEnterTown:                   dict.GetNumber("inTown", idx) > 0,
			IsUndeadLow:                    dict.GetNumber("lUndead", idx) > 0,
			IsUndeadHigh:                   dict.GetNumber("hUndead", idx) > 0,
			IsDemon:                        dict.GetNumber("demon", idx) > 0,
			IsFlying:                       dict.GetNumber("flying", idx) > 0,
			CanOpenDoors:                   dict.GetNumber("opendoors", idx) > 0,
			IsSpecialBoss:                  dict.GetNumber("boss", idx) > 0,
			IsActBoss:                      dict.GetNumber("primeevil", idx) > 0,
			IsKillable:                     dict.GetNumber("killable", idx) > 0,
			IsAiSwitchable:                 dict.GetNumber("switchai", idx) > 0,
			DisableAura:                    dict.GetNumber("noAura", idx) > 0,
			DisableMultiShot:               dict.GetNumber("nomultishot", idx) > 0,
			DisableCounting:                dict.GetNumber("neverCount", idx) > 0,
			IgnorePets:                     dict.GetNumber("petIgnore", idx) > 0,
			DealsDamageOnDeath:             dict.GetNumber("deathDmg", idx) > 0,
			GenericSpawn:                   dict.GetNumber("genericSpawn", idx) > 0,
			SkillId1:                       dict.GetString("Skill1", idx),
			SkillId2:                       dict.GetString("Skill2", idx),
			SkillId3:                       dict.GetString("Skill3", idx),
			SkillId4:                       dict.GetString("Skill4", idx),
			SkillId5:                       dict.GetString("Skill5", idx),
			SkillId6:                       dict.GetString("Skill6", idx),
			SkillId7:                       dict.GetString("Skill7", idx),
			SkillId8:                       dict.GetString("Skill8", idx),
			SkillAnimation1:                dict.GetString("Sk1mode", idx),
			SkillAnimation2:                dict.GetString("Sk2mode", idx),
			SkillAnimation3:                dict.GetString("Sk3mode", idx),
			SkillAnimation4:                dict.GetString("Sk4mode", idx),
			SkillAnimation5:                dict.GetString("Sk5mode", idx),
			SkillAnimation6:                dict.GetString("Sk6mode", idx),
			SkillAnimation7:                dict.GetString("Sk7mode", idx),
			SkillAnimation8:                dict.GetString("Sk8mode", idx),
			SkillLevel1:                    dict.GetNumber("Sk1lvl", idx),
			SkillLevel2:                    dict.GetNumber("Sk2lvl", idx),
			SkillLevel3:                    dict.GetNumber("Sk3lvl", idx),
			SkillLevel4:                    dict.GetNumber("Sk4lvl", idx),
			SkillLevel5:                    dict.GetNumber("Sk5lvl", idx),
			SkillLevel6:                    dict.GetNumber("Sk6lvl", idx),
			SkillLevel7:                    dict.GetNumber("Sk7lvl", idx),
			SkillLevel8:                    dict.GetNumber("Sk8lvl", idx),
			LeechSensitivityNormal:         dict.GetNumber("Drain", idx),
			LeechSensitivityNightmare:      dict.GetNumber("Drain(N)", idx),
			LeechSensitivityHell:           dict.GetNumber("Drain(H)", idx),
			ColdSensitivityNormal:          dict.GetNumber("coldeffect", idx),
			ColdSensitivityNightmare:       dict.GetNumber("coldeffect(N)", idx),
			ColdSensitivityHell:            dict.GetNumber("coldeffect(H)", idx),
			ResistancePhysicalNormal:       dict.GetNumber("ResDm", idx),
			ResistancePhysicalNightmare:    dict.GetNumber("ResDm(N)", idx),
			ResistancePhysicalHell:         dict.GetNumber("ResDm(H)", idx),
			ResistanceMagicNormal:          dict.GetNumber("ResMa", idx),
			ResistanceMagicNightmare:       dict.GetNumber("ResMa(N)", idx),
			ResistanceMagicHell:            dict.GetNumber("ResMa(H)", idx),
			ResistanceFireNormal:           dict.GetNumber("ResFi", idx),
			ResistanceFireNightmare:        dict.GetNumber("ResFi(N)", idx),
			ResistanceFireHell:             dict.GetNumber("ResFi(H)", idx),
			ResistanceLightningNormal:      dict.GetNumber("ResLi", idx),
			ResistanceLightningNightmare:   dict.GetNumber("ResLi(N)", idx),
			ResistanceLightningHell:        dict.GetNumber("ResLi(H)", idx),
			ResistanceColdNormal:           dict.GetNumber("ResCo", idx),
			ResistanceColdNightmare:        dict.GetNumber("ResCo(N)", idx),
			ResistanceColdHell:             dict.GetNumber("ResCo(H)", idx),
			ResistancePoisonNormal:         dict.GetNumber("ResPo", idx),
			ResistancePoisonNightmare:      dict.GetNumber("ResPo(N)", idx),
			ResistancePoisonHell:           dict.GetNumber("ResPo(H)", idx),
			HealthRegenPerFrame:            dict.GetNumber("DamageRegen", idx),
			DamageSkillId:                  dict.GetString("SkillDamage", idx),
			IgnoreMonLevelTxt:              dict.GetNumber("noRatio", idx) > 0,
			CanBlockWithoutShield:          dict.GetNumber("NoShldBlock", idx) > 0,
			ChanceToBlockNormal:            dict.GetNumber("ToBlock", idx),
			ChanceToBlockNightmare:         dict.GetNumber("ToBlock(N)", idx),
			ChanceToBlockHell:              dict.GetNumber("ToBlock(H)", idx),
			ChanceDeadlyStrike:             dict.GetNumber("Crit", idx),
			MinHPNormal:                    dict.GetNumber("minHP", idx),
			MinHPNightmare:                 dict.GetNumber("MinHP(N)", idx),
			MinHPHell:                      dict.GetNumber("MinHP(H)", idx),
			MaxHPNormal:                    dict.GetNumber("maxHP", idx),
			MaxHPNightmare:                 dict.GetNumber("MaxHP(N)", idx),
			MaxHPHell:                      dict.GetNumber("MaxHP(H)", idx),
			ArmorClassNormal:               dict.GetNumber("AC", idx),
			ArmorClassNightmare:            dict.GetNumber("AC(N)", idx),
			ArmorClassHell:                 dict.GetNumber("AC(H)", idx),
			ExperienceNormal:               dict.GetNumber("Exp", idx),
			ExperienceNightmare:            dict.GetNumber("Exp(N)", idx),
			ExperienceHell:                 dict.GetNumber("Exp(H)", idx),
			DamageMinA1Normal:              dict.GetNumber("A1MinD", idx),
			DamageMinA1Nightmare:           dict.GetNumber("A1MinD(N)", idx),
			DamageMinA1Hell:                dict.GetNumber("A1MinD(H)", idx),
			DamageMaxA1Normal:              dict.GetNumber("A1MaxD", idx),
			DamageMaxA1Nightmare:           dict.GetNumber("A1MaxD(N)", idx),
			DamageMaxA1Hell:                dict.GetNumber("A1MaxD(H)", idx),
			DamageMinA2Normal:              dict.GetNumber("A2MinD", idx),
			DamageMinA2Nightmare:           dict.GetNumber("A2MinD(N)", idx),
			DamageMinA2Hell:                dict.GetNumber("A2MinD(H)", idx),
			DamageMaxA2Normal:              dict.GetNumber("A2MaxD", idx),
			DamageMaxA2Nightmare:           dict.GetNumber("A2MaxD(N)", idx),
			DamageMaxA2Hell:                dict.GetNumber("A2MaxD(H)", idx),
			DamageMinS1Normal:              dict.GetNumber("S1MinD", idx),
			DamageMinS1Nightmare:           dict.GetNumber("S1MinD(N)", idx),
			DamageMinS1Hell:                dict.GetNumber("S1MinD(H)", idx),
			DamageMaxS1Normal:              dict.GetNumber("S1MaxD", idx),
			DamageMaxS1Nightmare:           dict.GetNumber("S1MaxD(N)", idx),
			DamageMaxS1Hell:                dict.GetNumber("S1MaxD(H)", idx),
			AttackRatingA1Normal:           dict.GetNumber("A1TH", idx),
			AttackRatingA1Nightmare:        dict.GetNumber("A1TH(N)", idx),
			AttackRatingA1Hell:             dict.GetNumber("A1TH(H)", idx),
			AttackRatingA2Normal:           dict.GetNumber("A2TH", idx),
			AttackRatingA2Nightmare:        dict.GetNumber("A2TH(N)", idx),
			AttackRatingA2Hell:             dict.GetNumber("A2TH(H)", idx),
			AttackRatingS1Normal:           dict.GetNumber("S1TH", idx),
			AttackRatingS1Nightmare:        dict.GetNumber("S1TH(N)", idx),
			AttackRatingS1Hell:             dict.GetNumber("S1TH(H)", idx),
			ElementAttackMode1:             dict.GetString("El1Mode", idx),
			ElementAttackMode2:             dict.GetString("El2Mode", idx),
			ElementAttackMode3:             dict.GetString("El3Mode", idx),
			ElementType1:                   dict.GetString("El1Type", idx),
			ElementType2:                   dict.GetString("El2Type", idx),
			ElementType3:                   dict.GetString("El3Type", idx),
			ElementChance1Normal:           dict.GetNumber("El1Pct", idx),
			ElementChance1Nightmare:        dict.GetNumber("El1Pct(N)", idx),
			ElementChance1Hell:             dict.GetNumber("El1Pct(H)", idx),
			ElementChance2Normal:           dict.GetNumber("El2Pct", idx),
			ElementChance2Nightmare:        dict.GetNumber("El2Pct(N)", idx),
			ElementChance2Hell:             dict.GetNumber("El2Pct(H)", idx),
			ElementChance3Normal:           dict.GetNumber("El3Pct", idx),
			ElementChance3Nightmare:        dict.GetNumber("El3Pct(N)", idx),
			ElementChance3Hell:             dict.GetNumber("El3Pct(H)", idx),
			ElementDamageMin1Normal:        dict.GetNumber("El1MinD", idx),
			ElementDamageMin1Nightmare:     dict.GetNumber("El1MinD(N)", idx),
			ElementDamageMin1Hell:          dict.GetNumber("El1MinD(H)", idx),
			ElementDamageMin2Normal:        dict.GetNumber("El2MinD", idx),
			ElementDamageMin2Nightmare:     dict.GetNumber("El2MinD(N)", idx),
			ElementDamageMin2Hell:          dict.GetNumber("El2MinD(H)", idx),
			ElementDamageMin3Normal:        dict.GetNumber("El3MinD", idx),
			ElementDamageMin3Nightmare:     dict.GetNumber("El3MinD(N)", idx),
			ElementDamageMin3Hell:          dict.GetNumber("El3MinD(H)", idx),
			ElementDamageMax1Normal:        dict.GetNumber("El1MaxD", idx),
			ElementDamageMax1Nightmare:     dict.GetNumber("El1MaxD(N)", idx),
			ElementDamageMax1Hell:          dict.GetNumber("El1MaxD(H)", idx),
			ElementDamageMax2Normal:        dict.GetNumber("El2MaxD", idx),
			ElementDamageMax2Nightmare:     dict.GetNumber("El2MaxD(N)", idx),
			ElementDamageMax2Hell:          dict.GetNumber("El2MaxD(H)", idx),
			ElementDamageMax3Normal:        dict.GetNumber("El3MaxD", idx),
			ElementDamageMax3Nightmare:     dict.GetNumber("El3MaxD(N)", idx),
			ElementDamageMax3Hell:          dict.GetNumber("El3MaxD(H)", idx),
			ElementDuration1Normal:         dict.GetNumber("El1Dur", idx),
			ElementDuration1Nightmare:      dict.GetNumber("El1Dur(N)", idx),
			ElementDuration1Hell:           dict.GetNumber("El1Dur(H)", idx),
			ElementDuration2Normal:         dict.GetNumber("El2Dur", idx),
			ElementDuration2Nightmare:      dict.GetNumber("El2Dur(N)", idx),
			ElementDuration2Hell:           dict.GetNumber("El2Dur(H)", idx),
			ElementDuration3Normal:         dict.GetNumber("El3Dur", idx),
			ElementDuration3Nightmare:      dict.GetNumber("El3Dur(N)", idx),
			ElementDuration3Hell:           dict.GetNumber("El3Dur(H)", idx),
			TreasureClassNormal:            dict.GetString("TreasureClass1", idx),
			TreasureClassNightmare:         dict.GetString("TreasureClass1(N)", idx),
			TreasureClassHell:              dict.GetString("TreasureClass1(H)", idx),
			TreasureClassChampionNormal:    dict.GetString("TreasureClass2", idx),
			TreasureClassChampionNightmare: dict.GetString("TreasureClass2(N)", idx),
			TreasureClassChampionHell:      dict.GetString("TreasureClass2(H)", idx),
			TreasureClass3UniqueNormal:     dict.GetString("TreasureClass3", idx),
			TreasureClass3UniqueNightmare:  dict.GetString("TreasureClass3(N)", idx),
			TreasureClass3UniqueHell:       dict.GetString("TreasureClass3(H)", idx),
			TreasureClassQuestNormal:       dict.GetString("TreasureClass4", idx),
			TreasureClassQuestNightmare:    dict.GetString("TreasureClass4(N)", idx),
			TreasureClassQuestHell:         dict.GetString("TreasureClass4(H)", idx),
			TreasureClassQuestTriggerId:    dict.GetString("TCQuestId", idx),
			TreasureClassQuestComleteId:    dict.GetString("TCQuestCP", idx),
			SpecialEndDeath:                dict.GetNumber("SplEndDeath", idx),
			SpecialGetModeChart:            dict.GetNumber("SplGetModeChart", idx) > 0,
			SpecialEndGeneric:              dict.GetNumber("SplEndGeneric", idx) > 0,
			SpecialClientEnd:               dict.GetNumber("SplClientEnd", idx) > 0,
		}
		MonStats[record.Key] = record
	}

	log.Printf("Loaded %d MonStats records", len(MonStats))
}
