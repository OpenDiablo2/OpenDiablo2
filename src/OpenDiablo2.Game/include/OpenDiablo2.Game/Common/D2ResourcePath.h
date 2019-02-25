#ifndef OPENDIABLO2_GAME_COMMON_D2RESOURCEPATH_H
#define OPENDIABLO2_GAME_COMMON_D2RESOURCEPATH_H

#include <string>

namespace OpenDiablo2::Game::Common {

class D2ResourcePath {
public:
// --- Loading Screen ---
        static const std::string LoadingScreen;

        // --- Main Menu ---
        static const std::string GameSelectScreen;
        static const std::string Diablo2LogoFireLeft;
        static const std::string Diablo2LogoFireRight;
        static const std::string Diablo2LogoBlackLeft;
        static const std::string Diablo2LogoBlackRight;

        // --- Credits ---
        static const std::string CreditsBackground;
        static const std::string CreditsText;

        // --- Character Select Screen ---
        static const std::string CharacterSelectBackground;
        static const std::string CharacterSelectCampfire;

        static const std::string CharacterSelectBarbarianUnselected;
        static const std::string CharacterSelectBarbarianUnselectedH;
        static const std::string CharacterSelectBarbarianSelected;
        static const std::string CharacterSelectBarbarianForwardWalk;
        static const std::string CharacterSelectBarbarianForwardWalkOverlay;
        static const std::string CharacterSelectBarbarianBackWalk;

        static const std::string CharacterSelecSorceressUnselected;
        static const std::string CharacterSelecSorceressUnselectedH;
        static const std::string CharacterSelecSorceressSelected;
        static const std::string CharacterSelecSorceressSelectedOverlay;
        static const std::string CharacterSelecSorceressForwardWalk;
        static const std::string CharacterSelecSorceressForwardWalkOverlay;
        static const std::string CharacterSelecSorceressBackWalk;
        static const std::string CharacterSelecSorceressBackWalkOverlay;

        static const std::string CharacterSelectNecromancerUnselected;
        static const std::string CharacterSelectNecromancerUnselectedH;
        static const std::string CharacterSelecNecromancerSelected;
        static const std::string CharacterSelecNecromancerSelectedOverlay;
        static const std::string CharacterSelecNecromancerForwardWalk;
        static const std::string CharacterSelecNecromancerForwardWalkOverlay;
        static const std::string CharacterSelecNecromancerBackWalk;
        static const std::string CharacterSelecNecromancerBackWalkOverlay;

        static const std::string CharacterSelectPaladinUnselected;
        static const std::string CharacterSelectPaladinUnselectedH;
        static const std::string CharacterSelecPaladinSelected;
        static const std::string CharacterSelecPaladinForwardWalk;
        static const std::string CharacterSelecPaladinForwardWalkOverlay;
        static const std::string CharacterSelecPaladinBackWalk;


        static const std::string CharacterSelectAmazonUnselected;
        static const std::string CharacterSelectAmazonUnselectedH;
        static const std::string CharacterSelecAmazonSelected;
        static const std::string CharacterSelecAmazonForwardWalk;
        static const std::string CharacterSelecAmazonForwardWalkOverlay;
        static const std::string CharacterSelecAmazonBackWalk;

        static const std::string CharacterSelectAssassinUnselected;
        static const std::string CharacterSelectAssassinUnselectedH;
        static const std::string CharacterSelectAssassinSelected;
        static const std::string CharacterSelectAssassinForwardWalk;
        static const std::string CharacterSelectAssassinBackWalk;

        static const std::string CharacterSelectDruidUnselected;
        static const std::string CharacterSelectDruidUnselectedH;
        static const std::string CharacterSelectDruidSelected;
        static const std::string CharacterSelectDruidForwardWalk;
        static const std::string CharacterSelectDruidBackWalk;

        // -- Character Selection
        static const std::string CharacterSelectionBackground;

        // --- Game ---
        static const std::string GamePanels;
        static const std::string GameGlobeOverlap;
        static const std::string HealthMana;
        static const std::string GameSmallMenuButton;
        static const std::string SkillIcon;
        static const std::string AddSkillButton;

