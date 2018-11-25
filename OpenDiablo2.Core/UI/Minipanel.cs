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
            Location = new Point(400, 600);

            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units);

            characterBtn = createButton(eButtonType.MinipanelCharacter);
            characterBtn.Location = new Point(0, 0);

            inventoryBtn = createButton(eButtonType.MinipanelInventory);
            inventoryBtn.Location = new Point(20, 0);

            skillBtn = createButton(eButtonType.MinipanelSkill);
            skillBtn.Location = new Point(40, 0);

            automapBtn = createButton(eButtonType.MinipanelAutomap);
            automapBtn.Location = new Point(60, 0);

            messageBtn = createButton(eButtonType.MinipanelMessage);
            messageBtn.Location = new Point(80, 0);

            questBtn = createButton(eButtonType.MinipanelQuest);
            questBtn.Location = new Point(100, 0);

            menuBtn = createButton(eButtonType.MinipanelMenu);
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
            
            characterBtn.Render();
            inventoryBtn.Render();
            skillBtn.Render();
            automapBtn.Render();
            messageBtn.Render();
            questBtn.Render();
            menuBtn.Render();
        }
    }
}
