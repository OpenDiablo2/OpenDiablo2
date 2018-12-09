/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    internal sealed class SDL2Sprite : ISprite
    {
        internal readonly ImageSet source;
        private readonly IntPtr renderer;
        private readonly bool cacheFrames;

        private bool disposed = false;

        private IntPtr[] texture;
        private bool[] frameLoaded;

        public IntPtr Texture
        {
            get
            {
                if (!frameLoaded[TextureIndex])
                    LoadFrame();

                return texture[TextureIndex];
            }
        }

        public Point Location { get; set; }
        public Size FrameSize { get; set; }


        private bool darken;
        public bool Darken
        {
            get => darken;
            set
            {
                if (darken == value)
                    return;
                darken = value;
                ClearAllFrames();
            }
        }

        private int frame;
        public int Frame
        {
            get => frame;
            set
            {
                if (frame == value)
                    return;

                frame = Math.Max(0, Math.Min(value, TotalFrames));

                if (cacheFrames)
                    return;

                frameLoaded[TextureIndex] = false;
            }
        }
        public int TotalFrames { get; internal set; }

        private bool blend = false;
        public bool Blend
        {
            get => blend;
            set
            {
                blend = value;
                if (cacheFrames)
                {
                    for (var i = 0; i < TotalFrames; i++)
                        if (texture[i] != IntPtr.Zero)
                            SDL.SDL_SetTextureBlendMode(texture[i], blend ? SDL.SDL_BlendMode.SDL_BLENDMODE_ADD : SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
                }
                else
                    if (texture[TextureIndex] != IntPtr.Zero)
                        SDL.SDL_SetTextureBlendMode(texture[TextureIndex], blend ? SDL.SDL_BlendMode.SDL_BLENDMODE_ADD : SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

            }
        }

        private Palette palette;
        public Palette CurrentPalette
        {
            get => palette;
            set
            {
                palette = value;
                ClearAllFrames();
            }
        }

        private int TextureIndex => cacheFrames ? frame : 0;

        public SDL2Sprite(ImageSet source, IntPtr renderer, bool cacheFrames = false)
        {
            this.source = source;
            this.renderer = renderer;
            this.cacheFrames = cacheFrames;

            texture = new IntPtr[cacheFrames ? source.Frames.Count() : 1];
            frameLoaded = new bool[cacheFrames ? source.Frames.Count() : 1];
            TotalFrames = source.Frames.Count();

            ClearAllFrames();

            Location = Point.Empty;
            FrameSize = new Size(Pow2((int)source.Frames.Max(x => x.Width)), Pow2((int)source.Frames.Max(x => x.Height)));

            Frame = 0;
        }


        internal Point GetRenderPoint()
        {
            return source == null
                ? Location
                : new Point(Location.X + source.Frames[Frame].OffsetX, Location.Y - FrameSize.Height + source.Frames[Frame].OffsetY);
        }

        public Size LocalFrameSize => new Size((int)source.Frames[Frame].Width, (int)source.Frames[Frame].Height);

        private void UpdateTextureData()
        {
            Frame = 0;
        }

        private unsafe void LoadFrame()
        {
            if (texture[TextureIndex] == IntPtr.Zero)
            {
                texture[TextureIndex] = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, Pow2(FrameSize.Width), Pow2(FrameSize.Height));

                if (texture[TextureIndex] == IntPtr.Zero)
                    throw new OpenDiablo2Exception("Unaple to initialize texture.");
                SDL.SDL_SetTextureBlendMode(texture[TextureIndex], blend ? SDL.SDL_BlendMode.SDL_BLENDMODE_ADD : SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            }

            var sourceFrame = source.Frames[frame];

            SDL.SDL_LockTexture(texture[TextureIndex], IntPtr.Zero, out IntPtr pixels, out int pitch);
            try
            {
                UInt32* data = (UInt32*)pixels;
                var frameOffset = FrameSize.Height - sourceFrame.Height;
                var frameWidth = FrameSize.Width;
                var frameHeight = FrameSize.Height;
                for (int y = 0; y < frameHeight; y++)
                {
                    for (int x = 0; x < frameWidth; x++)
                    {
                        if ((x >= sourceFrame.Width) || (y < frameOffset))
                        {
                            data[x + (y * (pitch / 4))] = 0;
                            continue;
                        }

                        var color = sourceFrame.GetColor(x, (int)(y - frameOffset), CurrentPalette);
                        if (darken)
                            color = ((color & 0xFF000000) > 0) ? (color >> 1) & 0xFF7F7F7F | 0xFF000000 : 0;
                        data[x + (y * (pitch / 4))] = color;
                    }
                }
            }
            finally
            {
                SDL.SDL_UnlockTexture(texture[TextureIndex]);
            }

            frameLoaded[TextureIndex] = true;
        }

        private int Pow2(int val)
        {
            int result = 1;
            while (result < val)
                result *= 2;
            return result;
        }

        private void ClearAllFrames()
        {
            var framestoClear = cacheFrames ? TotalFrames : 1;
            for (int i = 0; i < framestoClear; i++)
                frameLoaded[i] = false;
        }

        public void Dispose()
        {
            if (disposed)
                return;

            var framestoClear = cacheFrames ? TotalFrames : 1;
            for (var i = 0; i < framestoClear; i++)
            {
                SDL.SDL_DestroyTexture(texture[i]);
                texture[i] = IntPtr.Zero;
            }

            texture = Array.Empty<IntPtr>();
            disposed = true;
        }

    }
}
