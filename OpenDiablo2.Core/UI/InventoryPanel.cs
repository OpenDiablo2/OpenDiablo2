using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    public sealed class InventoryPanel : IInventoryPanel
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite, framesprite;

        private Point location = new Point();
        public Point Location
        {
            get => location;
            set
            {
                if (location == value)
                    return;
                location = value;
            }
        }

        public InventoryPanel(IRenderWindow renderWindow)
        {
            this.renderWindow = renderWindow;
            
            framesprite = renderWindow.LoadSprite(ResourcePaths.Frame, Palettes.Units, new Point(0, 0));

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, new Point(402,61));
            Location = new Point(400, 0);

           
        }

        private void DrawPanel()
        {
            renderWindow.Draw(framesprite, 5, new Point(location.X + 0,66));
            renderWindow.Draw(framesprite, 6, new Point(location.X + 145, 256));
            renderWindow.Draw(framesprite, 7, new Point(location.X + 145 + 169, 256+231));
            renderWindow.Draw(framesprite, 8, new Point(location.X + 145, 256 + 231 + 66));
            renderWindow.Draw(framesprite, 9, new Point(location.X + 0, 256 + 231 + 66));
            renderWindow.Draw(sprite, 2, 2, 1);
        }


        public void Update()
        {
            
            
        }

        public void Render()
        {
            DrawPanel();
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
