package common

import (
	"log"
	"strings"
)

type ObjectType int

const (
	ObjectTypeCharacter ObjectType = 1
	ObjectTypeItem      ObjectType = 2
)

type ObjectLookupRecord struct {
	Act           int
	Type          ObjectType
	Id            int
	Description   string
	ObjectsTxtId  int
	MonstatsTxtId int
	Direction     int
	Base          string
	Token         string
	Mode          string
	Class         string
	HD            string
	TR            string
	LG            string
	RA            string
	LA            string
	RH            string
	LH            string
	SH            string
	S1            string
	S2            string
	S3            string
	S4            string
	S5            string
	S6            string
	S7            string
	S8            string
	ColorMap      string
	Index         int
}

var ObjectLookups []ObjectLookupRecord

func LoadObjectLookups() {
	lines := strings.Split(objLookupRaw, "\n")
	ObjectLookups = make([]ObjectLookupRecord, len(lines))
	for lineIdx, line := range lines {
		lineParts := strings.Split(line, "\t")
		idx := -1
		inc := func() int {
			idx++
			return idx
		}
		for len(lineParts) < 29 {
			lineParts = append(lineParts, "")
		}
		ObjectLookups[lineIdx].Act = SafeStringToInt(lineParts[inc()])
		ObjectLookups[lineIdx].Type = ObjectType(SafeStringToInt(lineParts[inc()]))
		ObjectLookups[lineIdx].Id = SafeStringToInt(lineParts[inc()])
		ObjectLookups[lineIdx].Description = lineParts[inc()]
		ObjectLookups[lineIdx].ObjectsTxtId = SafeStringToInt(lineParts[inc()])
		ObjectLookups[lineIdx].MonstatsTxtId = SafeStringToInt(lineParts[inc()])
		ObjectLookups[lineIdx].Direction = SafeStringToInt(lineParts[inc()])
		ObjectLookups[lineIdx].Base = lineParts[inc()]
		ObjectLookups[lineIdx].Token = lineParts[inc()]
		ObjectLookups[lineIdx].Mode = lineParts[inc()]
		ObjectLookups[lineIdx].Class = lineParts[inc()]
		ObjectLookups[lineIdx].HD = lineParts[inc()]
		ObjectLookups[lineIdx].TR = lineParts[inc()]
		ObjectLookups[lineIdx].LG = lineParts[inc()]
		ObjectLookups[lineIdx].RA = lineParts[inc()]
		ObjectLookups[lineIdx].LA = lineParts[inc()]
		ObjectLookups[lineIdx].RH = lineParts[inc()]
		ObjectLookups[lineIdx].LH = lineParts[inc()]
		ObjectLookups[lineIdx].SH = lineParts[inc()]
		ObjectLookups[lineIdx].S1 = lineParts[inc()]
		ObjectLookups[lineIdx].S2 = lineParts[inc()]
		ObjectLookups[lineIdx].S3 = lineParts[inc()]
		ObjectLookups[lineIdx].S4 = lineParts[inc()]
		ObjectLookups[lineIdx].S5 = lineParts[inc()]
		ObjectLookups[lineIdx].S6 = lineParts[inc()]
		ObjectLookups[lineIdx].S7 = lineParts[inc()]
		ObjectLookups[lineIdx].S8 = lineParts[inc()]
		ObjectLookups[lineIdx].ColorMap = lineParts[inc()]
		ObjectLookups[lineIdx].Index = SafeStringToInt(lineParts[inc()])
	}
	log.Printf("Loaded %d object lookups", len(ObjectLookups))
}

func LookupObject(act, typ, id int) *ObjectLookupRecord {
	for _, lookup := range ObjectLookups {
		if lookup.Act != act || int(lookup.Type) != typ || lookup.Id != id {
			continue
		}
		return &lookup
	}
	log.Panicf("Failed to look up object Act: %d, Type: %d, Id: %d", act, typ, id)
	return nil
}

// TODO: This really should be a struct, but ain't nobody got time for that...
var objLookupRaw = `1	1	0	gheed-ACT 1 TABLE				/Data/Global/Monsters	GH	NU	HTH		LIT																	0
1	1	1	cain1-ACT 1 TABLE				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	1	2	akara-ACT 1 TABLE				/Data/Global/Monsters	PS	NU	HTH		LIT																	0
1	1	3	chicken-ACT 1 TABLE				/Data/Global/Monsters	CK	NU	HTH		LIT																	0
1	1	4	rogue1-ACT 1 TABLE				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
1	1	5	kashya-ACT 1 TABLE				/Data/Global/Monsters	RC	NU	HTH		LIT																	0
1	1	6	cow-ACT 1 TABLE				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
1	1	7	warriv1-ACT 1 TABLE				/Data/Global/Monsters	WA	NU	HTH		LIT																	0
1	1	8	charsi-ACT 1 TABLE				/Data/Global/Monsters	CI	NU	HTH		LIT																	0
1	1	9	andariel-ACT 1 TABLE				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
1	1	10	place_fallen-ACT 1 TABLE				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
1	1	11	place_fallenshaman-ACT 1 TABLE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	12	place_bloodraven-ACT 1 TABLE				/Data/Global/Monsters	CR	NU	BOW	BRV	HVY	BRV	HVY	HVY	LIT	LBB		HVY	HVY									0
1	1	13	cow-ACT 1 TABLE				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
1	1	14	camel-ACT 1 TABLE				/Data/Global/Monsters	CM	NU	HTH		LIT																	0
1	1	15	place_unique_pack-ACT 1 TABLE																										0
1	1	16	place_npc_pack-ACT 1 TABLE																										0
1	1	17	place_nothing-ACT 1 TABLE																										0
1	1	18	place_nothing-ACT 1 TABLE																										0
1	1	19	place_champion-ACT 1 TABLE																										0
1	1	20	navi-ACT 1 TABLE				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
1	1	21	rogue1-ACT 1 TABLE				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
1	1	22	rogue3-ACT 1 TABLE				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
1	1	23	gargoyletrap-ACT 1 TABLE				/Data/Global/Monsters	GT	A1	HTH		LIT																	0
1	1	24	place_fallennest-ACT 1 TABLE																										0
1	1	25	place_talkingrogue-ACT 1 TABLE																										0
1	1	26	place_fallen-ACT 1 TABLE				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		BUC	MED										0
1	1	27	place_fallenshaman-ACT 1 TABLE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	28	trap-horzmissile-ACT 1 TABLE																										0
1	1	29	trap-vertmissile-ACT 1 TABLE																										0
1	1	30	place_group25-ACT 1 TABLE																										0
1	1	31	place_group50-ACT 1 TABLE																										0
1	1	32	place_group75-ACT 1 TABLE																										0
1	1	33	place_group100-ACT 1 TABLE																										0
1	1	34	Bishibosh-ACT 1 TABLE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	35	Bonebreak-ACT 1 TABLE				/Data/Global/Monsters	SK	NU	1HS	HVY	LIT	MED	LIT	LIT	SCM			MED	MED									0
1	1	36	Coldcrow-ACT 1 TABLE				/Data/Global/Monsters	CR	NU	BOW	HVY	LIT	LIT	LIT	LIT	LIT	LBB		LIT	LIT									0
1	1	37	Rakanishu-ACT 1 TABLE				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
1	1	38	Treehead WoodFist-ACT 1 TABLE				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	39	Griswold-ACT 1 TABLE				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
1	1	40	The Countess-ACT 1 TABLE				/Data/Global/Monsters	CR	NU	1HS	MED	LIT	MED	LIT	LIT	WHM			LIT	LIT									0
1	1	41	Pitspawn Fouldog-ACT 1 TABLE				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	42	Flamespike the Crawler-ACT 1 TABLE				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	43	Boneash-ACT 1 TABLE				/Data/Global/Monsters	SK	NU	HTH	MED	MED	DES	LIT	LIT				LIT			POS	POS						0
1	1	44	The Smith-ACT 1 TABLE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
1	1	45	The Cow King-ACT 1 TABLE				/Data/Global/Monsters	EC	NU	HTH		LIT				HAL													0
1	1	46	Corpsefire-ACT 1 TABLE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	47	skeleton1-Skeleton-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	48	skeleton2-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	49	skeleton3-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	50	skeleton4-BurningDead-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	51	skeleton5-Horror-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	52	zombie1-Zombie-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	53	zombie2-HungryDead-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	54	zombie3-Ghoul-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	55	zombie4-DrownedCarcass-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	56	zombie5-PlagueBearer-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	57	bighead1-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	58	bighead2-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	59	bighead3-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	60	bighead4-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	61	bighead5-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	62	foulcrow1-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	63	foulcrow2-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	64	foulcrow3-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	65	foulcrow4-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	66	fallen1-Fallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
1	1	67	fallen2-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
1	1	68	fallen3-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
1	1	69	fallen4-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
1	1	70	fallen5-WarpedFallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
1	1	71	brute2-Brute-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	72	brute3-Yeti-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	73	brute4-Crusher-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	74	brute5-WailingBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	75	brute1-GargantuanBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	76	sandraider1-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	77	sandraider2-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	78	sandraider3-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	79	sandraider4-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	80	sandraider5-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	81	gorgon1-unused-Idle				/Data/Global/Monsters	GO																					0
1	1	82	gorgon2-unused-Idle				/Data/Global/Monsters	GO																					0
1	1	83	gorgon3-unused-Idle				/Data/Global/Monsters	GO																					0
1	1	84	gorgon4-unused-Idle				/Data/Global/Monsters	GO																					0
1	1	85	wraith1-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	86	wraith2-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	87	wraith3-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	88	wraith4-Apparition-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	89	wraith5-DarkShape-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	90	corruptrogue1-DarkHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	91	corruptrogue2-VileHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	92	corruptrogue3-DarkStalker-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	93	corruptrogue4-BlackRogue-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	94	corruptrogue5-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	95	baboon1-DuneBeast-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	96	baboon2-RockDweller-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	97	baboon3-JungleHunter-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	98	baboon4-DoomApe-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	99	baboon5-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	100	goatman1-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	101	goatman2-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	102	goatman3-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	103	goatman4-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	104	goatman5-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	105	fallenshaman1-FallenShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	106	fallenshaman2-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	107	fallenshaman3-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	108	fallenshaman4-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	109	fallenshaman5-WarpedShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	110	quillrat1-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	111	quillrat2-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	112	quillrat3-ThornBeast-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	113	quillrat4-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	114	quillrat5-JungleUrchin-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	115	sandmaggot1-SandMaggot-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	116	sandmaggot2-RockWorm-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	117	sandmaggot3-Devourer-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	118	sandmaggot4-GiantLamprey-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	119	sandmaggot5-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	120	clawviper1-TombViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	121	clawviper2-ClawViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	122	clawviper3-Salamander-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	123	clawviper4-PitViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	124	clawviper5-SerpentMagus-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	125	sandleaper1-SandLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	126	sandleaper2-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	127	sandleaper3-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	128	sandleaper4-TreeLurker-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	129	sandleaper5-RazorPitDemon-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	130	pantherwoman1-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	131	pantherwoman2-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	132	pantherwoman3-NightTiger-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	133	pantherwoman4-HellCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	134	swarm1-Itchies-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
1	1	135	swarm2-BlackLocusts-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
1	1	136	swarm3-PlagueBugs-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
1	1	137	swarm4-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
1	1	138	scarab1-DungSoldier-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	139	scarab2-SandWarrior-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	140	scarab3-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	141	scarab4-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	142	scarab5-AlbinoRoach-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	143	mummy1-DriedCorpse-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	144	mummy2-Decayed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	145	mummy3-Embalmed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	146	mummy4-PreservedDead-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	147	mummy5-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	148	unraveler1-HollowOne-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	149	unraveler2-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	150	unraveler3-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	151	unraveler4-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	152	unraveler5-Baal Subject Mummy-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	153	chaoshorde1-unused-Idle				/Data/Global/Monsters	CH																					0
1	1	154	chaoshorde2-unused-Idle				/Data/Global/Monsters	CH																					0
1	1	155	chaoshorde3-unused-Idle				/Data/Global/Monsters	CH																					0
1	1	156	chaoshorde4-unused-Idle				/Data/Global/Monsters	CH																					0
1	1	157	vulture1-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
1	1	158	vulture2-UndeadScavenger-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
1	1	159	vulture3-HellBuzzard-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
1	1	160	vulture4-WingedNightmare-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
1	1	161	mosquito1-Sucker-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
1	1	162	mosquito2-Feeder-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
1	1	163	mosquito3-BloodHook-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
1	1	164	mosquito4-BloodWing-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
1	1	165	willowisp1-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	166	willowisp2-SwampGhost-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	167	willowisp3-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	168	willowisp4-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	169	arach1-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	170	arach2-SandFisher-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	171	arach3-PoisonSpinner-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	172	arach4-FlameSpider-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	173	arach5-SpiderMagus-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	174	thornhulk1-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
1	1	175	thornhulk2-BrambleHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
1	1	176	thornhulk3-Thrasher-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
1	1	177	thornhulk4-Spikefist-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
1	1	178	vampire1-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	179	vampire2-NightLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	180	vampire3-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	181	vampire4-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	182	vampire5-Banished-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	183	batdemon1-DesertWing-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	184	batdemon2-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	185	batdemon3-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	186	batdemon4-BloodDiver-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	187	batdemon5-DarkFamiliar-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	188	fetish1-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	189	fetish2-Fetish-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	190	fetish3-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	191	fetish4-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	192	fetish5-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	193	cain1-DeckardCain-NpcOutOfTown				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	1	194	gheed-Gheed-Npc				/Data/Global/Monsters	GH	NU	HTH		LIT																	0
1	1	195	akara-Akara-Npc				/Data/Global/Monsters	PS	NU	HTH		LIT																	0
1	1	196	chicken-dummy-Idle				/Data/Global/Monsters	CK	NU	HTH		LIT																	0
1	1	197	kashya-Kashya-Npc				/Data/Global/Monsters	RC	NU	HTH		LIT																	0
1	1	198	rat-dummy-Idle				/Data/Global/Monsters	RT	NU	HTH		LIT																	0
1	1	199	rogue1-Dummy-Idle				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
1	1	200	hellmeteor-Dummy-HellMeteor				/Data/Global/Monsters	K9																					0
1	1	201	charsi-Charsi-Npc				/Data/Global/Monsters	CI	NU	HTH		LIT																	0
1	1	202	warriv1-Warriv-Npc				/Data/Global/Monsters	WA	NU	HTH		LIT																	0
1	1	203	andariel-Andariel-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
1	1	204	bird1-dummy-Idle				/Data/Global/Monsters	BS	WL	HTH		LIT																	0
1	1	205	bird2-dummy-Idle				/Data/Global/Monsters	BL																					0
1	1	206	bat-dummy-Idle				/Data/Global/Monsters	B9	WL	HTH		LIT																	0
1	1	207	cr_archer1-DarkRanger-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	208	cr_archer2-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	209	cr_archer3-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	210	cr_archer4-BlackArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	211	cr_archer5-FleshArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	212	cr_lancer1-DarkSpearwoman-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	213	cr_lancer2-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	214	cr_lancer3-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	215	cr_lancer4-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	216	cr_lancer5-FleshLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	217	sk_archer1-SkeletonArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	218	sk_archer2-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	219	sk_archer3-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	220	sk_archer4-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	221	sk_archer5-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	222	warriv2-Warriv-Npc				/Data/Global/Monsters	WX	NU	HTH		LIT																	0
1	1	223	atma-Atma-Npc				/Data/Global/Monsters	AS	NU	HTH		LIT																	0
1	1	224	drognan-Drognan-Npc				/Data/Global/Monsters	DR	NU	HTH		LIT																	0
1	1	225	fara-Fara-Npc				/Data/Global/Monsters	OF	NU	HTH		LIT																	0
1	1	226	cow-dummy-Idle				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
1	1	227	maggotbaby1-SandMaggotYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
1	1	228	maggotbaby2-RockWormYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
1	1	229	maggotbaby3-DevourerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
1	1	230	maggotbaby4-GiantLampreyYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
1	1	231	maggotbaby5-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
1	1	232	camel-dummy-Idle				/Data/Global/Monsters	CM	NU	HTH		LIT																	0
1	1	233	blunderbore1-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
1	1	234	blunderbore2-Gorbelly-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
1	1	235	blunderbore3-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
1	1	236	blunderbore4-Urdar-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
1	1	237	maggotegg1-SandMaggotEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
1	1	238	maggotegg2-RockWormEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
1	1	239	maggotegg3-DevourerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
1	1	240	maggotegg4-GiantLampreyEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
1	1	241	maggotegg5-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
1	1	242	act2male-dummy-Towner				/Data/Global/Monsters	2M	NU	HTH	OLD	MED	MED						TUR										0
1	1	243	act2female-Dummy-Towner				/Data/Global/Monsters	2F	NU	HTH	LIT	LIT	LIT																0
1	1	244	act2child-dummy-Towner				/Data/Global/Monsters	2C																					0
1	1	245	greiz-Greiz-Npc				/Data/Global/Monsters	GR	NU	HTH		LIT																	0
1	1	246	elzix-Elzix-Npc				/Data/Global/Monsters	EL	NU	HTH		LIT																	0
1	1	247	geglash-Geglash-Npc				/Data/Global/Monsters	GE	NU	HTH		LIT																	0
1	1	248	jerhyn-Jerhyn-Npc				/Data/Global/Monsters	JE	NU	HTH		LIT																	0
1	1	249	lysander-Lysander-Npc				/Data/Global/Monsters	LY	NU	HTH		LIT																	0
1	1	250	act2guard1-Dummy-Towner				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
1	1	251	act2vendor1-dummy-Vendor				/Data/Global/Monsters	M1	NU	HTH		LIT																	0
1	1	252	act2vendor2-dummy-Vendor				/Data/Global/Monsters	M2	NU	HTH		LIT																	0
1	1	253	crownest1-FoulCrowNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
1	1	254	crownest2-BloodHawkNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
1	1	255	crownest3-BlackVultureNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
1	1	256	crownest4-CloudStalkerNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
1	1	257	meshif1-Meshif-Npc				/Data/Global/Monsters	MS	NU	HTH		LIT																	0
1	1	258	duriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
1	1	259	bonefetish1-Undead RatMan-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	260	bonefetish2-Undead Fetish-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	261	bonefetish3-Undead Flayer-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	262	bonefetish4-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	263	bonefetish5-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	264	darkguard1-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	265	darkguard2-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	266	darkguard3-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	267	darkguard4-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	268	darkguard5-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	269	bloodmage1-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	270	bloodmage2-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	271	bloodmage3-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	272	bloodmage4-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	273	bloodmage5-unused-Idle				/Data/Global/Monsters	xx																					0
1	1	274	maggot-Maggot-Idle				/Data/Global/Monsters	MA	NU	HTH		LIT																	0
1	1	275	sarcophagus-MummyGenerator-Sarcophagus				/Data/Global/Monsters	MG	NU	HTH		LIT																	0
1	1	276	radament-Radament-GreaterMummy				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
1	1	277	firebeast-unused-ElementalBeast				/Data/Global/Monsters	FM	NU	HTH		LIT																	0
1	1	278	iceglobe-unused-ElementalBeast				/Data/Global/Monsters	IM	NU	HTH		LIT																	0
1	1	279	lightningbeast-unused-ElementalBeast				/Data/Global/Monsters	xx																					0
1	1	280	poisonorb-unused-ElementalBeast				/Data/Global/Monsters	PM	NU	HTH		LIT																	0
1	1	281	flyingscimitar-FlyingScimitar-FlyingScimitar				/Data/Global/Monsters	ST	NU	HTH		LIT																	0
1	1	282	zealot1-Zakarumite-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
1	1	283	zealot2-Faithful-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
1	1	284	zealot3-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
1	1	285	cantor1-Sexton-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
1	1	286	cantor2-Cantor-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
1	1	287	cantor3-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
1	1	288	cantor4-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
1	1	289	mephisto-Mephisto-Mephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
1	1	290	diablo-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
1	1	291	cain2-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	1	292	cain3-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	1	293	cain4-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	1	294	frogdemon1-Swamp Dweller-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
1	1	295	frogdemon2-Bog Creature-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
1	1	296	frogdemon3-Slime Prince-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
1	1	297	summoner-Summoner-Summoner				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
1	1	298	tyrael1-tyrael-NpcStationary				/Data/Global/Monsters	TX	NU	HTH		LIT		LIT	LIT														0
1	1	299	asheara-asheara-Npc				/Data/Global/Monsters	AH	NU	HTH		LIT																	0
1	1	300	hratli-hratli-Npc				/Data/Global/Monsters	HR	NU	HTH		LIT																	0
1	1	301	alkor-alkor-Npc				/Data/Global/Monsters	AL	NU	HTH		LIT																	0
1	1	302	ormus-ormus-Npc				/Data/Global/Monsters	OR	NU	HTH		LIT																	0
1	1	303	izual-izual-Izual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
1	1	304	halbu-halbu-Npc				/Data/Global/Monsters	20	NU	HTH		LIT																	0
1	1	305	tentacle1-WaterWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
1	1	306	tentacle2-RiverStalkerLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
1	1	307	tentacle3-StygianWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
1	1	308	tentaclehead1-WaterWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
1	1	309	tentaclehead2-RiverStalkerHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
1	1	310	tentaclehead3-StygianWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
1	1	311	meshif2-meshif-Npc				/Data/Global/Monsters	M3	NU	HTH		LIT																	0
1	1	312	cain5-DeckardCain-Npc				/Data/Global/Monsters	1D	NU	HTH		LIT																	0
1	1	313	navi-navi-Navi				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
1	1	314	bloodraven-Bloodraven-BloodRaven				/Data/Global/Monsters	CR	NU	BOW	BRV	HVY	BRV	HVY	HVY	LIT	LBB		HVY	HVY									0
1	1	315	bug-Dummy-Idle				/Data/Global/Monsters	BG	NU	HTH		LIT																	0
1	1	316	scorpion-Dummy-Idle				/Data/Global/Monsters	DS	NU	HTH		LIT																	0
1	1	317	rogue2-RogueScout-GoodNpcRanged				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
1	1	318	roguehire-Dummy-Hireable				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
1	1	319	rogue3-Dummy-TownRogue				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
1	1	320	gargoyletrap-GargoyleTrap-GargoyleTrap				/Data/Global/Monsters	GT	NU	HTH		LIT																	0
1	1	321	skmage_pois1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	322	skmage_pois2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	323	skmage_pois3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	324	skmage_pois4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	325	fetishshaman1-RatManShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	326	fetishshaman2-FetishShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	327	fetishshaman3-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	328	fetishshaman4-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	329	fetishshaman5-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	330	larva-larva-Idle				/Data/Global/Monsters	LV	NU	HTH		LIT																	0
1	1	331	maggotqueen1-SandMaggotQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
1	1	332	maggotqueen2-RockWormQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
1	1	333	maggotqueen3-DevourerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
1	1	334	maggotqueen4-GiantLampreyQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
1	1	335	maggotqueen5-WorldKillerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
1	1	336	claygolem-ClayGolem-NecroPet				/Data/Global/Monsters	G1	NU	HTH		LIT																	0
1	1	337	bloodgolem-BloodGolem-NecroPet				/Data/Global/Monsters	G2	NU	HTH		LIT																	0
1	1	338	irongolem-IronGolem-NecroPet				/Data/Global/Monsters	G4	NU	HTH		LIT																	0
1	1	339	firegolem-FireGolem-NecroPet				/Data/Global/Monsters	G3	NU	HTH		LIT																	0
1	1	340	familiar-Dummy-Idle				/Data/Global/Monsters	FI	NU	HTH		LIT																	0
1	1	341	act3male-Dummy-Towner				/Data/Global/Monsters	N4	NU	HTH	BRD	HVY	HVY	HEV	HEV	FSH	SAK		TKT										0
1	1	342	baboon6-NightMarauder-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	343	act3female-Dummy-Towner				/Data/Global/Monsters	N3	NU	HTH	LIT	MTP	SRT			BSK	BSK												0
1	1	344	natalya-Natalya-Npc				/Data/Global/Monsters	TZ	NU	HTH		LIT																	0
1	1	345	vilemother1-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
1	1	346	vilemother2-StygianHag-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
1	1	347	vilemother3-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
1	1	348	vilechild1-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
1	1	349	vilechild2-StygianDog-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
1	1	350	vilechild3-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
1	1	351	fingermage1-Groper-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
1	1	352	fingermage2-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
1	1	353	fingermage3-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
1	1	354	regurgitator1-Corpulent-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
1	1	355	regurgitator2-CorpseSpitter-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
1	1	356	regurgitator3-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
1	1	357	doomknight1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	358	doomknight2-AbyssKnight-AbyssKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	359	doomknight3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	360	quillbear1-QuillBear-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
1	1	361	quillbear2-SpikeGiant-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
1	1	362	quillbear3-ThornBrute-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
1	1	363	quillbear4-RazorBeast-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
1	1	364	quillbear5-GiantUrchin-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
1	1	365	snake-Dummy-Idle				/Data/Global/Monsters	CO	NU	HTH		LIT																	0
1	1	366	parrot-Dummy-Idle				/Data/Global/Monsters	PR	WL	HTH		LIT																	0
1	1	367	fish-Dummy-Idle				/Data/Global/Monsters	FJ																					0
1	1	368	evilhole1-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
1	1	369	evilhole2-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
1	1	370	evilhole3-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
1	1	371	evilhole4-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
1	1	372	evilhole5-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
1	1	373	trap-firebolt-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
1	1	374	trap-horzmissile-a trap-Trap-RightArrow				/Data/Global/Monsters	9A																					0
1	1	375	trap-vertmissile-a trap-Trap-LeftArrow				/Data/Global/Monsters	9A																					0
1	1	376	trap-poisoncloud-a trap-Trap-Poison				/Data/Global/Monsters	9A																					0
1	1	377	trap-lightning-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
1	1	378	act2guard2-Kaelan-JarJar				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
1	1	379	invisospawner-Dummy-InvisoSpawner				/Data/Global/Monsters	K9																					0
1	1	380	diabloclone-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
1	1	381	suckernest1-SuckerNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
1	1	382	suckernest2-FeederNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
1	1	383	suckernest3-BloodHookNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
1	1	384	suckernest4-BloodWingNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
1	1	385	act2hire-Guard-Hireable				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
1	1	386	minispider-Dummy-Idle				/Data/Global/Monsters	LS	NU	HTH		LIT																	0
1	1	387	boneprison1--Idle				/Data/Global/Monsters	67	NU	HTH		LIT																	0
1	1	388	boneprison2--Idle				/Data/Global/Monsters	66	NU	HTH		LIT																	0
1	1	389	boneprison3--Idle				/Data/Global/Monsters	69	NU	HTH		LIT																	0
1	1	390	boneprison4--Idle				/Data/Global/Monsters	68	NU	HTH		LIT																	0
1	1	391	bonewall-Dummy-BoneWall				/Data/Global/Monsters	BW	NU	HTH		LIT																	0
1	1	392	councilmember1-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	393	councilmember2-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	394	councilmember3-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	395	turret1-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
1	1	396	turret2-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
1	1	397	turret3-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
1	1	398	hydra1-Hydra-Hydra				/Data/Global/Monsters	HX	NU	HTH		LIT							LIT										0
1	1	399	hydra2-Hydra-Hydra				/Data/Global/Monsters	21	NU	HTH		LIT							LIT										0
1	1	400	hydra3-Hydra-Hydra				/Data/Global/Monsters	HZ	NU	HTH		LIT							LIT										0
1	1	401	trap-melee-a trap-Trap-Melee				/Data/Global/Monsters	M4	A1	HTH		LIT																	0
1	1	402	seventombs-Dummy-7TIllusion				/Data/Global/Monsters	9A																					0
1	1	403	dopplezon-Dopplezon-Idle				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
1	1	404	valkyrie-Valkyrie-NecroPet				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
1	1	405	act2guard3-Dummy-Idle				/Data/Global/Monsters	SK																					0
1	1	406	act3hire-Iron Wolf-Hireable				/Data/Global/Monsters	IW	NU	1HS	LIT	LIT				WND		KIT											0
1	1	407	megademon1-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	408	megademon2-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	409	megademon3-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	410	necroskeleton-NecroSkeleton-NecroPet				/Data/Global/Monsters	SK	NU	1HS	DES	HVY	DES	DES	DES	SCM		KIT	DES	DES	LIT								0
1	1	411	necromage-NecroMage-NecroPet				/Data/Global/Monsters	SK	NU	HTH	DES	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
1	1	412	griswold-Griswold-Griswold				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
1	1	413	compellingorb-compellingorb-Idle				/Data/Global/Monsters	9a																					0
1	1	414	tyrael2-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
1	1	415	darkwanderer-youngdiablo-DarkWanderer				/Data/Global/Monsters	1Z	NU	HTH		LIT																	0
1	1	416	trap-nova-a trap-Trap-Nova				/Data/Global/Monsters	9A																					0
1	1	417	spiritmummy-Dummy-Idle				/Data/Global/Monsters	xx																					0
1	1	418	lightningspire-LightningSpire-ArcaneTower				/Data/Global/Monsters	AE	NU	HTH		LIT							LIT										0
1	1	419	firetower-FireTower-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
1	1	420	slinger1-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
1	1	421	slinger2-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
1	1	422	slinger3-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
1	1	423	slinger4-HellSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
1	1	424	act2guard4-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
1	1	425	act2guard5-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
1	1	426	skmage_cold1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
1	1	427	skmage_cold2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
1	1	428	skmage_cold3-BaalColdMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
1	1	429	skmage_cold4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
1	1	430	skmage_fire1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
1	1	431	skmage_fire2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
1	1	432	skmage_fire3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
1	1	433	skmage_fire4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
1	1	434	skmage_ltng1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
1	1	435	skmage_ltng2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
1	1	436	skmage_ltng3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
1	1	437	skmage_ltng4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
1	1	438	hellbovine-Hell Bovine-Skeleton				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
1	1	439	window1--Idle				/Data/Global/Monsters	VH	NU	HTH		LIT							LIT										0
1	1	440	window2--Idle				/Data/Global/Monsters	VJ	NU	HTH		LIT							LIT										0
1	1	441	slinger5-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
1	1	442	slinger6-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
1	1	443	fetishblow1-RatMan-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	444	fetishblow2-Fetish-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	445	fetishblow3-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	446	fetishblow4-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	447	fetishblow5-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	448	mephistospirit-Dummy-Spirit				/Data/Global/Monsters	M6	A1	HTH		LIT																	0
1	1	449	smith-The Smith-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
1	1	450	trappedsoul1-TrappedSoul-TrappedSoul				/Data/Global/Monsters	10	NU	HTH		LIT																	0
1	1	451	trappedsoul2-TrappedSoul-TrappedSoul				/Data/Global/Monsters	13	NU	HTH		LIT																	0
1	1	452	jamella-Jamella-Npc				/Data/Global/Monsters	ja	NU	HTH		LIT																	0
1	1	453	izualghost-Izual-NpcStationary				/Data/Global/Monsters	17	NU	HTH		LIT							LIT										0
1	1	454	fetish11-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	455	malachai-Malachai-Buffy				/Data/Global/Monsters	36	NU	HTH		LIT							LIT										0
1	1	456	hephasto-The Feature Creep-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
1	1	457	wakeofdestruction-Wake of Destruction-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
1	1	458	chargeboltsentry-Charged Bolt Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
1	1	459	lightningsentry-Lightning Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
1	1	460	bladecreeper-Blade Creeper-BladeCreeper				/Data/Global/Monsters	b8	NU	HTH		LIT							LIT										0
1	1	461	invisopet-Invis Pet-InvisoPet				/Data/Global/Monsters	k9																					0
1	1	462	infernosentry-Inferno Sentry-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
1	1	463	deathsentry-Death Sentry-DeathSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
1	1	464	shadowwarrior-Shadow Warrior-ShadowWarrior				/Data/Global/Monsters	k9																					0
1	1	465	shadowmaster-Shadow Master-ShadowMaster				/Data/Global/Monsters	k9																					0
1	1	466	druidhawk-Druid Hawk-Raven				/Data/Global/Monsters	hk	NU	HTH		LIT																	0
1	1	467	spiritwolf-Druid Spirit Wolf-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
1	1	468	fenris-Druid Fenris-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
1	1	469	spiritofbarbs-Spirit of Barbs-Totem				/Data/Global/Monsters	x4	NU	HTH		LIT																	0
1	1	470	heartofwolverine-Heart of Wolverine-Totem				/Data/Global/Monsters	x3	NU	HTH		LIT																	0
1	1	471	oaksage-Oak Sage-Totem				/Data/Global/Monsters	xw	NU	HTH		LIT																	0
1	1	472	plaguepoppy-Druid Plague Poppy-Vines				/Data/Global/Monsters	k9																					0
1	1	473	cycleoflife-Druid Cycle of Life-CycleOfLife				/Data/Global/Monsters	k9																					0
1	1	474	vinecreature-Vine Creature-CycleOfLife				/Data/Global/Monsters	k9																					0
1	1	475	druidbear-Druid Bear-DruidBear				/Data/Global/Monsters	b7	NU	HTH		LIT																	0
1	1	476	eagle-Eagle-Idle				/Data/Global/Monsters	eg	NU	HTH		LIT							LIT										0
1	1	477	wolf-Wolf-NecroPet				/Data/Global/Monsters	40	NU	HTH		LIT																	0
1	1	478	bear-Bear-NecroPet				/Data/Global/Monsters	TG	NU	HTH		LIT							LIT										0
1	1	479	barricadedoor1-Barricade Door-Idle				/Data/Global/Monsters	AJ	NU	HTH		LIT																	0
1	1	480	barricadedoor2-Barricade Door-Idle				/Data/Global/Monsters	AG	NU	HTH		LIT																	0
1	1	481	prisondoor-Prison Door-Idle				/Data/Global/Monsters	2Q	NU	HTH		LIT																	0
1	1	482	barricadetower-Barricade Tower-SiegeTower				/Data/Global/Monsters	ac	NU	HTH		LIT							LIT						LIT				0
1	1	483	reanimatedhorde1-RotWalker-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	484	reanimatedhorde2-ReanimatedHorde-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	485	reanimatedhorde3-ProwlingDead-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	486	reanimatedhorde4-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	487	reanimatedhorde5-DefiledWarrior-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	488	siegebeast1-Siege Beast-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
1	1	489	siegebeast2-CrushBiest-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
1	1	490	siegebeast3-BloodBringer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
1	1	491	siegebeast4-GoreBearer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
1	1	492	siegebeast5-DeamonSteed-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
1	1	493	snowyeti1-SnowYeti1-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
1	1	494	snowyeti2-SnowYeti2-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
1	1	495	snowyeti3-SnowYeti3-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
1	1	496	snowyeti4-SnowYeti4-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
1	1	497	wolfrider1-WolfRider1-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
1	1	498	wolfrider2-WolfRider2-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
1	1	499	wolfrider3-WolfRider3-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
1	1	500	minion1-Minionexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	501	minion2-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	502	minion3-IceBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	503	minion4-FireBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	504	minion5-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	505	minion6-IceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	506	minion7-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	507	minion8-GreaterIceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	508	suicideminion1-FanaticMinion-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	509	suicideminion2-BerserkSlayer-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	510	suicideminion3-ConsumedIceBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	511	suicideminion4-ConsumedFireBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	512	suicideminion5-FrenziedHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	513	suicideminion6-FrenziedIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	514	suicideminion7-InsaneHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	515	suicideminion8-InsaneIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
1	1	516	succubus1-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	517	succubus2-VileTemptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	518	succubus3-StygianHarlot-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	519	succubus4-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	520	succubus5-Blood Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	521	succubuswitch1-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	522	succubuswitch2-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	523	succubuswitch3-StygianFury-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	524	succubuswitch4-Blood Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	525	succubuswitch5-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	526	overseer1-OverSeer-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	527	overseer2-Lasher-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	528	overseer3-OverLord-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	529	overseer4-BloodBoss-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	530	overseer5-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	531	minionspawner1-MinionSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	532	minionspawner2-MinionSlayerSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	533	minionspawner3-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	534	minionspawner4-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	535	minionspawner5-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	536	minionspawner6-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	537	minionspawner7-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	538	minionspawner8-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	539	imp1-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	540	imp2-Imp2-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	541	imp3-Imp3-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	542	imp4-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	543	imp5-Imp5-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	544	catapult1-CatapultS-Catapult				/Data/Global/Monsters	65	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
1	1	545	catapult2-CatapultE-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
1	1	546	catapult3-CatapultSiege-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
1	1	547	catapult4-CatapultW-Catapult				/Data/Global/Monsters	ua	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT	LIT								0
1	1	548	frozenhorror1-Frozen Horror1-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
1	1	549	frozenhorror2-Frozen Horror2-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
1	1	550	frozenhorror3-Frozen Horror3-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
1	1	551	frozenhorror4-Frozen Horror4-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
1	1	552	frozenhorror5-Frozen Horror5-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
1	1	553	bloodlord1-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	554	bloodlord2-Blood Lord2-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	555	bloodlord3-Blood Lord3-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	556	bloodlord4-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	557	bloodlord5-Blood Lord5-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	558	larzuk-Larzuk-Npc				/Data/Global/Monsters	XR	NU	HTH		LIT																	0
1	1	559	drehya-Drehya-Npc				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
1	1	560	malah-Malah-Npc				/Data/Global/Monsters	XT	NU	HTH		LIT																	0
1	1	561	nihlathak-Nihlathak Town-Npc				/Data/Global/Monsters	0J	NU	HTH		LIT																	0
1	1	562	qual-kehk-Qual-Kehk-Npc				/Data/Global/Monsters	XV	NU	HTH		LIT																	0
1	1	563	catapultspotter1-Catapult Spotter S-CatapultSpotter				/Data/Global/Monsters	k9																					0
1	1	564	catapultspotter2-Catapult Spotter E-CatapultSpotter				/Data/Global/Monsters	k9																					0
1	1	565	catapultspotter3-Catapult Spotter Siege-CatapultSpotter				/Data/Global/Monsters	k9																					0
1	1	566	catapultspotter4-Catapult Spotter W-CatapultSpotter				/Data/Global/Monsters	k9																					0
1	1	567	cain6-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	1	568	tyrael3-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
1	1	569	act5barb1-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
1	1	570	act5barb2-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
1	1	571	barricadewall1-Barricade Wall Right-Idle				/Data/Global/Monsters	A6	NU	HTH		LIT																	0
1	1	572	barricadewall2-Barricade Wall Left-Idle				/Data/Global/Monsters	AK	NU	HTH		LIT																	0
1	1	573	nihlathakboss-Nihlathak-Nihlathak				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
1	1	574	drehyaiced-Drehya-NpcOutOfTown				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
1	1	575	evilhut-Evil hut-GenericSpawner				/Data/Global/Monsters	2T	NU	HTH		LIT							LIT										0
1	1	576	deathmauler1-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	577	deathmauler2-Death Mauler2-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	578	deathmauler3-Death Mauler3-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	579	deathmauler4-Death Mauler4-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	580	deathmauler5-Death Mauler5-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	581	act5pow-POW-Wussie				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
1	1	582	act5barb3-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
1	1	583	act5barb4-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
1	1	584	ancientstatue1-Ancient Statue 1-AncientStatue				/Data/Global/Monsters	0G	NU	HTH		LIT																	0
1	1	585	ancientstatue2-Ancient Statue 2-AncientStatue				/Data/Global/Monsters	0H	NU	HTH		LIT																	0
1	1	586	ancientstatue3-Ancient Statue 3-AncientStatue				/Data/Global/Monsters	0I	NU	HTH		LIT																	0
1	1	587	ancientbarb1-Ancient Barbarian 1-Ancient				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
1	1	588	ancientbarb2-Ancient Barbarian 2-Ancient				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
1	1	589	ancientbarb3-Ancient Barbarian 3-Ancient				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
1	1	590	baalthrone-Baal Throne-BaalThrone				/Data/Global/Monsters	41	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
1	1	591	baalcrab-Baal Crab-BaalCrab				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
1	1	592	baaltaunt-Baal Taunt-BaalTaunt				/Data/Global/Monsters	K9																					0
1	1	593	putriddefiler1-Putrid Defiler1-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
1	1	594	putriddefiler2-Putrid Defiler2-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
1	1	595	putriddefiler3-Putrid Defiler3-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
1	1	596	putriddefiler4-Putrid Defiler4-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
1	1	597	putriddefiler5-Putrid Defiler5-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
1	1	598	painworm1-Pain Worm1-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
1	1	599	painworm2-Pain Worm2-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
1	1	600	painworm3-Pain Worm3-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
1	1	601	painworm4-Pain Worm4-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
1	1	602	painworm5-Pain Worm5-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
1	1	603	bunny-dummy-Idle				/Data/Global/Monsters	48	NU	HTH		LIT																	0
1	1	604	baalhighpriest-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	605	venomlord-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				FLB													0
1	1	606	baalcrabstairs-Baal Crab to Stairs-BaalToStairs				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
1	1	607	act5hire1-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
1	1	608	act5hire2-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
1	1	609	baaltentacle1-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
1	1	610	baaltentacle2-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
1	1	611	baaltentacle3-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
1	1	612	baaltentacle4-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
1	1	613	baaltentacle5-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
1	1	614	injuredbarb1-dummy-Idle				/Data/Global/Monsters	6z	NU	HTH		LIT																	0
1	1	615	injuredbarb2-dummy-Idle				/Data/Global/Monsters	7j	NU	HTH		LIT																	0
1	1	616	injuredbarb3-dummy-Idle				/Data/Global/Monsters	7i	NU	HTH		LIT																	0
1	1	617	baalclone-Baal Crab Clone-BaalCrabClone				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
1	1	618	baalminion1-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
1	1	619	baalminion2-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
1	1	620	baalminion3-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
1	1	621	worldstoneeffect-dummy-Idle				/Data/Global/Monsters	K9																					0
1	1	622	sk_archer6-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	623	sk_archer7-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	624	sk_archer8-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	625	sk_archer9-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	626	sk_archer10-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	627	bighead6-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	628	bighead7-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	629	bighead8-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	630	bighead9-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	631	bighead10-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	632	goatman6-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	633	goatman7-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	634	goatman8-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	635	goatman9-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	636	goatman10-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
1	1	637	foulcrow5-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	638	foulcrow6-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	639	foulcrow7-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	640	foulcrow8-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
1	1	641	clawviper6-ClawViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	642	clawviper7-PitViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	643	clawviper8-Salamander-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	644	clawviper9-TombViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	645	clawviper10-SerpentMagus-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	646	sandraider6-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	647	sandraider7-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	648	sandraider8-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	649	sandraider9-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	650	sandraider10-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	651	deathmauler6-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	652	quillrat6-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	653	quillrat7-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	654	quillrat8-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	655	vulture5-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
1	1	656	thornhulk5-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
1	1	657	slinger7-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	658	slinger8-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	659	slinger9-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	660	cr_archer6-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	661	cr_archer7-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	662	cr_lancer6-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	663	cr_lancer7-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	664	cr_lancer8-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	665	blunderbore5-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
1	1	666	blunderbore6-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
1	1	667	skmage_fire5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
1	1	668	skmage_fire6-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
1	1	669	skmage_ltng5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
1	1	670	skmage_ltng6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
1	1	671	skmage_cold5-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		CLD	CLD						0
1	1	672	skmage_pois5-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	673	skmage_pois6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	674	pantherwoman5-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	675	pantherwoman6-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	676	sandleaper6-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	677	sandleaper7-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
1	1	678	wraith6-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	679	wraith7-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	680	wraith8-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	681	succubus6-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	682	succubus7-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	683	succubuswitch6-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	684	succubuswitch7-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	685	succubuswitch8-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	686	willowisp5-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	687	willowisp6-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	688	willowisp7-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	689	fallen6-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
1	1	690	fallen7-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
1	1	691	fallen8-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
1	1	692	fallenshaman6-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	693	fallenshaman7-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	694	fallenshaman8-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	695	skeleton6-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	696	skeleton7-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	697	batdemon6-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	698	batdemon7-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	699	bloodlord6-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	700	bloodlord7-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	701	scarab6-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	702	scarab7-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	703	fetish6-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	704	fetish7-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	705	fetish8-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
1	1	706	fetishblow6-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	707	fetishblow7-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	708	fetishblow8-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
1	1	709	fetishshaman6-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	710	fetishshaman7-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	711	fetishshaman8-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	712	baboon7-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	713	baboon8-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
1	1	714	unraveler6-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	715	unraveler7-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	716	unraveler8-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	717	unraveler9-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	718	zealot4-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
1	1	719	zealot5-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
1	1	720	cantor5-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
1	1	721	cantor6-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
1	1	722	vilemother4-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
1	1	723	vilemother5-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
1	1	724	vilechild4-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
1	1	725	vilechild5-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
1	1	726	sandmaggot6-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	727	maggotbaby6-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
1	1	728	maggotegg6-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
1	1	729	minion9-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	730	minion10-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	731	minion11-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	732	arach6-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	733	megademon4-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	734	megademon5-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	735	imp6-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	736	imp7-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	737	bonefetish6-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	738	bonefetish7-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
1	1	739	fingermage4-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
1	1	740	fingermage5-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
1	1	741	regurgitator4-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
1	1	742	vampire6-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	743	vampire7-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	744	vampire8-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	745	reanimatedhorde6-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	746	dkfig1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	747	dkfig2-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	748	dkmag1-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	749	dkmag2-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	750	mummy6-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	751	ubermephisto-Mephisto-UberMephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
1	1	752	uberdiablo-Diablo-UberDiablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
1	1	753	uberizual-izual-UberIzual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
1	1	754	uberandariel-Lilith-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
1	1	755	uberduriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
1	1	756	uberbaal-Baal Crab-UberBaal				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
1	1	757	demonspawner-Evil hut-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
1	1	758	demonhole-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
1	1	759	megademon6-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	760	dkmag3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	761	imp8-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	762	swarm5-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
1	1	763	sandmaggot7-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
1	1	764	arach7-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	765	scarab8-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	766	succubus8-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	767	succubuswitch9-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	768	corruptrogue6-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	769	cr_archer8-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	770	cr_lancer9-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
1	1	771	overseer6-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	772	skeleton8-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	773	sk_archer11-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
1	1	774	skmage_fire7-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
1	1	775	skmage_ltng7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
1	1	776	skmage_cold6-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
1	1	777	skmage_pois7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		POS	POS						0
1	1	778	vampire9-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
1	1	779	wraith9-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
1	1	780	willowisp8-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	781	Bishibosh-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	782	Bonebreak-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
1	1	783	Coldcrow-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
1	1	784	Rakanishu-SUPER UNIQUE				/Data/Global/Monsters	FA	NU	HTH		LIT				SWD		TCH	LIT										0
1	1	785	Treehead WoodFist-SUPER UNIQUE				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
1	1	786	Griswold-SUPER UNIQUE				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
1	1	787	The Countess-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	MED	LIT	MED	LIT	LIT	WHM			LIT	LIT									0
1	1	788	Pitspawn Fouldog-SUPER UNIQUE				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
1	1	789	Flamespike the Crawler-SUPER UNIQUE				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
1	1	790	Boneash-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
1	1	791	Radament-SUPER UNIQUE				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
1	1	792	Bloodwitch the Wild-SUPER UNIQUE				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
1	1	793	Fangskin-SUPER UNIQUE				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
1	1	794	Beetleburst-SUPER UNIQUE				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
1	1	795	Leatherarm-SUPER UNIQUE				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
1	1	796	Coldworm the Burrower-SUPER UNIQUE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
1	1	797	Fire Eye-SUPER UNIQUE				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
1	1	798	Dark Elder-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	799	The Summoner-SUPER UNIQUE				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
1	1	800	Ancient Kaa the Soulless-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	801	The Smith-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
1	1	802	Web Mage the Burning-SUPER UNIQUE				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
1	1	803	Witch Doctor Endugu-SUPER UNIQUE				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
1	1	804	Stormtree-SUPER UNIQUE				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
1	1	805	Sarina the Battlemaid-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
1	1	806	Icehawk Riftwing-SUPER UNIQUE				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
1	1	807	Ismail Vilehand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	808	Geleb Flamefinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	809	Bremm Sparkfist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	810	Toorc Icefist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	811	Wyand Voidfinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	812	Maffer Dragonhand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	813	Winged Death-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	814	The Tormentor-SUPER UNIQUE				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
1	1	815	Taintbreeder-SUPER UNIQUE				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
1	1	816	Riftwraith the Cannibal-SUPER UNIQUE				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
1	1	817	Infector of Souls-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	818	Lord De Seis-SUPER UNIQUE				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
1	1	819	Grand Vizier of Chaos-SUPER UNIQUE				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
1	1	820	The Cow King-SUPER UNIQUE				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
1	1	821	Corpsefire-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
1	1	822	The Feature Creep-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
1	1	823	Siege Boss-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	824	Ancient Barbarian 1-SUPER UNIQUE				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
1	1	825	Ancient Barbarian 2-SUPER UNIQUE				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
1	1	826	Ancient Barbarian 3-SUPER UNIQUE				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
1	1	827	Axe Dweller-SUPER UNIQUE				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
1	1	828	Bonesaw Breaker-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	829	Dac Farren-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	830	Megaflow Rectifier-SUPER UNIQUE				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
1	1	831	Eyeback Unleashed-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	832	Threash Socket-SUPER UNIQUE				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
1	1	833	Pindleskin-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
1	1	834	Snapchip Shatter-SUPER UNIQUE				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
1	1	835	Anodized Elite-SUPER UNIQUE				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
1	1	836	Vinvear Molech-SUPER UNIQUE				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
1	1	837	Sharp Tooth Sayer-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
1	1	838	Magma Torquer-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
1	1	839	Blaze Ripper-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
1	1	840	Frozenstein-SUPER UNIQUE				/Data/Global/Monsters	io	NU	HTH		LIT																	0
1	1	841	Nihlathak Boss-SUPER UNIQUE				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
1	1	842	Baal Subject 1-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
1	1	843	Baal Subject 2-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
1	1	844	Baal Subject 3-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
1	1	845	Baal Subject 4-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
1	1	846	Baal Subject 5-SUPER UNIQUE				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
1	2	0	rogue fountain (12)	12			/Data/Global/Objects	FN	NU	HTH		LIT																	0
1	2	1	torch 1 tiki (37)	37			/Data/Global/Objects	TO	ON	HTH		LIT																	0
1	2	2	Fire, rogue camp (39)	39			/Data/Global/Objects	RB	ON	HTH		LIT																	0
1	2	3	flag 1 (35)	35			/Data/Global/Objects	N1	NU	HTH		LIT																	0
1	2	4	flag 2 (36)	36			/Data/Global/Objects	N2	NU	HTH		LIT																	0
1	2	5	Chest, R Large (5)	5			/Data/Global/Objects	L1	NU	HTH		LIT																	0
1	2	6	Cairn Stone, Alpha (17)	17			/Data/Global/Objects	S1	NU	HTH		LIT																	0
1	2	7	Cairn Stone, Beta (18)	18			/Data/Global/Objects	S2	NU	HTH		LIT																	0
1	2	8	Cairn Stone, Gamma (19)	19			/Data/Global/Objects	S3	NU	HTH		LIT																	0
1	2	9	Cairn Stone, Delta (20)	20			/Data/Global/Objects	S4	NU	HTH		LIT																	0
1	2	10	Cairn Stone, Lambda (21)	21			/Data/Global/Objects	S5	NU	HTH		LIT																	0
1	2	11	Cairn Stone, Theta (inactive) (22)	22			/Data/Global/Objects	S6	NU	HTH		LIT																	0
1	2	12	Tree of Inifuss (30)	30			/Data/Global/Objects	IT	NU	HTH		LIT																	0
1	2	13	water effect 2 (70)	70			/Data/Global/Objects	4R	NU	HTH		LIT																	0
1	2	14	water effect 2 (70)	70			/Data/Global/Objects	4R	NU	HTH		LIT																	0
1	2	15	water effect 1 (69)	69			/Data/Global/Objects	3R	NU	HTH		LIT																	0
1	2	16	water effect 1 (69)	69			/Data/Global/Objects	3R	NU	HTH		LIT																	0
1	2	17	brazier (29)	29			/Data/Global/Objects	BR	OP	HTH		LIT							LIT										0
1	2	18	bloody fountain (31)	31			/Data/Global/Objects	BF	NU	HTH		LIT																	0
1	2	19	candles 1 (33)	33			/Data/Global/Objects	A1	NU	HTH		LIT																	0
1	2	20	candles 2 (34)	34			/Data/Global/Objects	A2	NU	HTH		LIT																	0
1	2	21	torch 1 tiki (37)	37			/Data/Global/Objects	TO	ON	HTH		LIT																	0
1	2	22	Invisible object (61)	61			/Data/Global/Objects	SS	NU	HTH		LIT																	0
1	2	23	Invisible river sound 1 (65)	65			/Data/Global/Objects	X1	NU	HTH		LIT																	0
1	2	24	Invisible river sound 2 (66)	66			/Data/Global/Objects	X2	NU	HTH		LIT																	0
1	2	25	The Moldy Tome (8)	8			/Data/Global/Objects	TT	NU	HTH		LIT																	0
1	2	26	Cain's Gibbet (26)	26			/Data/Global/Objects	GI	NU	HTH		LIT																	0
1	2	27	Undefiled Grave (28)	28			/Data/Global/Objects	HI	NU	HTH		LIT																	0
1	2	28	bubbling pool of blood (82)	82			/Data/Global/Objects	B2	NU	HTH		LIT																	0
1	2	29	Shrine (2)	2			/Data/Global/Objects	SF	NU	HTH		LIT																	0
1	2	30	Shrine, forest altar (81)	81			/Data/Global/Objects	AF	NU	HTH		LIT																	0
1	2	31	Shrine, healing well (84)	84			/Data/Global/Objects	HW	NU	HTH		LIT																	0
1	2	32	Shrine, horn (83)	83			/Data/Global/Objects	HS	NU	HTH		LIT																	0
1	2	33	invisible town sound (78)	78																									0
1	2	34	invisible object (61)	61			/Data/Global/Objects	SS	NU	HTH		LIT																	0
1	2	35	flies (103)	103			/Data/Global/Objects	FL	NU	HTH		LIT																	0
1	2	36	Horadric Malus (108)	108			/Data/Global/Objects	HM	NU	HTH		LIT																	0
1	2	37	Waypoint (119)	119			/Data/Global/Objects	WP	ON	HTH		LIT							LIT										0
1	2	38	error ? (-580)	580																									0
1	2	39	Well, pool wilderness (130)	130			/Data/Global/Objects	ZW	NU	HTH		LIT																	0
1	2	40	Hidden Stash, rock wilderness (159)	159			/Data/Global/Objects	CQ	OP	HTH		LIT																	0
1	2	41	Hiding Spot, cliff wilderness (163)	163			/Data/Global/Objects	CF	NU	HTH		LIT																	0
1	2	42	Hollow Log (169)	169			/Data/Global/Objects	CZ	NU	HTH		LIT																	0
1	2	43	Fire, small (160)	160			/Data/Global/Objects	FX	OP	HTH		LIT																	0
1	2	44	Fire, medium (161)	161			/Data/Global/Objects	FY	OP	HTH		LIT																	0
1	2	45	Fire, large (162)	162			/Data/Global/Objects	FZ	OP	HTH		LIT																	0
1	2	46	Armor Stand, 1R (104)	104			/Data/Global/Objects	A3	NU	HTH		LIT																	0
1	2	47	Armor Stand, 2L (105)	105			/Data/Global/Objects	A4	NU	HTH		LIT																	0
1	2	48	Weapon Rack, 1R (106)	106			/Data/Global/Objects	W1	NU	HTH		LIT																	0
1	2	49	Weapon Rack, 2L (107)	107			/Data/Global/Objects	W2	NU	HTH		LIT																	0
1	2	50	Bookshelf L (179)	179			/Data/Global/Objects	B4	NU	HTH		LIT																	0
1	2	51	Bookshelf R (180)	180			/Data/Global/Objects	B5	NU	HTH		LIT																	0
1	2	52	Waypoint (119)	119			/Data/Global/Objects	WP	ON	HTH		LIT							LIT										0
1	2	53	Waypoint, wilderness (157)	157			/Data/Global/Objects	WN	ON	HTH		LIT							LIT										0
1	2	54	Bed R (247)	247			/Data/Global/Objects	QA	OP	HTH		LIT																	0
1	2	55	Bed L (248)	248			/Data/Global/Objects	QB	OP	HTH		LIT																	0
1	2	56	Hidden Stash, rock wilderness (155)	155			/Data/Global/Objects	C7	OP	HTH		LIT																	0
1	2	57	Loose Rock, wilderness (174)	174			/Data/Global/Objects	RY	OP	HTH		LIT																	0
1	2	58	Loose Boulder, wilderness (175)	175			/Data/Global/Objects	RZ	OP	HTH		LIT																	0
1	2	59	Chest, R Large (139)	139			/Data/Global/Objects	Q1	OP	HTH		LIT																	0
1	2	60	Chest, R Tallskinney (140)	140			/Data/Global/Objects	Q2	OP	HTH		LIT																	0
1	2	61	Chest, R Med (141)	141			/Data/Global/Objects	Q3	OP	HTH		LIT																	0
1	2	62	Chest, L (144)	144			/Data/Global/Objects	Q6	OP	HTH		LIT																	0
1	2	63	Chest, L Large (6)	6			/Data/Global/Objects	L2	OP	HTH		LIT																	0
1	2	64	Chest, 1L general (240)	240			/Data/Global/Objects	CY	OP	HTH		LIT																	0
1	2	65	Chest, 2R general (241)	241			/Data/Global/Objects	CX	OP	HTH		LIT																	0
1	2	66	Chest, 3R general (242)	242			/Data/Global/Objects	CU	OP	HTH		LIT																	0
1	2	67	Dead Rogue, 1 (54)	54			/Data/Global/Objects	Z1	NU	HTH		LIT																	0
1	2	68	Dead Rogue, 2 (55)	55			/Data/Global/Objects	Z2	NU	HTH		LIT																	0
1	2	69	Dead Rogue, rolling (56)	56			/Data/Global/Objects	Z5	OP	HTH		LIT																	0
1	2	70	Dead Rogue, on a stick 1 (57)	57			/Data/Global/Objects	Z3	OP	HTH		LIT																	0
1	2	71	Dead Rogue, on a stick 2 (58)	58			/Data/Global/Objects	Z4	OP	HTH		LIT																	0
1	2	72	Skeleton (171)	171			/Data/Global/Objects	SX	OP	HTH		LIT																	0
1	2	73	Guard Corpse, on a stick (178)	178			/Data/Global/Objects	GS	OP	HTH		LIT																	0
1	2	74	Body, burning town 1 (239)	239			/Data/Global/Objects	BZ	NU	HTH		LIT																	0
1	2	75	Body, burning town 2 (245)	245			/Data/Global/Objects	BY	NU	HTH		LIT							LIT										0
1	2	76	A Trap, exploding cow (250)	250			/Data/Global/Objects	EW	OP	HTH		LIT																	0
1	2	77	Well, fountain 1 (111)	111			/Data/Global/Objects	F3	NU	HTH		LIT																	0
1	2	78	Well, cavewell caves (138)	138			/Data/Global/Objects	ZY	NU	HTH		LIT																	0
1	2	79	Well, cathedralwell inside (132)	132			/Data/Global/Objects	ZC	NU	HTH		LIT																	0
1	2	80	Shrine, mana well 1 (164)	164			/Data/Global/Objects	MB	OP	HTH		LIT																	0
1	2	81	Shrine, mana well 2 (165)	165			/Data/Global/Objects	MD	OP	HTH		LIT																	0
1	2	82	Shrine, healthorama (77)	77			/Data/Global/Objects	SH	OP	HTH		LIT																	0
1	2	83	Shrine, bull shrine, health, tombs (85)	85			/Data/Global/Objects	BC	OP	HTH		LIT																	0
1	2	84	stele,magic shrine, stone, desert (86)	86			/Data/Global/Objects	SG	OP	HTH		LIT																	0
1	2	85	Shrine, cathedral (262)	262			/Data/Global/Objects	S0	OP	HTH		LIT							LIT										0
1	2	86	Shrine, jail 1 (263)	263			/Data/Global/Objects	JB	OP	HTH		LIT							LIT										0
1	2	87	Shrine, jail 2 (264)	264			/Data/Global/Objects	JD	OP	HTH		LIT							LIT										0
1	2	88	Shrine, jail 3 (265)	265			/Data/Global/Objects	JF	OP	HTH		LIT							LIT										0
1	2	89	Casket, 1 R (50)	50			/Data/Global/Objects	C1	OP	HTH		LIT																	0
1	2	90	Casket, 2 L (51)	51			/Data/Global/Objects	C2	OP	HTH		LIT																	0
1	2	91	Casket, 3 (79)	79			/Data/Global/Objects	C3	OP	HTH		LIT																	0
1	2	92	Casket, 4 (53)	53			/Data/Global/Objects	C4	OP	HTH		LIT																	0
1	2	93	Casket, 5 (1)	1			/Data/Global/Objects	C5	OP	HTH		LIT																	0
1	2	94	Casket, 6 (3)	3			/Data/Global/Objects	C6	OP	HTH		LIT																	0
1	2	95	Barrel (7)	7			/Data/Global/Objects	B1	OP	HTH		LIT																	0
1	2	96	Crate (46)	46			/Data/Global/Objects	CT	OP	HTH		LIT																	0
1	2	97	torch 2 wall (38)	38			/Data/Global/Objects	WT	ON	HTH		LIT																	0
1	2	98	cabin stool (256)	256			/Data/Global/Objects	S9	NU	HTH		LIT																	0
1	2	99	cabin wood (257)	257			/Data/Global/Objects	WG	NU	HTH		LIT																	0
1	2	100	cabin wood more (258)	258			/Data/Global/Objects	WH	NU	HTH		LIT																	0
1	2	101	Door, secret 1 (129)	129			/Data/Global/Objects	H2	OP	HTH		LIT																	0
1	2	102	Your Private Stash (267)	267			/Data/Global/Objects	B6	NU	HTH		LIT																	0
1	2	103	Wirt's body (268)	268			/Data/Global/Objects	BP	NU	HTH		LIT																	0
1	2	104	gold placeholder (269)	269			/Data/Global/Objects	1G																					0
1	2	105	ACT 1 TABLE DO NOT USE SKIP IT	581																									0
1	2	106	hell light source 1 (351)	351			/Data/Global/Objects	SS																					0
1	2	107	hell light source 2 (352)	352			/Data/Global/Objects	SS																					0
1	2	108	hell light source 3 (353)	353			/Data/Global/Objects	SS																					0
1	2	109	fog water (374)	374			/Data/Global/Objects	UD	NU	HTH		LIT																	0
1	2	110	cain start position (385)	385		1	/Data/Global/Monsters	DC	NU	HTH		LIT																	0
1	2	111	Chest, sparkly (397)	397			/Data/Global/Objects	YF	OP	HTH		LIT																	0
1	2	112	Red Portal (321)	321			/Data/Global/Objects	PP	NU	HTH	LIT	LIT																	0
1	2	113	ACT 1 TABLE DO NOT USE SKIP IT	0			/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									0
1	2	114	Myhrginoc's Book of Lore	8			/Data/Global/Objects	TT	NU	HTH		LIT																	0
1	2	115	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	116	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	117	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	118	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	119	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	120	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	121	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	122	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	123	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	124	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	125	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	126	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	127	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	128	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	129	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	130	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	131	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	132	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	133	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	134	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	135	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	136	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	137	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	138	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	139	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	140	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	141	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	142	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	143	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	144	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	145	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	146	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	147	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	148	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	149	ACT 1 TABLE DO NOT USE SKIP IT	0																									0
1	2	150	Dummy-test data SKIPT IT				/Data/Global/Objects	NU0																					
1	2	151	Casket-Casket #5				/Data/Global/Objects	C5	OP	HTH		LIT																	
1	2	152	Shrine-Shrine				/Data/Global/Objects	SF	OP	HTH		LIT																	
1	2	153	Casket-Casket #6				/Data/Global/Objects	C6	OP	HTH		LIT																	
1	2	154	LargeUrn-Urn #1				/Data/Global/Objects	U1	OP	HTH		LIT																	
1	2	155	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
1	2	156	chest-LargeChestL				/Data/Global/Objects	L2	OP	HTH		LIT																	
1	2	157	Barrel-Barrel				/Data/Global/Objects	B1	OP	HTH		LIT																	
1	2	158	TowerTome-Tower Tome				/Data/Global/Objects	TT	OP	HTH		LIT																	
1	2	159	Urn-Urn #2				/Data/Global/Objects	U2	OP	HTH		LIT																	
1	2	160	Dummy-Bench				/Data/Global/Objects	BE	NU	HTH		LIT																	
1	2	161	Barrel-BarrelExploding				/Data/Global/Objects	BX	OP	HTH		LIT							LIT	LIT									
1	2	162	Dummy-RogueFountain				/Data/Global/Objects	FN	NU	HTH		LIT																	
1	2	163	Door-Door Gate Left				/Data/Global/Objects	D1	OP	HTH		LIT																	
1	2	164	Door-Door Gate Right				/Data/Global/Objects	D2	OP	HTH		LIT																	
1	2	165	Door-Door Wooden Left				/Data/Global/Objects	D3	OP	HTH		LIT																	
1	2	166	Door-Door Wooden Right				/Data/Global/Objects	D4	OP	HTH		LIT																	
1	2	167	StoneAlpha-StoneAlpha				/Data/Global/Objects	S1	OP	HTH		LIT																	
1	2	168	StoneBeta-StoneBeta				/Data/Global/Objects	S2	OP	HTH		LIT																	
1	2	169	StoneGamma-StoneGamma				/Data/Global/Objects	S3	OP	HTH		LIT																	
1	2	170	StoneDelta-StoneDelta				/Data/Global/Objects	S4	OP	HTH		LIT																	
1	2	171	StoneLambda-StoneLambda				/Data/Global/Objects	S5	OP	HTH		LIT																	
1	2	172	StoneTheta-StoneTheta				/Data/Global/Objects	S6	OP	HTH		LIT																	
1	2	173	Door-Door Courtyard Left				/Data/Global/Objects	D5	OP	HTH		LIT																	
1	2	174	Door-Door Courtyard Right				/Data/Global/Objects	D6	OP	HTH		LIT																	
1	2	175	Door-Door Cathedral Double				/Data/Global/Objects	D7	OP	HTH		LIT																	
1	2	176	Gibbet-Cain's Been Captured				/Data/Global/Objects	GI	OP	HTH		LIT																	
1	2	177	Door-Door Monastery Double Right				/Data/Global/Objects	D8	OP	HTH		LIT																	
1	2	178	HoleAnim-Hole in Ground				/Data/Global/Objects	HI	OP	HTH		LIT																	
1	2	179	Dummy-Brazier				/Data/Global/Objects	BR	ON	HTH		LIT							LIT										
1	2	180	Inifuss-inifuss tree				/Data/Global/Objects	IT	NU	HTH		LIT																	
1	2	181	Dummy-Fountain				/Data/Global/Objects	BF	NU	HTH		LIT																	
1	2	182	Dummy-crucifix				/Data/Global/Objects	CL	NU	HTH		LIT																	
1	2	183	Dummy-Candles1				/Data/Global/Objects	A1	NU	HTH		LIT																	
1	2	184	Dummy-Candles2				/Data/Global/Objects	A2	NU	HTH		LIT																	
1	2	185	Dummy-Standard1				/Data/Global/Objects	N1	NU	HTH		LIT																	
1	2	186	Dummy-Standard2				/Data/Global/Objects	N2	NU	HTH		LIT																	
1	2	187	Dummy-Torch1 Tiki				/Data/Global/Objects	TO	ON	HTH		LIT																	
1	2	188	Dummy-Torch2 Wall				/Data/Global/Objects	WT	ON	HTH		LIT																	
1	2	189	fire-RogueBonfire				/Data/Global/Objects	RB	ON	HTH		LIT																	
1	2	190	Dummy-River1				/Data/Global/Objects	R1	NU	HTH		LIT																	
1	2	191	Dummy-River2				/Data/Global/Objects	R2	NU	HTH		LIT																	
1	2	192	Dummy-River3				/Data/Global/Objects	R3	NU	HTH		LIT																	
1	2	193	Dummy-River4				/Data/Global/Objects	R4	NU	HTH		LIT																	
1	2	194	Dummy-River5				/Data/Global/Objects	R5	NU	HTH		LIT																	
1	2	195	AmbientSound-ambient sound generator				/Data/Global/Objects	S1	OP	HTH		LIT																	
1	2	196	Crate-Crate				/Data/Global/Objects	CT	OP	HTH		LIT																	
1	2	197	Door-Andariel's Door				/Data/Global/Objects	AD	NU	HTH		LIT																	
1	2	198	Dummy-RogueTorch				/Data/Global/Objects	T1	NU	HTH		LIT																	
1	2	199	Dummy-RogueTorch				/Data/Global/Objects	T2	NU	HTH		LIT																	
1	2	200	Casket-CasketR				/Data/Global/Objects	C1	OP	HTH		LIT																	
1	2	201	Casket-CasketL				/Data/Global/Objects	C2	OP	HTH		LIT																	
1	2	202	Urn-Urn #3				/Data/Global/Objects	U3	OP	HTH		LIT																	
1	2	203	Casket-Casket				/Data/Global/Objects	C4	OP	HTH		LIT																	
1	2	204	RogueCorpse-Rogue corpse 1				/Data/Global/Objects	Z1	NU	HTH		LIT																	
1	2	205	RogueCorpse-Rogue corpse 2				/Data/Global/Objects	Z2	NU	HTH		LIT																	
1	2	206	RogueCorpse-rolling rogue corpse				/Data/Global/Objects	Z5	OP	HTH		LIT																	
1	2	207	CorpseOnStick-rogue on a stick 1				/Data/Global/Objects	Z3	OP	HTH		LIT																	
1	2	208	CorpseOnStick-rogue on a stick 2				/Data/Global/Objects	Z4	OP	HTH		LIT																	
1	2	209	Portal-Town portal				/Data/Global/Objects	TP	ON	HTH	LIT	LIT																	
1	2	210	Portal-Permanent town portal				/Data/Global/Objects	PP	ON	HTH	LIT	LIT																	
1	2	211	Dummy-Invisible object				/Data/Global/Objects	SS																					
1	2	212	Door-Door Cathedral Left				/Data/Global/Objects	D9	OP	HTH		LIT																	
1	2	213	Door-Door Cathedral Right				/Data/Global/Objects	DA	OP	HTH		LIT																	
1	2	214	Door-Door Wooden Left #2				/Data/Global/Objects	DB	OP	HTH		LIT																	
1	2	215	Dummy-invisible river sound1				/Data/Global/Objects	X1																					
1	2	216	Dummy-invisible river sound2				/Data/Global/Objects	X2																					
1	2	217	Dummy-ripple				/Data/Global/Objects	1R	NU	HTH		LIT																	
1	2	218	Dummy-ripple				/Data/Global/Objects	2R	NU	HTH		LIT																	
1	2	219	Dummy-ripple				/Data/Global/Objects	3R	NU	HTH		LIT																	
1	2	220	Dummy-ripple				/Data/Global/Objects	4R	NU	HTH		LIT																	
1	2	221	Dummy-forest night sound #1				/Data/Global/Objects	F1																					
1	2	222	Dummy-forest night sound #2				/Data/Global/Objects	F2																					
1	2	223	Dummy-yeti dung				/Data/Global/Objects	YD	NU	HTH		LIT																	
1	2	224	TrappDoor-Trap Door				/Data/Global/Objects	TD	ON	HTH		LIT																	
1	2	225	Door-Door by Dock, Act 2				/Data/Global/Objects	DD	ON	HTH		LIT																	
1	2	226	Dummy-sewer drip				/Data/Global/Objects	SZ																					
1	2	227	Shrine-healthorama				/Data/Global/Objects	SH	OP	HTH		LIT																	
1	2	228	Dummy-invisible town sound				/Data/Global/Objects	TA																					
1	2	229	Casket-casket #3				/Data/Global/Objects	C3	OP	HTH		LIT																	
1	2	230	Obelisk-obelisk				/Data/Global/Objects	OB	OP	HTH		LIT																	
1	2	231	Shrine-forest altar				/Data/Global/Objects	AF	OP	HTH		LIT																	
1	2	232	Dummy-bubbling pool of blood				/Data/Global/Objects	B2	NU	HTH		LIT																	
1	2	233	Shrine-horn shrine				/Data/Global/Objects	HS	OP	HTH		LIT																	
1	2	234	Shrine-healing well				/Data/Global/Objects	HW	OP	HTH		LIT																	
1	2	235	Shrine-bull shrine,health, tombs				/Data/Global/Objects	BC	OP	HTH		LIT																	
1	2	236	Dummy-stele,magic shrine, stone, desert				/Data/Global/Objects	SG	OP	HTH		LIT																	
1	2	237	Chest3-tombchest 1, largechestL				/Data/Global/Objects	CA	OP	HTH		LIT																	
1	2	238	Chest3-tombchest 2 largechestR				/Data/Global/Objects	CB	OP	HTH		LIT																	
1	2	239	Sarcophagus-mummy coffinL, tomb				/Data/Global/Objects	MC	OP	HTH		LIT																	
1	2	240	Obelisk-desert obelisk				/Data/Global/Objects	DO	OP	HTH		LIT																	
1	2	241	Door-tomb door left				/Data/Global/Objects	TL	OP	HTH		LIT																	
1	2	242	Door-tomb door right				/Data/Global/Objects	TR	OP	HTH		LIT																	
1	2	243	Shrine-mana shrineforinnerhell				/Data/Global/Objects	iz	OP	HTH		LIT							LIT										
1	2	244	LargeUrn-Urn #4				/Data/Global/Objects	U4	OP	HTH		LIT																	
1	2	245	LargeUrn-Urn #5				/Data/Global/Objects	U5	OP	HTH		LIT																	
1	2	246	Shrine-health shrineforinnerhell				/Data/Global/Objects	iy	OP	HTH		LIT							LIT										
1	2	247	Shrine-innershrinehell				/Data/Global/Objects	ix	OP	HTH		LIT							LIT										
1	2	248	Door-tomb door left 2				/Data/Global/Objects	TS	OP	HTH		LIT																	
1	2	249	Door-tomb door right 2				/Data/Global/Objects	TU	OP	HTH		LIT																	
1	2	250	Duriel's Lair-Portal to Duriel's Lair				/Data/Global/Objects	SJ	OP	HTH		LIT																	
1	2	251	Dummy-Brazier3				/Data/Global/Objects	B3	OP	HTH		LIT							LIT										
1	2	252	Dummy-Floor brazier				/Data/Global/Objects	FB	ON	HTH		LIT							LIT										
1	2	253	Dummy-flies				/Data/Global/Objects	FL	NU	HTH		LIT																	
1	2	254	ArmorStand-Armor Stand 1R				/Data/Global/Objects	A3	NU	HTH		LIT																	
1	2	255	ArmorStand-Armor Stand 2L				/Data/Global/Objects	A4	NU	HTH		LIT																	
1	2	256	WeaponRack-Weapon Rack 1R				/Data/Global/Objects	W1	NU	HTH		LIT																	
1	2	257	WeaponRack-Weapon Rack 2L				/Data/Global/Objects	W2	NU	HTH		LIT																	
1	2	258	Malus-Malus				/Data/Global/Objects	HM	NU	HTH		LIT																	
1	2	259	Shrine-palace shrine, healthR, harom, arcane Sanctuary				/Data/Global/Objects	P2	OP	HTH		LIT																	
1	2	260	not used-drinker				/Data/Global/Objects	n5	S1	HTH		LIT																	
1	2	261	well-Fountain 1				/Data/Global/Objects	F3	OP	HTH		LIT																	
1	2	262	not used-gesturer				/Data/Global/Objects	n6	S1	HTH		LIT																	
1	2	263	well-Fountain 2, well, desert, tomb				/Data/Global/Objects	F4	OP	HTH		LIT																	
1	2	264	not used-turner				/Data/Global/Objects	n7	S1	HTH		LIT																	
1	2	265	well-Fountain 3				/Data/Global/Objects	F5	OP	HTH		LIT																	
1	2	266	Shrine-snake woman, magic shrine, tomb, arcane sanctuary				/Data/Global/Objects	SN	OP	HTH		LIT							LIT										
1	2	267	Dummy-jungle torch				/Data/Global/Objects	JT	ON	HTH		LIT							LIT										
1	2	268	Well-Fountain 4				/Data/Global/Objects	F6	OP	HTH		LIT																	
1	2	269	Waypoint-waypoint portal				/Data/Global/Objects	wp	ON	HTH		LIT							LIT										
1	2	270	Dummy-healthshrine, act 3, dungeun				/Data/Global/Objects	dj	OP	HTH		LIT																	
1	2	271	jerhyn-placeholder #1				/Data/Global/Objects	ss																					
1	2	272	jerhyn-placeholder #2				/Data/Global/Objects	ss																					
1	2	273	Shrine-innershrinehell2				/Data/Global/Objects	iw	OP	HTH		LIT							LIT										
1	2	274	Shrine-innershrinehell3				/Data/Global/Objects	iv	OP	HTH		LIT																	
1	2	275	hidden stash-ihobject3 inner hell				/Data/Global/Objects	iu	OP	HTH		LIT																	
1	2	276	skull pile-skullpile inner hell				/Data/Global/Objects	is	OP	HTH		LIT																	
1	2	277	hidden stash-ihobject5 inner hell				/Data/Global/Objects	ir	OP	HTH		LIT																	
1	2	278	hidden stash-hobject4 inner hell				/Data/Global/Objects	hg	OP	HTH		LIT																	
1	2	279	Door-secret door 1				/Data/Global/Objects	h2	OP	HTH		LIT																	
1	2	280	Well-pool act 1 wilderness				/Data/Global/Objects	zw	NU	HTH		LIT																	
1	2	281	Dummy-vile dog afterglow				/Data/Global/Objects	9b	OP	HTH		LIT																	
1	2	282	Well-cathedralwell act 1 inside				/Data/Global/Objects	zc	NU	HTH		LIT																	
1	2	283	shrine-shrine1_arcane sanctuary				/Data/Global/Objects	xx																					
1	2	284	shrine-dshrine2 act 2 shrine				/Data/Global/Objects	zs	OP	HTH		LIT							LIT										
1	2	285	shrine-desertshrine3 act 2 shrine				/Data/Global/Objects	zr	OP	HTH		LIT																	
1	2	286	shrine-dshrine1 act 2 shrine				/Data/Global/Objects	zd	OP	HTH		LIT																	
1	2	287	Well-desertwell act 2 well, desert, tomb				/Data/Global/Objects	zl	NU	HTH		LIT																	
1	2	288	Well-cavewell act 1 caves 				/Data/Global/Objects	zy	NU	HTH		LIT																	
1	2	289	chest-chest-r-large act 1				/Data/Global/Objects	q1	OP	HTH		LIT																	
1	2	290	chest-chest-r-tallskinney act 1				/Data/Global/Objects	q2	OP	HTH		LIT																	
1	2	291	chest-chest-r-med act 1				/Data/Global/Objects	q3	OP	HTH		LIT																	
1	2	292	jug-jug1 act 2, desert				/Data/Global/Objects	q4	OP	HTH		LIT																	
1	2	293	jug-jug2 act 2, desert				/Data/Global/Objects	q5	OP	HTH		LIT																	
1	2	294	chest-Lchest1 act 1				/Data/Global/Objects	q6	OP	HTH		LIT																	
1	2	295	Waypoint-waypointi inner hell				/Data/Global/Objects	wi	ON	HTH		LIT							LIT										
1	2	296	chest-dchest2R act 2, desert, tomb, chest-r-med				/Data/Global/Objects	q9	OP	HTH		LIT																	
1	2	297	chest-dchestr act 2, desert, tomb, chest -r large				/Data/Global/Objects	q7	OP	HTH		LIT																	
1	2	298	chest-dchestL act 2, desert, tomb chest l large				/Data/Global/Objects	q8	OP	HTH		LIT																	
1	2	299	taintedsunaltar-tainted sun altar quest				/Data/Global/Objects	za	OP	HTH		LIT							LIT										
1	2	300	shrine-dshrine1 act 2 , desert				/Data/Global/Objects	zv	NU	HTH		LIT							LIT	LIT									
1	2	301	shrine-dshrine4 act 2, desert				/Data/Global/Objects	ze	OP	HTH		LIT							LIT										
1	2	302	orifice-Where you place the Horadric staff				/Data/Global/Objects	HA	NU	HTH		LIT																	
1	2	303	Door-tyrael's door				/Data/Global/Objects	DX	OP	HTH		LIT																	
1	2	304	corpse-guard corpse				/Data/Global/Objects	GC	OP	HTH		LIT																	
1	2	305	hidden stash-rock act 1 wilderness				/Data/Global/Objects	c7	OP	HTH		LIT																	
1	2	306	Waypoint-waypoint act 2				/Data/Global/Objects	wm	ON	HTH		LIT							LIT										
1	2	307	Waypoint-waypoint act 1 wilderness				/Data/Global/Objects	wn	ON	HTH		LIT							LIT										
1	2	308	skeleton-corpse				/Data/Global/Objects	cp	OP	HTH		LIT																	
1	2	309	hidden stash-rockb act 1 wilderness				/Data/Global/Objects	cq	OP	HTH		LIT																	
1	2	310	fire-fire small				/Data/Global/Objects	FX	NU	HTH		LIT																	
1	2	311	fire-fire medium				/Data/Global/Objects	FY	NU	HTH		LIT																	
1	2	312	fire-fire large				/Data/Global/Objects	FZ	NU	HTH		LIT																	
1	2	313	hiding spot-cliff act 1 wilderness				/Data/Global/Objects	cf	NU	HTH		LIT																	
1	2	314	Shrine-mana well1				/Data/Global/Objects	MB	OP	HTH		LIT																	
1	2	315	Shrine-mana well2				/Data/Global/Objects	MD	OP	HTH		LIT																	
1	2	316	Shrine-mana well3, act 2, tomb, 				/Data/Global/Objects	MF	OP	HTH		LIT																	
1	2	317	Shrine-mana well4, act 2, harom				/Data/Global/Objects	MH	OP	HTH		LIT																	
1	2	318	Shrine-mana well5				/Data/Global/Objects	MJ	OP	HTH		LIT																	
1	2	319	hollow log-log				/Data/Global/Objects	cz	NU	HTH		LIT																	
1	2	320	Shrine-jungle healwell act 3				/Data/Global/Objects	JH	OP	HTH		LIT																	
1	2	321	skeleton-corpseb				/Data/Global/Objects	sx	OP	HTH		LIT																	
1	2	322	Shrine-health well, health shrine, desert				/Data/Global/Objects	Mk	OP	HTH		LIT																	
1	2	323	Shrine-mana well7, mana shrine, desert				/Data/Global/Objects	Mi	OP	HTH		LIT																	
1	2	324	loose rock-rockc act 1 wilderness				/Data/Global/Objects	RY	OP	HTH		LIT																	
1	2	325	loose boulder-rockd act 1 wilderness				/Data/Global/Objects	RZ	OP	HTH		LIT																	
1	2	326	chest-chest-L-med				/Data/Global/Objects	c8	OP	HTH		LIT																	
1	2	327	chest-chest-L-large				/Data/Global/Objects	c9	OP	HTH		LIT																	
1	2	328	GuardCorpse-guard on a stick, desert, tomb, harom				/Data/Global/Objects	GS	OP	HTH		LIT																	
1	2	329	bookshelf-bookshelf1				/Data/Global/Objects	b4	OP	HTH		LIT																	
1	2	330	bookshelf-bookshelf2				/Data/Global/Objects	b5	OP	HTH		LIT																	
1	2	331	chest-jungle chest act 3				/Data/Global/Objects	JC	OP	HTH		LIT																	
1	2	332	coffin-tombcoffin				/Data/Global/Objects	tm	OP	HTH		LIT																	
1	2	333	chest-chest-L-med, jungle				/Data/Global/Objects	jz	OP	HTH		LIT																	
1	2	334	Shrine-jungle shrine2				/Data/Global/Objects	jy	OP	HTH		LIT							LIT	LIT									
1	2	335	stash-jungle object act3				/Data/Global/Objects	jx	OP	HTH		LIT																	
1	2	336	stash-jungle object act3				/Data/Global/Objects	jw	OP	HTH		LIT																	
1	2	337	stash-jungle object act3				/Data/Global/Objects	jv	OP	HTH		LIT																	
1	2	338	stash-jungle object act3				/Data/Global/Objects	ju	OP	HTH		LIT																	
1	2	339	Dummy-cain portal				/Data/Global/Objects	tP	OP	HTH	LIT	LIT																	
1	2	340	Shrine-jungle shrine3 act 3				/Data/Global/Objects	js	OP	HTH		LIT							LIT										
1	2	341	Shrine-jungle shrine4 act 3				/Data/Global/Objects	jr	OP	HTH		LIT							LIT										
1	2	342	teleport pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
1	2	343	LamTome-Lam Esen's Tome				/Data/Global/Objects	ab	OP	HTH		LIT																	
1	2	344	stair-stairsl				/Data/Global/Objects	sl																					
1	2	345	stair-stairsr				/Data/Global/Objects	sv																					
1	2	346	a trap-test data floortrap				/Data/Global/Objects	a5	OP	HTH		LIT																	
1	2	347	Shrine-jungleshrine act 3				/Data/Global/Objects	jq	OP	HTH		LIT							LIT										
1	2	348	chest-chest-L-tallskinney, general chest r?				/Data/Global/Objects	c0	OP	HTH		LIT																	
1	2	349	Shrine-mafistoshrine				/Data/Global/Objects	mz	OP	HTH		LIT							LIT										
1	2	350	Shrine-mafistoshrine				/Data/Global/Objects	my	OP	HTH		LIT							LIT										
1	2	351	Shrine-mafistoshrine				/Data/Global/Objects	mx	NU	HTH		LIT							LIT										
1	2	352	Shrine-mafistomana				/Data/Global/Objects	mw	OP	HTH		LIT							LIT										
1	2	353	stash-mafistolair				/Data/Global/Objects	mv	OP	HTH		LIT																	
1	2	354	stash-box				/Data/Global/Objects	mu	OP	HTH		LIT																	
1	2	355	stash-altar				/Data/Global/Objects	mt	OP	HTH		LIT																	
1	2	356	Shrine-mafistohealth				/Data/Global/Objects	mr	OP	HTH		LIT							LIT										
1	2	357	dummy-water rocks in act 3 wrok				/Data/Global/Objects	rw	NU	HTH		LIT																	
1	2	358	Basket-basket 1				/Data/Global/Objects	bd	OP	HTH		LIT																	
1	2	359	Basket-basket 2				/Data/Global/Objects	bj	OP	HTH		LIT																	
1	2	360	Dummy-water logs in act 3  ne logw				/Data/Global/Objects	lw	NU	HTH		LIT																	
1	2	361	Dummy-water rocks girl in act 3 wrob				/Data/Global/Objects	wb	NU	HTH		LIT																	
1	2	362	Dummy-bubbles in act3 water				/Data/Global/Objects	yb	NU	HTH		LIT																	
1	2	363	Dummy-water logs in act 3 logx				/Data/Global/Objects	wd	NU	HTH		LIT																	
1	2	364	Dummy-water rocks in act 3 rokb				/Data/Global/Objects	wc	NU	HTH		LIT																	
1	2	365	Dummy-water rocks girl in act 3 watc				/Data/Global/Objects	we	NU	HTH		LIT																	
1	2	366	Dummy-water rocks in act 3 waty				/Data/Global/Objects	wy	NU	HTH		LIT																	
1	2	367	Dummy-water logs in act 3  logz				/Data/Global/Objects	lx	NU	HTH		LIT																	
1	2	368	Dummy-web covered tree 1				/Data/Global/Objects	w3	NU	HTH		LIT							LIT										
1	2	369	Dummy-web covered tree 2				/Data/Global/Objects	w4	NU	HTH		LIT							LIT										
1	2	370	Dummy-web covered tree 3				/Data/Global/Objects	w5	NU	HTH		LIT							LIT										
1	2	371	Dummy-web covered tree 4				/Data/Global/Objects	w6	NU	HTH		LIT							LIT										
1	2	372	pillar-hobject1				/Data/Global/Objects	70	OP	HTH		LIT																	
1	2	373	cocoon-cacoon				/Data/Global/Objects	CN	OP	HTH		LIT																	
1	2	374	cocoon-cacoon 2				/Data/Global/Objects	CC	OP	HTH		LIT																	
1	2	375	skullpile-hobject1				/Data/Global/Objects	ib	OP	HTH		LIT																	
1	2	376	Shrine-outershrinehell				/Data/Global/Objects	ia	OP	HTH		LIT							LIT										
1	2	377	dummy-water rock girl act 3  nw  blgb				/Data/Global/Objects	QX	NU	HTH		LIT																	
1	2	378	dummy-big log act 3  sw blga				/Data/Global/Objects	qw	NU	HTH		LIT																	
1	2	379	door-slimedoor1				/Data/Global/Objects	SQ	OP	HTH		LIT																	
1	2	380	door-slimedoor2				/Data/Global/Objects	SY	OP	HTH		LIT																	
1	2	381	Shrine-outershrinehell2				/Data/Global/Objects	ht	OP	HTH		LIT							LIT										
1	2	382	Shrine-outershrinehell3				/Data/Global/Objects	hq	OP	HTH		LIT																	
1	2	383	pillar-hobject2				/Data/Global/Objects	hv	OP	HTH		LIT																	
1	2	384	dummy-Big log act 3 se blgc 				/Data/Global/Objects	Qy	NU	HTH		LIT																	
1	2	385	dummy-Big log act 3 nw blgd				/Data/Global/Objects	Qz	NU	HTH		LIT																	
1	2	386	Shrine-health wellforhell				/Data/Global/Objects	ho	OP	HTH		LIT																	
1	2	387	Waypoint-act3waypoint town				/Data/Global/Objects	wz	ON	HTH		LIT							LIT										
1	2	388	Waypoint-waypointh				/Data/Global/Objects	wv	ON	HTH		LIT							LIT										
1	2	389	body-burning town				/Data/Global/Objects	bz	ON	HTH		LIT							LIT										
1	2	390	chest-gchest1L general				/Data/Global/Objects	cy	OP	HTH		LIT																	
1	2	391	chest-gchest2R general				/Data/Global/Objects	cx	OP	HTH		LIT																	
1	2	392	chest-gchest3R general				/Data/Global/Objects	cu	OP	HTH		LIT																	
1	2	393	chest-glchest3L general				/Data/Global/Objects	cd	OP	HTH		LIT																	
1	2	394	ratnest-sewers				/Data/Global/Objects	rn	OP	HTH		LIT																	
1	2	395	body-burning town				/Data/Global/Objects	by	NU	HTH		LIT							LIT										
1	2	396	ratnest-sewers				/Data/Global/Objects	ra	OP	HTH		LIT																	
1	2	397	bed-bed act 1				/Data/Global/Objects	qa	OP	HTH		LIT																	
1	2	398	bed-bed act 1				/Data/Global/Objects	qb	OP	HTH		LIT																	
1	2	399	manashrine-mana wellforhell				/Data/Global/Objects	hn	OP	HTH		LIT							LIT										
1	2	400	a trap-exploding cow  for Tristan and ACT 3 only??Very Rare  1 or 2				/Data/Global/Objects	ew	OP	HTH		LIT																	
1	2	401	gidbinn altar-gidbinn altar				/Data/Global/Objects	ga	ON	HTH		LIT							LIT										
1	2	402	gidbinn-gidbinn decoy				/Data/Global/Objects	gd	ON	HTH		LIT							LIT										
1	2	403	Dummy-diablo right light				/Data/Global/Objects	11	NU	HTH		LIT																	
1	2	404	Dummy-diablo left light				/Data/Global/Objects	12	NU	HTH		LIT																	
1	2	405	Dummy-diablo start point				/Data/Global/Objects	ss																					
1	2	406	Dummy-stool for act 1 cabin				/Data/Global/Objects	s9	NU	HTH		LIT																	
1	2	407	Dummy-wood for act 1 cabin				/Data/Global/Objects	wg	NU	HTH		LIT																	
1	2	408	Dummy-more wood for act 1 cabin				/Data/Global/Objects	wh	NU	HTH		LIT																	
1	2	409	Dummy-skeleton spawn for hell   facing nw				/Data/Global/Objects	QS	OP	HTH		LIT							LIT										
1	2	410	Shrine-holyshrine for monastery,catacombs,jail				/Data/Global/Objects	HL	OP	HTH		LIT							LIT										
1	2	411	a trap-spikes for tombs floortrap				/Data/Global/Objects	A7	OP	HTH		LIT																	
1	2	412	Shrine-act 1 cathedral				/Data/Global/Objects	s0	OP	HTH		LIT							LIT										
1	2	413	Shrine-act 1 jail				/Data/Global/Objects	jb	NU	HTH		LIT							LIT										
1	2	414	Shrine-act 1 jail				/Data/Global/Objects	jd	OP	HTH		LIT							LIT										
1	2	415	Shrine-act 1 jail				/Data/Global/Objects	jf	OP	HTH		LIT							LIT										
1	2	416	goo pile-goo pile for sand maggot lair				/Data/Global/Objects	GP	OP	HTH		LIT																	
1	2	417	bank-bank				/Data/Global/Objects	b6	NU	HTH		LIT																	
1	2	418	wirt's body-wirt's body				/Data/Global/Objects	BP	NU	HTH		LIT																	
1	2	419	dummy-gold placeholder				/Data/Global/Objects	1g																					
1	2	420	corpse-guard corpse 2				/Data/Global/Objects	GF	OP	HTH		LIT																	
1	2	421	corpse-dead villager 1				/Data/Global/Objects	dg	OP	HTH		LIT																	
1	2	422	corpse-dead villager 2				/Data/Global/Objects	df	OP	HTH		LIT																	
1	2	423	Dummy-yet another flame, no damage				/Data/Global/Objects	f8	NU	HTH		LIT																	
1	2	424	hidden stash-tiny pixel shaped thingie				/Data/Global/Objects	f9																					
1	2	425	Shrine-health shrine for caves				/Data/Global/Objects	ce	OP	HTH		LIT																	
1	2	426	Shrine-mana shrine for caves				/Data/Global/Objects	cg	OP	HTH		LIT																	
1	2	427	Shrine-cave magic shrine				/Data/Global/Objects	cg	OP	HTH		LIT																	
1	2	428	Shrine-manashrine, act 3, dungeun				/Data/Global/Objects	de	OP	HTH		LIT																	
1	2	429	Shrine-magic shrine, act 3 sewers.				/Data/Global/Objects	wj	NU	HTH		LIT							LIT	LIT									
1	2	430	Shrine-healthwell, act 3, sewers				/Data/Global/Objects	wk	OP	HTH		LIT																	
1	2	431	Shrine-manawell, act 3, sewers				/Data/Global/Objects	wl	OP	HTH		LIT																	
1	2	432	Shrine-magic shrine, act 3 sewers, dungeon.				/Data/Global/Objects	ws	NU	HTH		LIT							LIT	LIT									
1	2	433	dummy-brazier_celler, act 2				/Data/Global/Objects	bi	NU	HTH		LIT							LIT										
1	2	434	sarcophagus-anubis coffin, act2, tomb				/Data/Global/Objects	qc	OP	HTH		LIT																	
1	2	435	dummy-brazier_general, act 2, sewers, tomb, desert				/Data/Global/Objects	bm	NU	HTH		LIT							LIT										
1	2	436	Dummy-brazier_tall, act 2, desert, town, tombs				/Data/Global/Objects	bo	NU	HTH		LIT							LIT										
1	2	437	Dummy-brazier_small, act 2, desert, town, tombs				/Data/Global/Objects	bq	NU	HTH		LIT							LIT										
1	2	438	Waypoint-waypoint, celler				/Data/Global/Objects	w7	ON	HTH		LIT							LIT										
1	2	439	bed-bed for harum				/Data/Global/Objects	ub	OP	HTH		LIT																	
1	2	440	door-iron grate door left				/Data/Global/Objects	dv	NU	HTH		LIT																	
1	2	441	door-iron grate door right				/Data/Global/Objects	dn	NU	HTH		LIT																	
1	2	442	door-wooden grate door left				/Data/Global/Objects	dp	NU	HTH		LIT																	
1	2	443	door-wooden grate door right				/Data/Global/Objects	dt	NU	HTH		LIT																	
1	2	444	door-wooden door left				/Data/Global/Objects	dk	NU	HTH		LIT																	
1	2	445	door-wooden door right				/Data/Global/Objects	dl	NU	HTH		LIT																	
1	2	446	Dummy-wall torch left for tombs				/Data/Global/Objects	qd	NU	HTH		LIT							LIT										
1	2	447	Dummy-wall torch right for tombs				/Data/Global/Objects	qe	NU	HTH		LIT							LIT										
1	2	448	portal-arcane sanctuary portal				/Data/Global/Objects	ay	ON	HTH		LIT							LIT	LIT									
1	2	449	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hb	OP	HTH		LIT							LIT										
1	2	450	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hc	OP	HTH		LIT							LIT										
1	2	451	Dummy-maggot well health				/Data/Global/Objects	qf	OP	HTH		LIT																	
1	2	452	manashrine-maggot well mana				/Data/Global/Objects	qg	OP	HTH		LIT																	
1	2	453	magic shrine-magic shrine, act 3 arcane sanctuary.				/Data/Global/Objects	hd	OP	HTH		LIT							LIT										
1	2	454	teleportation pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
1	2	455	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
1	2	456	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
1	2	457	Dummy-arcane thing				/Data/Global/Objects	7a	NU	HTH		LIT																	
1	2	458	Dummy-arcane thing				/Data/Global/Objects	7b	NU	HTH		LIT																	
1	2	459	Dummy-arcane thing				/Data/Global/Objects	7c	NU	HTH		LIT																	
1	2	460	Dummy-arcane thing				/Data/Global/Objects	7d	NU	HTH		LIT																	
1	2	461	Dummy-arcane thing				/Data/Global/Objects	7e	NU	HTH		LIT																	
1	2	462	Dummy-arcane thing				/Data/Global/Objects	7f	NU	HTH		LIT																	
1	2	463	Dummy-arcane thing				/Data/Global/Objects	7g	NU	HTH		LIT																	
1	2	464	dead guard-harem guard 1				/Data/Global/Objects	qh	NU	HTH		LIT																	
1	2	465	dead guard-harem guard 2				/Data/Global/Objects	qi	NU	HTH		LIT																	
1	2	466	dead guard-harem guard 3				/Data/Global/Objects	qj	NU	HTH		LIT																	
1	2	467	dead guard-harem guard 4				/Data/Global/Objects	qk	NU	HTH		LIT																	
1	2	468	eunuch-harem blocker				/Data/Global/Objects	ss																					
1	2	469	Dummy-healthwell, act 2, arcane				/Data/Global/Objects	ax	OP	HTH		LIT																	
1	2	470	manashrine-healthwell, act 2, arcane				/Data/Global/Objects	au	OP	HTH		LIT																	
1	2	471	Dummy-test data				/Data/Global/Objects	pp	S1	HTH	LIT	LIT																	
1	2	472	Well-tombwell act 2 well, tomb				/Data/Global/Objects	hu	NU	HTH		LIT																	
1	2	473	Waypoint-waypoint act2 sewer				/Data/Global/Objects	qm	ON	HTH		LIT							LIT										
1	2	474	Waypoint-waypoint act3 travincal				/Data/Global/Objects	ql	ON	HTH		LIT							LIT										
1	2	475	magic shrine-magic shrine, act 3, sewer				/Data/Global/Objects	qn	NU	HTH		LIT							LIT										
1	2	476	dead body-act3, sewer				/Data/Global/Objects	qo	OP	HTH		LIT																	
1	2	477	dummy-torch (act 3 sewer) stra				/Data/Global/Objects	V1	NU	HTH		LIT							LIT										
1	2	478	dummy-torch (act 3 kurast) strb				/Data/Global/Objects	V2	NU	HTH		LIT							LIT										
1	2	479	chest-mafistochestlargeLeft				/Data/Global/Objects	xb	OP	HTH		LIT																	
1	2	480	chest-mafistochestlargeright				/Data/Global/Objects	xc	OP	HTH		LIT																	
1	2	481	chest-mafistochestmedleft				/Data/Global/Objects	xd	OP	HTH		LIT																	
1	2	482	chest-mafistochestmedright				/Data/Global/Objects	xe	OP	HTH		LIT																	
1	2	483	chest-spiderlairchestlargeLeft				/Data/Global/Objects	xf	OP	HTH		LIT																	
1	2	484	chest-spiderlairchesttallLeft				/Data/Global/Objects	xg	OP	HTH		LIT																	
1	2	485	chest-spiderlairchestmedright				/Data/Global/Objects	xh	OP	HTH		LIT																	
1	2	486	chest-spiderlairchesttallright				/Data/Global/Objects	xi	OP	HTH		LIT																	
1	2	487	Steeg Stone-steeg stone				/Data/Global/Objects	y6	NU	HTH		LIT							LIT										
1	2	488	Guild Vault-guild vault				/Data/Global/Objects	y4	NU	HTH		LIT																	
1	2	489	Trophy Case-trophy case				/Data/Global/Objects	y2	NU	HTH		LIT																	
1	2	490	Message Board-message board				/Data/Global/Objects	y3	NU	HTH		LIT																	
1	2	491	Dummy-mephisto bridge				/Data/Global/Objects	xj	OP	HTH		LIT																	
1	2	492	portal-hellgate				/Data/Global/Objects	1y	ON	HTH		LIT								LIT	LIT								
1	2	493	Shrine-manawell, act 3, kurast				/Data/Global/Objects	xl	OP	HTH		LIT																	
1	2	494	Shrine-healthwell, act 3, kurast				/Data/Global/Objects	xm	OP	HTH		LIT																	
1	2	495	Dummy-hellfire1				/Data/Global/Objects	e3	NU	HTH		LIT																	
1	2	496	Dummy-hellfire2				/Data/Global/Objects	e4	NU	HTH		LIT																	
1	2	497	Dummy-hellfire3				/Data/Global/Objects	e5	NU	HTH		LIT																	
1	2	498	Dummy-helllava1				/Data/Global/Objects	e6	NU	HTH		LIT																	
1	2	499	Dummy-helllava2				/Data/Global/Objects	e7	NU	HTH		LIT																	
1	2	500	Dummy-helllava3				/Data/Global/Objects	e8	NU	HTH		LIT																	
1	2	501	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
1	2	502	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
1	2	503	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
1	2	504	chest-horadric cube chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
1	2	505	chest-horadric scroll chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
1	2	506	chest-staff of kings chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
1	2	507	Tome-yet another tome				/Data/Global/Objects	TT	NU	HTH		LIT																	
1	2	508	fire-hell brazier				/Data/Global/Objects	E1	NU	HTH	LIT	LIT																	
1	2	509	fire-hell brazier				/Data/Global/Objects	E2	NU	HTH	LIT	LIT																	
1	2	510	RockPIle-dungeon				/Data/Global/Objects	xn	OP	HTH		LIT																	
1	2	511	magic shrine-magic shrine, act 3,dundeon				/Data/Global/Objects	qo	OP	HTH		LIT																	
1	2	512	basket-dungeon				/Data/Global/Objects	xp	OP	HTH		LIT																	
1	2	513	HungSkeleton-outerhell skeleton				/Data/Global/Objects	jw	OP	HTH		LIT																	
1	2	514	Dummy-guy for dungeon				/Data/Global/Objects	ea	OP	HTH		LIT																	
1	2	515	casket-casket for Act 3 dungeon				/Data/Global/Objects	vb	OP	HTH		LIT																	
1	2	516	sewer stairs-stairs for act 3 sewer quest				/Data/Global/Objects	ve	OP	HTH		LIT																	
1	2	517	sewer lever-lever for act 3 sewer quest				/Data/Global/Objects	vf	OP	HTH		LIT																	
1	2	518	darkwanderer-start position				/Data/Global/Objects	ss																					
1	2	519	dummy-trapped soul placeholder				/Data/Global/Objects	ss																					
1	2	520	Dummy-torch for act3 town				/Data/Global/Objects	VG	NU	HTH		LIT							LIT										
1	2	521	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
1	2	522	BoneChest-innerhellbonepile				/Data/Global/Objects	y1	OP	HTH		LIT																	
1	2	523	Dummy-skeleton spawn for hell facing ne				/Data/Global/Objects	Qt	OP	HTH		LIT							LIT										
1	2	524	Dummy-fog act 3 water rfga				/Data/Global/Objects	ud	NU	HTH		LIT																	
1	2	525	Dummy-Not used				/Data/Global/Objects	xx																					
1	2	526	Hellforge-Forge  hell				/Data/Global/Objects	ux	ON	HTH		LIT							LIT	LIT	LIT								
1	2	527	Guild Portal-Portal to next guild level				/Data/Global/Objects	PP	NU	HTH	LIT	LIT																	
1	2	528	Dummy-hratli start				/Data/Global/Objects	ss																					
1	2	529	Dummy-hratli end				/Data/Global/Objects	ss																					
1	2	530	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	uy	OP	HTH		LIT							LIT										
1	2	531	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	15	OP	HTH		LIT							LIT										
1	2	532	Dummy-natalya start				/Data/Global/Objects	ss																					
1	2	533	TrappedSoul-guy stuck in hell				/Data/Global/Objects	18	OP	HTH		LIT																	
1	2	534	TrappedSoul-guy stuck in hell				/Data/Global/Objects	19	OP	HTH		LIT																	
1	2	535	Dummy-cain start position				/Data/Global/Objects	ss																					
1	2	536	Dummy-stairsr				/Data/Global/Objects	sv	OP	HTH		LIT																	
1	2	537	chest-arcanesanctuarybigchestLeft				/Data/Global/Objects	y7	OP	HTH		LIT																	
1	2	538	casket-arcanesanctuarycasket				/Data/Global/Objects	y8	OP	HTH		LIT																	
1	2	539	chest-arcanesanctuarybigchestRight				/Data/Global/Objects	y9	OP	HTH		LIT																	
1	2	540	chest-arcanesanctuarychestsmallLeft				/Data/Global/Objects	ya	OP	HTH		LIT																	
1	2	541	chest-arcanesanctuarychestsmallRight				/Data/Global/Objects	yc	OP	HTH		LIT																	
1	2	542	Seal-Diablo seal				/Data/Global/Objects	30	ON	HTH		LIT							LIT										
1	2	543	Seal-Diablo seal				/Data/Global/Objects	31	ON	HTH		LIT							LIT										
1	2	544	Seal-Diablo seal				/Data/Global/Objects	32	ON	HTH		LIT							LIT										
1	2	545	Seal-Diablo seal				/Data/Global/Objects	33	ON	HTH		LIT							LIT										
1	2	546	Seal-Diablo seal				/Data/Global/Objects	34	ON	HTH		LIT							LIT										
1	2	547	chest-sparklychest				/Data/Global/Objects	yf	OP	HTH		LIT																	
1	2	548	Waypoint-waypoint pandamonia fortress				/Data/Global/Objects	yg	ON	HTH		LIT							LIT										
1	2	549	fissure-fissure for act 4 inner hell				/Data/Global/Objects	fh	OP	HTH		LIT							LIT										
1	2	550	Dummy-brazier for act 4, hell mesa				/Data/Global/Objects	he	NU	HTH		LIT							LIT										
1	2	551	Dummy-smoke				/Data/Global/Objects	35	NU	HTH		LIT																	
1	2	552	Waypoint-waypoint valleywaypoint				/Data/Global/Objects	yi	ON	HTH		LIT							LIT										
1	2	553	fire-hell brazier				/Data/Global/Objects	9f	NU	HTH		LIT							LIT										
1	2	554	compellingorb-compelling orb				/Data/Global/Objects	55	NU	HTH		LIT							LIT	LIT									
1	2	555	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
1	2	556	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
1	2	557	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
1	2	558	Dummy-fortress brazier #1				/Data/Global/Objects	98	NU	HTH		LIT							LIT										
1	2	559	Dummy-fortress brazier #2				/Data/Global/Objects	99	NU	HTH		LIT							LIT										
1	2	560	Siege Control-To control siege machines				/Data/Global/Objects	zq	OP	HTH		LIT																	
1	2	561	ptox-Pot O Torch (level 1)				/Data/Global/Objects	px	NU	HTH		LIT							LIT	LIT									
1	2	562	pyox-fire pit  (level 1)				/Data/Global/Objects	py	NU	HTH		LIT							LIT										
1	2	563	chestR-expansion no snow				/Data/Global/Objects	6q	OP	HTH		LIT																	
1	2	564	Shrine3wilderness-expansion no snow				/Data/Global/Objects	6r	OP	HTH		LIT							LIT										
1	2	565	Shrine2wilderness-expansion no snow				/Data/Global/Objects	6s	NU	HTH		LIT							LIT										
1	2	566	hiddenstash-expansion no snow				/Data/Global/Objects	3w	OP	HTH		LIT																	
1	2	567	flag wilderness-expansion no snow				/Data/Global/Objects	ym	NU	HTH		LIT																	
1	2	568	barrel wilderness-expansion no snow				/Data/Global/Objects	yn	OP	HTH		LIT																	
1	2	569	barrel wilderness-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
1	2	570	woodchestL-expansion no snow				/Data/Global/Objects	yp	OP	HTH		LIT																	
1	2	571	Shrine3wilderness-expansion no snow				/Data/Global/Objects	yq	NU	HTH		LIT							LIT										
1	2	572	manashrine-expansion no snow				/Data/Global/Objects	yr	OP	HTH		LIT							LIT										
1	2	573	healthshrine-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
1	2	574	burialchestL-expansion no snow				/Data/Global/Objects	yt	OP	HTH		LIT																	
1	2	575	burialchestR-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
1	2	576	well-expansion no snow				/Data/Global/Objects	yv	NU	HTH		LIT																	
1	2	577	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yw	OP	HTH		LIT							LIT	LIT									
1	2	578	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yx	OP	HTH		LIT							LIT										
1	2	579	Waypoint-expansion no snow				/Data/Global/Objects	yy	ON	HTH		LIT							LIT										
1	2	580	ChestL-expansion no snow				/Data/Global/Objects	yz	OP	HTH		LIT																	
1	2	581	woodchestR-expansion no snow				/Data/Global/Objects	6a	OP	HTH		LIT																	
1	2	582	ChestSL-expansion no snow				/Data/Global/Objects	6b	OP	HTH		LIT																	
1	2	583	ChestSR-expansion no snow				/Data/Global/Objects	6c	OP	HTH		LIT																	
1	2	584	etorch1-expansion no snow				/Data/Global/Objects	6d	NU	HTH		LIT							LIT										
1	2	585	ecfra-camp fire				/Data/Global/Objects	2w	NU	HTH		LIT							LIT	LIT									
1	2	586	ettr-town torch				/Data/Global/Objects	2x	NU	HTH		LIT							LIT	LIT									
1	2	587	etorch2-expansion no snow				/Data/Global/Objects	6e	NU	HTH		LIT							LIT										
1	2	588	burningbodies-wilderness/siege				/Data/Global/Objects	6f	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
1	2	589	burningpit-wilderness/siege				/Data/Global/Objects	6g	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
1	2	590	tribal flag-wilderness/siege				/Data/Global/Objects	6h	NU	HTH		LIT																	
1	2	591	eflg-town flag				/Data/Global/Objects	2y	NU	HTH		LIT																	
1	2	592	chan-chandeleir				/Data/Global/Objects	2z	NU	HTH		LIT							LIT										
1	2	593	jar1-wilderness/siege				/Data/Global/Objects	6i	OP	HTH		LIT																	
1	2	594	jar2-wilderness/siege				/Data/Global/Objects	6j	OP	HTH		LIT																	
1	2	595	jar3-wilderness/siege				/Data/Global/Objects	6k	OP	HTH		LIT																	
1	2	596	swingingheads-wilderness				/Data/Global/Objects	6L	NU	HTH		LIT																	
1	2	597	pole-wilderness				/Data/Global/Objects	6m	NU	HTH		LIT																	
1	2	598	animated skulland rockpile-expansion no snow				/Data/Global/Objects	6n	OP	HTH		LIT																	
1	2	599	gate-town main gate				/Data/Global/Objects	2v	OP	HTH		LIT																	
1	2	600	pileofskullsandrocks-seige				/Data/Global/Objects	6o	NU	HTH		LIT																	
1	2	601	hellgate-seige				/Data/Global/Objects	6p	NU	HTH		LIT							LIT	LIT									
1	2	602	banner 1-preset in enemy camp				/Data/Global/Objects	ao	NU	HTH		LIT																	
1	2	603	banner 2-preset in enemy camp				/Data/Global/Objects	ap	NU	HTH		LIT																	
1	2	604	explodingchest-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
1	2	605	chest-specialchest				/Data/Global/Objects	6u	OP	HTH		LIT																	
1	2	606	deathpole-wilderness				/Data/Global/Objects	6v	NU	HTH		LIT																	
1	2	607	Ldeathpole-wilderness				/Data/Global/Objects	6w	NU	HTH		LIT																	
1	2	608	Altar-inside of temple				/Data/Global/Objects	6x	NU	HTH		LIT							LIT										
1	2	609	dummy-Drehya Start In Town				/Data/Global/Objects	ss																					
1	2	610	dummy-Drehya Start Outside Town				/Data/Global/Objects	ss																					
1	2	611	dummy-Nihlathak Start In Town				/Data/Global/Objects	ss																					
1	2	612	dummy-Nihlathak Start Outside Town				/Data/Global/Objects	ss																					
1	2	613	hidden stash-icecave_				/Data/Global/Objects	6y	OP	HTH		LIT																	
1	2	614	healthshrine-icecave_				/Data/Global/Objects	8a	OP	HTH		LIT																	
1	2	615	manashrine-icecave_				/Data/Global/Objects	8b	OP	HTH		LIT																	
1	2	616	evilurn-icecave_				/Data/Global/Objects	8c	OP	HTH		LIT																	
1	2	617	icecavejar1-icecave_				/Data/Global/Objects	8d	OP	HTH		LIT																	
1	2	618	icecavejar2-icecave_				/Data/Global/Objects	8e	OP	HTH		LIT																	
1	2	619	icecavejar3-icecave_				/Data/Global/Objects	8f	OP	HTH		LIT																	
1	2	620	icecavejar4-icecave_				/Data/Global/Objects	8g	OP	HTH		LIT																	
1	2	621	icecavejar4-icecave_				/Data/Global/Objects	8h	OP	HTH		LIT																	
1	2	622	icecaveshrine2-icecave_				/Data/Global/Objects	8i	NU	HTH		LIT							LIT										
1	2	623	cagedwussie1-caged fellow(A5-Prisonner)				/Data/Global/Objects	60	NU	HTH		LIT																	
1	2	624	Ancient Statue 3-statue				/Data/Global/Objects	60	NU	HTH		LIT																	
1	2	625	Ancient Statue 1-statue				/Data/Global/Objects	61	NU	HTH		LIT																	
1	2	626	Ancient Statue 2-statue				/Data/Global/Objects	62	NU	HTH		LIT																	
1	2	627	deadbarbarian-seige/wilderness				/Data/Global/Objects	8j	OP	HTH		LIT																	
1	2	628	clientsmoke-client smoke				/Data/Global/Objects	oz	NU	HTH		LIT																	
1	2	629	icecaveshrine2-icecave_				/Data/Global/Objects	8k	NU	HTH		LIT							LIT										
1	2	630	icecave_torch1-icecave_				/Data/Global/Objects	8L	NU	HTH		LIT							LIT										
1	2	631	icecave_torch2-icecave_				/Data/Global/Objects	8m	NU	HTH		LIT							LIT										
1	2	632	ttor-expansion tiki torch				/Data/Global/Objects	2p	NU	HTH		LIT							LIT										
1	2	633	manashrine-baals				/Data/Global/Objects	8n	OP	HTH		LIT																	
1	2	634	healthshrine-baals				/Data/Global/Objects	8o	OP	HTH		LIT																	
1	2	635	tomb1-baal's lair				/Data/Global/Objects	8p	OP	HTH		LIT																	
1	2	636	tomb2-baal's lair				/Data/Global/Objects	8q	OP	HTH		LIT																	
1	2	637	tomb3-baal's lair				/Data/Global/Objects	8r	OP	HTH		LIT																	
1	2	638	magic shrine-baal's lair				/Data/Global/Objects	8s	NU	HTH		LIT							LIT										
1	2	639	torch1-baal's lair				/Data/Global/Objects	8t	NU	HTH		LIT							LIT										
1	2	640	torch2-baal's lair				/Data/Global/Objects	8u	NU	HTH		LIT							LIT										
1	2	641	manashrine-snowy				/Data/Global/Objects	8v	OP	HTH		LIT							LIT										
1	2	642	healthshrine-snowy				/Data/Global/Objects	8w	OP	HTH		LIT							LIT										
1	2	643	well-snowy				/Data/Global/Objects	8x	NU	HTH		LIT																	
1	2	644	Waypoint-baals_waypoint				/Data/Global/Objects	8y	ON	HTH		LIT							LIT										
1	2	645	magic shrine-snowy_shrine3				/Data/Global/Objects	8z	NU	HTH		LIT							LIT										
1	2	646	Waypoint-wilderness_waypoint				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
1	2	647	magic shrine-snowy_shrine3				/Data/Global/Objects	5b	OP	HTH		LIT							LIT	LIT									
1	2	648	well-baalslair				/Data/Global/Objects	5c	NU	HTH		LIT																	
1	2	649	magic shrine2-baal's lair				/Data/Global/Objects	5d	NU	HTH		LIT							LIT										
1	2	650	object1-snowy				/Data/Global/Objects	5e	OP	HTH		LIT																	
1	2	651	woodchestL-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
1	2	652	woodchestR-snowy				/Data/Global/Objects	5g	OP	HTH		LIT																	
1	2	653	magic shrine-baals_shrine3				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
1	2	654	woodchest2L-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
1	2	655	woodchest2R-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
1	2	656	swingingheads-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
1	2	657	debris-snowy				/Data/Global/Objects	5l	NU	HTH		LIT																	
1	2	658	pene-Pen breakable door				/Data/Global/Objects	2q	NU	HTH		LIT																	
1	2	659	magic shrine-temple				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
1	2	660	mrpole-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
1	2	661	Waypoint-icecave 				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
1	2	662	magic shrine-temple				/Data/Global/Objects	5t	NU	HTH		LIT							LIT										
1	2	663	well-temple				/Data/Global/Objects	5q	NU	HTH		LIT																	
1	2	664	torch1-temple				/Data/Global/Objects	5r	NU	HTH		LIT							LIT										
1	2	665	torch1-temple				/Data/Global/Objects	5s	NU	HTH		LIT							LIT										
1	2	666	object1-temple				/Data/Global/Objects	5u	OP	HTH		LIT																	
1	2	667	object2-temple				/Data/Global/Objects	5v	OP	HTH		LIT																	
1	2	668	mrbox-baals				/Data/Global/Objects	5w	OP	HTH		LIT																	
1	2	669	well-icecave				/Data/Global/Objects	5x	NU	HTH		LIT																	
1	2	670	magic shrine-temple				/Data/Global/Objects	5y	NU	HTH		LIT							LIT										
1	2	671	healthshrine-temple				/Data/Global/Objects	5z	OP	HTH		LIT																	
1	2	672	manashrine-temple				/Data/Global/Objects	3a	OP	HTH		LIT																	
1	2	673	red light- (touch me)  for blacksmith				/Data/Global/Objects	ss																					
1	2	674	tomb1L-baal's lair				/Data/Global/Objects	3b	OP	HTH		LIT																	
1	2	675	tomb2L-baal's lair				/Data/Global/Objects	3c	OP	HTH		LIT																	
1	2	676	tomb3L-baal's lair				/Data/Global/Objects	3d	OP	HTH		LIT																	
1	2	677	ubub-Ice cave bubbles 01				/Data/Global/Objects	2u	NU	HTH		LIT																	
1	2	678	sbub-Ice cave bubbles 01				/Data/Global/Objects	2s	NU	HTH		LIT																	
1	2	679	tomb1-redbaal's lair				/Data/Global/Objects	3f	OP	HTH		LIT																	
1	2	680	tomb1L-redbaal's lair				/Data/Global/Objects	3g	OP	HTH		LIT																	
1	2	681	tomb2-redbaal's lair				/Data/Global/Objects	3h	OP	HTH		LIT																	
1	2	682	tomb2L-redbaal's lair				/Data/Global/Objects	3i	OP	HTH		LIT																	
1	2	683	tomb3-redbaal's lair				/Data/Global/Objects	3j	OP	HTH		LIT																	
1	2	684	tomb3L-redbaal's lair				/Data/Global/Objects	3k	OP	HTH		LIT																	
1	2	685	mrbox-redbaals				/Data/Global/Objects	3L	OP	HTH		LIT																	
1	2	686	torch1-redbaal's lair				/Data/Global/Objects	3m	NU	HTH		LIT							LIT										
1	2	687	torch2-redbaal's lair				/Data/Global/Objects	3n	NU	HTH		LIT							LIT										
1	2	688	candles-temple				/Data/Global/Objects	3o	NU	HTH		LIT							LIT										
1	2	689	Waypoint-temple				/Data/Global/Objects	3p	ON	HTH		LIT							LIT										
1	2	690	deadperson-everywhere				/Data/Global/Objects	3q	NU	HTH		LIT																	
1	2	691	groundtomb-temple				/Data/Global/Objects	3s	OP	HTH		LIT																	
1	2	692	Dummy-Larzuk Greeting				/Data/Global/Objects	ss																					
1	2	693	Dummy-Larzuk Standard				/Data/Global/Objects	ss																					
1	2	694	groundtombL-temple				/Data/Global/Objects	3t	OP	HTH		LIT																	
1	2	695	deadperson2-everywhere				/Data/Global/Objects	3u	OP	HTH		LIT																	
1	2	696	ancientsaltar-ancientsaltar				/Data/Global/Objects	4a	OP	HTH		LIT							LIT										
1	2	697	To The Worldstone Keep Level 1-ancientsdoor				/Data/Global/Objects	4b	OP	HTH		LIT																	
1	2	698	eweaponrackR-everywhere				/Data/Global/Objects	3x	NU	HTH		LIT																	
1	2	699	eweaponrackL-everywhere				/Data/Global/Objects	3y	NU	HTH		LIT																	
1	2	700	earmorstandR-everywhere				/Data/Global/Objects	3z	NU	HTH		LIT																	
1	2	701	earmorstandL-everywhere				/Data/Global/Objects	4c	NU	HTH		LIT																	
1	2	702	torch2-summit				/Data/Global/Objects	9g	NU	HTH		LIT							LIT										
1	2	703	funeralpire-outside				/Data/Global/Objects	9h	NU	HTH		LIT							LIT										
1	2	704	burninglogs-outside				/Data/Global/Objects	9i	NU	HTH		LIT							LIT										
1	2	705	stma-Ice cave steam				/Data/Global/Objects	2o	NU	HTH		LIT																	
1	2	706	deadperson2-everywhere				/Data/Global/Objects	3v	OP	HTH		LIT																	
1	2	707	Dummy-Baal's lair				/Data/Global/Objects	ss																					
1	2	708	fana-frozen anya				/Data/Global/Objects	2n	NU	HTH		LIT																	
1	2	709	BBQB-BBQ Bunny				/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									
1	2	710	btor-Baal Torch Big				/Data/Global/Objects	25	NU	HTH		LIT							LIT										
1	2	711	Dummy-invisible ancient				/Data/Global/Objects	ss																					
1	2	712	Dummy-invisible base				/Data/Global/Objects	ss																					
1	2	713	The Worldstone Chamber-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
1	2	714	Glacial Caves Level 1-summit door				/Data/Global/Objects	4u	OP	HTH		LIT																	
1	2	715	strlastcinematic-last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
1	2	716	Harrogath-last last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
1	2	717	Zoo-test data				/Data/Global/Objects	ss																					
1	2	718	Keeper-test data				/Data/Global/Objects	7z	NU	HTH		LIT																	
1	2	719	Throne of Destruction-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
1	2	720	Dummy-fire place guy				/Data/Global/Objects	7y	NU	HTH		LIT																	
1	2	721	Dummy-door blocker				/Data/Global/Objects	ss																					
1	2	722	Dummy-door blocker				/Data/Global/Objects	ss																					
2	1	0	warriv2-ACT 2 TABLE				/Data/Global/Monsters	ss	NU	HTH		LIT																	0
2	1	1	atma-ACT 2 TABLE				/Data/Global/Monsters	ss	NU	HTH		LIT																	0
2	1	2	drognan-ACT 2 TABLE				/Data/Global/Monsters	zv	NU	HTH		LIT																	0
2	1	3	fara-ACT 2 TABLE				/Data/Global/Monsters	OF	NU	HTH		LIT																	0
2	1	4	place_nothing-ACT 2 TABLE																										0
2	1	5	place_nothing-ACT 2 TABLE																										0
2	1	6	place_nothing-ACT 2 TABLE																										0
2	1	7	place_nothing-ACT 2 TABLE																										0
2	1	8	greiz-ACT 2 TABLE				/Data/Global/Monsters	GR	NU	HTH		LIT																	0
2	1	9	elzix-ACT 2 TABLE				/Data/Global/Monsters	EL	NU	HTH		LIT																	0
2	1	10	lysander-ACT 2 TABLE				/Data/Global/Monsters	LY	NU	HTH		LIT																	0
2	1	11	meshif1-ACT 2 TABLE				/Data/Global/Monsters	MS	NU	HTH		LIT																	0
2	1	12	geglash-ACT 2 TABLE				/Data/Global/Monsters	GE	NU	HTH		LIT																	0
2	1	13	jerhyn-ACT 2 TABLE				/Data/Global/Monsters	JE	NU	HTH		LIT																	0
2	1	14	place_unique_pack-ACT 2 TABLE																										0
2	1	15	place_npc_pack-ACT 2 TABLE																										0
2	1	16	place_nothing-ACT 2 TABLE																										0
2	1	17	summoner-ACT 2 TABLE				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
2	1	18	Radament-ACT 2 TABLE				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
2	1	19	duriel-ACT 2 TABLE			6	/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
2	1	20	cain2-ACT 2 TABLE				/Data/Global/Monsters	2D	NU	HTH		LIT																	0
2	1	21	place_champion-ACT 2 TABLE																										0
2	1	22	act2male-ACT 2 TABLE				/Data/Global/Monsters	2M	NU	HTH	YNG	MED	MED						FEZ										0
2	1	23	act2female-ACT 2 TABLE				/Data/Global/Monsters	2F	NU	HTH	MED	MED	HVY																0
2	1	24	act2guard1-ACT 2 TABLE			7	/Data/Global/Monsters	GU	NU	HTH	MED	MED	MED	MED	MED	GLV			LIT	LIT	LIT								0
2	1	25	act2vendor1-ACT 2 TABLE				/Data/Global/Monsters	M1	NU	HTH		LIT																	0
2	1	26	act2vendor2-ACT 2 TABLE				/Data/Global/Monsters	M2	NU	HTH		LIT																	0
2	1	27	place_tightspotboss-ACT 2 TABLE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	28	fish-ACT 2 TABLE																										0
2	1	29	place_talkingguard-ACT 2 TABLE																										0
2	1	30	place_dumbguard-ACT 2 TABLE																										0
2	1	31	place_maggot-ACT 2 TABLE																										0
2	1	32	place_maggotegg-ACT 2 TABLE																										0
2	1	33	place_nothing-ACT 2 TABLE																										0
2	1	34	gargoyletrap-ACT 2 TABLE				/Data/Global/Monsters	GT	A1	HTH		LIT																	0
2	1	35	trap-horzmissile-ACT 2 TABLE																										0
2	1	36	trap-vertmissile-ACT 2 TABLE																										0
2	1	37	place_group25-ACT 2 TABLE																										0
2	1	38	place_group50-ACT 2 TABLE																										0
2	1	39	place_group75-ACT 2 TABLE																										0
2	1	40	place_group100-ACT 2 TABLE																										0
2	1	41	lightningspire-ACT 2 TABLE				/Data/Global/Monsters	AE	A1	HTH		LIT							LIT										0
2	1	42	firetower-ACT 2 TABLE			13	/Data/Global/Monsters	PB	A1	HTH		LIT																	0
2	1	43	place_nothing-ACT 2 TABLE																										0
2	1	44	place_nothing-ACT 2 TABLE																										0
2	1	45	place_nothing-ACT 2 TABLE																										0
2	1	46	Bloodwitch the Wild-ACT 2 TABLE			7	/Data/Global/Monsters	PW	NU	HTH	LIT	MED		MED	MED		WHP	BUC	LIT	LIT	LIT	LIT					/Data/Global/Monsters/PW/COF/Palshift.dat	3	0
2	1	47	Fangskin-ACT 2 TABLE			7	/Data/Global/Monsters	SD	NU	HTH		LIT															/Data/Global/Monsters/SD/COF/Palshift.dat	6	0
2	1	48	Beetleburst-ACT 2 TABLE			5	/Data/Global/Monsters	SC	NU	HTH	MED	LIT		HVY													/Data/Global/Monsters/SC/COF/Palshift.dat	6	0
2	1	49	Leatherarm-ACT 2 TABLE			3	/Data/Global/Monsters	MM	NU	HTH		LIT															/Data/Global/Monsters/MM/COF/Palshift.dat	7	0
2	1	50	Coldworm the Burrower-ACT 2 TABLE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	51	Fire Eye-ACT 2 TABLE				/Data/Global/Monsters	SR	NU	HTH		LIT															/Data/Global/Monsters/SR/COF/Palshift.dat	5	0
2	1	52	Dark Elder-ACT 2 TABLE				/Data/Global/Monsters	ZM	NU	HTH	MED	MED	MED	MED	MED				MED	MED	BLD						/Data/Global/Monsters/ZM/COF/Palshift.dat	6	0
2	1	53	Ancient Kaa the Soulless-ACT 2 TABLE			7	/Data/Global/Monsters	GY	NU	HTH		LIT															/Data/Global/Monsters/GY/COF/Palshift.dat	4	0
2	1	54	act2guard4-ACT 2 TABLE			7	/Data/Global/Monsters	GU	NU	HTH	MED	MED	MED	MED	MED	GLV			LIT	LIT	LIT								0
2	1	55	act2guard5-ACT 2 TABLE			7	/Data/Global/Monsters	GU	NU	HTH	MED	MED	MED	MED	MED	SPR			LIT	LIT	LIT								0
2	1	56	sarcophagus-ACT 2 TABLE				/Data/Global/Monsters	MG	S1	HTH		LIT																	0
2	1	57	tyrael1-ACT 2 TABLE				/Data/Global/Monsters	TX	NU	HTH		LIT		LIT	LIT														0
2	1	58	skeleton5-ACT 2 TABLE			3	/Data/Global/Monsters	SK	NU	1HS	DES	HVY	DES	DES	DES	AXE		LRG	DES	DES	LIT						/Data/Global/Monsters/SK/COF/Palshift.dat	7	0
2	1	59	ACT 2 TABLE SKIP IT																										
2	1	60	ACT 2 TABLE SKIP IT																										
2	1	61	ACT 2 TABLE SKIP IT																										
2	1	62	ACT 2 TABLE SKIP IT																										
2	1	63	ACT 2 TABLE SKIP IT																										
2	1	64	ACT 2 TABLE SKIP IT																										
2	1	65	ACT 2 TABLE SKIP IT																										
2	1	66	ACT 2 TABLE SKIP IT																										
2	1	67	ACT 2 TABLE SKIP IT																										
2	1	68	ACT 2 TABLE SKIP IT																										
2	1	69	ACT 2 TABLE SKIP IT																										
2	1	70	ACT 2 TABLE SKIP IT																										
2	1	71	ACT 2 TABLE SKIP IT																										
2	1	72	ACT 2 TABLE SKIP IT																										
2	1	73	ACT 2 TABLE SKIP IT																										
2	1	74	ACT 2 TABLE SKIP IT																										
2	1	75	ACT 2 TABLE SKIP IT																										
2	1	76	ACT 2 TABLE SKIP IT																										
2	1	77	ACT 2 TABLE SKIP IT																										
2	1	78	ACT 2 TABLE SKIP IT																										
2	1	79	ACT 2 TABLE SKIP IT																										
2	1	80	ACT 2 TABLE SKIP IT																										
2	1	81	ACT 2 TABLE SKIP IT																										
2	1	82	ACT 2 TABLE SKIP IT																										
2	1	83	ACT 2 TABLE SKIP IT																										
2	1	84	ACT 2 TABLE SKIP IT																										
2	1	85	ACT 2 TABLE SKIP IT																										
2	1	86	ACT 2 TABLE SKIP IT																										
2	1	87	ACT 2 TABLE SKIP IT																										
2	1	88	ACT 2 TABLE SKIP IT																										
2	1	89	ACT 2 TABLE SKIP IT																										
2	1	90	ACT 2 TABLE SKIP IT																										
2	1	91	ACT 2 TABLE SKIP IT																										
2	1	92	ACT 2 TABLE SKIP IT																										
2	1	93	ACT 2 TABLE SKIP IT																										
2	1	94	ACT 2 TABLE SKIP IT																										
2	1	95	ACT 2 TABLE SKIP IT																										
2	1	96	ACT 2 TABLE SKIP IT																										
2	1	97	ACT 2 TABLE SKIP IT																										
2	1	98	ACT 2 TABLE SKIP IT																										
2	1	99	ACT 2 TABLE SKIP IT																										
2	1	100	ACT 2 TABLE SKIP IT																										
2	1	101	ACT 2 TABLE SKIP IT																										
2	1	102	ACT 2 TABLE SKIP IT																										
2	1	103	ACT 2 TABLE SKIP IT																										
2	1	104	ACT 2 TABLE SKIP IT																										
2	1	105	ACT 2 TABLE SKIP IT																										
2	1	106	skeleton1-Skeleton-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	107	skeleton2-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	108	skeleton3-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	109	skeleton4-BurningDead-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	110	skeleton5-Horror-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	111	zombie1-Zombie-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	112	zombie2-HungryDead-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	113	zombie3-Ghoul-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	114	zombie4-DrownedCarcass-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	115	zombie5-PlagueBearer-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	116	bighead1-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	117	bighead2-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	118	bighead3-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	119	bighead4-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	120	bighead5-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	121	foulcrow1-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	122	foulcrow2-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	123	foulcrow3-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	124	foulcrow4-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	125	fallen1-Fallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
2	1	126	fallen2-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
2	1	127	fallen3-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
2	1	128	fallen4-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
2	1	129	fallen5-WarpedFallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
2	1	130	brute2-Brute-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
2	1	131	brute3-Yeti-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
2	1	132	brute4-Crusher-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
2	1	133	brute5-WailingBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
2	1	134	brute1-GargantuanBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
2	1	135	sandraider1-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	136	sandraider2-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	137	sandraider3-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	138	sandraider4-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	139	sandraider5-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	140	gorgon1-unused-Idle				/Data/Global/Monsters	GO																					0
2	1	141	gorgon2-unused-Idle				/Data/Global/Monsters	GO																					0
2	1	142	gorgon3-unused-Idle				/Data/Global/Monsters	GO																					0
2	1	143	gorgon4-unused-Idle				/Data/Global/Monsters	GO																					0
2	1	144	wraith1-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	145	wraith2-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	146	wraith3-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	147	wraith4-Apparition-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	148	wraith5-DarkShape-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	149	corruptrogue1-DarkHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	150	corruptrogue2-VileHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	151	corruptrogue3-DarkStalker-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	152	corruptrogue4-BlackRogue-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	153	corruptrogue5-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	154	baboon1-DuneBeast-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	155	baboon2-RockDweller-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	156	baboon3-JungleHunter-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	157	baboon4-DoomApe-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	158	baboon5-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	159	goatman1-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	160	goatman2-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	161	goatman3-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	162	goatman4-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	163	goatman5-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	164	fallenshaman1-FallenShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	165	fallenshaman2-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	166	fallenshaman3-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	167	fallenshaman4-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	168	fallenshaman5-WarpedShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	169	quillrat1-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	170	quillrat2-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	171	quillrat3-ThornBeast-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	172	quillrat4-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	173	quillrat5-JungleUrchin-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	174	sandmaggot1-SandMaggot-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	175	sandmaggot2-RockWorm-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	176	sandmaggot3-Devourer-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	177	sandmaggot4-GiantLamprey-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	178	sandmaggot5-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	179	clawviper1-TombViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	180	clawviper2-ClawViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	181	clawviper3-Salamander-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	182	clawviper4-PitViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	183	clawviper5-SerpentMagus-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	184	sandleaper1-SandLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	185	sandleaper2-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	186	sandleaper3-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	187	sandleaper4-TreeLurker-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	188	sandleaper5-RazorPitDemon-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	189	pantherwoman1-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	190	pantherwoman2-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	191	pantherwoman3-NightTiger-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	192	pantherwoman4-HellCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	193	swarm1-Itchies-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
2	1	194	swarm2-BlackLocusts-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
2	1	195	swarm3-PlagueBugs-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
2	1	196	swarm4-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
2	1	197	scarab1-DungSoldier-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	198	scarab2-SandWarrior-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	199	scarab3-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	200	scarab4-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	201	scarab5-AlbinoRoach-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	202	mummy1-DriedCorpse-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	203	mummy2-Decayed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	204	mummy3-Embalmed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	205	mummy4-PreservedDead-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	206	mummy5-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	207	unraveler1-HollowOne-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	208	unraveler2-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	209	unraveler3-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	210	unraveler4-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	211	unraveler5-Baal Subject Mummy-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	212	chaoshorde1-unused-Idle				/Data/Global/Monsters	CH																					0
2	1	213	chaoshorde2-unused-Idle				/Data/Global/Monsters	CH																					0
2	1	214	chaoshorde3-unused-Idle				/Data/Global/Monsters	CH																					0
2	1	215	chaoshorde4-unused-Idle				/Data/Global/Monsters	CH																					0
2	1	216	vulture1-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
2	1	217	vulture2-UndeadScavenger-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
2	1	218	vulture3-HellBuzzard-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
2	1	219	vulture4-WingedNightmare-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
2	1	220	mosquito1-Sucker-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
2	1	221	mosquito2-Feeder-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
2	1	222	mosquito3-BloodHook-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
2	1	223	mosquito4-BloodWing-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
2	1	224	willowisp1-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	225	willowisp2-SwampGhost-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	226	willowisp3-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	227	willowisp4-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	228	arach1-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	229	arach2-SandFisher-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	230	arach3-PoisonSpinner-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	231	arach4-FlameSpider-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	232	arach5-SpiderMagus-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	233	thornhulk1-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
2	1	234	thornhulk2-BrambleHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
2	1	235	thornhulk3-Thrasher-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
2	1	236	thornhulk4-Spikefist-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
2	1	237	vampire1-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	238	vampire2-NightLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	239	vampire3-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	240	vampire4-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	241	vampire5-Banished-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	242	batdemon1-DesertWing-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	243	batdemon2-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	244	batdemon3-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	245	batdemon4-BloodDiver-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	246	batdemon5-DarkFamiliar-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	247	fetish1-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	248	fetish2-Fetish-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	249	fetish3-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	250	fetish4-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	251	fetish5-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	252	cain1-DeckardCain-NpcOutOfTown				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
2	1	253	gheed-Gheed-Npc				/Data/Global/Monsters	GH	NU	HTH		LIT																	0
2	1	254	akara-Akara-Npc				/Data/Global/Monsters	PS	NU	HTH		LIT																	0
2	1	255	chicken-dummy-Idle				/Data/Global/Monsters	CK	NU	HTH		LIT																	0
2	1	256	kashya-Kashya-Npc				/Data/Global/Monsters	RC	NU	HTH		LIT																	0
2	1	257	rat-dummy-Idle				/Data/Global/Monsters	RT	NU	HTH		LIT																	0
2	1	258	rogue1-Dummy-Idle				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
2	1	259	hellmeteor-Dummy-HellMeteor				/Data/Global/Monsters	K9																					0
2	1	260	charsi-Charsi-Npc				/Data/Global/Monsters	CI	NU	HTH		LIT																	0
2	1	261	warriv1-Warriv-Npc				/Data/Global/Monsters	WA	NU	HTH		LIT																	0
2	1	262	andariel-Andariel-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
2	1	263	bird1-dummy-Idle				/Data/Global/Monsters	BS	WL	HTH		LIT																	0
2	1	264	bird2-dummy-Idle				/Data/Global/Monsters	BL																					0
2	1	265	bat-dummy-Idle				/Data/Global/Monsters	B9	WL	HTH		LIT																	0
2	1	266	cr_archer1-DarkRanger-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	267	cr_archer2-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	268	cr_archer3-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	269	cr_archer4-BlackArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	270	cr_archer5-FleshArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	271	cr_lancer1-DarkSpearwoman-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	272	cr_lancer2-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	273	cr_lancer3-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	274	cr_lancer4-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	275	cr_lancer5-FleshLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	276	sk_archer1-SkeletonArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	277	sk_archer2-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	278	sk_archer3-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	279	sk_archer4-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	280	sk_archer5-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	281	warriv2-Warriv-Npc				/Data/Global/Monsters	WX	NU	HTH		LIT																	0
2	1	282	atma-Atma-Npc				/Data/Global/Monsters	AS	NU	HTH		LIT																	0
2	1	283	drognan-Drognan-Npc				/Data/Global/Monsters	DR	NU	HTH		LIT																	0
2	1	284	fara-Fara-Npc				/Data/Global/Monsters	OF	NU	HTH		LIT																	0
2	1	285	cow-dummy-Idle				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
2	1	286	maggotbaby1-SandMaggotYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
2	1	287	maggotbaby2-RockWormYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
2	1	288	maggotbaby3-DevourerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
2	1	289	maggotbaby4-GiantLampreyYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
2	1	290	maggotbaby5-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
2	1	291	camel-dummy-Idle				/Data/Global/Monsters	CM	NU	HTH		LIT																	0
2	1	292	blunderbore1-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
2	1	293	blunderbore2-Gorbelly-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
2	1	294	blunderbore3-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
2	1	295	blunderbore4-Urdar-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
2	1	296	maggotegg1-SandMaggotEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
2	1	297	maggotegg2-RockWormEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
2	1	298	maggotegg3-DevourerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
2	1	299	maggotegg4-GiantLampreyEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
2	1	300	maggotegg5-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
2	1	301	act2male-dummy-Towner				/Data/Global/Monsters	2M	NU	HTH	OLD	MED	MED						TUR										0
2	1	302	act2female-Dummy-Towner				/Data/Global/Monsters	2F	NU	HTH	LIT	LIT	LIT																0
2	1	303	act2child-dummy-Towner				/Data/Global/Monsters	2C																					0
2	1	304	greiz-Greiz-Npc				/Data/Global/Monsters	GR	NU	HTH		LIT																	0
2	1	305	elzix-Elzix-Npc				/Data/Global/Monsters	EL	NU	HTH		LIT																	0
2	1	306	geglash-Geglash-Npc				/Data/Global/Monsters	GE	NU	HTH		LIT																	0
2	1	307	jerhyn-Jerhyn-Npc				/Data/Global/Monsters	JE	NU	HTH		LIT																	0
2	1	308	lysander-Lysander-Npc				/Data/Global/Monsters	LY	NU	HTH		LIT																	0
2	1	309	act2guard1-Dummy-Towner				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
2	1	310	act2vendor1-dummy-Vendor				/Data/Global/Monsters	M1	NU	HTH		LIT																	0
2	1	311	act2vendor2-dummy-Vendor				/Data/Global/Monsters	M2	NU	HTH		LIT																	0
2	1	312	crownest1-FoulCrowNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
2	1	313	crownest2-BloodHawkNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
2	1	314	crownest3-BlackVultureNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
2	1	315	crownest4-CloudStalkerNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
2	1	316	meshif1-Meshif-Npc				/Data/Global/Monsters	MS	NU	HTH		LIT																	0
2	1	317	duriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
2	1	318	bonefetish1-Undead RatMan-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	319	bonefetish2-Undead Fetish-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	320	bonefetish3-Undead Flayer-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	321	bonefetish4-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	322	bonefetish5-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	323	darkguard1-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	324	darkguard2-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	325	darkguard3-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	326	darkguard4-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	327	darkguard5-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	328	bloodmage1-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	329	bloodmage2-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	330	bloodmage3-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	331	bloodmage4-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	332	bloodmage5-unused-Idle				/Data/Global/Monsters	xx																					0
2	1	333	maggot-Maggot-Idle				/Data/Global/Monsters	MA	NU	HTH		LIT																	0
2	1	334	sarcophagus-MummyGenerator-Sarcophagus				/Data/Global/Monsters	MG	NU	HTH		LIT																	0
2	1	335	radament-Radament-GreaterMummy				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
2	1	336	firebeast-unused-ElementalBeast				/Data/Global/Monsters	FM	NU	HTH		LIT																	0
2	1	337	iceglobe-unused-ElementalBeast				/Data/Global/Monsters	IM	NU	HTH		LIT																	0
2	1	338	lightningbeast-unused-ElementalBeast				/Data/Global/Monsters	xx																					0
2	1	339	poisonorb-unused-ElementalBeast				/Data/Global/Monsters	PM	NU	HTH		LIT																	0
2	1	340	flyingscimitar-FlyingScimitar-FlyingScimitar				/Data/Global/Monsters	ST	NU	HTH		LIT																	0
2	1	341	zealot1-Zakarumite-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
2	1	342	zealot2-Faithful-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
2	1	343	zealot3-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
2	1	344	cantor1-Sexton-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
2	1	345	cantor2-Cantor-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
2	1	346	cantor3-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
2	1	347	cantor4-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
2	1	348	mephisto-Mephisto-Mephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
2	1	349	diablo-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
2	1	350	cain2-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
2	1	351	cain3-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
2	1	352	cain4-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
2	1	353	frogdemon1-Swamp Dweller-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
2	1	354	frogdemon2-Bog Creature-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
2	1	355	frogdemon3-Slime Prince-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
2	1	356	summoner-Summoner-Summoner				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
2	1	357	tyrael1-tyrael-NpcStationary				/Data/Global/Monsters	TX	NU	HTH		LIT		LIT	LIT														0
2	1	358	asheara-asheara-Npc				/Data/Global/Monsters	AH	NU	HTH		LIT																	0
2	1	359	hratli-hratli-Npc				/Data/Global/Monsters	HR	NU	HTH		LIT																	0
2	1	360	alkor-alkor-Npc				/Data/Global/Monsters	AL	NU	HTH		LIT																	0
2	1	361	ormus-ormus-Npc				/Data/Global/Monsters	OR	NU	HTH		LIT																	0
2	1	362	izual-izual-Izual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
2	1	363	halbu-halbu-Npc				/Data/Global/Monsters	20	NU	HTH		LIT																	0
2	1	364	tentacle1-WaterWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
2	1	365	tentacle2-RiverStalkerLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
2	1	366	tentacle3-StygianWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
2	1	367	tentaclehead1-WaterWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
2	1	368	tentaclehead2-RiverStalkerHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
2	1	369	tentaclehead3-StygianWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
2	1	370	meshif2-meshif-Npc				/Data/Global/Monsters	M3	NU	HTH		LIT																	0
2	1	371	cain5-DeckardCain-Npc				/Data/Global/Monsters	1D	NU	HTH		LIT																	0
2	1	372	navi-navi-Navi				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
2	1	373	bloodraven-Bloodraven-BloodRaven				/Data/Global/Monsters	CR	NU	BOW	BRV	HVY	BRV	HVY	HVY	LIT	LBB		HVY	HVY									0
2	1	374	bug-Dummy-Idle				/Data/Global/Monsters	BG	NU	HTH		LIT																	0
2	1	375	scorpion-Dummy-Idle				/Data/Global/Monsters	DS	NU	HTH		LIT																	0
2	1	376	rogue2-RogueScout-GoodNpcRanged				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
2	1	377	roguehire-Dummy-Hireable				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
2	1	378	rogue3-Dummy-TownRogue				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
2	1	379	gargoyletrap-GargoyleTrap-GargoyleTrap				/Data/Global/Monsters	GT	NU	HTH		LIT																	0
2	1	380	skmage_pois1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	381	skmage_pois2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	382	skmage_pois3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	383	skmage_pois4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	384	fetishshaman1-RatManShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	385	fetishshaman2-FetishShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	386	fetishshaman3-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	387	fetishshaman4-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	388	fetishshaman5-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	389	larva-larva-Idle				/Data/Global/Monsters	LV	NU	HTH		LIT																	0
2	1	390	maggotqueen1-SandMaggotQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	391	maggotqueen2-RockWormQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	392	maggotqueen3-DevourerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	393	maggotqueen4-GiantLampreyQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	394	maggotqueen5-WorldKillerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	395	claygolem-ClayGolem-NecroPet				/Data/Global/Monsters	G1	NU	HTH		LIT																	0
2	1	396	bloodgolem-BloodGolem-NecroPet				/Data/Global/Monsters	G2	NU	HTH		LIT																	0
2	1	397	irongolem-IronGolem-NecroPet				/Data/Global/Monsters	G4	NU	HTH		LIT																	0
2	1	398	firegolem-FireGolem-NecroPet				/Data/Global/Monsters	G3	NU	HTH		LIT																	0
2	1	399	familiar-Dummy-Idle				/Data/Global/Monsters	FI	NU	HTH		LIT																	0
2	1	400	act3male-Dummy-Towner				/Data/Global/Monsters	N4	NU	HTH	BRD	HVY	HVY	HEV	HEV	FSH	SAK		TKT										0
2	1	401	baboon6-NightMarauder-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	402	act3female-Dummy-Towner				/Data/Global/Monsters	N3	NU	HTH	LIT	MTP	SRT			BSK	BSK												0
2	1	403	natalya-Natalya-Npc				/Data/Global/Monsters	TZ	NU	HTH		LIT																	0
2	1	404	vilemother1-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
2	1	405	vilemother2-StygianHag-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
2	1	406	vilemother3-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
2	1	407	vilechild1-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
2	1	408	vilechild2-StygianDog-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
2	1	409	vilechild3-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
2	1	410	fingermage1-Groper-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
2	1	411	fingermage2-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
2	1	412	fingermage3-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
2	1	413	regurgitator1-Corpulent-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
2	1	414	regurgitator2-CorpseSpitter-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
2	1	415	regurgitator3-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
2	1	416	doomknight1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	417	doomknight2-AbyssKnight-AbyssKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	418	doomknight3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	419	quillbear1-QuillBear-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
2	1	420	quillbear2-SpikeGiant-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
2	1	421	quillbear3-ThornBrute-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
2	1	422	quillbear4-RazorBeast-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
2	1	423	quillbear5-GiantUrchin-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
2	1	424	snake-Dummy-Idle				/Data/Global/Monsters	CO	NU	HTH		LIT																	0
2	1	425	parrot-Dummy-Idle				/Data/Global/Monsters	PR	WL	HTH		LIT																	0
2	1	426	fish-Dummy-Idle				/Data/Global/Monsters	FJ																					0
2	1	427	evilhole1-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
2	1	428	evilhole2-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
2	1	429	evilhole3-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
2	1	430	evilhole4-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
2	1	431	evilhole5-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
2	1	432	trap-firebolt-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
2	1	433	trap-horzmissile-a trap-Trap-RightArrow				/Data/Global/Monsters	9A																					0
2	1	434	trap-vertmissile-a trap-Trap-LeftArrow				/Data/Global/Monsters	9A																					0
2	1	435	trap-poisoncloud-a trap-Trap-Poison				/Data/Global/Monsters	9A																					0
2	1	436	trap-lightning-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
2	1	437	act2guard2-Kaelan-JarJar				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
2	1	438	invisospawner-Dummy-InvisoSpawner				/Data/Global/Monsters	K9																					0
2	1	439	diabloclone-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
2	1	440	suckernest1-SuckerNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
2	1	441	suckernest2-FeederNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
2	1	442	suckernest3-BloodHookNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
2	1	443	suckernest4-BloodWingNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
2	1	444	act2hire-Guard-Hireable				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
2	1	445	minispider-Dummy-Idle				/Data/Global/Monsters	LS	NU	HTH		LIT																	0
2	1	446	boneprison1--Idle				/Data/Global/Monsters	67	NU	HTH		LIT																	0
2	1	447	boneprison2--Idle				/Data/Global/Monsters	66	NU	HTH		LIT																	0
2	1	448	boneprison3--Idle				/Data/Global/Monsters	69	NU	HTH		LIT																	0
2	1	449	boneprison4--Idle				/Data/Global/Monsters	68	NU	HTH		LIT																	0
2	1	450	bonewall-Dummy-BoneWall				/Data/Global/Monsters	BW	NU	HTH		LIT																	0
2	1	451	councilmember1-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	452	councilmember2-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	453	councilmember3-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	454	turret1-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
2	1	455	turret2-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
2	1	456	turret3-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
2	1	457	hydra1-Hydra-Hydra				/Data/Global/Monsters	HX	NU	HTH		LIT							LIT										0
2	1	458	hydra2-Hydra-Hydra				/Data/Global/Monsters	21	NU	HTH		LIT							LIT										0
2	1	459	hydra3-Hydra-Hydra				/Data/Global/Monsters	HZ	NU	HTH		LIT							LIT										0
2	1	460	trap-melee-a trap-Trap-Melee				/Data/Global/Monsters	M4	A1	HTH		LIT																	0
2	1	461	seventombs-Dummy-7TIllusion				/Data/Global/Monsters	9A																					0
2	1	462	dopplezon-Dopplezon-Idle				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
2	1	463	valkyrie-Valkyrie-NecroPet				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
2	1	464	act2guard3-Dummy-Idle				/Data/Global/Monsters	SK																					0
2	1	465	act3hire-Iron Wolf-Hireable				/Data/Global/Monsters	IW	NU	1HS	LIT	LIT				WND		KIT											0
2	1	466	megademon1-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	467	megademon2-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	468	megademon3-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	469	necroskeleton-NecroSkeleton-NecroPet				/Data/Global/Monsters	SK	NU	1HS	DES	HVY	DES	DES	DES	SCM		KIT	DES	DES	LIT								0
2	1	470	necromage-NecroMage-NecroPet				/Data/Global/Monsters	SK	NU	HTH	DES	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
2	1	471	griswold-Griswold-Griswold				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
2	1	472	compellingorb-compellingorb-Idle				/Data/Global/Monsters	9a																					0
2	1	473	tyrael2-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
2	1	474	darkwanderer-youngdiablo-DarkWanderer				/Data/Global/Monsters	1Z	NU	HTH		LIT																	0
2	1	475	trap-nova-a trap-Trap-Nova				/Data/Global/Monsters	9A																					0
2	1	476	spiritmummy-Dummy-Idle				/Data/Global/Monsters	xx																					0
2	1	477	lightningspire-LightningSpire-ArcaneTower				/Data/Global/Monsters	AE	NU	HTH		LIT							LIT										0
2	1	478	firetower-FireTower-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
2	1	479	slinger1-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
2	1	480	slinger2-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
2	1	481	slinger3-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
2	1	482	slinger4-HellSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
2	1	483	act2guard4-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
2	1	484	act2guard5-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
2	1	485	skmage_cold1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
2	1	486	skmage_cold2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
2	1	487	skmage_cold3-BaalColdMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
2	1	488	skmage_cold4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
2	1	489	skmage_fire1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
2	1	490	skmage_fire2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
2	1	491	skmage_fire3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
2	1	492	skmage_fire4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
2	1	493	skmage_ltng1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
2	1	494	skmage_ltng2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
2	1	495	skmage_ltng3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
2	1	496	skmage_ltng4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
2	1	497	hellbovine-Hell Bovine-Skeleton				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
2	1	498	window1--Idle				/Data/Global/Monsters	VH	NU	HTH		LIT							LIT										0
2	1	499	window2--Idle				/Data/Global/Monsters	VJ	NU	HTH		LIT							LIT										0
2	1	500	slinger5-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
2	1	501	slinger6-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
2	1	502	fetishblow1-RatMan-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	503	fetishblow2-Fetish-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	504	fetishblow3-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	505	fetishblow4-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	506	fetishblow5-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	507	mephistospirit-Dummy-Spirit				/Data/Global/Monsters	M6	A1	HTH		LIT																	0
2	1	508	smith-The Smith-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
2	1	509	trappedsoul1-TrappedSoul-TrappedSoul				/Data/Global/Monsters	10	NU	HTH		LIT																	0
2	1	510	trappedsoul2-TrappedSoul-TrappedSoul				/Data/Global/Monsters	13	NU	HTH		LIT																	0
2	1	511	jamella-Jamella-Npc				/Data/Global/Monsters	ja	NU	HTH		LIT																	0
2	1	512	izualghost-Izual-NpcStationary				/Data/Global/Monsters	17	NU	HTH		LIT							LIT										0
2	1	513	fetish11-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	514	malachai-Malachai-Buffy				/Data/Global/Monsters	36	NU	HTH		LIT							LIT										0
2	1	515	hephasto-The Feature Creep-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
2	1	516	wakeofdestruction-Wake of Destruction-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
2	1	517	chargeboltsentry-Charged Bolt Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
2	1	518	lightningsentry-Lightning Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
2	1	519	bladecreeper-Blade Creeper-BladeCreeper				/Data/Global/Monsters	b8	NU	HTH		LIT							LIT										0
2	1	520	invisopet-Invis Pet-InvisoPet				/Data/Global/Monsters	k9																					0
2	1	521	infernosentry-Inferno Sentry-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
2	1	522	deathsentry-Death Sentry-DeathSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
2	1	523	shadowwarrior-Shadow Warrior-ShadowWarrior				/Data/Global/Monsters	k9																					0
2	1	524	shadowmaster-Shadow Master-ShadowMaster				/Data/Global/Monsters	k9																					0
2	1	525	druidhawk-Druid Hawk-Raven				/Data/Global/Monsters	hk	NU	HTH		LIT																	0
2	1	526	spiritwolf-Druid Spirit Wolf-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
2	1	527	fenris-Druid Fenris-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
2	1	528	spiritofbarbs-Spirit of Barbs-Totem				/Data/Global/Monsters	x4	NU	HTH		LIT																	0
2	1	529	heartofwolverine-Heart of Wolverine-Totem				/Data/Global/Monsters	x3	NU	HTH		LIT																	0
2	1	530	oaksage-Oak Sage-Totem				/Data/Global/Monsters	xw	NU	HTH		LIT																	0
2	1	531	plaguepoppy-Druid Plague Poppy-Vines				/Data/Global/Monsters	k9																					0
2	1	532	cycleoflife-Druid Cycle of Life-CycleOfLife				/Data/Global/Monsters	k9																					0
2	1	533	vinecreature-Vine Creature-CycleOfLife				/Data/Global/Monsters	k9																					0
2	1	534	druidbear-Druid Bear-DruidBear				/Data/Global/Monsters	b7	NU	HTH		LIT																	0
2	1	535	eagle-Eagle-Idle				/Data/Global/Monsters	eg	NU	HTH		LIT							LIT										0
2	1	536	wolf-Wolf-NecroPet				/Data/Global/Monsters	40	NU	HTH		LIT																	0
2	1	537	bear-Bear-NecroPet				/Data/Global/Monsters	TG	NU	HTH		LIT							LIT										0
2	1	538	barricadedoor1-Barricade Door-Idle				/Data/Global/Monsters	AJ	NU	HTH		LIT																	0
2	1	539	barricadedoor2-Barricade Door-Idle				/Data/Global/Monsters	AG	NU	HTH		LIT																	0
2	1	540	prisondoor-Prison Door-Idle				/Data/Global/Monsters	2Q	NU	HTH		LIT																	0
2	1	541	barricadetower-Barricade Tower-SiegeTower				/Data/Global/Monsters	ac	NU	HTH		LIT							LIT						LIT				0
2	1	542	reanimatedhorde1-RotWalker-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	543	reanimatedhorde2-ReanimatedHorde-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	544	reanimatedhorde3-ProwlingDead-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	545	reanimatedhorde4-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	546	reanimatedhorde5-DefiledWarrior-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	547	siegebeast1-Siege Beast-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
2	1	548	siegebeast2-CrushBiest-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
2	1	549	siegebeast3-BloodBringer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
2	1	550	siegebeast4-GoreBearer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
2	1	551	siegebeast5-DeamonSteed-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
2	1	552	snowyeti1-SnowYeti1-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
2	1	553	snowyeti2-SnowYeti2-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
2	1	554	snowyeti3-SnowYeti3-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
2	1	555	snowyeti4-SnowYeti4-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
2	1	556	wolfrider1-WolfRider1-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
2	1	557	wolfrider2-WolfRider2-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
2	1	558	wolfrider3-WolfRider3-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
2	1	559	minion1-Minionexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	560	minion2-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	561	minion3-IceBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	562	minion4-FireBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	563	minion5-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	564	minion6-IceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	565	minion7-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	566	minion8-GreaterIceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	567	suicideminion1-FanaticMinion-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	568	suicideminion2-BerserkSlayer-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	569	suicideminion3-ConsumedIceBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	570	suicideminion4-ConsumedFireBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	571	suicideminion5-FrenziedHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	572	suicideminion6-FrenziedIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	573	suicideminion7-InsaneHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	574	suicideminion8-InsaneIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
2	1	575	succubus1-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	576	succubus2-VileTemptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	577	succubus3-StygianHarlot-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	578	succubus4-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	579	succubus5-Blood Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	580	succubuswitch1-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	581	succubuswitch2-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	582	succubuswitch3-StygianFury-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	583	succubuswitch4-Blood Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	584	succubuswitch5-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	585	overseer1-OverSeer-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	586	overseer2-Lasher-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	587	overseer3-OverLord-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	588	overseer4-BloodBoss-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	589	overseer5-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	590	minionspawner1-MinionSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	591	minionspawner2-MinionSlayerSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	592	minionspawner3-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	593	minionspawner4-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	594	minionspawner5-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	595	minionspawner6-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	596	minionspawner7-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	597	minionspawner8-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	598	imp1-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	599	imp2-Imp2-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	600	imp3-Imp3-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	601	imp4-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	602	imp5-Imp5-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	603	catapult1-CatapultS-Catapult				/Data/Global/Monsters	65	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
2	1	604	catapult2-CatapultE-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
2	1	605	catapult3-CatapultSiege-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
2	1	606	catapult4-CatapultW-Catapult				/Data/Global/Monsters	ua	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT	LIT								0
2	1	607	frozenhorror1-Frozen Horror1-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
2	1	608	frozenhorror2-Frozen Horror2-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
2	1	609	frozenhorror3-Frozen Horror3-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
2	1	610	frozenhorror4-Frozen Horror4-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
2	1	611	frozenhorror5-Frozen Horror5-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
2	1	612	bloodlord1-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	613	bloodlord2-Blood Lord2-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	614	bloodlord3-Blood Lord3-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	615	bloodlord4-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	616	bloodlord5-Blood Lord5-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	617	larzuk-Larzuk-Npc				/Data/Global/Monsters	XR	NU	HTH		LIT																	0
2	1	618	drehya-Drehya-Npc				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
2	1	619	malah-Malah-Npc				/Data/Global/Monsters	XT	NU	HTH		LIT																	0
2	1	620	nihlathak-Nihlathak Town-Npc				/Data/Global/Monsters	0J	NU	HTH		LIT																	0
2	1	621	qual-kehk-Qual-Kehk-Npc				/Data/Global/Monsters	XV	NU	HTH		LIT																	0
2	1	622	catapultspotter1-Catapult Spotter S-CatapultSpotter				/Data/Global/Monsters	k9																					0
2	1	623	catapultspotter2-Catapult Spotter E-CatapultSpotter				/Data/Global/Monsters	k9																					0
2	1	624	catapultspotter3-Catapult Spotter Siege-CatapultSpotter				/Data/Global/Monsters	k9																					0
2	1	625	catapultspotter4-Catapult Spotter W-CatapultSpotter				/Data/Global/Monsters	k9																					0
2	1	626	cain6-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
2	1	627	tyrael3-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
2	1	628	act5barb1-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
2	1	629	act5barb2-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
2	1	630	barricadewall1-Barricade Wall Right-Idle				/Data/Global/Monsters	A6	NU	HTH		LIT																	0
2	1	631	barricadewall2-Barricade Wall Left-Idle				/Data/Global/Monsters	AK	NU	HTH		LIT																	0
2	1	632	nihlathakboss-Nihlathak-Nihlathak				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
2	1	633	drehyaiced-Drehya-NpcOutOfTown				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
2	1	634	evilhut-Evil hut-GenericSpawner				/Data/Global/Monsters	2T	NU	HTH		LIT							LIT										0
2	1	635	deathmauler1-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	636	deathmauler2-Death Mauler2-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	637	deathmauler3-Death Mauler3-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	638	deathmauler4-Death Mauler4-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	639	deathmauler5-Death Mauler5-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	640	act5pow-POW-Wussie				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
2	1	641	act5barb3-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
2	1	642	act5barb4-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
2	1	643	ancientstatue1-Ancient Statue 1-AncientStatue				/Data/Global/Monsters	0G	NU	HTH		LIT																	0
2	1	644	ancientstatue2-Ancient Statue 2-AncientStatue				/Data/Global/Monsters	0H	NU	HTH		LIT																	0
2	1	645	ancientstatue3-Ancient Statue 3-AncientStatue				/Data/Global/Monsters	0I	NU	HTH		LIT																	0
2	1	646	ancientbarb1-Ancient Barbarian 1-Ancient				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
2	1	647	ancientbarb2-Ancient Barbarian 2-Ancient				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
2	1	648	ancientbarb3-Ancient Barbarian 3-Ancient				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
2	1	649	baalthrone-Baal Throne-BaalThrone				/Data/Global/Monsters	41	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
2	1	650	baalcrab-Baal Crab-BaalCrab				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
2	1	651	baaltaunt-Baal Taunt-BaalTaunt				/Data/Global/Monsters	K9																					0
2	1	652	putriddefiler1-Putrid Defiler1-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
2	1	653	putriddefiler2-Putrid Defiler2-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
2	1	654	putriddefiler3-Putrid Defiler3-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
2	1	655	putriddefiler4-Putrid Defiler4-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
2	1	656	putriddefiler5-Putrid Defiler5-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
2	1	657	painworm1-Pain Worm1-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
2	1	658	painworm2-Pain Worm2-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
2	1	659	painworm3-Pain Worm3-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
2	1	660	painworm4-Pain Worm4-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
2	1	661	painworm5-Pain Worm5-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
2	1	662	bunny-dummy-Idle				/Data/Global/Monsters	48	NU	HTH		LIT																	0
2	1	663	baalhighpriest-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	664	venomlord-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				FLB													0
2	1	665	baalcrabstairs-Baal Crab to Stairs-BaalToStairs				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
2	1	666	act5hire1-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
2	1	667	act5hire2-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
2	1	668	baaltentacle1-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
2	1	669	baaltentacle2-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
2	1	670	baaltentacle3-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
2	1	671	baaltentacle4-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
2	1	672	baaltentacle5-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
2	1	673	injuredbarb1-dummy-Idle				/Data/Global/Monsters	6z	NU	HTH		LIT																	0
2	1	674	injuredbarb2-dummy-Idle				/Data/Global/Monsters	7j	NU	HTH		LIT																	0
2	1	675	injuredbarb3-dummy-Idle				/Data/Global/Monsters	7i	NU	HTH		LIT																	0
2	1	676	baalclone-Baal Crab Clone-BaalCrabClone				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
2	1	677	baalminion1-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
2	1	678	baalminion2-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
2	1	679	baalminion3-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
2	1	680	worldstoneeffect-dummy-Idle				/Data/Global/Monsters	K9																					0
2	1	681	sk_archer6-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	682	sk_archer7-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	683	sk_archer8-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	684	sk_archer9-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	685	sk_archer10-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	686	bighead6-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	687	bighead7-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	688	bighead8-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	689	bighead9-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	690	bighead10-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	691	goatman6-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	692	goatman7-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	693	goatman8-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	694	goatman9-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	695	goatman10-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
2	1	696	foulcrow5-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	697	foulcrow6-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	698	foulcrow7-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	699	foulcrow8-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
2	1	700	clawviper6-ClawViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	701	clawviper7-PitViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	702	clawviper8-Salamander-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	703	clawviper9-TombViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	704	clawviper10-SerpentMagus-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	705	sandraider6-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	706	sandraider7-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	707	sandraider8-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	708	sandraider9-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	709	sandraider10-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	710	deathmauler6-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	711	quillrat6-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	712	quillrat7-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	713	quillrat8-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	714	vulture5-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
2	1	715	thornhulk5-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
2	1	716	slinger7-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	717	slinger8-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	718	slinger9-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	719	cr_archer6-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	720	cr_archer7-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	721	cr_lancer6-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	722	cr_lancer7-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	723	cr_lancer8-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	724	blunderbore5-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
2	1	725	blunderbore6-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
2	1	726	skmage_fire5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
2	1	727	skmage_fire6-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
2	1	728	skmage_ltng5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
2	1	729	skmage_ltng6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
2	1	730	skmage_cold5-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		CLD	CLD						0
2	1	731	skmage_pois5-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	732	skmage_pois6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	733	pantherwoman5-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	734	pantherwoman6-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	735	sandleaper6-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	736	sandleaper7-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
2	1	737	wraith6-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	738	wraith7-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	739	wraith8-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	740	succubus6-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	741	succubus7-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	742	succubuswitch6-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	743	succubuswitch7-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	744	succubuswitch8-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	745	willowisp5-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	746	willowisp6-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	747	willowisp7-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	748	fallen6-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
2	1	749	fallen7-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
2	1	750	fallen8-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
2	1	751	fallenshaman6-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	752	fallenshaman7-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	753	fallenshaman8-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	754	skeleton6-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	755	skeleton7-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	756	batdemon6-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	757	batdemon7-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	758	bloodlord6-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	759	bloodlord7-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	760	scarab6-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	761	scarab7-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	762	fetish6-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	763	fetish7-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	764	fetish8-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
2	1	765	fetishblow6-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	766	fetishblow7-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	767	fetishblow8-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
2	1	768	fetishshaman6-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	769	fetishshaman7-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	770	fetishshaman8-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	771	baboon7-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	772	baboon8-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
2	1	773	unraveler6-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	774	unraveler7-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	775	unraveler8-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	776	unraveler9-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	777	zealot4-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
2	1	778	zealot5-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
2	1	779	cantor5-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
2	1	780	cantor6-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
2	1	781	vilemother4-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
2	1	782	vilemother5-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
2	1	783	vilechild4-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
2	1	784	vilechild5-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
2	1	785	sandmaggot6-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	786	maggotbaby6-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
2	1	787	maggotegg6-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
2	1	788	minion9-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	789	minion10-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	790	minion11-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	791	arach6-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	792	megademon4-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	793	megademon5-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	794	imp6-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	795	imp7-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	796	bonefetish6-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	797	bonefetish7-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
2	1	798	fingermage4-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
2	1	799	fingermage5-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
2	1	800	regurgitator4-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
2	1	801	vampire6-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	802	vampire7-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	803	vampire8-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	804	reanimatedhorde6-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	805	dkfig1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	806	dkfig2-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	807	dkmag1-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	808	dkmag2-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	809	mummy6-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	810	ubermephisto-Mephisto-UberMephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
2	1	811	uberdiablo-Diablo-UberDiablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
2	1	812	uberizual-izual-UberIzual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
2	1	813	uberandariel-Lilith-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
2	1	814	uberduriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
2	1	815	uberbaal-Baal Crab-UberBaal				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
2	1	816	demonspawner-Evil hut-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
2	1	817	demonhole-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
2	1	818	megademon6-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	819	dkmag3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	820	imp8-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	821	swarm5-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
2	1	822	sandmaggot7-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
2	1	823	arach7-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	824	scarab8-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	825	succubus8-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	826	succubuswitch9-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	827	corruptrogue6-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	828	cr_archer8-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	829	cr_lancer9-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
2	1	830	overseer6-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	831	skeleton8-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	832	sk_archer11-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
2	1	833	skmage_fire7-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
2	1	834	skmage_ltng7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
2	1	835	skmage_cold6-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
2	1	836	skmage_pois7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		POS	POS						0
2	1	837	vampire9-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
2	1	838	wraith9-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
2	1	839	willowisp8-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	840	Bishibosh-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	841	Bonebreak-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
2	1	842	Coldcrow-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
2	1	843	Rakanishu-SUPER UNIQUE				/Data/Global/Monsters	FA	NU	HTH		LIT				SWD		TCH	LIT										0
2	1	844	Treehead WoodFist-SUPER UNIQUE				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
2	1	845	Griswold-SUPER UNIQUE				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
2	1	846	The Countess-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	MED	LIT	MED	LIT	LIT	WHM			LIT	LIT									0
2	1	847	Pitspawn Fouldog-SUPER UNIQUE				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
2	1	848	Flamespike the Crawler-SUPER UNIQUE				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
2	1	849	Boneash-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
2	1	850	Radament-SUPER UNIQUE				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
2	1	851	Bloodwitch the Wild-SUPER UNIQUE				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
2	1	852	Fangskin-SUPER UNIQUE				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
2	1	853	Beetleburst-SUPER UNIQUE				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
2	1	854	Leatherarm-SUPER UNIQUE				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
2	1	855	Coldworm the Burrower-SUPER UNIQUE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
2	1	856	Fire Eye-SUPER UNIQUE				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
2	1	857	Dark Elder-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	858	The Summoner-SUPER UNIQUE				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
2	1	859	Ancient Kaa the Soulless-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	860	The Smith-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
2	1	861	Web Mage the Burning-SUPER UNIQUE				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
2	1	862	Witch Doctor Endugu-SUPER UNIQUE				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
2	1	863	Stormtree-SUPER UNIQUE				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
2	1	864	Sarina the Battlemaid-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
2	1	865	Icehawk Riftwing-SUPER UNIQUE				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
2	1	866	Ismail Vilehand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	867	Geleb Flamefinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	868	Bremm Sparkfist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	869	Toorc Icefist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	870	Wyand Voidfinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	871	Maffer Dragonhand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	872	Winged Death-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	873	The Tormentor-SUPER UNIQUE				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
2	1	874	Taintbreeder-SUPER UNIQUE				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
2	1	875	Riftwraith the Cannibal-SUPER UNIQUE				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
2	1	876	Infector of Souls-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	877	Lord De Seis-SUPER UNIQUE				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
2	1	878	Grand Vizier of Chaos-SUPER UNIQUE				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
2	1	879	The Cow King-SUPER UNIQUE				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
2	1	880	Corpsefire-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
2	1	881	The Feature Creep-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
2	1	882	Siege Boss-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	883	Ancient Barbarian 1-SUPER UNIQUE				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
2	1	884	Ancient Barbarian 2-SUPER UNIQUE				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
2	1	885	Ancient Barbarian 3-SUPER UNIQUE				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
2	1	886	Axe Dweller-SUPER UNIQUE				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
2	1	887	Bonesaw Breaker-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	888	Dac Farren-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	889	Megaflow Rectifier-SUPER UNIQUE				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
2	1	890	Eyeback Unleashed-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	891	Threash Socket-SUPER UNIQUE				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
2	1	892	Pindleskin-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
2	1	893	Snapchip Shatter-SUPER UNIQUE				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
2	1	894	Anodized Elite-SUPER UNIQUE				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
2	1	895	Vinvear Molech-SUPER UNIQUE				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
2	1	896	Sharp Tooth Sayer-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
2	1	897	Magma Torquer-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
2	1	898	Blaze Ripper-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
2	1	899	Frozenstein-SUPER UNIQUE				/Data/Global/Monsters	io	NU	HTH		LIT																	0
2	1	900	Nihlathak Boss-SUPER UNIQUE				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
2	1	901	Baal Subject 1-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
2	1	902	Baal Subject 2-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
2	1	903	Baal Subject 3-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
2	1	904	Baal Subject 4-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
2	1	905	Baal Subject 5-SUPER UNIQUE				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
2	2	0	Trap Door (74)	74			/Data/Global/Objects	TD	NU	HTH		LIT																	0
2	2	1	torch 1 tiki (37)	37			/Data/Global/Objects	TO	ON	HTH		LIT																	0
2	2	2	Teleport Pad 1 (192)	192			/Data/Global/Objects	7H	NU	HTH		LIT							LIT	LIT									0
2	2	3	Teleport Pad 2 (304)	304			/Data/Global/Objects	7H	NU	HTH		LIT							LIT	LIT									0
2	2	4	Teleport Pad 3 (305)	305			/Data/Global/Objects	AA	NU	HTH		LIT							LIT	LIT									0
2	2	5	Teleport Pad 4 (306)	306			/Data/Global/Objects	AA	NU	HTH		LIT							LIT	LIT									0
2	2	6	brazier 3 (101)	101			/Data/Global/Objects	B3	OP	HTH		LIT							LIT										0
2	2	7	brazier floor (102)	102			/Data/Global/Objects	FB	ON	HTH		LIT							LIT										0
2	2	8	invisible town sound (78)	78			/Data/Global/Objects	TA																					0
2	2	9	flies (103)	103			/Data/Global/Objects	FL	NU	HTH		LIT																	0
2	2	10	waypoint (156)	156			/Data/Global/Objects	WM	ON	HTH		LIT							LIT										0
2	2	11	-580	580																									0
2	2	12	Well, cathedralwell inside (132)	132			/Data/Global/Objects	ZC	NU	HTH		LIT																	0
2	2	13	Door, secret 1 (129)	129			/Data/Global/Objects	H2	OP	HTH		LIT																	0
2	2	14	Horazon's Journal (357)	357			/Data/Global/Objects	TT	NU	HTH		LIT																	0
2	2	15	Door, Tyrael's door (153)	153			/Data/Global/Objects	DX	OP	HTH		LIT																	0
2	2	16	Jerhyn, placeholder 1 (121)	121		1	/Data/Global/Monsters	JE	NU	HTH		LIT																	0
2	2	17	Jerhyn, placeholder 2 (122)	122		7	/Data/Global/Monsters	JE	NU	HTH		LIT																	0
2	2	18	Closed Door, slimedoor R (229)	229			/Data/Global/Objects	SQ	OP	HTH		LIT																	0
2	2	19	Closed Door, slimedoor L (230)	230			/Data/Global/Objects	SY	OP	HTH		LIT																	0
2	2	20	a Trap, test data floortrap (196)	196			/Data/Global/Objects	A5	OP	HTH		LIT																	0
2	2	21	Your Private Stash (267)	267			/Data/Global/Objects	B6	NU	HTH		LIT																	0
2	2	22	a Trap, spikes tombs floortrap (261)	261			/Data/Global/Objects	A7	OP	HTH		LIT																	0
2	2	23	Tainted Sun Altar (149)	149			/Data/Global/Objects	ZA	OP	HTH		LIT							LIT										0
2	2	24	gold placeholder (269)	269			/Data/Global/Objects	1G																					0
2	2	25	Large Urn, urn 1 (4)	4			/Data/Global/Objects	U1	OP	HTH		LIT																	0
2	2	26	Corona, urn 2 (9)	9			/Data/Global/Objects	U2	OP	HTH		LIT																	0
2	2	27	Urn, urn 3 (52)	52			/Data/Global/Objects	U3	OP	HTH		LIT																	0
2	2	28	Large Urn, urn 4 (94)	94			/Data/Global/Objects	U4	OP	HTH		LIT																	0
2	2	29	Large Urn, urn 5 (95)	95			/Data/Global/Objects	U5	OP	HTH		LIT																	0
2	2	30	Jug, desert 1 (142)	142			/Data/Global/Objects	Q4	OP	HTH		LIT																	0
2	2	31	Jug, desert 2 (143)	143			/Data/Global/Objects	Q5	OP	HTH		LIT																	0
2	2	32	Chest, R Large (5)	5			/Data/Global/Objects	L1	OP	HTH		LIT																	0
2	2	33	Chest, L Large 1 (6)	6			/Data/Global/Objects	L2	OP	HTH		LIT																	0
2	2	34	Chest, L Large tomb 1 (87)	87			/Data/Global/Objects	CA	OP	HTH		LIT																	0
2	2	35	Chest, R Large tomb 2 (88)	88			/Data/Global/Objects	CB	OP	HTH		LIT																	0
2	2	36	Chest, R Med  (146)	146			/Data/Global/Objects	Q9	OP	HTH		LIT																	0
2	2	37	Chest, R Med  (146)	146			/Data/Global/Objects	Q9	OP	HTH		LIT																	0
2	2	38	Chest, R Large desert tomb (147)	147			/Data/Global/Objects	Q7	OP	HTH		LIT																	0
2	2	39	Chest, L Large desert tomb (148)	148			/Data/Global/Objects	Q8	OP	HTH		LIT																	0
2	2	40	Chest, 1L general (240)	240			/Data/Global/Objects	CY	OP	HTH		LIT																	0
2	2	41	Chest, 2R general (241)	241			/Data/Global/Objects	CX	OP	HTH		LIT																	0
2	2	42	Chest, 3R general (242)	242			/Data/Global/Objects	CU	OP	HTH		LIT																	0
2	2	43	Chest, 3L general (243)	243			/Data/Global/Objects	CD	OP	HTH		LIT																	0
2	2	44	Chest, L Med (176)	176			/Data/Global/Objects	C8	OP	HTH		LIT																	0
2	2	45	Chest, L Large 2 (177)	177			/Data/Global/Objects	C9	OP	HTH		LIT																	0
2	2	46	Chest, L Tallskinney (198)	198			/Data/Global/Objects	C0	OP	HTH		LIT																	0
2	2	47	Rat's Nest (246)	246			/Data/Global/Objects	RA	OP	HTH		LIT																	0
2	2	48	brazier (29)	29			/Data/Global/Objects	BR	OP	HTH		LIT							LIT										0
2	2	49	Flame, fire small (160)	160			/Data/Global/Objects	FX	OP	HTH		LIT																	0
2	2	50	Flame, fire medium (161)	161			/Data/Global/Objects	FY	OP	HTH		LIT																	0
2	2	51	Fire, fire large (162)	162			/Data/Global/Objects	FZ	OP	HTH		LIT																	0
2	2	52	flame, no damage (273)	273			/Data/Global/Objects	F8	NU	HTH		LIT																	0
2	2	53	brazier celler (283)	283			/Data/Global/Objects	BI	OP	HTH		LIT																	0
2	2	54	Shrine, bull health tombs (85)	85			/Data/Global/Objects	BC	OP	HTH		LIT																	0
2	2	55	stele, magic shrine stone desert (86)	86			/Data/Global/Objects	SG	OP	HTH		LIT																	0
2	2	56	Shrine, palace health R harom arcane (109)	109			/Data/Global/Objects	P2	OP	HTH		LIT																	0
2	2	57	Shrine, snake woman magic tomb arcane (116)	116			/Data/Global/Objects	SN	OP	HTH		LIT							LIT										0
2	2	58	Shrine, dshrine2 (134)	134			/Data/Global/Objects	ZS	OP	HTH		LIT							LIT										0
2	2	59	Shrine, desertshrine 3 (135)	135			/Data/Global/Objects	ZR	OP	HTH		LIT							LIT										0
2	2	60	Shrine, dshrine 1 a (136)	136			/Data/Global/Objects	ZD	OP	HTH		LIT																	0
2	2	61	Shrine, dshrine 1 b (150)	150			/Data/Global/Objects	ZV	OP	HTH		LIT							LIT	LIT									0
2	2	62	Shrine, dshrine 4 (151)	151			/Data/Global/Objects	ZE	OP	HTH		LIT							LIT										0
2	2	63	Shrine, health well desert (172)	172			/Data/Global/Objects	MK	OP	HTH		LIT																	0
2	2	64	Shrine, mana well7 desert (173)	173			/Data/Global/Objects	MI	OP	HTH		LIT																	0
2	2	65	Shrine, magic shrine sewers (279)	279			/Data/Global/Objects	WJ	OP	HTH		LIT							LIT	LIT									0
2	2	66	Shrine, healthwell sewers (280)	280			/Data/Global/Objects	WK	OP	HTH		LIT																	0
2	2	67	Shrine, manawell sewers (281)	281			/Data/Global/Objects	WL	OP	HTH		LIT																	0
2	2	68	Shrine, magic shrine sewers dungeon (282)	282			/Data/Global/Objects	WS	OP	HTH		LIT							LIT	LIT	LIT								0
2	2	69	Shrine, mana well3 tomb (166)	166			/Data/Global/Objects	MF	OP	HTH		LIT																	0
2	2	70	Shrine, mana well4 harom (167)	167			/Data/Global/Objects	MH	OP	HTH		LIT																	0
2	2	71	Well, fountain2 desert tomb (113)	113			/Data/Global/Objects	F4	NU	HTH		LIT																	0
2	2	72	Well, desertwell tomb (137)	137			/Data/Global/Objects	ZL	NU	HTH		LIT																	0
2	2	73	Sarcophagus, mummy coffin L tomb (89)	89			/Data/Global/Objects	MC	OP	HTH		LIT																	0
2	2	74	Armor stand, 1 R (104)	104			/Data/Global/Objects	A3	NU	HTH		LIT																	0
2	2	75	Armor stand, 2 L (105)	105			/Data/Global/Objects	A4	NU	HTH		LIT																	0
2	2	76	Weapon Rack, 1 R (106)	106			/Data/Global/Objects	W1	NU	HTH		LIT																	0
2	2	77	Weapon Rack, 2 L (107)	107			/Data/Global/Objects	W2	NU	HTH		LIT																	0
2	2	78	Corpse, guard (154)	154			/Data/Global/Objects	GC	OP	HTH		LIT																	0
2	2	79	Skeleton (171)	171			/Data/Global/Objects	SX	OP	HTH		LIT																	0
2	2	80	Guard Corpse, on stick (178)	178			/Data/Global/Objects	GS	OP	HTH		LIT																	0
2	2	81	Corpse, guard 2 (270)	270			/Data/Global/Objects	GF	OP	HTH		LIT																	0
2	2	82	Corpse, villager 1 (271)	271			/Data/Global/Objects	DG	OP	HTH		LIT																	0
2	2	83	Corpse, villager 2 (272)	272			/Data/Global/Objects	DF	OP	HTH		LIT																	0
2	2	84	Goo Pile, for sand maggot lair (266)	266			/Data/Global/Objects	GP	OP	HTH		LIT																	0
2	2	85	Hidden Stash, tiny pixel shaped (274)	274			/Data/Global/Objects	F9	NU	HTH		LIT																	0
2	2	86	Rat's Nest, sewers (244)	244			/Data/Global/Objects	RN	OP	HTH		LIT																	0
2	2	87	Sarcophagus, anubis coffin tomb (284)	284			/Data/Global/Objects	QC	OP	HTH		LIT																	0
2	2	88	waypoint, celler (288)	288			/Data/Global/Objects	W7	ON	HTH		LIT							LIT										0
2	2	89	Portal to, arcane portal (298)	298			/Data/Global/Objects	AY	ON	HTH		LIT							LIT	LIT									0
2	2	90	Bed, harum (289)	289			/Data/Global/Objects	UB	OP	HTH		LIT																	0
2	2	91	wall torch L for tombs (296)	296			/Data/Global/Objects	QD	NU	HTH		LIT							LIT										0
2	2	92	wall torch R for tombs (297)	297			/Data/Global/Objects	QE	NU	HTH		LIT							LIT										0
2	2	93	brazier small desert town tombs (287)	287			/Data/Global/Objects	BQ	NU	HTH		LIT							LIT										0
2	2	94	brazier tall desert town tombs (286)	286			/Data/Global/Objects	BO	NU	HTH		LIT							LIT										0
2	2	95	brazier general sewers tomb desert (285)	285			/Data/Global/Objects	BM	NU	HTH		LIT							LIT										0
2	2	96	Closed Door, iron grate L (290)	290			/Data/Global/Objects	DV	OP	HTH		LIT																	0
2	2	97	Closed Door, iron grate R (291)	291			/Data/Global/Objects	DN	OP	HTH		LIT																	0
2	2	98	Door, wooden grate L (292)	292			/Data/Global/Objects	DP	OP	HTH		LIT																	0
2	2	99	Door, wooden grate R (293)	293			/Data/Global/Objects	DT	OP	HTH		LIT																	0
2	2	100	Door, wooden L (294)	294			/Data/Global/Objects	DK	OP	HTH		LIT																	0
2	2	101	Closed Door, wooden R (295)	295			/Data/Global/Objects	DL	OP	HTH		LIT																	0
2	2	102	Shrine, arcane (133)	133			/Data/Global/Objects	AZ	NU	HTH		LIT							LIT	LIT	LIT								0
2	2	103	Magic Shrine, arcane (303)	303			/Data/Global/Objects	HD	OP	HTH		LIT							LIT										0
2	2	104	Magic Shrine, haram 1 (299)	299			/Data/Global/Objects	HB	OP	HTH		LIT							LIT										0
2	2	105	Magic Shrine, haram 2 (300)	300			/Data/Global/Objects	HC	OP	HTH		LIT							LIT										0
2	2	106	maggot well health (301)	301			/Data/Global/Objects	QF	OP	HTH		LIT																	0
2	2	107	Shrine, maggot well mana (302)	302			/Data/Global/Objects	QG	OP	HTH		LIT																	0
2	2	108	-581	581																									0
2	2	109	Chest, horadric cube (354)	354			/Data/Global/Objects	XK	OP	HTH		LIT																	0
2	2	110	Tomb signs in Arcane (582)	582			/Data/Global/Objects	7C	NU	HTH		LIT																	0
2	2	111	Dead Guard, harem 1 (314)	314			/Data/Global/Objects	QH	NU	HTH		LIT																	0
2	2	112	Dead Guard, harem 2 (315)	315			/Data/Global/Objects	QI	NU	HTH		LIT																	0
2	2	113	Dead Guard, harem 3 (316)	316			/Data/Global/Objects	QJ	NU	HTH		LIT																	0
2	2	114	Dead Guard, harem 4 (317)	317			/Data/Global/Objects	QK	NU	HTH		LIT																	0
2	2	115	Waypoint, sewer (323)	323			/Data/Global/Objects	QM	ON	HTH		LIT							LIT										0
2	2	116	Well, tomb (322)	322			/Data/Global/Objects	HU	NU	HTH		LIT																	0
2	2	117	drinker (110)	110			/Data/Global/Objects	N5	S1	HTH		LIT																	0
2	2	118	gesturer (112)	112			/Data/Global/Objects	N6	S2	HTH		LIT																	0
2	2	119	turner (114)	114			/Data/Global/Objects	N7	S1	HTH		LIT																	0
2	2	120	Chest, horadric scroll (355)	355			/Data/Global/Objects	XK	OP	HTH		LIT																	0
2	2	121	Chest, staff of kings (356)	356			/Data/Global/Objects	XK	OP	HTH		LIT																	0
2	2	122	Horazon's Journal (357)	357			/Data/Global/Objects	TT	OP	HTH		LIT																	0
2	2	123	helllight source 1 (351)	351			/Data/Global/Objects	SS																					0
2	2	124	helllight source 2 (352)	352			/Data/Global/Objects	SS																					0
2	2	125	helllight source 3 (353)	353			/Data/Global/Objects	SS																					0
2	2	126	orifice, place Horadric Staff (152)	152			/Data/Global/Objects	HA	NU	HTH		LIT																	0
2	2	127	fog water (374)	374			/Data/Global/Objects	UD	NU	HTH		LIT																	0
2	2	128	Chest, arcane big L (387)	387			/Data/Global/Objects	Y7	OP	HTH		LIT																	0
2	2	129	Chest, arcane big R (389)	389			/Data/Global/Objects	Y9	OP	HTH		LIT																	0
2	2	130	Chest, arcane small L (390)	390			/Data/Global/Objects	YA	OP	HTH		LIT																	0
2	2	131	Chest, arcane small R (391)	391			/Data/Global/Objects	YC	OP	HTH		LIT																	0
2	2	132	Casket, arcane (388)	388			/Data/Global/Objects	Y8	OP	HTH		LIT																	0
2	2	133	Chest, sparkly (397)	397			/Data/Global/Objects	YF	OP	HTH		LIT																	0
2	2	134	Waypoint, valley (402)	402			/Data/Global/Objects	YI	ON	HTH		LIT							LIT										0
2	2	135	ACT 2 TABLE SKIP IT	0																									0
2	2	136	ACT 2 TABLE SKIP IT	0																									0
2	2	137	ACT 2 TABLE SKIP IT	0																									0
2	2	138	ACT 2 TABLE SKIP IT	0																									0
2	2	139	ACT 2 TABLE SKIP IT	0																									0
2	2	140	ACT 2 TABLE SKIP IT	0																									0
2	2	141	ACT 2 TABLE SKIP IT	0																									0
2	2	142	ACT 2 TABLE SKIP IT	0																									0
2	2	143	ACT 2 TABLE SKIP IT	0																									0
2	2	144	ACT 2 TABLE SKIP IT	0																									0
2	2	145	ACT 2 TABLE SKIP IT	0																									0
2	2	146	ACT 2 TABLE SKIP IT	0																									0
2	2	147	ACT 2 TABLE SKIP IT	0																									0
2	2	148	ACT 2 TABLE SKIP IT	0																									0
2	2	149	ACT 2 TABLE SKIP IT	0																									0
2	2	150	Dummy-test data SKIPT IT				/Data/Global/Objects	NU0																					
2	2	151	Casket-Casket #5				/Data/Global/Objects	C5	OP	HTH		LIT																	
2	2	152	Shrine-Shrine				/Data/Global/Objects	SF	OP	HTH		LIT																	
2	2	153	Casket-Casket #6				/Data/Global/Objects	C6	OP	HTH		LIT																	
2	2	154	LargeUrn-Urn #1				/Data/Global/Objects	U1	OP	HTH		LIT																	
2	2	155	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
2	2	156	chest-LargeChestL				/Data/Global/Objects	L2	OP	HTH		LIT																	
2	2	157	Barrel-Barrel				/Data/Global/Objects	B1	OP	HTH		LIT																	
2	2	158	TowerTome-Tower Tome				/Data/Global/Objects	TT	OP	HTH		LIT																	
2	2	159	Urn-Urn #2				/Data/Global/Objects	U2	OP	HTH		LIT																	
2	2	160	Dummy-Bench				/Data/Global/Objects	BE	NU	HTH		LIT																	
2	2	161	Barrel-BarrelExploding				/Data/Global/Objects	BX	OP	HTH		LIT							LIT	LIT									
2	2	162	Dummy-RogueFountain				/Data/Global/Objects	FN	NU	HTH		LIT																	
2	2	163	Door-Door Gate Left				/Data/Global/Objects	D1	OP	HTH		LIT																	
2	2	164	Door-Door Gate Right				/Data/Global/Objects	D2	OP	HTH		LIT																	
2	2	165	Door-Door Wooden Left				/Data/Global/Objects	D3	OP	HTH		LIT																	
2	2	166	Door-Door Wooden Right				/Data/Global/Objects	D4	OP	HTH		LIT																	
2	2	167	StoneAlpha-StoneAlpha				/Data/Global/Objects	S1	OP	HTH		LIT																	
2	2	168	StoneBeta-StoneBeta				/Data/Global/Objects	S2	OP	HTH		LIT																	
2	2	169	StoneGamma-StoneGamma				/Data/Global/Objects	S3	OP	HTH		LIT																	
2	2	170	StoneDelta-StoneDelta				/Data/Global/Objects	S4	OP	HTH		LIT																	
2	2	171	StoneLambda-StoneLambda				/Data/Global/Objects	S5	OP	HTH		LIT																	
2	2	172	StoneTheta-StoneTheta				/Data/Global/Objects	S6	OP	HTH		LIT																	
2	2	173	Door-Door Courtyard Left				/Data/Global/Objects	D5	OP	HTH		LIT																	
2	2	174	Door-Door Courtyard Right				/Data/Global/Objects	D6	OP	HTH		LIT																	
2	2	175	Door-Door Cathedral Double				/Data/Global/Objects	D7	OP	HTH		LIT																	
2	2	176	Gibbet-Cain's Been Captured				/Data/Global/Objects	GI	OP	HTH		LIT																	
2	2	177	Door-Door Monastery Double Right				/Data/Global/Objects	D8	OP	HTH		LIT																	
2	2	178	HoleAnim-Hole in Ground				/Data/Global/Objects	HI	OP	HTH		LIT																	
2	2	179	Dummy-Brazier				/Data/Global/Objects	BR	ON	HTH		LIT							LIT										
2	2	180	Inifuss-inifuss tree				/Data/Global/Objects	IT	NU	HTH		LIT																	
2	2	181	Dummy-Fountain				/Data/Global/Objects	BF	NU	HTH		LIT																	
2	2	182	Dummy-crucifix				/Data/Global/Objects	CL	NU	HTH		LIT																	
2	2	183	Dummy-Candles1				/Data/Global/Objects	A1	NU	HTH		LIT																	
2	2	184	Dummy-Candles2				/Data/Global/Objects	A2	NU	HTH		LIT																	
2	2	185	Dummy-Standard1				/Data/Global/Objects	N1	NU	HTH		LIT																	
2	2	186	Dummy-Standard2				/Data/Global/Objects	N2	NU	HTH		LIT																	
2	2	187	Dummy-Torch1 Tiki				/Data/Global/Objects	TO	ON	HTH		LIT																	
2	2	188	Dummy-Torch2 Wall				/Data/Global/Objects	WT	ON	HTH		LIT																	
2	2	189	fire-RogueBonfire				/Data/Global/Objects	RB	ON	HTH		LIT																	
2	2	190	Dummy-River1				/Data/Global/Objects	R1	NU	HTH		LIT																	
2	2	191	Dummy-River2				/Data/Global/Objects	R2	NU	HTH		LIT																	
2	2	192	Dummy-River3				/Data/Global/Objects	R3	NU	HTH		LIT																	
2	2	193	Dummy-River4				/Data/Global/Objects	R4	NU	HTH		LIT																	
2	2	194	Dummy-River5				/Data/Global/Objects	R5	NU	HTH		LIT																	
2	2	195	AmbientSound-ambient sound generator				/Data/Global/Objects	S1	OP	HTH		LIT																	
2	2	196	Crate-Crate				/Data/Global/Objects	CT	OP	HTH		LIT																	
2	2	197	Door-Andariel's Door				/Data/Global/Objects	AD	NU	HTH		LIT																	
2	2	198	Dummy-RogueTorch				/Data/Global/Objects	T1	NU	HTH		LIT																	
2	2	199	Dummy-RogueTorch				/Data/Global/Objects	T2	NU	HTH		LIT																	
2	2	200	Casket-CasketR				/Data/Global/Objects	C1	OP	HTH		LIT																	
2	2	201	Casket-CasketL				/Data/Global/Objects	C2	OP	HTH		LIT																	
2	2	202	Urn-Urn #3				/Data/Global/Objects	U3	OP	HTH		LIT																	
2	2	203	Casket-Casket				/Data/Global/Objects	C4	OP	HTH		LIT																	
2	2	204	RogueCorpse-Rogue corpse 1				/Data/Global/Objects	Z1	NU	HTH		LIT																	
2	2	205	RogueCorpse-Rogue corpse 2				/Data/Global/Objects	Z2	NU	HTH		LIT																	
2	2	206	RogueCorpse-rolling rogue corpse				/Data/Global/Objects	Z5	OP	HTH		LIT																	
2	2	207	CorpseOnStick-rogue on a stick 1				/Data/Global/Objects	Z3	OP	HTH		LIT																	
2	2	208	CorpseOnStick-rogue on a stick 2				/Data/Global/Objects	Z4	OP	HTH		LIT																	
2	2	209	Portal-Town portal				/Data/Global/Objects	TP	ON	HTH	LIT	LIT																	
2	2	210	Portal-Permanent town portal				/Data/Global/Objects	PP	ON	HTH	LIT	LIT																	
2	2	211	Dummy-Invisible object				/Data/Global/Objects	SS																					
2	2	212	Door-Door Cathedral Left				/Data/Global/Objects	D9	OP	HTH		LIT																	
2	2	213	Door-Door Cathedral Right				/Data/Global/Objects	DA	OP	HTH		LIT																	
2	2	214	Door-Door Wooden Left #2				/Data/Global/Objects	DB	OP	HTH		LIT																	
2	2	215	Dummy-invisible river sound1				/Data/Global/Objects	X1																					
2	2	216	Dummy-invisible river sound2				/Data/Global/Objects	X2																					
2	2	217	Dummy-ripple				/Data/Global/Objects	1R	NU	HTH		LIT																	
2	2	218	Dummy-ripple				/Data/Global/Objects	2R	NU	HTH		LIT																	
2	2	219	Dummy-ripple				/Data/Global/Objects	3R	NU	HTH		LIT																	
2	2	220	Dummy-ripple				/Data/Global/Objects	4R	NU	HTH		LIT																	
2	2	221	Dummy-forest night sound #1				/Data/Global/Objects	F1																					
2	2	222	Dummy-forest night sound #2				/Data/Global/Objects	F2																					
2	2	223	Dummy-yeti dung				/Data/Global/Objects	YD	NU	HTH		LIT																	
2	2	224	TrappDoor-Trap Door				/Data/Global/Objects	TD	ON	HTH		LIT																	
2	2	225	Door-Door by Dock, Act 2				/Data/Global/Objects	DD	ON	HTH		LIT																	
2	2	226	Dummy-sewer drip				/Data/Global/Objects	SZ																					
2	2	227	Shrine-healthorama				/Data/Global/Objects	SH	OP	HTH		LIT																	
2	2	228	Dummy-invisible town sound				/Data/Global/Objects	TA																					
2	2	229	Casket-casket #3				/Data/Global/Objects	C3	OP	HTH		LIT																	
2	2	230	Obelisk-obelisk				/Data/Global/Objects	OB	OP	HTH		LIT																	
2	2	231	Shrine-forest altar				/Data/Global/Objects	AF	OP	HTH		LIT																	
2	2	232	Dummy-bubbling pool of blood				/Data/Global/Objects	B2	NU	HTH		LIT																	
2	2	233	Shrine-horn shrine				/Data/Global/Objects	HS	OP	HTH		LIT																	
2	2	234	Shrine-healing well				/Data/Global/Objects	HW	OP	HTH		LIT																	
2	2	235	Shrine-bull shrine,health, tombs				/Data/Global/Objects	BC	OP	HTH		LIT																	
2	2	236	Dummy-stele,magic shrine, stone, desert				/Data/Global/Objects	SG	OP	HTH		LIT																	
2	2	237	Chest3-tombchest 1, largechestL				/Data/Global/Objects	CA	OP	HTH		LIT																	
2	2	238	Chest3-tombchest 2 largechestR				/Data/Global/Objects	CB	OP	HTH		LIT																	
2	2	239	Sarcophagus-mummy coffinL, tomb				/Data/Global/Objects	MC	OP	HTH		LIT																	
2	2	240	Obelisk-desert obelisk				/Data/Global/Objects	DO	OP	HTH		LIT																	
2	2	241	Door-tomb door left				/Data/Global/Objects	TL	OP	HTH		LIT																	
2	2	242	Door-tomb door right				/Data/Global/Objects	TR	OP	HTH		LIT																	
2	2	243	Shrine-mana shrineforinnerhell				/Data/Global/Objects	iz	OP	HTH		LIT							LIT										
2	2	244	LargeUrn-Urn #4				/Data/Global/Objects	U4	OP	HTH		LIT																	
2	2	245	LargeUrn-Urn #5				/Data/Global/Objects	U5	OP	HTH		LIT																	
2	2	246	Shrine-health shrineforinnerhell				/Data/Global/Objects	iy	OP	HTH		LIT							LIT										
2	2	247	Shrine-innershrinehell				/Data/Global/Objects	ix	OP	HTH		LIT							LIT										
2	2	248	Door-tomb door left 2				/Data/Global/Objects	TS	OP	HTH		LIT																	
2	2	249	Door-tomb door right 2				/Data/Global/Objects	TU	OP	HTH		LIT																	
2	2	250	Duriel's Lair-Portal to Duriel's Lair				/Data/Global/Objects	SJ	OP	HTH		LIT																	
2	2	251	Dummy-Brazier3				/Data/Global/Objects	B3	OP	HTH		LIT							LIT										
2	2	252	Dummy-Floor brazier				/Data/Global/Objects	FB	ON	HTH		LIT							LIT										
2	2	253	Dummy-flies				/Data/Global/Objects	FL	NU	HTH		LIT																	
2	2	254	ArmorStand-Armor Stand 1R				/Data/Global/Objects	A3	NU	HTH		LIT																	
2	2	255	ArmorStand-Armor Stand 2L				/Data/Global/Objects	A4	NU	HTH		LIT																	
2	2	256	WeaponRack-Weapon Rack 1R				/Data/Global/Objects	W1	NU	HTH		LIT																	
2	2	257	WeaponRack-Weapon Rack 2L				/Data/Global/Objects	W2	NU	HTH		LIT																	
2	2	258	Malus-Malus				/Data/Global/Objects	HM	NU	HTH		LIT																	
2	2	259	Shrine-palace shrine, healthR, harom, arcane Sanctuary				/Data/Global/Objects	P2	OP	HTH		LIT																	
2	2	260	not used-drinker				/Data/Global/Objects	n5	S1	HTH		LIT																	
2	2	261	well-Fountain 1				/Data/Global/Objects	F3	OP	HTH		LIT																	
2	2	262	not used-gesturer				/Data/Global/Objects	n6	S1	HTH		LIT																	
2	2	263	well-Fountain 2, well, desert, tomb				/Data/Global/Objects	F4	OP	HTH		LIT																	
2	2	264	not used-turner				/Data/Global/Objects	n7	S1	HTH		LIT																	
2	2	265	well-Fountain 3				/Data/Global/Objects	F5	OP	HTH		LIT																	
2	2	266	Shrine-snake woman, magic shrine, tomb, arcane sanctuary				/Data/Global/Objects	SN	OP	HTH		LIT							LIT										
2	2	267	Dummy-jungle torch				/Data/Global/Objects	JT	ON	HTH		LIT							LIT										
2	2	268	Well-Fountain 4				/Data/Global/Objects	F6	OP	HTH		LIT																	
2	2	269	Waypoint-waypoint portal				/Data/Global/Objects	wp	ON	HTH		LIT							LIT										
2	2	270	Dummy-healthshrine, act 3, dungeun				/Data/Global/Objects	dj	OP	HTH		LIT																	
2	2	271	jerhyn-placeholder #1				/Data/Global/Objects	ss																					
2	2	272	jerhyn-placeholder #2				/Data/Global/Objects	ss																					
2	2	273	Shrine-innershrinehell2				/Data/Global/Objects	iw	OP	HTH		LIT							LIT										
2	2	274	Shrine-innershrinehell3				/Data/Global/Objects	iv	OP	HTH		LIT																	
2	2	275	hidden stash-ihobject3 inner hell				/Data/Global/Objects	iu	OP	HTH		LIT																	
2	2	276	skull pile-skullpile inner hell				/Data/Global/Objects	is	OP	HTH		LIT																	
2	2	277	hidden stash-ihobject5 inner hell				/Data/Global/Objects	ir	OP	HTH		LIT																	
2	2	278	hidden stash-hobject4 inner hell				/Data/Global/Objects	hg	OP	HTH		LIT																	
2	2	279	Door-secret door 1				/Data/Global/Objects	h2	OP	HTH		LIT																	
2	2	280	Well-pool act 1 wilderness				/Data/Global/Objects	zw	NU	HTH		LIT																	
2	2	281	Dummy-vile dog afterglow				/Data/Global/Objects	9b	OP	HTH		LIT																	
2	2	282	Well-cathedralwell act 1 inside				/Data/Global/Objects	zc	NU	HTH		LIT																	
2	2	283	shrine-shrine1_arcane sanctuary				/Data/Global/Objects	xx																					
2	2	284	shrine-dshrine2 act 2 shrine				/Data/Global/Objects	zs	OP	HTH		LIT							LIT										
2	2	285	shrine-desertshrine3 act 2 shrine				/Data/Global/Objects	zr	OP	HTH		LIT																	
2	2	286	shrine-dshrine1 act 2 shrine				/Data/Global/Objects	zd	OP	HTH		LIT																	
2	2	287	Well-desertwell act 2 well, desert, tomb				/Data/Global/Objects	zl	NU	HTH		LIT																	
2	2	288	Well-cavewell act 1 caves 				/Data/Global/Objects	zy	NU	HTH		LIT																	
2	2	289	chest-chest-r-large act 1				/Data/Global/Objects	q1	OP	HTH		LIT																	
2	2	290	chest-chest-r-tallskinney act 1				/Data/Global/Objects	q2	OP	HTH		LIT																	
2	2	291	chest-chest-r-med act 1				/Data/Global/Objects	q3	OP	HTH		LIT																	
2	2	292	jug-jug1 act 2, desert				/Data/Global/Objects	q4	OP	HTH		LIT																	
2	2	293	jug-jug2 act 2, desert				/Data/Global/Objects	q5	OP	HTH		LIT																	
2	2	294	chest-Lchest1 act 1				/Data/Global/Objects	q6	OP	HTH		LIT																	
2	2	295	Waypoint-waypointi inner hell				/Data/Global/Objects	wi	ON	HTH		LIT							LIT										
2	2	296	chest-dchest2R act 2, desert, tomb, chest-r-med				/Data/Global/Objects	q9	OP	HTH		LIT																	
2	2	297	chest-dchestr act 2, desert, tomb, chest -r large				/Data/Global/Objects	q7	OP	HTH		LIT																	
2	2	298	chest-dchestL act 2, desert, tomb chest l large				/Data/Global/Objects	q8	OP	HTH		LIT																	
2	2	299	taintedsunaltar-tainted sun altar quest				/Data/Global/Objects	za	OP	HTH		LIT							LIT										
2	2	300	shrine-dshrine1 act 2 , desert				/Data/Global/Objects	zv	NU	HTH		LIT							LIT	LIT									
2	2	301	shrine-dshrine4 act 2, desert				/Data/Global/Objects	ze	OP	HTH		LIT							LIT										
2	2	302	orifice-Where you place the Horadric staff				/Data/Global/Objects	HA	NU	HTH		LIT																	
2	2	303	Door-tyrael's door				/Data/Global/Objects	DX	OP	HTH		LIT																	
2	2	304	corpse-guard corpse				/Data/Global/Objects	GC	OP	HTH		LIT																	
2	2	305	hidden stash-rock act 1 wilderness				/Data/Global/Objects	c7	OP	HTH		LIT																	
2	2	306	Waypoint-waypoint act 2				/Data/Global/Objects	wm	ON	HTH		LIT							LIT										
2	2	307	Waypoint-waypoint act 1 wilderness				/Data/Global/Objects	wn	ON	HTH		LIT							LIT										
2	2	308	skeleton-corpse				/Data/Global/Objects	cp	OP	HTH		LIT																	
2	2	309	hidden stash-rockb act 1 wilderness				/Data/Global/Objects	cq	OP	HTH		LIT																	
2	2	310	fire-fire small				/Data/Global/Objects	FX	NU	HTH		LIT																	
2	2	311	fire-fire medium				/Data/Global/Objects	FY	NU	HTH		LIT																	
2	2	312	fire-fire large				/Data/Global/Objects	FZ	NU	HTH		LIT																	
2	2	313	hiding spot-cliff act 1 wilderness				/Data/Global/Objects	cf	NU	HTH		LIT																	
2	2	314	Shrine-mana well1				/Data/Global/Objects	MB	OP	HTH		LIT																	
2	2	315	Shrine-mana well2				/Data/Global/Objects	MD	OP	HTH		LIT																	
2	2	316	Shrine-mana well3, act 2, tomb, 				/Data/Global/Objects	MF	OP	HTH		LIT																	
2	2	317	Shrine-mana well4, act 2, harom				/Data/Global/Objects	MH	OP	HTH		LIT																	
2	2	318	Shrine-mana well5				/Data/Global/Objects	MJ	OP	HTH		LIT																	
2	2	319	hollow log-log				/Data/Global/Objects	cz	NU	HTH		LIT																	
2	2	320	Shrine-jungle healwell act 3				/Data/Global/Objects	JH	OP	HTH		LIT																	
2	2	321	skeleton-corpseb				/Data/Global/Objects	sx	OP	HTH		LIT																	
2	2	322	Shrine-health well, health shrine, desert				/Data/Global/Objects	Mk	OP	HTH		LIT																	
2	2	323	Shrine-mana well7, mana shrine, desert				/Data/Global/Objects	Mi	OP	HTH		LIT																	
2	2	324	loose rock-rockc act 1 wilderness				/Data/Global/Objects	RY	OP	HTH		LIT																	
2	2	325	loose boulder-rockd act 1 wilderness				/Data/Global/Objects	RZ	OP	HTH		LIT																	
2	2	326	chest-chest-L-med				/Data/Global/Objects	c8	OP	HTH		LIT																	
2	2	327	chest-chest-L-large				/Data/Global/Objects	c9	OP	HTH		LIT																	
2	2	328	GuardCorpse-guard on a stick, desert, tomb, harom				/Data/Global/Objects	GS	OP	HTH		LIT																	
2	2	329	bookshelf-bookshelf1				/Data/Global/Objects	b4	OP	HTH		LIT																	
2	2	330	bookshelf-bookshelf2				/Data/Global/Objects	b5	OP	HTH		LIT																	
2	2	331	chest-jungle chest act 3				/Data/Global/Objects	JC	OP	HTH		LIT																	
2	2	332	coffin-tombcoffin				/Data/Global/Objects	tm	OP	HTH		LIT																	
2	2	333	chest-chest-L-med, jungle				/Data/Global/Objects	jz	OP	HTH		LIT																	
2	2	334	Shrine-jungle shrine2				/Data/Global/Objects	jy	OP	HTH		LIT							LIT	LIT									
2	2	335	stash-jungle object act3				/Data/Global/Objects	jx	OP	HTH		LIT																	
2	2	336	stash-jungle object act3				/Data/Global/Objects	jw	OP	HTH		LIT																	
2	2	337	stash-jungle object act3				/Data/Global/Objects	jv	OP	HTH		LIT																	
2	2	338	stash-jungle object act3				/Data/Global/Objects	ju	OP	HTH		LIT																	
2	2	339	Dummy-cain portal				/Data/Global/Objects	tP	OP	HTH	LIT	LIT																	
2	2	340	Shrine-jungle shrine3 act 3				/Data/Global/Objects	js	OP	HTH		LIT							LIT										
2	2	341	Shrine-jungle shrine4 act 3				/Data/Global/Objects	jr	OP	HTH		LIT							LIT										
2	2	342	teleport pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
2	2	343	LamTome-Lam Esen's Tome				/Data/Global/Objects	ab	OP	HTH		LIT																	
2	2	344	stair-stairsl				/Data/Global/Objects	sl																					
2	2	345	stair-stairsr				/Data/Global/Objects	sv																					
2	2	346	a trap-test data floortrap				/Data/Global/Objects	a5	OP	HTH		LIT																	
2	2	347	Shrine-jungleshrine act 3				/Data/Global/Objects	jq	OP	HTH		LIT							LIT										
2	2	348	chest-chest-L-tallskinney, general chest r?				/Data/Global/Objects	c0	OP	HTH		LIT																	
2	2	349	Shrine-mafistoshrine				/Data/Global/Objects	mz	OP	HTH		LIT							LIT										
2	2	350	Shrine-mafistoshrine				/Data/Global/Objects	my	OP	HTH		LIT							LIT										
2	2	351	Shrine-mafistoshrine				/Data/Global/Objects	mx	NU	HTH		LIT							LIT										
2	2	352	Shrine-mafistomana				/Data/Global/Objects	mw	OP	HTH		LIT							LIT										
2	2	353	stash-mafistolair				/Data/Global/Objects	mv	OP	HTH		LIT																	
2	2	354	stash-box				/Data/Global/Objects	mu	OP	HTH		LIT																	
2	2	355	stash-altar				/Data/Global/Objects	mt	OP	HTH		LIT																	
2	2	356	Shrine-mafistohealth				/Data/Global/Objects	mr	OP	HTH		LIT							LIT										
2	2	357	dummy-water rocks in act 3 wrok				/Data/Global/Objects	rw	NU	HTH		LIT																	
2	2	358	Basket-basket 1				/Data/Global/Objects	bd	OP	HTH		LIT																	
2	2	359	Basket-basket 2				/Data/Global/Objects	bj	OP	HTH		LIT																	
2	2	360	Dummy-water logs in act 3  ne logw				/Data/Global/Objects	lw	NU	HTH		LIT																	
2	2	361	Dummy-water rocks girl in act 3 wrob				/Data/Global/Objects	wb	NU	HTH		LIT																	
2	2	362	Dummy-bubbles in act3 water				/Data/Global/Objects	yb	NU	HTH		LIT																	
2	2	363	Dummy-water logs in act 3 logx				/Data/Global/Objects	wd	NU	HTH		LIT																	
2	2	364	Dummy-water rocks in act 3 rokb				/Data/Global/Objects	wc	NU	HTH		LIT																	
2	2	365	Dummy-water rocks girl in act 3 watc				/Data/Global/Objects	we	NU	HTH		LIT																	
2	2	366	Dummy-water rocks in act 3 waty				/Data/Global/Objects	wy	NU	HTH		LIT																	
2	2	367	Dummy-water logs in act 3  logz				/Data/Global/Objects	lx	NU	HTH		LIT																	
2	2	368	Dummy-web covered tree 1				/Data/Global/Objects	w3	NU	HTH		LIT							LIT										
2	2	369	Dummy-web covered tree 2				/Data/Global/Objects	w4	NU	HTH		LIT							LIT										
2	2	370	Dummy-web covered tree 3				/Data/Global/Objects	w5	NU	HTH		LIT							LIT										
2	2	371	Dummy-web covered tree 4				/Data/Global/Objects	w6	NU	HTH		LIT							LIT										
2	2	372	pillar-hobject1				/Data/Global/Objects	70	OP	HTH		LIT																	
2	2	373	cocoon-cacoon				/Data/Global/Objects	CN	OP	HTH		LIT																	
2	2	374	cocoon-cacoon 2				/Data/Global/Objects	CC	OP	HTH		LIT																	
2	2	375	skullpile-hobject1				/Data/Global/Objects	ib	OP	HTH		LIT																	
2	2	376	Shrine-outershrinehell				/Data/Global/Objects	ia	OP	HTH		LIT							LIT										
2	2	377	dummy-water rock girl act 3  nw  blgb				/Data/Global/Objects	QX	NU	HTH		LIT																	
2	2	378	dummy-big log act 3  sw blga				/Data/Global/Objects	qw	NU	HTH		LIT																	
2	2	379	door-slimedoor1				/Data/Global/Objects	SQ	OP	HTH		LIT																	
2	2	380	door-slimedoor2				/Data/Global/Objects	SY	OP	HTH		LIT																	
2	2	381	Shrine-outershrinehell2				/Data/Global/Objects	ht	OP	HTH		LIT							LIT										
2	2	382	Shrine-outershrinehell3				/Data/Global/Objects	hq	OP	HTH		LIT																	
2	2	383	pillar-hobject2				/Data/Global/Objects	hv	OP	HTH		LIT																	
2	2	384	dummy-Big log act 3 se blgc 				/Data/Global/Objects	Qy	NU	HTH		LIT																	
2	2	385	dummy-Big log act 3 nw blgd				/Data/Global/Objects	Qz	NU	HTH		LIT																	
2	2	386	Shrine-health wellforhell				/Data/Global/Objects	ho	OP	HTH		LIT																	
2	2	387	Waypoint-act3waypoint town				/Data/Global/Objects	wz	ON	HTH		LIT							LIT										
2	2	388	Waypoint-waypointh				/Data/Global/Objects	wv	ON	HTH		LIT							LIT										
2	2	389	body-burning town				/Data/Global/Objects	bz	ON	HTH		LIT							LIT										
2	2	390	chest-gchest1L general				/Data/Global/Objects	cy	OP	HTH		LIT																	
2	2	391	chest-gchest2R general				/Data/Global/Objects	cx	OP	HTH		LIT																	
2	2	392	chest-gchest3R general				/Data/Global/Objects	cu	OP	HTH		LIT																	
2	2	393	chest-glchest3L general				/Data/Global/Objects	cd	OP	HTH		LIT																	
2	2	394	ratnest-sewers				/Data/Global/Objects	rn	OP	HTH		LIT																	
2	2	395	body-burning town				/Data/Global/Objects	by	NU	HTH		LIT							LIT										
2	2	396	ratnest-sewers				/Data/Global/Objects	ra	OP	HTH		LIT																	
2	2	397	bed-bed act 1				/Data/Global/Objects	qa	OP	HTH		LIT																	
2	2	398	bed-bed act 1				/Data/Global/Objects	qb	OP	HTH		LIT																	
2	2	399	manashrine-mana wellforhell				/Data/Global/Objects	hn	OP	HTH		LIT							LIT										
2	2	400	a trap-exploding cow  for Tristan and ACT 3 only??Very Rare  1 or 2				/Data/Global/Objects	ew	OP	HTH		LIT																	
2	2	401	gidbinn altar-gidbinn altar				/Data/Global/Objects	ga	ON	HTH		LIT							LIT										
2	2	402	gidbinn-gidbinn decoy				/Data/Global/Objects	gd	ON	HTH		LIT							LIT										
2	2	403	Dummy-diablo right light				/Data/Global/Objects	11	NU	HTH		LIT																	
2	2	404	Dummy-diablo left light				/Data/Global/Objects	12	NU	HTH		LIT																	
2	2	405	Dummy-diablo start point				/Data/Global/Objects	ss																					
2	2	406	Dummy-stool for act 1 cabin				/Data/Global/Objects	s9	NU	HTH		LIT																	
2	2	407	Dummy-wood for act 1 cabin				/Data/Global/Objects	wg	NU	HTH		LIT																	
2	2	408	Dummy-more wood for act 1 cabin				/Data/Global/Objects	wh	NU	HTH		LIT																	
2	2	409	Dummy-skeleton spawn for hell   facing nw				/Data/Global/Objects	QS	OP	HTH		LIT							LIT										
2	2	410	Shrine-holyshrine for monastery,catacombs,jail				/Data/Global/Objects	HL	OP	HTH		LIT							LIT										
2	2	411	a trap-spikes for tombs floortrap				/Data/Global/Objects	A7	OP	HTH		LIT																	
2	2	412	Shrine-act 1 cathedral				/Data/Global/Objects	s0	OP	HTH		LIT							LIT										
2	2	413	Shrine-act 1 jail				/Data/Global/Objects	jb	NU	HTH		LIT							LIT										
2	2	414	Shrine-act 1 jail				/Data/Global/Objects	jd	OP	HTH		LIT							LIT										
2	2	415	Shrine-act 1 jail				/Data/Global/Objects	jf	OP	HTH		LIT							LIT										
2	2	416	goo pile-goo pile for sand maggot lair				/Data/Global/Objects	GP	OP	HTH		LIT																	
2	2	417	bank-bank				/Data/Global/Objects	b6	NU	HTH		LIT																	
2	2	418	wirt's body-wirt's body				/Data/Global/Objects	BP	NU	HTH		LIT																	
2	2	419	dummy-gold placeholder				/Data/Global/Objects	1g																					
2	2	420	corpse-guard corpse 2				/Data/Global/Objects	GF	OP	HTH		LIT																	
2	2	421	corpse-dead villager 1				/Data/Global/Objects	dg	OP	HTH		LIT																	
2	2	422	corpse-dead villager 2				/Data/Global/Objects	df	OP	HTH		LIT																	
2	2	423	Dummy-yet another flame, no damage				/Data/Global/Objects	f8	NU	HTH		LIT																	
2	2	424	hidden stash-tiny pixel shaped thingie				/Data/Global/Objects	f9																					
2	2	425	Shrine-health shrine for caves				/Data/Global/Objects	ce	OP	HTH		LIT																	
2	2	426	Shrine-mana shrine for caves				/Data/Global/Objects	cg	OP	HTH		LIT																	
2	2	427	Shrine-cave magic shrine				/Data/Global/Objects	cg	OP	HTH		LIT																	
2	2	428	Shrine-manashrine, act 3, dungeun				/Data/Global/Objects	de	OP	HTH		LIT																	
2	2	429	Shrine-magic shrine, act 3 sewers.				/Data/Global/Objects	wj	NU	HTH		LIT							LIT	LIT									
2	2	430	Shrine-healthwell, act 3, sewers				/Data/Global/Objects	wk	OP	HTH		LIT																	
2	2	431	Shrine-manawell, act 3, sewers				/Data/Global/Objects	wl	OP	HTH		LIT																	
2	2	432	Shrine-magic shrine, act 3 sewers, dungeon.				/Data/Global/Objects	ws	NU	HTH		LIT							LIT	LIT									
2	2	433	dummy-brazier_celler, act 2				/Data/Global/Objects	bi	NU	HTH		LIT							LIT										
2	2	434	sarcophagus-anubis coffin, act2, tomb				/Data/Global/Objects	qc	OP	HTH		LIT																	
2	2	435	dummy-brazier_general, act 2, sewers, tomb, desert				/Data/Global/Objects	bm	NU	HTH		LIT							LIT										
2	2	436	Dummy-brazier_tall, act 2, desert, town, tombs				/Data/Global/Objects	bo	NU	HTH		LIT							LIT										
2	2	437	Dummy-brazier_small, act 2, desert, town, tombs				/Data/Global/Objects	bq	NU	HTH		LIT							LIT										
2	2	438	Waypoint-waypoint, celler				/Data/Global/Objects	w7	ON	HTH		LIT							LIT										
2	2	439	bed-bed for harum				/Data/Global/Objects	ub	OP	HTH		LIT																	
2	2	440	door-iron grate door left				/Data/Global/Objects	dv	NU	HTH		LIT																	
2	2	441	door-iron grate door right				/Data/Global/Objects	dn	NU	HTH		LIT																	
2	2	442	door-wooden grate door left				/Data/Global/Objects	dp	NU	HTH		LIT																	
2	2	443	door-wooden grate door right				/Data/Global/Objects	dt	NU	HTH		LIT																	
2	2	444	door-wooden door left				/Data/Global/Objects	dk	NU	HTH		LIT																	
2	2	445	door-wooden door right				/Data/Global/Objects	dl	NU	HTH		LIT																	
2	2	446	Dummy-wall torch left for tombs				/Data/Global/Objects	qd	NU	HTH		LIT							LIT										
2	2	447	Dummy-wall torch right for tombs				/Data/Global/Objects	qe	NU	HTH		LIT							LIT										
2	2	448	portal-arcane sanctuary portal				/Data/Global/Objects	ay	ON	HTH		LIT							LIT	LIT									
2	2	449	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hb	OP	HTH		LIT							LIT										
2	2	450	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hc	OP	HTH		LIT							LIT										
2	2	451	Dummy-maggot well health				/Data/Global/Objects	qf	OP	HTH		LIT																	
2	2	452	manashrine-maggot well mana				/Data/Global/Objects	qg	OP	HTH		LIT																	
2	2	453	magic shrine-magic shrine, act 3 arcane sanctuary.				/Data/Global/Objects	hd	OP	HTH		LIT							LIT										
2	2	454	teleportation pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
2	2	455	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
2	2	456	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
2	2	457	Dummy-arcane thing				/Data/Global/Objects	7a	NU	HTH		LIT																	
2	2	458	Dummy-arcane thing				/Data/Global/Objects	7b	NU	HTH		LIT																	
2	2	459	Dummy-arcane thing				/Data/Global/Objects	7c	NU	HTH		LIT																	
2	2	460	Dummy-arcane thing				/Data/Global/Objects	7d	NU	HTH		LIT																	
2	2	461	Dummy-arcane thing				/Data/Global/Objects	7e	NU	HTH		LIT																	
2	2	462	Dummy-arcane thing				/Data/Global/Objects	7f	NU	HTH		LIT																	
2	2	463	Dummy-arcane thing				/Data/Global/Objects	7g	NU	HTH		LIT																	
2	2	464	dead guard-harem guard 1				/Data/Global/Objects	qh	NU	HTH		LIT																	
2	2	465	dead guard-harem guard 2				/Data/Global/Objects	qi	NU	HTH		LIT																	
2	2	466	dead guard-harem guard 3				/Data/Global/Objects	qj	NU	HTH		LIT																	
2	2	467	dead guard-harem guard 4				/Data/Global/Objects	qk	NU	HTH		LIT																	
2	2	468	eunuch-harem blocker				/Data/Global/Objects	ss																					
2	2	469	Dummy-healthwell, act 2, arcane				/Data/Global/Objects	ax	OP	HTH		LIT																	
2	2	470	manashrine-healthwell, act 2, arcane				/Data/Global/Objects	au	OP	HTH		LIT																	
2	2	471	Dummy-test data				/Data/Global/Objects	pp	S1	HTH	LIT	LIT																	
2	2	472	Well-tombwell act 2 well, tomb				/Data/Global/Objects	hu	NU	HTH		LIT																	
2	2	473	Waypoint-waypoint act2 sewer				/Data/Global/Objects	qm	ON	HTH		LIT							LIT										
2	2	474	Waypoint-waypoint act3 travincal				/Data/Global/Objects	ql	ON	HTH		LIT							LIT										
2	2	475	magic shrine-magic shrine, act 3, sewer				/Data/Global/Objects	qn	NU	HTH		LIT							LIT										
2	2	476	dead body-act3, sewer				/Data/Global/Objects	qo	OP	HTH		LIT																	
2	2	477	dummy-torch (act 3 sewer) stra				/Data/Global/Objects	V1	NU	HTH		LIT							LIT										
2	2	478	dummy-torch (act 3 kurast) strb				/Data/Global/Objects	V2	NU	HTH		LIT							LIT										
2	2	479	chest-mafistochestlargeLeft				/Data/Global/Objects	xb	OP	HTH		LIT																	
2	2	480	chest-mafistochestlargeright				/Data/Global/Objects	xc	OP	HTH		LIT																	
2	2	481	chest-mafistochestmedleft				/Data/Global/Objects	xd	OP	HTH		LIT																	
2	2	482	chest-mafistochestmedright				/Data/Global/Objects	xe	OP	HTH		LIT																	
2	2	483	chest-spiderlairchestlargeLeft				/Data/Global/Objects	xf	OP	HTH		LIT																	
2	2	484	chest-spiderlairchesttallLeft				/Data/Global/Objects	xg	OP	HTH		LIT																	
2	2	485	chest-spiderlairchestmedright				/Data/Global/Objects	xh	OP	HTH		LIT																	
2	2	486	chest-spiderlairchesttallright				/Data/Global/Objects	xi	OP	HTH		LIT																	
2	2	487	Steeg Stone-steeg stone				/Data/Global/Objects	y6	NU	HTH		LIT							LIT										
2	2	488	Guild Vault-guild vault				/Data/Global/Objects	y4	NU	HTH		LIT																	
2	2	489	Trophy Case-trophy case				/Data/Global/Objects	y2	NU	HTH		LIT																	
2	2	490	Message Board-message board				/Data/Global/Objects	y3	NU	HTH		LIT																	
2	2	491	Dummy-mephisto bridge				/Data/Global/Objects	xj	OP	HTH		LIT																	
2	2	492	portal-hellgate				/Data/Global/Objects	1y	ON	HTH		LIT								LIT	LIT								
2	2	493	Shrine-manawell, act 3, kurast				/Data/Global/Objects	xl	OP	HTH		LIT																	
2	2	494	Shrine-healthwell, act 3, kurast				/Data/Global/Objects	xm	OP	HTH		LIT																	
2	2	495	Dummy-hellfire1				/Data/Global/Objects	e3	NU	HTH		LIT																	
2	2	496	Dummy-hellfire2				/Data/Global/Objects	e4	NU	HTH		LIT																	
2	2	497	Dummy-hellfire3				/Data/Global/Objects	e5	NU	HTH		LIT																	
2	2	498	Dummy-helllava1				/Data/Global/Objects	e6	NU	HTH		LIT																	
2	2	499	Dummy-helllava2				/Data/Global/Objects	e7	NU	HTH		LIT																	
2	2	500	Dummy-helllava3				/Data/Global/Objects	e8	NU	HTH		LIT																	
2	2	501	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
2	2	502	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
2	2	503	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
2	2	504	chest-horadric cube chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
2	2	505	chest-horadric scroll chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
2	2	506	chest-staff of kings chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
2	2	507	Tome-yet another tome				/Data/Global/Objects	TT	NU	HTH		LIT																	
2	2	508	fire-hell brazier				/Data/Global/Objects	E1	NU	HTH	LIT	LIT																	
2	2	509	fire-hell brazier				/Data/Global/Objects	E2	NU	HTH	LIT	LIT																	
2	2	510	RockPIle-dungeon				/Data/Global/Objects	xn	OP	HTH		LIT																	
2	2	511	magic shrine-magic shrine, act 3,dundeon				/Data/Global/Objects	qo	OP	HTH		LIT																	
2	2	512	basket-dungeon				/Data/Global/Objects	xp	OP	HTH		LIT																	
2	2	513	HungSkeleton-outerhell skeleton				/Data/Global/Objects	jw	OP	HTH		LIT																	
2	2	514	Dummy-guy for dungeon				/Data/Global/Objects	ea	OP	HTH		LIT																	
2	2	515	casket-casket for Act 3 dungeon				/Data/Global/Objects	vb	OP	HTH		LIT																	
2	2	516	sewer stairs-stairs for act 3 sewer quest				/Data/Global/Objects	ve	OP	HTH		LIT																	
2	2	517	sewer lever-lever for act 3 sewer quest				/Data/Global/Objects	vf	OP	HTH		LIT																	
2	2	518	darkwanderer-start position				/Data/Global/Objects	ss																					
2	2	519	dummy-trapped soul placeholder				/Data/Global/Objects	ss																					
2	2	520	Dummy-torch for act3 town				/Data/Global/Objects	VG	NU	HTH		LIT							LIT										
2	2	521	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
2	2	522	BoneChest-innerhellbonepile				/Data/Global/Objects	y1	OP	HTH		LIT																	
2	2	523	Dummy-skeleton spawn for hell facing ne				/Data/Global/Objects	Qt	OP	HTH		LIT							LIT										
2	2	524	Dummy-fog act 3 water rfga				/Data/Global/Objects	ud	NU	HTH		LIT																	
2	2	525	Dummy-Not used				/Data/Global/Objects	xx																					
2	2	526	Hellforge-Forge  hell				/Data/Global/Objects	ux	ON	HTH		LIT							LIT	LIT	LIT								
2	2	527	Guild Portal-Portal to next guild level				/Data/Global/Objects	PP	NU	HTH	LIT	LIT																	
2	2	528	Dummy-hratli start				/Data/Global/Objects	ss																					
2	2	529	Dummy-hratli end				/Data/Global/Objects	ss																					
2	2	530	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	uy	OP	HTH		LIT							LIT										
2	2	531	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	15	OP	HTH		LIT							LIT										
2	2	532	Dummy-natalya start				/Data/Global/Objects	ss																					
2	2	533	TrappedSoul-guy stuck in hell				/Data/Global/Objects	18	OP	HTH		LIT																	
2	2	534	TrappedSoul-guy stuck in hell				/Data/Global/Objects	19	OP	HTH		LIT																	
2	2	535	Dummy-cain start position				/Data/Global/Objects	ss																					
2	2	536	Dummy-stairsr				/Data/Global/Objects	sv	OP	HTH		LIT																	
2	2	537	chest-arcanesanctuarybigchestLeft				/Data/Global/Objects	y7	OP	HTH		LIT																	
2	2	538	casket-arcanesanctuarycasket				/Data/Global/Objects	y8	OP	HTH		LIT																	
2	2	539	chest-arcanesanctuarybigchestRight				/Data/Global/Objects	y9	OP	HTH		LIT																	
2	2	540	chest-arcanesanctuarychestsmallLeft				/Data/Global/Objects	ya	OP	HTH		LIT																	
2	2	541	chest-arcanesanctuarychestsmallRight				/Data/Global/Objects	yc	OP	HTH		LIT																	
2	2	542	Seal-Diablo seal				/Data/Global/Objects	30	ON	HTH		LIT							LIT										
2	2	543	Seal-Diablo seal				/Data/Global/Objects	31	ON	HTH		LIT							LIT										
2	2	544	Seal-Diablo seal				/Data/Global/Objects	32	ON	HTH		LIT							LIT										
2	2	545	Seal-Diablo seal				/Data/Global/Objects	33	ON	HTH		LIT							LIT										
2	2	546	Seal-Diablo seal				/Data/Global/Objects	34	ON	HTH		LIT							LIT										
2	2	547	chest-sparklychest				/Data/Global/Objects	yf	OP	HTH		LIT																	
2	2	548	Waypoint-waypoint pandamonia fortress				/Data/Global/Objects	yg	ON	HTH		LIT							LIT										
2	2	549	fissure-fissure for act 4 inner hell				/Data/Global/Objects	fh	OP	HTH		LIT							LIT										
2	2	550	Dummy-brazier for act 4, hell mesa				/Data/Global/Objects	he	NU	HTH		LIT							LIT										
2	2	551	Dummy-smoke				/Data/Global/Objects	35	NU	HTH		LIT																	
2	2	552	Waypoint-waypoint valleywaypoint				/Data/Global/Objects	yi	ON	HTH		LIT							LIT										
2	2	553	fire-hell brazier				/Data/Global/Objects	9f	NU	HTH		LIT							LIT										
2	2	554	compellingorb-compelling orb				/Data/Global/Objects	55	NU	HTH		LIT							LIT	LIT									
2	2	555	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
2	2	556	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
2	2	557	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
2	2	558	Dummy-fortress brazier #1				/Data/Global/Objects	98	NU	HTH		LIT							LIT										
2	2	559	Dummy-fortress brazier #2				/Data/Global/Objects	99	NU	HTH		LIT							LIT										
2	2	560	Siege Control-To control siege machines				/Data/Global/Objects	zq	OP	HTH		LIT																	
2	2	561	ptox-Pot O Torch (level 1)				/Data/Global/Objects	px	NU	HTH		LIT							LIT	LIT									
2	2	562	pyox-fire pit  (level 1)				/Data/Global/Objects	py	NU	HTH		LIT							LIT										
2	2	563	chestR-expansion no snow				/Data/Global/Objects	6q	OP	HTH		LIT																	
2	2	564	Shrine3wilderness-expansion no snow				/Data/Global/Objects	6r	OP	HTH		LIT							LIT										
2	2	565	Shrine2wilderness-expansion no snow				/Data/Global/Objects	6s	NU	HTH		LIT							LIT										
2	2	566	hiddenstash-expansion no snow				/Data/Global/Objects	3w	OP	HTH		LIT																	
2	2	567	flag wilderness-expansion no snow				/Data/Global/Objects	ym	NU	HTH		LIT																	
2	2	568	barrel wilderness-expansion no snow				/Data/Global/Objects	yn	OP	HTH		LIT																	
2	2	569	barrel wilderness-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
2	2	570	woodchestL-expansion no snow				/Data/Global/Objects	yp	OP	HTH		LIT																	
2	2	571	Shrine3wilderness-expansion no snow				/Data/Global/Objects	yq	NU	HTH		LIT							LIT										
2	2	572	manashrine-expansion no snow				/Data/Global/Objects	yr	OP	HTH		LIT							LIT										
2	2	573	healthshrine-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
2	2	574	burialchestL-expansion no snow				/Data/Global/Objects	yt	OP	HTH		LIT																	
2	2	575	burialchestR-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
2	2	576	well-expansion no snow				/Data/Global/Objects	yv	NU	HTH		LIT																	
2	2	577	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yw	OP	HTH		LIT							LIT	LIT									
2	2	578	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yx	OP	HTH		LIT							LIT										
2	2	579	Waypoint-expansion no snow				/Data/Global/Objects	yy	ON	HTH		LIT							LIT										
2	2	580	ChestL-expansion no snow				/Data/Global/Objects	yz	OP	HTH		LIT																	
2	2	581	woodchestR-expansion no snow				/Data/Global/Objects	6a	OP	HTH		LIT																	
2	2	582	ChestSL-expansion no snow				/Data/Global/Objects	6b	OP	HTH		LIT																	
2	2	583	ChestSR-expansion no snow				/Data/Global/Objects	6c	OP	HTH		LIT																	
2	2	584	etorch1-expansion no snow				/Data/Global/Objects	6d	NU	HTH		LIT							LIT										
2	2	585	ecfra-camp fire				/Data/Global/Objects	2w	NU	HTH		LIT							LIT	LIT									
2	2	586	ettr-town torch				/Data/Global/Objects	2x	NU	HTH		LIT							LIT	LIT									
2	2	587	etorch2-expansion no snow				/Data/Global/Objects	6e	NU	HTH		LIT							LIT										
2	2	588	burningbodies-wilderness/siege				/Data/Global/Objects	6f	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
2	2	589	burningpit-wilderness/siege				/Data/Global/Objects	6g	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
2	2	590	tribal flag-wilderness/siege				/Data/Global/Objects	6h	NU	HTH		LIT																	
2	2	591	eflg-town flag				/Data/Global/Objects	2y	NU	HTH		LIT																	
2	2	592	chan-chandeleir				/Data/Global/Objects	2z	NU	HTH		LIT							LIT										
2	2	593	jar1-wilderness/siege				/Data/Global/Objects	6i	OP	HTH		LIT																	
2	2	594	jar2-wilderness/siege				/Data/Global/Objects	6j	OP	HTH		LIT																	
2	2	595	jar3-wilderness/siege				/Data/Global/Objects	6k	OP	HTH		LIT																	
2	2	596	swingingheads-wilderness				/Data/Global/Objects	6L	NU	HTH		LIT																	
2	2	597	pole-wilderness				/Data/Global/Objects	6m	NU	HTH		LIT																	
2	2	598	animated skulland rockpile-expansion no snow				/Data/Global/Objects	6n	OP	HTH		LIT																	
2	2	599	gate-town main gate				/Data/Global/Objects	2v	OP	HTH		LIT																	
2	2	600	pileofskullsandrocks-seige				/Data/Global/Objects	6o	NU	HTH		LIT																	
2	2	601	hellgate-seige				/Data/Global/Objects	6p	NU	HTH		LIT							LIT	LIT									
2	2	602	banner 1-preset in enemy camp				/Data/Global/Objects	ao	NU	HTH		LIT																	
2	2	603	banner 2-preset in enemy camp				/Data/Global/Objects	ap	NU	HTH		LIT																	
2	2	604	explodingchest-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
2	2	605	chest-specialchest				/Data/Global/Objects	6u	OP	HTH		LIT																	
2	2	606	deathpole-wilderness				/Data/Global/Objects	6v	NU	HTH		LIT																	
2	2	607	Ldeathpole-wilderness				/Data/Global/Objects	6w	NU	HTH		LIT																	
2	2	608	Altar-inside of temple				/Data/Global/Objects	6x	NU	HTH		LIT							LIT										
2	2	609	dummy-Drehya Start In Town				/Data/Global/Objects	ss																					
2	2	610	dummy-Drehya Start Outside Town				/Data/Global/Objects	ss																					
2	2	611	dummy-Nihlathak Start In Town				/Data/Global/Objects	ss																					
2	2	612	dummy-Nihlathak Start Outside Town				/Data/Global/Objects	ss																					
2	2	613	hidden stash-icecave_				/Data/Global/Objects	6y	OP	HTH		LIT																	
2	2	614	healthshrine-icecave_				/Data/Global/Objects	8a	OP	HTH		LIT																	
2	2	615	manashrine-icecave_				/Data/Global/Objects	8b	OP	HTH		LIT																	
2	2	616	evilurn-icecave_				/Data/Global/Objects	8c	OP	HTH		LIT																	
2	2	617	icecavejar1-icecave_				/Data/Global/Objects	8d	OP	HTH		LIT																	
2	2	618	icecavejar2-icecave_				/Data/Global/Objects	8e	OP	HTH		LIT																	
2	2	619	icecavejar3-icecave_				/Data/Global/Objects	8f	OP	HTH		LIT																	
2	2	620	icecavejar4-icecave_				/Data/Global/Objects	8g	OP	HTH		LIT																	
2	2	621	icecavejar4-icecave_				/Data/Global/Objects	8h	OP	HTH		LIT																	
2	2	622	icecaveshrine2-icecave_				/Data/Global/Objects	8i	NU	HTH		LIT							LIT										
2	2	623	cagedwussie1-caged fellow(A5-Prisonner)				/Data/Global/Objects	60	NU	HTH		LIT																	
2	2	624	Ancient Statue 3-statue				/Data/Global/Objects	60	NU	HTH		LIT																	
2	2	625	Ancient Statue 1-statue				/Data/Global/Objects	61	NU	HTH		LIT																	
2	2	626	Ancient Statue 2-statue				/Data/Global/Objects	62	NU	HTH		LIT																	
2	2	627	deadbarbarian-seige/wilderness				/Data/Global/Objects	8j	OP	HTH		LIT																	
2	2	628	clientsmoke-client smoke				/Data/Global/Objects	oz	NU	HTH		LIT																	
2	2	629	icecaveshrine2-icecave_				/Data/Global/Objects	8k	NU	HTH		LIT							LIT										
2	2	630	icecave_torch1-icecave_				/Data/Global/Objects	8L	NU	HTH		LIT							LIT										
2	2	631	icecave_torch2-icecave_				/Data/Global/Objects	8m	NU	HTH		LIT							LIT										
2	2	632	ttor-expansion tiki torch				/Data/Global/Objects	2p	NU	HTH		LIT							LIT										
2	2	633	manashrine-baals				/Data/Global/Objects	8n	OP	HTH		LIT																	
2	2	634	healthshrine-baals				/Data/Global/Objects	8o	OP	HTH		LIT																	
2	2	635	tomb1-baal's lair				/Data/Global/Objects	8p	OP	HTH		LIT																	
2	2	636	tomb2-baal's lair				/Data/Global/Objects	8q	OP	HTH		LIT																	
2	2	637	tomb3-baal's lair				/Data/Global/Objects	8r	OP	HTH		LIT																	
2	2	638	magic shrine-baal's lair				/Data/Global/Objects	8s	NU	HTH		LIT							LIT										
2	2	639	torch1-baal's lair				/Data/Global/Objects	8t	NU	HTH		LIT							LIT										
2	2	640	torch2-baal's lair				/Data/Global/Objects	8u	NU	HTH		LIT							LIT										
2	2	641	manashrine-snowy				/Data/Global/Objects	8v	OP	HTH		LIT							LIT										
2	2	642	healthshrine-snowy				/Data/Global/Objects	8w	OP	HTH		LIT							LIT										
2	2	643	well-snowy				/Data/Global/Objects	8x	NU	HTH		LIT																	
2	2	644	Waypoint-baals_waypoint				/Data/Global/Objects	8y	ON	HTH		LIT							LIT										
2	2	645	magic shrine-snowy_shrine3				/Data/Global/Objects	8z	NU	HTH		LIT							LIT										
2	2	646	Waypoint-wilderness_waypoint				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
2	2	647	magic shrine-snowy_shrine3				/Data/Global/Objects	5b	OP	HTH		LIT							LIT	LIT									
2	2	648	well-baalslair				/Data/Global/Objects	5c	NU	HTH		LIT																	
2	2	649	magic shrine2-baal's lair				/Data/Global/Objects	5d	NU	HTH		LIT							LIT										
2	2	650	object1-snowy				/Data/Global/Objects	5e	OP	HTH		LIT																	
2	2	651	woodchestL-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
2	2	652	woodchestR-snowy				/Data/Global/Objects	5g	OP	HTH		LIT																	
2	2	653	magic shrine-baals_shrine3				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
2	2	654	woodchest2L-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
2	2	655	woodchest2R-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
2	2	656	swingingheads-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
2	2	657	debris-snowy				/Data/Global/Objects	5l	NU	HTH		LIT																	
2	2	658	pene-Pen breakable door				/Data/Global/Objects	2q	NU	HTH		LIT																	
2	2	659	magic shrine-temple				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
2	2	660	mrpole-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
2	2	661	Waypoint-icecave 				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
2	2	662	magic shrine-temple				/Data/Global/Objects	5t	NU	HTH		LIT							LIT										
2	2	663	well-temple				/Data/Global/Objects	5q	NU	HTH		LIT																	
2	2	664	torch1-temple				/Data/Global/Objects	5r	NU	HTH		LIT							LIT										
2	2	665	torch1-temple				/Data/Global/Objects	5s	NU	HTH		LIT							LIT										
2	2	666	object1-temple				/Data/Global/Objects	5u	OP	HTH		LIT																	
2	2	667	object2-temple				/Data/Global/Objects	5v	OP	HTH		LIT																	
2	2	668	mrbox-baals				/Data/Global/Objects	5w	OP	HTH		LIT																	
2	2	669	well-icecave				/Data/Global/Objects	5x	NU	HTH		LIT																	
2	2	670	magic shrine-temple				/Data/Global/Objects	5y	NU	HTH		LIT							LIT										
2	2	671	healthshrine-temple				/Data/Global/Objects	5z	OP	HTH		LIT																	
2	2	672	manashrine-temple				/Data/Global/Objects	3a	OP	HTH		LIT																	
2	2	673	red light- (touch me)  for blacksmith				/Data/Global/Objects	ss																					
2	2	674	tomb1L-baal's lair				/Data/Global/Objects	3b	OP	HTH		LIT																	
2	2	675	tomb2L-baal's lair				/Data/Global/Objects	3c	OP	HTH		LIT																	
2	2	676	tomb3L-baal's lair				/Data/Global/Objects	3d	OP	HTH		LIT																	
2	2	677	ubub-Ice cave bubbles 01				/Data/Global/Objects	2u	NU	HTH		LIT																	
2	2	678	sbub-Ice cave bubbles 01				/Data/Global/Objects	2s	NU	HTH		LIT																	
2	2	679	tomb1-redbaal's lair				/Data/Global/Objects	3f	OP	HTH		LIT																	
2	2	680	tomb1L-redbaal's lair				/Data/Global/Objects	3g	OP	HTH		LIT																	
2	2	681	tomb2-redbaal's lair				/Data/Global/Objects	3h	OP	HTH		LIT																	
2	2	682	tomb2L-redbaal's lair				/Data/Global/Objects	3i	OP	HTH		LIT																	
2	2	683	tomb3-redbaal's lair				/Data/Global/Objects	3j	OP	HTH		LIT																	
2	2	684	tomb3L-redbaal's lair				/Data/Global/Objects	3k	OP	HTH		LIT																	
2	2	685	mrbox-redbaals				/Data/Global/Objects	3L	OP	HTH		LIT																	
2	2	686	torch1-redbaal's lair				/Data/Global/Objects	3m	NU	HTH		LIT							LIT										
2	2	687	torch2-redbaal's lair				/Data/Global/Objects	3n	NU	HTH		LIT							LIT										
2	2	688	candles-temple				/Data/Global/Objects	3o	NU	HTH		LIT							LIT										
2	2	689	Waypoint-temple				/Data/Global/Objects	3p	ON	HTH		LIT							LIT										
2	2	690	deadperson-everywhere				/Data/Global/Objects	3q	NU	HTH		LIT																	
2	2	691	groundtomb-temple				/Data/Global/Objects	3s	OP	HTH		LIT																	
2	2	692	Dummy-Larzuk Greeting				/Data/Global/Objects	ss																					
2	2	693	Dummy-Larzuk Standard				/Data/Global/Objects	ss																					
2	2	694	groundtombL-temple				/Data/Global/Objects	3t	OP	HTH		LIT																	
2	2	695	deadperson2-everywhere				/Data/Global/Objects	3u	OP	HTH		LIT																	
2	2	696	ancientsaltar-ancientsaltar				/Data/Global/Objects	4a	OP	HTH		LIT							LIT										
2	2	697	To The Worldstone Keep Level 1-ancientsdoor				/Data/Global/Objects	4b	OP	HTH		LIT																	
2	2	698	eweaponrackR-everywhere				/Data/Global/Objects	3x	NU	HTH		LIT																	
2	2	699	eweaponrackL-everywhere				/Data/Global/Objects	3y	NU	HTH		LIT																	
2	2	700	earmorstandR-everywhere				/Data/Global/Objects	3z	NU	HTH		LIT																	
2	2	701	earmorstandL-everywhere				/Data/Global/Objects	4c	NU	HTH		LIT																	
2	2	702	torch2-summit				/Data/Global/Objects	9g	NU	HTH		LIT							LIT										
2	2	703	funeralpire-outside				/Data/Global/Objects	9h	NU	HTH		LIT							LIT										
2	2	704	burninglogs-outside				/Data/Global/Objects	9i	NU	HTH		LIT							LIT										
2	2	705	stma-Ice cave steam				/Data/Global/Objects	2o	NU	HTH		LIT																	
2	2	706	deadperson2-everywhere				/Data/Global/Objects	3v	OP	HTH		LIT																	
2	2	707	Dummy-Baal's lair				/Data/Global/Objects	ss																					
2	2	708	fana-frozen anya				/Data/Global/Objects	2n	NU	HTH		LIT																	
2	2	709	BBQB-BBQ Bunny				/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									
2	2	710	btor-Baal Torch Big				/Data/Global/Objects	25	NU	HTH		LIT							LIT										
2	2	711	Dummy-invisible ancient				/Data/Global/Objects	ss																					
2	2	712	Dummy-invisible base				/Data/Global/Objects	ss																					
2	2	713	The Worldstone Chamber-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
2	2	714	Glacial Caves Level 1-summit door				/Data/Global/Objects	4u	OP	HTH		LIT																	
2	2	715	strlastcinematic-last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
2	2	716	Harrogath-last last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
2	2	717	Zoo-test data				/Data/Global/Objects	ss																					
2	2	718	Keeper-test data				/Data/Global/Objects	7z	NU	HTH		LIT																	
2	2	719	Throne of Destruction-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
2	2	720	Dummy-fire place guy				/Data/Global/Objects	7y	NU	HTH		LIT																	
2	2	721	Dummy-door blocker				/Data/Global/Objects	ss																					
2	2	722	Dummy-door blocker				/Data/Global/Objects	ss																					
3	1	0	cain3-ACT 3 TABLE				/Data/Global/Monsters	2D	NU	HTH		LIT																	0
3	1	1	place_champion-ACT 3 TABLE																										0
3	1	2	act3male-ACT 3 TABLE				/Data/Global/Monsters	N4	NU	HTH	BRD	HVY	MED	MED	MED	BAN	BUK		HBD										0
3	1	3	act3female-ACT 3 TABLE				/Data/Global/Monsters	N3	NU	HTH	LIT	BTP	DLN			BSK	BSK												0
3	1	4	asheara-ACT 3 TABLE				/Data/Global/Monsters	AH	NU	HTH		LIT																	0
3	1	5	hratli-ACT 3 TABLE				/Data/Global/Monsters	HR	NU	HTH		LIT																	0
3	1	6	alkor-ACT 3 TABLE				/Data/Global/Monsters	AL	NU	HTH		LIT																	0
3	1	7	ormus-ACT 3 TABLE				/Data/Global/Monsters	OR	NU	HTH		LIT																	0
3	1	8	meshif2-ACT 3 TABLE				/Data/Global/Monsters	M3	NU	HTH		LIT																	0
3	1	9	place_amphibian-ACT 3 TABLE																										0
3	1	10	place_tentacle_ns-ACT 3 TABLE			7	/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
3	1	11	place_tentacle_ew-ACT 3 TABLE			5	/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
3	1	12	place_fetishnest-ACT 3 TABLE																										0
3	1	13	trap-horzmissile-ACT 3 TABLE																										0
3	1	14	trap-vertmissile-ACT 3 TABLE																										0
3	1	15	natalya-ACT 3 TABLE				/Data/Global/Monsters	TZ	NU	HTH		LIT																	0
3	1	16	place_mosquitonest-ACT 3 TABLE																										0
3	1	17	place_group25-ACT 3 TABLE																										0
3	1	18	place_group50-ACT 3 TABLE																										0
3	1	19	place_group75-ACT 3 TABLE																										0
3	1	20	place_group100-ACT 3 TABLE																										0
3	1	21	compellingorb-ACT 3 TABLE				/Data/Global/Objects	55	NU	HTH		LIT							LIT	LIT									0
3	1	22	mephisto-ACT 3 TABLE				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
3	1	23	trap-melee-ACT 3 TABLE				/Data/Global/Monsters	M4	A1	HTH		LIT																	0
3	1	24	mephistospirit-ACT 3 TABLE			2	/Data/Global/Monsters	M6	A1	HTH		LIT																	0
3	1	25	act3hire-ACT 3 TABLE				/Data/Global/Monsters	IW	NU	1HS	LIT	LIT				LSD		KIT											0
3	1	26	place_fetish-ACT 3 TABLE				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	27	place_fetishshaman-ACT 3 TABLE				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	28	Web Mage the Burning-ACT 3 TABLE																										0
3	1	29	Witch Doctor Endugu-ACT 3 TABLE																										0
3	1	30	Stormtree-ACT 3 TABLE																										0
3	1	31	Sarina the Battlemaid-ACT 3 TABLE																										0
3	1	32	Icehawk Riftwing-ACT 3 TABLE																										0
3	1	33	Ismail Vilehand-ACT 3 TABLE																										0
3	1	34	Geleb Flamefinger-ACT 3 TABLE																										0
3	1	35	Bremm Sparkfist-ACT 3 TABLE																										0
3	1	36	Toorc Icefist-ACT 3 TABLE																										0
3	1	37	Wyand Voidfinger-ACT 3 TABLE																										0
3	1	38	Maffer Dragonhand-ACT 3 TABLE																										0
3	1	39	skeleton1-Skeleton-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	40	skeleton2-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	41	skeleton3-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	42	skeleton4-BurningDead-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	43	skeleton5-Horror-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	44	zombie1-Zombie-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	45	zombie2-HungryDead-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	46	zombie3-Ghoul-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	47	zombie4-DrownedCarcass-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	48	zombie5-PlagueBearer-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	49	bighead1-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	50	bighead2-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	51	bighead3-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	52	bighead4-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	53	bighead5-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	54	foulcrow1-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	55	foulcrow2-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	56	foulcrow3-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	57	foulcrow4-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	58	fallen1-Fallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
3	1	59	fallen2-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
3	1	60	fallen3-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
3	1	61	fallen4-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
3	1	62	fallen5-WarpedFallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
3	1	63	brute2-Brute-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
3	1	64	brute3-Yeti-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
3	1	65	brute4-Crusher-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
3	1	66	brute5-WailingBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
3	1	67	brute1-GargantuanBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
3	1	68	sandraider1-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	69	sandraider2-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	70	sandraider3-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	71	sandraider4-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	72	sandraider5-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	73	gorgon1-unused-Idle				/Data/Global/Monsters	GO																					0
3	1	74	gorgon2-unused-Idle				/Data/Global/Monsters	GO																					0
3	1	75	gorgon3-unused-Idle				/Data/Global/Monsters	GO																					0
3	1	76	gorgon4-unused-Idle				/Data/Global/Monsters	GO																					0
3	1	77	wraith1-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	78	wraith2-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	79	wraith3-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	80	wraith4-Apparition-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	81	wraith5-DarkShape-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	82	corruptrogue1-DarkHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	83	corruptrogue2-VileHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	84	corruptrogue3-DarkStalker-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	85	corruptrogue4-BlackRogue-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	86	corruptrogue5-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	87	baboon1-DuneBeast-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	88	baboon2-RockDweller-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	89	baboon3-JungleHunter-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	90	baboon4-DoomApe-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	91	baboon5-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	92	goatman1-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	93	goatman2-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	94	goatman3-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	95	goatman4-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	96	goatman5-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	97	fallenshaman1-FallenShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	98	fallenshaman2-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	99	fallenshaman3-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	100	fallenshaman4-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	101	fallenshaman5-WarpedShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	102	quillrat1-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	103	quillrat2-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	104	quillrat3-ThornBeast-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	105	quillrat4-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	106	quillrat5-JungleUrchin-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	107	sandmaggot1-SandMaggot-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	108	sandmaggot2-RockWorm-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	109	sandmaggot3-Devourer-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	110	sandmaggot4-GiantLamprey-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	111	sandmaggot5-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	112	clawviper1-TombViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	113	clawviper2-ClawViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	114	clawviper3-Salamander-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	115	clawviper4-PitViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	116	clawviper5-SerpentMagus-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	117	sandleaper1-SandLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	118	sandleaper2-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	119	sandleaper3-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	120	sandleaper4-TreeLurker-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	121	sandleaper5-RazorPitDemon-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	122	pantherwoman1-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	123	pantherwoman2-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	124	pantherwoman3-NightTiger-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	125	pantherwoman4-HellCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	126	swarm1-Itchies-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
3	1	127	swarm2-BlackLocusts-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
3	1	128	swarm3-PlagueBugs-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
3	1	129	swarm4-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
3	1	130	scarab1-DungSoldier-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	131	scarab2-SandWarrior-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	132	scarab3-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	133	scarab4-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	134	scarab5-AlbinoRoach-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	135	mummy1-DriedCorpse-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	136	mummy2-Decayed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	137	mummy3-Embalmed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	138	mummy4-PreservedDead-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	139	mummy5-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	140	unraveler1-HollowOne-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	141	unraveler2-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	142	unraveler3-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	143	unraveler4-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	144	unraveler5-Baal Subject Mummy-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	145	chaoshorde1-unused-Idle				/Data/Global/Monsters	CH																					0
3	1	146	chaoshorde2-unused-Idle				/Data/Global/Monsters	CH																					0
3	1	147	chaoshorde3-unused-Idle				/Data/Global/Monsters	CH																					0
3	1	148	chaoshorde4-unused-Idle				/Data/Global/Monsters	CH																					0
3	1	149	vulture1-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
3	1	150	vulture2-UndeadScavenger-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
3	1	151	vulture3-HellBuzzard-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
3	1	152	vulture4-WingedNightmare-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
3	1	153	mosquito1-Sucker-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
3	1	154	mosquito2-Feeder-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
3	1	155	mosquito3-BloodHook-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
3	1	156	mosquito4-BloodWing-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
3	1	157	willowisp1-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	158	willowisp2-SwampGhost-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	159	willowisp3-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	160	willowisp4-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	161	arach1-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	162	arach2-SandFisher-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	163	arach3-PoisonSpinner-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	164	arach4-FlameSpider-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	165	arach5-SpiderMagus-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	166	thornhulk1-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
3	1	167	thornhulk2-BrambleHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
3	1	168	thornhulk3-Thrasher-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
3	1	169	thornhulk4-Spikefist-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
3	1	170	vampire1-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	171	vampire2-NightLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	172	vampire3-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	173	vampire4-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	174	vampire5-Banished-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	175	batdemon1-DesertWing-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	176	batdemon2-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	177	batdemon3-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	178	batdemon4-BloodDiver-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	179	batdemon5-DarkFamiliar-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	180	fetish1-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	181	fetish2-Fetish-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	182	fetish3-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	183	fetish4-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	184	fetish5-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	185	cain1-DeckardCain-NpcOutOfTown				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
3	1	186	gheed-Gheed-Npc				/Data/Global/Monsters	GH	NU	HTH		LIT																	0
3	1	187	akara-Akara-Npc				/Data/Global/Monsters	PS	NU	HTH		LIT																	0
3	1	188	chicken-dummy-Idle				/Data/Global/Monsters	CK	NU	HTH		LIT																	0
3	1	189	kashya-Kashya-Npc				/Data/Global/Monsters	RC	NU	HTH		LIT																	0
3	1	190	rat-dummy-Idle				/Data/Global/Monsters	RT	NU	HTH		LIT																	0
3	1	191	rogue1-Dummy-Idle				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
3	1	192	hellmeteor-Dummy-HellMeteor				/Data/Global/Monsters	K9																					0
3	1	193	charsi-Charsi-Npc				/Data/Global/Monsters	CI	NU	HTH		LIT																	0
3	1	194	warriv1-Warriv-Npc				/Data/Global/Monsters	WA	NU	HTH		LIT																	0
3	1	195	andariel-Andariel-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
3	1	196	bird1-dummy-Idle				/Data/Global/Monsters	BS	WL	HTH		LIT																	0
3	1	197	bird2-dummy-Idle				/Data/Global/Monsters	BL																					0
3	1	198	bat-dummy-Idle				/Data/Global/Monsters	B9	WL	HTH		LIT																	0
3	1	199	cr_archer1-DarkRanger-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	200	cr_archer2-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	201	cr_archer3-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	202	cr_archer4-BlackArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	203	cr_archer5-FleshArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	204	cr_lancer1-DarkSpearwoman-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	205	cr_lancer2-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	206	cr_lancer3-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	207	cr_lancer4-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	208	cr_lancer5-FleshLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	209	sk_archer1-SkeletonArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	210	sk_archer2-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	211	sk_archer3-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	212	sk_archer4-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	213	sk_archer5-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	214	warriv2-Warriv-Npc				/Data/Global/Monsters	WX	NU	HTH		LIT																	0
3	1	215	atma-Atma-Npc				/Data/Global/Monsters	AS	NU	HTH		LIT																	0
3	1	216	drognan-Drognan-Npc				/Data/Global/Monsters	DR	NU	HTH		LIT																	0
3	1	217	fara-Fara-Npc				/Data/Global/Monsters	OF	NU	HTH		LIT																	0
3	1	218	cow-dummy-Idle				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
3	1	219	maggotbaby1-SandMaggotYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
3	1	220	maggotbaby2-RockWormYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
3	1	221	maggotbaby3-DevourerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
3	1	222	maggotbaby4-GiantLampreyYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
3	1	223	maggotbaby5-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
3	1	224	camel-dummy-Idle				/Data/Global/Monsters	CM	NU	HTH		LIT																	0
3	1	225	blunderbore1-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
3	1	226	blunderbore2-Gorbelly-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
3	1	227	blunderbore3-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
3	1	228	blunderbore4-Urdar-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
3	1	229	maggotegg1-SandMaggotEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
3	1	230	maggotegg2-RockWormEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
3	1	231	maggotegg3-DevourerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
3	1	232	maggotegg4-GiantLampreyEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
3	1	233	maggotegg5-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
3	1	234	act2male-dummy-Towner				/Data/Global/Monsters	2M	NU	HTH	OLD	MED	MED						TUR										0
3	1	235	act2female-Dummy-Towner				/Data/Global/Monsters	2F	NU	HTH	LIT	LIT	LIT																0
3	1	236	act2child-dummy-Towner				/Data/Global/Monsters	2C																					0
3	1	237	greiz-Greiz-Npc				/Data/Global/Monsters	GR	NU	HTH		LIT																	0
3	1	238	elzix-Elzix-Npc				/Data/Global/Monsters	EL	NU	HTH		LIT																	0
3	1	239	geglash-Geglash-Npc				/Data/Global/Monsters	GE	NU	HTH		LIT																	0
3	1	240	jerhyn-Jerhyn-Npc				/Data/Global/Monsters	JE	NU	HTH		LIT																	0
3	1	241	lysander-Lysander-Npc				/Data/Global/Monsters	LY	NU	HTH		LIT																	0
3	1	242	act2guard1-Dummy-Towner				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
3	1	243	act2vendor1-dummy-Vendor				/Data/Global/Monsters	M1	NU	HTH		LIT																	0
3	1	244	act2vendor2-dummy-Vendor				/Data/Global/Monsters	M2	NU	HTH		LIT																	0
3	1	245	crownest1-FoulCrowNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
3	1	246	crownest2-BloodHawkNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
3	1	247	crownest3-BlackVultureNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
3	1	248	crownest4-CloudStalkerNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
3	1	249	meshif1-Meshif-Npc				/Data/Global/Monsters	MS	NU	HTH		LIT																	0
3	1	250	duriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
3	1	251	bonefetish1-Undead RatMan-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	252	bonefetish2-Undead Fetish-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	253	bonefetish3-Undead Flayer-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	254	bonefetish4-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	255	bonefetish5-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	256	darkguard1-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	257	darkguard2-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	258	darkguard3-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	259	darkguard4-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	260	darkguard5-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	261	bloodmage1-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	262	bloodmage2-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	263	bloodmage3-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	264	bloodmage4-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	265	bloodmage5-unused-Idle				/Data/Global/Monsters	xx																					0
3	1	266	maggot-Maggot-Idle				/Data/Global/Monsters	MA	NU	HTH		LIT																	0
3	1	267	sarcophagus-MummyGenerator-Sarcophagus				/Data/Global/Monsters	MG	NU	HTH		LIT																	0
3	1	268	radament-Radament-GreaterMummy				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
3	1	269	firebeast-unused-ElementalBeast				/Data/Global/Monsters	FM	NU	HTH		LIT																	0
3	1	270	iceglobe-unused-ElementalBeast				/Data/Global/Monsters	IM	NU	HTH		LIT																	0
3	1	271	lightningbeast-unused-ElementalBeast				/Data/Global/Monsters	xx																					0
3	1	272	poisonorb-unused-ElementalBeast				/Data/Global/Monsters	PM	NU	HTH		LIT																	0
3	1	273	flyingscimitar-FlyingScimitar-FlyingScimitar				/Data/Global/Monsters	ST	NU	HTH		LIT																	0
3	1	274	zealot1-Zakarumite-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
3	1	275	zealot2-Faithful-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
3	1	276	zealot3-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
3	1	277	cantor1-Sexton-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
3	1	278	cantor2-Cantor-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
3	1	279	cantor3-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
3	1	280	cantor4-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
3	1	281	mephisto-Mephisto-Mephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
3	1	282	diablo-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
3	1	283	cain2-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
3	1	284	cain3-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
3	1	285	cain4-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
3	1	286	frogdemon1-Swamp Dweller-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
3	1	287	frogdemon2-Bog Creature-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
3	1	288	frogdemon3-Slime Prince-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
3	1	289	summoner-Summoner-Summoner				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
3	1	290	tyrael1-tyrael-NpcStationary				/Data/Global/Monsters	TX	NU	HTH		LIT		LIT	LIT														0
3	1	291	asheara-asheara-Npc				/Data/Global/Monsters	AH	NU	HTH		LIT																	0
3	1	292	hratli-hratli-Npc				/Data/Global/Monsters	HR	NU	HTH		LIT																	0
3	1	293	alkor-alkor-Npc				/Data/Global/Monsters	AL	NU	HTH		LIT																	0
3	1	294	ormus-ormus-Npc				/Data/Global/Monsters	OR	NU	HTH		LIT																	0
3	1	295	izual-izual-Izual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
3	1	296	halbu-halbu-Npc				/Data/Global/Monsters	20	NU	HTH		LIT																	0
3	1	297	tentacle1-WaterWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
3	1	298	tentacle2-RiverStalkerLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
3	1	299	tentacle3-StygianWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
3	1	300	tentaclehead1-WaterWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
3	1	301	tentaclehead2-RiverStalkerHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
3	1	302	tentaclehead3-StygianWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
3	1	303	meshif2-meshif-Npc				/Data/Global/Monsters	M3	NU	HTH		LIT																	0
3	1	304	cain5-DeckardCain-Npc				/Data/Global/Monsters	1D	NU	HTH		LIT																	0
3	1	305	navi-navi-Navi				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
3	1	306	bloodraven-Bloodraven-BloodRaven				/Data/Global/Monsters	CR	NU	BOW	BRV	HVY	BRV	HVY	HVY	LIT	LBB		HVY	HVY									0
3	1	307	bug-Dummy-Idle				/Data/Global/Monsters	BG	NU	HTH		LIT																	0
3	1	308	scorpion-Dummy-Idle				/Data/Global/Monsters	DS	NU	HTH		LIT																	0
3	1	309	rogue2-RogueScout-GoodNpcRanged				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
3	1	310	roguehire-Dummy-Hireable				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
3	1	311	rogue3-Dummy-TownRogue				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
3	1	312	gargoyletrap-GargoyleTrap-GargoyleTrap				/Data/Global/Monsters	GT	NU	HTH		LIT																	0
3	1	313	skmage_pois1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	314	skmage_pois2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	315	skmage_pois3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	316	skmage_pois4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	317	fetishshaman1-RatManShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	318	fetishshaman2-FetishShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	319	fetishshaman3-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	320	fetishshaman4-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	321	fetishshaman5-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	322	larva-larva-Idle				/Data/Global/Monsters	LV	NU	HTH		LIT																	0
3	1	323	maggotqueen1-SandMaggotQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
3	1	324	maggotqueen2-RockWormQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
3	1	325	maggotqueen3-DevourerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
3	1	326	maggotqueen4-GiantLampreyQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
3	1	327	maggotqueen5-WorldKillerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
3	1	328	claygolem-ClayGolem-NecroPet				/Data/Global/Monsters	G1	NU	HTH		LIT																	0
3	1	329	bloodgolem-BloodGolem-NecroPet				/Data/Global/Monsters	G2	NU	HTH		LIT																	0
3	1	330	irongolem-IronGolem-NecroPet				/Data/Global/Monsters	G4	NU	HTH		LIT																	0
3	1	331	firegolem-FireGolem-NecroPet				/Data/Global/Monsters	G3	NU	HTH		LIT																	0
3	1	332	familiar-Dummy-Idle				/Data/Global/Monsters	FI	NU	HTH		LIT																	0
3	1	333	act3male-Dummy-Towner				/Data/Global/Monsters	N4	NU	HTH	BRD	HVY	HVY	HEV	HEV	FSH	SAK		TKT										0
3	1	334	baboon6-NightMarauder-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	335	act3female-Dummy-Towner				/Data/Global/Monsters	N3	NU	HTH	LIT	MTP	SRT			BSK	BSK												0
3	1	336	natalya-Natalya-Npc				/Data/Global/Monsters	TZ	NU	HTH		LIT																	0
3	1	337	vilemother1-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
3	1	338	vilemother2-StygianHag-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
3	1	339	vilemother3-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
3	1	340	vilechild1-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
3	1	341	vilechild2-StygianDog-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
3	1	342	vilechild3-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
3	1	343	fingermage1-Groper-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
3	1	344	fingermage2-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
3	1	345	fingermage3-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
3	1	346	regurgitator1-Corpulent-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
3	1	347	regurgitator2-CorpseSpitter-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
3	1	348	regurgitator3-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
3	1	349	doomknight1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	350	doomknight2-AbyssKnight-AbyssKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	351	doomknight3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	352	quillbear1-QuillBear-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
3	1	353	quillbear2-SpikeGiant-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
3	1	354	quillbear3-ThornBrute-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
3	1	355	quillbear4-RazorBeast-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
3	1	356	quillbear5-GiantUrchin-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
3	1	357	snake-Dummy-Idle				/Data/Global/Monsters	CO	NU	HTH		LIT																	0
3	1	358	parrot-Dummy-Idle				/Data/Global/Monsters	PR	WL	HTH		LIT																	0
3	1	359	fish-Dummy-Idle				/Data/Global/Monsters	FJ																					0
3	1	360	evilhole1-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
3	1	361	evilhole2-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
3	1	362	evilhole3-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
3	1	363	evilhole4-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
3	1	364	evilhole5-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
3	1	365	trap-firebolt-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
3	1	366	trap-horzmissile-a trap-Trap-RightArrow				/Data/Global/Monsters	9A																					0
3	1	367	trap-vertmissile-a trap-Trap-LeftArrow				/Data/Global/Monsters	9A																					0
3	1	368	trap-poisoncloud-a trap-Trap-Poison				/Data/Global/Monsters	9A																					0
3	1	369	trap-lightning-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
3	1	370	act2guard2-Kaelan-JarJar				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
3	1	371	invisospawner-Dummy-InvisoSpawner				/Data/Global/Monsters	K9																					0
3	1	372	diabloclone-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
3	1	373	suckernest1-SuckerNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
3	1	374	suckernest2-FeederNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
3	1	375	suckernest3-BloodHookNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
3	1	376	suckernest4-BloodWingNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
3	1	377	act2hire-Guard-Hireable				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
3	1	378	minispider-Dummy-Idle				/Data/Global/Monsters	LS	NU	HTH		LIT																	0
3	1	379	boneprison1--Idle				/Data/Global/Monsters	67	NU	HTH		LIT																	0
3	1	380	boneprison2--Idle				/Data/Global/Monsters	66	NU	HTH		LIT																	0
3	1	381	boneprison3--Idle				/Data/Global/Monsters	69	NU	HTH		LIT																	0
3	1	382	boneprison4--Idle				/Data/Global/Monsters	68	NU	HTH		LIT																	0
3	1	383	bonewall-Dummy-BoneWall				/Data/Global/Monsters	BW	NU	HTH		LIT																	0
3	1	384	councilmember1-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	385	councilmember2-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	386	councilmember3-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	387	turret1-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
3	1	388	turret2-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
3	1	389	turret3-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
3	1	390	hydra1-Hydra-Hydra				/Data/Global/Monsters	HX	NU	HTH		LIT							LIT										0
3	1	391	hydra2-Hydra-Hydra				/Data/Global/Monsters	21	NU	HTH		LIT							LIT										0
3	1	392	hydra3-Hydra-Hydra				/Data/Global/Monsters	HZ	NU	HTH		LIT							LIT										0
3	1	393	trap-melee-a trap-Trap-Melee				/Data/Global/Monsters	M4	A1	HTH		LIT																	0
3	1	394	seventombs-Dummy-7TIllusion				/Data/Global/Monsters	9A																					0
3	1	395	dopplezon-Dopplezon-Idle				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
3	1	396	valkyrie-Valkyrie-NecroPet				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
3	1	397	act2guard3-Dummy-Idle				/Data/Global/Monsters	SK																					0
3	1	398	act3hire-Iron Wolf-Hireable				/Data/Global/Monsters	IW	NU	1HS	LIT	LIT				WND		KIT											0
3	1	399	megademon1-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	400	megademon2-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	401	megademon3-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	402	necroskeleton-NecroSkeleton-NecroPet				/Data/Global/Monsters	SK	NU	1HS	DES	HVY	DES	DES	DES	SCM		KIT	DES	DES	LIT								0
3	1	403	necromage-NecroMage-NecroPet				/Data/Global/Monsters	SK	NU	HTH	DES	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
3	1	404	griswold-Griswold-Griswold				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
3	1	405	compellingorb-compellingorb-Idle				/Data/Global/Monsters	9a																					0
3	1	406	tyrael2-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
3	1	407	darkwanderer-youngdiablo-DarkWanderer				/Data/Global/Monsters	1Z	NU	HTH		LIT																	0
3	1	408	trap-nova-a trap-Trap-Nova				/Data/Global/Monsters	9A																					0
3	1	409	spiritmummy-Dummy-Idle				/Data/Global/Monsters	xx																					0
3	1	410	lightningspire-LightningSpire-ArcaneTower				/Data/Global/Monsters	AE	NU	HTH		LIT							LIT										0
3	1	411	firetower-FireTower-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
3	1	412	slinger1-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
3	1	413	slinger2-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
3	1	414	slinger3-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
3	1	415	slinger4-HellSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
3	1	416	act2guard4-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
3	1	417	act2guard5-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
3	1	418	skmage_cold1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
3	1	419	skmage_cold2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
3	1	420	skmage_cold3-BaalColdMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
3	1	421	skmage_cold4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
3	1	422	skmage_fire1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
3	1	423	skmage_fire2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
3	1	424	skmage_fire3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
3	1	425	skmage_fire4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
3	1	426	skmage_ltng1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
3	1	427	skmage_ltng2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
3	1	428	skmage_ltng3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
3	1	429	skmage_ltng4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
3	1	430	hellbovine-Hell Bovine-Skeleton				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
3	1	431	window1--Idle				/Data/Global/Monsters	VH	NU	HTH		LIT							LIT										0
3	1	432	window2--Idle				/Data/Global/Monsters	VJ	NU	HTH		LIT							LIT										0
3	1	433	slinger5-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
3	1	434	slinger6-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
3	1	435	fetishblow1-RatMan-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	436	fetishblow2-Fetish-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	437	fetishblow3-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	438	fetishblow4-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	439	fetishblow5-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	440	mephistospirit-Dummy-Spirit				/Data/Global/Monsters	M6	A1	HTH		LIT																	0
3	1	441	smith-The Smith-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
3	1	442	trappedsoul1-TrappedSoul-TrappedSoul				/Data/Global/Monsters	10	NU	HTH		LIT																	0
3	1	443	trappedsoul2-TrappedSoul-TrappedSoul				/Data/Global/Monsters	13	NU	HTH		LIT																	0
3	1	444	jamella-Jamella-Npc				/Data/Global/Monsters	ja	NU	HTH		LIT																	0
3	1	445	izualghost-Izual-NpcStationary				/Data/Global/Monsters	17	NU	HTH		LIT							LIT										0
3	1	446	fetish11-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	447	malachai-Malachai-Buffy				/Data/Global/Monsters	36	NU	HTH		LIT							LIT										0
3	1	448	hephasto-The Feature Creep-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
3	1	449	wakeofdestruction-Wake of Destruction-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
3	1	450	chargeboltsentry-Charged Bolt Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
3	1	451	lightningsentry-Lightning Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
3	1	452	bladecreeper-Blade Creeper-BladeCreeper				/Data/Global/Monsters	b8	NU	HTH		LIT							LIT										0
3	1	453	invisopet-Invis Pet-InvisoPet				/Data/Global/Monsters	k9																					0
3	1	454	infernosentry-Inferno Sentry-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
3	1	455	deathsentry-Death Sentry-DeathSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
3	1	456	shadowwarrior-Shadow Warrior-ShadowWarrior				/Data/Global/Monsters	k9																					0
3	1	457	shadowmaster-Shadow Master-ShadowMaster				/Data/Global/Monsters	k9																					0
3	1	458	druidhawk-Druid Hawk-Raven				/Data/Global/Monsters	hk	NU	HTH		LIT																	0
3	1	459	spiritwolf-Druid Spirit Wolf-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
3	1	460	fenris-Druid Fenris-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
3	1	461	spiritofbarbs-Spirit of Barbs-Totem				/Data/Global/Monsters	x4	NU	HTH		LIT																	0
3	1	462	heartofwolverine-Heart of Wolverine-Totem				/Data/Global/Monsters	x3	NU	HTH		LIT																	0
3	1	463	oaksage-Oak Sage-Totem				/Data/Global/Monsters	xw	NU	HTH		LIT																	0
3	1	464	plaguepoppy-Druid Plague Poppy-Vines				/Data/Global/Monsters	k9																					0
3	1	465	cycleoflife-Druid Cycle of Life-CycleOfLife				/Data/Global/Monsters	k9																					0
3	1	466	vinecreature-Vine Creature-CycleOfLife				/Data/Global/Monsters	k9																					0
3	1	467	druidbear-Druid Bear-DruidBear				/Data/Global/Monsters	b7	NU	HTH		LIT																	0
3	1	468	eagle-Eagle-Idle				/Data/Global/Monsters	eg	NU	HTH		LIT							LIT										0
3	1	469	wolf-Wolf-NecroPet				/Data/Global/Monsters	40	NU	HTH		LIT																	0
3	1	470	bear-Bear-NecroPet				/Data/Global/Monsters	TG	NU	HTH		LIT							LIT										0
3	1	471	barricadedoor1-Barricade Door-Idle				/Data/Global/Monsters	AJ	NU	HTH		LIT																	0
3	1	472	barricadedoor2-Barricade Door-Idle				/Data/Global/Monsters	AG	NU	HTH		LIT																	0
3	1	473	prisondoor-Prison Door-Idle				/Data/Global/Monsters	2Q	NU	HTH		LIT																	0
3	1	474	barricadetower-Barricade Tower-SiegeTower				/Data/Global/Monsters	ac	NU	HTH		LIT							LIT						LIT				0
3	1	475	reanimatedhorde1-RotWalker-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	476	reanimatedhorde2-ReanimatedHorde-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	477	reanimatedhorde3-ProwlingDead-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	478	reanimatedhorde4-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	479	reanimatedhorde5-DefiledWarrior-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	480	siegebeast1-Siege Beast-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
3	1	481	siegebeast2-CrushBiest-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
3	1	482	siegebeast3-BloodBringer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
3	1	483	siegebeast4-GoreBearer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
3	1	484	siegebeast5-DeamonSteed-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
3	1	485	snowyeti1-SnowYeti1-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
3	1	486	snowyeti2-SnowYeti2-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
3	1	487	snowyeti3-SnowYeti3-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
3	1	488	snowyeti4-SnowYeti4-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
3	1	489	wolfrider1-WolfRider1-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
3	1	490	wolfrider2-WolfRider2-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
3	1	491	wolfrider3-WolfRider3-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
3	1	492	minion1-Minionexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	493	minion2-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	494	minion3-IceBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	495	minion4-FireBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	496	minion5-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	497	minion6-IceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	498	minion7-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	499	minion8-GreaterIceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	500	suicideminion1-FanaticMinion-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	501	suicideminion2-BerserkSlayer-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	502	suicideminion3-ConsumedIceBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	503	suicideminion4-ConsumedFireBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	504	suicideminion5-FrenziedHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	505	suicideminion6-FrenziedIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	506	suicideminion7-InsaneHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	507	suicideminion8-InsaneIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
3	1	508	succubus1-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	509	succubus2-VileTemptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	510	succubus3-StygianHarlot-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	511	succubus4-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	512	succubus5-Blood Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	513	succubuswitch1-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	514	succubuswitch2-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	515	succubuswitch3-StygianFury-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	516	succubuswitch4-Blood Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	517	succubuswitch5-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	518	overseer1-OverSeer-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	519	overseer2-Lasher-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	520	overseer3-OverLord-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	521	overseer4-BloodBoss-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	522	overseer5-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	523	minionspawner1-MinionSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	524	minionspawner2-MinionSlayerSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	525	minionspawner3-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	526	minionspawner4-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	527	minionspawner5-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	528	minionspawner6-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	529	minionspawner7-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	530	minionspawner8-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	531	imp1-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	532	imp2-Imp2-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	533	imp3-Imp3-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	534	imp4-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	535	imp5-Imp5-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	536	catapult1-CatapultS-Catapult				/Data/Global/Monsters	65	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
3	1	537	catapult2-CatapultE-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
3	1	538	catapult3-CatapultSiege-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
3	1	539	catapult4-CatapultW-Catapult				/Data/Global/Monsters	ua	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT	LIT								0
3	1	540	frozenhorror1-Frozen Horror1-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
3	1	541	frozenhorror2-Frozen Horror2-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
3	1	542	frozenhorror3-Frozen Horror3-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
3	1	543	frozenhorror4-Frozen Horror4-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
3	1	544	frozenhorror5-Frozen Horror5-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
3	1	545	bloodlord1-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	546	bloodlord2-Blood Lord2-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	547	bloodlord3-Blood Lord3-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	548	bloodlord4-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	549	bloodlord5-Blood Lord5-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	550	larzuk-Larzuk-Npc				/Data/Global/Monsters	XR	NU	HTH		LIT																	0
3	1	551	drehya-Drehya-Npc				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
3	1	552	malah-Malah-Npc				/Data/Global/Monsters	XT	NU	HTH		LIT																	0
3	1	553	nihlathak-Nihlathak Town-Npc				/Data/Global/Monsters	0J	NU	HTH		LIT																	0
3	1	554	qual-kehk-Qual-Kehk-Npc				/Data/Global/Monsters	XV	NU	HTH		LIT																	0
3	1	555	catapultspotter1-Catapult Spotter S-CatapultSpotter				/Data/Global/Monsters	k9																					0
3	1	556	catapultspotter2-Catapult Spotter E-CatapultSpotter				/Data/Global/Monsters	k9																					0
3	1	557	catapultspotter3-Catapult Spotter Siege-CatapultSpotter				/Data/Global/Monsters	k9																					0
3	1	558	catapultspotter4-Catapult Spotter W-CatapultSpotter				/Data/Global/Monsters	k9																					0
3	1	559	cain6-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
3	1	560	tyrael3-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
3	1	561	act5barb1-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
3	1	562	act5barb2-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
3	1	563	barricadewall1-Barricade Wall Right-Idle				/Data/Global/Monsters	A6	NU	HTH		LIT																	0
3	1	564	barricadewall2-Barricade Wall Left-Idle				/Data/Global/Monsters	AK	NU	HTH		LIT																	0
3	1	565	nihlathakboss-Nihlathak-Nihlathak				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
3	1	566	drehyaiced-Drehya-NpcOutOfTown				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
3	1	567	evilhut-Evil hut-GenericSpawner				/Data/Global/Monsters	2T	NU	HTH		LIT							LIT										0
3	1	568	deathmauler1-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	569	deathmauler2-Death Mauler2-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	570	deathmauler3-Death Mauler3-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	571	deathmauler4-Death Mauler4-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	572	deathmauler5-Death Mauler5-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	573	act5pow-POW-Wussie				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
3	1	574	act5barb3-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
3	1	575	act5barb4-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
3	1	576	ancientstatue1-Ancient Statue 1-AncientStatue				/Data/Global/Monsters	0G	NU	HTH		LIT																	0
3	1	577	ancientstatue2-Ancient Statue 2-AncientStatue				/Data/Global/Monsters	0H	NU	HTH		LIT																	0
3	1	578	ancientstatue3-Ancient Statue 3-AncientStatue				/Data/Global/Monsters	0I	NU	HTH		LIT																	0
3	1	579	ancientbarb1-Ancient Barbarian 1-Ancient				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
3	1	580	ancientbarb2-Ancient Barbarian 2-Ancient				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
3	1	581	ancientbarb3-Ancient Barbarian 3-Ancient				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
3	1	582	baalthrone-Baal Throne-BaalThrone				/Data/Global/Monsters	41	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
3	1	583	baalcrab-Baal Crab-BaalCrab				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
3	1	584	baaltaunt-Baal Taunt-BaalTaunt				/Data/Global/Monsters	K9																					0
3	1	585	putriddefiler1-Putrid Defiler1-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
3	1	586	putriddefiler2-Putrid Defiler2-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
3	1	587	putriddefiler3-Putrid Defiler3-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
3	1	588	putriddefiler4-Putrid Defiler4-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
3	1	589	putriddefiler5-Putrid Defiler5-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
3	1	590	painworm1-Pain Worm1-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
3	1	591	painworm2-Pain Worm2-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
3	1	592	painworm3-Pain Worm3-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
3	1	593	painworm4-Pain Worm4-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
3	1	594	painworm5-Pain Worm5-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
3	1	595	bunny-dummy-Idle				/Data/Global/Monsters	48	NU	HTH		LIT																	0
3	1	596	baalhighpriest-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	597	venomlord-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				FLB													0
3	1	598	baalcrabstairs-Baal Crab to Stairs-BaalToStairs				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
3	1	599	act5hire1-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
3	1	600	act5hire2-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
3	1	601	baaltentacle1-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
3	1	602	baaltentacle2-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
3	1	603	baaltentacle3-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
3	1	604	baaltentacle4-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
3	1	605	baaltentacle5-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
3	1	606	injuredbarb1-dummy-Idle				/Data/Global/Monsters	6z	NU	HTH		LIT																	0
3	1	607	injuredbarb2-dummy-Idle				/Data/Global/Monsters	7j	NU	HTH		LIT																	0
3	1	608	injuredbarb3-dummy-Idle				/Data/Global/Monsters	7i	NU	HTH		LIT																	0
3	1	609	baalclone-Baal Crab Clone-BaalCrabClone				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
3	1	610	baalminion1-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
3	1	611	baalminion2-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
3	1	612	baalminion3-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
3	1	613	worldstoneeffect-dummy-Idle				/Data/Global/Monsters	K9																					0
3	1	614	sk_archer6-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	615	sk_archer7-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	616	sk_archer8-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	617	sk_archer9-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	618	sk_archer10-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	619	bighead6-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	620	bighead7-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	621	bighead8-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	622	bighead9-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	623	bighead10-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	624	goatman6-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	625	goatman7-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	626	goatman8-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	627	goatman9-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	628	goatman10-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
3	1	629	foulcrow5-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	630	foulcrow6-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	631	foulcrow7-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	632	foulcrow8-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
3	1	633	clawviper6-ClawViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	634	clawviper7-PitViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	635	clawviper8-Salamander-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	636	clawviper9-TombViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	637	clawviper10-SerpentMagus-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	638	sandraider6-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	639	sandraider7-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	640	sandraider8-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	641	sandraider9-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	642	sandraider10-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	643	deathmauler6-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	644	quillrat6-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	645	quillrat7-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	646	quillrat8-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	647	vulture5-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
3	1	648	thornhulk5-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
3	1	649	slinger7-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	650	slinger8-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	651	slinger9-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	652	cr_archer6-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	653	cr_archer7-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	654	cr_lancer6-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	655	cr_lancer7-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	656	cr_lancer8-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	657	blunderbore5-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
3	1	658	blunderbore6-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
3	1	659	skmage_fire5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
3	1	660	skmage_fire6-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
3	1	661	skmage_ltng5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
3	1	662	skmage_ltng6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
3	1	663	skmage_cold5-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		CLD	CLD						0
3	1	664	skmage_pois5-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	665	skmage_pois6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	666	pantherwoman5-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	667	pantherwoman6-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	668	sandleaper6-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	669	sandleaper7-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
3	1	670	wraith6-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	671	wraith7-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	672	wraith8-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	673	succubus6-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	674	succubus7-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	675	succubuswitch6-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	676	succubuswitch7-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	677	succubuswitch8-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	678	willowisp5-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	679	willowisp6-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	680	willowisp7-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	681	fallen6-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
3	1	682	fallen7-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
3	1	683	fallen8-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
3	1	684	fallenshaman6-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	685	fallenshaman7-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	686	fallenshaman8-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	687	skeleton6-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	688	skeleton7-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	689	batdemon6-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	690	batdemon7-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	691	bloodlord6-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	692	bloodlord7-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	693	scarab6-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	694	scarab7-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	695	fetish6-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	696	fetish7-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	697	fetish8-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
3	1	698	fetishblow6-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	699	fetishblow7-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	700	fetishblow8-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
3	1	701	fetishshaman6-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	702	fetishshaman7-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	703	fetishshaman8-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	704	baboon7-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	705	baboon8-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
3	1	706	unraveler6-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	707	unraveler7-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	708	unraveler8-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	709	unraveler9-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	710	zealot4-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
3	1	711	zealot5-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
3	1	712	cantor5-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
3	1	713	cantor6-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
3	1	714	vilemother4-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
3	1	715	vilemother5-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
3	1	716	vilechild4-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
3	1	717	vilechild5-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
3	1	718	sandmaggot6-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	719	maggotbaby6-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
3	1	720	maggotegg6-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
3	1	721	minion9-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	722	minion10-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	723	minion11-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	724	arach6-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	725	megademon4-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	726	megademon5-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	727	imp6-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	728	imp7-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	729	bonefetish6-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	730	bonefetish7-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
3	1	731	fingermage4-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
3	1	732	fingermage5-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
3	1	733	regurgitator4-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
3	1	734	vampire6-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	735	vampire7-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	736	vampire8-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	737	reanimatedhorde6-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	738	dkfig1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	739	dkfig2-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	740	dkmag1-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	741	dkmag2-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	742	mummy6-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	743	ubermephisto-Mephisto-UberMephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
3	1	744	uberdiablo-Diablo-UberDiablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
3	1	745	uberizual-izual-UberIzual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
3	1	746	uberandariel-Lilith-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
3	1	747	uberduriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
3	1	748	uberbaal-Baal Crab-UberBaal				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
3	1	749	demonspawner-Evil hut-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
3	1	750	demonhole-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
3	1	751	megademon6-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	752	dkmag3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	753	imp8-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	754	swarm5-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
3	1	755	sandmaggot7-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
3	1	756	arach7-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	757	scarab8-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	758	succubus8-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	759	succubuswitch9-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	760	corruptrogue6-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	761	cr_archer8-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	762	cr_lancer9-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
3	1	763	overseer6-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	764	skeleton8-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	765	sk_archer11-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
3	1	766	skmage_fire7-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
3	1	767	skmage_ltng7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
3	1	768	skmage_cold6-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
3	1	769	skmage_pois7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		POS	POS						0
3	1	770	vampire9-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
3	1	771	wraith9-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
3	1	772	willowisp8-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	773	Bishibosh-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	774	Bonebreak-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
3	1	775	Coldcrow-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
3	1	776	Rakanishu-SUPER UNIQUE				/Data/Global/Monsters	FA	NU	HTH		LIT				SWD		TCH	LIT										0
3	1	777	Treehead WoodFist-SUPER UNIQUE				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
3	1	778	Griswold-SUPER UNIQUE				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
3	1	779	The Countess-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	MED	LIT	MED	LIT	LIT	WHM			LIT	LIT									0
3	1	780	Pitspawn Fouldog-SUPER UNIQUE				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
3	1	781	Flamespike the Crawler-SUPER UNIQUE				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
3	1	782	Boneash-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
3	1	783	Radament-SUPER UNIQUE				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
3	1	784	Bloodwitch the Wild-SUPER UNIQUE				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
3	1	785	Fangskin-SUPER UNIQUE				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
3	1	786	Beetleburst-SUPER UNIQUE				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
3	1	787	Leatherarm-SUPER UNIQUE				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
3	1	788	Coldworm the Burrower-SUPER UNIQUE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
3	1	789	Fire Eye-SUPER UNIQUE				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
3	1	790	Dark Elder-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	791	The Summoner-SUPER UNIQUE				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
3	1	792	Ancient Kaa the Soulless-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	793	The Smith-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
3	1	794	Web Mage the Burning-SUPER UNIQUE				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
3	1	795	Witch Doctor Endugu-SUPER UNIQUE				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
3	1	796	Stormtree-SUPER UNIQUE				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
3	1	797	Sarina the Battlemaid-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
3	1	798	Icehawk Riftwing-SUPER UNIQUE				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
3	1	799	Ismail Vilehand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	800	Geleb Flamefinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	801	Bremm Sparkfist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	802	Toorc Icefist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	803	Wyand Voidfinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	804	Maffer Dragonhand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	805	Winged Death-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	806	The Tormentor-SUPER UNIQUE				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
3	1	807	Taintbreeder-SUPER UNIQUE				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
3	1	808	Riftwraith the Cannibal-SUPER UNIQUE				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
3	1	809	Infector of Souls-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	810	Lord De Seis-SUPER UNIQUE				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
3	1	811	Grand Vizier of Chaos-SUPER UNIQUE				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
3	1	812	The Cow King-SUPER UNIQUE				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
3	1	813	Corpsefire-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
3	1	814	The Feature Creep-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
3	1	815	Siege Boss-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	816	Ancient Barbarian 1-SUPER UNIQUE				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
3	1	817	Ancient Barbarian 2-SUPER UNIQUE				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
3	1	818	Ancient Barbarian 3-SUPER UNIQUE				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
3	1	819	Axe Dweller-SUPER UNIQUE				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
3	1	820	Bonesaw Breaker-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	821	Dac Farren-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	822	Megaflow Rectifier-SUPER UNIQUE				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
3	1	823	Eyeback Unleashed-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	824	Threash Socket-SUPER UNIQUE				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
3	1	825	Pindleskin-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
3	1	826	Snapchip Shatter-SUPER UNIQUE				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
3	1	827	Anodized Elite-SUPER UNIQUE				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
3	1	828	Vinvear Molech-SUPER UNIQUE				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
3	1	829	Sharp Tooth Sayer-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
3	1	830	Magma Torquer-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
3	1	831	Blaze Ripper-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
3	1	832	Frozenstein-SUPER UNIQUE				/Data/Global/Monsters	io	NU	HTH		LIT																	0
3	1	833	Nihlathak Boss-SUPER UNIQUE				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
3	1	834	Baal Subject 1-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
3	1	835	Baal Subject 2-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
3	1	836	Baal Subject 3-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
3	1	837	Baal Subject 4-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
3	1	838	Baal Subject 5-SUPER UNIQUE				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
3	2	0	jungle torch (117)	117			/Data/Global/Objects	JT	ON	HTH		LIT							LIT										0
3	2	1	Waypoint (237)	237			/Data/Global/Objects	WZ	ON	HTH		LIT							LIT										0
3	2	2	-580	580																									0
3	2	3	Well, pool wilderness (130)	130			/Data/Global/Objects	ZW	NU	HTH		LIT																	0
3	2	4	brazier floor (102)	102			/Data/Global/Objects	FB	ON	HTH		LIT							LIT										0
3	2	5	torch 1 tiki (37)	37			/Data/Global/Objects	TO	ON	HTH		LIT																	0
3	2	6	Fire, small (160)	160			/Data/Global/Objects	FX	NU	HTH		LIT																	0
3	2	7	Fire, medium (161)	161			/Data/Global/Objects	FY	NU	HTH		LIT																	0
3	2	8	Fire, large (162)	162			/Data/Global/Objects	FZ	ON	HTH		LIT																	0
3	2	9	Armor Stand, 1 R (104)	104			/Data/Global/Objects	A3	NU	HTH		LIT																	0
3	2	10	Armor Stand, 2 L (105)	105			/Data/Global/Objects	A4	NU	HTH		LIT																	0
3	2	11	Weapon Rack, 1 R (106)	106			/Data/Global/Objects	W1	NU	HTH		LIT																	0
3	2	12	Weapon Rack, 2 L (107)	107			/Data/Global/Objects	W2	NU	HTH		LIT																	0
3	2	13	Stair, L altar to underground (194)	194			/Data/Global/Objects	9C	OP	HTH		LIT																	0
3	2	14	Stair, R altar to underground (195)	195			/Data/Global/Objects	SV	OP	HTH		LIT																	0
3	2	15	Lam Esen's Tome (193)	193			/Data/Global/Objects	AB	NU	HTH		LIT																	0
3	2	16	water rocks 1 (207)	207			/Data/Global/Objects	RW	NU	HTH		LIT																	0
3	2	17	water rocks girl 1 (211)	211			/Data/Global/Objects	WB	NU	HTH		LIT																	0
3	2	18	water logs 1 (210)	210			/Data/Global/Objects	LW	NU	HTH		LIT																	0
3	2	19	water log bigger 1 (234)	234			/Data/Global/Objects	QY	NU	HTH		LIT																	0
3	2	20	water rocks 2 (214)	214			/Data/Global/Objects	WC	NU	HTH		LIT																	0
3	2	21	water rocks girl 2 (215)	215			/Data/Global/Objects	WE	NU	HTH		LIT																	0
3	2	22	water logs 2 (213)	213			/Data/Global/Objects	WD	NU	HTH		LIT																	0
3	2	23	water log bigger 2 (228)	228			/Data/Global/Objects	QW	NU	HTH		LIT																	0
3	2	24	water rocks 3 (216)	216			/Data/Global/Objects	WY	NU	HTH		LIT																	0
3	2	25	water rocks girl 3 (227)	227			/Data/Global/Objects	QX	NU	HTH		LIT																	0
3	2	26	water logs 3 (217)	217			/Data/Global/Objects	LX	NU	HTH		LIT																	0
3	2	27	water log bigger 3 (235)	235			/Data/Global/Objects	QZ	NU	HTH		LIT																	0
3	2	28	web between 2 trees L (218)	218			/Data/Global/Objects	W3	NU	HTH		LIT							LIT										0
3	2	29	web between 2 trees R (219)	219			/Data/Global/Objects	W4	NU	HTH		LIT							LIT										0
3	2	30	web around 1 tree L (220)	220			/Data/Global/Objects	W5	NU	HTH		LIT							LIT										0
3	2	31	web around 1 tree R (221)	221			/Data/Global/Objects	W6	NU	HTH		LIT							LIT										0
3	2	32	Cocoon, living (223)	223			/Data/Global/Objects	CN	NU	HTH		LIT																	0
3	2	33	Cocoon, static (224)	224			/Data/Global/Objects	CC	NU	HTH		LIT																	0
3	2	34	Your Private Stash (267)	267			/Data/Global/Objects	B6	NU	HTH		LIT																	0
3	2	35	gold placeholder (269)	269			/Data/Global/Objects	1G	NU	HTH		LIT																	0
3	2	36	-581	581																									0
3	2	37	Shrine, jungle heal well (170)	170			/Data/Global/Objects	JH	OP	HTH		LIT																	0
3	2	38	Magic Shrine, sewer (325)	325			/Data/Global/Objects	QN	NU	HTH		LIT							LIT										0
3	2	39	Shrine, jungle 1 (184)	184			/Data/Global/Objects	JY	NU	HTH		LIT							LIT										0
3	2	40	Shrine, jungle 2 (190)	190			/Data/Global/Objects	JS	NU	HTH		LIT							LIT										0
3	2	41	Shrine, jungle 3 (191)	191			/Data/Global/Objects	JR	NU	HTH		LIT							LIT										0
3	2	42	Shrine, jungle 4 (197)	197			/Data/Global/Objects	JQ	NU	HTH		LIT							LIT										0
3	2	43	Shrine, mephisto 1 (199)	199			/Data/Global/Objects	MZ	NU	HTH		LIT							LIT										0
3	2	44	Shrine, mephisto 2 (200)	200			/Data/Global/Objects	MY	NU	HTH		LIT							LIT										0
3	2	45	Shrine, mephisto 3 (201)	201			/Data/Global/Objects	MX	NU	HTH		LIT							LIT										0
3	2	46	Shrine, mephisto mana (202)	202			/Data/Global/Objects	MW	OP	HTH		LIT							LIT										0
3	2	47	Shrine, mephisto health (206)	206			/Data/Global/Objects	MR	OP	HTH		LIT							LIT										0
3	2	48	Shrine, mana dungeon (278)	278			/Data/Global/Objects	DE	OP	HTH		LIT																	0
3	2	49	dummy shrine health dungeon (120)	120			/Data/Global/Objects	DJ	OP	HTH		LIT																	0
3	2	50	Well, pool wilderness (130)	130			/Data/Global/Objects	ZW	NU	HTH		LIT																	0
3	2	51	Dead Body, sewer (326)	326			/Data/Global/Objects	QO	OP	HTH		LIT																	0
3	2	52	Skeleton (158)	158			/Data/Global/Objects	CP	OP	HTH		LIT																	0
3	2	53	Corpse, villager 1 (271)	271			/Data/Global/Objects	DG	OP	HTH		LIT																	0
3	2	54	Corpse, villager 2 (272)	272			/Data/Global/Objects	DF	OP	HTH		LIT																	0
3	2	55	torch 1 (327)	327			/Data/Global/Objects	V1	NU	HTH		LIT							LIT										0
3	2	56	torch 2 (328)	328			/Data/Global/Objects	V2	NU	HTH		LIT							LIT										0
3	2	57	Chest, Mephisto L Large (329)	329			/Data/Global/Objects	XB	OP	HTH		LIT																	0
3	2	58	Chest, Mephisto R Large (330)	330			/Data/Global/Objects	XC	OP	HTH		LIT																	0
3	2	59	Chest, Mephisto L Med (331)	331			/Data/Global/Objects	XD	OP	HTH		LIT																	0
3	2	60	Chest, Mephisto R Med (332)	332			/Data/Global/Objects	XE	OP	HTH		LIT																	0
3	2	61	Chest, spider lair L Large (333)	333			/Data/Global/Objects	XF	OP	HTH		LIT																	0
3	2	62	Chest, spider lair L Tall (334)	334			/Data/Global/Objects	XG	OP	HTH		LIT																	0
3	2	63	Chest, spider lair R Med (335)	335			/Data/Global/Objects	XH	OP	HTH		LIT																	0
3	2	64	Chest, spider lair R Tall (336)	336			/Data/Global/Objects	XI	OP	HTH		LIT																	0
3	2	65	Chest, R Large (5)	5			/Data/Global/Objects	L1	OP	HTH		LIT																	0
3	2	66	Chest, L Large (6)	6			/Data/Global/Objects	L2	OP	HTH		LIT																	0
3	2	67	Chest, L Med (176)	176			/Data/Global/Objects	C8	OP	HTH		LIT																	0
3	2	68	Chest, general L (240)	240			/Data/Global/Objects	CY	OP	HTH		LIT																	0
3	2	69	Chest, general R (241)	241			/Data/Global/Objects	CX	OP	HTH		LIT																	0
3	2	70	Chest, jungle (181)	181			/Data/Global/Objects	JC	OP	HTH		LIT																	0
3	2	71	Chest, L Med jungle (183)	183			/Data/Global/Objects	JZ	OP	HTH		LIT																	0
3	2	72	Rat's Nest, sewers (246)	246			/Data/Global/Objects	RA	OP	HTH		LIT																	0
3	2	73	Stash, jungle 1 (185)	185			/Data/Global/Objects	JX	OP	HTH		LIT																	0
3	2	74	Stash, jungle 2 (186)	186			/Data/Global/Objects	JW	OP	HTH		LIT																	0
3	2	75	Stash, jungle 3 (187)	187			/Data/Global/Objects	JV	OP	HTH		LIT																	0
3	2	76	Stash, jungle 4 (188)	188			/Data/Global/Objects	JU	OP	HTH		LIT																	0
3	2	77	Stash, Mephisto lair (203)	203			/Data/Global/Objects	MV	OP	HTH		LIT																	0
3	2	78	Stash, box (204)	204			/Data/Global/Objects	MU	OP	HTH		LIT																	0
3	2	79	Stash, altar (205)	205			/Data/Global/Objects	MT	OP	HTH		LIT																	0
3	2	80	Basket, 1 say 'not here' (208)	208			/Data/Global/Objects	BD	OP	HTH		LIT																	0
3	2	81	Basket, 2 say 'not here' (209)	209			/Data/Global/Objects	BJ	OP	HTH		LIT																	0
3	2	82	Hollow Log (169)	169			/Data/Global/Objects	CZ	NU	HTH		LIT																	0
3	2	83	Waypoint, sewer (323)	323			/Data/Global/Objects	QM	ON	HTH		LIT							LIT										0
3	2	84	Waypoint, Travincal (324)	324			/Data/Global/Objects	QL	ON	HTH		LIT							LIT										0
3	2	85	a Trap, test data floortrap (196)	196			/Data/Global/Objects	A5	OP	HTH		LIT																	0
3	2	86	bubbles in water (212)	212			/Data/Global/Objects	YB	NU	HTH		LIT																	0
3	2	87	Skullpile (225)	225			/Data/Global/Objects	IB	OP	HTH		LIT																	0
3	2	88	Rat's Nest, sewers (244)	244			/Data/Global/Objects	RN	OP	HTH		LIT																	0
3	2	89	helllight source 1 (351)	351																									0
3	2	90	helllight source 2 (352)	352																									0
3	2	91	helllight source 3 (353)	353																									0
3	2	92	Rock Pile, dungeon (360)	360			/Data/Global/Objects	XN	OP	HTH		LIT																	0
3	2	93	Magic Shrine, dungeon (361)	361			/Data/Global/Objects	XO	OP	HTH		LIT							LIT										0
3	2	94	Basket, dungeon (362)	362			/Data/Global/Objects	XP	OP	HTH		LIT																	0
3	2	95	Casket, dungeon (365)	365			/Data/Global/Objects	VB	OP	HTH		LIT																	0
3	2	96	Gidbinn Altar (251)	251			/Data/Global/Objects	GA	ON	HTH		LIT							LIT										0
3	2	97	Gidbinn, decoy (252)	252			/Data/Global/Objects	GD	ON	HTH		LIT							LIT										0
3	2	98	Basket, 1 (208)	208			/Data/Global/Objects	BD	OP	HTH		LIT																	0
3	2	99	brazier celler (283)	283			/Data/Global/Objects	BI	NU	HTH		LIT							LIT										0
3	2	100	Sewer Lever (367)	367			/Data/Global/Objects	VF	OP	HTH		LIT																	0
3	2	101	Sewer Stairs (366)	366			/Data/Global/Objects	VE	OP	HTH		LIT																	0
3	2	102	Dark Wanderer (368)	368																									0
3	2	103	Mephisto bridge (341)	341			/Data/Global/Objects	XJ	ON	HTH		LIT																	0
3	2	342	Portal to hellgate (342)	342			/Data/Global/Objects	1Y	ON	HTH		LIT								LIT	LIT								0
3	2	105	Shrine, mana well kurast (343)	343			/Data/Global/Objects	XL	OP	HTH		LIT																	0
3	2	106	Shrine, health well kurast (344)	344			/Data/Global/Objects	XM	OP	HTH		LIT																	0
3	2	107	fog water (374)	374			/Data/Global/Objects	UD	NU	HTH		LIT																	0
3	2	108	torch town (370)	370			/Data/Global/Objects	VG	NU	HTH		LIT							LIT										0
3	2	109	Hratli start (378)	378		1	/Data/Global/Monsters	HR	NU	HTH		LIT																	0
3	2	110	Hratli end (379)	379		7	/Data/Global/Monsters	HR	NU	HTH		LIT																	0
3	2	111	stairs of Compelling Orb (386)	386			/Data/Global/Objects	SV	OP	HTH		LIT																	0
3	2	112	Chest, sparkly (397)	397			/Data/Global/Objects	YF	OP	HTH		LIT																	0
3	2	113	Chest, Khalim's Heart (405)	405			/Data/Global/Objects	XK	OP	HTH		LIT																	0
3	2	114	Chest, Khalim's Eye (407)	407			/Data/Global/Objects	XK	OP	HTH		LIT																	0
3	2	115	Chest, Khalim's Brain (406)	406			/Data/Global/Objects	XK	OP	HTH		LIT																	0
3	2	116	ACT 3 TABLE SKIP IT	0																									0
3	2	117	ACT 3 TABLE SKIP IT	0																									0
3	2	118	ACT 3 TABLE SKIP IT	0																									0
3	2	119	ACT 3 TABLE SKIP IT	0																									0
3	2	120	ACT 3 TABLE SKIP IT	0																									0
3	2	121	ACT 3 TABLE SKIP IT	0																									0
3	2	122	ACT 3 TABLE SKIP IT	0																									0
3	2	123	ACT 3 TABLE SKIP IT	0																									0
3	2	124	ACT 3 TABLE SKIP IT	0																									0
3	2	125	ACT 3 TABLE SKIP IT	0																									0
3	2	126	ACT 3 TABLE SKIP IT	0																									0
3	2	127	ACT 3 TABLE SKIP IT	0																									0
3	2	128	ACT 3 TABLE SKIP IT	0																									0
3	2	129	ACT 3 TABLE SKIP IT	0																									0
3	2	130	ACT 3 TABLE SKIP IT	0																									0
3	2	131	ACT 3 TABLE SKIP IT	0																									0
3	2	132	ACT 3 TABLE SKIP IT	0																									0
3	2	133	ACT 3 TABLE SKIP IT	0																									0
3	2	134	ACT 3 TABLE SKIP IT	0																									0
3	2	135	ACT 3 TABLE SKIP IT	0																									0
3	2	136	ACT 3 TABLE SKIP IT	0																									0
3	2	137	ACT 3 TABLE SKIP IT	0																									0
3	2	138	ACT 3 TABLE SKIP IT	0																									0
3	2	139	ACT 3 TABLE SKIP IT	0																									0
3	2	140	ACT 3 TABLE SKIP IT	0																									0
3	2	141	ACT 3 TABLE SKIP IT	0																									0
3	2	142	ACT 3 TABLE SKIP IT	0																									0
3	2	143	ACT 3 TABLE SKIP IT	0																									0
3	2	144	ACT 3 TABLE SKIP IT	0																									0
3	2	145	ACT 3 TABLE SKIP IT	0																									0
3	2	146	ACT 3 TABLE SKIP IT	0																									0
3	2	147	ACT 3 TABLE SKIP IT	0																									0
3	2	148	ACT 3 TABLE SKIP IT	0																									0
3	2	149	ACT 3 TABLE SKIP IT	0																									0
3	2	150	Dummy-test data SKIPT IT				/Data/Global/Objects	NU0																					
3	2	151	Casket-Casket #5				/Data/Global/Objects	C5	OP	HTH		LIT																	
3	2	152	Shrine-Shrine				/Data/Global/Objects	SF	OP	HTH		LIT																	
3	2	153	Casket-Casket #6				/Data/Global/Objects	C6	OP	HTH		LIT																	
3	2	154	LargeUrn-Urn #1				/Data/Global/Objects	U1	OP	HTH		LIT																	
3	2	155	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
3	2	156	chest-LargeChestL				/Data/Global/Objects	L2	OP	HTH		LIT																	
3	2	157	Barrel-Barrel				/Data/Global/Objects	B1	OP	HTH		LIT																	
3	2	158	TowerTome-Tower Tome				/Data/Global/Objects	TT	OP	HTH		LIT																	
3	2	159	Urn-Urn #2				/Data/Global/Objects	U2	OP	HTH		LIT																	
3	2	160	Dummy-Bench				/Data/Global/Objects	BE	NU	HTH		LIT																	
3	2	161	Barrel-BarrelExploding				/Data/Global/Objects	BX	OP	HTH		LIT							LIT	LIT									
3	2	162	Dummy-RogueFountain				/Data/Global/Objects	FN	NU	HTH		LIT																	
3	2	163	Door-Door Gate Left				/Data/Global/Objects	D1	OP	HTH		LIT																	
3	2	164	Door-Door Gate Right				/Data/Global/Objects	D2	OP	HTH		LIT																	
3	2	165	Door-Door Wooden Left				/Data/Global/Objects	D3	OP	HTH		LIT																	
3	2	166	Door-Door Wooden Right				/Data/Global/Objects	D4	OP	HTH		LIT																	
3	2	167	StoneAlpha-StoneAlpha				/Data/Global/Objects	S1	OP	HTH		LIT																	
3	2	168	StoneBeta-StoneBeta				/Data/Global/Objects	S2	OP	HTH		LIT																	
3	2	169	StoneGamma-StoneGamma				/Data/Global/Objects	S3	OP	HTH		LIT																	
3	2	170	StoneDelta-StoneDelta				/Data/Global/Objects	S4	OP	HTH		LIT																	
3	2	171	StoneLambda-StoneLambda				/Data/Global/Objects	S5	OP	HTH		LIT																	
3	2	172	StoneTheta-StoneTheta				/Data/Global/Objects	S6	OP	HTH		LIT																	
3	2	173	Door-Door Courtyard Left				/Data/Global/Objects	D5	OP	HTH		LIT																	
3	2	174	Door-Door Courtyard Right				/Data/Global/Objects	D6	OP	HTH		LIT																	
3	2	175	Door-Door Cathedral Double				/Data/Global/Objects	D7	OP	HTH		LIT																	
3	2	176	Gibbet-Cain's Been Captured				/Data/Global/Objects	GI	OP	HTH		LIT																	
3	2	177	Door-Door Monastery Double Right				/Data/Global/Objects	D8	OP	HTH		LIT																	
3	2	178	HoleAnim-Hole in Ground				/Data/Global/Objects	HI	OP	HTH		LIT																	
3	2	179	Dummy-Brazier				/Data/Global/Objects	BR	ON	HTH		LIT							LIT										
3	2	180	Inifuss-inifuss tree				/Data/Global/Objects	IT	NU	HTH		LIT																	
3	2	181	Dummy-Fountain				/Data/Global/Objects	BF	NU	HTH		LIT																	
3	2	182	Dummy-crucifix				/Data/Global/Objects	CL	NU	HTH		LIT																	
3	2	183	Dummy-Candles1				/Data/Global/Objects	A1	NU	HTH		LIT																	
3	2	184	Dummy-Candles2				/Data/Global/Objects	A2	NU	HTH		LIT																	
3	2	185	Dummy-Standard1				/Data/Global/Objects	N1	NU	HTH		LIT																	
3	2	186	Dummy-Standard2				/Data/Global/Objects	N2	NU	HTH		LIT																	
3	2	187	Dummy-Torch1 Tiki				/Data/Global/Objects	TO	ON	HTH		LIT																	
3	2	188	Dummy-Torch2 Wall				/Data/Global/Objects	WT	ON	HTH		LIT																	
3	2	189	fire-RogueBonfire				/Data/Global/Objects	RB	ON	HTH		LIT																	
3	2	190	Dummy-River1				/Data/Global/Objects	R1	NU	HTH		LIT																	
3	2	191	Dummy-River2				/Data/Global/Objects	R2	NU	HTH		LIT																	
3	2	192	Dummy-River3				/Data/Global/Objects	R3	NU	HTH		LIT																	
3	2	193	Dummy-River4				/Data/Global/Objects	R4	NU	HTH		LIT																	
3	2	194	Dummy-River5				/Data/Global/Objects	R5	NU	HTH		LIT																	
3	2	195	AmbientSound-ambient sound generator				/Data/Global/Objects	S1	OP	HTH		LIT																	
3	2	196	Crate-Crate				/Data/Global/Objects	CT	OP	HTH		LIT																	
3	2	197	Door-Andariel's Door				/Data/Global/Objects	AD	NU	HTH		LIT																	
3	2	198	Dummy-RogueTorch				/Data/Global/Objects	T1	NU	HTH		LIT																	
3	2	199	Dummy-RogueTorch				/Data/Global/Objects	T2	NU	HTH		LIT																	
3	2	200	Casket-CasketR				/Data/Global/Objects	C1	OP	HTH		LIT																	
3	2	201	Casket-CasketL				/Data/Global/Objects	C2	OP	HTH		LIT																	
3	2	202	Urn-Urn #3				/Data/Global/Objects	U3	OP	HTH		LIT																	
3	2	203	Casket-Casket				/Data/Global/Objects	C4	OP	HTH		LIT																	
3	2	204	RogueCorpse-Rogue corpse 1				/Data/Global/Objects	Z1	NU	HTH		LIT																	
3	2	205	RogueCorpse-Rogue corpse 2				/Data/Global/Objects	Z2	NU	HTH		LIT																	
3	2	206	RogueCorpse-rolling rogue corpse				/Data/Global/Objects	Z5	OP	HTH		LIT																	
3	2	207	CorpseOnStick-rogue on a stick 1				/Data/Global/Objects	Z3	OP	HTH		LIT																	
3	2	208	CorpseOnStick-rogue on a stick 2				/Data/Global/Objects	Z4	OP	HTH		LIT																	
3	2	209	Portal-Town portal				/Data/Global/Objects	TP	ON	HTH	LIT	LIT																	
3	2	210	Portal-Permanent town portal				/Data/Global/Objects	PP	ON	HTH	LIT	LIT																	
3	2	211	Dummy-Invisible object				/Data/Global/Objects	SS																					
3	2	212	Door-Door Cathedral Left				/Data/Global/Objects	D9	OP	HTH		LIT																	
3	2	213	Door-Door Cathedral Right				/Data/Global/Objects	DA	OP	HTH		LIT																	
3	2	214	Door-Door Wooden Left #2				/Data/Global/Objects	DB	OP	HTH		LIT																	
3	2	215	Dummy-invisible river sound1				/Data/Global/Objects	X1																					
3	2	216	Dummy-invisible river sound2				/Data/Global/Objects	X2																					
3	2	217	Dummy-ripple				/Data/Global/Objects	1R	NU	HTH		LIT																	
3	2	218	Dummy-ripple				/Data/Global/Objects	2R	NU	HTH		LIT																	
3	2	219	Dummy-ripple				/Data/Global/Objects	3R	NU	HTH		LIT																	
3	2	220	Dummy-ripple				/Data/Global/Objects	4R	NU	HTH		LIT																	
3	2	221	Dummy-forest night sound #1				/Data/Global/Objects	F1																					
3	2	222	Dummy-forest night sound #2				/Data/Global/Objects	F2																					
3	2	223	Dummy-yeti dung				/Data/Global/Objects	YD	NU	HTH		LIT																	
3	2	224	TrappDoor-Trap Door				/Data/Global/Objects	TD	ON	HTH		LIT																	
3	2	225	Door-Door by Dock, Act 2				/Data/Global/Objects	DD	ON	HTH		LIT																	
3	2	226	Dummy-sewer drip				/Data/Global/Objects	SZ																					
3	2	227	Shrine-healthorama				/Data/Global/Objects	SH	OP	HTH		LIT																	
3	2	228	Dummy-invisible town sound				/Data/Global/Objects	TA																					
3	2	229	Casket-casket #3				/Data/Global/Objects	C3	OP	HTH		LIT																	
3	2	230	Obelisk-obelisk				/Data/Global/Objects	OB	OP	HTH		LIT																	
3	2	231	Shrine-forest altar				/Data/Global/Objects	AF	OP	HTH		LIT																	
3	2	232	Dummy-bubbling pool of blood				/Data/Global/Objects	B2	NU	HTH		LIT																	
3	2	233	Shrine-horn shrine				/Data/Global/Objects	HS	OP	HTH		LIT																	
3	2	234	Shrine-healing well				/Data/Global/Objects	HW	OP	HTH		LIT																	
3	2	235	Shrine-bull shrine,health, tombs				/Data/Global/Objects	BC	OP	HTH		LIT																	
3	2	236	Dummy-stele,magic shrine, stone, desert				/Data/Global/Objects	SG	OP	HTH		LIT																	
3	2	237	Chest3-tombchest 1, largechestL				/Data/Global/Objects	CA	OP	HTH		LIT																	
3	2	238	Chest3-tombchest 2 largechestR				/Data/Global/Objects	CB	OP	HTH		LIT																	
3	2	239	Sarcophagus-mummy coffinL, tomb				/Data/Global/Objects	MC	OP	HTH		LIT																	
3	2	240	Obelisk-desert obelisk				/Data/Global/Objects	DO	OP	HTH		LIT																	
3	2	241	Door-tomb door left				/Data/Global/Objects	TL	OP	HTH		LIT																	
3	2	242	Door-tomb door right				/Data/Global/Objects	TR	OP	HTH		LIT																	
3	2	243	Shrine-mana shrineforinnerhell				/Data/Global/Objects	iz	OP	HTH		LIT							LIT										
3	2	244	LargeUrn-Urn #4				/Data/Global/Objects	U4	OP	HTH		LIT																	
3	2	245	LargeUrn-Urn #5				/Data/Global/Objects	U5	OP	HTH		LIT																	
3	2	246	Shrine-health shrineforinnerhell				/Data/Global/Objects	iy	OP	HTH		LIT							LIT										
3	2	247	Shrine-innershrinehell				/Data/Global/Objects	ix	OP	HTH		LIT							LIT										
3	2	248	Door-tomb door left 2				/Data/Global/Objects	TS	OP	HTH		LIT																	
3	2	249	Door-tomb door right 2				/Data/Global/Objects	TU	OP	HTH		LIT																	
3	2	250	Duriel's Lair-Portal to Duriel's Lair				/Data/Global/Objects	SJ	OP	HTH		LIT																	
3	2	251	Dummy-Brazier3				/Data/Global/Objects	B3	OP	HTH		LIT							LIT										
3	2	252	Dummy-Floor brazier				/Data/Global/Objects	FB	ON	HTH		LIT							LIT										
3	2	253	Dummy-flies				/Data/Global/Objects	FL	NU	HTH		LIT																	
3	2	254	ArmorStand-Armor Stand 1R				/Data/Global/Objects	A3	NU	HTH		LIT																	
3	2	255	ArmorStand-Armor Stand 2L				/Data/Global/Objects	A4	NU	HTH		LIT																	
3	2	256	WeaponRack-Weapon Rack 1R				/Data/Global/Objects	W1	NU	HTH		LIT																	
3	2	257	WeaponRack-Weapon Rack 2L				/Data/Global/Objects	W2	NU	HTH		LIT																	
3	2	258	Malus-Malus				/Data/Global/Objects	HM	NU	HTH		LIT																	
3	2	259	Shrine-palace shrine, healthR, harom, arcane Sanctuary				/Data/Global/Objects	P2	OP	HTH		LIT																	
3	2	260	not used-drinker				/Data/Global/Objects	n5	S1	HTH		LIT																	
3	2	261	well-Fountain 1				/Data/Global/Objects	F3	OP	HTH		LIT																	
3	2	262	not used-gesturer				/Data/Global/Objects	n6	S1	HTH		LIT																	
3	2	263	well-Fountain 2, well, desert, tomb				/Data/Global/Objects	F4	OP	HTH		LIT																	
3	2	264	not used-turner				/Data/Global/Objects	n7	S1	HTH		LIT																	
3	2	265	well-Fountain 3				/Data/Global/Objects	F5	OP	HTH		LIT																	
3	2	266	Shrine-snake woman, magic shrine, tomb, arcane sanctuary				/Data/Global/Objects	SN	OP	HTH		LIT							LIT										
3	2	267	Dummy-jungle torch				/Data/Global/Objects	JT	ON	HTH		LIT							LIT										
3	2	268	Well-Fountain 4				/Data/Global/Objects	F6	OP	HTH		LIT																	
3	2	269	Waypoint-waypoint portal				/Data/Global/Objects	wp	ON	HTH		LIT							LIT										
3	2	270	Dummy-healthshrine, act 3, dungeun				/Data/Global/Objects	dj	OP	HTH		LIT																	
3	2	271	jerhyn-placeholder #1				/Data/Global/Objects	ss																					
3	2	272	jerhyn-placeholder #2				/Data/Global/Objects	ss																					
3	2	273	Shrine-innershrinehell2				/Data/Global/Objects	iw	OP	HTH		LIT							LIT										
3	2	274	Shrine-innershrinehell3				/Data/Global/Objects	iv	OP	HTH		LIT																	
3	2	275	hidden stash-ihobject3 inner hell				/Data/Global/Objects	iu	OP	HTH		LIT																	
3	2	276	skull pile-skullpile inner hell				/Data/Global/Objects	is	OP	HTH		LIT																	
3	2	277	hidden stash-ihobject5 inner hell				/Data/Global/Objects	ir	OP	HTH		LIT																	
3	2	278	hidden stash-hobject4 inner hell				/Data/Global/Objects	hg	OP	HTH		LIT																	
3	2	279	Door-secret door 1				/Data/Global/Objects	h2	OP	HTH		LIT																	
3	2	280	Well-pool act 1 wilderness				/Data/Global/Objects	zw	NU	HTH		LIT																	
3	2	281	Dummy-vile dog afterglow				/Data/Global/Objects	9b	OP	HTH		LIT																	
3	2	282	Well-cathedralwell act 1 inside				/Data/Global/Objects	zc	NU	HTH		LIT																	
3	2	283	shrine-shrine1_arcane sanctuary				/Data/Global/Objects	xx																					
3	2	284	shrine-dshrine2 act 2 shrine				/Data/Global/Objects	zs	OP	HTH		LIT							LIT										
3	2	285	shrine-desertshrine3 act 2 shrine				/Data/Global/Objects	zr	OP	HTH		LIT																	
3	2	286	shrine-dshrine1 act 2 shrine				/Data/Global/Objects	zd	OP	HTH		LIT																	
3	2	287	Well-desertwell act 2 well, desert, tomb				/Data/Global/Objects	zl	NU	HTH		LIT																	
3	2	288	Well-cavewell act 1 caves 				/Data/Global/Objects	zy	NU	HTH		LIT																	
3	2	289	chest-chest-r-large act 1				/Data/Global/Objects	q1	OP	HTH		LIT																	
3	2	290	chest-chest-r-tallskinney act 1				/Data/Global/Objects	q2	OP	HTH		LIT																	
3	2	291	chest-chest-r-med act 1				/Data/Global/Objects	q3	OP	HTH		LIT																	
3	2	292	jug-jug1 act 2, desert				/Data/Global/Objects	q4	OP	HTH		LIT																	
3	2	293	jug-jug2 act 2, desert				/Data/Global/Objects	q5	OP	HTH		LIT																	
3	2	294	chest-Lchest1 act 1				/Data/Global/Objects	q6	OP	HTH		LIT																	
3	2	295	Waypoint-waypointi inner hell				/Data/Global/Objects	wi	ON	HTH		LIT							LIT										
3	2	296	chest-dchest2R act 2, desert, tomb, chest-r-med				/Data/Global/Objects	q9	OP	HTH		LIT																	
3	2	297	chest-dchestr act 2, desert, tomb, chest -r large				/Data/Global/Objects	q7	OP	HTH		LIT																	
3	2	298	chest-dchestL act 2, desert, tomb chest l large				/Data/Global/Objects	q8	OP	HTH		LIT																	
3	2	299	taintedsunaltar-tainted sun altar quest				/Data/Global/Objects	za	OP	HTH		LIT							LIT										
3	2	300	shrine-dshrine1 act 2 , desert				/Data/Global/Objects	zv	NU	HTH		LIT							LIT	LIT									
3	2	301	shrine-dshrine4 act 2, desert				/Data/Global/Objects	ze	OP	HTH		LIT							LIT										
3	2	302	orifice-Where you place the Horadric staff				/Data/Global/Objects	HA	NU	HTH		LIT																	
3	2	303	Door-tyrael's door				/Data/Global/Objects	DX	OP	HTH		LIT																	
3	2	304	corpse-guard corpse				/Data/Global/Objects	GC	OP	HTH		LIT																	
3	2	305	hidden stash-rock act 1 wilderness				/Data/Global/Objects	c7	OP	HTH		LIT																	
3	2	306	Waypoint-waypoint act 2				/Data/Global/Objects	wm	ON	HTH		LIT							LIT										
3	2	307	Waypoint-waypoint act 1 wilderness				/Data/Global/Objects	wn	ON	HTH		LIT							LIT										
3	2	308	skeleton-corpse				/Data/Global/Objects	cp	OP	HTH		LIT																	
3	2	309	hidden stash-rockb act 1 wilderness				/Data/Global/Objects	cq	OP	HTH		LIT																	
3	2	310	fire-fire small				/Data/Global/Objects	FX	NU	HTH		LIT																	
3	2	311	fire-fire medium				/Data/Global/Objects	FY	NU	HTH		LIT																	
3	2	312	fire-fire large				/Data/Global/Objects	FZ	NU	HTH		LIT																	
3	2	313	hiding spot-cliff act 1 wilderness				/Data/Global/Objects	cf	NU	HTH		LIT																	
3	2	314	Shrine-mana well1				/Data/Global/Objects	MB	OP	HTH		LIT																	
3	2	315	Shrine-mana well2				/Data/Global/Objects	MD	OP	HTH		LIT																	
3	2	316	Shrine-mana well3, act 2, tomb, 				/Data/Global/Objects	MF	OP	HTH		LIT																	
3	2	317	Shrine-mana well4, act 2, harom				/Data/Global/Objects	MH	OP	HTH		LIT																	
3	2	318	Shrine-mana well5				/Data/Global/Objects	MJ	OP	HTH		LIT																	
3	2	319	hollow log-log				/Data/Global/Objects	cz	NU	HTH		LIT																	
3	2	320	Shrine-jungle healwell act 3				/Data/Global/Objects	JH	OP	HTH		LIT																	
3	2	321	skeleton-corpseb				/Data/Global/Objects	sx	OP	HTH		LIT																	
3	2	322	Shrine-health well, health shrine, desert				/Data/Global/Objects	Mk	OP	HTH		LIT																	
3	2	323	Shrine-mana well7, mana shrine, desert				/Data/Global/Objects	Mi	OP	HTH		LIT																	
3	2	324	loose rock-rockc act 1 wilderness				/Data/Global/Objects	RY	OP	HTH		LIT																	
3	2	325	loose boulder-rockd act 1 wilderness				/Data/Global/Objects	RZ	OP	HTH		LIT																	
3	2	326	chest-chest-L-med				/Data/Global/Objects	c8	OP	HTH		LIT																	
3	2	327	chest-chest-L-large				/Data/Global/Objects	c9	OP	HTH		LIT																	
3	2	328	GuardCorpse-guard on a stick, desert, tomb, harom				/Data/Global/Objects	GS	OP	HTH		LIT																	
3	2	329	bookshelf-bookshelf1				/Data/Global/Objects	b4	OP	HTH		LIT																	
3	2	330	bookshelf-bookshelf2				/Data/Global/Objects	b5	OP	HTH		LIT																	
3	2	331	chest-jungle chest act 3				/Data/Global/Objects	JC	OP	HTH		LIT																	
3	2	332	coffin-tombcoffin				/Data/Global/Objects	tm	OP	HTH		LIT																	
3	2	333	chest-chest-L-med, jungle				/Data/Global/Objects	jz	OP	HTH		LIT																	
3	2	334	Shrine-jungle shrine2				/Data/Global/Objects	jy	OP	HTH		LIT							LIT	LIT									
3	2	335	stash-jungle object act3				/Data/Global/Objects	jx	OP	HTH		LIT																	
3	2	336	stash-jungle object act3				/Data/Global/Objects	jw	OP	HTH		LIT																	
3	2	337	stash-jungle object act3				/Data/Global/Objects	jv	OP	HTH		LIT																	
3	2	338	stash-jungle object act3				/Data/Global/Objects	ju	OP	HTH		LIT																	
3	2	339	Dummy-cain portal				/Data/Global/Objects	tP	OP	HTH	LIT	LIT																	
3	2	340	Shrine-jungle shrine3 act 3				/Data/Global/Objects	js	OP	HTH		LIT							LIT										
3	2	341	Shrine-jungle shrine4 act 3				/Data/Global/Objects	jr	OP	HTH		LIT							LIT										
3	2	342	teleport pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
3	2	343	LamTome-Lam Esen's Tome				/Data/Global/Objects	ab	OP	HTH		LIT																	
3	2	344	stair-stairsl				/Data/Global/Objects	sl																					
3	2	345	stair-stairsr				/Data/Global/Objects	sv																					
3	2	346	a trap-test data floortrap				/Data/Global/Objects	a5	OP	HTH		LIT																	
3	2	347	Shrine-jungleshrine act 3				/Data/Global/Objects	jq	OP	HTH		LIT							LIT										
3	2	348	chest-chest-L-tallskinney, general chest r?				/Data/Global/Objects	c0	OP	HTH		LIT																	
3	2	349	Shrine-mafistoshrine				/Data/Global/Objects	mz	OP	HTH		LIT							LIT										
3	2	350	Shrine-mafistoshrine				/Data/Global/Objects	my	OP	HTH		LIT							LIT										
3	2	351	Shrine-mafistoshrine				/Data/Global/Objects	mx	NU	HTH		LIT							LIT										
3	2	352	Shrine-mafistomana				/Data/Global/Objects	mw	OP	HTH		LIT							LIT										
3	2	353	stash-mafistolair				/Data/Global/Objects	mv	OP	HTH		LIT																	
3	2	354	stash-box				/Data/Global/Objects	mu	OP	HTH		LIT																	
3	2	355	stash-altar				/Data/Global/Objects	mt	OP	HTH		LIT																	
3	2	356	Shrine-mafistohealth				/Data/Global/Objects	mr	OP	HTH		LIT							LIT										
3	2	357	dummy-water rocks in act 3 wrok				/Data/Global/Objects	rw	NU	HTH		LIT																	
3	2	358	Basket-basket 1				/Data/Global/Objects	bd	OP	HTH		LIT																	
3	2	359	Basket-basket 2				/Data/Global/Objects	bj	OP	HTH		LIT																	
3	2	360	Dummy-water logs in act 3  ne logw				/Data/Global/Objects	lw	NU	HTH		LIT																	
3	2	361	Dummy-water rocks girl in act 3 wrob				/Data/Global/Objects	wb	NU	HTH		LIT																	
3	2	362	Dummy-bubbles in act3 water				/Data/Global/Objects	yb	NU	HTH		LIT																	
3	2	363	Dummy-water logs in act 3 logx				/Data/Global/Objects	wd	NU	HTH		LIT																	
3	2	364	Dummy-water rocks in act 3 rokb				/Data/Global/Objects	wc	NU	HTH		LIT																	
3	2	365	Dummy-water rocks girl in act 3 watc				/Data/Global/Objects	we	NU	HTH		LIT																	
3	2	366	Dummy-water rocks in act 3 waty				/Data/Global/Objects	wy	NU	HTH		LIT																	
3	2	367	Dummy-water logs in act 3  logz				/Data/Global/Objects	lx	NU	HTH		LIT																	
3	2	368	Dummy-web covered tree 1				/Data/Global/Objects	w3	NU	HTH		LIT							LIT										
3	2	369	Dummy-web covered tree 2				/Data/Global/Objects	w4	NU	HTH		LIT							LIT										
3	2	370	Dummy-web covered tree 3				/Data/Global/Objects	w5	NU	HTH		LIT							LIT										
3	2	371	Dummy-web covered tree 4				/Data/Global/Objects	w6	NU	HTH		LIT							LIT										
3	2	372	pillar-hobject1				/Data/Global/Objects	70	OP	HTH		LIT																	
3	2	373	cocoon-cacoon				/Data/Global/Objects	CN	OP	HTH		LIT																	
3	2	374	cocoon-cacoon 2				/Data/Global/Objects	CC	OP	HTH		LIT																	
3	2	375	skullpile-hobject1				/Data/Global/Objects	ib	OP	HTH		LIT																	
3	2	376	Shrine-outershrinehell				/Data/Global/Objects	ia	OP	HTH		LIT							LIT										
3	2	377	dummy-water rock girl act 3  nw  blgb				/Data/Global/Objects	QX	NU	HTH		LIT																	
3	2	378	dummy-big log act 3  sw blga				/Data/Global/Objects	qw	NU	HTH		LIT																	
3	2	379	door-slimedoor1				/Data/Global/Objects	SQ	OP	HTH		LIT																	
3	2	380	door-slimedoor2				/Data/Global/Objects	SY	OP	HTH		LIT																	
3	2	381	Shrine-outershrinehell2				/Data/Global/Objects	ht	OP	HTH		LIT							LIT										
3	2	382	Shrine-outershrinehell3				/Data/Global/Objects	hq	OP	HTH		LIT																	
3	2	383	pillar-hobject2				/Data/Global/Objects	hv	OP	HTH		LIT																	
3	2	384	dummy-Big log act 3 se blgc 				/Data/Global/Objects	Qy	NU	HTH		LIT																	
3	2	385	dummy-Big log act 3 nw blgd				/Data/Global/Objects	Qz	NU	HTH		LIT																	
3	2	386	Shrine-health wellforhell				/Data/Global/Objects	ho	OP	HTH		LIT																	
3	2	387	Waypoint-act3waypoint town				/Data/Global/Objects	wz	ON	HTH		LIT							LIT										
3	2	388	Waypoint-waypointh				/Data/Global/Objects	wv	ON	HTH		LIT							LIT										
3	2	389	body-burning town				/Data/Global/Objects	bz	ON	HTH		LIT							LIT										
3	2	390	chest-gchest1L general				/Data/Global/Objects	cy	OP	HTH		LIT																	
3	2	391	chest-gchest2R general				/Data/Global/Objects	cx	OP	HTH		LIT																	
3	2	392	chest-gchest3R general				/Data/Global/Objects	cu	OP	HTH		LIT																	
3	2	393	chest-glchest3L general				/Data/Global/Objects	cd	OP	HTH		LIT																	
3	2	394	ratnest-sewers				/Data/Global/Objects	rn	OP	HTH		LIT																	
3	2	395	body-burning town				/Data/Global/Objects	by	NU	HTH		LIT							LIT										
3	2	396	ratnest-sewers				/Data/Global/Objects	ra	OP	HTH		LIT																	
3	2	397	bed-bed act 1				/Data/Global/Objects	qa	OP	HTH		LIT																	
3	2	398	bed-bed act 1				/Data/Global/Objects	qb	OP	HTH		LIT																	
3	2	399	manashrine-mana wellforhell				/Data/Global/Objects	hn	OP	HTH		LIT							LIT										
3	2	400	a trap-exploding cow  for Tristan and ACT 3 only??Very Rare  1 or 2				/Data/Global/Objects	ew	OP	HTH		LIT																	
3	2	401	gidbinn altar-gidbinn altar				/Data/Global/Objects	ga	ON	HTH		LIT							LIT										
3	2	402	gidbinn-gidbinn decoy				/Data/Global/Objects	gd	ON	HTH		LIT							LIT										
3	2	403	Dummy-diablo right light				/Data/Global/Objects	11	NU	HTH		LIT																	
3	2	404	Dummy-diablo left light				/Data/Global/Objects	12	NU	HTH		LIT																	
3	2	405	Dummy-diablo start point				/Data/Global/Objects	ss																					
3	2	406	Dummy-stool for act 1 cabin				/Data/Global/Objects	s9	NU	HTH		LIT																	
3	2	407	Dummy-wood for act 1 cabin				/Data/Global/Objects	wg	NU	HTH		LIT																	
3	2	408	Dummy-more wood for act 1 cabin				/Data/Global/Objects	wh	NU	HTH		LIT																	
3	2	409	Dummy-skeleton spawn for hell   facing nw				/Data/Global/Objects	QS	OP	HTH		LIT							LIT										
3	2	410	Shrine-holyshrine for monastery,catacombs,jail				/Data/Global/Objects	HL	OP	HTH		LIT							LIT										
3	2	411	a trap-spikes for tombs floortrap				/Data/Global/Objects	A7	OP	HTH		LIT																	
3	2	412	Shrine-act 1 cathedral				/Data/Global/Objects	s0	OP	HTH		LIT							LIT										
3	2	413	Shrine-act 1 jail				/Data/Global/Objects	jb	NU	HTH		LIT							LIT										
3	2	414	Shrine-act 1 jail				/Data/Global/Objects	jd	OP	HTH		LIT							LIT										
3	2	415	Shrine-act 1 jail				/Data/Global/Objects	jf	OP	HTH		LIT							LIT										
3	2	416	goo pile-goo pile for sand maggot lair				/Data/Global/Objects	GP	OP	HTH		LIT																	
3	2	417	bank-bank				/Data/Global/Objects	b6	NU	HTH		LIT																	
3	2	418	wirt's body-wirt's body				/Data/Global/Objects	BP	NU	HTH		LIT																	
3	2	419	dummy-gold placeholder				/Data/Global/Objects	1g																					
3	2	420	corpse-guard corpse 2				/Data/Global/Objects	GF	OP	HTH		LIT																	
3	2	421	corpse-dead villager 1				/Data/Global/Objects	dg	OP	HTH		LIT																	
3	2	422	corpse-dead villager 2				/Data/Global/Objects	df	OP	HTH		LIT																	
3	2	423	Dummy-yet another flame, no damage				/Data/Global/Objects	f8	NU	HTH		LIT																	
3	2	424	hidden stash-tiny pixel shaped thingie				/Data/Global/Objects	f9																					
3	2	425	Shrine-health shrine for caves				/Data/Global/Objects	ce	OP	HTH		LIT																	
3	2	426	Shrine-mana shrine for caves				/Data/Global/Objects	cg	OP	HTH		LIT																	
3	2	427	Shrine-cave magic shrine				/Data/Global/Objects	cg	OP	HTH		LIT																	
3	2	428	Shrine-manashrine, act 3, dungeun				/Data/Global/Objects	de	OP	HTH		LIT																	
3	2	429	Shrine-magic shrine, act 3 sewers.				/Data/Global/Objects	wj	NU	HTH		LIT							LIT	LIT									
3	2	430	Shrine-healthwell, act 3, sewers				/Data/Global/Objects	wk	OP	HTH		LIT																	
3	2	431	Shrine-manawell, act 3, sewers				/Data/Global/Objects	wl	OP	HTH		LIT																	
3	2	432	Shrine-magic shrine, act 3 sewers, dungeon.				/Data/Global/Objects	ws	NU	HTH		LIT							LIT	LIT									
3	2	433	dummy-brazier_celler, act 2				/Data/Global/Objects	bi	NU	HTH		LIT							LIT										
3	2	434	sarcophagus-anubis coffin, act2, tomb				/Data/Global/Objects	qc	OP	HTH		LIT																	
3	2	435	dummy-brazier_general, act 2, sewers, tomb, desert				/Data/Global/Objects	bm	NU	HTH		LIT							LIT										
3	2	436	Dummy-brazier_tall, act 2, desert, town, tombs				/Data/Global/Objects	bo	NU	HTH		LIT							LIT										
3	2	437	Dummy-brazier_small, act 2, desert, town, tombs				/Data/Global/Objects	bq	NU	HTH		LIT							LIT										
3	2	438	Waypoint-waypoint, celler				/Data/Global/Objects	w7	ON	HTH		LIT							LIT										
3	2	439	bed-bed for harum				/Data/Global/Objects	ub	OP	HTH		LIT																	
3	2	440	door-iron grate door left				/Data/Global/Objects	dv	NU	HTH		LIT																	
3	2	441	door-iron grate door right				/Data/Global/Objects	dn	NU	HTH		LIT																	
3	2	442	door-wooden grate door left				/Data/Global/Objects	dp	NU	HTH		LIT																	
3	2	443	door-wooden grate door right				/Data/Global/Objects	dt	NU	HTH		LIT																	
3	2	444	door-wooden door left				/Data/Global/Objects	dk	NU	HTH		LIT																	
3	2	445	door-wooden door right				/Data/Global/Objects	dl	NU	HTH		LIT																	
3	2	446	Dummy-wall torch left for tombs				/Data/Global/Objects	qd	NU	HTH		LIT							LIT										
3	2	447	Dummy-wall torch right for tombs				/Data/Global/Objects	qe	NU	HTH		LIT							LIT										
3	2	448	portal-arcane sanctuary portal				/Data/Global/Objects	ay	ON	HTH		LIT							LIT	LIT									
3	2	449	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hb	OP	HTH		LIT							LIT										
3	2	450	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hc	OP	HTH		LIT							LIT										
3	2	451	Dummy-maggot well health				/Data/Global/Objects	qf	OP	HTH		LIT																	
3	2	452	manashrine-maggot well mana				/Data/Global/Objects	qg	OP	HTH		LIT																	
3	2	453	magic shrine-magic shrine, act 3 arcane sanctuary.				/Data/Global/Objects	hd	OP	HTH		LIT							LIT										
3	2	454	teleportation pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
3	2	455	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
3	2	456	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
3	2	457	Dummy-arcane thing				/Data/Global/Objects	7a	NU	HTH		LIT																	
3	2	458	Dummy-arcane thing				/Data/Global/Objects	7b	NU	HTH		LIT																	
3	2	459	Dummy-arcane thing				/Data/Global/Objects	7c	NU	HTH		LIT																	
3	2	460	Dummy-arcane thing				/Data/Global/Objects	7d	NU	HTH		LIT																	
3	2	461	Dummy-arcane thing				/Data/Global/Objects	7e	NU	HTH		LIT																	
3	2	462	Dummy-arcane thing				/Data/Global/Objects	7f	NU	HTH		LIT																	
3	2	463	Dummy-arcane thing				/Data/Global/Objects	7g	NU	HTH		LIT																	
3	2	464	dead guard-harem guard 1				/Data/Global/Objects	qh	NU	HTH		LIT																	
3	2	465	dead guard-harem guard 2				/Data/Global/Objects	qi	NU	HTH		LIT																	
3	2	466	dead guard-harem guard 3				/Data/Global/Objects	qj	NU	HTH		LIT																	
3	2	467	dead guard-harem guard 4				/Data/Global/Objects	qk	NU	HTH		LIT																	
3	2	468	eunuch-harem blocker				/Data/Global/Objects	ss																					
3	2	469	Dummy-healthwell, act 2, arcane				/Data/Global/Objects	ax	OP	HTH		LIT																	
3	2	470	manashrine-healthwell, act 2, arcane				/Data/Global/Objects	au	OP	HTH		LIT																	
3	2	471	Dummy-test data				/Data/Global/Objects	pp	S1	HTH	LIT	LIT																	
3	2	472	Well-tombwell act 2 well, tomb				/Data/Global/Objects	hu	NU	HTH		LIT																	
3	2	473	Waypoint-waypoint act2 sewer				/Data/Global/Objects	qm	ON	HTH		LIT							LIT										
3	2	474	Waypoint-waypoint act3 travincal				/Data/Global/Objects	ql	ON	HTH		LIT							LIT										
3	2	475	magic shrine-magic shrine, act 3, sewer				/Data/Global/Objects	qn	NU	HTH		LIT							LIT										
3	2	476	dead body-act3, sewer				/Data/Global/Objects	qo	OP	HTH		LIT																	
3	2	477	dummy-torch (act 3 sewer) stra				/Data/Global/Objects	V1	NU	HTH		LIT							LIT										
3	2	478	dummy-torch (act 3 kurast) strb				/Data/Global/Objects	V2	NU	HTH		LIT							LIT										
3	2	479	chest-mafistochestlargeLeft				/Data/Global/Objects	xb	OP	HTH		LIT																	
3	2	480	chest-mafistochestlargeright				/Data/Global/Objects	xc	OP	HTH		LIT																	
3	2	481	chest-mafistochestmedleft				/Data/Global/Objects	xd	OP	HTH		LIT																	
3	2	482	chest-mafistochestmedright				/Data/Global/Objects	xe	OP	HTH		LIT																	
3	2	483	chest-spiderlairchestlargeLeft				/Data/Global/Objects	xf	OP	HTH		LIT																	
3	2	484	chest-spiderlairchesttallLeft				/Data/Global/Objects	xg	OP	HTH		LIT																	
3	2	485	chest-spiderlairchestmedright				/Data/Global/Objects	xh	OP	HTH		LIT																	
3	2	486	chest-spiderlairchesttallright				/Data/Global/Objects	xi	OP	HTH		LIT																	
3	2	487	Steeg Stone-steeg stone				/Data/Global/Objects	y6	NU	HTH		LIT							LIT										
3	2	488	Guild Vault-guild vault				/Data/Global/Objects	y4	NU	HTH		LIT																	
3	2	489	Trophy Case-trophy case				/Data/Global/Objects	y2	NU	HTH		LIT																	
3	2	490	Message Board-message board				/Data/Global/Objects	y3	NU	HTH		LIT																	
3	2	491	Dummy-mephisto bridge				/Data/Global/Objects	xj	OP	HTH		LIT																	
3	2	492	portal-hellgate				/Data/Global/Objects	1y	ON	HTH		LIT								LIT	LIT								
3	2	493	Shrine-manawell, act 3, kurast				/Data/Global/Objects	xl	OP	HTH		LIT																	
3	2	494	Shrine-healthwell, act 3, kurast				/Data/Global/Objects	xm	OP	HTH		LIT																	
3	2	495	Dummy-hellfire1				/Data/Global/Objects	e3	NU	HTH		LIT																	
3	2	496	Dummy-hellfire2				/Data/Global/Objects	e4	NU	HTH		LIT																	
3	2	497	Dummy-hellfire3				/Data/Global/Objects	e5	NU	HTH		LIT																	
3	2	498	Dummy-helllava1				/Data/Global/Objects	e6	NU	HTH		LIT																	
3	2	499	Dummy-helllava2				/Data/Global/Objects	e7	NU	HTH		LIT																	
3	2	500	Dummy-helllava3				/Data/Global/Objects	e8	NU	HTH		LIT																	
3	2	501	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
3	2	502	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
3	2	503	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
3	2	504	chest-horadric cube chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
3	2	505	chest-horadric scroll chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
3	2	506	chest-staff of kings chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
3	2	507	Tome-yet another tome				/Data/Global/Objects	TT	NU	HTH		LIT																	
3	2	508	fire-hell brazier				/Data/Global/Objects	E1	NU	HTH	LIT	LIT																	
3	2	509	fire-hell brazier				/Data/Global/Objects	E2	NU	HTH	LIT	LIT																	
3	2	510	RockPIle-dungeon				/Data/Global/Objects	xn	OP	HTH		LIT																	
3	2	511	magic shrine-magic shrine, act 3,dundeon				/Data/Global/Objects	qo	OP	HTH		LIT																	
3	2	512	basket-dungeon				/Data/Global/Objects	xp	OP	HTH		LIT																	
3	2	513	HungSkeleton-outerhell skeleton				/Data/Global/Objects	jw	OP	HTH		LIT																	
3	2	514	Dummy-guy for dungeon				/Data/Global/Objects	ea	OP	HTH		LIT																	
3	2	515	casket-casket for Act 3 dungeon				/Data/Global/Objects	vb	OP	HTH		LIT																	
3	2	516	sewer stairs-stairs for act 3 sewer quest				/Data/Global/Objects	ve	OP	HTH		LIT																	
3	2	517	sewer lever-lever for act 3 sewer quest				/Data/Global/Objects	vf	OP	HTH		LIT																	
3	2	518	darkwanderer-start position				/Data/Global/Objects	ss																					
3	2	519	dummy-trapped soul placeholder				/Data/Global/Objects	ss																					
3	2	520	Dummy-torch for act3 town				/Data/Global/Objects	VG	NU	HTH		LIT							LIT										
3	2	521	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
3	2	522	BoneChest-innerhellbonepile				/Data/Global/Objects	y1	OP	HTH		LIT																	
3	2	523	Dummy-skeleton spawn for hell facing ne				/Data/Global/Objects	Qt	OP	HTH		LIT							LIT										
3	2	524	Dummy-fog act 3 water rfga				/Data/Global/Objects	ud	NU	HTH		LIT																	
3	2	525	Dummy-Not used				/Data/Global/Objects	xx																					
3	2	526	Hellforge-Forge  hell				/Data/Global/Objects	ux	ON	HTH		LIT							LIT	LIT	LIT								
3	2	527	Guild Portal-Portal to next guild level				/Data/Global/Objects	PP	NU	HTH	LIT	LIT																	
3	2	528	Dummy-hratli start				/Data/Global/Objects	ss																					
3	2	529	Dummy-hratli end				/Data/Global/Objects	ss																					
3	2	530	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	uy	OP	HTH		LIT							LIT										
3	2	531	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	15	OP	HTH		LIT							LIT										
3	2	532	Dummy-natalya start				/Data/Global/Objects	ss																					
3	2	533	TrappedSoul-guy stuck in hell				/Data/Global/Objects	18	OP	HTH		LIT																	
3	2	534	TrappedSoul-guy stuck in hell				/Data/Global/Objects	19	OP	HTH		LIT																	
3	2	535	Dummy-cain start position				/Data/Global/Objects	ss																					
3	2	536	Dummy-stairsr				/Data/Global/Objects	sv	OP	HTH		LIT																	
3	2	537	chest-arcanesanctuarybigchestLeft				/Data/Global/Objects	y7	OP	HTH		LIT																	
3	2	538	casket-arcanesanctuarycasket				/Data/Global/Objects	y8	OP	HTH		LIT																	
3	2	539	chest-arcanesanctuarybigchestRight				/Data/Global/Objects	y9	OP	HTH		LIT																	
3	2	540	chest-arcanesanctuarychestsmallLeft				/Data/Global/Objects	ya	OP	HTH		LIT																	
3	2	541	chest-arcanesanctuarychestsmallRight				/Data/Global/Objects	yc	OP	HTH		LIT																	
3	2	542	Seal-Diablo seal				/Data/Global/Objects	30	ON	HTH		LIT							LIT										
3	2	543	Seal-Diablo seal				/Data/Global/Objects	31	ON	HTH		LIT							LIT										
3	2	544	Seal-Diablo seal				/Data/Global/Objects	32	ON	HTH		LIT							LIT										
3	2	545	Seal-Diablo seal				/Data/Global/Objects	33	ON	HTH		LIT							LIT										
3	2	546	Seal-Diablo seal				/Data/Global/Objects	34	ON	HTH		LIT							LIT										
3	2	547	chest-sparklychest				/Data/Global/Objects	yf	OP	HTH		LIT																	
3	2	548	Waypoint-waypoint pandamonia fortress				/Data/Global/Objects	yg	ON	HTH		LIT							LIT										
3	2	549	fissure-fissure for act 4 inner hell				/Data/Global/Objects	fh	OP	HTH		LIT							LIT										
3	2	550	Dummy-brazier for act 4, hell mesa				/Data/Global/Objects	he	NU	HTH		LIT							LIT										
3	2	551	Dummy-smoke				/Data/Global/Objects	35	NU	HTH		LIT																	
3	2	552	Waypoint-waypoint valleywaypoint				/Data/Global/Objects	yi	ON	HTH		LIT							LIT										
3	2	553	fire-hell brazier				/Data/Global/Objects	9f	NU	HTH		LIT							LIT										
3	2	554	compellingorb-compelling orb				/Data/Global/Objects	55	NU	HTH		LIT							LIT	LIT									
3	2	555	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
3	2	556	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
3	2	557	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
3	2	558	Dummy-fortress brazier #1				/Data/Global/Objects	98	NU	HTH		LIT							LIT										
3	2	559	Dummy-fortress brazier #2				/Data/Global/Objects	99	NU	HTH		LIT							LIT										
3	2	560	Siege Control-To control siege machines				/Data/Global/Objects	zq	OP	HTH		LIT																	
3	2	561	ptox-Pot O Torch (level 1)				/Data/Global/Objects	px	NU	HTH		LIT							LIT	LIT									
3	2	562	pyox-fire pit  (level 1)				/Data/Global/Objects	py	NU	HTH		LIT							LIT										
3	2	563	chestR-expansion no snow				/Data/Global/Objects	6q	OP	HTH		LIT																	
3	2	564	Shrine3wilderness-expansion no snow				/Data/Global/Objects	6r	OP	HTH		LIT							LIT										
3	2	565	Shrine2wilderness-expansion no snow				/Data/Global/Objects	6s	NU	HTH		LIT							LIT										
3	2	566	hiddenstash-expansion no snow				/Data/Global/Objects	3w	OP	HTH		LIT																	
3	2	567	flag wilderness-expansion no snow				/Data/Global/Objects	ym	NU	HTH		LIT																	
3	2	568	barrel wilderness-expansion no snow				/Data/Global/Objects	yn	OP	HTH		LIT																	
3	2	569	barrel wilderness-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
3	2	570	woodchestL-expansion no snow				/Data/Global/Objects	yp	OP	HTH		LIT																	
3	2	571	Shrine3wilderness-expansion no snow				/Data/Global/Objects	yq	NU	HTH		LIT							LIT										
3	2	572	manashrine-expansion no snow				/Data/Global/Objects	yr	OP	HTH		LIT							LIT										
3	2	573	healthshrine-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
3	2	574	burialchestL-expansion no snow				/Data/Global/Objects	yt	OP	HTH		LIT																	
3	2	575	burialchestR-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
3	2	576	well-expansion no snow				/Data/Global/Objects	yv	NU	HTH		LIT																	
3	2	577	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yw	OP	HTH		LIT							LIT	LIT									
3	2	578	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yx	OP	HTH		LIT							LIT										
3	2	579	Waypoint-expansion no snow				/Data/Global/Objects	yy	ON	HTH		LIT							LIT										
3	2	580	ChestL-expansion no snow				/Data/Global/Objects	yz	OP	HTH		LIT																	
3	2	581	woodchestR-expansion no snow				/Data/Global/Objects	6a	OP	HTH		LIT																	
3	2	582	ChestSL-expansion no snow				/Data/Global/Objects	6b	OP	HTH		LIT																	
3	2	583	ChestSR-expansion no snow				/Data/Global/Objects	6c	OP	HTH		LIT																	
3	2	584	etorch1-expansion no snow				/Data/Global/Objects	6d	NU	HTH		LIT							LIT										
3	2	585	ecfra-camp fire				/Data/Global/Objects	2w	NU	HTH		LIT							LIT	LIT									
3	2	586	ettr-town torch				/Data/Global/Objects	2x	NU	HTH		LIT							LIT	LIT									
3	2	587	etorch2-expansion no snow				/Data/Global/Objects	6e	NU	HTH		LIT							LIT										
3	2	588	burningbodies-wilderness/siege				/Data/Global/Objects	6f	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
3	2	589	burningpit-wilderness/siege				/Data/Global/Objects	6g	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
3	2	590	tribal flag-wilderness/siege				/Data/Global/Objects	6h	NU	HTH		LIT																	
3	2	591	eflg-town flag				/Data/Global/Objects	2y	NU	HTH		LIT																	
3	2	592	chan-chandeleir				/Data/Global/Objects	2z	NU	HTH		LIT							LIT										
3	2	593	jar1-wilderness/siege				/Data/Global/Objects	6i	OP	HTH		LIT																	
3	2	594	jar2-wilderness/siege				/Data/Global/Objects	6j	OP	HTH		LIT																	
3	2	595	jar3-wilderness/siege				/Data/Global/Objects	6k	OP	HTH		LIT																	
3	2	596	swingingheads-wilderness				/Data/Global/Objects	6L	NU	HTH		LIT																	
3	2	597	pole-wilderness				/Data/Global/Objects	6m	NU	HTH		LIT																	
3	2	598	animated skulland rockpile-expansion no snow				/Data/Global/Objects	6n	OP	HTH		LIT																	
3	2	599	gate-town main gate				/Data/Global/Objects	2v	OP	HTH		LIT																	
3	2	600	pileofskullsandrocks-seige				/Data/Global/Objects	6o	NU	HTH		LIT																	
3	2	601	hellgate-seige				/Data/Global/Objects	6p	NU	HTH		LIT							LIT	LIT									
3	2	602	banner 1-preset in enemy camp				/Data/Global/Objects	ao	NU	HTH		LIT																	
3	2	603	banner 2-preset in enemy camp				/Data/Global/Objects	ap	NU	HTH		LIT																	
3	2	604	explodingchest-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
3	2	605	chest-specialchest				/Data/Global/Objects	6u	OP	HTH		LIT																	
3	2	606	deathpole-wilderness				/Data/Global/Objects	6v	NU	HTH		LIT																	
3	2	607	Ldeathpole-wilderness				/Data/Global/Objects	6w	NU	HTH		LIT																	
3	2	608	Altar-inside of temple				/Data/Global/Objects	6x	NU	HTH		LIT							LIT										
3	2	609	dummy-Drehya Start In Town				/Data/Global/Objects	ss																					
3	2	610	dummy-Drehya Start Outside Town				/Data/Global/Objects	ss																					
3	2	611	dummy-Nihlathak Start In Town				/Data/Global/Objects	ss																					
3	2	612	dummy-Nihlathak Start Outside Town				/Data/Global/Objects	ss																					
3	2	613	hidden stash-icecave_				/Data/Global/Objects	6y	OP	HTH		LIT																	
3	2	614	healthshrine-icecave_				/Data/Global/Objects	8a	OP	HTH		LIT																	
3	2	615	manashrine-icecave_				/Data/Global/Objects	8b	OP	HTH		LIT																	
3	2	616	evilurn-icecave_				/Data/Global/Objects	8c	OP	HTH		LIT																	
3	2	617	icecavejar1-icecave_				/Data/Global/Objects	8d	OP	HTH		LIT																	
3	2	618	icecavejar2-icecave_				/Data/Global/Objects	8e	OP	HTH		LIT																	
3	2	619	icecavejar3-icecave_				/Data/Global/Objects	8f	OP	HTH		LIT																	
3	2	620	icecavejar4-icecave_				/Data/Global/Objects	8g	OP	HTH		LIT																	
3	2	621	icecavejar4-icecave_				/Data/Global/Objects	8h	OP	HTH		LIT																	
3	2	622	icecaveshrine2-icecave_				/Data/Global/Objects	8i	NU	HTH		LIT							LIT										
3	2	623	cagedwussie1-caged fellow(A5-Prisonner)				/Data/Global/Objects	60	NU	HTH		LIT																	
3	2	624	Ancient Statue 3-statue				/Data/Global/Objects	60	NU	HTH		LIT																	
3	2	625	Ancient Statue 1-statue				/Data/Global/Objects	61	NU	HTH		LIT																	
3	2	626	Ancient Statue 2-statue				/Data/Global/Objects	62	NU	HTH		LIT																	
3	2	627	deadbarbarian-seige/wilderness				/Data/Global/Objects	8j	OP	HTH		LIT																	
3	2	628	clientsmoke-client smoke				/Data/Global/Objects	oz	NU	HTH		LIT																	
3	2	629	icecaveshrine2-icecave_				/Data/Global/Objects	8k	NU	HTH		LIT							LIT										
3	2	630	icecave_torch1-icecave_				/Data/Global/Objects	8L	NU	HTH		LIT							LIT										
3	2	631	icecave_torch2-icecave_				/Data/Global/Objects	8m	NU	HTH		LIT							LIT										
3	2	632	ttor-expansion tiki torch				/Data/Global/Objects	2p	NU	HTH		LIT							LIT										
3	2	633	manashrine-baals				/Data/Global/Objects	8n	OP	HTH		LIT																	
3	2	634	healthshrine-baals				/Data/Global/Objects	8o	OP	HTH		LIT																	
3	2	635	tomb1-baal's lair				/Data/Global/Objects	8p	OP	HTH		LIT																	
3	2	636	tomb2-baal's lair				/Data/Global/Objects	8q	OP	HTH		LIT																	
3	2	637	tomb3-baal's lair				/Data/Global/Objects	8r	OP	HTH		LIT																	
3	2	638	magic shrine-baal's lair				/Data/Global/Objects	8s	NU	HTH		LIT							LIT										
3	2	639	torch1-baal's lair				/Data/Global/Objects	8t	NU	HTH		LIT							LIT										
3	2	640	torch2-baal's lair				/Data/Global/Objects	8u	NU	HTH		LIT							LIT										
3	2	641	manashrine-snowy				/Data/Global/Objects	8v	OP	HTH		LIT							LIT										
3	2	642	healthshrine-snowy				/Data/Global/Objects	8w	OP	HTH		LIT							LIT										
3	2	643	well-snowy				/Data/Global/Objects	8x	NU	HTH		LIT																	
3	2	644	Waypoint-baals_waypoint				/Data/Global/Objects	8y	ON	HTH		LIT							LIT										
3	2	645	magic shrine-snowy_shrine3				/Data/Global/Objects	8z	NU	HTH		LIT							LIT										
3	2	646	Waypoint-wilderness_waypoint				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
3	2	647	magic shrine-snowy_shrine3				/Data/Global/Objects	5b	OP	HTH		LIT							LIT	LIT									
3	2	648	well-baalslair				/Data/Global/Objects	5c	NU	HTH		LIT																	
3	2	649	magic shrine2-baal's lair				/Data/Global/Objects	5d	NU	HTH		LIT							LIT										
3	2	650	object1-snowy				/Data/Global/Objects	5e	OP	HTH		LIT																	
3	2	651	woodchestL-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
3	2	652	woodchestR-snowy				/Data/Global/Objects	5g	OP	HTH		LIT																	
3	2	653	magic shrine-baals_shrine3				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
3	2	654	woodchest2L-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
3	2	655	woodchest2R-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
3	2	656	swingingheads-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
3	2	657	debris-snowy				/Data/Global/Objects	5l	NU	HTH		LIT																	
3	2	658	pene-Pen breakable door				/Data/Global/Objects	2q	NU	HTH		LIT																	
3	2	659	magic shrine-temple				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
3	2	660	mrpole-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
3	2	661	Waypoint-icecave 				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
3	2	662	magic shrine-temple				/Data/Global/Objects	5t	NU	HTH		LIT							LIT										
3	2	663	well-temple				/Data/Global/Objects	5q	NU	HTH		LIT																	
3	2	664	torch1-temple				/Data/Global/Objects	5r	NU	HTH		LIT							LIT										
3	2	665	torch1-temple				/Data/Global/Objects	5s	NU	HTH		LIT							LIT										
3	2	666	object1-temple				/Data/Global/Objects	5u	OP	HTH		LIT																	
3	2	667	object2-temple				/Data/Global/Objects	5v	OP	HTH		LIT																	
3	2	668	mrbox-baals				/Data/Global/Objects	5w	OP	HTH		LIT																	
3	2	669	well-icecave				/Data/Global/Objects	5x	NU	HTH		LIT																	
3	2	670	magic shrine-temple				/Data/Global/Objects	5y	NU	HTH		LIT							LIT										
3	2	671	healthshrine-temple				/Data/Global/Objects	5z	OP	HTH		LIT																	
3	2	672	manashrine-temple				/Data/Global/Objects	3a	OP	HTH		LIT																	
3	2	673	red light- (touch me)  for blacksmith				/Data/Global/Objects	ss																					
3	2	674	tomb1L-baal's lair				/Data/Global/Objects	3b	OP	HTH		LIT																	
3	2	675	tomb2L-baal's lair				/Data/Global/Objects	3c	OP	HTH		LIT																	
3	2	676	tomb3L-baal's lair				/Data/Global/Objects	3d	OP	HTH		LIT																	
3	2	677	ubub-Ice cave bubbles 01				/Data/Global/Objects	2u	NU	HTH		LIT																	
3	2	678	sbub-Ice cave bubbles 01				/Data/Global/Objects	2s	NU	HTH		LIT																	
3	2	679	tomb1-redbaal's lair				/Data/Global/Objects	3f	OP	HTH		LIT																	
3	2	680	tomb1L-redbaal's lair				/Data/Global/Objects	3g	OP	HTH		LIT																	
3	2	681	tomb2-redbaal's lair				/Data/Global/Objects	3h	OP	HTH		LIT																	
3	2	682	tomb2L-redbaal's lair				/Data/Global/Objects	3i	OP	HTH		LIT																	
3	2	683	tomb3-redbaal's lair				/Data/Global/Objects	3j	OP	HTH		LIT																	
3	2	684	tomb3L-redbaal's lair				/Data/Global/Objects	3k	OP	HTH		LIT																	
3	2	685	mrbox-redbaals				/Data/Global/Objects	3L	OP	HTH		LIT																	
3	2	686	torch1-redbaal's lair				/Data/Global/Objects	3m	NU	HTH		LIT							LIT										
3	2	687	torch2-redbaal's lair				/Data/Global/Objects	3n	NU	HTH		LIT							LIT										
3	2	688	candles-temple				/Data/Global/Objects	3o	NU	HTH		LIT							LIT										
3	2	689	Waypoint-temple				/Data/Global/Objects	3p	ON	HTH		LIT							LIT										
3	2	690	deadperson-everywhere				/Data/Global/Objects	3q	NU	HTH		LIT																	
3	2	691	groundtomb-temple				/Data/Global/Objects	3s	OP	HTH		LIT																	
3	2	692	Dummy-Larzuk Greeting				/Data/Global/Objects	ss																					
3	2	693	Dummy-Larzuk Standard				/Data/Global/Objects	ss																					
3	2	694	groundtombL-temple				/Data/Global/Objects	3t	OP	HTH		LIT																	
3	2	695	deadperson2-everywhere				/Data/Global/Objects	3u	OP	HTH		LIT																	
3	2	696	ancientsaltar-ancientsaltar				/Data/Global/Objects	4a	OP	HTH		LIT							LIT										
3	2	697	To The Worldstone Keep Level 1-ancientsdoor				/Data/Global/Objects	4b	OP	HTH		LIT																	
3	2	698	eweaponrackR-everywhere				/Data/Global/Objects	3x	NU	HTH		LIT																	
3	2	699	eweaponrackL-everywhere				/Data/Global/Objects	3y	NU	HTH		LIT																	
3	2	700	earmorstandR-everywhere				/Data/Global/Objects	3z	NU	HTH		LIT																	
3	2	701	earmorstandL-everywhere				/Data/Global/Objects	4c	NU	HTH		LIT																	
3	2	702	torch2-summit				/Data/Global/Objects	9g	NU	HTH		LIT							LIT										
3	2	703	funeralpire-outside				/Data/Global/Objects	9h	NU	HTH		LIT							LIT										
3	2	704	burninglogs-outside				/Data/Global/Objects	9i	NU	HTH		LIT							LIT										
3	2	705	stma-Ice cave steam				/Data/Global/Objects	2o	NU	HTH		LIT																	
3	2	706	deadperson2-everywhere				/Data/Global/Objects	3v	OP	HTH		LIT																	
3	2	707	Dummy-Baal's lair				/Data/Global/Objects	ss																					
3	2	708	fana-frozen anya				/Data/Global/Objects	2n	NU	HTH		LIT																	
3	2	709	BBQB-BBQ Bunny				/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									
3	2	710	btor-Baal Torch Big				/Data/Global/Objects	25	NU	HTH		LIT							LIT										
3	2	711	Dummy-invisible ancient				/Data/Global/Objects	ss																					
3	2	712	Dummy-invisible base				/Data/Global/Objects	ss																					
3	2	713	The Worldstone Chamber-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
3	2	714	Glacial Caves Level 1-summit door				/Data/Global/Objects	4u	OP	HTH		LIT																	
3	2	715	strlastcinematic-last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
3	2	716	Harrogath-last last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
3	2	717	Zoo-test data				/Data/Global/Objects	ss																					
3	2	718	Keeper-test data				/Data/Global/Objects	7z	NU	HTH		LIT																	
3	2	719	Throne of Destruction-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
3	2	720	Dummy-fire place guy				/Data/Global/Objects	7y	NU	HTH		LIT																	
3	2	721	Dummy-door blocker				/Data/Global/Objects	ss																					
3	2	722	Dummy-door blocker				/Data/Global/Objects	ss																					
4	1	0	place_champion-ACT 4 TABLE																										0
4	1	1	trap-horzmissile-ACT 4 TABLE																										0
4	1	2	trap-vertmissile-ACT 4 TABLE																										0
4	1	3	place_group25-ACT 4 TABLE																										0
4	1	4	place_group50-ACT 4 TABLE																										0
4	1	5	place_group75-ACT 4 TABLE																										0
4	1	6	place_group100-ACT 4 TABLE																										0
4	1	7	tyrael2-ACT 4 TABLE				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
4	1	8	window2-ACT 4 TABLE				/Data/Global/Monsters	VJ	DT	HTH		LIT							S1										0
4	1	9	window1-ACT 4 TABLE				/Data/Global/Monsters	VH	DT	HTH		LIT							S1										0
4	1	10	jamella-ACT 4 TABLE				/Data/Global/Monsters	JA	NU	HTH		LIT																	0
4	1	11	halbu-ACT 4 TABLE				/Data/Global/Monsters	20	NU	HTH		LIT																	0
4	1	12	hellmeteor-ACT 4 TABLE																										0
4	1	13	izual-ACT 4 TABLE			6	/Data/Global/Monsters	22	NU	HTH		LIT																	0
4	1	14	diablo-ACT 4 TABLE			1	/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
4	1	15	Winged Death-ACT 4 TABLE																										0
4	1	16	The Tormentor-ACT 4 TABLE																										0
4	1	17	Taintbreeder-ACT 4 TABLE																										0
4	1	18	Riftwraith the Cannibal-ACT 4 TABLE																										0
4	1	19	Infector of Souls-ACT 4 TABLE																										0
4	1	20	Lord De Seis-ACT 4 TABLE																										0
4	1	21	Grand Vizier of Chaos-ACT 4 TABLE																										0
4	1	22	trappedsoul1-ACT 4 TABLE				/Data/Global/Monsters	10	NU	HTH		LIT																	0
4	1	23	trappedsoul2-ACT 4 TABLE				/Data/Global/Monsters	13	S1	HTH		LIT																	0
4	1	24	regurgitator3-ACT 4 TABLE			6	/Data/Global/Monsters	CS	NU	HTH		LIT																	0
4	1	25	cain4-ACT 4 TABLE				/Data/Global/Monsters	4D	NU	HTH		LIT																	0
4	1	26	malachai-ACT 4 TABLE				/Data/Global/Monsters	36	NU	HTH		LIT							LIT										0
4	1	27	The Feature Creep-ACT 4 TABLE																										0
4	1	28	skeleton1-Skeleton-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	29	skeleton2-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	30	skeleton3-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	31	skeleton4-BurningDead-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	32	skeleton5-Horror-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	33	zombie1-Zombie-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	34	zombie2-HungryDead-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	35	zombie3-Ghoul-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	36	zombie4-DrownedCarcass-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	37	zombie5-PlagueBearer-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	38	bighead1-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	39	bighead2-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	40	bighead3-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	41	bighead4-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	42	bighead5-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	43	foulcrow1-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	44	foulcrow2-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	45	foulcrow3-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	46	foulcrow4-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	47	fallen1-Fallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
4	1	48	fallen2-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
4	1	49	fallen3-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
4	1	50	fallen4-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
4	1	51	fallen5-WarpedFallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
4	1	52	brute2-Brute-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
4	1	53	brute3-Yeti-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
4	1	54	brute4-Crusher-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
4	1	55	brute5-WailingBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
4	1	56	brute1-GargantuanBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
4	1	57	sandraider1-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	58	sandraider2-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	59	sandraider3-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	60	sandraider4-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	61	sandraider5-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	62	gorgon1-unused-Idle				/Data/Global/Monsters	GO																					0
4	1	63	gorgon2-unused-Idle				/Data/Global/Monsters	GO																					0
4	1	64	gorgon3-unused-Idle				/Data/Global/Monsters	GO																					0
4	1	65	gorgon4-unused-Idle				/Data/Global/Monsters	GO																					0
4	1	66	wraith1-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	67	wraith2-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	68	wraith3-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	69	wraith4-Apparition-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	70	wraith5-DarkShape-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	71	corruptrogue1-DarkHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	72	corruptrogue2-VileHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	73	corruptrogue3-DarkStalker-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	74	corruptrogue4-BlackRogue-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	75	corruptrogue5-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	76	baboon1-DuneBeast-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	77	baboon2-RockDweller-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	78	baboon3-JungleHunter-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	79	baboon4-DoomApe-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	80	baboon5-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	81	goatman1-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	82	goatman2-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	83	goatman3-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	84	goatman4-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	85	goatman5-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	86	fallenshaman1-FallenShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	87	fallenshaman2-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	88	fallenshaman3-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	89	fallenshaman4-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	90	fallenshaman5-WarpedShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	91	quillrat1-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	92	quillrat2-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	93	quillrat3-ThornBeast-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	94	quillrat4-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	95	quillrat5-JungleUrchin-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	96	sandmaggot1-SandMaggot-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	97	sandmaggot2-RockWorm-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	98	sandmaggot3-Devourer-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	99	sandmaggot4-GiantLamprey-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	100	sandmaggot5-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	101	clawviper1-TombViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	102	clawviper2-ClawViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	103	clawviper3-Salamander-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	104	clawviper4-PitViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	105	clawviper5-SerpentMagus-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	106	sandleaper1-SandLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	107	sandleaper2-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	108	sandleaper3-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	109	sandleaper4-TreeLurker-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	110	sandleaper5-RazorPitDemon-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	111	pantherwoman1-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	112	pantherwoman2-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	113	pantherwoman3-NightTiger-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	114	pantherwoman4-HellCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	115	swarm1-Itchies-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
4	1	116	swarm2-BlackLocusts-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
4	1	117	swarm3-PlagueBugs-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
4	1	118	swarm4-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
4	1	119	scarab1-DungSoldier-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	120	scarab2-SandWarrior-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	121	scarab3-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	122	scarab4-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	123	scarab5-AlbinoRoach-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	124	mummy1-DriedCorpse-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	125	mummy2-Decayed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	126	mummy3-Embalmed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	127	mummy4-PreservedDead-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	128	mummy5-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	129	unraveler1-HollowOne-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	130	unraveler2-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	131	unraveler3-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	132	unraveler4-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	133	unraveler5-Baal Subject Mummy-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	134	chaoshorde1-unused-Idle				/Data/Global/Monsters	CH																					0
4	1	135	chaoshorde2-unused-Idle				/Data/Global/Monsters	CH																					0
4	1	136	chaoshorde3-unused-Idle				/Data/Global/Monsters	CH																					0
4	1	137	chaoshorde4-unused-Idle				/Data/Global/Monsters	CH																					0
4	1	138	vulture1-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
4	1	139	vulture2-UndeadScavenger-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
4	1	140	vulture3-HellBuzzard-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
4	1	141	vulture4-WingedNightmare-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
4	1	142	mosquito1-Sucker-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
4	1	143	mosquito2-Feeder-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
4	1	144	mosquito3-BloodHook-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
4	1	145	mosquito4-BloodWing-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
4	1	146	willowisp1-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	147	willowisp2-SwampGhost-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	148	willowisp3-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	149	willowisp4-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	150	arach1-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	151	arach2-SandFisher-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	152	arach3-PoisonSpinner-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	153	arach4-FlameSpider-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	154	arach5-SpiderMagus-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	155	thornhulk1-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
4	1	156	thornhulk2-BrambleHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
4	1	157	thornhulk3-Thrasher-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
4	1	158	thornhulk4-Spikefist-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
4	1	159	vampire1-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	160	vampire2-NightLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	161	vampire3-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	162	vampire4-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	163	vampire5-Banished-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	164	batdemon1-DesertWing-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	165	batdemon2-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	166	batdemon3-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	167	batdemon4-BloodDiver-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	168	batdemon5-DarkFamiliar-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	169	fetish1-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	170	fetish2-Fetish-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	171	fetish3-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	172	fetish4-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	173	fetish5-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	174	cain1-DeckardCain-NpcOutOfTown				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
4	1	175	gheed-Gheed-Npc				/Data/Global/Monsters	GH	NU	HTH		LIT																	0
4	1	176	akara-Akara-Npc				/Data/Global/Monsters	PS	NU	HTH		LIT																	0
4	1	177	chicken-dummy-Idle				/Data/Global/Monsters	CK	NU	HTH		LIT																	0
4	1	178	kashya-Kashya-Npc				/Data/Global/Monsters	RC	NU	HTH		LIT																	0
4	1	179	rat-dummy-Idle				/Data/Global/Monsters	RT	NU	HTH		LIT																	0
4	1	180	rogue1-Dummy-Idle				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
4	1	181	hellmeteor-Dummy-HellMeteor				/Data/Global/Monsters	K9																					0
4	1	182	charsi-Charsi-Npc				/Data/Global/Monsters	CI	NU	HTH		LIT																	0
4	1	183	warriv1-Warriv-Npc				/Data/Global/Monsters	WA	NU	HTH		LIT																	0
4	1	184	andariel-Andariel-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
4	1	185	bird1-dummy-Idle				/Data/Global/Monsters	BS	WL	HTH		LIT																	0
4	1	186	bird2-dummy-Idle				/Data/Global/Monsters	BL																					0
4	1	187	bat-dummy-Idle				/Data/Global/Monsters	B9	WL	HTH		LIT																	0
4	1	188	cr_archer1-DarkRanger-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	189	cr_archer2-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	190	cr_archer3-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	191	cr_archer4-BlackArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	192	cr_archer5-FleshArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	193	cr_lancer1-DarkSpearwoman-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	194	cr_lancer2-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	195	cr_lancer3-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	196	cr_lancer4-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	197	cr_lancer5-FleshLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	198	sk_archer1-SkeletonArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	199	sk_archer2-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	200	sk_archer3-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	201	sk_archer4-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	202	sk_archer5-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	203	warriv2-Warriv-Npc				/Data/Global/Monsters	WX	NU	HTH		LIT																	0
4	1	204	atma-Atma-Npc				/Data/Global/Monsters	AS	NU	HTH		LIT																	0
4	1	205	drognan-Drognan-Npc				/Data/Global/Monsters	DR	NU	HTH		LIT																	0
4	1	206	fara-Fara-Npc				/Data/Global/Monsters	OF	NU	HTH		LIT																	0
4	1	207	cow-dummy-Idle				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
4	1	208	maggotbaby1-SandMaggotYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
4	1	209	maggotbaby2-RockWormYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
4	1	210	maggotbaby3-DevourerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
4	1	211	maggotbaby4-GiantLampreyYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
4	1	212	maggotbaby5-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
4	1	213	camel-dummy-Idle				/Data/Global/Monsters	CM	NU	HTH		LIT																	0
4	1	214	blunderbore1-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
4	1	215	blunderbore2-Gorbelly-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
4	1	216	blunderbore3-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
4	1	217	blunderbore4-Urdar-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
4	1	218	maggotegg1-SandMaggotEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
4	1	219	maggotegg2-RockWormEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
4	1	220	maggotegg3-DevourerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
4	1	221	maggotegg4-GiantLampreyEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
4	1	222	maggotegg5-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
4	1	223	act2male-dummy-Towner				/Data/Global/Monsters	2M	NU	HTH	OLD	MED	MED						TUR										0
4	1	224	act2female-Dummy-Towner				/Data/Global/Monsters	2F	NU	HTH	LIT	LIT	LIT																0
4	1	225	act2child-dummy-Towner				/Data/Global/Monsters	2C																					0
4	1	226	greiz-Greiz-Npc				/Data/Global/Monsters	GR	NU	HTH		LIT																	0
4	1	227	elzix-Elzix-Npc				/Data/Global/Monsters	EL	NU	HTH		LIT																	0
4	1	228	geglash-Geglash-Npc				/Data/Global/Monsters	GE	NU	HTH		LIT																	0
4	1	229	jerhyn-Jerhyn-Npc				/Data/Global/Monsters	JE	NU	HTH		LIT																	0
4	1	230	lysander-Lysander-Npc				/Data/Global/Monsters	LY	NU	HTH		LIT																	0
4	1	231	act2guard1-Dummy-Towner				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
4	1	232	act2vendor1-dummy-Vendor				/Data/Global/Monsters	M1	NU	HTH		LIT																	0
4	1	233	act2vendor2-dummy-Vendor				/Data/Global/Monsters	M2	NU	HTH		LIT																	0
4	1	234	crownest1-FoulCrowNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
4	1	235	crownest2-BloodHawkNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
4	1	236	crownest3-BlackVultureNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
4	1	237	crownest4-CloudStalkerNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
4	1	238	meshif1-Meshif-Npc				/Data/Global/Monsters	MS	NU	HTH		LIT																	0
4	1	239	duriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
4	1	240	bonefetish1-Undead RatMan-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	241	bonefetish2-Undead Fetish-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	242	bonefetish3-Undead Flayer-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	243	bonefetish4-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	244	bonefetish5-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	245	darkguard1-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	246	darkguard2-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	247	darkguard3-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	248	darkguard4-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	249	darkguard5-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	250	bloodmage1-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	251	bloodmage2-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	252	bloodmage3-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	253	bloodmage4-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	254	bloodmage5-unused-Idle				/Data/Global/Monsters	xx																					0
4	1	255	maggot-Maggot-Idle				/Data/Global/Monsters	MA	NU	HTH		LIT																	0
4	1	256	sarcophagus-MummyGenerator-Sarcophagus				/Data/Global/Monsters	MG	NU	HTH		LIT																	0
4	1	257	radament-Radament-GreaterMummy				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
4	1	258	firebeast-unused-ElementalBeast				/Data/Global/Monsters	FM	NU	HTH		LIT																	0
4	1	259	iceglobe-unused-ElementalBeast				/Data/Global/Monsters	IM	NU	HTH		LIT																	0
4	1	260	lightningbeast-unused-ElementalBeast				/Data/Global/Monsters	xx																					0
4	1	261	poisonorb-unused-ElementalBeast				/Data/Global/Monsters	PM	NU	HTH		LIT																	0
4	1	262	flyingscimitar-FlyingScimitar-FlyingScimitar				/Data/Global/Monsters	ST	NU	HTH		LIT																	0
4	1	263	zealot1-Zakarumite-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
4	1	264	zealot2-Faithful-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
4	1	265	zealot3-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
4	1	266	cantor1-Sexton-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
4	1	267	cantor2-Cantor-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
4	1	268	cantor3-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
4	1	269	cantor4-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
4	1	270	mephisto-Mephisto-Mephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
4	1	271	diablo-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
4	1	272	cain2-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
4	1	273	cain3-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
4	1	274	cain4-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
4	1	275	frogdemon1-Swamp Dweller-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
4	1	276	frogdemon2-Bog Creature-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
4	1	277	frogdemon3-Slime Prince-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
4	1	278	summoner-Summoner-Summoner				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
4	1	279	tyrael1-tyrael-NpcStationary				/Data/Global/Monsters	TX	NU	HTH		LIT		LIT	LIT														0
4	1	280	asheara-asheara-Npc				/Data/Global/Monsters	AH	NU	HTH		LIT																	0
4	1	281	hratli-hratli-Npc				/Data/Global/Monsters	HR	NU	HTH		LIT																	0
4	1	282	alkor-alkor-Npc				/Data/Global/Monsters	AL	NU	HTH		LIT																	0
4	1	283	ormus-ormus-Npc				/Data/Global/Monsters	OR	NU	HTH		LIT																	0
4	1	284	izual-izual-Izual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
4	1	285	halbu-halbu-Npc				/Data/Global/Monsters	20	NU	HTH		LIT																	0
4	1	286	tentacle1-WaterWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
4	1	287	tentacle2-RiverStalkerLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
4	1	288	tentacle3-StygianWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
4	1	289	tentaclehead1-WaterWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
4	1	290	tentaclehead2-RiverStalkerHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
4	1	291	tentaclehead3-StygianWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
4	1	292	meshif2-meshif-Npc				/Data/Global/Monsters	M3	NU	HTH		LIT																	0
4	1	293	cain5-DeckardCain-Npc				/Data/Global/Monsters	1D	NU	HTH		LIT																	0
4	1	294	navi-navi-Navi				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
4	1	295	bloodraven-Bloodraven-BloodRaven				/Data/Global/Monsters	CR	NU	BOW	BRV	HVY	BRV	HVY	HVY	LIT	LBB		HVY	HVY									0
4	1	296	bug-Dummy-Idle				/Data/Global/Monsters	BG	NU	HTH		LIT																	0
4	1	297	scorpion-Dummy-Idle				/Data/Global/Monsters	DS	NU	HTH		LIT																	0
4	1	298	rogue2-RogueScout-GoodNpcRanged				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
4	1	299	roguehire-Dummy-Hireable				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
4	1	300	rogue3-Dummy-TownRogue				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
4	1	301	gargoyletrap-GargoyleTrap-GargoyleTrap				/Data/Global/Monsters	GT	NU	HTH		LIT																	0
4	1	302	skmage_pois1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	303	skmage_pois2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	304	skmage_pois3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	305	skmage_pois4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	306	fetishshaman1-RatManShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	307	fetishshaman2-FetishShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	308	fetishshaman3-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	309	fetishshaman4-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	310	fetishshaman5-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	311	larva-larva-Idle				/Data/Global/Monsters	LV	NU	HTH		LIT																	0
4	1	312	maggotqueen1-SandMaggotQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
4	1	313	maggotqueen2-RockWormQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
4	1	314	maggotqueen3-DevourerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
4	1	315	maggotqueen4-GiantLampreyQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
4	1	316	maggotqueen5-WorldKillerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
4	1	317	claygolem-ClayGolem-NecroPet				/Data/Global/Monsters	G1	NU	HTH		LIT																	0
4	1	318	bloodgolem-BloodGolem-NecroPet				/Data/Global/Monsters	G2	NU	HTH		LIT																	0
4	1	319	irongolem-IronGolem-NecroPet				/Data/Global/Monsters	G4	NU	HTH		LIT																	0
4	1	320	firegolem-FireGolem-NecroPet				/Data/Global/Monsters	G3	NU	HTH		LIT																	0
4	1	321	familiar-Dummy-Idle				/Data/Global/Monsters	FI	NU	HTH		LIT																	0
4	1	322	act3male-Dummy-Towner				/Data/Global/Monsters	N4	NU	HTH	BRD	HVY	HVY	HEV	HEV	FSH	SAK		TKT										0
4	1	323	baboon6-NightMarauder-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	324	act3female-Dummy-Towner				/Data/Global/Monsters	N3	NU	HTH	LIT	MTP	SRT			BSK	BSK												0
4	1	325	natalya-Natalya-Npc				/Data/Global/Monsters	TZ	NU	HTH		LIT																	0
4	1	326	vilemother1-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
4	1	327	vilemother2-StygianHag-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
4	1	328	vilemother3-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
4	1	329	vilechild1-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
4	1	330	vilechild2-StygianDog-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
4	1	331	vilechild3-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
4	1	332	fingermage1-Groper-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
4	1	333	fingermage2-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
4	1	334	fingermage3-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
4	1	335	regurgitator1-Corpulent-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
4	1	336	regurgitator2-CorpseSpitter-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
4	1	337	regurgitator3-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
4	1	338	doomknight1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	339	doomknight2-AbyssKnight-AbyssKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	340	doomknight3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	341	quillbear1-QuillBear-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
4	1	342	quillbear2-SpikeGiant-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
4	1	343	quillbear3-ThornBrute-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
4	1	344	quillbear4-RazorBeast-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
4	1	345	quillbear5-GiantUrchin-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
4	1	346	snake-Dummy-Idle				/Data/Global/Monsters	CO	NU	HTH		LIT																	0
4	1	347	parrot-Dummy-Idle				/Data/Global/Monsters	PR	WL	HTH		LIT																	0
4	1	348	fish-Dummy-Idle				/Data/Global/Monsters	FJ																					0
4	1	349	evilhole1-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
4	1	350	evilhole2-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
4	1	351	evilhole3-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
4	1	352	evilhole4-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
4	1	353	evilhole5-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
4	1	354	trap-firebolt-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
4	1	355	trap-horzmissile-a trap-Trap-RightArrow				/Data/Global/Monsters	9A																					0
4	1	356	trap-vertmissile-a trap-Trap-LeftArrow				/Data/Global/Monsters	9A																					0
4	1	357	trap-poisoncloud-a trap-Trap-Poison				/Data/Global/Monsters	9A																					0
4	1	358	trap-lightning-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
4	1	359	act2guard2-Kaelan-JarJar				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
4	1	360	invisospawner-Dummy-InvisoSpawner				/Data/Global/Monsters	K9																					0
4	1	361	diabloclone-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
4	1	362	suckernest1-SuckerNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
4	1	363	suckernest2-FeederNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
4	1	364	suckernest3-BloodHookNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
4	1	365	suckernest4-BloodWingNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
4	1	366	act2hire-Guard-Hireable				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
4	1	367	minispider-Dummy-Idle				/Data/Global/Monsters	LS	NU	HTH		LIT																	0
4	1	368	boneprison1--Idle				/Data/Global/Monsters	67	NU	HTH		LIT																	0
4	1	369	boneprison2--Idle				/Data/Global/Monsters	66	NU	HTH		LIT																	0
4	1	370	boneprison3--Idle				/Data/Global/Monsters	69	NU	HTH		LIT																	0
4	1	371	boneprison4--Idle				/Data/Global/Monsters	68	NU	HTH		LIT																	0
4	1	372	bonewall-Dummy-BoneWall				/Data/Global/Monsters	BW	NU	HTH		LIT																	0
4	1	373	councilmember1-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	374	councilmember2-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	375	councilmember3-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	376	turret1-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
4	1	377	turret2-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
4	1	378	turret3-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
4	1	379	hydra1-Hydra-Hydra				/Data/Global/Monsters	HX	NU	HTH		LIT							LIT										0
4	1	380	hydra2-Hydra-Hydra				/Data/Global/Monsters	21	NU	HTH		LIT							LIT										0
4	1	381	hydra3-Hydra-Hydra				/Data/Global/Monsters	HZ	NU	HTH		LIT							LIT										0
4	1	382	trap-melee-a trap-Trap-Melee				/Data/Global/Monsters	M4	A1	HTH		LIT																	0
4	1	383	seventombs-Dummy-7TIllusion				/Data/Global/Monsters	9A																					0
4	1	384	dopplezon-Dopplezon-Idle				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
4	1	385	valkyrie-Valkyrie-NecroPet				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
4	1	386	act2guard3-Dummy-Idle				/Data/Global/Monsters	SK																					0
4	1	387	act3hire-Iron Wolf-Hireable				/Data/Global/Monsters	IW	NU	1HS	LIT	LIT				WND		KIT											0
4	1	388	megademon1-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	389	megademon2-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	390	megademon3-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	391	necroskeleton-NecroSkeleton-NecroPet				/Data/Global/Monsters	SK	NU	1HS	DES	HVY	DES	DES	DES	SCM		KIT	DES	DES	LIT								0
4	1	392	necromage-NecroMage-NecroPet				/Data/Global/Monsters	SK	NU	HTH	DES	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
4	1	393	griswold-Griswold-Griswold				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
4	1	394	compellingorb-compellingorb-Idle				/Data/Global/Monsters	9a																					0
4	1	395	tyrael2-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
4	1	396	darkwanderer-youngdiablo-DarkWanderer				/Data/Global/Monsters	1Z	NU	HTH		LIT																	0
4	1	397	trap-nova-a trap-Trap-Nova				/Data/Global/Monsters	9A																					0
4	1	398	spiritmummy-Dummy-Idle				/Data/Global/Monsters	xx																					0
4	1	399	lightningspire-LightningSpire-ArcaneTower				/Data/Global/Monsters	AE	NU	HTH		LIT							LIT										0
4	1	400	firetower-FireTower-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
4	1	401	slinger1-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
4	1	402	slinger2-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
4	1	403	slinger3-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
4	1	404	slinger4-HellSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
4	1	405	act2guard4-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
4	1	406	act2guard5-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
4	1	407	skmage_cold1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
4	1	408	skmage_cold2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
4	1	409	skmage_cold3-BaalColdMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
4	1	410	skmage_cold4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
4	1	411	skmage_fire1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
4	1	412	skmage_fire2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
4	1	413	skmage_fire3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
4	1	414	skmage_fire4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
4	1	415	skmage_ltng1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
4	1	416	skmage_ltng2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
4	1	417	skmage_ltng3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
4	1	418	skmage_ltng4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
4	1	419	hellbovine-Hell Bovine-Skeleton				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
4	1	420	window1--Idle				/Data/Global/Monsters	VH	NU	HTH		LIT							LIT										0
4	1	421	window2--Idle				/Data/Global/Monsters	VJ	NU	HTH		LIT							LIT										0
4	1	422	slinger5-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
4	1	423	slinger6-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
4	1	424	fetishblow1-RatMan-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	425	fetishblow2-Fetish-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	426	fetishblow3-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	427	fetishblow4-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	428	fetishblow5-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	429	mephistospirit-Dummy-Spirit				/Data/Global/Monsters	M6	A1	HTH		LIT																	0
4	1	430	smith-The Smith-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
4	1	431	trappedsoul1-TrappedSoul-TrappedSoul				/Data/Global/Monsters	10	NU	HTH		LIT																	0
4	1	432	trappedsoul2-TrappedSoul-TrappedSoul				/Data/Global/Monsters	13	NU	HTH		LIT																	0
4	1	433	jamella-Jamella-Npc				/Data/Global/Monsters	ja	NU	HTH		LIT																	0
4	1	434	izualghost-Izual-NpcStationary				/Data/Global/Monsters	17	NU	HTH		LIT							LIT										0
4	1	435	fetish11-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	436	malachai-Malachai-Buffy				/Data/Global/Monsters	36	NU	HTH		LIT							LIT										0
4	1	437	hephasto-The Feature Creep-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
4	1	438	wakeofdestruction-Wake of Destruction-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
4	1	439	chargeboltsentry-Charged Bolt Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
4	1	440	lightningsentry-Lightning Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
4	1	441	bladecreeper-Blade Creeper-BladeCreeper				/Data/Global/Monsters	b8	NU	HTH		LIT							LIT										0
4	1	442	invisopet-Invis Pet-InvisoPet				/Data/Global/Monsters	k9																					0
4	1	443	infernosentry-Inferno Sentry-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
4	1	444	deathsentry-Death Sentry-DeathSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
4	1	445	shadowwarrior-Shadow Warrior-ShadowWarrior				/Data/Global/Monsters	k9																					0
4	1	446	shadowmaster-Shadow Master-ShadowMaster				/Data/Global/Monsters	k9																					0
4	1	447	druidhawk-Druid Hawk-Raven				/Data/Global/Monsters	hk	NU	HTH		LIT																	0
4	1	448	spiritwolf-Druid Spirit Wolf-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
4	1	449	fenris-Druid Fenris-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
4	1	450	spiritofbarbs-Spirit of Barbs-Totem				/Data/Global/Monsters	x4	NU	HTH		LIT																	0
4	1	451	heartofwolverine-Heart of Wolverine-Totem				/Data/Global/Monsters	x3	NU	HTH		LIT																	0
4	1	452	oaksage-Oak Sage-Totem				/Data/Global/Monsters	xw	NU	HTH		LIT																	0
4	1	453	plaguepoppy-Druid Plague Poppy-Vines				/Data/Global/Monsters	k9																					0
4	1	454	cycleoflife-Druid Cycle of Life-CycleOfLife				/Data/Global/Monsters	k9																					0
4	1	455	vinecreature-Vine Creature-CycleOfLife				/Data/Global/Monsters	k9																					0
4	1	456	druidbear-Druid Bear-DruidBear				/Data/Global/Monsters	b7	NU	HTH		LIT																	0
4	1	457	eagle-Eagle-Idle				/Data/Global/Monsters	eg	NU	HTH		LIT							LIT										0
4	1	458	wolf-Wolf-NecroPet				/Data/Global/Monsters	40	NU	HTH		LIT																	0
4	1	459	bear-Bear-NecroPet				/Data/Global/Monsters	TG	NU	HTH		LIT							LIT										0
4	1	460	barricadedoor1-Barricade Door-Idle				/Data/Global/Monsters	AJ	NU	HTH		LIT																	0
4	1	461	barricadedoor2-Barricade Door-Idle				/Data/Global/Monsters	AG	NU	HTH		LIT																	0
4	1	462	prisondoor-Prison Door-Idle				/Data/Global/Monsters	2Q	NU	HTH		LIT																	0
4	1	463	barricadetower-Barricade Tower-SiegeTower				/Data/Global/Monsters	ac	NU	HTH		LIT							LIT						LIT				0
4	1	464	reanimatedhorde1-RotWalker-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	465	reanimatedhorde2-ReanimatedHorde-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	466	reanimatedhorde3-ProwlingDead-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	467	reanimatedhorde4-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	468	reanimatedhorde5-DefiledWarrior-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	469	siegebeast1-Siege Beast-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
4	1	470	siegebeast2-CrushBiest-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
4	1	471	siegebeast3-BloodBringer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
4	1	472	siegebeast4-GoreBearer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
4	1	473	siegebeast5-DeamonSteed-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
4	1	474	snowyeti1-SnowYeti1-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
4	1	475	snowyeti2-SnowYeti2-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
4	1	476	snowyeti3-SnowYeti3-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
4	1	477	snowyeti4-SnowYeti4-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
4	1	478	wolfrider1-WolfRider1-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
4	1	479	wolfrider2-WolfRider2-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
4	1	480	wolfrider3-WolfRider3-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
4	1	481	minion1-Minionexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	482	minion2-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	483	minion3-IceBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	484	minion4-FireBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	485	minion5-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	486	minion6-IceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	487	minion7-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	488	minion8-GreaterIceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	489	suicideminion1-FanaticMinion-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	490	suicideminion2-BerserkSlayer-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	491	suicideminion3-ConsumedIceBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	492	suicideminion4-ConsumedFireBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	493	suicideminion5-FrenziedHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	494	suicideminion6-FrenziedIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	495	suicideminion7-InsaneHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	496	suicideminion8-InsaneIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
4	1	497	succubus1-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	498	succubus2-VileTemptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	499	succubus3-StygianHarlot-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	500	succubus4-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	501	succubus5-Blood Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	502	succubuswitch1-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	503	succubuswitch2-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	504	succubuswitch3-StygianFury-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	505	succubuswitch4-Blood Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	506	succubuswitch5-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	507	overseer1-OverSeer-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	508	overseer2-Lasher-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	509	overseer3-OverLord-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	510	overseer4-BloodBoss-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	511	overseer5-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	512	minionspawner1-MinionSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	513	minionspawner2-MinionSlayerSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	514	minionspawner3-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	515	minionspawner4-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	516	minionspawner5-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	517	minionspawner6-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	518	minionspawner7-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	519	minionspawner8-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	520	imp1-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	521	imp2-Imp2-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	522	imp3-Imp3-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	523	imp4-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	524	imp5-Imp5-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	525	catapult1-CatapultS-Catapult				/Data/Global/Monsters	65	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
4	1	526	catapult2-CatapultE-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
4	1	527	catapult3-CatapultSiege-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
4	1	528	catapult4-CatapultW-Catapult				/Data/Global/Monsters	ua	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT	LIT								0
4	1	529	frozenhorror1-Frozen Horror1-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
4	1	530	frozenhorror2-Frozen Horror2-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
4	1	531	frozenhorror3-Frozen Horror3-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
4	1	532	frozenhorror4-Frozen Horror4-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
4	1	533	frozenhorror5-Frozen Horror5-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
4	1	534	bloodlord1-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	535	bloodlord2-Blood Lord2-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	536	bloodlord3-Blood Lord3-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	537	bloodlord4-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	538	bloodlord5-Blood Lord5-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	539	larzuk-Larzuk-Npc				/Data/Global/Monsters	XR	NU	HTH		LIT																	0
4	1	540	drehya-Drehya-Npc				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
4	1	541	malah-Malah-Npc				/Data/Global/Monsters	XT	NU	HTH		LIT																	0
4	1	542	nihlathak-Nihlathak Town-Npc				/Data/Global/Monsters	0J	NU	HTH		LIT																	0
4	1	543	qual-kehk-Qual-Kehk-Npc				/Data/Global/Monsters	XV	NU	HTH		LIT																	0
4	1	544	catapultspotter1-Catapult Spotter S-CatapultSpotter				/Data/Global/Monsters	k9																					0
4	1	545	catapultspotter2-Catapult Spotter E-CatapultSpotter				/Data/Global/Monsters	k9																					0
4	1	546	catapultspotter3-Catapult Spotter Siege-CatapultSpotter				/Data/Global/Monsters	k9																					0
4	1	547	catapultspotter4-Catapult Spotter W-CatapultSpotter				/Data/Global/Monsters	k9																					0
4	1	548	cain6-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
4	1	549	tyrael3-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
4	1	550	act5barb1-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
4	1	551	act5barb2-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
4	1	552	barricadewall1-Barricade Wall Right-Idle				/Data/Global/Monsters	A6	NU	HTH		LIT																	0
4	1	553	barricadewall2-Barricade Wall Left-Idle				/Data/Global/Monsters	AK	NU	HTH		LIT																	0
4	1	554	nihlathakboss-Nihlathak-Nihlathak				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
4	1	555	drehyaiced-Drehya-NpcOutOfTown				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
4	1	556	evilhut-Evil hut-GenericSpawner				/Data/Global/Monsters	2T	NU	HTH		LIT							LIT										0
4	1	557	deathmauler1-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	558	deathmauler2-Death Mauler2-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	559	deathmauler3-Death Mauler3-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	560	deathmauler4-Death Mauler4-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	561	deathmauler5-Death Mauler5-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	562	act5pow-POW-Wussie				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
4	1	563	act5barb3-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
4	1	564	act5barb4-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
4	1	565	ancientstatue1-Ancient Statue 1-AncientStatue				/Data/Global/Monsters	0G	NU	HTH		LIT																	0
4	1	566	ancientstatue2-Ancient Statue 2-AncientStatue				/Data/Global/Monsters	0H	NU	HTH		LIT																	0
4	1	567	ancientstatue3-Ancient Statue 3-AncientStatue				/Data/Global/Monsters	0I	NU	HTH		LIT																	0
4	1	568	ancientbarb1-Ancient Barbarian 1-Ancient				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
4	1	569	ancientbarb2-Ancient Barbarian 2-Ancient				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
4	1	570	ancientbarb3-Ancient Barbarian 3-Ancient				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
4	1	571	baalthrone-Baal Throne-BaalThrone				/Data/Global/Monsters	41	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
4	1	572	baalcrab-Baal Crab-BaalCrab				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
4	1	573	baaltaunt-Baal Taunt-BaalTaunt				/Data/Global/Monsters	K9																					0
4	1	574	putriddefiler1-Putrid Defiler1-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
4	1	575	putriddefiler2-Putrid Defiler2-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
4	1	576	putriddefiler3-Putrid Defiler3-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
4	1	577	putriddefiler4-Putrid Defiler4-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
4	1	578	putriddefiler5-Putrid Defiler5-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
4	1	579	painworm1-Pain Worm1-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
4	1	580	painworm2-Pain Worm2-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
4	1	581	painworm3-Pain Worm3-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
4	1	582	painworm4-Pain Worm4-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
4	1	583	painworm5-Pain Worm5-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
4	1	584	bunny-dummy-Idle				/Data/Global/Monsters	48	NU	HTH		LIT																	0
4	1	585	baalhighpriest-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	586	venomlord-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				FLB													0
4	1	587	baalcrabstairs-Baal Crab to Stairs-BaalToStairs				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
4	1	588	act5hire1-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
4	1	589	act5hire2-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
4	1	590	baaltentacle1-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
4	1	591	baaltentacle2-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
4	1	592	baaltentacle3-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
4	1	593	baaltentacle4-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
4	1	594	baaltentacle5-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
4	1	595	injuredbarb1-dummy-Idle				/Data/Global/Monsters	6z	NU	HTH		LIT																	0
4	1	596	injuredbarb2-dummy-Idle				/Data/Global/Monsters	7j	NU	HTH		LIT																	0
4	1	597	injuredbarb3-dummy-Idle				/Data/Global/Monsters	7i	NU	HTH		LIT																	0
4	1	598	baalclone-Baal Crab Clone-BaalCrabClone				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
4	1	599	baalminion1-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
4	1	600	baalminion2-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
4	1	601	baalminion3-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
4	1	602	worldstoneeffect-dummy-Idle				/Data/Global/Monsters	K9																					0
4	1	603	sk_archer6-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	604	sk_archer7-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	605	sk_archer8-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	606	sk_archer9-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	607	sk_archer10-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	608	bighead6-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	609	bighead7-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	610	bighead8-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	611	bighead9-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	612	bighead10-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	613	goatman6-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	614	goatman7-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	615	goatman8-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	616	goatman9-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	617	goatman10-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
4	1	618	foulcrow5-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	619	foulcrow6-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	620	foulcrow7-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	621	foulcrow8-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
4	1	622	clawviper6-ClawViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	623	clawviper7-PitViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	624	clawviper8-Salamander-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	625	clawviper9-TombViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	626	clawviper10-SerpentMagus-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	627	sandraider6-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	628	sandraider7-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	629	sandraider8-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	630	sandraider9-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	631	sandraider10-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	632	deathmauler6-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	633	quillrat6-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	634	quillrat7-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	635	quillrat8-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	636	vulture5-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
4	1	637	thornhulk5-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
4	1	638	slinger7-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	639	slinger8-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	640	slinger9-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	641	cr_archer6-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	642	cr_archer7-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	643	cr_lancer6-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	644	cr_lancer7-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	645	cr_lancer8-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	646	blunderbore5-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
4	1	647	blunderbore6-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
4	1	648	skmage_fire5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
4	1	649	skmage_fire6-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
4	1	650	skmage_ltng5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
4	1	651	skmage_ltng6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
4	1	652	skmage_cold5-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		CLD	CLD						0
4	1	653	skmage_pois5-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	654	skmage_pois6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	655	pantherwoman5-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	656	pantherwoman6-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	657	sandleaper6-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	658	sandleaper7-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
4	1	659	wraith6-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	660	wraith7-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	661	wraith8-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	662	succubus6-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	663	succubus7-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	664	succubuswitch6-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	665	succubuswitch7-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	666	succubuswitch8-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	667	willowisp5-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	668	willowisp6-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	669	willowisp7-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	670	fallen6-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
4	1	671	fallen7-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
4	1	672	fallen8-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
4	1	673	fallenshaman6-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	674	fallenshaman7-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	675	fallenshaman8-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	676	skeleton6-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	677	skeleton7-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	678	batdemon6-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	679	batdemon7-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	680	bloodlord6-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	681	bloodlord7-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	682	scarab6-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	683	scarab7-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	684	fetish6-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	685	fetish7-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	686	fetish8-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
4	1	687	fetishblow6-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	688	fetishblow7-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	689	fetishblow8-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
4	1	690	fetishshaman6-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	691	fetishshaman7-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	692	fetishshaman8-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	693	baboon7-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	694	baboon8-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
4	1	695	unraveler6-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	696	unraveler7-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	697	unraveler8-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	698	unraveler9-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	699	zealot4-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
4	1	700	zealot5-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
4	1	701	cantor5-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
4	1	702	cantor6-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
4	1	703	vilemother4-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
4	1	704	vilemother5-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
4	1	705	vilechild4-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
4	1	706	vilechild5-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
4	1	707	sandmaggot6-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	708	maggotbaby6-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
4	1	709	maggotegg6-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
4	1	710	minion9-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	711	minion10-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	712	minion11-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	713	arach6-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	714	megademon4-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	715	megademon5-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	716	imp6-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	717	imp7-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	718	bonefetish6-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	719	bonefetish7-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
4	1	720	fingermage4-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
4	1	721	fingermage5-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
4	1	722	regurgitator4-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
4	1	723	vampire6-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	724	vampire7-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	725	vampire8-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	726	reanimatedhorde6-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	727	dkfig1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	728	dkfig2-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	729	dkmag1-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	730	dkmag2-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	731	mummy6-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	732	ubermephisto-Mephisto-UberMephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
4	1	733	uberdiablo-Diablo-UberDiablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
4	1	734	uberizual-izual-UberIzual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
4	1	735	uberandariel-Lilith-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
4	1	736	uberduriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
4	1	737	uberbaal-Baal Crab-UberBaal				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
4	1	738	demonspawner-Evil hut-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
4	1	739	demonhole-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
4	1	740	megademon6-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	741	dkmag3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	742	imp8-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	743	swarm5-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
4	1	744	sandmaggot7-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
4	1	745	arach7-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	746	scarab8-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	747	succubus8-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	748	succubuswitch9-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	749	corruptrogue6-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	750	cr_archer8-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	751	cr_lancer9-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
4	1	752	overseer6-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	753	skeleton8-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	754	sk_archer11-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
4	1	755	skmage_fire7-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
4	1	756	skmage_ltng7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
4	1	757	skmage_cold6-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
4	1	758	skmage_pois7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		POS	POS						0
4	1	759	vampire9-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
4	1	760	wraith9-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
4	1	761	willowisp8-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	762	Bishibosh-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	763	Bonebreak-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
4	1	764	Coldcrow-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
4	1	765	Rakanishu-SUPER UNIQUE				/Data/Global/Monsters	FA	NU	HTH		LIT				SWD		TCH	LIT										0
4	1	766	Treehead WoodFist-SUPER UNIQUE				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
4	1	767	Griswold-SUPER UNIQUE				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
4	1	768	The Countess-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	MED	LIT	MED	LIT	LIT	WHM			LIT	LIT									0
4	1	769	Pitspawn Fouldog-SUPER UNIQUE				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
4	1	770	Flamespike the Crawler-SUPER UNIQUE				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
4	1	771	Boneash-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
4	1	772	Radament-SUPER UNIQUE				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
4	1	773	Bloodwitch the Wild-SUPER UNIQUE				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
4	1	774	Fangskin-SUPER UNIQUE				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
4	1	775	Beetleburst-SUPER UNIQUE				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
4	1	776	Leatherarm-SUPER UNIQUE				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
4	1	777	Coldworm the Burrower-SUPER UNIQUE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
4	1	778	Fire Eye-SUPER UNIQUE				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
4	1	779	Dark Elder-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	780	The Summoner-SUPER UNIQUE				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
4	1	781	Ancient Kaa the Soulless-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	782	The Smith-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
4	1	783	Web Mage the Burning-SUPER UNIQUE				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
4	1	784	Witch Doctor Endugu-SUPER UNIQUE				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
4	1	785	Stormtree-SUPER UNIQUE				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
4	1	786	Sarina the Battlemaid-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
4	1	787	Icehawk Riftwing-SUPER UNIQUE				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
4	1	788	Ismail Vilehand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	789	Geleb Flamefinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	790	Bremm Sparkfist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	791	Toorc Icefist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	792	Wyand Voidfinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	793	Maffer Dragonhand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	794	Winged Death-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	795	The Tormentor-SUPER UNIQUE				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
4	1	796	Taintbreeder-SUPER UNIQUE				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
4	1	797	Riftwraith the Cannibal-SUPER UNIQUE				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
4	1	798	Infector of Souls-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	799	Lord De Seis-SUPER UNIQUE				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
4	1	800	Grand Vizier of Chaos-SUPER UNIQUE				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
4	1	801	The Cow King-SUPER UNIQUE				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
4	1	802	Corpsefire-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
4	1	803	The Feature Creep-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
4	1	804	Siege Boss-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	805	Ancient Barbarian 1-SUPER UNIQUE				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
4	1	806	Ancient Barbarian 2-SUPER UNIQUE				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
4	1	807	Ancient Barbarian 3-SUPER UNIQUE				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
4	1	808	Axe Dweller-SUPER UNIQUE				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
4	1	809	Bonesaw Breaker-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	810	Dac Farren-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	811	Megaflow Rectifier-SUPER UNIQUE				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
4	1	812	Eyeback Unleashed-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	813	Threash Socket-SUPER UNIQUE				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
4	1	814	Pindleskin-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
4	1	815	Snapchip Shatter-SUPER UNIQUE				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
4	1	816	Anodized Elite-SUPER UNIQUE				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
4	1	817	Vinvear Molech-SUPER UNIQUE				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
4	1	818	Sharp Tooth Sayer-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
4	1	819	Magma Torquer-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
4	1	820	Blaze Ripper-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
4	1	821	Frozenstein-SUPER UNIQUE				/Data/Global/Monsters	io	NU	HTH		LIT																	0
4	1	822	Nihlathak Boss-SUPER UNIQUE				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
4	1	823	Baal Subject 1-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
4	1	824	Baal Subject 2-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
4	1	825	Baal Subject 3-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
4	1	826	Baal Subject 4-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
4	1	827	Baal Subject 5-SUPER UNIQUE				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
4	2	0	Waypoint (238)	238			/Data/Global/Objects	WV	ON	HTH		LIT							LIT										0
4	2	1	-580	580																									0
4	2	2	Your Private Stash (267)	267			/Data/Global/Objects	B6	NU	HTH		LIT																	0
4	2	3	gold placeholder (269)	269			/Data/Global/Objects	1G	NU	HTH		LIT																	0
4	2	4	-581	581																									0
4	2	5	-573	573																									0
4	2	6	-573	573																									0
4	2	7	-573	573																									0
4	2	8	hell fire 1 (345)	345			/Data/Global/Objects	E3	NU	HTH		LIT																	0
4	2	9	hell fire 2 (346)	346			/Data/Global/Objects	E4	NU	HTH		LIT																	0
4	2	10	hell fire 3 (347)	347			/Data/Global/Objects	E5	NU	HTH		LIT																	0
4	2	11	hell lava 1 (348)	348			/Data/Global/Objects	E6	NU	HTH		LIT																	0
4	2	12	hell lava 2 (349)	349			/Data/Global/Objects	E7	NU	HTH		LIT																	0
4	2	13	hell lava 3 (350)	350			/Data/Global/Objects	E8	NU	HTH		LIT																	0
4	2	14	hell light source 1 (351)	351																									0
4	2	15	hell light source 2 (352)	352																									0
4	2	16	hell light source 3 (353)	353																									0
4	2	17	Fire, hell brazier 1 (358)	358			/Data/Global/Objects	E1	NU	HTH	LIT	LIT																	0
4	2	18	Fire, hell brazier 2 (359)	359			/Data/Global/Objects	E2	NU	HTH	LIT	LIT																	0
4	2	19	Hung Skeleton (363)	363			/Data/Global/Objects	XQ	OP	HTH		LIT																	0
4	2	20	skeleton rising from lava L (259)	259			/Data/Global/Objects	QS	OP	HTH		LIT							LIT										0
4	2	21	skeleton rising from lava R (373)	373			/Data/Global/Objects	QT	OP	HTH		LIT							LIT										0
4	2	22	Bone Chest (372)	372			/Data/Global/Objects	Y1	OP	HTH		LIT																	0
4	2	23	fog water (374)	374			/Data/Global/Objects	UD	NU	HTH		LIT																	0
4	2	24	Shrine, hell well (236)	236			/Data/Global/Objects	HO	OP	HTH		LIT							LIT										0
4	2	25	Shrine, mana well (249)	249			/Data/Global/Objects	HN	OP	HTH		LIT							LIT										0
4	2	26	Shrine, outer hell 1 (226)	226			/Data/Global/Objects	IA	OP	HTH		LIT							LIT										0
4	2	27	Shrine, outer hell 2 (231)	231			/Data/Global/Objects	HT	OP	HTH		LIT							LIT										0
4	2	28	Shrine, outer hell 3 (232)	232			/Data/Global/Objects	HQ	OP	HTH		LIT							LIT										0
4	2	29	Shrine, mana inner hell (93)	93			/Data/Global/Objects	IZ	OP	HTH		LIT							LIT										0
4	2	30	Shrine, inner hell 1 (97)	97			/Data/Global/Objects	IX	OP	HTH		LIT							LIT										0
4	2	31	Shrine, inner hell 2 (123)	123			/Data/Global/Objects	IW	OP	HTH		LIT							LIT										0
4	2	32	Shrine, inner hell 3 (124)	124			/Data/Global/Objects	IV	OP	HTH		LIT							LIT										0
4	2	33	Shrine, health inner hell (96)	96			/Data/Global/Objects	IY	OP	HTH		LIT							LIT										0
4	2	34	Skullpile (225)	225			/Data/Global/Objects	IB	OP	HTH		LIT																	0
4	2	35	Pillar 1 (233)	233			/Data/Global/Objects	HV	OP	HTH		LIT																	0
4	2	36	Pillar 2 (222)	222			/Data/Global/Objects	70	OP	HTH		LIT																	0
4	2	37	Hidden Stash, inner hell (125)	125			/Data/Global/Objects	IU	OP	HTH		LIT																	0
4	2	38	Skull Pile, inner hell (126)	126			/Data/Global/Objects	IS	OP	HTH		LIT																	0
4	2	39	Hidden Stash, inner hell 1 (127)	127			/Data/Global/Objects	IR	OP	HTH		LIT																	0
4	2	40	Hidden Stash, inner hell 2 (128)	128			/Data/Global/Objects	HG	ON	HTH		LIT																	0
4	2	41	CRASH THE GAME ! (375)	375																									0
4	2	42	Hellforge (376)	376			/Data/Global/Objects	UX	ON	HTH		LIT							LIT	LIT	LIT								0
4	2	43	ray of light L Diablo (254)	254			/Data/Global/Objects	12	NU	HTH		LIT																	0
4	2	44	ray of light R Diablo (253)	253			/Data/Global/Objects	11	NU	HTH		LIT																	0
4	2	45	Portal to hell (342)	342			/Data/Global/Objects	1Y	ON	HTH		LIT								LIT	LIT								0
4	2	46	Diablo start point (255)	255		7	/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
4	2	47	Diablo seal 1 (392)	392			/Data/Global/Objects	30	ON	HTH		LIT							LIT										0
4	2	48	Diablo seal 2 (393)	393			/Data/Global/Objects	31	ON	HTH		LIT							LIT										0
4	2	49	Diablo seal 3 (394)	394			/Data/Global/Objects	32	ON	HTH		LIT							LIT										0
4	2	50	Diablo seal 4 (395)	395			/Data/Global/Objects	33	ON	HTH		LIT							LIT										0
4	2	51	Diablo seal 5 (396)	396			/Data/Global/Objects	34	ON	HTH		LIT							LIT										0
4	2	52	Waypoint, fortress (398)	398			/Data/Global/Objects	YG	ON	HTH		LIT							LIT										0
4	2	53	Chest, sparkly (397)	397			/Data/Global/Objects	YF	OP	HTH		LIT																	0
4	2	54	fissure (399)	399			/Data/Global/Objects	FH	OP	HTH		LIT							LIT										0
4	2	55	smoke (401)	401			/Data/Global/Objects	35	NU	HTH		LIT																	0
4	2	56	brazier hell mesa (400)	400			/Data/Global/Objects	HE	NU	HTH		LIT							LIT										0
4	2	57	Trapped Soul, burning guy (380)	380			/Data/Global/Objects	UY	ON	HTH		LIT							LIT										0
4	2	58	Trapped Soul, guy stuck 1 (383)	383			/Data/Global/Objects	18	OP	HTH		LIT																	0
4	2	59	Trapped Soul, guy stuck 2 (384)	384			/Data/Global/Objects	19	OP	HTH		LIT																	0
4	2	60	wall torch L tombs (296)	296			/Data/Global/Objects	QD	NU	HTH		LIT							LIT										0
4	2	61	wall torch R tombs (297)	297			/Data/Global/Objects	QE	NU	HTH		LIT							LIT										0
4	2	62	Fire, hell brazier (403)	403			/Data/Global/Objects	9F	NU	HTH		LIT							LIT										0
4	2	63	floor brazier (102)	102			/Data/Global/Objects	FB	ON	HTH		LIT							LIT										0
4	2	64	fortress brazier (408)	408			/Data/Global/Objects	98	NU	HTH		LIT							LIT										0
4	2	65	Torch Pit (409)	409			/Data/Global/Objects	99	NU	HTH		LIT							LIT										0
4	2	66	ACT 4 TABLE SKIP IT	0																									0
4	2	67	ACT 4 TABLE SKIP IT	0																									0
4	2	68	ACT 4 TABLE SKIP IT	0																									0
4	2	69	ACT 4 TABLE SKIP IT	0																									0
4	2	70	ACT 4 TABLE SKIP IT	0																									0
4	2	71	ACT 4 TABLE SKIP IT	0																									0
4	2	72	ACT 4 TABLE SKIP IT	0																									0
4	2	73	ACT 4 TABLE SKIP IT	0																									0
4	2	74	ACT 4 TABLE SKIP IT	0																									0
4	2	75	ACT 4 TABLE SKIP IT	0																									0
4	2	76	ACT 4 TABLE SKIP IT	0																									0
4	2	77	ACT 4 TABLE SKIP IT	0																									0
4	2	78	ACT 4 TABLE SKIP IT	0																									0
4	2	79	ACT 4 TABLE SKIP IT	0																									0
4	2	80	ACT 4 TABLE SKIP IT	0																									0
4	2	81	ACT 4 TABLE SKIP IT	0																									0
4	2	82	ACT 4 TABLE SKIP IT	0																									0
4	2	83	ACT 4 TABLE SKIP IT	0																									0
4	2	84	ACT 4 TABLE SKIP IT	0																									0
4	2	85	ACT 4 TABLE SKIP IT	0																									0
4	2	86	ACT 4 TABLE SKIP IT	0																									0
4	2	87	ACT 4 TABLE SKIP IT	0																									0
4	2	88	ACT 4 TABLE SKIP IT	0																									0
4	2	89	ACT 4 TABLE SKIP IT	0																									0
4	2	90	ACT 4 TABLE SKIP IT	0																									0
4	2	91	ACT 4 TABLE SKIP IT	0																									0
4	2	92	ACT 4 TABLE SKIP IT	0																									0
4	2	93	ACT 4 TABLE SKIP IT	0																									0
4	2	94	ACT 4 TABLE SKIP IT	0																									0
4	2	95	ACT 4 TABLE SKIP IT	0																									0
4	2	96	ACT 4 TABLE SKIP IT	0																									0
4	2	97	ACT 4 TABLE SKIP IT	0																									0
4	2	98	ACT 4 TABLE SKIP IT	0																									0
4	2	99	ACT 4 TABLE SKIP IT	0																									0
4	2	100	ACT 4 TABLE SKIP IT	0																									0
4	2	101	ACT 4 TABLE SKIP IT	0																									0
4	2	102	ACT 4 TABLE SKIP IT	0																									0
4	2	103	ACT 4 TABLE SKIP IT	0																									0
4	2	104	ACT 4 TABLE SKIP IT	0																									0
4	2	105	ACT 4 TABLE SKIP IT	0																									0
4	2	106	ACT 4 TABLE SKIP IT	0																									0
4	2	107	ACT 4 TABLE SKIP IT	0																									0
4	2	108	ACT 4 TABLE SKIP IT	0																									0
4	2	109	ACT 4 TABLE SKIP IT	0																									0
4	2	110	ACT 4 TABLE SKIP IT	0																									0
4	2	111	ACT 4 TABLE SKIP IT	0																									0
4	2	112	ACT 4 TABLE SKIP IT	0																									0
4	2	113	ACT 4 TABLE SKIP IT	0																									0
4	2	114	ACT 4 TABLE SKIP IT	0																									0
4	2	115	ACT 4 TABLE SKIP IT	0																									0
4	2	116	ACT 4 TABLE SKIP IT	0																									0
4	2	117	ACT 4 TABLE SKIP IT	0																									0
4	2	118	ACT 4 TABLE SKIP IT	0																									0
4	2	119	ACT 4 TABLE SKIP IT	0																									0
4	2	120	ACT 4 TABLE SKIP IT	0																									0
4	2	121	ACT 4 TABLE SKIP IT	0																									0
4	2	122	ACT 4 TABLE SKIP IT	0																									0
4	2	123	ACT 4 TABLE SKIP IT	0																									0
4	2	124	ACT 4 TABLE SKIP IT	0																									0
4	2	125	ACT 4 TABLE SKIP IT	0																									0
4	2	126	ACT 4 TABLE SKIP IT	0																									0
4	2	127	ACT 4 TABLE SKIP IT	0																									0
4	2	128	ACT 4 TABLE SKIP IT	0																									0
4	2	129	ACT 4 TABLE SKIP IT	0																									0
4	2	130	ACT 4 TABLE SKIP IT	0																									0
4	2	131	ACT 4 TABLE SKIP IT	0																									0
4	2	132	ACT 4 TABLE SKIP IT	0																									0
4	2	133	ACT 4 TABLE SKIP IT	0																									0
4	2	134	ACT 4 TABLE SKIP IT	0																									0
4	2	135	ACT 4 TABLE SKIP IT	0																									0
4	2	136	ACT 4 TABLE SKIP IT	0																									0
4	2	137	ACT 4 TABLE SKIP IT	0																									0
4	2	138	ACT 4 TABLE SKIP IT	0																									0
4	2	139	ACT 4 TABLE SKIP IT	0																									0
4	2	140	ACT 4 TABLE SKIP IT	0																									0
4	2	141	ACT 4 TABLE SKIP IT	0																									0
4	2	142	ACT 4 TABLE SKIP IT	0																									0
4	2	143	ACT 4 TABLE SKIP IT	0																									0
4	2	144	ACT 4 TABLE SKIP IT	0																									0
4	2	145	ACT 4 TABLE SKIP IT	0																									0
4	2	146	ACT 4 TABLE SKIP IT	0																									0
4	2	147	ACT 4 TABLE SKIP IT	0																									0
4	2	148	ACT 4 TABLE SKIP IT	0																									0
4	2	149	ACT 4 TABLE SKIP IT	0																									0
4	2	150	Dummy-test data SKIPT IT				/Data/Global/Objects	NU0																					
4	2	151	Casket-Casket #5				/Data/Global/Objects	C5	OP	HTH		LIT																	
4	2	152	Shrine-Shrine				/Data/Global/Objects	SF	OP	HTH		LIT																	
4	2	153	Casket-Casket #6				/Data/Global/Objects	C6	OP	HTH		LIT																	
4	2	154	LargeUrn-Urn #1				/Data/Global/Objects	U1	OP	HTH		LIT																	
4	2	155	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
4	2	156	chest-LargeChestL				/Data/Global/Objects	L2	OP	HTH		LIT																	
4	2	157	Barrel-Barrel				/Data/Global/Objects	B1	OP	HTH		LIT																	
4	2	158	TowerTome-Tower Tome				/Data/Global/Objects	TT	OP	HTH		LIT																	
4	2	159	Urn-Urn #2				/Data/Global/Objects	U2	OP	HTH		LIT																	
4	2	160	Dummy-Bench				/Data/Global/Objects	BE	NU	HTH		LIT																	
4	2	161	Barrel-BarrelExploding				/Data/Global/Objects	BX	OP	HTH		LIT							LIT	LIT									
4	2	162	Dummy-RogueFountain				/Data/Global/Objects	FN	NU	HTH		LIT																	
4	2	163	Door-Door Gate Left				/Data/Global/Objects	D1	OP	HTH		LIT																	
4	2	164	Door-Door Gate Right				/Data/Global/Objects	D2	OP	HTH		LIT																	
4	2	165	Door-Door Wooden Left				/Data/Global/Objects	D3	OP	HTH		LIT																	
4	2	166	Door-Door Wooden Right				/Data/Global/Objects	D4	OP	HTH		LIT																	
4	2	167	StoneAlpha-StoneAlpha				/Data/Global/Objects	S1	OP	HTH		LIT																	
4	2	168	StoneBeta-StoneBeta				/Data/Global/Objects	S2	OP	HTH		LIT																	
4	2	169	StoneGamma-StoneGamma				/Data/Global/Objects	S3	OP	HTH		LIT																	
4	2	170	StoneDelta-StoneDelta				/Data/Global/Objects	S4	OP	HTH		LIT																	
4	2	171	StoneLambda-StoneLambda				/Data/Global/Objects	S5	OP	HTH		LIT																	
4	2	172	StoneTheta-StoneTheta				/Data/Global/Objects	S6	OP	HTH		LIT																	
4	2	173	Door-Door Courtyard Left				/Data/Global/Objects	D5	OP	HTH		LIT																	
4	2	174	Door-Door Courtyard Right				/Data/Global/Objects	D6	OP	HTH		LIT																	
4	2	175	Door-Door Cathedral Double				/Data/Global/Objects	D7	OP	HTH		LIT																	
4	2	176	Gibbet-Cain's Been Captured				/Data/Global/Objects	GI	OP	HTH		LIT																	
4	2	177	Door-Door Monastery Double Right				/Data/Global/Objects	D8	OP	HTH		LIT																	
4	2	178	HoleAnim-Hole in Ground				/Data/Global/Objects	HI	OP	HTH		LIT																	
4	2	179	Dummy-Brazier				/Data/Global/Objects	BR	ON	HTH		LIT							LIT										
4	2	180	Inifuss-inifuss tree				/Data/Global/Objects	IT	NU	HTH		LIT																	
4	2	181	Dummy-Fountain				/Data/Global/Objects	BF	NU	HTH		LIT																	
4	2	182	Dummy-crucifix				/Data/Global/Objects	CL	NU	HTH		LIT																	
4	2	183	Dummy-Candles1				/Data/Global/Objects	A1	NU	HTH		LIT																	
4	2	184	Dummy-Candles2				/Data/Global/Objects	A2	NU	HTH		LIT																	
4	2	185	Dummy-Standard1				/Data/Global/Objects	N1	NU	HTH		LIT																	
4	2	186	Dummy-Standard2				/Data/Global/Objects	N2	NU	HTH		LIT																	
4	2	187	Dummy-Torch1 Tiki				/Data/Global/Objects	TO	ON	HTH		LIT																	
4	2	188	Dummy-Torch2 Wall				/Data/Global/Objects	WT	ON	HTH		LIT																	
4	2	189	fire-RogueBonfire				/Data/Global/Objects	RB	ON	HTH		LIT																	
4	2	190	Dummy-River1				/Data/Global/Objects	R1	NU	HTH		LIT																	
4	2	191	Dummy-River2				/Data/Global/Objects	R2	NU	HTH		LIT																	
4	2	192	Dummy-River3				/Data/Global/Objects	R3	NU	HTH		LIT																	
4	2	193	Dummy-River4				/Data/Global/Objects	R4	NU	HTH		LIT																	
4	2	194	Dummy-River5				/Data/Global/Objects	R5	NU	HTH		LIT																	
4	2	195	AmbientSound-ambient sound generator				/Data/Global/Objects	S1	OP	HTH		LIT																	
4	2	196	Crate-Crate				/Data/Global/Objects	CT	OP	HTH		LIT																	
4	2	197	Door-Andariel's Door				/Data/Global/Objects	AD	NU	HTH		LIT																	
4	2	198	Dummy-RogueTorch				/Data/Global/Objects	T1	NU	HTH		LIT																	
4	2	199	Dummy-RogueTorch				/Data/Global/Objects	T2	NU	HTH		LIT																	
4	2	200	Casket-CasketR				/Data/Global/Objects	C1	OP	HTH		LIT																	
4	2	201	Casket-CasketL				/Data/Global/Objects	C2	OP	HTH		LIT																	
4	2	202	Urn-Urn #3				/Data/Global/Objects	U3	OP	HTH		LIT																	
4	2	203	Casket-Casket				/Data/Global/Objects	C4	OP	HTH		LIT																	
4	2	204	RogueCorpse-Rogue corpse 1				/Data/Global/Objects	Z1	NU	HTH		LIT																	
4	2	205	RogueCorpse-Rogue corpse 2				/Data/Global/Objects	Z2	NU	HTH		LIT																	
4	2	206	RogueCorpse-rolling rogue corpse				/Data/Global/Objects	Z5	OP	HTH		LIT																	
4	2	207	CorpseOnStick-rogue on a stick 1				/Data/Global/Objects	Z3	OP	HTH		LIT																	
4	2	208	CorpseOnStick-rogue on a stick 2				/Data/Global/Objects	Z4	OP	HTH		LIT																	
4	2	209	Portal-Town portal				/Data/Global/Objects	TP	ON	HTH	LIT	LIT																	
4	2	210	Portal-Permanent town portal				/Data/Global/Objects	PP	ON	HTH	LIT	LIT																	
4	2	211	Dummy-Invisible object				/Data/Global/Objects	SS																					
4	2	212	Door-Door Cathedral Left				/Data/Global/Objects	D9	OP	HTH		LIT																	
4	2	213	Door-Door Cathedral Right				/Data/Global/Objects	DA	OP	HTH		LIT																	
4	2	214	Door-Door Wooden Left #2				/Data/Global/Objects	DB	OP	HTH		LIT																	
4	2	215	Dummy-invisible river sound1				/Data/Global/Objects	X1																					
4	2	216	Dummy-invisible river sound2				/Data/Global/Objects	X2																					
4	2	217	Dummy-ripple				/Data/Global/Objects	1R	NU	HTH		LIT																	
4	2	218	Dummy-ripple				/Data/Global/Objects	2R	NU	HTH		LIT																	
4	2	219	Dummy-ripple				/Data/Global/Objects	3R	NU	HTH		LIT																	
4	2	220	Dummy-ripple				/Data/Global/Objects	4R	NU	HTH		LIT																	
4	2	221	Dummy-forest night sound #1				/Data/Global/Objects	F1																					
4	2	222	Dummy-forest night sound #2				/Data/Global/Objects	F2																					
4	2	223	Dummy-yeti dung				/Data/Global/Objects	YD	NU	HTH		LIT																	
4	2	224	TrappDoor-Trap Door				/Data/Global/Objects	TD	ON	HTH		LIT																	
4	2	225	Door-Door by Dock, Act 2				/Data/Global/Objects	DD	ON	HTH		LIT																	
4	2	226	Dummy-sewer drip				/Data/Global/Objects	SZ																					
4	2	227	Shrine-healthorama				/Data/Global/Objects	SH	OP	HTH		LIT																	
4	2	228	Dummy-invisible town sound				/Data/Global/Objects	TA																					
4	2	229	Casket-casket #3				/Data/Global/Objects	C3	OP	HTH		LIT																	
4	2	230	Obelisk-obelisk				/Data/Global/Objects	OB	OP	HTH		LIT																	
4	2	231	Shrine-forest altar				/Data/Global/Objects	AF	OP	HTH		LIT																	
4	2	232	Dummy-bubbling pool of blood				/Data/Global/Objects	B2	NU	HTH		LIT																	
4	2	233	Shrine-horn shrine				/Data/Global/Objects	HS	OP	HTH		LIT																	
4	2	234	Shrine-healing well				/Data/Global/Objects	HW	OP	HTH		LIT																	
4	2	235	Shrine-bull shrine,health, tombs				/Data/Global/Objects	BC	OP	HTH		LIT																	
4	2	236	Dummy-stele,magic shrine, stone, desert				/Data/Global/Objects	SG	OP	HTH		LIT																	
4	2	237	Chest3-tombchest 1, largechestL				/Data/Global/Objects	CA	OP	HTH		LIT																	
4	2	238	Chest3-tombchest 2 largechestR				/Data/Global/Objects	CB	OP	HTH		LIT																	
4	2	239	Sarcophagus-mummy coffinL, tomb				/Data/Global/Objects	MC	OP	HTH		LIT																	
4	2	240	Obelisk-desert obelisk				/Data/Global/Objects	DO	OP	HTH		LIT																	
4	2	241	Door-tomb door left				/Data/Global/Objects	TL	OP	HTH		LIT																	
4	2	242	Door-tomb door right				/Data/Global/Objects	TR	OP	HTH		LIT																	
4	2	243	Shrine-mana shrineforinnerhell				/Data/Global/Objects	iz	OP	HTH		LIT							LIT										
4	2	244	LargeUrn-Urn #4				/Data/Global/Objects	U4	OP	HTH		LIT																	
4	2	245	LargeUrn-Urn #5				/Data/Global/Objects	U5	OP	HTH		LIT																	
4	2	246	Shrine-health shrineforinnerhell				/Data/Global/Objects	iy	OP	HTH		LIT							LIT										
4	2	247	Shrine-innershrinehell				/Data/Global/Objects	ix	OP	HTH		LIT							LIT										
4	2	248	Door-tomb door left 2				/Data/Global/Objects	TS	OP	HTH		LIT																	
4	2	249	Door-tomb door right 2				/Data/Global/Objects	TU	OP	HTH		LIT																	
4	2	250	Duriel's Lair-Portal to Duriel's Lair				/Data/Global/Objects	SJ	OP	HTH		LIT																	
4	2	251	Dummy-Brazier3				/Data/Global/Objects	B3	OP	HTH		LIT							LIT										
4	2	252	Dummy-Floor brazier				/Data/Global/Objects	FB	ON	HTH		LIT							LIT										
4	2	253	Dummy-flies				/Data/Global/Objects	FL	NU	HTH		LIT																	
4	2	254	ArmorStand-Armor Stand 1R				/Data/Global/Objects	A3	NU	HTH		LIT																	
4	2	255	ArmorStand-Armor Stand 2L				/Data/Global/Objects	A4	NU	HTH		LIT																	
4	2	256	WeaponRack-Weapon Rack 1R				/Data/Global/Objects	W1	NU	HTH		LIT																	
4	2	257	WeaponRack-Weapon Rack 2L				/Data/Global/Objects	W2	NU	HTH		LIT																	
4	2	258	Malus-Malus				/Data/Global/Objects	HM	NU	HTH		LIT																	
4	2	259	Shrine-palace shrine, healthR, harom, arcane Sanctuary				/Data/Global/Objects	P2	OP	HTH		LIT																	
4	2	260	not used-drinker				/Data/Global/Objects	n5	S1	HTH		LIT																	
4	2	261	well-Fountain 1				/Data/Global/Objects	F3	OP	HTH		LIT																	
4	2	262	not used-gesturer				/Data/Global/Objects	n6	S1	HTH		LIT																	
4	2	263	well-Fountain 2, well, desert, tomb				/Data/Global/Objects	F4	OP	HTH		LIT																	
4	2	264	not used-turner				/Data/Global/Objects	n7	S1	HTH		LIT																	
4	2	265	well-Fountain 3				/Data/Global/Objects	F5	OP	HTH		LIT																	
4	2	266	Shrine-snake woman, magic shrine, tomb, arcane sanctuary				/Data/Global/Objects	SN	OP	HTH		LIT							LIT										
4	2	267	Dummy-jungle torch				/Data/Global/Objects	JT	ON	HTH		LIT							LIT										
4	2	268	Well-Fountain 4				/Data/Global/Objects	F6	OP	HTH		LIT																	
4	2	269	Waypoint-waypoint portal				/Data/Global/Objects	wp	ON	HTH		LIT							LIT										
4	2	270	Dummy-healthshrine, act 3, dungeun				/Data/Global/Objects	dj	OP	HTH		LIT																	
4	2	271	jerhyn-placeholder #1				/Data/Global/Objects	ss																					
4	2	272	jerhyn-placeholder #2				/Data/Global/Objects	ss																					
4	2	273	Shrine-innershrinehell2				/Data/Global/Objects	iw	OP	HTH		LIT							LIT										
4	2	274	Shrine-innershrinehell3				/Data/Global/Objects	iv	OP	HTH		LIT																	
4	2	275	hidden stash-ihobject3 inner hell				/Data/Global/Objects	iu	OP	HTH		LIT																	
4	2	276	skull pile-skullpile inner hell				/Data/Global/Objects	is	OP	HTH		LIT																	
4	2	277	hidden stash-ihobject5 inner hell				/Data/Global/Objects	ir	OP	HTH		LIT																	
4	2	278	hidden stash-hobject4 inner hell				/Data/Global/Objects	hg	OP	HTH		LIT																	
4	2	279	Door-secret door 1				/Data/Global/Objects	h2	OP	HTH		LIT																	
4	2	280	Well-pool act 1 wilderness				/Data/Global/Objects	zw	NU	HTH		LIT																	
4	2	281	Dummy-vile dog afterglow				/Data/Global/Objects	9b	OP	HTH		LIT																	
4	2	282	Well-cathedralwell act 1 inside				/Data/Global/Objects	zc	NU	HTH		LIT																	
4	2	283	shrine-shrine1_arcane sanctuary				/Data/Global/Objects	xx																					
4	2	284	shrine-dshrine2 act 2 shrine				/Data/Global/Objects	zs	OP	HTH		LIT							LIT										
4	2	285	shrine-desertshrine3 act 2 shrine				/Data/Global/Objects	zr	OP	HTH		LIT																	
4	2	286	shrine-dshrine1 act 2 shrine				/Data/Global/Objects	zd	OP	HTH		LIT																	
4	2	287	Well-desertwell act 2 well, desert, tomb				/Data/Global/Objects	zl	NU	HTH		LIT																	
4	2	288	Well-cavewell act 1 caves 				/Data/Global/Objects	zy	NU	HTH		LIT																	
4	2	289	chest-chest-r-large act 1				/Data/Global/Objects	q1	OP	HTH		LIT																	
4	2	290	chest-chest-r-tallskinney act 1				/Data/Global/Objects	q2	OP	HTH		LIT																	
4	2	291	chest-chest-r-med act 1				/Data/Global/Objects	q3	OP	HTH		LIT																	
4	2	292	jug-jug1 act 2, desert				/Data/Global/Objects	q4	OP	HTH		LIT																	
4	2	293	jug-jug2 act 2, desert				/Data/Global/Objects	q5	OP	HTH		LIT																	
4	2	294	chest-Lchest1 act 1				/Data/Global/Objects	q6	OP	HTH		LIT																	
4	2	295	Waypoint-waypointi inner hell				/Data/Global/Objects	wi	ON	HTH		LIT							LIT										
4	2	296	chest-dchest2R act 2, desert, tomb, chest-r-med				/Data/Global/Objects	q9	OP	HTH		LIT																	
4	2	297	chest-dchestr act 2, desert, tomb, chest -r large				/Data/Global/Objects	q7	OP	HTH		LIT																	
4	2	298	chest-dchestL act 2, desert, tomb chest l large				/Data/Global/Objects	q8	OP	HTH		LIT																	
4	2	299	taintedsunaltar-tainted sun altar quest				/Data/Global/Objects	za	OP	HTH		LIT							LIT										
4	2	300	shrine-dshrine1 act 2 , desert				/Data/Global/Objects	zv	NU	HTH		LIT							LIT	LIT									
4	2	301	shrine-dshrine4 act 2, desert				/Data/Global/Objects	ze	OP	HTH		LIT							LIT										
4	2	302	orifice-Where you place the Horadric staff				/Data/Global/Objects	HA	NU	HTH		LIT																	
4	2	303	Door-tyrael's door				/Data/Global/Objects	DX	OP	HTH		LIT																	
4	2	304	corpse-guard corpse				/Data/Global/Objects	GC	OP	HTH		LIT																	
4	2	305	hidden stash-rock act 1 wilderness				/Data/Global/Objects	c7	OP	HTH		LIT																	
4	2	306	Waypoint-waypoint act 2				/Data/Global/Objects	wm	ON	HTH		LIT							LIT										
4	2	307	Waypoint-waypoint act 1 wilderness				/Data/Global/Objects	wn	ON	HTH		LIT							LIT										
4	2	308	skeleton-corpse				/Data/Global/Objects	cp	OP	HTH		LIT																	
4	2	309	hidden stash-rockb act 1 wilderness				/Data/Global/Objects	cq	OP	HTH		LIT																	
4	2	310	fire-fire small				/Data/Global/Objects	FX	NU	HTH		LIT																	
4	2	311	fire-fire medium				/Data/Global/Objects	FY	NU	HTH		LIT																	
4	2	312	fire-fire large				/Data/Global/Objects	FZ	NU	HTH		LIT																	
4	2	313	hiding spot-cliff act 1 wilderness				/Data/Global/Objects	cf	NU	HTH		LIT																	
4	2	314	Shrine-mana well1				/Data/Global/Objects	MB	OP	HTH		LIT																	
4	2	315	Shrine-mana well2				/Data/Global/Objects	MD	OP	HTH		LIT																	
4	2	316	Shrine-mana well3, act 2, tomb, 				/Data/Global/Objects	MF	OP	HTH		LIT																	
4	2	317	Shrine-mana well4, act 2, harom				/Data/Global/Objects	MH	OP	HTH		LIT																	
4	2	318	Shrine-mana well5				/Data/Global/Objects	MJ	OP	HTH		LIT																	
4	2	319	hollow log-log				/Data/Global/Objects	cz	NU	HTH		LIT																	
4	2	320	Shrine-jungle healwell act 3				/Data/Global/Objects	JH	OP	HTH		LIT																	
4	2	321	skeleton-corpseb				/Data/Global/Objects	sx	OP	HTH		LIT																	
4	2	322	Shrine-health well, health shrine, desert				/Data/Global/Objects	Mk	OP	HTH		LIT																	
4	2	323	Shrine-mana well7, mana shrine, desert				/Data/Global/Objects	Mi	OP	HTH		LIT																	
4	2	324	loose rock-rockc act 1 wilderness				/Data/Global/Objects	RY	OP	HTH		LIT																	
4	2	325	loose boulder-rockd act 1 wilderness				/Data/Global/Objects	RZ	OP	HTH		LIT																	
4	2	326	chest-chest-L-med				/Data/Global/Objects	c8	OP	HTH		LIT																	
4	2	327	chest-chest-L-large				/Data/Global/Objects	c9	OP	HTH		LIT																	
4	2	328	GuardCorpse-guard on a stick, desert, tomb, harom				/Data/Global/Objects	GS	OP	HTH		LIT																	
4	2	329	bookshelf-bookshelf1				/Data/Global/Objects	b4	OP	HTH		LIT																	
4	2	330	bookshelf-bookshelf2				/Data/Global/Objects	b5	OP	HTH		LIT																	
4	2	331	chest-jungle chest act 3				/Data/Global/Objects	JC	OP	HTH		LIT																	
4	2	332	coffin-tombcoffin				/Data/Global/Objects	tm	OP	HTH		LIT																	
4	2	333	chest-chest-L-med, jungle				/Data/Global/Objects	jz	OP	HTH		LIT																	
4	2	334	Shrine-jungle shrine2				/Data/Global/Objects	jy	OP	HTH		LIT							LIT	LIT									
4	2	335	stash-jungle object act3				/Data/Global/Objects	jx	OP	HTH		LIT																	
4	2	336	stash-jungle object act3				/Data/Global/Objects	jw	OP	HTH		LIT																	
4	2	337	stash-jungle object act3				/Data/Global/Objects	jv	OP	HTH		LIT																	
4	2	338	stash-jungle object act3				/Data/Global/Objects	ju	OP	HTH		LIT																	
4	2	339	Dummy-cain portal				/Data/Global/Objects	tP	OP	HTH	LIT	LIT																	
4	2	340	Shrine-jungle shrine3 act 3				/Data/Global/Objects	js	OP	HTH		LIT							LIT										
4	2	341	Shrine-jungle shrine4 act 3				/Data/Global/Objects	jr	OP	HTH		LIT							LIT										
4	2	342	teleport pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
4	2	343	LamTome-Lam Esen's Tome				/Data/Global/Objects	ab	OP	HTH		LIT																	
4	2	344	stair-stairsl				/Data/Global/Objects	sl																					
4	2	345	stair-stairsr				/Data/Global/Objects	sv																					
4	2	346	a trap-test data floortrap				/Data/Global/Objects	a5	OP	HTH		LIT																	
4	2	347	Shrine-jungleshrine act 3				/Data/Global/Objects	jq	OP	HTH		LIT							LIT										
4	2	348	chest-chest-L-tallskinney, general chest r?				/Data/Global/Objects	c0	OP	HTH		LIT																	
4	2	349	Shrine-mafistoshrine				/Data/Global/Objects	mz	OP	HTH		LIT							LIT										
4	2	350	Shrine-mafistoshrine				/Data/Global/Objects	my	OP	HTH		LIT							LIT										
4	2	351	Shrine-mafistoshrine				/Data/Global/Objects	mx	NU	HTH		LIT							LIT										
4	2	352	Shrine-mafistomana				/Data/Global/Objects	mw	OP	HTH		LIT							LIT										
4	2	353	stash-mafistolair				/Data/Global/Objects	mv	OP	HTH		LIT																	
4	2	354	stash-box				/Data/Global/Objects	mu	OP	HTH		LIT																	
4	2	355	stash-altar				/Data/Global/Objects	mt	OP	HTH		LIT																	
4	2	356	Shrine-mafistohealth				/Data/Global/Objects	mr	OP	HTH		LIT							LIT										
4	2	357	dummy-water rocks in act 3 wrok				/Data/Global/Objects	rw	NU	HTH		LIT																	
4	2	358	Basket-basket 1				/Data/Global/Objects	bd	OP	HTH		LIT																	
4	2	359	Basket-basket 2				/Data/Global/Objects	bj	OP	HTH		LIT																	
4	2	360	Dummy-water logs in act 3  ne logw				/Data/Global/Objects	lw	NU	HTH		LIT																	
4	2	361	Dummy-water rocks girl in act 3 wrob				/Data/Global/Objects	wb	NU	HTH		LIT																	
4	2	362	Dummy-bubbles in act3 water				/Data/Global/Objects	yb	NU	HTH		LIT																	
4	2	363	Dummy-water logs in act 3 logx				/Data/Global/Objects	wd	NU	HTH		LIT																	
4	2	364	Dummy-water rocks in act 3 rokb				/Data/Global/Objects	wc	NU	HTH		LIT																	
4	2	365	Dummy-water rocks girl in act 3 watc				/Data/Global/Objects	we	NU	HTH		LIT																	
4	2	366	Dummy-water rocks in act 3 waty				/Data/Global/Objects	wy	NU	HTH		LIT																	
4	2	367	Dummy-water logs in act 3  logz				/Data/Global/Objects	lx	NU	HTH		LIT																	
4	2	368	Dummy-web covered tree 1				/Data/Global/Objects	w3	NU	HTH		LIT							LIT										
4	2	369	Dummy-web covered tree 2				/Data/Global/Objects	w4	NU	HTH		LIT							LIT										
4	2	370	Dummy-web covered tree 3				/Data/Global/Objects	w5	NU	HTH		LIT							LIT										
4	2	371	Dummy-web covered tree 4				/Data/Global/Objects	w6	NU	HTH		LIT							LIT										
4	2	372	pillar-hobject1				/Data/Global/Objects	70	OP	HTH		LIT																	
4	2	373	cocoon-cacoon				/Data/Global/Objects	CN	OP	HTH		LIT																	
4	2	374	cocoon-cacoon 2				/Data/Global/Objects	CC	OP	HTH		LIT																	
4	2	375	skullpile-hobject1				/Data/Global/Objects	ib	OP	HTH		LIT																	
4	2	376	Shrine-outershrinehell				/Data/Global/Objects	ia	OP	HTH		LIT							LIT										
4	2	377	dummy-water rock girl act 3  nw  blgb				/Data/Global/Objects	QX	NU	HTH		LIT																	
4	2	378	dummy-big log act 3  sw blga				/Data/Global/Objects	qw	NU	HTH		LIT																	
4	2	379	door-slimedoor1				/Data/Global/Objects	SQ	OP	HTH		LIT																	
4	2	380	door-slimedoor2				/Data/Global/Objects	SY	OP	HTH		LIT																	
4	2	381	Shrine-outershrinehell2				/Data/Global/Objects	ht	OP	HTH		LIT							LIT										
4	2	382	Shrine-outershrinehell3				/Data/Global/Objects	hq	OP	HTH		LIT																	
4	2	383	pillar-hobject2				/Data/Global/Objects	hv	OP	HTH		LIT																	
4	2	384	dummy-Big log act 3 se blgc 				/Data/Global/Objects	Qy	NU	HTH		LIT																	
4	2	385	dummy-Big log act 3 nw blgd				/Data/Global/Objects	Qz	NU	HTH		LIT																	
4	2	386	Shrine-health wellforhell				/Data/Global/Objects	ho	OP	HTH		LIT																	
4	2	387	Waypoint-act3waypoint town				/Data/Global/Objects	wz	ON	HTH		LIT							LIT										
4	2	388	Waypoint-waypointh				/Data/Global/Objects	wv	ON	HTH		LIT							LIT										
4	2	389	body-burning town				/Data/Global/Objects	bz	ON	HTH		LIT							LIT										
4	2	390	chest-gchest1L general				/Data/Global/Objects	cy	OP	HTH		LIT																	
4	2	391	chest-gchest2R general				/Data/Global/Objects	cx	OP	HTH		LIT																	
4	2	392	chest-gchest3R general				/Data/Global/Objects	cu	OP	HTH		LIT																	
4	2	393	chest-glchest3L general				/Data/Global/Objects	cd	OP	HTH		LIT																	
4	2	394	ratnest-sewers				/Data/Global/Objects	rn	OP	HTH		LIT																	
4	2	395	body-burning town				/Data/Global/Objects	by	NU	HTH		LIT							LIT										
4	2	396	ratnest-sewers				/Data/Global/Objects	ra	OP	HTH		LIT																	
4	2	397	bed-bed act 1				/Data/Global/Objects	qa	OP	HTH		LIT																	
4	2	398	bed-bed act 1				/Data/Global/Objects	qb	OP	HTH		LIT																	
4	2	399	manashrine-mana wellforhell				/Data/Global/Objects	hn	OP	HTH		LIT							LIT										
4	2	400	a trap-exploding cow  for Tristan and ACT 3 only??Very Rare  1 or 2				/Data/Global/Objects	ew	OP	HTH		LIT																	
4	2	401	gidbinn altar-gidbinn altar				/Data/Global/Objects	ga	ON	HTH		LIT							LIT										
4	2	402	gidbinn-gidbinn decoy				/Data/Global/Objects	gd	ON	HTH		LIT							LIT										
4	2	403	Dummy-diablo right light				/Data/Global/Objects	11	NU	HTH		LIT																	
4	2	404	Dummy-diablo left light				/Data/Global/Objects	12	NU	HTH		LIT																	
4	2	405	Dummy-diablo start point				/Data/Global/Objects	ss																					
4	2	406	Dummy-stool for act 1 cabin				/Data/Global/Objects	s9	NU	HTH		LIT																	
4	2	407	Dummy-wood for act 1 cabin				/Data/Global/Objects	wg	NU	HTH		LIT																	
4	2	408	Dummy-more wood for act 1 cabin				/Data/Global/Objects	wh	NU	HTH		LIT																	
4	2	409	Dummy-skeleton spawn for hell   facing nw				/Data/Global/Objects	QS	OP	HTH		LIT							LIT										
4	2	410	Shrine-holyshrine for monastery,catacombs,jail				/Data/Global/Objects	HL	OP	HTH		LIT							LIT										
4	2	411	a trap-spikes for tombs floortrap				/Data/Global/Objects	A7	OP	HTH		LIT																	
4	2	412	Shrine-act 1 cathedral				/Data/Global/Objects	s0	OP	HTH		LIT							LIT										
4	2	413	Shrine-act 1 jail				/Data/Global/Objects	jb	NU	HTH		LIT							LIT										
4	2	414	Shrine-act 1 jail				/Data/Global/Objects	jd	OP	HTH		LIT							LIT										
4	2	415	Shrine-act 1 jail				/Data/Global/Objects	jf	OP	HTH		LIT							LIT										
4	2	416	goo pile-goo pile for sand maggot lair				/Data/Global/Objects	GP	OP	HTH		LIT																	
4	2	417	bank-bank				/Data/Global/Objects	b6	NU	HTH		LIT																	
4	2	418	wirt's body-wirt's body				/Data/Global/Objects	BP	NU	HTH		LIT																	
4	2	419	dummy-gold placeholder				/Data/Global/Objects	1g																					
4	2	420	corpse-guard corpse 2				/Data/Global/Objects	GF	OP	HTH		LIT																	
4	2	421	corpse-dead villager 1				/Data/Global/Objects	dg	OP	HTH		LIT																	
4	2	422	corpse-dead villager 2				/Data/Global/Objects	df	OP	HTH		LIT																	
4	2	423	Dummy-yet another flame, no damage				/Data/Global/Objects	f8	NU	HTH		LIT																	
4	2	424	hidden stash-tiny pixel shaped thingie				/Data/Global/Objects	f9																					
4	2	425	Shrine-health shrine for caves				/Data/Global/Objects	ce	OP	HTH		LIT																	
4	2	426	Shrine-mana shrine for caves				/Data/Global/Objects	cg	OP	HTH		LIT																	
4	2	427	Shrine-cave magic shrine				/Data/Global/Objects	cg	OP	HTH		LIT																	
4	2	428	Shrine-manashrine, act 3, dungeun				/Data/Global/Objects	de	OP	HTH		LIT																	
4	2	429	Shrine-magic shrine, act 3 sewers.				/Data/Global/Objects	wj	NU	HTH		LIT							LIT	LIT									
4	2	430	Shrine-healthwell, act 3, sewers				/Data/Global/Objects	wk	OP	HTH		LIT																	
4	2	431	Shrine-manawell, act 3, sewers				/Data/Global/Objects	wl	OP	HTH		LIT																	
4	2	432	Shrine-magic shrine, act 3 sewers, dungeon.				/Data/Global/Objects	ws	NU	HTH		LIT							LIT	LIT									
4	2	433	dummy-brazier_celler, act 2				/Data/Global/Objects	bi	NU	HTH		LIT							LIT										
4	2	434	sarcophagus-anubis coffin, act2, tomb				/Data/Global/Objects	qc	OP	HTH		LIT																	
4	2	435	dummy-brazier_general, act 2, sewers, tomb, desert				/Data/Global/Objects	bm	NU	HTH		LIT							LIT										
4	2	436	Dummy-brazier_tall, act 2, desert, town, tombs				/Data/Global/Objects	bo	NU	HTH		LIT							LIT										
4	2	437	Dummy-brazier_small, act 2, desert, town, tombs				/Data/Global/Objects	bq	NU	HTH		LIT							LIT										
4	2	438	Waypoint-waypoint, celler				/Data/Global/Objects	w7	ON	HTH		LIT							LIT										
4	2	439	bed-bed for harum				/Data/Global/Objects	ub	OP	HTH		LIT																	
4	2	440	door-iron grate door left				/Data/Global/Objects	dv	NU	HTH		LIT																	
4	2	441	door-iron grate door right				/Data/Global/Objects	dn	NU	HTH		LIT																	
4	2	442	door-wooden grate door left				/Data/Global/Objects	dp	NU	HTH		LIT																	
4	2	443	door-wooden grate door right				/Data/Global/Objects	dt	NU	HTH		LIT																	
4	2	444	door-wooden door left				/Data/Global/Objects	dk	NU	HTH		LIT																	
4	2	445	door-wooden door right				/Data/Global/Objects	dl	NU	HTH		LIT																	
4	2	446	Dummy-wall torch left for tombs				/Data/Global/Objects	qd	NU	HTH		LIT							LIT										
4	2	447	Dummy-wall torch right for tombs				/Data/Global/Objects	qe	NU	HTH		LIT							LIT										
4	2	448	portal-arcane sanctuary portal				/Data/Global/Objects	ay	ON	HTH		LIT							LIT	LIT									
4	2	449	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hb	OP	HTH		LIT							LIT										
4	2	450	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hc	OP	HTH		LIT							LIT										
4	2	451	Dummy-maggot well health				/Data/Global/Objects	qf	OP	HTH		LIT																	
4	2	452	manashrine-maggot well mana				/Data/Global/Objects	qg	OP	HTH		LIT																	
4	2	453	magic shrine-magic shrine, act 3 arcane sanctuary.				/Data/Global/Objects	hd	OP	HTH		LIT							LIT										
4	2	454	teleportation pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
4	2	455	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
4	2	456	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
4	2	457	Dummy-arcane thing				/Data/Global/Objects	7a	NU	HTH		LIT																	
4	2	458	Dummy-arcane thing				/Data/Global/Objects	7b	NU	HTH		LIT																	
4	2	459	Dummy-arcane thing				/Data/Global/Objects	7c	NU	HTH		LIT																	
4	2	460	Dummy-arcane thing				/Data/Global/Objects	7d	NU	HTH		LIT																	
4	2	461	Dummy-arcane thing				/Data/Global/Objects	7e	NU	HTH		LIT																	
4	2	462	Dummy-arcane thing				/Data/Global/Objects	7f	NU	HTH		LIT																	
4	2	463	Dummy-arcane thing				/Data/Global/Objects	7g	NU	HTH		LIT																	
4	2	464	dead guard-harem guard 1				/Data/Global/Objects	qh	NU	HTH		LIT																	
4	2	465	dead guard-harem guard 2				/Data/Global/Objects	qi	NU	HTH		LIT																	
4	2	466	dead guard-harem guard 3				/Data/Global/Objects	qj	NU	HTH		LIT																	
4	2	467	dead guard-harem guard 4				/Data/Global/Objects	qk	NU	HTH		LIT																	
4	2	468	eunuch-harem blocker				/Data/Global/Objects	ss																					
4	2	469	Dummy-healthwell, act 2, arcane				/Data/Global/Objects	ax	OP	HTH		LIT																	
4	2	470	manashrine-healthwell, act 2, arcane				/Data/Global/Objects	au	OP	HTH		LIT																	
4	2	471	Dummy-test data				/Data/Global/Objects	pp	S1	HTH	LIT	LIT																	
4	2	472	Well-tombwell act 2 well, tomb				/Data/Global/Objects	hu	NU	HTH		LIT																	
4	2	473	Waypoint-waypoint act2 sewer				/Data/Global/Objects	qm	ON	HTH		LIT							LIT										
4	2	474	Waypoint-waypoint act3 travincal				/Data/Global/Objects	ql	ON	HTH		LIT							LIT										
4	2	475	magic shrine-magic shrine, act 3, sewer				/Data/Global/Objects	qn	NU	HTH		LIT							LIT										
4	2	476	dead body-act3, sewer				/Data/Global/Objects	qo	OP	HTH		LIT																	
4	2	477	dummy-torch (act 3 sewer) stra				/Data/Global/Objects	V1	NU	HTH		LIT							LIT										
4	2	478	dummy-torch (act 3 kurast) strb				/Data/Global/Objects	V2	NU	HTH		LIT							LIT										
4	2	479	chest-mafistochestlargeLeft				/Data/Global/Objects	xb	OP	HTH		LIT																	
4	2	480	chest-mafistochestlargeright				/Data/Global/Objects	xc	OP	HTH		LIT																	
4	2	481	chest-mafistochestmedleft				/Data/Global/Objects	xd	OP	HTH		LIT																	
4	2	482	chest-mafistochestmedright				/Data/Global/Objects	xe	OP	HTH		LIT																	
4	2	483	chest-spiderlairchestlargeLeft				/Data/Global/Objects	xf	OP	HTH		LIT																	
4	2	484	chest-spiderlairchesttallLeft				/Data/Global/Objects	xg	OP	HTH		LIT																	
4	2	485	chest-spiderlairchestmedright				/Data/Global/Objects	xh	OP	HTH		LIT																	
4	2	486	chest-spiderlairchesttallright				/Data/Global/Objects	xi	OP	HTH		LIT																	
4	2	487	Steeg Stone-steeg stone				/Data/Global/Objects	y6	NU	HTH		LIT							LIT										
4	2	488	Guild Vault-guild vault				/Data/Global/Objects	y4	NU	HTH		LIT																	
4	2	489	Trophy Case-trophy case				/Data/Global/Objects	y2	NU	HTH		LIT																	
4	2	490	Message Board-message board				/Data/Global/Objects	y3	NU	HTH		LIT																	
4	2	491	Dummy-mephisto bridge				/Data/Global/Objects	xj	OP	HTH		LIT																	
4	2	492	portal-hellgate				/Data/Global/Objects	1y	ON	HTH		LIT								LIT	LIT								
4	2	493	Shrine-manawell, act 3, kurast				/Data/Global/Objects	xl	OP	HTH		LIT																	
4	2	494	Shrine-healthwell, act 3, kurast				/Data/Global/Objects	xm	OP	HTH		LIT																	
4	2	495	Dummy-hellfire1				/Data/Global/Objects	e3	NU	HTH		LIT																	
4	2	496	Dummy-hellfire2				/Data/Global/Objects	e4	NU	HTH		LIT																	
4	2	497	Dummy-hellfire3				/Data/Global/Objects	e5	NU	HTH		LIT																	
4	2	498	Dummy-helllava1				/Data/Global/Objects	e6	NU	HTH		LIT																	
4	2	499	Dummy-helllava2				/Data/Global/Objects	e7	NU	HTH		LIT																	
4	2	500	Dummy-helllava3				/Data/Global/Objects	e8	NU	HTH		LIT																	
4	2	501	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
4	2	502	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
4	2	503	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
4	2	504	chest-horadric cube chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
4	2	505	chest-horadric scroll chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
4	2	506	chest-staff of kings chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
4	2	507	Tome-yet another tome				/Data/Global/Objects	TT	NU	HTH		LIT																	
4	2	508	fire-hell brazier				/Data/Global/Objects	E1	NU	HTH	LIT	LIT																	
4	2	509	fire-hell brazier				/Data/Global/Objects	E2	NU	HTH	LIT	LIT																	
4	2	510	RockPIle-dungeon				/Data/Global/Objects	xn	OP	HTH		LIT																	
4	2	511	magic shrine-magic shrine, act 3,dundeon				/Data/Global/Objects	qo	OP	HTH		LIT																	
4	2	512	basket-dungeon				/Data/Global/Objects	xp	OP	HTH		LIT																	
4	2	513	HungSkeleton-outerhell skeleton				/Data/Global/Objects	jw	OP	HTH		LIT																	
4	2	514	Dummy-guy for dungeon				/Data/Global/Objects	ea	OP	HTH		LIT																	
4	2	515	casket-casket for Act 3 dungeon				/Data/Global/Objects	vb	OP	HTH		LIT																	
4	2	516	sewer stairs-stairs for act 3 sewer quest				/Data/Global/Objects	ve	OP	HTH		LIT																	
4	2	517	sewer lever-lever for act 3 sewer quest				/Data/Global/Objects	vf	OP	HTH		LIT																	
4	2	518	darkwanderer-start position				/Data/Global/Objects	ss																					
4	2	519	dummy-trapped soul placeholder				/Data/Global/Objects	ss																					
4	2	520	Dummy-torch for act3 town				/Data/Global/Objects	VG	NU	HTH		LIT							LIT										
4	2	521	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
4	2	522	BoneChest-innerhellbonepile				/Data/Global/Objects	y1	OP	HTH		LIT																	
4	2	523	Dummy-skeleton spawn for hell facing ne				/Data/Global/Objects	Qt	OP	HTH		LIT							LIT										
4	2	524	Dummy-fog act 3 water rfga				/Data/Global/Objects	ud	NU	HTH		LIT																	
4	2	525	Dummy-Not used				/Data/Global/Objects	xx																					
4	2	526	Hellforge-Forge  hell				/Data/Global/Objects	ux	ON	HTH		LIT							LIT	LIT	LIT								
4	2	527	Guild Portal-Portal to next guild level				/Data/Global/Objects	PP	NU	HTH	LIT	LIT																	
4	2	528	Dummy-hratli start				/Data/Global/Objects	ss																					
4	2	529	Dummy-hratli end				/Data/Global/Objects	ss																					
4	2	530	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	uy	OP	HTH		LIT							LIT										
4	2	531	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	15	OP	HTH		LIT							LIT										
4	2	532	Dummy-natalya start				/Data/Global/Objects	ss																					
4	2	533	TrappedSoul-guy stuck in hell				/Data/Global/Objects	18	OP	HTH		LIT																	
4	2	534	TrappedSoul-guy stuck in hell				/Data/Global/Objects	19	OP	HTH		LIT																	
4	2	535	Dummy-cain start position				/Data/Global/Objects	ss																					
4	2	536	Dummy-stairsr				/Data/Global/Objects	sv	OP	HTH		LIT																	
4	2	537	chest-arcanesanctuarybigchestLeft				/Data/Global/Objects	y7	OP	HTH		LIT																	
4	2	538	casket-arcanesanctuarycasket				/Data/Global/Objects	y8	OP	HTH		LIT																	
4	2	539	chest-arcanesanctuarybigchestRight				/Data/Global/Objects	y9	OP	HTH		LIT																	
4	2	540	chest-arcanesanctuarychestsmallLeft				/Data/Global/Objects	ya	OP	HTH		LIT																	
4	2	541	chest-arcanesanctuarychestsmallRight				/Data/Global/Objects	yc	OP	HTH		LIT																	
4	2	542	Seal-Diablo seal				/Data/Global/Objects	30	ON	HTH		LIT							LIT										
4	2	543	Seal-Diablo seal				/Data/Global/Objects	31	ON	HTH		LIT							LIT										
4	2	544	Seal-Diablo seal				/Data/Global/Objects	32	ON	HTH		LIT							LIT										
4	2	545	Seal-Diablo seal				/Data/Global/Objects	33	ON	HTH		LIT							LIT										
4	2	546	Seal-Diablo seal				/Data/Global/Objects	34	ON	HTH		LIT							LIT										
4	2	547	chest-sparklychest				/Data/Global/Objects	yf	OP	HTH		LIT																	
4	2	548	Waypoint-waypoint pandamonia fortress				/Data/Global/Objects	yg	ON	HTH		LIT							LIT										
4	2	549	fissure-fissure for act 4 inner hell				/Data/Global/Objects	fh	OP	HTH		LIT							LIT										
4	2	550	Dummy-brazier for act 4, hell mesa				/Data/Global/Objects	he	NU	HTH		LIT							LIT										
4	2	551	Dummy-smoke				/Data/Global/Objects	35	NU	HTH		LIT																	
4	2	552	Waypoint-waypoint valleywaypoint				/Data/Global/Objects	yi	ON	HTH		LIT							LIT										
4	2	553	fire-hell brazier				/Data/Global/Objects	9f	NU	HTH		LIT							LIT										
4	2	554	compellingorb-compelling orb				/Data/Global/Objects	55	NU	HTH		LIT							LIT	LIT									
4	2	555	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
4	2	556	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
4	2	557	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
4	2	558	Dummy-fortress brazier #1				/Data/Global/Objects	98	NU	HTH		LIT							LIT										
4	2	559	Dummy-fortress brazier #2				/Data/Global/Objects	99	NU	HTH		LIT							LIT										
4	2	560	Siege Control-To control siege machines				/Data/Global/Objects	zq	OP	HTH		LIT																	
4	2	561	ptox-Pot O Torch (level 1)				/Data/Global/Objects	px	NU	HTH		LIT							LIT	LIT									
4	2	562	pyox-fire pit  (level 1)				/Data/Global/Objects	py	NU	HTH		LIT							LIT										
4	2	563	chestR-expansion no snow				/Data/Global/Objects	6q	OP	HTH		LIT																	
4	2	564	Shrine3wilderness-expansion no snow				/Data/Global/Objects	6r	OP	HTH		LIT							LIT										
4	2	565	Shrine2wilderness-expansion no snow				/Data/Global/Objects	6s	NU	HTH		LIT							LIT										
4	2	566	hiddenstash-expansion no snow				/Data/Global/Objects	3w	OP	HTH		LIT																	
4	2	567	flag wilderness-expansion no snow				/Data/Global/Objects	ym	NU	HTH		LIT																	
4	2	568	barrel wilderness-expansion no snow				/Data/Global/Objects	yn	OP	HTH		LIT																	
4	2	569	barrel wilderness-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
4	2	570	woodchestL-expansion no snow				/Data/Global/Objects	yp	OP	HTH		LIT																	
4	2	571	Shrine3wilderness-expansion no snow				/Data/Global/Objects	yq	NU	HTH		LIT							LIT										
4	2	572	manashrine-expansion no snow				/Data/Global/Objects	yr	OP	HTH		LIT							LIT										
4	2	573	healthshrine-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
4	2	574	burialchestL-expansion no snow				/Data/Global/Objects	yt	OP	HTH		LIT																	
4	2	575	burialchestR-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
4	2	576	well-expansion no snow				/Data/Global/Objects	yv	NU	HTH		LIT																	
4	2	577	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yw	OP	HTH		LIT							LIT	LIT									
4	2	578	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yx	OP	HTH		LIT							LIT										
4	2	579	Waypoint-expansion no snow				/Data/Global/Objects	yy	ON	HTH		LIT							LIT										
4	2	580	ChestL-expansion no snow				/Data/Global/Objects	yz	OP	HTH		LIT																	
4	2	581	woodchestR-expansion no snow				/Data/Global/Objects	6a	OP	HTH		LIT																	
4	2	582	ChestSL-expansion no snow				/Data/Global/Objects	6b	OP	HTH		LIT																	
4	2	583	ChestSR-expansion no snow				/Data/Global/Objects	6c	OP	HTH		LIT																	
4	2	584	etorch1-expansion no snow				/Data/Global/Objects	6d	NU	HTH		LIT							LIT										
4	2	585	ecfra-camp fire				/Data/Global/Objects	2w	NU	HTH		LIT							LIT	LIT									
4	2	586	ettr-town torch				/Data/Global/Objects	2x	NU	HTH		LIT							LIT	LIT									
4	2	587	etorch2-expansion no snow				/Data/Global/Objects	6e	NU	HTH		LIT							LIT										
4	2	588	burningbodies-wilderness/siege				/Data/Global/Objects	6f	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
4	2	589	burningpit-wilderness/siege				/Data/Global/Objects	6g	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
4	2	590	tribal flag-wilderness/siege				/Data/Global/Objects	6h	NU	HTH		LIT																	
4	2	591	eflg-town flag				/Data/Global/Objects	2y	NU	HTH		LIT																	
4	2	592	chan-chandeleir				/Data/Global/Objects	2z	NU	HTH		LIT							LIT										
4	2	593	jar1-wilderness/siege				/Data/Global/Objects	6i	OP	HTH		LIT																	
4	2	594	jar2-wilderness/siege				/Data/Global/Objects	6j	OP	HTH		LIT																	
4	2	595	jar3-wilderness/siege				/Data/Global/Objects	6k	OP	HTH		LIT																	
4	2	596	swingingheads-wilderness				/Data/Global/Objects	6L	NU	HTH		LIT																	
4	2	597	pole-wilderness				/Data/Global/Objects	6m	NU	HTH		LIT																	
4	2	598	animated skulland rockpile-expansion no snow				/Data/Global/Objects	6n	OP	HTH		LIT																	
4	2	599	gate-town main gate				/Data/Global/Objects	2v	OP	HTH		LIT																	
4	2	600	pileofskullsandrocks-seige				/Data/Global/Objects	6o	NU	HTH		LIT																	
4	2	601	hellgate-seige				/Data/Global/Objects	6p	NU	HTH		LIT							LIT	LIT									
4	2	602	banner 1-preset in enemy camp				/Data/Global/Objects	ao	NU	HTH		LIT																	
4	2	603	banner 2-preset in enemy camp				/Data/Global/Objects	ap	NU	HTH		LIT																	
4	2	604	explodingchest-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
4	2	605	chest-specialchest				/Data/Global/Objects	6u	OP	HTH		LIT																	
4	2	606	deathpole-wilderness				/Data/Global/Objects	6v	NU	HTH		LIT																	
4	2	607	Ldeathpole-wilderness				/Data/Global/Objects	6w	NU	HTH		LIT																	
4	2	608	Altar-inside of temple				/Data/Global/Objects	6x	NU	HTH		LIT							LIT										
4	2	609	dummy-Drehya Start In Town				/Data/Global/Objects	ss																					
4	2	610	dummy-Drehya Start Outside Town				/Data/Global/Objects	ss																					
4	2	611	dummy-Nihlathak Start In Town				/Data/Global/Objects	ss																					
4	2	612	dummy-Nihlathak Start Outside Town				/Data/Global/Objects	ss																					
4	2	613	hidden stash-icecave_				/Data/Global/Objects	6y	OP	HTH		LIT																	
4	2	614	healthshrine-icecave_				/Data/Global/Objects	8a	OP	HTH		LIT																	
4	2	615	manashrine-icecave_				/Data/Global/Objects	8b	OP	HTH		LIT																	
4	2	616	evilurn-icecave_				/Data/Global/Objects	8c	OP	HTH		LIT																	
4	2	617	icecavejar1-icecave_				/Data/Global/Objects	8d	OP	HTH		LIT																	
4	2	618	icecavejar2-icecave_				/Data/Global/Objects	8e	OP	HTH		LIT																	
4	2	619	icecavejar3-icecave_				/Data/Global/Objects	8f	OP	HTH		LIT																	
4	2	620	icecavejar4-icecave_				/Data/Global/Objects	8g	OP	HTH		LIT																	
4	2	621	icecavejar4-icecave_				/Data/Global/Objects	8h	OP	HTH		LIT																	
4	2	622	icecaveshrine2-icecave_				/Data/Global/Objects	8i	NU	HTH		LIT							LIT										
4	2	623	cagedwussie1-caged fellow(A5-Prisonner)				/Data/Global/Objects	60	NU	HTH		LIT																	
4	2	624	Ancient Statue 3-statue				/Data/Global/Objects	60	NU	HTH		LIT																	
4	2	625	Ancient Statue 1-statue				/Data/Global/Objects	61	NU	HTH		LIT																	
4	2	626	Ancient Statue 2-statue				/Data/Global/Objects	62	NU	HTH		LIT																	
4	2	627	deadbarbarian-seige/wilderness				/Data/Global/Objects	8j	OP	HTH		LIT																	
4	2	628	clientsmoke-client smoke				/Data/Global/Objects	oz	NU	HTH		LIT																	
4	2	629	icecaveshrine2-icecave_				/Data/Global/Objects	8k	NU	HTH		LIT							LIT										
4	2	630	icecave_torch1-icecave_				/Data/Global/Objects	8L	NU	HTH		LIT							LIT										
4	2	631	icecave_torch2-icecave_				/Data/Global/Objects	8m	NU	HTH		LIT							LIT										
4	2	632	ttor-expansion tiki torch				/Data/Global/Objects	2p	NU	HTH		LIT							LIT										
4	2	633	manashrine-baals				/Data/Global/Objects	8n	OP	HTH		LIT																	
4	2	634	healthshrine-baals				/Data/Global/Objects	8o	OP	HTH		LIT																	
4	2	635	tomb1-baal's lair				/Data/Global/Objects	8p	OP	HTH		LIT																	
4	2	636	tomb2-baal's lair				/Data/Global/Objects	8q	OP	HTH		LIT																	
4	2	637	tomb3-baal's lair				/Data/Global/Objects	8r	OP	HTH		LIT																	
4	2	638	magic shrine-baal's lair				/Data/Global/Objects	8s	NU	HTH		LIT							LIT										
4	2	639	torch1-baal's lair				/Data/Global/Objects	8t	NU	HTH		LIT							LIT										
4	2	640	torch2-baal's lair				/Data/Global/Objects	8u	NU	HTH		LIT							LIT										
4	2	641	manashrine-snowy				/Data/Global/Objects	8v	OP	HTH		LIT							LIT										
4	2	642	healthshrine-snowy				/Data/Global/Objects	8w	OP	HTH		LIT							LIT										
4	2	643	well-snowy				/Data/Global/Objects	8x	NU	HTH		LIT																	
4	2	644	Waypoint-baals_waypoint				/Data/Global/Objects	8y	ON	HTH		LIT							LIT										
4	2	645	magic shrine-snowy_shrine3				/Data/Global/Objects	8z	NU	HTH		LIT							LIT										
4	2	646	Waypoint-wilderness_waypoint				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
4	2	647	magic shrine-snowy_shrine3				/Data/Global/Objects	5b	OP	HTH		LIT							LIT	LIT									
4	2	648	well-baalslair				/Data/Global/Objects	5c	NU	HTH		LIT																	
4	2	649	magic shrine2-baal's lair				/Data/Global/Objects	5d	NU	HTH		LIT							LIT										
4	2	650	object1-snowy				/Data/Global/Objects	5e	OP	HTH		LIT																	
4	2	651	woodchestL-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
4	2	652	woodchestR-snowy				/Data/Global/Objects	5g	OP	HTH		LIT																	
4	2	653	magic shrine-baals_shrine3				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
4	2	654	woodchest2L-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
4	2	655	woodchest2R-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
4	2	656	swingingheads-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
4	2	657	debris-snowy				/Data/Global/Objects	5l	NU	HTH		LIT																	
4	2	658	pene-Pen breakable door				/Data/Global/Objects	2q	NU	HTH		LIT																	
4	2	659	magic shrine-temple				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
4	2	660	mrpole-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
4	2	661	Waypoint-icecave 				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
4	2	662	magic shrine-temple				/Data/Global/Objects	5t	NU	HTH		LIT							LIT										
4	2	663	well-temple				/Data/Global/Objects	5q	NU	HTH		LIT																	
4	2	664	torch1-temple				/Data/Global/Objects	5r	NU	HTH		LIT							LIT										
4	2	665	torch1-temple				/Data/Global/Objects	5s	NU	HTH		LIT							LIT										
4	2	666	object1-temple				/Data/Global/Objects	5u	OP	HTH		LIT																	
4	2	667	object2-temple				/Data/Global/Objects	5v	OP	HTH		LIT																	
4	2	668	mrbox-baals				/Data/Global/Objects	5w	OP	HTH		LIT																	
4	2	669	well-icecave				/Data/Global/Objects	5x	NU	HTH		LIT																	
4	2	670	magic shrine-temple				/Data/Global/Objects	5y	NU	HTH		LIT							LIT										
4	2	671	healthshrine-temple				/Data/Global/Objects	5z	OP	HTH		LIT																	
4	2	672	manashrine-temple				/Data/Global/Objects	3a	OP	HTH		LIT																	
4	2	673	red light- (touch me)  for blacksmith				/Data/Global/Objects	ss																					
4	2	674	tomb1L-baal's lair				/Data/Global/Objects	3b	OP	HTH		LIT																	
4	2	675	tomb2L-baal's lair				/Data/Global/Objects	3c	OP	HTH		LIT																	
4	2	676	tomb3L-baal's lair				/Data/Global/Objects	3d	OP	HTH		LIT																	
4	2	677	ubub-Ice cave bubbles 01				/Data/Global/Objects	2u	NU	HTH		LIT																	
4	2	678	sbub-Ice cave bubbles 01				/Data/Global/Objects	2s	NU	HTH		LIT																	
4	2	679	tomb1-redbaal's lair				/Data/Global/Objects	3f	OP	HTH		LIT																	
4	2	680	tomb1L-redbaal's lair				/Data/Global/Objects	3g	OP	HTH		LIT																	
4	2	681	tomb2-redbaal's lair				/Data/Global/Objects	3h	OP	HTH		LIT																	
4	2	682	tomb2L-redbaal's lair				/Data/Global/Objects	3i	OP	HTH		LIT																	
4	2	683	tomb3-redbaal's lair				/Data/Global/Objects	3j	OP	HTH		LIT																	
4	2	684	tomb3L-redbaal's lair				/Data/Global/Objects	3k	OP	HTH		LIT																	
4	2	685	mrbox-redbaals				/Data/Global/Objects	3L	OP	HTH		LIT																	
4	2	686	torch1-redbaal's lair				/Data/Global/Objects	3m	NU	HTH		LIT							LIT										
4	2	687	torch2-redbaal's lair				/Data/Global/Objects	3n	NU	HTH		LIT							LIT										
4	2	688	candles-temple				/Data/Global/Objects	3o	NU	HTH		LIT							LIT										
4	2	689	Waypoint-temple				/Data/Global/Objects	3p	ON	HTH		LIT							LIT										
4	2	690	deadperson-everywhere				/Data/Global/Objects	3q	NU	HTH		LIT																	
4	2	691	groundtomb-temple				/Data/Global/Objects	3s	OP	HTH		LIT																	
4	2	692	Dummy-Larzuk Greeting				/Data/Global/Objects	ss																					
4	2	693	Dummy-Larzuk Standard				/Data/Global/Objects	ss																					
4	2	694	groundtombL-temple				/Data/Global/Objects	3t	OP	HTH		LIT																	
4	2	695	deadperson2-everywhere				/Data/Global/Objects	3u	OP	HTH		LIT																	
4	2	696	ancientsaltar-ancientsaltar				/Data/Global/Objects	4a	OP	HTH		LIT							LIT										
4	2	697	To The Worldstone Keep Level 1-ancientsdoor				/Data/Global/Objects	4b	OP	HTH		LIT																	
4	2	698	eweaponrackR-everywhere				/Data/Global/Objects	3x	NU	HTH		LIT																	
4	2	699	eweaponrackL-everywhere				/Data/Global/Objects	3y	NU	HTH		LIT																	
4	2	700	earmorstandR-everywhere				/Data/Global/Objects	3z	NU	HTH		LIT																	
4	2	701	earmorstandL-everywhere				/Data/Global/Objects	4c	NU	HTH		LIT																	
4	2	702	torch2-summit				/Data/Global/Objects	9g	NU	HTH		LIT							LIT										
4	2	703	funeralpire-outside				/Data/Global/Objects	9h	NU	HTH		LIT							LIT										
4	2	704	burninglogs-outside				/Data/Global/Objects	9i	NU	HTH		LIT							LIT										
4	2	705	stma-Ice cave steam				/Data/Global/Objects	2o	NU	HTH		LIT																	
4	2	706	deadperson2-everywhere				/Data/Global/Objects	3v	OP	HTH		LIT																	
4	2	707	Dummy-Baal's lair				/Data/Global/Objects	ss																					
4	2	708	fana-frozen anya				/Data/Global/Objects	2n	NU	HTH		LIT																	
4	2	709	BBQB-BBQ Bunny				/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									
4	2	710	btor-Baal Torch Big				/Data/Global/Objects	25	NU	HTH		LIT							LIT										
4	2	711	Dummy-invisible ancient				/Data/Global/Objects	ss																					
4	2	712	Dummy-invisible base				/Data/Global/Objects	ss																					
4	2	713	The Worldstone Chamber-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
4	2	714	Glacial Caves Level 1-summit door				/Data/Global/Objects	4u	OP	HTH		LIT																	
4	2	715	strlastcinematic-last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
4	2	716	Harrogath-last last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
4	2	717	Zoo-test data				/Data/Global/Objects	ss																					
4	2	718	Keeper-test data				/Data/Global/Objects	7z	NU	HTH		LIT																	
4	2	719	Throne of Destruction-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
4	2	720	Dummy-fire place guy				/Data/Global/Objects	7y	NU	HTH		LIT																	
4	2	721	Dummy-door blocker				/Data/Global/Objects	ss																					
4	2	722	Dummy-door blocker				/Data/Global/Objects	ss																					
5	1	0	larzuk-ACT 5 TABLE			7	/Data/Global/Monsters	XR	NU	HTH		LIT																	0
5	1	1	drehya-ACT 5 TABLE				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
5	1	2	malah-ACT 5 TABLE				/Data/Global/Monsters	XT	NU	HTH		LIT																	0
5	1	3	nihlathak-ACT 5 TABLE				/Data/Global/Monsters	0J	NU	HTH		LIT																	0
5	1	4	qual-kehk-ACT 5 TABLE				/Data/Global/Monsters	XV	NU	HTH		LIT																	0
5	1	5	place_impgroup-ACT 5 TABLE				/Data/Global/Monsters	IP	NU	HTH		LIT																	0
5	1	6	Siege Boss-ACT 5 TABLE			1	/Data/Global/Monsters	OS	NU	HTH	HVY	HVY		HVY	HVY		HVY	HVY	MED	MED									0
5	1	7	tyrael3-ACT 5 TABLE				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
5	1	8	cain6-ACT 5 TABLE				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	9	place_imp-ACT 5 TABLE																										0
5	1	10	place_minion-ACT 5 TABLE																										0
5	1	11	place_miniongroup-ACT 5 TABLE																										0
5	1	12	catapult2-ACT 5 TABLE				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT									0
5	1	13	catapult1-ACT 5 TABLE				/Data/Global/Monsters	65	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT									0
5	1	14	place_bloodlord-ACT 5 TABLE																										0
5	1	15	catapultspotter2-ACT 5 TABLE																	LIT									0
5	1	16	catapultspotter1-ACT 5 TABLE																										0
5	1	17	act5barb1-ACT 5 TABLE			3	/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
5	1	18	place_deadbarb-ACT 5 TABLE																										0
5	1	19	place_deadminion-ACT 5 TABLE																										0
5	1	20	place_deadimp-ACT 5 TABLE																										0
5	1	21	cain6-ACT 5 TABLE				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	22	act5barb3-ACT 5 TABLE				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
5	1	23	place_reanimateddead-ACT 5 TABLE																										0
5	1	24	ancientstatue1-ACT 5 TABLE				/Data/Global/Monsters	0H	NU	HTH		LIT																	0
5	1	25	ancientstatue2-ACT 5 TABLE				/Data/Global/Monsters	0G	NU	HTH		LIT																	0
5	1	26	ancientstatue3-ACT 5 TABLE				/Data/Global/Monsters	0I	NU	HTH		LIT																	0
5	1	27	Dac Farren-ACT 5 TABLE																										0
5	1	28	baalthrone-ACT 5 TABLE				/Data/Global/Monsters	41	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
5	1	29	baaltaunt-ACT 5 TABLE																										0
5	1	30	injuredbarb1-ACT 5 TABLE				/Data/Global/Monsters	6Z	NU	HTH		LIT																	0
5	1	31	injuredbarb2-ACT 5 TABLE				/Data/Global/Monsters	7J	NU	HTH		LIT																	0
5	1	32	injuredbarb3-ACT 5 TABLE				/Data/Global/Monsters	7I	NU	HTH		LIT																	0
5	1	33	baalcrab-ACT 5 TABLE			7	/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT								/Data/Global/Monsters/42/COF/Palshift.dat	5	0
5	1	34	Axe Dweller-ACT 5 TABLE																										0
5	1	35	Bonesaw Breaker-ACT 5 TABLE																										0
5	1	36	Megaflow Rectifier-ACT 5 TABLE																										0
5	1	37	Eyeback Unleashed-ACT 5 TABLE																										0
5	1	38	Threash Socket-ACT 5 TABLE																										0
5	1	39	Pindleskin-ACT 5 TABLE																										0
5	1	40	Snapchip Shatter-ACT 5 TABLE																										0
5	1	41	Anodized Elite-ACT 5 TABLE																										0
5	1	42	Vinvear Molech-ACT 5 TABLE																										0
5	1	43	Sharp Tooth Sayer-ACT 5 TABLE																										0
5	1	44	Magma Torquer-ACT 5 TABLE																										0
5	1	45	Blaze Ripper-ACT 5 TABLE																										0
5	1	46	Frozenstein-ACT 5 TABLE																										0
5	1	47	worldstoneeffect-ACT 5 TABLE																										0
5	1	48	chicken-ACT 5 TABLE			2	/Data/Global/Monsters	CK	NU	HTH		LIT																	0
5	1	49	place_champion-ACT 5 TABLE																										0
5	1	50	evilhut-ACT 5 TABLE				/Data/Global/Monsters	2T	S1	HTH		LIT							LIT	LIT									0
5	1	51	place_nothing-ACT 5 TABLE																										0
5	1	52	place_nothing-ACT 5 TABLE																										0
5	1	53	place_nothing-ACT 5 TABLE																										0
5	1	54	place_nothing-ACT 5 TABLE																										0
5	1	55	place_nothing-ACT 5 TABLE																										0
5	1	56	skeleton1-Skeleton-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	57	skeleton2-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	58	skeleton3-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	59	skeleton4-BurningDead-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	60	skeleton5-Horror-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	61	zombie1-Zombie-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	62	zombie2-HungryDead-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	63	zombie3-Ghoul-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	64	zombie4-DrownedCarcass-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	65	zombie5-PlagueBearer-Zombie				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	66	bighead1-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	67	bighead2-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	68	bighead3-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	69	bighead4-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	70	bighead5-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	71	foulcrow1-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	72	foulcrow2-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	73	foulcrow3-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	74	foulcrow4-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	75	fallen1-Fallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
5	1	76	fallen2-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
5	1	77	fallen3-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
5	1	78	fallen4-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
5	1	79	fallen5-WarpedFallen-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				AXE		TCH	LIT										0
5	1	80	brute2-Brute-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
5	1	81	brute3-Yeti-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
5	1	82	brute4-Crusher-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
5	1	83	brute5-WailingBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
5	1	84	brute1-GargantuanBeast-Brute				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
5	1	85	sandraider1-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	86	sandraider2-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	87	sandraider3-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	88	sandraider4-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	89	sandraider5-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	90	gorgon1-unused-Idle				/Data/Global/Monsters	GO																					0
5	1	91	gorgon2-unused-Idle				/Data/Global/Monsters	GO																					0
5	1	92	gorgon3-unused-Idle				/Data/Global/Monsters	GO																					0
5	1	93	gorgon4-unused-Idle				/Data/Global/Monsters	GO																					0
5	1	94	wraith1-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	95	wraith2-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	96	wraith3-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	97	wraith4-Apparition-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	98	wraith5-DarkShape-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	99	corruptrogue1-DarkHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	100	corruptrogue2-VileHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	101	corruptrogue3-DarkStalker-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	102	corruptrogue4-BlackRogue-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	103	corruptrogue5-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	104	baboon1-DuneBeast-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	105	baboon2-RockDweller-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	106	baboon3-JungleHunter-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	107	baboon4-DoomApe-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	108	baboon5-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	109	goatman1-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	110	goatman2-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	111	goatman3-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	112	goatman4-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	113	goatman5-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	114	fallenshaman1-FallenShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	115	fallenshaman2-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	116	fallenshaman3-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	117	fallenshaman4-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	118	fallenshaman5-WarpedShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	119	quillrat1-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	120	quillrat2-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	121	quillrat3-ThornBeast-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	122	quillrat4-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	123	quillrat5-JungleUrchin-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	124	sandmaggot1-SandMaggot-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	125	sandmaggot2-RockWorm-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	126	sandmaggot3-Devourer-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	127	sandmaggot4-GiantLamprey-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	128	sandmaggot5-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	129	clawviper1-TombViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	130	clawviper2-ClawViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	131	clawviper3-Salamander-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	132	clawviper4-PitViper-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	133	clawviper5-SerpentMagus-ClawViper				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	134	sandleaper1-SandLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	135	sandleaper2-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	136	sandleaper3-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	137	sandleaper4-TreeLurker-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	138	sandleaper5-RazorPitDemon-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	139	pantherwoman1-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	140	pantherwoman2-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	141	pantherwoman3-NightTiger-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	142	pantherwoman4-HellCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	143	swarm1-Itchies-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
5	1	144	swarm2-BlackLocusts-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
5	1	145	swarm3-PlagueBugs-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
5	1	146	swarm4-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
5	1	147	scarab1-DungSoldier-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	148	scarab2-SandWarrior-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	149	scarab3-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	150	scarab4-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	151	scarab5-AlbinoRoach-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	152	mummy1-DriedCorpse-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	153	mummy2-Decayed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	154	mummy3-Embalmed-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	155	mummy4-PreservedDead-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	156	mummy5-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	157	unraveler1-HollowOne-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	158	unraveler2-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	159	unraveler3-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	160	unraveler4-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	161	unraveler5-Baal Subject Mummy-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	162	chaoshorde1-unused-Idle				/Data/Global/Monsters	CH																					0
5	1	163	chaoshorde2-unused-Idle				/Data/Global/Monsters	CH																					0
5	1	164	chaoshorde3-unused-Idle				/Data/Global/Monsters	CH																					0
5	1	165	chaoshorde4-unused-Idle				/Data/Global/Monsters	CH																					0
5	1	166	vulture1-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
5	1	167	vulture2-UndeadScavenger-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
5	1	168	vulture3-HellBuzzard-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
5	1	169	vulture4-WingedNightmare-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
5	1	170	mosquito1-Sucker-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
5	1	171	mosquito2-Feeder-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
5	1	172	mosquito3-BloodHook-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
5	1	173	mosquito4-BloodWing-Mosquito				/Data/Global/Monsters	MO	NU	HTH		LIT							LIT										0
5	1	174	willowisp1-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	175	willowisp2-SwampGhost-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	176	willowisp3-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	177	willowisp4-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	178	arach1-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	179	arach2-SandFisher-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	180	arach3-PoisonSpinner-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	181	arach4-FlameSpider-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	182	arach5-SpiderMagus-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	183	thornhulk1-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
5	1	184	thornhulk2-BrambleHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
5	1	185	thornhulk3-Thrasher-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
5	1	186	thornhulk4-Spikefist-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
5	1	187	vampire1-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	188	vampire2-NightLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	189	vampire3-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	190	vampire4-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	191	vampire5-Banished-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	192	batdemon1-DesertWing-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	193	batdemon2-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	194	batdemon3-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	195	batdemon4-BloodDiver-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	196	batdemon5-DarkFamiliar-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	197	fetish1-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	198	fetish2-Fetish-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	199	fetish3-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	200	fetish4-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	201	fetish5-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	202	cain1-DeckardCain-NpcOutOfTown				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	203	gheed-Gheed-Npc				/Data/Global/Monsters	GH	NU	HTH		LIT																	0
5	1	204	akara-Akara-Npc				/Data/Global/Monsters	PS	NU	HTH		LIT																	0
5	1	205	chicken-dummy-Idle				/Data/Global/Monsters	CK	NU	HTH		LIT																	0
5	1	206	kashya-Kashya-Npc				/Data/Global/Monsters	RC	NU	HTH		LIT																	0
5	1	207	rat-dummy-Idle				/Data/Global/Monsters	RT	NU	HTH		LIT																	0
5	1	208	rogue1-Dummy-Idle				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
5	1	209	hellmeteor-Dummy-HellMeteor				/Data/Global/Monsters	K9																					0
5	1	210	charsi-Charsi-Npc				/Data/Global/Monsters	CI	NU	HTH		LIT																	0
5	1	211	warriv1-Warriv-Npc				/Data/Global/Monsters	WA	NU	HTH		LIT																	0
5	1	212	andariel-Andariel-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
5	1	213	bird1-dummy-Idle				/Data/Global/Monsters	BS	WL	HTH		LIT																	0
5	1	214	bird2-dummy-Idle				/Data/Global/Monsters	BL																					0
5	1	215	bat-dummy-Idle				/Data/Global/Monsters	B9	WL	HTH		LIT																	0
5	1	216	cr_archer1-DarkRanger-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	217	cr_archer2-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	218	cr_archer3-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	219	cr_archer4-BlackArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	220	cr_archer5-FleshArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	221	cr_lancer1-DarkSpearwoman-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	222	cr_lancer2-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	223	cr_lancer3-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	224	cr_lancer4-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	225	cr_lancer5-FleshLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	226	sk_archer1-SkeletonArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	227	sk_archer2-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	228	sk_archer3-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	229	sk_archer4-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	230	sk_archer5-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	231	warriv2-Warriv-Npc				/Data/Global/Monsters	WX	NU	HTH		LIT																	0
5	1	232	atma-Atma-Npc				/Data/Global/Monsters	AS	NU	HTH		LIT																	0
5	1	233	drognan-Drognan-Npc				/Data/Global/Monsters	DR	NU	HTH		LIT																	0
5	1	234	fara-Fara-Npc				/Data/Global/Monsters	OF	NU	HTH		LIT																	0
5	1	235	cow-dummy-Idle				/Data/Global/Monsters	CW	NU	HTH		LIT																	0
5	1	236	maggotbaby1-SandMaggotYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
5	1	237	maggotbaby2-RockWormYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
5	1	238	maggotbaby3-DevourerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
5	1	239	maggotbaby4-GiantLampreyYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
5	1	240	maggotbaby5-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
5	1	241	camel-dummy-Idle				/Data/Global/Monsters	CM	NU	HTH		LIT																	0
5	1	242	blunderbore1-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
5	1	243	blunderbore2-Gorbelly-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
5	1	244	blunderbore3-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
5	1	245	blunderbore4-Urdar-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
5	1	246	maggotegg1-SandMaggotEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
5	1	247	maggotegg2-RockWormEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
5	1	248	maggotegg3-DevourerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
5	1	249	maggotegg4-GiantLampreyEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
5	1	250	maggotegg5-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
5	1	251	act2male-dummy-Towner				/Data/Global/Monsters	2M	NU	HTH	OLD	MED	MED						TUR										0
5	1	252	act2female-Dummy-Towner				/Data/Global/Monsters	2F	NU	HTH	LIT	LIT	LIT																0
5	1	253	act2child-dummy-Towner				/Data/Global/Monsters	2C																					0
5	1	254	greiz-Greiz-Npc				/Data/Global/Monsters	GR	NU	HTH		LIT																	0
5	1	255	elzix-Elzix-Npc				/Data/Global/Monsters	EL	NU	HTH		LIT																	0
5	1	256	geglash-Geglash-Npc				/Data/Global/Monsters	GE	NU	HTH		LIT																	0
5	1	257	jerhyn-Jerhyn-Npc				/Data/Global/Monsters	JE	NU	HTH		LIT																	0
5	1	258	lysander-Lysander-Npc				/Data/Global/Monsters	LY	NU	HTH		LIT																	0
5	1	259	act2guard1-Dummy-Towner				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
5	1	260	act2vendor1-dummy-Vendor				/Data/Global/Monsters	M1	NU	HTH		LIT																	0
5	1	261	act2vendor2-dummy-Vendor				/Data/Global/Monsters	M2	NU	HTH		LIT																	0
5	1	262	crownest1-FoulCrowNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
5	1	263	crownest2-BloodHawkNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
5	1	264	crownest3-BlackVultureNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
5	1	265	crownest4-CloudStalkerNest-FoulCrowNest				/Data/Global/Monsters	BN	NU	HTH		LIT																	0
5	1	266	meshif1-Meshif-Npc				/Data/Global/Monsters	MS	NU	HTH		LIT																	0
5	1	267	duriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
5	1	268	bonefetish1-Undead RatMan-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	269	bonefetish2-Undead Fetish-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	270	bonefetish3-Undead Flayer-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	271	bonefetish4-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	272	bonefetish5-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	273	darkguard1-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	274	darkguard2-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	275	darkguard3-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	276	darkguard4-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	277	darkguard5-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	278	bloodmage1-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	279	bloodmage2-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	280	bloodmage3-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	281	bloodmage4-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	282	bloodmage5-unused-Idle				/Data/Global/Monsters	xx																					0
5	1	283	maggot-Maggot-Idle				/Data/Global/Monsters	MA	NU	HTH		LIT																	0
5	1	284	sarcophagus-MummyGenerator-Sarcophagus				/Data/Global/Monsters	MG	NU	HTH		LIT																	0
5	1	285	radament-Radament-GreaterMummy				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
5	1	286	firebeast-unused-ElementalBeast				/Data/Global/Monsters	FM	NU	HTH		LIT																	0
5	1	287	iceglobe-unused-ElementalBeast				/Data/Global/Monsters	IM	NU	HTH		LIT																	0
5	1	288	lightningbeast-unused-ElementalBeast				/Data/Global/Monsters	xx																					0
5	1	289	poisonorb-unused-ElementalBeast				/Data/Global/Monsters	PM	NU	HTH		LIT																	0
5	1	290	flyingscimitar-FlyingScimitar-FlyingScimitar				/Data/Global/Monsters	ST	NU	HTH		LIT																	0
5	1	291	zealot1-Zakarumite-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
5	1	292	zealot2-Faithful-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
5	1	293	zealot3-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
5	1	294	cantor1-Sexton-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
5	1	295	cantor2-Cantor-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
5	1	296	cantor3-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
5	1	297	cantor4-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
5	1	298	mephisto-Mephisto-Mephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
5	1	299	diablo-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
5	1	300	cain2-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	301	cain3-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	302	cain4-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	303	frogdemon1-Swamp Dweller-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
5	1	304	frogdemon2-Bog Creature-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
5	1	305	frogdemon3-Slime Prince-FrogDemon				/Data/Global/Monsters	FD	NU	HTH		LIT																	0
5	1	306	summoner-Summoner-Summoner				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
5	1	307	tyrael1-tyrael-NpcStationary				/Data/Global/Monsters	TX	NU	HTH		LIT		LIT	LIT														0
5	1	308	asheara-asheara-Npc				/Data/Global/Monsters	AH	NU	HTH		LIT																	0
5	1	309	hratli-hratli-Npc				/Data/Global/Monsters	HR	NU	HTH		LIT																	0
5	1	310	alkor-alkor-Npc				/Data/Global/Monsters	AL	NU	HTH		LIT																	0
5	1	311	ormus-ormus-Npc				/Data/Global/Monsters	OR	NU	HTH		LIT																	0
5	1	312	izual-izual-Izual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
5	1	313	halbu-halbu-Npc				/Data/Global/Monsters	20	NU	HTH		LIT																	0
5	1	314	tentacle1-WaterWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
5	1	315	tentacle2-RiverStalkerLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
5	1	316	tentacle3-StygianWatcherLimb-Tentacle				/Data/Global/Monsters	TN	NU	HTH		LIT							LIT										0
5	1	317	tentaclehead1-WaterWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
5	1	318	tentaclehead2-RiverStalkerHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
5	1	319	tentaclehead3-StygianWatcherHead-TentacleHead				/Data/Global/Monsters	TE	NU	HTH		LIT							LIT										0
5	1	320	meshif2-meshif-Npc				/Data/Global/Monsters	M3	NU	HTH		LIT																	0
5	1	321	cain5-DeckardCain-Npc				/Data/Global/Monsters	1D	NU	HTH		LIT																	0
5	1	322	navi-navi-Navi				/Data/Global/Monsters	RG	NU	HTH	LIT	LIT		LIT	LIT		LBW		LIT	LIT									0
5	1	323	bloodraven-Bloodraven-BloodRaven				/Data/Global/Monsters	CR	NU	BOW	BRV	HVY	BRV	HVY	HVY	LIT	LBB		HVY	HVY									0
5	1	324	bug-Dummy-Idle				/Data/Global/Monsters	BG	NU	HTH		LIT																	0
5	1	325	scorpion-Dummy-Idle				/Data/Global/Monsters	DS	NU	HTH		LIT																	0
5	1	326	rogue2-RogueScout-GoodNpcRanged				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
5	1	327	roguehire-Dummy-Hireable				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
5	1	328	rogue3-Dummy-TownRogue				/Data/Global/Monsters	RG	NU	HTH	MED	MED		LIT	LIT		LBW		MED	MED									0
5	1	329	gargoyletrap-GargoyleTrap-GargoyleTrap				/Data/Global/Monsters	GT	NU	HTH		LIT																	0
5	1	330	skmage_pois1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	331	skmage_pois2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	332	skmage_pois3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	333	skmage_pois4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	334	fetishshaman1-RatManShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	335	fetishshaman2-FetishShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	336	fetishshaman3-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	337	fetishshaman4-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	338	fetishshaman5-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	339	larva-larva-Idle				/Data/Global/Monsters	LV	NU	HTH		LIT																	0
5	1	340	maggotqueen1-SandMaggotQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
5	1	341	maggotqueen2-RockWormQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
5	1	342	maggotqueen3-DevourerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
5	1	343	maggotqueen4-GiantLampreyQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
5	1	344	maggotqueen5-WorldKillerQueen-SandMaggotQueen				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
5	1	345	claygolem-ClayGolem-NecroPet				/Data/Global/Monsters	G1	NU	HTH		LIT																	0
5	1	346	bloodgolem-BloodGolem-NecroPet				/Data/Global/Monsters	G2	NU	HTH		LIT																	0
5	1	347	irongolem-IronGolem-NecroPet				/Data/Global/Monsters	G4	NU	HTH		LIT																	0
5	1	348	firegolem-FireGolem-NecroPet				/Data/Global/Monsters	G3	NU	HTH		LIT																	0
5	1	349	familiar-Dummy-Idle				/Data/Global/Monsters	FI	NU	HTH		LIT																	0
5	1	350	act3male-Dummy-Towner				/Data/Global/Monsters	N4	NU	HTH	BRD	HVY	HVY	HEV	HEV	FSH	SAK		TKT										0
5	1	351	baboon6-NightMarauder-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	352	act3female-Dummy-Towner				/Data/Global/Monsters	N3	NU	HTH	LIT	MTP	SRT			BSK	BSK												0
5	1	353	natalya-Natalya-Npc				/Data/Global/Monsters	TZ	NU	HTH		LIT																	0
5	1	354	vilemother1-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
5	1	355	vilemother2-StygianHag-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
5	1	356	vilemother3-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
5	1	357	vilechild1-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
5	1	358	vilechild2-StygianDog-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
5	1	359	vilechild3-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
5	1	360	fingermage1-Groper-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
5	1	361	fingermage2-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
5	1	362	fingermage3-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
5	1	363	regurgitator1-Corpulent-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
5	1	364	regurgitator2-CorpseSpitter-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
5	1	365	regurgitator3-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
5	1	366	doomknight1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	367	doomknight2-AbyssKnight-AbyssKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	368	doomknight3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	369	quillbear1-QuillBear-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
5	1	370	quillbear2-SpikeGiant-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
5	1	371	quillbear3-ThornBrute-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
5	1	372	quillbear4-RazorBeast-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
5	1	373	quillbear5-GiantUrchin-QuillMother				/Data/Global/Monsters	S7	NU	HTH		LIT																	0
5	1	374	snake-Dummy-Idle				/Data/Global/Monsters	CO	NU	HTH		LIT																	0
5	1	375	parrot-Dummy-Idle				/Data/Global/Monsters	PR	WL	HTH		LIT																	0
5	1	376	fish-Dummy-Idle				/Data/Global/Monsters	FJ																					0
5	1	377	evilhole1-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
5	1	378	evilhole2-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
5	1	379	evilhole3-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
5	1	380	evilhole4-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
5	1	381	evilhole5-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
5	1	382	trap-firebolt-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
5	1	383	trap-horzmissile-a trap-Trap-RightArrow				/Data/Global/Monsters	9A																					0
5	1	384	trap-vertmissile-a trap-Trap-LeftArrow				/Data/Global/Monsters	9A																					0
5	1	385	trap-poisoncloud-a trap-Trap-Poison				/Data/Global/Monsters	9A																					0
5	1	386	trap-lightning-a trap-Trap-Missile				/Data/Global/Monsters	9A																					0
5	1	387	act2guard2-Kaelan-JarJar				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
5	1	388	invisospawner-Dummy-InvisoSpawner				/Data/Global/Monsters	K9																					0
5	1	389	diabloclone-Diablo-Diablo				/Data/Global/Monsters	DI	NU	HTH		LIT	LIT	LIT	LIT														0
5	1	390	suckernest1-SuckerNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
5	1	391	suckernest2-FeederNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
5	1	392	suckernest3-BloodHookNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
5	1	393	suckernest4-BloodWingNest-MosquitoNest				/Data/Global/Monsters	DH	NU	HTH		LIT																	0
5	1	394	act2hire-Guard-Hireable				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	GLV			LIT	LIT	LIT								0
5	1	395	minispider-Dummy-Idle				/Data/Global/Monsters	LS	NU	HTH		LIT																	0
5	1	396	boneprison1--Idle				/Data/Global/Monsters	67	NU	HTH		LIT																	0
5	1	397	boneprison2--Idle				/Data/Global/Monsters	66	NU	HTH		LIT																	0
5	1	398	boneprison3--Idle				/Data/Global/Monsters	69	NU	HTH		LIT																	0
5	1	399	boneprison4--Idle				/Data/Global/Monsters	68	NU	HTH		LIT																	0
5	1	400	bonewall-Dummy-BoneWall				/Data/Global/Monsters	BW	NU	HTH		LIT																	0
5	1	401	councilmember1-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	402	councilmember2-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	403	councilmember3-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	404	turret1-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
5	1	405	turret2-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
5	1	406	turret3-Turret-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
5	1	407	hydra1-Hydra-Hydra				/Data/Global/Monsters	HX	NU	HTH		LIT							LIT										0
5	1	408	hydra2-Hydra-Hydra				/Data/Global/Monsters	21	NU	HTH		LIT							LIT										0
5	1	409	hydra3-Hydra-Hydra				/Data/Global/Monsters	HZ	NU	HTH		LIT							LIT										0
5	1	410	trap-melee-a trap-Trap-Melee				/Data/Global/Monsters	M4	A1	HTH		LIT																	0
5	1	411	seventombs-Dummy-7TIllusion				/Data/Global/Monsters	9A																					0
5	1	412	dopplezon-Dopplezon-Idle				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
5	1	413	valkyrie-Valkyrie-NecroPet				/Data/Global/Monsters	VK	DT	HTH		LIT							LIT										0
5	1	414	act2guard3-Dummy-Idle				/Data/Global/Monsters	SK																					0
5	1	415	act3hire-Iron Wolf-Hireable				/Data/Global/Monsters	IW	NU	1HS	LIT	LIT				WND		KIT											0
5	1	416	megademon1-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	417	megademon2-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	418	megademon3-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	419	necroskeleton-NecroSkeleton-NecroPet				/Data/Global/Monsters	SK	NU	1HS	DES	HVY	DES	DES	DES	SCM		KIT	DES	DES	LIT								0
5	1	420	necromage-NecroMage-NecroPet				/Data/Global/Monsters	SK	NU	HTH	DES	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
5	1	421	griswold-Griswold-Griswold				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
5	1	422	compellingorb-compellingorb-Idle				/Data/Global/Monsters	9a																					0
5	1	423	tyrael2-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
5	1	424	darkwanderer-youngdiablo-DarkWanderer				/Data/Global/Monsters	1Z	NU	HTH		LIT																	0
5	1	425	trap-nova-a trap-Trap-Nova				/Data/Global/Monsters	9A																					0
5	1	426	spiritmummy-Dummy-Idle				/Data/Global/Monsters	xx																					0
5	1	427	lightningspire-LightningSpire-ArcaneTower				/Data/Global/Monsters	AE	NU	HTH		LIT							LIT										0
5	1	428	firetower-FireTower-DesertTurret				/Data/Global/Monsters	PB	NU	HTH		LIT																	0
5	1	429	slinger1-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
5	1	430	slinger2-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
5	1	431	slinger3-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
5	1	432	slinger4-HellSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
5	1	433	act2guard4-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
5	1	434	act2guard5-Dummy-Idle				/Data/Global/Monsters	GU	NU	HTH	LIT	LIT	LIT	LIT	LIT	SPR			LIT	LIT	LIT								0
5	1	435	skmage_cold1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
5	1	436	skmage_cold2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
5	1	437	skmage_cold3-BaalColdMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
5	1	438	skmage_cold4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
5	1	439	skmage_fire1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
5	1	440	skmage_fire2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
5	1	441	skmage_fire3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
5	1	442	skmage_fire4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
5	1	443	skmage_ltng1-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
5	1	444	skmage_ltng2-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
5	1	445	skmage_ltng3-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
5	1	446	skmage_ltng4-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
5	1	447	hellbovine-Hell Bovine-Skeleton				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
5	1	448	window1--Idle				/Data/Global/Monsters	VH	NU	HTH		LIT							LIT										0
5	1	449	window2--Idle				/Data/Global/Monsters	VJ	NU	HTH		LIT							LIT										0
5	1	450	slinger5-SpearCat-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
5	1	451	slinger6-NightSlinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	PHA	HVY		HVY	HVY		JAV	BUC	HVY	HVY	HVY	HVY							0
5	1	452	fetishblow1-RatMan-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	453	fetishblow2-Fetish-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	454	fetishblow3-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	455	fetishblow4-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	456	fetishblow5-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	457	mephistospirit-Dummy-Spirit				/Data/Global/Monsters	M6	A1	HTH		LIT																	0
5	1	458	smith-The Smith-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
5	1	459	trappedsoul1-TrappedSoul-TrappedSoul				/Data/Global/Monsters	10	NU	HTH		LIT																	0
5	1	460	trappedsoul2-TrappedSoul-TrappedSoul				/Data/Global/Monsters	13	NU	HTH		LIT																	0
5	1	461	jamella-Jamella-Npc				/Data/Global/Monsters	ja	NU	HTH		LIT																	0
5	1	462	izualghost-Izual-NpcStationary				/Data/Global/Monsters	17	NU	HTH		LIT							LIT										0
5	1	463	fetish11-RatMan-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	464	malachai-Malachai-Buffy				/Data/Global/Monsters	36	NU	HTH		LIT							LIT										0
5	1	465	hephasto-The Feature Creep-Smith				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
5	1	466	wakeofdestruction-Wake of Destruction-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
5	1	467	chargeboltsentry-Charged Bolt Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
5	1	468	lightningsentry-Lightning Sentry-AssassinSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
5	1	469	bladecreeper-Blade Creeper-BladeCreeper				/Data/Global/Monsters	b8	NU	HTH		LIT							LIT										0
5	1	470	invisopet-Invis Pet-InvisoPet				/Data/Global/Monsters	k9																					0
5	1	471	infernosentry-Inferno Sentry-AssassinSentry				/Data/Global/Monsters	e9	NU	HTH		LIT																	0
5	1	472	deathsentry-Death Sentry-DeathSentry				/Data/Global/Monsters	lg	NU	HTH		LIT							LIT										0
5	1	473	shadowwarrior-Shadow Warrior-ShadowWarrior				/Data/Global/Monsters	k9																					0
5	1	474	shadowmaster-Shadow Master-ShadowMaster				/Data/Global/Monsters	k9																					0
5	1	475	druidhawk-Druid Hawk-Raven				/Data/Global/Monsters	hk	NU	HTH		LIT																	0
5	1	476	spiritwolf-Druid Spirit Wolf-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
5	1	477	fenris-Druid Fenris-DruidWolf				/Data/Global/Monsters	wf	NU	HTH		LIT																	0
5	1	478	spiritofbarbs-Spirit of Barbs-Totem				/Data/Global/Monsters	x4	NU	HTH		LIT																	0
5	1	479	heartofwolverine-Heart of Wolverine-Totem				/Data/Global/Monsters	x3	NU	HTH		LIT																	0
5	1	480	oaksage-Oak Sage-Totem				/Data/Global/Monsters	xw	NU	HTH		LIT																	0
5	1	481	plaguepoppy-Druid Plague Poppy-Vines				/Data/Global/Monsters	k9																					0
5	1	482	cycleoflife-Druid Cycle of Life-CycleOfLife				/Data/Global/Monsters	k9																					0
5	1	483	vinecreature-Vine Creature-CycleOfLife				/Data/Global/Monsters	k9																					0
5	1	484	druidbear-Druid Bear-DruidBear				/Data/Global/Monsters	b7	NU	HTH		LIT																	0
5	1	485	eagle-Eagle-Idle				/Data/Global/Monsters	eg	NU	HTH		LIT							LIT										0
5	1	486	wolf-Wolf-NecroPet				/Data/Global/Monsters	40	NU	HTH		LIT																	0
5	1	487	bear-Bear-NecroPet				/Data/Global/Monsters	TG	NU	HTH		LIT							LIT										0
5	1	488	barricadedoor1-Barricade Door-Idle				/Data/Global/Monsters	AJ	NU	HTH		LIT																	0
5	1	489	barricadedoor2-Barricade Door-Idle				/Data/Global/Monsters	AG	NU	HTH		LIT																	0
5	1	490	prisondoor-Prison Door-Idle				/Data/Global/Monsters	2Q	NU	HTH		LIT																	0
5	1	491	barricadetower-Barricade Tower-SiegeTower				/Data/Global/Monsters	ac	NU	HTH		LIT							LIT						LIT				0
5	1	492	reanimatedhorde1-RotWalker-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	493	reanimatedhorde2-ReanimatedHorde-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	494	reanimatedhorde3-ProwlingDead-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	495	reanimatedhorde4-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	496	reanimatedhorde5-DefiledWarrior-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	497	siegebeast1-Siege Beast-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
5	1	498	siegebeast2-CrushBiest-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
5	1	499	siegebeast3-BloodBringer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
5	1	500	siegebeast4-GoreBearer-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
5	1	501	siegebeast5-DeamonSteed-SiegeBeast				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
5	1	502	snowyeti1-SnowYeti1-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
5	1	503	snowyeti2-SnowYeti2-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
5	1	504	snowyeti3-SnowYeti3-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
5	1	505	snowyeti4-SnowYeti4-Brute				/Data/Global/Monsters	io	NU	HTH		LIT																	0
5	1	506	wolfrider1-WolfRider1-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
5	1	507	wolfrider2-WolfRider2-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
5	1	508	wolfrider3-WolfRider3-Idle				/Data/Global/Monsters	wr	NU	HTH		LIT																	0
5	1	509	minion1-Minionexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	510	minion2-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	511	minion3-IceBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	512	minion4-FireBoar-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	513	minion5-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	514	minion6-IceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	515	minion7-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	516	minion8-GreaterIceSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	517	suicideminion1-FanaticMinion-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	518	suicideminion2-BerserkSlayer-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	519	suicideminion3-ConsumedIceBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	520	suicideminion4-ConsumedFireBoar-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	521	suicideminion5-FrenziedHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	522	suicideminion6-FrenziedIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	523	suicideminion7-InsaneHellSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	524	suicideminion8-InsaneIceSpawn-SuicideMinion				/Data/Global/Monsters	xy	NU	HTH	HVY	LIT																	0
5	1	525	succubus1-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	526	succubus2-VileTemptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	527	succubus3-StygianHarlot-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	528	succubus4-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	529	succubus5-Blood Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	530	succubuswitch1-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	531	succubuswitch2-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	532	succubuswitch3-StygianFury-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	533	succubuswitch4-Blood Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	534	succubuswitch5-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	535	overseer1-OverSeer-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	536	overseer2-Lasher-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	537	overseer3-OverLord-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	538	overseer4-BloodBoss-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	539	overseer5-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	540	minionspawner1-MinionSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	541	minionspawner2-MinionSlayerSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	542	minionspawner3-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	543	minionspawner4-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	544	minionspawner5-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	545	minionspawner6-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	546	minionspawner7-MinionIce/fireBoarSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	547	minionspawner8-Minionice/hellSpawnSpawner-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	548	imp1-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	549	imp2-Imp2-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	550	imp3-Imp3-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	551	imp4-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	552	imp5-Imp5-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	553	catapult1-CatapultS-Catapult				/Data/Global/Monsters	65	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
5	1	554	catapult2-CatapultE-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
5	1	555	catapult3-CatapultSiege-Catapult				/Data/Global/Monsters	64	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT					LIT				0
5	1	556	catapult4-CatapultW-Catapult				/Data/Global/Monsters	ua	NU	HTH	LIT	LIT	LIT	LIT	LIT					LIT	LIT								0
5	1	557	frozenhorror1-Frozen Horror1-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
5	1	558	frozenhorror2-Frozen Horror2-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
5	1	559	frozenhorror3-Frozen Horror3-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
5	1	560	frozenhorror4-Frozen Horror4-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
5	1	561	frozenhorror5-Frozen Horror5-FrozenHorror				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
5	1	562	bloodlord1-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	563	bloodlord2-Blood Lord2-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	564	bloodlord3-Blood Lord3-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	565	bloodlord4-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	566	bloodlord5-Blood Lord5-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	567	larzuk-Larzuk-Npc				/Data/Global/Monsters	XR	NU	HTH		LIT																	0
5	1	568	drehya-Drehya-Npc				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
5	1	569	malah-Malah-Npc				/Data/Global/Monsters	XT	NU	HTH		LIT																	0
5	1	570	nihlathak-Nihlathak Town-Npc				/Data/Global/Monsters	0J	NU	HTH		LIT																	0
5	1	571	qual-kehk-Qual-Kehk-Npc				/Data/Global/Monsters	XV	NU	HTH		LIT																	0
5	1	572	catapultspotter1-Catapult Spotter S-CatapultSpotter				/Data/Global/Monsters	k9																					0
5	1	573	catapultspotter2-Catapult Spotter E-CatapultSpotter				/Data/Global/Monsters	k9																					0
5	1	574	catapultspotter3-Catapult Spotter Siege-CatapultSpotter				/Data/Global/Monsters	k9																					0
5	1	575	catapultspotter4-Catapult Spotter W-CatapultSpotter				/Data/Global/Monsters	k9																					0
5	1	576	cain6-DeckardCain-Npc				/Data/Global/Monsters	DC	NU	HTH		LIT																	0
5	1	577	tyrael3-tyrael-NpcStationary				/Data/Global/Monsters	TY	NU	HTH		LIT		LIT	LIT														0
5	1	578	act5barb1-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
5	1	579	act5barb2-Act 5 Combatant-NpcBarb				/Data/Global/Monsters	0A	NU	1HS	FHM	HVY				AXE	AXE		HVY	HVY									0
5	1	580	barricadewall1-Barricade Wall Right-Idle				/Data/Global/Monsters	A6	NU	HTH		LIT																	0
5	1	581	barricadewall2-Barricade Wall Left-Idle				/Data/Global/Monsters	AK	NU	HTH		LIT																	0
5	1	582	nihlathakboss-Nihlathak-Nihlathak				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
5	1	583	drehyaiced-Drehya-NpcOutOfTown				/Data/Global/Monsters	XS	NU	HTH		LIT																	0
5	1	584	evilhut-Evil hut-GenericSpawner				/Data/Global/Monsters	2T	NU	HTH		LIT							LIT										0
5	1	585	deathmauler1-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	586	deathmauler2-Death Mauler2-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	587	deathmauler3-Death Mauler3-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	588	deathmauler4-Death Mauler4-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	589	deathmauler5-Death Mauler5-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	590	act5pow-POW-Wussie				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
5	1	591	act5barb3-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
5	1	592	act5barb4-Act 5 Townguard-Npc				/Data/Global/Monsters	0A	NU	HTH	HED	LIT				BHN	BHN		LIT	LIT									0
5	1	593	ancientstatue1-Ancient Statue 1-AncientStatue				/Data/Global/Monsters	0G	NU	HTH		LIT																	0
5	1	594	ancientstatue2-Ancient Statue 2-AncientStatue				/Data/Global/Monsters	0H	NU	HTH		LIT																	0
5	1	595	ancientstatue3-Ancient Statue 3-AncientStatue				/Data/Global/Monsters	0I	NU	HTH		LIT																	0
5	1	596	ancientbarb1-Ancient Barbarian 1-Ancient				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
5	1	597	ancientbarb2-Ancient Barbarian 2-Ancient				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
5	1	598	ancientbarb3-Ancient Barbarian 3-Ancient				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
5	1	599	baalthrone-Baal Throne-BaalThrone				/Data/Global/Monsters	41	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
5	1	600	baalcrab-Baal Crab-BaalCrab				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
5	1	601	baaltaunt-Baal Taunt-BaalTaunt				/Data/Global/Monsters	K9																					0
5	1	602	putriddefiler1-Putrid Defiler1-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
5	1	603	putriddefiler2-Putrid Defiler2-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
5	1	604	putriddefiler3-Putrid Defiler3-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
5	1	605	putriddefiler4-Putrid Defiler4-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
5	1	606	putriddefiler5-Putrid Defiler5-PutridDefiler				/Data/Global/Monsters	45	NU	HTH		LIT																	0
5	1	607	painworm1-Pain Worm1-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
5	1	608	painworm2-Pain Worm2-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
5	1	609	painworm3-Pain Worm3-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
5	1	610	painworm4-Pain Worm4-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
5	1	611	painworm5-Pain Worm5-VileDog				/Data/Global/Monsters	46	NU	HTH		LIT																	0
5	1	612	bunny-dummy-Idle				/Data/Global/Monsters	48	NU	HTH		LIT																	0
5	1	613	baalhighpriest-Council Member-HighPriest				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	614	venomlord-VenomLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				FLB													0
5	1	615	baalcrabstairs-Baal Crab to Stairs-BaalToStairs				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
5	1	616	act5hire1-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
5	1	617	act5hire2-dummy-Hireable				/Data/Global/Monsters	0A	NU	1HS	FHM	LIT				AXE	AXE		MED	MED									0
5	1	618	baaltentacle1-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
5	1	619	baaltentacle2-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
5	1	620	baaltentacle3-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
5	1	621	baaltentacle4-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
5	1	622	baaltentacle5-Baal Tentacle-BaalTentacle				/Data/Global/Monsters	44	NU	HTH		LIT							LIT										0
5	1	623	injuredbarb1-dummy-Idle				/Data/Global/Monsters	6z	NU	HTH		LIT																	0
5	1	624	injuredbarb2-dummy-Idle				/Data/Global/Monsters	7j	NU	HTH		LIT																	0
5	1	625	injuredbarb3-dummy-Idle				/Data/Global/Monsters	7i	NU	HTH		LIT																	0
5	1	626	baalclone-Baal Crab Clone-BaalCrabClone				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
5	1	627	baalminion1-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
5	1	628	baalminion2-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
5	1	629	baalminion3-Baals Minion-BaalMinion				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
5	1	630	worldstoneeffect-dummy-Idle				/Data/Global/Monsters	K9																					0
5	1	631	sk_archer6-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	632	sk_archer7-BoneArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	633	sk_archer8-BurningDeadArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	634	sk_archer9-ReturnedArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	635	sk_archer10-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	636	bighead6-Afflicted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	637	bighead7-Tainted-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	638	bighead8-Misshapen-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	639	bighead9-Disfigured-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	640	bighead10-Damned-Bighead				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	641	goatman6-MoonClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	642	goatman7-NightClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	643	goatman8-HellClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	644	goatman9-BloodClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	645	goatman10-DeathClan-Goatman				/Data/Global/Monsters	GM	NU	2HS		LIT				HAL													0
5	1	646	foulcrow5-FoulCrow-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	647	foulcrow6-BloodHawk-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	648	foulcrow7-BlackRaptor-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	649	foulcrow8-CloudStalker-BloodHawk				/Data/Global/Monsters	BK	NU	HTH		LIT																	0
5	1	650	clawviper6-ClawViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	651	clawviper7-PitViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	652	clawviper8-Salamander-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	653	clawviper9-TombViper-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	654	clawviper10-SerpentMagus-ClawViperEx				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	655	sandraider6-Marauder-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	656	sandraider7-Infidel-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	657	sandraider8-SandRaider-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	658	sandraider9-Invader-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	659	sandraider10-Assailant-SandRaider				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	660	deathmauler6-Death Mauler1-DeathMauler				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	661	quillrat6-QuillRat-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	662	quillrat7-SpikeFiend-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	663	quillrat8-RazorSpine-QuillRat				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	664	vulture5-CarrionBird-Vulture				/Data/Global/Monsters	VD	NU	HTH		LIT																	0
5	1	665	thornhulk5-ThornedHulk-ThornHulk				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
5	1	666	slinger7-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	667	slinger8-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	668	slinger9-Slinger-PantherJavelin				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	669	cr_archer6-VileArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	670	cr_archer7-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	671	cr_lancer6-VileLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	672	cr_lancer7-DarkLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	673	cr_lancer8-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	674	blunderbore5-Blunderbore-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
5	1	675	blunderbore6-Mauler-PinHead				/Data/Global/Monsters	PN	NU	HTH		LIT																	0
5	1	676	skmage_fire5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
5	1	677	skmage_fire6-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		FIR	FIR						0
5	1	678	skmage_ltng5-ReturnedMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
5	1	679	skmage_ltng6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		LHT	LHT						0
5	1	680	skmage_cold5-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		CLD	CLD						0
5	1	681	skmage_pois5-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	682	skmage_pois6-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	683	pantherwoman5-Huntress-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	684	pantherwoman6-SaberCat-PantherWoman				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	685	sandleaper6-CaveLeaper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	686	sandleaper7-TombCreeper-SandLeaper				/Data/Global/Monsters	SL	NU	HTH		LIT																	0
5	1	687	wraith6-Ghost-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	688	wraith7-Wraith-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	689	wraith8-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	690	succubus6-Succubusexp-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	691	succubus7-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	692	succubuswitch6-Dominus-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	693	succubuswitch7-Hell Witch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	694	succubuswitch8-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	695	willowisp5-Gloam-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	696	willowisp6-BlackSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	697	willowisp7-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	698	fallen6-Carver-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
5	1	699	fallen7-Devilkin-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
5	1	700	fallen8-DarkOne-Fallen				/Data/Global/Monsters	FA	NU	HTH		LIT				CLB		BUC	LIT										0
5	1	701	fallenshaman6-CarverShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	702	fallenshaman7-DevilkinShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	703	fallenshaman8-DarkShaman-FallenShaman				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	704	skeleton6-BoneWarrior-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	705	skeleton7-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	706	batdemon6-Gloombat-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	707	batdemon7-Fiend-BatDemon				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	708	bloodlord6-Blood Lord1-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	709	bloodlord7-Blood Lord4-BloodLord				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	710	scarab6-Scarab-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	711	scarab7-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	712	fetish6-Flayer-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	713	fetish7-StygianDoll-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	714	fetish8-SoulKiller-Fetish				/Data/Global/Monsters	FE	NU	HTH		LIT				FBL													0
5	1	715	fetishblow6-Flayer-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	716	fetishblow7-StygianDoll-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	717	fetishblow8-SoulKiller-FetishBlowgun				/Data/Global/Monsters	FC	NU	HTH		LIT																	0
5	1	718	fetishshaman6-FlayerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	719	fetishshaman7-StygianDollShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	720	fetishshaman8-SoulKillerShaman-FetishShaman				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	721	baboon7-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	722	baboon8-TempleGuard-Baboon				/Data/Global/Monsters	BB	NU	HTH		LIT																	0
5	1	723	unraveler6-Guardian-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	724	unraveler7-Unraveler-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	725	unraveler8-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	726	unraveler9-Horadrim Ancient-GreaterMummy				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	727	zealot4-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
5	1	728	zealot5-Zealot-ZakarumZealot				/Data/Global/Monsters	ZZ	NU	HTH	HD1	ZZ5							HAL										0
5	1	729	cantor5-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
5	1	730	cantor6-Heirophant-ZakarumPriest				/Data/Global/Monsters	ZP	NU	HTH		LIT																	0
5	1	731	vilemother4-Grotesque-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
5	1	732	vilemother5-FleshSpawner-VileMother				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
5	1	733	vilechild4-GrotesqueWyrm-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
5	1	734	vilechild5-FleshBeast-VileDog				/Data/Global/Monsters	VC	NU	HTH		LIT																	0
5	1	735	sandmaggot6-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	736	maggotbaby6-WorldKillerYoung-MaggotLarva				/Data/Global/Monsters	SB	NU	HTH		LIT																	0
5	1	737	maggotegg6-WorldKillerEgg-MaggotEgg				/Data/Global/Monsters	SE	NU	HTH		LIT																	0
5	1	738	minion9-Slayerexp-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	739	minion10-HellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	740	minion11-GreaterHellSpawn-Minion				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	741	arach6-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	742	megademon4-Balrog-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	743	megademon5-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	744	imp6-Imp1-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	745	imp7-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	746	bonefetish6-Undead StygianDoll-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	747	bonefetish7-Undead SoulKiller-Fetish				/Data/Global/Monsters	FK	NU	1HS		LIT				FBL													0
5	1	748	fingermage4-Strangler-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
5	1	749	fingermage5-StormCaster-FingerMage				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
5	1	750	regurgitator4-MawFiend-Regurgitator				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
5	1	751	vampire6-BloodLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	752	vampire7-GhoulLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	753	vampire8-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	754	reanimatedhorde6-UnholyCorpse-ReanimatedHorde				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	755	dkfig1-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	756	dkfig2-DoomKnight-DoomKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	757	dkmag1-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	758	dkmag2-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	759	mummy6-Cadaver-Mummy				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	760	ubermephisto-Mephisto-UberMephisto				/Data/Global/Monsters	MP	NU	HTH		LIT		LIT	LIT														0
5	1	761	uberdiablo-Diablo-UberDiablo				/Data/Global/Monsters	DI	NU	HTH	LIT	LIT	LIT	LIT	LIT														0
5	1	762	uberizual-izual-UberIzual				/Data/Global/Monsters	22	NU	HTH		LIT																	0
5	1	763	uberandariel-Lilith-Andariel				/Data/Global/Monsters	AN	NU	HTH		LIT																	0
5	1	764	uberduriel-Duriel-Duriel				/Data/Global/Monsters	DU	NU	HTH		LIT	LIT	LIT	LIT														0
5	1	765	uberbaal-Baal Crab-UberBaal				/Data/Global/Monsters	42	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT										0
5	1	766	demonspawner-Evil hut-MinionSpawner				/Data/Global/Monsters	xa	NU	HTH		LIT							LIT	LIT	LIT								0
5	1	767	demonhole-Dummy-EvilHole				/Data/Global/Monsters	EH	S4	HTH		LIT							LIT										0
5	1	768	megademon6-PitLord-Megademon				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	769	dkmag3-OblivionKnight-OblivionKnight				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	770	imp8-Imp4-Imp				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	771	swarm5-HellSwarm-Swarm				/Data/Global/Monsters	SW	NU	HTH		LIT																	0
5	1	772	sandmaggot7-WorldKiller-SandMaggot				/Data/Global/Monsters	SM	NU	HTH		LIT																	0
5	1	773	arach7-Arach-Arach				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	774	scarab8-SteelWeevil-Scarab				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	775	succubus8-Hell Temptress-Succubus				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	776	succubuswitch9-VileWitch-SuccubusWitch				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	777	corruptrogue6-FleshHunter-CorruptRogue				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	778	cr_archer8-DarkArcher-CorruptArcher				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	779	cr_lancer9-BlackLancer-CorruptLancer				/Data/Global/Monsters	CR	NU	2HT	HVY	HVY	HVY	HVY	HVY	PIK			HVY	HVY									0
5	1	780	overseer6-HellWhip-Overseer				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	781	skeleton8-Returned-Skeleton				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	782	sk_archer11-HorrorArcher-SkeletonBow				/Data/Global/Monsters	SK	NU	BOW	HVY	HVY	HVY	HVY	HVY		SBW		HVY	HVY									0
5	1	783	skmage_fire7-BurningDeadMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		FIR	FIR						0
5	1	784	skmage_ltng7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		LHT	LHT						0
5	1	785	skmage_cold6-BoneMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		CLD	CLD						0
5	1	786	skmage_pois7-HorrorMage-SkeletonMage				/Data/Global/Monsters	SK	NU	HTH	HVY	HVY	DES	DES	DES				DES	DES		POS	POS						0
5	1	787	vampire9-DarkLord-Vampire				/Data/Global/Monsters	VA	NU	HTH		LIT																	0
5	1	788	wraith9-Specter-Wraith				/Data/Global/Monsters	WR	NU	HTH		LIT																	0
5	1	789	willowisp8-BurningSoul-WillOWisp				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	790	Bishibosh-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	791	Bonebreak-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BUC	HVY	HVY	LIT								0
5	1	792	Coldcrow-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	BOW	HVY	HVY	HVY	HVY	HVY	LIT	LBW		HVY	HVY									0
5	1	793	Rakanishu-SUPER UNIQUE				/Data/Global/Monsters	FA	NU	HTH		LIT				SWD		TCH	LIT										0
5	1	794	Treehead WoodFist-SUPER UNIQUE				/Data/Global/Monsters	YE	NU	HTH		LIT																	0
5	1	795	Griswold-SUPER UNIQUE				/Data/Global/Monsters	GZ	NU	HTH		LIT																	0
5	1	796	The Countess-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	MED	LIT	MED	LIT	LIT	WHM			LIT	LIT									0
5	1	797	Pitspawn Fouldog-SUPER UNIQUE				/Data/Global/Monsters	BH	NU	HTH		LIT																	0
5	1	798	Flamespike the Crawler-SUPER UNIQUE				/Data/Global/Monsters	SI	NU	HTH		LIT																	0
5	1	799	Boneash-SUPER UNIQUE				/Data/Global/Monsters	SK	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT		POS	POS						0
5	1	800	Radament-SUPER UNIQUE				/Data/Global/Monsters	RD	NU	HTH		LIT																	0
5	1	801	Bloodwitch the Wild-SUPER UNIQUE				/Data/Global/Monsters	PW	NU	1HT	BAB	HVY		HVY	HVY		GPL	BUC	HVY	HVY	HVY	HVY							0
5	1	802	Fangskin-SUPER UNIQUE				/Data/Global/Monsters	SD	NU	HTH		LIT																	0
5	1	803	Beetleburst-SUPER UNIQUE				/Data/Global/Monsters	SC	NU	HTH	LIT	LIT		HVY															0
5	1	804	Leatherarm-SUPER UNIQUE				/Data/Global/Monsters	MM	NU	HTH		LIT							LIT										0
5	1	805	Coldworm the Burrower-SUPER UNIQUE				/Data/Global/Monsters	MQ	NU	HTH		LIT																	0
5	1	806	Fire Eye-SUPER UNIQUE				/Data/Global/Monsters	SR	NU	HTH		LIT																	0
5	1	807	Dark Elder-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	808	The Summoner-SUPER UNIQUE				/Data/Global/Monsters	SU	NU	HTH		LIT																	0
5	1	809	Ancient Kaa the Soulless-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	810	The Smith-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
5	1	811	Web Mage the Burning-SUPER UNIQUE				/Data/Global/Monsters	SP	NU	HTH		LIT																	0
5	1	812	Witch Doctor Endugu-SUPER UNIQUE				/Data/Global/Monsters	FW	NU	HTH		LIT																	0
5	1	813	Stormtree-SUPER UNIQUE				/Data/Global/Monsters	TH	NU	HTH	LIT	LIT		LIT	LIT														0
5	1	814	Sarina the Battlemaid-SUPER UNIQUE				/Data/Global/Monsters	CR	NU	1HS	HVY	HVY	HVY	HVY	HVY	AXE		BRV	HVY	HVY									0
5	1	815	Icehawk Riftwing-SUPER UNIQUE				/Data/Global/Monsters	BT	NU	HTH		LIT																	0
5	1	816	Ismail Vilehand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	817	Geleb Flamefinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	818	Bremm Sparkfist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	819	Toorc Icefist-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	820	Wyand Voidfinger-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	821	Maffer Dragonhand-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	822	Winged Death-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	823	The Tormentor-SUPER UNIQUE				/Data/Global/Monsters	WW	NU	HTH		LIT																	0
5	1	824	Taintbreeder-SUPER UNIQUE				/Data/Global/Monsters	VM	NU	HTH		LIT																	0
5	1	825	Riftwraith the Cannibal-SUPER UNIQUE				/Data/Global/Monsters	CS	NU	HTH		LIT																	0
5	1	826	Infector of Souls-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	827	Lord De Seis-SUPER UNIQUE				/Data/Global/Monsters	UM	NU	HTH	HRN	LIT		MED	MED		BSD		RSP	LSP	UNH	POS							0
5	1	828	Grand Vizier of Chaos-SUPER UNIQUE				/Data/Global/Monsters	FR	NU	HTH		LIT							LIT										0
5	1	829	The Cow King-SUPER UNIQUE				/Data/Global/Monsters	EC	NU	HTH		LIT				BTX													0
5	1	830	Corpsefire-SUPER UNIQUE				/Data/Global/Monsters	ZM	NU	HTH	HVY	HVY	LIT	LIT	LIT				LIT	LIT	BLD								0
5	1	831	The Feature Creep-SUPER UNIQUE				/Data/Global/Monsters	5P	NU	HTH		LIT																	0
5	1	832	Siege Boss-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	833	Ancient Barbarian 1-SUPER UNIQUE				/Data/Global/Monsters	0D	NU	HTH		LIT							LIT	LIT									0
5	1	834	Ancient Barbarian 2-SUPER UNIQUE				/Data/Global/Monsters	0F	NU	HTH		LIT								LIT									0
5	1	835	Ancient Barbarian 3-SUPER UNIQUE				/Data/Global/Monsters	0E	NU	HTH		LIT								LIT									0
5	1	836	Axe Dweller-SUPER UNIQUE				/Data/Global/Monsters	L3	NU	HTH	HEV	LIT	HEV	HEV	HEV	FLA	FLA		HEV	HEV									0
5	1	837	Bonesaw Breaker-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	838	Dac Farren-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	839	Megaflow Rectifier-SUPER UNIQUE				/Data/Global/Monsters	xx	NU	HTH	HVY	LIT				HVY		HVY											0
5	1	840	Eyeback Unleashed-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	841	Threash Socket-SUPER UNIQUE				/Data/Global/Monsters	ox	NU	HTH		LIT		LIT	LIT				LIT	LIT	LIT	LIT			LIT				0
5	1	842	Pindleskin-SUPER UNIQUE				/Data/Global/Monsters	re	NU	HTH	HVY	LIT	HVY	HVY	HVY	CLM			HVY	HVY									0
5	1	843	Snapchip Shatter-SUPER UNIQUE				/Data/Global/Monsters	f0	NU	HTH		LIT							LIT										0
5	1	844	Anodized Elite-SUPER UNIQUE				/Data/Global/Monsters	0B	NU	HTH		LIT																	0
5	1	845	Vinvear Molech-SUPER UNIQUE				/Data/Global/Monsters	0C	NU	HTH		LIT																	0
5	1	846	Sharp Tooth Sayer-SUPER UNIQUE				/Data/Global/Monsters	os	NU	HTH	HVY	HVY		HVY	HVY		LIT		HVY	HVY									0
5	1	847	Magma Torquer-SUPER UNIQUE				/Data/Global/Monsters	ip	NU	HTH		LIT																	0
5	1	848	Blaze Ripper-SUPER UNIQUE				/Data/Global/Monsters	m5	NU	HTH		LIT																	0
5	1	849	Frozenstein-SUPER UNIQUE				/Data/Global/Monsters	io	NU	HTH		LIT																	0
5	1	850	Nihlathak Boss-SUPER UNIQUE				/Data/Global/Monsters	XU	NU	HTH		LIT																	0
5	1	851	Baal Subject 1-SUPER UNIQUE				/Data/Global/Monsters	FS	NU	HTH		LIT																	0
5	1	852	Baal Subject 2-SUPER UNIQUE				/Data/Global/Monsters	GY	NU	HTH		LIT																	0
5	1	853	Baal Subject 3-SUPER UNIQUE				/Data/Global/Monsters	HP	NU	HTH		LIT																	0
5	1	854	Baal Subject 4-SUPER UNIQUE				/Data/Global/Monsters	DM	NU	HTH		LIT				WSC													0
5	1	855	Baal Subject 5-SUPER UNIQUE				/Data/Global/Monsters	43	NU	HTH	LIT	LIT	LIT	LIT	LIT				LIT	LIT									0
5	2	0	banner 1  (452)	452			/Data/Global/Objects	AO	NU	HTH		LIT																	0
5	2	1	banner 2  (453)	453			/Data/Global/Objects	AP	NU	HTH		LIT																	0
5	2	2	guild vault  (338)	338			/Data/Global/Objects	Y4	NU	HTH		LIT																	0
5	2	3	steeg stone  (337)	337			/Data/Global/Objects	Y6	NU	HTH		LIT							LIT										0
5	2	4	Your Private Stash  (267)	267			/Data/Global/Objects	B6	NU	HTH		LIT																	0
5	2	5	fog water  (374)	374			/Data/Global/Objects	UD	NU	HTH		LIT																	0
5	2	6	torch, expansion tiki 1 (482)	482			/Data/Global/Objects	2P	NU	HTH		LIT							LIT										0
5	2	7	Fire, Rogue camp (39)	39			/Data/Global/Objects	RB	ON	HTH		LIT																	0
5	2	8	standard / direction (35)	35			/Data/Global/Objects	N1	NU	HTH		LIT																	0
5	2	9	standard / direction ((36)	36			/Data/Global/Objects	N2	NU	HTH		LIT																	0
5	2	10	candles R (33)	33			/Data/Global/Objects	A1	NU	HTH		LIT																	0
5	2	11	candles L (34)	34			/Data/Global/Objects	A2	NU	HTH		LIT																	0
5	2	12	torch 2 wall  (38)	38			/Data/Global/Objects	WT	ON	HTH		LIT																	0
5	2	13	floor brazier  (102)	102			/Data/Global/Objects	FB	ON	HTH		LIT							LIT										0
5	2	14	Pot O Torch, level 1 (411)	411			/Data/Global/Objects	PX	NU	HTH		LIT							LIT	LIT									0
5	2	15	burning bodies  (438)	438			/Data/Global/Objects	6F	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					0
5	2	16	fire pit, level 1 (412)	412			/Data/Global/Objects	PY	NU	HTH		LIT							LIT										0
5	2	17	camp fire (435)	435			/Data/Global/Objects	2W	NU	HTH		LIT							LIT	LIT									0
5	2	18	town torch  (436)	436			/Data/Global/Objects	2X	NU	HTH		LIT							LIT	LIT									0
5	2	19	Tribal Flag  (440)	440			/Data/Global/Objects	6H	NU	HTH		LIT																	0
5	2	20	Town Flag  (441)	441			/Data/Global/Objects	2Y	NU	HTH		LIT																	0
5	2	21	Town Flag  (441)	441			/Data/Global/Objects	2Y	NU	HTH		LIT																	0
5	2	22	Chandelier  (442)	442			/Data/Global/Objects	2Z	NU	HTH		LIT							LIT										0
5	2	23	Waypoint  (429)	429			/Data/Global/Objects	YY	ON	HTH		LIT							LIT										0
5	2	24	Wooden Chest L  (420)	420			/Data/Global/Objects	YP	OP	HTH		LIT																	0
5	2	25	Wooden Chest R  (431)	431			/Data/Global/Objects	6A	OP	HTH		LIT																	0
5	2	26	Chest L  (430)	430			/Data/Global/Objects	YZ	OP	HTH		LIT																	0
5	2	27	Chest R  (413)	413			/Data/Global/Objects	6Q	OP	HTH		LIT																	0
5	2	28	Chest S L  (432)	432			/Data/Global/Objects	6B	OP	HTH		LIT																	0
5	2	29	Chest S R  (433)	433			/Data/Global/Objects	6C	OP	HTH		LIT																	0
5	2	30	wilderness barrel  (418)	418			/Data/Global/Objects	YN	OP	HTH		LIT																	0
5	2	31	exploding wildernes barrel  (419)	419			/Data/Global/Objects	YO	OP	HTH		LIT							LIT										0
5	2	32	Burial Chest  (424)	424			/Data/Global/Objects	YT	OP	HTH		LIT																	0
5	2	33	Burial Chest R (425)	425			/Data/Global/Objects	YU	OP	HTH		LIT																	0
5	2	34	Hidden Stash  (416)	416			/Data/Global/Objects	3W	OP	HTH		LIT																	0
5	2	35	Shrine 1  (414)	414			/Data/Global/Objects	6R	OP	HTH		LIT							LIT										0
5	2	36	Shrine 2  (415)	415			/Data/Global/Objects	6S	NU	HTH		LIT							LIT										0
5	2	37	Shrine 3  (427)	427			/Data/Global/Objects	YW	OP	HTH		LIT							LIT	LIT									0
5	2	38	Shrine 4  (428)	428			/Data/Global/Objects	YX	OP	HTH		LIT							LIT										0
5	2	39	Shrine 5  (421)	421			/Data/Global/Objects	YQ	NU	HTH		LIT							LIT										0
5	2	40	Shrine, mana  (422)	422			/Data/Global/Objects	YR	OP	HTH		LIT							LIT										0
5	2	41	Health Shrine  (423)	423			/Data/Global/Objects	YS	OP	HTH		LIT																	0
5	2	42	Well  (426)	426			/Data/Global/Objects	YV	NU	HTH		LIT																	0
5	2	43	Hell Gate  (451)	451			/Data/Global/Objects	6P	NU	HTH		LIT							LIT	LIT									0
5	2	44	Your Private Stash  (267)	267			/Data/Global/Objects	B6	NU	HTH		LIT																	0
5	2	45	Jar 1  (443)	443			/Data/Global/Objects	6I	OP	HTH		LIT																	0
5	2	46	Jar 2  (444)	444			/Data/Global/Objects	6J	OP	HTH		LIT																	0
5	2	47	Jar 3  (445)	445			/Data/Global/Objects	6K	OP	HTH		LIT																	0
5	2	48	Swinging Heads  (446)	446			/Data/Global/Objects	6L	NU	HTH		LIT																	0
5	2	49	Pole  (447)	447			/Data/Global/Objects	6M	NU	HTH		LIT																	0
5	2	50	Skulls and Rocks, no snow  (448)	448			/Data/Global/Objects	6N	NU	HTH		LIT																	0
5	2	51	Skulls and Rocks, siege  (450)	450			/Data/Global/Objects	6O	NU	HTH		LIT																	0
5	2	52	Hell Gate  (451)	451			/Data/Global/Objects	6P	NU	HTH		LIT							LIT	LIT									0
5	2	53	Anya start in town  (459)	459		1	/Data/Global/Monsters	XS	NU	HTH		LIT																	0
5	2	54	Anya start outside town  (460)	460			/Data/Global/Objects	2N	NU	HTH		LIT																	0
5	2	55	Nihlathak start in town  (461)	461		1	/Data/Global/Monsters	0J	NU	HTH		LIT																	0
5	2	56	Nihlathak outside in town  (462)	462		7	/Data/Global/Monsters	0J	NU	HTH		LIT																	0
5	2	57	Torch, expansion tiki torch  (482)	482			/Data/Global/Objects	2P	NU	HTH		LIT							LIT										0
5	2	58	Cage, caged fellow  (473)	473																									0
5	2	59	Chest, specialchest  (455)	455			/Data/Global/Objects	6U	OP	HTH		LIT																	0
5	2	60	Death Pole 1, wilderness (456)	456			/Data/Global/Objects	6V	NU	HTH		LIT																	0
5	2	61	Death Pole 2, wilderness (457)	457			/Data/Global/Objects	6W	NU	HTH		LIT																	0
5	2	62	Altar, inside of temple  (458)	458			/Data/Global/Objects	6X	NU	HTH		LIT							LIT										0
5	2	63	Hidden Stash, icecave  (463)	463			/Data/Global/Objects	6Y	OP	HTH		LIT																	0
5	2	64	Health Shrine, icecave  (464)	464			/Data/Global/Objects	8A	OP	HTH		LIT																	0
5	2	65	Shrine, icecave  (465)	465			/Data/Global/Objects	8B	OP	HTH		LIT																	0
5	2	66	Evil Urn, icecave  (466)	466			/Data/Global/Objects	8C	OP	HTH		LIT																	0
5	2	67	Jar, icecave 1  (467)	467			/Data/Global/Objects	8D	OP	HTH		LIT																	0
5	2	68	Jar, icecave 2  (468)	468			/Data/Global/Objects	8E	OP	HTH		LIT																	0
5	2	69	Jar, icecave 3  (469)	469			/Data/Global/Objects	8F	OP	HTH		LIT																	0
5	2	70	Jar, icecave 4  (470)	470			/Data/Global/Objects	8G	OP	HTH		LIT																	0
5	2	71	Jar, icecave 5  (471)	471			/Data/Global/Objects	8H	OP	HTH		LIT																	0
5	2	72	Shrine, icecave 1  (472)	472			/Data/Global/Objects	8I	OP	HTH		LIT							LIT	LIT									0
5	2	73	Dead Barbarian, seige/wilderness  (477)	477			/Data/Global/Objects	8J	OP	HTH		LIT																	0
5	2	74	Shrine, icecave 2  (479)	479			/Data/Global/Objects	8K	OP	HTH		LIT							LIT	LIT									0
5	2	75	Torch, icecave 1  (480)	480			/Data/Global/Objects	8L	NU	HTH		LIT							LIT										0
5	2	76	Torch, icecave 2  (481)	481			/Data/Global/Objects	8M	NU	HTH		LIT							LIT										0
5	2	77	Shrine, baals  (483)	483			/Data/Global/Objects	8N	OP	HTH		LIT																	0
5	2	78	Health Shrine, baals  (484)	484			/Data/Global/Objects	8O	OP	HTH		LIT																	0
5	2	79	Tomb, baal's lair 1  (485)	485			/Data/Global/Objects	8P	OP	HTH		LIT																	0
5	2	80	Tomb, baal's lair 2  (486)	486			/Data/Global/Objects	8Q	OP	HTH		LIT																	0
5	2	81	Tomb, baal's lair 3  (487)	487			/Data/Global/Objects	8R	OP	HTH		LIT																	0
5	2	82	Chest, wilderness/siege exploding  (454)	454			/Data/Global/Objects	6T	OP	HTH		LIT							LIT										0
5	2	83	Torch, expansion no snow  (437)	437			/Data/Global/Objects	6E	NU	HTH		LIT							LIT										0
5	2	84	Stash, Pen breakable door  (508)	508			/Data/Global/Objects	2Q	OP	HTH		LIT																	0
5	2	85	Magic Shrine, baal's lair  (488)	488			/Data/Global/Objects	8S	OP	HTH		LIT							LIT										0
5	2	86	Well, snowy  (493)	493			/Data/Global/Objects	8X	NU	HTH		LIT																	0
5	2	87	Well, snowy  (493)	493			/Data/Global/Objects	8X	NU	HTH		LIT																	0
5	2	88	Magic Shrine, snowy_shrine3 a  (495)	495			/Data/Global/Objects	8Z	OP	HTH		LIT							LIT										0
5	2	89	Magic Shrine, snowy_shrine3 b  (497)	497			/Data/Global/Objects	5B	OP	HTH		LIT							LIT	LIT									0
5	2	90	Magic Shrine, baal's lair  (499)	499			/Data/Global/Objects	5D	OP	HTH		LIT							LIT										0
5	2	91	Magic Shrine, baals_shrine3  (503)	503			/Data/Global/Objects	5H	OP	HTH		LIT							LIT										0
5	2	92	Magic Shrine, temple 1  (509)	509			/Data/Global/Objects	5M	OP	HTH		LIT							LIT										0
5	2	93	Magic Shrine, temple 2  (512)	512			/Data/Global/Objects	5T	OP	HTH		LIT							LIT										0
5	2	94	Torch, baal's lair 1  (489)	489			/Data/Global/Objects	8T	NU	HTH		LIT							LIT										0
5	2	95	Torch, baal's lair 2  (490)	490			/Data/Global/Objects	8U	NU	HTH		LIT							LIT										0
5	2	96	Torch, temple 1  (514)	514			/Data/Global/Objects	5R	NU	HTH		LIT							LIT										0
5	2	97	Torch, temple 2  (515)	515			/Data/Global/Objects	5S	NU	HTH		LIT							LIT										0
5	2	98	Well, snowy  (493)	493			/Data/Global/Objects	8X	NU	HTH		LIT																	0
5	2	99	Well, baalslair  (498)	498			/Data/Global/Objects	5C	NU	HTH		LIT																	0
5	2	100	Well, temple  (513)	513			/Data/Global/Objects	5Q	NU	HTH		LIT																	0
5	2	101	Waypoint, baals_waypoint  (494)	494			/Data/Global/Objects	8Y	ON	HTH		LIT							LIT										0
5	2	102	Waypoint, wilderness_waypoint  (496)	496			/Data/Global/Objects	5A	ON	HTH		LIT							LIT										0
5	2	103	Waypoint, icecave  (511)	511			/Data/Global/Objects	5O	ON	HTH		LIT							LIT										0
5	2	104	Hidden Stash, snowy  (500)	500			/Data/Global/Objects	5E	OP	HTH		LIT																	0
5	2	105	Wooden Chest, snowy L  (501)	501			/Data/Global/Objects	5F	OP	HTH		LIT																	0
5	2	106	Wooden Chest, snowy R  (502)	502			/Data/Global/Objects	5G	OP	HTH		LIT																	0
5	2	107	Wooden Chest, snowy L 2  (504)	504			/Data/Global/Objects	5I	OP	HTH		LIT																	0
5	2	108	Wooden Chest, snowy R 2  (505)	505			/Data/Global/Objects	5J	OP	HTH		LIT																	0
5	2	109	Swinging Heads, snowy  (506)	506			/Data/Global/Objects	5K	NU	HTH		LIT																	0
5	2	110	Debris, snowy  (507)	507			/Data/Global/Objects	5L	NU	HTH		LIT																	0
5	2	111	Pole, snowy  (510)	510			/Data/Global/Objects	5N	NU	HTH		LIT																	0
5	2	112	Fire, fire small  (160)	160			/Data/Global/Objects	FX	NU	HTH		LIT																	0
5	2	113	Fire, fire medium  (161)	161			/Data/Global/Objects	FY	NU	HTH		LIT																	0
5	2	114	Fire, fire large  (162)	162			/Data/Global/Objects	FZ	NU	HTH		LIT																	0
5	2	115	gold placeholder  (269)	269																									0
5	2	116	Red Light, (touch me) for blacksmith  (523)	523																									0
5	2	117	Torch, expansion no snow  (434)	434			/Data/Global/Objects	6D	NU	HTH		LIT							LIT										0
5	2	118	Waypoint, wilderness_waypoint 1  (496)	496			/Data/Global/Objects	5A	ON	HTH		LIT							LIT										0
5	2	119	Waypoint, wilderness_waypoint 2  (496)	496			/Data/Global/Objects	5A	ON	HTH		LIT							LIT										0
5	2	120	Waypoint, wilderness_waypoint 3  (496)	496			/Data/Global/Objects	5A	ON	HTH		LIT							LIT										0
5	2	121	Shrub, Ice cave bubbles 01  (527)	527			/Data/Global/Objects	2U	NU	HTH		LIT																	0
5	2	122	Shrub, Ice cave bubbles 01  (528)	528			/Data/Global/Objects	2S	OP	HTH		LIT																	0
5	2	123	Candles, temple  (538)	538			/Data/Global/Objects	3O	NU	HTH		LIT							LIT										0
5	2	124	Waypoint, temple  (539)	539			/Data/Global/Objects	3P	ON	HTH		LIT							LIT										0
5	2	125	Larzuk Greeting  (542)	542		1	/Data/Global/Monsters	XR	NU	HTH		LIT																	0
5	2	126	Larzuk Standard  (543)	543			/Data/Global/Monsters	XR	NU	HTH		LIT																	0
5	2	127	Altar of the Heavens, ancientsaltar  (546)	546			/Data/Global/Objects	4A	OP	HTH		LIT																	0
5	2	128	door, ancient To Worldstone lev 1  (547)	547			/Data/Global/Objects	4B	OP	HTH		LIT																	0
5	2	129	Weapon Rack, R  (548)	548			/Data/Global/Objects	3X	NU	HTH		LIT																	0
5	2	130	Weapon Rack, L  (549)	549			/Data/Global/Objects	3Y	NU	HTH		LIT																	0
5	2	131	Armor Stand, R  (550)	550			/Data/Global/Objects	3Z	NU	HTH		LIT																	0
5	2	132	Armor Stand, L  (551)	551			/Data/Global/Objects	4C	NU	HTH		LIT																	0
5	2	133	Torch, summit  (552)	552			/Data/Global/Objects	9G	NU	HTH		LIT							LIT										0
5	2	134	Ice cave steam  (555)	555			/Data/Global/Objects	2O	NU	HTH		LIT																	0
5	2	135	funeralpire  (553)	553			/Data/Global/Objects	9H	NU	HTH		LIT							LIT										0
5	2	136	burninglogs  (554)	554			/Data/Global/Objects	9I	NU	HTH		LIT							LIT										0
5	2	137	dummy, Baal's lair  (557)	557																									0
5	2	138	Tomb, temple ground  (541)	541			/Data/Global/Objects	3S	OP	HTH		LIT																	0
5	2	139	Tomb, temple ground L  (544)	544			/Data/Global/Objects	3T	OP	HTH		LIT																	0
5	2	140	BBQ Bunny  (559)	559			/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									0
5	2	141	Baal Torch Big  (560)	560			/Data/Global/Objects	25	NU	HTH		LIT							LIT										0
5	2	142	The Ancients' Way, summit door  (564)	564			/Data/Global/Objects	4U	OP	HTH		LIT																	0
5	2	143	test data, zoo  (567)	567																									0
5	2	144	test data, keeper  (568)	568			/Data/Global/Objects	7Z	NU	HTH		LIT																	0
5	2	145	Torch, redbaal's lair 1  (536)	536			/Data/Global/Objects	3M	NU	HTH		LIT							LIT										0
5	2	146	Torch, redbaal's lair 2  (537)	537			/Data/Global/Objects	3N	NU	HTH		LIT							LIT										0
5	2	147	The Worldstone Chamber, baals portal  (563)	563			/Data/Global/Objects	4X	ON	HTH		LIT							LIT										0
5	2	148	fire place guy  (570)	570			/Data/Global/Objects	7Y	NU	HTH		LIT																	0
5	2	149	Chest, spark (397)	397			/Data/Global/Objects	YF	OP	HTH		LIT																	0
5	2	150	Dummy-test data SKIPT IT				/Data/Global/Objects	NU0																					
5	2	151	Casket-Casket #5				/Data/Global/Objects	C5	OP	HTH		LIT																	
5	2	152	Shrine-Shrine				/Data/Global/Objects	SF	OP	HTH		LIT																	
5	2	153	Casket-Casket #6				/Data/Global/Objects	C6	OP	HTH		LIT																	
5	2	154	LargeUrn-Urn #1				/Data/Global/Objects	U1	OP	HTH		LIT																	
5	2	155	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
5	2	156	chest-LargeChestL				/Data/Global/Objects	L2	OP	HTH		LIT																	
5	2	157	Barrel-Barrel				/Data/Global/Objects	B1	OP	HTH		LIT																	
5	2	158	TowerTome-Tower Tome				/Data/Global/Objects	TT	OP	HTH		LIT																	
5	2	159	Urn-Urn #2				/Data/Global/Objects	U2	OP	HTH		LIT																	
5	2	160	Dummy-Bench				/Data/Global/Objects	BE	NU	HTH		LIT																	
5	2	161	Barrel-BarrelExploding				/Data/Global/Objects	BX	OP	HTH		LIT							LIT	LIT									
5	2	162	Dummy-RogueFountain				/Data/Global/Objects	FN	NU	HTH		LIT																	
5	2	163	Door-Door Gate Left				/Data/Global/Objects	D1	OP	HTH		LIT																	
5	2	164	Door-Door Gate Right				/Data/Global/Objects	D2	OP	HTH		LIT																	
5	2	165	Door-Door Wooden Left				/Data/Global/Objects	D3	OP	HTH		LIT																	
5	2	166	Door-Door Wooden Right				/Data/Global/Objects	D4	OP	HTH		LIT																	
5	2	167	StoneAlpha-StoneAlpha				/Data/Global/Objects	S1	OP	HTH		LIT																	
5	2	168	StoneBeta-StoneBeta				/Data/Global/Objects	S2	OP	HTH		LIT																	
5	2	169	StoneGamma-StoneGamma				/Data/Global/Objects	S3	OP	HTH		LIT																	
5	2	170	StoneDelta-StoneDelta				/Data/Global/Objects	S4	OP	HTH		LIT																	
5	2	171	StoneLambda-StoneLambda				/Data/Global/Objects	S5	OP	HTH		LIT																	
5	2	172	StoneTheta-StoneTheta				/Data/Global/Objects	S6	OP	HTH		LIT																	
5	2	173	Door-Door Courtyard Left				/Data/Global/Objects	D5	OP	HTH		LIT																	
5	2	174	Door-Door Courtyard Right				/Data/Global/Objects	D6	OP	HTH		LIT																	
5	2	175	Door-Door Cathedral Double				/Data/Global/Objects	D7	OP	HTH		LIT																	
5	2	176	Gibbet-Cain's Been Captured				/Data/Global/Objects	GI	OP	HTH		LIT																	
5	2	177	Door-Door Monastery Double Right				/Data/Global/Objects	D8	OP	HTH		LIT																	
5	2	178	HoleAnim-Hole in Ground				/Data/Global/Objects	HI	OP	HTH		LIT																	
5	2	179	Dummy-Brazier				/Data/Global/Objects	BR	ON	HTH		LIT							LIT										
5	2	180	Inifuss-inifuss tree				/Data/Global/Objects	IT	NU	HTH		LIT																	
5	2	181	Dummy-Fountain				/Data/Global/Objects	BF	NU	HTH		LIT																	
5	2	182	Dummy-crucifix				/Data/Global/Objects	CL	NU	HTH		LIT																	
5	2	183	Dummy-Candles1				/Data/Global/Objects	A1	NU	HTH		LIT																	
5	2	184	Dummy-Candles2				/Data/Global/Objects	A2	NU	HTH		LIT																	
5	2	185	Dummy-Standard1				/Data/Global/Objects	N1	NU	HTH		LIT																	
5	2	186	Dummy-Standard2				/Data/Global/Objects	N2	NU	HTH		LIT																	
5	2	187	Dummy-Torch1 Tiki				/Data/Global/Objects	TO	ON	HTH		LIT																	
5	2	188	Dummy-Torch2 Wall				/Data/Global/Objects	WT	ON	HTH		LIT																	
5	2	189	fire-RogueBonfire				/Data/Global/Objects	RB	ON	HTH		LIT																	
5	2	190	Dummy-River1				/Data/Global/Objects	R1	NU	HTH		LIT																	
5	2	191	Dummy-River2				/Data/Global/Objects	R2	NU	HTH		LIT																	
5	2	192	Dummy-River3				/Data/Global/Objects	R3	NU	HTH		LIT																	
5	2	193	Dummy-River4				/Data/Global/Objects	R4	NU	HTH		LIT																	
5	2	194	Dummy-River5				/Data/Global/Objects	R5	NU	HTH		LIT																	
5	2	195	AmbientSound-ambient sound generator				/Data/Global/Objects	S1	OP	HTH		LIT																	
5	2	196	Crate-Crate				/Data/Global/Objects	CT	OP	HTH		LIT																	
5	2	197	Door-Andariel's Door				/Data/Global/Objects	AD	NU	HTH		LIT																	
5	2	198	Dummy-RogueTorch				/Data/Global/Objects	T1	NU	HTH		LIT																	
5	2	199	Dummy-RogueTorch				/Data/Global/Objects	T2	NU	HTH		LIT																	
5	2	200	Casket-CasketR				/Data/Global/Objects	C1	OP	HTH		LIT																	
5	2	201	Casket-CasketL				/Data/Global/Objects	C2	OP	HTH		LIT																	
5	2	202	Urn-Urn #3				/Data/Global/Objects	U3	OP	HTH		LIT																	
5	2	203	Casket-Casket				/Data/Global/Objects	C4	OP	HTH		LIT																	
5	2	204	RogueCorpse-Rogue corpse 1				/Data/Global/Objects	Z1	NU	HTH		LIT																	
5	2	205	RogueCorpse-Rogue corpse 2				/Data/Global/Objects	Z2	NU	HTH		LIT																	
5	2	206	RogueCorpse-rolling rogue corpse				/Data/Global/Objects	Z5	OP	HTH		LIT																	
5	2	207	CorpseOnStick-rogue on a stick 1				/Data/Global/Objects	Z3	OP	HTH		LIT																	
5	2	208	CorpseOnStick-rogue on a stick 2				/Data/Global/Objects	Z4	OP	HTH		LIT																	
5	2	209	Portal-Town portal				/Data/Global/Objects	TP	ON	HTH	LIT	LIT																	
5	2	210	Portal-Permanent town portal				/Data/Global/Objects	PP	ON	HTH	LIT	LIT																	
5	2	211	Dummy-Invisible object				/Data/Global/Objects	SS																					
5	2	212	Door-Door Cathedral Left				/Data/Global/Objects	D9	OP	HTH		LIT																	
5	2	213	Door-Door Cathedral Right				/Data/Global/Objects	DA	OP	HTH		LIT																	
5	2	214	Door-Door Wooden Left #2				/Data/Global/Objects	DB	OP	HTH		LIT																	
5	2	215	Dummy-invisible river sound1				/Data/Global/Objects	X1																					
5	2	216	Dummy-invisible river sound2				/Data/Global/Objects	X2																					
5	2	217	Dummy-ripple				/Data/Global/Objects	1R	NU	HTH		LIT																	
5	2	218	Dummy-ripple				/Data/Global/Objects	2R	NU	HTH		LIT																	
5	2	219	Dummy-ripple				/Data/Global/Objects	3R	NU	HTH		LIT																	
5	2	220	Dummy-ripple				/Data/Global/Objects	4R	NU	HTH		LIT																	
5	2	221	Dummy-forest night sound #1				/Data/Global/Objects	F1																					
5	2	222	Dummy-forest night sound #2				/Data/Global/Objects	F2																					
5	2	223	Dummy-yeti dung				/Data/Global/Objects	YD	NU	HTH		LIT																	
5	2	224	TrappDoor-Trap Door				/Data/Global/Objects	TD	ON	HTH		LIT																	
5	2	225	Door-Door by Dock, Act 2				/Data/Global/Objects	DD	ON	HTH		LIT																	
5	2	226	Dummy-sewer drip				/Data/Global/Objects	SZ																					
5	2	227	Shrine-healthorama				/Data/Global/Objects	SH	OP	HTH		LIT																	
5	2	228	Dummy-invisible town sound				/Data/Global/Objects	TA																					
5	2	229	Casket-casket #3				/Data/Global/Objects	C3	OP	HTH		LIT																	
5	2	230	Obelisk-obelisk				/Data/Global/Objects	OB	OP	HTH		LIT																	
5	2	231	Shrine-forest altar				/Data/Global/Objects	AF	OP	HTH		LIT																	
5	2	232	Dummy-bubbling pool of blood				/Data/Global/Objects	B2	NU	HTH		LIT																	
5	2	233	Shrine-horn shrine				/Data/Global/Objects	HS	OP	HTH		LIT																	
5	2	234	Shrine-healing well				/Data/Global/Objects	HW	OP	HTH		LIT																	
5	2	235	Shrine-bull shrine,health, tombs				/Data/Global/Objects	BC	OP	HTH		LIT																	
5	2	236	Dummy-stele,magic shrine, stone, desert				/Data/Global/Objects	SG	OP	HTH		LIT																	
5	2	237	Chest3-tombchest 1, largechestL				/Data/Global/Objects	CA	OP	HTH		LIT																	
5	2	238	Chest3-tombchest 2 largechestR				/Data/Global/Objects	CB	OP	HTH		LIT																	
5	2	239	Sarcophagus-mummy coffinL, tomb				/Data/Global/Objects	MC	OP	HTH		LIT																	
5	2	240	Obelisk-desert obelisk				/Data/Global/Objects	DO	OP	HTH		LIT																	
5	2	241	Door-tomb door left				/Data/Global/Objects	TL	OP	HTH		LIT																	
5	2	242	Door-tomb door right				/Data/Global/Objects	TR	OP	HTH		LIT																	
5	2	243	Shrine-mana shrineforinnerhell				/Data/Global/Objects	iz	OP	HTH		LIT							LIT										
5	2	244	LargeUrn-Urn #4				/Data/Global/Objects	U4	OP	HTH		LIT																	
5	2	245	LargeUrn-Urn #5				/Data/Global/Objects	U5	OP	HTH		LIT																	
5	2	246	Shrine-health shrineforinnerhell				/Data/Global/Objects	iy	OP	HTH		LIT							LIT										
5	2	247	Shrine-innershrinehell				/Data/Global/Objects	ix	OP	HTH		LIT							LIT										
5	2	248	Door-tomb door left 2				/Data/Global/Objects	TS	OP	HTH		LIT																	
5	2	249	Door-tomb door right 2				/Data/Global/Objects	TU	OP	HTH		LIT																	
5	2	250	Duriel's Lair-Portal to Duriel's Lair				/Data/Global/Objects	SJ	OP	HTH		LIT																	
5	2	251	Dummy-Brazier3				/Data/Global/Objects	B3	OP	HTH		LIT							LIT										
5	2	252	Dummy-Floor brazier				/Data/Global/Objects	FB	ON	HTH		LIT							LIT										
5	2	253	Dummy-flies				/Data/Global/Objects	FL	NU	HTH		LIT																	
5	2	254	ArmorStand-Armor Stand 1R				/Data/Global/Objects	A3	NU	HTH		LIT																	
5	2	255	ArmorStand-Armor Stand 2L				/Data/Global/Objects	A4	NU	HTH		LIT																	
5	2	256	WeaponRack-Weapon Rack 1R				/Data/Global/Objects	W1	NU	HTH		LIT																	
5	2	257	WeaponRack-Weapon Rack 2L				/Data/Global/Objects	W2	NU	HTH		LIT																	
5	2	258	Malus-Malus				/Data/Global/Objects	HM	NU	HTH		LIT																	
5	2	259	Shrine-palace shrine, healthR, harom, arcane Sanctuary				/Data/Global/Objects	P2	OP	HTH		LIT																	
5	2	260	not used-drinker				/Data/Global/Objects	n5	S1	HTH		LIT																	
5	2	261	well-Fountain 1				/Data/Global/Objects	F3	OP	HTH		LIT																	
5	2	262	not used-gesturer				/Data/Global/Objects	n6	S1	HTH		LIT																	
5	2	263	well-Fountain 2, well, desert, tomb				/Data/Global/Objects	F4	OP	HTH		LIT																	
5	2	264	not used-turner				/Data/Global/Objects	n7	S1	HTH		LIT																	
5	2	265	well-Fountain 3				/Data/Global/Objects	F5	OP	HTH		LIT																	
5	2	266	Shrine-snake woman, magic shrine, tomb, arcane sanctuary				/Data/Global/Objects	SN	OP	HTH		LIT							LIT										
5	2	267	Dummy-jungle torch				/Data/Global/Objects	JT	ON	HTH		LIT							LIT										
5	2	268	Well-Fountain 4				/Data/Global/Objects	F6	OP	HTH		LIT																	
5	2	269	Waypoint-waypoint portal				/Data/Global/Objects	wp	ON	HTH		LIT							LIT										
5	2	270	Dummy-healthshrine, act 3, dungeun				/Data/Global/Objects	dj	OP	HTH		LIT																	
5	2	271	jerhyn-placeholder #1				/Data/Global/Objects	ss																					
5	2	272	jerhyn-placeholder #2				/Data/Global/Objects	ss																					
5	2	273	Shrine-innershrinehell2				/Data/Global/Objects	iw	OP	HTH		LIT							LIT										
5	2	274	Shrine-innershrinehell3				/Data/Global/Objects	iv	OP	HTH		LIT																	
5	2	275	hidden stash-ihobject3 inner hell				/Data/Global/Objects	iu	OP	HTH		LIT																	
5	2	276	skull pile-skullpile inner hell				/Data/Global/Objects	is	OP	HTH		LIT																	
5	2	277	hidden stash-ihobject5 inner hell				/Data/Global/Objects	ir	OP	HTH		LIT																	
5	2	278	hidden stash-hobject4 inner hell				/Data/Global/Objects	hg	OP	HTH		LIT																	
5	2	279	Door-secret door 1				/Data/Global/Objects	h2	OP	HTH		LIT																	
5	2	280	Well-pool act 1 wilderness				/Data/Global/Objects	zw	NU	HTH		LIT																	
5	2	281	Dummy-vile dog afterglow				/Data/Global/Objects	9b	OP	HTH		LIT																	
5	2	282	Well-cathedralwell act 1 inside				/Data/Global/Objects	zc	NU	HTH		LIT																	
5	2	283	shrine-shrine1_arcane sanctuary				/Data/Global/Objects	xx																					
5	2	284	shrine-dshrine2 act 2 shrine				/Data/Global/Objects	zs	OP	HTH		LIT							LIT										
5	2	285	shrine-desertshrine3 act 2 shrine				/Data/Global/Objects	zr	OP	HTH		LIT																	
5	2	286	shrine-dshrine1 act 2 shrine				/Data/Global/Objects	zd	OP	HTH		LIT																	
5	2	287	Well-desertwell act 2 well, desert, tomb				/Data/Global/Objects	zl	NU	HTH		LIT																	
5	2	288	Well-cavewell act 1 caves 				/Data/Global/Objects	zy	NU	HTH		LIT																	
5	2	289	chest-chest-r-large act 1				/Data/Global/Objects	q1	OP	HTH		LIT																	
5	2	290	chest-chest-r-tallskinney act 1				/Data/Global/Objects	q2	OP	HTH		LIT																	
5	2	291	chest-chest-r-med act 1				/Data/Global/Objects	q3	OP	HTH		LIT																	
5	2	292	jug-jug1 act 2, desert				/Data/Global/Objects	q4	OP	HTH		LIT																	
5	2	293	jug-jug2 act 2, desert				/Data/Global/Objects	q5	OP	HTH		LIT																	
5	2	294	chest-Lchest1 act 1				/Data/Global/Objects	q6	OP	HTH		LIT																	
5	2	295	Waypoint-waypointi inner hell				/Data/Global/Objects	wi	ON	HTH		LIT							LIT										
5	2	296	chest-dchest2R act 2, desert, tomb, chest-r-med				/Data/Global/Objects	q9	OP	HTH		LIT																	
5	2	297	chest-dchestr act 2, desert, tomb, chest -r large				/Data/Global/Objects	q7	OP	HTH		LIT																	
5	2	298	chest-dchestL act 2, desert, tomb chest l large				/Data/Global/Objects	q8	OP	HTH		LIT																	
5	2	299	taintedsunaltar-tainted sun altar quest				/Data/Global/Objects	za	OP	HTH		LIT							LIT										
5	2	300	shrine-dshrine1 act 2 , desert				/Data/Global/Objects	zv	NU	HTH		LIT							LIT	LIT									
5	2	301	shrine-dshrine4 act 2, desert				/Data/Global/Objects	ze	OP	HTH		LIT							LIT										
5	2	302	orifice-Where you place the Horadric staff				/Data/Global/Objects	HA	NU	HTH		LIT																	
5	2	303	Door-tyrael's door				/Data/Global/Objects	DX	OP	HTH		LIT																	
5	2	304	corpse-guard corpse				/Data/Global/Objects	GC	OP	HTH		LIT																	
5	2	305	hidden stash-rock act 1 wilderness				/Data/Global/Objects	c7	OP	HTH		LIT																	
5	2	306	Waypoint-waypoint act 2				/Data/Global/Objects	wm	ON	HTH		LIT							LIT										
5	2	307	Waypoint-waypoint act 1 wilderness				/Data/Global/Objects	wn	ON	HTH		LIT							LIT										
5	2	308	skeleton-corpse				/Data/Global/Objects	cp	OP	HTH		LIT																	
5	2	309	hidden stash-rockb act 1 wilderness				/Data/Global/Objects	cq	OP	HTH		LIT																	
5	2	310	fire-fire small				/Data/Global/Objects	FX	NU	HTH		LIT																	
5	2	311	fire-fire medium				/Data/Global/Objects	FY	NU	HTH		LIT																	
5	2	312	fire-fire large				/Data/Global/Objects	FZ	NU	HTH		LIT																	
5	2	313	hiding spot-cliff act 1 wilderness				/Data/Global/Objects	cf	NU	HTH		LIT																	
5	2	314	Shrine-mana well1				/Data/Global/Objects	MB	OP	HTH		LIT																	
5	2	315	Shrine-mana well2				/Data/Global/Objects	MD	OP	HTH		LIT																	
5	2	316	Shrine-mana well3, act 2, tomb, 				/Data/Global/Objects	MF	OP	HTH		LIT																	
5	2	317	Shrine-mana well4, act 2, harom				/Data/Global/Objects	MH	OP	HTH		LIT																	
5	2	318	Shrine-mana well5				/Data/Global/Objects	MJ	OP	HTH		LIT																	
5	2	319	hollow log-log				/Data/Global/Objects	cz	NU	HTH		LIT																	
5	2	320	Shrine-jungle healwell act 3				/Data/Global/Objects	JH	OP	HTH		LIT																	
5	2	321	skeleton-corpseb				/Data/Global/Objects	sx	OP	HTH		LIT																	
5	2	322	Shrine-health well, health shrine, desert				/Data/Global/Objects	Mk	OP	HTH		LIT																	
5	2	323	Shrine-mana well7, mana shrine, desert				/Data/Global/Objects	Mi	OP	HTH		LIT																	
5	2	324	loose rock-rockc act 1 wilderness				/Data/Global/Objects	RY	OP	HTH		LIT																	
5	2	325	loose boulder-rockd act 1 wilderness				/Data/Global/Objects	RZ	OP	HTH		LIT																	
5	2	326	chest-chest-L-med				/Data/Global/Objects	c8	OP	HTH		LIT																	
5	2	327	chest-chest-L-large				/Data/Global/Objects	c9	OP	HTH		LIT																	
5	2	328	GuardCorpse-guard on a stick, desert, tomb, harom				/Data/Global/Objects	GS	OP	HTH		LIT																	
5	2	329	bookshelf-bookshelf1				/Data/Global/Objects	b4	OP	HTH		LIT																	
5	2	330	bookshelf-bookshelf2				/Data/Global/Objects	b5	OP	HTH		LIT																	
5	2	331	chest-jungle chest act 3				/Data/Global/Objects	JC	OP	HTH		LIT																	
5	2	332	coffin-tombcoffin				/Data/Global/Objects	tm	OP	HTH		LIT																	
5	2	333	chest-chest-L-med, jungle				/Data/Global/Objects	jz	OP	HTH		LIT																	
5	2	334	Shrine-jungle shrine2				/Data/Global/Objects	jy	OP	HTH		LIT							LIT	LIT									
5	2	335	stash-jungle object act3				/Data/Global/Objects	jx	OP	HTH		LIT																	
5	2	336	stash-jungle object act3				/Data/Global/Objects	jw	OP	HTH		LIT																	
5	2	337	stash-jungle object act3				/Data/Global/Objects	jv	OP	HTH		LIT																	
5	2	338	stash-jungle object act3				/Data/Global/Objects	ju	OP	HTH		LIT																	
5	2	339	Dummy-cain portal				/Data/Global/Objects	tP	OP	HTH	LIT	LIT																	
5	2	340	Shrine-jungle shrine3 act 3				/Data/Global/Objects	js	OP	HTH		LIT							LIT										
5	2	341	Shrine-jungle shrine4 act 3				/Data/Global/Objects	jr	OP	HTH		LIT							LIT										
5	2	342	teleport pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
5	2	343	LamTome-Lam Esen's Tome				/Data/Global/Objects	ab	OP	HTH		LIT																	
5	2	344	stair-stairsl				/Data/Global/Objects	sl																					
5	2	345	stair-stairsr				/Data/Global/Objects	sv																					
5	2	346	a trap-test data floortrap				/Data/Global/Objects	a5	OP	HTH		LIT																	
5	2	347	Shrine-jungleshrine act 3				/Data/Global/Objects	jq	OP	HTH		LIT							LIT										
5	2	348	chest-chest-L-tallskinney, general chest r?				/Data/Global/Objects	c0	OP	HTH		LIT																	
5	2	349	Shrine-mafistoshrine				/Data/Global/Objects	mz	OP	HTH		LIT							LIT										
5	2	350	Shrine-mafistoshrine				/Data/Global/Objects	my	OP	HTH		LIT							LIT										
5	2	351	Shrine-mafistoshrine				/Data/Global/Objects	mx	NU	HTH		LIT							LIT										
5	2	352	Shrine-mafistomana				/Data/Global/Objects	mw	OP	HTH		LIT							LIT										
5	2	353	stash-mafistolair				/Data/Global/Objects	mv	OP	HTH		LIT																	
5	2	354	stash-box				/Data/Global/Objects	mu	OP	HTH		LIT																	
5	2	355	stash-altar				/Data/Global/Objects	mt	OP	HTH		LIT																	
5	2	356	Shrine-mafistohealth				/Data/Global/Objects	mr	OP	HTH		LIT							LIT										
5	2	357	dummy-water rocks in act 3 wrok				/Data/Global/Objects	rw	NU	HTH		LIT																	
5	2	358	Basket-basket 1				/Data/Global/Objects	bd	OP	HTH		LIT																	
5	2	359	Basket-basket 2				/Data/Global/Objects	bj	OP	HTH		LIT																	
5	2	360	Dummy-water logs in act 3  ne logw				/Data/Global/Objects	lw	NU	HTH		LIT																	
5	2	361	Dummy-water rocks girl in act 3 wrob				/Data/Global/Objects	wb	NU	HTH		LIT																	
5	2	362	Dummy-bubbles in act3 water				/Data/Global/Objects	yb	NU	HTH		LIT																	
5	2	363	Dummy-water logs in act 3 logx				/Data/Global/Objects	wd	NU	HTH		LIT																	
5	2	364	Dummy-water rocks in act 3 rokb				/Data/Global/Objects	wc	NU	HTH		LIT																	
5	2	365	Dummy-water rocks girl in act 3 watc				/Data/Global/Objects	we	NU	HTH		LIT																	
5	2	366	Dummy-water rocks in act 3 waty				/Data/Global/Objects	wy	NU	HTH		LIT																	
5	2	367	Dummy-water logs in act 3  logz				/Data/Global/Objects	lx	NU	HTH		LIT																	
5	2	368	Dummy-web covered tree 1				/Data/Global/Objects	w3	NU	HTH		LIT							LIT										
5	2	369	Dummy-web covered tree 2				/Data/Global/Objects	w4	NU	HTH		LIT							LIT										
5	2	370	Dummy-web covered tree 3				/Data/Global/Objects	w5	NU	HTH		LIT							LIT										
5	2	371	Dummy-web covered tree 4				/Data/Global/Objects	w6	NU	HTH		LIT							LIT										
5	2	372	pillar-hobject1				/Data/Global/Objects	70	OP	HTH		LIT																	
5	2	373	cocoon-cacoon				/Data/Global/Objects	CN	OP	HTH		LIT																	
5	2	374	cocoon-cacoon 2				/Data/Global/Objects	CC	OP	HTH		LIT																	
5	2	375	skullpile-hobject1				/Data/Global/Objects	ib	OP	HTH		LIT																	
5	2	376	Shrine-outershrinehell				/Data/Global/Objects	ia	OP	HTH		LIT							LIT										
5	2	377	dummy-water rock girl act 3  nw  blgb				/Data/Global/Objects	QX	NU	HTH		LIT																	
5	2	378	dummy-big log act 3  sw blga				/Data/Global/Objects	qw	NU	HTH		LIT																	
5	2	379	door-slimedoor1				/Data/Global/Objects	SQ	OP	HTH		LIT																	
5	2	380	door-slimedoor2				/Data/Global/Objects	SY	OP	HTH		LIT																	
5	2	381	Shrine-outershrinehell2				/Data/Global/Objects	ht	OP	HTH		LIT							LIT										
5	2	382	Shrine-outershrinehell3				/Data/Global/Objects	hq	OP	HTH		LIT																	
5	2	383	pillar-hobject2				/Data/Global/Objects	hv	OP	HTH		LIT																	
5	2	384	dummy-Big log act 3 se blgc 				/Data/Global/Objects	Qy	NU	HTH		LIT																	
5	2	385	dummy-Big log act 3 nw blgd				/Data/Global/Objects	Qz	NU	HTH		LIT																	
5	2	386	Shrine-health wellforhell				/Data/Global/Objects	ho	OP	HTH		LIT																	
5	2	387	Waypoint-act3waypoint town				/Data/Global/Objects	wz	ON	HTH		LIT							LIT										
5	2	388	Waypoint-waypointh				/Data/Global/Objects	wv	ON	HTH		LIT							LIT										
5	2	389	body-burning town				/Data/Global/Objects	bz	ON	HTH		LIT							LIT										
5	2	390	chest-gchest1L general				/Data/Global/Objects	cy	OP	HTH		LIT																	
5	2	391	chest-gchest2R general				/Data/Global/Objects	cx	OP	HTH		LIT																	
5	2	392	chest-gchest3R general				/Data/Global/Objects	cu	OP	HTH		LIT																	
5	2	393	chest-glchest3L general				/Data/Global/Objects	cd	OP	HTH		LIT																	
5	2	394	ratnest-sewers				/Data/Global/Objects	rn	OP	HTH		LIT																	
5	2	395	body-burning town				/Data/Global/Objects	by	NU	HTH		LIT							LIT										
5	2	396	ratnest-sewers				/Data/Global/Objects	ra	OP	HTH		LIT																	
5	2	397	bed-bed act 1				/Data/Global/Objects	qa	OP	HTH		LIT																	
5	2	398	bed-bed act 1				/Data/Global/Objects	qb	OP	HTH		LIT																	
5	2	399	manashrine-mana wellforhell				/Data/Global/Objects	hn	OP	HTH		LIT							LIT										
5	2	400	a trap-exploding cow  for Tristan and ACT 3 only??Very Rare  1 or 2				/Data/Global/Objects	ew	OP	HTH		LIT																	
5	2	401	gidbinn altar-gidbinn altar				/Data/Global/Objects	ga	ON	HTH		LIT							LIT										
5	2	402	gidbinn-gidbinn decoy				/Data/Global/Objects	gd	ON	HTH		LIT							LIT										
5	2	403	Dummy-diablo right light				/Data/Global/Objects	11	NU	HTH		LIT																	
5	2	404	Dummy-diablo left light				/Data/Global/Objects	12	NU	HTH		LIT																	
5	2	405	Dummy-diablo start point				/Data/Global/Objects	ss																					
5	2	406	Dummy-stool for act 1 cabin				/Data/Global/Objects	s9	NU	HTH		LIT																	
5	2	407	Dummy-wood for act 1 cabin				/Data/Global/Objects	wg	NU	HTH		LIT																	
5	2	408	Dummy-more wood for act 1 cabin				/Data/Global/Objects	wh	NU	HTH		LIT																	
5	2	409	Dummy-skeleton spawn for hell   facing nw				/Data/Global/Objects	QS	OP	HTH		LIT							LIT										
5	2	410	Shrine-holyshrine for monastery,catacombs,jail				/Data/Global/Objects	HL	OP	HTH		LIT							LIT										
5	2	411	a trap-spikes for tombs floortrap				/Data/Global/Objects	A7	OP	HTH		LIT																	
5	2	412	Shrine-act 1 cathedral				/Data/Global/Objects	s0	OP	HTH		LIT							LIT										
5	2	413	Shrine-act 1 jail				/Data/Global/Objects	jb	NU	HTH		LIT							LIT										
5	2	414	Shrine-act 1 jail				/Data/Global/Objects	jd	OP	HTH		LIT							LIT										
5	2	415	Shrine-act 1 jail				/Data/Global/Objects	jf	OP	HTH		LIT							LIT										
5	2	416	goo pile-goo pile for sand maggot lair				/Data/Global/Objects	GP	OP	HTH		LIT																	
5	2	417	bank-bank				/Data/Global/Objects	b6	NU	HTH		LIT																	
5	2	418	wirt's body-wirt's body				/Data/Global/Objects	BP	NU	HTH		LIT																	
5	2	419	dummy-gold placeholder				/Data/Global/Objects	1g																					
5	2	420	corpse-guard corpse 2				/Data/Global/Objects	GF	OP	HTH		LIT																	
5	2	421	corpse-dead villager 1				/Data/Global/Objects	dg	OP	HTH		LIT																	
5	2	422	corpse-dead villager 2				/Data/Global/Objects	df	OP	HTH		LIT																	
5	2	423	Dummy-yet another flame, no damage				/Data/Global/Objects	f8	NU	HTH		LIT																	
5	2	424	hidden stash-tiny pixel shaped thingie				/Data/Global/Objects	f9																					
5	2	425	Shrine-health shrine for caves				/Data/Global/Objects	ce	OP	HTH		LIT																	
5	2	426	Shrine-mana shrine for caves				/Data/Global/Objects	cg	OP	HTH		LIT																	
5	2	427	Shrine-cave magic shrine				/Data/Global/Objects	cg	OP	HTH		LIT																	
5	2	428	Shrine-manashrine, act 3, dungeun				/Data/Global/Objects	de	OP	HTH		LIT																	
5	2	429	Shrine-magic shrine, act 3 sewers.				/Data/Global/Objects	wj	NU	HTH		LIT							LIT	LIT									
5	2	430	Shrine-healthwell, act 3, sewers				/Data/Global/Objects	wk	OP	HTH		LIT																	
5	2	431	Shrine-manawell, act 3, sewers				/Data/Global/Objects	wl	OP	HTH		LIT																	
5	2	432	Shrine-magic shrine, act 3 sewers, dungeon.				/Data/Global/Objects	ws	NU	HTH		LIT							LIT	LIT									
5	2	433	dummy-brazier_celler, act 2				/Data/Global/Objects	bi	NU	HTH		LIT							LIT										
5	2	434	sarcophagus-anubis coffin, act2, tomb				/Data/Global/Objects	qc	OP	HTH		LIT																	
5	2	435	dummy-brazier_general, act 2, sewers, tomb, desert				/Data/Global/Objects	bm	NU	HTH		LIT							LIT										
5	2	436	Dummy-brazier_tall, act 2, desert, town, tombs				/Data/Global/Objects	bo	NU	HTH		LIT							LIT										
5	2	437	Dummy-brazier_small, act 2, desert, town, tombs				/Data/Global/Objects	bq	NU	HTH		LIT							LIT										
5	2	438	Waypoint-waypoint, celler				/Data/Global/Objects	w7	ON	HTH		LIT							LIT										
5	2	439	bed-bed for harum				/Data/Global/Objects	ub	OP	HTH		LIT																	
5	2	440	door-iron grate door left				/Data/Global/Objects	dv	NU	HTH		LIT																	
5	2	441	door-iron grate door right				/Data/Global/Objects	dn	NU	HTH		LIT																	
5	2	442	door-wooden grate door left				/Data/Global/Objects	dp	NU	HTH		LIT																	
5	2	443	door-wooden grate door right				/Data/Global/Objects	dt	NU	HTH		LIT																	
5	2	444	door-wooden door left				/Data/Global/Objects	dk	NU	HTH		LIT																	
5	2	445	door-wooden door right				/Data/Global/Objects	dl	NU	HTH		LIT																	
5	2	446	Dummy-wall torch left for tombs				/Data/Global/Objects	qd	NU	HTH		LIT							LIT										
5	2	447	Dummy-wall torch right for tombs				/Data/Global/Objects	qe	NU	HTH		LIT							LIT										
5	2	448	portal-arcane sanctuary portal				/Data/Global/Objects	ay	ON	HTH		LIT							LIT	LIT									
5	2	449	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hb	OP	HTH		LIT							LIT										
5	2	450	magic shrine-magic shrine, act 2, haram				/Data/Global/Objects	hc	OP	HTH		LIT							LIT										
5	2	451	Dummy-maggot well health				/Data/Global/Objects	qf	OP	HTH		LIT																	
5	2	452	manashrine-maggot well mana				/Data/Global/Objects	qg	OP	HTH		LIT																	
5	2	453	magic shrine-magic shrine, act 3 arcane sanctuary.				/Data/Global/Objects	hd	OP	HTH		LIT							LIT										
5	2	454	teleportation pad-teleportation pad				/Data/Global/Objects	7h	NU	HTH		LIT							LIT	LIT									
5	2	455	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
5	2	456	teleportation pad-teleportation pad				/Data/Global/Objects	aa	NU	HTH		LIT							LIT	LIT									
5	2	457	Dummy-arcane thing				/Data/Global/Objects	7a	NU	HTH		LIT																	
5	2	458	Dummy-arcane thing				/Data/Global/Objects	7b	NU	HTH		LIT																	
5	2	459	Dummy-arcane thing				/Data/Global/Objects	7c	NU	HTH		LIT																	
5	2	460	Dummy-arcane thing				/Data/Global/Objects	7d	NU	HTH		LIT																	
5	2	461	Dummy-arcane thing				/Data/Global/Objects	7e	NU	HTH		LIT																	
5	2	462	Dummy-arcane thing				/Data/Global/Objects	7f	NU	HTH		LIT																	
5	2	463	Dummy-arcane thing				/Data/Global/Objects	7g	NU	HTH		LIT																	
5	2	464	dead guard-harem guard 1				/Data/Global/Objects	qh	NU	HTH		LIT																	
5	2	465	dead guard-harem guard 2				/Data/Global/Objects	qi	NU	HTH		LIT																	
5	2	466	dead guard-harem guard 3				/Data/Global/Objects	qj	NU	HTH		LIT																	
5	2	467	dead guard-harem guard 4				/Data/Global/Objects	qk	NU	HTH		LIT																	
5	2	468	eunuch-harem blocker				/Data/Global/Objects	ss																					
5	2	469	Dummy-healthwell, act 2, arcane				/Data/Global/Objects	ax	OP	HTH		LIT																	
5	2	470	manashrine-healthwell, act 2, arcane				/Data/Global/Objects	au	OP	HTH		LIT																	
5	2	471	Dummy-test data				/Data/Global/Objects	pp	S1	HTH	LIT	LIT																	
5	2	472	Well-tombwell act 2 well, tomb				/Data/Global/Objects	hu	NU	HTH		LIT																	
5	2	473	Waypoint-waypoint act2 sewer				/Data/Global/Objects	qm	ON	HTH		LIT							LIT										
5	2	474	Waypoint-waypoint act3 travincal				/Data/Global/Objects	ql	ON	HTH		LIT							LIT										
5	2	475	magic shrine-magic shrine, act 3, sewer				/Data/Global/Objects	qn	NU	HTH		LIT							LIT										
5	2	476	dead body-act3, sewer				/Data/Global/Objects	qo	OP	HTH		LIT																	
5	2	477	dummy-torch (act 3 sewer) stra				/Data/Global/Objects	V1	NU	HTH		LIT							LIT										
5	2	478	dummy-torch (act 3 kurast) strb				/Data/Global/Objects	V2	NU	HTH		LIT							LIT										
5	2	479	chest-mafistochestlargeLeft				/Data/Global/Objects	xb	OP	HTH		LIT																	
5	2	480	chest-mafistochestlargeright				/Data/Global/Objects	xc	OP	HTH		LIT																	
5	2	481	chest-mafistochestmedleft				/Data/Global/Objects	xd	OP	HTH		LIT																	
5	2	482	chest-mafistochestmedright				/Data/Global/Objects	xe	OP	HTH		LIT																	
5	2	483	chest-spiderlairchestlargeLeft				/Data/Global/Objects	xf	OP	HTH		LIT																	
5	2	484	chest-spiderlairchesttallLeft				/Data/Global/Objects	xg	OP	HTH		LIT																	
5	2	485	chest-spiderlairchestmedright				/Data/Global/Objects	xh	OP	HTH		LIT																	
5	2	486	chest-spiderlairchesttallright				/Data/Global/Objects	xi	OP	HTH		LIT																	
5	2	487	Steeg Stone-steeg stone				/Data/Global/Objects	y6	NU	HTH		LIT							LIT										
5	2	488	Guild Vault-guild vault				/Data/Global/Objects	y4	NU	HTH		LIT																	
5	2	489	Trophy Case-trophy case				/Data/Global/Objects	y2	NU	HTH		LIT																	
5	2	490	Message Board-message board				/Data/Global/Objects	y3	NU	HTH		LIT																	
5	2	491	Dummy-mephisto bridge				/Data/Global/Objects	xj	OP	HTH		LIT																	
5	2	492	portal-hellgate				/Data/Global/Objects	1y	ON	HTH		LIT								LIT	LIT								
5	2	493	Shrine-manawell, act 3, kurast				/Data/Global/Objects	xl	OP	HTH		LIT																	
5	2	494	Shrine-healthwell, act 3, kurast				/Data/Global/Objects	xm	OP	HTH		LIT																	
5	2	495	Dummy-hellfire1				/Data/Global/Objects	e3	NU	HTH		LIT																	
5	2	496	Dummy-hellfire2				/Data/Global/Objects	e4	NU	HTH		LIT																	
5	2	497	Dummy-hellfire3				/Data/Global/Objects	e5	NU	HTH		LIT																	
5	2	498	Dummy-helllava1				/Data/Global/Objects	e6	NU	HTH		LIT																	
5	2	499	Dummy-helllava2				/Data/Global/Objects	e7	NU	HTH		LIT																	
5	2	500	Dummy-helllava3				/Data/Global/Objects	e8	NU	HTH		LIT																	
5	2	501	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
5	2	502	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
5	2	503	Dummy-helllightsource1				/Data/Global/Objects	ss		HTH		LIT																	
5	2	504	chest-horadric cube chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
5	2	505	chest-horadric scroll chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
5	2	506	chest-staff of kings chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
5	2	507	Tome-yet another tome				/Data/Global/Objects	TT	NU	HTH		LIT																	
5	2	508	fire-hell brazier				/Data/Global/Objects	E1	NU	HTH	LIT	LIT																	
5	2	509	fire-hell brazier				/Data/Global/Objects	E2	NU	HTH	LIT	LIT																	
5	2	510	RockPIle-dungeon				/Data/Global/Objects	xn	OP	HTH		LIT																	
5	2	511	magic shrine-magic shrine, act 3,dundeon				/Data/Global/Objects	qo	OP	HTH		LIT																	
5	2	512	basket-dungeon				/Data/Global/Objects	xp	OP	HTH		LIT																	
5	2	513	HungSkeleton-outerhell skeleton				/Data/Global/Objects	jw	OP	HTH		LIT																	
5	2	514	Dummy-guy for dungeon				/Data/Global/Objects	ea	OP	HTH		LIT																	
5	2	515	casket-casket for Act 3 dungeon				/Data/Global/Objects	vb	OP	HTH		LIT																	
5	2	516	sewer stairs-stairs for act 3 sewer quest				/Data/Global/Objects	ve	OP	HTH		LIT																	
5	2	517	sewer lever-lever for act 3 sewer quest				/Data/Global/Objects	vf	OP	HTH		LIT																	
5	2	518	darkwanderer-start position				/Data/Global/Objects	ss																					
5	2	519	dummy-trapped soul placeholder				/Data/Global/Objects	ss																					
5	2	520	Dummy-torch for act3 town				/Data/Global/Objects	VG	NU	HTH		LIT							LIT										
5	2	521	chest-LargeChestR				/Data/Global/Objects	L1	OP	HTH		LIT																	
5	2	522	BoneChest-innerhellbonepile				/Data/Global/Objects	y1	OP	HTH		LIT																	
5	2	523	Dummy-skeleton spawn for hell facing ne				/Data/Global/Objects	Qt	OP	HTH		LIT							LIT										
5	2	524	Dummy-fog act 3 water rfga				/Data/Global/Objects	ud	NU	HTH		LIT																	
5	2	525	Dummy-Not used				/Data/Global/Objects	xx																					
5	2	526	Hellforge-Forge  hell				/Data/Global/Objects	ux	ON	HTH		LIT							LIT	LIT	LIT								
5	2	527	Guild Portal-Portal to next guild level				/Data/Global/Objects	PP	NU	HTH	LIT	LIT																	
5	2	528	Dummy-hratli start				/Data/Global/Objects	ss																					
5	2	529	Dummy-hratli end				/Data/Global/Objects	ss																					
5	2	530	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	uy	OP	HTH		LIT							LIT										
5	2	531	TrappedSoul-Burning guy for outer hell				/Data/Global/Objects	15	OP	HTH		LIT							LIT										
5	2	532	Dummy-natalya start				/Data/Global/Objects	ss																					
5	2	533	TrappedSoul-guy stuck in hell				/Data/Global/Objects	18	OP	HTH		LIT																	
5	2	534	TrappedSoul-guy stuck in hell				/Data/Global/Objects	19	OP	HTH		LIT																	
5	2	535	Dummy-cain start position				/Data/Global/Objects	ss																					
5	2	536	Dummy-stairsr				/Data/Global/Objects	sv	OP	HTH		LIT																	
5	2	537	chest-arcanesanctuarybigchestLeft				/Data/Global/Objects	y7	OP	HTH		LIT																	
5	2	538	casket-arcanesanctuarycasket				/Data/Global/Objects	y8	OP	HTH		LIT																	
5	2	539	chest-arcanesanctuarybigchestRight				/Data/Global/Objects	y9	OP	HTH		LIT																	
5	2	540	chest-arcanesanctuarychestsmallLeft				/Data/Global/Objects	ya	OP	HTH		LIT																	
5	2	541	chest-arcanesanctuarychestsmallRight				/Data/Global/Objects	yc	OP	HTH		LIT																	
5	2	542	Seal-Diablo seal				/Data/Global/Objects	30	ON	HTH		LIT							LIT										
5	2	543	Seal-Diablo seal				/Data/Global/Objects	31	ON	HTH		LIT							LIT										
5	2	544	Seal-Diablo seal				/Data/Global/Objects	32	ON	HTH		LIT							LIT										
5	2	545	Seal-Diablo seal				/Data/Global/Objects	33	ON	HTH		LIT							LIT										
5	2	546	Seal-Diablo seal				/Data/Global/Objects	34	ON	HTH		LIT							LIT										
5	2	547	chest-sparklychest				/Data/Global/Objects	yf	OP	HTH		LIT																	
5	2	548	Waypoint-waypoint pandamonia fortress				/Data/Global/Objects	yg	ON	HTH		LIT							LIT										
5	2	549	fissure-fissure for act 4 inner hell				/Data/Global/Objects	fh	OP	HTH		LIT							LIT										
5	2	550	Dummy-brazier for act 4, hell mesa				/Data/Global/Objects	he	NU	HTH		LIT							LIT										
5	2	551	Dummy-smoke				/Data/Global/Objects	35	NU	HTH		LIT																	
5	2	552	Waypoint-waypoint valleywaypoint				/Data/Global/Objects	yi	ON	HTH		LIT							LIT										
5	2	553	fire-hell brazier				/Data/Global/Objects	9f	NU	HTH		LIT							LIT										
5	2	554	compellingorb-compelling orb				/Data/Global/Objects	55	NU	HTH		LIT							LIT	LIT									
5	2	555	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
5	2	556	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
5	2	557	chest-khalim chest				/Data/Global/Objects	xk	OP	HTH		LIT																	
5	2	558	Dummy-fortress brazier #1				/Data/Global/Objects	98	NU	HTH		LIT							LIT										
5	2	559	Dummy-fortress brazier #2				/Data/Global/Objects	99	NU	HTH		LIT							LIT										
5	2	560	Siege Control-To control siege machines				/Data/Global/Objects	zq	OP	HTH		LIT																	
5	2	561	ptox-Pot O Torch (level 1)				/Data/Global/Objects	px	NU	HTH		LIT							LIT	LIT									
5	2	562	pyox-fire pit  (level 1)				/Data/Global/Objects	py	NU	HTH		LIT							LIT										
5	2	563	chestR-expansion no snow				/Data/Global/Objects	6q	OP	HTH		LIT																	
5	2	564	Shrine3wilderness-expansion no snow				/Data/Global/Objects	6r	OP	HTH		LIT							LIT										
5	2	565	Shrine2wilderness-expansion no snow				/Data/Global/Objects	6s	NU	HTH		LIT							LIT										
5	2	566	hiddenstash-expansion no snow				/Data/Global/Objects	3w	OP	HTH		LIT																	
5	2	567	flag wilderness-expansion no snow				/Data/Global/Objects	ym	NU	HTH		LIT																	
5	2	568	barrel wilderness-expansion no snow				/Data/Global/Objects	yn	OP	HTH		LIT																	
5	2	569	barrel wilderness-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
5	2	570	woodchestL-expansion no snow				/Data/Global/Objects	yp	OP	HTH		LIT																	
5	2	571	Shrine3wilderness-expansion no snow				/Data/Global/Objects	yq	NU	HTH		LIT							LIT										
5	2	572	manashrine-expansion no snow				/Data/Global/Objects	yr	OP	HTH		LIT							LIT										
5	2	573	healthshrine-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
5	2	574	burialchestL-expansion no snow				/Data/Global/Objects	yt	OP	HTH		LIT																	
5	2	575	burialchestR-expansion no snow				/Data/Global/Objects	ys	OP	HTH		LIT							LIT										
5	2	576	well-expansion no snow				/Data/Global/Objects	yv	NU	HTH		LIT																	
5	2	577	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yw	OP	HTH		LIT							LIT	LIT									
5	2	578	Shrine2wilderness-expansion no snow				/Data/Global/Objects	yx	OP	HTH		LIT							LIT										
5	2	579	Waypoint-expansion no snow				/Data/Global/Objects	yy	ON	HTH		LIT							LIT										
5	2	580	ChestL-expansion no snow				/Data/Global/Objects	yz	OP	HTH		LIT																	
5	2	581	woodchestR-expansion no snow				/Data/Global/Objects	6a	OP	HTH		LIT																	
5	2	582	ChestSL-expansion no snow				/Data/Global/Objects	6b	OP	HTH		LIT																	
5	2	583	ChestSR-expansion no snow				/Data/Global/Objects	6c	OP	HTH		LIT																	
5	2	584	etorch1-expansion no snow				/Data/Global/Objects	6d	NU	HTH		LIT							LIT										
5	2	585	ecfra-camp fire				/Data/Global/Objects	2w	NU	HTH		LIT							LIT	LIT									
5	2	586	ettr-town torch				/Data/Global/Objects	2x	NU	HTH		LIT							LIT	LIT									
5	2	587	etorch2-expansion no snow				/Data/Global/Objects	6e	NU	HTH		LIT							LIT										
5	2	588	burningbodies-wilderness/siege				/Data/Global/Objects	6f	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
5	2	589	burningpit-wilderness/siege				/Data/Global/Objects	6g	NU	HTH		LIT							LIT	LIT	LIT	LIT	LIT	LIT					
5	2	590	tribal flag-wilderness/siege				/Data/Global/Objects	6h	NU	HTH		LIT																	
5	2	591	eflg-town flag				/Data/Global/Objects	2y	NU	HTH		LIT																	
5	2	592	chan-chandeleir				/Data/Global/Objects	2z	NU	HTH		LIT							LIT										
5	2	593	jar1-wilderness/siege				/Data/Global/Objects	6i	OP	HTH		LIT																	
5	2	594	jar2-wilderness/siege				/Data/Global/Objects	6j	OP	HTH		LIT																	
5	2	595	jar3-wilderness/siege				/Data/Global/Objects	6k	OP	HTH		LIT																	
5	2	596	swingingheads-wilderness				/Data/Global/Objects	6L	NU	HTH		LIT																	
5	2	597	pole-wilderness				/Data/Global/Objects	6m	NU	HTH		LIT																	
5	2	598	animated skulland rockpile-expansion no snow				/Data/Global/Objects	6n	OP	HTH		LIT																	
5	2	599	gate-town main gate				/Data/Global/Objects	2v	OP	HTH		LIT																	
5	2	600	pileofskullsandrocks-seige				/Data/Global/Objects	6o	NU	HTH		LIT																	
5	2	601	hellgate-seige				/Data/Global/Objects	6p	NU	HTH		LIT							LIT	LIT									
5	2	602	banner 1-preset in enemy camp				/Data/Global/Objects	ao	NU	HTH		LIT																	
5	2	603	banner 2-preset in enemy camp				/Data/Global/Objects	ap	NU	HTH		LIT																	
5	2	604	explodingchest-wilderness/siege				/Data/Global/Objects	6t	OP	HTH		LIT							LIT										
5	2	605	chest-specialchest				/Data/Global/Objects	6u	OP	HTH		LIT																	
5	2	606	deathpole-wilderness				/Data/Global/Objects	6v	NU	HTH		LIT																	
5	2	607	Ldeathpole-wilderness				/Data/Global/Objects	6w	NU	HTH		LIT																	
5	2	608	Altar-inside of temple				/Data/Global/Objects	6x	NU	HTH		LIT							LIT										
5	2	609	dummy-Drehya Start In Town				/Data/Global/Objects	ss																					
5	2	610	dummy-Drehya Start Outside Town				/Data/Global/Objects	ss																					
5	2	611	dummy-Nihlathak Start In Town				/Data/Global/Objects	ss																					
5	2	612	dummy-Nihlathak Start Outside Town				/Data/Global/Objects	ss																					
5	2	613	hidden stash-icecave_				/Data/Global/Objects	6y	OP	HTH		LIT																	
5	2	614	healthshrine-icecave_				/Data/Global/Objects	8a	OP	HTH		LIT																	
5	2	615	manashrine-icecave_				/Data/Global/Objects	8b	OP	HTH		LIT																	
5	2	616	evilurn-icecave_				/Data/Global/Objects	8c	OP	HTH		LIT																	
5	2	617	icecavejar1-icecave_				/Data/Global/Objects	8d	OP	HTH		LIT																	
5	2	618	icecavejar2-icecave_				/Data/Global/Objects	8e	OP	HTH		LIT																	
5	2	619	icecavejar3-icecave_				/Data/Global/Objects	8f	OP	HTH		LIT																	
5	2	620	icecavejar4-icecave_				/Data/Global/Objects	8g	OP	HTH		LIT																	
5	2	621	icecavejar4-icecave_				/Data/Global/Objects	8h	OP	HTH		LIT																	
5	2	622	icecaveshrine2-icecave_				/Data/Global/Objects	8i	NU	HTH		LIT							LIT										
5	2	623	cagedwussie1-caged fellow(A5-Prisonner)				/Data/Global/Objects	60	NU	HTH		LIT																	
5	2	624	Ancient Statue 3-statue				/Data/Global/Objects	60	NU	HTH		LIT																	
5	2	625	Ancient Statue 1-statue				/Data/Global/Objects	61	NU	HTH		LIT																	
5	2	626	Ancient Statue 2-statue				/Data/Global/Objects	62	NU	HTH		LIT																	
5	2	627	deadbarbarian-seige/wilderness				/Data/Global/Objects	8j	OP	HTH		LIT																	
5	2	628	clientsmoke-client smoke				/Data/Global/Objects	oz	NU	HTH		LIT																	
5	2	629	icecaveshrine2-icecave_				/Data/Global/Objects	8k	NU	HTH		LIT							LIT										
5	2	630	icecave_torch1-icecave_				/Data/Global/Objects	8L	NU	HTH		LIT							LIT										
5	2	631	icecave_torch2-icecave_				/Data/Global/Objects	8m	NU	HTH		LIT							LIT										
5	2	632	ttor-expansion tiki torch				/Data/Global/Objects	2p	NU	HTH		LIT							LIT										
5	2	633	manashrine-baals				/Data/Global/Objects	8n	OP	HTH		LIT																	
5	2	634	healthshrine-baals				/Data/Global/Objects	8o	OP	HTH		LIT																	
5	2	635	tomb1-baal's lair				/Data/Global/Objects	8p	OP	HTH		LIT																	
5	2	636	tomb2-baal's lair				/Data/Global/Objects	8q	OP	HTH		LIT																	
5	2	637	tomb3-baal's lair				/Data/Global/Objects	8r	OP	HTH		LIT																	
5	2	638	magic shrine-baal's lair				/Data/Global/Objects	8s	NU	HTH		LIT							LIT										
5	2	639	torch1-baal's lair				/Data/Global/Objects	8t	NU	HTH		LIT							LIT										
5	2	640	torch2-baal's lair				/Data/Global/Objects	8u	NU	HTH		LIT							LIT										
5	2	641	manashrine-snowy				/Data/Global/Objects	8v	OP	HTH		LIT							LIT										
5	2	642	healthshrine-snowy				/Data/Global/Objects	8w	OP	HTH		LIT							LIT										
5	2	643	well-snowy				/Data/Global/Objects	8x	NU	HTH		LIT																	
5	2	644	Waypoint-baals_waypoint				/Data/Global/Objects	8y	ON	HTH		LIT							LIT										
5	2	645	magic shrine-snowy_shrine3				/Data/Global/Objects	8z	NU	HTH		LIT							LIT										
5	2	646	Waypoint-wilderness_waypoint				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
5	2	647	magic shrine-snowy_shrine3				/Data/Global/Objects	5b	OP	HTH		LIT							LIT	LIT									
5	2	648	well-baalslair				/Data/Global/Objects	5c	NU	HTH		LIT																	
5	2	649	magic shrine2-baal's lair				/Data/Global/Objects	5d	NU	HTH		LIT							LIT										
5	2	650	object1-snowy				/Data/Global/Objects	5e	OP	HTH		LIT																	
5	2	651	woodchestL-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
5	2	652	woodchestR-snowy				/Data/Global/Objects	5g	OP	HTH		LIT																	
5	2	653	magic shrine-baals_shrine3				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
5	2	654	woodchest2L-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
5	2	655	woodchest2R-snowy				/Data/Global/Objects	5f	OP	HTH		LIT																	
5	2	656	swingingheads-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
5	2	657	debris-snowy				/Data/Global/Objects	5l	NU	HTH		LIT																	
5	2	658	pene-Pen breakable door				/Data/Global/Objects	2q	NU	HTH		LIT																	
5	2	659	magic shrine-temple				/Data/Global/Objects	5h	NU	HTH		LIT							LIT										
5	2	660	mrpole-snowy				/Data/Global/Objects	5k	NU	HTH		LIT																	
5	2	661	Waypoint-icecave 				/Data/Global/Objects	5a	ON	HTH		LIT							LIT										
5	2	662	magic shrine-temple				/Data/Global/Objects	5t	NU	HTH		LIT							LIT										
5	2	663	well-temple				/Data/Global/Objects	5q	NU	HTH		LIT																	
5	2	664	torch1-temple				/Data/Global/Objects	5r	NU	HTH		LIT							LIT										
5	2	665	torch1-temple				/Data/Global/Objects	5s	NU	HTH		LIT							LIT										
5	2	666	object1-temple				/Data/Global/Objects	5u	OP	HTH		LIT																	
5	2	667	object2-temple				/Data/Global/Objects	5v	OP	HTH		LIT																	
5	2	668	mrbox-baals				/Data/Global/Objects	5w	OP	HTH		LIT																	
5	2	669	well-icecave				/Data/Global/Objects	5x	NU	HTH		LIT																	
5	2	670	magic shrine-temple				/Data/Global/Objects	5y	NU	HTH		LIT							LIT										
5	2	671	healthshrine-temple				/Data/Global/Objects	5z	OP	HTH		LIT																	
5	2	672	manashrine-temple				/Data/Global/Objects	3a	OP	HTH		LIT																	
5	2	673	red light- (touch me)  for blacksmith				/Data/Global/Objects	ss																					
5	2	674	tomb1L-baal's lair				/Data/Global/Objects	3b	OP	HTH		LIT																	
5	2	675	tomb2L-baal's lair				/Data/Global/Objects	3c	OP	HTH		LIT																	
5	2	676	tomb3L-baal's lair				/Data/Global/Objects	3d	OP	HTH		LIT																	
5	2	677	ubub-Ice cave bubbles 01				/Data/Global/Objects	2u	NU	HTH		LIT																	
5	2	678	sbub-Ice cave bubbles 01				/Data/Global/Objects	2s	NU	HTH		LIT																	
5	2	679	tomb1-redbaal's lair				/Data/Global/Objects	3f	OP	HTH		LIT																	
5	2	680	tomb1L-redbaal's lair				/Data/Global/Objects	3g	OP	HTH		LIT																	
5	2	681	tomb2-redbaal's lair				/Data/Global/Objects	3h	OP	HTH		LIT																	
5	2	682	tomb2L-redbaal's lair				/Data/Global/Objects	3i	OP	HTH		LIT																	
5	2	683	tomb3-redbaal's lair				/Data/Global/Objects	3j	OP	HTH		LIT																	
5	2	684	tomb3L-redbaal's lair				/Data/Global/Objects	3k	OP	HTH		LIT																	
5	2	685	mrbox-redbaals				/Data/Global/Objects	3L	OP	HTH		LIT																	
5	2	686	torch1-redbaal's lair				/Data/Global/Objects	3m	NU	HTH		LIT							LIT										
5	2	687	torch2-redbaal's lair				/Data/Global/Objects	3n	NU	HTH		LIT							LIT										
5	2	688	candles-temple				/Data/Global/Objects	3o	NU	HTH		LIT							LIT										
5	2	689	Waypoint-temple				/Data/Global/Objects	3p	ON	HTH		LIT							LIT										
5	2	690	deadperson-everywhere				/Data/Global/Objects	3q	NU	HTH		LIT																	
5	2	691	groundtomb-temple				/Data/Global/Objects	3s	OP	HTH		LIT																	
5	2	692	Dummy-Larzuk Greeting				/Data/Global/Objects	ss																					
5	2	693	Dummy-Larzuk Standard				/Data/Global/Objects	ss																					
5	2	694	groundtombL-temple				/Data/Global/Objects	3t	OP	HTH		LIT																	
5	2	695	deadperson2-everywhere				/Data/Global/Objects	3u	OP	HTH		LIT																	
5	2	696	ancientsaltar-ancientsaltar				/Data/Global/Objects	4a	OP	HTH		LIT							LIT										
5	2	697	To The Worldstone Keep Level 1-ancientsdoor				/Data/Global/Objects	4b	OP	HTH		LIT																	
5	2	698	eweaponrackR-everywhere				/Data/Global/Objects	3x	NU	HTH		LIT																	
5	2	699	eweaponrackL-everywhere				/Data/Global/Objects	3y	NU	HTH		LIT																	
5	2	700	earmorstandR-everywhere				/Data/Global/Objects	3z	NU	HTH		LIT																	
5	2	701	earmorstandL-everywhere				/Data/Global/Objects	4c	NU	HTH		LIT																	
5	2	702	torch2-summit				/Data/Global/Objects	9g	NU	HTH		LIT							LIT										
5	2	703	funeralpire-outside				/Data/Global/Objects	9h	NU	HTH		LIT							LIT										
5	2	704	burninglogs-outside				/Data/Global/Objects	9i	NU	HTH		LIT							LIT										
5	2	705	stma-Ice cave steam				/Data/Global/Objects	2o	NU	HTH		LIT																	
5	2	706	deadperson2-everywhere				/Data/Global/Objects	3v	OP	HTH		LIT																	
5	2	707	Dummy-Baal's lair				/Data/Global/Objects	ss																					
5	2	708	fana-frozen anya				/Data/Global/Objects	2n	NU	HTH		LIT																	
5	2	709	BBQB-BBQ Bunny				/Data/Global/Objects	29	NU	HTH		LIT							LIT	LIT									
5	2	710	btor-Baal Torch Big				/Data/Global/Objects	25	NU	HTH		LIT							LIT										
5	2	711	Dummy-invisible ancient				/Data/Global/Objects	ss																					
5	2	712	Dummy-invisible base				/Data/Global/Objects	ss																					
5	2	713	The Worldstone Chamber-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
5	2	714	Glacial Caves Level 1-summit door				/Data/Global/Objects	4u	OP	HTH		LIT																	
5	2	715	strlastcinematic-last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
5	2	716	Harrogath-last last portal				/Data/Global/Objects	pp	NU	HTH	LIT	LIT																	
5	2	717	Zoo-test data				/Data/Global/Objects	ss																					
5	2	718	Keeper-test data				/Data/Global/Objects	7z	NU	HTH		LIT																	
5	2	719	Throne of Destruction-baals portal				/Data/Global/Objects	4x	ON	HTH		LIT							LIT										
5	2	720	Dummy-fire place guy				/Data/Global/Objects	7y	NU	HTH		LIT																	
5	2	721	Dummy-door blocker				/Data/Global/Objects	ss																					
5	2	722	Dummy-door blocker				/Data/Global/Objects	ss																					
`
