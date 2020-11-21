package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	logPrefix = "Record Manager"
)

// NewRecordManager creates a new record manager (no loaders are bound!)
func NewRecordManager(l d2util.LogLevel) (*RecordManager, error) {
	rm := &RecordManager{
		boundLoaders: make(map[string][]recordLoader),
	}

	rm.Logger = d2util.NewLogger()
	rm.Logger.SetPrefix(logPrefix)
	rm.Logger.SetLevel(l)

	err := rm.init()
	if err != nil {
		return nil, err
	}

	return rm, nil
}

// RecordManager stores all of the records loaded from txt files
type RecordManager struct {
	*d2util.Logger
	boundLoaders map[string][]recordLoader // there can be more than one loader bound for a file
	Animation    struct {
		Data  d2data.AnimationData
		Token struct {
			Player    PlayerTypes
			Composite CompositeTypes
			Armor     ArmorTypes
			Weapon    WeaponClasses
			HitClass  HitClasses
		}
	}
	BodyLocations
	Calculation struct {
		Skills   Calculations
		Missiles Calculations
	}
	Character struct {
		Classes PlayerClasses
		Events
		Experience ExperienceBreakpoints
		MaxLevel   ExperienceMaxLevels
		Modes      PlayerModes
		Stats      CharStats
	}
	ComponentCodes
	Colors
	DifficultyLevels
	ElemTypes
	Gamble
	Hireling struct {
		Details      Hirelings
		Descriptions HirelingDescriptions
	}
	Item struct {
		All CommonItems // NOTE: populated when armor, weapons, and misc items are ALL loaded

		Armors  CommonItems
		Misc    CommonItems
		Weapons CommonItems

		Equivalency         ItemEquivalenceMap      // NOTE: populated when all items are loaded
		EquivalenceByRecord ItemEquivalenceByRecord // NOTE: populated when all items are loaded

		AutoMagic
		Belts
		Books
		Gems
		Magic struct {
			Prefix MagicPrefix
			Suffix MagicSuffix
		}
		MagicPrefixGroups  ItemAffixGroups
		MagicSuffixGroups  ItemAffixGroups
		Quality            ItemQualities
		LowQualityPrefixes LowQualities
		Rare               struct {
			Prefix RarePrefixes
			Suffix RareSuffixes
		}
		Ratios ItemRatios
		Cube   struct {
			Recipes   CubeRecipes
			Modifiers CubeModifiers
			Types     CubeTypes
		}
		Runewords
		Sets
		SetItems
		Stats    ItemStatCosts
		Treasure struct {
			Normal    TreasureClass
			Expansion TreasureClass
		}
		Types  ItemTypes
		Unique UniqueItems
		StorePages
	}
	Layout struct {
		Inventory
		Overlays
	}
	Level struct {
		AutoMaps
		Details LevelDetails
		Maze    LevelMazeDetails
		Presets LevelPresets
		Sub     LevelSubstitutions
		Types   LevelTypes
		Warp    LevelWarps
	}
	Missiles
	missilesByName
	Monster struct {
		AI        MonsterAI
		Equipment MonsterEquipment
		Levels    MonsterLevels
		Modes     MonModes
		Name      struct {
			Prefix UniqueMonsterAffixes
			Suffix UniqueMonsterAffixes
		}
		Placements MonsterPlacements
		Presets    MonPresets
		Props      MonsterProperties
		Sequences  MonsterSequences
		Sounds     MonsterSounds
		Stats      MonStats
		Stats2     MonStats2
		Types      MonsterTypes
		Unique     struct {
			Appellations UniqueAppellations
			Mods         MonsterUniqueModifiers
			Constants    MonsterUniqueModifierConstants
			Super        SuperUniques
		}
	}
	NPCs
	Object struct {
		Details ObjectDetails
		Lookup  IndexedObjects
		Modes   ObjectModes
		Shrines
		Types ObjectTypes
	}
	PetTypes
	Properties
	Skill struct {
		Details      SkillDetails
		Descriptions SkillDescriptions
	}
	Sound struct {
		Details     SoundDetails
		Environment SoundEnvironments
	}
	States
}

