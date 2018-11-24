﻿using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.SDL2_
{
    internal sealed class SDL2Font : IFont
    {
        internal readonly MPQFont font;
        internal readonly SDL2Sprite sprite;

        public Palette CurrentPalette
        {
            get => sprite.CurrentPalette;
            set => sprite.CurrentPalette = value;
        }

        internal SDL2Font(MPQFont font, IntPtr renderer)
        {
            this.font = font;

            sprite = new SDL2Sprite(font.FontImageSet, renderer);
        }

        public void Dispose()
        {
            sprite.Dispose();
        }

        public Size CalculateSize(string text)
        {
            int w = 0;
            int h = 0;
            foreach(byte ch in text)
            {
                w += font.CharacterMetric[ch].Width;
                h = Math.Max(h, font.CharacterMetric[ch].Height);
            }

            return new Size(w, h);
        }
    }
}
