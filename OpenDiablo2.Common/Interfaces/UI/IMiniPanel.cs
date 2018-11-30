using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMiniPanel : IDisposable
    {
        void Render();
        void Update();
    }
}
