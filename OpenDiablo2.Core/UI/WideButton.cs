using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Core.UI
{
    public sealed class WideButton
    {
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IRenderWindow renderWindow;

        public delegate void OnActivateDelegate();
        public OnActivateDelegate OnActivate { get; set; }

        private Point location = new Point();
        public Point Location
        {
            get => location;
            set
            {
                location = value;
                sprite.Location = value;
            }
        }

        private readonly int buttonWidth, buttonHeight;
        private ISprite sprite;
        private IFont font;
        private ILabel label;
        private bool pressed = false;
        private Point labelOffset = new Point();

        private string text;
        public string Text
        {
            get => text;
            set
            {
                text = value;
                UpdateText();
            }
        }

        public WideButton(IRenderWindow renderWindow, IMouseInfoProvider mouseInfoProvider)
        {
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;

            sprite = renderWindow.LoadSprite(ResourcePaths.WideButtonBlank, Palettes.Units);
            font = renderWindow.LoadFont(ResourcePaths.FontExocet10, Palettes.Units);
            label = renderWindow.CreateLabel(font);

            // TODO: Less stupid way of doing this would be nice
            sprite.Frame = 0;
            buttonWidth = sprite.LocalFrameSize.Width;
            buttonHeight = sprite.LocalFrameSize.Height;
            sprite.Frame = 1;
            buttonWidth += sprite.LocalFrameSize.Width;

        }

        public void Update()
        {
            var hovered = (mouseInfoProvider.MouseX >= location.X && mouseInfoProvider.MouseX < (location.X + buttonWidth))
                && (mouseInfoProvider.MouseY >= location.Y && mouseInfoProvider.MouseY < (location.Y + buttonHeight));

            pressed = hovered;
        }

        public void Render()
        {
            renderWindow.Draw(sprite, 2, 1, pressed ? 1 : 0);
            var offset = pressed ? -3 : 0;

            label.Location = new Point(location.X + offset + labelOffset.X, location.Y - offset + labelOffset.Y);
            renderWindow.Draw(label);
        }

        private void UpdateText()
        {
            label.Text = text;
            label.TextColor = Color.FromArgb(75, 75, 75);

            var offsetX = (buttonWidth / 2) - (label.TextArea.Width / 2);
            labelOffset = new Point(offsetX, -5);
        }
    }
}
