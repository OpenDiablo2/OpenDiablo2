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

using System.Drawing;
using SDL2;

namespace OpenDiablo2.SDL2_
{
    public static class SDL2Extensions
    {
        public static SDL.SDL_Rect ToSDL2Rect(this Rectangle src) => new SDL.SDL_Rect { x = src.X, y = src.Y, w = src.Width, h = src.Height };
        public static Rectangle ToRectangle(this SDL.SDL_Rect src) => new Rectangle { X = src.x, Y = src.y, Width = src.w, Height = src.h };
    }
}
