using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IScene : IDisposable
    {
        void Update(long ms);
        void Render();
    }
}
