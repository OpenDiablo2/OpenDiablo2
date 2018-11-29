using System;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2MouseCursor : IMouseCursor
    {
        public IntPtr Surface { get; set; }
    }
}
