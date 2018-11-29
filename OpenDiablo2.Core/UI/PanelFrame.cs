using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    public sealed class PanelFrame : IPanelFrame
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite;
        private ePanelFrameType panelFrameType;

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

        public PanelFrame(IRenderWindow renderWindow, ePanelFrameType panelFrameType)
        {
            this.renderWindow = renderWindow;
            this.panelFrameType = panelFrameType;
            
            sprite = renderWindow.LoadSprite(ResourcePaths.Frame, Palettes.Units, new Point(0, 0));

            Location = new Point(0, 0);

           
        }

        private void DrawPanel()
        {
            switch(this.panelFrameType)
            {
                case ePanelFrameType.Left:
                    renderWindow.Draw(sprite, 0, new Point(0, 256));
                    renderWindow.Draw(sprite, 1, new Point(256, 66));
                    renderWindow.Draw(sprite, 2, new Point(0, 256 + 231));
                    renderWindow.Draw(sprite, 3, new Point(0, 256 + 231 + 66));
                    renderWindow.Draw(sprite, 4, new Point(256, 256 + 231 + 66));
                    break;
                case ePanelFrameType.Right:
                    renderWindow.Draw(sprite, 5, new Point(400 + 0, 66));
                    renderWindow.Draw(sprite, 6, new Point(400 + 145, 256));
                    renderWindow.Draw(sprite, 7, new Point(400 + 145 + 169, 256 + 231));
                    renderWindow.Draw(sprite, 8, new Point(400 + 145, 256 + 231 + 66));
                    renderWindow.Draw(sprite, 9, new Point(400 + 0, 256 + 231 + 66));
                    break;
            }
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
