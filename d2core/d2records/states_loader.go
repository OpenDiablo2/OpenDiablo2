package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func statesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string]*StateRecord)

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

		records[record.State] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.States = records

	r.Logger.Infof("Loaded %d State records", len(records))

	return nil
}
