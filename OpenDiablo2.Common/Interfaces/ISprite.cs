using OpenDiablo2.Common.Models;
using System;
using System.Drawing;

namespace OpenDiablo2.Common.Interfaces
{
    public interface ISprite : IDisposable
    {
        Point Location { get; set; }
        Size FrameSize { get; set; }
        Size LocalFrameSize { get; }
        int Frame { get; set; }
        int TotalFrames { get; }
        Palette CurrentPalette { get; set; }
        bool Blend { get; set; }
    }
}
