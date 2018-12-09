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

using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Extensions;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Drawing;

namespace OpenDiablo2.Core.UI
{
    public sealed class CharacterPanel : ICharacterPanel
    {
        private readonly IRenderWindow renderWindow;
        private readonly ISprite panelSprite;

        private readonly IButton closeButton;

        public event OnPanelClosedEvent OnPanelClosed;

        public CharacterPanel(
            IRenderWindow renderWindow,
            Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;

            panelSprite = renderWindow.LoadSprite(ResourcePaths.InventoryCharacterPanel, Palettes.Act1, FrameType.GetOffset(), true);

            closeButton = createButton(eButtonType.Close);
            closeButton.Location = panelSprite.Location + new Size(128, 388);
            closeButton.OnActivate = () => OnPanelClosed?.Invoke(this);
        }

        public eButtonType PanelType => eButtonType.MinipanelCharacter;
        public ePanelFrameType FrameType => ePanelFrameType.Left;

        public void Update()
        {
            closeButton.Update();
        }

        public void Render()
        {
            renderWindow.Draw(panelSprite, 2, 2, 0);

            closeButton.Render();
        }

        public void Dispose()
        {
            panelSprite.Dispose();
        }
    }
}
