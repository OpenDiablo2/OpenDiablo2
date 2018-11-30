using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ICharacterPanel : IDisposable
    {
        void Render();
        void Update();
    }
}
