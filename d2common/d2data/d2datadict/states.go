package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// StateRecord describes a body location that items can be equipped to
type StateRecord struct {
	// Name of status effect (Line # is used for ID purposes)
	State string

	// Exact usage depends on the state and how the code accesses it,
	// overlay1 however is the one you should generally be using.
	Overlay1 string
	Overlay2 string
	Overlay3 string
	Overlay4 string

	// Overlay shown on target of progressive skill when chargeup triggers.
	PgOverlay string

	// Overlay displayed when the state is applied initially
	// (later replaced by overlay1 or whatever applicable by code).
	CastOverlay string

	// Like castoverlay, just this one is displayed when the state expires.
	RemOverlay string

	// Primary stat associated with the state, mostly used for display purposes
	// (you should generally use skills.txt for assigning stats to states).
	Stat string

	// The missile that this state will utilize for certain events,
	// how this is used depends entirely on the functions associated with the state.
	Missile string

	// The skill that will be queried for this state in some sections of code,
	// strangely enough this contradicts the fact states store their assigner skill anyway
	// (via STAT_MODIFIERLIST_SKILL)
	Skill string

	// What item type is effected by this states color change.
	ItemType string

	// The color being applied to this item
	// (only going to have an effect on alternate gfx, inventory gfx isn't effected).
	ItemTrans string

	// Sound played respectively when the state is initially applied
	OnSound string

	// and when it expires
	OffSound string

	// States can be grouped together, so they cannot stack
	Group int

	// Clientside callback function invoked when the state is applied initially.
	SetFunc int

	// Clientside callback function invoked when the state expires or is otherwise removed.
	RemFunc int

	// The color priority for this states color change, the, this can range from 0 to 255,
	// the state with the highest color priority will always be used should more then
	// one co-exist on the unit. If two states exist with the same priority the one with the
	// lowest id is used (IIRC).
	ColorPri int

	// Index for the color shift palette picked from the *.PL2 files.
	ColorShift int

	// Change the color of the light radius to what is indicated here,
	// (only has an effect in D3D and glide of course).
	LightR int

	// Change the color of the light radius to what is indicated here,
	// (only has an effect in D3D and glide of course).
	LightG int

	// Change the color of the light radius to what is indicated here,
	// (only has an effect in D3D and glide of course).
	LightB int

	// What unit type is used for the disguise gfx
	// (1 being monsters, 2 being players, contrary to internal game logic).
	GfxType int

	// The unit class used for disguise gfx, this corresponds with the index
	// from monstats.txt and charstats.txt
	GfxClass int

	// When 'gfxtype' is set to 1, the "class" represents an hcIdx from MonStats.txt.
	// If it's set to 2 then it will indicate a character class the unit with this state will be
	// morphed into.

	// Clientside event callback for this state
	// (likely very inconsistent with the server side events, beware).
	CltEvent string

	// Callback function invoked when the client event triggers.
	CltEventFunc int

	// CltDoFunc called every frame the state is active
	CltActiveFunc int

	// Srvdofunc called every frame the state is active
	SrvActiveFunc int

	// If a monster gets hit, the state will be dispelled
	RemHit bool

	// Not yet analyzed in detail
	NoSend bool

	// Whenever a state transforms the appearance of a unit
	Transform bool

	// Aura states will stack on-screen. Aura states are dispelled when a monster is
	// affected by conversion
	Aura bool

	// Can a heal enabled npc remove this state when you talk to them?
	Cureable bool

	// Curse states can't stack. Controls duration reduction from cleansing, and curse resistance.
	// When a new curse is applied, the old one is removed.
	Curse bool

	// State has a StateActiveFunc associated with it. The active func is called every frame while
	// the state is active.
	Active bool

	// State restricts skill usage (druid shapeshift)
	Restrict bool

	// State makes the game load another sprite (use with Transform)
	Disguise bool

	// State applies a color change that overrides all others
	Blue bool

	// Control when attack rating is displayed in blue
	AttBlue bool

	// Control when damage is displayed in blue
	DmgBlue bool

	// Control when armor class is displayed in blue
	ArmBlue bool

	// Control when fire resistance is displayed in blue
	RfBlue bool

	// Control when lightning resistance is displayed in blue
	RlBlue bool

	// Control when cold resistance is displayed in blue
	RcBlue bool

	// Control when poison resistance is displayed in blue
	RpBlue bool

	// Control when attack rating is displayed in red
	AttRed bool

	// Control when damage is displayed in red
	DmgRed bool

	// Control when armor class is displayed in red
	ArmRed bool

	// Control when fire resistance is displayed in red
	RfRed bool

	// Control when lightning resistance is displayed in red
	RlRed bool

	// Control when cold resistance is displayed in red
	RcRed bool

	// Control when poison resistance is displayed in red
	RpRed bool

	// Control when stamina bar color is changed to blue
	StamBarBlue bool

	// When a unit effected by this state kills another unit,
	// the summon owner will receive experience
	Exp bool

	// Whenever the state is removed when the player dies
	PlrStayDeath bool

	// Whenever the state is removed when the monster dies
	MonStayDeath bool

	// Whenever the state is removed when the boss dies. Prevents bosses from shattering?
	BossStayDeath bool

	// When the unit dies, the corpse and death animation will not be drawn
	Hide bool

	// Whenever the unit shatters or explodes when it dies. This is heavily hardcoded,
	// it will always use the ice shatter for all states other than STATE_UBERMINION
	Shatter bool

	// Whenever this state prevents the corpse from being selected by spells and the ai
	UDead bool

	// When a state with this is active, it cancels out the native life regen of monsters.
	// (using only the mod part instead of accr).
	Life bool

	// Whenever this state applies a color change that overrides all others (such as from items).
	// (see blue column, which seams to do the same).
	Green bool

	// Whenever this state is associated with progressive spells and will be
	// looked up when the charges are triggered.
	Pgsv bool

	// Related to assigning overlays to the unit, not extensively researched yet.
	NoOverlays bool

	// Like the previous column, also only used on states with the previous column enabled.
	NoClear bool

	// whenever this state will use the minion owners inventory clientside (this is what makes
	// the decoy always show up with your own equipment,
	// even when you change what you wear after summoning one).
	BossInv bool

	// Prevents druids that wield a bow or crossbow while shape shifted from
	// firing missiles, and will rather attack in melee.
	MeleeOnly bool

	// Not researched yet
	NotOnDead bool
}

