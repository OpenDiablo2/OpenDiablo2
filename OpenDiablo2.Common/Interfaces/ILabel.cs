using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ILabel : IDisposable
    {
        string Text { get; set; }
        Point Location { get; set; }
        Size TextArea { get; set; }
        Color TextColor { get; set; }
    }
}
