using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common
{
    public static class ResourcePaths
    {
        // --- Loading Screen ---
        public static string LoadingScreen { get; }  = "data\\global\\ui\\Loading\\loadingscreen.dc6";

        // --- Main Menu ---
        public static string GameSelectScreen { get; } = "data\\global\\ui\\FrontEnd\\gameselectscreenEXP.dc6";
        public static string Diablo2LogoFireLeft { get; } = "data\\global\\ui\\FrontEnd\\D2logoFireLeft.DC6";
        public static string Diablo2LogoFireRight { get; } = "data\\global\\ui\\FrontEnd\\D2logoFireRight.DC6";
        public static string Diablo2LogoBlackLeft { get; } = "data\\global\\ui\\FrontEnd\\D2logoBlackLeft.DC6";
        public static string Diablo2LogoBlackRight { get; } = "data\\global\\ui\\FrontEnd\\D2logoBlackRight.DC6";

        // --- Character Select Screen ---
        public static string CharacterSelectBackground { get; } = "data\\global\\ui\\FrontEnd\\charactercreationscreenEXP.dc6";
        public static string CharacterSelectCampfire { get; } = "data\\global\\ui\\FrontEnd\\fire.DC6";

        public static string CharacterSelectBarbarianUnselected { get; } = "data\\global\\ui\\FrontEnd\\barbarian\\banu1.DC6";
        public static string CharacterSelectBarbarianUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\barbarian\\banu2.DC6";
        public static string CharacterSelectBarbarianSelected { get; } = "data\\global\\ui\\FrontEnd\\barbarian\\banu3.DC6";
        public static string CharacterSelectBarbarianForwardWalk { get; } = "data\\global\\ui\\FrontEnd\\barbarian\\bafw.DC6";
        public static string CharacterSelectBarbarianForwardWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\barbarian\\BAFWs.DC6";
        public static string CharacterSelectBarbarianBackWalk { get; } = "data\\global\\ui\\FrontEnd\\barbarian\\babw.DC6";

        public static string CharacterSelecSorceressUnselected { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SONU1.DC6";
        public static string CharacterSelecSorceressUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SONU2.DC6";
        public static string CharacterSelecSorceressSelected { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SONU3.DC6";
        public static string CharacterSelecSorceressSelectedOverlay { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SONU3s.DC6";
        public static string CharacterSelecSorceressForwardWalk{ get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SOFW.DC6";
        public static string CharacterSelecSorceressForwardWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SOFWs.DC6";
        public static string CharacterSelecSorceressBackWalk { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SOBW.DC6";
        public static string CharacterSelecSorceressBackWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\sorceress\\SOBWs.DC6";

        public static string CharacterSelectNecromancerUnselected { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NENU1.DC6";
        public static string CharacterSelectNecromancerUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NENU2.DC6";
        public static string CharacterSelecNecromancerSelected { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NENU3.DC6";
        public static string CharacterSelecNecromancerSelectedOverlay { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NENU3s.DC6";
        public static string CharacterSelecNecromancerForwardWalk { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NEFW.DC6";
        public static string CharacterSelecNecromancerForwardWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NEFWs.DC6";
        public static string CharacterSelecNecromancerBackWalk { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NEBW.DC6";
        public static string CharacterSelecNecromancerBackWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\necromancer\\NEBWs.DC6";

        public static string CharacterSelectPaladinUnselected { get; } = "data\\global\\ui\\FrontEnd\\paladin\\PANU1.DC6";
        public static string CharacterSelectPaladinUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\paladin\\PANU2.DC6";
        public static string CharacterSelecPaladinSelected { get; } = "data\\global\\ui\\FrontEnd\\paladin\\PANU3.DC6";
        public static string CharacterSelecPaladinForwardWalk { get; } = "data\\global\\ui\\FrontEnd\\paladin\\PAFW.DC6";
        public static string CharacterSelecPaladinForwardWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\paladin\\PAFWs.DC6";
        public static string CharacterSelecPaladinBackWalk { get; } = "data\\global\\ui\\FrontEnd\\paladin\\PABW.DC6";


        public static string CharacterSelectAmazonUnselected { get; } = "data\\global\\ui\\FrontEnd\\amazon\\AMNU1.DC6";
        public static string CharacterSelectAmazonUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\amazon\\AMNU2.DC6";
        public static string CharacterSelecAmazonSelected { get; } = "data\\global\\ui\\FrontEnd\\amazon\\AMNU3.DC6";
        public static string CharacterSelecAmazonForwardWalk { get; } = "data\\global\\ui\\FrontEnd\\amazon\\AMFW.DC6";
        public static string CharacterSelecAmazonForwardWalkOverlay { get; } = "data\\global\\ui\\FrontEnd\\amazon\\AMFWs.DC6";
        public static string CharacterSelecAmazonBackWalk { get; } = "data\\global\\ui\\FrontEnd\\amazon\\AMBW.DC6";

        public static string CharacterSelectAssassinUnselected { get; } = "data\\global\\ui\\FrontEnd\\assassin\\ASNU1.DC6";
        public static string CharacterSelectAssassinUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\assassin\\ASNU2.DC6";
        public static string CharacterSelectAssassinSelected { get; } = "data\\global\\ui\\FrontEnd\\assassin\\ASNU3.DC6";
        public static string CharacterSelectAssassinForwardWalk { get; } = "data\\global\\ui\\FrontEnd\\assassin\\ASFW.DC6";
        public static string CharacterSelectAssassinBackWalk { get; } = "data\\global\\ui\\FrontEnd\\assassin\\ASBW.DC6";

        public static string CharacterSelectDruidUnselected { get; } = "data\\global\\ui\\FrontEnd\\druid\\DZNU1.dc6";
        public static string CharacterSelectDruidUnselectedH { get; } = "data\\global\\ui\\FrontEnd\\druid\\DZNU2.dc6";
        public static string CharacterSelectDruidSelected { get; } = "data\\global\\ui\\FrontEnd\\druid\\DZNU3.DC6";
        public static string CharacterSelectDruidForwardWalk { get; } = "data\\global\\ui\\FrontEnd\\druid\\DZFW.DC6";
        public static string CharacterSelectDruidBackWalk { get; } = "data\\global\\ui\\FrontEnd\\druid\\DZBW.DC6";

        // -- Character Selection
        public static string CharacterSelectionBackground { get; } = "data\\global\\ui\\CharSelect\\characterselectscreenEXP.dc6";
        
        // --- Game ---
        public static string GamePanels { get; } = "data\\global\\ui\\PANEL\\800ctrlpnl7.dc6";
        public static string GameGlobeOverlap { get; } = "data\\global\\ui\\PANEL\\overlap.DC6";
        public static string HealthMana { get; } = "data\\global\\ui\\PANEL\\hlthmana.DC6";
        public static string GameSmallMenuButton { get; } = "data\\global\\ui\\PANEL\\menubutton.DC6"; // TODO: Used for inventory popout
        public static string SkillIcon { get; } = "data\\global\\ui\\PANEL\\Skillicon.DC6"; // TODO: Used for skill icon button

        // --- Mouse Pointers ---
        public static string CursorDefault { get; } = "data\\global\\ui\\CURSOR\\ohand.DC6";

        // --- Fonts ---
        public static string Font6 { get; } = "data\\local\\font\\latin\\font6";
        public static string Font8 { get; } = "data\\local\\font\\latin\\font8";
        public static string Font16 { get; } = "data\\local\\font\\latin\\font16";
        public static string Font24 { get; } = "data\\local\\font\\latin\\font24";
        public static string Font30 { get; } = "data\\local\\font\\latin\\font30";
        public static string FontFormal12 { get; } = "data\\local\\font\\latin\\fontformal12";
        public static string FontFormal11 { get; } = "data\\local\\font\\latin\\fontformal11";
        public static string FontFormal10 { get; } = "data\\local\\font\\latin\\fontformal10";
        public static string FontExocet10 { get; } = "data\\local\\font\\latin\\fontexocet10";
        public static string FontExocet8 { get; } = "data\\local\\font\\latin\\fontexocet8";

        // --- UI ---
        public static string WideButtonBlank { get; } = "data\\global\\ui\\FrontEnd\\WideButtonBlank.dc6";
        public static string MediumButtonBlank { get; } = "data\\global\\ui\\FrontEnd\\MediumButtonBlank.dc6";
        public static string CancelButton { get; } = "data\\global\\ui\\FrontEnd\\CancelButtonBlank.dc6";
        public static string NarrowButtonBlank { get; } = "data\\global\\ui\\FrontEnd\\NarrowButtonBlank.dc6";
        public static string TextBox2 { get; } = "data\\global\\ui\\FrontEnd\\textbox2.dc6";
        public static string TallButtonBlank { get; } = "data\\global\\ui\\CharSelect\\TallButtonBlank.dc6";

        // --- GAME UI ---
        public static string MinipanelSmall { get; } = "data\\global\\ui\\PANEL\\minipanel_s.dc6";
        public static string MinipanelButton { get; } = "data\\global\\ui\\PANEL\\minipanelbtn.DC6";

        public static string Frame { get; } = "data\\global\\ui\\PANEL\\800borderframe.dc6";
        public static string InventoryCharacterPanel { get; } = "data\\global\\ui\\PANEL\\invchar.DC6";

        public static string RunButton { get; } = "data\\global\\ui\\PANEL\\runbutton.dc6";
        public static string MenuButton { get; } = "data\\global\\ui\\PANEL\\menubutton.DC6";

        public static string ArmorPlaceholder { get; } = "data\\global\\ui\\PANEL\\inv_armor.DC6";
        public static string BeltPlaceholder { get; } = "data\\global\\ui\\PANEL\\inv_belt.DC6";
        public static string BootsPlaceholder { get; } = "data\\global\\ui\\PANEL\\inv_boots.DC6";
        public static string HelmGlovePlaceholder { get; } = "data\\global\\ui\\PANEL\\inv_helm_glove.DC6";
        public static string RingAmuletPlaceholder { get; } = "data\\global\\ui\\PANEL\\inv_ring_amulet.DC6";
        public static string WeaponsPlaceholder { get; } = "data\\global\\ui\\PANEL\\inv_weapons.DC6";

        // --- Data ---
        // TODO: Doesn't sound right :)
        public static string EnglishTable { get; } = "data\\local\\lng\\eng\\English.txt";
        public static string ExpansionStringTable { get; } = "data\\local\\lng\\eng\\expansionstring.tbl";
        public static string LevelPreset { get; } = "data\\global\\excel\\LvlPrest.txt";
        public static string LevelType { get; } = "data\\global\\excel\\LvlTypes.txt";
        public static string LevelDetails { get; } = "data\\global\\excel\\Levels.txt";

        // --- Animations ---
        public static string ObjectData { get; } = "data\\global\\objects";
        public static string AnimationData { get; } = "data\\global\\animdata.d2";
        public static string PlayerAnimationBase { get; } = "data\\global\\CHARS";

        // --- Inventory Data ---
        public static string Weapons { get; } = "data\\global\\excel\\weapons.txt";
        public static string Armor { get; } = "data\\global\\excel\\armor.txt";
        public static string Misc { get; } = "data\\global\\excel\\misc.txt";

        // --- Character Data ---
        public static string Experience { get; } = "data\\global\\excel\\experience.txt";
        public static string CharStats { get; } = "data\\global\\excel\\charstats.txt";

        public static string GeneratePathForItem(string spriteName)
        {
            return $"data\\global\\items\\{spriteName}.dc6";
        }
    }
}
