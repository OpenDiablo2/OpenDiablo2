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
        private readonly ISprite sprite;

        public eButtonType PanelType => eButtonType.MinipanelInventory;
        public ePanelFrameType FrameType => ePanelFrameType.Right;

        // Test vars
        public IItemContainer helmContainer, armorContainer, weaponLeftContainer, weaponRightContainer, beltContainer, gloveContainer, bootsContainer;
        private readonly IItemContainer ringtLeftContainer;
        private readonly IItemContainer ringtRightContainer;
        private readonly IItemContainer amuletContainer;

        private readonly IButton closeButton, goldButton;

        public event OnPanelClosedEvent OnPanelClosed;

        public InventoryPanel(IRenderWindow renderWindow, 
            IItemManager itemManager, 
            Func<eItemContainerType, IItemContainer> createItemContainer,
            Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;

            sprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Units, FrameType.GetOffset(), true);

            closeButton = createButton(eButtonType.Close);
            closeButton.Location = sprite.Location + new Size(18, 384);
            closeButton.OnActivate = () => OnPanelClosed?.Invoke(this);

            goldButton = createButton(eButtonType.GoldCoin);
            goldButton.Location = sprite.Location + new Size(84, 391);
            goldButton.OnActivate = OpenGoldDrop;

            helmContainer = createItemContainer(eItemContainerType.Helm);
            helmContainer.Location = sprite.Location + new Size(135, 5);
            
            amuletContainer = createItemContainer(eItemContainerType.Amulet);
            amuletContainer.Location = sprite.Location + new Size(209, 34);
            
            armorContainer = createItemContainer(eItemContainerType.Armor);
            armorContainer.Location = sprite.Location + new Size(135, 75);
            
            weaponLeftContainer = createItemContainer(eItemContainerType.Weapon);
            weaponLeftContainer.Location = sprite.Location + new Size(20, 47);
            
            weaponRightContainer = createItemContainer(eItemContainerType.Weapon);
            weaponRightContainer.Location = sprite.Location + new Size(253, 47);
            
            beltContainer = createItemContainer(eItemContainerType.Belt);
            beltContainer.Location = sprite.Location + new Size(136, 178);
            
            ringtLeftContainer = createItemContainer(eItemContainerType.Ring);
            ringtLeftContainer.Location = sprite.Location + new Size(95, 179);
            
            ringtRightContainer = createItemContainer(eItemContainerType.Ring);
            ringtRightContainer.Location = sprite.Location + new Size(209, 179);
            
            gloveContainer = createItemContainer(eItemContainerType.Glove);
            gloveContainer.Location = sprite.Location + new Size(20, 179);
            
            bootsContainer = createItemContainer(eItemContainerType.Boots);
            bootsContainer.Location = sprite.Location + new Size(251, 178);
        }

        public void Update()
        {
            closeButton.Update();
            goldButton.Update();

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

            closeButton.Render();
            goldButton.Render();

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

        private void OpenGoldDrop()
        {
            // todo;
        }
    }
}
