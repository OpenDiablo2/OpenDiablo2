using System;
using System.Drawing;
using System.Linq;
using System.Runtime.InteropServices;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2RenderWindow : IRenderWindow, IMouseInfoProvider, IKeyboardInfoProvider
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private IntPtr window, renderer;
        private bool fullscreen;

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
        private readonly Func<IMapEngine> getMapEngine;

        public SDL2RenderWindow(
            GlobalConfiguration globalConfig,
            IMPQProvider mpqProvider,
            IPaletteProvider paletteProvider,
            IResourceManager resourceManager,
            Func<IGameState> getGameState,
            Func<IMapEngine> getMapEngine
            )
        {
            this.globalConfig = globalConfig;
            this.mpqProvider = mpqProvider;
            this.paletteProvider = paletteProvider;
            this.resourceManager = resourceManager;
            this.getGameState = getGameState;
            this.getMapEngine = getMapEngine;
            this.fullscreen = globalConfig.FullScreen;

            SDL.SDL_Init(SDL.SDL_INIT_EVERYTHING);
            if (SDL.SDL_SetHint(SDL.SDL_HINT_RENDER_SCALE_QUALITY, "0") == SDL.SDL_bool.SDL_FALSE)
                throw new ApplicationException($"Unable to Init hinting: {SDL.SDL_GetError()}");

            window = SDL.SDL_CreateWindow("OpenDiablo2", SDL.SDL_WINDOWPOS_UNDEFINED, SDL.SDL_WINDOWPOS_UNDEFINED, 800, 600,
                SDL.SDL_WindowFlags.SDL_WINDOW_SHOWN | (fullscreen ? SDL.SDL_WindowFlags.SDL_WINDOW_FULLSCREEN : 0));
            if (window == IntPtr.Zero)
                throw new ApplicationException($"Unable to create SDL Window: {SDL.SDL_GetError()}");

            renderer = SDL.SDL_CreateRenderer(window, -1, SDL.SDL_RendererFlags.SDL_RENDERER_ACCELERATED | SDL.SDL_RendererFlags.SDL_RENDERER_PRESENTVSYNC);
            if (renderer == IntPtr.Zero)
                throw new ApplicationException($"Unable to create SDL Window: {SDL.SDL_GetError()}");


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

        public unsafe MapCellInfo CacheMapCell(MPQDT1Tile mapCell)
        {
            var minX = mapCell.Blocks.Min(x => x.PositionX);
            var minY = mapCell.Blocks.Min(x => x.PositionY);
            var maxX = mapCell.Blocks.Max(x => x.PositionX + 32);
            var maxY = mapCell.Blocks.Max(x => x.PositionY + 32);
            var diffX = maxX - minX;
            var diffY = maxY - minY;

            var offX = -minX;
            var offY = -minY;

            var frameSize = new Size(diffX, Math.Abs(diffY));

            var srcRect = new SDL.SDL_Rect { x = 0, y = 0, w = frameSize.Width, h = Math.Abs(frameSize.Height) };
            var frameSizeMax = diffX * Math.Abs(diffY);


            var texId = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_STREAMING, frameSize.Width, frameSize.Height);
            SDL.SDL_SetTextureBlendMode(texId, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);

            if (SDL.SDL_LockTexture(texId, IntPtr.Zero, out IntPtr pixels, out int pitch) != 0)
                throw new ApplicationException("Could not lock texture for map rendering");

            try
            {
                UInt32* data = (UInt32*)pixels;

                var pitchChange = (pitch / 4);

                for (var i = 0; i < frameSize.Height * pitchChange; i++)
                    data[i] = 0x0;


                foreach (var block in mapCell.Blocks)
                {
                    var index = block.PositionX + offX + ((block.PositionY + offY) * pitchChange);
                    var xx = 0;
                    var yy = 0;
                    foreach (var colorIndex in block.PixelData)
                    {
                        try
                        {
                            if (colorIndex == 0)
                                continue;
                            var color = getGameState().CurrentPalette.Colors[colorIndex];

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

            return new MapCellInfo
            {
                Tile = mapCell,
                FrameHeight = frameSize.Height,
                FrameWidth = frameSize.Width,
                OffX = offX,
                OffY = offY,
                Rect = srcRect.ToRectangle(),
                Texture = new SDL2Texture { Pointer = texId }
            };
        }

        public void DrawMapCell(MapCellInfo mapCellInfo, int xPixel, int yPixel)
        {
            var srcRect = new SDL.SDL_Rect { x = 0, y = 0, w = mapCellInfo.FrameWidth, h = Math.Abs(mapCellInfo.FrameHeight) };
            var destRect = new SDL.SDL_Rect { x = xPixel - mapCellInfo.OffX, y = yPixel - mapCellInfo.OffY, w = mapCellInfo.FrameWidth, h = mapCellInfo.FrameHeight };
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
                SDL.SDL_RenderCopy(renderer, (sprite as SDL2Sprite).texture, IntPtr.Zero, IntPtr.Zero);
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
                throw new ApplicationException($"Unable to set the cursor cursor: {SDL.SDL_GetError()}"); // TODO: Is this supported everywhere? May need to still support software cursors.
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

    }
}
