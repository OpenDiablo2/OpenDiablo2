using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Core.UI
{


    public sealed class Button : IDisposable
    {
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IRenderWindow renderWindow;
        private readonly ButtonLayout buttonLayout;

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

        private int buttonWidth, buttonHeight;
        private ISprite sprite;
        private IFont font;
        private ILabel label;
        private bool pressed = false;
        private bool active = false; // When true, button is actively being focus pressed
        private bool activeLock = false; // When true, something else is being pressed so ignore everything
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


        public Button(
            ButtonLayout buttonLayout,
            IRenderWindow renderWindow, 
            IMouseInfoProvider mouseInfoProvider
            )
        {
            this.buttonLayout = buttonLayout;
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;

            font = renderWindow.LoadFont(ResourcePaths.FontExocet10, Palettes.Units);
            label = renderWindow.CreateLabel(font);


            sprite = renderWindow.LoadSprite(buttonLayout.ResourceName, buttonLayout.PaletteName);

            // TODO: Less stupid way of doing this would be nice
            buttonWidth = 0;
            buttonHeight = 0;
            for (int i = 0; i < buttonLayout.XSegments; i++)
            {
                sprite.Frame = i;
                buttonWidth += sprite.LocalFrameSize.Width;
                buttonHeight = Math.Max(buttonHeight, sprite.LocalFrameSize.Height);
            }
        }

        public void Update()
        {
            var hovered = (mouseInfoProvider.MouseX >= location.X && mouseInfoProvider.MouseX < (location.X + buttonWidth))
                && (mouseInfoProvider.MouseY >= location.Y && mouseInfoProvider.MouseY < (location.Y + buttonHeight));


            if (!activeLock && hovered && mouseInfoProvider.LeftMouseDown && !mouseInfoProvider.ReserveMouse)
            {
                // The button is being pressed down
                mouseInfoProvider.ReserveMouse = true;
                active = true;

            }
            else if (active && !mouseInfoProvider.LeftMouseDown)
            {
                mouseInfoProvider.ReserveMouse = false;
                active = false;

                if (hovered)
                    OnActivate?.Invoke();
            }
            else if (!active && mouseInfoProvider.LeftMouseDown)
            {
                activeLock = true;
            }
            else if (activeLock && !mouseInfoProvider.LeftMouseDown)
            {
                activeLock = false;
            }

            pressed = (hovered && mouseInfoProvider.LeftMouseDown && active);
        }

        public void Render()
        {
            renderWindow.Draw(sprite, buttonLayout.XSegments, 1, pressed ? 1 : 0);
            var offset = pressed ? -2 : 0;

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

        public void Dispose()
        {
            sprite.Dispose();
            font.Dispose();
            label.Dispose();
        }
    }        
}
