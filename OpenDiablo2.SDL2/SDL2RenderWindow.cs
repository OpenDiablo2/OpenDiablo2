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

namespace OpenDiablo2.SDL2_
{
    public sealed class SDL2RenderWindow : IRenderWindow, IRenderTarget
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);
        private IntPtr window, renderer;
        private bool running;
        public bool IsRunning => running;
        private readonly ILifetimeScope lifetimeScope;

        public SDL2RenderWindow(ILifetimeScope lifetimeScope)
        {
            this.lifetimeScope = lifetimeScope;

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

            running = true;

        }
        public void Dispose()
        {
            SDL.SDL_DestroyRenderer(renderer);
            SDL.SDL_DestroyWindow(window);
            SDL.SDL_Quit();
        }

        public void Clear()
        {
            SDL.SDL_RenderClear(renderer);

        }

        public void Sync()
        {
            SDL.SDL_RenderPresent(renderer);
        }

        public void Update()
        {
            while (SDL.SDL_PollEvent(out SDL.SDL_Event evt) != 0)
            {
                if (evt.type == SDL.SDL_EventType.SDL_QUIT)
                    running = false;
            }
        }


        public void Draw(ISprite sprite)
        {
            var spr = sprite as SDL2Sprite;
            var loc = spr.GetRenderPoint();

            var destRect = new SDL.SDL_Rect
            {
                x = loc.X,
                y = loc.Y,
                w = spr.FrameSize.Width,
                h = spr.FrameSize.Height
            };
            SDL.SDL_RenderCopy(renderer, spr.textures[spr.Frame], IntPtr.Zero, ref destRect);

        }

        public ISprite LoadSprite(ImageSet source)
            => new SDL2Sprite(source, renderer);
    }
}
