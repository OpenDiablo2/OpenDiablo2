using System;
using System.Drawing;
using OpenDiablo2.Common.Interfaces;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2MouseCursor : IMouseCursor
    {
        public IntPtr HWSurface { get; set; }
        public IntPtr SWTexture { get; set; }
        public Size ImageSize { get; set; }
        public Point Hotspot { get; set; }
    }
}
