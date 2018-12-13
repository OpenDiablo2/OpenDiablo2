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
using OpenDiablo2.Common.Interfaces.UI;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Core.UI
{
    public class SkillsPanel : ISkillsPanel
    {
        private readonly IRenderWindow renderWindow;
        private readonly ISprite panelSprite;

        private readonly IButton[] treeButtons;

        public event OnPanelClosedEvent OnPanelClosed;

        // Test fields
        private readonly eHero hero = eHero.Barbarian;

        public SkillsPanel(
            IRenderWindow renderWindow,
            Func<eHero, int, IButton> createTreeButton,
            Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;

            panelSprite = renderWindow.LoadSprite(ResourcePaths.GetHeroSkillPanel(hero), Palettes.Act1, FrameType.GetOffset(), true);

            treeButtons = Enumerable.Range(0, 3).Select(o =>
            {
                var btn = createTreeButton(hero, o);
                btn.Location = FrameType.GetOffset();
                btn.OnActivate = () => { ActivePanelIndex = o; };
                return btn;
            }).ToArray();
        }

        public ePanelType PanelType => ePanelType.Skill;
        public ePanelFrameType FrameType => ePanelFrameType.Right;

        public int ActivePanelIndex { get; private set; }

        public void Update()
        {
            foreach (var button in treeButtons)
                button.Update();
        }

        public void Render()
        {
            renderWindow.Draw(panelSprite, 2, 2, 0);

            treeButtons[ActivePanelIndex].Render();
        }

        public void Dispose()
        {
            panelSprite.Dispose();
        }
    }
}
