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
using OpenDiablo2.Common.Interfaces;
using System;
using System.Diagnostics;
using System.Drawing;

namespace OpenDiablo2.Core.UI
{
    public class GameHUD : IGameHUD
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IMiniPanel minipanel;

        private readonly IButton runButton, menuButton, addStatButton, addSkillButton;
        private readonly ISprite panelSprite, healthManaSprite, gameGlobeOverlapSprite;
        private readonly IPanelFrame leftPanelFrame, rightPanelFrame;

        // Test fields
        private bool addNewAttribute = true;
        private bool addNewSkill = true;

        public GameHUD(
            IRenderWindow renderWindow,
            IMouseInfoProvider mouseInfoProvider,
            Func<IMiniPanel> createMiniPanel,
            Func<eButtonType, IButton> createButton,
            Func<ePanelFrameType, IPanelFrame> createPanelFrame)
        {
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;
            minipanel = createMiniPanel();
            minipanel.OnPanelToggled += TogglePanel;

            leftPanelFrame = createPanelFrame(ePanelFrameType.Left);
            rightPanelFrame = createPanelFrame(ePanelFrameType.Right);

            runButton = createButton(eButtonType.Run);
            runButton.Location = new Point(256, 570);
            runButton.OnToggle = OnRunToggle;

            menuButton = createButton(eButtonType.Menu);
            menuButton.Location = new Point(393, 561);
            menuButton.OnToggle = minipanel.OnMenuToggle;
            menuButton.Toggle();

            addStatButton = createButton(eButtonType.Skill);
            addStatButton.Location = new Point(207, 562);
            addStatButton.Enabled = addNewAttribute;
            addStatButton.OnActivate = () => TogglePanel(ePanelType.Character);

            addSkillButton = createButton(eButtonType.Skill);
            addSkillButton.Location = new Point(564, 562);
            addSkillButton.Enabled = addNewSkill;
            addSkillButton.OnActivate = () => TogglePanel(ePanelType.Skill);

            panelSprite = renderWindow.LoadSprite(ResourcePaths.GamePanels, Palettes.Act1, true);
            healthManaSprite = renderWindow.LoadSprite(ResourcePaths.HealthMana, Palettes.Act1, true);
            gameGlobeOverlapSprite = renderWindow.LoadSprite(ResourcePaths.GameGlobeOverlap, Palettes.Act1, true);
        }
        
        public IPanel LeftPanel { get; private set; }
        public IPanel RightPanel { get; private set; }
        public bool ArePanelsBounded { get; private set; } = false;

        public bool IsLeftPanelVisible => LeftPanel != null;
        public bool IsRightPanelVisible => RightPanel != null;
        public bool IsRunningEnabled => runButton.Toggled;

        private bool needsCameraUpdate = true;

        public void TogglePanel(ePanelType panelType)
        {
            TogglePanel(minipanel.GetPanel(panelType));
        }

        public void TogglePanel(IPanel panel)
        {
            switch (panel.FrameType)
            {
                case ePanelFrameType.Left:
                    LeftPanel = LeftPanel == panel ? null : panel;
                    UpdateCameraOffset();
                    if (ArePanelsBounded)
                        RightPanel = null;
                    ArePanelsBounded = false;
                    break;

                case ePanelFrameType.Right:
                    RightPanel = RightPanel == panel ? null : panel;
                    UpdateCameraOffset();
                    if (ArePanelsBounded)
                        LeftPanel = null;
                    ArePanelsBounded = false;
                    break;

                case ePanelFrameType.Center:
                    // todo; write logic for "center" panels
                    break;

                default:
                    Debug.Fail("Unknown frame type");
                    break;
            }
        }

        // used when panels are bounded with each other (shops/chest)
        public void OpenPanels(IPanel leftPanel, IPanel rightPanel)
        {
            if (leftPanel.FrameType != ePanelFrameType.Left || rightPanel.FrameType != ePanelFrameType.Right)
                throw new ArgumentException("wrong panel position.");

            LeftPanel = leftPanel;
            RightPanel = rightPanel;
            UpdateCameraOffset();
            ArePanelsBounded = true;
        }

        public void ClosePanels()
        {
            LeftPanel = null;
            RightPanel = null;
            UpdateCameraOffset();
            ArePanelsBounded = false;
        }

        public bool IsMouseOver()
        {
            return mouseInfoProvider.MouseY >= 550
                || minipanel.IsMouseOver()
                || IsRightPanelVisible && mouseInfoProvider.MouseX >= 400
                || IsLeftPanelVisible && mouseInfoProvider.MouseX < 400;
        }

        public void Render()
        {
            if (IsLeftPanelVisible)
            {
                LeftPanel.Render();
                leftPanelFrame.Render();
            }

            if (IsRightPanelVisible)
            {
                RightPanel.Render();
                rightPanelFrame.Render();
            }
            
            // Render the background bottom bar
            renderWindow.Draw(panelSprite, 0, new Point(0, 600));
            renderWindow.Draw(panelSprite, 1, new Point(166, 600));
            renderWindow.Draw(panelSprite, 2, new Point(294, 600));
            renderWindow.Draw(panelSprite, 3, new Point(422, 600));
            renderWindow.Draw(panelSprite, 4, new Point(550, 600));
            renderWindow.Draw(panelSprite, 5, new Point(685, 600));

            // Render the health bar
            renderWindow.Draw(healthManaSprite, 0, new Point(30, 587));
            renderWindow.Draw(gameGlobeOverlapSprite, 0, new Point(28, 595));

            // Render the mana bar
            renderWindow.Draw(healthManaSprite, 1, new Point(692, 588));
            renderWindow.Draw(gameGlobeOverlapSprite, 1, new Point(693, 591));
            
            if (!IsLeftPanelVisible || !IsRightPanelVisible)
                minipanel.Render();

            runButton.Render();
            menuButton.Render();
            addStatButton.Render();
            addSkillButton.Render();
        }

        public void Update()
        {
            if (needsCameraUpdate)
            {
                needsCameraUpdate = false;
                UpdateCameraOffset();
            }

            runButton.Update();
            menuButton.Update();
            addStatButton.Update();
            addSkillButton.Update();

            if (IsLeftPanelVisible)
            {
                LeftPanel.Update();
                leftPanelFrame.Update();
            }

            if (IsRightPanelVisible)
            {
                RightPanel.Update();
                rightPanelFrame.Update();
            }

            if(!IsLeftPanelVisible || !IsRightPanelVisible)
                minipanel.Update();
        }

        private void UpdateCameraOffset()
            => minipanel.UpdatePanelLocation();

        private void OnRunToggle(bool isToggled)
            => log.Debug("Run Toggle: " + isToggled);
    }
}
