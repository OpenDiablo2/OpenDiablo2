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
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Drawing;

namespace OpenDiablo2.Common.Models
{
    public class ButtonLayout
    {
        public int XSegments { get; set; } = 1;
        public int YSegments { get; set; } = 1;
        public string ResourceName { get; set; }
        public string PaletteName { get; set; }
        public bool Toggleable { get; set; } = false;
        public int BaseFrame { get; set; } = 0;
        public int DisabledFrame { get; set; } = -1;
        public string FontPath { get; set; } = ResourcePaths.FontExocet10;
        public Rectangle ClickableRect { get; set; }
        public bool AllowFrameChange { get; set; } = true;

        public bool IsDarkenedWhenDisabled => DisabledFrame == -1;

        public static ImmutableDictionary<eButtonType, ButtonLayout> Values { get; } = new Dictionary<eButtonType, ButtonLayout>
        {
            {eButtonType.Wide,  new ButtonLayout { XSegments = 2, ResourceName = ResourcePaths.WideButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Medium, new ButtonLayout{ ResourceName = ResourcePaths.MediumButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Narrow, new ButtonLayout { ResourceName = ResourcePaths.NarrowButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Tall, new ButtonLayout { ResourceName = ResourcePaths.TallButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Short, new ButtonLayout { ResourceName = ResourcePaths.ShortButtonBlank, PaletteName = Palettes.Units, FontPath = ResourcePaths.FontExocet10 } },
            {eButtonType.Cancel, new ButtonLayout { ResourceName = ResourcePaths.CancelButton, PaletteName = Palettes.Units } },
            // Minipanel
            {eButtonType.MinipanelCharacter, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 0 } },
            {eButtonType.MinipanelInventory, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 2 } },
            {eButtonType.MinipanelSkill, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 4 } },
            {eButtonType.MinipanelAutomap, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 8 } },
            {eButtonType.MinipanelMessage, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 10 } },
            {eButtonType.MinipanelQuest, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 12 } },
            {eButtonType.MinipanelMenu, new ButtonLayout { ResourceName = ResourcePaths.MinipanelButton, PaletteName = Palettes.Units, BaseFrame = 14 } },
            
            {eButtonType.SecondaryInvHand, new ButtonLayout { ResourceName = ResourcePaths.InventoryWeaponsTab, PaletteName = Palettes.Units,
                ClickableRect = new Rectangle(0, 0, 0, 20), AllowFrameChange = false } },
            {eButtonType.Run, new ButtonLayout { ResourceName = ResourcePaths.RunButton, PaletteName = Palettes.Units, Toggleable = true } },
            {eButtonType.Menu, new ButtonLayout { ResourceName = ResourcePaths.MenuButton, PaletteName = Palettes.Units, Toggleable = true } },
            {eButtonType.GoldCoin, new ButtonLayout { ResourceName = ResourcePaths.GoldCoinButton, PaletteName = Palettes.Units } },
            {eButtonType.Close, new ButtonLayout { ResourceName = ResourcePaths.SquareButton, PaletteName = Palettes.Units, BaseFrame = 10 } },
            {eButtonType.Skill, new ButtonLayout { ResourceName = ResourcePaths.AddSkillButton, PaletteName = Palettes.Units, DisabledFrame = 2 } },
        }.ToImmutableDictionary();
    }

}
