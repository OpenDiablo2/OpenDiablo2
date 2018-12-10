/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common
{
    public static class ResourcePaths
    {
        // --- Loading Screen ---
        public const string LoadingScreen  = @"data\global\ui\Loading\loadingscreen.dc6";

        // --- Main Menu ---
        public const string GameSelectScreen = @"data\global\ui\FrontEnd\gameselectscreenEXP.dc6";
        public const string Diablo2LogoFireLeft = @"data\global\ui\FrontEnd\D2logoFireLeft.DC6";
        public const string Diablo2LogoFireRight = @"data\global\ui\FrontEnd\D2logoFireRight.DC6";
        public const string Diablo2LogoBlackLeft = @"data\global\ui\FrontEnd\D2logoBlackLeft.DC6";
        public const string Diablo2LogoBlackRight = @"data\global\ui\FrontEnd\D2logoBlackRight.DC6";

        // --- Character Select Screen ---
        public const string CharacterSelectBackground = @"data\global\ui\FrontEnd\charactercreationscreenEXP.dc6";
        public const string CharacterSelectCampfire = @"data\global\ui\FrontEnd\fire.DC6";

        public const string CharacterSelectBarbarianUnselected = @"data\global\ui\FrontEnd\barbarian\banu1.DC6";
        public const string CharacterSelectBarbarianUnselectedH = @"data\global\ui\FrontEnd\barbarian\banu2.DC6";
        public const string CharacterSelectBarbarianSelected = @"data\global\ui\FrontEnd\barbarian\banu3.DC6";
        public const string CharacterSelectBarbarianForwardWalk = @"data\global\ui\FrontEnd\barbarian\bafw.DC6";
        public const string CharacterSelectBarbarianForwardWalkOverlay = @"data\global\ui\FrontEnd\barbarian\BAFWs.DC6";
        public const string CharacterSelectBarbarianBackWalk = @"data\global\ui\FrontEnd\barbarian\babw.DC6";

        public const string CharacterSelecSorceressUnselected = @"data\global\ui\FrontEnd\sorceress\SONU1.DC6";
        public const string CharacterSelecSorceressUnselectedH = @"data\global\ui\FrontEnd\sorceress\SONU2.DC6";
        public const string CharacterSelecSorceressSelected = @"data\global\ui\FrontEnd\sorceress\SONU3.DC6";
        public const string CharacterSelecSorceressSelectedOverlay = @"data\global\ui\FrontEnd\sorceress\SONU3s.DC6";
        public const string CharacterSelecSorceressForwardWalk = @"data\global\ui\FrontEnd\sorceress\SOFW.DC6";
        public const string CharacterSelecSorceressForwardWalkOverlay = @"data\global\ui\FrontEnd\sorceress\SOFWs.DC6";
        public const string CharacterSelecSorceressBackWalk = @"data\global\ui\FrontEnd\sorceress\SOBW.DC6";
        public const string CharacterSelecSorceressBackWalkOverlay = @"data\global\ui\FrontEnd\sorceress\SOBWs.DC6";

        public const string CharacterSelectNecromancerUnselected = @"data\global\ui\FrontEnd\necromancer\NENU1.DC6";
        public const string CharacterSelectNecromancerUnselectedH = @"data\global\ui\FrontEnd\necromancer\NENU2.DC6";
        public const string CharacterSelecNecromancerSelected = @"data\global\ui\FrontEnd\necromancer\NENU3.DC6";
        public const string CharacterSelecNecromancerSelectedOverlay = @"data\global\ui\FrontEnd\necromancer\NENU3s.DC6";
        public const string CharacterSelecNecromancerForwardWalk = @"data\global\ui\FrontEnd\necromancer\NEFW.DC6";
        public const string CharacterSelecNecromancerForwardWalkOverlay = @"data\global\ui\FrontEnd\necromancer\NEFWs.DC6";
        public const string CharacterSelecNecromancerBackWalk = @"data\global\ui\FrontEnd\necromancer\NEBW.DC6";
        public const string CharacterSelecNecromancerBackWalkOverlay = @"data\global\ui\FrontEnd\necromancer\NEBWs.DC6";

        public const string CharacterSelectPaladinUnselected = @"data\global\ui\FrontEnd\paladin\PANU1.DC6";
        public const string CharacterSelectPaladinUnselectedH = @"data\global\ui\FrontEnd\paladin\PANU2.DC6";
        public const string CharacterSelecPaladinSelected = @"data\global\ui\FrontEnd\paladin\PANU3.DC6";
        public const string CharacterSelecPaladinForwardWalk = @"data\global\ui\FrontEnd\paladin\PAFW.DC6";
        public const string CharacterSelecPaladinForwardWalkOverlay = @"data\global\ui\FrontEnd\paladin\PAFWs.DC6";
        public const string CharacterSelecPaladinBackWalk = @"data\global\ui\FrontEnd\paladin\PABW.DC6";


        public const string CharacterSelectAmazonUnselected = @"data\global\ui\FrontEnd\amazon\AMNU1.DC6";
        public const string CharacterSelectAmazonUnselectedH = @"data\global\ui\FrontEnd\amazon\AMNU2.DC6";
        public const string CharacterSelecAmazonSelected = @"data\global\ui\FrontEnd\amazon\AMNU3.DC6";
        public const string CharacterSelecAmazonForwardWalk = @"data\global\ui\FrontEnd\amazon\AMFW.DC6";
        public const string CharacterSelecAmazonForwardWalkOverlay = @"data\global\ui\FrontEnd\amazon\AMFWs.DC6";
        public const string CharacterSelecAmazonBackWalk = @"data\global\ui\FrontEnd\amazon\AMBW.DC6";

        public const string CharacterSelectAssassinUnselected = @"data\global\ui\FrontEnd\assassin\ASNU1.DC6";
        public const string CharacterSelectAssassinUnselectedH = @"data\global\ui\FrontEnd\assassin\ASNU2.DC6";
        public const string CharacterSelectAssassinSelected = @"data\global\ui\FrontEnd\assassin\ASNU3.DC6";
        public const string CharacterSelectAssassinForwardWalk = @"data\global\ui\FrontEnd\assassin\ASFW.DC6";
        public const string CharacterSelectAssassinBackWalk = @"data\global\ui\FrontEnd\assassin\ASBW.DC6";

        public const string CharacterSelectDruidUnselected = @"data\global\ui\FrontEnd\druid\DZNU1.dc6";
        public const string CharacterSelectDruidUnselectedH = @"data\global\ui\FrontEnd\druid\DZNU2.dc6";
        public const string CharacterSelectDruidSelected = @"data\global\ui\FrontEnd\druid\DZNU3.DC6";
        public const string CharacterSelectDruidForwardWalk = @"data\global\ui\FrontEnd\druid\DZFW.DC6";
        public const string CharacterSelectDruidBackWalk = @"data\global\ui\FrontEnd\druid\DZBW.DC6";

        // -- Character Selection
        public const string CharacterSelectionBackground = @"data\global\ui\CharSelect\characterselectscreenEXP.dc6";
        
        // --- Game ---
        public const string GamePanels = @"data\global\ui\PANEL\800ctrlpnl7.dc6";
        public const string GameGlobeOverlap = @"data\global\ui\PANEL\overlap.DC6";
        public const string HealthMana = @"data\global\ui\PANEL\hlthmana.DC6";
        public const string GameSmallMenuButton = @"data\global\ui\PANEL\menubutton.DC6"; // TODO: Used for inventory popout
        public const string SkillIcon = @"data\global\ui\PANEL\Skillicon.DC6"; // TODO: Used for skill icon button

        // --- Mouse Pointers ---
        public const string CursorDefault = @"data\global\ui\CURSOR\ohand.DC6";

        // --- Fonts ---
        public const string Font6 = @"data\local\font\latin\font6";
        public const string Font8 = @"data\local\font\latin\font8";
        public const string Font16 = @"data\local\font\latin\font16";
        public const string Font24 = @"data\local\font\latin\font24";
        public const string Font30 = @"data\local\font\latin\font30";
        public const string FontFormal12 = @"data\local\font\latin\fontformal12";
        public const string FontFormal11 = @"data\local\font\latin\fontformal11";
        public const string FontFormal10 = @"data\local\font\latin\fontformal10";
        public const string FontExocet10 = @"data\local\font\latin\fontexocet10";
        public const string FontExocet8 = @"data\local\font\latin\fontexocet8";

        // --- UI ---
        public const string WideButtonBlank = @"data\global\ui\FrontEnd\WideButtonBlank.dc6";
        public const string MediumButtonBlank = @"data\global\ui\FrontEnd\MediumButtonBlank.dc6";
        public const string CancelButton = @"data\global\ui\FrontEnd\CancelButtonBlank.dc6";
        public const string NarrowButtonBlank = @"data\global\ui\FrontEnd\NarrowButtonBlank.dc6";
        public const string TextBox2 = @"data\global\ui\FrontEnd\textbox2.dc6";
        public const string TallButtonBlank = @"data\global\ui\CharSelect\TallButtonBlank.dc6";

        // --- GAME UI ---
        public const string MinipanelSmall = @"data\global\ui\PANEL\minipanel_s.dc6";
        public const string MinipanelButton = @"data\global\ui\PANEL\minipanelbtn.DC6";

        public const string Frame = @"data\global\ui\PANEL\800borderframe.dc6";
        public const string InventoryCharacterPanel = @"data\global\ui\PANEL\invchar6.DC6";
        public const string InventoryWeaponsTab = @"data\global\ui\PANEL\invchar6Tab.DC6";

        public const string RunButton = @"data\global\ui\PANEL\runbutton.dc6";
        public const string MenuButton = @"data\global\ui\PANEL\menubutton.DC6";
        public const string GoldCoinButton = @"data\global\ui\panel\goldcoinbtn.dc6";
        public const string SquareButton = @"data\global\ui\panel\buysellbtn.dc6";

        public const string ArmorPlaceholder = @"data\global\ui\PANEL\inv_armor.DC6";
        public const string BeltPlaceholder = @"data\global\ui\PANEL\inv_belt.DC6";
        public const string BootsPlaceholder = @"data\global\ui\PANEL\inv_boots.DC6";
        public const string HelmGlovePlaceholder = @"data\global\ui\PANEL\inv_helm_glove.DC6";
        public const string RingAmuletPlaceholder = @"data\global\ui\PANEL\inv_ring_amulet.DC6";
        public const string WeaponsPlaceholder = @"data\global\ui\PANEL\inv_weapons.DC6";

        // --- Data ---
        // TODO: Doesn't sound right :)
        public const string EnglishTable = @"data\local\lng\eng\English.txt";
        public const string ExpansionStringTable = @"data\local\lng\eng\expansionstring.tbl";
        public const string LevelPreset = @"data\global\excel\LvlPrest.txt";
        public const string LevelType = @"data\global\excel\LvlTypes.txt";
        public const string LevelDetails = @"data\global\excel\Levels.txt";

        // --- Animations ---
        public const string ObjectData = @"data\global\objects";
        public const string AnimationData = @"data\global\animdata.d2";
        public const string PlayerAnimationBase = @"data\global\CHARS";

        // --- Inventory Data ---
        public const string Weapons = @"data\global\excel\weapons.txt";
        public const string Armor = @"data\global\excel\armor.txt";
        public const string Misc = @"data\global\excel\misc.txt";

        // --- Character Data ---
        public const string Experience = @"data\global\excel\experience.txt";
        public const string CharStats = @"data\global\excel\charstats.txt";

        // --- Music ---
        public const string BGMTitle = @"data\global\music\introedit.wav";
        public const string BGMOptions = @"data\global\music\Common\options.wav";
        public const string BGMAct1AndarielAction = @"data\global\music\Act1\andarielaction.wav";
        public const string BGMAct1BloodRavenResolution = @"data\global\music\Act1\bloodravenresolution.wav";
        public const string BGMAct1Caves = @"data\global\music\Act1\caves.wav";
        public const string BGMAct1Crypt = @"data\global\music\Act1\crypt.wav";
        public const string BGMAct1DenOfEvilAction = @"data\global\music\Act1\denofevilaction.wav";
        public const string BGMAct1Monastery = @"data\global\music\Act1\monastery.wav";
        public const string BGMAct1Town1 = @"data\global\music\Act1\town1.wav";
        public const string BGMAct1Tristram = @"data\global\music\Act1\tristram.wav";
        public const string BGMAct1Wild = @"data\global\music\Act1\wild.wav";
        public const string BGMAct2Desert = @"data\global\music\Act2\desert.wav";
        public const string BGMAct2Harem = @"data\global\music\Act2\harem.wav";
        public const string BGMAct2HoradricAction = @"data\global\music\Act2\horadricaction.wav";
        public const string BGMAct2Lair = @"data\global\music\Act2\lair.wav";
        public const string BGMAct2RadamentResolution = @"data\global\music\Act2\radamentresolution.wav";
        public const string BGMAct2Sanctuary = @"data\global\music\Act2\sanctuary.wav";
        public const string BGMAct2Sewer = @"data\global\music\Act2\sewer.wav";
        public const string BGMAct2TaintedSunAction = @"data\global\music\Act2\taintedsunaction.wav";
        public const string BGMAct2Tombs = @"data\global\music\Act2\tombs.wav";
        public const string BGMAct2Town2 = @"data\global\music\Act2\town2.wav";
        public const string BGMAct2Valley = @"data\global\music\Act2\valley.wav";
        public const string BGMAct3Jungle = @"data\global\music\Act3\jungle.wav";
        public const string BGMAct3Kurast = @"data\global\music\Act3\kurast.wav";
        public const string BGMAct3KurastSewer = @"data\global\music\Act3\kurastsewer.wav";
        public const string BGMAct3MefDeathAction = @"data\global\music\Act3\mefdeathaction.wav";
        public const string BGMAct3OrbAction = @"data\global\music\Act3\orbaction.wav";
        public const string BGMAct3Spider = @"data\global\music\Act3\spider.wav";
        public const string BGMAct3Town3 = @"data\global\music\Act3\town3.wav";
        public const string BGMAct4Diablo = @"data\global\music\Act4\diablo.wav";
        public const string BGMAct4DiabloAction = @"data\global\music\Act4\diabloaction.wav";
        public const string BGMAct4ForgeAction = @"data\global\music\Act4\forgeaction.wav";
        public const string BGMAct4IzualAction = @"data\global\music\Act4\izualaction.wav";
        public const string BGMAct4Mesa = @"data\global\music\Act4\mesa.wav";
        public const string BGMAct4Town4 = @"data\global\music\Act4\town4.wav";
        public const string BGMAct5Baal = @"data\global\music\Act5\baal.wav";
        public const string BGMAct5XTown = @"data\global\music\Act5\xtown.wav";


        // --- Sound Effects ---
        public const string SFXButtonClick = @"data\global\sfx\Cursor\button.wav";
        public const string SFXAmazonDeselect = @"data\global\sfx\Cursor\intro\amazon deselect.wav";
        public const string SFXAmazonSelect = @"data\global\sfx\Cursor\intro\amazon select.wav";
        public const string SFXAssassinDeselect = @"data\global\sfx\Cursor\intro\assassin deselect.wav";
        public const string SFXAssassinSelect = @"data\global\sfx\Cursor\intro\assassin select.wav";
        public const string SFXBarbarianDeselect = @"data\global\sfx\Cursor\intro\barbarian deselect.wav";
        public const string SFXBarbarianSelect = @"data\global\sfx\Cursor\intro\barbarian select.wav";
        public const string SFXDruidDeselect = @"data\global\sfx\Cursor\intro\druid deselect.wav";
        public const string SFXDruidSelect = @"data\global\sfx\Cursor\intro\druid select.wav";
        public const string SFXNecromancerDeselect = @"data\global\sfx\Cursor\intro\necromancer deselect.wav";
        public const string SFXNecromancerSelect = @"data\global\sfx\Cursor\intro\necromancer select.wav";
        public const string SFXPaladinDeselect = @"data\global\sfx\Cursor\intro\paladin deselect.wav";
        public const string SFXPaladinSelect = @"data\global\sfx\Cursor\intro\paladin select.wav";
        public const string SFXSorceressDeselect = @"data\global\sfx\Cursor\intro\sorceress deselect.wav";
        public const string SFXSorceressSelect = @"data\global\sfx\Cursor\intro\sorceress select.wav";

        public static string GeneratePathForItem(string spriteName)
        {
            return $@"data\global\items\{spriteName}.dc6";
        }

        public static string GetMusicPathForLevel(eLevelId levelId)
        {
            switch (levelId)
            {
                case eLevelId.None:
                    return string.Empty;
                case eLevelId.Act1_Town1:
                    return BGMAct1Town1;
                case eLevelId.Act1_CaveTreasure2:
                    return BGMAct1Caves;
                case eLevelId.Act1_CaveTreasure3:
                    return BGMAct1Caves;
                case eLevelId.Act1_CaveTreasure4:
                    return BGMAct1Caves;
                case eLevelId.Act1_CaveTreasure5:
                    return BGMAct1Caves;
                case eLevelId.Act1_CryptCountessX:
                    return BGMAct1BloodRavenResolution; // TODO: Verify
                case eLevelId.Act1_Tower2:
                    return BGMAct1Caves; // TODO: Verify
                case eLevelId.Act1_MonFront:
                    return BGMAct1DenOfEvilAction; // TODO: Verify
                case eLevelId.Act1_Courtyard1:
                    return BGMAct1Monastery; // TODO: Verify
                case eLevelId.Act1_Courtyard2:
                    return BGMAct1Monastery; // TODO: Verify
                case eLevelId.Act1_Cathedral:
                    return BGMAct1Monastery; // TODO: Verify
                case eLevelId.Act1_Andariel:
                    return BGMAct1AndarielAction; 
                case eLevelId.Act1_Tristram:
                    return BGMAct1Tristram;
                case eLevelId.Act2_Town:
                    return BGMAct2Town2;
                case eLevelId.Act2_Harem:
                    return BGMAct2Harem;
                case eLevelId.Act2_DurielsLair:
                    return BGMAct2Lair;
                case eLevelId.Act3_Town:
                    return BGMAct3Town3;
                case eLevelId.Act3_DungeonTreasure1:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_DungeonTreasure2:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_SewerTreasureX:
                    return BGMAct3KurastSewer; // TODO: Verify
                case eLevelId.Act3_Temple1:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_Temple2:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_Temple3:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_Temple4:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_Temple5:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_Temple6:
                    return BGMAct3Kurast; // TODO: Verify
                case eLevelId.Act3_MephistoComplex:
                    return BGMAct3MefDeathAction; // TODO: Verify
                case eLevelId.Act4_Fortress:
                    return BGMAct4Mesa; // TODO: Verify
                case eLevelId.Act5_Town:
                    return BGMAct5XTown;
                case eLevelId.Act5_TempleFinalRoom:
                    return BGMAct2Sanctuary; // TODO: Verify
                case eLevelId.Act5_ThroneRoom:
                    return BGMAct2Sanctuary; // TODO: Verify
                case eLevelId.Act5_WorldStone:
                    return BGMAct4ForgeAction; // TODO: Verify
                case eLevelId.Act5_TempleEntrance:
                    return BGMAct5Baal; // TODO: Verify
                case eLevelId.Act5_BaalEntrance:
                    return BGMAct5Baal; // TODO: Verify
                default:
                    return string.Empty;
            }
        }
    }
}
