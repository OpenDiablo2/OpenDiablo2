package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type LevelDetailsRecord struct {

	// This column has no function, it only serves as a comment field to make it
	// easier to identify the Level name
	Name string // Name <-- the corresponding column name in the txt

	// Level ID (used in columns like VIS0-7)
	Id int // Id

	// Act Palette . Reference only
	Palette int // Pal

	// The Act the Level is located in (internal enumeration ranges from 0 to 4)
	Act int // Act

	// Used the first one in Classic games and the latter in Expansion games ,
	// they set a questflag. If this flag is set, a character must have
	// completed the quest associated with the flag to take a town portal to
	// the area in question. A character can always use a portal to get back to
	// town.
	QuestFlag          int // QuestFlag
	QuestFlagExpansion int // QuestFlagEx

	// Each layer is an unique ID. This number is used to store each automap on
	// a character. This is used by the game to remember what level the automap
	// are for.
	// NOTE: you need to use the extended levels plugin to be able to add
	// additional layers.
	AutomapIndex int // Layer

	// sizeX - SizeY in each difficuly. If this is a preset area this sets the
	// X size for the area. Othervise use the same value here that are used in
	// lvlprest.txt to set the size for the .ds1 file.
	SizeXNormal int // SizeX
	SizeYNormal int // SizeY

	SizeXNightmare int // SizeX(N)
	SizeYNightmare int // SizeY(N)

	SizeXHell int // SizeX(H)
	SizeYHell int // SizeY(H)

	// They set the X\Y position in the world space
	WorldOffsetX int // OffsetX
	WorldOffsetY int // OffsetY

	// This set what level id's are the Depended level.
	// Example: Monastery uses this field to place its entrance always at same
	// location.
	DependantLevelID int // Depend

	// Controls if teleport is allowed in that level.
	// 0 = Teleport not allowed
	// 1 = Teleport allowed
	// 2 = Teleport allowed, but not able to use teleport throu walls/objects
	// (maybe for objects this is controlled by IsDoor column in objects.txt)
	TeleportFlag d2enum.TeleportFlag // Teleport

	// It sets whether rain or snow (in act 5 only) can fall . Set it to 1 in
	// order to enable it, 0 to disable it.
	EnableRain bool // Rain

	// Unused setting (In pre beta D2 Blizzard planned Rain to generate Mud
	// which would have slowed your character's speed down, but this never made
	// it into the final game). the field is read by the code but the return
	// value is never utilized.
	EnableMud bool // Mud

	// Setting for 3D Enhanced D2 that disables Perspective Mode for a specific
	// level. A value of 1 enables the users to choose between normal and
	// Perspective view, while 0 disables that choice.
	EnablePerspective bool // NoPer

	// Allows you to look through objects and walls even if they are not in a
	// wilderness level. 1 enables it, 0 disables it.
	EnableLineOfSightDraw bool // LOSDraw

	// Unknown. Probably has to do with Tiles and their Placement.
	// 1 enables it, 0 disables it.
	EnableFloorFliter bool // FloorFilter

	// Unknown. Probably has to do with tiles and their placement.
	// 1 enables it, 0 disables it.
	// TODO: needs a better name
	EnableBlankScreen bool // BlankScreen

	// for levels bordered with mountains or walls, like the act 1 wildernesses.
	// 1 enables it, 0 disables it.
	EnableDrawEdges bool // DrawEdges

	// Setting it to 1 makes the level to be treated as an indoor area, while
	// 0 makes this level an outdoor. Indoor areas are not affected by day-night
	// cycles, because they always use the light values specified in Intensity,
	// Red, Green, Blue. this field also controls whenever sounds will echo if
	// you're running the game with a sound card capable of it and have
	// enviroment sound effects set to true.
	IsInside bool // IsInside

	// Setting for Level Generation: You have 3 possibilities here:
	// 1 Random Maze
	// 2 Preset Area
	// 3 Wilderness level
	LevelGenerationType d2enum.LevelGenerationType // DrlgType

	// The type of the Level (Id from lvltypes.txt)
	LevelType int // LevelType

	// NOTE
	// IDs from LvlSub.txt, which is used to randomize outdoor areas, such as
	// spawning ponds in the blood moor and more stones in the Stoney Field.
	// This is all changeable, the other subcolumns are explained in this post.

	// Setting Regarding the level sub-type.
	// Example: 6=wilderness, 9=desert etc, -1=no subtype.
	SubType int // SubType

	// TODO this may need an enumeration.. ?
	// Tells which subtheme a wilderness area should use.
	// Themes ranges from -1 (no subtheme) to 4.
	SubTheme int // SubTheme

	// Setting Regarding Waypoints
	// NOTE: it does NOT control waypoint placement.
	SubWaypoint int // SubWaypoint

	// Setting Regarding Shrines.
	// NOTE: it does NOT control which Shrine will spawn.
	SubShrine int // SubShrine

	// These fields allow linking level serverside, allowing you to travel
	// through areas. The Vis must be filled in with the LevelID your level is
	// linked with, but the actuall number of Vis ( 0 - 7 ) is determined by
	// your actual map (the .ds1 fle).
	// Example: Normally Cave levels are only using vis 0-3 and wilderness areas 4-7 .
	LevelLinkId0 int // Vis0
	LevelLinkId1 int // Vis1
	LevelLinkId2 int // Vis2
	LevelLinkId3 int // Vis3
	LevelLinkId4 int // Vis4
	LevelLinkId5 int // Vis5
	LevelLinkId6 int // Vis6
	LevelLinkId7 int // Vis7

	// This controls the visual graphics then you move the mouse pointer over
	// an entrance. To show the graphics you use an ID from lvlwarp.txt and the
	// behavior on the graphics is controlled by lvlwarp.txt. Your Warps must
	// match your Vis.
	// Example: If your level uses Vis 3,5,7 then you must also use Warp 3,5,7 .
	WarpGraphicsId0 int // Warp0
	WarpGraphicsId1 int // Warp1
	WarpGraphicsId2 int // Warp2
	WarpGraphicsId3 int // Warp3
	WarpGraphicsId4 int // Warp4
	WarpGraphicsId5 int // Warp5
	WarpGraphicsId6 int // Warp6
	WarpGraphicsId7 int // Warp7

	// These settings handle the light intensity as well as its RGB components
	LightIntensity int // Intensity
	Red            int // Red
	Green          int // Green
	Blue           int // Blue

	// This field is required for some levels, entering those levels when portal
	// field isn't set will often crash the game. This also applies to
	// duplicates of those levels created with both of the extended level
	// plugins.
	PortalEnable bool // Portal
	// TODO: this field needs a better name

	// This controls if you can re-position a portal in a level or not. If it's
	// set to 1 you will be able to reposition the portal by using either map
	// entry#76 Tp Location #79. If both tiles are in the level it will use Tp
	// Location #79. If set to 0 the map won't allow repositioning.
	PortalRepositionEnable bool // Position

	// Setting this field to 1 will make the monsters status saved in the map.
	// Setting it to 0 will allow some useful things like NPC refreshing their
	// stores.
	// WARNING: Do not set this to 1 for non-town areas, or the monsters you'll
	// flee from will simply vanish and never reappear. They won't even be
	// replaced by new ones
	// Gravestench - this funcionality should not be in one field
	SaveMonsterStates  bool // SaveMonsters
	SaveMerchantStates bool // SaveMonsters

	// What quest is this level related to. This is the quest id (as example the
	// first quest Den of Evil are set to 1, since its the first quest).
	QuestId int // Quest

	// This sets the minimum distance from a VisX or WarpX location that a
	// monster, object or tile can be spawned at. (also applies to waypoints and
	// some preset portals).
	WarpClearanceDistance int // WarpDist

	//  Area Level on Normal-Nightmare-Hell in Classic and Expansion.
	// It controls the item level of items that drop from chests etc.
	MonsterLevelNormal      int // MonLvl1
	MonsterLevelNightmare   int // MonLvl2
	MonsterLevelHell        int // MonLvl3
	MonsterLevelNormalEx    int // MonLvl1Ex
	MonsterLevelNightmareEx int // MonLvl2Ex
	MonsterLevelHellEx      int // MonLvl3Ex

	// This is a chance in 100000ths that a monster pack will spawn on a tile.
	// The maximum chance the game allows is 10% (aka 10000) in v1.10+,
	MonsterDensityNormal    int // MonDen
	MonsterDensityNightmare int // MonDen(N)
	MonsterDensityHell      int // MonDen(H)

	// Minimum - Maximum Unique and Champion Monsters Spawned in this Level.
	// Whenever any spawn at all however is bound to MonDen.
	MonsterUniqueMinNormal    int // MonUMin
	MonsterUniqueMinNightmare int // MonUMin(N)
	MonsterUniqueMinHell      int // MonUMin(H)

	MonsterUniqueMaxNormal    int // MonUMax
	MonsterUniqueMaxNightmare int // MonUMax(N)
	MonsterUniqueMaxHell      int // MonUMax(H)

	// No info on the PK page, but I'm guessing it's for monster wandering
	MonsterWanderEnable bool // MonWndr

	// This setting is hardcoded to certain level Ids, like the River Of Flame,
	// enabling it in other places can glitch up the game, so leave it alone.
	// It is not known what exactly it does however.
	MonsterSpecialWalk bool // MonSpcWalk

	// Number of different Monster Types that will be present in this area, the
	// maximum is 13. You can have up to 13 different monster types at a time in
	// Nightmare and Hell difficulties, selected randomly from nmon1-nmon25. In
	// Normal difficulty you can have up to 13 normal monster types selected
	// randomly from mon1-mon25, and the same number of champion and unique
	// types selected randomly from umon1-umon25.
	NumMonsterTypes int // NumMon

	// mon1-mon25 work in Normal difficulty, while nmon1-nmon25 in Nightmare and
	// Hell. They tell the game which monster ID taken from MonStats.txt.
	// NOTE: you need to manually add from mon11 to mon25 and from nmon11 to
	// nmon25 !
	MonsterId1Normal  string // mon1
	MonsterId2Normal  string // mon2
	MonsterId3Normal  string // mon3
	MonsterId4Normal  string // mon4
	MonsterId5Normal  string // mon5
	MonsterId6Normal  string // mon6
	MonsterId7Normal  string // mon7
	MonsterId8Normal  string // mon8
	MonsterId9Normal  string // mon9
	MonsterId10Normal string // mon10

	MonsterId1Nightmare  string // nmon1
	MonsterId2Nightmare  string // nmon2
	MonsterId3Nightmare  string // nmon3
	MonsterId4Nightmare  string // nmon4
	MonsterId5Nightmare  string // nmon5
	MonsterId6Nightmare  string // nmon6
	MonsterId7Nightmare  string // nmon7
	MonsterId8Nightmare  string // nmon8
	MonsterId9Nightmare  string // nmon9
	MonsterId10Nightmare string // nmon10

	// Gravestench - adding additional fields for Hell, original txt combined
	// the nighmare and hell ID's stringo the same field
	MonsterId1Hell  string // nmon1
	MonsterId2Hell  string // nmon2
	MonsterId3Hell  string // nmon3
	MonsterId4Hell  string // nmon4
	MonsterId5Hell  string // nmon5
	MonsterId6Hell  string // nmon6
	MonsterId7Hell  string // nmon7
	MonsterId8Hell  string // nmon8
	MonsterId9Hell  string // nmon9
	MonsterId10Hell string // nmon10

	//  Give preference to monsters set to ranged=1 in MonStats.txt on Nightmare
	// and Hell difficulties when picking something to spawn.
	MonsterPreferRanged bool // rangedspawn

	// Works only in normal and it tells which ID will be used for Champion and
	// Random Uniques. The ID is taken from MonStats.txtOnly the first ten
	// columns appear in the unmodded file. In 1.10 final, beta 1.10s and
	// v1.11+ you can add the missing umon11-umon25 columns.
	// NOTE: you can allow umon1-25 to also work in Nightmare and Hell by
	// following this simple ASM edit
	// (https://d2mods.info/forum/viewtopic.php?f=8&t=53969&p=425179&hilit=umon#p425179)
	MonsterUniqueId1  string // umon1
	MonsterUniqueId2  string // umon2
	MonsterUniqueId3  string // umon3
	MonsterUniqueId4  string // umon4
	MonsterUniqueId5  string // umon5
	MonsterUniqueId6  string // umon6
	MonsterUniqueId7  string // umon7
	MonsterUniqueId8  string // umon8
	MonsterUniqueId9  string // umon9
	MonsterUniqueId10 string // umon10

	// Critter Species 1-4. Uses the Id from monstats2.txt and only monsters
	// with critter column set to 1 can spawn here. critter column is also found
	// in monstats2.txt. Critters are in reality only present clientside.
	MonsterCritterId1 string // cmon1
	MonsterCritterId2 string // cmon2
	MonsterCritterId3 string // cmon3
	MonsterCritterId4 string // cmon4

	// Controls the chance for a critter to spawn.
	MonsterCritter1SpawnChance int // cpct1
	MonsterCritter2SpawnChance int // cpct2
	MonsterCritter3SpawnChance int // cpct3
	MonsterCritter4SpawnChance int // cpct4

	// Unknown. These columns are bugged, as the game overrides the contents of
	// columns 3-4 with the value from column 1 when it compiles the bin files.
	// camt1
	// camt2
	// camt3
	// camt4

	// Unknown. It states which theme is used by the area and this field is
	// accessed by the code but it is not exactly known what it does.
	// Themes

	// Referes to a entry in SoundEnviron.txt (for the Levels Music)
	SoundEnvironmentId int // SoundEnv

	// 255 means no Waipoint for this level, while others state the Waypoint' ID
	// for the level
	// NOTE: you can switch waypoint destinations between areas this way, not
	// between acts however so don't even bother to try.
	WaypointId int // Waypoint

	// String Code for the Display name of the Level
	LevelDisplayName string // LevelName

	LevelWarpName string // LevelWarp

	// Which *.DC6 Title Image is loaded when you enter this area. this file
	// MUST exist, otherwise you will crash with an exception when you enter the
	// level (for all levels below the expansion row, the files must be
	// present in the expension folders)
	TitleImageName string // EntryFile

	// this field uses the ID of the ObjectGroup you want to Spawn in this Area,
	// taken from Objgroup.txt.
	ObjectGroupId0 int // ObjGrp0
	ObjectGroupId1 int // ObjGrp1
	ObjectGroupId2 int // ObjGrp2
	ObjectGroupId3 int // ObjGrp3
	ObjectGroupId4 int // ObjGrp4
	ObjectGroupId5 int // ObjGrp5
	ObjectGroupId6 int // ObjGrp6
	ObjectGroupId7 int // ObjGrp7

	// These fields indicates the chance for each object group to spawn (if you
	// use ObjGrp0 then set ObjPrb0 to a value below 100)
	ObjectGroupSpawnChance0 int // ObjPrb0
	ObjectGroupSpawnChance1 int // ObjPrb1
	ObjectGroupSpawnChance2 int // ObjPrb2
	ObjectGroupSpawnChance3 int // ObjPrb3
	ObjectGroupSpawnChance4 int // ObjPrb4
	ObjectGroupSpawnChance5 int // ObjPrb5
	ObjectGroupSpawnChance6 int // ObjPrb6
	ObjectGroupSpawnChance7 int // ObjPrb7

	// Reference Only (can be used for comments)
	// Beta
}

