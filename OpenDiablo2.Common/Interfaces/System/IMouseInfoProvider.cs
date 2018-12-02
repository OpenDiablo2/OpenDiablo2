namespace OpenDiablo2.Common.Interfaces
{
    public interface IMouseInfoProvider
    {
        int MouseX { get; }
        int MouseY { get; }
        bool LeftMouseDown { get; }
        bool LeftMousePressed { get; }
        bool RightMouseDown { get; }
        bool ReserveMouse { get; set; }
    }
}
