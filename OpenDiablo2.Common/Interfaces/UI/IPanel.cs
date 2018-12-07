using OpenDiablo2.Common.Enums;
using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IPanel : IDisposable
    {
        eButtonType PanelType { get; }
        ePanelFrameType FrameType { get; }
        void Render();
        void Update();
    }
}
