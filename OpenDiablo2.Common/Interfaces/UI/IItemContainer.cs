using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IItemContainer : IDisposable
    {
        void Render();
        void Update();
    }
}
