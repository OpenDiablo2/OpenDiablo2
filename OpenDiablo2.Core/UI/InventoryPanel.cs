/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Drawing;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Extensions;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    /**
     * TODO: Add logic so it can be used as an element in inventory grid
     **/
    public sealed class InventoryPanel : IInventoryPanel
    {
        private readonly IRenderWindow renderWindow;
        private readonly ISprite panelSprite;
        
        public IItemContainer helmContainer, armorContainer, beltContainer, gloveContainer, bootsContainer,
            leftHandContainer, rightHandContainer, secondaryLeftHandContainer, secondaryRightHandContainer,
            ringtLeftContainer, ringtRightContainer, amuletContainer;

        private readonly IButton closeButton, secondaryLeftButton, secondaryRightButton, goldButton;

        public event OnPanelClosedEvent OnPanelClosed;

        public InventoryPanel(IRenderWindow renderWindow, 
            IItemManager itemManager, 
            Func<eItemContainerType, IItemContainer> createItemContainer,
            Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;

            panelSprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, FrameType.GetOffset());

            closeButton = createButton(eButtonType.Close);
            closeButton.Location = panelSprite.Location + new Size(18, 384);
            closeButton.OnActivate = () => OnPanelClosed?.Invoke(this);

            secondaryLeftButton = createButton(eButtonType.SecondaryInvHand);
            secondaryLeftButton.Location = panelSprite.Location + new Size(15, 22);
            secondaryLeftButton.OnActivate = ToggleWeaponsSlot;
            secondaryLeftButton.ClickableRect = new Size(0, 20);
            secondaryLeftButton.AllowFrameChange = false;

            secondaryRightButton = createButton(eButtonType.SecondaryInvHand);
            secondaryRightButton.Location = panelSprite.Location + new Size(246, 22);
            secondaryRightButton.OnActivate = ToggleWeaponsSlot;
            secondaryRightButton.ClickableRect = new Size(0, 20);
            secondaryRightButton.AllowFrameChange = false;

            goldButton = createButton(eButtonType.GoldCoin);
            goldButton.Location = panelSprite.Location + new Size(84, 391);
            goldButton.OnActivate = OpenGoldDrop;

            helmContainer = createItemContainer(eItemContainerType.Helm);
            helmContainer.Location = panelSprite.Location + new Size(135, 5);
            helmContainer.SetContainedItem(itemManager.getItem("cap"));
            
            amuletContainer = createItemContainer(eItemContainerType.Amulet);
            amuletContainer.Location = panelSprite.Location + new Size(209, 34);
            amuletContainer.SetContainedItem(itemManager.getItem("vip"));
            
            armorContainer = createItemContainer(eItemContainerType.Armor);
            armorContainer.Location = panelSprite.Location + new Size(135, 75);
            armorContainer.SetContainedItem(itemManager.getItem("hla"));

            leftHandContainer = createItemContainer(eItemContainerType.Weapon);
            leftHandContainer.Location = panelSprite.Location + new Size(20, 47);
            leftHandContainer.SetContainedItem(itemManager.getItem("ame"));

            rightHandContainer = createItemContainer(eItemContainerType.Weapon);
            rightHandContainer.Location = panelSprite.Location + new Size(253, 47);
            rightHandContainer.SetContainedItem(itemManager.getItem("paf"));

            secondaryLeftHandContainer = createItemContainer(eItemContainerType.Weapon);
            secondaryLeftHandContainer.Location = panelSprite.Location + new Size(24, 45);
            secondaryLeftHandContainer.SetContainedItem(itemManager.getItem("crs"));

            secondaryRightHandContainer = createItemContainer(eItemContainerType.Weapon);
            secondaryRightHandContainer.Location = panelSprite.Location + new Size(257, 45);
            secondaryRightHandContainer.SetContainedItem(itemManager.getItem("kit"));

            beltContainer = createItemContainer(eItemContainerType.Belt);
            beltContainer.Location = panelSprite.Location + new Size(136, 178);
            beltContainer.SetContainedItem(itemManager.getItem("vbl"));
            
            ringtLeftContainer = createItemContainer(eItemContainerType.Ring);
            ringtLeftContainer.Location = panelSprite.Location + new Size(95, 179);
            ringtLeftContainer.SetContainedItem(itemManager.getItem("rin"));
            
            ringtRightContainer = createItemContainer(eItemContainerType.Ring);
            ringtRightContainer.Location = panelSprite.Location + new Size(209, 179);
            ringtRightContainer.SetContainedItem(itemManager.getItem("rin"));
            
            gloveContainer = createItemContainer(eItemContainerType.Glove);
            gloveContainer.Location = panelSprite.Location + new Size(20, 179);
            gloveContainer.SetContainedItem(itemManager.getItem("tgl"));
            
            bootsContainer = createItemContainer(eItemContainerType.Boots);
            bootsContainer.Location = panelSprite.Location + new Size(251, 178);
            bootsContainer.SetContainedItem(itemManager.getItem("lbt"));
        }

        public eButtonType PanelType => eButtonType.MinipanelInventory;
        public ePanelFrameType FrameType => ePanelFrameType.Right;

        public bool IsSecondaryEquipped { get; private set; }

        public void Update()
        {
            if (IsSecondaryEquipped)
            {
                secondaryLeftHandContainer.Update();
                secondaryRightHandContainer.Update();
            }
            else
            {
                leftHandContainer.Update();
                rightHandContainer.Update();
            }

            secondaryLeftButton.Update();
            secondaryRightButton.Update();

            closeButton.Update();
            goldButton.Update();

            helmContainer.Update();
            amuletContainer.Update();
            armorContainer.Update();
            beltContainer.Update();
            ringtLeftContainer.Update();
            ringtRightContainer.Update();
            gloveContainer.Update();
            bootsContainer.Update();
        }

        public void Render()
        {
            renderWindow.Draw(panelSprite, 2, 2, 1);

            if (IsSecondaryEquipped)
            {
                secondaryLeftButton.Render();
                secondaryRightButton.Render();
                secondaryLeftHandContainer.Render();
                secondaryRightHandContainer.Render();
            }
            else
            {
                leftHandContainer.Render();
                rightHandContainer.Render();
            }

            closeButton.Render();
            goldButton.Render();

            helmContainer.Render();
            amuletContainer.Render();
            armorContainer.Render();
            beltContainer.Render();
            ringtLeftContainer.Render();
            ringtRightContainer.Render();
            gloveContainer.Render();
            bootsContainer.Render();
        }

        public void Dispose()
        {
            panelSprite.Dispose();
        }

        private void ToggleWeaponsSlot()
        {
            IsSecondaryEquipped = !IsSecondaryEquipped;
        }

        private void OpenGoldDrop()
        {

        }
    }
}
