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
using System.Collections.Generic;
using System.Drawing;
using System.Linq;

namespace OpenDiablo2.Core.UI
{
    public sealed class MiniPanel : IMiniPanel
    {
        private static readonly IEnumerable<eButtonType> panelButtons = new[] { eButtonType.MinipanelCharacter, eButtonType.MinipanelInventory,
            eButtonType.MinipanelSkill, eButtonType.MinipanelAutomap, eButtonType.MinipanelMessage, eButtonType.MinipanelQuest, eButtonType.MinipanelMenu };

        private readonly IRenderWindow renderWindow;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly ISprite sprite;
        private readonly IReadOnlyList<IButton> buttons;
        private readonly IEnumerable<IPanel> panels;
        private readonly Func<IGameHUD> _getGameHud;

        private bool isPanelVisible;

        public event OnPanelToggledEvent OnPanelToggled;

        public MiniPanel(IRenderWindow renderWindow,
            Func<IGameHUD> getGameHud,
            IMouseInfoProvider mouseInfoProvider,
            IEnumerable<IPanel> panels, 
            Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;
            this._getGameHud = getGameHud;
            this.panels = panels;

            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units, true);

            buttons = panelButtons.Select((x, i) =>
            {
                var newBtn = createButton(x);
                var panel = panels.SingleOrDefault(o => o.PanelType == x.GetPanelType());

                if (panel != null)
                {
                    newBtn.OnActivate = () => OnPanelToggled?.Invoke(panel);
                    panel.OnPanelClosed += Panel_OnPanelClosed;
                }
                return newBtn;
            }).ToList().AsReadOnly();
        }

        public void OnMenuToggle(bool isToggled) => isPanelVisible = isToggled;

        public IPanel GetPanel(ePanelType panelType)
        {
            return panels.SingleOrDefault(o => o.PanelType == panelType);
        }

        public bool IsMouseOver()
        {
            int xDiff = mouseInfoProvider.MouseX - sprite.Location.X;
            int yDiff = mouseInfoProvider.MouseY - sprite.Location.Y + sprite.LocalFrameSize.Height;

            return isPanelVisible
                && xDiff >= 0 && xDiff <= sprite.LocalFrameSize.Width
                && yDiff >= 0 && yDiff <= sprite.LocalFrameSize.Height;
        }

        public void Render()
        {
            if (!isPanelVisible)
                return;

            renderWindow.Draw(sprite);

            foreach (var button in buttons)
                button.Render();
        }

        public void Update()
        {
            if (!isPanelVisible)
                return;

            foreach (var button in buttons)
                button.Update();
        }

        public void Dispose()
        {
            foreach (var button in buttons)
                button.Dispose();

            sprite.Dispose();
        }

        public void UpdatePanelLocation()
        {
            var cameraOffset = (_getGameHud().IsRightPanelVisible ? -200 : 0) + (_getGameHud().IsLeftPanelVisible ? 200 : 0);
            sprite.Location = new Point((800 - sprite.LocalFrameSize.Width + (int)(cameraOffset * 1.3f)) / 2, 
                526 + sprite.LocalFrameSize.Height);

            for (int i = 0; i < buttons.Count; i++)
                buttons[i].Location = new Point(3 + 21 * i + sprite.Location.X, 3 + sprite.Location.Y - sprite.LocalFrameSize.Height);
        }

        private void Panel_OnPanelClosed(IPanel panel)
        {
            OnPanelToggled?.Invoke(panel);
        }
    }
}