// States contains the state records
//nolint:gochecknoglobals // Currently global by design, only written once
var States map[string]*StateRecord

// LoadStates loads states from the supplied file
func LoadStates(file []byte) {
	States = make(map[string]*StateRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &StateRecord{
			State:         d.String("state"),
			Group:         d.Number("group"),
			RemHit:        d.Number("remhit") > 0,
			NoSend:        d.Number("nosend") > 0,
			Transform:     d.Number("transform") > 0,
			Aura:          d.Number("aura") > 0,
			Cureable:      d.Number("cureable") > 0,
			Curse:         d.Number("curse") > 0,
			Active:        d.Number("active") > 0,
			Restrict:      d.Number("restrict") > 0,
			Disguise:      d.Number("disguise") > 0,
			Blue:          d.Number("blue") > 0,
			AttBlue:       d.Number("attblue") > 0,
			DmgBlue:       d.Number("dmgblue") > 0,
			ArmBlue:       d.Number("armblue") > 0,
			RfBlue:        d.Number("rfblue") > 0,
			RlBlue:        d.Number("rlblue") > 0,
			RcBlue:        d.Number("rcblue") > 0,
			RpBlue:        d.Number("rpblue") > 0,
			AttRed:        d.Number("attred") > 0,
			DmgRed:        d.Number("dmgred") > 0,
			ArmRed:        d.Number("armred") > 0,
			RfRed:         d.Number("rfred") > 0,
			RlRed:         d.Number("rlred") > 0,
			RcRed:         d.Number("rcred") > 0,
			RpRed:         d.Number("rpred") > 0,
			StamBarBlue:   d.Number("stambarblue") > 0,
			Exp:           d.Number("exp") > 0,
			PlrStayDeath:  d.Number("plrstaydeath") > 0,
			MonStayDeath:  d.Number("monstaydeath") > 0,
			BossStayDeath: d.Number("bossstaydeath") > 0,
			Hide:          d.Number("hide") > 0,
			Shatter:       d.Number("shatter") > 0,
			UDead:         d.Number("udead") > 0,
			Life:          d.Number("life") > 0,
			Green:         d.Number("green") > 0,
			Pgsv:          d.Number("pgsv") > 0,
			NoOverlays:    d.Number("nooverlays") > 0,
			NoClear:       d.Number("noclear") > 0,
			BossInv:       d.Number("bossinv") > 0,
			MeleeOnly:     d.Number("meleeonly") > 0,
			NotOnDead:     d.Number("notondead") > 0,
			Overlay1:      d.String("overlay1"),
			Overlay2:      d.String("overlay2"),
			Overlay3:      d.String("overlay3"),
			Overlay4:      d.String("overlay4"),
			PgOverlay:     d.String("pgoverlay"),
			CastOverlay:   d.String("castoverlay"),
			RemOverlay:    d.String("removerlay"),
			Stat:          d.String("stat"),
			SetFunc:       d.Number("setfunc"),
			RemFunc:       d.Number("remfun"),
			Missile:       d.String("missile"),
			Skill:         d.String("skill"),
			ItemType:      d.String("itemtype"),
			ItemTrans:     d.String("itemtrans"),
			ColorPri:      d.Number("colorpri"),
			ColorShift:    d.Number("colorshift"),
			LightR:        d.Number("light-r"),
			LightG:        d.Number("light-g"),
			LightB:        d.Number("light-b"),
			OnSound:       d.String("onsound"),
			OffSound:      d.String("offsound"),
			GfxType:       d.Number("gfxtype"),
			GfxClass:      d.Number("gfxclass"),
			CltEvent:      d.String("cltevent"),
			CltEventFunc:  d.Number("clteventfunc"),
			CltActiveFunc: d.Number("cltactivefun"),
			SrvActiveFunc: d.Number("srvactivefunc"),
		}
		States[record.State] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d State records", len(States))
}
