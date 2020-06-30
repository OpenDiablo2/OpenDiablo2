package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// LevelDetailsRecord is a representation of a row from levels.txt
// it describes lots of things about the levels, like where they are connected,
// what kinds of monsters spawn, the level generator type, and lots of other stuff.
type LevelDetailsRecord struct {

	// Name
	// This column has no function, it only serves as a comment field to make it
	// easier to identify the Level name
	Name string // Name <-- the corresponding column name in the txt

	// mon1-mon25 work in Normal difficulty, while nmon1-nmon25 in Nightmare and
	// Hell. They tell the game which monster ID taken from MonStats.txt.
	// NOTE: you need to manually add from mon11 to mon25 and from nmon11 to
	// nmon25 !
	MonsterID1Normal  string // mon1
	MonsterID2Normal  string // mon2
	MonsterID3Normal  string // mon3
	MonsterID4Normal  string // mon4
	MonsterID5Normal  string // mon5
	MonsterID6Normal  string // mon6
	MonsterID7Normal  string // mon7
	MonsterID8Normal  string // mon8
	MonsterID9Normal  string // mon9
	MonsterID10Normal string // mon10

	MonsterID1Nightmare  string // nmon1
	MonsterID2Nightmare  string // nmon2
	MonsterID3Nightmare  string // nmon3
	MonsterID4Nightmare  string // nmon4
	MonsterID5Nightmare  string // nmon5
	MonsterID6Nightmare  string // nmon6
	MonsterID7Nightmare  string // nmon7
	MonsterID8Nightmare  string // nmon8
	MonsterID9Nightmare  string // nmon9
	MonsterID10Nightmare string // nmon10

	// Gravestench - adding additional fields for Hell, original txt combined
	// the nighmare and hell ID's stringo the same field
	MonsterID1Hell  string // nmon1
	MonsterID2Hell  string // nmon2
	MonsterID3Hell  string // nmon3
	MonsterID4Hell  string // nmon4
	MonsterID5Hell  string // nmon5
	MonsterID6Hell  string // nmon6
	MonsterID7Hell  string // nmon7
	MonsterID8Hell  string // nmon8
	MonsterID9Hell  string // nmon9
	MonsterID10Hell string // nmon10

	// Works only in normal and it tells which ID will be used for Champion and
	// Random Uniques. The ID is taken from MonStats.txtOnly the first ten
	// columns appear in the unmodded file. In 1.10 final, beta 1.10s and
	// v1.11+ you can add the missing umon11-umon25 columns.
	// NOTE: you can allow umon1-25 to also work in Nightmare and Hell by
	// following this simple ASM edit
	// (https://d2mods.info/forum/viewtopic.php?f=8&t=53969&p=425179&hilit=umon#p425179)
	MonsterUniqueID1  string // umon1
	MonsterUniqueID2  string // umon2
	MonsterUniqueID3  string // umon3
	MonsterUniqueID4  string // umon4
	MonsterUniqueID5  string // umon5
	MonsterUniqueID6  string // umon6
	MonsterUniqueID7  string // umon7
	MonsterUniqueID8  string // umon8
	MonsterUniqueID9  string // umon9
	MonsterUniqueID10 string // umon10

	// Critter Species 1-4. Uses the Id from monstats2.txt and only monsters
	// with critter column set to 1 can spawn here. critter column is also found
	// in monstats2.txt. Critters are in reality only present clientside.
	MonsterCritterID1 string // cmon1
	MonsterCritterID2 string // cmon2
	MonsterCritterID3 string // cmon3
	MonsterCritterID4 string // cmon4

	// String Code for the Display name of the Level
	LevelDisplayName string // LevelName

	LevelWarpName string // LevelWarp

	// Which *.DC6 Title Image is loaded when you enter this area. this file
	// MUST exist, otherwise you will crash with an exception when you enter the
	// level (for all levels below the expansion row, the files must be
	// present in the expension folders)
	TitleImageName string // EntryFile

	// Id
	// Level ID (used in columns like VIS0-7)
	Id int //nolint:golint Id is the right key

	// Palette is the Act Palette . Reference only
	Palette int // Pal

	// Act that the Level is located in (internal enumeration ranges from 0 to 4)
	Act int // Act

	// QuestFlag, QuestExpansionFlag
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

	// SizeXNormal -- SizeYHell If this is a preset area this sets the
	// X size for the area. Othervise use the same value here that are used in
	// lvlprest.txt to set the size for the .ds1 file.
	SizeXNormal    int // SizeX
	SizeYNormal    int // SizeY
	SizeXNightmare int // SizeX(N)
	SizeYNightmare int // SizeY(N)
	SizeXHell      int // SizeX(H)
	SizeYHell      int // SizeY(H)

	// They set the X\Y position in the world space
	WorldOffsetX int // OffsetX
	WorldOffsetY int // OffsetY

	// This set what level id's are the Depended level.
	// Example: Monastery uses this field to place its entrance always at same
	// location.
	DependantLevelID int // Depend

	// The type of the Level (Id from lvltypes.txt)
	LevelType int // LevelType

	// Controls if teleport is allowed in that level.
	// 0 = Teleport not allowed
	// 1 = Teleport allowed
	// 2 = Teleport allowed, but not able to use teleport throu walls/objects
	// (maybe for objects this is controlled by IsDoor column in objects.txt)
	TeleportFlag d2enum.TeleportFlag // Teleport

	// Setting for Level Generation: You have 3 possibilities here:
	// 1 Random Maze
	// 2 Preset Area
	// 3 Wilderness level
	LevelGenerationType d2enum.LevelGenerationType // DrlgType

	// NOTE
	// IDs from LvlSub.txt, which is used to randomize outdoor areas, such as
	// spawning ponds in the blood moor and more stones in the Stoney Field.
	// This is all changeable, the other subcolumns are explained in this post.

	// Setting Regarding the level sub-type.
	// Example: 6=wilderness, 9=desert etc, -1=no subtype.
	SubType int // SubType

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
	LevelLinkID0 int // Vis0
	LevelLinkID1 int // Vis1
	LevelLinkID2 int // Vis2
	LevelLinkID3 int // Vis3
	LevelLinkID4 int // Vis4
	LevelLinkID5 int // Vis5
	LevelLinkID6 int // Vis6
	LevelLinkID7 int // Vis7

	// This controls the visual graphics then you move the mouse pointer over
	// an entrance. To show the graphics you use an ID from lvlwarp.txt and the
	// behavior on the graphics is controlled by lvlwarp.txt. Your Warps must
	// match your Vis.
	// Example: If your level uses Vis 3,5,7 then you must also use Warp 3,5,7 .
	WarpGraphicsID0 int // Warp0
	WarpGraphicsID1 int // Warp1
	WarpGraphicsID2 int // Warp2
	WarpGraphicsID3 int // Warp3
	WarpGraphicsID4 int // Warp4
	WarpGraphicsID5 int // Warp5
	WarpGraphicsID6 int // Warp6
	WarpGraphicsID7 int // Warp7

	// These settings handle the light intensity as well as its RGB components
	LightIntensity int // Intensity
	Red            int // Red
	Green          int // Green
	Blue           int // Blue

	// What quest is this level related to. This is the quest id (as example the
	// first quest Den of Evil are set to 1, since its the first quest).
	QuestID int // Quest

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

	// Number of different Monster Types that will be present in this area, the
	// maximum is 13. You can have up to 13 different monster types at a time in
	// Nightmare and Hell difficulties, selected randomly from nmon1-nmon25. In
	// Normal difficulty you can have up to 13 normal monster types selected
	// randomly from mon1-mon25, and the same number of champion and unique
	// types selected randomly from umon1-umon25.
	NumMonsterTypes int // NumMon

	// Controls the chance for a critter to spawn.
	MonsterCritter1SpawnChance int // cpct1
	MonsterCritter2SpawnChance int // cpct2
	MonsterCritter3SpawnChance int // cpct3
	MonsterCritter4SpawnChance int // cpct4

	// Referes to a entry in SoundEnviron.txt (for the Levels Music)
	SoundEnvironmentID int // SoundEnv

	// 255 means no Waipoint for this level, while others state the Waypoint' ID
	// for the level
	// NOTE: you can switch waypoint destinations between areas this way, not
	// between acts however so don't even bother to try.
	WaypointID int // Waypoint

	// this field uses the ID of the ObjectGroup you want to Spawn in this Area,
	// taken from Objgroup.txt.
	ObjectGroupID0 int // ObjGrp0
	ObjectGroupID1 int // ObjGrp1
	ObjectGroupID2 int // ObjGrp2
	ObjectGroupID3 int // ObjGrp3
	ObjectGroupID4 int // ObjGrp4
	ObjectGroupID5 int // ObjGrp5
	ObjectGroupID6 int // ObjGrp6
	ObjectGroupID7 int // ObjGrp7

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
	EnableBlankScreen bool // BlankScreen

	// for levels bordered with mountains or walls, like the act 1 wildernesses.
	// 1 enables it, 0 disables it.
	EnableDrawEdges bool // DrawEdges

	// Setting it to 1 makes the level to be treated as an indoor area, while
	// 0 makes this level an outdoor. Indoor areas are not affected by day-night
	// cycles, because they always use the light values specified in Intensity,
	// Red, Green, Blue. this field also controls whenever sounds will echo if
	// you're running the game with a sound card capable of it and have
	// environment sound effects set to true.
	IsInside bool // IsInside

	// This field is required for some levels, entering those levels when portal
	// field isn't set will often crash the game. This also applies to
	// duplicates of those levels created with both of the extended level
	// plugins.
	PortalEnable bool // Portal

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

	// No info on the PK page, but I'm guessing it's for monster wandering
	MonsterWanderEnable bool // MonWndr

	// This setting is hardcoded to certain level Ids, like the River Of Flame,
	// enabling it in other places can glitch up the game, so leave it alone.
	// It is not known what exactly it does however.
	MonsterSpecialWalk bool // MonSpcWalk

	//  Give preference to monsters set to ranged=1 in MonStats.txt on Nightmare
	// and Hell difficulties when picking something to spawn.
	MonsterPreferRanged bool // rangedspawn

}