var LevelDetails map[int]*LevelDetailsRecord

func LoadLevelDetails(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	LevelDetails = make(map[int]*LevelDetailsRecord, numRecords)

	for idx := range dict.Data {
		record := &LevelDetailsRecord{
			Name:                       dict.GetString("Name ", idx),
			Id:                         dict.GetNumber("Id", idx),
			Palette:                    dict.GetNumber("Pal", idx),
			Act:                        dict.GetNumber("Act", idx),
			QuestFlag:                  dict.GetNumber("QuestFlag", idx),
			QuestFlagExpansion:         dict.GetNumber("QuestFlagEx", idx),
			AutomapIndex:               dict.GetNumber("Layer", idx),
			SizeXNormal:                dict.GetNumber("SizeX", idx),
			SizeYNormal:                dict.GetNumber("SizeY", idx),
			SizeXNightmare:             dict.GetNumber("SizeX(N)", idx),
			SizeYNightmare:             dict.GetNumber("SizeY(N)", idx),
			SizeXHell:                  dict.GetNumber("SizeX(H)", idx),
			SizeYHell:                  dict.GetNumber("SizeY(H)", idx),
			WorldOffsetX:               dict.GetNumber("OffsetX", idx),
			WorldOffsetY:               dict.GetNumber("OffsetY", idx),
			DependantLevelID:           dict.GetNumber("Depend", idx),
			TeleportFlag:               d2enum.TeleportFlag(dict.GetNumber("Teleport", idx)),
			EnableRain:                 dict.GetNumber("Rain", idx) > 0,
			EnableMud:                  dict.GetNumber("Mud", idx) > 0,
			EnablePerspective:          dict.GetNumber("NoPer", idx) > 0,
			EnableLineOfSightDraw:      dict.GetNumber("LOSDraw", idx) > 0,
			EnableFloorFliter:          dict.GetNumber("FloorFilter", idx) > 0,
			EnableBlankScreen:          dict.GetNumber("BlankScreen", idx) > 0,
			EnableDrawEdges:            dict.GetNumber("DrawEdges", idx) > 0,
			IsInside:                   dict.GetNumber("IsInside", idx) > 0,
			LevelGenerationType:        d2enum.LevelGenerationType(dict.GetNumber("DrlgType", idx)),
			LevelType:                  dict.GetNumber("LevelType", idx),
			SubType:                    dict.GetNumber("SubType", idx),
			SubTheme:                   dict.GetNumber("SubTheme", idx),
			SubWaypoint:                dict.GetNumber("SubWaypoint", idx),
			SubShrine:                  dict.GetNumber("SubShrine", idx),
			LevelLinkId0:               dict.GetNumber("Vis0", idx),
			LevelLinkId1:               dict.GetNumber("Vis1", idx),
			LevelLinkId2:               dict.GetNumber("Vis2", idx),
			LevelLinkId3:               dict.GetNumber("Vis3", idx),
			LevelLinkId4:               dict.GetNumber("Vis4", idx),
			LevelLinkId5:               dict.GetNumber("Vis5", idx),
			LevelLinkId6:               dict.GetNumber("Vis6", idx),
			LevelLinkId7:               dict.GetNumber("Vis7", idx),
			WarpGraphicsId0:            dict.GetNumber("Warp0", idx),
			WarpGraphicsId1:            dict.GetNumber("Warp1", idx),
			WarpGraphicsId2:            dict.GetNumber("Warp2", idx),
			WarpGraphicsId3:            dict.GetNumber("Warp3", idx),
			WarpGraphicsId4:            dict.GetNumber("Warp4", idx),
			WarpGraphicsId5:            dict.GetNumber("Warp5", idx),
			WarpGraphicsId6:            dict.GetNumber("Warp6", idx),
			WarpGraphicsId7:            dict.GetNumber("Warp7", idx),
			LightIntensity:             dict.GetNumber("Intensity", idx),
			Red:                        dict.GetNumber("Red", idx),
			Green:                      dict.GetNumber("Green", idx),
			Blue:                       dict.GetNumber("Blue", idx),
			PortalEnable:               dict.GetNumber("Portal", idx) > 0,
			PortalRepositionEnable:     dict.GetNumber("Position", idx) > 0,
			SaveMonsterStates:          dict.GetNumber("SaveMonsters", idx) > 0,
			SaveMerchantStates:         dict.GetNumber("SaveMonsters", idx) > 0,
			QuestId:                    dict.GetNumber("Quest", idx),
			WarpClearanceDistance:      dict.GetNumber("WarpDist", idx),
			MonsterLevelNormal:         dict.GetNumber("MonLvl1", idx),
			MonsterLevelNightmare:      dict.GetNumber("MonLvl2", idx),
			MonsterLevelHell:           dict.GetNumber("MonLvl3", idx),
			MonsterLevelNormalEx:       dict.GetNumber("MonLvl1Ex", idx),
			MonsterLevelNightmareEx:    dict.GetNumber("MonLvl2Ex", idx),
			MonsterLevelHellEx:         dict.GetNumber("MonLvl3Ex", idx),
			MonsterDensityNormal:       dict.GetNumber("MonDen", idx),
			MonsterDensityNightmare:    dict.GetNumber("MonDen(N)", idx),
			MonsterDensityHell:         dict.GetNumber("MonDen(H)", idx),
			MonsterUniqueMinNormal:     dict.GetNumber("MonUMin", idx),
			MonsterUniqueMinNightmare:  dict.GetNumber("MonUMin(N)", idx),
			MonsterUniqueMinHell:       dict.GetNumber("MonUMin(H)", idx),
			MonsterUniqueMaxNormal:     dict.GetNumber("MonUMax", idx),
			MonsterUniqueMaxNightmare:  dict.GetNumber("MonUMax(N)", idx),
			MonsterUniqueMaxHell:       dict.GetNumber("MonUMax(H)", idx),
			MonsterWanderEnable:        dict.GetNumber("MonWndr", idx) > 0,
			MonsterSpecialWalk:         dict.GetNumber("MonSpcWalk", idx) > 0,
			NumMonsterTypes:            dict.GetNumber("NumMon", idx),
			MonsterId1Normal:           dict.GetString("mon1", idx),
			MonsterId2Normal:           dict.GetString("mon2", idx),
			MonsterId3Normal:           dict.GetString("mon3", idx),
			MonsterId4Normal:           dict.GetString("mon4", idx),
			MonsterId5Normal:           dict.GetString("mon5", idx),
			MonsterId6Normal:           dict.GetString("mon6", idx),
			MonsterId7Normal:           dict.GetString("mon7", idx),
			MonsterId8Normal:           dict.GetString("mon8", idx),
			MonsterId9Normal:           dict.GetString("mon9", idx),
			MonsterId10Normal:          dict.GetString("mon10", idx),
			MonsterId1Nightmare:        dict.GetString("nmon1", idx),
			MonsterId2Nightmare:        dict.GetString("nmon2", idx),
			MonsterId3Nightmare:        dict.GetString("nmon3", idx),
			MonsterId4Nightmare:        dict.GetString("nmon4", idx),
			MonsterId5Nightmare:        dict.GetString("nmon5", idx),
			MonsterId6Nightmare:        dict.GetString("nmon6", idx),
			MonsterId7Nightmare:        dict.GetString("nmon7", idx),
			MonsterId8Nightmare:        dict.GetString("nmon8", idx),
			MonsterId9Nightmare:        dict.GetString("nmon9", idx),
			MonsterId10Nightmare:       dict.GetString("nmon10", idx),
			MonsterId1Hell:             dict.GetString("nmon1", idx),
			MonsterId2Hell:             dict.GetString("nmon2", idx),
			MonsterId3Hell:             dict.GetString("nmon3", idx),
			MonsterId4Hell:             dict.GetString("nmon4", idx),
			MonsterId5Hell:             dict.GetString("nmon5", idx),
			MonsterId6Hell:             dict.GetString("nmon6", idx),
			MonsterId7Hell:             dict.GetString("nmon7", idx),
			MonsterId8Hell:             dict.GetString("nmon8", idx),
			MonsterId9Hell:             dict.GetString("nmon9", idx),
			MonsterId10Hell:            dict.GetString("nmon10", idx),
			MonsterPreferRanged:        dict.GetNumber("rangedspawn", idx) > 0,
			MonsterUniqueId1:           dict.GetString("umon1", idx),
			MonsterUniqueId2:           dict.GetString("umon2", idx),
			MonsterUniqueId3:           dict.GetString("umon3", idx),
			MonsterUniqueId4:           dict.GetString("umon4", idx),
			MonsterUniqueId5:           dict.GetString("umon5", idx),
			MonsterUniqueId6:           dict.GetString("umon6", idx),
			MonsterUniqueId7:           dict.GetString("umon7", idx),
			MonsterUniqueId8:           dict.GetString("umon8", idx),
			MonsterUniqueId9:           dict.GetString("umon9", idx),
			MonsterUniqueId10:          dict.GetString("umon10", idx),
			MonsterCritterId1:          dict.GetString("cmon1", idx),
			MonsterCritterId2:          dict.GetString("cmon2", idx),
			MonsterCritterId3:          dict.GetString("cmon3", idx),
			MonsterCritterId4:          dict.GetString("cmon4", idx),
			MonsterCritter1SpawnChance: dict.GetNumber("cpct1", idx),
			MonsterCritter2SpawnChance: dict.GetNumber("cpct2", idx),
			MonsterCritter3SpawnChance: dict.GetNumber("cpct3", idx),
			MonsterCritter4SpawnChance: dict.GetNumber("cpct4", idx),
			SoundEnvironmentId:         dict.GetNumber("SoundEnv", idx),
			WaypointId:                 dict.GetNumber("Waypoint", idx),
			LevelDisplayName:           dict.GetString("LevelName", idx),
			LevelWarpName:              dict.GetString("LevelWarp", idx),
			TitleImageName:             dict.GetString("EntryFile", idx),
			ObjectGroupId0:             dict.GetNumber("ObjGrp0", idx),
			ObjectGroupId1:             dict.GetNumber("ObjGrp1", idx),
			ObjectGroupId2:             dict.GetNumber("ObjGrp2", idx),
			ObjectGroupId3:             dict.GetNumber("ObjGrp3", idx),
			ObjectGroupId4:             dict.GetNumber("ObjGrp4", idx),
			ObjectGroupId5:             dict.GetNumber("ObjGrp5", idx),
			ObjectGroupId6:             dict.GetNumber("ObjGrp6", idx),
			ObjectGroupId7:             dict.GetNumber("ObjGrp7", idx),
			ObjectGroupSpawnChance0:    dict.GetNumber("ObjPrb0", idx),
			ObjectGroupSpawnChance1:    dict.GetNumber("ObjPrb1", idx),
			ObjectGroupSpawnChance2:    dict.GetNumber("ObjPrb2", idx),
			ObjectGroupSpawnChance3:    dict.GetNumber("ObjPrb3", idx),
			ObjectGroupSpawnChance4:    dict.GetNumber("ObjPrb4", idx),
			ObjectGroupSpawnChance5:    dict.GetNumber("ObjPrb5", idx),
			ObjectGroupSpawnChance6:    dict.GetNumber("ObjPrb6", idx),
			ObjectGroupSpawnChance7:    dict.GetNumber("ObjPrb7", idx),
		}
		LevelDetails[idx] = record
	}
	log.Printf("Loaded %d LevelDetails records", len(LevelDetails))
}
