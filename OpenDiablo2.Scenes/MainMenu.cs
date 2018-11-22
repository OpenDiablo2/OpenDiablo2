using OpenDiablo2.Common.Attributes;
using OpenDiablo2.Common.Interfaces;
using OpenDiablo2.Common.Models;
using System;
using System.Collections.Generic;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading;
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
        private readonly IMusicProvider musicProvider;

        private float logoFrame;
        private ISprite backgroundSprite, diabloLogoLeft, diabloLogoRight, diabloLogoLeftBlack, diabloLogoRightBlack, mouseSprite, wideButton;

        public MainMenu(
            IRenderWindow renderWindow,
            IPaletteProvider paletteProvider,
            IMPQProvider mpqProvider,
            IMouseInfoProvider mouseInfoProvider,
            IMusicProvider musicProvider
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

            wideButton = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\FrontEnd\\WideButtonBlank.dc6")));
            wideButton.CurrentPalette = paletteProvider.PaletteTable["ACT1"];

            logoFrame = 0f;

            diabloLogoLeft.Location = new Point(400, 120);
            diabloLogoRight.Location = new Point(400, 120);
            diabloLogoLeftBlack.Location = new Point(400, 120);
            diabloLogoRightBlack.Location = new Point(400, 120);


            var loadingSprite = renderWindow.LoadSprite(ImageSet.LoadFromStream(mpqProvider.GetStream("data\\global\\ui\\Loading\\loadingscreen.dc6")));
            loadingSprite.CurrentPalette = paletteProvider.PaletteTable["loading"];
            loadingSprite.Location = new Point(300, 400);
            renderWindow.Clear();
            renderWindow.Draw(loadingSprite);
            renderWindow.Sync();

            musicProvider.LoadSong(mpqProvider.GetStream("data\\global\\music\\introedit.wav"));

            // TODO: Fake loading for now, this should be in its own scene as we start loading real stuff
            var r = new Random();
            for(int i = 1; i < 10; i++)
            {
                renderWindow.Clear();
                loadingSprite.Frame = i;
                renderWindow.Draw(loadingSprite);
                renderWindow.Sync();
                Thread.Sleep(r.Next(150));

            }

            musicProvider.PlaySong();
        }

        public void Render()
        {
            renderWindow.Clear();

            for (int y = 0; y < 3; y++)
            {
                for (int x = 0; x < 4; x++)
                {
                    backgroundSprite.Frame = x + (y * 4);
                    backgroundSprite.Location = new Point(x * 256, ((y+1) * 256) - (backgroundSprite.FrameSize.Height - backgroundSprite.LocalFrameSize.Height));
                    renderWindow.Draw(backgroundSprite);
                }

            }

            diabloLogoLeftBlack.Frame = (int)((float)diabloLogoLeftBlack.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoLeftBlack);
            diabloLogoRightBlack.Frame = (int)((float)diabloLogoRightBlack.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoRightBlack);

            diabloLogoLeft.Frame = (int)((float)diabloLogoLeft.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoLeft);
            diabloLogoRight.Frame = (int)((float)diabloLogoRight.TotalFrames * logoFrame);
            renderWindow.Draw(diabloLogoRight);


            wideButton.Location = new Point(260, 320);
            wideButton.Frame = 0;
            renderWindow.Draw(wideButton);
            wideButton.Frame = 1;
            wideButton.Location = new Point(260 + 256, 320);
            renderWindow.Draw(wideButton);

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
