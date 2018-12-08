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
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

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
            foreach(var ch in text)
            {
                w += font.CharacterMetric[ch].Width;
                h = Math.Max(h, font.CharacterMetric[ch].Height);
            }

            return new Size(w, h);
        }
    }
}
