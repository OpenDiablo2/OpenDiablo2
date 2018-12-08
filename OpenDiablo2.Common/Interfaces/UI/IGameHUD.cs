namespace OpenDiablo2.Common.Interfaces
{
    public interface IGameHUD
    {
        bool IsRunningEnabled { get; }
        bool IsLeftPanelVisible { get; }
        bool IsRightPanelVisible { get; }

        bool IsMouseOver();
        void OpenPanel(IPanel panel);
        void OpenPanels(IPanel leftPanel, IPanel rightPanel);

        void Render();
        void Update();
    }
}