        // --- Mouse Pointers ---
        static const std::string CursorDefault;

        // --- Fonts ---
        static const std::string Font6;
        static const std::string Font8;
        static const std::string Font16;
        static const std::string Font24;
        static const std::string Font30;
        static const std::string FontFormal12;
        static const std::string FontFormal11;
        static const std::string FontFormal10;
        static const std::string FontExocet10;
        static const std::string FontExocet8;

        // --- UI ---
        static const std::string WideButtonBlank;
        static const std::string MediumButtonBlank;
        static const std::string CancelButton;
        static const std::string NarrowButtonBlank;
        static const std::string ShortButtonBlank;
        static const std::string TextBox2;
        static const std::string TallButtonBlank;

        // --- GAME UI ---
        static const std::string MinipanelSmall;
        static const std::string MinipanelButton;

        static const std::string Frame;
        static const std::string InventoryCharacterPanel;
        static const std::string InventoryWeaponsTab;
        static const std::string SkillsPanelAmazon;
        static const std::string SkillsPanelBarbarian;
        static const std::string SkillsPanelDruid;
        static const std::string SkillsPanelAssassin;
        static const std::string SkillsPanelNecromancer;
        static const std::string SkillsPanelPaladin;
        static const std::string SkillsPanelSorcerer;

        static const std::string GenericSkills;
        static const std::string AmazonSkills;
        static const std::string BarbarianSkills;
        static const std::string DruidSkills;
        static const std::string AssassinSkills;
        static const std::string NecromancerSkills;
        static const std::string PaladinSkills;
        static const std::string SorcererSkills;

        static const std::string RunButton;
        static const std::string MenuButton;
        static const std::string GoldCoinButton;
        static const std::string SquareButton;

        static const std::string ArmorPlaceholder;
        static const std::string BeltPlaceholder;
        static const std::string BootsPlaceholder;
        static const std::string HelmGlovePlaceholder;
        static const std::string RingAmuletPlaceholder;
        static const std::string WeaponsPlaceholder;

        // --- Data ---
        static const std::string EnglishTable;
        static const std::string ExpansionStringTable;
        static const std::string LevelPreset;
        static const std::string LevelType;
        static const std::string LevelDetails;
        static const std::string ObjectDetails;

        // --- Animations ---
        static const std::string ObjectData;
        static const std::string AnimationData;
        static const std::string PlayerAnimationBase;

        // --- Inventory Data ---
        static const std::string Weapons;
        static const std::string Armor;
        static const std::string Misc;

        // --- Character Data ---
        static const std::string Experience;
        static const std::string CharStats;

        // --- Music ---
        static const std::string BGMTitle;
        static const std::string BGMOptions;
        static const std::string BGMAct1AndarielAction;
        static const std::string BGMAct1BloodRavenResolution;
        static const std::string BGMAct1Caves;
        static const std::string BGMAct1Crypt;
        static const std::string BGMAct1DenOfEvilAction;
        static const std::string BGMAct1Monastery;
        static const std::string BGMAct1Town1;
        static const std::string BGMAct1Tristram;
        static const std::string BGMAct1Wild;
        static const std::string BGMAct2Desert;
        static const std::string BGMAct2Harem;
        static const std::string BGMAct2HoradricAction;
        static const std::string BGMAct2Lair;
        static const std::string BGMAct2RadamentResolution;
        static const std::string BGMAct2Sanctuary;
        static const std::string BGMAct2Sewer;
        static const std::string BGMAct2TaintedSunAction;
        static const std::string BGMAct2Tombs;
        static const std::string BGMAct2Town2;
        static const std::string BGMAct2Valley;
        static const std::string BGMAct3Jungle;
        static const std::string BGMAct3Kurast;
        static const std::string BGMAct3KurastSewer;
        static const std::string BGMAct3MefDeathAction;
        static const std::string BGMAct3OrbAction;
        static const std::string BGMAct3Spider;
        static const std::string BGMAct3Town3;
        static const std::string BGMAct4Diablo;
        static const std::string BGMAct4DiabloAction;
        static const std::string BGMAct4ForgeAction;
        static const std::string BGMAct4IzualAction;
        static const std::string BGMAct4Mesa;
        static const std::string BGMAct4Town4;
        static const std::string BGMAct5Baal;
        static const std::string BGMAct5XTown;


