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
        private readonly IGameState gameState;
        private ISprite sprite;

        private eItemContainerType itemContainerType;
        private IMouseInfoProvider mouseInfoProvider;

        public Item ContainedItem { get; set; }

        public bool waitForMouseUp = false;

        private Dictionary<eItemContainerType, ISprite> sprites = new Dictionary<eItemContainerType, ISprite>();

        private Point location = new Point();
        private Item previouslyContainedItem;

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

        public ItemContainer(IRenderWindow renderWindow, IGameState gameState, eItemContainerType itemContainerType, IMouseInfoProvider mouseInfoProvider)
        {
            this.renderWindow = renderWindow;
            this.gameState = gameState;
            this.itemContainerType = itemContainerType;
            this.mouseInfoProvider = mouseInfoProvider;

            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units); // Ignore for now

            sprites.Add(eItemContainerType.Helm, renderWindow.LoadSprite(ResourcePaths.HelmGlovePlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Glove, renderWindow.LoadSprite(ResourcePaths.HelmGlovePlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Armor, renderWindow.LoadSprite(ResourcePaths.ArmorPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Belt, renderWindow.LoadSprite(ResourcePaths.BeltPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Boots, renderWindow.LoadSprite(ResourcePaths.BootsPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Weapon, renderWindow.LoadSprite(ResourcePaths.WeaponsPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Amulet, renderWindow.LoadSprite(ResourcePaths.RingAmuletPlaceholder, Palettes.Units));
            sprites.Add(eItemContainerType.Ring, renderWindow.LoadSprite(ResourcePaths.RingAmuletPlaceholder, Palettes.Units));
        }

        public void Update()
        {
            previouslyContainedItem = ContainedItem;
            var hovered = (mouseInfoProvider.MouseX >= location.X && mouseInfoProvider.MouseX < (location.X + this.sprites[this.itemContainerType].FrameSize.Width))
                && (mouseInfoProvider.MouseY >= location.Y && mouseInfoProvider.MouseY < (location.Y + this.sprites[this.itemContainerType].FrameSize.Height));

            if (hovered && mouseInfoProvider.LeftMousePressed)
            {
                // If there is an item contained, remove from container and send to mouse
                if (this.ContainedItem != null)
                {
                    if (this.gameState.SelectedItem != null)
                    {
                        var switchItem = this.gameState.SelectedItem;

                        this.gameState.SelectItem(this.ContainedItem);
                        this.ContainedItem = switchItem;
                    } else
                    {
                        this.gameState.SelectItem(this.ContainedItem);
                        this.ContainedItem = null;
                    }
                    
                }
                else if (this.gameState.SelectedItem != null)
                {
                    this.ContainedItem = this.gameState.SelectedItem;
                    this.gameState.SelectItem(null);
                }
            }

            if (this.ContainedItem == null)
            {
                if (this.itemContainerType != eItemContainerType.Generic && sprite != sprites[this.itemContainerType])
                {
                    sprite = sprites[this.itemContainerType];
                    sprite.Location = new Point(location.X, location.Y + sprite.LocalFrameSize.Height);
                }
            }
            else
            {
                if (ContainedItem != previouslyContainedItem)
                {
                    sprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(this.ContainedItem.InvFile), Palettes.Units);
                    sprite.Location = new Point(location.X, location.Y + sprite.LocalFrameSize.Height);
                }
            }

        }

        public void Render()
        {
            if (this.ContainedItem == null)
            {
                switch (this.itemContainerType)
                {
                    case eItemContainerType.Helm:
                        renderWindow.Draw(sprite, 1);
                        break;
                    case eItemContainerType.Ring:
                        renderWindow.Draw(sprite, 1);
                        break;
                    case eItemContainerType.Glove:
                    case eItemContainerType.Armor:
                    case eItemContainerType.Belt:
                    case eItemContainerType.Weapon:
                    case eItemContainerType.Amulet:
                    case eItemContainerType.Boots:
                        renderWindow.Draw(sprite);
                        break;
                }
            }
            else
            {
                renderWindow.Draw(sprite);
            }
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
