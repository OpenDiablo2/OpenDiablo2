package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2animdata"
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
	Item struct {
		Cube struct {
			Modifiers CubeModifiers
			Types     CubeTypes
			Recipes   CubeRecipes
		}
		Magic struct {
			Prefix MagicPrefix
			Suffix MagicSuffix
		}
		Treasure struct {
			Normal    TreasureClass
			Expansion TreasureClass
		}
		All                 CommonItems
		Equivalency         ItemEquivalenceMap
		EquivalenceByRecord ItemEquivalenceByRecord
		Unique              UniqueItems
		Belts
		Books
		Gems
		Weapons           CommonItems
		MagicPrefixGroups ItemAffixGroups
		MagicSuffixGroups ItemAffixGroups
		Quality           ItemQualities
		Types             ItemTypes
		Misc              CommonItems
		Ratios            ItemRatios
		Armors            CommonItems
		Runewords
		Sets
		SetItems
		Stats ItemStatCosts
		StorePages
		Rare struct {
			Prefix RarePrefixes
			Suffix RareSuffixes
		}
		LowQualityPrefixes LowQualities
		AutoMagic
	}
	Monster struct {
		Unique struct {
			Appellations UniqueAppellations
			Mods         MonsterUniqueModifiers
			Super        SuperUniques
			Constants    MonsterUniqueModifierConstants
		}
		Name struct {
			Prefix UniqueMonsterAffixes
			Suffix UniqueMonsterAffixes
		}
		Props      MonsterProperties
		Modes      MonModes
		Levels     MonsterLevels
		Types      MonsterTypes
		Presets    MonPresets
		Equipment  MonsterEquipment
		Sequences  MonsterSequences
		Sounds     MonsterSounds
		Stats      MonStats
		Stats2     MonStats2
		AI         MonsterAI
		Placements MonsterPlacements
	}
	Level struct {
		Warp    LevelWarps
		Details LevelDetails
		Maze    LevelMazeDetails
		Presets LevelPresets
		Sub     LevelSubstitutions
		Types   LevelTypes
		AutoMaps
	}
	Character struct {
		Classes    PlayerClasses
		Experience ExperienceBreakpoints
		MaxLevel   ExperienceMaxLevels
		Modes      PlayerModes
		Stats      CharStats
		Events
	}
	Animation struct {
		Data  *d2animdata.AnimationData
		Token struct {
			Player    PlayerTypes
			Composite CompositeTypes
			Armor     ArmorTypes
			Weapon    WeaponClasses
			HitClass  HitClasses
		}
	}
	Hireling struct {
		Descriptions HirelingDescriptions
		Details      Hirelings
	}
	Calculation struct {
		Skills   Calculations
		Missiles Calculations
	}
	Skill struct {
		Details      SkillDetails
		Descriptions SkillDescriptions
	}
	Layout struct {
		Inventory
		Overlays
	}
	Sound struct {
		Details     SoundDetails
		Environment SoundEnvironments
	}
	*d2util.Logger
	Gamble
	ElemTypes
	DifficultyLevels
	Colors
	Missiles
	missilesByName
	ComponentCodes
	NPCs
	BodyLocations
	PetTypes
	Properties
	boundLoaders map[ // NOTE: populated when armor, weapons, and misc items are ALL loaded
	// NOTE: populated when all items are loaded
	string][]recordLoader
	States
	Object struct {
		Details ObjectDetails
		Lookup  IndexedObjects
		Modes   ObjectModes
		Shrines
		Types ObjectTypes
	}
}

