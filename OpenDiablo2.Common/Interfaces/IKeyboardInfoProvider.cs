namespace OpenDiablo2.Common.Interfaces
{
    public delegate void OnKeyPressed(char charcode);

    public interface IKeyboardInfoProvider
    {
        OnKeyPressed KeyPressCallback { get; set; }
        bool KeyIsPressed(int scancode);
    }
}
