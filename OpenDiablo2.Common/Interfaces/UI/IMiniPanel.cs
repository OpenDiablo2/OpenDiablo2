using OpenDiablo2.Common.Enums;
using System;

namespace OpenDiablo2.Common.Interfaces
{
    public delegate void OnPanelToggledEvent(IPanel panel);

    public interface IMiniPanel : IDisposable
    {
        event OnPanelToggledEvent OnPanelToggled;

        IPanel GetPanel(ePanelType panelType);
        bool IsMouseOver();
        void UpdatePanelLocation();
        void OnMenuToggle(bool isToggled);
        void Render();
        void Update();
    }
}
