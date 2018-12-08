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

        private readonly ItemContainerLayout itemContainerLayout;
        private readonly IMouseInfoProvider mouseInfoProvider;

        public Item ContainedItem { get; internal set; }

        private readonly Dictionary<eItemContainerType, ISprite> sprites = new Dictionary<eItemContainerType, ISprite>();

        private Point location = new Point();

        public Point Location
        {
            get => location;
            set
            {
                if (location == value)
                    return;
                location = value;

                placeholderSprite.Location = new Point(value.X, value.Y + placeholderSprite.LocalFrameSize.Height);
            }
        }

        private readonly ISprite placeholderSprite;

        public Size Size { get; internal set; }
        
        public ItemContainer(IRenderWindow renderWindow, IGameState gameState, ItemContainerLayout itemContainerLayout, IMouseInfoProvider mouseInfoProvider)
        {
            this.renderWindow = renderWindow;
            this.gameState = gameState;
            this.itemContainerLayout = itemContainerLayout;
            this.mouseInfoProvider = mouseInfoProvider;

            placeholderSprite = renderWindow.LoadSprite(itemContainerLayout.ResourceName, itemContainerLayout.PaletteName);
            placeholderSprite.Location = new Point(location.X, location.Y + placeholderSprite.LocalFrameSize.Height);
            this.Size = placeholderSprite.FrameSize; // For all but generic size is equal to the placeholder size. Source: me.
        }

        public void SetContainedItem(Item containedItem)
        {
            ContainedItem = containedItem;

            if (ContainedItem != null)
            {
                sprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(this.ContainedItem.InvFile), Palettes.Units);
                sprite.Location = new Point(location.X, location.Y + sprite.LocalFrameSize.Height);
            }
        }

        public void Update()
        {
            var hovered = (mouseInfoProvider.MouseX >= location.X && mouseInfoProvider.MouseX < (location.X + this.Size.Width))
                && (mouseInfoProvider.MouseY >= location.Y && mouseInfoProvider.MouseY < (location.Y + this.Size.Height));

            if (hovered && mouseInfoProvider.LeftMousePressed)
            {
                // If there is an item contained, remove from container and send to mouse
                if (this.ContainedItem != null)
                {
                    if (this.gameState.SelectedItem != null)
                    {
                        var switchItem = this.gameState.SelectedItem;

                        this.gameState.SelectItem(this.ContainedItem);
                        this.SetContainedItem(switchItem);
                    } else
                    {
                        this.gameState.SelectItem(this.ContainedItem);
                        this.SetContainedItem(null);
                    }
                    
                }
                else if (this.gameState.SelectedItem != null)
                {
                    this.SetContainedItem(this.gameState.SelectedItem);
                    this.gameState.SelectItem(null);
                }
            }
        }

        public void Render()
        {
            if (this.ContainedItem == null)
            {
                renderWindow.Draw(placeholderSprite, this.itemContainerLayout.BaseFrame);
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
