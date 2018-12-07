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

        public Point Location { get; set; }

        public CharacterPanel(IRenderWindow renderWindow)
        {
            this.renderWindow = renderWindow;

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, new Point(79,61));
            Location = new Point(0, 0);
        }

        public eButtonType PanelType => eButtonType.MinipanelCharacter;
        public ePanelFrameType FrameType => ePanelFrameType.Left;

        public void Update()
        {
            
        }

        public void Render()
        {
            renderWindow.Draw(sprite, 2, 2, 0);
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
