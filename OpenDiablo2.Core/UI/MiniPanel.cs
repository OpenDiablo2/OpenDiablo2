using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Drawing;
using System.Linq;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    public sealed class MiniPanel : IMiniPanel
    {
        private static readonly IEnumerable<eButtonType> panelButtons = new[] { eButtonType.MinipanelCharacter, eButtonType.MinipanelInventory,
            eButtonType.MinipanelSkill, eButtonType.MinipanelAutomap, eButtonType.MinipanelMessage, eButtonType.MinipanelQuest, eButtonType.MinipanelMenu };

        private readonly IRenderWindow renderWindow;
        private readonly IGameState gameState;
        private readonly ISprite sprite;
        private readonly IReadOnlyList<IButton> buttons;
        private readonly IEnumerable<IPanel> panels;
        private readonly IPanelFrame leftPanelFrame;
        private readonly IPanelFrame rightPanelFrame;

        private bool isPanelVisible;

        public IPanel LeftPanel { get; private set; }
        public IPanel RightPanel { get; private set; }
        public bool IsLeftPanelVisible => LeftPanel != null;
        public bool IsRightPanelVisible => RightPanel != null;

        public MiniPanel(IRenderWindow renderWindow, IGameState gameState, 
            IEnumerable<IPanel> panels, Func<eButtonType, IButton> createButton,
            Func<ePanelFrameType, IPanelFrame> createPanelFrame)
        {
            this.renderWindow = renderWindow;
            this.gameState = gameState;
            this.panels = panels;
            leftPanelFrame = createPanelFrame(ePanelFrameType.Left);
            rightPanelFrame = createPanelFrame(ePanelFrameType.Right);

            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units);

            buttons = panelButtons.Select((x, i) =>
            {
                var newBtn = createButton(x);
                newBtn.OnActivate = () =>
                {
                    var panel = panels.SingleOrDefault(o => o.PanelType == x);
                    if (panel == null) return;
                    TogglePanel(panel);
                };
                return newBtn;
            }).ToList().AsReadOnly();

            UpdatePanelLocation();
            OnMenuToggle(true);
        }

        public void OnMenuToggle(bool isToggled) => isPanelVisible = isToggled;

        public void Update()
        {
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
            
            if (!isPanelVisible || (IsLeftPanelVisible && IsRightPanelVisible))
                return;

            foreach (var button in buttons)
                button.Update();
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

            if (!isPanelVisible || (IsLeftPanelVisible && IsRightPanelVisible))
                return;

            renderWindow.Draw(sprite);

            foreach (var button in buttons)
                button.Render();
        }

        public void Dispose()
        {
            foreach (var button in buttons)
                button.Dispose();

            sprite.Dispose();
        }

        private void TogglePanel(IPanel panel)
        {
            switch (panel.FrameType)
            {
                case ePanelFrameType.Left:
                    LeftPanel = LeftPanel == panel ? null : panel;
                    break;
                case ePanelFrameType.Right:
                    RightPanel = RightPanel == panel ? null : panel;
                    break;
                case ePanelFrameType.Center:
                    // todo; stack center panels
                    break;
                default:
                    Debug.Fail("Unknown frame type");
                    break;
            }

            UpdateCameraOffset();
        }

        private void UpdateCameraOffset()
        {
            gameState.CameraOffset = (IsRightPanelVisible ? -200 : 0) + (IsLeftPanelVisible ? 200 : 0);
            UpdatePanelLocation();
        }

        private void UpdatePanelLocation()
        {
            sprite.Location = new Point((800 - sprite.LocalFrameSize.Width + (int)(gameState.CameraOffset * 1.3f)) / 2, 
                526 + sprite.LocalFrameSize.Height);

            for (int i = 0; i < buttons.Count; i++)
                buttons[i].Location = new Point(3 + 21 * i + sprite.Location.X, 3 + sprite.Location.Y - sprite.LocalFrameSize.Height);
        }
    }
}
