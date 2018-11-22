using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using SDL2;
using System.Runtime.InteropServices;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.SDL2_
{
    internal sealed class SDL2Sprite : ISprite
    {
        public Point Location { get; set; } = new Point();
        public Size FrameSize { get; set; } = new Size();
        public int Frame { get; set; }
        public int TotalFrames { get; internal set; }

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
        internal readonly ImageSet source;
        private readonly IntPtr renderer;
        internal IntPtr[] textures = new IntPtr[0];

        public SDL2Sprite(ImageSet source, IntPtr renderer)
        {
            this.source = source;
            this.renderer = renderer;


            TotalFrames = source.Frames.Count();
            FrameSize = new Size(Pow2((int)source.Frames.Max(x => x.Width)), Pow2((int)source.Frames.Max(x => x.Height)));

        }

        internal Point GetRenderPoint()
            => new Point(
                Location.X + source.Frames[Frame].OffsetX,
                (Location.Y - FrameSize.Height) + source.Frames[Frame].OffsetY
            );

        public Size LocalFrameSize => new Size((int)source.Frames[Frame].Width, (int)source.Frames[Frame].Height);

        private void UpdateTextureData()
        {
            foreach (var texture in textures)
            {
                SDL.SDL_DestroyTexture(texture);
            }

            textures = new IntPtr[TotalFrames];

            for (var i = 0; i < source.Frames.Count(); i++)
                textures[i] = LoadFrame(source.Frames[i], renderer);


        }

        // TODO: Less dumb color correction
        private Color AdjustColor(Color source)
        => Color.FromArgb(
                source.A,
                (byte)Math.Min((float)source.R * 1.2, 255),
                (byte)Math.Min((float)source.G * 1.2, 255),
                (byte)Math.Min((float)source.B * 1.2, 255)
            );

        private IntPtr LoadFrame(ImageFrame frame, IntPtr renderer)
        {
            var texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_TARGET, Pow2(FrameSize.Width), Pow2(FrameSize.Height));

            if (texture == IntPtr.Zero)
                throw new ApplicationException("Unaple to initialize texture.");

            SDL.SDL_SetTextureBlendMode(texture, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            SDL.SDL_SetRenderTarget(renderer, texture);
            SDL.SDL_SetRenderDrawColor(renderer, 0, 0, 0, 0);
            SDL.SDL_RenderFillRect(renderer, IntPtr.Zero);
            SDL.SDL_SetRenderTarget(renderer, IntPtr.Zero);

            var binaryData = new UInt32[frame.Width * frame.Height];
            for (int y = 0; y < frame.Height; y++)
                for (int x = 0; x < frame.Width; x++)
                {
                    var col = AdjustColor(frame.GetColor(x, y, CurrentPalette));
                    binaryData[x + y * frame.Width] = (uint)col.ToArgb();
                }
            var rect = new SDL.SDL_Rect { x = 0, y = FrameSize.Height - (int)frame.Height, w = (int)frame.Width, h = (int)frame.Height };
            GCHandle pinnedArray = GCHandle.Alloc(binaryData, GCHandleType.Pinned);
            SDL.SDL_UpdateTexture(texture, ref rect, pinnedArray.AddrOfPinnedObject(), (int)frame.Width * 4);
            pinnedArray.Free();

            return texture;

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
            foreach (var texture in textures)
            {
                SDL.SDL_DestroyTexture(texture);

            }
        }
    }
}
