using OpenDiablo2.Common.Enums;
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

        public int MaxWidth { get; set; }
        public eTextAlign Alignment { get; set; }

        internal SDL2Label(IFont font, IntPtr renderer)
        {
            this.renderer = renderer;
            this.font = font as SDL2Font;
            this.texture = IntPtr.Zero;
            this.MaxWidth = -1;
        }

        internal Size CalculateSize()
        {
            if (MaxWidth == -1)
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

            if (MaxWidth < (font.sprite.FrameSize.Width))
                throw new ApplicationException("Max label width cannot be smaller than a single character.");

            var lastWordIndex = 0;
            var width = 0;
            var maxWidth = 0;
            var height = font.sprite.FrameSize.Height;
            for (int idx = 0; idx < text.Length; idx++)
            {
                width += font.font.CharacterMetric[(byte)text[idx]].Width;

                if (width >= MaxWidth)
                {
                    idx = lastWordIndex;
                    height += font.font.CharacterMetric[(byte)'|'].Height + 6;
                    width = 0;
                    continue;
                }
                maxWidth = Math.Max(width, maxWidth);

                if (idx > 0 && (text[idx - 1] == ' ') && (text[idx] != ' '))
                    lastWordIndex = idx;
            }

            return new Size(maxWidth, height);

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
            SDL.SDL_RenderClear(renderer);
            SDL.SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255);

            if (MaxWidth == -1)
            {

                int cx = 0;
                int cy = 0;
                foreach (var ch in text)
                {
                    WriteCharacter(cx, cy, (byte)ch);
                    cx += font.font.CharacterMetric[(byte)ch].Width;
                }
            }
            else
            {
                var linesToRender = new List<string>();

                var lastWordIndex = 0;
                var width = 0;
                var lastStartX = 0;
                for (int idx = 0; idx < text.Length; idx++)
                {
                    width += font.font.CharacterMetric[(byte)text[idx]].Width;

                    if (width >= MaxWidth)
                    {
                        idx = lastWordIndex;
                        linesToRender.Add(text.Substring(lastStartX, lastWordIndex- lastStartX));
                        lastStartX = idx;
                        width = 0;
                        continue;
                    }

                    if (idx > 0 && text[idx - 1] == ' ' && text[idx] != ' ')
                        lastWordIndex = idx;
                }

                var lastLine = text.Substring(lastStartX)?.Trim();
                if (!String.IsNullOrEmpty(lastLine))
                    linesToRender.Add(lastLine);

                var y = 0;
                foreach(var line in linesToRender)
                {
                    var lineWidth = (line.Sum(c => font.font.CharacterMetric[(byte)c].Width));
                    var x = 0;

                    if (Alignment == eTextAlign.Centered)
                        x = (TextArea.Width / 2) - (lineWidth / 2);
                    else if (Alignment == eTextAlign.Right)
                        x = TextArea.Width - lineWidth;

                    foreach (var ch in line)
                    {
                        WriteCharacter(x, y, (byte)ch);
                        x += font.font.CharacterMetric[(byte)ch].Width;
                    }

                    y += font.font.CharacterMetric[(byte)'|'].Height + 6;
                }
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

            font.sprite.Frame = character;
            SDL.SDL_SetTextureColorMod(font.sprite.texture, TextColor.R, TextColor.G, TextColor.B);
            SDL.SDL_RenderCopy(renderer, font.sprite.texture, IntPtr.Zero, ref rect);

        }

        public void Dispose()
        {
            if (texture != IntPtr.Zero)
                SDL.SDL_DestroyTexture(texture);
            texture = IntPtr.Zero;
        }
    }
}