func (r *RecordManager) init() error { // nolint:funlen // can't reduce
	loaders := []struct {
		path   string
		loader recordLoader
	}{
		{d2resource.LevelType, levelTypesLoader},
		{d2resource.LevelPreset, levelPresetLoader},
		{d2resource.LevelWarp, levelWarpsLoader},
		{d2resource.ObjectType, objectTypesLoader},
		{d2resource.ObjectDetails, objectDetailsLoader},
		{d2resource.ObjectMode, objectModesLoader},
		{d2resource.Weapons, weaponsLoader},
		{d2resource.Armor, armorLoader},
		{d2resource.Misc, miscItemsLoader},
		{d2resource.Books, booksLoader},
		{d2resource.Belts, beltsLoader},
		{d2resource.Colors, colorsLoader},
		{d2resource.ItemTypes, itemTypesLoader}, // WARN: needs to be after weapons, armor, and misc
		{d2resource.UniqueItems, uniqueItemsLoader},
		{d2resource.Missiles, missilesLoader},
		{d2resource.SoundSettings, soundDetailsLoader},
		{d2resource.MonStats, monsterStatsLoader},
		{d2resource.MonStats2, monsterStats2Loader},
		{d2resource.MonPreset, monsterPresetLoader},
		{d2resource.MonProp, monsterPropertiesLoader},
		{d2resource.MonType, monsterTypesLoader},
		{d2resource.MonMode, monsterModeLoader},
		{d2resource.MagicPrefix, magicPrefixLoader},
		{d2resource.MagicSuffix, magicSuffixLoader},
		{d2resource.ItemStatCost, itemStatCostLoader},
		{d2resource.ItemRatio, itemRatioLoader},
		{d2resource.StorePage, storePagesLoader},
		{d2resource.Overlays, overlaysLoader},
		{d2resource.CharStats, charStatsLoader},
		{d2resource.Gamble, gambleLoader},
		{d2resource.Hireling, hirelingLoader},
		{d2resource.Experience, experienceLoader},
		{d2resource.Gems, gemsLoader},
		{d2resource.QualityItems, itemQualityLoader},
		{d2resource.Runes, runewordLoader},
		{d2resource.DifficultyLevels, difficultyLevelsLoader},
		{d2resource.AutoMap, autoMapLoader},
		{d2resource.LevelDetails, levelDetailsLoader},
		{d2resource.LevelMaze, levelMazeDetailsLoader},
		{d2resource.LevelSubstitutions, levelSubstitutionsLoader},
		{d2resource.CubeRecipes, cubeRecipeLoader},
		{d2resource.SuperUniques, monsterSuperUniqeLoader},
		{d2resource.Inventory, inventoryLoader},
		{d2resource.Skills, skillDetailsLoader},
		{d2resource.SkillCalc, skillCalcLoader},
		{d2resource.MissileCalc, missileCalcLoader},
		{d2resource.Properties, propertyLoader},
		{d2resource.SkillDesc, skillDescriptionLoader},
		{d2resource.BodyLocations, bodyLocationsLoader},
		{d2resource.Sets, setLoader},
		{d2resource.SetItems, setItemLoader},
		{d2resource.AutoMagic, autoMagicLoader},
		{d2resource.TreasureClass, treasureClassLoader},
		{d2resource.TreasureClassEx, treasureClassExLoader},
		{d2resource.States, statesLoader},
		{d2resource.SoundEnvirons, soundEnvironmentLoader},
		{d2resource.Shrines, shrineLoader},
		{d2resource.ElemType, elemTypesLoader},
		{d2resource.PlrMode, playerModesLoader},
		{d2resource.PetType, petTypesLoader},
		{d2resource.NPC, npcLoader},
		{d2resource.MonsterUniqueModifier, monsterUniqModifiersLoader},
		{d2resource.MonsterEquipment, monsterEquipmentLoader},
		{d2resource.UniqueAppellation, uniqueAppellationsLoader},
		{d2resource.MonsterLevel, monsterLevelsLoader},
		{d2resource.MonsterSound, monsterSoundsLoader},
		{d2resource.MonsterSequence, monsterSequencesLoader},
		{d2resource.PlayerClass, playerClassLoader},
		{d2resource.MonsterPlacement, monsterPlacementsLoader},
		{d2resource.ObjectGroup, objectGroupsLoader},
		{d2resource.CompCode, componentCodesLoader},
		{d2resource.MonsterAI, monsterAiLoader},
		{d2resource.RarePrefix, rareItemPrefixLoader},
		{d2resource.RareSuffix, rareItemSuffixLoader},
		{d2resource.Events, eventsLoader},
		{d2resource.ArmorType, armorTypesLoader},      // anim mode tokens
		{d2resource.WeaponClass, weaponClassesLoader}, // anim mode tokens
		{d2resource.PlayerType, playerTypeLoader},     // anim mode tokens
		{d2resource.Composite, compositeTypeLoader},   // anim mode tokens
		{d2resource.HitClass, hitClassLoader},         // anim mode tokens
		{d2resource.UniquePrefix, uniqueMonsterPrefixLoader},
		{d2resource.UniqueSuffix, uniqueMonsterSuffixLoader},
		{d2resource.CubeModifier, cubeModifierLoader},
		{d2resource.CubeType, cubeTypeLoader},
		{d2resource.HirelingDescription, hirelingDescriptionLoader},
		{d2resource.LowQualityItems, lowQualityLoader},
	}

	for idx := range loaders {
		err := r.AddLoader(loaders[idx].path, loaders[idx].loader)
		if err != nil {
			return err
		}
	}

	r.initObjectRecords(objectLookups)

	return nil
}

