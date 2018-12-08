using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using OpenDiablo2.Common;
using OpenDiablo2.Common.Interfaces;

namespace OpenDiablo2.Core.UI
{
    public sealed class TextBox : ITextBox
    {
        private readonly IRenderWindow renderWindow;
        private readonly ISprite sprite;
        private readonly IFont font;
        private readonly ILabel label, linebar;
        private float frameTime = 0f;

        private Point location = new Point();
        public Point Location
        {
            get => location;
            set
            {
                if (location == value)
                    return;
                location = value;
                label.Location = new Point(value.X + 6, value.Y + 3);
                linebar.Location = new Point(value.X + 6 + label.TextArea.Width, value.Y + 3);
                sprite.Location = new Point(value.X, value.Y + sprite.LocalFrameSize.Height);
            }
        }

        private string text = "";
        public string Text
        {
            get => text;
            set
            {
                if (text == value)
                    return;
                text = value;

                // Max width is 130
                var newSize = font.CalculateSize(value);
                if (newSize.Width < 130)
                {
                    label.Text = value;
                    linebar.Location = new Point(location.X + 6 + newSize.Width, location.Y + 3);
                    return;
                }

                var newStr = value.Substring(1);
                while(true)
                {
                    newSize = font.CalculateSize(newStr);
                    if (newSize.Width >= 130)
                    {
                        newStr = newStr.Substring(1);
                        continue;
                    }

                    label.Text = newStr;
                    linebar.Location = new Point(location.X + 6 + newSize.Width, location.Y + 3);
                    break;
                }
                
                
                
            }
        }

        public TextBox(IRenderWindow renderWindow)
        {
            this.renderWindow = renderWindow;

            sprite = renderWindow.LoadSprite(ResourcePaths.TextBox2, Palettes.Units);
            font = renderWindow.LoadFont(ResourcePaths.FontFormal11, Palettes.Units);
            label = renderWindow.CreateLabel(font);
            linebar = renderWindow.CreateLabel(font);
            linebar.Text = "_";
        }


        public void Update(long ms)
        {
            frameTime += ms / 500f;
            while (frameTime >= 1f)
                frameTime -= 1f;
        }

        public void Render()
        {
            renderWindow.Draw(sprite);
            renderWindow.Draw(label);
            if (frameTime < 0.5)
                renderWindow.Draw(linebar);
        }

        public void Dispose()
        {

        }
    }
}
