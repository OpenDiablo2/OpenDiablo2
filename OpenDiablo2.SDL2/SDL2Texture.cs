using System;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2Texture : ITexture
    {
        public IntPtr Pointer { get; set; }
    }
}
