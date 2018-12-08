using System;

namespace OpenDiablo2.Common.Interfaces
{
    public delegate void PanelSelectedEvent(IPanel panel);

    public interface IMiniPanel : IDisposable
    {
        event PanelSelectedEvent PanelSelected;

        bool IsMouseOver();
        void UpdatePanelLocation();
        void OnMenuToggle(bool isToggled);
        void Render();
        void Update();
    }
}
