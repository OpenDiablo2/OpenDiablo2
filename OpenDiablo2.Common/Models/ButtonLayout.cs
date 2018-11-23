﻿using System;
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

        public static Dictionary<eButtonType, ButtonLayout> Values = new Dictionary<eButtonType, ButtonLayout>
        {
            {eButtonType.Wide,  new ButtonLayout { XSegments = 2, ResourceName = ResourcePaths.WideButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Medium, new ButtonLayout{ XSegments = 1, ResourceName=ResourcePaths.MediumButtonBlank, PaletteName = Palettes.Units } },
            {eButtonType.Cancel, new ButtonLayout {XSegments = 1,ResourceName = ResourcePaths.CancelButton,PaletteName = Palettes.Units } }
        };
    }

}
