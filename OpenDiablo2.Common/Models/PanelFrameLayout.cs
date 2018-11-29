using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Models
{
    public class PanelFrameLayout
    {
        public ePanelFrameType panelFrameType { get; internal set; }

        public static Dictionary<ePanelFrameType, PanelFrameLayout> Values = new Dictionary<ePanelFrameType, PanelFrameLayout>
        {
            {ePanelFrameType.Left, new PanelFrameLayout { panelFrameType = ePanelFrameType.Left } },
            {ePanelFrameType.Right, new PanelFrameLayout { panelFrameType = ePanelFrameType.Right } }
        };
    }

}
