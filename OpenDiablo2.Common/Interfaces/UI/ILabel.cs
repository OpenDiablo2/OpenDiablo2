using System;
using System.Drawing;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ILabel : IDisposable
    {
        string Text { get; set; }
        Point Location { get; set; }
        Size TextArea { get; set; }
        Color TextColor { get; set; }
        int MaxWidth { get; set; }
        eTextAlign Alignment { get; set; }
    }
}