// AddLoader associates a file path with a record loader
func (r *RecordManager) AddLoader(path string, loader recordLoader) error {
	if _, found := r.boundLoaders[path]; !found {
		r.boundLoaders[path] = make([]recordLoader, 0)
	}

	r.boundLoaders[path] = append(r.boundLoaders[path], loader)

	return nil
}

// Load will pass the dictionary to any bound loaders and populate the record entries
func (r *RecordManager) Load(path string, dict *d2txt.DataDictionary) error {
	loaders, found := r.boundLoaders[path]
	if !found {
		return fmt.Errorf("no loader bound for `%s`", path)
	}

	for idx := range loaders {
		err := loaders[idx](r, dict)
		if err != nil {
			return err
		}
	}

	// as soon as Armor, Weapons, and Misc items are loaded, we merge into r.Item.All
	if r.Item.All == nil && r.Item.Armors != nil && r.Item.Weapons != nil && r.Item.Misc != nil {
		r.Item.All = make(CommonItems)

		for code := range r.Item.Armors {
			r.Item.All[code] = r.Item.Armors[code]
		}

		for code := range r.Item.Weapons {
			r.Item.All[code] = r.Item.Weapons[code]
		}

		for code := range r.Item.Misc {
			r.Item.All[code] = r.Item.Misc[code]
		}
	}

	return nil
}

// GetMaxLevelByHero returns the highest level attainable for a hero type
func (r *RecordManager) GetMaxLevelByHero(heroType d2enum.Hero) int {
	return r.Character.MaxLevel[heroType]
}

// GetExperienceBreakpoint given a hero type and a level, returns the experience required for the level
func (r *RecordManager) GetExperienceBreakpoint(heroType d2enum.Hero, level int) int {
	return r.Character.Experience[level].HeroBreakpoints[heroType]
}

