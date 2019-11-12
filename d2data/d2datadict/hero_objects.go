package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

var HeroObjects = map[d2enum.Hero]*ObjectLookupRecord{
	d2enum.HeroBarbarian: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "BA", Class: "HTH", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
	d2enum.HeroNecromancer: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "NE", Class: "HTH", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
	d2enum.HeroPaladin: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "PA", Class: "1HS", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
	d2enum.HeroAssassin: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "AI", Class: "HTH", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
	d2enum.HeroSorceress: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "SO", Class: "HTH", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
	d2enum.HeroAmazon: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "AM", Class: "1HT", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
	d2enum.HeroDruid: &ObjectLookupRecord{
		Mode:  d2enum.AnimationModePlayerNeutral.String(),
		Base:  "/data/global/chars",
		Token: "DZ", Class: "HTH", HD: "LIT", TR: "LIT", LG: "LIT", RA: "LIT", LA: "LIT", RH: "", LH: "",
	},
}
