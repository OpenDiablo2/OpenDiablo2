package d2mapentity

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

const (
	subtilesPerTile       = 5
	retailFps             = 25.0
	millisecondsPerSecond = 1000.0
)

// NewMapEntityFactory creates a MapEntityFactory instance with the given asset manager
func NewMapEntityFactory(asset *d2asset.AssetManager) (*MapEntityFactory, error) {
	itemFactory, err := diablo2item.NewItemFactory(asset)
	if err != nil {
		return nil, err
	}

	stateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	entityFactory := &MapEntityFactory{
		stateFactory,
		asset,
		itemFactory,
	}

	return entityFactory, nil
}

// MapEntityFactory creates map entities for the MapEngine
type MapEntityFactory struct {
	*d2hero.HeroStateFactory
	asset *d2asset.AssetManager
	item  *diablo2item.ItemFactory
}

// NewAnimatedEntity creates an instance of AnimatedEntity
func NewAnimatedEntity(x, y int, animation d2interface.Animation) *AnimatedEntity {
	entity := &AnimatedEntity{
		mapEntity: newMapEntity(x, y),
		animation: animation,
	}
	entity.mapEntity.directioner = entity.rotate

	return entity
}

// NewPlayer creates a new player entity and returns a pointer to it.
func (f *MapEntityFactory) NewPlayer(id, name string, x, y, direction int, heroType d2enum.Hero,
	stats *d2hero.HeroStatsState, skills map[int]*d2hero.HeroSkill, equipment *d2inventory.CharacterEquipment,
	leftSkill, rightSkill int, gold int) *Player {
	layerEquipment := &[d2enum.CompositeTypeMax]string{
		d2enum.CompositeTypeHead:      equipment.Head.GetArmorClass(),
		d2enum.CompositeTypeTorso:     equipment.Torso.GetArmorClass(),
		d2enum.CompositeTypeLegs:      equipment.Legs.GetArmorClass(),
		d2enum.CompositeTypeRightArm:  equipment.RightArm.GetArmorClass(),
		d2enum.CompositeTypeLeftArm:   equipment.LeftArm.GetArmorClass(),
		d2enum.CompositeTypeRightHand: equipment.RightHand.GetItemCode(),
		d2enum.CompositeTypeLeftHand:  equipment.LeftHand.GetItemCode(),
		d2enum.CompositeTypeShield:    equipment.Shield.GetItemCode(),
	}

	composite, err := f.asset.LoadComposite(d2enum.ObjectTypePlayer, heroType.GetToken(),
		d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	stats.NextLevelExp = f.asset.Records.GetExperienceBreakpoint(heroType, stats.Level)
	stats.Stamina = float64(stats.MaxStamina)

	heroState, _ := f.CreateHeroState(name, heroType, stats)

	result := &Player{
		mapEntity:  newMapEntity(x, y),
		composite:  composite,
		Equipment:  equipment,
		Stats:      heroState.Stats,
		Skills:     heroState.Skills,
		LeftSkill:  heroState.Skills[leftSkill],
		RightSkill: heroState.Skills[rightSkill],
		name:       name,
		Class:      heroType,
		//nameLabel:    d2ui.NewLabel(d2resource.FontFormal11, d2resource.PaletteStatic),
		isRunToggled: false,
		isInTown:     true,
		isRunning:    false,
		Gold:         gold,
		Act:          1,
	}

	result.mapEntity.uuid = id
	result.SetSpeed(baseWalkSpeed)
	result.mapEntity.directioner = result.rotate
	err = composite.SetMode(d2enum.PlayerAnimationModeTownNeutral, equipment.RightHand.GetWeaponClass())

	if err != nil {
		panic(err)
	}

	composite.SetDirection(direction)

	if err := composite.Equip(layerEquipment); err != nil {
		fmt.Printf("failed to equip, err: %v\n", err)
	}

	return result
}

// NewMissile creates a new Missile and initializes it's animation.
func (f *MapEntityFactory) NewMissile(x, y int, record *d2records.MissileRecord) (*Missile, error) {
	animation, err := f.asset.LoadAnimation(
		fmt.Sprintf("%s/%s.dcc", d2resource.MissileData, record.Animation.CelFileName),
		d2resource.PaletteUnits,
	)
	if err != nil {
		return nil, err
	}

	if record.Animation.HasSubLoop {
		animation.SetSubLoop(record.Animation.SubStartingFrame, record.Animation.SubEndingFrame)
	}

	animation.SetEffect(d2enum.DrawEffectModulate)
	animation.SetPlayLoop(record.Animation.LoopAnimation)
	animation.PlayForward()
	entity := NewAnimatedEntity(x, y, animation)

	result := &Missile{
		AnimatedEntity: entity,
		record:         record,
	}
	result.Speed = float64(record.Velocity)

	return result, nil
}

// NewItem creates an item map entity
func (f *MapEntityFactory) NewItem(x, y int, codes ...string) (*Item, error) {
	item, err := f.item.NewItem(codes...)

	if err != nil {
		return nil, err
	}

	filename := item.CommonRecord().FlippyFile
	filepath := fmt.Sprintf("%s/%s.DC6", d2resource.ItemGraphics, filename)
	animation, err := f.asset.LoadAnimation(filepath, d2resource.PaletteUnits)

	if err != nil {
		return nil, err
	}

	animation.PlayForward()
	animation.SetPlayLoop(false)
	entity := NewAnimatedEntity(x*subtilesPerTile, y*subtilesPerTile, animation)

	result := &Item{
		AnimatedEntity: entity,
		Item:           item,
	}

	return result, nil
}

// NewNPC creates a new NPC and returns a pointer to it.
func (f *MapEntityFactory) NewNPC(x, y int, monstat *d2records.MonStatsRecord, direction int) (*NPC, error) {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/803
	result := &NPC{
		mapEntity:     newMapEntity(x, y),
		HasPaths:      false,
		monstatRecord: monstat,
		monstatEx:     f.asset.Records.Monster.Stats2[monstat.ExtraDataKey],
	}

	var equipment [16]string

	for compType, opts := range result.monstatEx.EquipmentOptions {
		equipment[compType] = selectEquip(opts)
	}

	composite, err := f.asset.LoadComposite(d2enum.ObjectTypeCharacter, monstat.AnimationDirectoryToken,
		d2resource.PaletteUnits)

	if err != nil {
		return nil, err
	}

	result.composite = composite

	if err := composite.SetMode(d2enum.MonsterAnimationModeNeutral,
		result.monstatEx.BaseWeaponClass); err != nil {
		return nil, err
	}

	if err := composite.Equip(&equipment); err != nil {
		return nil, err
	}

	result.SetSpeed(float64(monstat.SpeedBase))
	result.mapEntity.directioner = result.rotate

	result.composite.SetDirection(direction)

	if result.monstatRecord != nil && result.monstatRecord.IsInteractable {
		result.name = f.asset.TranslateString(result.monstatRecord.NameString)
	}

	return result, nil
}

// NewCastOverlay creates a cast overlay map entity
func (f *MapEntityFactory) NewCastOverlay(x, y int, overlayRecord *d2records.OverlayRecord) (*CastOverlay, error) {
	animation, err := f.asset.LoadAnimationWithEffect(
		fmt.Sprintf("/data/Global/Overlays/%s.dcc", overlayRecord.Filename),
		d2resource.PaletteUnits,
		d2enum.DrawEffectModulate,
	)

	if err != nil {
		return nil, err
	}

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/767
	animation.Rewind()
	animation.ResetPlayedCount()

	animationSpeed := float64(overlayRecord.AnimRate*retailFps) / millisecondsPerSecond
	playLoop := false // https://github.com/OpenDiablo2/OpenDiablo2/issues/804

	animation.SetPlayLength(animationSpeed)
	animation.SetPlayLoop(playLoop)
	animation.PlayForward()

	targetX := x + overlayRecord.XOffset
	targetY := y + overlayRecord.YOffset

	entity := NewAnimatedEntity(targetX, targetY, animation)

	result := &CastOverlay{
		AnimatedEntity: entity,
		record:         overlayRecord,
		playLoop:       playLoop,
	}

	return result, nil
}

// NewObject creates an instance of AnimatedComposite
func (f *MapEntityFactory) NewObject(x, y int, objectRec *d2records.ObjectDetailsRecord,
	palettePath string) (*Object, error) {
	locX, locY := float64(x), float64(y)
	entity := &Object{
		uuid:         uuid.New().String(),
		objectRecord: objectRec,
		Position:     d2vector.NewPosition(locX, locY),
		name:         f.asset.TranslateString(objectRec.Name),
	}
	objectType := f.asset.Records.Object.Types[objectRec.Index]

	composite, err := f.asset.LoadComposite(d2enum.ObjectTypeItem, objectType.Token,
		palettePath)
	if err != nil {
		return nil, err
	}

	entity.composite = composite

	err = entity.setMode(d2enum.ObjectAnimationModeNeutral, 0, false)
	if err != nil {
		return nil, err
	}

	_, err = initObject(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
