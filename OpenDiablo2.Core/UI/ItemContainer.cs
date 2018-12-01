using System;
using System.Collections.Generic;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;

namespace OpenDiablo2.Core.UI
{
    // TODO: Self-align when side panels are open
    public sealed class ItemContainer : IItemContainer
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite;

        private IButton characterBtn, inventoryBtn, skillBtn, automapBtn, messageBtn, questBtn, menuBtn;

        private eItemContainerType itemContainerType;

        public Item ContainedItem { get; internal set; }

        private Dictionary<eItemContainerType, ISprite> sprites = new Dictionary<eItemContainerType, ISprite>();

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

        public ItemContainer(IRenderWindow renderWindow, IGameState gameState, eItemContainerType itemContainerType)
        {
            this.renderWindow = renderWindow;
            this.itemContainerType = itemContainerType;
            
            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units);
            Location = new Point(800/2-sprite.LocalFrameSize.Width/2, 526);


            sprites.Add(eItemContainerType.Helm, renderWindow.LoadSprite(ResourcePaths.HelmGlovePlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Armor, renderWindow.LoadSprite(ResourcePaths.ArmorPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Belt, renderWindow.LoadSprite(ResourcePaths.BeltPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Weapon, renderWindow.LoadSprite(ResourcePaths.WeaponsPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Amulet, renderWindow.LoadSprite(ResourcePaths.RingAmuletPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Ring, renderWindow.LoadSprite(ResourcePaths.RingAmuletPlaceholder, Palettes.Units));
        }


        public void Update()
        {
            if(this.ContainedItem == null)
            {
                switch(this.itemContainerType)
                {
                    case eItemContainerType.Helm:
                    case eItemContainerType.Glove:
                    case eItemContainerType.Armor:
                    case eItemContainerType.Belt:
                    case eItemContainerType.Weapon:
                    case eItemContainerType.Amulet:
                    case eItemContainerType.Ring:
                        sprite = sprites[this.itemContainerType];
                        break;
                }
            } else
            {
                // Nothing for now
            }
        }

        public void Render()
        {
            if (this.ContainedItem == null)
            {
                switch (this.itemContainerType)
                {
                    case eItemContainerType.Helm:
                        renderWindow.Draw(sprite,1);
                        break;
                    case eItemContainerType.Ring:
                        renderWindow.Draw(sprite, 1);
                        break;
                    case eItemContainerType.Glove:
                    case eItemContainerType.Armor:
                    case eItemContainerType.Belt:
                    case eItemContainerType.Weapon:
                    case eItemContainerType.Amulet:
                        renderWindow.Draw(sprite);
                        break;
                }
            }
            else
            {
                // Nothing for now
            }

            renderWindow.Draw(sprite);
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
