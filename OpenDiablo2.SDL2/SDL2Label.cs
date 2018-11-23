using OpenDiablo2.Common.Interfaces;
using SDL2;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.SDL2_
{
    internal sealed class SDL2Label : ILabel
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly SDL2Font font;
        private readonly IntPtr renderer;
        internal IntPtr texture;
        internal Size textureSize = new Size();
        public Point Location { get; set; }
        public Size TextArea { get; set; } = new Size();

        private Color textColor = Color.White;
        public Color TextColor
        {
            get => textColor;
            set
            {
                textColor = value;
                RegenerateTexture();
            }
        }

        private string text;
        public string Text
        {
            get => text;
            set
            {
                text = value;
                RegenerateTexture();
            }
        }

        internal SDL2Label(IFont font, IntPtr renderer)
        {
            this.renderer = renderer;
            this.font = font as SDL2Font;
            this.texture = IntPtr.Zero;
        }

        internal Size CalculateSize()
        {
            int w = 0;
            int h = 0;
            foreach (var ch in text)
            {
                var metric = font.font.CharacterMetric[(byte)ch];
                w += metric.Width;
                h = Math.Max(Math.Max(h, metric.Height), font.sprite.FrameSize.Height);
            }

            return new Size(w, h);
        }

        internal int Pow2(int input)
        {
            var result = 1;
            while (result < input)
                result *= 2;
            return result;
        }

        internal void RegenerateTexture()
        {
            if (texture != IntPtr.Zero)
                SDL.SDL_DestroyTexture(texture);

            TextArea = CalculateSize();
            textureSize = new Size(Pow2(TextArea.Width), Pow2(TextArea.Height));
            texture = SDL.SDL_CreateTexture(renderer, SDL.SDL_PIXELFORMAT_ARGB8888, (int)SDL.SDL_TextureAccess.SDL_TEXTUREACCESS_TARGET, textureSize.Width, textureSize.Height);

            if (texture == IntPtr.Zero)
                throw new ApplicationException("Unaple to initialize texture.");

            SDL.SDL_SetTextureBlendMode(texture, SDL.SDL_BlendMode.SDL_BLENDMODE_BLEND);
            SDL.SDL_SetRenderTarget(renderer, texture);
            SDL.SDL_SetRenderDrawColor(renderer, 0, 0, 0, 0);
            SDL.SDL_RenderFillRect(renderer, IntPtr.Zero);
            SDL.SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255);

            int cx = 0;
            int cy = 0;
            foreach (var ch in text)
            {
                WriteCharacter(cx, cy, (byte)ch);
                cx += font.font.CharacterMetric[(byte)ch].Width;
            }

            SDL.SDL_SetRenderTarget(renderer, IntPtr.Zero);
        }

        internal void WriteCharacter(int cx, int cy, byte character)
        {
            var rect = new SDL.SDL_Rect
            {
                x = cx,
                y = cy,
                w = font.sprite.FrameSize.Width,
                h = font.sprite.FrameSize.Height
            };

            SDL.SDL_SetTextureColorMod(font.sprite.textures[character], TextColor.R, TextColor.G, TextColor.B);
            SDL.SDL_RenderCopy(renderer, font.sprite.textures[character], IntPtr.Zero, ref rect);

        }

        public void Dispose()
        {
            if (texture != IntPtr.Zero)
                SDL.SDL_DestroyTexture(texture);
            texture = IntPtr.Zero;
        }
    }
}
