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
using System.Runtime.InteropServices;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Exceptions;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Interfaces.Drawing;
using OpenDiablo2.Common.Models;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2RenderWindow : IRenderWindow, IMouseInfoProvider, IKeyboardInfoProvider
    {
        private readonly IntPtr window, renderer;
        private readonly bool fullscreen;

        public bool IsRunning { get; private set; }

        public int MouseX { get; internal set; } = 0;
        public int MouseY { get; internal set; } = 0;
        public bool LeftMouseDown { get; internal set; } = false;
        public bool LeftMousePressed { get; internal set; } = false;
        public bool RightMouseDown { get; internal set; } = false;
        public bool ReserveMouse { get; set; } = false;

        public OnKeyPressed KeyPressCallback { get; set; }

        private IMouseCursor mouseCursor = null;
        public IMouseCursor MouseCursor
        {
            get => mouseCursor;
            set
            {
                if (mouseCursor == value)
                    return;

                SetCursor(value);
            }
        }

        private readonly IMPQProvider mpqProvider;
        private readonly IPaletteProvider paletteProvider;
        private readonly IResourceManager resourceManager;
        private readonly GlobalConfiguration globalConfig;
        private readonly Func<IGameState> getGameState;
        private readonly Func<IMapRenderer> getMapRenderer;

        public SDL2RenderWindow(
            GlobalConfiguration globalConfig,
            IMPQProvider mpqProvider,
            IPaletteProvider paletteProvider,
            IResourceManager resourceManager,
            Func<IGameState> getGameState,
            Func<IMapRenderer> getMapEngine
            )
        {
            this.globalConfig = globalConfig;
            this.mpqProvider = mpqProvider;
            this.paletteProvider = paletteProvider;
            this.resourceManager = resourceManager;
            this.getGameState = getGameState;
            this.getMapRenderer = getMapEngine;
            this.fullscreen = globalConfig.FullScreen;

            SDL.SDL_Init(SDL.SDL_INIT_EVERYTHING);
            if (SDL.SDL_SetHint(SDL.SDL_HINT_RENDER_SCALE_QUALITY, "0") == SDL.SDL_bool.SDL_FALSE)
                throw new OpenDiablo2Exception($"Unable to Init hinting: {SDL.SDL_GetError()}");

            window = SDL.SDL_CreateWindow("OpenDiablo2", SDL.SDL_WINDOWPOS_UNDEFINED, SDL.SDL_WINDOWPOS_UNDEFINED, 800, 600,
                SDL.SDL_WindowFlags.SDL_WINDOW_SHOWN | (fullscreen ? SDL.SDL_WindowFlags.SDL_WINDOW_FULLSCREEN : 0));
            if (window == IntPtr.Zero)
                throw new OpenDiablo2Exception($"Unable to create SDL Window: {SDL.SDL_GetError()}");

            renderer = SDL.SDL_CreateRenderer(window, -1, SDL.SDL_RendererFlags.SDL_RENDERER_ACCELERATED | SDL.SDL_RendererFlags.SDL_RENDERER_PRESENTVSYNC);
            if (renderer == IntPtr.Zero)
                throw new OpenDiablo2Exception($"Unable to create SDL Window: {SDL.SDL_GetError()}");


            SDL.SDL_SetRenderDrawBlendMode(renderer, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            SDL.SDL_GL_SetSwapInterval(1);
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
            byte* keys = (byte*)SDL.SDL_GetKeyboardState(out int numKeys);
            return keys[scancode] > 0;

        }

        public void Clear()
        {
            SDL.SDL_SetRenderTarget(renderer, IntPtr.Zero);
            SDL.SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255);
            SDL.SDL_RenderClear(renderer);

        }

        public unsafe void Sync()
        {
            if (globalConfig.MouseMode == eMouseMode.Software)
            {
                var cursor = mouseCursor as SDL2MouseCursor;
                var texture = cursor.SWTexture;

                var srcRect = new SDL.SDL_Rect
                {
                    x = 0,
                    y = 0,
                    w = cursor.ImageSize.Width,
                    h = cursor.ImageSize.Height
                };

                var destRect = new SDL.SDL_Rect
                {
                    x = MouseX - cursor.Hotspot.X,
                    y = MouseY - cursor.Hotspot.Y,
                    w = cursor.ImageSize.Width,
                    h = cursor.ImageSize.Height
                };

                SDL.SDL_RenderCopy(renderer, texture, ref srcRect, ref destRect);
            }


            SDL.SDL_RenderPresent(renderer);
        }

        public unsafe void Update()
        {
            LeftMousePressed = false;

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
                            LeftMousePressed = true; // Cannot find a better to handle a single press
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
                    /* NOTE: This absolutely trashes rendering for some reason...
                    if (evt.key.keysym.mod.HasFlag(SDL.SDL_Keymod.KMOD_RALT) && evt.key.keysym.scancode.HasFlag(SDL.SDL_Scancode.SDL_SCANCODE_RETURN))
                    {
                        fullscreen = !fullscreen;
                        SDL.SDL_SetWindowFullscreen(window, (uint)(fullscreen ? SDL.SDL_WindowFlags.SDL_WINDOW_FULLSCREEN : 0));
                    }
                    else*/
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
            if (spr.Texture == IntPtr.Zero)
                return;

            var loc = spr.GetRenderPoint();

            var destRect = new SDL.SDL_Rect
            {
                x = loc.X,
                y = loc.Y,
                w = spr.FrameSize.Width,
                h = spr.FrameSize.Height
            };
            SDL.SDL_RenderCopy(renderer, spr.Texture, IntPtr.Zero, ref destRect);

        }

        public void Draw(ISprite sprite, int xSegments, int ySegments, int offset)
        {
            var spr = sprite as SDL2Sprite;

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
                    SDL.SDL_RenderCopy(renderer, spr.Texture, IntPtr.Zero, ref destRect);
                }
            }
        }

        public ISprite LoadSprite(string resourcePath, string palette, bool cacheFrames = false) => LoadSprite(resourcePath, palette, Point.Empty, cacheFrames);
        public ISprite LoadSprite(string resourcePath, string palette, Point location, bool cacheFrames = false)
        {
            var result = new SDL2Sprite(resourceManager.GetImageSet(resourcePath), renderer, cacheFrames)
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

        public unsafe MapCellInfo CacheMapCell(MPQDT1Tile mapCell, eRenderCellType cellType)
        {
            var minX = mapCell.Blocks.Min(x => x.PositionX);
            var minY = mapCell.Blocks.Min(x => x.PositionY);
            var maxX = mapCell.Blocks.Max(x => x.PositionX + 32);
            var maxY = mapCell.Blocks.Max(x => x.PositionY + 32);
            var frameWidth = maxX - minX;
            var frameHeight = maxY - minY;

            var dx = minX;
            var dy = minY;

            dx -= 80;

            
            var srcRect = new SDL.SDL_Rect { x = dx, y = dy, w = frameWidth, h = frameHeight };

            var texId = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, frameWidth, frameHeight);

            if (texId == IntPtr.Zero)
                throw new OpenDiablo2Exception("Could not create texture");

            SDL.SDL_SetTextureBlendMode(texId, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

            if (SDL.SDL_LockTexture(texId, IntPtr.Zero, out IntPtr pixels, out int pitch) != 0)
                throw new OpenDiablo2Exception("Could not lock texture for map rendering");
            
            try
            {
                UInt32* data = (UInt32*)pixels;

                var pitchChange = pitch / 4;
                var colors = getGameState().CurrentPalette.Colors;

                foreach (var block in mapCell.Blocks)
                {
                    var index = (block.PositionX - minX) + ((block.PositionY - minY) * pitchChange);
                    var xx = 0;
                    foreach (var colorIndex in block.PixelData)
                    {
                        try
                        {
                            if (colorIndex == 0)
                                continue;

                            var color = colors[colorIndex];

                            if (color > 0)
                                data[index] = color;

                        }
                        finally
                        {
                            index++;
                            xx++;
                            if (xx == 32)
                            {
                                index += pitchChange - 32;
                                xx = 0;
                            }
                        }
                    }
                }
            }
            finally
            {
                SDL.SDL_UnlockTexture(texId);
            }

            return new MapCellInfo
            {
                Tile = mapCell,
                FrameHeight = frameHeight,
                FrameWidth = frameWidth,
                Rect = srcRect.ToRectangle(),
                Texture = new SDL2Texture { Pointer = texId }
            };
        }

        public void DrawMapCell(MapCellInfo mapCellInfo, int xPixel, int yPixel)
        {
            var srcRect = new SDL.SDL_Rect { x = 0, y = 0, w = mapCellInfo.FrameWidth, h = Math.Abs(mapCellInfo.FrameHeight) };
            var destRect = new SDL.SDL_Rect { x = xPixel + mapCellInfo.Rect.X, y = yPixel + mapCellInfo.Rect.Y, w = mapCellInfo.FrameWidth, h = mapCellInfo.FrameHeight };
            SDL.SDL_RenderCopy(renderer, (mapCellInfo.Texture as SDL2Texture).Pointer, ref srcRect, ref destRect);
        }

        public unsafe IMouseCursor LoadCursor(ISprite sprite, int frame, Point hotspot)
        {
            if (globalConfig.MouseMode == eMouseMode.Software)
            {
                sprite.Frame = frame;

                var texId = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888,
                    (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_TARGET, sprite.LocalFrameSize.Width, sprite.LocalFrameSize.Height);
                SDL.SDL_SetTextureBlendMode(texId, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

                SDL.SDL_SetRenderTarget(renderer, texId);
                SDL.SDL_RenderCopy(renderer, (sprite as SDL2Sprite).Texture, IntPtr.Zero, IntPtr.Zero);
                SDL.SDL_SetRenderTarget(renderer, IntPtr.Zero);

                return new SDL2MouseCursor
                {
                    Hotspot = hotspot,
                    ImageSize = sprite.LocalFrameSize,
                    SWTexture = texId
                };
            }

            var multiple = globalConfig.HardwareMouseScale;
            var spr = sprite as SDL2Sprite;
            var surface = SDL.SDL_CreateRGBSurface(0, spr.LocalFrameSize.Width * multiple, spr.LocalFrameSize.Height * multiple, 32, 0xFF0000, 0xFF00, 0xFF, 0xFF000000);
            var yOffset = 0; //(spr.FrameSize.Height - spr.LocalFrameSize.Height);
            var XOffset = 0; //(spr.FrameSize.Width - spr.LocalFrameSize.Width);
            var pixels = (UInt32*)((SDL.SDL_Surface*)surface)->pixels;
            for (var y = 0; y < (spr.LocalFrameSize.Height * multiple) - 1; y++)
                for (var x = 0; x < (spr.LocalFrameSize.Width * multiple) - 1; x++)
                {
                    pixels[x + XOffset + ((y + yOffset) * spr.LocalFrameSize.Width * multiple)] = spr.source.Frames[frame].GetColor(x / multiple, y / multiple, sprite.CurrentPalette);
                }

            var cursor = SDL.SDL_CreateColorCursor(surface, hotspot.X * multiple, hotspot.Y * multiple);
            if (cursor == IntPtr.Zero)
                throw new OpenDiablo2Exception($"Unable to set the cursor cursor: {SDL.SDL_GetError()}"); // TODO: Is this supported everywhere? May need to still support software cursors.
            return new SDL2MouseCursor { HWSurface = cursor };
        }

        private void SetCursor(IMouseCursor mouseCursor)
        {
            this.mouseCursor = mouseCursor;

            if (globalConfig.MouseMode != eMouseMode.Hardware)
                return;

            SDL.SDL_SetCursor((mouseCursor as SDL2MouseCursor).HWSurface);
        }

        public uint GetTicks() => SDL.SDL_GetTicks();

        public ICharacterRenderer CreateCharacterRenderer()
            => new SDL2CharacterRenderer(this.renderer, resourceManager, paletteProvider, getGameState());
    }
}
