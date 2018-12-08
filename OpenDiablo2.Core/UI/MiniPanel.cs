using OpenDiablo2.Common;
using OpenDiablo2.Common.Enums;
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
        private readonly IGameState gameState;
        private readonly ISprite sprite;
        private readonly IReadOnlyList<IButton> buttons;
        private readonly IEnumerable<IPanel> panels;

        private bool isPanelVisible;

        public event OnPanelToggledEvent OnPanelToggled;

        public MiniPanel(IRenderWindow renderWindow, 
            IGameState gameState,
            IMouseInfoProvider mouseInfoProvider,
            IEnumerable<IPanel> panels, 
            Func<eButtonType, IButton> createButton)
        {
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;
            this.gameState = gameState;
            this.panels = panels;

            sprite = renderWindow.LoadSprite(ResourcePaths.MinipanelSmall, Palettes.Units);

            buttons = panelButtons.Select((x, i) =>
            {
                var newBtn = createButton(x);
                var panel = panels.SingleOrDefault(o => o.PanelType == x);

                if (panel != null)
                {
                    newBtn.OnActivate = () => OnPanelToggled?.Invoke(panel);
                    panel.OnPanelClosed += Panel_OnPanelClosed;
                }
                return newBtn;
            }).ToList().AsReadOnly();

            UpdatePanelLocation();
        }

        public void OnMenuToggle(bool isToggled) => isPanelVisible = isToggled;

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
            sprite.Location = new Point((800 - sprite.LocalFrameSize.Width + (int)(gameState.CameraOffset * 1.3f)) / 2, 
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