func (r *RecordManager) init() error { // nolint:funlen // can't reduce
	loaders := []struct {
		loader recordLoader
		path   string
	}{
		{levelTypesLoader, d2resource.LevelType},
		{levelPresetLoader, d2resource.LevelPreset},
		{levelWarpsLoader, d2resource.LevelWarp},
		{objectTypesLoader, d2resource.ObjectType},
		{objectDetailsLoader, d2resource.ObjectDetails},
		{objectModesLoader, d2resource.ObjectMode},
		{weaponsLoader, d2resource.Weapons},
		{armorLoader, d2resource.Armor},
		{miscItemsLoader, d2resource.Misc},
		{booksLoader, d2resource.Books},
		{beltsLoader, d2resource.Belts},
		{colorsLoader, d2resource.Colors},
		{itemTypesLoader, d2resource.ItemTypes}, // WARN: needs to be after weapons, armor, and misc
		{uniqueItemsLoader, d2resource.UniqueItems},
		{missilesLoader, d2resource.Missiles},
		{soundDetailsLoader, d2resource.SoundSettings},
		{monsterStatsLoader, d2resource.MonStats},
		{monsterStats2Loader, d2resource.MonStats2},
		{monsterPresetLoader, d2resource.MonPreset},
		{monsterPropertiesLoader, d2resource.MonProp},
		{monsterTypesLoader, d2resource.MonType},
		{monsterModeLoader, d2resource.MonMode},
		{magicPrefixLoader, d2resource.MagicPrefix},
		{magicSuffixLoader, d2resource.MagicSuffix},
		{itemStatCostLoader, d2resource.ItemStatCost},
		{itemRatioLoader, d2resource.ItemRatio},
		{storePagesLoader, d2resource.StorePage},
		{overlaysLoader, d2resource.Overlays},
		{charStatsLoader, d2resource.CharStats},
		{gambleLoader, d2resource.Gamble},
		{hirelingLoader, d2resource.Hireling},
		{experienceLoader, d2resource.Experience},
		{gemsLoader, d2resource.Gems},
		{itemQualityLoader, d2resource.QualityItems},
		{runewordLoader, d2resource.Runes},
		{difficultyLevelsLoader, d2resource.DifficultyLevels},
		{autoMapLoader, d2resource.AutoMap},
		{levelDetailsLoader, d2resource.LevelDetails},
		{levelMazeDetailsLoader, d2resource.LevelMaze},
		{levelSubstitutionsLoader, d2resource.LevelSubstitutions},
		{cubeRecipeLoader, d2resource.CubeRecipes},
		{monsterSuperUniqeLoader, d2resource.SuperUniques},
		{inventoryLoader, d2resource.Inventory},
		{skillDetailsLoader, d2resource.Skills},
		{skillCalcLoader, d2resource.SkillCalc},
		{missileCalcLoader, d2resource.MissileCalc},
		{propertyLoader, d2resource.Properties},
		{skillDescriptionLoader, d2resource.SkillDesc},
		{bodyLocationsLoader, d2resource.BodyLocations},
		{setLoader, d2resource.Sets},
		{setItemLoader, d2resource.SetItems},
		{autoMagicLoader, d2resource.AutoMagic},
		{treasureClassLoader, d2resource.TreasureClass},
		{treasureClassExLoader, d2resource.TreasureClassEx},
		{statesLoader, d2resource.States},
		{soundEnvironmentLoader, d2resource.SoundEnvirons},
		{shrineLoader, d2resource.Shrines},
		{elemTypesLoader, d2resource.ElemType},
		{playerModesLoader, d2resource.PlrMode},
		{petTypesLoader, d2resource.PetType},
		{npcLoader, d2resource.NPC},
		{monsterUniqModifiersLoader, d2resource.MonsterUniqueModifier},
		{monsterEquipmentLoader, d2resource.MonsterEquipment},
		{uniqueAppellationsLoader, d2resource.UniqueAppellation},
		{monsterLevelsLoader, d2resource.MonsterLevel},
		{monsterSoundsLoader, d2resource.MonsterSound},
		{monsterSequencesLoader, d2resource.MonsterSequence},
		{playerClassLoader, d2resource.PlayerClass},
		{monsterPlacementsLoader, d2resource.MonsterPlacement},
		{objectGroupsLoader, d2resource.ObjectGroup},
		{componentCodesLoader, d2resource.CompCode},
		{monsterAiLoader, d2resource.MonsterAI},
		{rareItemPrefixLoader, d2resource.RarePrefix},
		{rareItemSuffixLoader, d2resource.RareSuffix},
		{eventsLoader, d2resource.Events},
		{armorTypesLoader, d2resource.ArmorType},      // anim mode tokens
		{weaponClassesLoader, d2resource.WeaponClass}, // anim mode tokens
		{playerTypeLoader, d2resource.PlayerType},     // anim mode tokens
		{compositeTypeLoader, d2resource.Composite},   // anim mode tokens
		{hitClassLoader, d2resource.HitClass},         // anim mode tokens
		{uniqueMonsterPrefixLoader, d2resource.UniquePrefix},
		{uniqueMonsterSuffixLoader, d2resource.UniqueSuffix},
		{cubeModifierLoader, d2resource.CubeModifier},
		{cubeTypeLoader, d2resource.CubeType},
		{hirelingDescriptionLoader, d2resource.HirelingDescription},
		{lowQualityLoader, d2resource.LowQualityItems},
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

// GetLevelDetails gets a LevelDetailRecord by the record Id
func (r *RecordManager) GetLevelDetails(id int) *LevelDetailRecord {
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

		records[record.Act][record.Type][record.ID] = record
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
func (r *RecordManager) SelectSoundByIndex(index int) *SoundDetailRecord {
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
