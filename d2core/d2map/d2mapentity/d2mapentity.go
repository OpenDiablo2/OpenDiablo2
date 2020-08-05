package d2mapentity

import (
	"errors"
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
)

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
func NewPlayer(id, name string, x, y, direction int, heroType d2enum.Hero,
	stats *d2hero.HeroStatsState, equipment *d2inventory.CharacterEquipment) *Player {
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

	composite, err := d2asset.LoadComposite(d2enum.ObjectTypePlayer, heroType.GetToken(),
		d2resource.PaletteUnits)
	if err != nil {
		panic(err)
	}

	stats.NextLevelExp = d2datadict.GetExperienceBreakpoint(heroType, stats.Level)
	stats.Stamina = stats.MaxStamina

	result := &Player{
		mapEntity: newMapEntity(x, y),
		composite: composite,
		Equipment: equipment,
		Stats:     stats,
		name:      name,
		Class:     heroType,
		//nameLabel:    d2ui.CreateLabel(d2resource.FontFormal11, d2resource.PaletteStatic),
		isRunToggled: true,
		isInTown:     true,
		isRunning:    true,
	}

	result.mapEntity.uuid = id
	result.SetSpeed(baseRunSpeed)
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
func NewMissile(x, y int, record *d2datadict.MissileRecord) (*Missile, error) {
	animation, err := d2asset.LoadAnimation(
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
func NewItem(x, y int, codes ...string) (*Item, error) {
	item := diablo2item.NewItem(codes...)

	if item == nil {
		return nil, errors.New(errInvalidItemCodes)
	}

	filename := item.CommonRecord().FlippyFile
	filepath := fmt.Sprintf("%s/%s.DC6", d2resource.ItemGraphics, filename)
	animation, err := d2asset.LoadAnimation(filepath, d2resource.PaletteUnits)

	if err != nil {
		return nil, err
	}

	animation.PlayForward()
	animation.SetPlayLoop(false)
	entity := NewAnimatedEntity(x*5, y*5, animation)

	result := &Item{
		AnimatedEntity: entity,
		Item:           item,
	}

	return result, nil
}

// NewNPC creates a new NPC and returns a pointer to it.
func NewNPC(x, y int, monstat *d2datadict.MonStatsRecord, direction int) (*NPC, error) {
	result := &NPC{
		mapEntity:     newMapEntity(x, y),
		HasPaths:      false,
		monstatRecord: monstat,
		monstatEx:     d2datadict.MonStats2[monstat.ExtraDataKey],
	}

	var equipment [16]string

	for compType, opts := range result.monstatEx.EquipmentOptions {
		equipment[compType] = selectEquip(opts)
	}

	composite, _ := d2asset.LoadComposite(d2enum.ObjectTypeCharacter, monstat.AnimationDirectoryToken,
		d2resource.PaletteUnits)
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
		result.name = d2common.TranslateString(result.monstatRecord.NameString)
	}

	return result, nil
}
