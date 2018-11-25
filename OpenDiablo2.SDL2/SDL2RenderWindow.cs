using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using SDL2;
using System.IO;
using System.Drawing;
using OpenDiablo2.Common.Models;
using Autofac;
using System.Runtime.InteropServices;
using OpenDiablo2.Common.Enums;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2RenderWindow : IRenderWindow, IRenderTarget, IMouseInfoProvider, IKeyboardInfoProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private IntPtr window, renderer;
        public bool IsRunning { get; private set; }

        public int MouseX { get; internal set; } = 0;
        public int MouseY { get; internal set; } = 0;
        public bool LeftMouseDown { get; internal set; } = false;
        public bool RightMouseDown { get; internal set; } = false;
        public bool ReserveMouse { get; set; } = false;

        public OnKeyPressed KeyPressCallback { get; set; }

        private readonly IMPQProvider mpqProvider;
        private readonly IPaletteProvider paletteProvider;
        private readonly IResourceManager resourceManager;

        private IntPtr cellTexture;

        public SDL2RenderWindow(
            IMPQProvider mpqProvider,
            IPaletteProvider paletteProvider,
            IResourceManager resourceManager
            )
        {
            this.mpqProvider = mpqProvider;
            this.paletteProvider = paletteProvider;
            this.resourceManager = resourceManager;

            SDL.SDL_Init(SDL.SDL_INIT_EVERYTHING);
            if (SDL.SDL_SetHint(SDL.SDL_HINT_RENDER_SCALE_QUALITY, "0") == SDL.SDL_bool.SDL_FALSE)
                throw new ApplicationException($"Unable to Init hinting: {SDL.SDL_GetError()}");

            window = SDL.SDL_CreateWindow("OpenDiablo2", SDL.SDL_WINDOWPOS_UNDEFINED, SDL.SDL_WINDOWPOS_UNDEFINED, 800, 600, SDL.SDL_WindowFlags.SDL_WINDOW_SHOWN);
            if (window == IntPtr.Zero)
                throw new ApplicationException($"Unable to create SDL Window: {SDL.SDL_GetError()}");

            renderer = SDL.SDL_CreateRenderer(window, -1, SDL.SDL_RendererFlags.SDL_RENDERER_ACCELERATED);
            if (renderer == IntPtr.Zero)
                throw new ApplicationException($"Unable to create SDL Window: {SDL.SDL_GetError()}");


            SDL.SDL_SetRenderDrawBlendMode(renderer, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            SDL.SDL_ShowCursor(0);

            cellTexture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, 256, 256);
            SDL.SDL_SetTextureBlendMode(cellTexture, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);


            IsRunning = true;

        }

        public void Quit() => IsRunning = false;

        public void Dispose()
        {
            SDL.SDL_DestroyRenderer(renderer);
            SDL.SDL_DestroyWindow(window);
            SDL.SDL_Quit();
        }

        public unsafe bool KeyIsPressed(int scancode)
        {
            int numKeys;
            byte* keys = (byte*)SDL.SDL_GetKeyboardState(out numKeys);
            return keys[scancode] > 0;

        }

        public void Clear()
        {
            SDL.SDL_SetRenderTarget(renderer, IntPtr.Zero);
            SDL.SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255);
            SDL.SDL_RenderClear(renderer);

        }

        public void Sync()
        {
            SDL.SDL_RenderPresent(renderer);
        }

        public unsafe void Update()
        {
            while (SDL.SDL_PollEvent(out SDL.SDL_Event evt) != 0)
            {
                if (evt.type == SDL.SDL_EventType.SDL_MOUSEMOTION)
                {
                    MouseX = evt.motion.x;
                    MouseY = evt.motion.y;
                    continue;
                }

                else if (evt.type == SDL.SDL_EventType.SDL_MOUSEBUTTONDOWN)
                {
                    switch ((uint)evt.button.button)
                    {
                        case SDL.SDL_BUTTON_LEFT:
                            LeftMouseDown = true;
                            break;
                        case SDL.SDL_BUTTON_RIGHT:
                            RightMouseDown = true;
                            break;
                    }
                }

                else if (evt.type == SDL.SDL_EventType.SDL_MOUSEBUTTONUP)
                {
                    switch ((uint)evt.button.button)
                    {
                        case SDL.SDL_BUTTON_LEFT:
                            LeftMouseDown = false;
                            break;
                        case SDL.SDL_BUTTON_RIGHT:
                            RightMouseDown = false;
                            break;
                    }
                }
                else if (evt.type == SDL.SDL_EventType.SDL_KEYDOWN)
                {
                    if (evt.key.keysym.sym == SDL.SDL_Keycode.SDLK_BACKSPACE && KeyPressCallback != null)
                        KeyPressCallback('\b');
                }
                else if (evt.type == SDL.SDL_EventType.SDL_TEXTINPUT)
                {
                    KeyPressCallback?.Invoke(Marshal.PtrToStringAnsi((IntPtr)evt.text.text)[0]);
                    continue;
                }

                else if (evt.type == SDL.SDL_EventType.SDL_QUIT)
                {
                    IsRunning = false;
                    continue;
                }
            }
        }

        public void Draw(ISprite sprite, Point location)
        {
            sprite.Location = location;
            Draw(sprite);
        }

        public void Draw(ISprite sprite, int frame, Point location)
        {
            sprite.Location = location;
            sprite.Frame = frame;
            Draw(sprite);
        }

        public void Draw(ISprite sprite, int frame)
        {
            sprite.Frame = frame;
            Draw(sprite);
        }

        public void Draw(ISprite sprite)
        {
            var spr = sprite as SDL2Sprite;
            if (spr.texture == IntPtr.Zero)
                return;

            var loc = spr.GetRenderPoint();

            var destRect = new SDL.SDL_Rect
            {
                x = loc.X,
                y = loc.Y,
                w = spr.FrameSize.Width,
                h = spr.FrameSize.Height
            };
            SDL.SDL_RenderCopy(renderer, spr.texture, IntPtr.Zero, ref destRect);

        }

        public void Draw(ISprite sprite, int xSegments, int ySegments, int offset)
        {
            var spr = sprite as SDL2Sprite;
            var segSize = xSegments * ySegments;

            for (var y = 0; y < ySegments; y++)
            {
                for (var x = 0; x < xSegments; x++)
                {
                    var textureIndex = x + (y * xSegments) + (offset * xSegments * ySegments);
                    spr.Frame = Math.Min(spr.TotalFrames - 1, Math.Max(0, textureIndex));

                    var destRect = new SDL.SDL_Rect
                    {
                        x = sprite.Location.X + (x * 256),
                        y = sprite.Location.Y + (y * 256) - (int)(spr.FrameSize.Height - spr.source.Frames[textureIndex].Height),
                        w = spr.FrameSize.Width,
                        h = spr.FrameSize.Height
                    };
                    SDL.SDL_RenderCopy(renderer, spr.texture, IntPtr.Zero, ref destRect);
                }
            }
        }

        public ISprite LoadSprite(string resourcePath, string palette) => LoadSprite(resourcePath, palette, Point.Empty);
        public ISprite LoadSprite(string resourcePath, string palette, Point location)
        {
            var result = new SDL2Sprite(resourceManager.GetImageSet(resourcePath), renderer)
            {
                CurrentPalette = paletteProvider.PaletteTable[palette],
                Location = location
            };
            return result;
        }

        public IFont LoadFont(string resourcePath, string palette)
        {
            var result = new SDL2Font(resourceManager.GetMPQFont(resourcePath), renderer)
            {
                CurrentPalette = paletteProvider.PaletteTable[palette]
            };
            return result;
        }


        public ILabel CreateLabel(IFont font)
        {
            var result = new SDL2Label(font, renderer);
            return result;
        }

        public ILabel CreateLabel(IFont font, string text)
        {
            var result = CreateLabel(font);
            result.Text = text;
            return result;
        }

        public ILabel CreateLabel(IFont font, Point position, string text)
        {
            var result = new SDL2Label(font, renderer)
            {
                Text = text,
                Location = position
            };

            return result;
        }

        public void Draw(ILabel label)
        {
            var lbl = label as SDL2Label;
            var loc = lbl.Location;

            var destRect = new SDL.SDL_Rect
            {
                x = loc.X,
                y = loc.Y,
                w = lbl.textureSize.Width,
                h = lbl.textureSize.Height
            };

            SDL.SDL_RenderCopy(renderer, lbl.texture, IntPtr.Zero, ref destRect);
        }

        public unsafe void DrawMapCell(int xCell, int yCell, int xPixel, int yPixel, MPQDS1 mapData)
        {
            var floorLayer = mapData.FloorLayers.First();
            var floor = floorLayer.Props[xCell + (yCell * mapData.Width)];

            if (floor.Prop1 == 0)
                return;

            var palette = paletteProvider.PaletteTable[$"ACT{mapData.Act}"];
            var sub_index = floor.Prop2;
            var main_index = (floor.Prop3 >> 4) + ((floor.Prop4 & 0x03) << 4);

            MPQDT1Tile tile = null;
            for (int i = 0; i < mapData.DT1s.Count(); i++)
            {
                if (mapData.DT1s[i] == null)
                    continue;

                tile = mapData.DT1s[i].Tiles.FirstOrDefault(z => z.MainIndex == main_index && z.SubIndex == sub_index);
                if (tile != null)
                    break;
            }

            if (tile == null)
                throw new ApplicationException("Could not locate tile!");


            var frameSize = new Size(tile.Width, Math.Abs(tile.Height));
            var srcRect = new SDL.SDL_Rect { x = 0, y = 0, w = frameSize.Width, h = frameSize.Height };
            var frameSizeMax = frameSize.Width * frameSize.Height;
            SDL.SDL_LockTexture(cellTexture, ref srcRect, out IntPtr pixels, out int pitch);
            try
            {
                UInt32* data = (UInt32*)pixels;
                for (var i = 0; i < frameSizeMax; i++)
                    data[i] = 0x0;

                var pitchChange = (pitch / 4);

                foreach (var block in tile.Blocks)
                {
                    for (int yy = 0; yy < 32; yy++)
                    {
                        var index = block.PositionX + ((block.PositionY + yy) * pitchChange);

                        for (int xx = 0; xx < 32; xx++)
                        {
                            index++;

                            if (index > frameSizeMax)
                                continue;
                            if (index < 0)
                                continue;

                            var color = palette.Colors[block.PixelData[xx + (yy * 32)]];

                            if ((color & 0xFFFFFF) > 0)
                                data[index] = color;

                        }
                    }
                }
            }
            finally
            {
                SDL.SDL_UnlockTexture(cellTexture);
            }


            var dstRect = new SDL.SDL_Rect { x = xPixel, y = yPixel, w = frameSize.Width, h = frameSize.Height };
            SDL.SDL_RenderCopy(renderer, cellTexture, ref srcRect, ref dstRect);
        }

    }
}
