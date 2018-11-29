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
        private ISprite sprite;
        private IPanelFrame panelFrame;

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

        public InventoryPanel(Func<ePanelFrameType, IPanelFrame> createPanelFrame, IRenderWindow renderWindow)
        {
            this.renderWindow = renderWindow;
            this.panelFrame = createPanelFrame(ePanelFrameType.Right);

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, new Point(402,61));
            Location = new Point(400, 0);
        }

        public void Update()
        {
            
        }

        public void Render()
        {
            panelFrame.Render();
            renderWindow.Draw(sprite, 2, 2, 1);
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
