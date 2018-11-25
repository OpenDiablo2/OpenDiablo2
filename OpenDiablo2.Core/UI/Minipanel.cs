using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    // TODO: Allow to set Minipanel.buttons.character.OnAction or similar
    public sealed class Minipanel
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite;

        private Button characterBtn, inventoryBtn, skillBtn, automapBtn, messageBtn, questBtn, menuBtn;

        private Point location = new Point();
        public Point Location
        {
            get => location;
            set
            {
                if (location == value)
                    return;
                location = value;

                sprite.Location = new Point(value.X, value.Y + sprite.LocalFrameSize.Height);
            }
        }

        public Minipanel(IRenderWindow renderWindow, Func<eButtonType, Button> createButton)
        {
            this.renderWindow = renderWindow;

            sprite = renderWindow.LoadSprite(ResourcePaths.TextBox2, Palettes.Units);

            // TODO: Move button configuration to a readonly dictionary indexed by name
            characterBtn = createButton(eButtonType.Minimap);
            characterBtn.BaseFrame = 0;
            characterBtn.Location = new Point(0, 0);
            
            inventoryBtn = createButton(eButtonType.Minimap);
            inventoryBtn.BaseFrame = 2;
            inventoryBtn.Location = new Point(20, 0);

            skillBtn = createButton(eButtonType.Minimap);
            skillBtn.BaseFrame = 4;
            skillBtn.Location = new Point(40, 0);
            
            automapBtn = createButton(eButtonType.Minimap);
            automapBtn.BaseFrame = 8;
            automapBtn.Location = new Point(60, 0);

            messageBtn = createButton(eButtonType.Minimap);
            messageBtn.BaseFrame = 10;
            messageBtn.Location = new Point(80, 0);
            

            questBtn = createButton(eButtonType.Minimap);
            questBtn.BaseFrame = 12;
            questBtn.Location = new Point(100, 0);
            

            menuBtn = createButton(eButtonType.Minimap);
            menuBtn.BaseFrame = 14;
            menuBtn.Location = new Point(120, 0);
            
        }


        public void Update()
        {
            characterBtn.Update();
            inventoryBtn.Update();
            skillBtn.Update();
            automapBtn.Update();
            messageBtn.Update();
            questBtn.Update();
            menuBtn.Update();
        }

        public void Render()
        {
            renderWindow.Draw(sprite);

            characterBtn.Location.Offset(Location);
            characterBtn.Render();

            inventoryBtn.Location.Offset(Location);
            inventoryBtn.Render();

            skillBtn.Location.Offset(Location);
            skillBtn.Render();

            automapBtn.Location.Offset(Location);
            automapBtn.Render();

            messageBtn.Location.Offset(Location);
            messageBtn.Render();

            questBtn.Location.Offset(Location);
            questBtn.Render();

            menuBtn.Location.Offset(Location);
            menuBtn.Render();
        }
    }
}
