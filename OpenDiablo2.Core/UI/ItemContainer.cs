using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System.Collections.Generic;
using System.Drawing;

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
        private readonly ISessionManager sessionManager;

        public ItemInstance ContainedItem { get; internal set; }

        private readonly Dictionary<eItemContainerType, ISprite> sprites = new Dictionary<eItemContainerType, ISprite>();

        private Point location = new Point();
        private IMapRenderer mapRenderer;

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
        public string Slot { get; set; }
        
        public ItemContainer(IRenderWindow renderWindow, IGameState gameState, ItemContainerLayout itemContainerLayout, IMouseInfoProvider mouseInfoProvider, ISessionManager sessionManager, IMapRenderer mapRenderer)
        {
            this.renderWindow = renderWindow;
            this.gameState = gameState;
            this.itemContainerLayout = itemContainerLayout;
            this.mouseInfoProvider = mouseInfoProvider;
            this.sessionManager = sessionManager;
            this.mapRenderer = mapRenderer;

            placeholderSprite = renderWindow.LoadSprite(itemContainerLayout.ResourceName, itemContainerLayout.PaletteName, true);
            placeholderSprite.Location = new Point(location.X, location.Y + placeholderSprite.LocalFrameSize.Height);
            this.Size = placeholderSprite.FrameSize; // For all but generic size is equal to the placeholder size. Source: me.
        }

        public void SetContainedItem(ItemInstance containedItem)
        {
            ContainedItem = containedItem;

            if (ContainedItem != null)
            {
                sprite = renderWindow.LoadSprite(ResourcePaths.GeneratePathForItem(this.ContainedItem.Item.InvFile), Palettes.Units, true);
                sprite.Location = new Point(location.X, location.Y + sprite.LocalFrameSize.Height);
            }
        }

        // TODO: Add all restrictions.
        public bool CanEquip(ItemInstance itemInstance)
        {
            if (Slot == "rarm" || Slot == "larm")
                return (itemInstance.Item is Weapon || (itemInstance.Item is Armor && (itemInstance.Item as Armor).Type == "shie"));

            if (Slot == "tors")
                return (itemInstance.Item is Armor && (itemInstance.Item as Armor).Type == "tors");

            if (Slot == "head")
                return (itemInstance.Item is Armor && (itemInstance.Item as Armor).Type == "helm");

            return true;
        }

        public void Update()
        {
            var hovered = mouseInfoProvider.MouseX >= location.X && mouseInfoProvider.MouseX < (location.X + this.Size.Width)
                && mouseInfoProvider.MouseY >= location.Y && mouseInfoProvider.MouseY < (location.Y + this.Size.Height);

            if (hovered && mouseInfoProvider.LeftMousePressed)
            {
                // If there is an item contained, remove from container and send to mouse
                if (this.ContainedItem != null)
                {
                    if (this.gameState.SelectedItem != null)
                    {
                        if(CanEquip(this.gameState.SelectedItem))
                        {
                            var switchItem = this.gameState.SelectedItem;

                            this.gameState.SelectItem(this.ContainedItem);
                            this.SetContainedItem(switchItem);
                        }
                    } else
                    {
                        this.gameState.SelectItem(this.ContainedItem);
                        this.SetContainedItem(null);
                    }

                    sessionManager.UpdateEquipment(Slot, ContainedItem);
                }
                else if (this.gameState.SelectedItem != null)
                {
                    if (CanEquip(this.gameState.SelectedItem))
                    {
                        this.SetContainedItem(this.gameState.SelectedItem);
                        this.gameState.SelectItem(null);

                        sessionManager.UpdateEquipment(Slot, ContainedItem);
                    }
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
            if(sprite != null)
            {
                sprite.Dispose();
            }
        }
    }
}