        // --- Sound Effects ---
        static const std::string SFXButtonClick;
        static const std::string SFXAmazonDeselect;
        static const std::string SFXAmazonSelect;
        static const std::string SFXAssassinDeselect;
        static const std::string SFXAssassinSelect;
        static const std::string SFXBarbarianDeselect;
        static const std::string SFXBarbarianSelect;
        static const std::string SFXDruidDeselect;
        static const std::string SFXDruidSelect;
        static const std::string SFXNecromancerDeselect;
        static const std::string SFXNecromancerSelect;
        static const std::string SFXPaladinDeselect;
        static const std::string SFXPaladinSelect;
        static const std::string SFXSorceressDeselect;
        static const std::string SFXSorceressSelect;

        // --- Enemy Data ---
        static const std::string MonStats;

        // --- Skill Data ---
        static const std::string Missiles;
private:
	D2ResourcePath() { }
};

// --- Loading Screen ---
const std::string D2ResourcePath::LoadingScreen  = "data\\global\\ui\\Loading\\loadingscreen.dc6";

// --- Main Menu ---
const std::string D2ResourcePath::GameSelectScreen = "data\\global\\ui\\FrontEnd\\gameselectscreenEXP.dc6";
const std::string D2ResourcePath::Diablo2LogoFireLeft = "data\\global\\ui\\FrontEnd\\D2logoFireLeft.DC6";
const std::string D2ResourcePath::Diablo2LogoFireRight = "data\\global\\ui\\FrontEnd\\D2logoFireRight.DC6";
const std::string D2ResourcePath::Diablo2LogoBlackLeft = "data\\global\\ui\\FrontEnd\\D2logoBlackLeft.DC6";
const std::string D2ResourcePath::Diablo2LogoBlackRight = "data\\global\\ui\\FrontEnd\\D2logoBlackRight.DC6";

// --- Credits ---
const std::string D2ResourcePath::CreditsBackground = "data\\global\\ui\\CharSelect\\creditsbckgexpand.dc6";
const std::string D2ResourcePath::CreditsText = "data\\local\\ui\\eng\\ExpansionCredits.txt";

// --- Character Select Screen ---
const std::string D2ResourcePath::CharacterSelectBackground = "data\\global\\ui\\FrontEnd\\charactercreationscreenEXP.dc6";
const std::string D2ResourcePath::CharacterSelectCampfire = "data\\global\\ui\\FrontEnd\\fire.DC6";

const std::string D2ResourcePath::CharacterSelectBarbarianUnselected = "data\\global\\ui\\FrontEnd\\barbarian\\banu1.DC6";
const std::string D2ResourcePath::CharacterSelectBarbarianUnselectedH = "data\\global\\ui\\FrontEnd\\barbarian\\banu2.DC6";
const std::string D2ResourcePath::CharacterSelectBarbarianSelected = "data\\global\\ui\\FrontEnd\\barbarian\\banu3.DC6";
const std::string D2ResourcePath::CharacterSelectBarbarianForwardWalk = "data\\global\\ui\\FrontEnd\\barbarian\\bafw.DC6";
const std::string D2ResourcePath::CharacterSelectBarbarianForwardWalkOverlay = "data\\global\\ui\\FrontEnd\\barbarian\\BAFWs.DC6";
const std::string D2ResourcePath::CharacterSelectBarbarianBackWalk = "data\\global\\ui\\FrontEnd\\barbarian\\babw.DC6";

const std::string D2ResourcePath::CharacterSelecSorceressUnselected = "data\\global\\ui\\FrontEnd\\sorceress\\SONU1.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressUnselectedH = "data\\global\\ui\\FrontEnd\\sorceress\\SONU2.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressSelected = "data\\global\\ui\\FrontEnd\\sorceress\\SONU3.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressSelectedOverlay = "data\\global\\ui\\FrontEnd\\sorceress\\SONU3s.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressForwardWalk = "data\\global\\ui\\FrontEnd\\sorceress\\SOFW.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressForwardWalkOverlay = "data\\global\\ui\\FrontEnd\\sorceress\\SOFWs.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressBackWalk = "data\\global\\ui\\FrontEnd\\sorceress\\SOBW.DC6";
const std::string D2ResourcePath::CharacterSelecSorceressBackWalkOverlay = "data\\global\\ui\\FrontEnd\\sorceress\\SOBWs.DC6";

const std::string D2ResourcePath::CharacterSelectNecromancerUnselected = "data\\global\\ui\\FrontEnd\\necromancer\\NENU1.DC6";
const std::string D2ResourcePath::CharacterSelectNecromancerUnselectedH = "data\\global\\ui\\FrontEnd\\necromancer\\NENU2.DC6";
const std::string D2ResourcePath::CharacterSelecNecromancerSelected = "data\\global\\ui\\FrontEnd\\necromancer\\NENU3.DC6";
const std::string D2ResourcePath::CharacterSelecNecromancerSelectedOverlay = "data\\global\\ui\\FrontEnd\\necromancer\\NENU3s.DC6";
const std::string D2ResourcePath::CharacterSelecNecromancerForwardWalk = "data\\global\\ui\\FrontEnd\\necromancer\\NEFW.DC6";
const std::string D2ResourcePath::CharacterSelecNecromancerForwardWalkOverlay = "data\\global\\ui\\FrontEnd\\necromancer\\NEFWs.DC6";
const std::string D2ResourcePath::CharacterSelecNecromancerBackWalk = "data\\global\\ui\\FrontEnd\\necromancer\\NEBW.DC6";
const std::string D2ResourcePath::CharacterSelecNecromancerBackWalkOverlay = "data\\global\\ui\\FrontEnd\\necromancer\\NEBWs.DC6";

const std::string D2ResourcePath::CharacterSelectPaladinUnselected = "data\\global\\ui\\FrontEnd\\paladin\\PANU1.DC6";
const std::string D2ResourcePath::CharacterSelectPaladinUnselectedH = "data\\global\\ui\\FrontEnd\\paladin\\PANU2.DC6";
const std::string D2ResourcePath::CharacterSelecPaladinSelected = "data\\global\\ui\\FrontEnd\\paladin\\PANU3.DC6";
const std::string D2ResourcePath::CharacterSelecPaladinForwardWalk = "data\\global\\ui\\FrontEnd\\paladin\\PAFW.DC6";
const std::string D2ResourcePath::CharacterSelecPaladinForwardWalkOverlay = "data\\global\\ui\\FrontEnd\\paladin\\PAFWs.DC6";
const std::string D2ResourcePath::CharacterSelecPaladinBackWalk = "data\\global\\ui\\FrontEnd\\paladin\\PABW.DC6";


const std::string D2ResourcePath::CharacterSelectAmazonUnselected = "data\\global\\ui\\FrontEnd\\amazon\\AMNU1.DC6";
const std::string D2ResourcePath::CharacterSelectAmazonUnselectedH = "data\\global\\ui\\FrontEnd\\amazon\\AMNU2.DC6";
const std::string D2ResourcePath::CharacterSelecAmazonSelected = "data\\global\\ui\\FrontEnd\\amazon\\AMNU3.DC6";
const std::string D2ResourcePath::CharacterSelecAmazonForwardWalk = "data\\global\\ui\\FrontEnd\\amazon\\AMFW.DC6";
const std::string D2ResourcePath::CharacterSelecAmazonForwardWalkOverlay = "data\\global\\ui\\FrontEnd\\amazon\\AMFWs.DC6";
const std::string D2ResourcePath::CharacterSelecAmazonBackWalk = "data\\global\\ui\\FrontEnd\\amazon\\AMBW.DC6";

const std::string D2ResourcePath::CharacterSelectAssassinUnselected = "data\\global\\ui\\FrontEnd\\assassin\\ASNU1.DC6";
const std::string D2ResourcePath::CharacterSelectAssassinUnselectedH = "data\\global\\ui\\FrontEnd\\assassin\\ASNU2.DC6";
const std::string D2ResourcePath::CharacterSelectAssassinSelected = "data\\global\\ui\\FrontEnd\\assassin\\ASNU3.DC6";
const std::string D2ResourcePath::CharacterSelectAssassinForwardWalk = "data\\global\\ui\\FrontEnd\\assassin\\ASFW.DC6";
const std::string D2ResourcePath::CharacterSelectAssassinBackWalk = "data\\global\\ui\\FrontEnd\\assassin\\ASBW.DC6";

const std::string D2ResourcePath::CharacterSelectDruidUnselected = "data\\global\\ui\\FrontEnd\\druid\\DZNU1.dc6";
const std::string D2ResourcePath::CharacterSelectDruidUnselectedH = "data\\global\\ui\\FrontEnd\\druid\\DZNU2.dc6";
const std::string D2ResourcePath::CharacterSelectDruidSelected = "data\\global\\ui\\FrontEnd\\druid\\DZNU3.DC6";
const std::string D2ResourcePath::CharacterSelectDruidForwardWalk = "data\\global\\ui\\FrontEnd\\druid\\DZFW.DC6";
const std::string D2ResourcePath::CharacterSelectDruidBackWalk = "data\\global\\ui\\FrontEnd\\druid\\DZBW.DC6";

// -- Character Selection
const std::string D2ResourcePath::CharacterSelectionBackground = "data\\global\\ui\\CharSelect\\characterselectscreenEXP.dc6";

// --- Game ---
const std::string D2ResourcePath::GamePanels = "data\\global\\ui\\PANEL\\800ctrlpnl7.dc6";
const std::string D2ResourcePath::GameGlobeOverlap = "data\\global\\ui\\PANEL\\overlap.DC6";
const std::string D2ResourcePath::HealthMana = "data\\global\\ui\\PANEL\\hlthmana.DC6";
const std::string D2ResourcePath::GameSmallMenuButton = "data\\global\\ui\\PANEL\\menubutton.DC6"; // TODO: Used for inventory popout
const std::string D2ResourcePath::SkillIcon = "data\\global\\ui\\PANEL\\Skillicon.DC6"; // TODO: Used for skill icon button
const std::string D2ResourcePath::AddSkillButton = "data\\global\\ui\\PANEL\\level.DC6";

// --- Mouse Pointers ---
const std::string D2ResourcePath::CursorDefault = "data\\global\\ui\\CURSOR\\ohand.DC6";

// --- Fonts ---
const std::string D2ResourcePath::Font6 = "data\\local\\font\\latin\\font6";
const std::string D2ResourcePath::Font8 = "data\\local\\font\\latin\\font8";
const std::string D2ResourcePath::Font16 = "data\\local\\font\\latin\\font16";
const std::string D2ResourcePath::Font24 = "data\\local\\font\\latin\\font24";
const std::string D2ResourcePath::Font30 = "data\\local\\font\\latin\\font30";
const std::string D2ResourcePath::FontFormal12 = "data\\local\\font\\latin\\fontformal12";
const std::string D2ResourcePath::FontFormal11 = "data\\local\\font\\latin\\fontformal11";
const std::string D2ResourcePath::FontFormal10 = "data\\local\\font\\latin\\fontformal10";
const std::string D2ResourcePath::FontExocet10 = "data\\local\\font\\latin\\fontexocet10";
const std::string D2ResourcePath::FontExocet8 = "data\\local\\font\\latin\\fontexocet8";

// --- UI ---
const std::string D2ResourcePath::WideButtonBlank = "data\\global\\ui\\FrontEnd\\WideButtonBlank.dc6";
const std::string D2ResourcePath::MediumButtonBlank = "data\\global\\ui\\FrontEnd\\MediumButtonBlank.dc6";
const std::string D2ResourcePath::CancelButton = "data\\global\\ui\\FrontEnd\\CancelButtonBlank.dc6";
const std::string D2ResourcePath::NarrowButtonBlank = "data\\global\\ui\\FrontEnd\\NarrowButtonBlank.dc6";
const std::string D2ResourcePath::ShortButtonBlank = "data\\global\\ui\\CharSelect\\ShortButtonBlank.dc6";
const std::string D2ResourcePath::TextBox2 = "data\\global\\ui\\FrontEnd\\textbox2.dc6";
const std::string D2ResourcePath::TallButtonBlank = "data\\global\\ui\\CharSelect\\TallButtonBlank.dc6";

// --- GAME UI ---
const std::string D2ResourcePath::MinipanelSmall = "data\\global\\ui\\PANEL\\minipanel_s.dc6";
const std::string D2ResourcePath::MinipanelButton = "data\\global\\ui\\PANEL\\minipanelbtn.DC6";

const std::string D2ResourcePath::Frame = "data\\global\\ui\\PANEL\\800borderframe.dc6";
const std::string D2ResourcePath::InventoryCharacterPanel = "data\\global\\ui\\PANEL\\invchar6.DC6";
const std::string D2ResourcePath::InventoryWeaponsTab = "data\\global\\ui\\PANEL\\invchar6Tab.DC6";
const std::string D2ResourcePath::SkillsPanelAmazon = "data\\global\\ui\\SPELLS\\skltree_a_back.DC6";
const std::string D2ResourcePath::SkillsPanelBarbarian = "data\\global\\ui\\SPELLS\\skltree_b_back.DC6";
const std::string D2ResourcePath::SkillsPanelDruid = "data\\global\\ui\\SPELLS\\skltree_d_back.DC6";
const std::string D2ResourcePath::SkillsPanelAssassin = "data\\global\\ui\\SPELLS\\skltree_i_back.DC6";
const std::string D2ResourcePath::SkillsPanelNecromancer = "data\\global\\ui\\SPELLS\\skltree_n_back.DC6";
const std::string D2ResourcePath::SkillsPanelPaladin = "data\\global\\ui\\SPELLS\\skltree_p_back.DC6";
const std::string D2ResourcePath::SkillsPanelSorcerer = "data\\global\\ui\\SPELLS\\skltree_s_back.DC6";

const std::string D2ResourcePath::GenericSkills = "data\\global\\ui\\SPELLS\\Skillicon.DC6";
const std::string D2ResourcePath::AmazonSkills = "data\\global\\ui\\SPELLS\\AmSkillicon.DC6";
const std::string D2ResourcePath::BarbarianSkills = "data\\global\\ui\\SPELLS\\BaSkillicon.DC6";
const std::string D2ResourcePath::DruidSkills = "data\\global\\ui\\SPELLS\\DrSkillicon.DC6";
const std::string D2ResourcePath::AssassinSkills = "data\\global\\ui\\SPELLS\\AsSkillicon.DC6";
const std::string D2ResourcePath::NecromancerSkills = "data\\global\\ui\\SPELLS\\NeSkillicon.DC6";
const std::string D2ResourcePath::PaladinSkills = "data\\global\\ui\\SPELLS\\PaSkillicon.DC6";
const std::string D2ResourcePath::SorcererSkills = "data\\global\\ui\\SPELLS\\SoSkillicon.DC6";

const std::string D2ResourcePath::RunButton = "data\\global\\ui\\PANEL\\runbutton.dc6";
const std::string D2ResourcePath::MenuButton = "data\\global\\ui\\PANEL\\menubutton.DC6";
const std::string D2ResourcePath::GoldCoinButton = "data\\global\\ui\\panel\\goldcoinbtn.dc6";
const std::string D2ResourcePath::SquareButton = "data\\global\\ui\\panel\\buysellbtn.dc6";

const std::string D2ResourcePath::ArmorPlaceholder = "data\\global\\ui\\PANEL\\inv_armor.DC6";
const std::string D2ResourcePath::BeltPlaceholder = "data\\global\\ui\\PANEL\\inv_belt.DC6";
const std::string D2ResourcePath::BootsPlaceholder = "data\\global\\ui\\PANEL\\inv_boots.DC6";
const std::string D2ResourcePath::HelmGlovePlaceholder = "data\\global\\ui\\PANEL\\inv_helm_glove.DC6";
const std::string D2ResourcePath::RingAmuletPlaceholder = "data\\global\\ui\\PANEL\\inv_ring_amulet.DC6";
const std::string D2ResourcePath::WeaponsPlaceholder = "data\\global\\ui\\PANEL\\inv_weapons.DC6";

// --- Data ---
const std::string D2ResourcePath::EnglishTable = "data\\local\\lng\\eng\\English.txt";
const std::string D2ResourcePath::ExpansionStringTable = "data\\local\\lng\\eng\\expansionstring.tbl";
const std::string D2ResourcePath::LevelPreset = "data\\global\\excel\\LvlPrest.txt";
const std::string D2ResourcePath::LevelType = "data\\global\\excel\\LvlTypes.txt";
const std::string D2ResourcePath::LevelDetails = "data\\global\\excel\\Levels.txt";
const std::string D2ResourcePath::ObjectDetails = "data\\global\\excel\\Objects.txt";

// --- Animations ---
const std::string D2ResourcePath::ObjectData = "data\\global\\objects";
const std::string D2ResourcePath::AnimationData = "data\\global\\animdata.d2";
const std::string D2ResourcePath::PlayerAnimationBase = "data\\global\\CHARS";

// --- Inventory Data ---
const std::string D2ResourcePath::Weapons = "data\\global\\excel\\weapons.txt";
const std::string D2ResourcePath::Armor = "data\\global\\excel\\armor.txt";
const std::string D2ResourcePath::Misc = "data\\global\\excel\\misc.txt";

// --- Character Data ---
const std::string D2ResourcePath::Experience = "data\\global\\excel\\experience.txt";
const std::string D2ResourcePath::CharStats = "data\\global\\excel\\charstats.txt";

// --- Music ---
const std::string D2ResourcePath::BGMTitle = "data\\global\\music\\introedit.wav";
const std::string D2ResourcePath::BGMOptions = "data\\global\\music\\Common\\options.wav";
const std::string D2ResourcePath::BGMAct1AndarielAction = "data\\global\\music\\Act1\\andarielaction.wav";
const std::string D2ResourcePath::BGMAct1BloodRavenResolution = "data\\global\\music\\Act1\\bloodravenresolution.wav";
const std::string D2ResourcePath::BGMAct1Caves = "data\\global\\music\\Act1\\caves.wav";
const std::string D2ResourcePath::BGMAct1Crypt = "data\\global\\music\\Act1\\crypt.wav";
const std::string D2ResourcePath::BGMAct1DenOfEvilAction = "data\\global\\music\\Act1\\denofevilaction.wav";
const std::string D2ResourcePath::BGMAct1Monastery = "data\\global\\music\\Act1\\monastery.wav";
const std::string D2ResourcePath::BGMAct1Town1 = "data\\global\\music\\Act1\\town1.wav";
const std::string D2ResourcePath::BGMAct1Tristram = "data\\global\\music\\Act1\\tristram.wav";
const std::string D2ResourcePath::BGMAct1Wild = "data\\global\\music\\Act1\\wild.wav";
const std::string D2ResourcePath::BGMAct2Desert = "data\\global\\music\\Act2\\desert.wav";
const std::string D2ResourcePath::BGMAct2Harem = "data\\global\\music\\Act2\\harem.wav";
const std::string D2ResourcePath::BGMAct2HoradricAction = "data\\global\\music\\Act2\\horadricaction.wav";
const std::string D2ResourcePath::BGMAct2Lair = "data\\global\\music\\Act2\\lair.wav";
const std::string D2ResourcePath::BGMAct2RadamentResolution = "data\\global\\music\\Act2\\radamentresolution.wav";
const std::string D2ResourcePath::BGMAct2Sanctuary = "data\\global\\music\\Act2\\sanctuary.wav";
const std::string D2ResourcePath::BGMAct2Sewer = "data\\global\\music\\Act2\\sewer.wav";
const std::string D2ResourcePath::BGMAct2TaintedSunAction = "data\\global\\music\\Act2\\taintedsunaction.wav";
const std::string D2ResourcePath::BGMAct2Tombs = "data\\global\\music\\Act2\\tombs.wav";
const std::string D2ResourcePath::BGMAct2Town2 = "data\\global\\music\\Act2\\town2.wav";
const std::string D2ResourcePath::BGMAct2Valley = "data\\global\\music\\Act2\\valley.wav";
const std::string D2ResourcePath::BGMAct3Jungle = "data\\global\\music\\Act3\\jungle.wav";
const std::string D2ResourcePath::BGMAct3Kurast = "data\\global\\music\\Act3\\kurast.wav";
const std::string D2ResourcePath::BGMAct3KurastSewer = "data\\global\\music\\Act3\\kurastsewer.wav";
const std::string D2ResourcePath::BGMAct3MefDeathAction = "data\\global\\music\\Act3\\mefdeathaction.wav";
const std::string D2ResourcePath::BGMAct3OrbAction = "data\\global\\music\\Act3\\orbaction.wav";
const std::string D2ResourcePath::BGMAct3Spider = "data\\global\\music\\Act3\\spider.wav";
const std::string D2ResourcePath::BGMAct3Town3 = "data\\global\\music\\Act3\\town3.wav";
const std::string D2ResourcePath::BGMAct4Diablo = "data\\global\\music\\Act4\\diablo.wav";
const std::string D2ResourcePath::BGMAct4DiabloAction = "data\\global\\music\\Act4\\diabloaction.wav";
const std::string D2ResourcePath::BGMAct4ForgeAction = "data\\global\\music\\Act4\\forgeaction.wav";
const std::string D2ResourcePath::BGMAct4IzualAction = "data\\global\\music\\Act4\\izualaction.wav";
const std::string D2ResourcePath::BGMAct4Mesa = "data\\global\\music\\Act4\\mesa.wav";
const std::string D2ResourcePath::BGMAct4Town4 = "data\\global\\music\\Act4\\town4.wav";
const std::string D2ResourcePath::BGMAct5Baal = "data\\global\\music\\Act5\\baal.wav";
const std::string D2ResourcePath::BGMAct5XTown = "data\\global\\music\\Act5\\xtown.wav";


// --- Sound Effects ---
const std::string D2ResourcePath::SFXButtonClick = "data\\global\\sfx\\Cursor\\button.wav";
const std::string D2ResourcePath::SFXAmazonDeselect = "data\\global\\sfx\\Cursor\\intro\\amazon deselect.wav";
const std::string D2ResourcePath::SFXAmazonSelect = "data\\global\\sfx\\Cursor\\intro\\amazon select.wav";
const std::string D2ResourcePath::SFXAssassinDeselect = "data\\global\\sfx\\Cursor\\intro\\assassin deselect.wav";
const std::string D2ResourcePath::SFXAssassinSelect = "data\\global\\sfx\\Cursor\\intro\\assassin select.wav";
const std::string D2ResourcePath::SFXBarbarianDeselect = "data\\global\\sfx\\Cursor\\intro\\barbarian deselect.wav";
const std::string D2ResourcePath::SFXBarbarianSelect = "data\\global\\sfx\\Cursor\\intro\\barbarian select.wav";
const std::string D2ResourcePath::SFXDruidDeselect = "data\\global\\sfx\\Cursor\\intro\\druid deselect.wav";
const std::string D2ResourcePath::SFXDruidSelect = "data\\global\\sfx\\Cursor\\intro\\druid select.wav";
const std::string D2ResourcePath::SFXNecromancerDeselect = "data\\global\\sfx\\Cursor\\intro\\necromancer deselect.wav";
const std::string D2ResourcePath::SFXNecromancerSelect = "data\\global\\sfx\\Cursor\\intro\\necromancer select.wav";
const std::string D2ResourcePath::SFXPaladinDeselect = "data\\global\\sfx\\Cursor\\intro\\paladin deselect.wav";
const std::string D2ResourcePath::SFXPaladinSelect = "data\\global\\sfx\\Cursor\\intro\\paladin select.wav";
const std::string D2ResourcePath::SFXSorceressDeselect = "data\\global\\sfx\\Cursor\\intro\\sorceress deselect.wav";
const std::string D2ResourcePath::SFXSorceressSelect = "data\\global\\sfx\\Cursor\\intro\\sorceress select.wav";

// --- Enemy Data ---
const std::string D2ResourcePath::MonStats = "data\\global\\excel\\monstats.txt";

// --- Skill Data ---
const std::string D2ResourcePath::Missiles = "data\\global\\excel\\missiles.txt";

}

#endif // OPENDIABLO2_GAME_COMMON_D2RESOURCEPATH_H