// GetLevelDetails gets a LevelDetailsRecord by the record Id
func (r *RecordManager) GetLevelDetails(id int) *LevelDetailsRecord {
	for i := 0; i < len(r.Level.Details); i++ {
		if r.Level.Details[i].ID == id {
			return r.Level.Details[i]
		}
	}

	return nil
}

// LevelPreset looks up a LevelPresetRecord by ID
func (r *RecordManager) LevelPreset(id int) LevelPresetRecord {
	for i := 0; i < len(r.Level.Presets); i++ {
		if r.Level.Presets[i].DefinitionID == id {
			return r.Level.Presets[i]
		}
	}

	panic("Unknown level preset")
}

// FindEquivalentTypesByItemCommonRecord returns itemtype codes that are equivalent
// to the given item common record
func (r *RecordManager) FindEquivalentTypesByItemCommonRecord(
	icr *ItemCommonRecord,
) []string {
	if r.Item.EquivalenceByRecord == nil {
		r.Item.EquivalenceByRecord = make(map[*ItemCommonRecord][]string)
	}

	// the first lookup generates the lookup table entry, next time will just use the table
	if r.Item.EquivalenceByRecord[icr] == nil {
		r.Item.EquivalenceByRecord[icr] = make([]string, 0)

		for code := range r.Item.Equivalency {
			icrList := r.Item.Equivalency[code]
			for idx := range icrList {
				if icr == icrList[idx] {
					r.Item.EquivalenceByRecord[icr] = append(r.Item.EquivalenceByRecord[icr], code)
					break
				}
			}
		}
	}

	return r.Item.EquivalenceByRecord[icr]
}

func (r *RecordManager) initObjectRecords(lookups []ObjectLookupRecord) {
	// Allocating 6 to allow Acts 1-5 without requiring a -1 at every read.
	records := make(IndexedObjects, 6)

	for i := range lookups {
		record := &lookups[i]
		if records[record.Act] == nil {
			// Likewise allocating 3 so a -1 isn't necessary.
			records[record.Act] = make([][]*ObjectLookupRecord, 3)
		}

		if records[record.Act][record.Type] == nil {
			// For simplicity, allocating with length 1000 then filling the values in by index.
			// If ids in the dictionary ever surpass 1000, raise this number.
			records[record.Act][record.Type] = make([]*ObjectLookupRecord, 1000)
		}

		records[record.Act][record.Type][record.Id] = record
	}

	r.Object.Lookup = records
}

// LookupObject looks up an object record
func (r *RecordManager) LookupObject(act, typ, id int) *ObjectLookupRecord {
	object := r.lookupObject(act, typ, id)
	if object == nil {
		r.Fatalf("Failed to look up object Act: %d, Type: %d, ID: %d", act, typ, id)
	}

	return object
}

func (r *RecordManager) lookupObject(act, typ, id int) *ObjectLookupRecord {
	if len(r.Object.Lookup) < act {
		return nil
	}

	if len(r.Object.Lookup[act]) < typ {
		return nil
	}

	if len(r.Object.Lookup[act][typ]) < id {
		return nil
	}

	return r.Object.Lookup[act][typ][id]
}

// SelectSoundByIndex selects a sound by its ID
func (r *RecordManager) SelectSoundByIndex(index int) *SoundDetailsRecord {
	for idx := range r.Sound.Details {
		if r.Sound.Details[idx].Index == index {
			return r.Sound.Details[idx]
		}
	}

	return nil
}

// GetSkillByName returns the skill record for the given Skill name.
func (r *RecordManager) GetSkillByName(skillName string) *SkillRecord {
	for idx := range r.Skill.Details {
		if r.Skill.Details[idx].Skill == skillName {
			return r.Skill.Details[idx]
		}
	}

	return nil
}

// GetMissileByName allows lookup of a MissileRecord by a given name. The name will be lowercased and stripped of whitespaces.
func (r *RecordManager) GetMissileByName(missileName string) *MissileRecord {
	return r.missilesByName[sanitizeMissilesKey(missileName)]
}
