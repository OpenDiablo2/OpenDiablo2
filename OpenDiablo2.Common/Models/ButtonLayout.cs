using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public class ButtonLayout
    {
        public int XSegments { get; internal set; }
        public string ResourceName { get; internal set; }
        public string PaletteName { get; internal set; }
        public bool Toggleable { get; internal set; } = false;
        public int BaseFrame { get; internal set; } = 0;

        public static Dictionary<eButtonType, ButtonLayout> Values = new Dictionary<eButtonType, ButtonLayout>
        {
            {eButtonType.Wide,  new ButtonLayout { XSegments = 2, ResourceName = ResourcePaths.WideButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Medium, new ButtonLayout{ XSegments = 1, ResourceName=ResourcePaths.MediumButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Narrow, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.NarrowButtonBlank,PaletteName = Palettes.Units } },
            {eButtonType.Tall, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.TallButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Cancel, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.CancelButton,PaletteName = Palettes.Units } },
            // Minipanel
            {eButtonType.MinipanelCharacter, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 0 } },
            {eButtonType.MinipanelInventory, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 2 } },
            {eButtonType.MinipanelSkill, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 4 } },
            {eButtonType.MinipanelAutomap, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 8 } },
            {eButtonType.MinipanelMessage, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 10 } },
            {eButtonType.MinipanelQuest, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 12 } },
            {eButtonType.MinipanelMenu, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MinipanelButton,PaletteName = Palettes.Units, BaseFrame = 14 } },

            {eButtonType.Run, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.RunButton,PaletteName = Palettes.Units, Toggleable = true } },
            {eButtonType.Menu, new ButtonLayout {XSegments = 1, ResourceName = ResourcePaths.MenuButton,PaletteName = Palettes.Units, Toggleable = true } },
        };
    }

}
