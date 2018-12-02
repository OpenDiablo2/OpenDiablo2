using OpenDiablo2.Common.Enums;
using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IPanelFrame : IDisposable
    {
        void Render();
        void Update();
    }
}
