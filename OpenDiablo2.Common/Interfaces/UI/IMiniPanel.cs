using System;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IMiniPanel : IDisposable
    {
        bool IsLeftPanelVisible { get; }
        bool IsRightPanelVisible { get; }

        void OnMenuToggle(bool isToggled);
        void Render();
        void Update();
    }
}
