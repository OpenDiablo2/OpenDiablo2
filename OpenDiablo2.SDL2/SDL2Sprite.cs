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
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        internal readonly ImageSet source;
        private readonly IntPtr renderer;
        internal IntPtr texture = IntPtr.Zero;

        public Point Location { get; set; } = new Point();
        public Size FrameSize { get; set; } = new Size();


        private bool darken;
        public bool Darken
        {
            get => darken;
            set
            {
                if (darken == value)
                    return;
                darken = value;
                LoadFrame(frame);
            }
        }

        private int frame = -1;
        public int Frame
        {
            get => frame;
            set
            {
                if (frame == value && texture != IntPtr.Zero)
                    return;

                frame = Math.Max(0, Math.Min(value, TotalFrames));
                LoadFrame(frame);
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
                SDL.SDL_SetTextureBlendMode(texture, blend ? SDL.SDL_BlendMode.SDL_BLENDMODE_ADD : SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            }
        }

        private Palette palette;
        public Palette CurrentPalette
        {
            get => palette;
            set
            {
                palette = value;
                UpdateTextureData();
            }
        }


        public SDL2Sprite(ImageSet source, IntPtr renderer)
        {
            this.source = source;
            this.renderer = renderer;


            TotalFrames = source.Frames.Count();
            FrameSize = new Size(Pow2((int)source.Frames.Max(x => x.Width)), Pow2((int)source.Frames.Max(x => x.Height)));
        }

        
        internal Point GetRenderPoint()
        {
            return source == null
                ? Location
                : new Point(Location.X + source.Frames[Frame].OffsetX, (Location.Y - FrameSize.Height) + source.Frames[Frame].OffsetY);
        }

        public Size LocalFrameSize => new Size((int)source.Frames[Frame].Width, (int)source.Frames[Frame].Height);

        private void UpdateTextureData()
        {
            if (texture == IntPtr.Zero)
            {
                texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, Pow2(FrameSize.Width), Pow2(FrameSize.Height));

                if (texture == IntPtr.Zero)
                    throw new OpenDiablo2Exception("Unaple to initialize texture.");

                Frame = 0;
            }
        }

        private unsafe void LoadFrame(int index)
        {
            var frame = source.Frames[index];
            var fullRect = new SDL.SDL_Rect { x = 0, y = 0, w = FrameSize.Width, h = FrameSize.Height };
            SDL.SDL_SetTextureBlendMode(texture, blend ? SDL.SDL_BlendMode.SDL_BLENDMODE_ADD : SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

            SDL.SDL_LockTexture(texture, IntPtr.Zero, out IntPtr pixels, out int pitch);
            try
            {
                UInt32* data = (UInt32*)pixels;
                var frameOffset = FrameSize.Height - frame.Height;
                var frameWidth = FrameSize.Width;
                var frameHeight = FrameSize.Height;
                for (var y = 0; y < frameHeight; y++)
                {
                    for (int x = 0; x < frameWidth; x++)
                    {
                        if ((x >= frame.Width) || (y < frameOffset))
                        {
                            data[x + (y * (pitch / 4))] = 0;
                            continue;
                        }

                        var color = frame.GetColor(x, (int)(y - frameOffset), CurrentPalette);
                        if (darken)
                            color = ((color & 0xFF000000) > 0) ? (color >> 1) & 0xFF7F7F7F | 0xFF000000 : 0;
                        data[x + (y * (pitch / 4))] = color;
                    }
                }
            }
            finally
            {
                SDL.SDL_UnlockTexture(texture);
            }


        }

        private int Pow2(int val)
        {
            int result = 1;
            while (result < val)
                result *= 2;
            return result;
        }

        public void Dispose()
        {
            SDL.SDL_DestroyTexture(texture);
        }

    }
}
