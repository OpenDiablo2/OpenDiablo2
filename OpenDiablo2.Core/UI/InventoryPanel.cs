using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    public sealed class InventoryPanel : IInventoryPanel
    {
        private readonly IRenderWindow renderWindow;
        private ISprite sprite;
        private IPanelFrame panelFrame;
        private Point location;

        public Point Location {
            get => location;
            set {
                previouslyContainedItem = location;
                location = value;
            }
        }

        // Test vars
        public IItemContainer helmContainer, armorContainer, beltContainer, gloveContainer, bootsContainer;
        private Point previouslyContainedItem;

        public InventoryPanel(Func<ePanelFrameType, IPanelFrame> createPanelFrame, IRenderWindow renderWindow, IItemManager itemManager, Func<eItemContainerType, IItemContainer> createItemContainer)
        {
            this.renderWindow = renderWindow;
            this.panelFrame = createPanelFrame(ePanelFrameType.Right);

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, new Point(402,61));
            Location = new Point(400, 0);

            this.helmContainer = createItemContainer(eItemContainerType.Helm);
            this.helmContainer.Location = new Point(Location.X + 138, Location.Y + 68);
            this.helmContainer.ContainedItem = itemManager.getItem("cap");

            this.armorContainer = createItemContainer(eItemContainerType.Armor);
            this.armorContainer.Location = new Point(Location.X + 138, Location.Y + 138);
            this.armorContainer.ContainedItem = itemManager.getItem("hla");

            this.beltContainer = createItemContainer(eItemContainerType.Belt);
            this.beltContainer.Location = new Point(Location.X + 138, Location.Y + 238);
            this.beltContainer.ContainedItem = itemManager.getItem("vbl");

            this.gloveContainer = createItemContainer(eItemContainerType.Glove);
            this.gloveContainer.Location = new Point(Location.X + 22, Location.Y + 238);
            this.gloveContainer.ContainedItem = itemManager.getItem("tgl");

            this.bootsContainer = createItemContainer(eItemContainerType.Boots);
            this.bootsContainer.Location = new Point(Location.X + 255, Location.Y + 238);
            this.bootsContainer.ContainedItem = itemManager.getItem("lbt");
        }

        public void Update()
        {

            helmContainer.Update();
            armorContainer.Update();
            beltContainer.Update();
            gloveContainer.Update();
            bootsContainer.Update();
        }

        public void Render()
        {
            panelFrame.Render();
            renderWindow.Draw(sprite, 2, 2, 1);

            helmContainer.Render();
            armorContainer.Render();
            beltContainer.Render();
            gloveContainer.Render();
            bootsContainer.Render();
        }

        public void Dispose()
        {
            sprite.Dispose();
        }
    }
}
