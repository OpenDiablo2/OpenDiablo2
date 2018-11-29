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