// LevelDetails has all of the LevelDetailsRecords
//nolint:gochecknoglobals // Currently global by design, only written once
var LevelDetails map[int]*LevelDetailsRecord

// GetLevelDetails gets a LevelDetailsRecord by the record Id
func GetLevelDetails(id int) *LevelDetailsRecord {
	for i := 0; i < len(LevelDetails); i++ {
		if LevelDetails[i].Id == id {
			return LevelDetails[i]
		}
	}

	return nil
}

// LoadLevelDetails loads level details records from levels.txt
//nolint:funlen // Txt loader, makes no sense to split
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
			LevelLinkID0:               dict.GetNumber("Vis0", idx),
			LevelLinkID1:               dict.GetNumber("Vis1", idx),
			LevelLinkID2:               dict.GetNumber("Vis2", idx),
			LevelLinkID3:               dict.GetNumber("Vis3", idx),
			LevelLinkID4:               dict.GetNumber("Vis4", idx),
			LevelLinkID5:               dict.GetNumber("Vis5", idx),
			LevelLinkID6:               dict.GetNumber("Vis6", idx),
			LevelLinkID7:               dict.GetNumber("Vis7", idx),
			WarpGraphicsID0:            dict.GetNumber("Warp0", idx),
			WarpGraphicsID1:            dict.GetNumber("Warp1", idx),
			WarpGraphicsID2:            dict.GetNumber("Warp2", idx),
			WarpGraphicsID3:            dict.GetNumber("Warp3", idx),
			WarpGraphicsID4:            dict.GetNumber("Warp4", idx),
			WarpGraphicsID5:            dict.GetNumber("Warp5", idx),
			WarpGraphicsID6:            dict.GetNumber("Warp6", idx),
			WarpGraphicsID7:            dict.GetNumber("Warp7", idx),
			LightIntensity:             dict.GetNumber("Intensity", idx),
			Red:                        dict.GetNumber("Red", idx),
			Green:                      dict.GetNumber("Green", idx),
			Blue:                       dict.GetNumber("Blue", idx),
			PortalEnable:               dict.GetNumber("Portal", idx) > 0,
			PortalRepositionEnable:     dict.GetNumber("Position", idx) > 0,
			SaveMonsterStates:          dict.GetNumber("SaveMonsters", idx) > 0,
			SaveMerchantStates:         dict.GetNumber("SaveMonsters", idx) > 0,
			QuestID:                    dict.GetNumber("Quest", idx),
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
			MonsterID1Normal:           dict.GetString("mon1", idx),
			MonsterID2Normal:           dict.GetString("mon2", idx),
			MonsterID3Normal:           dict.GetString("mon3", idx),
			MonsterID4Normal:           dict.GetString("mon4", idx),
			MonsterID5Normal:           dict.GetString("mon5", idx),
			MonsterID6Normal:           dict.GetString("mon6", idx),
			MonsterID7Normal:           dict.GetString("mon7", idx),
			MonsterID8Normal:           dict.GetString("mon8", idx),
			MonsterID9Normal:           dict.GetString("mon9", idx),
			MonsterID10Normal:          dict.GetString("mon10", idx),
			MonsterID1Nightmare:        dict.GetString("nmon1", idx),
			MonsterID2Nightmare:        dict.GetString("nmon2", idx),
			MonsterID3Nightmare:        dict.GetString("nmon3", idx),
			MonsterID4Nightmare:        dict.GetString("nmon4", idx),
			MonsterID5Nightmare:        dict.GetString("nmon5", idx),
			MonsterID6Nightmare:        dict.GetString("nmon6", idx),
			MonsterID7Nightmare:        dict.GetString("nmon7", idx),
			MonsterID8Nightmare:        dict.GetString("nmon8", idx),
			MonsterID9Nightmare:        dict.GetString("nmon9", idx),
			MonsterID10Nightmare:       dict.GetString("nmon10", idx),
			MonsterID1Hell:             dict.GetString("nmon1", idx),
			MonsterID2Hell:             dict.GetString("nmon2", idx),
			MonsterID3Hell:             dict.GetString("nmon3", idx),
			MonsterID4Hell:             dict.GetString("nmon4", idx),
			MonsterID5Hell:             dict.GetString("nmon5", idx),
			MonsterID6Hell:             dict.GetString("nmon6", idx),
			MonsterID7Hell:             dict.GetString("nmon7", idx),
			MonsterID8Hell:             dict.GetString("nmon8", idx),
			MonsterID9Hell:             dict.GetString("nmon9", idx),
			MonsterID10Hell:            dict.GetString("nmon10", idx),
			MonsterPreferRanged:        dict.GetNumber("rangedspawn", idx) > 0,
			MonsterUniqueID1:           dict.GetString("umon1", idx),
			MonsterUniqueID2:           dict.GetString("umon2", idx),
			MonsterUniqueID3:           dict.GetString("umon3", idx),
			MonsterUniqueID4:           dict.GetString("umon4", idx),
			MonsterUniqueID5:           dict.GetString("umon5", idx),
			MonsterUniqueID6:           dict.GetString("umon6", idx),
			MonsterUniqueID7:           dict.GetString("umon7", idx),
			MonsterUniqueID8:           dict.GetString("umon8", idx),
			MonsterUniqueID9:           dict.GetString("umon9", idx),
			MonsterUniqueID10:          dict.GetString("umon10", idx),
			MonsterCritterID1:          dict.GetString("cmon1", idx),
			MonsterCritterID2:          dict.GetString("cmon2", idx),
			MonsterCritterID3:          dict.GetString("cmon3", idx),
			MonsterCritterID4:          dict.GetString("cmon4", idx),
			MonsterCritter1SpawnChance: dict.GetNumber("cpct1", idx),
			MonsterCritter2SpawnChance: dict.GetNumber("cpct2", idx),
			MonsterCritter3SpawnChance: dict.GetNumber("cpct3", idx),
			MonsterCritter4SpawnChance: dict.GetNumber("cpct4", idx),
			SoundEnvironmentID:         dict.GetNumber("SoundEnv", idx),
			WaypointID:                 dict.GetNumber("Waypoint", idx),
			LevelDisplayName:           dict.GetString("LevelName", idx),
			LevelWarpName:              dict.GetString("LevelWarp", idx),
			TitleImageName:             dict.GetString("EntryFile", idx),
			ObjectGroupID0:             dict.GetNumber("ObjGrp0", idx),
			ObjectGroupID1:             dict.GetNumber("ObjGrp1", idx),
			ObjectGroupID2:             dict.GetNumber("ObjGrp2", idx),
			ObjectGroupID3:             dict.GetNumber("ObjGrp3", idx),
			ObjectGroupID4:             dict.GetNumber("ObjGrp4", idx),
			ObjectGroupID5:             dict.GetNumber("ObjGrp5", idx),
			ObjectGroupID6:             dict.GetNumber("ObjGrp6", idx),
			ObjectGroupID7:             dict.GetNumber("ObjGrp7", idx),
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
