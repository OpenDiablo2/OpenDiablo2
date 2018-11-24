using System;
using System.Drawing;
using System.Linq;
using System.Runtime.InteropServices;
using OpenDiablo2.Common.Enums;
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
        private Point globalFrameOffset = new Point(0, 0);

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

        public unsafe SDL2Sprite(IntPtr renderer, Palette palette, MPQDS1 mapData, int x, int y, eRenderCellType cellType)
        {
            this.renderer = renderer;
            // TODO: Cell types

            // Floor cell types
            // Todo: multiple floor layers
            var floorLayer = mapData.FloorLayers.First();
            var floor = floorLayer.Props[x + (y * mapData.Width)];

            if (floor.Prop1 == 0)
                return;

            var sub_index = floor.Prop2;
            var main_index = (floor.Prop3 / 16) + ((floor.Prop4 & 0x0F) * 16);
            var tile = mapData.DT1s[main_index].Tiles[sub_index];

            FrameSize = new Size(tile.Width, Math.Abs(tile.Height));
            TotalFrames = 1;
            frame = 0;
            IntPtr pixels;
            int pitch;

            texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, FrameSize.Width, FrameSize.Height);

            if (texture == IntPtr.Zero)
                throw new ApplicationException($"Unaple to initialize texture: {SDL.SDL_GetError()}");

            SDL.SDL_LockTexture(texture, IntPtr.Zero, out pixels, out pitch);
            try
            {
                UInt32* data = (UInt32*)pixels;
                for (var i = 0; i < FrameSize.Width * FrameSize.Height; i++)
                    data[i] = 0;


                foreach (var block in tile.Blocks)
                {
                    var px = block.PositionX;
                    var py = FrameSize.Height + block.PositionY;
                    for (int yy = 0; yy < 32; yy++)
                    {
                        for (int xx = 0; xx < 32; xx++)
                        {
                            data[px + xx + ((py + yy) * pitch / 4)] = palette.Colors[block.PixelData[xx + (yy * 32)]];
                        }
                    }
                }

            }
            finally
            {
                SDL.SDL_UnlockTexture(texture);
            }
        }

        internal Point GetRenderPoint()
        {
            return source == null
                ? globalFrameOffset
                : new Point(Location.X + source.Frames[Frame].OffsetX, (Location.Y - FrameSize.Height) + source.Frames[Frame].OffsetY);
        }

        public Size LocalFrameSize => new Size((int)source.Frames[Frame].Width, (int)source.Frames[Frame].Height);

        private void UpdateTextureData()
        {
            if (texture == IntPtr.Zero)
            {
                texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, Pow2(FrameSize.Width), Pow2(FrameSize.Height));

                if (texture == IntPtr.Zero)
                    throw new ApplicationException("Unaple to initialize texture.");

                Frame = 0;
            }
        }

        private unsafe void LoadFrame(int index)
        {
            var frame = source.Frames[index];

            IntPtr pixels;
            int pitch;
            var fullRect = new SDL.SDL_Rect { x = 0, y = 0, w = FrameSize.Width, h = FrameSize.Height };
            SDL.SDL_SetTextureBlendMode(texture, blend ? SDL.SDL_BlendMode.SDL_BLENDMODE_ADD : SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

            SDL.SDL_LockTexture(texture, IntPtr.Zero, out pixels, out pitch);
            try
            {
                UInt32* data = (UInt32*)pixels;
                var frameOffset = FrameSize.Height - frame.Height;
                for (var y = 0; y < FrameSize.Height; y++)
                {
                    for (int x = 0; x < FrameSize.Width; x++)
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
