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
        public Point Position { get; set; } = new Point();

        ISprite sprite;
        IFont font;
        ILabel label;

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
            font = renderWindow.LoadFont(ResourcePaths.Font24, Palettes.Static);
        }

        public void Update()
        {

        }

        public void Draw()
        {

        }

        private void UpdateText()
        {

        }
    }
}
