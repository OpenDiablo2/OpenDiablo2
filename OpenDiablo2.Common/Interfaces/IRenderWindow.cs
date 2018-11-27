using System;
using System.Drawing;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Common.Interfaces
{
    public interface IRenderWindow : IDisposable
    {
        bool IsRunning { get; }
        void Update();
        void Clear();
        void Sync();
        void Quit();
        ISprite LoadSprite(string resourcePath, string palette, Point location);
        ISprite LoadSprite(string resourcePath, string palette);
        IFont LoadFont(string resourcePath, string palette);
        ILabel CreateLabel(IFont font);
        ILabel CreateLabel(IFont font, string text);
        ILabel CreateLabel(IFont font, Point position, string text);
        void Draw(ISprite sprite);
        void Draw(ISprite sprite, Point location);
        void Draw(ISprite sprite, int frame, Point location);
        void Draw(ISprite sprite, int frame);
        void Draw(ISprite sprite, int xSegments, int ySegments, int offset);
        void Draw(ILabel label);
        void DrawMapCell(int xCell, int yCell, int xPixel, int yPixel, MPQDS1 mapData, int main_index, int sub_index, Palette palette, int orientation = -1);
    }
}
