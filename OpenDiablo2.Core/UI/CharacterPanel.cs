using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    public sealed class CharacterPanel : ICharacterPanel
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite;
        private IPanelFrame panelFrame;

        public Point Location { get; set; }

        public CharacterPanel(IRenderWindow renderWindow, Func<ePanelFrameType, IPanelFrame> createPanelFrame)
        {
            this.renderWindow = renderWindow;
            this.panelFrame = createPanelFrame(ePanelFrameType.Left);

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, new Point(79,61));
            Location = new Point(0, 0);

           
        }

        public void Update()
        {
            
        }

        public void Render()
        {
            panelFrame.Render();
            renderWindow.Draw(sprite, 2, 2, 0);
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
