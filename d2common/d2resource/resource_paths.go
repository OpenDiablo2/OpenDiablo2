package d2resource

var LanguageCode string

const (
	// --- Screens ---

	LoadingScreen = "/data/global/ui/Loading/loadingscreen.dc6"

	// --- Main Menu ---

	TrademarkScreen       = "/data/global/ui/FrontEnd/trademarkscreenEXP.dc6"
	GameSelectScreen      = "/data/global/ui/FrontEnd/gameselectscreenEXP.dc6"
	Diablo2LogoFireLeft   = "/data/global/ui/FrontEnd/D2logoFireLeft.DC6"
	Diablo2LogoFireRight  = "/data/global/ui/FrontEnd/D2logoFireRight.DC6"
	Diablo2LogoBlackLeft  = "/data/global/ui/FrontEnd/D2logoBlackLeft.DC6"
	Diablo2LogoBlackRight = "/data/global/ui/FrontEnd/D2logoBlackRight.DC6"

	// --- Credits ---

	CreditsBackground = "/data/global/ui/CharSelect/creditsbckgexpand.dc6"
	CreditsText       = "/data/local/ui/{LANG}/ExpansionCredits.txt"

	// --- Character Select Screen ---

	CharacterSelectBackground = "/data/global/ui/FrontEnd/charactercreationscreenEXP.dc6"
	CharacterSelectCampfire   = "/data/global/ui/FrontEnd/fire.DC6"

	CharacterSelectBarbarianUnselected         = "/data/global/ui/FrontEnd/barbarian/banu1.DC6"
	CharacterSelectBarbarianUnselectedH        = "/data/global/ui/FrontEnd/barbarian/banu2.DC6"
	CharacterSelectBarbarianSelected           = "/data/global/ui/FrontEnd/barbarian/banu3.DC6"
	CharacterSelectBarbarianForwardWalk        = "/data/global/ui/FrontEnd/barbarian/bafw.DC6"
	CharacterSelectBarbarianForwardWalkOverlay = "/data/global/ui/FrontEnd/barbarian/BAFWs.DC6"
	CharacterSelectBarbarianBackWalk           = "/data/global/ui/FrontEnd/barbarian/babw.DC6"

	CharacterSelecSorceressUnselected         = "/data/global/ui/FrontEnd/sorceress/SONU1.DC6"
	CharacterSelecSorceressUnselectedH        = "/data/global/ui/FrontEnd/sorceress/SONU2.DC6"
	CharacterSelecSorceressSelected           = "/data/global/ui/FrontEnd/sorceress/SONU3.DC6"
	CharacterSelecSorceressSelectedOverlay    = "/data/global/ui/FrontEnd/sorceress/SONU3s.DC6"
	CharacterSelecSorceressForwardWalk        = "/data/global/ui/FrontEnd/sorceress/SOFW.DC6"
	CharacterSelecSorceressForwardWalkOverlay = "/data/global/ui/FrontEnd/sorceress/SOFWs.DC6"
	CharacterSelecSorceressBackWalk           = "/data/global/ui/FrontEnd/sorceress/SOBW.DC6"
	CharacterSelecSorceressBackWalkOverlay    = "/data/global/ui/FrontEnd/sorceress/SOBWs.DC6"

	CharacterSelectNecromancerUnselected        = "/data/global/ui/FrontEnd/necromancer/NENU1.DC6"
	CharacterSelectNecromancerUnselectedH       = "/data/global/ui/FrontEnd/necromancer/NENU2.DC6"
	CharacterSelecNecromancerSelected           = "/data/global/ui/FrontEnd/necromancer/NENU3.DC6"
	CharacterSelecNecromancerSelectedOverlay    = "/data/global/ui/FrontEnd/necromancer/NENU3s.DC6"
	CharacterSelecNecromancerForwardWalk        = "/data/global/ui/FrontEnd/necromancer/NEFW.DC6"
	CharacterSelecNecromancerForwardWalkOverlay = "/data/global/ui/FrontEnd/necromancer/NEFWs.DC6"
	CharacterSelecNecromancerBackWalk           = "/data/global/ui/FrontEnd/necromancer/NEBW.DC6"
	CharacterSelecNecromancerBackWalkOverlay    = "/data/global/ui/FrontEnd/necromancer/NEBWs.DC6"

	CharacterSelectPaladinUnselected        = "/data/global/ui/FrontEnd/paladin/PANU1.DC6"
	CharacterSelectPaladinUnselectedH       = "/data/global/ui/FrontEnd/paladin/PANU2.DC6"
	CharacterSelecPaladinSelected           = "/data/global/ui/FrontEnd/paladin/PANU3.DC6"
	CharacterSelecPaladinForwardWalk        = "/data/global/ui/FrontEnd/paladin/PAFW.DC6"
	CharacterSelecPaladinForwardWalkOverlay = "/data/global/ui/FrontEnd/paladin/PAFWs.DC6"
	CharacterSelecPaladinBackWalk           = "/data/global/ui/FrontEnd/paladin/PABW.DC6"

	CharacterSelectAmazonUnselected        = "/data/global/ui/FrontEnd/amazon/AMNU1.DC6"
	CharacterSelectAmazonUnselectedH       = "/data/global/ui/FrontEnd/amazon/AMNU2.DC6"
	CharacterSelecAmazonSelected           = "/data/global/ui/FrontEnd/amazon/AMNU3.DC6"
	CharacterSelecAmazonForwardWalk        = "/data/global/ui/FrontEnd/amazon/AMFW.DC6"
	CharacterSelecAmazonForwardWalkOverlay = "/data/global/ui/FrontEnd/amazon/AMFWs.DC6"
	CharacterSelecAmazonBackWalk           = "/data/global/ui/FrontEnd/amazon/AMBW.DC6"

	CharacterSelectAssassinUnselected  = "/data/global/ui/FrontEnd/assassin/ASNU1.DC6"
	CharacterSelectAssassinUnselectedH = "/data/global/ui/FrontEnd/assassin/ASNU2.DC6"
	CharacterSelectAssassinSelected    = "/data/global/ui/FrontEnd/assassin/ASNU3.DC6"
	CharacterSelectAssassinForwardWalk = "/data/global/ui/FrontEnd/assassin/ASFW.DC6"
	CharacterSelectAssassinBackWalk    = "/data/global/ui/FrontEnd/assassin/ASBW.DC6"

	CharacterSelectDruidUnselected  = "/data/global/ui/FrontEnd/druid/DZNU1.dc6"
	CharacterSelectDruidUnselectedH = "/data/global/ui/FrontEnd/druid/DZNU2.dc6"
	CharacterSelectDruidSelected    = "/data/global/ui/FrontEnd/druid/DZNU3.DC6"
	CharacterSelectDruidForwardWalk = "/data/global/ui/FrontEnd/druid/DZFW.DC6"
	CharacterSelectDruidBackWalk    = "/data/global/ui/FrontEnd/druid/DZBW.DC6"

	// -- Character Selection

	CharacterSelectionBackground = "/data/global/ui/CharSelect/characterselectscreenEXP.dc6"
	CharacterSelectionSelectBox  = "/data/global/ui/CharSelect/charselectbox.dc6"
	PopUpOkCancel                = "/data/global/ui/FrontEnd/PopUpOKCancel.dc6"

	// --- Game ---

	GamePanels          = "/data/global/ui/PANEL/800ctrlpnl7.dc6"
	GameGlobeOverlap    = "/data/global/ui/PANEL/overlap.DC6"
	HealthMana          = "/data/global/ui/PANEL/hlthmana.DC6"
	GameSmallMenuButton = "/data/global/ui/PANEL/menubutton.DC6" // TODO: Used for inventory popout
	SkillIcon           = "/data/global/ui/PANEL/Skillicon.DC6"  // TODO: Used for skill icon button
	AddSkillButton      = "/data/global/ui/PANEL/level.DC6"

	// --- Mouse Pointers ---

	CursorDefault = "/data/global/ui/CURSOR/ohand.DC6"

	// --- Fonts ---

	Font6          = "/data/local/font/{LANG_FONT}/font6"
	Font8          = "/data/local/font/{LANG_FONT}/font8"
	Font16         = "/data/local/font/{LANG_FONT}/font16"
	Font24         = "/data/local/font/{LANG_FONT}/font24"
	Font30         = "/data/local/font/{LANG_FONT}/font30"
	Font42         = "/data/local/font/{LANG_FONT}/font42"
	FontFormal12   = "/data/local/font/{LANG_FONT}/fontformal12"
	FontFormal11   = "/data/local/font/{LANG_FONT}/fontformal11"
	FontFormal10   = "/data/local/font/{LANG_FONT}/fontformal10"
	FontExocet10   = "/data/local/font/{LANG_FONT}/fontexocet10"
	FontExocet8    = "/data/local/font/{LANG_FONT}/fontexocet8"
	FontSucker     = "/data/local/font/{LANG_FONT}/ReallyTheLastSucker"
	FontRediculous = "/data/local/font/{LANG_FONT}/fontridiculous"

	// --- UI ---

	WideButtonBlank   = "/data/global/ui/FrontEnd/WideButtonBlank.dc6"
	MediumButtonBlank = "/data/global/ui/FrontEnd/MediumButtonBlank.dc6"
	CancelButton      = "/data/global/ui/FrontEnd/CancelButtonBlank.dc6"
	NarrowButtonBlank = "/data/global/ui/FrontEnd/NarrowButtonBlank.dc6"
	ShortButtonBlank  = "/data/global/ui/CharSelect/ShortButtonBlank.dc6"
	TextBox2          = "/data/global/ui/FrontEnd/textbox2.dc6"
	TallButtonBlank   = "/data/global/ui/CharSelect/TallButtonBlank.dc6"
	Checkbox          = "/data/global/ui/FrontEnd/clickbox.dc6"
	Scrollbar         = "/data/global/ui/PANEL/scrollbar.dc6"

	// --- GAME UI ---

	PentSpin        = "/data/global/ui/CURSOR/pentspin.DC6"
	MinipanelSmall  = "/data/global/ui/PANEL/minipanel_s.dc6"
	MinipanelButton = "/data/global/ui/PANEL/minipanelbtn.DC6"

	Frame                   = "/data/global/ui/PANEL/800borderframe.dc6"
	InventoryCharacterPanel = "/data/global/ui/PANEL/invchar6.DC6"
	InventoryWeaponsTab     = "/data/global/ui/PANEL/invchar6Tab.DC6"
	SkillsPanelAmazon       = "/data/global/ui/SPELLS/skltree_a_back.DC6"
	SkillsPanelBarbarian    = "/data/global/ui/SPELLS/skltree_b_back.DC6"
	SkillsPanelDruid        = "/data/global/ui/SPELLS/skltree_d_back.DC6"
	SkillsPanelAssassin     = "/data/global/ui/SPELLS/skltree_i_back.DC6"
	SkillsPanelNecromancer  = "/data/global/ui/SPELLS/skltree_n_back.DC6"
	SkillsPanelPaladin      = "/data/global/ui/SPELLS/skltree_p_back.DC6"
	SkillsPanelSorcerer     = "/data/global/ui/SPELLS/skltree_s_back.DC6"

	GenericSkills     = "/data/global/ui/SPELLS/Skillicon.DC6"
	AmazonSkills      = "/data/global/ui/SPELLS/AmSkillicon.DC6"
	BarbarianSkills   = "/data/global/ui/SPELLS/BaSkillicon.DC6"
	DruidSkills       = "/data/global/ui/SPELLS/DrSkillicon.DC6"
	AssassinSkills    = "/data/global/ui/SPELLS/AsSkillicon.DC6"
	NecromancerSkills = "/data/global/ui/SPELLS/NeSkillicon.DC6"
	PaladinSkills     = "/data/global/ui/SPELLS/PaSkillicon.DC6"
	SorcererSkills    = "/data/global/ui/SPELLS/SoSkillicon.DC6"

	RunButton      = "/data/global/ui/PANEL/runbutton.dc6"
	MenuButton     = "/data/global/ui/PANEL/menubutton.DC6"
	GoldCoinButton = "/data/global/ui/panel/goldcoinbtn.dc6"
	SquareButton   = "/data/global/ui/panel/buysellbtn.dc6"

	ArmorPlaceholder      = "/data/global/ui/PANEL/inv_armor.DC6"
	BeltPlaceholder       = "/data/global/ui/PANEL/inv_belt.DC6"
	BootsPlaceholder      = "/data/global/ui/PANEL/inv_boots.DC6"
	HelmGlovePlaceholder  = "/data/global/ui/PANEL/inv_helm_glove.DC6"
	RingAmuletPlaceholder = "/data/global/ui/PANEL/inv_ring_amulet.DC6"
	WeaponsPlaceholder    = "/data/global/ui/PANEL/inv_weapons.DC6"

	// --- Data ---

	ExpansionStringTable = "/data/local/lng/{LANG}/expansionstring.tbl"
	StringTable          = "/data/local/lng/{LANG}/string.tbl"
	PatchStringTable     = "/data/local/lng/{LANG}/patchstring.tbl"
	LevelPreset          = "/data/global/excel/LvlPrest.txt"
	LevelType            = "/data/global/excel/LvlTypes.txt"
	ObjectType           = "/data/global/excel/objtype.bin"
	LevelWarp            = "/data/global/excel/LvlWarp.bin"
	LevelDetails         = "/data/global/excel/Levels.bin"
	ObjectDetails        = "/data/global/excel/Objects.txt"
	SoundSettings        = "/data/global/excel/Sounds.txt"

	// --- Animations ---

	ObjectData          = "/data/global/objects"
	AnimationData       = "/data/global/animdata.d2"
	PlayerAnimationBase = "/data/global/CHARS"
	MissileData         = "/data/global/missiles"

	// --- Inventory Data ---

	Weapons     = "/data/global/excel/weapons.txt"
	Armor       = "/data/global/excel/armor.txt"
	Misc        = "/data/global/excel/misc.txt"
	UniqueItems = "/data/global/excel/UniqueItems.txt"

	// --- Character Data ---

	Experience = "/data/global/excel/experience.txt"
	CharStats  = "/data/global/excel/charstats.txt"

	// --- Music ---

	BGMTitle                    = "/data/global/music/introedit.wav"
	BGMOptions                  = "/data/global/music/Common/options.wav"
	BGMAct1AndarielAction       = "/data/global/music/Act1/andarielaction.wav"
	BGMAct1BloodRavenResolution = "/data/global/music/Act1/bloodravenresolution.wav"
	BGMAct1Caves                = "/data/global/music/Act1/caves.wav"
	BGMAct1Crypt                = "/data/global/music/Act1/crypt.wav"
	BGMAct1DenOfEvilAction      = "/data/global/music/Act1/denofevilaction.wav"
	BGMAct1Monastery            = "/data/global/music/Act1/monastery.wav"
	BGMAct1Town1                = "/data/global/music/Act1/town1.wav"
	BGMAct1Tristram             = "/data/global/music/Act1/tristram.wav"
	BGMAct1Wild                 = "/data/global/music/Act1/wild.wav"
	BGMAct2Desert               = "/data/global/music/Act2/desert.wav"
	BGMAct2Harem                = "/data/global/music/Act2/harem.wav"
	BGMAct2HoradricAction       = "/data/global/music/Act2/horadricaction.wav"
	BGMAct2Lair                 = "/data/global/music/Act2/lair.wav"
	BGMAct2RadamentResolution   = "/data/global/music/Act2/radamentresolution.wav"
	BGMAct2Sanctuary            = "/data/global/music/Act2/sanctuary.wav"
	BGMAct2Sewer                = "/data/global/music/Act2/sewer.wav"
	BGMAct2TaintedSunAction     = "/data/global/music/Act2/taintedsunaction.wav"
	BGMAct2Tombs                = "/data/global/music/Act2/tombs.wav"
	BGMAct2Town2                = "/data/global/music/Act2/town2.wav"
	BGMAct2Valley               = "/data/global/music/Act2/valley.wav"
	BGMAct3Jungle               = "/data/global/music/Act3/jungle.wav"
	BGMAct3Kurast               = "/data/global/music/Act3/kurast.wav"
	BGMAct3KurastSewer          = "/data/global/music/Act3/kurastsewer.wav"
	BGMAct3MefDeathAction       = "/data/global/music/Act3/mefdeathaction.wav"
	BGMAct3OrbAction            = "/data/global/music/Act3/orbaction.wav"
	BGMAct3Spider               = "/data/global/music/Act3/spider.wav"
	BGMAct3Town3                = "/data/global/music/Act3/town3.wav"
	BGMAct4Diablo               = "/data/global/music/Act4/diablo.wav"
	BGMAct4DiabloAction         = "/data/global/music/Act4/diabloaction.wav"
	BGMAct4ForgeAction          = "/data/global/music/Act4/forgeaction.wav"
	BGMAct4IzualAction          = "/data/global/music/Act4/izualaction.wav"
	BGMAct4Mesa                 = "/data/global/music/Act4/mesa.wav"
	BGMAct4Town4                = "/data/global/music/Act4/town4.wav"
	BGMAct5Baal                 = "/data/global/music/Act5/baal.wav"
	BGMAct5XTown                = "/data/global/music/Act5/xtown.wav"

	// --- Sound Effects ---

	SFXButtonClick         = "cursor_button_click"
	SFXAmazonDeselect      = "cursor_amazon_deselect"
	SFXAmazonSelect        = "cursor_amazon_select"
	SFXAssassinDeselect    = "/data/global/sfx/Cursor/intro/assassin deselect.wav"
	SFXAssassinSelect      = "/data/global/sfx/Cursor/intro/assassin select.wav"
	SFXBarbarianDeselect   = "cursor_barbarian_deselect"
	SFXBarbarianSelect     = "cursor_barbarian_select"
	SFXDruidDeselect       = "/data/global/sfx/Cursor/intro/druid deselect.wav"
	SFXDruidSelect         = "/data/global/sfx/Cursor/intro/druid select.wav"
	SFXNecromancerDeselect = "cursor_necromancer_deselect"
	SFXNecromancerSelect   = "cursor_necromancer_select"
	SFXPaladinDeselect     = "cursor_paladin_deselect"
	SFXPaladinSelect       = "cursor_paladin_select"
	SFXSorceressDeselect   = "cursor_sorceress_deselect"
	SFXSorceressSelect     = "cursor_sorceress_select"

	// --- Enemy Data ---

	MonStats = "/data/global/excel/monstats.txt"

	// --- Skill Data ---

	Missiles = "/data/global/excel/Missiles.txt"

	// --- Palettes ---

	PaletteAct1      = "/data/global/palette/act1/pal.dat"
	PaletteAct2      = "/data/global/palette/act2/pal.dat"
	PaletteAct3      = "/data/global/palette/act3/pal.dat"
	PaletteAct4      = "/data/global/palette/act4/pal.dat"
	PaletteAct5      = "/data/global/palette/act5/pal.dat"
	PaletteEndGame   = "/data/global/palette/endgame/pal.dat"
	PaletteEndGame2  = "/data/global/palette/endgame2/pal.dat"
	PaletteFechar    = "/data/global/palette/fechar/pal.dat"
	PaletteLoading   = "/data/global/palette/loading/pal.dat"
	PaletteMenu0     = "/data/global/palette/menu0/pal.dat"
	PaletteMenu1     = "/data/global/palette/menu1/pal.dat"
	PaletteMenu2     = "/data/global/palette/menu2/pal.dat"
	PaletteMenu3     = "/data/global/palette/menu3/pal.dat"
	PaletteMenu4     = "/data/global/palette/menu4/pal.dat"
	PaletteSky       = "/data/global/palette/sky/pal.dat"
	PaletteStatic    = "/data/global/palette/static/pal.dat"
	PaletteTrademark = "/data/global/palette/trademark/pal.dat"
	PaletteUnits     = "/data/global/palette/units/pal.dat"

	// --- Palette Transforms ---

	PaletteTransformAct1      = "/data/global/palette/act1/Pal.pl2"
	PaletteTransformAct2      = "/data/global/palette/act2/Pal.pl2"
	PaletteTransformAct3      = "/data/global/palette/act3/Pal.pl2"
	PaletteTransformAct4      = "/data/global/palette/act4/Pal.pl2"
	PaletteTransformAct5      = "/data/global/palette/act5/Pal.pl2"
	PaletteTransformEndGame   = "/data/global/palette/endgame/Pal.pl2"
	PaletteTransformEndGame2  = "/data/global/palette/endgame2/Pal.pl2"
	PaletteTransformFechar    = "/data/global/palette/fechar/Pal.pl2"
	PaletteTransformLoading   = "/data/global/palette/loading/Pal.pl2"
	PaletteTransformMenu0     = "/data/global/palette/menu0/Pal.pl2"
	PaletteTransformMenu1     = "/data/global/palette/menu1/Pal.pl2"
	PaletteTransformMenu2     = "/data/global/palette/menu2/Pal.pl2"
	PaletteTransformMenu3     = "/data/global/palette/menu3/Pal.pl2"
	PaletteTransformMenu4     = "/data/global/palette/menu4/Pal.pl2"
	PaletteTransformSky       = "/data/global/palette/sky/Pal.pl2"
	PaletteTransformTrademark = "/data/global/palette/trademark/Pal.pl2"
)
