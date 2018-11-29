using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IInventoryPanel : IDisposable
    {
        void Render();
        void Update();
    }
}
