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

        public CharacterPanel(IRenderWindow renderWindow)
        {
            this.renderWindow = renderWindow;
            
            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units);
            Location = new Point(0, 0);

           
        }

        public void DrawPanel()
        {
            renderWindow.Draw(sprite, 0, new Point(location.X, location.Y));
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
