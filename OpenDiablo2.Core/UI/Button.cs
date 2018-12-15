using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System.Drawing;

namespace OpenDiablo2.Core.UI
{
    public sealed class Button : IButton
    {
        private readonly IMouseInfoProvider mouseInfoProvider;
        private readonly IRenderWindow renderWindow;
        private readonly ISoundProvider musicProvider;
        private readonly ButtonLayout buttonLayout;
        private readonly byte[] sfxButtonClick;

        public OnActivateDelegate OnActivate { get; set; }
        public OnToggleDelegate OnToggle { get; set; }

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

        private readonly int buttonWidth;
        private readonly int buttonHeight;
        private readonly ISprite sprite;
        private readonly IFont font;
        private readonly ILabel label;
        private bool pressed = false;
        private bool active = false; // When true, button is actively being focus pressed
        private bool activeLock = false; // When true, we have locked the mouse from everything else
        public bool Toggled { get; private set; } = false;

        private Point labelOffset = new Point();

        private bool enabled = true;
        public bool Enabled
        {
            get => enabled;
            set
            {
                if (value == enabled)
                    return;
                enabled = value;

                if(buttonLayout.IsDarkenedWhenDisabled)
                    sprite.Darken = !enabled;
            }
        }
        
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
            IMouseInfoProvider mouseInfoProvider,
            ISoundProvider soundProvider,
            IMPQProvider mpqProvider
            )
        {
            this.buttonLayout = buttonLayout;
            this.renderWindow = renderWindow;
            this.mouseInfoProvider = mouseInfoProvider;
            this.musicProvider = soundProvider;

            font = renderWindow.LoadFont(buttonLayout.FontPath, Palettes.Units);
            label = renderWindow.CreateLabel(font);

            sprite = renderWindow.LoadSprite(buttonLayout.ResourceName, buttonLayout.PaletteName, true);

            // TODO: Less stupid way of doing this would be super nice
            buttonWidth = 0;
            buttonHeight = 0;
            for (int i = 0; i < buttonLayout.XSegments; i++)
            {
                sprite.Frame = i;
                buttonWidth += sprite.LocalFrameSize.Width;
            }
            for(int i = 0; i < buttonLayout.YSegments; i++)
            {
                sprite.Frame = i * buttonLayout.YSegments;
                buttonHeight += sprite.LocalFrameSize.Height;
            }

            label.MaxWidth = buttonWidth - 8;
            label.Alignment = Common.Enums.eTextAlign.Centered;

            sfxButtonClick = mpqProvider.GetBytes(ResourcePaths.SFXButtonClick);
        }

        public bool Toggle()
        {
            Toggled = !Toggled;

            OnToggle?.Invoke(Toggled);

            return Toggled;
        }

        public bool Toggle(bool isToggled)
        {
            if(Toggled != isToggled)
            {
                OnToggle?.Invoke(isToggled);

                Toggled = isToggled;
            }

            return isToggled;
        }

        public void Update()
        {
            if (!enabled)
            {
                // Prevent sticky locks
                if (activeLock && mouseInfoProvider.ReserveMouse)
                {
                    activeLock = false;
                    mouseInfoProvider.ReserveMouse = false;
                }

                active = false;
                return;
            }

            int clickWidth = buttonLayout.ClickableRect.Width > 0 ? buttonLayout.ClickableRect.Width : buttonWidth;
            int clickHeight = buttonLayout.ClickableRect.Height > 0 ? buttonLayout.ClickableRect.Height : buttonHeight;

            var hovered = mouseInfoProvider.MouseX >= location.X + buttonLayout.ClickableRect.X 
                && mouseInfoProvider.MouseY >= location.Y + buttonLayout.ClickableRect.Y
                && mouseInfoProvider.MouseX < location.X + clickWidth + buttonLayout.ClickableRect.X
                && mouseInfoProvider.MouseY < location.Y + clickHeight + buttonLayout.ClickableRect.Y;

            if (!activeLock && hovered && mouseInfoProvider.LeftMouseDown && !mouseInfoProvider.ReserveMouse)
            {
                // The button is being pressed down
                mouseInfoProvider.ReserveMouse = true;
                active = true;
                musicProvider.PlaySfx(sfxButtonClick);
            }
            else if (active && !mouseInfoProvider.LeftMouseDown)
            {
                mouseInfoProvider.ReserveMouse = false;
                active = false;

                if (hovered)
                {
                    OnActivate?.Invoke();

                    if (buttonLayout.Toggleable)
                    {
                        Toggle();
                    }
                }
                    
            }
            else if (!active && mouseInfoProvider.LeftMouseDown)
            {
                activeLock = true;
            }
            else if (activeLock && !mouseInfoProvider.LeftMouseDown)
            {
                activeLock = false;
            }

            pressed = hovered && mouseInfoProvider.LeftMouseDown && active;
        }

        public void Render()
        {
            var frame = buttonLayout.BaseFrame;

            if (buttonLayout.AllowFrameChange)
            {
                if(!Enabled && buttonLayout.DisabledFrame >= 0)
                {
                    frame = buttonLayout.DisabledFrame;
                }
                else if (Toggled && pressed)
                {
                    frame = buttonLayout.BaseFrame + 3;
                }
                else if (pressed)
                {
                    frame = buttonLayout.BaseFrame + 1;
                }
                else if (Toggled)
                {
                    frame = buttonLayout.BaseFrame + 2;
                }
            }

            renderWindow.Draw(sprite, buttonLayout.XSegments, buttonLayout.YSegments, frame);
            var offset = pressed ? -2 : 0;

            label.Location = new Point(location.X + offset + labelOffset.X, location.Y - offset + labelOffset.Y);
            renderWindow.Draw(label);
        }

        private void UpdateText()
        {
            label.Text = text;
            label.TextColor = Color.FromArgb(75, 75, 75);

            var offsetX = (buttonWidth / 2) - (label.TextArea.Width / 2);
            var offsetY = (buttonHeight / 2) - (label.TextArea.Height / 2);
            labelOffset = new Point(offsetX, offsetY - 5);
        }

        public void Dispose()
        {
            sprite.Dispose();
            font.Dispose();
            label.Dispose();
        }
    }        
}
