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

        private ISprite sprite;
        private IFont font;
        private ILabel label;
        private bool pressed;
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

            sprite = renderWindow.LoadSprite(ResourcePaths.WideButtonBlank, Palettes.Act1);
            font = renderWindow.LoadFont(ResourcePaths.FontExocet10, Palettes.Menu4);
            label = renderWindow.CreateLabel(font);
        }

        public void Update()
        {

        }

        public void Render()
        {
            renderWindow.Draw(sprite, 2, 1, pressed ? 1 : 0);
            var offset = pressed ? -5 : 0;

            label.Location = new Point(location.X + offset + labelOffset.X, location.Y + offset + labelOffset.Y);
            renderWindow.Draw(label);
        }

        private void UpdateText()
        {
            label.Text = text;
            label.TextColor = Color.FromArgb(128, 128, 128);

            // TODO: Less stupid way of doing this would be nice
            sprite.Frame = 0;
            int btnWidth = sprite.LocalFrameSize.Width;
            sprite.Frame = 1;
            btnWidth += sprite.LocalFrameSize.Width;

            var offsetX = (btnWidth / 2) - (label.TextArea.Width / 2);
            labelOffset = new Point(offsetX, -5);
        }
    }
}
