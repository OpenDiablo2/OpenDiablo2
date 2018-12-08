using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    /**
     * TODO: Check positioning, it's probably not exact
     * TODO: Add logic so it can be used as an element in inventory grid
     **/
    public sealed class InventoryPanel : IInventoryPanel
    {
        private readonly IRenderWindow renderWindow;
        private readonly ISprite sprite;
        private Point location;

        public Point Location {
            get => location;
            set {
#pragma warning disable S4275 // Getters and setters should access the expected fields
                previouslyContainedItem = location;
#pragma warning restore S4275 // Getters and setters should access the expected fields
                location = value;
            }
        }

        public eButtonType PanelType => eButtonType.MinipanelInventory;
        public ePanelFrameType FrameType => ePanelFrameType.Right;

        // Test vars
        public IItemContainer helmContainer, armorContainer, weaponLeftContainer, weaponRightContainer, beltContainer, gloveContainer, bootsContainer;
        private Point previouslyContainedItem;
        private readonly IItemContainer ringtLeftContainer;
        private readonly IItemContainer ringtRightContainer;
        private readonly IItemContainer amuletContainer;

        public InventoryPanel(IRenderWindow renderWindow, IItemManager itemManager, Func<eItemContainerType, IItemContainer> createItemContainer)
        {
            this.renderWindow = renderWindow;

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, new Point(402,61));
            Location = new Point(400, 0);

            this.helmContainer = createItemContainer(eItemContainerType.Helm);
            this.helmContainer.Location = new Point(Location.X + 138, Location.Y + 68);
            this.helmContainer.SetContainedItem(itemManager.getItem("cap"));

            this.amuletContainer = createItemContainer(eItemContainerType.Amulet);
            this.amuletContainer.Location = new Point(Location.X + 211, Location.Y + 92);
            this.amuletContainer.SetContainedItem(itemManager.getItem("vip"));

            this.armorContainer = createItemContainer(eItemContainerType.Armor);
            this.armorContainer.Location = new Point(Location.X + 138, Location.Y + 138);
            this.armorContainer.SetContainedItem(itemManager.getItem("hla"));

            this.weaponLeftContainer = createItemContainer(eItemContainerType.Weapon);
            this.weaponLeftContainer.Location = new Point(Location.X + 22, Location.Y + 108);
            this.weaponLeftContainer.SetContainedItem(itemManager.getItem("ame"));

            this.weaponRightContainer = createItemContainer(eItemContainerType.Weapon);
            this.weaponRightContainer.Location = new Point(Location.X + 255, Location.Y + 108);
            this.weaponRightContainer.SetContainedItem(itemManager.getItem("paf"));

            this.beltContainer = createItemContainer(eItemContainerType.Belt);
            this.beltContainer.Location = new Point(Location.X + 138, Location.Y + 238);
            this.beltContainer.SetContainedItem(itemManager.getItem("vbl"));

            this.ringtLeftContainer = createItemContainer(eItemContainerType.Ring);
            this.ringtLeftContainer.Location = new Point(Location.X + 97, Location.Y + 238);
            this.ringtLeftContainer.SetContainedItem(itemManager.getItem("rin"));

            this.ringtRightContainer = createItemContainer(eItemContainerType.Ring);
            this.ringtRightContainer.Location = new Point(Location.X + 211, Location.Y + 238);
            this.ringtRightContainer.SetContainedItem(itemManager.getItem("rin"));

            this.gloveContainer = createItemContainer(eItemContainerType.Glove);
            this.gloveContainer.Location = new Point(Location.X + 22, Location.Y + 238);
            this.gloveContainer.SetContainedItem(itemManager.getItem("tgl"));

            this.bootsContainer = createItemContainer(eItemContainerType.Boots);
            this.bootsContainer.Location = new Point(Location.X + 255, Location.Y + 238);
            this.bootsContainer.SetContainedItem(itemManager.getItem("lbt"));
        }

        public void Update()
        {
            helmContainer.Update();
            amuletContainer.Update();
            armorContainer.Update();
            weaponLeftContainer.Update();
            weaponRightContainer.Update();
            beltContainer.Update();
            ringtLeftContainer.Update();
            ringtRightContainer.Update();
            gloveContainer.Update();
            bootsContainer.Update();
        }

        public void Render()
        {
            renderWindow.Draw(sprite, 2, 2, 1);

            helmContainer.Render();
            amuletContainer.Render();
            armorContainer.Render();
            weaponLeftContainer.Render();
            weaponRightContainer.Render();
            beltContainer.Render();
            ringtLeftContainer.Render();
            ringtRightContainer.Render();
            gloveContainer.Render();
            bootsContainer.Render();
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
