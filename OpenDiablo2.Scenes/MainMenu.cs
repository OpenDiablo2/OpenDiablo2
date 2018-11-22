using OpenDiablo2.Common;
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

            backgroundSprite = renderWindow.LoadSprite(ResourcePaths.GameSelectScreen, Palettes.Sky);
            diabloLogoLeft = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoFireLeft, Palettes.Units, new Point(400, 120));
            diabloLogoRight = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoFireRight, Palettes.Units, new Point(400, 120));
            diabloLogoLeftBlack = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoBlackLeft, Palettes.Units, new Point(400, 120));
            diabloLogoRightBlack = renderWindow.LoadSprite(ResourcePaths.Diablo2LogoBlackRight, Palettes.Units, new Point(400, 120));

            mouseSprite = renderWindow.LoadSprite(ResourcePaths.CursorDefault, Palettes.Units);
            wideButton = renderWindow.LoadSprite("data\\global\\ui\\FrontEnd\\WideButtonBlank.dc6", "ACT1");


            var loadingSprite = renderWindow.LoadSprite(ResourcePaths.LoadingScreen, Palettes.Loading, new Point(300, 400));

            renderWindow.Clear();
            renderWindow.Draw(loadingSprite);
            renderWindow.Sync();

            musicProvider.LoadSong(mpqProvider.GetStream("data\\global\\music\\introedit.wav"));

            // TODO: Fake loading for now, this should be in its own scene as we start loading real stuff
            var r = new Random();
            for (int i = 1; i < 10; i++)
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

            renderWindow.Draw(backgroundSprite, 4, 3, 0);

            renderWindow.Draw(diabloLogoLeftBlack, (int)((float)diabloLogoLeftBlack.TotalFrames * logoFrame));
            renderWindow.Draw(diabloLogoRightBlack, (int)((float)diabloLogoRightBlack.TotalFrames * logoFrame));

            renderWindow.Draw(diabloLogoLeft, (int)((float)diabloLogoLeft.TotalFrames * logoFrame));
            renderWindow.Draw(diabloLogoRight, (int)((float)diabloLogoRight.TotalFrames * logoFrame));

            wideButton.Location = new Point(264, 290);
            renderWindow.Draw(wideButton, 2, 1, 0);

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
