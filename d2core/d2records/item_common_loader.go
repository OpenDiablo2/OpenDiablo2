package d2records

import (
	"fmt"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// nolint:funlen // cant reduce
func loadCommonItems(d *d2txt.DataDictionary, source d2enum.InventoryItemType) (CommonItems, error) {
	records := make(CommonItems)

	for d.Next() {
		record := &ItemCommonRecord{
			Source: source,

			Name: d.String("name"),

			Version:     d.Number("version"),
			CompactSave: d.Number("compactsave") > 0,
			Rarity:      d.Number("rarity"),
			Spawnable:   d.Number("spawnable") > 0,

			MinAC:            d.Number("minac"),
			MaxAC:            d.Number("maxac"),
			Absorbs:          d.Number("absorbs"),
			Speed:            d.Number("speed"),
			RequiredStrength: d.Number("reqstr"),
			Block:            d.Number("block"),
			Durability:       d.Number("durability"),
			NoDurability:     d.Number("nodurability") > 0,

			Level:         d.Number("level"),
			RequiredLevel: d.Number("levelreq"),
			Cost:          d.Number("cost"),
			GambleCost:    d.Number("gamble cost"),
			Code:          d.String("code"),
			NameString:    d.String("namestr"),
			MagicLevel:    d.Number("magic lvl"),
			AutoPrefix:    d.Number("auto prefix"),

			AlternateGfx: d.String("alternategfx"),
			OpenBetaGfx:  d.String("OpenBetaGfx"),
			NormalCode:   d.String("normcode"),
			UberCode:     d.String("ubercode"),
			UltraCode:    d.String("ultracode"),

			SpellOffset: d.Number("spelloffset"),

			Component:       d.Number("component"),
			InventoryWidth:  d.Number("invwidth"),
			InventoryHeight: d.Number("invheight"),
			HasInventory:    d.Number("hasinv") > 0,
			GemSockets:      d.Number("gemsockets"),
			GemApplyType:    d.Number("gemapplytype"),

			FlippyFile:          d.String("flippyfile"),
			InventoryFile:       d.String("invfile"),
			UniqueInventoryFile: d.String("uniqueinvfile"),
			SetInventoryFile:    d.String("setinvfile"),

			AnimRightArm:         d.Number("rArm"),
			AnimLeftArm:          d.Number("lArm"),
			AnimTorso:            d.Number("Torso"),
			AnimLegs:             d.Number("Legs"),
			AnimRightShoulderPad: d.Number("rSPad"),
			AnimLeftShoulderPad:  d.Number("lSPad"),

			Useable: d.Number("useable") > 0,

			Throwable: d.Number("throwable") > 0,
			Stackable: d.Number("stackable") > 0,
			MinStack:  d.Number("minstack"),
			MaxStack:  d.Number("maxstack"),

			Type:  d.String("type"),
			Type2: d.String("type2"),

			DropSound:    d.String("dropsound"),
			DropSfxFrame: d.Number("dropsfxframe"),
			UseSound:     d.String("usesound"),

			Unique:      d.Number("unique") > 0,
			Transparent: d.Number("transparent") > 0,
			TransTable:  d.Number("transtbl"),
			Quivered:    d.Number("quivered") > 0,
			LightRadius: d.Number("lightradius"),
			Belt:        d.Number("belt") > 0,

			Quest: d.Number("quest"),

			MissileType:       d.Number("missiletype"),
			DurabilityWarning: d.Number("durwarning"),
			QuantityWarning:   d.Number("qntwarning"),

			MinDamage:      d.Number("mindam"),
			MaxDamage:      d.Number("maxdam"),
			StrengthBonus:  d.Number("StrBonus"),
			DexterityBonus: d.Number("DexBonus"),

			GemOffset: d.Number("gemoffset"),
			BitField1: d.Number("bitfield1"),

			Vendors: createItemVendorParams(d),

			SourceArt:               d.String("Source Art"),
			GameArt:                 d.String("Game Art"),
			ColorTransform:          d.Number("Transform"),
			InventoryColorTransform: d.Number("InvTrans"),

			SkipName:         d.Number("SkipName") > 0,
			NightmareUpgrade: d.String("NightmareUpgrade"),
			HellUpgrade:      d.String("HellUpgrade"),

			Nameable: d.Number("Nameable") > 0,

			// weapon params
			BarbOneOrTwoHanded: d.Number("1or2handed") > 0,
			UsesTwoHands:       d.Number("2handed") > 0,
			Min2HandDamage:     d.Number("2handmindam"),
			Max2HandDamage:     d.Number("2handmaxdam"),
			MinMissileDamage:   d.Number("minmisdam"),
			MaxMissileDamage:   d.Number("maxmisdam"),
			MissileSpeed:       d.Number("misspeed"),
			ExtraRange:         d.Number("rangeadder"),

			RequiredDexterity: d.Number("reqdex"),

			WeaponClass:      d.String("wclass"),
			WeaponClass2Hand: d.String("2handedwclass"),

			HitClass:   d.String("hit class"),
			SpawnStack: d.Number("spawnstack"),

			SpecialFeature: d.String("special"),

			QuestDifficultyCheck: d.Number("questdiffcheck") > 0,

			PermStoreItem: d.Number("PermStoreItem") > 0,

			// misc params
			FlavorText: d.String("szFlavorText"),

			Transmogrify: d.Number("Transmogrify") > 0,
			TransmogCode: d.String("TMogType"),
			TransmogMin:  d.Number("TMogMin"),
			TransmogMax:  d.Number("TMogMax"),

			AutoBelt: d.Number("autobelt") > 0,

			SpellIcon:    d.Number("spellicon"),
			SpellType:    d.Number("pSpell"),
			OverlayState: d.String("state"),
			CureOverlayStates: [2]string{
				d.String("cstate1"),
				d.String("cstate2"),
			},
			EffectLength: d.Number("len"),
			UsageStats:   createItemUsageStats(d),

			SpellDescriptionType: d.Number("spelldesc"),
			// 0 = none, 1 = use desc string, 2 = use desc string + calc value
			SpellDescriptionString: d.String("spelldescstr"),
			SpellDescriptionCalc:   d2calculation.CalcString(d.String("spelldesccalc")),

			BetterGem: d.String("BetterGem"),

			Multibuy: d.Number("multibuy") > 0,
		}

		records[record.Code] = record
	}

	if d.Err != nil {
		return nil, d.Err
	}

	return records, nil
}

func createItemVendorParams(d *d2txt.DataDictionary) map[string]*ItemVendorParams {
	vs := []string{
		"Charsi",
		"Gheed",
		"Akara",
		"Fara",
		"Lysander",
		"Drognan",
		"Hralti",
		"Alkor",
		"Ormus",
		"Elzix",
		"Asheara",
		"Cain",
		"Halbu",
		"Jamella",
		"Larzuk",
		"Malah",
		"Drehya",
	}

	result := make(map[string]*ItemVendorParams)

	for _, name := range vs {
		wvp := ItemVendorParams{
			Min:        d.Number(fmt.Sprintf("%s%s", name, "Min")),
			Max:        d.Number(fmt.Sprintf("%s%s", name, "Max")),
			MagicMin:   d.Number(fmt.Sprintf("%s%s", name, "MagicMin")),
			MagicMax:   d.Number(fmt.Sprintf("%s%s", name, "MagicMax")),
			MagicLevel: d.Number(fmt.Sprintf("%s%s", name, "MagicLvl")),
		}
		result[name] = &wvp
	}

	return result
}

func createItemUsageStats(d *d2txt.DataDictionary) [3]ItemUsageStat {
	result := [3]ItemUsageStat{}
	for i := 0; i < 3; i++ {
		result[i].Stat = d.String("stat" + strconv.Itoa(i))
		result[i].Calc = d2calculation.CalcString(d.String("calc" + strconv.Itoa(i)))
	}

	return result
}
