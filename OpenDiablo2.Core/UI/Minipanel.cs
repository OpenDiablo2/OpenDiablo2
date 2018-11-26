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
    // TODO: Allow to set Minipanel.buttons.character.OnAction or similar for button delegates
    public sealed class MiniPanel : IMiniPanel
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite;

        private IButton characterBtn, inventoryBtn, skillBtn, automapBtn, messageBtn, questBtn, menuBtn;

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

        public MiniPanel(IRenderWindow renderWindow, Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;
            
            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units);
            Location = new Point(800/2-sprite.LocalFrameSize.Width/2, 526);

            characterBtn = createButton(eButtonType.MinipanelCharacter);
            characterBtn.Location = new Point(3 + Location.X, 3 + Location.Y);

            inventoryBtn = createButton(eButtonType.MinipanelInventory);
            inventoryBtn.Location = new Point(24 + Location.X, 3 + Location.Y);

            skillBtn = createButton(eButtonType.MinipanelSkill);
            skillBtn.Location = new Point(45 + Location.X, 3 + Location.Y);

            automapBtn = createButton(eButtonType.MinipanelAutomap);
            automapBtn.Location = new Point(66 + Location.X, 3 + Location.Y);

            messageBtn = createButton(eButtonType.MinipanelMessage);
            messageBtn.Location = new Point(87 + Location.X, 3 + Location.Y);

            questBtn = createButton(eButtonType.MinipanelQuest);
            questBtn.Location = new Point(108 + Location.X, 3 + Location.Y);

            menuBtn = createButton(eButtonType.MinipanelMenu);
            menuBtn.Location = new Point(129 + Location.X, 3 + Location.Y);
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

        public void Dispose()
        {
            characterBtn.Dispose();
            inventoryBtn.Dispose();
            skillBtn.Dispose();
            automapBtn.Dispose();
            messageBtn.Dispose();
            questBtn.Dispose();
            menuBtn.Dispose();

            sprite.Dispose();
        }
    }
}
