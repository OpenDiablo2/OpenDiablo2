using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Common.Enums
{
    public enum eDrawEffect
    {
        //75 % transparency (colormaps 561-816 in a .pl2)
        PctTransparency75,

        //50 % transparency (colormaps 305-560 in a .pl2)
        PctTransparency50,

        //25 % transparency (colormaps 49-304 in a .pl2)
        PctTransparency25,

        //Screen (colormaps 817-1072 in a .pl2)
        Screen,

        //luminance (colormaps 1073-1328 in a .pl2)
        Luminance,

        //bright alpha blending (colormaps 1457-1712 in a .pl2)
        BringAlphaBlending
    }
}
