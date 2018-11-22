using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OpenDiablo2.Scenes
{
    [Scene("Main Menu")]
    public class MainMenu : IScene
    {
        static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private readonly IRenderWindow renderWindow;
        private readonly IPaletteProvider paletteProvider;
        private readonly IMPQProvider mpqProvider;
        private readonly IMouseInfoProvider mouseInfoProvider;

        private float logoFrame;
        private ISprite backgroundSprite, diabloLogoLeft, diabloLogoRight, diabloLogoLeftBlack, diabloLogoRightBlack, mouseSprite;

        public MainMenu(
            IRenderWindow renderWindow,
            IPaletteProvider paletteProvider,
            IMPQProvider mpqProvider,
            IMouseInfoProvider mouseInfoProvider
            )
        {
            this.renderWindow = renderWindow;
            this.paletteProvider = paletteProvider;
            this.mpqProvider = mpqProvider;
            this.mouseInfoProvider = mouseInfoProvider;

            //var texture = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\Logo\\logo.DC6")));
            backgroundSprite = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\FrontEnd\\gameselectscreenEXP.dc6")));
            backgroundSprite.CurrentPalette = paletteProvider.PaletteTable["Sky"];

            diabloLogoLeft = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\FrontEnd\\D2logoFireLeft.DC6")));
            diabloLogoLeft.CurrentPalette = paletteProvider.PaletteTable["Units"];

            diabloLogoRight = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\FrontEnd\\D2logoFireRight.DC6")));
            diabloLogoRight.CurrentPalette = paletteProvider.PaletteTable["Units"];

            diabloLogoLeftBlack = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\FrontEnd\\D2logoBlackLeft.DC6")));
            diabloLogoLeftBlack.CurrentPalette = paletteProvider.PaletteTable["Units"];

            diabloLogoRightBlack = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\FrontEnd\\D2logoBlackRight.DC6")));
            diabloLogoRightBlack.CurrentPalette = paletteProvider.PaletteTable["Units"];

            mouseSprite = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\CURSOR\\ohand.DC6")));
            mouseSprite.CurrentPalette = paletteProvider.PaletteTable["STATIC"];

            logoFrame = 0f;

            diabloLogoLeft.Location = new Point(400, 120);
            diabloLogoRight.Location = new Point(400, 120);
            diabloLogoLeftBlack.Location = new Point(400, 120);
            diabloLogoRightBlack.Location = new Point(400, 120);

        }

        public void Render()
        {
            renderWindow.Clear();

            for (int y = 0; y < 3; y++)
                for (int x = 0; x < 4; x++)
                {
                    backgroundSprite.Frame = x + (y * 4);
                    backgroundSprite.Location = new Point(x * backgroundSprite.FrameSize.Width, (y + 1) * backgroundSprite.FrameSize.Height);
                    renderWindow.Draw(backgroundSprite);
                }

            diabloLogoLeftBlack.Frame = (int)((float)diabloLogoLeftBlack.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoLeftBlack);
            diabloLogoRightBlack.Frame = (int)((float)diabloLogoRightBlack.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoRightBlack);

            diabloLogoLeft.Frame = (int)((float)diabloLogoLeft.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoLeft);
            diabloLogoRight.Frame = (int)((float)diabloLogoRight.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoRight);

            mouseSprite.Location = new Point(mouseInfoProvider.MouseX, mouseInfoProvider.MouseY + mouseSprite.FrameSize.Height - 1);
            renderWindow.Draw(mouseSprite);

            renderWindow.Sync();
        }

        public void Update(long ms)
        {
            float seconds = ((float)ms / 1000f);
            logoFrame += seconds;
            while (logoFrame >= 1f)
                logoFrame -= 1f;

        }

        public void Dispose()
        {

        }
    }
}
