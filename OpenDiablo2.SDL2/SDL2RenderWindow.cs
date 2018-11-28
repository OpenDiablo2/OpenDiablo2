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
        private readonly GlobalConfiguration globalConfig;
        private readonly IGameState gameState;
        private readonly Func<IMapEngine> getMapEngine;

        public SDL2RenderWindow(
            GlobalConfiguration globalConfig,
            IMPQProvider mpqProvider,
            IPaletteProvider paletteProvider,
            IResourceManager resourceManager,
            IGameState gameState,
            Func<IMapEngine> getMapEngine
            )
        {
            this.globalConfig = globalConfig;
            this.mpqProvider = mpqProvider;
            this.paletteProvider = paletteProvider;
            this.resourceManager = resourceManager;
            this.gameState = gameState;
            this.getMapEngine = getMapEngine;

            SDL.SDL_Init(SDL.SDL_INIT_EVERYTHING);
            if (SDL.SDL_SetHint(SDL.SDL_HINT_RENDER_SCALE_QUALITY, "0") == SDL.SDL_bool.SDL_FALSE)
                throw new ApplicationException($"Unable to Init hinting: {SDL.SDL_GetError()}");

            window = SDL.SDL_CreateWindow("OpenDiablo2", SDL.SDL_WINDOWPOS_UNDEFINED, SDL.SDL_WINDOWPOS_UNDEFINED, 800, 600, SDL.SDL_WindowFlags.SDL_WINDOW_SHOWN | SDL.SDL_WindowFlags.SDL_WINDOW_ALLOW_HIGHDPI);
            if (window == IntPtr.Zero)
                throw new ApplicationException($"Unable to create SDL Window: {SDL.SDL_GetError()}");

            renderer = SDL.SDL_CreateRenderer(window, -1, SDL.SDL_RendererFlags.SDL_RENDERER_SOFTWARE);
            if (renderer == IntPtr.Zero)
                throw new ApplicationException($"Unable to create SDL Window: {SDL.SDL_GetError()}");


            SDL.SDL_SetRenderDrawBlendMode(renderer, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            
            SDL.SDL_ShowCursor(globalConfig.MouseMode == eMouseMode.Hardware ? 1 : 0);

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

        public unsafe void DrawMapCell(int xCell, int yCell, int xPixel, int yPixel, MPQDS1 mapData, int main_index, int sub_index, Palette palette, int orientation)
        {
            var tiles = mapData.LookupTable.Where(x =>
                 x.MainIndex == main_index &&
                 x.SubIndex == sub_index &&
                 (orientation == -1 || x.Orientation == orientation)).Select(x => x.TileRef);

            if (!tiles.Any())
                throw new ApplicationException("Invalid tile id found!");


            // TODO: This isn't good.. should be remembered in the map engine layer
            MPQDT1Tile tile = null;
            if (tiles.Count() > 0)
            {
                var totalRarity = tiles.Sum(q => q.RarityOrFrameIndex);
                var random = new Random(gameState.Seed + xCell + (mapData.Width * yCell));
                var x = random.Next(totalRarity);
                var z = 0;
                foreach (var t in tiles)
                {
                    z += t.RarityOrFrameIndex;
                    if (x <= z)
                    {
                        tile = t;
                        break;
                    }
                }
            }
            else tile = tiles.First();

            // This WILL happen to you
            if (tile.Width == 0 || tile.Height == 0)
                return;

            var mapCellInfo = getMapEngine().GetMapCellInfo(mapData.Id, tile.Id); ;
            if (mapCellInfo != null)
            {
                var xd = new SDL.SDL_Rect { x = xPixel - mapCellInfo.OffX, y = yPixel - mapCellInfo.OffY, w = mapCellInfo.FrameWidth, h = mapCellInfo.FrameHeight };
                var xs = mapCellInfo.Rect.ToSDL2Rect();
                SDL.SDL_RenderCopy(renderer, ((SDL2Texture)mapCellInfo.Texture).Pointer, ref xs, ref xd);
                return;
            }


            var minX = tile.Blocks.Min(x => x.PositionX);
            var minY = tile.Blocks.Min(x => x.PositionY);
            var maxX = tile.Blocks.Max(x => x.PositionX + 32);
            var maxY = tile.Blocks.Max(x => x.PositionY + 32);
            var diffX = maxX - minX;
            var diffY = maxY - minY;

            var offX = -minX;
            var offy = -minY;

            var frameSize = new Size(diffX, Math.Abs(diffY));

            var srcRect = new SDL.SDL_Rect { x = 0, y = 0, w = frameSize.Width, h = Math.Abs(frameSize.Height) };
            var frameSizeMax = diffX * Math.Abs(diffY);


            var texId = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, frameSize.Width, frameSize.Height);
            SDL.SDL_SetTextureBlendMode(texId, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);


            if (SDL.SDL_LockTexture(texId, IntPtr.Zero, out IntPtr pixels, out int pitch) != 0)
            {
                log.Error("Could not lock texture for map rendering");
                return;
            }
            try
            {
                UInt32* data = (UInt32*)pixels;

                var pitchChange = (pitch / 4);

                for (var i = 0; i < frameSize.Height * pitchChange; i++)
                    data[i] = 0x0;


                foreach (var block in tile.Blocks)
                {
                    var index = block.PositionX + offX + ((block.PositionY + offy) * pitchChange);
                    var xx = 0;
                    var yy = 0;
                    foreach (var colorIndex in block.PixelData)
                    {
                        try
                        {
                            if (colorIndex == 0)
                                continue;
                            var color = palette.Colors[colorIndex];

                            if (color > 0)
                                data[index] = color;

                        }
                        finally
                        {
                            index++;
                            xx++;
                            if (xx == 32)
                            {
                                index -= 32;
                                index += pitchChange;
                                xx = 0;
                                yy++;
                            }
                        }
                    }
                }
            }
            finally
            {
                SDL.SDL_UnlockTexture(texId);
            }

            var lookup = new MapCellInfo
            {
                FrameHeight = frameSize.Height,
                FrameWidth = frameSize.Width,
                OffX = offX,
                OffY = offy,
                Rect = srcRect.ToRectangle(),
                TileId = tile.Id,
                Texture = new SDL2Texture { Pointer = texId }
            };

            getMapEngine().SetMapCellInfo(mapData.Id, lookup);

            var dr = new SDL.SDL_Rect { x = xPixel - lookup.OffX, y = yPixel - lookup.OffY, w = lookup.FrameWidth, h = lookup.FrameHeight };
            SDL.SDL_RenderCopy(renderer, texId, ref srcRect, ref dr);
        }

        public unsafe IMouseCursor LoadCursor(ISprite sprite, int frame, Point hotspot)
        {
            if (globalConfig.MouseMode != eMouseMode.Hardware)
                throw new ApplicationException("Tried to set a hardware cursor, but we are using software cursors!");

            var multiple = globalConfig.HardwareMouseScale;
            var spr = sprite as SDL2Sprite;
            var surface = SDL.SDL_CreateRGBSurface(0, spr.FrameSize.Width * multiple, spr.FrameSize.Height * multiple, 32, 0xFF0000, 0xFF00, 0xFF, 0xFF000000);

            var pixels = (UInt32*)((SDL.SDL_Surface*)surface)->pixels;
            for (var y = 0; y < (spr.FrameSize.Height * multiple) - 1; y ++)
                for (var x = 0; x < (spr.FrameSize.Width * multiple) - 1; x ++)
                {
                    pixels[x + (y * spr.FrameSize.Width * multiple)] = spr.source.Frames[frame].GetColor(x / multiple, y / multiple, sprite.CurrentPalette);
                }

            var cursor = SDL.SDL_CreateColorCursor(surface, hotspot.X, hotspot.Y);
            if (cursor == IntPtr.Zero)
                throw new ApplicationException($"Unable to set the cursor cursor: {SDL.SDL_GetError()}"); // TODO: Is this supported everywhere? May need to still support software cursors.
            return new SDL2MouseCursor { Surface = cursor };
        }

        public void SetCursor(IMouseCursor mouseCursor)
        {
            if (globalConfig.MouseMode != eMouseMode.Hardware)
                throw new ApplicationException("Tried to set a hardware cursor, but we are using software cursors!");

            SDL.SDL_SetCursor((mouseCursor as SDL2MouseCursor).Surface);
        }
    }
}
