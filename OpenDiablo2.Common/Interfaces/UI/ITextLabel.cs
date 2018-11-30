using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ITextLabel
    {
        string Text { get; set; }
        Point Position { get; set; }
        IFont Font { get; set; }
    }
}
